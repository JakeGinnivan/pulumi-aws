// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Iam.Inputs
{

    /// <summary>
    /// IAM roles that can be assumed by an AWS service are called service roles. Service roles must include a trust policy.
    /// Trust policies are resource-based policies that are attached to a role that define which principals can assume the
    /// role. Some service role have predefined trust policies. However, in some cases, you must specify the service
    /// principal in the trust policy. A service principal is an identifier that is used to grant permissions to a service.
    /// The identifier includes the long version of a service name, e.g. long_service_name.amazonaws.com.The service
    /// principal is defined by the service. To learn the service principal for a service, see the documentation for that
    /// service.
    /// </summary>
    public sealed class ServicePrincipalGetArgs : Pulumi.ResourceArgs
    {
        [Input("Service", required: true)]
        public InputUnion<string, ImmutableArray<string>> Service { get; set; } = null!;

        public ServicePrincipalGetArgs()
        {
        }
    }
}
