// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// This resource can be useful for getting back a list of route table ids to be referenced elsewhere.
func GetRouteTables(ctx *pulumi.Context, args *GetRouteTablesArgs, opts ...pulumi.InvokeOption) (*GetRouteTablesResult, error) {
	var rv GetRouteTablesResult
	err := ctx.Invoke("aws:ec2/getRouteTables:getRouteTables", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRouteTables.
type GetRouteTablesArgs struct {
	// Custom filter block as described below.
	Filters []GetRouteTablesFilter `pulumi:"filters"`
	// A map of tags, each pair of which must exactly match
	// a pair on the desired route tables.
	Tags map[string]interface{} `pulumi:"tags"`
	// The VPC ID that you want to filter from.
	VpcId *string `pulumi:"vpcId"`
}

// A collection of values returned by getRouteTables.
type GetRouteTablesResult struct {
	Filters []GetRouteTablesFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A set of all the route table ids found. This data source will fail if none are found.
	Ids   []string               `pulumi:"ids"`
	Tags  map[string]interface{} `pulumi:"tags"`
	VpcId *string                `pulumi:"vpcId"`
}
