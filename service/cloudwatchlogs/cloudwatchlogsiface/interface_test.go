// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package cloudwatchlogsiface_test

import (
	"testing"

	"github.com/gunosy/aws-sdk-go/service/cloudwatchlogs"
	"github.com/gunosy/aws-sdk-go/service/cloudwatchlogs/cloudwatchlogsiface"
	"github.com/stretchr/testify/assert"
)

func TestInterface(t *testing.T) {
	assert.Implements(t, (*cloudwatchlogsiface.CloudWatchLogsAPI)(nil), cloudwatchlogs.New(nil))
}
