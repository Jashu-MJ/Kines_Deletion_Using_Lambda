package awssns_test

import (
	"awslambdamonitor/awssns"
	"testing"

	"github.com/stretchr/testify/assert"
)

//TestSnsClientCreation is used for testing the sns client creation
func TestSnsClientCreation(t *testing.T) {

	_, err := awssns.SnsClientCreation()
	assert.NoError(t, err, "Error occurred at TestSnsClientCreation")

}
