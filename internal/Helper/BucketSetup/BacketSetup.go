package bucketsetup

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var (
	S3Client        *s3.Client
	Presigner       *s3.PresignClient
	miniobucketName string
	s3bucketName    string
)

var MinioClient *minio.Client

func InitMinioClient() error {
	switch os.Getenv("BUCKETACCESS") {

	case "minio":
		endpoint := os.Getenv("MINIO_ENDPOINT")
		portStr := os.Getenv("MINIO_PORT")
		useSSL := os.Getenv("MINIO_USE_SSL") == "true"
		accessKey := os.Getenv("MINIO_ACCESS_KEY")
		secretKey := os.Getenv("MINIO_SECRET_KEY")

		miniobucketName = os.Getenv("MINIO_BUCKET") // ✅ set global

		port, err := strconv.Atoi(portStr)
		if err != nil {
			return fmt.Errorf("invalid MINIO_PORT: %v", err)
		}

		fullEndpoint := fmt.Sprintf("%s:%d", endpoint, port)

		client, err := minio.New(fullEndpoint, &minio.Options{
			Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
			Secure: useSSL,
		})
		if err != nil {
			return fmt.Errorf("failed to initialize MinIO client: %v", err)
		}

		MinioClient = client
		fmt.Println("✅ MinIO client initialized")

		ctx := context.Background()
		exists, err := client.BucketExists(ctx, miniobucketName)
		if err != nil {
			return fmt.Errorf("error checking bucket: %v", err)
		}

		if !exists {
			err = client.MakeBucket(ctx, miniobucketName, minio.MakeBucketOptions{Region: "us-east-1"})
			if err != nil {
				return fmt.Errorf("failed to create bucket: %v", err)
			}
			fmt.Printf("✅ Bucket %s created\n", miniobucketName)
		}

		return nil

	case "aws":
		ctx := context.Background()

		s3bucketName = os.Getenv("AWS_S3_BUCKET") // ✅ set global

		cfg, err := config.LoadDefaultConfig(ctx)
		if err != nil {
			return fmt.Errorf("failed to load AWS config: %w", err)
		}

		S3Client = s3.NewFromConfig(cfg)
		Presigner = s3.NewPresignClient(S3Client)

		_, err = S3Client.HeadBucket(ctx, &s3.HeadBucketInput{
			Bucket: aws.String(s3bucketName),
		})
		if err != nil {
			return fmt.Errorf("bucket does not exist or access denied: %w", err)
		}

		fmt.Println("✅ AWS S3 client initialized")
		return nil

	default:
		return fmt.Errorf("invalid BUCKETACCESS value: %s (use 'minio' or 'aws')", os.Getenv("BUCKETACCESS"))
	}
}

func GetFileURL(fileName string, expireMins int) (string, error) {

	switch os.Getenv("BUCKETACCESS") {

	case "minio":
		if MinioClient == nil {
			return "", fmt.Errorf("Minio client is not initialized, call InitMinioClient() first")
		}

		ctx := context.Background()
		expiry := time.Duration(expireMins) * time.Minute

		url, err := MinioClient.PresignedGetObject(ctx, miniobucketName, fileName, expiry, nil)
		if err != nil {
			return "", fmt.Errorf("failed to generate file URL: %w", err)
		}
		return url.String(), nil
	case "aws":
		ctx := context.Background()
		expiry := time.Duration(expireMins) * time.Minute

		req, err := Presigner.PresignGetObject(ctx, &s3.GetObjectInput{
			Bucket:                     aws.String(s3bucketName),
			Key:                        aws.String(fileName),
			ResponseContentDisposition: aws.String("inline"),
		}, s3.WithPresignExpires(expiry))

		if err != nil {
			return "", err
		}

		return req.URL, nil

	default:
		return "", fmt.Errorf("invalid BUCKETACCESS value: %s (use 'minio' or 'aws')", os.Getenv("BUCKETACCESS"))

	}
}

