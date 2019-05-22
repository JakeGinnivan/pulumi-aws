// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a revision of an ECS task definition to be used in `aws_ecs_service`.
 */
export class TaskDefinition extends pulumi.CustomResource {
    /**
     * Get an existing TaskDefinition resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TaskDefinitionState, opts?: pulumi.CustomResourceOptions): TaskDefinition {
        return new TaskDefinition(name, <any>state, { ...opts, id: id });
    }

    /**
     * Full ARN of the Task Definition (including both `family` and `revision`).
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * A list of valid [container definitions]
     * (http://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_ContainerDefinition.html) provided as a
     * single valid JSON document. Please note that you should only provide values that are part of the container
     * definition document. For a detailed description of what parameters are available, see the [Task Definition Parameters]
     * (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_definition_parameters.html) section from the
     * official [Developer Guide](https://docs.aws.amazon.com/AmazonECS/latest/developerguide).
     */
    public readonly containerDefinitions!: pulumi.Output<string>;
    /**
     * The number of cpu units used by the task. If the `requires_compatibilities` is `FARGATE` this field is required.
     */
    public readonly cpu!: pulumi.Output<string | undefined>;
    /**
     * The Amazon Resource Name (ARN) of the task execution role that the Amazon ECS container agent and the Docker daemon can assume.
     */
    public readonly executionRoleArn!: pulumi.Output<string | undefined>;
    /**
     * A unique name for your task definition.
     */
    public readonly family!: pulumi.Output<string>;
    /**
     * The IPC resource namespace to be used for the containers in the task The valid values are `host`, `task`, and `none`.
     */
    public readonly ipcMode!: pulumi.Output<string | undefined>;
    /**
     * The amount (in MiB) of memory used by the task. If the `requires_compatibilities` is `FARGATE` this field is required.
     */
    public readonly memory!: pulumi.Output<string | undefined>;
    /**
     * The Docker networking mode to use for the containers in the task. The valid values are `none`, `bridge`, `awsvpc`, and `host`.
     */
    public readonly networkMode!: pulumi.Output<string>;
    /**
     * The process namespace to use for the containers in the task. The valid values are `host` and `task`.
     */
    public readonly pidMode!: pulumi.Output<string | undefined>;
    /**
     * A set of placement constraints rules that are taken into consideration during task placement. Maximum number of `placement_constraints` is `10`.
     */
    public readonly placementConstraints!: pulumi.Output<{ expression?: string, type: string }[] | undefined>;
    /**
     * A set of launch types required by the task. The valid values are `EC2` and `FARGATE`.
     */
    public readonly requiresCompatibilities!: pulumi.Output<string[] | undefined>;
    /**
     * The revision of the task in a particular family.
     */
    public /*out*/ readonly revision!: pulumi.Output<number>;
    /**
     * Key-value mapping of resource tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The ARN of IAM role that allows your Amazon ECS container task to make calls to other AWS services.
     */
    public readonly taskRoleArn!: pulumi.Output<string | undefined>;
    /**
     * A set of volume blocks that containers in your task may use.
     */
    public readonly volumes!: pulumi.Output<{ dockerVolumeConfiguration?: { autoprovision?: boolean, driver?: string, driverOpts?: {[key: string]: string}, labels?: {[key: string]: string}, scope: string }, hostPath?: string, name: string }[] | undefined>;

