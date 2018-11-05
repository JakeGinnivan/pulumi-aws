# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class GetExportResult(object):
    """
    A collection of values returned by getExport.
    """
    def __init__(__self__, exporting_stack_id=None, value=None, id=None):
        if exporting_stack_id and not isinstance(exporting_stack_id, str):
            raise TypeError('Expected argument exporting_stack_id to be a str')
        __self__.exporting_stack_id = exporting_stack_id
        """
        The exporting_stack_id (AWS ARNs) equivalent `ExportingStackId` from [list-exports](http://docs.aws.amazon.com/cli/latest/reference/cloudformation/list-exports.html) 
        """
        if value and not isinstance(value, str):
            raise TypeError('Expected argument value to be a str')
        __self__.value = value
        """
        The value from Cloudformation export identified by the export name found from [list-exports](http://docs.aws.amazon.com/cli/latest/reference/cloudformation/list-exports.html)
        """
        if id and not isinstance(id, str):
            raise TypeError('Expected argument id to be a str')
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

def get_export(name=None):
    """
    The CloudFormation Export data source allows access to stack
    exports specified in the [Output](http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/outputs-section-structure.html) section of the Cloudformation Template using the optional Export Property.
    
     -> Note: If you are trying to use a value from a Cloudformation Stack in the same Terraform run please use normal interpolation or Cloudformation Outputs. 
    """
    __args__ = dict()

    __args__['name'] = name
    __ret__ = pulumi.runtime.invoke('aws:cloudformation/getExport:getExport', __args__)

    return GetExportResult(
        exporting_stack_id=__ret__.get('exportingStackId'),
        value=__ret__.get('value'),
        id=__ret__.get('id'))
