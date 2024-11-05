package vkapi

import (
	"github.com/SevereCloud/vksdk/v2/api"
)

// VKClient структура для работы с VK API
type VKClient struct {
	vk     *api.VK
	config VKApiConfig
}

// NewVKClient инициализация нового клиента VK
func NewVKClient(accessToken string, apiConfig VKApiConfig) *VKClient {
	vk := api.NewVK(accessToken)
	return &VKClient{
		vk:     vk,
		config: apiConfig,
	}
}
