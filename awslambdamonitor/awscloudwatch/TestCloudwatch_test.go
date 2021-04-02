package awscloudwatch_test

import (
	"awslambdamonitor/awscloudwatch"
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/aws/aws-sdk-go/service/cloudwatch/cloudwatchiface"
	"github.com/stretchr/testify/assert"
)

type mockCloudWatchClient struct {
	cloudwatchiface.CloudWatchAPI
	resp cloudwatch.GetMetricStatisticsOutput
	err  error
}

func (m *mockCloudWatchClient) GetMetricStatistics(input *cloudwatch.GetMetricStatisticsInput) (*cloudwatch.GetMetricStatisticsOutput, error) {

	return &m.resp, m.err
}
func TestMonitorCloudKinesisStreamLogsWithNonZeroDataPoints(t *testing.T) {

	mock := &mockCloudWatchClient{
		resp: cloudwatch.GetMetricStatisticsOutput{Datapoints: []*cloudwatch.Datapoint{
			&cloudwatch.Datapoint{Sum: aws.Float64(0)},
			&cloudwatch.Datapoint{Sum: aws.Float64(1)},
			&cloudwatch.Datapoint{Sum: aws.Float64(2)},
			&cloudwatch.Datapoint{Sum: aws.Float64(0)}}},
		err: nil,
	}
	value := false
	val, _ := awscloudwatch.MonitorCloudKinesisStreamLogs(mock)
	assert.Equal(t, value, val, "")
}
func TestMonitorCloudKinesisStreamLogsWithZeroDataPoints(t *testing.T) {
	mock := &mockCloudWatchClient{
		resp: cloudwatch.GetMetricStatisticsOutput{},
		err:  errors.New("error occurred"),
	}
	value := true
	val, _ := awscloudwatch.MonitorCloudKinesisStreamLogs(mock)
	assert.Equal(t, value, val, "")
}
