package main

import (
	"cdk-infra/packages/config"
	"cdk-infra/packages/lambda"
	"cdk-infra/packages/tags"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type CdkInfraStackProps struct {
	awscdk.StackProps
}

// read config file
var input = config.Conf()

func NewCdkInfraStack(scope constructs.Construct, id string, props *CdkInfraStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	// // Creates Common Tags
	tags.AddCDKTags(stack, input)
	//create a lambda function
	lambda.CreateLambda(stack, input)

	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	NewCdkInfraStack(app, "CdkInfraStack", &CdkInfraStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

func env() *awscdk.Environment {
	return &awscdk.Environment{
		Account: jsii.String(input.LambdaVariables.AccountNumber),
		Region:  jsii.String(input.LambdaVariables.Region),
	}
}
