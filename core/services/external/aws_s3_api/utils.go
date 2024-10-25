package awss3api

import (
	"fmt"
	"strings"
)

func (m *AWSS3CloudStorageImpl) pathSlashChecker(path, fileName string) string {
	if !strings.HasSuffix(path, "/") {
		path += "/"
	}
	// Проверка начала строки
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	if !strings.HasPrefix(path, fmt.Sprintf("/%s/", m.StorageConfig.RunMode)) {
		path = fmt.Sprintf("/%s%s", m.StorageConfig.RunMode, path)
	}

	if fileName != "" {
		path = fmt.Sprintf("%s%s", path, fileName)
	}

	// Замена всех двойных слэшей на одинарные
	path = strings.ReplaceAll(path, "//", "/")

	return path
}
