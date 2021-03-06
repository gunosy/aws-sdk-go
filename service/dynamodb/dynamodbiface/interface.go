// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package dynamodbiface provides an interface for the Amazon DynamoDB.
package dynamodbiface

import (
	"github.com/gunosy/aws-sdk-go/aws"
	"github.com/gunosy/aws-sdk-go/service/dynamodb"
)

// DynamoDBAPI is the interface type for dynamodb.DynamoDB.
type DynamoDBAPI interface {
	BatchGetItemRequest(*dynamodb.BatchGetItemInput) (*aws.Request, *dynamodb.BatchGetItemOutput)

	BatchGetItem(*dynamodb.BatchGetItemInput) (*dynamodb.BatchGetItemOutput, error)

	BatchGetItemPages(*dynamodb.BatchGetItemInput, func(*dynamodb.BatchGetItemOutput, bool) bool) error

	BatchWriteItemRequest(*dynamodb.BatchWriteItemInput) (*aws.Request, *dynamodb.BatchWriteItemOutput)

	BatchWriteItem(*dynamodb.BatchWriteItemInput) (*dynamodb.BatchWriteItemOutput, error)

	CreateTableRequest(*dynamodb.CreateTableInput) (*aws.Request, *dynamodb.CreateTableOutput)

	CreateTable(*dynamodb.CreateTableInput) (*dynamodb.CreateTableOutput, error)

	DeleteItemRequest(*dynamodb.DeleteItemInput) (*aws.Request, *dynamodb.DeleteItemOutput)

	DeleteItem(*dynamodb.DeleteItemInput) (*dynamodb.DeleteItemOutput, error)

	DeleteTableRequest(*dynamodb.DeleteTableInput) (*aws.Request, *dynamodb.DeleteTableOutput)

	DeleteTable(*dynamodb.DeleteTableInput) (*dynamodb.DeleteTableOutput, error)

	DescribeTableRequest(*dynamodb.DescribeTableInput) (*aws.Request, *dynamodb.DescribeTableOutput)

	DescribeTable(*dynamodb.DescribeTableInput) (*dynamodb.DescribeTableOutput, error)

	GetItemRequest(*dynamodb.GetItemInput) (*aws.Request, *dynamodb.GetItemOutput)

	GetItem(*dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error)

	ListTablesRequest(*dynamodb.ListTablesInput) (*aws.Request, *dynamodb.ListTablesOutput)

	ListTables(*dynamodb.ListTablesInput) (*dynamodb.ListTablesOutput, error)

	ListTablesPages(*dynamodb.ListTablesInput, func(*dynamodb.ListTablesOutput, bool) bool) error

	PutItemRequest(*dynamodb.PutItemInput) (*aws.Request, *dynamodb.PutItemOutput)

	PutItem(*dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error)

	QueryRequest(*dynamodb.QueryInput) (*aws.Request, *dynamodb.QueryOutput)

	Query(*dynamodb.QueryInput) (*dynamodb.QueryOutput, error)

	QueryPages(*dynamodb.QueryInput, func(*dynamodb.QueryOutput, bool) bool) error

	ScanRequest(*dynamodb.ScanInput) (*aws.Request, *dynamodb.ScanOutput)

	Scan(*dynamodb.ScanInput) (*dynamodb.ScanOutput, error)

	ScanPages(*dynamodb.ScanInput, func(*dynamodb.ScanOutput, bool) bool) error

	UpdateItemRequest(*dynamodb.UpdateItemInput) (*aws.Request, *dynamodb.UpdateItemOutput)

	UpdateItem(*dynamodb.UpdateItemInput) (*dynamodb.UpdateItemOutput, error)

	UpdateTableRequest(*dynamodb.UpdateTableInput) (*aws.Request, *dynamodb.UpdateTableOutput)

	UpdateTable(*dynamodb.UpdateTableInput) (*dynamodb.UpdateTableOutput, error)
}
