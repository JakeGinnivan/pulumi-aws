# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities, tables

class Association(pulumi.CustomResource):
    """
    Associates an SSM Document to an instance or EC2 tag.
    """
    def __init__(__self__, __name__, __opts__=None, association_name=None, document_version=None, instance_id=None, name=None, output_location=None, parameters=None, schedule_expression=None, targets=None):
        """Create a Association resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['association_name'] = association_name

        __props__['document_version'] = document_version

        __props__['instance_id'] = instance_id

        __props__['name'] = name

        __props__['output_location'] = output_location

        __props__['parameters'] = parameters

        __props__['schedule_expression'] = schedule_expression

        __props__['targets'] = targets

        __props__['association_id'] = None

        super(Association, __self__).__init__(
            'aws:ssm/association:Association',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

