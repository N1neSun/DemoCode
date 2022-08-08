package MinioClient

import (
	"context"
	"errors"

	"github.com/minio/minio-go/v7"
)

type MinioClient struct {
	client *minio.Client
}

const (
	location    string = "location" //Region,defalut is location
	contentType string = "application/zip"
)

var ctx = context.Background()

func (c *MinioClient) Createbucket(bucketName string) error {
	err := c.client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		exists, errBucketExists := c.client.BucketExists(context.Background(), bucketName)
		if errBucketExists == nil && exists {
			return errors.New("bucketname exists")
		} else {
			return errors.New("Createbucket error")
		}
	} else {
		return nil
	}
}

func (c *MinioClient) UpdateFile(bucketName string, objectName string, filePath string) (string, int64, error) {
	info, err := c.client.FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return "", 0, errors.New("UpdateFile failed")
	}
	return objectName, info.Size, nil
}
