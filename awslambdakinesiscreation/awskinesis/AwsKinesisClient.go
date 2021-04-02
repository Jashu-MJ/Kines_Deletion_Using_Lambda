//Package awskinesis has kinesis client creation.
package awskinesis

import (
	"awslambdakinesiscreation/awssession"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/kinesis"
)

//KinesisClientCreation function takes the reference of thesession(*session.Session)
//(once the session is established)as input and creates a kinesis client for further interaction with kinesis
func KinesisClientCreation() (*kinesis.Kinesis, events.APIGatewayProxyResponse, error) {
	awsSess, err := awssession.CreateSession()
	return kinesis.New(awsSess), events.APIGatewayProxyResponse{}, err
}
