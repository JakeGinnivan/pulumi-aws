# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class GetResourceShareResult:
    """
    A collection of values returned by getResourceShare.
    """
    def __init__(__self__, arn=None, filters=None, id=None, name=None, resource_owner=None, status=None, tags=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        __self__.arn = arn
        """
        The Amazon Resource Name (ARN) of the resource share.
        """
        if filters and not isinstance(filters, list):
            raise TypeError("Expected argument 'filters' to be a list")
        __self__.filters = filters
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The Amazon Resource Name (ARN) of the resource share.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        if resource_owner and not isinstance(resource_owner, str):
            raise TypeError("Expected argument 'resource_owner' to be a str")
        __self__.resource_owner = resource_owner
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        __self__.status = status
        """
        The Status of the RAM share.
        """
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        __self__.tags = tags
        """
        The Tags attached to the RAM share
        """

async def get_resource_share(filters=None,name=None,resource_owner=None,opts=None):
    """
    `aws_ram_resource_share` Retrieve information about a RAM Resource Share.
    """
    __args__ = dict()

    __args__['filters'] = filters
    __args__['name'] = name
    __args__['resourceOwner'] = resource_owner
    __ret__ = await pulumi.runtime.invoke('aws:ram/getResourceShare:getResourceShare', __args__, opts=opts)

    return GetResourceShareResult(
        arn=__ret__.get('arn'),
        filters=__ret__.get('filters'),
        id=__ret__.get('id'),
        name=__ret__.get('name'),
        resource_owner=__ret__.get('resourceOwner'),
        status=__ret__.get('status'),
        tags=__ret__.get('tags'))
