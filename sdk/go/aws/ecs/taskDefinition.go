// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ecs

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a revision of an ECS task definition to be used in `ecs.Service`.
type TaskDefinition struct {
	pulumi.CustomResourceState

	// Full ARN of the Task Definition (including both `family` and `revision`).
	Arn pulumi.StringOutput `pulumi:"arn"`
	// A list of valid [container definitions]
	// (http://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_ContainerDefinition.html) provided as a
	// single valid JSON document. Please note that you should only provide values that are part of the container
	// definition document. For a detailed description of what parameters are available, see the [Task Definition Parameters]
	// (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_definition_parameters.html) section from the
	// official [Developer Guide](https://docs.aws.amazon.com/AmazonECS/latest/developerguide).
	ContainerDefinitions pulumi.StringOutput `pulumi:"containerDefinitions"`
	// The number of cpu units used by the task. If the `requiresCompatibilities` is `FARGATE` this field is required.
	Cpu pulumi.StringPtrOutput `pulumi:"cpu"`
	// The Amazon Resource Name (ARN) of the task execution role that the Amazon ECS container agent and the Docker daemon can assume.
	ExecutionRoleArn pulumi.StringPtrOutput `pulumi:"executionRoleArn"`
	// A unique name for your task definition.
	Family pulumi.StringOutput `pulumi:"family"`
	// Configuration block(s) with Inference Accelerators settings. Detailed below.
	InferenceAccelerators TaskDefinitionInferenceAcceleratorArrayOutput `pulumi:"inferenceAccelerators"`
	// The IPC resource namespace to be used for the containers in the task The valid values are `host`, `task`, and `none`.
	IpcMode pulumi.StringPtrOutput `pulumi:"ipcMode"`
	// The amount (in MiB) of memory used by the task. If the `requiresCompatibilities` is `FARGATE` this field is required.
	Memory pulumi.StringPtrOutput `pulumi:"memory"`
	// The Docker networking mode to use for the containers in the task. The valid values are `none`, `bridge`, `awsvpc`, and `host`.
	NetworkMode pulumi.StringOutput `pulumi:"networkMode"`
	// The process namespace to use for the containers in the task. The valid values are `host` and `task`.
	PidMode pulumi.StringPtrOutput `pulumi:"pidMode"`
	// A set of placement constraints rules that are taken into consideration during task placement. Maximum number of `placementConstraints` is `10`.
	PlacementConstraints TaskDefinitionPlacementConstraintArrayOutput `pulumi:"placementConstraints"`
	// The proxy configuration details for the App Mesh proxy.
	ProxyConfiguration TaskDefinitionProxyConfigurationPtrOutput `pulumi:"proxyConfiguration"`
	// A set of launch types required by the task. The valid values are `EC2` and `FARGATE`.
	RequiresCompatibilities pulumi.StringArrayOutput `pulumi:"requiresCompatibilities"`
	// The revision of the task in a particular family.
	Revision pulumi.IntOutput `pulumi:"revision"`
	// Key-value map of resource tags
	Tags pulumi.MapOutput `pulumi:"tags"`
	// The ARN of IAM role that allows your Amazon ECS container task to make calls to other AWS services.
	TaskRoleArn pulumi.StringPtrOutput `pulumi:"taskRoleArn"`
	// A set of volume blocks that containers in your task may use.
	Volumes TaskDefinitionVolumeArrayOutput `pulumi:"volumes"`
}

