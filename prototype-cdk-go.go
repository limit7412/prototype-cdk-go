package main

import (
	"os"
	"os/exec"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awslambda"
	"github.com/aws/constructs-go/constructs/v3"
	"github.com/aws/jsii-runtime-go"
)

type PrototypeCdkGoStackProps struct {
	awscdk.StackProps
}

func NewPrototypeCdkGoSimpleStack(scope constructs.Construct, id string, props *PrototypeCdkGoStackProps) (awscdk.Stack, error) {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	simpleCmd := exec.Command("go", "build", "-o", "bin/handler/simple/main", "lambda/simple/main.go")
	simpleCmd.Env = append(os.Environ(), "GOOS=linux", "CGO_ENABLED=0")
	_, err := simpleCmd.CombinedOutput()
	if err != nil {
		return nil, err
	}
	_ = awslambda.NewFunction(stack, jsii.String("prototype-go-cdk-simple-lambda"), &awslambda.FunctionProps{
		FunctionName: jsii.String("prototype-go-cdk-function"),
		Runtime:      awslambda.Runtime_GO_1_X(),
		Code:         awslambda.Code_Asset(jsii.String("bin/handler/simple/")),
		Handler:      jsii.String("main"),
	})

	return stack, nil
}

// func NewPrototypeCdkGoAPIStack(scope constructs.Construct, id string, props *PrototypeCdkGoStackProps) (awscdk.Stack, error) {
// 	var sprops awscdk.StackProps
// 	if props != nil {
// 		sprops = props.StackProps
// 	}
// 	stack := awscdk.NewStack(scope, &id, &sprops)

// 	apiGW := awsapigatewayv2.NewHttpApi(stack, jsii.String("prototype-go-cdk-api-gw"), &awsapigatewayv2.HttpApiProps{})

// 	cmd := exec.Command("go", "build", "-o", "bin/handler/api/main", "lambda/api/main.go")
// 	cmd.Env = append(os.Environ(), "GOOS=linux", "CGO_ENABLED=0")
// 	_, err := cmd.CombinedOutput()
// 	if err != nil {
// 		return nil, err
// 	}
// 	apiFn := awslambda.NewFunction(stack, jsii.String("prototype-go-cdk-api-lambda"), &awslambda.FunctionProps{
// 		FunctionName: jsii.String("prototype-go-cdk-function"),
// 		Runtime:      awslambda.Runtime_GO_1_X(),
// 		Code:         awslambda.Code_Asset(jsii.String("bin/handler/api/")),
// 		Handler:      jsii.String("main"),
// 	})
// 	apiGW.AddRoutes(&awsapigatewayv2.AddRoutesOptions{
// 		// FIXME:
// 		// Integration: &awsapigatewayv2.Lambda,
// 		Path:    jsii.String("test"),
// 		Methods: &[]awsapigatewayv2.HttpMethod{"POST"},
// 	})

// 	return stack, nil
// }

// func NewPrototypeCdkGoCronStack(scope constructs.Construct, id string, props *PrototypeCdkGoStackProps) (awscdk.Stack, error) {
// 	var sprops awscdk.StackProps
// 	if props != nil {
// 		sprops = props.StackProps
// 	}
// 	stack := awscdk.NewStack(scope, &id, &sprops)

// 	cronCmd := exec.Command("go", "build", "-o", "bin/handler/cron/main", "lambda/cron/main.go")
// 	cronCmd.Env = append(os.Environ(), "GOOS=linux", "CGO_ENABLED=0")
// 	_, err := cronCmd.CombinedOutput()
// 	if err != nil {
// 		return nil, err
// 	}
// 	cronFn := awslambda.NewFunction(stack, jsii.String("prototype-go-cdk-cron-lambda"), &awslambda.FunctionProps{
// 		FunctionName: jsii.String("prototype-go-cdk-function"),
// 		Runtime:      awslambda.Runtime_GO_1_X(),
// 		Code:         awslambda.Code_Asset(jsii.String("bin/handler/cron/")),
// 		Handler:      jsii.String("main"),
// 	})
// 	awsevents.NewRule(stack, jsii.String("prototype-go-cdk-cron-rule"), &awsevents.RuleProps{
// 		Schedule: awsevents.Schedule_Cron(&awsevents.CronOptions{
// 			Minute:  jsii.String("0"),
// 			Month:   jsii.String("12"),
// 			WeekDay: jsii.String("*"),
// 			Year:    jsii.String("*"),
// 		}),
// 		// FIXME:
// 		// Targets: nil,
// 	})

// 	return stack, nil
// }

func main() {
	app := awscdk.NewApp(nil)

	_, err := NewPrototypeCdkGoSimpleStack(app, "prototype-cdk-go-simple-stack", nil)
	if err != nil {
		panic(err)
	}
	// _, err = NewPrototypeCdkGoAPIStack(app, "prototype-cdk-go-api-stack", nil)
	// if err != nil {
	// 	panic(err)
	// }
	// _, err = NewPrototypeCdkGoAPIStack(app, "prototype-cdk-go-cron-stack", nil)
	// if err != nil {
	// 	panic(err)
	// }

	app.Synth(nil)
}
