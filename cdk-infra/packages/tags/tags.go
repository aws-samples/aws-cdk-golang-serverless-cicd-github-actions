package tags

import (
	"cdk-infra/packages/types"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/jsii-runtime-go"
)

func AddCDKTags(stack awscdk.Stack, Common_Tags types.Config) {

	for k, v := range Common_Tags.Common_Tags {

		awscdk.Tags_Of(stack).Add(jsii.String(k), jsii.String(v), &awscdk.TagProps{})
	}
}
