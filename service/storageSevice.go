package service

import (
	"github.com/MiteshSharma/StorageService/service/provider"
	"github.com/MiteshSharma/StorageService/service/data"
	"github.com/Sirupsen/logrus"
	"net/http"
)

type StorageService struct {
	provider provider.Storage
}

func NewStorageService(providerType string) *StorageService {
	factory := &provider.StorageProviderFactory{}
	provider, err := factory.CreateProvider(providerType)
	if (err != nil) {
		logrus.Debug("Error creating provider ", err)
	}
	return &StorageService{provider: provider}
}

func (fs StorageService) GetBuckets() ([]*data.Bucket, error) {
	return fs.provider.GetBuckets()
}

func (fs StorageService)GetBucket(name string) (*data.Bucket, error) {
	return fs.provider.GetBucket(name)
}

func (fs StorageService)CreateBucket(name string) (*data.Bucket, error) {
	return fs.provider.CreateBucket(name)
}

func (fs StorageService)DestroyBucket(name string) (error)  {
	return fs.provider.DestroyBucket(name)
}

func (fs StorageService)GetFiles(containerName string) ([]*data.File, error)  {
	return fs.provider.GetFiles(containerName)
}

func (fs StorageService)GetFile(containerName, name string) (*data.File, error)  {
	return fs.provider.GetFile(containerName, name)
}

func (fs StorageService)RemoveFile(containerName, name string) (error)  {
	return fs.provider.RemoveFile(containerName, name)
}

func (fs StorageService) UploadFile(bucketName string, request *http.Request) ([]*data.File, error) {
	return fs.provider.UploadFile(bucketName, request)
}