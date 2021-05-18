package plugins

import (
	"github.com/debeando/lightflow/common"
	"github.com/debeando/lightflow/common/log"
	"github.com/debeando/lightflow/flow/template"
	"github.com/debeando/lightflow/plugins/aws"
	"github.com/debeando/lightflow/plugins/aws/s3"
	"github.com/debeando/lightflow/variables"
)

type PluginAWS struct {
	Config    aws.AWS
	Variables variables.List
}

func (p *PluginAWS) Load() {
	p.Variables = *variables.Load()
	p.Config.S3.Bucket = p.Render(p.Config.S3.Bucket)
	p.Config.S3.Prefix = p.Render(p.Config.S3.Prefix)

	if !p.isValid() {
		return
	}

	s := s3.S3{
		Bucket: p.Config.S3.Bucket,
		Prefix: p.Config.S3.Prefix,
	}

	size, err := s.Size()
	if err != nil {
		log.Error(err.Error(), nil)
	}

	p.Variables.Set(map[string]interface{}{
		"aws_s3_objects_size": size,
	})
}

func (c *PluginAWS) Render(s string) string {
	r, err := template.Render(s, c.Variables.GetItems())
	if err != nil {
		log.Warning(err.Error(), nil)
	}
	return r
}

func (c *PluginAWS) isValid() bool {
	if len(c.Bucket()) == 0 {
		return false
	}

	if len(c.Config.S3.Prefix) == 0 {
		return false
	}

	return true
}

func (c *PluginAWS) Bucket() string {
	bucket := c.Config.S3.Bucket

	if len(bucket) == 0 {
		return common.InterfaceToString(c.Variables.Get("aws_s3_bucket"))
	}

	return bucket
}
