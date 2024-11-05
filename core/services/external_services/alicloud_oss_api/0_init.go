package alicloudossapi

import (
	"io"
	"log"
	"net/http"
	"sync"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var (
	ClientStorage     *oss.Bucket
	onceClientStorage sync.Once
)

type AliCloudOSSStorageI interface { // удаляет файл из указанное хранилище AWS S3.
	DeleteFile(path, fileName string) (string, error)
	// удаляет папку из указанное хранилище AWS S3.
	DeleteFolder(path string) (string, error)
	//  загружает файл в указанное хранилище AWS S3 и делает его доступным только для владельца.
	UploadFilePrivate(reader io.Reader, path, fileName string) (string, error)
	//  потоково отправляет файл из указанное хранилище AWS S3.
	StreamFile(w http.ResponseWriter, objectName string) (http.ResponseWriter, error)
	//  загружает файл в указанное хранилище AWS S3 и делает его публично доступным.
	UploadPublicFile(reader io.Reader, path, fileName string) (string, error)

	// возвращает содержимое файла в виде массива байтов.
	GetFileContentAsBytes(objectName string) ([]byte, error)
}

type StorageConfigSt struct {
	Key      string `yaml:"key" env-required:"true"`
	Secret   string `yaml:"secret" env-required:"true"`
	Endpoint string `yaml:"endpoint" env-required:"true"`
	Bucket   string `yaml:"bucket" env-required:"true"`
}

type AliCloudOSSStorageImpl struct {
	StorageConfig StorageConfigSt
	ClientStorage *oss.Bucket
}

func NewAliCloudOSSStorage(storageConfig StorageConfigSt) AliCloudOSSStorageI {
	// logrus.Info("🟨 NewAliCloudOSSStorage")
	return &AliCloudOSSStorageImpl{
		StorageConfig: storageConfig,
		ClientStorage: GetClient(storageConfig),
	}
}

func GetClient(storageConfig StorageConfigSt) *oss.Bucket {
	onceClientStorage.Do(func() {
		ClientStorage = createClient(storageConfig)
	})
	return ClientStorage
}

func createClient(storageConfig StorageConfigSt) *oss.Bucket {
	// logrus.Info("🟨 createClient")
	client, err := oss.New(storageConfig.Endpoint,
		storageConfig.Key, storageConfig.Secret)
	if err != nil {
		log.Fatalf("failed to create session AliCloudOSS: %v", err)
		return nil
	}

	bucket, err := client.Bucket(storageConfig.Bucket)
	if err != nil {
		log.Fatalf("error connect bucket AliCloudOSS: %s %v", storageConfig.Bucket, err)
	}
	return bucket
}
