//Package glacier provides gucumber integration tests suppport.
package glacier

import (
	"github.com/gunosy/aws-sdk-go/internal/features/shared"
	"github.com/gunosy/aws-sdk-go/service/glacier"
	. "github.com/lsegal/gucumber"
)

var _ = shared.Imported

func init() {
	Before("@glacier", func() {
		World["client"] = glacier.New(nil)
	})
}
