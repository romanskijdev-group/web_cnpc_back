package awss3api

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/s3"
)

// потоково отправляет файл из указанного бакета бакет AWS S3.
func (m *AWSS3CloudStorageImpl) StreamFile(w http.ResponseWriter, path, fileName string) (http.ResponseWriter, error) {
	readerFIle, err := m.GetFileFromS3(path, fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to get file from S3: %v", err)
	}
	defer readerFIle.Close() // Ensure the response body is closed

	// Set the response header
	w.Header().Set("Content-Type", "application/octet-stream")

	// Копируйте содержимое файла в http.ResponseWriter.
	if _, err := io.Copy(w, readerFIle); err != nil {
		return nil, fmt.Errorf("failed to stream file: %v", err)
	}

	return w, nil
}

// Проверяет, существует ли файл по указанному пути в бакете S3
func (m *AWSS3CloudStorageImpl) FileExists(path string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	sess := m.ClientStorage

	headObjectInput := &s3.HeadObjectInput{
		Bucket: aws.String(m.StorageConfig.Bucket),
		Key:    aws.String(path),
	}

	_, err := sess.HeadObjectWithContext(ctx, headObjectInput)
	if err != nil {
		if s3Err, ok := err.(awserr.Error); ok && s3Err.Code() == s3.ErrCodeNoSuchKey {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

// Получает файл из S3 и возвращает его как io.ReadCloser.
func (m *AWSS3CloudStorageImpl) GetFileFromS3(path, fileName string) (io.ReadCloser, error) {
	// Assuming m.ClientStorage is an initialized AWS S3 client.
	svc := m.ClientStorage
	if fileName != ""{
		path = m.pathSlashChecker(path, fileName)
	}
	// Create the request with the context
	req, resp := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(m.StorageConfig.Bucket),
		Key:    aws.String(path),
	})

	// Send the request
	err := req.Send()
	if err != nil {
		return nil, err // Return the error if the request failed
	}

	return resp.Body, nil
}
