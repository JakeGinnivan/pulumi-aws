// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package lightsail

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Allocates a static IP address.
//
// > **Note:** Lightsail is currently only supported in a limited number of AWS Regions, please see ["Regions and Availability Zones in Amazon Lightsail"](https://lightsail.aws.amazon.com/ls/docs/overview/article/understanding-regions-and-availability-zones-in-amazon-lightsail) for more details
//
// ## Example Usage
//
//
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/lightsail"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		test, err := lightsail.NewStaticIp(ctx, "test", nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type StaticIp struct {
	pulumi.CustomResourceState

	// The ARN of the Lightsail static IP
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The allocated static IP address
	IpAddress pulumi.StringOutput `pulumi:"ipAddress"`
	// The name for the allocated static IP
	Name pulumi.StringOutput `pulumi:"name"`
	// The support code.
	SupportCode pulumi.StringOutput `pulumi:"supportCode"`
}

// NewStaticIp registers a new resource with the given unique name, arguments, and options.
func NewStaticIp(ctx *pulumi.Context,
	name string, args *StaticIpArgs, opts ...pulumi.ResourceOption) (*StaticIp, error) {
	if args == nil {
		args = &StaticIpArgs{}
	}
	var resource StaticIp
	err := ctx.RegisterResource("aws:lightsail/staticIp:StaticIp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStaticIp gets an existing StaticIp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStaticIp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StaticIpState, opts ...pulumi.ResourceOption) (*StaticIp, error) {
	var resource StaticIp
	err := ctx.ReadResource("aws:lightsail/staticIp:StaticIp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StaticIp resources.
type staticIpState struct {
	// The ARN of the Lightsail static IP
	Arn *string `pulumi:"arn"`
	// The allocated static IP address
	IpAddress *string `pulumi:"ipAddress"`
	// The name for the allocated static IP
	Name *string `pulumi:"name"`
	// The support code.
	SupportCode *string `pulumi:"supportCode"`
}

type StaticIpState struct {
	// The ARN of the Lightsail static IP
	Arn pulumi.StringPtrInput
	// The allocated static IP address
	IpAddress pulumi.StringPtrInput
	// The name for the allocated static IP
	Name pulumi.StringPtrInput
	// The support code.
	SupportCode pulumi.StringPtrInput
}

func (StaticIpState) ElementType() reflect.Type {
	return reflect.TypeOf((*staticIpState)(nil)).Elem()
}

type staticIpArgs struct {
	// The name for the allocated static IP
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a StaticIp resource.
type StaticIpArgs struct {
	// The name for the allocated static IP
	Name pulumi.StringPtrInput
}

func (StaticIpArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*staticIpArgs)(nil)).Elem()
}
