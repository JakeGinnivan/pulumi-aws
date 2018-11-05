# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class Function(pulumi.CustomResource):
    """
    Provides a Lambda Function resource. Lambda allows you to trigger execution of code in response to events in AWS. The Lambda Function itself includes source code and runtime configuration.
    
    For information about Lambda and how to use it, see [What is AWS Lambda?][1]
    """
    def __init__(__self__, __name__, __opts__=None, dead_letter_config=None, description=None, environment=None, code=None, name=None, handler=None, kms_key_arn=None, memory_size=None, publish=None, reserved_concurrent_executions=None, role=None, runtime=None, s3_bucket=None, s3_key=None, s3_object_version=None, source_code_hash=None, tags=None, timeout=None, tracing_config=None, vpc_config=None):
        """Create a Function resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['deadLetterConfig'] = dead_letter_config

        __props__['description'] = description

        __props__['environment'] = environment

        __props__['code'] = code

        __props__['name'] = name

        if not handler:
            raise TypeError('Missing required property handler')
        __props__['handler'] = handler

        __props__['kmsKeyArn'] = kms_key_arn

        __props__['memorySize'] = memory_size

        __props__['publish'] = publish

        __props__['reservedConcurrentExecutions'] = reserved_concurrent_executions

        if not role:
            raise TypeError('Missing required property role')
        __props__['role'] = role

        if not runtime:
            raise TypeError('Missing required property runtime')
        __props__['runtime'] = runtime

        __props__['s3Bucket'] = s3_bucket

        __props__['s3Key'] = s3_key

        __props__['s3ObjectVersion'] = s3_object_version

        __props__['sourceCodeHash'] = source_code_hash

        __props__['tags'] = tags

        __props__['timeout'] = timeout

        __props__['tracingConfig'] = tracing_config

        __props__['vpcConfig'] = vpc_config

        __props__['arn'] = None
        __props__['invoke_arn'] = None
        __props__['last_modified'] = None
        __props__['qualified_arn'] = None
        __props__['source_code_size'] = None
        __props__['version'] = None

        super(Function, __self__).__init__(
            'aws:lambda/function:Function',
            __name__,
            __props__,
            __opts__)

