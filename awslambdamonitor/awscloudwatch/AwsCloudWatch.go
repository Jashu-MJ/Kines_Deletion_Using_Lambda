//Package awscloudwatch has monitoring of cloud watch logs.
package awscloudwatch

import "github.com/aws/aws-sdk-go/service/cloudwatch/cloudwatchiface"

//MonitorCloudKinesisStreamLogs is used for monitoring the cloud watch logs of kinesis stream
func MonitorCloudKinesisStreamLogs(cloudWatchClient cloudwatchiface.CloudWatchAPI) (bool, error) {
	datapointsOfGetRecordsBytes, datapointsOfIncomingBytes, err := getDatapoints(cloudWatchClient)

	for i := range datapointsOfGetRecordsBytes { // checks for non zero points
		if *datapointsOfGetRecordsBytes[i].Sum != 0 && *datapointsOfIncomingBytes[i].Sum != 0 {
			return false, err
		}
	}
	return true, err
}
