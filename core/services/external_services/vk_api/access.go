package vkapi

import (
	"encoding/json"
	"net/http"
	"net/url"
)

// GetAccessToken получение access_token по коду авторизации
func (client *VKClient) GetAccessToken(code string) (string, error) {
	values := url.Values{}
	values.Set("client_id", client.config.clientID)
	values.Set("client_secret", client.config.clientSecret)
	values.Set("redirect_uri", client.config.redirectURI)
	values.Set("code", code)

	response, err := http.PostForm("https://oauth.vk.com/access_token", values)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	var vkResponse struct {
		AccessToken string `json:"access_token"`
		UserID      int    `json:"user_id"`
	}

	if err := json.NewDecoder(response.Body).Decode(&vkResponse); err != nil {
		return "", err
	}

	return vkResponse.AccessToken, nil
}
