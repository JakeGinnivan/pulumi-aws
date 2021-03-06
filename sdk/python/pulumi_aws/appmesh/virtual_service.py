# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class VirtualService(pulumi.CustomResource):
    arn: pulumi.Output[str]
    """
    The ARN of the virtual service.
    """
    created_date: pulumi.Output[str]
    """
    The creation date of the virtual service.
    """
    last_updated_date: pulumi.Output[str]
    """
    The last update date of the virtual service.
    """
    mesh_name: pulumi.Output[str]
    """
    The name of the service mesh in which to create the virtual service.
    """
    name: pulumi.Output[str]
    """
    The name to use for the virtual service.
    """
    spec: pulumi.Output[dict]
    """
    The virtual service specification to apply.

      * `provider` (`dict`) - The App Mesh object that is acting as the provider for a virtual service. You can specify a single virtual node or virtual router.
        * `virtualNode` (`dict`) - The virtual node associated with a virtual service.
          * `virtualNodeName` (`str`) - The name of the virtual node that is acting as a service provider.

        * `virtualRouter` (`dict`) - The virtual router associated with a virtual service.
          * `virtual_router_name` (`str`) - The name of the virtual router that is acting as a service provider.
    """
    tags: pulumi.Output[dict]
    """
    A map of tags to assign to the resource.
    """
    def __init__(__self__, resource_name, opts=None, mesh_name=None, name=None, spec=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides an AWS App Mesh virtual service resource.

        ## Example Usage

        ### Virtual Node Provider

        ```python
        import pulumi
        import pulumi_aws as aws

        servicea = aws.appmesh.VirtualService("servicea",
            mesh_name=aws_appmesh_mesh["simple"]["id"],
            spec={
                "provider": {
                    "virtualNode": {
                        "virtualNodeName": aws_appmesh_virtual_node["serviceb1"]["name"],
                    },
                },
            })
        ```

        ### Virtual Router Provider

        ```python
        import pulumi
        import pulumi_aws as aws

        servicea = aws.appmesh.VirtualService("servicea",
            mesh_name=aws_appmesh_mesh["simple"]["id"],
            spec={
                "provider": {
                    "virtualRouter": {
                        "virtual_router_name": aws_appmesh_virtual_router["serviceb"]["name"],
                    },
                },
            })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] mesh_name: The name of the service mesh in which to create the virtual service.
        :param pulumi.Input[str] name: The name to use for the virtual service.
        :param pulumi.Input[dict] spec: The virtual service specification to apply.
        :param pulumi.Input[dict] tags: A map of tags to assign to the resource.

        The **spec** object supports the following:

          * `provider` (`pulumi.Input[dict]`) - The App Mesh object that is acting as the provider for a virtual service. You can specify a single virtual node or virtual router.
            * `virtualNode` (`pulumi.Input[dict]`) - The virtual node associated with a virtual service.
              * `virtualNodeName` (`pulumi.Input[str]`) - The name of the virtual node that is acting as a service provider.

            * `virtualRouter` (`pulumi.Input[dict]`) - The virtual router associated with a virtual service.
              * `virtual_router_name` (`pulumi.Input[str]`) - The name of the virtual router that is acting as a service provider.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if mesh_name is None:
                raise TypeError("Missing required property 'mesh_name'")
            __props__['mesh_name'] = mesh_name
            __props__['name'] = name
            if spec is None:
                raise TypeError("Missing required property 'spec'")
            __props__['spec'] = spec
            __props__['tags'] = tags
            __props__['arn'] = None
            __props__['created_date'] = None
            __props__['last_updated_date'] = None
        super(VirtualService, __self__).__init__(
            'aws:appmesh/virtualService:VirtualService',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, arn=None, created_date=None, last_updated_date=None, mesh_name=None, name=None, spec=None, tags=None):
        """
        Get an existing VirtualService resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The ARN of the virtual service.
        :param pulumi.Input[str] created_date: The creation date of the virtual service.
        :param pulumi.Input[str] last_updated_date: The last update date of the virtual service.
        :param pulumi.Input[str] mesh_name: The name of the service mesh in which to create the virtual service.
        :param pulumi.Input[str] name: The name to use for the virtual service.
        :param pulumi.Input[dict] spec: The virtual service specification to apply.
        :param pulumi.Input[dict] tags: A map of tags to assign to the resource.

        The **spec** object supports the following:

          * `provider` (`pulumi.Input[dict]`) - The App Mesh object that is acting as the provider for a virtual service. You can specify a single virtual node or virtual router.
            * `virtualNode` (`pulumi.Input[dict]`) - The virtual node associated with a virtual service.
              * `virtualNodeName` (`pulumi.Input[str]`) - The name of the virtual node that is acting as a service provider.

            * `virtualRouter` (`pulumi.Input[dict]`) - The virtual router associated with a virtual service.
              * `virtual_router_name` (`pulumi.Input[str]`) - The name of the virtual router that is acting as a service provider.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["created_date"] = created_date
        __props__["last_updated_date"] = last_updated_date
        __props__["mesh_name"] = mesh_name
        __props__["name"] = name
        __props__["spec"] = spec
        __props__["tags"] = tags
        return VirtualService(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

