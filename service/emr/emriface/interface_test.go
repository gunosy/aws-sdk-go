// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package emriface_test

import (
	"testing"

	"github.com/gunosy/aws-sdk-go/service/emr"
	"github.com/gunosy/aws-sdk-go/service/emr/emriface"
	"github.com/stretchr/testify/assert"
)

func TestInterface(t *testing.T) {
	assert.Implements(t, (*emriface.EMRAPI)(nil), emr.New(nil))
}
