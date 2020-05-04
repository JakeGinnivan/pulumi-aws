// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ec2
{
    public static class GetVpcPeeringConnection
    {
        /// <summary>
        /// The VPC Peering Connection data source provides details about
        /// a specific VPC peering connection.
        /// 
        /// {{% examples %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetVpcPeeringConnectionResult> InvokeAsync(GetVpcPeeringConnectionArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetVpcPeeringConnectionResult>("aws:ec2/getVpcPeeringConnection:getVpcPeeringConnection", args ?? new GetVpcPeeringConnectionArgs(), options.WithVersion());
    }


    public sealed class GetVpcPeeringConnectionArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The CIDR block of the requester VPC of the specific VPC Peering Connection to retrieve.
        /// </summary>
        [Input("cidrBlock")]
        public string? CidrBlock { get; set; }

        [Input("filters")]
        private List<Inputs.GetVpcPeeringConnectionFilterArgs>? _filters;

        /// <summary>
        /// Custom filter block as described below.
        /// </summary>
        public List<Inputs.GetVpcPeeringConnectionFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetVpcPeeringConnectionFilterArgs>());
            set => _filters = value;
        }

        /// <summary>
        /// The ID of the specific VPC Peering Connection to retrieve.
        /// </summary>
        [Input("id")]
        public string? Id { get; set; }

        /// <summary>
        /// The AWS account ID of the owner of the requester VPC of the specific VPC Peering Connection to retrieve.
        /// </summary>
        [Input("ownerId")]
        public string? OwnerId { get; set; }

        /// <summary>
        /// The CIDR block of the accepter VPC of the specific VPC Peering Connection to retrieve.
        /// </summary>
        [Input("peerCidrBlock")]
        public string? PeerCidrBlock { get; set; }

        /// <summary>
        /// The AWS account ID of the owner of the accepter VPC of the specific VPC Peering Connection to retrieve.
        /// </summary>
        [Input("peerOwnerId")]
        public string? PeerOwnerId { get; set; }

        /// <summary>
        /// The region of the accepter VPC of the specific VPC Peering Connection to retrieve.
        /// </summary>
        [Input("peerRegion")]
        public string? PeerRegion { get; set; }

        /// <summary>
        /// The ID of the accepter VPC of the specific VPC Peering Connection to retrieve.
        /// </summary>
        [Input("peerVpcId")]
        public string? PeerVpcId { get; set; }

        /// <summary>
        /// The region of the requester VPC of the specific VPC Peering Connection to retrieve.
        /// </summary>
        [Input("region")]
        public string? Region { get; set; }

        /// <summary>
        /// The status of the specific VPC Peering Connection to retrieve.
        /// </summary>
        [Input("status")]
        public string? Status { get; set; }

        [Input("tags")]
        private Dictionary<string, object>? _tags;

        /// <summary>
        /// A map of tags, each pair of which must exactly match
        /// a pair on the desired VPC Peering Connection.
        /// </summary>
        public Dictionary<string, object> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, object>());
            set => _tags = value;
        }

        /// <summary>
        /// The ID of the requester VPC of the specific VPC Peering Connection to retrieve.
        /// </summary>
        [Input("vpcId")]
        public string? VpcId { get; set; }

        public GetVpcPeeringConnectionArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetVpcPeeringConnectionResult
    {
        /// <summary>
        /// A configuration block that describes [VPC Peering Connection]
        /// (http://docs.aws.amazon.com/AmazonVPC/latest/PeeringGuide) options set for the accepter VPC.
        /// </summary>
        public readonly ImmutableDictionary<string, bool> Accepter;
        public readonly string CidrBlock;
        public readonly ImmutableArray<Outputs.GetVpcPeeringConnectionFilterResult> Filters;
        public readonly string Id;
        public readonly string OwnerId;
        public readonly string PeerCidrBlock;
        public readonly string PeerOwnerId;
        public readonly string PeerRegion;
        public readonly string PeerVpcId;
        public readonly string Region;
        /// <summary>
        /// A configuration block that describes [VPC Peering Connection]
        /// (http://docs.aws.amazon.com/AmazonVPC/latest/PeeringGuide) options set for the requester VPC.
        /// </summary>
        public readonly ImmutableDictionary<string, bool> Requester;
        public readonly string Status;
        public readonly ImmutableDictionary<string, object> Tags;
        public readonly string VpcId;

        [OutputConstructor]
        private GetVpcPeeringConnectionResult(
            ImmutableDictionary<string, bool> accepter,

            string cidrBlock,

            ImmutableArray<Outputs.GetVpcPeeringConnectionFilterResult> filters,

            string id,

            string ownerId,

            string peerCidrBlock,

            string peerOwnerId,

            string peerRegion,

            string peerVpcId,

            string region,

            ImmutableDictionary<string, bool> requester,

            string status,

            ImmutableDictionary<string, object> tags,

            string vpcId)
        {
            Accepter = accepter;
            CidrBlock = cidrBlock;
            Filters = filters;
            Id = id;
            OwnerId = ownerId;
            PeerCidrBlock = peerCidrBlock;
            PeerOwnerId = peerOwnerId;
            PeerRegion = peerRegion;
            PeerVpcId = peerVpcId;
            Region = region;
            Requester = requester;
            Status = status;
            Tags = tags;
            VpcId = vpcId;
        }
    }
}
