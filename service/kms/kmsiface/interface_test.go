// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package kmsiface_test

import (
	"testing"

	"github.com/gunosy/aws-sdk-go/service/kms"
	"github.com/gunosy/aws-sdk-go/service/kms/kmsiface"
	"github.com/stretchr/testify/assert"
)

func TestInterface(t *testing.T) {
	assert.Implements(t, (*kmsiface.KMSAPI)(nil), kms.New(nil))
}
