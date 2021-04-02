// Package awssession has aws session creation.
package awssession

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

// CreateSession function is passed with the  config details in form of configDetails
// Structure and which are used to establish a session with Aws server and returns a
// reference to the session object(*session.Session)
func CreateSession() (*session.Session, error) {
	setConfigDetails()
	Region, AccessKeyID, SecretAccessKey, AwsSessionToken := getConfigDetails()
	AwsSession, err := session.NewSession(&aws.Config{
		Region: aws.String(Region),
		Credentials: credentials.NewStaticCredentials(
			AccessKeyID,
			SecretAccessKey,
			AwsSessionToken,
		),
	})
	return AwsSession, err
}
