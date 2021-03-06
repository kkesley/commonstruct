package es

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws/signer/v4"
	"github.com/olivere/elastic"
	"github.com/sha1sum/aws_signing_client"
)

//New generate new client
func New(url string, region string) (*elastic.Client, error) {
	config := aws.Config{
		Region: aws.String(region),
	}
	sess := session.Must(session.NewSession(&config))
	return NewWithSession(url, region, sess)
}

//NewWithCreds new client with existing credentials
func NewWithCreds(url string, region string, creds *credentials.Credentials) (*elastic.Client, error) {
	signer := v4.NewSigner(creds)
	awsClient, err := aws_signing_client.New(signer, nil, "es", region)
	if err != nil {
		return nil, err
	}
	return elastic.NewClient(
		elastic.SetURL(url),
		elastic.SetScheme("https"),
		elastic.SetHttpClient(awsClient),
		elastic.SetSniff(false), // See note below
	)
}

//NewWithSession generate new client with existing session
func NewWithSession(url string, region string, sess *session.Session) (*elastic.Client, error) {
	return NewWithCreds(url, region, sess.Config.Credentials)

}
