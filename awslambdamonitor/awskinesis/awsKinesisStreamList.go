//Package awskinesis contais the list of kinesis streams
package awskinesis

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/kinesis"
	"github.com/aws/aws-sdk-go/service/kinesis/kinesisiface"
)

//IsStreamExists is used for whethere the kinesis stream exists or not in aws
func IsStreamExists(kinesisClient kinesisiface.KinesisAPI) (bool, error) {
	stream := os.Getenv("StreamName")
	streams, err := getListOfStreams(kinesisClient)
	if err != nil {
		return false, err
	}

	for i := range streams {
		if stream == *streams[i] {
			return true, nil
		}
	}
	return false, nil

}

func getListOfStreams(kinesisClient kinesisiface.KinesisAPI) ([]*string, error) {
	list := []*string{}

	params := &kinesis.ListStreamsInput{Limit: aws.Int64(500)} // limiting the streams to 500
	resp, err := kinesisClient.ListStreams(params)
	if err != nil {
		return nil, err
	}
	list = append(list, resp.StreamNames...)
	if aws.BoolValue(resp.HasMoreStreams) != true {
		return resp.StreamNames, nil
	}
	for aws.BoolValue(resp.HasMoreStreams) == true {
		resp, err := kinesisClient.ListStreams(params)
		if err != nil {
			return nil, err
		}
		list = append(list, resp.StreamNames...)

	}
	return list, nil
}
