//Package codecommit provides gucumber integration tests suppport.
package codecommit

import (
	"github.com/gunosy/aws-sdk-go/internal/features/shared"
	"github.com/gunosy/aws-sdk-go/service/codecommit"
	. "github.com/lsegal/gucumber"
)

var _ = shared.Imported

func init() {
	Before("@codecommit", func() {
		World["client"] = codecommit.New(nil)
	})
}
