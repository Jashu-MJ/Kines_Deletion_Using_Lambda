package awssession_test

import (
	//"awslambda/awssession"
	"awslambdakinesiscreation/awssession"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAwsSession(t *testing.T) { //Done
	_, err := awssession.CreateSession()
	assert.NoError(t, err, "aws session test is failed, session could not be established")
}
