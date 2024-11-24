package marshallerusers

// распаковка структур grpc в структуры модуля

import (
	marshallerutils "cnpc_backend/core/grpc_core/marshaller_utils"
	protoobj "cnpc_backend/core/proto"
	"cnpc_backend/core/typescore"
)

func RegionInfoDetectedDeserialization(obj *protoobj.RegionInfoDetected) *typescore.RegionInfoDetected {
	if obj == nil {
		return nil
	}
	d := marshallerutils.InitDeserializationUtils()
	return &typescore.RegionInfoDetected{
		City:        d.OptionalString(obj.City),
		Region:      d.OptionalString(obj.Region),
		CountryCode: d.OptionalString(obj.CountryCode),
		CountryName: d.OptionalString(obj.CountryName),
	}
}

func DetectorIPStructDeserialization(obj *protoobj.DetectorIPStruct) *typescore.DetectorIPStruct {
	if obj == nil {
		return nil
	}
	d := marshallerutils.InitDeserializationUtils()
	regionInfo := RegionInfoDetectedDeserialization(obj.RegionInfo)

	return &typescore.DetectorIPStruct{
		IP:            d.OptionalString(obj.IP),
		IsINBlackList: d.OptionalBool(obj.IsINBlackList),
		RegionInfo:    regionInfo,
	}
}

func UserAuthReqAccountReqDeserialization(obj *protoobj.UserAuthReqAccountReq) *typescore.UserAuthReqAccountReq {
	if obj == nil {
		return nil
	}
	d := marshallerutils.InitDeserializationUtils()

	return &typescore.UserAuthReqAccountReq{
		Email:             d.OptionalString(obj.Email),
		TemporaryPassword: d.OptionalString(obj.TemporaryPassword),
		TelegramID:        d.OptionalInt64(obj.TelegramID),
		EmailCode:         d.OptionalString(obj.EmailCode),
		SystemID:          d.OptionalString(obj.SystemID),
		VKID:              d.OptionalInt64(obj.VKID),
		Code:              d.OptionalString(obj.Code),
		Secret:            d.OptionalString(obj.Secret),
		DetectorIPStruct:  DetectorIPStructDeserialization(obj.DetectorIPStruct),
		AuthType:          mapTypeAuthDeserialization(obj.AuthType),
	}
}

func LogInInfoResDeserialization(obj *protoobj.LogInInfoRes) *typescore.LogInInfoRes {
	if obj == nil {
		return nil
	}

	tokenAuth := TokenInfoDeserialization(obj.TokenAuth)
	params := UserParamsLoginDeserialization(obj.Params)

	return &typescore.LogInInfoRes{
		TokenAuth: tokenAuth,
		Params:    params,
	}
}

func TokenInfoDeserialization(obj *protoobj.TokenInfo) *typescore.TokenInfo {
	if obj == nil {
		return nil
	}
	return &typescore.TokenInfo{
		AccessToken: obj.AccessToken,
		ExpiresIn:   obj.ExpiresIn,
	}
}

func UserParamsLoginDeserialization(obj *protoobj.UserParamsLogin) *typescore.UserParamsLogin {
	if obj == nil {
		return nil
	}
	d := marshallerutils.InitDeserializationUtils()

	return &typescore.UserParamsLogin{
		Language:  d.OptionalString(obj.Language),
		IsNewUser: d.OptionalBool(obj.IsNewUser),
	}
}

func UsersProviderControlMsgListDeserialization(obj *protoobj.UsersMsgList) []*typescore.UsersProviderControl {
	if obj == nil {
		return nil
	}
	var res []*typescore.UsersProviderControl
	for _, item := range obj.UsersMsg {
		deserializedItem := UsersProviderControlDeserialization(item)
		if deserializedItem != nil {
			res = append(res, deserializedItem)
		}
	}
	return res
}

func UpdateUserAvatarURLReqDeserialization(obj *protoobj.UpdateUserAvatarURLReq) (*string, *string) {
	if obj == nil {
		return nil, nil
	}
	d := marshallerutils.InitDeserializationUtils()
	return d.OptionalString(obj.UserSystemID), d.OptionalString(obj.AvatarURL)
}

