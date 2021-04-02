package awskinesis

import (
	"awslambdakinesisdeletion/api"
	"errors"
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go/service/kinesis"
	"github.com/aws/aws-sdk-go/service/kinesis/kinesisiface"
	"github.com/stretchr/testify/assert"
)

type mockKinesisClient struct {
	kinesisiface.KinesisAPI
	err  error
	resp *kinesis.DeleteStreamOutput
}

func (m *mockKinesisClient) DeleteStream(*kinesis.DeleteStreamInput) (*kinesis.DeleteStreamOutput, error) {
	return m.resp, m.err
}

//TestDeleteKinesisStream is used for testing the kinesis stream deletion
func TestDeleteKinesisStreamWithoutError(t *testing.T) {
	mock := &mockKinesisClient{
		err:  nil,
		resp: &kinesis.DeleteStreamOutput{},
	}
	response := api.Resp200
	streamName := "jaswanth"

	resp, _ := DeleteKinesisStreams(mock, streamName)
	assert.Equal(t, response, resp)
	if !reflect.DeepEqual(resp, response) {
		t.Error("test case is not passed in delete kinesis stream at ")
	}

}

func TestDeleteKinesisStreamWithError(t *testing.T) {
	mock := &mockKinesisClient{
		err:  errors.New("error occurred"),
		resp: &kinesis.DeleteStreamOutput{},
	}
	response := api.Resp501
	streamName := "jaswanth"
	resp, _ := DeleteKinesisStreams(mock, streamName)
	assert.Equal(t, response, resp)

}
