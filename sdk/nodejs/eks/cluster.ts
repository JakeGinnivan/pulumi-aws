// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an EKS Cluster.
 * 
 * ## Example Usage
 * 
 * ### Basic Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const example = new aws.eks.Cluster("example", {
 *     roleArn: aws_iam_role_example.arn,
 *     vpcConfig: {
 *         subnetIds: [
 *             aws_subnet_example1.id,
 *             aws_subnet_example2.id,
 *         ],
 *     },
 * });
 * 
 * export const endpoint = example.endpoint;
 * export const kubeconfig_certificate_authority_data = example.certificateAuthority.data;
 * ```
 * 
 * ### Enabling Control Plane Logging
 * 
 * [EKS Control Plane Logging](https://docs.aws.amazon.com/eks/latest/userguide/control-plane-logs.html) can be enabled via the `enabled_cluster_log_types` argument. To manage the CloudWatch Log Group retention period, the [`aws_cloudwatch_log_group` resource](https://www.terraform.io/docs/providers/aws/r/cloudwatch_log_group.html) can be used.
 * 
 * > The below configuration uses [`depends_on`](https://www.terraform.io/docs/configuration/resources.html#depends_on-explicit-resource-dependencies) to prevent ordering issues with EKS automatically creating the log group first and a variable for naming consistency. Other ordering and naming methodologies may be more appropriate for your environment.
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const config = new pulumi.Config();
 * const clusterName = config.get("clusterName") || "example";
 * 
 * const exampleLogGroup = new aws.cloudwatch.LogGroup("example", {
 *     retentionInDays: 7,
 * });
 * const exampleCluster = new aws.eks.Cluster("example", {
 *     enabledClusterLogTypes: [
 *         "api",
 *         "audit",
 *     ],
 * }, {dependsOn: [exampleLogGroup]});
 * ```
 */
