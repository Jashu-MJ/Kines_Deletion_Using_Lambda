//Package awskinesis has kinesis stream creation and parameters for creation of stream
package awskinesis

import (
	"log"
	"strconv"

	"awslambdakinesiscreation/api"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/kinesis"
	"github.com/aws/aws-sdk-go/service/kinesis/kinesisiface"
)

//CreateKinesisStream function is used for creating the kinesis stream which
//takes in kinesis client,no of shards and stream name as parameters and
//creates a kinesis stream with passed params and upon succeful creation it
// will return response
func CreateKinesisStream(apiReqBody api.RequestBody, kinesisClient kinesisiface.KinesisAPI) events.APIGatewayProxyResponse {

	//geting the details from the request body in the form of Api.RequestBody
	StreamName, NoOfShards := getKinesisStreamDetails(apiReqBody)

	//Creating the stream begins
	params := &kinesis.CreateStreamInput{
		ShardCount: aws.Int64(NoOfShards),  // Required,*kinesisiface.KinesisAPI
		StreamName: aws.String(StreamName), // Required
	}
	_, err := kinesisClient.CreateStream(params)
	if err != nil {
		return api.Resp501
	}
	log.Println("creating kinesis stream")
	return api.Resp200
}

func getKinesisStreamDetails(apireq api.RequestBody) (string, int64) {
	NoOfShards, _ := strconv.Atoi(apireq.ClientSize)
	return apireq.ClientName, int64(NoOfShards)

}
