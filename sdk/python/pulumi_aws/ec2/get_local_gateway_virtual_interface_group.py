# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetLocalGatewayVirtualInterfaceGroupResult:
    """
    A collection of values returned by getLocalGatewayVirtualInterfaceGroup.
    """
    def __init__(__self__, filters=None, id=None, local_gateway_id=None, local_gateway_virtual_interface_ids=None, tags=None):
        if filters and not isinstance(filters, list):
            raise TypeError("Expected argument 'filters' to be a list")
        __self__.filters = filters
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        if local_gateway_id and not isinstance(local_gateway_id, str):
            raise TypeError("Expected argument 'local_gateway_id' to be a str")
        __self__.local_gateway_id = local_gateway_id
        if local_gateway_virtual_interface_ids and not isinstance(local_gateway_virtual_interface_ids, list):
            raise TypeError("Expected argument 'local_gateway_virtual_interface_ids' to be a list")
        __self__.local_gateway_virtual_interface_ids = local_gateway_virtual_interface_ids
        """
        Set of EC2 Local Gateway Virtual Interface identifiers.
        """
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        __self__.tags = tags
class AwaitableGetLocalGatewayVirtualInterfaceGroupResult(GetLocalGatewayVirtualInterfaceGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetLocalGatewayVirtualInterfaceGroupResult(
            filters=self.filters,
            id=self.id,
            local_gateway_id=self.local_gateway_id,
            local_gateway_virtual_interface_ids=self.local_gateway_virtual_interface_ids,
            tags=self.tags)

def get_local_gateway_virtual_interface_group(filters=None,id=None,local_gateway_id=None,tags=None,opts=None):
    """
    Provides details about an EC2 Local Gateway Virtual Interface Group. More information can be found in the [Outposts User Guide](https://docs.aws.amazon.com/outposts/latest/userguide/outposts-networking-components.html#routing).

    ## Example Usage



    ```python
    import pulumi
    import pulumi_aws as aws

    example = aws.ec2.get_local_gateway_virtual_interface_group(local_gateway_id=data["aws_ec2_local_gateway"]["example"]["id"])
    ```


    :param list filters: One or more configuration blocks containing name-values filters. See the [EC2 API Reference](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeLocalGatewayVirtualInterfaceGroups.html) for supported filters. Detailed below.
    :param str id: Identifier of EC2 Local Gateway Virtual Interface Group.
    :param str local_gateway_id: Identifier of EC2 Local Gateway.
    :param dict tags: Key-value map of resource tags, each pair of which must exactly match a pair on the desired local gateway route table.

    The **filters** object supports the following:

      * `name` (`str`) - Name of the filter.
      * `values` (`list`) - List of one or more values for the filter.
    """
    __args__ = dict()


    __args__['filters'] = filters
    __args__['id'] = id
    __args__['localGatewayId'] = local_gateway_id
    __args__['tags'] = tags
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:ec2/getLocalGatewayVirtualInterfaceGroup:getLocalGatewayVirtualInterfaceGroup', __args__, opts=opts).value

    return AwaitableGetLocalGatewayVirtualInterfaceGroupResult(
        filters=__ret__.get('filters'),
        id=__ret__.get('id'),
        local_gateway_id=__ret__.get('localGatewayId'),
        local_gateway_virtual_interface_ids=__ret__.get('localGatewayVirtualInterfaceIds'),
        tags=__ret__.get('tags'))
