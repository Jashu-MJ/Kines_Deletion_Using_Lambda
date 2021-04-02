package awslambda_test

import (
	"awslambdamonitor/awslambda"
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/aws/aws-sdk-go/service/lambda/lambdaiface"
	"github.com/stretchr/testify/assert"
)

type mockLambdaClient struct {
	lambdaiface.LambdaAPI
	permission lambda.AddPermissionOutput
	err        error
}

func (m *mockLambdaClient) AddPermission(*lambda.AddPermissionInput) (*lambda.AddPermissionOutput, error) {

	return &m.permission, m.err
}

func TestGivePermissionsForIntegration(t *testing.T) {

	mock := &mockLambdaClient{}
	statementId := "12"
	ARNLambda := "1234567890"
	//var err error
	err := awslambda.GivePermissionsForIntegration(mock, statementId, ARNLambda)
	assert.NoError(t, err, "Error occurred at TestGivePermissionsForIntegration")
}
func TestGivePermissionsForIntegrationWithError(t *testing.T) {
	errOccurred := errors.New("error at granting permissionn")
	mock := &mockLambdaClient{err: errOccurred}
	statementId := "34"
	ARNLambda := "1234567890"
	err := awslambda.GivePermissionsForIntegration(mock, statementId, ARNLambda)
	assert.Error(t, err, "Error occurred at TestGivePermissionsForIntegration")
}
