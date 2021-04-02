//Package awscloudwatch has function for fetching data points from cloud watch logs.
package awscloudwatch

import (
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/aws/aws-sdk-go/service/cloudwatch/cloudwatchiface"
)

//StreamName is the dimension in cloudwatch
const StreamName string = "StreamName"

// getRecordBytes is metric type for kinesis stream
const getRecordBytes string = "GetRecords.Bytes"

// name space for kinesis stream in cloud watch
const nameSpaceForKinesis string = "AWS/Kinesis"

// time period of every 30 minutes
const timePeriod int64 = 1800

// time unit
const seconds string = "Seconds"

// incomingBytes is metric type for kinesis stream
const incomingBytes string = "IncomingBytes"

// getDatapoints fetches data points of incoming bytes and outgoing bytes  from cloud watch logs.
func getDatapoints(cloudWatchClient cloudwatchiface.CloudWatchAPI) ([]*cloudwatch.Datapoint, []*cloudwatch.Datapoint, error) {

	paramsForGetRecordsBytes := &cloudwatch.GetMetricStatisticsInput{

		Dimensions: []*cloudwatch.Dimension{&cloudwatch.Dimension{Name: aws.String(StreamName), Value: aws.String(os.Getenv("StreamName"))}},

		EndTime: aws.Time(time.Now()),

		MetricName: aws.String(getRecordBytes),

		Namespace: aws.String(nameSpaceForKinesis),

		Period: aws.Int64(timePeriod),

		StartTime: aws.Time(time.Now().Add(-10 * time.Hour)),

		Statistics: aws.StringSlice([]string{"Sum"}),

		Unit: aws.String(seconds),
	}
	paramsForIncomingBytes := &cloudwatch.GetMetricStatisticsInput{

		Dimensions: []*cloudwatch.Dimension{&cloudwatch.Dimension{Name: aws.String(StreamName), Value: aws.String(os.Getenv("StreamName"))}},

		EndTime: aws.Time(time.Now()),

		MetricName: aws.String(incomingBytes),

		Namespace: aws.String(nameSpaceForKinesis),

		Period: aws.Int64(timePeriod),

		StartTime: aws.Time(time.Now().Add(-10 * time.Hour)),

		Statistics: aws.StringSlice([]string{"Sum"}),

		Unit: aws.String(seconds),
	}

	// getting metric statistics for outgoing bytes of a stream
	respOfGetRecordsBytes, err := cloudWatchClient.GetMetricStatistics(paramsForGetRecordsBytes)
	if err != nil {
		return nil, nil, err
	}

	// getting metric statistics for incoming bytes of a stream
	respOfIncomingBytes, err := cloudWatchClient.GetMetricStatistics(paramsForIncomingBytes)
	if err != nil {
		return nil, nil, err
	}
	return respOfGetRecordsBytes.Datapoints, respOfIncomingBytes.Datapoints, nil
}
