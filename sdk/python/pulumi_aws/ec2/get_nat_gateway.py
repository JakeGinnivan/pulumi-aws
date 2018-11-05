# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class GetNatGatewayResult(object):
    """
    A collection of values returned by getNatGateway.
    """
    def __init__(__self__, allocation_id=None, id=None, network_interface_id=None, private_ip=None, public_ip=None, state=None, subnet_id=None, tags=None, vpc_id=None):
        if allocation_id and not isinstance(allocation_id, str):
            raise TypeError('Expected argument allocation_id to be a str')
        __self__.allocation_id = allocation_id
        """
        The Id of the EIP allocated to the selected Nat Gateway.
        """
        if id and not isinstance(id, str):
            raise TypeError('Expected argument id to be a str')
        __self__.id = id
        if network_interface_id and not isinstance(network_interface_id, str):
            raise TypeError('Expected argument network_interface_id to be a str')
        __self__.network_interface_id = network_interface_id
        """
        The Id of the ENI allocated to the selected Nat Gateway.
        """
        if private_ip and not isinstance(private_ip, str):
            raise TypeError('Expected argument private_ip to be a str')
        __self__.private_ip = private_ip
        """
        The private Ip address of the selected Nat Gateway.
        """
        if public_ip and not isinstance(public_ip, str):
            raise TypeError('Expected argument public_ip to be a str')
        __self__.public_ip = public_ip
        """
        The public Ip (EIP) address of the selected Nat Gateway.
        """
        if state and not isinstance(state, str):
            raise TypeError('Expected argument state to be a str')
        __self__.state = state
        if subnet_id and not isinstance(subnet_id, str):
            raise TypeError('Expected argument subnet_id to be a str')
        __self__.subnet_id = subnet_id
        if tags and not isinstance(tags, dict):
            raise TypeError('Expected argument tags to be a dict')
        __self__.tags = tags
        if vpc_id and not isinstance(vpc_id, str):
            raise TypeError('Expected argument vpc_id to be a str')
        __self__.vpc_id = vpc_id

def get_nat_gateway(filters=None, id=None, state=None, subnet_id=None, tags=None, vpc_id=None):
    """
    Provides details about a specific Nat Gateway.
    """
    __args__ = dict()

    __args__['filters'] = filters
    __args__['id'] = id
    __args__['state'] = state
    __args__['subnetId'] = subnet_id
    __args__['tags'] = tags
    __args__['vpcId'] = vpc_id
    __ret__ = pulumi.runtime.invoke('aws:ec2/getNatGateway:getNatGateway', __args__)

    return GetNatGatewayResult(
        allocation_id=__ret__.get('allocationId'),
        id=__ret__.get('id'),
        network_interface_id=__ret__.get('networkInterfaceId'),
        private_ip=__ret__.get('privateIp'),
        public_ip=__ret__.get('publicIp'),
        state=__ret__.get('state'),
        subnet_id=__ret__.get('subnetId'),
        tags=__ret__.get('tags'),
        vpc_id=__ret__.get('vpcId'))
