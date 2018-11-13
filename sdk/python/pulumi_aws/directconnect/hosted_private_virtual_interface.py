# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities, tables

class HostedPrivateVirtualInterface(pulumi.CustomResource):
    """
    Provides a Direct Connect hosted private virtual interface resource. This resource represents the allocator's side of the hosted virtual interface.
    A hosted virtual interface is a virtual interface that is owned by another AWS account.
    """
    def __init__(__self__, __name__, __opts__=None, address_family=None, amazon_address=None, bgp_asn=None, bgp_auth_key=None, connection_id=None, customer_address=None, mtu=None, name=None, owner_account_id=None, vlan=None):
        """Create a HostedPrivateVirtualInterface resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not address_family:
            raise TypeError('Missing required property address_family')
        __props__['address_family'] = address_family

        __props__['amazon_address'] = amazon_address

        if not bgp_asn:
            raise TypeError('Missing required property bgp_asn')
        __props__['bgp_asn'] = bgp_asn

        __props__['bgp_auth_key'] = bgp_auth_key

        if not connection_id:
            raise TypeError('Missing required property connection_id')
        __props__['connection_id'] = connection_id

        __props__['customer_address'] = customer_address

        __props__['mtu'] = mtu

        __props__['name'] = name

        if not owner_account_id:
            raise TypeError('Missing required property owner_account_id')
        __props__['owner_account_id'] = owner_account_id

        if not vlan:
            raise TypeError('Missing required property vlan')
        __props__['vlan'] = vlan

        __props__['arn'] = None
        __props__['jumbo_frame_capable'] = None

        super(HostedPrivateVirtualInterface, __self__).__init__(
            'aws:directconnect/hostedPrivateVirtualInterface:HostedPrivateVirtualInterface',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

