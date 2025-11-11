package main

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var (
	s3Client       *s3.S3
	uploader       *s3manager.Uploader
	downloader     *s3manager.Downloader
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

	uploader = s3manager.NewUploader(sess)
	downloader = s3manager.NewDownloader(sess)
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
		// uploadFile(files[0].Name())
		// uploadMultipartFileNaUnha(files[0].Name())
		uploadMultipartFileComUploaderManager(files[0].Name())
		// downloadFile(files[0].Name())
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

func uploadMultipartFileNaUnha(filename string) {
	completeFileName := fmt.Sprintf("%s/%s", filesSourceDir, filename)

	fmt.Printf("Uploading file %s to bucket %s\n", completeFileName, s3Bucket)

	f, err := os.Open(completeFileName)
	if err != nil {
		fmt.Printf("Error opening file %s: %v\n", completeFileName, err)
		return
	}
	defer f.Close()

	key := fmt.Sprintf("%s/%s", filesTargetDir, filename)

	// 1) cria upload multipart
	createResp, err := s3Client.CreateMultipartUpload(&s3.CreateMultipartUploadInput{
		Bucket: aws.String(s3Bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		fmt.Printf("Error creating multipart upload for file %s: %v\n", completeFileName, err)
		return
	}

	var completedParts []*s3.CompletedPart
	partNumber := int64(1)

	// tamanho de cada parte (5MB mínimo exigido pelo S3 depois da primeira parte)
	const partSize = 5 * 1024 * 1024 // 5MB

	buf := make([]byte, partSize)

	for {
		n, err := f.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Printf("Error reading file %s: %v\n", completeFileName, err)
			return
		}
		if n == 0 {
			break
		}

		// 2) envia cada parte
		uploadResp, err := s3Client.UploadPart(&s3.UploadPartInput{
			Body:          aws.ReadSeekCloser(bytes.NewReader(buf[:n])),
			Bucket:        aws.String(s3Bucket),
			Key:           aws.String(key),
			PartNumber:    aws.Int64(partNumber),
			UploadId:      createResp.UploadId,
			ContentLength: aws.Int64(int64(n)),
		})

		if err != nil {
			// se falhar, aborte o upload multipart
			s3Client.AbortMultipartUpload(&s3.AbortMultipartUploadInput{
				Bucket:   aws.String(s3Bucket),
				Key:      aws.String(key),
				UploadId: createResp.UploadId,
			})
			panic(err)
		}

		completedParts = append(completedParts, &s3.CompletedPart{
			ETag:       uploadResp.ETag,
			PartNumber: aws.Int64(partNumber),
		})

		partNumber++
	}

	// 3) finaliza
	_, err = s3Client.CompleteMultipartUpload(&s3.CompleteMultipartUploadInput{
		Bucket:   aws.String(s3Bucket),
		Key:      aws.String(key),
		UploadId: createResp.UploadId,
		MultipartUpload: &s3.CompletedMultipartUpload{
			Parts: completedParts,
		},
	})

	if err != nil {
		panic(err)
	}

	fmt.Printf("File %s uploaded successfully\n", completeFileName)
}

func uploadMultipartFileComUploaderManager(filename string) {
	completeFileName := fmt.Sprintf("%s/%s", filesSourceDir, filename)

	fmt.Printf("Uploading file %s to bucket %s\n", completeFileName, s3Bucket)

	f, err := os.Open(completeFileName)
	if err != nil {
		fmt.Printf("Error opening file %s: %v\n", completeFileName, err)
		return
	}
	defer f.Close()

	key := fmt.Sprintf("%s/%s", filesTargetDir, filename)

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(s3Bucket),
		Key:    aws.String(key),
		Body:   f,
	})
	if err != nil {
		fmt.Printf("Error uploading file %s: %v\n", completeFileName, err)
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

func downloadFile(filename string) {
	completeFileName := fmt.Sprintf("%s/%s", filesTargetDir, filename)

	f, err := os.Create("./baixado/" + filename)
	if err != nil {
		fmt.Printf("Error creating file %s: %v\n", completeFileName, err)
		return
	}
	defer f.Close()

	bytes, err := downloader.Download(f, &s3.GetObjectInput{
		Bucket: aws.String(s3Bucket),
		Key:    aws.String(completeFileName),
	})
	if err != nil {
		fmt.Printf("Error downloading file %s: %v\n", completeFileName, err)
		return
	}

	fmt.Printf("File %s downloaded successfully, %d bytes\n", completeFileName, bytes)
}
