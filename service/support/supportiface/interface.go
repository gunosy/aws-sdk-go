// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package supportiface provides an interface for the AWS Support.
package supportiface

import (
	"github.com/gunosy/aws-sdk-go/aws"
	"github.com/gunosy/aws-sdk-go/service/support"
)

// SupportAPI is the interface type for support.Support.
type SupportAPI interface {
	AddAttachmentsToSetRequest(*support.AddAttachmentsToSetInput) (*aws.Request, *support.AddAttachmentsToSetOutput)

	AddAttachmentsToSet(*support.AddAttachmentsToSetInput) (*support.AddAttachmentsToSetOutput, error)

	AddCommunicationToCaseRequest(*support.AddCommunicationToCaseInput) (*aws.Request, *support.AddCommunicationToCaseOutput)

	AddCommunicationToCase(*support.AddCommunicationToCaseInput) (*support.AddCommunicationToCaseOutput, error)

	CreateCaseRequest(*support.CreateCaseInput) (*aws.Request, *support.CreateCaseOutput)

	CreateCase(*support.CreateCaseInput) (*support.CreateCaseOutput, error)

	DescribeAttachmentRequest(*support.DescribeAttachmentInput) (*aws.Request, *support.DescribeAttachmentOutput)

	DescribeAttachment(*support.DescribeAttachmentInput) (*support.DescribeAttachmentOutput, error)

	DescribeCasesRequest(*support.DescribeCasesInput) (*aws.Request, *support.DescribeCasesOutput)

	DescribeCases(*support.DescribeCasesInput) (*support.DescribeCasesOutput, error)

	DescribeCasesPages(*support.DescribeCasesInput, func(*support.DescribeCasesOutput, bool) bool) error

	DescribeCommunicationsRequest(*support.DescribeCommunicationsInput) (*aws.Request, *support.DescribeCommunicationsOutput)

	DescribeCommunications(*support.DescribeCommunicationsInput) (*support.DescribeCommunicationsOutput, error)

	DescribeCommunicationsPages(*support.DescribeCommunicationsInput, func(*support.DescribeCommunicationsOutput, bool) bool) error

	DescribeServicesRequest(*support.DescribeServicesInput) (*aws.Request, *support.DescribeServicesOutput)

	DescribeServices(*support.DescribeServicesInput) (*support.DescribeServicesOutput, error)

	DescribeSeverityLevelsRequest(*support.DescribeSeverityLevelsInput) (*aws.Request, *support.DescribeSeverityLevelsOutput)

	DescribeSeverityLevels(*support.DescribeSeverityLevelsInput) (*support.DescribeSeverityLevelsOutput, error)

	DescribeTrustedAdvisorCheckRefreshStatusesRequest(*support.DescribeTrustedAdvisorCheckRefreshStatusesInput) (*aws.Request, *support.DescribeTrustedAdvisorCheckRefreshStatusesOutput)

	DescribeTrustedAdvisorCheckRefreshStatuses(*support.DescribeTrustedAdvisorCheckRefreshStatusesInput) (*support.DescribeTrustedAdvisorCheckRefreshStatusesOutput, error)

	DescribeTrustedAdvisorCheckResultRequest(*support.DescribeTrustedAdvisorCheckResultInput) (*aws.Request, *support.DescribeTrustedAdvisorCheckResultOutput)

	DescribeTrustedAdvisorCheckResult(*support.DescribeTrustedAdvisorCheckResultInput) (*support.DescribeTrustedAdvisorCheckResultOutput, error)

	DescribeTrustedAdvisorCheckSummariesRequest(*support.DescribeTrustedAdvisorCheckSummariesInput) (*aws.Request, *support.DescribeTrustedAdvisorCheckSummariesOutput)

	DescribeTrustedAdvisorCheckSummaries(*support.DescribeTrustedAdvisorCheckSummariesInput) (*support.DescribeTrustedAdvisorCheckSummariesOutput, error)

	DescribeTrustedAdvisorChecksRequest(*support.DescribeTrustedAdvisorChecksInput) (*aws.Request, *support.DescribeTrustedAdvisorChecksOutput)

	DescribeTrustedAdvisorChecks(*support.DescribeTrustedAdvisorChecksInput) (*support.DescribeTrustedAdvisorChecksOutput, error)

	RefreshTrustedAdvisorCheckRequest(*support.RefreshTrustedAdvisorCheckInput) (*aws.Request, *support.RefreshTrustedAdvisorCheckOutput)

	RefreshTrustedAdvisorCheck(*support.RefreshTrustedAdvisorCheckInput) (*support.RefreshTrustedAdvisorCheckOutput, error)

	ResolveCaseRequest(*support.ResolveCaseInput) (*aws.Request, *support.ResolveCaseOutput)

	ResolveCase(*support.ResolveCaseInput) (*support.ResolveCaseOutput, error)
}
