# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities, tables

class Endpoint(pulumi.CustomResource):
    """
    Provides a DMS (Data Migration Service) endpoint resource. DMS endpoints can be created, updated, deleted, and imported.
    
    ~> **Note:** All arguments including the password will be stored in the raw state as plain-text.
    [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).
    """
    def __init__(__self__, __name__, __opts__=None, certificate_arn=None, database_name=None, endpoint_id=None, endpoint_type=None, engine_name=None, extra_connection_attributes=None, kms_key_arn=None, mongodb_settings=None, password=None, port=None, s3_settings=None, server_name=None, service_access_role=None, ssl_mode=None, tags=None, username=None):
        """Create a Endpoint resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['certificate_arn'] = certificate_arn

        __props__['database_name'] = database_name

        if not endpoint_id:
            raise TypeError('Missing required property endpoint_id')
        __props__['endpoint_id'] = endpoint_id

        if not endpoint_type:
            raise TypeError('Missing required property endpoint_type')
        __props__['endpoint_type'] = endpoint_type

        if not engine_name:
            raise TypeError('Missing required property engine_name')
        __props__['engine_name'] = engine_name

        __props__['extra_connection_attributes'] = extra_connection_attributes

        __props__['kms_key_arn'] = kms_key_arn

        __props__['mongodb_settings'] = mongodb_settings

        __props__['password'] = password

        __props__['port'] = port

        __props__['s3_settings'] = s3_settings

        __props__['server_name'] = server_name

        __props__['service_access_role'] = service_access_role

        __props__['ssl_mode'] = ssl_mode

        __props__['tags'] = tags

        __props__['username'] = username

        __props__['endpoint_arn'] = None

        super(Endpoint, __self__).__init__(
            'aws:dms/endpoint:Endpoint',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

