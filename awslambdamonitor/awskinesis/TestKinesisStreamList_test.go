package awskinesis_test

import (
	"awslambdamonitor/awskinesis"
	"errors"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/kinesis"
	"github.com/aws/aws-sdk-go/service/kinesis/kinesisiface"
	"github.com/stretchr/testify/assert"
)

type mockKinesisClient struct {
	kinesisiface.KinesisAPI
	err  error
	resp kinesis.ListStreamsOutput
}

var x int = 0

func (m *mockKinesisClient) ListStreams(*kinesis.ListStreamsInput) (*kinesis.ListStreamsOutput, error) {
	x = x + 1
	if x == 4 {
		m.resp.SetHasMoreStreams(false)
	}
	return &m.resp, m.err
}

func TestIsStreamExistsWithNonExistingStream(t *testing.T) {
	mock := &mockKinesisClient{resp: kinesis.ListStreamsOutput{StreamNames: []*string{aws.String("jaswanth"), aws.String("stream1"), aws.String("stream2")}}}
	os.Setenv("StreamName", "stream4")

	actualResponse, _ := awskinesis.IsStreamExists(mock)
	expectedResponse := false
	assert.Equal(t, expectedResponse, actualResponse, "Error in searching the kinesis stream")

}

func TestIsStreamExistsWithNonExistingStreams(t *testing.T) {
	mock := &mockKinesisClient{resp: kinesis.ListStreamsOutput{StreamNames: []*string{aws.String("jaswanth"), aws.String("stream1"), aws.String("stream2")}, HasMoreStreams: aws.Bool(true)}}
	os.Setenv("StreamName", "stream4")

	actualResponse, _ := awskinesis.IsStreamExists(mock)
	expectedResponse := false
	assert.Equal(t, expectedResponse, actualResponse, "Error in searching the kinesis stream")

}

func TestIsStreamExistsWithExistingStream(t *testing.T) {
	mock := &mockKinesisClient{resp: kinesis.ListStreamsOutput{StreamNames: []*string{aws.String("jaswanth"), aws.String("stream1"), aws.String("stream2")}}}
	os.Setenv("StreamName", "stream1")

	actualResponse, _ := awskinesis.IsStreamExists(mock)
	expectedResponse := true
	assert.Equal(t, expectedResponse, actualResponse, "Error in searching the kinesis stream")

}
func TestIsStreamExistsWithExistingStreamsAndError(t *testing.T) {
	mock := &mockKinesisClient{resp: kinesis.ListStreamsOutput{StreamNames: []*string{aws.String("jaswanth"), aws.String("stream1"), aws.String("stream2")}, HasMoreStreams: aws.Bool(true)}, err: errors.New("LimitExceededException")}
	os.Setenv("StreamName", "stream1")
	_, err := awskinesis.IsStreamExists(mock)
	expectedResponse := "LimitExceededException"
	assert.Equal(t, expectedResponse, err.Error(), "Error in searching the kinesis stream")

}
