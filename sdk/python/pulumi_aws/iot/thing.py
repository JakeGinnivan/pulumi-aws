# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Thing(pulumi.CustomResource):
    arn: pulumi.Output[str]
    """
    The ARN of the thing.
    """
    attributes: pulumi.Output[dict]
    """
    Map of attributes of the thing.
    """
    default_client_id: pulumi.Output[str]
    """
    The default client ID.
    """
    name: pulumi.Output[str]
    """
    The name of the thing.
    """
    thing_type_name: pulumi.Output[str]
    """
    The thing type name.
    """
    version: pulumi.Output[float]
    """
    The current version of the thing record in the registry.
    """
    def __init__(__self__, resource_name, opts=None, attributes=None, name=None, thing_type_name=None, __name__=None, __opts__=None):
        """
        Creates and manages an AWS IoT Thing.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] attributes: Map of attributes of the thing.
        :param pulumi.Input[str] name: The name of the thing.
        :param pulumi.Input[str] thing_type_name: The thing type name.
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

        __props__['attributes'] = attributes

        __props__['name'] = name

        __props__['thing_type_name'] = thing_type_name

        __props__['arn'] = None
        __props__['default_client_id'] = None
        __props__['version'] = None

        super(Thing, __self__).__init__(
            'aws:iot/thing:Thing',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

