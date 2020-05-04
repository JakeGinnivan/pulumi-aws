// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This resource can be useful for getting back a list of route table ids to be referenced elsewhere.
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const rts = pulumi.output(aws.ec2.getRouteTables({
 *     filters: [{
 *         name: "tag:kubernetes.io/kops/role",
 *         values: ["private*"],
 *     }],
 *     vpcId: var_vpc_id,
 * }, { async: true }));
 * const route: aws.ec2.Route[] = [];
 * for (let i = 0; i < rts.apply(rts => rts.ids.length); i++) {
 *     route.push(new aws.ec2.Route(`r-${i}`, {
 *         destinationCidrBlock: "10.0.1.0/22",
 *         routeTableId: rts.apply(rts => rts.ids[i]),
 *         vpcPeeringConnectionId: "pcx-0e9a7a9ecd137dc54",
 *     }));
 * }
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/route_tables.html.markdown.
 */
export function getRouteTables(args?: GetRouteTablesArgs, opts?: pulumi.InvokeOptions): Promise<GetRouteTablesResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:ec2/getRouteTables:getRouteTables", {
        "filters": args.filters,
        "tags": args.tags,
        "vpcId": args.vpcId,
    }, opts);
}

/**
 * A collection of arguments for invoking getRouteTables.
 */
export interface GetRouteTablesArgs {
    /**
     * Custom filter block as described below.
     */
    readonly filters?: inputs.ec2.GetRouteTablesFilter[];
    /**
     * A map of tags, each pair of which must exactly match
     * a pair on the desired route tables.
     */
    readonly tags?: {[key: string]: any};
    /**
     * The VPC ID that you want to filter from.
     */
    readonly vpcId?: string;
}

/**
 * A collection of values returned by getRouteTables.
 */
export interface GetRouteTablesResult {
    readonly filters?: outputs.ec2.GetRouteTablesFilter[];
    /**
     * A set of all the route table ids found. This data source will fail if none are found.
     */
    readonly ids: string[];
    readonly tags: {[key: string]: any};
    readonly vpcId?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
