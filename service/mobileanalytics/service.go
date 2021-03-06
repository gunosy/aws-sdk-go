// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package mobileanalytics

import (
	"github.com/gunosy/aws-sdk-go/aws"
	"github.com/gunosy/aws-sdk-go/internal/protocol/restjson"
	"github.com/gunosy/aws-sdk-go/internal/signer/v4"
)

// Amazon Mobile Analytics is a service for collecting, visualizing, and understanding
// app usage data at scale.
type MobileAnalytics struct {
	*aws.Service
}

// Used for custom service initialization logic
var initService func(*aws.Service)

// Used for custom request initialization logic
var initRequest func(*aws.Request)

// New returns a new MobileAnalytics client.
func New(config *aws.Config) *MobileAnalytics {
	service := &aws.Service{
		Config:      aws.DefaultConfig.Merge(config),
		ServiceName: "mobileanalytics",
		APIVersion:  "2014-06-05",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(restjson.Build)
	service.Handlers.Unmarshal.PushBack(restjson.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(restjson.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(restjson.UnmarshalError)

	// Run custom service initialization if present
	if initService != nil {
		initService(service)
	}

	return &MobileAnalytics{service}
}

// newRequest creates a new request for a MobileAnalytics operation and runs any
// custom request initialization.
func (c *MobileAnalytics) newRequest(op *aws.Operation, params, data interface{}) *aws.Request {
	req := aws.NewRequest(c.Service, op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
