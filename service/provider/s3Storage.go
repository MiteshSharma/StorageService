package provider

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/MiteshSharma/StorageService/service/data"
	log "github.com/Sirupsen/logrus"
	"net/http"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/MiteshSharma/StorageService/utils"
	"time"
	"fmt"
)

type S3Storage struct {
	Session *session.Session
	S3Connection *s3.S3
}

func NewS3Storage() *S3Storage {
	credentials.NewStaticCredentials(
		utils.ConfigParam.StorageConfig.S3Storage.KeyId,
		utils.ConfigParam.StorageConfig.S3Storage.AccessKey, "")
	awsConfig := &aws.Config{
		Region: aws.String("us-west-1"),
	}
	session := session.New(awsConfig)
	connection := s3.New(session)
	return &S3Storage{Session: session, S3Connection: connection}
}

func (s3S S3Storage) GetBuckets() ([]*data.Bucket, error) {
	var params *s3.ListBucketsInput
	resp, err := s3S.S3Connection.ListBuckets(params)
	if err != nil {
		log.Debug("Error fetching bucket list %v", err)
		return nil, err
	}

	buckets := make([]*data.Bucket, len(resp.Buckets))

	index := 0

	for _, bucket := range resp.Buckets {
		bucket := data.NewBucket(*bucket.Name)
		buckets[index] = bucket
		index++
	}
	return buckets, nil
}

func (s3S S3Storage)GetBucket(name string) (*data.Bucket, error) {
	req := s3.HeadBucketInput{
		Bucket: &name,
	}
	_, err := s3S.S3Connection.HeadBucket(&req)
	if err != nil {
		log.Debug("No bucket exist %v", err)
		return &data.Bucket{}, err
	}

	return data.NewBucket(*req.Bucket), nil
}

func (s3S S3Storage)CreateBucket(name string) (*data.Bucket, error) {
	req := s3.CreateBucketInput{
		Bucket: aws.String(name),
		ACL: aws.String("public-read"),
	}
	_, err := s3S.S3Connection.CreateBucket(&req)
	if err != nil {
		log.Debug("Error on creating bucket %v", err)
		return nil, err
	}
	if err = s3S.S3Connection.WaitUntilBucketExists(&s3.HeadBucketInput{Bucket: &name}); err != nil {
		log.Debug("Failed to wait for bucket to exist %s, %s\n", name, err)
		return nil, err
	}
	return data.NewBucket(name), nil
}

func (s3S S3Storage)DestroyBucket(name string) (error)  {
	params := &s3.DeleteBucketInput{
		Bucket: &name, // Required
	}
	_, err := s3S.S3Connection.DeleteBucket(params)
	if err != nil {
		log.Debug("Error on deleting bucket %v", err)
		return err
	}
	return nil
}

func (s3S S3Storage)GetFiles(bucketName string) ([]*data.File, error)  {
	bucket, err := s3S.GetBucket(bucketName)

	if (&data.Bucket{}) == bucket {
		log.Debug("Bucket name doesn't exist %v", err)
		return nil, err
	}

	params := &s3.ListObjectsInput{
		Bucket:       aws.String(bucketName), // Required
	}
	resp, err := s3S.S3Connection.ListObjects(params)
	if err != nil {
		log.Debug("Error fetching bucket content %v", err)
		return nil, err
	}
	fileList := make([]*data.File, len(resp.Contents))
	index := 0
	for _, content := range resp.Contents {
		file := data.NewFile(*content.Key, *content.Size, time.Now())
		fileList[index] = file
		index++
	}

	return fileList, nil
}

func (s3S S3Storage)GetFile(bucketName, name string) (*data.File, error)  {
	params := &s3.GetObjectInput{
		Bucket:                     aws.String(bucketName), // Required
		Key:                        aws.String(name), // Required
	}
	resp, err := s3S.S3Connection.GetObject(params)
	if err != nil {
		log.Debug("Error fetching bucket file %v", err)
		return &data.File{}, err
	}
	return data.NewFile(name, *resp.ContentLength, *resp.LastModified), nil
}

func (s3S S3Storage)RemoveFile(bucketName, name string) (error)  {
	params := &s3.DeleteObjectInput{
		Bucket:       aws.String(bucketName), // Required
		Key:          aws.String(name),  // Required
	}
	_, err := s3S.S3Connection.DeleteObject(params)
	if err != nil {
		log.Debug("Error deleting bucket content %v", err)
		return err
	}
	return nil
}

func (s3S S3Storage) UploadFile(bucketName string, request *http.Request) ([]*data.File, error)  {
	uploader := s3manager.NewUploader(s3S.Session)
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
		result, err := uploader.Upload(&s3manager.UploadInput{
			Bucket: aws.String(bucketName),
			Key:    aws.String(fileName),
			ACL: 	aws.String("public-read"),
			Body:   file,
		})
		if err != nil {
			log.Debug("Failed to upload %v", err)
		} else {
			fmt.Println("Upload successful with uploadId "+result.UploadID)
		}
		fileObj, _ := s3S.GetFile(bucketName, fileName)
		fileList[index] = fileObj
		index++
	}

	return fileList, nil
}