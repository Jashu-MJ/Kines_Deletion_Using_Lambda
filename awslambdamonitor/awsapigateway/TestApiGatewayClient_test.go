package awsapigateway_test

import (
	"awslambdamonitor/awsapigateway"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAPIGatewayClient(t *testing.T) {
	_, err := awsapigateway.CreateAPIGatewayClient()
	assert.NoError(t, err, "error occured at TestCreateAPIGatewayClient")
}
