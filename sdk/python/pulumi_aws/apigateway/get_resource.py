# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class GetResourceResult(object):
    """
    A collection of values returned by getResource.
    """
    def __init__(__self__, parent_id=None, path_part=None, id=None):
        if parent_id and not isinstance(parent_id, str):
            raise TypeError('Expected argument parent_id to be a str')
        __self__.parent_id = parent_id
        """
        Set to the ID of the parent Resource.
        """
        if path_part and not isinstance(path_part, str):
            raise TypeError('Expected argument path_part to be a str')
        __self__.path_part = path_part
        """
        Set to the path relative to the parent Resource.
        """
        if id and not isinstance(id, str):
            raise TypeError('Expected argument id to be a str')
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

def get_resource(path=None, rest_api_id=None):
    """
    Use this data source to get the id of a Resource in API Gateway. 
    To fetch the Resource, you must provide the REST API id as well as the full path.  
    """
    __args__ = dict()

    __args__['path'] = path
    __args__['restApiId'] = rest_api_id
    __ret__ = pulumi.runtime.invoke('aws:apigateway/getResource:getResource', __args__)

    return GetResourceResult(
        parent_id=__ret__.get('parentId'),
        path_part=__ret__.get('pathPart'),
        id=__ret__.get('id'))
