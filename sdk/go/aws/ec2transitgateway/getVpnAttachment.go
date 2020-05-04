// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2transitgateway

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Get information on an EC2 Transit Gateway VPN Attachment.
func GetVpnAttachment(ctx *pulumi.Context, args *GetVpnAttachmentArgs, opts ...pulumi.InvokeOption) (*GetVpnAttachmentResult, error) {
	var rv GetVpnAttachmentResult
	err := ctx.Invoke("aws:ec2transitgateway/getVpnAttachment:getVpnAttachment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVpnAttachment.
type GetVpnAttachmentArgs struct {
	// Configuration block(s) for filtering. Detailed below.
	Filters []GetVpnAttachmentFilter `pulumi:"filters"`
	// A map of tags, each pair of which must exactly match a pair on the desired Transit Gateway VPN Attachment.
	Tags map[string]interface{} `pulumi:"tags"`
	// Identifier of the EC2 Transit Gateway.
	TransitGatewayId *string `pulumi:"transitGatewayId"`
	// Identifier of the EC2 VPN Connection.
	VpnConnectionId *string `pulumi:"vpnConnectionId"`
}

// A collection of values returned by getVpnAttachment.
type GetVpnAttachmentResult struct {
	Filters []GetVpnAttachmentFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Key-value tags for the EC2 Transit Gateway VPN Attachment
	Tags             map[string]interface{} `pulumi:"tags"`
	TransitGatewayId *string                `pulumi:"transitGatewayId"`
	VpnConnectionId  *string                `pulumi:"vpnConnectionId"`
}
