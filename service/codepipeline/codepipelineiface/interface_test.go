// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package codepipelineiface_test

import (
	"testing"

	"github.com/gunosy/aws-sdk-go/service/codepipeline"
	"github.com/gunosy/aws-sdk-go/service/codepipeline/codepipelineiface"
	"github.com/stretchr/testify/assert"
)

func TestInterface(t *testing.T) {
	assert.Implements(t, (*codepipelineiface.CodePipelineAPI)(nil), codepipeline.New(nil))
}
