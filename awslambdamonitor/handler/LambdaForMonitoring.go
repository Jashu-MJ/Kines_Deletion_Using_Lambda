//Package handler has lambda function for monitoring cloud watch logs
package handler

import (
	"awslambdamonitor/awsapigateway"
	"awslambdamonitor/awscloudwatch"
	"awslambdamonitor/awskinesis"
	"awslambdamonitor/awslambda"
	"awslambdamonitor/awssns"
	"errors"
)

//Handler function is used for deleting the kinesis stream, this function is triggered by an api request.
//ctx context.Context, cld events.CloudWatchEvent
func Handler() error {

	kinesisClient, err := awskinesis.KinesisClientCreation()
	if err != nil {
		return err
	}
	isStreamExists, err := awskinesis.IsStreamExists(kinesisClient)
	if err != nil {
		return err
	}
	if isStreamExists != true {
		return errors.New("Stream doesn't exists")
	}

	//cloud watch client creation
	cloudWatchClient, err := awscloudwatch.CreateCloudWatchClient()
	if err != nil {
		return err
	}

	//checking the cloud watch for kinesis stream activity
	emailToBeSent, err := awscloudwatch.MonitorCloudKinesisStreamLogs(cloudWatchClient)
	if err != nil {
		return err
	}

	if !emailToBeSent {
		return err //if toBeSent is false then function will exit with error as nil
	}

	// or else starts generates a link and sends an email.
	//creating api gateway client
	apiGatewayClient, err := awsapigateway.CreateAPIGatewayClient()
	if err != nil {
		return err
	}
	// lambda client creation
	lambdaClient, err := awslambda.LambdaClientCreation()
	if err != nil {
		return err
	}
	link, err := awsapigateway.CreationOfURL(apiGatewayClient, lambdaClient)
	if err != nil {
		return err
	}

	//Creating sns client
	snsClient, err := awssns.SnsClientCreation()
	if err != nil {
		return err
	}

	// sends email with link to users who have subscribed to the topic
	awssns.SendEmailToTopicWithURL(snsClient, link)

	return err
}
