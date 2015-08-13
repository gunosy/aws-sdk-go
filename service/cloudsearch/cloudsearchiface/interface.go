// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package cloudsearchiface provides an interface for the Amazon CloudSearch.
package cloudsearchiface

import (
	"github.com/gunosy/aws-sdk-go/aws"
	"github.com/gunosy/aws-sdk-go/service/cloudsearch"
)

// CloudSearchAPI is the interface type for cloudsearch.CloudSearch.
type CloudSearchAPI interface {
	BuildSuggestersRequest(*cloudsearch.BuildSuggestersInput) (*aws.Request, *cloudsearch.BuildSuggestersOutput)

	BuildSuggesters(*cloudsearch.BuildSuggestersInput) (*cloudsearch.BuildSuggestersOutput, error)

	CreateDomainRequest(*cloudsearch.CreateDomainInput) (*aws.Request, *cloudsearch.CreateDomainOutput)

	CreateDomain(*cloudsearch.CreateDomainInput) (*cloudsearch.CreateDomainOutput, error)

	DefineAnalysisSchemeRequest(*cloudsearch.DefineAnalysisSchemeInput) (*aws.Request, *cloudsearch.DefineAnalysisSchemeOutput)

	DefineAnalysisScheme(*cloudsearch.DefineAnalysisSchemeInput) (*cloudsearch.DefineAnalysisSchemeOutput, error)

	DefineExpressionRequest(*cloudsearch.DefineExpressionInput) (*aws.Request, *cloudsearch.DefineExpressionOutput)

	DefineExpression(*cloudsearch.DefineExpressionInput) (*cloudsearch.DefineExpressionOutput, error)

	DefineIndexFieldRequest(*cloudsearch.DefineIndexFieldInput) (*aws.Request, *cloudsearch.DefineIndexFieldOutput)

	DefineIndexField(*cloudsearch.DefineIndexFieldInput) (*cloudsearch.DefineIndexFieldOutput, error)

	DefineSuggesterRequest(*cloudsearch.DefineSuggesterInput) (*aws.Request, *cloudsearch.DefineSuggesterOutput)

	DefineSuggester(*cloudsearch.DefineSuggesterInput) (*cloudsearch.DefineSuggesterOutput, error)

	DeleteAnalysisSchemeRequest(*cloudsearch.DeleteAnalysisSchemeInput) (*aws.Request, *cloudsearch.DeleteAnalysisSchemeOutput)

	DeleteAnalysisScheme(*cloudsearch.DeleteAnalysisSchemeInput) (*cloudsearch.DeleteAnalysisSchemeOutput, error)

	DeleteDomainRequest(*cloudsearch.DeleteDomainInput) (*aws.Request, *cloudsearch.DeleteDomainOutput)

	DeleteDomain(*cloudsearch.DeleteDomainInput) (*cloudsearch.DeleteDomainOutput, error)

	DeleteExpressionRequest(*cloudsearch.DeleteExpressionInput) (*aws.Request, *cloudsearch.DeleteExpressionOutput)

	DeleteExpression(*cloudsearch.DeleteExpressionInput) (*cloudsearch.DeleteExpressionOutput, error)

	DeleteIndexFieldRequest(*cloudsearch.DeleteIndexFieldInput) (*aws.Request, *cloudsearch.DeleteIndexFieldOutput)

	DeleteIndexField(*cloudsearch.DeleteIndexFieldInput) (*cloudsearch.DeleteIndexFieldOutput, error)

	DeleteSuggesterRequest(*cloudsearch.DeleteSuggesterInput) (*aws.Request, *cloudsearch.DeleteSuggesterOutput)

	DeleteSuggester(*cloudsearch.DeleteSuggesterInput) (*cloudsearch.DeleteSuggesterOutput, error)

	DescribeAnalysisSchemesRequest(*cloudsearch.DescribeAnalysisSchemesInput) (*aws.Request, *cloudsearch.DescribeAnalysisSchemesOutput)

	DescribeAnalysisSchemes(*cloudsearch.DescribeAnalysisSchemesInput) (*cloudsearch.DescribeAnalysisSchemesOutput, error)

	DescribeAvailabilityOptionsRequest(*cloudsearch.DescribeAvailabilityOptionsInput) (*aws.Request, *cloudsearch.DescribeAvailabilityOptionsOutput)

	DescribeAvailabilityOptions(*cloudsearch.DescribeAvailabilityOptionsInput) (*cloudsearch.DescribeAvailabilityOptionsOutput, error)

	DescribeDomainsRequest(*cloudsearch.DescribeDomainsInput) (*aws.Request, *cloudsearch.DescribeDomainsOutput)

	DescribeDomains(*cloudsearch.DescribeDomainsInput) (*cloudsearch.DescribeDomainsOutput, error)

	DescribeExpressionsRequest(*cloudsearch.DescribeExpressionsInput) (*aws.Request, *cloudsearch.DescribeExpressionsOutput)

	DescribeExpressions(*cloudsearch.DescribeExpressionsInput) (*cloudsearch.DescribeExpressionsOutput, error)

	DescribeIndexFieldsRequest(*cloudsearch.DescribeIndexFieldsInput) (*aws.Request, *cloudsearch.DescribeIndexFieldsOutput)

	DescribeIndexFields(*cloudsearch.DescribeIndexFieldsInput) (*cloudsearch.DescribeIndexFieldsOutput, error)

	DescribeScalingParametersRequest(*cloudsearch.DescribeScalingParametersInput) (*aws.Request, *cloudsearch.DescribeScalingParametersOutput)

	DescribeScalingParameters(*cloudsearch.DescribeScalingParametersInput) (*cloudsearch.DescribeScalingParametersOutput, error)

	DescribeServiceAccessPoliciesRequest(*cloudsearch.DescribeServiceAccessPoliciesInput) (*aws.Request, *cloudsearch.DescribeServiceAccessPoliciesOutput)

	DescribeServiceAccessPolicies(*cloudsearch.DescribeServiceAccessPoliciesInput) (*cloudsearch.DescribeServiceAccessPoliciesOutput, error)

	DescribeSuggestersRequest(*cloudsearch.DescribeSuggestersInput) (*aws.Request, *cloudsearch.DescribeSuggestersOutput)

	DescribeSuggesters(*cloudsearch.DescribeSuggestersInput) (*cloudsearch.DescribeSuggestersOutput, error)

	IndexDocumentsRequest(*cloudsearch.IndexDocumentsInput) (*aws.Request, *cloudsearch.IndexDocumentsOutput)

	IndexDocuments(*cloudsearch.IndexDocumentsInput) (*cloudsearch.IndexDocumentsOutput, error)

	ListDomainNamesRequest(*cloudsearch.ListDomainNamesInput) (*aws.Request, *cloudsearch.ListDomainNamesOutput)

	ListDomainNames(*cloudsearch.ListDomainNamesInput) (*cloudsearch.ListDomainNamesOutput, error)

	UpdateAvailabilityOptionsRequest(*cloudsearch.UpdateAvailabilityOptionsInput) (*aws.Request, *cloudsearch.UpdateAvailabilityOptionsOutput)

	UpdateAvailabilityOptions(*cloudsearch.UpdateAvailabilityOptionsInput) (*cloudsearch.UpdateAvailabilityOptionsOutput, error)

	UpdateScalingParametersRequest(*cloudsearch.UpdateScalingParametersInput) (*aws.Request, *cloudsearch.UpdateScalingParametersOutput)

	UpdateScalingParameters(*cloudsearch.UpdateScalingParametersInput) (*cloudsearch.UpdateScalingParametersOutput, error)

	UpdateServiceAccessPoliciesRequest(*cloudsearch.UpdateServiceAccessPoliciesInput) (*aws.Request, *cloudsearch.UpdateServiceAccessPoliciesOutput)

	UpdateServiceAccessPolicies(*cloudsearch.UpdateServiceAccessPoliciesInput) (*cloudsearch.UpdateServiceAccessPoliciesOutput, error)
}
