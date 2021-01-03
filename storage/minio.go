package storage

import (
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"io"
	"log"
	"os"
	"time"
)

var minioClient *minio.Client

func init() {
	endpoint := os.Getenv("MINIO_ENDPOINT")
	accessKeyID := os.Getenv("MINIO_ACCESS_KEY_ID")
	secretAccessKey := os.Getenv("MINIO_SECRET_ACCESS_KEY_ID")
	useSSL := false

	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})

	if err != nil {
		panic(err)
	}

	minioClient = client

	err = MakeBucket("user")
	if err != nil {
		panic(err)
	}
}

func MakeBucket(bucketName string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	err := minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
	if err != nil {
		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			log.Printf("minio : We already own bucket %s\n", bucketName)
		} else {
			return err
		}
	} else {
		log.Printf("minio : Successfully created bucket %s\n", bucketName)
	}
	return nil
}

func Upload(bucketName, fileName string, file io.Reader) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_, err := minioClient.PutObject(ctx, bucketName, fileName, file, -1, minio.PutObjectOptions{})
	if err != nil {
		log.Println(err)
		return "", err
	}
	return fileName, nil
}
