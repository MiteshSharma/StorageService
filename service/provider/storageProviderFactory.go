package provider

import (
	"errors"
	"github.com/MiteshSharma/StorageService/utils"
)

type StorageProviderFactory struct{}

func (f StorageProviderFactory) CreateProvider(provider string) (Storage, error)  {
	switch provider {
	case "File":
		fp := NewFileStorage(utils.ConfigParam.StorageConfig.FileStorage.RootPath)
		return fp, nil
	case "AWS":
		s3p := NewS3Storage()
		return s3p, nil
	default:
		//if type is invalid, return an error
		return nil, errors.New("Invalid provider Type")
	}
}