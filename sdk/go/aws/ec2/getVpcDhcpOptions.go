// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Retrieve information about an EC2 DHCP Options configuration.
func LookupVpcDhcpOptions(ctx *pulumi.Context, args *GetVpcDhcpOptionsArgs) (*GetVpcDhcpOptionsResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["dhcpOptionsId"] = args.DhcpOptionsId
		inputs["filters"] = args.Filters
		inputs["tags"] = args.Tags
	}
	outputs, err := ctx.Invoke("aws:ec2/getVpcDhcpOptions:getVpcDhcpOptions", inputs)
	if err != nil {
		return nil, err
	}
	return &GetVpcDhcpOptionsResult{
		DhcpOptionsId: outputs["dhcpOptionsId"],
		DomainName: outputs["domainName"],
		DomainNameServers: outputs["domainNameServers"],
		NetbiosNameServers: outputs["netbiosNameServers"],
		NetbiosNodeType: outputs["netbiosNodeType"],
		NtpServers: outputs["ntpServers"],
		OwnerId: outputs["ownerId"],
		Tags: outputs["tags"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getVpcDhcpOptions.
type GetVpcDhcpOptionsArgs struct {
	// The EC2 DHCP Options ID.
	DhcpOptionsId interface{}
	// List of custom filters as described below.
	Filters interface{}
	Tags interface{}
}

// A collection of values returned by getVpcDhcpOptions.
type GetVpcDhcpOptionsResult struct {
	// EC2 DHCP Options ID
	DhcpOptionsId interface{}
	// The suffix domain name to used when resolving non Fully Qualified Domain Names. e.g. the `search` value in the `/etc/resolv.conf` file.
	DomainName interface{}
	// List of name servers.
	DomainNameServers interface{}
	// List of NETBIOS name servers.
	NetbiosNameServers interface{}
	// The NetBIOS node type (1, 2, 4, or 8). For more information about these node types, see [RFC 2132](http://www.ietf.org/rfc/rfc2132.txt).
	NetbiosNodeType interface{}
	// List of NTP servers.
	NtpServers interface{}
	// The ID of the AWS account that owns the DHCP options set.
	OwnerId interface{}
	// A mapping of tags assigned to the resource.
	Tags interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
