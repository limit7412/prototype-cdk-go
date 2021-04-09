package main

import (
	"os/exec"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awslambda"
	"github.com/aws/constructs-go/constructs/v3"
	"github.com/aws/jsii-runtime-go"
)

type PrototypeCdkGoStackProps struct {
	awscdk.StackProps
}

func NewPrototypeCdkGoStack(scope constructs.Construct, id string, props *PrototypeCdkGoStackProps) (awscdk.Stack, error) {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	err := exec.Command("GOOS=linux", "CGO_ENABLED=0", "go", "build", "main.go", "-o", "bin/main").Run()
	if err != nil {
		return nil, err
	}

	awslambda.NewFunction(stack, jsii.String("prototype-go-cdk-id"), &awslambda.FunctionProps{
		FunctionName: jsii.String("prototype-go-cdk-function"),
		Runtime:      awslambda.Runtime_GO_1_X(),
		Code:         awslambda.Code_Asset(jsii.String("bin/")),
		Handler:      jsii.String("main"),
	})

	return stack, nil
}

func main() {
	app := awscdk.NewApp(nil)

	_, err := NewPrototypeCdkGoStack(app, "PrototypeCdkGoStack", &PrototypeCdkGoStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})
	if err != nil {
		panic(err)
	}

	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	// If unspecified, this stack will be "environment-agnostic".
	// Account/Region-dependent features and context lookups will not work, but a
	// single synthesized template can be deployed anywhere.
	//---------------------------------------------------------------------------
	return nil

	// Uncomment if you know exactly what account and region you want to deploy
	// the stack to. This is the recommendation for production stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String("123456789012"),
	//  Region:  jsii.String("us-east-1"),
	// }

	// Uncomment to specialize this stack for the AWS Account and Region that are
	// implied by the current CLI configuration. This is recommended for dev
	// stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
	//  Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	// }
}
