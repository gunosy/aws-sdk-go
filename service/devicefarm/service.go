// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package devicefarm

import (
	"github.com/gunosy/aws-sdk-go/aws"
	"github.com/gunosy/aws-sdk-go/internal/protocol/jsonrpc"
	"github.com/gunosy/aws-sdk-go/internal/signer/v4"
)

// AWS Device Farm is a service that enables mobile app developers to test Android,
// iOS, and Fire OS apps on physical phones, tablets, and other devices in the
// cloud.
type DeviceFarm struct {
	*aws.Service
}

// Used for custom service initialization logic
var initService func(*aws.Service)

// Used for custom request initialization logic
var initRequest func(*aws.Request)

// New returns a new DeviceFarm client.
func New(config *aws.Config) *DeviceFarm {
	service := &aws.Service{
		Config:       aws.DefaultConfig.Merge(config),
		ServiceName:  "devicefarm",
		APIVersion:   "2015-06-23",
		JSONVersion:  "1.1",
		TargetPrefix: "DeviceFarm_20150623",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(jsonrpc.Build)
	service.Handlers.Unmarshal.PushBack(jsonrpc.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(jsonrpc.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(jsonrpc.UnmarshalError)

	// Run custom service initialization if present
	if initService != nil {
		initService(service)
	}

	return &DeviceFarm{service}
}

// newRequest creates a new request for a DeviceFarm operation and runs any
// custom request initialization.
func (c *DeviceFarm) newRequest(op *aws.Operation, params, data interface{}) *aws.Request {
	req := aws.NewRequest(c.Service, op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
