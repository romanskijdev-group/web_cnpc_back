package awss3api

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

// загружает файл в указанный бакет AWS S3 и делает его доступным только для владельца.
func (m *AWSS3CloudStorageImpl) UploadFilePrivate(reader io.Reader, path, fileName string) (string, error) {
	// logrus.Info("🟨 UploadFilePrivate")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Создаем новую сессию с контекстом
	sess := m.ClientStorage
	// Читаем данные из reader в буфер
	data, err := io.ReadAll(reader)
	if err != nil {
		return "", err
	}
	pathBase := path
	path = m.pathSlashChecker(path, fileName)
	awsPath := aws.String(path)
	isExit, err := m.FileExists(*awsPath)
	if err != nil {
		log.Println("Error checking file exists", err)
	}
	log.Println("isExit", isExit)
	if isExit {
		pthDelete, err := m.DeleteFile(pathBase, fileName)
		log.Println("pthDelete", pthDelete)
		if err != nil {
			log.Println("Error deleting file", err)
			return "", err
		}
	}
	log.Println("path", path)
	log.Println("pathBase", pathBase)

	// // Создаем io.ReadSeeker из прочитанных данных
	readSeeker := bytes.NewReader(data)

	object := &s3.PutObjectInput{
		Bucket: aws.String(m.StorageConfig.Bucket), // Имя бакета
		Key:    aws.String(path),                   // Имя объекта (путь к файлу в бакете)
		Body:   readSeeker,                         // Содержимое файла
		ACL:    aws.String("private"),              // Делаем файл приватным
	}
	// Загружаем файл в S3 с использованием контекста
	_, err = sess.PutObjectWithContext(ctx, object)
	if err != nil {
		fmt.Println(err.Error())
	}

	// Возвращаем полный путь к файлу в S3
	return path, nil
}

// загружает файл в указанный бакет бакет AWS S3 и делает его публично доступным.
func (m *AWSS3CloudStorageImpl) UploadPublicFile(reader io.Reader, path, fileName string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Создаем клиента S3 из сессии
	sess := m.ClientStorage
	// Читаем данные из reader в буфер
	data, err := io.ReadAll(reader)
	if err != nil {
		return "", err
	}
	// Определяем contentType из первых 512 байт данных
	contentType := http.DetectContentType(data[:512])
	// Создаем io.ReadSeeker из прочитанных данных
	readSeeker := bytes.NewReader(data)
	// Полный путь к файлу в бакете

	path = m.pathSlashChecker(path, fileName)

	object := &s3.PutObjectInput{
		Bucket:             aws.String(m.StorageConfig.Bucket), // Имя бакета
		Key:                aws.String(path),                   // Имя объекта (путь к файлу в бакете, включая имя файла)
		Body:               readSeeker,                         // Содержимое файла
		ACL:                aws.String("public-read"),          // Делаем файл доступным для чтения
		ContentType:        aws.String(contentType),            // Устанавливаем MIME тип
		ContentDisposition: aws.String("inline"),               // Указываем, что файл должен отображаться в браузере
	}
	// Загружаем файл в S3 с использованием контекста
	_, err = sess.PutObjectWithContext(ctx, object)
	if err != nil {
		return "", err
	}

	// Формируем публичный URL файла
	publicURL := fmt.Sprintf("%s%s%s%s/%s", "https://", m.StorageConfig.Bucket, ".", m.StorageConfig.Endpoint, path)
	return publicURL, nil
}
