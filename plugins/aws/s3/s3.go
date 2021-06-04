package s3

import (
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type Action string

const (
	SIZE    Action = "SIZE"
	UPLOAD  Action = "UPLOAD"
	PRESIGN Action = "PRESIGN"
)

type S3 struct {
	Bucket string `yaml:"bucket"`
	Prefix string `yaml:"prefix"`
	Object string `yaml:"object"`
	Action Action
}

func (s *S3) Size() (int64, error) {
	s3svc := s3.New(session.New())

	input := s3.ListObjectsInput{
		Bucket: aws.String(s.Bucket),
		Prefix: aws.String(s.Prefix),
	}

	result, err := s3svc.ListObjects(&input)
	if err != nil {
		return 0, err
	}

	var size int64

	for _, object := range result.Contents {
		size += int64(*(object.Size))
	}

	return size, nil
}

func (s *S3) Upload() (int64, error) {
	file, err := os.Open(s.Object)
    if err != nil {
        return 0, err
    }
    defer file.Close()

	fileInfo, _ := file.Stat()
    var size int64 = fileInfo.Size()

	_, err = s3.New(session.New()).PutObject(&s3.PutObjectInput{
		Bucket:             aws.String(s.Bucket),
		Key:                aws.String(s.Prefix),
		ACL:                aws.String("private"),
		Body:               file,
		ContentDisposition: aws.String("attachment"),
	})

    return size, err
}

func (s *S3) Presign() (string, error) {
	req, _ := s3.New(session.New()).GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(s.Prefix),
    })

    url, err := req.Presign(15 * time.Minute)

	return url, err
}
