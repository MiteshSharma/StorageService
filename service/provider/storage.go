package provider

import (
	"github.com/MiteshSharma/StorageService/service/data"
	"net/http"
)

type Storage interface {
	GetBuckets() ([]*data.Bucket, error)
	GetBucket(name string) (*data.Bucket, error)
	CreateBucket(name string) (*data.Bucket, error)
	DestroyBucket(string string) (error)
	GetFiles(bucketName string) ([]*data.File, error)
	GetFile(bucketName, name string) (*data.File, error)
	RemoveFile(bucketName, name string) (error)
	UploadFile(bucketName string, request *http.Request, isStreaming bool) (*data.File, error)
}