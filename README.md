# StorageService
StorageService provide flexible way to manage files on server. We are using multipart upload which handles file upload in limited memory and has capability to upload large files.

# Storage service has two main components
1. Buckets: Buckets are group of files which are placed together. This is like a directory in file system where we can place files.
2. Files: These are any kind of files which can be stored.

REST calls available :

GET    /bucket               : Get all bucket lists

GET    /bucket/:bucketName   : Check if bucket bucketName exists and fetch it

POST   /bucket/:bucketName   : Create new bucket with bucketName

DELETE /bucket/:bucketName   : Delete bucket with name bucketName

GET    /bucket/:bucketName/file : Get all files belong to bucket bucketName

GET    /bucket/:bucketName/file/:fileName : Get file details with file name as fileName in bucket name as bucketName

POST   /bucket/:bucketName/file : Upload file in bucket name as bucketName

DELETE /bucket/:bucketName/file/:fileName : Delete file with name fileName in bucket name as bucketName
