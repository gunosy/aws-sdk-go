// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package codedeployiface_test

import (
	"testing"

	"github.com/gunosy/aws-sdk-go/service/codedeploy"
	"github.com/gunosy/aws-sdk-go/service/codedeploy/codedeployiface"
	"github.com/stretchr/testify/assert"
)

func TestInterface(t *testing.T) {
	assert.Implements(t, (*codedeployiface.CodeDeployAPI)(nil), codedeploy.New(nil))
}
