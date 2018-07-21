package main

import (
	"os"
	"log"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

const region = "eu-west-1"

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

// Expects two arguments, the name of the bucket and the path to the file to upload to the s3 bucket
func main() {
	if len(os.Args) != 3 {
		log.Fatalf("bucket and file name required. Usage: %s <bucket_name> <filename>", os.Args[0])
	}

	bucket := os.Args[1]
	filename := os.Args[2]

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("%v", err)
	}

	defer file.Close()

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)

	uploader := s3manager.NewUploader(sess)

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key: 	aws.String(filename),
		Body: 	file,
	})
	if err != nil {
		log.Fatalf("unable to upload %q to %q, %v", filename, bucket, err)
	}

	log.Printf("successfully uploaded %q to %q", filename, bucket)
}