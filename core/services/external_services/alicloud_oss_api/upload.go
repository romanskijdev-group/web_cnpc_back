package alicloudossapi

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

// загружает файл в указанный бакет AliCloud OSS и делает его доступным только для владельца.
func (m *AliCloudOSSStorageImpl) UploadFilePrivate(reader io.Reader, path, fileName string) (string, error) {
	// logrus.Info("🟨 UploadFilePrivate")
	// Читаем данные из reader в буфер
	data, err := io.ReadAll(reader)
	if err != nil {
		return "", err
	}

	// Создаем io.Reader из прочитанных данных
	readSeeker := bytes.NewReader(data)

	if !strings.HasSuffix(path, "/") {
		path += "/"
	}

	fullPath := path + fileName

	// Загружаем файл в OSS
	err = m.ClientStorage.PutObject(fullPath, readSeeker, oss.ObjectACL(oss.ACLPrivate))
	if err != nil {
		return "", err
	}

	// Возвращаем полный путь к файлу в OSS
	return fullPath, nil
}

// потоково отправляет файл из указанного бакета AliCloud OSS.
func (m *AliCloudOSSStorageImpl) StreamFile(w http.ResponseWriter, objectName string) (http.ResponseWriter, error) {
	// logrus.Info("🟨 StreamFile")
	// Получаем объект из бакета
	body, err := m.ClientStorage.GetObject(objectName)
	if err != nil {
		return nil, err
	}
	defer body.Close() // Убедитесь, что тело ответа закрыто

	// Устанавливаем заголовок ответа
	w.Header().Set("Content-Type", "application/octet-stream")

	// Копируем содержимое файла в http.ResponseWriter
	if _, err := io.Copy(w, body); err != nil {
		return nil, fmt.Errorf("failed to stream file: %v", err)
	}

	return w, nil
}

func (m *AliCloudOSSStorageImpl) GetFileContentAsBytes(objectName string) ([]byte, error) {
	// logrus.Info("🟨 GetFileContentAsBytes")
	body, err := m.ClientStorage.GetObject(objectName)
	if err != nil {
		return nil, err
	}
	defer body.Close()

	bytes, err := io.ReadAll(body)
	if err != nil {
		return nil, fmt.Errorf("failed to read file content: %v", err)
	}

	return bytes, nil
}

// загружает файл в указанный бакет AliCloud OSS и делает его публично доступным.
func (m *AliCloudOSSStorageImpl) UploadPublicFile(reader io.Reader, path, fileName string) (string, error) {
	// logrus.Info("🟨 UploadPublicFile")
	// Читаем данные из reader в буфер
	data, err := io.ReadAll(reader)
	if err != nil {
		return "", err
	}

	// Определяем contentType из первых 512 байт данных
	contentType := http.DetectContentType(data[:512])

	// Создаем io.Reader из прочитанных данных для повторного использования
	readSeeker := bytes.NewReader(data)
	// Полный путь к файлу в бакете
	fullPath := path + fileName

	// Загружаем файл в OSS с автоматически определенным contentType
	err = m.ClientStorage.PutObject(fullPath, readSeeker, oss.ObjectACL(oss.ACLPublicRead), oss.ContentType(contentType), oss.ContentDisposition("inline"))
	if err != nil {
		return "", err
	}

	// Формируем публичный URL файла
	publicURL := fmt.Sprintf("https://%s.%s/%s", m.StorageConfig.Bucket, m.StorageConfig.Endpoint, fullPath)
	return publicURL, nil
}
