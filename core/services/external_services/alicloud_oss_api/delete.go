package alicloudossapi

import (
	"fmt"
	"strings"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

// удаляет файл из указанного бакета AliCloud OSS.
func (m *AliCloudOSSStorageImpl) DeleteFile(path, fileName string) (string, error) {
	// logrus.Info("🟨 DeleteFile")

	if !strings.HasSuffix(path, "/") {
		path += "/"
	}

	// Полный путь к файлу в бакете
	fullPath := path + fileName

	// Выполняем операцию удаления файла из бакета
	err := m.ClientStorage.DeleteObject(fullPath)
	if err != nil {
		return "", err
	}

	return fullPath, nil
}

// удаляет папку из указанного бакета AliCloud OSS.
func (m *AliCloudOSSStorageImpl) DeleteFolder(path string) (string, error) {
	// logrus.Info("🟨 DeleteFolder")
	// Убедитесь, что путь к папке заканчивается на "/"
	if !strings.HasSuffix(path, "/") {
		path += "/"
	}

	// Получаем список всех объектов в папке
	marker := oss.Marker("")
	for {
		lor, err := m.ClientStorage.ListObjects(oss.Prefix(path), marker)
		if err != nil {
			return "", fmt.Errorf("failed to list objects for deletion: %w", err)
		}

		// Удаляем объекты
		for _, object := range lor.Objects {
			err := m.ClientStorage.DeleteObject(object.Key)
			if err != nil {
				return "", fmt.Errorf("failed to delete object %s: %w", object.Key, err)
			}
		}

		if !lor.IsTruncated {
			break
		}
		marker = oss.Marker(lor.NextMarker)
	}

	return path, nil
}
