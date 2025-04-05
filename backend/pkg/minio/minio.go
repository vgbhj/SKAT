package minio

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/vgbhj/SKAT/pkg/setting"
)

var minioClient *minio.Client
var ctx context.Context

func Init() {
	ctx = context.Background()
	var err error
	minioClient, err = minio.New(setting.MinioSetting.EndPoint, &minio.Options{
		Creds:  credentials.NewStaticV4(setting.MinioSetting.AccessKeyID, setting.MinioSetting.SecretAccessKey, ""),
		Secure: false,
	})
	if err != nil {
		fmt.Println("Error initializing MinIO client:", err)
	}

	err = minioClient.MakeBucket(ctx, setting.MinioSetting.BucketName, minio.MakeBucketOptions{})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := minioClient.BucketExists(ctx, setting.MinioSetting.BucketName)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", setting.MinioSetting.BucketName)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("Successfully created %s\n", setting.MinioSetting.BucketName)
	}
}

func AddFile(fileName string, reader io.Reader, fileSize int64) error {
	_, err := minioClient.PutObject(ctx, setting.MinioSetting.BucketName, fileName, reader, fileSize, minio.PutObjectOptions{})
	if err != nil {
		log.Fatalln(err)
		return err
	}

	return nil
}
