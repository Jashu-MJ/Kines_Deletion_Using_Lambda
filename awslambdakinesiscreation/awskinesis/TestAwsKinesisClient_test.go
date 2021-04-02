package awskinesis_test

import (
	"awslambdakinesiscreation/awskinesis"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAwsKinesisClient(t *testing.T) {
	_, _, err := awskinesis.KinesisClientCreation()
	assert.NoError(t, err, "error occured at kinesis client creation")
}
