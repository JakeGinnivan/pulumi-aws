# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class MysqlLayer(pulumi.CustomResource):
    """
    Provides an OpsWorks MySQL layer resource.
    
    ~> **Note:** All arguments including the root password will be stored in the raw state as plain-text.
    [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).
    """
    def __init__(__self__, __name__, __opts__=None, auto_assign_elastic_ips=None, auto_assign_public_ips=None, auto_healing=None, custom_configure_recipes=None, custom_deploy_recipes=None, custom_instance_profile_arn=None, custom_json=None, custom_security_group_ids=None, custom_setup_recipes=None, custom_shutdown_recipes=None, custom_undeploy_recipes=None, drain_elb_on_shutdown=None, ebs_volumes=None, elastic_load_balancer=None, install_updates_on_boot=None, instance_shutdown_timeout=None, name=None, root_password=None, root_password_on_all_instances=None, stack_id=None, system_packages=None, use_ebs_optimized_instances=None):
        """Create a MysqlLayer resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['autoAssignElasticIps'] = auto_assign_elastic_ips

        __props__['autoAssignPublicIps'] = auto_assign_public_ips

        __props__['autoHealing'] = auto_healing

        __props__['customConfigureRecipes'] = custom_configure_recipes

        __props__['customDeployRecipes'] = custom_deploy_recipes

        __props__['customInstanceProfileArn'] = custom_instance_profile_arn

        __props__['customJson'] = custom_json

        __props__['customSecurityGroupIds'] = custom_security_group_ids

        __props__['customSetupRecipes'] = custom_setup_recipes

        __props__['customShutdownRecipes'] = custom_shutdown_recipes

        __props__['customUndeployRecipes'] = custom_undeploy_recipes

        __props__['drainElbOnShutdown'] = drain_elb_on_shutdown

        __props__['ebsVolumes'] = ebs_volumes

        __props__['elasticLoadBalancer'] = elastic_load_balancer

        __props__['installUpdatesOnBoot'] = install_updates_on_boot

        __props__['instanceShutdownTimeout'] = instance_shutdown_timeout

        __props__['name'] = name

        __props__['rootPassword'] = root_password

        __props__['rootPasswordOnAllInstances'] = root_password_on_all_instances

        if not stack_id:
            raise TypeError('Missing required property stack_id')
        __props__['stackId'] = stack_id

        __props__['systemPackages'] = system_packages

        __props__['useEbsOptimizedInstances'] = use_ebs_optimized_instances

        super(MysqlLayer, __self__).__init__(
            'aws:opsworks/mysqlLayer:MysqlLayer',
            __name__,
            __props__,
            __opts__)

