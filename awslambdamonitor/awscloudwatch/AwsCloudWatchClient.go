//Package awscloudwatch has cloud watch client creation,
package awscloudwatch

import (
	"awslambdamonitor/awssession"

	"github.com/aws/aws-sdk-go/service/cloudwatch"
)

//CreateCloudWatchClient used for creating a cloud watch client to communicate with aws cloud watch.
func CreateCloudWatchClient() (*cloudwatch.CloudWatch, error) {
	awsSession, err := awssession.CreateSession()

	return cloudwatch.New(awsSession), err
}
