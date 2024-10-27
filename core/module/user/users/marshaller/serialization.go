package marshalleruser

import (
	marshallerutils "zod_backend_dev/core/grpc_core/marshaller_utils"
	"zod_backend_dev/core/models"
	marshallersystem "zod_backend_dev/core/module/system/marshaller"
	protoobj "zod_backend_dev/core/proto"
)

func UserAuthReqAccountReqSerialization(obj *models.UserAuthReqAccountReq) *protoobj.UserAuthReqAccountReq {
	if obj == nil {
		return nil
	}
	s := marshallerutils.InitSerializationUtils()

	return &protoobj.UserAuthReqAccountReq{
		TelegramID: s.Int64ToWrapperInt64Value(obj.TelegramID),
		Username:   s.StringToWrapperStringValue(obj.Username),
		FirstName:  s.StringToWrapperStringValue(obj.FirstName),
		IsPremium:  s.BoolToWrapperBoolValue(obj.IsPremium),
		SystemID:   s.StringToWrapperStringValue(obj.SystemID),
		Language:   s.StringToWrapperStringValue(obj.Language),
		ReferralID: s.Int64ToWrapperInt64Value(obj.ReferralID),
		AuthType:   mapTypeAuthSerialization(obj.AuthType),
	}
}

func LogInInfoResSerialization(obj *models.LogInInfoRes) *protoobj.LogInInfoRes {
	if obj == nil {
		return nil
	}
	tokenAuth := TokenInfoSerialization(obj.TokenAuth)
	params := UserParamsLoginSerialization(obj.Params)

	return &protoobj.LogInInfoRes{
		TokenAuth: tokenAuth,
		Params:    params,
	}
}

func UserMsgReqSerialization(user *models.User,
	offset *uint64, limit *uint64, likesFiled map[string]string) *protoobj.UsersMsgReq {
	if user == nil {
		return nil
	}
	s := marshallerutils.InitSerializationUtils()
	userObj := UsersSerialization(user)

	return &protoobj.UsersMsgReq{
		ParamsFiltering: userObj,
		Offset:          s.Uint64ToWrapperUInt64Value(offset),
		Limit:           s.Uint64ToWrapperUInt64Value(limit),
		LikeFields:      likesFiled,
	}
}

func UsersSerialization(obj *models.User) *protoobj.UsersMsg {
	if obj == nil {
		return nil
	}
	s := marshallerutils.InitSerializationUtils()

	return &protoobj.UsersMsg{
		SystemId:            s.StringToWrapperStringValue(obj.SystemID),
		SerialId:            s.Uint64ToWrapperUInt64Value(obj.SerialID),
		Role:                mapRoleSerialization(obj.Role),
		TelegramId:          s.Int64ToWrapperInt64Value(obj.TelegramID),
		Username:            s.StringToWrapperStringValue(obj.Username),
		FirstName:           s.StringToWrapperStringValue(obj.FirstName),
		LastName:            s.StringToWrapperStringValue(obj.LastName),
		Gender:              s.StringToWrapperStringValue(obj.Gender),
		Zodiac:              marshallersystem.MapZodiacSignsSerialization(obj.Zodiac),
		IsBlocked:           s.BoolToWrapperBoolValue(obj.IsBlocked),
		IsPremium:           s.BoolToWrapperBoolValue(obj.IsBlocked),
		BirthDate:           s.TimePtrToTimestampPB(obj.BirthDate),
		BirthTime:           s.TimePtrToTimestampPB(obj.BirthTime),
		BirthPlace:          s.StringToWrapperStringValue(obj.BirthPlace),
		NotificationEnabled: s.BoolToWrapperBoolValue(obj.NotificationsEnabled),
		Language:            s.StringToWrapperStringValue(obj.Language),
		ReferralId:          s.StringToWrapperStringValue(obj.ReferralID),
		ReferralCount:       s.DecimalToWrapperStringValue(obj.ReferralCount),
		Balance:             s.DecimalToWrapperStringValue(obj.Balance),
		CreatedAt:           s.TimePtrToTimestampPB(obj.CreatedAt),
		LastLogin:           s.TimePtrToTimestampPB(obj.LastLogin),
		AvatarUrl:           s.StringToWrapperStringValue(obj.AvatarURL),
	}
}

func UsersListSerialization(obj []*models.User) *protoobj.UsersMsgList {
	if obj == nil {
		return nil
	}
	var res []*protoobj.UsersMsg
	for _, item := range obj {
		protoItem := UsersSerialization(item)
		if protoItem != nil {
			res = append(res, protoItem)
		}
	}
	return &protoobj.UsersMsgList{
		UsersMsg: res,
	}
}

func TokenInfoSerialization(obj *models.TokenInfo) *protoobj.TokenInfo {
	if obj == nil {
		return nil
	}
	return &protoobj.TokenInfo{
		AccessToken: obj.AccessToken,
		ExpiresIn:   obj.ExpiresIn,
	}
}

func UserParamsLoginSerialization(obj *models.UserParamsLogin) *protoobj.UserParamsLogin {
	if obj == nil {
		return nil
	}
	s := marshallerutils.InitSerializationUtils()

	return &protoobj.UserParamsLogin{
		Language:  s.StringToWrapperStringValue(obj.Language),
		IsNewUser: s.BoolToWrapperBoolValue(obj.IsNewUser),
	}
}

func mapRoleSerialization(roleType *models.UserRoleTypes) protoobj.UserRole {
	if roleType == nil {
		return protoobj.UserRole_USER
	}
	switch *roleType {
	case models.UserRole:
		return protoobj.UserRole_USER
	case models.AdminRole:
		return protoobj.UserRole_ADMIN
	case models.AstrologerRole:
		return protoobj.UserRole_ASTROLOGER
	default:
		return protoobj.UserRole_USER
	}
}

func mapTypeAuthSerialization(authType *models.TypeAuth) protoobj.TypeAuth {
	if authType == nil {
		return protoobj.TypeAuth_TypeAuth_NULL
	}
	switch *authType {
	case models.TelegramType:
		return protoobj.TypeAuth_TypeAuth_telegram
	case models.AuthTokenType:
		return protoobj.TypeAuth_TypeAuth_token_auth
	default:
		return protoobj.TypeAuth_TypeAuth_NULL
	}
}

func ShortUserInfoSerialization(obj *models.ShortUserInfo) *protoobj.ShortUserInfo {
	if obj == nil {
		return nil
	}
	s := marshallerutils.InitSerializationUtils()
	return &protoobj.ShortUserInfo{
		SerialId:   s.Uint64ToWrapperUInt64Value(obj.SerialID),
		TelegramId: s.Int64ToWrapperInt64Value(obj.TelegramID),
		Nickname:   s.StringToWrapperStringValue(obj.Username),
		FirstName:  s.StringToWrapperStringValue(obj.FirstName),
		LastName:   s.StringToWrapperStringValue(obj.LastName),
	}
}
