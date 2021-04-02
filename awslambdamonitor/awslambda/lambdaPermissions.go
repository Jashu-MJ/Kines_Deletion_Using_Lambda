//Package awslambda has provision to add permissions.
package awslambda

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/aws/aws-sdk-go/service/lambda/lambdaiface"
)

// service which is going to use the lambda function
const service = "apigateway.amazonaws.com"

// action performed on lambda by the service or principal.
const actionPerformed string = "lambda:InvokeFunction"

//GivePermissionsForIntegration for granting api gateway with permission for invoking the lambda function
func GivePermissionsForIntegration(lambdaClient lambdaiface.LambdaAPI, policyID string, ARNOfDeleteFunction string) error {

	params := &lambda.AddPermissionInput{
		Action:       aws.String(actionPerformed),
		FunctionName: aws.String(ARNOfDeleteFunction),
		Principal:    aws.String(service),
		StatementId:  aws.String(policyID),
	}
	_, err := lambdaClient.AddPermission(params)
	return err
}
