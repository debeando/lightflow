package plugins

import (
	"errors"

	"github.com/debeando/lightflow/common"
	"github.com/debeando/lightflow/common/log"
	"github.com/debeando/lightflow/common/template"
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

	p.Variables.Set(map[string]interface{}{
		"aws_s3_objects_size": int64(0),
		"aws_s3_objects_presign": "",
	})

	err := p.isValid()
	if err != nil {
		return
	}

	p.Config.S3.Bucket = p.Render(p.Bucket())
	p.Config.S3.Prefix = p.Render(p.Config.S3.Prefix)
	p.Config.S3.Object = p.Render(p.Config.S3.Object)

	s := s3.S3{
		Bucket: p.Config.S3.Bucket,
		Prefix: p.Config.S3.Prefix,
		Object: p.Config.S3.Object,
		Action: p.Config.S3.Action,
	}

	switch action := p.Config.S3.Action; action {
	case s3.SIZE:
		size, err := s.Size()
		if err != nil {
			log.Error(err.Error(), nil)
		}

		p.Variables.Set(map[string]interface{}{
			"aws_s3_objects_size": size,
		})
	case s3.UPLOAD:
		size, err := s.Upload()
		if err != nil {
			log.Error(err.Error(), nil)
		}
		p.Variables.Set(map[string]interface{}{
			"aws_s3_objects_size": size,
		})
	case s3.PRESIGN:
		url, err := s.Presign()
		if err != nil {
			log.Error(err.Error(), nil)
		}
		p.Variables.Set(map[string]interface{}{
			"aws_s3_objects_presign": url,
		})
	}
}

func (c *PluginAWS) Render(s string) string {
	r, err := template.Render(s, c.Variables.GetItems())
	if err != nil {
		log.Warning(err.Error(), nil)
	}
	return r
}

func (c *PluginAWS) isValid() error {
	if len(c.Bucket()) == 0 {
		return errors.New("Plugin AWS/S3: Require valid Bucket name.")
	}

	if len(c.Config.S3.Prefix) == 0 {
		return errors.New("Plugin AWS/S3: Require valid Prefix name.")
	}

	return nil
}

func (c *PluginAWS) Bucket() string {
	bucket := c.Config.S3.Bucket

	if len(bucket) == 0 {
		return common.InterfaceToString(c.Variables.Get("aws_s3_bucket"))
	}

	return bucket
}
