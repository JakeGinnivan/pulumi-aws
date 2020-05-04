// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storagegateway

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an AWS Storage Gateway cached iSCSI volume.
//
// > **NOTE:** The gateway must have cache added (e.g. via the [`storagegateway.Cache`](https://www.terraform.io/docs/providers/aws/r/storagegateway_cache.html) resource) before creating volumes otherwise the Storage Gateway API will return an error.
//
// > **NOTE:** The gateway must have an upload buffer added (e.g. via the [`storagegateway.UploadBuffer`](https://www.terraform.io/docs/providers/aws/r/storagegateway_upload_buffer.html) resource) before the volume is operational to clients, however the Storage Gateway API will allow volume creation without error in that case and return volume status as `UPLOAD BUFFER NOT CONFIGURED`.
type CachesIscsiVolume struct {
	pulumi.CustomResourceState

	// Volume Amazon Resource Name (ARN), e.g. `arn:aws:storagegateway:us-east-1:123456789012:gateway/sgw-12345678/volume/vol-12345678`.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Whether mutual CHAP is enabled for the iSCSI target.
	ChapEnabled pulumi.BoolOutput `pulumi:"chapEnabled"`
	// The Amazon Resource Name (ARN) of the gateway.
	GatewayArn pulumi.StringOutput `pulumi:"gatewayArn"`
	// Logical disk number.
	LunNumber pulumi.IntOutput `pulumi:"lunNumber"`
	// The network interface of the gateway on which to expose the iSCSI target. Only IPv4 addresses are accepted.
	NetworkInterfaceId pulumi.StringOutput `pulumi:"networkInterfaceId"`
	// The port used to communicate with iSCSI targets.
	NetworkInterfacePort pulumi.IntOutput `pulumi:"networkInterfacePort"`
	// The snapshot ID of the snapshot to restore as the new cached volume. e.g. `snap-1122aabb`.
	SnapshotId pulumi.StringPtrOutput `pulumi:"snapshotId"`
	// The ARN for an existing volume. Specifying this ARN makes the new volume into an exact copy of the specified existing volume's latest recovery point. The `volumeSizeInBytes` value for this new volume must be equal to or larger than the size of the existing volume, in bytes.
	SourceVolumeArn pulumi.StringPtrOutput `pulumi:"sourceVolumeArn"`
	// Key-value map of resource tags
	Tags pulumi.MapOutput `pulumi:"tags"`
	// Target Amazon Resource Name (ARN), e.g. `arn:aws:storagegateway:us-east-1:123456789012:gateway/sgw-12345678/target/iqn.1997-05.com.amazon:TargetName`.
	TargetArn pulumi.StringOutput `pulumi:"targetArn"`
	// The name of the iSCSI target used by initiators to connect to the target and as a suffix for the target ARN. The target name must be unique across all volumes of a gateway.
	TargetName pulumi.StringOutput `pulumi:"targetName"`
	// Volume Amazon Resource Name (ARN), e.g. `arn:aws:storagegateway:us-east-1:123456789012:gateway/sgw-12345678/volume/vol-12345678`.
	VolumeArn pulumi.StringOutput `pulumi:"volumeArn"`
	// Volume ID, e.g. `vol-12345678`.
	VolumeId pulumi.StringOutput `pulumi:"volumeId"`
	// The size of the volume in bytes.
	VolumeSizeInBytes pulumi.IntOutput `pulumi:"volumeSizeInBytes"`
}

