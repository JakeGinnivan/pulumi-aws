# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime

class EgressOnlyInternetGateway(pulumi.CustomResource):
    """
    [IPv6 only] Creates an egress-only Internet gateway for your VPC.
    An egress-only Internet gateway is used to enable outbound communication
    over IPv6 from instances in your VPC to the Internet, and prevents hosts
    outside of your VPC from initiating an IPv6 connection with your instance.
    """
    def __init__(__self__, __name__, __opts__=None, vpc_id=None):
        """Create a EgressOnlyInternetGateway resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, basestring):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not vpc_id:
            raise TypeError('Missing required property vpc_id')
        elif not isinstance(vpc_id, basestring):
            raise TypeError('Expected property vpc_id to be a basestring')
        __self__.vpc_id = vpc_id
        """
        The VPC ID to create in.
        """
        __props__['vpcId'] = vpc_id

        super(EgressOnlyInternetGateway, __self__).__init__(
            'aws:ec2/egressOnlyInternetGateway:EgressOnlyInternetGateway',
            __name__,
            __props__,
            __opts__)

    def set_outputs(self, outs):
        if 'vpcId' in outs:
            self.vpc_id = outs['vpcId']