// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Get information on an EC2 Transit Gateway VPC Attachment.
 * 
 * ## Example Usage
 * 
 * ### By Filter
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const example = pulumi.output(aws.ec2transitgateway.getVpcAttachment({
 *     filters: [{
 *         name: "vpc-id",
 *         values: ["vpc-12345678"],
 *     }],
 * }));
 * ```
 * 
 * ### By Identifier
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const example = pulumi.output(aws.ec2transitgateway.getVpcAttachment({
 *     id: "tgw-attach-12345678",
 * }));
 * ```
 */
export function getVpcAttachment(args?: GetVpcAttachmentArgs, opts?: pulumi.InvokeOptions): Promise<GetVpcAttachmentResult> {
    args = args || {};
    return pulumi.runtime.invoke("aws:ec2transitgateway/getVpcAttachment:getVpcAttachment", {
        "filters": args.filters,
        "id": args.id,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getVpcAttachment.
 */
export interface GetVpcAttachmentArgs {
    /**
     * One or more configuration blocks containing name-values filters. Detailed below.
     */
    readonly filters?: { name: string, values: string[] }[];
    /**
     * Identifier of the EC2 Transit Gateway VPC Attachment.
     */
    readonly id?: string;
    readonly tags?: {[key: string]: any};
}

/**
 * A collection of values returned by getVpcAttachment.
 */
export interface GetVpcAttachmentResult {
    /**
     * Whether DNS support is enabled.
     */
    readonly dnsSupport: string;
    readonly filters?: { name: string, values: string[] }[];
    /**
     * EC2 Transit Gateway VPC Attachment identifier
     */
    readonly id?: string;
    /**
     * Whether IPv6 support is enabled.
     */
    readonly ipv6Support: string;
    /**
     * Identifiers of EC2 Subnets.
     */
    readonly subnetIds: string[];
    /**
     * Key-value tags for the EC2 Transit Gateway VPC Attachment
     */
    readonly tags: {[key: string]: any};
    /**
     * EC2 Transit Gateway identifier
     */
    readonly transitGatewayId: string;
    /**
     * Identifier of EC2 VPC.
     */
    readonly vpcId: string;
    /**
     * Identifier of the AWS account that owns the EC2 VPC.
     */
    readonly vpcOwnerId: string;
}
