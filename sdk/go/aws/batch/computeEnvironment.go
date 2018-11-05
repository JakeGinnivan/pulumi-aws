// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package batch

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Creates a AWS Batch compute environment. Compute environments contain the Amazon ECS container instances that are used to run containerized batch jobs.
// 
// For information about AWS Batch, see [What is AWS Batch?][1] .
// For information about compute environment, see [Compute Environments][2] .
// 
// ~> **Note:** To prevent a race condition during environment deletion, make sure to set `depends_on` to the related `aws_iam_role_policy_attachment`;
//    otherwise, the policy may be destroyed too soon and the compute environment will then get stuck in the `DELETING` state, see [Troubleshooting AWS Batch][3] .
type ComputeEnvironment struct {
	s *pulumi.ResourceState
}

// NewComputeEnvironment registers a new resource with the given unique name, arguments, and options.
func NewComputeEnvironment(ctx *pulumi.Context,
	name string, args *ComputeEnvironmentArgs, opts ...pulumi.ResourceOpt) (*ComputeEnvironment, error) {
	if args == nil || args.ComputeEnvironmentName == nil {
		return nil, errors.New("missing required argument 'ComputeEnvironmentName'")
	}
	if args == nil || args.ServiceRole == nil {
		return nil, errors.New("missing required argument 'ServiceRole'")
	}
	if args == nil || args.Type == nil {
		return nil, errors.New("missing required argument 'Type'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["computeEnvironmentName"] = nil
		inputs["computeResources"] = nil
		inputs["serviceRole"] = nil
		inputs["state"] = nil
		inputs["type"] = nil
	} else {
		inputs["computeEnvironmentName"] = args.ComputeEnvironmentName
		inputs["computeResources"] = args.ComputeResources
		inputs["serviceRole"] = args.ServiceRole
		inputs["state"] = args.State
		inputs["type"] = args.Type
	}
	inputs["arn"] = nil
	inputs["eccClusterArn"] = nil
	inputs["ecsClusterArn"] = nil
	inputs["status"] = nil
	inputs["statusReason"] = nil
	s, err := ctx.RegisterResource("aws:batch/computeEnvironment:ComputeEnvironment", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ComputeEnvironment{s: s}, nil
}

// GetComputeEnvironment gets an existing ComputeEnvironment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetComputeEnvironment(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ComputeEnvironmentState, opts ...pulumi.ResourceOpt) (*ComputeEnvironment, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["arn"] = state.Arn
		inputs["computeEnvironmentName"] = state.ComputeEnvironmentName
		inputs["computeResources"] = state.ComputeResources
		inputs["eccClusterArn"] = state.EccClusterArn
		inputs["ecsClusterArn"] = state.EcsClusterArn
		inputs["serviceRole"] = state.ServiceRole
		inputs["state"] = state.State
		inputs["status"] = state.Status
		inputs["statusReason"] = state.StatusReason
		inputs["type"] = state.Type
	}
	s, err := ctx.ReadResource("aws:batch/computeEnvironment:ComputeEnvironment", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ComputeEnvironment{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *ComputeEnvironment) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *ComputeEnvironment) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The Amazon Resource Name (ARN) of the compute environment.
func (r *ComputeEnvironment) Arn() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["arn"])
}

// The name for your compute environment. Up to 128 letters (uppercase and lowercase), numbers, and underscores are allowed.
func (r *ComputeEnvironment) ComputeEnvironmentName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["computeEnvironmentName"])
}

// Details of the compute resources managed by the compute environment. This parameter is required for managed compute environments. See details below.
func (r *ComputeEnvironment) ComputeResources() *pulumi.Output {
	return r.s.State["computeResources"]
}

func (r *ComputeEnvironment) EccClusterArn() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["eccClusterArn"])
}

// The Amazon Resource Name (ARN) of the underlying Amazon ECS cluster used by the compute environment.
func (r *ComputeEnvironment) EcsClusterArn() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["ecsClusterArn"])
}

// The full Amazon Resource Name (ARN) of the IAM role that allows AWS Batch to make calls to other AWS services on your behalf.
func (r *ComputeEnvironment) ServiceRole() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["serviceRole"])
}

// The state of the compute environment. If the state is `ENABLED`, then the compute environment accepts jobs from a queue and can scale out automatically based on queues. Valid items are `ENABLED` or `DISABLED`. Defaults to `ENABLED`.
func (r *ComputeEnvironment) State() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["state"])
}

// The current status of the compute environment (for example, CREATING or VALID).
func (r *ComputeEnvironment) Status() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["status"])
}

// A short, human-readable string to provide additional details about the current status of the compute environment.
func (r *ComputeEnvironment) StatusReason() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["statusReason"])
}

// The type of compute environment. Valid items are `EC2` or `SPOT`.
func (r *ComputeEnvironment) Type() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["type"])
}

// Input properties used for looking up and filtering ComputeEnvironment resources.
type ComputeEnvironmentState struct {
	// The Amazon Resource Name (ARN) of the compute environment.
	Arn interface{}
	// The name for your compute environment. Up to 128 letters (uppercase and lowercase), numbers, and underscores are allowed.
	ComputeEnvironmentName interface{}
	// Details of the compute resources managed by the compute environment. This parameter is required for managed compute environments. See details below.
	ComputeResources interface{}
	EccClusterArn interface{}
	// The Amazon Resource Name (ARN) of the underlying Amazon ECS cluster used by the compute environment.
	EcsClusterArn interface{}
	// The full Amazon Resource Name (ARN) of the IAM role that allows AWS Batch to make calls to other AWS services on your behalf.
	ServiceRole interface{}
	// The state of the compute environment. If the state is `ENABLED`, then the compute environment accepts jobs from a queue and can scale out automatically based on queues. Valid items are `ENABLED` or `DISABLED`. Defaults to `ENABLED`.
	State interface{}
	// The current status of the compute environment (for example, CREATING or VALID).
	Status interface{}
	// A short, human-readable string to provide additional details about the current status of the compute environment.
	StatusReason interface{}
	// The type of compute environment. Valid items are `EC2` or `SPOT`.
	Type interface{}
}

// The set of arguments for constructing a ComputeEnvironment resource.
type ComputeEnvironmentArgs struct {
	// The name for your compute environment. Up to 128 letters (uppercase and lowercase), numbers, and underscores are allowed.
	ComputeEnvironmentName interface{}
	// Details of the compute resources managed by the compute environment. This parameter is required for managed compute environments. See details below.
	ComputeResources interface{}
	// The full Amazon Resource Name (ARN) of the IAM role that allows AWS Batch to make calls to other AWS services on your behalf.
	ServiceRole interface{}
	// The state of the compute environment. If the state is `ENABLED`, then the compute environment accepts jobs from a queue and can scale out automatically based on queues. Valid items are `ENABLED` or `DISABLED`. Defaults to `ENABLED`.
	State interface{}
	// The type of compute environment. Valid items are `EC2` or `SPOT`.
	Type interface{}
}
