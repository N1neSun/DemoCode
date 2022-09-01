package MinioClient

import (
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinioClient struct {
	Client *minio.Client
}

const (
	location          string = "location" //Region,defalut is location
	contentTypeFile   string = "application/zip"
	contentTypeStream string = "application/octet-stream"
)

//var ctx = context.Background()

func (c *MinioClient) Connect() error {
	endpoint := ""
	accessKeyID := ""
	secretAccessKey := ""
	useSSL := false

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		return errors.New("connect Minio error")
	}
	c.Client = minioClient
	return nil
}

func (c *MinioClient) Createbucket(ctx context.Context, bucketName string) error {
	err := c.Client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		exists, errBucketExists := c.Client.BucketExists(context.Background(), bucketName)
		if errBucketExists == nil && exists {
			return errors.New("bucketname exists")
		} else {
			return fmt.Errorf("createbucket error, %v", err)
		}
	} else {
		return nil
	}
}

func (c *MinioClient) RemoveBucket(ctx context.Context, bucketName string) error {
	err := c.Client.RemoveBucket(ctx, bucketName)
	if err != nil {
		return fmt.Errorf("remove bucket error, %v", err)
	}
	return nil
}

func (c *MinioClient) UploadFile(ctx context.Context, bucketName string, objectName string, filePath string) (string, int64, error) {
	info, err := c.Client.FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentTypeFile})
	if err != nil {
		return "", 0, fmt.Errorf("uploadFile failed, %v", err)
	}
	return objectName, info.Size, nil
}

func (c *MinioClient) DownloadFile(ctx context.Context, bucketName string, objectName string, filePath string) error {
	err := c.Client.FGetObject(ctx, bucketName, objectName, filePath, minio.GetObjectOptions{})
	if err != nil {
		return fmt.Errorf("downloadFile failed, %v", err)
	}
	return nil
}

func (c *MinioClient) UploadFileStream(ctx context.Context, bucketName string, objectName string, fileStream multipart.File, fileSize int64) error {
	_, err := c.Client.PutObject(ctx, bucketName, objectName, fileStream, fileSize, minio.PutObjectOptions{ContentType: contentTypeStream})
	if err != nil {
		return fmt.Errorf("uploadFileStream failed, %v", err)
	}
	return nil
}

func (c *MinioClient) DownloadFileStream(ctx context.Context, bucketName string, objectName string) (io.Reader, error) {
	object, err := c.Client.GetObject(ctx, bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		return nil, fmt.Errorf("downloadFileStream failed, %v", err)
	}
	return object, nil
}

func (c *MinioClient) RemoveFile(ctx context.Context, bucketName string, objectName string) error {
	objInfo, err := c.Client.StatObject(ctx, bucketName, objectName, minio.StatObjectOptions{})
	if err != nil {
		return fmt.Errorf("removeFile failed, %v", err)
	}
	opts := minio.RemoveObjectOptions{
		GovernanceBypass: true,
		VersionID:        objInfo.VersionID,
	}
	err = c.Client.RemoveObject(ctx, bucketName, objectName, opts)
	if err != nil {
		return fmt.Errorf("removeFile failed, %v", err)
	}
	return nil
}
