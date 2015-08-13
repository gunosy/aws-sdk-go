//Package elasticbeanstalk provides gucumber integration tests suppport.
package elasticbeanstalk

import (
	"github.com/gunosy/aws-sdk-go/internal/features/shared"
	"github.com/gunosy/aws-sdk-go/service/elasticbeanstalk"
	. "github.com/lsegal/gucumber"
)

var _ = shared.Imported

func init() {
	Before("@elasticbeanstalk", func() {
		World["client"] = elasticbeanstalk.New(nil)
	})
}
