# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class RouteTableAssociation(pulumi.CustomResource):
    resource_id: pulumi.Output[str]
    """
    Identifier of the resource
    """
    resource_type: pulumi.Output[str]
    """
    Type of the resource
    """
    transit_gateway_attachment_id: pulumi.Output[str]
    """
    Identifier of EC2 Transit Gateway Attachment.
    """
    transit_gateway_route_table_id: pulumi.Output[str]
    """
    Identifier of EC2 Transit Gateway Route Table.
    """
    def __init__(__self__, resource_name, opts=None, transit_gateway_attachment_id=None, transit_gateway_route_table_id=None, __name__=None, __opts__=None):
        """
        Manages an EC2 Transit Gateway Route Table association.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] transit_gateway_attachment_id: Identifier of EC2 Transit Gateway Attachment.
        :param pulumi.Input[str] transit_gateway_route_table_id: Identifier of EC2 Transit Gateway Route Table.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if transit_gateway_attachment_id is None:
            raise TypeError("Missing required property 'transit_gateway_attachment_id'")
        __props__['transit_gateway_attachment_id'] = transit_gateway_attachment_id

        if transit_gateway_route_table_id is None:
            raise TypeError("Missing required property 'transit_gateway_route_table_id'")
        __props__['transit_gateway_route_table_id'] = transit_gateway_route_table_id

        __props__['resource_id'] = None
        __props__['resource_type'] = None

        super(RouteTableAssociation, __self__).__init__(
            'aws:ec2transitgateway/routeTableAssociation:RouteTableAssociation',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

