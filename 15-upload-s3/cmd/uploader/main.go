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
	s3Client       *s3.S3
	s3Bucket       = "goexpert-bucket-exemplo-bg"
	filesSourceDir = "./tmp"
	filesTargetDir = "pending"
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
	dir, err := os.Open(filesSourceDir)
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

	// listFiles := listFiles()
	// deleteFile(*listFiles[0].Key)
	// moveFile()
}

func uploadFile(filename string) {
	completeFileName := fmt.Sprintf("%s/%s", filesSourceDir, filename)

	fmt.Printf("Uploading file %s to bucket %s\n", completeFileName, s3Bucket)

	f, err := os.Open(completeFileName)
	if err != nil {
		fmt.Printf("Error opening file %s: %v\n", completeFileName, err)
		return
	}
	defer f.Close()

	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(s3Bucket),
		Key:    aws.String(fmt.Sprintf("%s/%s", filesTargetDir, filename)),
		Body:   f,
	})
	if err != nil {
		fmt.Printf("Error uploading file %s\n", completeFileName)
		return
	}

	fmt.Printf("File %s uploaded successfully\n", completeFileName)
}

func listFiles() []*s3.Object {
	resp, err := s3Client.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket:    aws.String(s3Bucket),
		Prefix:    aws.String("pending/"),
		Delimiter: aws.String("/"),
	})
	if err != nil {
		fmt.Printf("Error listing objects: %v\n", err)
		return nil
	}

	fmt.Println(resp.Contents)
	return resp.Contents
}

func deleteFile(fileName string) {
	_, err := s3Client.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(s3Bucket),
		Key:    aws.String(fileName),
	})
	if err != nil {
		fmt.Printf("Error deleting object: %v\n", err)
		return
	}
	fmt.Printf("File %s deleted successfully\n", fileName)
}

func moveFile() {
	target := "processed/meuarquivo.txt"
	source := "/pending/file0.txt"

	_, err := s3Client.CopyObject(&s3.CopyObjectInput{
		Bucket:     aws.String(s3Bucket),
		Key:        aws.String(target),
		CopySource: aws.String(s3Bucket + source),
	})
	if err != nil {
		fmt.Printf("Error copying object: %v\n", err)
		return
	}

	_, err = s3Client.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(s3Bucket),
		Key:    aws.String(source),
	})
	if err != nil {
		fmt.Printf("Error deleting original object: %v\n", err)
		return
	}
	fmt.Printf("File %s copied to %s successfully\n", source, target)
}
