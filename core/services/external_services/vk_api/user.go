package vkapi

import "github.com/SevereCloud/vksdk/v2/api"

// GetUserInfo получение информации о пользователе по userID
func (client *VKClient) GetUserInfo(userID int) (*api.UsersGetResponse, error) {
	users, err := client.vk.UsersGet(api.Params{
		"user_ids": userID,
		"fields":   "first_name,last_name",
	})
	if err != nil {
		return nil, err
	}
	return &users, nil
}
