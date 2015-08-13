// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package machinelearningiface provides an interface for the Amazon Machine Learning.
package machinelearningiface

import (
	"github.com/gunosy/aws-sdk-go/aws"
	"github.com/gunosy/aws-sdk-go/service/machinelearning"
)

// MachineLearningAPI is the interface type for machinelearning.MachineLearning.
type MachineLearningAPI interface {
	CreateBatchPredictionRequest(*machinelearning.CreateBatchPredictionInput) (*aws.Request, *machinelearning.CreateBatchPredictionOutput)

	CreateBatchPrediction(*machinelearning.CreateBatchPredictionInput) (*machinelearning.CreateBatchPredictionOutput, error)

	CreateDataSourceFromRDSRequest(*machinelearning.CreateDataSourceFromRDSInput) (*aws.Request, *machinelearning.CreateDataSourceFromRDSOutput)

	CreateDataSourceFromRDS(*machinelearning.CreateDataSourceFromRDSInput) (*machinelearning.CreateDataSourceFromRDSOutput, error)

	CreateDataSourceFromRedshiftRequest(*machinelearning.CreateDataSourceFromRedshiftInput) (*aws.Request, *machinelearning.CreateDataSourceFromRedshiftOutput)

	CreateDataSourceFromRedshift(*machinelearning.CreateDataSourceFromRedshiftInput) (*machinelearning.CreateDataSourceFromRedshiftOutput, error)

	CreateDataSourceFromS3Request(*machinelearning.CreateDataSourceFromS3Input) (*aws.Request, *machinelearning.CreateDataSourceFromS3Output)

	CreateDataSourceFromS3(*machinelearning.CreateDataSourceFromS3Input) (*machinelearning.CreateDataSourceFromS3Output, error)

	CreateEvaluationRequest(*machinelearning.CreateEvaluationInput) (*aws.Request, *machinelearning.CreateEvaluationOutput)

	CreateEvaluation(*machinelearning.CreateEvaluationInput) (*machinelearning.CreateEvaluationOutput, error)

	CreateMLModelRequest(*machinelearning.CreateMLModelInput) (*aws.Request, *machinelearning.CreateMLModelOutput)

	CreateMLModel(*machinelearning.CreateMLModelInput) (*machinelearning.CreateMLModelOutput, error)

	CreateRealtimeEndpointRequest(*machinelearning.CreateRealtimeEndpointInput) (*aws.Request, *machinelearning.CreateRealtimeEndpointOutput)

	CreateRealtimeEndpoint(*machinelearning.CreateRealtimeEndpointInput) (*machinelearning.CreateRealtimeEndpointOutput, error)

	DeleteBatchPredictionRequest(*machinelearning.DeleteBatchPredictionInput) (*aws.Request, *machinelearning.DeleteBatchPredictionOutput)

	DeleteBatchPrediction(*machinelearning.DeleteBatchPredictionInput) (*machinelearning.DeleteBatchPredictionOutput, error)

	DeleteDataSourceRequest(*machinelearning.DeleteDataSourceInput) (*aws.Request, *machinelearning.DeleteDataSourceOutput)

	DeleteDataSource(*machinelearning.DeleteDataSourceInput) (*machinelearning.DeleteDataSourceOutput, error)

	DeleteEvaluationRequest(*machinelearning.DeleteEvaluationInput) (*aws.Request, *machinelearning.DeleteEvaluationOutput)

	DeleteEvaluation(*machinelearning.DeleteEvaluationInput) (*machinelearning.DeleteEvaluationOutput, error)

	DeleteMLModelRequest(*machinelearning.DeleteMLModelInput) (*aws.Request, *machinelearning.DeleteMLModelOutput)

	DeleteMLModel(*machinelearning.DeleteMLModelInput) (*machinelearning.DeleteMLModelOutput, error)

	DeleteRealtimeEndpointRequest(*machinelearning.DeleteRealtimeEndpointInput) (*aws.Request, *machinelearning.DeleteRealtimeEndpointOutput)

	DeleteRealtimeEndpoint(*machinelearning.DeleteRealtimeEndpointInput) (*machinelearning.DeleteRealtimeEndpointOutput, error)

	DescribeBatchPredictionsRequest(*machinelearning.DescribeBatchPredictionsInput) (*aws.Request, *machinelearning.DescribeBatchPredictionsOutput)

	DescribeBatchPredictions(*machinelearning.DescribeBatchPredictionsInput) (*machinelearning.DescribeBatchPredictionsOutput, error)

	DescribeBatchPredictionsPages(*machinelearning.DescribeBatchPredictionsInput, func(*machinelearning.DescribeBatchPredictionsOutput, bool) bool) error

	DescribeDataSourcesRequest(*machinelearning.DescribeDataSourcesInput) (*aws.Request, *machinelearning.DescribeDataSourcesOutput)

	DescribeDataSources(*machinelearning.DescribeDataSourcesInput) (*machinelearning.DescribeDataSourcesOutput, error)

	DescribeDataSourcesPages(*machinelearning.DescribeDataSourcesInput, func(*machinelearning.DescribeDataSourcesOutput, bool) bool) error

	DescribeEvaluationsRequest(*machinelearning.DescribeEvaluationsInput) (*aws.Request, *machinelearning.DescribeEvaluationsOutput)

	DescribeEvaluations(*machinelearning.DescribeEvaluationsInput) (*machinelearning.DescribeEvaluationsOutput, error)

	DescribeEvaluationsPages(*machinelearning.DescribeEvaluationsInput, func(*machinelearning.DescribeEvaluationsOutput, bool) bool) error

	DescribeMLModelsRequest(*machinelearning.DescribeMLModelsInput) (*aws.Request, *machinelearning.DescribeMLModelsOutput)

	DescribeMLModels(*machinelearning.DescribeMLModelsInput) (*machinelearning.DescribeMLModelsOutput, error)

	DescribeMLModelsPages(*machinelearning.DescribeMLModelsInput, func(*machinelearning.DescribeMLModelsOutput, bool) bool) error

	GetBatchPredictionRequest(*machinelearning.GetBatchPredictionInput) (*aws.Request, *machinelearning.GetBatchPredictionOutput)

	GetBatchPrediction(*machinelearning.GetBatchPredictionInput) (*machinelearning.GetBatchPredictionOutput, error)

	GetDataSourceRequest(*machinelearning.GetDataSourceInput) (*aws.Request, *machinelearning.GetDataSourceOutput)

	GetDataSource(*machinelearning.GetDataSourceInput) (*machinelearning.GetDataSourceOutput, error)

	GetEvaluationRequest(*machinelearning.GetEvaluationInput) (*aws.Request, *machinelearning.GetEvaluationOutput)

	GetEvaluation(*machinelearning.GetEvaluationInput) (*machinelearning.GetEvaluationOutput, error)

	GetMLModelRequest(*machinelearning.GetMLModelInput) (*aws.Request, *machinelearning.GetMLModelOutput)

	GetMLModel(*machinelearning.GetMLModelInput) (*machinelearning.GetMLModelOutput, error)

	PredictRequest(*machinelearning.PredictInput) (*aws.Request, *machinelearning.PredictOutput)

	Predict(*machinelearning.PredictInput) (*machinelearning.PredictOutput, error)

	UpdateBatchPredictionRequest(*machinelearning.UpdateBatchPredictionInput) (*aws.Request, *machinelearning.UpdateBatchPredictionOutput)

	UpdateBatchPrediction(*machinelearning.UpdateBatchPredictionInput) (*machinelearning.UpdateBatchPredictionOutput, error)

	UpdateDataSourceRequest(*machinelearning.UpdateDataSourceInput) (*aws.Request, *machinelearning.UpdateDataSourceOutput)

	UpdateDataSource(*machinelearning.UpdateDataSourceInput) (*machinelearning.UpdateDataSourceOutput, error)

	UpdateEvaluationRequest(*machinelearning.UpdateEvaluationInput) (*aws.Request, *machinelearning.UpdateEvaluationOutput)

	UpdateEvaluation(*machinelearning.UpdateEvaluationInput) (*machinelearning.UpdateEvaluationOutput, error)

	UpdateMLModelRequest(*machinelearning.UpdateMLModelInput) (*aws.Request, *machinelearning.UpdateMLModelOutput)

	UpdateMLModel(*machinelearning.UpdateMLModelInput) (*machinelearning.UpdateMLModelOutput, error)
}
