package awslambda_test

import (
	"awslambdamonitor/awslambda"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLambdaClientCreation(t *testing.T) {
	_, err := awslambda.LambdaClientCreation()
	assert.NoError(t, err, "Error occurred at TestLambdaClientCreation")
}
