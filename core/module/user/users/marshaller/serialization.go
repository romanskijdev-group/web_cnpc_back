package marshallerusers

import (
	marshallerutils "cnpc_backend/core/grpc_core/marshaller_utils"
	protoobj "cnpc_backend/core/proto"
	"cnpc_backend/core/typescore"
)

func LogInInfoResSerialization(obj *typescore.LogInInfoRes) *protoobj.LogInInfoRes {
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

func TokenInfoSerialization(obj *typescore.TokenInfo) *protoobj.TokenInfo {
	if obj == nil {
		return nil
	}
	return &protoobj.TokenInfo{
		AccessToken: obj.AccessToken,
		ExpiresIn:   obj.ExpiresIn,
	}
}

func UserParamsLoginSerialization(obj *typescore.UserParamsLogin) *protoobj.UserParamsLogin {
	if obj == nil {
		return nil
	}
	s := marshallerutils.InitSerializationUtils()

	return &protoobj.UserParamsLogin{
		Language:  s.StringToWrapperStringValue(obj.Language),
		IsNewUser: s.BoolToWrapperBoolValue(obj.IsNewUser),
	}
}

func RegionInfoDetectedSerialization(obj *typescore.RegionInfoDetected) *protoobj.RegionInfoDetected {
	if obj == nil {
		return nil
	}
	s := marshallerutils.InitSerializationUtils()
	return &protoobj.RegionInfoDetected{
		City:        s.StringToWrapperStringValue(obj.City),
		Region:      s.StringToWrapperStringValue(obj.Region),
		CountryCode: s.StringToWrapperStringValue(obj.CountryCode),
		CountryName: s.StringToWrapperStringValue(obj.CountryName),
	}
}

func DetectorIPStructSerialization(obj *typescore.DetectorIPStruct) *protoobj.DetectorIPStruct {
	if obj == nil {
		return nil
	}
	s := marshallerutils.InitSerializationUtils()
	regionInfo := RegionInfoDetectedSerialization(obj.RegionInfo)

	return &protoobj.DetectorIPStruct{
		IP:            s.StringToWrapperStringValue(obj.IP),
		IsINBlackList: s.BoolToWrapperBoolValue(obj.IsINBlackList),
		RegionInfo:    regionInfo,
	}
}

func UserAuthReqAccountReqSerialization(obj *typescore.UserAuthReqAccountReq) *protoobj.UserAuthReqAccountReq {
	if obj == nil {
		return nil
	}
	s := marshallerutils.InitSerializationUtils()

	return &protoobj.UserAuthReqAccountReq{
		Email:             s.StringToWrapperStringValue(obj.Email),
		TemporaryPassword: s.StringToWrapperStringValue(obj.TemporaryPassword),
		TelegramID:        s.Int64ToWrapperInt64Value(obj.TelegramID),
		EmailCode:         s.StringToWrapperStringValue(obj.EmailCode),
		SystemID:          s.StringToWrapperStringValue(obj.SystemID),
		VKID:              s.Int64ToWrapperInt64Value(obj.VKID),
		Code:              s.StringToWrapperStringValue(obj.Code),
		Secret:            s.StringToWrapperStringValue(obj.Secret),
		DetectorIPStruct:  DetectorIPStructSerialization(obj.DetectorIPStruct),
		AuthType:          mapTypeAuthSerialization(obj.AuthType),
	}
}

func UsersProviderControlMsgListSerialization(obj []*typescore.UsersProviderControl) *protoobj.UsersMsgList {
	if obj == nil {
		return nil
	}
	var res []*protoobj.UsersMsg
	for _, item := range obj {
		protoItem := UsersProviderControlSerialization(item)
		if protoItem != nil {
			res = append(res, protoItem)
		}
	}
	return &protoobj.UsersMsgList{
		UsersMsg: res,
	}
}

func UpdateUserAvatarURLReqSerialization(userSystemID *string, avatarURL *string) *protoobj.UpdateUserAvatarURLReq {
	if userSystemID == nil {
		return nil
	}
	s := marshallerutils.InitSerializationUtils()
	return &protoobj.UpdateUserAvatarURLReq{
		UserSystemID: s.StringToWrapperStringValue(userSystemID),
		AvatarURL:    s.StringToWrapperStringValue(avatarURL),
	}
}

func UsersProviderControlSerialization(obj *typescore.UsersProviderControl) *protoobj.UsersMsg {
	if obj == nil {
		return nil
	}
	s := marshallerutils.InitSerializationUtils()

	return &protoobj.UsersMsg{
		SystemId:            s.StringToWrapperStringValue(obj.SystemID),
		SerialId:            s.Uint64ToWrapperUInt64Value(obj.SerialID),
		Role:                mapRoleSerialization(obj.Role),
		Email:               s.StringToWrapperStringValue(obj.Email),
		TelegramId:          s.Int64ToWrapperInt64Value(obj.TelegramID),
		VkId:                s.Int64ToWrapperInt64Value(obj.VKID),
		Nickname:            s.StringToWrapperStringValue(obj.Nickname),
		FirstName:           s.StringToWrapperStringValue(obj.FirstName),
		LastName:            s.StringToWrapperStringValue(obj.LastName),
		BirthDate:           s.TimePtrToTimestampPB(obj.BirthDate),
		PhoneNumber:         s.Uint64ToWrapperUInt64Value(obj.PhoneNumber),
		AvatarUrl:           s.StringToWrapperStringValue(obj.AvatarURL),
		Language:            s.StringToWrapperStringValue(obj.Language),
		NotificationEnabled: s.BoolToWrapperBoolValue(obj.NotificationEnabled),
		IsBlocked:           s.BoolToWrapperBoolValue(obj.IsBlocked),

		ReferralId:   s.StringToWrapperStringValue(obj.ReferralID),
		ReferralCode: s.StringToWrapperStringValue(obj.ReferralCode),
		LastIp:       s.StringToWrapperStringValue(obj.LastIP),
		CreatedAt:    s.TimePtrToTimestampPB(obj.CreatedAt),
		LastLogin:    s.TimePtrToTimestampPB(obj.LastLogin),

		IsOnline:   s.BoolToWrapperBoolValue(obj.IsOnline),
		LastOnline: s.TimePtrToTimestampPB(obj.LastOnline),
	}
}

func ShortUserInfoSerialization(obj *typescore.ShortUserInfo) *protoobj.ShortUserInfo {
	if obj == nil {
		return nil
	}
	s := marshallerutils.InitSerializationUtils()
	return &protoobj.ShortUserInfo{
		SerialId:   s.Uint64ToWrapperUInt64Value(obj.SerialID),
		Email:      s.StringToWrapperStringValue(obj.Email),
		TelegramId: s.Int64ToWrapperInt64Value(obj.TelegramID),
		Nickname:   s.StringToWrapperStringValue(obj.Nickname),
		FirstName:  s.StringToWrapperStringValue(obj.FirstName),
		LastName:   s.StringToWrapperStringValue(obj.LastName),
		ParentName: s.StringToWrapperStringValue(obj.ParentName),
		AvatarUrl:  s.StringToWrapperStringValue(obj.AvatarURL),
		CreatedAt:  s.StringTimeToWrapperOnlyDate(obj.CreatedAt),
		IsOnline:   s.BoolToWrapperBoolValue(obj.IsOnline),
		LastOnline: s.TimePtrToTimestampPB(obj.LastOnline),
	}
}

func mapRoleSerialization(roleType *typescore.UserRoleTypes) protoobj.UserRole {
	if roleType == nil {
		return protoobj.UserRole_USER
	}
	switch *roleType {
	case typescore.UserRole:
		return protoobj.UserRole_USER
	case typescore.AdminRole:
		return protoobj.UserRole_ADMIN
	case typescore.SuperAdminRole:
		return protoobj.UserRole_SUPER_ADMIN
	case typescore.SupportRole:
		return protoobj.UserRole_SUPPORT
	default:
		return protoobj.UserRole_USER
	}
}

func mapTypeAuthSerialization(authType *typescore.TypeAuth) protoobj.TypeAuth {
	if authType == nil {
		return protoobj.TypeAuth_TypeAuth_NULL
	}
	switch *authType {
	case typescore.EmailType:
		return protoobj.TypeAuth_TypeAuth_email_auth
	case typescore.VKType:
		return protoobj.TypeAuth_TypeAuth_vk
	case typescore.TelegramType:
		return protoobj.TypeAuth_TypeAuth_telegram
	case typescore.AuthTokenType:
		return protoobj.TypeAuth_TypeAuth_token_auth
	default:
		return protoobj.TypeAuth_TypeAuth_NULL
	}
}

func UserMsgReqSerialization(user *typescore.UsersProviderControl,
	offset *uint64, limit *uint64, likesFiled map[string]string) *protoobj.UsersMsgReq {
	if user == nil {
		return nil
	}
	s := marshallerutils.InitSerializationUtils()
	userObj := UsersProviderControlSerialization(user)

	return &protoobj.UsersMsgReq{
		ParamsFiltering: userObj,
		Offset:          s.Uint64ToWrapperUInt64Value(offset),
		Limit:           s.Uint64ToWrapperUInt64Value(limit),
		LikeFields:      likesFiled,
	}
}

func UserMsgReqDeserialization(obj *protoobj.UsersMsgReq) (*typescore.UsersProviderControl, *uint64, *uint64, map[string]string) {
	if obj == nil {
		return nil, nil, nil, nil
	}
	d := marshallerutils.InitDeserializationUtils()
	user := UsersProviderControlDeserialization(obj.ParamsFiltering)
	return user,
		d.OptionalUint64(obj.Offset),
		d.OptionalUint64(obj.Limit), obj.LikeFields

}
