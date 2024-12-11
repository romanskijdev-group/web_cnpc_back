package userprofile

import (
	marshallerusers "cnpc_backend/core/module/user/users/marshaller"
	"cnpc_backend/core/typescore"
	"cnpc_backend/rest_user_service/types"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"path/filepath"
	"strings"
	"time"
)

type HandlerUserProfile struct {
	ipc *types.InternalProviderControl
}

const (
	userControlBaseURI     = "/api/control/user"
	userAvatarURLUpdateURI = "/api/profile/user/avatar"
)

func (h *HandlerUserProfile) UpdateUserProfileHandler(w http.ResponseWriter, r *http.Request, userObj *typescore.UsersProviderControl, detectorIP *typescore.DetectorIPStruct) (interface{}, *uint64, *typescore.WEvent) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	defer r.Body.Close()
	// Проверка на пустое тело запроса
	if r.ContentLength == 0 {
		errW := &typescore.WEvent{Err: errors.New("empty request body"), Text: "invalid_request_body"}
		return nil, nil, errW
	}
	userObjUpdates := &types.UpdateUserProfileHandlerReq{}

	err := json.NewDecoder(r.Body).Decode(userObjUpdates)
	if err != nil {
		errW := &typescore.WEvent{Err: err, Text: "system_error"}
		return nil, nil, errW
	}

	objPr := marshallerusers.UsersProviderControlSerialization(&typescore.UsersProviderControl{
		SystemID:  userObj.SystemID,
		SerialID:  userObj.SerialID,
		Nickname:  userObjUpdates.Nickname, // Псевдоним или никнейм пользователя
		FirstName: userObjUpdates.FirstName,
		LastName:  userObjUpdates.LastName,
		BirthDate: userObjUpdates.BirthDate,
		Language:  userObjUpdates.Language,
	})

	objPrRes, err := h.ipc.Clients.UserAccountServiceProto.UpdateUserProfile(ctx, objPr)
	if err != nil {
		errW := &typescore.WEvent{Err: err, Text: "system_error"}
		return nil, nil, errW
	}

	result := marshallerusers.UsersProviderControlDeserialization(objPrRes)

	return result, nil, nil
}

// удаление пользователя
func (h *HandlerUserProfile) DeleteUserHandler(w http.ResponseWriter, r *http.Request, userObj *typescore.UsersProviderControl, detectorIP *typescore.DetectorIPStruct) (interface{}, *uint64, *typescore.WEvent) {
	defer r.Body.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if userObj.SystemID == nil {
		errW := &typescore.WEvent{Err: errors.New("user_is_not_found"), Text: "invalid_request_body"}
		return nil, nil, errW
	}
	userObjPr := marshallerusers.UsersProviderControlSerialization(userObj)

	_, err := h.ipc.Clients.UserAccountServiceProto.DeleteUser(ctx, userObjPr)
	if err != nil {
		errW := &typescore.WEvent{Err: err, Text: "system_error"}
		return nil, nil, errW
	}
	return nil, nil, nil
}

// получение информации о пользователе
func (h *HandlerUserProfile) GetUserProfile(w http.ResponseWriter, r *http.Request, userObj *typescore.UsersProviderControl, detectorIP *typescore.DetectorIPStruct) (interface{}, *uint64, *typescore.WEvent) {
	defer r.Body.Close()
	//  все данные мы уже получаем на стадии проверки токена при запросе к этому методу и передаем в userObj
	return userObj, nil, nil
}

