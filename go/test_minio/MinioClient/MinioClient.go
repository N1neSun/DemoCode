package MinioClient

import (
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/minio/minio-go/v7"
)

type MinioClient struct {
	client *minio.Client
}

const (
	location          string = "location" //Region,defalut is location
	contentTypeFile   string = "application/zip"
	contentTypeStream string = "application/octet-stream"
)

var ctx = context.Background()

func (c *MinioClient) Createbucket(bucketName string) error {
	err := c.client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		exists, errBucketExists := c.client.BucketExists(context.Background(), bucketName)
		if errBucketExists == nil && exists {
			return errors.New("bucketname exists")
		} else {
			return errors.New(fmt.Sprintf("Createbucket error, %v", err))
		}
	} else {
		return nil
	}
}

func (c *MinioClient) RemoveBucket(bucketName string) error {
	err := c.client.RemoveBucket(ctx, bucketName)
	if err != nil {
		return errors.New(fmt.Sprintf("Remove bucket error, %v", err))
	}
	return nil
}

func (c *MinioClient) UploadFile(bucketName string, objectName string, filePath string) (string, int64, error) {
	info, err := c.client.FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentTypeFile})
	if err != nil {
		return "", 0, errors.New(fmt.Sprintf("UploadFile failed, %v", err))
	}
	return objectName, info.Size, nil
}

func (c *MinioClient) DownloadFile(bucketName string, objectName string, filePath string) error {
	err := c.client.FGetObject(ctx, bucketName, objectName, filePath, minio.GetObjectOptions{})
	if err != nil {
		return errors.New(fmt.Sprintf("DownloadFile failed, %v", err))
	}
	return nil
}

func (c *MinioClient) UploadFileStream(bucketName string, objectName string, fileStream io.Reader, fileSize int64) error {
	_, err := c.client.PutObject(ctx, bucketName, objectName, fileStream, fileSize, minio.PutObjectOptions{ContentType: contentTypeStream})
	if err != nil {
		return errors.New(fmt.Sprintf("UploadFileStream failed, %v", err))
	}
	return nil
}

func (c *MinioClient) DownloadFileStream(bucketName string, objectName string) (io.Reader, error) {
	object, err := c.client.GetObject(ctx, bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		return nil, errors.New(fmt.Sprintf("DownloadFileStream failed, %v", err))
	}
	return object, nil
}

func (c *MinioClient) RemoveFile(bucketName string, objectName string) error {
	objInfo, err := c.client.StatObject(ctx, bucketName, objectName, minio.StatObjectOptions{})
	if err != nil {
		return errors.New(fmt.Sprintf("RemoveFile failed, %v", err))
	}
	opts := minio.RemoveObjectOptions{
		GovernanceBypass: true,
		VersionID:        objInfo.VersionID,
	}
	err = c.client.RemoveObject(ctx, bucketName, objectName, opts)
	if err != nil {
		return errors.New(fmt.Sprintf("RemoveFile failed, %v", err))
	}
	return nil
}
