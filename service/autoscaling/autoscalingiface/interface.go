// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package autoscalingiface provides an interface for the Auto Scaling.
package autoscalingiface

import (
	"github.com/gunosy/aws-sdk-go/aws"
	"github.com/gunosy/aws-sdk-go/service/autoscaling"
)

// AutoScalingAPI is the interface type for autoscaling.AutoScaling.
type AutoScalingAPI interface {
	AttachInstancesRequest(*autoscaling.AttachInstancesInput) (*aws.Request, *autoscaling.AttachInstancesOutput)

	AttachInstances(*autoscaling.AttachInstancesInput) (*autoscaling.AttachInstancesOutput, error)

	AttachLoadBalancersRequest(*autoscaling.AttachLoadBalancersInput) (*aws.Request, *autoscaling.AttachLoadBalancersOutput)

	AttachLoadBalancers(*autoscaling.AttachLoadBalancersInput) (*autoscaling.AttachLoadBalancersOutput, error)

	CompleteLifecycleActionRequest(*autoscaling.CompleteLifecycleActionInput) (*aws.Request, *autoscaling.CompleteLifecycleActionOutput)

	CompleteLifecycleAction(*autoscaling.CompleteLifecycleActionInput) (*autoscaling.CompleteLifecycleActionOutput, error)

	CreateAutoScalingGroupRequest(*autoscaling.CreateAutoScalingGroupInput) (*aws.Request, *autoscaling.CreateAutoScalingGroupOutput)

	CreateAutoScalingGroup(*autoscaling.CreateAutoScalingGroupInput) (*autoscaling.CreateAutoScalingGroupOutput, error)

	CreateLaunchConfigurationRequest(*autoscaling.CreateLaunchConfigurationInput) (*aws.Request, *autoscaling.CreateLaunchConfigurationOutput)

	CreateLaunchConfiguration(*autoscaling.CreateLaunchConfigurationInput) (*autoscaling.CreateLaunchConfigurationOutput, error)

	CreateOrUpdateTagsRequest(*autoscaling.CreateOrUpdateTagsInput) (*aws.Request, *autoscaling.CreateOrUpdateTagsOutput)

	CreateOrUpdateTags(*autoscaling.CreateOrUpdateTagsInput) (*autoscaling.CreateOrUpdateTagsOutput, error)

	DeleteAutoScalingGroupRequest(*autoscaling.DeleteAutoScalingGroupInput) (*aws.Request, *autoscaling.DeleteAutoScalingGroupOutput)

	DeleteAutoScalingGroup(*autoscaling.DeleteAutoScalingGroupInput) (*autoscaling.DeleteAutoScalingGroupOutput, error)

	DeleteLaunchConfigurationRequest(*autoscaling.DeleteLaunchConfigurationInput) (*aws.Request, *autoscaling.DeleteLaunchConfigurationOutput)

	DeleteLaunchConfiguration(*autoscaling.DeleteLaunchConfigurationInput) (*autoscaling.DeleteLaunchConfigurationOutput, error)

	DeleteLifecycleHookRequest(*autoscaling.DeleteLifecycleHookInput) (*aws.Request, *autoscaling.DeleteLifecycleHookOutput)

	DeleteLifecycleHook(*autoscaling.DeleteLifecycleHookInput) (*autoscaling.DeleteLifecycleHookOutput, error)

	DeleteNotificationConfigurationRequest(*autoscaling.DeleteNotificationConfigurationInput) (*aws.Request, *autoscaling.DeleteNotificationConfigurationOutput)

	DeleteNotificationConfiguration(*autoscaling.DeleteNotificationConfigurationInput) (*autoscaling.DeleteNotificationConfigurationOutput, error)

	DeletePolicyRequest(*autoscaling.DeletePolicyInput) (*aws.Request, *autoscaling.DeletePolicyOutput)

	DeletePolicy(*autoscaling.DeletePolicyInput) (*autoscaling.DeletePolicyOutput, error)

	DeleteScheduledActionRequest(*autoscaling.DeleteScheduledActionInput) (*aws.Request, *autoscaling.DeleteScheduledActionOutput)

	DeleteScheduledAction(*autoscaling.DeleteScheduledActionInput) (*autoscaling.DeleteScheduledActionOutput, error)

	DeleteTagsRequest(*autoscaling.DeleteTagsInput) (*aws.Request, *autoscaling.DeleteTagsOutput)

	DeleteTags(*autoscaling.DeleteTagsInput) (*autoscaling.DeleteTagsOutput, error)

	DescribeAccountLimitsRequest(*autoscaling.DescribeAccountLimitsInput) (*aws.Request, *autoscaling.DescribeAccountLimitsOutput)

	DescribeAccountLimits(*autoscaling.DescribeAccountLimitsInput) (*autoscaling.DescribeAccountLimitsOutput, error)

	DescribeAdjustmentTypesRequest(*autoscaling.DescribeAdjustmentTypesInput) (*aws.Request, *autoscaling.DescribeAdjustmentTypesOutput)

	DescribeAdjustmentTypes(*autoscaling.DescribeAdjustmentTypesInput) (*autoscaling.DescribeAdjustmentTypesOutput, error)

	DescribeAutoScalingGroupsRequest(*autoscaling.DescribeAutoScalingGroupsInput) (*aws.Request, *autoscaling.DescribeAutoScalingGroupsOutput)

	DescribeAutoScalingGroups(*autoscaling.DescribeAutoScalingGroupsInput) (*autoscaling.DescribeAutoScalingGroupsOutput, error)

	DescribeAutoScalingGroupsPages(*autoscaling.DescribeAutoScalingGroupsInput, func(*autoscaling.DescribeAutoScalingGroupsOutput, bool) bool) error

	DescribeAutoScalingInstancesRequest(*autoscaling.DescribeAutoScalingInstancesInput) (*aws.Request, *autoscaling.DescribeAutoScalingInstancesOutput)

	DescribeAutoScalingInstances(*autoscaling.DescribeAutoScalingInstancesInput) (*autoscaling.DescribeAutoScalingInstancesOutput, error)

	DescribeAutoScalingInstancesPages(*autoscaling.DescribeAutoScalingInstancesInput, func(*autoscaling.DescribeAutoScalingInstancesOutput, bool) bool) error

	DescribeAutoScalingNotificationTypesRequest(*autoscaling.DescribeAutoScalingNotificationTypesInput) (*aws.Request, *autoscaling.DescribeAutoScalingNotificationTypesOutput)

	DescribeAutoScalingNotificationTypes(*autoscaling.DescribeAutoScalingNotificationTypesInput) (*autoscaling.DescribeAutoScalingNotificationTypesOutput, error)

	DescribeLaunchConfigurationsRequest(*autoscaling.DescribeLaunchConfigurationsInput) (*aws.Request, *autoscaling.DescribeLaunchConfigurationsOutput)

	DescribeLaunchConfigurations(*autoscaling.DescribeLaunchConfigurationsInput) (*autoscaling.DescribeLaunchConfigurationsOutput, error)

	DescribeLaunchConfigurationsPages(*autoscaling.DescribeLaunchConfigurationsInput, func(*autoscaling.DescribeLaunchConfigurationsOutput, bool) bool) error

	DescribeLifecycleHookTypesRequest(*autoscaling.DescribeLifecycleHookTypesInput) (*aws.Request, *autoscaling.DescribeLifecycleHookTypesOutput)

	DescribeLifecycleHookTypes(*autoscaling.DescribeLifecycleHookTypesInput) (*autoscaling.DescribeLifecycleHookTypesOutput, error)

	DescribeLifecycleHooksRequest(*autoscaling.DescribeLifecycleHooksInput) (*aws.Request, *autoscaling.DescribeLifecycleHooksOutput)

	DescribeLifecycleHooks(*autoscaling.DescribeLifecycleHooksInput) (*autoscaling.DescribeLifecycleHooksOutput, error)

	DescribeLoadBalancersRequest(*autoscaling.DescribeLoadBalancersInput) (*aws.Request, *autoscaling.DescribeLoadBalancersOutput)

	DescribeLoadBalancers(*autoscaling.DescribeLoadBalancersInput) (*autoscaling.DescribeLoadBalancersOutput, error)

	DescribeMetricCollectionTypesRequest(*autoscaling.DescribeMetricCollectionTypesInput) (*aws.Request, *autoscaling.DescribeMetricCollectionTypesOutput)

	DescribeMetricCollectionTypes(*autoscaling.DescribeMetricCollectionTypesInput) (*autoscaling.DescribeMetricCollectionTypesOutput, error)

	DescribeNotificationConfigurationsRequest(*autoscaling.DescribeNotificationConfigurationsInput) (*aws.Request, *autoscaling.DescribeNotificationConfigurationsOutput)

	DescribeNotificationConfigurations(*autoscaling.DescribeNotificationConfigurationsInput) (*autoscaling.DescribeNotificationConfigurationsOutput, error)

	DescribeNotificationConfigurationsPages(*autoscaling.DescribeNotificationConfigurationsInput, func(*autoscaling.DescribeNotificationConfigurationsOutput, bool) bool) error

	DescribePoliciesRequest(*autoscaling.DescribePoliciesInput) (*aws.Request, *autoscaling.DescribePoliciesOutput)

	DescribePolicies(*autoscaling.DescribePoliciesInput) (*autoscaling.DescribePoliciesOutput, error)

	DescribePoliciesPages(*autoscaling.DescribePoliciesInput, func(*autoscaling.DescribePoliciesOutput, bool) bool) error

	DescribeScalingActivitiesRequest(*autoscaling.DescribeScalingActivitiesInput) (*aws.Request, *autoscaling.DescribeScalingActivitiesOutput)

	DescribeScalingActivities(*autoscaling.DescribeScalingActivitiesInput) (*autoscaling.DescribeScalingActivitiesOutput, error)

	DescribeScalingActivitiesPages(*autoscaling.DescribeScalingActivitiesInput, func(*autoscaling.DescribeScalingActivitiesOutput, bool) bool) error

	DescribeScalingProcessTypesRequest(*autoscaling.DescribeScalingProcessTypesInput) (*aws.Request, *autoscaling.DescribeScalingProcessTypesOutput)

	DescribeScalingProcessTypes(*autoscaling.DescribeScalingProcessTypesInput) (*autoscaling.DescribeScalingProcessTypesOutput, error)

	DescribeScheduledActionsRequest(*autoscaling.DescribeScheduledActionsInput) (*aws.Request, *autoscaling.DescribeScheduledActionsOutput)

	DescribeScheduledActions(*autoscaling.DescribeScheduledActionsInput) (*autoscaling.DescribeScheduledActionsOutput, error)

	DescribeScheduledActionsPages(*autoscaling.DescribeScheduledActionsInput, func(*autoscaling.DescribeScheduledActionsOutput, bool) bool) error

	DescribeTagsRequest(*autoscaling.DescribeTagsInput) (*aws.Request, *autoscaling.DescribeTagsOutput)

	DescribeTags(*autoscaling.DescribeTagsInput) (*autoscaling.DescribeTagsOutput, error)

	DescribeTagsPages(*autoscaling.DescribeTagsInput, func(*autoscaling.DescribeTagsOutput, bool) bool) error

	DescribeTerminationPolicyTypesRequest(*autoscaling.DescribeTerminationPolicyTypesInput) (*aws.Request, *autoscaling.DescribeTerminationPolicyTypesOutput)

	DescribeTerminationPolicyTypes(*autoscaling.DescribeTerminationPolicyTypesInput) (*autoscaling.DescribeTerminationPolicyTypesOutput, error)

	DetachInstancesRequest(*autoscaling.DetachInstancesInput) (*aws.Request, *autoscaling.DetachInstancesOutput)

	DetachInstances(*autoscaling.DetachInstancesInput) (*autoscaling.DetachInstancesOutput, error)

	DetachLoadBalancersRequest(*autoscaling.DetachLoadBalancersInput) (*aws.Request, *autoscaling.DetachLoadBalancersOutput)

	DetachLoadBalancers(*autoscaling.DetachLoadBalancersInput) (*autoscaling.DetachLoadBalancersOutput, error)

	DisableMetricsCollectionRequest(*autoscaling.DisableMetricsCollectionInput) (*aws.Request, *autoscaling.DisableMetricsCollectionOutput)

	DisableMetricsCollection(*autoscaling.DisableMetricsCollectionInput) (*autoscaling.DisableMetricsCollectionOutput, error)

	EnableMetricsCollectionRequest(*autoscaling.EnableMetricsCollectionInput) (*aws.Request, *autoscaling.EnableMetricsCollectionOutput)

	EnableMetricsCollection(*autoscaling.EnableMetricsCollectionInput) (*autoscaling.EnableMetricsCollectionOutput, error)

	EnterStandbyRequest(*autoscaling.EnterStandbyInput) (*aws.Request, *autoscaling.EnterStandbyOutput)

	EnterStandby(*autoscaling.EnterStandbyInput) (*autoscaling.EnterStandbyOutput, error)

	ExecutePolicyRequest(*autoscaling.ExecutePolicyInput) (*aws.Request, *autoscaling.ExecutePolicyOutput)

	ExecutePolicy(*autoscaling.ExecutePolicyInput) (*autoscaling.ExecutePolicyOutput, error)

	ExitStandbyRequest(*autoscaling.ExitStandbyInput) (*aws.Request, *autoscaling.ExitStandbyOutput)

	ExitStandby(*autoscaling.ExitStandbyInput) (*autoscaling.ExitStandbyOutput, error)

	PutLifecycleHookRequest(*autoscaling.PutLifecycleHookInput) (*aws.Request, *autoscaling.PutLifecycleHookOutput)

	PutLifecycleHook(*autoscaling.PutLifecycleHookInput) (*autoscaling.PutLifecycleHookOutput, error)

	PutNotificationConfigurationRequest(*autoscaling.PutNotificationConfigurationInput) (*aws.Request, *autoscaling.PutNotificationConfigurationOutput)

	PutNotificationConfiguration(*autoscaling.PutNotificationConfigurationInput) (*autoscaling.PutNotificationConfigurationOutput, error)

	PutScalingPolicyRequest(*autoscaling.PutScalingPolicyInput) (*aws.Request, *autoscaling.PutScalingPolicyOutput)

	PutScalingPolicy(*autoscaling.PutScalingPolicyInput) (*autoscaling.PutScalingPolicyOutput, error)

	PutScheduledUpdateGroupActionRequest(*autoscaling.PutScheduledUpdateGroupActionInput) (*aws.Request, *autoscaling.PutScheduledUpdateGroupActionOutput)

	PutScheduledUpdateGroupAction(*autoscaling.PutScheduledUpdateGroupActionInput) (*autoscaling.PutScheduledUpdateGroupActionOutput, error)

	RecordLifecycleActionHeartbeatRequest(*autoscaling.RecordLifecycleActionHeartbeatInput) (*aws.Request, *autoscaling.RecordLifecycleActionHeartbeatOutput)

	RecordLifecycleActionHeartbeat(*autoscaling.RecordLifecycleActionHeartbeatInput) (*autoscaling.RecordLifecycleActionHeartbeatOutput, error)

	ResumeProcessesRequest(*autoscaling.ScalingProcessQuery) (*aws.Request, *autoscaling.ResumeProcessesOutput)

	ResumeProcesses(*autoscaling.ScalingProcessQuery) (*autoscaling.ResumeProcessesOutput, error)

	SetDesiredCapacityRequest(*autoscaling.SetDesiredCapacityInput) (*aws.Request, *autoscaling.SetDesiredCapacityOutput)

	SetDesiredCapacity(*autoscaling.SetDesiredCapacityInput) (*autoscaling.SetDesiredCapacityOutput, error)

	SetInstanceHealthRequest(*autoscaling.SetInstanceHealthInput) (*aws.Request, *autoscaling.SetInstanceHealthOutput)

	SetInstanceHealth(*autoscaling.SetInstanceHealthInput) (*autoscaling.SetInstanceHealthOutput, error)

	SuspendProcessesRequest(*autoscaling.ScalingProcessQuery) (*aws.Request, *autoscaling.SuspendProcessesOutput)

	SuspendProcesses(*autoscaling.ScalingProcessQuery) (*autoscaling.SuspendProcessesOutput, error)

	TerminateInstanceInAutoScalingGroupRequest(*autoscaling.TerminateInstanceInAutoScalingGroupInput) (*aws.Request, *autoscaling.TerminateInstanceInAutoScalingGroupOutput)

	TerminateInstanceInAutoScalingGroup(*autoscaling.TerminateInstanceInAutoScalingGroupInput) (*autoscaling.TerminateInstanceInAutoScalingGroupOutput, error)

	UpdateAutoScalingGroupRequest(*autoscaling.UpdateAutoScalingGroupInput) (*aws.Request, *autoscaling.UpdateAutoScalingGroupOutput)

	UpdateAutoScalingGroup(*autoscaling.UpdateAutoScalingGroupInput) (*autoscaling.UpdateAutoScalingGroupOutput, error)
}
