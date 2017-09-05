package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"hue-ec-swat-golang/common/awsctl"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	ctx := context.Background()
	awssvc := awsctl.NewAWS("", "", "")
	filePath := "/home/weiming/go/src/go-learning/awstest/seq_attack_result_20180125173458.png"
	awssvc.UploadFile(ctx, "load-tester-result", filePath, "newtest/seq_attack_result_20180125173458.png", time.Minute)
}

func awsLib() {
	accessKeyID := ""
	secretAccessKey := ""

	bucket := "load-tester-result"
	key := "report/test/seq_attack_result_20180125173458.png"
	timeout := time.Duration(60 * time.Second)

	// All clients require a Session. The Session provides the client with
	// shared configuration such as region, endpoint, and credentials. A
	// Session should be shared where possible to take advantage of
	// configuration and credential caching. See the session package for
	// more information.
	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(endpoints.ApNortheast1RegionID),
		Credentials: credentials.NewStaticCredentials(accessKeyID, secretAccessKey, ""),
	}))

	filePath := "/home/weiming/go/src/go-learning/awstest/seq_attack_result_20180125173458.png"
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Get file size and read the file content into a buffer
	fileInfo, _ := file.Stat()
	var size int64 = fileInfo.Size()
	buffer := make([]byte, size)
	file.Read(buffer)

	// Create a new instance of the service's client with a Session.
	// Optional aws.Config values can also be provided as variadic arguments
	// to the New function. This option allows you to provide service
	// specific configuration.
	svc := s3.New(sess)

	// Create a context with a timeout that will abort the upload if it takes
	// more than the passed in timeout.
	ctx := context.Background()
	var cancelFn func()
	if timeout > 0 {
		ctx, cancelFn = context.WithTimeout(ctx, timeout)
	}
	// Ensure the context is canceled to prevent leaking.
	// See context package for more information, https://golang.org/pkg/context/
	defer cancelFn()

	// Uploads the object to S3. The Context will interrupt the request if the
	// timeout expires.
	_, err = svc.PutObjectWithContext(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		//ContentType: aws.String("image/png"),
		Body:                 bytes.NewReader(buffer),
		ContentLength:        aws.Int64(size),
		ContentType:          aws.String(http.DetectContentType(buffer)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
	})

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok && aerr.Code() == request.CanceledErrorCode {
			// If the SDK can determine the request or retry delay was canceled
			// by a context the CanceledErrorCode error code will be returned.
			fmt.Fprintf(os.Stderr, "upload canceled due to timeout, %v\n", err)
		} else {
			fmt.Fprintf(os.Stderr, "failed to upload object, %v\n", err)
		}
		os.Exit(1)
	}

	fmt.Printf("successfully uploaded file to %s/%s\n", bucket, key)

}

func minioLib() {
	endpoint := "s3.amazonaws.com"
	accessKeyID := ""
	secretAccessKey := ""
	useSSL := true
	ctx := context.Background()

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		log.Fatalln("Error:", err)
	}

	log.Printf("Success: %#v\n", minioClient) // minioClient is now setup

	bucketName := "load-tester-result"

	exists, err := minioClient.BucketExists(bucketName)
	if err != nil {
		log.Fatalln(err)
	}

	if exists {
		log.Printf("We already own %s\n", bucketName)
	} else {
		location := "ap-northeast-1"
		err = minioClient.MakeBucket(bucketName, location)
		if err != nil {
			log.Fatalln(err)
		}
		log.Printf("Successfully created %s\n", bucketName)
	}

	objectName := "seq_attack_result_20180125173458.png"
	filePath := "/home/weiming/go/src/go-learning/awstest/seq_attack_result_20180125173458.png"
	contentType := "image/png"

	putOpts := minio.PutObjectOptions{
		ContentType: contentType,
	}
	n, err := minioClient.FPutObjectWithContext(ctx, bucketName, objectName, filePath, putOpts)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Successfully uploaded %s of size %d\n", objectName, n)
}