func UsersProviderControlDeserialization(obj *protoobj.UsersMsg) *typescore.UsersProviderControl {
	if obj == nil {
		return nil
	}
	d := marshallerutils.InitDeserializationUtils()

	return &typescore.UsersProviderControl{
		SystemID:            d.OptionalString(obj.SystemId),
		SerialID:            d.OptionalUint64(obj.SerialId),
		Role:                mapRoleDeserialization(obj.Role),
		Email:               d.OptionalString(obj.Email),
		TelegramID:          d.OptionalInt64(obj.TelegramId),
		VKID:                d.OptionalInt64(obj.VkId),
		Nickname:            d.OptionalString(obj.Nickname),
		FirstName:           d.OptionalString(obj.FirstName),
		LastName:            d.OptionalString(obj.LastName),
		BirthDate:           d.OptionalTime(obj.BirthDate),
		PhoneNumber:         d.OptionalUint64(obj.PhoneNumber),
		AvatarURL:           d.OptionalString(obj.AvatarUrl),
		Language:            d.OptionalString(obj.Language),
		NotificationEnabled: d.OptionalBool(obj.NotificationEnabled),
		IsBlocked:           d.OptionalBool(obj.IsBlocked),
		ReferralID:          d.OptionalString(obj.ReferralId),
		ReferralCode:        d.OptionalString(obj.ReferralCode),
		LastIP:              d.OptionalString(obj.LastIp),
		CreatedAt:           d.OptionalTime(obj.CreatedAt),
		LastLogin:           d.OptionalTime(obj.LastLogin),

		IsOnline:   d.OptionalBool(obj.IsOnline),
		LastOnline: d.OptionalTime(obj.LastOnline),
	}
}

func ShortUserInfoDeserialization(obj *protoobj.ShortUserInfo) *typescore.ShortUserInfo {
	if obj == nil {
		return nil
	}
	d := marshallerutils.InitDeserializationUtils()

	return &typescore.ShortUserInfo{
		SerialID:   d.OptionalUint64(obj.SerialId),
		Email:      d.OptionalString(obj.Email),
		TelegramID: d.OptionalInt64(obj.TelegramId),
		Nickname:   d.OptionalString(obj.Nickname),
		FirstName:  d.OptionalString(obj.FirstName),
		LastName:   d.OptionalString(obj.LastName),
		ParentName: d.OptionalString(obj.ParentName),
		AvatarURL:  d.OptionalString(obj.AvatarUrl),

		CreatedAt: d.OptionalStringTimeOnlyDate(obj.CreatedAt),

		IsOnline:   d.OptionalBool(obj.IsOnline),
		LastOnline: d.OptionalTime(obj.LastOnline),
	}
}

func mapRoleDeserialization(roleTypeMsg protoobj.UserRole) *typescore.UserRoleTypes {
	switch roleTypeMsg {
	case protoobj.UserRole_USER:
		st := typescore.UserRole
		return &st
	case protoobj.UserRole_ADMIN:
		st := typescore.AdminRole
		return &st
	case protoobj.UserRole_SUPER_ADMIN:
		st := typescore.SuperAdminRole
		return &st
	case protoobj.UserRole_SUPPORT:
		st := typescore.SupportRole
		return &st
	default:
		return nil
	}
}

func mapTypeAuthDeserialization(authType protoobj.TypeAuth) *typescore.TypeAuth {
	switch authType {
	case protoobj.TypeAuth_TypeAuth_email_auth:
		t := typescore.EmailType
		return &t
	case protoobj.TypeAuth_TypeAuth_vk:
		t := typescore.VKType
		return &t
	case protoobj.TypeAuth_TypeAuth_telegram:
		t := typescore.TelegramType
		return &t
	case protoobj.TypeAuth_TypeAuth_token_auth:
		t := typescore.AuthTokenType
		return &t
	default:
		return nil
	}
}
