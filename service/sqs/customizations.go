package sqs

import "github.com/gunosy/aws-sdk-go/aws"

func init() {
	initRequest = func(r *aws.Request) {
		setupChecksumValidation(r)
	}
}
