# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class RolePolicyAttachment(pulumi.CustomResource):
    """
    Attaches a Managed IAM Policy to an IAM role
    """
    def __init__(__self__, __name__, __opts__=None, policy_arn=None, role=None):
        """Create a RolePolicyAttachment resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not policy_arn:
            raise TypeError('Missing required property policy_arn')
        __props__['policyArn'] = policy_arn

        if not role:
            raise TypeError('Missing required property role')
        __props__['role'] = role

        super(RolePolicyAttachment, __self__).__init__(
            'aws:iam/rolePolicyAttachment:RolePolicyAttachment',
            __name__,
            __props__,
            __opts__)

