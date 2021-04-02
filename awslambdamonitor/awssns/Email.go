// Package awssns has publish to topic .
package awssns

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sns/snsiface"
)

const message string = "Mentioned mail in the link can be used for deleting the kinesis stream which is unsed for past two hours"

// SendEmailToTopicWithURL is used for sending emails to the users who have
//subscribed to the topic of sns.
func SendEmailToTopicWithURL(snsClient snsiface.SNSAPI, url string) error {

	params := &sns.PublishInput{
		Message: aws.String(message + "	\n	" + url),
		Subject:   aws.String("Regarding the unused kinesis stream in Aws"),
		TargetArn: aws.String(os.Getenv("TopicArnOfSNS")),
	}
	_, err := snsClient.Publish(params) //publishing to the topic
	return err
}
