// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides an SWF Domain resource.
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const foo = new aws.swf.Domain("foo", {
 *     description: "SWF Domain",
 *     workflowExecutionRetentionPeriodInDays: "30",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/swf_domain.html.markdown.
 */
export class Domain extends pulumi.CustomResource {
    /**
     * Get an existing Domain resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DomainState, opts?: pulumi.CustomResourceOptions): Domain {
        return new Domain(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:swf/domain:Domain';

    /**
     * Returns true if the given object is an instance of Domain.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Domain {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Domain.__pulumiType;
    }

    /**
     * Amazon Resource Name (ARN)
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The domain description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The name of the domain. If omitted, this provider will assign a random, unique name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Creates a unique name beginning with the specified prefix. Conflicts with `name`.
     */
    public readonly namePrefix!: pulumi.Output<string | undefined>;
    /**
     * Key-value map of resource tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * Length of time that SWF will continue to retain information about the workflow execution after the workflow execution is complete, must be between 0 and 90 days.
     */
    public readonly workflowExecutionRetentionPeriodInDays!: pulumi.Output<string>;

    /**
     * Create a Domain resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DomainArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DomainArgs | DomainState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as DomainState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["namePrefix"] = state ? state.namePrefix : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["workflowExecutionRetentionPeriodInDays"] = state ? state.workflowExecutionRetentionPeriodInDays : undefined;
        } else {
            const args = argsOrState as DomainArgs | undefined;
            if (!args || args.workflowExecutionRetentionPeriodInDays === undefined) {
                throw new Error("Missing required property 'workflowExecutionRetentionPeriodInDays'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["namePrefix"] = args ? args.namePrefix : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["workflowExecutionRetentionPeriodInDays"] = args ? args.workflowExecutionRetentionPeriodInDays : undefined;
            inputs["arn"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Domain.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Domain resources.
 */
export interface DomainState {
    /**
     * Amazon Resource Name (ARN)
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * The domain description.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The name of the domain. If omitted, this provider will assign a random, unique name.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Creates a unique name beginning with the specified prefix. Conflicts with `name`.
     */
    readonly namePrefix?: pulumi.Input<string>;
    /**
     * Key-value map of resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * Length of time that SWF will continue to retain information about the workflow execution after the workflow execution is complete, must be between 0 and 90 days.
     */
    readonly workflowExecutionRetentionPeriodInDays?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Domain resource.
 */
export interface DomainArgs {
    /**
     * The domain description.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The name of the domain. If omitted, this provider will assign a random, unique name.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Creates a unique name beginning with the specified prefix. Conflicts with `name`.
     */
    readonly namePrefix?: pulumi.Input<string>;
    /**
     * Key-value map of resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * Length of time that SWF will continue to retain information about the workflow execution after the workflow execution is complete, must be between 0 and 90 days.
     */
    readonly workflowExecutionRetentionPeriodInDays: pulumi.Input<string>;
}
