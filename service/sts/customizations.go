package sts

import "github.com/gunosy/aws-sdk-go/aws"

func init() {
	initRequest = func(r *aws.Request) {
		switch r.Operation.Name {
		case opAssumeRoleWithSAML, opAssumeRoleWithWebIdentity:
			r.Handlers.Sign.Clear() // these operations are unsigned
		}
	}
}
