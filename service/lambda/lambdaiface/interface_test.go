// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package lambdaiface_test

import (
	"testing"

	"github.com/gunosy/aws-sdk-go/service/lambda"
	"github.com/gunosy/aws-sdk-go/service/lambda/lambdaiface"
	"github.com/stretchr/testify/assert"
)

func TestInterface(t *testing.T) {
	assert.Implements(t, (*lambdaiface.LambdaAPI)(nil), lambda.New(nil))
}
