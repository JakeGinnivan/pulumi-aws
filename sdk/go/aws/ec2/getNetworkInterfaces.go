// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ## Example Usage
//
//
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := ec2.LookupNetworkInterfaces(ctx, &ec2.LookupNetworkInterfacesArgs{
// 			Tags: map[string]interface{}{
// 				"Name": "test",
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("example1", example.Ids)
// 		return nil
// 	})
// }
// ```
func GetNetworkInterfaces(ctx *pulumi.Context, args *GetNetworkInterfacesArgs, opts ...pulumi.InvokeOption) (*GetNetworkInterfacesResult, error) {
	var rv GetNetworkInterfacesResult
	err := ctx.Invoke("aws:ec2/getNetworkInterfaces:getNetworkInterfaces", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNetworkInterfaces.
type GetNetworkInterfacesArgs struct {
	// Custom filter block as described below.
	Filters []GetNetworkInterfacesFilter `pulumi:"filters"`
	// A map of tags, each pair of which must exactly match
	// a pair on the desired network interfaces.
	Tags map[string]string `pulumi:"tags"`
}

// A collection of values returned by getNetworkInterfaces.
type GetNetworkInterfacesResult struct {
	Filters []GetNetworkInterfacesFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of all the network interface ids found. This data source will fail if none are found.
	Ids  []string          `pulumi:"ids"`
	Tags map[string]string `pulumi:"tags"`
}
