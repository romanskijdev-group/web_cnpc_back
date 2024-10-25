package awss3api

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/s3"
)

// удаляет файл из указанного бакета AWS S3.
func (m *AWSS3CloudStorageImpl) DeleteFile(path, fileName string) (string, error) {
	bucketName := m.StorageConfig.Bucket
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	// Проверка существования файла
	path = m.pathSlashChecker(path, fileName)
	headObjectInput := &s3.HeadObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(path),
	}
	headObjectOutput, err := m.ClientStorage.HeadObjectWithContext(ctx, headObjectInput)
	if err != nil {
		if s3Err, ok := err.(awserr.Error); ok && s3Err.Code() == s3.ErrCodeNoSuchKey {
			return "", fmt.Errorf("file does not exist: %s", path)
		}
		return "", fmt.Errorf("failed to check if file exists: %w", err)
	}

	if headObjectOutput.VersionId != nil {
		log.Println("VersionId:", *headObjectOutput.VersionId)
	}
	// Выполняем операцию удаления файла из бакета
	_, err = m.ClientStorage.DeleteObjectWithContext(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(path),
		VersionId: headObjectOutput.VersionId,
	})
	if err != nil {
		return "", err
	}

	// Ждем, пока файл будет полностью удален
	err = m.ClientStorage.WaitUntilObjectNotExists(&s3.HeadObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(path),
	})
	if err != nil {
		return "", err
	}

	
	return path, nil
}

// удаляет папку из указанного бакета AWS S3.
func (m *AWSS3CloudStorageImpl) DeleteFolder(path string) (string, error) {
	// logrus.Info("🟨 DeleteFolder")
	ctxObjects, cancelObjects := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancelObjects()

	bucketName := m.StorageConfig.Bucket

	path = m.pathSlashChecker(path, "")

	// Получаем список всех объектов в папке
	listObjectsInput := &s3.ListObjectsV2Input{
		Bucket: aws.String(bucketName),
		Prefix: aws.String(path),
	}
	resp, err := m.ClientStorage.ListObjectsV2WithContext(ctxObjects, listObjectsInput)
	if err != nil {
		return "", fmt.Errorf("failed to list objects for deletion: %w", err)
	}

	// Удаляем каждый объект в папке
	for _, item := range resp.Contents {
		_, err := m.ClientStorage.DeleteObject(&s3.DeleteObjectInput{
			Bucket: aws.String(bucketName),
			Key:    item.Key,
		})
		if err != nil {
			return "", fmt.Errorf("failed to delete object %s: %w", *item.Key, err)
		}
	}

	// Удаляем саму папку (если она существует как объект)
	_, err = m.ClientStorage.DeleteObjectWithContext(ctxObjects, &s3.DeleteObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(path),
	})
	if err != nil {
		return "", fmt.Errorf("failed to delete folder %s: %w", path, err)
	}

	return path, nil
}
