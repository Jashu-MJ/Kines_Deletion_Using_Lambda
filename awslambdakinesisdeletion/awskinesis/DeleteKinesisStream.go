//Package awskinesis provide kinesis stream deletion
package awskinesis

import (
	"awslambdakinesisdeletion/api"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/kinesis"
	"github.com/aws/aws-sdk-go/service/kinesis/kinesisiface"
)

//DeleteKinesisStreams takes kinesis client and stream name and returns corresponding response and
//error used for deleting kinesis stream which are not required any more.
func DeleteKinesisStreams(kinesisClient kinesisiface.KinesisAPI, streamName string) (events.APIGatewayProxyResponse, error) {
	params := &kinesis.DeleteStreamInput{
		EnforceConsumerDeletion: aws.Bool(true), // ensures that stream deletes even if there is a consumer
		StreamName:              aws.String(streamName),
	}
	_, err := kinesisClient.DeleteStream(params)
	if err != nil {
		log.Println(err)
		return api.Resp501, err
	}
	return api.Resp200, err
}
