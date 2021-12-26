package common_services

type FileTransferService interface {
	Upload(dto FileDto) FileTransferResponse
	Remove(fileName string, bucketName string) error
	Download(fileName string, bucketName string) (error, []byte)
	CreateBucket(name string) error
}
