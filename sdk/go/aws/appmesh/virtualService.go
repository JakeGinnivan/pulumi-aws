// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appmesh

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an AWS App Mesh virtual service resource.
type VirtualService struct {
	pulumi.CustomResourceState

	// The ARN of the virtual service.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The creation date of the virtual service.
	CreatedDate pulumi.StringOutput `pulumi:"createdDate"`
	// The last update date of the virtual service.
	LastUpdatedDate pulumi.StringOutput `pulumi:"lastUpdatedDate"`
	// The name of the service mesh in which to create the virtual service.
	MeshName pulumi.StringOutput `pulumi:"meshName"`
	// The name to use for the virtual service.
	Name pulumi.StringOutput `pulumi:"name"`
	// The virtual service specification to apply.
	Spec VirtualServiceSpecOutput `pulumi:"spec"`
	// A map of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
}

// NewVirtualService registers a new resource with the given unique name, arguments, and options.
func NewVirtualService(ctx *pulumi.Context,
	name string, args *VirtualServiceArgs, opts ...pulumi.ResourceOption) (*VirtualService, error) {
	if args == nil || args.MeshName == nil {
		return nil, errors.New("missing required argument 'MeshName'")
	}
	if args == nil || args.Spec == nil {
		return nil, errors.New("missing required argument 'Spec'")
	}
	if args == nil {
		args = &VirtualServiceArgs{}
	}
	var resource VirtualService
	err := ctx.RegisterResource("aws:appmesh/virtualService:VirtualService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualService gets an existing VirtualService resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualServiceState, opts ...pulumi.ResourceOption) (*VirtualService, error) {
	var resource VirtualService
	err := ctx.ReadResource("aws:appmesh/virtualService:VirtualService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualService resources.
type virtualServiceState struct {
	// The ARN of the virtual service.
	Arn *string `pulumi:"arn"`
	// The creation date of the virtual service.
	CreatedDate *string `pulumi:"createdDate"`
	// The last update date of the virtual service.
	LastUpdatedDate *string `pulumi:"lastUpdatedDate"`
	// The name of the service mesh in which to create the virtual service.
	MeshName *string `pulumi:"meshName"`
	// The name to use for the virtual service.
	Name *string `pulumi:"name"`
	// The virtual service specification to apply.
	Spec *VirtualServiceSpec `pulumi:"spec"`
	// A map of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
}

type VirtualServiceState struct {
	// The ARN of the virtual service.
	Arn pulumi.StringPtrInput
	// The creation date of the virtual service.
	CreatedDate pulumi.StringPtrInput
	// The last update date of the virtual service.
	LastUpdatedDate pulumi.StringPtrInput
	// The name of the service mesh in which to create the virtual service.
	MeshName pulumi.StringPtrInput
	// The name to use for the virtual service.
	Name pulumi.StringPtrInput
	// The virtual service specification to apply.
	Spec VirtualServiceSpecPtrInput
	// A map of tags to assign to the resource.
	Tags pulumi.MapInput
}

func (VirtualServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualServiceState)(nil)).Elem()
}

type virtualServiceArgs struct {
	// The name of the service mesh in which to create the virtual service.
	MeshName string `pulumi:"meshName"`
	// The name to use for the virtual service.
	Name *string `pulumi:"name"`
	// The virtual service specification to apply.
	Spec VirtualServiceSpec `pulumi:"spec"`
	// A map of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a VirtualService resource.
type VirtualServiceArgs struct {
	// The name of the service mesh in which to create the virtual service.
	MeshName pulumi.StringInput
	// The name to use for the virtual service.
	Name pulumi.StringPtrInput
	// The virtual service specification to apply.
	Spec VirtualServiceSpecInput
	// A map of tags to assign to the resource.
	Tags pulumi.MapInput
}

func (VirtualServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualServiceArgs)(nil)).Elem()
}
