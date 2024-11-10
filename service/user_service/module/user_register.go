package module

import (
	marshallerusers "cnpc_backend/core/module/user/users/marshaller"
	"cnpc_backend/core/typescore"
	"context"
	"log"
	"time"
)

// создает нового пользователя и возвращает информацию о нем
func (s *UserAccountServiceProto) newUser(ctx context.Context, userObj *typescore.UsersProviderControl, userAuthReqAInfo *typescore.UserAuthReqAccountReq) (*typescore.UsersProviderControl, error) {

	if userAuthReqAInfo != nil && userAuthReqAInfo.DetectorIPStruct != nil {
		// Устанавливаем IP-адрес последнего входа для нового пользователя
		userObj.LastIP = userAuthReqAInfo.DetectorIPStruct.IP
		utc := time.Now().UTC()
		userObj.LastLogin = &utc
	}

	// Сериализуем объект пользователя для передачи в сервис создания пользователя
	paramsNewUserObj := marshallerusers.UsersProviderControlSerialization(userObj)

	// Вызываем метод создания нового пользователя с переданными параметрами
	userNewPr, err := s.CreateNewUser(ctx, paramsNewUserObj)
	if err != nil {
		log.Println("🔴 error UserLoginAccount: CreateNewUser: ", err)
		return nil, err // Возвращаем ошибку, если создание пользователя не удалось
	}

	// Десериализуем ответ от сервиса создания пользователя в объект пользователя
	userNewInfo := marshallerusers.UsersProviderControlDeserialization(userNewPr)

	// Возвращаем информацию о новом пользователе
	return userNewInfo, nil
}
