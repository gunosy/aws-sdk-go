//Package dynamodbstreams provides gucumber integration tests suppport.
package dynamodbstreams

import (
	"github.com/gunosy/aws-sdk-go/internal/features/shared"
	"github.com/gunosy/aws-sdk-go/service/dynamodbstreams"
	. "github.com/lsegal/gucumber"
)

var _ = shared.Imported

func init() {
	Before("@dynamodbstreams", func() {
		World["client"] = dynamodbstreams.New(nil)
	})
}
