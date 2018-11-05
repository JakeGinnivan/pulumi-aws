# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class ApnsVoipChannel(pulumi.CustomResource):
    def __init__(__self__, __name__, __opts__=None, application_id=None, bundle_id=None, certificate=None, default_authentication_method=None, enabled=None, private_key=None, team_id=None, token_key=None, token_key_id=None):
        """Create a ApnsVoipChannel resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not application_id:
            raise TypeError('Missing required property application_id')
        __props__['applicationId'] = application_id

        __props__['bundleId'] = bundle_id

        __props__['certificate'] = certificate

        __props__['defaultAuthenticationMethod'] = default_authentication_method

        __props__['enabled'] = enabled

        __props__['privateKey'] = private_key

        __props__['teamId'] = team_id

        __props__['tokenKey'] = token_key

        __props__['tokenKeyId'] = token_key_id

        super(ApnsVoipChannel, __self__).__init__(
            'aws:pinpoint/apnsVoipChannel:ApnsVoipChannel',
            __name__,
            __props__,
            __opts__)

