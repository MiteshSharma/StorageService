package provider

import (
	"os"
	"errors"
	log "github.com/Sirupsen/logrus"
	"github.com/MiteshSharma/StorageService/service/data"
	"path"
	"io/ioutil"
	"io"
	"net/http"
	"fmt"
)

type FileStorage struct {
	rootPath string
}

func NewFileStorage(rootPath string) *FileStorage {
	return &FileStorage{rootPath: rootPath}
}

func (fs FileStorage) GetBuckets() ([]*data.Bucket, error) {
	if (!fs.isDirectory(fs.rootPath)) {
		return nil, errors.New("Root path is not a directory")
	}
	files, _ := ioutil.ReadDir(fs.rootPath)

	buckets := make([]*data.Bucket, len(files))

	index := 0
	for _, f := range files {
		bucket := data.NewBucket(f.Name())
		buckets[index] = bucket
		index++
	}
	return buckets, nil
}

func (fs FileStorage)GetBucket(name string) (*data.Bucket, error) {
	if (!fs.isDirectory(fs.rootPath)) {
		return &data.Bucket{}, errors.New("Root path is not a directory")
	}
	f, err := os.Open(path.Join(fs.rootPath, name))
	if (err != nil) {
		log.Debug("Error on opening file path %v", err)
		return &data.Bucket{}, err
	}
	fInfo, err := f.Stat()
	if (err != nil) {
		log.Debug("Error on reading file stat %v", err)
		return &data.Bucket{}, err
	}
	return data.NewBucket(fInfo.Name()), nil
}

func (fs FileStorage)CreateBucket(name string) (*data.Bucket, error) {
	err := os.MkdirAll(path.Join(fs.rootPath, name), 0711)
	if err != nil {
		log.Debug("Error creating directory %v", err)
		return &data.Bucket{}, err
	}
	return fs.GetBucket(name)
}

func (fs FileStorage)DestroyBucket(name string) (error)  {
	err := os.RemoveAll(path.Join(fs.rootPath, name))
	if err != nil {
		log.Debug("Error deleting directory %v", err)
		return err
	}
	return nil
}

func (fs FileStorage)GetFiles(bucketName string) ([]*data.File, error)  {
	files, _ := ioutil.ReadDir(path.Join(fs.rootPath, bucketName))

	fileList := make([]*data.File, len(files))

	index := 0
	for _, f := range files {
		file := data.NewFile(f.Name(), f.Size(), f.ModTime())
		fileList[index] = file
		index++
	}
	return fileList, nil
}

func (fs FileStorage)GetFile(bucketName, name string) (*data.File, error)  {
	f, err := os.Open(path.Join(fs.rootPath, bucketName, "/", name))
	if (err != nil) {
		log.Debug("Error on opening file path %v", err)
		return &data.File{}, err
	}
	fInfo, err := f.Stat()
	if (err != nil) {
		log.Debug("Error on reading file stat %v", err)
		return &data.File{}, err
	}
	return data.NewFile(fInfo.Name(), fInfo.Size(), fInfo.ModTime()), nil
}

func (fs FileStorage)RemoveFile(bucketName, name string) (error)  {
	err := os.Remove(path.Join(fs.rootPath, bucketName, "/", name))
	if err != nil {
		log.Debug("Error deleting directory %v", err)
		return err
	}
	return nil
}

func (fs FileStorage) UploadFile(bucketName string, request *http.Request) ([]*data.File, error)  {
	var fileName string
	err := request.ParseMultipartForm(100000)
	if err != nil {
		log.Debug("Error parsing multipart form %v", err)
		return nil, err
	}

	//get a ref to the parsed multipart form
	m := request.MultipartForm
	fmt.Println("Reading form data")

	//get the *fileheaders
	files := m.File["file"]
	fmt.Println("File size %ld", len(files))
	fileList := make([]*data.File, len(files))
	index := 0
	for i, _ := range files {
		fileName = files[i].Filename
		//for each fileheader, get a handle to the actual file
		file, err := files[i].Open()
		defer file.Close()
		if err != nil {
			log.Debug("Error opening file received %v", err)
			return nil, err
		}
		//create destination file making sure the path is writeable.
		var filePath = path.Join(fs.rootPath, bucketName, "/", files[i].Filename)
		fmt.Println("Saving file at path %s", filePath)
		dst, err := os.Create(filePath)
		defer dst.Close()
		if err != nil {
			log.Debug("Error creating destination file %v", err)
			return nil, err
		}
		//copy the uploaded file to the destination file
		if _, err := io.Copy(dst, file); err != nil {
			log.Debug("Error copying source file to destination file %v", err)
			return nil, err
		}
		fileObj, _ := fs.GetFile(bucketName, fileName)
		fileList[index] = fileObj
		index++
	}
	return fileList, nil
}

func (fs FileStorage) isDirectory(filePath string) bool  {
	fi, err := os.Stat(filePath)
	if err != nil {
		log.Debug("Error on reading file path %v", err)
		return false
	}
	switch mode := fi.Mode(); {
	case mode.IsDir():
		return true
	case mode.IsRegular():
		log.Debug("Given path %s is a file.", filePath)
		return false
	}
	return false
}