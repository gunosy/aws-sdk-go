// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package devicefarmiface_test

import (
	"testing"

	"github.com/gunosy/aws-sdk-go/service/devicefarm"
	"github.com/gunosy/aws-sdk-go/service/devicefarm/devicefarmiface"
	"github.com/stretchr/testify/assert"
)

func TestInterface(t *testing.T) {
	assert.Implements(t, (*devicefarmiface.DeviceFarmAPI)(nil), devicefarm.New(nil))
}
