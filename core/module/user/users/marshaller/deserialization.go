package marshalleruser

import (
	marshallerutils "zod_backend_dev/core/grpc_core/marshaller_utils"
	"zod_backend_dev/core/models"
	marshallersystem "zod_backend_dev/core/module/system/marshaller"
	protoobj "zod_backend_dev/core/proto"
)

func LogInInfoResDeserialization(obj *protoobj.LogInInfoRes) *models.LogInInfoRes {
	if obj == nil {
		return nil
	}

	tokenAuth := TokenInfoDeserialization(obj.TokenAuth)
	params := UserParamsLoginDeserialization(obj.Params)

	return &models.LogInInfoRes{
		TokenAuth: tokenAuth,
		Params:    params,
	}
}

func UserAuthReqAccountReqDeserialization(obj *protoobj.UserAuthReqAccountReq) *models.UserAuthReqAccountReq {
	if obj == nil {
		return nil
	}
	d := marshallerutils.InitDeserializationUtils()

	return &models.UserAuthReqAccountReq{
		TelegramID: d.OptionalInt64(obj.TelegramID),
		SystemID:   d.OptionalString(obj.SystemID),
		IsPremium:  d.OptionalBool(obj.IsPremium),
		Username:   d.OptionalString(obj.Username),
		PhotoURL:   d.OptionalString(obj.PhotoURL),
		ReferralID: d.OptionalInt64(obj.ReferralID),
		Language:   d.OptionalString(obj.Language),
		AuthType:   mapTypeAuthDeserialization(obj.AuthType),
	}
}

func UsersDeserialization(obj *protoobj.UsersMsg) *models.User {
	if obj == nil {
		return nil
	}
	d := marshallerutils.InitDeserializationUtils()

	return &models.User{
		SystemID:             d.OptionalString(obj.SystemId),
		SerialID:             d.OptionalUint64(obj.SerialId),
		Role:                 mapRoleDeserialization(obj.Role),
		TelegramID:           d.OptionalInt64(obj.TelegramId),
		Username:             d.OptionalString(obj.Username),
		FirstName:            d.OptionalString(obj.FirstName),
		LastName:             d.OptionalString(obj.LastName),
		Gender:               d.OptionalString(obj.Gender),
		Zodiac:               marshallersystem.MapZodiacSignsDeserialization(obj.Zodiac),
		IsBlocked:            d.OptionalBool(obj.IsBlocked),
		IsPremium:            d.OptionalBool(obj.IsPremium),
		BirthDate:            d.OptionalTime(obj.BirthDate),
		BirthTime:            d.OptionalTime(obj.BirthTime),
		BirthPlace:           d.OptionalString(obj.BirthPlace),
		NotificationsEnabled: d.OptionalBool(obj.NotificationEnabled),
		Language:             d.OptionalString(obj.Language),
		ReferralID:           d.OptionalString(obj.ReferralId),
		ReferralCount:        d.OptionalDecimal(obj.ReferralCount),
		Balance:              d.OptionalDecimal(obj.Balance),
		CreatedAt:            d.OptionalTime(obj.CreatedAt),
		LastLogin:            d.OptionalTime(obj.LastLogin),
		AvatarURL:            d.OptionalString(obj.AvatarUrl),
	}

}

func ShortUserInfoDeserialization(obj *protoobj.ShortUserInfo) *models.ShortUserInfo {
	if obj == nil {
		return nil
	}
	d := marshallerutils.InitDeserializationUtils()

	return &models.ShortUserInfo{
		SerialID:   d.OptionalUint64(obj.SerialId),
		TelegramID: d.OptionalInt64(obj.TelegramId),
		Username:   d.OptionalString(obj.Nickname),
		FirstName:  d.OptionalString(obj.FirstName),
		LastName:   d.OptionalString(obj.LastName),
	}
}

func UserMsgReqDeserialization(obj *protoobj.UsersMsgReq) (*models.User, *uint64, *uint64, map[string]string) {
	if obj == nil {
		return nil, nil, nil, nil
	}
	d := marshallerutils.InitDeserializationUtils()
	user := UsersDeserialization(obj.ParamsFiltering)
	return user,
		d.OptionalUint64(obj.Offset),
		d.OptionalUint64(obj.Limit), obj.LikeFields

}

func TokenInfoDeserialization(obj *protoobj.TokenInfo) *models.TokenInfo {
	if obj == nil {
		return nil
	}
	return &models.TokenInfo{
		AccessToken: obj.AccessToken,
		ExpiresIn:   obj.ExpiresIn,
	}
}

func UserParamsLoginDeserialization(obj *protoobj.UserParamsLogin) *models.UserParamsLogin {
	if obj == nil {
		return nil
	}
	d := marshallerutils.InitDeserializationUtils()

	return &models.UserParamsLogin{
		Language:  d.OptionalString(obj.Language),
		IsNewUser: d.OptionalBool(obj.IsNewUser),
	}
}

func mapTypeAuthDeserialization(authType protoobj.TypeAuth) *models.TypeAuth {
	switch authType {
	case protoobj.TypeAuth_TypeAuth_telegram:
		t := models.TelegramType
		return &t
	case protoobj.TypeAuth_TypeAuth_token_auth:
		t := models.AuthTokenType
		return &t
	default:
		return nil
	}
}

func mapRoleDeserialization(roleTypeMsg protoobj.UserRole) *models.UserRoleTypes {
	switch roleTypeMsg {
	case protoobj.UserRole_USER:
		st := models.UserRole
		return &st
	case protoobj.UserRole_ADMIN:
		st := models.AdminRole
		return &st
	case protoobj.UserRole_ASTROLOGER:
		st := models.AstrologerRole
		return &st
	default:
		return nil
	}
}

func UsersListDeserialization(obj *protoobj.UsersMsgList) []*models.User {
	if obj == nil {
		return nil
	}
	var res []*models.User
	for _, item := range obj.UsersMsg {
		protoItem := UsersDeserialization(item)
		if protoItem != nil {
			res = append(res, protoItem)
		}
	}
	return res
}
