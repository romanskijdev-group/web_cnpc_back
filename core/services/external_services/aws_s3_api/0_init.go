package awss3api

import (
	"io"
	"log"
	"net/http"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	ClientStorage     *s3.S3
	onceClientStorage sync.Once
)

type AWSS3CloudStorageI interface { // удаляет файл из указанное хранилище AWS S3.
	DeleteFile(path, fileName string) (string, error)
	// удаляет папку из указанное хранилище AWS S3.
	DeleteFolder(path string) (string, error)
	//  загружает файл в указанное хранилище AWS S3 и делает его доступным только для владельца.
	UploadFilePrivate(reader io.Reader, path, fileName string) (string, error)
	//  потоково отправляет файл из указанное хранилище AWS S3.
	StreamFile(w http.ResponseWriter, path, fileName string) (http.ResponseWriter, error)
	//  загружает файл в указанное хранилище AWS S3 и делает его публично доступным.
	UploadPublicFile(reader io.Reader, path, fileName string) (string, error)
	// Получает файл из S3 и возвращает его как io.ReadCloser.
	GetFileFromS3(path, fileName string) (io.ReadCloser, error)
}

type StorageConfigSt struct {
	Key      string `yaml:"key" env-required:"true"`
	Secret   string `yaml:"secret" env-required:"true"`
	Region   string `yaml:"region" env-required:"true"`
	Endpoint string `yaml:"endpoint" env-required:"true"`
	Bucket   string `yaml:"bucket" env-required:"true"`
	RunMode  string `yaml:"run_mode" env-required:"true"`
}

type AWSS3CloudStorageImpl struct {
	StorageConfig StorageConfigSt
	ClientStorage *s3.S3
}

func NewAWSS3CloudStorage(storageConfig StorageConfigSt) AWSS3CloudStorageI {
	// logrus.Info("🟨 NewAWSS3CloudStorage")
	return &AWSS3CloudStorageImpl{
		StorageConfig: storageConfig,
		ClientStorage: GetClient(storageConfig),
	}
}

func GetClient(storageConfig StorageConfigSt) *s3.S3 {
	onceClientStorage.Do(func() {
		ClientStorage = createClient(storageConfig)
	})
	return ClientStorage
}

func createClient(storageConfig StorageConfigSt) *s3.S3 {
	// logrus.Info("🟨 createClient")
	s3Config := &aws.Config{
		Credentials:      credentials.NewStaticCredentials(storageConfig.Key, storageConfig.Secret, ""),
		Endpoint:         aws.String(storageConfig.Endpoint),
		S3ForcePathStyle: aws.Bool(false),
		Region:           aws.String(storageConfig.Region),
	}
	newSession, err := session.NewSession(s3Config)
	if err != nil {
		log.Fatalf("Failed to create session: %v", err)
		return nil
	}
	s3Client := s3.New(newSession)
	return s3Client
}
