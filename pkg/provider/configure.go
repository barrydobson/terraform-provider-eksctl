package provider

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/mumoshu/terraform-provider-eksctl/pkg/resource"
)

type ProviderInstance struct {
	AWSSession *session.Session
}

func providerConfigure() func(*schema.ResourceData) (interface{}, error) {
	return func(d *schema.ResourceData) (interface{}, error) {
		s := resource.AWSSessionFromResourceData(d)

		return &ProviderInstance{
			AWSSession: s,
		}, nil
	}
}
