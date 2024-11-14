package lambda

import (
	"cdk-infra/packages/types"
	"os"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/jsii-runtime-go"
)

func CreateLambda(stack awscdk.Stack, input types.Config) awslambda.DockerImageFunction {

	repo := awsecr.Repository_FromRepositoryArn(stack, jsii.String("ECRRepoURI"), jsii.String(os.Getenv("ECR_ARN")))
	imageTagStr := os.Getenv("IMAGETAG")
	envmap := make(map[string]*string)
	for k, v := range input.LambdaVariables.LambdaEnvVar {
		envmap[k] = jsii.String(v)
	}
	// Creates Lambda Function from ECR Image
	lambdaFn := awslambda.NewDockerImageFunction(stack, jsii.String("HandleRequest"), &awslambda.DockerImageFunctionProps{
		FunctionName: jsii.String(input.LambdaVariables.Name),
		Code: awslambda.DockerImageCode_FromEcr(repo, &awslambda.EcrImageCodeProps{
			Tag: jsii.String(imageTagStr),
		}),
		Architecture: awslambda.Architecture_ARM_64(),
		Environment:  &envmap,
		Timeout:      awscdk.Duration_Seconds(jsii.Number(input.LambdaVariables.Timeout)),
	})
	return lambdaFn
}
