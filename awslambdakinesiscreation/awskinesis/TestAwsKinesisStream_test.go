package awskinesis_test

import (
	"awslambdakinesiscreation/api"
	"awslambdakinesiscreation/awskinesis"
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go/service/kinesis"
	"github.com/aws/aws-sdk-go/service/kinesis/kinesisiface"
	"github.com/magiconair/properties/assert"
)

type mockKinesisClient struct {
	kinesisiface.KinesisAPI
	err  error
	resp kinesis.CreateStreamOutput
}

func (m *mockKinesisClient) CreateStream(*kinesis.CreateStreamInput) (*kinesis.CreateStreamOutput, error) {

	return &m.resp, m.err
}

func TestAwsCreateKinesisStreamWithNoError(t *testing.T) {
	mock := &mockKinesisClient{
		err:  nil,
		resp: kinesis.CreateStreamOutput{},
	}
	request := api.RequestBody{ClientName: "jaswanth", ClientSize: "1"}
	response := api.Resp200
	resp := awskinesis.CreateKinesisStream(request, mock)
	assert.Equal(t, response, resp, "Error in creation of kinesis stream")
}

func TestAwsCreateKinesisStreamWithError(t *testing.T) {
	mock := &mockKinesisClient{
		err:  errors.New("insuffiecient params"),
		resp: kinesis.CreateStreamOutput{},
	}
	request := api.RequestBody{ClientName: "", ClientSize: "1"}
	response := api.Resp501
	resp := awskinesis.CreateKinesisStream(request, mock)
	assert.Equal(t, response, resp, "Error in creation of kinesis stream")

}
