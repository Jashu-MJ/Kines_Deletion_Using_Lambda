package awsapigateway_test

import (
	"awslambdamonitor/awsapigateway"
	"errors"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/apigateway"
	"github.com/aws/aws-sdk-go/service/apigateway/apigatewayiface"
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

type mockAPIGatewayClient struct {
	apigatewayiface.APIGatewayAPI
	resourceList          apigateway.GetResourcesOutput
	deployment            apigateway.Deployment
	method                apigateway.Method
	integration           apigateway.Integration
	GetResourcesError     error
	CreateResourceError   error
	PutMethodError        error
	PutIntegrationError   error
	CreateDeploymentError error
}

func (m *mockAPIGatewayClient) GetResources(resrcsInput *apigateway.GetResourcesInput) (*apigateway.GetResourcesOutput, error) {
	return &m.resourceList, m.GetResourcesError
}

func (m *mockAPIGatewayClient) CreateResource(resrcInput *apigateway.CreateResourceInput) (*apigateway.Resource, error) {
	// passing some random string as resource id
	return &apigateway.Resource{Id: aws.String("2345"), Path: aws.String("/" + *resrcInput.PathPart), PathPart: resrcInput.PathPart, ParentId: resrcInput.ParentId}, m.CreateResourceError

}
func (m *mockAPIGatewayClient) PutMethod(methodInput *apigateway.PutMethodInput) (*apigateway.Method, error) {

	return &m.method, m.PutMethodError
}

func (m *mockAPIGatewayClient) PutIntegration(intInput *apigateway.PutIntegrationInput) (*apigateway.Integration, error) {

	return &m.integration, m.PutIntegrationError
}
func (m *mockAPIGatewayClient) CreateDeployment(deployInput *apigateway.CreateDeploymentInput) (*apigateway.Deployment, error) {

	return &m.deployment, m.CreateDeploymentError
}

func TestCreationOfURLWithCreateResourceError(t *testing.T) {
	os.Setenv("StreamName", "stream22")
	os.Setenv("RESTAPI_ID", "12345")
	os.Setenv("AwsRegion", "us-west-2")

	CreateResourceErr := errors.New("CreateResourceError")
	mockAPI := &mockAPIGatewayClient{CreateResourceError: CreateResourceErr}
	mockLambda := &mockLambdaClient{permission: lambda.AddPermissionOutput{}}
	link := ""
	url, _ := awsapigateway.CreationOfURL(mockAPI, mockLambda)
	assert.Equal(t, url, link, "error occurred in TestCreationOfURL")

}

func TestCreationOfURLWithPutMethodError(t *testing.T) {
	os.Setenv("StreamName", "stream22")
	os.Setenv("RESTAPI_ID", "12345")
	os.Setenv("AwsRegion", "us-west-2")

	PutMethodErr := errors.New("PutMethodError")
	mockAPI := &mockAPIGatewayClient{PutMethodError: PutMethodErr}
	mockLambda := &mockLambdaClient{permission: lambda.AddPermissionOutput{}}
	url := ""
	link, _ := awsapigateway.CreationOfURL(mockAPI, mockLambda)
	assert.Equal(t, url, link, "error occurred in TestCreationOfURL")

}

func TestCreationOfURLWithPutIntegrationError(t *testing.T) {
	os.Setenv("StreamName", "stream22")
	os.Setenv("RESTAPI_ID", "12345")
	os.Setenv("AwsRegion", "us-west-2")

	PutIntegrationErr := errors.New("PutIntegrationError")
	mockAPI := &mockAPIGatewayClient{PutIntegrationError: PutIntegrationErr}
	mockLambda := &mockLambdaClient{permission: lambda.AddPermissionOutput{}}
	url := ""
	link, _ := awsapigateway.CreationOfURL(mockAPI, mockLambda)
	assert.Equal(t, url, link, "error occurred in TestCreationOfURL")

}

func TestCreationOfURLWithCreateDeploymentError(t *testing.T) {
	os.Setenv("StreamName", "stream22")
	os.Setenv("RESTAPI_ID", "12345")
	os.Setenv("AwsRegion", "us-west-2")

	CreateDeploymentErr := errors.New("CreateDeploymentErr")
	mockAPI := &mockAPIGatewayClient{CreateDeploymentError: CreateDeploymentErr}
	mockLambda := &mockLambdaClient{permission: lambda.AddPermissionOutput{}}
	url := "https://12345.execute-api.us-west-2.amazonaws.com/KinesisDeletion/stream22/stream22"
	link, _ := awsapigateway.CreationOfURL(mockAPI, mockLambda)
	assert.Equal(t, url, link, "error occurred in TestCreationOfURL")

}

func TestCreationOfURLWithLambdaPermissionError(t *testing.T) {

	mockAPI := &mockAPIGatewayClient{}
	mockLambda := &mockLambdaClient{permission: lambda.AddPermissionOutput{}, err: errors.New("granting permission error")}
	url := ""
	link, _ := awsapigateway.CreationOfURL(mockAPI, mockLambda)
	assert.Equal(t, url, link, "error occurred in TestCreationOfURL")

}

// case where url is already present and while retrieving the resources from api error occurs
func TestCreationOfURLWithGetResourcesError(t *testing.T) {
	os.Setenv("StreamName", "stream2")
	GetResourcesErr := errors.New("GetResourcesError")

	//resources present at API
	resources := apigateway.GetResourcesOutput{Items: []*apigateway.Resource{
		&apigateway.Resource{Id: aws.String("1234"), Path: aws.String("/stream1"), PathPart: aws.String("stream1"), ParentId: aws.String("0000")},
		&apigateway.Resource{Id: aws.String("2345"), Path: aws.String("/stream2"), PathPart: aws.String("stream2"), ParentId: aws.String("1234")},
		&apigateway.Resource{}},
	}

	mockAPI := &mockAPIGatewayClient{resourceList: resources, GetResourcesError: GetResourcesErr}
	mockLambda := &mockLambdaClient{permission: lambda.AddPermissionOutput{}}
	link := ""
	url, _ := awsapigateway.CreationOfURL(mockAPI, mockLambda)
	assert.Equal(t, url, link, "error occurred in TestCreationOfURL")
}

//case where url already exists and no need to create one more
func TestCreationOfURLWithURLAlreadyExisting(t *testing.T) {
	os.Setenv("StreamName", "stream2")

	resources := apigateway.GetResourcesOutput{Items: []*apigateway.Resource{
		&apigateway.Resource{Id: aws.String("1234"), Path: aws.String("/stream1"), PathPart: aws.String("stream1"), ParentId: aws.String("0000")},
		&apigateway.Resource{Id: aws.String("2345"), Path: aws.String("/stream2"), PathPart: aws.String("stream2"), ParentId: aws.String("1234")},
		&apigateway.Resource{}},
	}
	mockAPI := &mockAPIGatewayClient{resourceList: resources}
	mockLambda := &mockLambdaClient{permission: lambda.AddPermissionOutput{}}
	link := "https://12345.execute-api.us-west-2.amazonaws.com/KinesisDeletion/stream2/stream2"
	url, _ := awsapigateway.CreationOfURL(mockAPI, mockLambda)
	assert.Equal(t, url, link, "error occurred in TestCreationOfURL")
}