export class Cluster extends pulumi.CustomResource {
    /**
     * Get an existing Cluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ClusterState, opts?: pulumi.CustomResourceOptions): Cluster {
        return new Cluster(name, <any>state, { ...opts, id: id });
    }

    /**
     * The Amazon Resource Name (ARN) of the cluster.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Nested attribute containing `certificate-authority-data` for your cluster.
     */
    public /*out*/ readonly certificateAuthority!: pulumi.Output<{ data: string }>;
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * A list of the desired control plane logging to enable. For more information, see [Amazon EKS Control Plane Logging](https://docs.aws.amazon.com/eks/latest/userguide/control-plane-logs.html)
     */
    public readonly enabledClusterLogTypes!: pulumi.Output<string[] | undefined>;
    /**
     * The endpoint for your Kubernetes API server.
     */
    public /*out*/ readonly endpoint!: pulumi.Output<string>;
    /**
     * Name of the cluster.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The platform version for the cluster.
     */
    public /*out*/ readonly platformVersion!: pulumi.Output<string>;
    /**
     * The Amazon Resource Name (ARN) of the IAM role that provides permissions for the Kubernetes control plane to make calls to AWS API operations on your behalf.
     */
    public readonly roleArn!: pulumi.Output<string>;
    /**
     * Desired Kubernetes master version. If you do not specify a value, the latest available version at resource creation is used and no upgrades will occur except those automatically triggered by EKS. The value must be configured and increased to upgrade the version when desired. Downgrades are not supported by EKS.
     */
    public readonly version!: pulumi.Output<string>;
    /**
     * Nested argument for the VPC associated with your cluster. Amazon EKS VPC resources have specific requirements to work properly with Kubernetes. For more information, see [Cluster VPC Considerations](https://docs.aws.amazon.com/eks/latest/userguide/network_reqs.html) and [Cluster Security Group Considerations](https://docs.aws.amazon.com/eks/latest/userguide/sec-group-reqs.html) in the Amazon EKS User Guide. Configuration detailed below.
     */
    public readonly vpcConfig!: pulumi.Output<{ endpointPrivateAccess?: boolean, endpointPublicAccess?: boolean, securityGroupIds?: string[], subnetIds: string[], vpcId: string }>;

    /**
     * Create a Cluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ClusterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ClusterArgs | ClusterState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ClusterState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["certificateAuthority"] = state ? state.certificateAuthority : undefined;
            inputs["createdAt"] = state ? state.createdAt : undefined;
            inputs["enabledClusterLogTypes"] = state ? state.enabledClusterLogTypes : undefined;
            inputs["endpoint"] = state ? state.endpoint : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["platformVersion"] = state ? state.platformVersion : undefined;
            inputs["roleArn"] = state ? state.roleArn : undefined;
            inputs["version"] = state ? state.version : undefined;
            inputs["vpcConfig"] = state ? state.vpcConfig : undefined;
        } else {
            const args = argsOrState as ClusterArgs | undefined;
            if (!args || args.roleArn === undefined) {
                throw new Error("Missing required property 'roleArn'");
            }
            if (!args || args.vpcConfig === undefined) {
                throw new Error("Missing required property 'vpcConfig'");
            }
            inputs["enabledClusterLogTypes"] = args ? args.enabledClusterLogTypes : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["roleArn"] = args ? args.roleArn : undefined;
            inputs["version"] = args ? args.version : undefined;
            inputs["vpcConfig"] = args ? args.vpcConfig : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["certificateAuthority"] = undefined /*out*/;
            inputs["createdAt"] = undefined /*out*/;
            inputs["endpoint"] = undefined /*out*/;
            inputs["platformVersion"] = undefined /*out*/;
        }
        super("aws:eks/cluster:Cluster", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Cluster resources.
 */
export interface ClusterState {
    /**
     * The Amazon Resource Name (ARN) of the cluster.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * Nested attribute containing `certificate-authority-data` for your cluster.
     */
    readonly certificateAuthority?: pulumi.Input<{ data?: pulumi.Input<string> }>;
    readonly createdAt?: pulumi.Input<string>;
    /**
     * A list of the desired control plane logging to enable. For more information, see [Amazon EKS Control Plane Logging](https://docs.aws.amazon.com/eks/latest/userguide/control-plane-logs.html)
     */
    readonly enabledClusterLogTypes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The endpoint for your Kubernetes API server.
     */
    readonly endpoint?: pulumi.Input<string>;
    /**
     * Name of the cluster.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The platform version for the cluster.
     */
    readonly platformVersion?: pulumi.Input<string>;
    /**
     * The Amazon Resource Name (ARN) of the IAM role that provides permissions for the Kubernetes control plane to make calls to AWS API operations on your behalf.
     */
    readonly roleArn?: pulumi.Input<string>;
    /**
     * Desired Kubernetes master version. If you do not specify a value, the latest available version at resource creation is used and no upgrades will occur except those automatically triggered by EKS. The value must be configured and increased to upgrade the version when desired. Downgrades are not supported by EKS.
     */
    readonly version?: pulumi.Input<string>;
    /**
     * Nested argument for the VPC associated with your cluster. Amazon EKS VPC resources have specific requirements to work properly with Kubernetes. For more information, see [Cluster VPC Considerations](https://docs.aws.amazon.com/eks/latest/userguide/network_reqs.html) and [Cluster Security Group Considerations](https://docs.aws.amazon.com/eks/latest/userguide/sec-group-reqs.html) in the Amazon EKS User Guide. Configuration detailed below.
     */
    readonly vpcConfig?: pulumi.Input<{ endpointPrivateAccess?: pulumi.Input<boolean>, endpointPublicAccess?: pulumi.Input<boolean>, securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>, subnetIds: pulumi.Input<pulumi.Input<string>[]>, vpcId?: pulumi.Input<string> }>;
}

/**
 * The set of arguments for constructing a Cluster resource.
 */
export interface ClusterArgs {
    /**
     * A list of the desired control plane logging to enable. For more information, see [Amazon EKS Control Plane Logging](https://docs.aws.amazon.com/eks/latest/userguide/control-plane-logs.html)
     */
    readonly enabledClusterLogTypes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name of the cluster.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The Amazon Resource Name (ARN) of the IAM role that provides permissions for the Kubernetes control plane to make calls to AWS API operations on your behalf.
     */
    readonly roleArn: pulumi.Input<string>;
    /**
     * Desired Kubernetes master version. If you do not specify a value, the latest available version at resource creation is used and no upgrades will occur except those automatically triggered by EKS. The value must be configured and increased to upgrade the version when desired. Downgrades are not supported by EKS.
     */
    readonly version?: pulumi.Input<string>;
    /**
     * Nested argument for the VPC associated with your cluster. Amazon EKS VPC resources have specific requirements to work properly with Kubernetes. For more information, see [Cluster VPC Considerations](https://docs.aws.amazon.com/eks/latest/userguide/network_reqs.html) and [Cluster Security Group Considerations](https://docs.aws.amazon.com/eks/latest/userguide/sec-group-reqs.html) in the Amazon EKS User Guide. Configuration detailed below.
     */
    readonly vpcConfig: pulumi.Input<{ endpointPrivateAccess?: pulumi.Input<boolean>, endpointPublicAccess?: pulumi.Input<boolean>, securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>, subnetIds: pulumi.Input<pulumi.Input<string>[]>, vpcId?: pulumi.Input<string> }>;
}
