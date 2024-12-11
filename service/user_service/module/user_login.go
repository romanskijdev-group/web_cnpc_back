package module

import (
	protoobj "cnpc_backend/core/proto"
	"cnpc_backend/core/typescore"
	"context"
	"errors"
	"log"
	"time"
)

// обрабатывает обнаруженного пользователя и возвращает информацию о входе
func (s *UserAccountServiceProto) userDetected(ctx context.Context, userObj *typescore.UsersProviderControl, isNewUser bool, userAuthReqAInfo *typescore.UserAuthReqAccountReq) (*protoobj.LogInInfoRes, error) {
	var err error

	if isNewUser {
		userObj, err = s.newUser(ctx, userObj, userAuthReqAInfo)
		if err != nil {
			log.Println("🔴 error UserLoginAccount: newUser: ", err)
			return nil, err
		}
	} else {
		err := s.userAuthNewInfoCombat(userObj, userAuthReqAInfo)
		if err != nil {
			log.Println("🔴 error UserLoginAccount: userAuthNewInfoCombat: ", err)
			return nil, err
		}
	}

	// Генерация токена для пользователя и возврат информации о входе
	return s.generateToken(userObj, isNewUser)
}

func (s *UserAccountServiceProto) userAuthNewInfoCombat(userObj *typescore.UsersProviderControl, userAuthReqAInfo *typescore.UserAuthReqAccountReq) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	if userObj == nil {
		return errors.New("user_not_finder")
	}
	// Проверка, заблокирован ли пользователь
	if userObj.IsBlocked != nil {
		if *userObj.IsBlocked {
			return errors.New("user_blocked")
		}
	}

	var newUserIp *string
	updateIpLogin := false

	if userAuthReqAInfo.DetectorIPStruct != nil && userAuthReqAInfo.DetectorIPStruct.IP != nil {
		if userObj.LastIP == nil {
			newUserIp = userAuthReqAInfo.DetectorIPStruct.IP
		} else if *userObj.LastIP != *userAuthReqAInfo.DetectorIPStruct.IP {
			newUserIp = userAuthReqAInfo.DetectorIPStruct.IP
		}
	}
	if newUserIp == nil {
		newUserIp = userObj.LastIP
	} else {
		updateIpLogin = true
	}

	userObj.LastIP = newUserIp
	lastLogin := time.Now().UTC()
	userObj.LastLogin = &lastLogin

	if updateIpLogin {
		err := s.sendDeviceAlertNotification(userObj, newUserIp)
		if err != nil {
			log.Println("🔴 error userAuthNewInfoCombat: sendDeviceAlertNotification: ", err)
		}
	}

	errW := s.ipc.Database.UsersActions.UpdateUserLastLoginInfoDB(ctx, userObj)
	if errW != nil {
		log.Println("🔴 error UserLoginAccount: UpdateUserLastLoginInfoDB: ", errW.Err)
	}

	return nil
}

// обрабатывает вход пользователя по email
func (s *UserAccountServiceProto) emailLogin(req *typescore.UserAuthReqAccountReq) (*typescore.UsersProviderControl, bool, error) {
	// Если тип входа - Email
	if req.Email == nil {
		return nil, true, errors.New("user_not_finder")
	}

	/// Проверка и удаление временного пароля из Redis
	err := s.ipc.RedisClient.CheckAndDeleteFromRedis(*req.Email, *req.TemporaryPassword, typescore.TempPassRedisType)
	if err != nil {
		log.Println("🔴 error UserLoginAccount: CheckAndDeleteFromRedis: ", err)
		return nil, true, err
	}

	findObjUer := &typescore.UsersProviderControl{
		Email: req.Email,
	}

	userObj := s.findUserInfo(findObjUer)
	if userObj == nil || userObj.SystemID == nil {
		return findObjUer, true, nil
	}

	err = s.sendLoginAlertNotification(userObj, req.AuthType)
	if err != nil {
		log.Println("🔴 error vkLogin: sendLoginAlertNotification: ", err)
	}

	return userObj, false, nil
}

// обрабатывает вход пользователя по vk
func (s *UserAccountServiceProto) vkLogin(req *typescore.UserAuthReqAccountReq) (*typescore.UsersProviderControl, bool, error) {
	// Если тип входа - Email
	if req.VKID == nil {
		return nil, true, errors.New("user_not_finder")
	}

	findObjUer := &typescore.UsersProviderControl{
		VKID: req.VKID,
	}

	userObj := s.findUserInfo(findObjUer)
	if userObj == nil || userObj.SystemID == nil {
		return findObjUer, true, nil
	}

	err := s.sendLoginAlertNotification(userObj, req.AuthType)
	if err != nil {
		log.Println("🔴 error vkLogin: sendLoginAlertNotification: ", err)
	}
	return userObj, false, nil
}
