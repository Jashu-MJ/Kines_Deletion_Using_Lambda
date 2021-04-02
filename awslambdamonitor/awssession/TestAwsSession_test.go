package awssession_test

import (
	"awslambdamonitor/awssession"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAwsSession(t *testing.T) { //Done
	_, err := awssession.CreateSession()
	assert.NoError(t, err, "Error occurred at TestAwsSession")
}
