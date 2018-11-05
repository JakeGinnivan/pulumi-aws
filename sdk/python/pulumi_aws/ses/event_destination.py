# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class EventDestination(pulumi.CustomResource):
    """
    Provides an SES event destination
    """
    def __init__(__self__, __name__, __opts__=None, cloudwatch_destination=None, configuration_set_name=None, enabled=None, kinesis_destination=None, matching_types=None, name=None, sns_destination=None):
        """Create a EventDestination resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['cloudwatchDestination'] = cloudwatch_destination

        if not configuration_set_name:
            raise TypeError('Missing required property configuration_set_name')
        __props__['configurationSetName'] = configuration_set_name

        __props__['enabled'] = enabled

        __props__['kinesisDestination'] = kinesis_destination

        if not matching_types:
            raise TypeError('Missing required property matching_types')
        __props__['matchingTypes'] = matching_types

        __props__['name'] = name

        __props__['snsDestination'] = sns_destination

        super(EventDestination, __self__).__init__(
            'aws:ses/eventDestination:EventDestination',
            __name__,
            __props__,
            __opts__)

