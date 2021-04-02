//Package awslambda has lambda client
package awslambda

import (
	"awslambdamonitor/awssession"

	"github.com/aws/aws-sdk-go/service/lambda"
)

//LambdaClientCreation is used for creating lambda client
func LambdaClientCreation() (*lambda.Lambda, error) {
	awsSession, err := awssession.CreateSession()
	return lambda.New(awsSession), err
}
