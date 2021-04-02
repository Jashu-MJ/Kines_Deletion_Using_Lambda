// Package awssession has configuration details of aws
package awssession

import "os"

// ConfigDetails consists of details required for establishing a session with aws and
// it can be used for setting config details and for retrieving it also
var ConfigDetails struct{ Region, AccessKeyID, SecretAccessKey, AwsSessionToken string }

// SetConfigDetails function is used for setting the configuration detailsin the form
func setConfigDetails() {
	ConfigDetails.Region = os.Getenv("AwsRegion")
	ConfigDetails.AccessKeyID = os.Getenv("AwsAccessKeyID")
	ConfigDetails.SecretAccessKey = os.Getenv("AwsSecretAccessKey")
	ConfigDetails.AwsSessionToken = os.Getenv("AwsSessionToken")
}

// GetConfigDetails function is used for retrieving the the aws configuration details
//which are availabe in ConfigDetails object
func getConfigDetails() (string, string, string, string) {
	return ConfigDetails.Region, ConfigDetails.AccessKeyID, ConfigDetails.SecretAccessKey, ConfigDetails.AwsSessionToken
}
