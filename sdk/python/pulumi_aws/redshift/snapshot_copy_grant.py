# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class SnapshotCopyGrant(pulumi.CustomResource):
    """
    Creates a snapshot copy grant that allows AWS Redshift to encrypt copied snapshots with a customer master key from AWS KMS in a destination region.
    
    Note that the grant must exist in the destination region, and not in the region of the cluster.
    """
    def __init__(__self__, __name__, __opts__=None, kms_key_id=None, snapshot_copy_grant_name=None, tags=None):
        """Create a SnapshotCopyGrant resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['kmsKeyId'] = kms_key_id

        if not snapshot_copy_grant_name:
            raise TypeError('Missing required property snapshot_copy_grant_name')
        __props__['snapshotCopyGrantName'] = snapshot_copy_grant_name

        __props__['tags'] = tags

        super(SnapshotCopyGrant, __self__).__init__(
            'aws:redshift/snapshotCopyGrant:SnapshotCopyGrant',
            __name__,
            __props__,
            __opts__)

