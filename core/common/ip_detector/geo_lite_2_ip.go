package ipdetector

import (
	"cnpc_backend/core/typescore"
	_ "embed"
	"net"
	"os"
	"strings"
	"sync"

	"github.com/oschwald/geoip2-golang"
)

//go:embed GeoLite2-City.mmdb
var geoLite2CityMMDB []byte

var (
	dbInstance *geoip2.Reader
	once       sync.Once
)

// initDB инициализирует базу данных GeoIP2
func initGeoip2DB() error {
	var err error
	once.Do(func() {
		// Создание временного файла для базы данных
		tmpFile, err := os.CreateTemp("", "GeoLite2-City.mmdb")
		if err != nil {
			return
		}
		defer os.Remove(tmpFile.Name())

		// Запись встроенного файла в временный файл
		if _, err = tmpFile.Write(geoLite2CityMMDB); err != nil {
			return
		}
		if err = tmpFile.Close(); err != nil {
			return
		}

		// Открытие базы данных GeoIP2 из временного файла
		dbInstance, err = geoip2.Open(tmpFile.Name())
	})
	return err
}

// IPdetectorM - структура для работы с IP-адресами
func (m *IPdetectorM) IpPositionDetect(ip string) (*typescore.RegionInfoDetected, error) {
	if err := initGeoip2DB(); err != nil {
		return nil, err
	}

	// Пример IP-адреса для поиска
	ipNet := net.ParseIP(ip)

	// Получение информации о городе
	record, err := dbInstance.City(ipNet)
	if err != nil {
		return nil, err
	}

	city := record.City.Names["en"]
	countyCode := strings.ToLower(record.Country.IsoCode)
	countryName := record.Country.Names["en"]
	continent := strings.ToLower(record.Continent.Names["en"])

	return &typescore.RegionInfoDetected{
		City:        &city,
		CountryCode: &countyCode,
		CountryName: &countryName,
		Region:      &continent,
	}, nil
}