    /**
     * Create a TaskDefinition resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TaskDefinitionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TaskDefinitionArgs | TaskDefinitionState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as TaskDefinitionState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["containerDefinitions"] = state ? state.containerDefinitions : undefined;
            inputs["cpu"] = state ? state.cpu : undefined;
            inputs["executionRoleArn"] = state ? state.executionRoleArn : undefined;
            inputs["family"] = state ? state.family : undefined;
            inputs["ipcMode"] = state ? state.ipcMode : undefined;
            inputs["memory"] = state ? state.memory : undefined;
            inputs["networkMode"] = state ? state.networkMode : undefined;
            inputs["pidMode"] = state ? state.pidMode : undefined;
            inputs["placementConstraints"] = state ? state.placementConstraints : undefined;
            inputs["requiresCompatibilities"] = state ? state.requiresCompatibilities : undefined;
            inputs["revision"] = state ? state.revision : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["taskRoleArn"] = state ? state.taskRoleArn : undefined;
            inputs["volumes"] = state ? state.volumes : undefined;
        } else {
            const args = argsOrState as TaskDefinitionArgs | undefined;
            if (!args || args.containerDefinitions === undefined) {
                throw new Error("Missing required property 'containerDefinitions'");
            }
            if (!args || args.family === undefined) {
                throw new Error("Missing required property 'family'");
            }
            inputs["containerDefinitions"] = args ? args.containerDefinitions : undefined;
            inputs["cpu"] = args ? args.cpu : undefined;
            inputs["executionRoleArn"] = args ? args.executionRoleArn : undefined;
            inputs["family"] = args ? args.family : undefined;
            inputs["ipcMode"] = args ? args.ipcMode : undefined;
            inputs["memory"] = args ? args.memory : undefined;
            inputs["networkMode"] = args ? args.networkMode : undefined;
            inputs["pidMode"] = args ? args.pidMode : undefined;
            inputs["placementConstraints"] = args ? args.placementConstraints : undefined;
            inputs["requiresCompatibilities"] = args ? args.requiresCompatibilities : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["taskRoleArn"] = args ? args.taskRoleArn : undefined;
            inputs["volumes"] = args ? args.volumes : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["revision"] = undefined /*out*/;
        }
        super("aws:ecs/taskDefinition:TaskDefinition", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TaskDefinition resources.
 */
export interface TaskDefinitionState {
    /**
     * Full ARN of the Task Definition (including both `family` and `revision`).
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * A list of valid [container definitions]
     * (http://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_ContainerDefinition.html) provided as a
     * single valid JSON document. Please note that you should only provide values that are part of the container
     * definition document. For a detailed description of what parameters are available, see the [Task Definition Parameters]
     * (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_definition_parameters.html) section from the
     * official [Developer Guide](https://docs.aws.amazon.com/AmazonECS/latest/developerguide).
     */
    readonly containerDefinitions?: pulumi.Input<string>;
    /**
     * The number of cpu units used by the task. If the `requires_compatibilities` is `FARGATE` this field is required.
     */
    readonly cpu?: pulumi.Input<string>;
    /**
     * The Amazon Resource Name (ARN) of the task execution role that the Amazon ECS container agent and the Docker daemon can assume.
     */
    readonly executionRoleArn?: pulumi.Input<string>;
    /**
     * A unique name for your task definition.
     */
    readonly family?: pulumi.Input<string>;
    /**
     * The IPC resource namespace to be used for the containers in the task The valid values are `host`, `task`, and `none`.
     */
    readonly ipcMode?: pulumi.Input<string>;
    /**
     * The amount (in MiB) of memory used by the task. If the `requires_compatibilities` is `FARGATE` this field is required.
     */
    readonly memory?: pulumi.Input<string>;
    /**
     * The Docker networking mode to use for the containers in the task. The valid values are `none`, `bridge`, `awsvpc`, and `host`.
     */
    readonly networkMode?: pulumi.Input<string>;
    /**
     * The process namespace to use for the containers in the task. The valid values are `host` and `task`.
     */
    readonly pidMode?: pulumi.Input<string>;
    /**
     * A set of placement constraints rules that are taken into consideration during task placement. Maximum number of `placement_constraints` is `10`.
     */
    readonly placementConstraints?: pulumi.Input<pulumi.Input<{ expression?: pulumi.Input<string>, type: pulumi.Input<string> }>[]>;
    /**
     * A set of launch types required by the task. The valid values are `EC2` and `FARGATE`.
     */
    readonly requiresCompatibilities?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The revision of the task in a particular family.
     */
    readonly revision?: pulumi.Input<number>;
    /**
     * Key-value mapping of resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The ARN of IAM role that allows your Amazon ECS container task to make calls to other AWS services.
     */
    readonly taskRoleArn?: pulumi.Input<string>;
    /**
     * A set of volume blocks that containers in your task may use.
     */
    readonly volumes?: pulumi.Input<pulumi.Input<{ dockerVolumeConfiguration?: pulumi.Input<{ autoprovision?: pulumi.Input<boolean>, driver?: pulumi.Input<string>, driverOpts?: pulumi.Input<{[key: string]: pulumi.Input<string>}>, labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>, scope?: pulumi.Input<string> }>, hostPath?: pulumi.Input<string>, name: pulumi.Input<string> }>[]>;
}

/**
 * The set of arguments for constructing a TaskDefinition resource.
 */
export interface TaskDefinitionArgs {
    /**
     * A list of valid [container definitions]
     * (http://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_ContainerDefinition.html) provided as a
     * single valid JSON document. Please note that you should only provide values that are part of the container
     * definition document. For a detailed description of what parameters are available, see the [Task Definition Parameters]
     * (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_definition_parameters.html) section from the
     * official [Developer Guide](https://docs.aws.amazon.com/AmazonECS/latest/developerguide).
     */
    readonly containerDefinitions: pulumi.Input<string>;
    /**
     * The number of cpu units used by the task. If the `requires_compatibilities` is `FARGATE` this field is required.
     */
    readonly cpu?: pulumi.Input<string>;
    /**
     * The Amazon Resource Name (ARN) of the task execution role that the Amazon ECS container agent and the Docker daemon can assume.
     */
    readonly executionRoleArn?: pulumi.Input<string>;
    /**
     * A unique name for your task definition.
     */
    readonly family: pulumi.Input<string>;
    /**
     * The IPC resource namespace to be used for the containers in the task The valid values are `host`, `task`, and `none`.
     */
    readonly ipcMode?: pulumi.Input<string>;
    /**
     * The amount (in MiB) of memory used by the task. If the `requires_compatibilities` is `FARGATE` this field is required.
     */
    readonly memory?: pulumi.Input<string>;
    /**
     * The Docker networking mode to use for the containers in the task. The valid values are `none`, `bridge`, `awsvpc`, and `host`.
     */
    readonly networkMode?: pulumi.Input<string>;
    /**
     * The process namespace to use for the containers in the task. The valid values are `host` and `task`.
     */
    readonly pidMode?: pulumi.Input<string>;
    /**
     * A set of placement constraints rules that are taken into consideration during task placement. Maximum number of `placement_constraints` is `10`.
     */
    readonly placementConstraints?: pulumi.Input<pulumi.Input<{ expression?: pulumi.Input<string>, type: pulumi.Input<string> }>[]>;
    /**
     * A set of launch types required by the task. The valid values are `EC2` and `FARGATE`.
     */
    readonly requiresCompatibilities?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Key-value mapping of resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The ARN of IAM role that allows your Amazon ECS container task to make calls to other AWS services.
     */
    readonly taskRoleArn?: pulumi.Input<string>;
    /**
     * A set of volume blocks that containers in your task may use.
     */
    readonly volumes?: pulumi.Input<pulumi.Input<{ dockerVolumeConfiguration?: pulumi.Input<{ autoprovision?: pulumi.Input<boolean>, driver?: pulumi.Input<string>, driverOpts?: pulumi.Input<{[key: string]: pulumi.Input<string>}>, labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>, scope?: pulumi.Input<string> }>, hostPath?: pulumi.Input<string>, name: pulumi.Input<string> }>[]>;
}
