package utilscore

import (
	"cnpc_backend/core/typescore"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/shopspring/decimal"
	"log"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

func ParseParamsGetRequest(queryParams url.Values, result interface{}) (*uint64, *uint64, map[string]string, *typescore.WEvent) {
	likeFields := map[string]string{}
	singleValueQueryParams := make(map[string]interface{})
	offset := PointerToUint64(0)
	limit := PointerToUint64(50)
	likeFieldsMode := queryParams.Get("like_fields_mode") == "true"

	paramTypes := make(map[string]string)
	// Получаем типы полей структуры result
	resultType := reflect.TypeOf(result).Elem()
	for i := 0; i < resultType.NumField(); i++ {
		field := resultType.Field(i)
		tag := field.Tag.Get("mapstructure")
		if tag != "" {
			fieldType := field.Type
			if fieldType.Kind() == reflect.Ptr {
				fieldType = fieldType.Elem()
			}
			paramTypes[tag] = fieldType.String()
		}
	}

	for key, values := range queryParams {
		if len(values) > 0 {
			value := values[0]
			switch key {
			case "offset":
				offset = parseUintParam(value)
			case "limit":
				limit = parseUintParam(value)
			default:
				singleValueQueryParams[key] = parseQueryParam(value, likeFieldsMode, &likeFields, key, paramTypes)
			}
		}
	}

	decoder, wEvent := setupDecoder(result)
	if wEvent != nil {
		return offset, limit, likeFields, wEvent
	}

	err := decoder.Decode(singleValueQueryParams)
	if err != nil {
		return offset, limit, likeFields, &typescore.WEvent{Err: err, Text: "system_error"}
	}

	return offset, limit, likeFields, nil
}

func parseUintParam(value string) *uint64 {
	val, err := strconv.ParseUint(value, 10, 64)
	if err == nil {
		return &val
	}
	return nil
}

func parseQueryParam(value string, likeFieldsMode bool, likeFields *map[string]string, key string, paramTypes map[string]string) interface{} {
	switch value {
	case "true":
		return true
	case "false":
		return false
	case "all":
		return nil
	default:
		if strings.Contains(value, ",") {
			return strings.Split(value, ",")
		}
		if likeFieldsMode {
			(*likeFields)[key] = value
		}
		if paramTypes[key] == "decimal.Decimal" {
			if decVal, err := decimal.NewFromString(value); err == nil {
				return decVal
			} else {
				// Логирование ошибки
				log.Printf("Error converting %s to decimal: %v", value, err)
				return nil
			}
		}

		if paramTypes[key] == "uint64" {
			if uintVal, err := strconv.ParseUint(value, 10, 64); err == nil {
				return uintVal
			} else {
				return nil
			}
		}

		// Проверяем тип данных из paramTypes
		if paramTypes[key] == "string" {
			return value
		}

		if floatVal, err := strconv.ParseFloat(value, 64); err == nil {
			return floatVal
		}
		return value
	}
}

func setupDecoder(result interface{}) (*mapstructure.Decoder, *typescore.WEvent) {
	decoderConfig := &mapstructure.DecoderConfig{
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			decodeHookStringSliceToString,
			decodeHookStringToPtrStringSlice,
			decodeHookStringToDecimal,
		),
		Result: result,
	}
	decoder, err := mapstructure.NewDecoder(decoderConfig)
	if err != nil {
		return nil, &typescore.WEvent{Err: err, Text: "system_error"}
	}
	return decoder, nil
}

func decodeHookStringSliceToString(f reflect.Type, t reflect.Type, data interface{}) (interface{}, error) {
	if t == reflect.TypeOf(new(string)) && f.Kind() == reflect.Slice {
		slice := reflect.ValueOf(data)
		strSlice := make([]string, slice.Len())
		for i := 0; i < slice.Len(); i++ {
			strSlice[i] = fmt.Sprint(slice.Index(i).Interface())
		}
		result := strings.Join(strSlice, ",")
		return &result, nil
	}
	return data, nil
}

func decodeHookStringToPtrStringSlice(f reflect.Type, t reflect.Type, data interface{}) (interface{}, error) {
	if t == reflect.TypeOf([]*string{}) && f.Kind() == reflect.String {
		str := data.(string)
		return []*string{&str}, nil
	}
	return data, nil
}

func decodeHookStringToDecimal(from reflect.Type, to reflect.Type, data interface{}) (interface{}, error) {
	if from.Kind() == reflect.String && to == reflect.TypeOf(&decimal.Decimal{}) {
		dec, err := decimal.NewFromString(data.(string))
		if err != nil {
			return nil, err
		}
		return &dec, nil
	}
	return data, nil
}