func CreateUploadURL(fileName string, expireMins int) (uploadUrl string, fileUrl string, err error) {

	switch os.Getenv("BUCKETACCESS") {
	case "minio":
		if MinioClient == nil {
			return "", "", fmt.Errorf("Minio client is not initialized, call InitMinioClient() first")
		}

		ctx := context.Background()
		expiry := time.Duration(expireMins) * time.Minute

		uploadUrlObj, err := MinioClient.PresignedPutObject(ctx, miniobucketName, fileName, expiry)
		if err != nil {
			return "", "", fmt.Errorf("failed to generate upload URL: %w", err)
		}
		uploadUrl = uploadUrlObj.String()

		fileUrl, err = GetFileURL(fileName, expireMins)
		if err != nil {
			return "", "", err
		}

		return uploadUrl, fileUrl, nil

	case "aws":
		if Presigner == nil {
			return "", "", fmt.Errorf("S3 client not initialized")
		}

		ctx := context.Background()
		expiry := time.Duration(expireMins) * time.Minute

		putReq, err := Presigner.PresignPutObject(ctx, &s3.PutObjectInput{
			Bucket: aws.String(s3bucketName),
			Key:    aws.String(fileName),
		}, s3.WithPresignExpires(expiry))
		if err != nil {
			return "", "", err
		}

		fileUrl, err = GetFileURL(fileName, expireMins)
		if err != nil {
			return "", "", err
		}

		return putReq.URL, fileUrl, nil

	default:
		return "", "", fmt.Errorf("invalid BUCKETACCESS value: %s (use 'minio' or 'aws')", os.Getenv("BUCKETACCESS"))

	}
}

func CreateDownloadURL(objectName, downloadName string, expireMins int) (string, error) {

	switch os.Getenv("BUCKETACCESS") {
	case "minio":
		if MinioClient == nil {
			return "", fmt.Errorf("Minio client is not initialized, call InitMinioClient() first")
		}

		ctx := context.Background()
		expiry := time.Duration(expireMins) * time.Minute

		// Use url.Values to set custom response headers
		reqParams := make(url.Values)
		// Force browser to download with given name
		reqParams.Set("response-content-disposition", fmt.Sprintf("attachment; filename=\"%s\"", downloadName))

		downloadURL, err := MinioClient.PresignedGetObject(ctx, miniobucketName, objectName, expiry, reqParams)
		if err != nil {
			return "", fmt.Errorf("failed to generate download URL: %w", err)
		}

		return downloadURL.String(), nil

	case "aws":
		ctx := context.Background()
		expiry := time.Duration(expireMins) * time.Minute

		req, err := Presigner.PresignGetObject(ctx, &s3.GetObjectInput{
			Bucket: aws.String(s3bucketName),
			Key:    aws.String(objectName),
			ResponseContentDisposition: aws.String(
				fmt.Sprintf("attachment; filename=\"%s\"", downloadName),
			),
		}, s3.WithPresignExpires(expiry))

		if err != nil {
			return "", err
		}

		return req.URL, nil

	default:
		return "", fmt.Errorf("invalid BUCKETACCESS value: %s (use 'minio' or 'aws')", os.Getenv("BUCKETACCESS"))

	}

}

func DeleteFile(fileName string) error {

	switch os.Getenv("BUCKETACCESS") {
	case "minio":
		if MinioClient == nil {
			return fmt.Errorf("Minio client is not initialized, call InitMinioClient() first")
		}

		ctx := context.Background()
		err := MinioClient.RemoveObject(ctx, miniobucketName, fileName, minio.RemoveObjectOptions{})
		if err != nil {
			return fmt.Errorf("failed to delete file %s: %w", fileName, err)
		}

		fmt.Printf("✅ File %s deleted successfully\n", fileName)
		return err
	case "aws":
		_, err := S3Client.DeleteObject(context.Background(), &s3.DeleteObjectInput{
			Bucket: aws.String(s3bucketName),
			Key:    aws.String(fileName),
		})
		return err
	default:
		return fmt.Errorf("invalid BUCKETACCESS value: %s (use 'minio' or 'aws')", os.Getenv("BUCKETACCESS"))
	}
}
