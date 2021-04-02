package awssns_test

import (
	"awslambdamonitor/awssns"
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sns/snsiface"
	"github.com/stretchr/testify/assert"
)

type mockSNSClient struct {
	snsiface.SNSAPI
	resp sns.PublishOutput
	err  error
}

func (m *mockSNSClient) Publish(*sns.PublishInput) (*sns.PublishOutput, error) {

	return &m.resp, m.err
}

//TestSendEmailToTopicWithURL is used to test whether email is sent or not
func TestSendEmailToTopicWithURL(t *testing.T) {
	mock := &mockSNSClient{}

	url := "this is url"

	err := awssns.SendEmailToTopicWithURL(mock, url)

	assert.NoError(t, err, "Error occurred at TestSendEmailToTopicWithURL")

}
func TestSendEmailToTopicWithURLWithError(t *testing.T) {
	Error := errors.New("error occurred")

	mock := &mockSNSClient{err: Error}
	url := "this is url"

	err := awssns.SendEmailToTopicWithURL(mock, url)

	assert.Error(t, err, "Error occurred at TestSendEmailToTopicWithURL")
	assert.Equal(t, Error, err)
}
