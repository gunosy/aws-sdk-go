// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package datapipelineiface_test

import (
	"testing"

	"github.com/gunosy/aws-sdk-go/service/datapipeline"
	"github.com/gunosy/aws-sdk-go/service/datapipeline/datapipelineiface"
	"github.com/stretchr/testify/assert"
)

func TestInterface(t *testing.T) {
	assert.Implements(t, (*datapipelineiface.DataPipelineAPI)(nil), datapipeline.New(nil))
}
