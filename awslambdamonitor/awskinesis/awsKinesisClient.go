//Package awskinesis has kinesis client creation.
package awskinesis

import (
	"awslambdamonitor/awssession"

	"github.com/aws/aws-sdk-go/service/kinesis"
)

//KinesisClientCreation function takes the reference of thesession(*session.Session)
//(once the session is established)as input and creates a kinesis client for further interaction with kinesis
func KinesisClientCreation() (*kinesis.Kinesis, error) {
	awsSess, err := awssession.CreateSession()
	return kinesis.New(awsSess), err
}
