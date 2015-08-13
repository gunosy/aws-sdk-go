//Package cloudsearch provides gucumber integration tests suppport.
package cloudsearch

import (
	"github.com/gunosy/aws-sdk-go/internal/features/shared"
	"github.com/gunosy/aws-sdk-go/service/cloudsearch"
	. "github.com/lsegal/gucumber"
)

var _ = shared.Imported

func init() {
	Before("@cloudsearch", func() {
		World["client"] = cloudsearch.New(nil)
	})
}
