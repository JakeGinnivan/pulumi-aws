# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class AmiFromInstance(pulumi.CustomResource):
    """
    The "AMI from instance" resource allows the creation of an Amazon Machine
    Image (AMI) modelled after an existing EBS-backed EC2 instance.
    
    The created AMI will refer to implicitly-created snapshots of the instance's
    EBS volumes and mimick its assigned block device configuration at the time
    the resource is created.
    
    This resource is best applied to an instance that is stopped when this instance
    is created, so that the contents of the created image are predictable. When
    applied to an instance that is running, *the instance will be stopped before taking
    the snapshots and then started back up again*, resulting in a period of
    downtime.
    
    Note that the source instance is inspected only at the initial creation of this
    resource. Ongoing updates to the referenced instance will not be propagated into
    the generated AMI. Users may taint or otherwise recreate the resource in order
    to produce a fresh snapshot.
    """
    def __init__(__self__, __name__, __opts__=None, description=None, ebs_block_devices=None, ephemeral_block_devices=None, name=None, snapshot_without_reboot=None, source_instance_id=None, tags=None):
        """Create a AmiFromInstance resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['description'] = description

        __props__['ebsBlockDevices'] = ebs_block_devices

        __props__['ephemeralBlockDevices'] = ephemeral_block_devices

        __props__['name'] = name

        __props__['snapshotWithoutReboot'] = snapshot_without_reboot

        if not source_instance_id:
            raise TypeError('Missing required property source_instance_id')
        __props__['sourceInstanceId'] = source_instance_id

        __props__['tags'] = tags

        __props__['architecture'] = None
        __props__['ena_support'] = None
        __props__['image_location'] = None
        __props__['kernel_id'] = None
        __props__['manage_ebs_snapshots'] = None
        __props__['ramdisk_id'] = None
        __props__['root_device_name'] = None
        __props__['root_snapshot_id'] = None
        __props__['sriov_net_support'] = None
        __props__['virtualization_type'] = None

        super(AmiFromInstance, __self__).__init__(
            'aws:ec2/amiFromInstance:AmiFromInstance',
            __name__,
            __props__,
            __opts__)

