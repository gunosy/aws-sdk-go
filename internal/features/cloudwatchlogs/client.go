//Package cloudwatchlogs provides gucumber integration tests suppport.
package cloudwatchlogs

import (
	"github.com/gunosy/aws-sdk-go/internal/features/shared"
	"github.com/gunosy/aws-sdk-go/service/cloudwatchlogs"
	. "github.com/lsegal/gucumber"
)

var _ = shared.Imported

func init() {
	Before("@cloudwatchlogs", func() {
		World["client"] = cloudwatchlogs.New(nil)
	})
}
