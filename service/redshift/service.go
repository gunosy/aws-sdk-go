// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package redshift

import (
	"github.com/gunosy/aws-sdk-go/aws"
	"github.com/gunosy/aws-sdk-go/internal/protocol/query"
	"github.com/gunosy/aws-sdk-go/internal/signer/v4"
)

// Overview  This is an interface reference for Amazon Redshift. It contains
// documentation for one of the programming or command line interfaces you can
// use to manage Amazon Redshift clusters. Note that Amazon Redshift is asynchronous,
// which means that some interfaces may require techniques, such as polling
// or asynchronous callback handlers, to determine when a command has been applied.
// In this reference, the parameter descriptions indicate whether a change is
// applied immediately, on the next instance reboot, or during the next maintenance
// window. For a summary of the Amazon Redshift cluster management interfaces,
// go to Using the Amazon Redshift Management Interfaces  (http://docs.aws.amazon.com/redshift/latest/mgmt/using-aws-sdk.html).
//
//  Amazon Redshift manages all the work of setting up, operating, and scaling
// a data warehouse: provisioning capacity, monitoring and backing up the cluster,
// and applying patches and upgrades to the Amazon Redshift engine. You can
// focus on using your data to acquire new insights for your business and customers.
//
// If you are a first-time user of Amazon Redshift, we recommend that you begin
// by reading the The Amazon Redshift Getting Started Guide (http://docs.aws.amazon.com/redshift/latest/gsg/getting-started.html)
//
// If you are a database developer, the Amazon Redshift Database Developer
// Guide (http://docs.aws.amazon.com/redshift/latest/dg/welcome.html) explains
// how to design, build, query, and maintain the databases that make up your
// data warehouse.
type Redshift struct {
	*aws.Service
}

// Used for custom service initialization logic
var initService func(*aws.Service)

// Used for custom request initialization logic
var initRequest func(*aws.Request)

// New returns a new Redshift client.
func New(config *aws.Config) *Redshift {
	service := &aws.Service{
		Config:      aws.DefaultConfig.Merge(config),
		ServiceName: "redshift",
		APIVersion:  "2012-12-01",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(query.Build)
	service.Handlers.Unmarshal.PushBack(query.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(query.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(query.UnmarshalError)

	// Run custom service initialization if present
	if initService != nil {
		initService(service)
	}

	return &Redshift{service}
}

// newRequest creates a new request for a Redshift operation and runs any
// custom request initialization.
func (c *Redshift) newRequest(op *aws.Operation, params, data interface{}) *aws.Request {
	req := aws.NewRequest(c.Service, op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
