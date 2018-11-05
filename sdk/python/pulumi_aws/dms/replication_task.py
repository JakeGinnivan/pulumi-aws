# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class ReplicationTask(pulumi.CustomResource):
    """
    Provides a DMS (Data Migration Service) replication task resource. DMS replication tasks can be created, updated, deleted, and imported.
    """
    def __init__(__self__, __name__, __opts__=None, cdc_start_time=None, migration_type=None, replication_instance_arn=None, replication_task_id=None, replication_task_settings=None, source_endpoint_arn=None, table_mappings=None, tags=None, target_endpoint_arn=None):
        """Create a ReplicationTask resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['cdcStartTime'] = cdc_start_time

        if not migration_type:
            raise TypeError('Missing required property migration_type')
        __props__['migrationType'] = migration_type

        if not replication_instance_arn:
            raise TypeError('Missing required property replication_instance_arn')
        __props__['replicationInstanceArn'] = replication_instance_arn

        if not replication_task_id:
            raise TypeError('Missing required property replication_task_id')
        __props__['replicationTaskId'] = replication_task_id

        __props__['replicationTaskSettings'] = replication_task_settings

        if not source_endpoint_arn:
            raise TypeError('Missing required property source_endpoint_arn')
        __props__['sourceEndpointArn'] = source_endpoint_arn

        if not table_mappings:
            raise TypeError('Missing required property table_mappings')
        __props__['tableMappings'] = table_mappings

        __props__['tags'] = tags

        if not target_endpoint_arn:
            raise TypeError('Missing required property target_endpoint_arn')
        __props__['targetEndpointArn'] = target_endpoint_arn

        __props__['replication_task_arn'] = None

        super(ReplicationTask, __self__).__init__(
            'aws:dms/replicationTask:ReplicationTask',
            __name__,
            __props__,
            __opts__)

