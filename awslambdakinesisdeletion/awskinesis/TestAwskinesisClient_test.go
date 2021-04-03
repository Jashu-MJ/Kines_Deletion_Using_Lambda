package awskinesis

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//TestAwsKinesisClient used for testing kinesis client creation
func TestAwsKinesisClient(t *testing.T) {
	_, _, err := KinesisClientCreation()
	assert.NoError(t, nil, "test case at kinesis client has failed")
	if err != nil {
		t.Error()
	}
}
