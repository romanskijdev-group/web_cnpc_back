package module

import (
	marshallerusers "cnpc_backend/core/module/user/users/marshaller"
	protoobj "cnpc_backend/core/proto"
	"cnpc_backend/core/typescore"
	"cnpc_backend/core/utilscore"
	"context"
	"errors"
	"log"
	"time"
)

// Ищет информацию о пользователе на основе типа входа и данных пользователя
func (s *UserAccountServiceProto) authUser(req *typescore.UserAuthReqAccountReq) (*typescore.UsersProviderControl, bool, error) {
	if req == nil {
		return nil, true, errors.New("invalid_request_body")
	}
	if req.AuthType == nil {
		return nil, true, errors.New("invalid login type")
	}

	// Определяем действия в зависимости от типа входа
	switch *req.AuthType {
	case typescore.AuthTokenType:
		// Если тип входа - токен аутентификации
		return s.authTokenLogin(req)
	//case typescore.TelegramType:
	//	return s.telegramLogin(req)
	case typescore.EmailType:
		// Если тип входа - Email
		return s.emailLogin(req)
	case typescore.VKType:
		// Если тип входа - VK
		return s.vkLogin(req)
	default:
		// Если тип входа не распознан
		return nil, true, errors.New("invalid login type")
	}
}

// Обрабатывает вход пользователя по токену аутентификации
func (s *UserAccountServiceProto) authTokenLogin(userAuthReqAccountReq *typescore.UserAuthReqAccountReq) (*typescore.UsersProviderControl, bool, error) {
	if userAuthReqAccountReq.SystemID == nil {
		return nil, true, errors.New("user_not_finder")
	}

	findObjUer := &typescore.UsersProviderControl{
		SystemID: userAuthReqAccountReq.SystemID,
	}
	userObj := s.findUserInfo(findObjUer)

	if userObj == nil || userObj.SystemID == nil {
		return findObjUer, true, nil
	}
	return userObj, false, nil
}

// Ищет информацию о пользователе в базе данных
func (s *UserAccountServiceProto) findUserInfo(userAuthReqAccountReq *typescore.UsersProviderControl) *typescore.UsersProviderControl {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Поиск пользователя в базе данных
	userObj, errW := s.ipc.Database.UsersActions.GetUserDB(ctx, userAuthReqAccountReq)
	if errW != nil {
		log.Println("🔴 error findUserInfo UserLoginAccount: GetUserDB: ", errW.Err)
		userObj = nil
	}
	// Возвращаем найденного пользователя или nil, если пользователь не найден
	return userObj
}

// Генерирует токен для пользователя и возвращает информацию о входе
func (s *UserAccountServiceProto) generateToken(userBaseData *typescore.UsersProviderControl, isNewUser bool) (*protoobj.LogInInfoRes, error) {
	// Получаем токен аутентификации для пользователя
	tokenObj, errW := s.ipc.Modules.RestAuth.TokenRestAuth.GetAuthToken(userBaseData)
	if errW != nil {
		log.Println("🔴 error UserLoginAccount: GetAuthToken: ", errW.Err)
		return nil, errW.Err // Возвращаем ошибку, если не удалось получить токен
	}

	// Сериализуем информацию о входе пользователя и возвращаем результат
	return marshallerusers.LogInInfoResSerialization(&typescore.LogInInfoRes{
		TokenAuth: tokenObj, // Токен аутентификации
		Params: &typescore.UserParamsLogin{
			Language:  userBaseData.Language,              // Язык пользователя
			IsNewUser: utilscore.PointerToBool(isNewUser), // Является ли пользователь новым
		},
	}), nil
}