// Установка аватара пользователя
// @Summary Установка аватара пользователя
// @Description Загружает аватар пользователя
// @Tags user profile
// @Accept  multipart/form-data
// @Produce  json
// @Param   file  formData  file  true  "Файл аватара"
// @Success 200 {object} AvatarURLSetter
// @Failure 400 {object} typescore.WEvent
// @Failure 500 {object} typescore.WEvent
// @Router /api/profile/user/avatar [post]
func (h *HandlerUserProfile) SetUserAvatarHandler(w http.ResponseWriter, r *http.Request, userObj *typescore.UsersProviderControl, detectorIP *typescore.DetectorIPStruct) (interface{}, *uint64, *typescore.WEvent) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	defer r.Body.Close()
	if userObj.SystemID == nil {
		errW := &typescore.WEvent{Err: errors.New("user_is_not_found"), Text: "invalid_request_body"}

		return nil, nil, errW
	}

	// Ограничение размера файла
	const maxFileSize = 5 << 20 // 5 MB
	if r.ContentLength > maxFileSize {
		errW := &typescore.WEvent{Err: errors.New("file_too_large"), Text: "invalid_request_body"}
		return nil, nil, errW
	}

	// Чтение файла из запроса
	file, header, err := r.FormFile("file")
	if err != nil {
		return nil, nil, &typescore.WEvent{Err: err, Text: "invalid_request_body"}
	}
	defer file.Close()

	// Получаем расширение файла из оригинального имени файла
	ext := strings.ToLower(filepath.Ext(header.Filename))
	// Список разрешенных форматов файлов
	allowedExtensions := map[string]bool{".jpg": true, ".jpeg": true, ".png": true, ".heif": true}

	if _, allowed := allowedExtensions[ext]; !allowed {
		errW := &typescore.WEvent{Err: errors.New("invalid_file_format"), Text: "invalid_request_body"}
		return nil, nil, errW
	}

	// Генерируем новое имя файла, используя текущее время и добавляем расширение файла
	fileName := time.Now().UTC().Format("20060102_150405") + ext
	path := fmt.Sprintf("users/%d/avatar/", *userObj.SerialID)

	publicURL, err := h.ipc.Storage.UploadPublicFile(file, path, fileName)
	if err != nil {
		fmt.Errorf(err.Error())
		return nil, nil, &typescore.WEvent{Err: err, Text: "invalid_request_body"}
	}

	objPrS := marshallerusers.UpdateUserAvatarURLReqSerialization(userObj.SystemID, &publicURL)

	_, err = h.ipc.Clients.UserAccountServiceProto.UpdateUserAvatarURL(ctx, objPrS)
	if err != nil {
		return nil, nil, &typescore.WEvent{Err: err, Text: "system_error"}
	}
	// публичный URL загруженного файла.
	return &types.AvatarURLSetter{
		URL: publicURL,
	}, nil, nil
}

// Удаление текущего аватара пользователя
// @Summary Удаление текущего аватара пользователя
// @Description Удаление текущего аватара пользователя
// @Tags user profile
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]bool
// @Failure 400 {object} typescore.WEvent
// @Failure 500 {object} typescore.WEvent
// @Router /api/profile/user/avatar [delete]
func (h *HandlerUserProfile) DeleteUserAvatarHandler(w http.ResponseWriter, r *http.Request, userObj *typescore.UsersProviderControl, detectorIP *typescore.DetectorIPStruct) (interface{}, *uint64, *typescore.WEvent) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	defer r.Body.Close()
	if userObj.SystemID == nil {
		errW := &typescore.WEvent{Err: errors.New("user_is_not_found"), Text: "invalid_request_body"}
		return nil, nil, errW
	}

	objPrS := marshallerusers.UpdateUserAvatarURLReqSerialization(userObj.SystemID, nil)

	_, err := h.ipc.Clients.UserAccountServiceProto.UpdateUserAvatarURL(ctx, objPrS)
	if err != nil {
		errW := &typescore.WEvent{Err: err, Text: "system_error"}
		return nil, nil, errW
	}

	path := fmt.Sprintf("users/%d/avatar/", *userObj.SerialID)
	_, err = h.ipc.Storage.DeleteFolder(path)
	if err != nil {
		return nil, nil, &typescore.WEvent{Err: err, Text: "system_error"}
	}

	return nil, nil, nil
}