// NewTaskDefinition registers a new resource with the given unique name, arguments, and options.
func NewTaskDefinition(ctx *pulumi.Context,
	name string, args *TaskDefinitionArgs, opts ...pulumi.ResourceOption) (*TaskDefinition, error) {
	if args == nil || args.ContainerDefinitions == nil {
		return nil, errors.New("missing required argument 'ContainerDefinitions'")
	}
	if args == nil || args.Family == nil {
		return nil, errors.New("missing required argument 'Family'")
	}
	if args == nil {
		args = &TaskDefinitionArgs{}
	}
	var resource TaskDefinition
	err := ctx.RegisterResource("aws:ecs/taskDefinition:TaskDefinition", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTaskDefinition gets an existing TaskDefinition resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTaskDefinition(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TaskDefinitionState, opts ...pulumi.ResourceOption) (*TaskDefinition, error) {
	var resource TaskDefinition
	err := ctx.ReadResource("aws:ecs/taskDefinition:TaskDefinition", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TaskDefinition resources.
type taskDefinitionState struct {
	// Full ARN of the Task Definition (including both `family` and `revision`).
	Arn *string `pulumi:"arn"`
	// A list of valid [container definitions]
	// (http://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_ContainerDefinition.html) provided as a
	// single valid JSON document. Please note that you should only provide values that are part of the container
	// definition document. For a detailed description of what parameters are available, see the [Task Definition Parameters]
	// (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_definition_parameters.html) section from the
	// official [Developer Guide](https://docs.aws.amazon.com/AmazonECS/latest/developerguide).
	ContainerDefinitions *string `pulumi:"containerDefinitions"`
	// The number of cpu units used by the task. If the `requiresCompatibilities` is `FARGATE` this field is required.
	Cpu *string `pulumi:"cpu"`
	// The Amazon Resource Name (ARN) of the task execution role that the Amazon ECS container agent and the Docker daemon can assume.
	ExecutionRoleArn *string `pulumi:"executionRoleArn"`
	// A unique name for your task definition.
	Family *string `pulumi:"family"`
	// Configuration block(s) with Inference Accelerators settings. Detailed below.
	InferenceAccelerators []TaskDefinitionInferenceAccelerator `pulumi:"inferenceAccelerators"`
	// The IPC resource namespace to be used for the containers in the task The valid values are `host`, `task`, and `none`.
	IpcMode *string `pulumi:"ipcMode"`
	// The amount (in MiB) of memory used by the task. If the `requiresCompatibilities` is `FARGATE` this field is required.
	Memory *string `pulumi:"memory"`
	// The Docker networking mode to use for the containers in the task. The valid values are `none`, `bridge`, `awsvpc`, and `host`.
	NetworkMode *string `pulumi:"networkMode"`
	// The process namespace to use for the containers in the task. The valid values are `host` and `task`.
	PidMode *string `pulumi:"pidMode"`
	// A set of placement constraints rules that are taken into consideration during task placement. Maximum number of `placementConstraints` is `10`.
	PlacementConstraints []TaskDefinitionPlacementConstraint `pulumi:"placementConstraints"`
	// The proxy configuration details for the App Mesh proxy.
	ProxyConfiguration *TaskDefinitionProxyConfiguration `pulumi:"proxyConfiguration"`
	// A set of launch types required by the task. The valid values are `EC2` and `FARGATE`.
	RequiresCompatibilities []string `pulumi:"requiresCompatibilities"`
	// The revision of the task in a particular family.
	Revision *int `pulumi:"revision"`
	// Key-value map of resource tags
	Tags map[string]interface{} `pulumi:"tags"`
	// The ARN of IAM role that allows your Amazon ECS container task to make calls to other AWS services.
	TaskRoleArn *string `pulumi:"taskRoleArn"`
	// A set of volume blocks that containers in your task may use.
	Volumes []TaskDefinitionVolume `pulumi:"volumes"`
}

type TaskDefinitionState struct {
	// Full ARN of the Task Definition (including both `family` and `revision`).
	Arn pulumi.StringPtrInput
	// A list of valid [container definitions]
	// (http://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_ContainerDefinition.html) provided as a
	// single valid JSON document. Please note that you should only provide values that are part of the container
	// definition document. For a detailed description of what parameters are available, see the [Task Definition Parameters]
	// (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_definition_parameters.html) section from the
	// official [Developer Guide](https://docs.aws.amazon.com/AmazonECS/latest/developerguide).
	ContainerDefinitions pulumi.StringPtrInput
	// The number of cpu units used by the task. If the `requiresCompatibilities` is `FARGATE` this field is required.
	Cpu pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) of the task execution role that the Amazon ECS container agent and the Docker daemon can assume.
	ExecutionRoleArn pulumi.StringPtrInput
	// A unique name for your task definition.
	Family pulumi.StringPtrInput
	// Configuration block(s) with Inference Accelerators settings. Detailed below.
	InferenceAccelerators TaskDefinitionInferenceAcceleratorArrayInput
	// The IPC resource namespace to be used for the containers in the task The valid values are `host`, `task`, and `none`.
	IpcMode pulumi.StringPtrInput
	// The amount (in MiB) of memory used by the task. If the `requiresCompatibilities` is `FARGATE` this field is required.
	Memory pulumi.StringPtrInput
	// The Docker networking mode to use for the containers in the task. The valid values are `none`, `bridge`, `awsvpc`, and `host`.
	NetworkMode pulumi.StringPtrInput
	// The process namespace to use for the containers in the task. The valid values are `host` and `task`.
	PidMode pulumi.StringPtrInput
	// A set of placement constraints rules that are taken into consideration during task placement. Maximum number of `placementConstraints` is `10`.
	PlacementConstraints TaskDefinitionPlacementConstraintArrayInput
	// The proxy configuration details for the App Mesh proxy.
	ProxyConfiguration TaskDefinitionProxyConfigurationPtrInput
	// A set of launch types required by the task. The valid values are `EC2` and `FARGATE`.
	RequiresCompatibilities pulumi.StringArrayInput
	// The revision of the task in a particular family.
	Revision pulumi.IntPtrInput
	// Key-value map of resource tags
	Tags pulumi.MapInput
	// The ARN of IAM role that allows your Amazon ECS container task to make calls to other AWS services.
	TaskRoleArn pulumi.StringPtrInput
	// A set of volume blocks that containers in your task may use.
	Volumes TaskDefinitionVolumeArrayInput
}

func (TaskDefinitionState) ElementType() reflect.Type {
	return reflect.TypeOf((*taskDefinitionState)(nil)).Elem()
}

type taskDefinitionArgs struct {
	// A list of valid [container definitions]
	// (http://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_ContainerDefinition.html) provided as a
	// single valid JSON document. Please note that you should only provide values that are part of the container
	// definition document. For a detailed description of what parameters are available, see the [Task Definition Parameters]
	// (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_definition_parameters.html) section from the
	// official [Developer Guide](https://docs.aws.amazon.com/AmazonECS/latest/developerguide).
	ContainerDefinitions string `pulumi:"containerDefinitions"`
	// The number of cpu units used by the task. If the `requiresCompatibilities` is `FARGATE` this field is required.
	Cpu *string `pulumi:"cpu"`
	// The Amazon Resource Name (ARN) of the task execution role that the Amazon ECS container agent and the Docker daemon can assume.
	ExecutionRoleArn *string `pulumi:"executionRoleArn"`
	// A unique name for your task definition.
	Family string `pulumi:"family"`
	// Configuration block(s) with Inference Accelerators settings. Detailed below.
	InferenceAccelerators []TaskDefinitionInferenceAccelerator `pulumi:"inferenceAccelerators"`
	// The IPC resource namespace to be used for the containers in the task The valid values are `host`, `task`, and `none`.
	IpcMode *string `pulumi:"ipcMode"`
	// The amount (in MiB) of memory used by the task. If the `requiresCompatibilities` is `FARGATE` this field is required.
	Memory *string `pulumi:"memory"`
	// The Docker networking mode to use for the containers in the task. The valid values are `none`, `bridge`, `awsvpc`, and `host`.
	NetworkMode *string `pulumi:"networkMode"`
	// The process namespace to use for the containers in the task. The valid values are `host` and `task`.
	PidMode *string `pulumi:"pidMode"`
	// A set of placement constraints rules that are taken into consideration during task placement. Maximum number of `placementConstraints` is `10`.
	PlacementConstraints []TaskDefinitionPlacementConstraint `pulumi:"placementConstraints"`
	// The proxy configuration details for the App Mesh proxy.
	ProxyConfiguration *TaskDefinitionProxyConfiguration `pulumi:"proxyConfiguration"`
	// A set of launch types required by the task. The valid values are `EC2` and `FARGATE`.
	RequiresCompatibilities []string `pulumi:"requiresCompatibilities"`
	// Key-value map of resource tags
	Tags map[string]interface{} `pulumi:"tags"`
	// The ARN of IAM role that allows your Amazon ECS container task to make calls to other AWS services.
	TaskRoleArn *string `pulumi:"taskRoleArn"`
	// A set of volume blocks that containers in your task may use.
	Volumes []TaskDefinitionVolume `pulumi:"volumes"`
}

// The set of arguments for constructing a TaskDefinition resource.
type TaskDefinitionArgs struct {
	// A list of valid [container definitions]
	// (http://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_ContainerDefinition.html) provided as a
	// single valid JSON document. Please note that you should only provide values that are part of the container
	// definition document. For a detailed description of what parameters are available, see the [Task Definition Parameters]
	// (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_definition_parameters.html) section from the
	// official [Developer Guide](https://docs.aws.amazon.com/AmazonECS/latest/developerguide).
	ContainerDefinitions pulumi.StringInput
	// The number of cpu units used by the task. If the `requiresCompatibilities` is `FARGATE` this field is required.
	Cpu pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) of the task execution role that the Amazon ECS container agent and the Docker daemon can assume.
	ExecutionRoleArn pulumi.StringPtrInput
	// A unique name for your task definition.
	Family pulumi.StringInput
	// Configuration block(s) with Inference Accelerators settings. Detailed below.
	InferenceAccelerators TaskDefinitionInferenceAcceleratorArrayInput
	// The IPC resource namespace to be used for the containers in the task The valid values are `host`, `task`, and `none`.
	IpcMode pulumi.StringPtrInput
	// The amount (in MiB) of memory used by the task. If the `requiresCompatibilities` is `FARGATE` this field is required.
	Memory pulumi.StringPtrInput
	// The Docker networking mode to use for the containers in the task. The valid values are `none`, `bridge`, `awsvpc`, and `host`.
	NetworkMode pulumi.StringPtrInput
	// The process namespace to use for the containers in the task. The valid values are `host` and `task`.
	PidMode pulumi.StringPtrInput
	// A set of placement constraints rules that are taken into consideration during task placement. Maximum number of `placementConstraints` is `10`.
	PlacementConstraints TaskDefinitionPlacementConstraintArrayInput
	// The proxy configuration details for the App Mesh proxy.
	ProxyConfiguration TaskDefinitionProxyConfigurationPtrInput
	// A set of launch types required by the task. The valid values are `EC2` and `FARGATE`.
	RequiresCompatibilities pulumi.StringArrayInput
	// Key-value map of resource tags
	Tags pulumi.MapInput
	// The ARN of IAM role that allows your Amazon ECS container task to make calls to other AWS services.
	TaskRoleArn pulumi.StringPtrInput
	// A set of volume blocks that containers in your task may use.
	Volumes TaskDefinitionVolumeArrayInput
}

func (TaskDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*taskDefinitionArgs)(nil)).Elem()
}
