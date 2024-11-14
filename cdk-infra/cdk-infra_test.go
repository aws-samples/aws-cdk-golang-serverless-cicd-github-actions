package main

import (
	"testing"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/assertions"
	"github.com/aws/jsii-runtime-go"
)

func TestCdkInfraStack(t *testing.T) {
	// GIVEN
	app := awscdk.NewApp(nil)

	// WHEN
	stack := NewCdkInfraStack(app, "TestStack", nil)

	// THEN
	template := assertions.Template_FromStack(stack, nil)

	// Test if Lambda function is created with correct properties
	template.HasResourceProperties(jsii.String("AWS::Lambda::Function"), map[string]interface{}{
		"Handler": "bootstrap",  // Go Lambda functions use "bootstrap" as the handler
    	"Runtime": "provided.al2",  // For Go Lambda functions
	})

	// Test if Lambda function has the correct tags
	template.HasResourceProperties(jsii.String("AWS::Lambda::Function"), map[string]interface{}{
		"Tags": map[string]interface{}{
			"Environment": "sample",
			"Project":     "CDKInfra",
		},
	})

	// Test if the stack has the expected number of resources
	template.ResourceCountIs(jsii.String("AWS::Lambda::Function"), jsii.Number(1))

	// Test if the Lambda function has the correct IAM role attached
	template.HasResourceProperties(jsii.String("AWS::IAM::Role"), map[string]interface{}{
		"AssumeRolePolicyDocument": map[string]interface{}{
			"Statement": []interface{}{
				map[string]interface{}{
					"Action": "sts:AssumeRole",
					"Effect": "Allow",
					"Principal": map[string]interface{}{
						"Service": "lambda.amazonaws.com",
					},
				},
			},
		},
	})

	// Test if the Lambda function has basic execution role permissions
	template.HasResourceProperties(jsii.String("AWS::IAM::Policy"), map[string]interface{}{
		"PolicyDocument": map[string]interface{}{
			"Statement": []interface{}{
				map[string]interface{}{
					"Action": []string{
						"logs:CreateLogGroup",
						"logs:CreateLogStream",
						"logs:PutLogEvents",
					},
					"Effect":   "Allow",
					"Resource": "arn:aws:logs:*:*:*",
				},
			},
		},
	})

	// Test if the stack has the correct tags
	template.HasResourceProperties(jsii.String("AWS::CloudFormation::Stack"), map[string]interface{}{
		"Tags": []interface{}{
			map[string]interface{}{
				"Key":   "Environment",
				"Value": "sample",
			},
			map[string]interface{}{
				"Key":   "Project",
				"Value": "CDKInfra",
			},
		},
	})

	// Test if the Lambda function has the correct name
	template.HasResourceProperties(jsii.String("AWS::Lambda::Function"), map[string]interface{}{
		"FunctionName": "TestStack-LambdaFunction",
	})

	// Test if the Lambda function has the expected environment variables
	template.HasResourceProperties(jsii.String("AWS::Lambda::Function"), map[string]interface{}{
		"Environment": map[string]interface{}{
			"Variables": map[string]interface{}{
				"ENV": "sample",
			},
		},
	})
}
