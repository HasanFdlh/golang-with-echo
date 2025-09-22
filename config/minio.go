package config

import (
	"context"
	"log"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var MinioClient *minio.Client

func InitMinio() {
	endpoint := os.Getenv("MINIO_ENDPOINT")
	accessKey := os.Getenv("MINIO_ACCESS_KEY")
	secretKey := os.Getenv("MINIO_SECRET_KEY")
	bucketName := os.Getenv("MINIO_BUCKET")

	// koneksi ke minio (S3 API)
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: false, // false karena pakai http
	})
	if err != nil {
		log.Fatalf("[SYSTEM ERROR] MinIO connection failed: %v", err)
	}

	log.Println("[SYSTEM] MinIO connected")

	// cek bucket
	ctx := context.Background()
	exists, err := minioClient.BucketExists(ctx, bucketName)
	if err != nil {
		log.Fatalf("[SYSTEM ERROR] Bucket check error: %v", err)
	}

	if !exists {
		err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
		if err != nil {
			log.Fatalf("[SYSTEM ERROR] Failed to create bucket %s: %v", bucketName, err)
		}
		log.Printf("[SYSTEM] Bucket %s created\n", bucketName)
	} else {
		log.Printf("[SYSTEM] Bucket %s already exists\n", bucketName)
	}

	MinioClient = minioClient
}
