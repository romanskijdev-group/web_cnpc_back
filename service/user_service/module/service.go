package module

import (
	marshallerusers "cnpc_backend/core/module/user/users/marshaller"
	protoobj "cnpc_backend/core/proto"
	"context"
	"errors"
	"log"
	"userservice/types"
)

func NewUserAccountServiceProto(ipc *types.InternalProviderControl) protoobj.UserAccountServiceProtoServer {
	return &UserAccountServiceProto{
		ipc: ipc,
	}
}

type UserAccountServiceProto struct {
	ipc *types.InternalProviderControl
	protoobj.UnimplementedUserAccountServiceProtoServer
}

// вход пользователя в аккаунт
// обрабатывает запрос на вход пользователя в аккаунт
func (s *UserAccountServiceProto) UserLoginAccount(ctx context.Context, obj *protoobj.UserAuthReqAccountReq) (*protoobj.LogInInfoRes, error) {
	// Десериализация запроса на вход пользователя
	userAuthReqAInfo := marshallerusers.UserAuthReqAccountReqDeserialization(obj)
	if userAuthReqAInfo == nil {
		return nil, errors.New("invalid_request_body")
	}

	// Поиск информации о пользователе на основе запроса и данных пользователя
	userObj, isNewUser, err := s.authUser(userAuthReqAInfo)
	if err != nil {
		log.Println("🔴 error UserLoginAccount: findUserInfo: ", err)
		return nil, errors.New("not_found")
	}

	// Обработка обнаруженного пользователя и возврат информации о входе
	return s.userDetected(ctx, userObj, isNewUser, userAuthReqAInfo)
}

// получение информаций о профиле пользователя
func (s *UserAccountServiceProto) GetUserProfile(ctx context.Context, obj *protoobj.UsersMsg) (*protoobj.UsersMsg, error) {
	// logrus.Info("🚀 GetUserProfile")
	paramsObj := marshallerusers.UsersProviderControlDeserialization(obj)

	objRes, errW := s.ipc.Database.UsersActions.GetUserDB(ctx, paramsObj)
	if errW != nil {
		return nil, errW.Err
	}
	return marshallerusers.UsersProviderControlSerialization(objRes), nil
}

// получение информаций о профилях пользователей
func (s *UserAccountServiceProto) GetUsersInfoList(ctx context.Context, obj *protoobj.UsersMsg) (*protoobj.UsersMsgList, error) {
	// logrus.Info("🚀 GetUsersInfoList")
	paramsObj := marshallerusers.UsersProviderControlDeserialization(obj)

	objList, errW := s.ipc.Database.UsersActions.GetUsersListDB(ctx, paramsObj, map[string]string{}, nil, nil)
	if errW != nil {
		return nil, errW.Err
	}

	return marshallerusers.UsersProviderControlMsgListSerialization(objList), nil
}

// обновление информаций о пользователе
func (s *UserAccountServiceProto) UpdateUserProfile(ctx context.Context, obj *protoobj.UsersMsg) (*protoobj.UsersMsg, error) {
	// logrus.Info("🚀 UpdateUserProfile")
	paramsObj := marshallerusers.UsersProviderControlDeserialization(obj)

	userUp, errW := s.ipc.Database.UsersActions.UpdateUserDB(ctx, paramsObj)
	if errW != nil {
		return nil, errW.Err
	}
	return marshallerusers.UsersProviderControlSerialization(userUp), nil
}

// создание пользователя
func (s *UserAccountServiceProto) CreateNewUser(ctx context.Context, obj *protoobj.UsersMsg) (*protoobj.UsersMsg, error) {
	// logrus.Info("🚀 CreateNewUser")
	paramsObj := marshallerusers.UsersProviderControlDeserialization(obj)

	userUp, errW := s.ipc.Database.UsersActions.CreateUserDB(ctx, paramsObj)
	if errW != nil {
		log.Println("🔴 error CreateNewUser: CreateUserDB: ", errW.Err)
		return nil, errW.Err
	}

	return marshallerusers.UsersProviderControlSerialization(userUp), nil
}

// удаление пользователя
func (s *UserAccountServiceProto) DeleteUser(ctx context.Context, obj *protoobj.UsersMsg) (*protoobj.Empty, error) {
	// logrus.Info("🚀 DeleteUser")
	paramsObj := marshallerusers.UsersProviderControlDeserialization(obj)

	errW := s.ipc.Database.UsersActions.DeleteUserDB(ctx, paramsObj)
	if errW != nil {
		return nil, errW.Err
	}
	return nil, nil
}
