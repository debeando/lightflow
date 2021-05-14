package flow

import (
	"github.com/debeando/lightflow/common"
	"github.com/debeando/lightflow/common/log"

	"github.com/debeando/lightflow/flow/aws/s3"
)

func (f *Flow) aws_s3() {
	if ! f.isValidAWSS3() {
		return
	}

	s := s3.S3{
		Bucket: f.Render(f.GetAWSS3Bucket()),
		Prefix: f.Render(f.GetAWSS3Prefix()),
	}

	size, err := s.Size()
	if err != nil {
		log.Error(err.Error(), nil)
	}

	f.Variables.Set(map[string]interface{}{
		"aws_s3_objects_size": size,
	})
}

func (f *Flow) isValidAWSS3() bool {
	bucket := f.GetAWSS3Bucket()
	prefix := f.GetAWSS3Prefix()

	if len(bucket) == 0 {
		return false
	}

	if len(prefix) == 0 {
		return false
	}

	return true
}

func (f *Flow) GetAWSS3Bucket() string {
	bucket := f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].AWS.S3.Bucket

	if len(bucket) == 0 {
		return common.InterfaceToString(f.Variables.Get("aws_s3_bucket"))
	}

	return bucket
}

func (f *Flow) GetAWSS3Prefix() string {
	return f.Config.Tasks[f.Index.Task].Pipes[f.Index.Pipe].AWS.S3.Prefix
}
