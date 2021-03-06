// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package redshift

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Creates a snapshot copy grant that allows AWS Redshift to encrypt copied snapshots with a customer master key from AWS KMS in a destination region.
//
// Note that the grant must exist in the destination region, and not in the region of the cluster.
//
// ## Example Usage
//
//
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/redshift"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		testSnapshotCopyGrant, err := redshift.NewSnapshotCopyGrant(ctx, "testSnapshotCopyGrant", &redshift.SnapshotCopyGrantArgs{
// 			SnapshotCopyGrantName: pulumi.String("my-grant"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		testCluster, err := redshift.NewCluster(ctx, "testCluster", &redshift.ClusterArgs{
// 			SnapshotCopy: &redshift.ClusterSnapshotCopyArgs{
// 				DestinationRegion: pulumi.String("us-east-2"),
// 				GrantName:         testSnapshotCopyGrant.SnapshotCopyGrantName,
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type SnapshotCopyGrant struct {
	pulumi.CustomResourceState

	// Amazon Resource Name (ARN) of snapshot copy grant
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The unique identifier for the customer master key (CMK) that the grant applies to. Specify the key ID or the Amazon Resource Name (ARN) of the CMK. To specify a CMK in a different AWS account, you must use the key ARN. If not specified, the default key is used.
	KmsKeyId pulumi.StringOutput `pulumi:"kmsKeyId"`
	// A friendly name for identifying the grant.
	SnapshotCopyGrantName pulumi.StringOutput `pulumi:"snapshotCopyGrantName"`
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewSnapshotCopyGrant registers a new resource with the given unique name, arguments, and options.
func NewSnapshotCopyGrant(ctx *pulumi.Context,
	name string, args *SnapshotCopyGrantArgs, opts ...pulumi.ResourceOption) (*SnapshotCopyGrant, error) {
	if args == nil || args.SnapshotCopyGrantName == nil {
		return nil, errors.New("missing required argument 'SnapshotCopyGrantName'")
	}
	if args == nil {
		args = &SnapshotCopyGrantArgs{}
	}
	var resource SnapshotCopyGrant
	err := ctx.RegisterResource("aws:redshift/snapshotCopyGrant:SnapshotCopyGrant", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSnapshotCopyGrant gets an existing SnapshotCopyGrant resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSnapshotCopyGrant(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SnapshotCopyGrantState, opts ...pulumi.ResourceOption) (*SnapshotCopyGrant, error) {
	var resource SnapshotCopyGrant
	err := ctx.ReadResource("aws:redshift/snapshotCopyGrant:SnapshotCopyGrant", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SnapshotCopyGrant resources.
type snapshotCopyGrantState struct {
	// Amazon Resource Name (ARN) of snapshot copy grant
	Arn *string `pulumi:"arn"`
	// The unique identifier for the customer master key (CMK) that the grant applies to. Specify the key ID or the Amazon Resource Name (ARN) of the CMK. To specify a CMK in a different AWS account, you must use the key ARN. If not specified, the default key is used.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// A friendly name for identifying the grant.
	SnapshotCopyGrantName *string `pulumi:"snapshotCopyGrantName"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type SnapshotCopyGrantState struct {
	// Amazon Resource Name (ARN) of snapshot copy grant
	Arn pulumi.StringPtrInput
	// The unique identifier for the customer master key (CMK) that the grant applies to. Specify the key ID or the Amazon Resource Name (ARN) of the CMK. To specify a CMK in a different AWS account, you must use the key ARN. If not specified, the default key is used.
	KmsKeyId pulumi.StringPtrInput
	// A friendly name for identifying the grant.
	SnapshotCopyGrantName pulumi.StringPtrInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (SnapshotCopyGrantState) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotCopyGrantState)(nil)).Elem()
}

type snapshotCopyGrantArgs struct {
	// The unique identifier for the customer master key (CMK) that the grant applies to. Specify the key ID or the Amazon Resource Name (ARN) of the CMK. To specify a CMK in a different AWS account, you must use the key ARN. If not specified, the default key is used.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// A friendly name for identifying the grant.
	SnapshotCopyGrantName string `pulumi:"snapshotCopyGrantName"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a SnapshotCopyGrant resource.
type SnapshotCopyGrantArgs struct {
	// The unique identifier for the customer master key (CMK) that the grant applies to. Specify the key ID or the Amazon Resource Name (ARN) of the CMK. To specify a CMK in a different AWS account, you must use the key ARN. If not specified, the default key is used.
	KmsKeyId pulumi.StringPtrInput
	// A friendly name for identifying the grant.
	SnapshotCopyGrantName pulumi.StringInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (SnapshotCopyGrantArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotCopyGrantArgs)(nil)).Elem()
}
