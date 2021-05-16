package s3

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3 struct {
	Bucket string `yaml:"bucket"`
	Prefix string `yaml:"prefix"`
}

func (s *S3) Size() (int, error) {
	s3svc := s3.New(session.New())

	input := s3.ListObjectsInput {
    	Bucket: aws.String(s.Bucket),
    	Prefix: aws.String(s.Prefix),
	}

	result, err := s3svc.ListObjects(&input)
	if err != nil {
		return 0, err
	}

	var size int

	for _, object := range result.Contents {
		size += int(*(object.Size))
	}

	return size, nil
}
