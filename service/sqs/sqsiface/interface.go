// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package sqsiface provides an interface for the Amazon Simple Queue Service.
package sqsiface

import (
	"github.com/gunosy/aws-sdk-go/aws"
	"github.com/gunosy/aws-sdk-go/service/sqs"
)

// SQSAPI is the interface type for sqs.SQS.
type SQSAPI interface {
	AddPermissionRequest(*sqs.AddPermissionInput) (*aws.Request, *sqs.AddPermissionOutput)

	AddPermission(*sqs.AddPermissionInput) (*sqs.AddPermissionOutput, error)

	ChangeMessageVisibilityRequest(*sqs.ChangeMessageVisibilityInput) (*aws.Request, *sqs.ChangeMessageVisibilityOutput)

	ChangeMessageVisibility(*sqs.ChangeMessageVisibilityInput) (*sqs.ChangeMessageVisibilityOutput, error)

	ChangeMessageVisibilityBatchRequest(*sqs.ChangeMessageVisibilityBatchInput) (*aws.Request, *sqs.ChangeMessageVisibilityBatchOutput)

	ChangeMessageVisibilityBatch(*sqs.ChangeMessageVisibilityBatchInput) (*sqs.ChangeMessageVisibilityBatchOutput, error)

	CreateQueueRequest(*sqs.CreateQueueInput) (*aws.Request, *sqs.CreateQueueOutput)

	CreateQueue(*sqs.CreateQueueInput) (*sqs.CreateQueueOutput, error)

	DeleteMessageRequest(*sqs.DeleteMessageInput) (*aws.Request, *sqs.DeleteMessageOutput)

	DeleteMessage(*sqs.DeleteMessageInput) (*sqs.DeleteMessageOutput, error)

	DeleteMessageBatchRequest(*sqs.DeleteMessageBatchInput) (*aws.Request, *sqs.DeleteMessageBatchOutput)

	DeleteMessageBatch(*sqs.DeleteMessageBatchInput) (*sqs.DeleteMessageBatchOutput, error)

	DeleteQueueRequest(*sqs.DeleteQueueInput) (*aws.Request, *sqs.DeleteQueueOutput)

	DeleteQueue(*sqs.DeleteQueueInput) (*sqs.DeleteQueueOutput, error)

	GetQueueAttributesRequest(*sqs.GetQueueAttributesInput) (*aws.Request, *sqs.GetQueueAttributesOutput)

	GetQueueAttributes(*sqs.GetQueueAttributesInput) (*sqs.GetQueueAttributesOutput, error)

	GetQueueURLRequest(*sqs.GetQueueURLInput) (*aws.Request, *sqs.GetQueueURLOutput)

	GetQueueURL(*sqs.GetQueueURLInput) (*sqs.GetQueueURLOutput, error)

	ListDeadLetterSourceQueuesRequest(*sqs.ListDeadLetterSourceQueuesInput) (*aws.Request, *sqs.ListDeadLetterSourceQueuesOutput)

	ListDeadLetterSourceQueues(*sqs.ListDeadLetterSourceQueuesInput) (*sqs.ListDeadLetterSourceQueuesOutput, error)

	ListQueuesRequest(*sqs.ListQueuesInput) (*aws.Request, *sqs.ListQueuesOutput)

	ListQueues(*sqs.ListQueuesInput) (*sqs.ListQueuesOutput, error)

	PurgeQueueRequest(*sqs.PurgeQueueInput) (*aws.Request, *sqs.PurgeQueueOutput)

	PurgeQueue(*sqs.PurgeQueueInput) (*sqs.PurgeQueueOutput, error)

	ReceiveMessageRequest(*sqs.ReceiveMessageInput) (*aws.Request, *sqs.ReceiveMessageOutput)

	ReceiveMessage(*sqs.ReceiveMessageInput) (*sqs.ReceiveMessageOutput, error)

	RemovePermissionRequest(*sqs.RemovePermissionInput) (*aws.Request, *sqs.RemovePermissionOutput)

	RemovePermission(*sqs.RemovePermissionInput) (*sqs.RemovePermissionOutput, error)

	SendMessageRequest(*sqs.SendMessageInput) (*aws.Request, *sqs.SendMessageOutput)

	SendMessage(*sqs.SendMessageInput) (*sqs.SendMessageOutput, error)

	SendMessageBatchRequest(*sqs.SendMessageBatchInput) (*aws.Request, *sqs.SendMessageBatchOutput)

	SendMessageBatch(*sqs.SendMessageBatchInput) (*sqs.SendMessageBatchOutput, error)

	SetQueueAttributesRequest(*sqs.SetQueueAttributesInput) (*aws.Request, *sqs.SetQueueAttributesOutput)

	SetQueueAttributes(*sqs.SetQueueAttributesInput) (*sqs.SetQueueAttributesOutput, error)
}
