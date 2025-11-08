package main

import (
	"fmt"
	"io"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	s3Client *s3.S3
	s3Bucket = "goexpert-bucket-exemplo-bg"
	filesDir = "./tmp"
)

func init() {
	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String("us-east-1"),
			Credentials: credentials.NewStaticCredentials(
				"---",
				"---",
				"",
			),
		},
	)
	if err != nil {
		panic(err)
	}
	s3Client = s3.New(sess)
}

func main() {
	dir, err := os.Open(filesDir)
	if err != nil {
		panic(err)
	}
	defer dir.Close()

	for {
		files, err := dir.Readdir(1)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Printf("Error reading directory: %v\n", err)
			continue
		}
		uploadFile(files[0].Name())
	}
}

func uploadFile(filename string) {
	completeFileName := fmt.Sprintf("%s/%s", filesDir, filename)

	fmt.Printf("Uploading file %s to bucket %s\n", completeFileName, s3Bucket)

	f, err := os.Open(completeFileName)
	if err != nil {
		fmt.Printf("Error opening file %s: %v\n", completeFileName, err)
		return
	}
	defer f.Close()

	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(s3Bucket),
		Key:    aws.String(filename),
		Body:   f,
	})
	if err != nil {
		fmt.Printf("Error uploading file %s\n", completeFileName)
		return
	}

	fmt.Printf("File %s uploaded successfully\n", completeFileName)
}
