package awscloudwatch_test

import (
	"awslambdamonitor/awscloudwatch"
	"testing"

	"github.com/stretchr/testify/assert"
)

//TestCreateCloudWatchClient is used for testing the creation of cloudwatch client creation
func TestCreateCloudWatchClient(t *testing.T) {

	_, err := awscloudwatch.CreateCloudWatchClient()
	assert.NoError(t, err, "error occurred at TestCreateCloudWatchClient ")
}