// NewCachesIscsiVolume registers a new resource with the given unique name, arguments, and options.
func NewCachesIscsiVolume(ctx *pulumi.Context,
	name string, args *CachesIscsiVolumeArgs, opts ...pulumi.ResourceOption) (*CachesIscsiVolume, error) {
	if args == nil || args.GatewayArn == nil {
		return nil, errors.New("missing required argument 'GatewayArn'")
	}
	if args == nil || args.NetworkInterfaceId == nil {
		return nil, errors.New("missing required argument 'NetworkInterfaceId'")
	}
	if args == nil || args.TargetName == nil {
		return nil, errors.New("missing required argument 'TargetName'")
	}
	if args == nil || args.VolumeSizeInBytes == nil {
		return nil, errors.New("missing required argument 'VolumeSizeInBytes'")
	}
	if args == nil {
		args = &CachesIscsiVolumeArgs{}
	}
	var resource CachesIscsiVolume
	err := ctx.RegisterResource("aws:storagegateway/cachesIscsiVolume:CachesIscsiVolume", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCachesIscsiVolume gets an existing CachesIscsiVolume resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCachesIscsiVolume(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CachesIscsiVolumeState, opts ...pulumi.ResourceOption) (*CachesIscsiVolume, error) {
	var resource CachesIscsiVolume
	err := ctx.ReadResource("aws:storagegateway/cachesIscsiVolume:CachesIscsiVolume", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CachesIscsiVolume resources.
type cachesIscsiVolumeState struct {
	// Volume Amazon Resource Name (ARN), e.g. `arn:aws:storagegateway:us-east-1:123456789012:gateway/sgw-12345678/volume/vol-12345678`.
	Arn *string `pulumi:"arn"`
	// Whether mutual CHAP is enabled for the iSCSI target.
	ChapEnabled *bool `pulumi:"chapEnabled"`
	// The Amazon Resource Name (ARN) of the gateway.
	GatewayArn *string `pulumi:"gatewayArn"`
	// Logical disk number.
	LunNumber *int `pulumi:"lunNumber"`
	// The network interface of the gateway on which to expose the iSCSI target. Only IPv4 addresses are accepted.
	NetworkInterfaceId *string `pulumi:"networkInterfaceId"`
	// The port used to communicate with iSCSI targets.
	NetworkInterfacePort *int `pulumi:"networkInterfacePort"`
	// The snapshot ID of the snapshot to restore as the new cached volume. e.g. `snap-1122aabb`.
	SnapshotId *string `pulumi:"snapshotId"`
	// The ARN for an existing volume. Specifying this ARN makes the new volume into an exact copy of the specified existing volume's latest recovery point. The `volumeSizeInBytes` value for this new volume must be equal to or larger than the size of the existing volume, in bytes.
	SourceVolumeArn *string `pulumi:"sourceVolumeArn"`
	// Key-value map of resource tags
	Tags map[string]interface{} `pulumi:"tags"`
	// Target Amazon Resource Name (ARN), e.g. `arn:aws:storagegateway:us-east-1:123456789012:gateway/sgw-12345678/target/iqn.1997-05.com.amazon:TargetName`.
	TargetArn *string `pulumi:"targetArn"`
	// The name of the iSCSI target used by initiators to connect to the target and as a suffix for the target ARN. The target name must be unique across all volumes of a gateway.
	TargetName *string `pulumi:"targetName"`
	// Volume Amazon Resource Name (ARN), e.g. `arn:aws:storagegateway:us-east-1:123456789012:gateway/sgw-12345678/volume/vol-12345678`.
	VolumeArn *string `pulumi:"volumeArn"`
	// Volume ID, e.g. `vol-12345678`.
	VolumeId *string `pulumi:"volumeId"`
	// The size of the volume in bytes.
	VolumeSizeInBytes *int `pulumi:"volumeSizeInBytes"`
}

type CachesIscsiVolumeState struct {
	// Volume Amazon Resource Name (ARN), e.g. `arn:aws:storagegateway:us-east-1:123456789012:gateway/sgw-12345678/volume/vol-12345678`.
	Arn pulumi.StringPtrInput
	// Whether mutual CHAP is enabled for the iSCSI target.
	ChapEnabled pulumi.BoolPtrInput
	// The Amazon Resource Name (ARN) of the gateway.
	GatewayArn pulumi.StringPtrInput
	// Logical disk number.
	LunNumber pulumi.IntPtrInput
	// The network interface of the gateway on which to expose the iSCSI target. Only IPv4 addresses are accepted.
	NetworkInterfaceId pulumi.StringPtrInput
	// The port used to communicate with iSCSI targets.
	NetworkInterfacePort pulumi.IntPtrInput
	// The snapshot ID of the snapshot to restore as the new cached volume. e.g. `snap-1122aabb`.
	SnapshotId pulumi.StringPtrInput
	// The ARN for an existing volume. Specifying this ARN makes the new volume into an exact copy of the specified existing volume's latest recovery point. The `volumeSizeInBytes` value for this new volume must be equal to or larger than the size of the existing volume, in bytes.
	SourceVolumeArn pulumi.StringPtrInput
	// Key-value map of resource tags
	Tags pulumi.MapInput
	// Target Amazon Resource Name (ARN), e.g. `arn:aws:storagegateway:us-east-1:123456789012:gateway/sgw-12345678/target/iqn.1997-05.com.amazon:TargetName`.
	TargetArn pulumi.StringPtrInput
	// The name of the iSCSI target used by initiators to connect to the target and as a suffix for the target ARN. The target name must be unique across all volumes of a gateway.
	TargetName pulumi.StringPtrInput
	// Volume Amazon Resource Name (ARN), e.g. `arn:aws:storagegateway:us-east-1:123456789012:gateway/sgw-12345678/volume/vol-12345678`.
	VolumeArn pulumi.StringPtrInput
	// Volume ID, e.g. `vol-12345678`.
	VolumeId pulumi.StringPtrInput
	// The size of the volume in bytes.
	VolumeSizeInBytes pulumi.IntPtrInput
}

func (CachesIscsiVolumeState) ElementType() reflect.Type {
	return reflect.TypeOf((*cachesIscsiVolumeState)(nil)).Elem()
}

type cachesIscsiVolumeArgs struct {
	// The Amazon Resource Name (ARN) of the gateway.
	GatewayArn string `pulumi:"gatewayArn"`
	// The network interface of the gateway on which to expose the iSCSI target. Only IPv4 addresses are accepted.
	NetworkInterfaceId string `pulumi:"networkInterfaceId"`
	// The snapshot ID of the snapshot to restore as the new cached volume. e.g. `snap-1122aabb`.
	SnapshotId *string `pulumi:"snapshotId"`
	// The ARN for an existing volume. Specifying this ARN makes the new volume into an exact copy of the specified existing volume's latest recovery point. The `volumeSizeInBytes` value for this new volume must be equal to or larger than the size of the existing volume, in bytes.
	SourceVolumeArn *string `pulumi:"sourceVolumeArn"`
	// Key-value map of resource tags
	Tags map[string]interface{} `pulumi:"tags"`
	// The name of the iSCSI target used by initiators to connect to the target and as a suffix for the target ARN. The target name must be unique across all volumes of a gateway.
	TargetName string `pulumi:"targetName"`
	// The size of the volume in bytes.
	VolumeSizeInBytes int `pulumi:"volumeSizeInBytes"`
}

// The set of arguments for constructing a CachesIscsiVolume resource.
type CachesIscsiVolumeArgs struct {
	// The Amazon Resource Name (ARN) of the gateway.
	GatewayArn pulumi.StringInput
	// The network interface of the gateway on which to expose the iSCSI target. Only IPv4 addresses are accepted.
	NetworkInterfaceId pulumi.StringInput
	// The snapshot ID of the snapshot to restore as the new cached volume. e.g. `snap-1122aabb`.
	SnapshotId pulumi.StringPtrInput
	// The ARN for an existing volume. Specifying this ARN makes the new volume into an exact copy of the specified existing volume's latest recovery point. The `volumeSizeInBytes` value for this new volume must be equal to or larger than the size of the existing volume, in bytes.
	SourceVolumeArn pulumi.StringPtrInput
	// Key-value map of resource tags
	Tags pulumi.MapInput
	// The name of the iSCSI target used by initiators to connect to the target and as a suffix for the target ARN. The target name must be unique across all volumes of a gateway.
	TargetName pulumi.StringInput
	// The size of the volume in bytes.
	VolumeSizeInBytes pulumi.IntInput
}

func (CachesIscsiVolumeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cachesIscsiVolumeArgs)(nil)).Elem()
}
