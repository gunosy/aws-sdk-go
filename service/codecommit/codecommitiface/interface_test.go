// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package codecommitiface_test

import (
	"testing"

	"github.com/gunosy/aws-sdk-go/service/codecommit"
	"github.com/gunosy/aws-sdk-go/service/codecommit/codecommitiface"
	"github.com/stretchr/testify/assert"
)

func TestInterface(t *testing.T) {
	assert.Implements(t, (*codecommitiface.CodeCommitAPI)(nil), codecommit.New(nil))
}
