// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a VPC Endpoint resource.
//
// > **NOTE on VPC Endpoints and VPC Endpoint Associations:** This provider provides both standalone VPC Endpoint Associations for
// Route Tables - (an association between a VPC endpoint and a single `routeTableId`) and
// Subnets - (an association between a VPC endpoint and a single `subnetId`) and
// a VPC Endpoint resource with `routeTableIds` and `subnetIds` attributes.
// Do not use the same resource ID in both a VPC Endpoint resource and a VPC Endpoint Association resource.
// Doing so will cause a conflict of associations and will overwrite the association.
//
// ## Example Usage
//
// ### Basic
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		s3, err := ec2.NewVpcEndpoint(ctx, "s3", &ec2.VpcEndpointArgs{
// 			ServiceName: pulumi.String("com.amazonaws.us-west-2.s3"),
// 			VpcId:       pulumi.String(aws_vpc.Main.Id),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ### Basic w/ Tags
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		s3, err := ec2.NewVpcEndpoint(ctx, "s3", &ec2.VpcEndpointArgs{
// 			ServiceName: pulumi.String("com.amazonaws.us-west-2.s3"),
// 			Tags: map[string]interface{}{
// 				"Environment": "test",
// 			},
// 			VpcId: pulumi.String(aws_vpc.Main.Id),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type VpcEndpoint struct {
	pulumi.CustomResourceState

	// Accept the VPC endpoint (the VPC endpoint and service need to be in the same AWS account).
	AutoAccept pulumi.BoolPtrOutput `pulumi:"autoAccept"`
	// The list of CIDR blocks for the exposed AWS service. Applicable for endpoints of type `Gateway`.
	CidrBlocks pulumi.StringArrayOutput `pulumi:"cidrBlocks"`
	// The DNS entries for the VPC Endpoint. Applicable for endpoints of type `Interface`. DNS blocks are documented below.
	DnsEntries VpcEndpointDnsEntryArrayOutput `pulumi:"dnsEntries"`
	// One or more network interfaces for the VPC Endpoint. Applicable for endpoints of type `Interface`.
	NetworkInterfaceIds pulumi.StringArrayOutput `pulumi:"networkInterfaceIds"`
	// The ID of the AWS account that owns the VPC endpoint.
	OwnerId pulumi.StringOutput `pulumi:"ownerId"`
	// A policy to attach to the endpoint that controls access to the service. Defaults to full access. All `Gateway` and some `Interface` endpoints support policies - see the [relevant AWS documentation](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-endpoints-access.html) for more details.
	Policy pulumi.StringOutput `pulumi:"policy"`
	// The prefix list ID of the exposed AWS service. Applicable for endpoints of type `Gateway`.
	PrefixListId pulumi.StringOutput `pulumi:"prefixListId"`
	// Whether or not to associate a private hosted zone with the specified VPC. Applicable for endpoints of type `Interface`.
	// Defaults to `false`.
	PrivateDnsEnabled pulumi.BoolPtrOutput `pulumi:"privateDnsEnabled"`
	// Whether or not the VPC Endpoint is being managed by its service - `true` or `false`.
	RequesterManaged pulumi.BoolOutput `pulumi:"requesterManaged"`
	// One or more route table IDs. Applicable for endpoints of type `Gateway`.
	RouteTableIds pulumi.StringArrayOutput `pulumi:"routeTableIds"`
	// The ID of one or more security groups to associate with the network interface. Required for endpoints of type `Interface`.
	SecurityGroupIds pulumi.StringArrayOutput `pulumi:"securityGroupIds"`
	// The service name. For AWS services the service name is usually in the form `com.amazonaws.<region>.<service>` (the SageMaker Notebook service is an exception to this rule, the service name is in the form `aws.sagemaker.<region>.notebook`).
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// The state of the VPC endpoint.
	State pulumi.StringOutput `pulumi:"state"`
	// The ID of one or more subnets in which to create a network interface for the endpoint. Applicable for endpoints of type `Interface`.
	SubnetIds pulumi.StringArrayOutput `pulumi:"subnetIds"`
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The VPC endpoint type, `Gateway` or `Interface`. Defaults to `Gateway`.
	VpcEndpointType pulumi.StringPtrOutput `pulumi:"vpcEndpointType"`
	// The ID of the VPC in which the endpoint will be used.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewVpcEndpoint registers a new resource with the given unique name, arguments, and options.
func NewVpcEndpoint(ctx *pulumi.Context,
	name string, args *VpcEndpointArgs, opts ...pulumi.ResourceOption) (*VpcEndpoint, error) {
	if args == nil || args.ServiceName == nil {
		return nil, errors.New("missing required argument 'ServiceName'")
	}
	if args == nil || args.VpcId == nil {
		return nil, errors.New("missing required argument 'VpcId'")
	}
	if args == nil {
		args = &VpcEndpointArgs{}
	}
	var resource VpcEndpoint
	err := ctx.RegisterResource("aws:ec2/vpcEndpoint:VpcEndpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpcEndpoint gets an existing VpcEndpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpcEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpcEndpointState, opts ...pulumi.ResourceOption) (*VpcEndpoint, error) {
	var resource VpcEndpoint
	err := ctx.ReadResource("aws:ec2/vpcEndpoint:VpcEndpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpcEndpoint resources.
type vpcEndpointState struct {
	// Accept the VPC endpoint (the VPC endpoint and service need to be in the same AWS account).
	AutoAccept *bool `pulumi:"autoAccept"`
	// The list of CIDR blocks for the exposed AWS service. Applicable for endpoints of type `Gateway`.
	CidrBlocks []string `pulumi:"cidrBlocks"`
	// The DNS entries for the VPC Endpoint. Applicable for endpoints of type `Interface`. DNS blocks are documented below.
	DnsEntries []VpcEndpointDnsEntry `pulumi:"dnsEntries"`
	// One or more network interfaces for the VPC Endpoint. Applicable for endpoints of type `Interface`.
	NetworkInterfaceIds []string `pulumi:"networkInterfaceIds"`
	// The ID of the AWS account that owns the VPC endpoint.
	OwnerId *string `pulumi:"ownerId"`
	// A policy to attach to the endpoint that controls access to the service. Defaults to full access. All `Gateway` and some `Interface` endpoints support policies - see the [relevant AWS documentation](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-endpoints-access.html) for more details.
	Policy *string `pulumi:"policy"`
	// The prefix list ID of the exposed AWS service. Applicable for endpoints of type `Gateway`.
	PrefixListId *string `pulumi:"prefixListId"`
	// Whether or not to associate a private hosted zone with the specified VPC. Applicable for endpoints of type `Interface`.
	// Defaults to `false`.
	PrivateDnsEnabled *bool `pulumi:"privateDnsEnabled"`
	// Whether or not the VPC Endpoint is being managed by its service - `true` or `false`.
	RequesterManaged *bool `pulumi:"requesterManaged"`
	// One or more route table IDs. Applicable for endpoints of type `Gateway`.
	RouteTableIds []string `pulumi:"routeTableIds"`
	// The ID of one or more security groups to associate with the network interface. Required for endpoints of type `Interface`.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// The service name. For AWS services the service name is usually in the form `com.amazonaws.<region>.<service>` (the SageMaker Notebook service is an exception to this rule, the service name is in the form `aws.sagemaker.<region>.notebook`).
	ServiceName *string `pulumi:"serviceName"`
	// The state of the VPC endpoint.
	State *string `pulumi:"state"`
	// The ID of one or more subnets in which to create a network interface for the endpoint. Applicable for endpoints of type `Interface`.
	SubnetIds []string `pulumi:"subnetIds"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The VPC endpoint type, `Gateway` or `Interface`. Defaults to `Gateway`.
	VpcEndpointType *string `pulumi:"vpcEndpointType"`
	// The ID of the VPC in which the endpoint will be used.
	VpcId *string `pulumi:"vpcId"`
}

type VpcEndpointState struct {
	// Accept the VPC endpoint (the VPC endpoint and service need to be in the same AWS account).
	AutoAccept pulumi.BoolPtrInput
	// The list of CIDR blocks for the exposed AWS service. Applicable for endpoints of type `Gateway`.
	CidrBlocks pulumi.StringArrayInput
	// The DNS entries for the VPC Endpoint. Applicable for endpoints of type `Interface`. DNS blocks are documented below.
	DnsEntries VpcEndpointDnsEntryArrayInput
	// One or more network interfaces for the VPC Endpoint. Applicable for endpoints of type `Interface`.
	NetworkInterfaceIds pulumi.StringArrayInput
	// The ID of the AWS account that owns the VPC endpoint.
	OwnerId pulumi.StringPtrInput
	// A policy to attach to the endpoint that controls access to the service. Defaults to full access. All `Gateway` and some `Interface` endpoints support policies - see the [relevant AWS documentation](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-endpoints-access.html) for more details.
	Policy pulumi.StringPtrInput
	// The prefix list ID of the exposed AWS service. Applicable for endpoints of type `Gateway`.
	PrefixListId pulumi.StringPtrInput
	// Whether or not to associate a private hosted zone with the specified VPC. Applicable for endpoints of type `Interface`.
	// Defaults to `false`.
	PrivateDnsEnabled pulumi.BoolPtrInput
	// Whether or not the VPC Endpoint is being managed by its service - `true` or `false`.
	RequesterManaged pulumi.BoolPtrInput
	// One or more route table IDs. Applicable for endpoints of type `Gateway`.
	RouteTableIds pulumi.StringArrayInput
	// The ID of one or more security groups to associate with the network interface. Required for endpoints of type `Interface`.
	SecurityGroupIds pulumi.StringArrayInput
	// The service name. For AWS services the service name is usually in the form `com.amazonaws.<region>.<service>` (the SageMaker Notebook service is an exception to this rule, the service name is in the form `aws.sagemaker.<region>.notebook`).
	ServiceName pulumi.StringPtrInput
	// The state of the VPC endpoint.
	State pulumi.StringPtrInput
	// The ID of one or more subnets in which to create a network interface for the endpoint. Applicable for endpoints of type `Interface`.
	SubnetIds pulumi.StringArrayInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The VPC endpoint type, `Gateway` or `Interface`. Defaults to `Gateway`.
	VpcEndpointType pulumi.StringPtrInput
	// The ID of the VPC in which the endpoint will be used.
	VpcId pulumi.StringPtrInput
}

func (VpcEndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcEndpointState)(nil)).Elem()
}

type vpcEndpointArgs struct {
	// Accept the VPC endpoint (the VPC endpoint and service need to be in the same AWS account).
	AutoAccept *bool `pulumi:"autoAccept"`
	// A policy to attach to the endpoint that controls access to the service. Defaults to full access. All `Gateway` and some `Interface` endpoints support policies - see the [relevant AWS documentation](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-endpoints-access.html) for more details.
	Policy *string `pulumi:"policy"`
	// Whether or not to associate a private hosted zone with the specified VPC. Applicable for endpoints of type `Interface`.
	// Defaults to `false`.
	PrivateDnsEnabled *bool `pulumi:"privateDnsEnabled"`
	// One or more route table IDs. Applicable for endpoints of type `Gateway`.
	RouteTableIds []string `pulumi:"routeTableIds"`
	// The ID of one or more security groups to associate with the network interface. Required for endpoints of type `Interface`.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// The service name. For AWS services the service name is usually in the form `com.amazonaws.<region>.<service>` (the SageMaker Notebook service is an exception to this rule, the service name is in the form `aws.sagemaker.<region>.notebook`).
	ServiceName string `pulumi:"serviceName"`
	// The ID of one or more subnets in which to create a network interface for the endpoint. Applicable for endpoints of type `Interface`.
	SubnetIds []string `pulumi:"subnetIds"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The VPC endpoint type, `Gateway` or `Interface`. Defaults to `Gateway`.
	VpcEndpointType *string `pulumi:"vpcEndpointType"`
	// The ID of the VPC in which the endpoint will be used.
	VpcId string `pulumi:"vpcId"`
}

// The set of arguments for constructing a VpcEndpoint resource.
type VpcEndpointArgs struct {
	// Accept the VPC endpoint (the VPC endpoint and service need to be in the same AWS account).
	AutoAccept pulumi.BoolPtrInput
	// A policy to attach to the endpoint that controls access to the service. Defaults to full access. All `Gateway` and some `Interface` endpoints support policies - see the [relevant AWS documentation](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-endpoints-access.html) for more details.
	Policy pulumi.StringPtrInput
	// Whether or not to associate a private hosted zone with the specified VPC. Applicable for endpoints of type `Interface`.
	// Defaults to `false`.
	PrivateDnsEnabled pulumi.BoolPtrInput
	// One or more route table IDs. Applicable for endpoints of type `Gateway`.
	RouteTableIds pulumi.StringArrayInput
	// The ID of one or more security groups to associate with the network interface. Required for endpoints of type `Interface`.
	SecurityGroupIds pulumi.StringArrayInput
	// The service name. For AWS services the service name is usually in the form `com.amazonaws.<region>.<service>` (the SageMaker Notebook service is an exception to this rule, the service name is in the form `aws.sagemaker.<region>.notebook`).
	ServiceName pulumi.StringInput
	// The ID of one or more subnets in which to create a network interface for the endpoint. Applicable for endpoints of type `Interface`.
	SubnetIds pulumi.StringArrayInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The VPC endpoint type, `Gateway` or `Interface`. Defaults to `Gateway`.
	VpcEndpointType pulumi.StringPtrInput
	// The ID of the VPC in which the endpoint will be used.
	VpcId pulumi.StringInput
}

func (VpcEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcEndpointArgs)(nil)).Elem()
}
