//Package awssns has sns client.
package awssns

import (
	"awslambdamonitor/awssession"

	"github.com/aws/aws-sdk-go/service/sns"
)

//SnsClientCreation is used
func SnsClientCreation() (*sns.SNS, error) {
	awsSess, err := awssession.CreateSession()
	return sns.New(awsSess), err
}
