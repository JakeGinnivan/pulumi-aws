// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ec2
{
    public static partial class Invokes
    {
        /// <summary>
        /// Provides information about a Launch Template.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/launch_template.html.markdown.
        /// </summary>
        [Obsolete("Use GetLaunchTemplate.InvokeAsync() instead")]
        public static Task<GetLaunchTemplateResult> GetLaunchTemplate(GetLaunchTemplateArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetLaunchTemplateResult>("aws:ec2/getLaunchTemplate:getLaunchTemplate", args ?? InvokeArgs.Empty, options.WithVersion());
    }
    public static class GetLaunchTemplate
    {
        /// <summary>
        /// Provides information about a Launch Template.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/launch_template.html.markdown.
        /// </summary>
        public static Task<GetLaunchTemplateResult> InvokeAsync(GetLaunchTemplateArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetLaunchTemplateResult>("aws:ec2/getLaunchTemplate:getLaunchTemplate", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetLaunchTemplateArgs : Pulumi.InvokeArgs
    {
        [Input("filters")]
        private List<Inputs.GetLaunchTemplateFiltersArgs>? _filters;

        /// <summary>
        /// Configuration block(s) for filtering. Detailed below.
        /// </summary>
        public List<Inputs.GetLaunchTemplateFiltersArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetLaunchTemplateFiltersArgs>());
            set => _filters = value;
        }

        /// <summary>
        /// The name of the filter field. Valid values can be found in the [EC2 DescribeLaunchTemplates API Reference](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeLaunchTemplates.html).
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        [Input("tags")]
        private Dictionary<string, object>? _tags;

        /// <summary>
        /// A mapping of tags, each pair of which must exactly match a pair on the desired Launch Template.
        /// </summary>
        public Dictionary<string, object> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, object>());
            set => _tags = value;
        }

        public GetLaunchTemplateArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetLaunchTemplateResult
    {
        /// <summary>
        /// Amazon Resource Name (ARN) of the launch template.
        /// </summary>
        public readonly string Arn;
        /// <summary>
        /// Specify volumes to attach to the instance besides the volumes specified by the AMI.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetLaunchTemplateBlockDeviceMappingsResult> BlockDeviceMappings;
        /// <summary>
        /// Customize the credit specification of the instance. See Credit
        /// Specification below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetLaunchTemplateCreditSpecificationsResult> CreditSpecifications;
        /// <summary>
        /// The default version of the launch template.
        /// </summary>
        public readonly int DefaultVersion;
        /// <summary>
        /// Description of the launch template.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// If `true`, enables [EC2 Instance
        /// Termination Protection](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/terminating-instances.html#Using_ChangingDisableAPITermination)
        /// </summary>
        public readonly bool DisableApiTermination;
        /// <summary>
        /// If `true`, the launched EC2 instance will be EBS-optimized.
        /// </summary>
        public readonly string EbsOptimized;
        /// <summary>
        /// The elastic GPU to attach to the instance. See Elastic GPU
        /// below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetLaunchTemplateElasticGpuSpecificationsResult> ElasticGpuSpecifications;
        public readonly ImmutableArray<Outputs.GetLaunchTemplateFiltersResult> Filters;
        /// <summary>
        /// The IAM Instance Profile to launch the instance with. See Instance Profile
        /// below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetLaunchTemplateIamInstanceProfilesResult> IamInstanceProfiles;
        /// <summary>
        /// The AMI from which to launch the instance.
        /// </summary>
        public readonly string ImageId;
        /// <summary>
        /// Shutdown behavior for the instance. Can be `stop` or `terminate`.
        /// (Default: `stop`).
        /// </summary>
        public readonly string InstanceInitiatedShutdownBehavior;
        /// <summary>
        /// The market (purchasing) option for the instance.
        /// below for details.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetLaunchTemplateInstanceMarketOptionsResult> InstanceMarketOptions;
        /// <summary>
        /// The type of the instance.
        /// </summary>
        public readonly string InstanceType;
        /// <summary>
        /// The kernel ID.
        /// </summary>
        public readonly string KernelId;
        /// <summary>
        /// The key name to use for the instance.
        /// </summary>
        public readonly string KeyName;
        /// <summary>
        /// The latest version of the launch template.
        /// </summary>
        public readonly int LatestVersion;
        /// <summary>
        /// The metadata options for the instance.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetLaunchTemplateMetadataOptionsResult> MetadataOptions;
        /// <summary>
        /// The monitoring option for the instance.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetLaunchTemplateMonitoringsResult> Monitorings;
        public readonly string? Name;
        /// <summary>
        /// Customize network interfaces to be attached at instance boot time. See Network
        /// Interfaces below for more details.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetLaunchTemplateNetworkInterfacesResult> NetworkInterfaces;
        /// <summary>
        /// The placement of the instance.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetLaunchTemplatePlacementsResult> Placements;
        /// <summary>
        /// The ID of the RAM disk.
        /// </summary>
        public readonly string RamDiskId;
        /// <summary>
        /// A list of security group names to associate with. If you are creating Instances in a VPC, use
        /// `vpc_security_group_ids` instead.
        /// </summary>
        public readonly ImmutableArray<string> SecurityGroupNames;
        /// <summary>
        /// The tags to apply to the resources during launch.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetLaunchTemplateTagSpecificationsResult> TagSpecifications;
        /// <summary>
        /// (Optional) A mapping of tags to assign to the launch template.
        /// </summary>
        public readonly ImmutableDictionary<string, object> Tags;
        /// <summary>
        /// The Base64-encoded user data to provide when launching the instance.
        /// </summary>
        public readonly string UserData;
        /// <summary>
        /// A list of security group IDs to associate with.
        /// </summary>
        public readonly ImmutableArray<string> VpcSecurityGroupIds;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetLaunchTemplateResult(
            string arn,
            ImmutableArray<Outputs.GetLaunchTemplateBlockDeviceMappingsResult> blockDeviceMappings,
            ImmutableArray<Outputs.GetLaunchTemplateCreditSpecificationsResult> creditSpecifications,
            int defaultVersion,
            string description,
            bool disableApiTermination,
            string ebsOptimized,
            ImmutableArray<Outputs.GetLaunchTemplateElasticGpuSpecificationsResult> elasticGpuSpecifications,
            ImmutableArray<Outputs.GetLaunchTemplateFiltersResult> filters,
            ImmutableArray<Outputs.GetLaunchTemplateIamInstanceProfilesResult> iamInstanceProfiles,
            string imageId,
            string instanceInitiatedShutdownBehavior,
            ImmutableArray<Outputs.GetLaunchTemplateInstanceMarketOptionsResult> instanceMarketOptions,
            string instanceType,
            string kernelId,
            string keyName,
            int latestVersion,
            ImmutableArray<Outputs.GetLaunchTemplateMetadataOptionsResult> metadataOptions,
            ImmutableArray<Outputs.GetLaunchTemplateMonitoringsResult> monitorings,
            string? name,
            ImmutableArray<Outputs.GetLaunchTemplateNetworkInterfacesResult> networkInterfaces,
            ImmutableArray<Outputs.GetLaunchTemplatePlacementsResult> placements,
            string ramDiskId,
            ImmutableArray<string> securityGroupNames,
            ImmutableArray<Outputs.GetLaunchTemplateTagSpecificationsResult> tagSpecifications,
            ImmutableDictionary<string, object> tags,
            string userData,
            ImmutableArray<string> vpcSecurityGroupIds,
            string id)
        {
            Arn = arn;
            BlockDeviceMappings = blockDeviceMappings;
            CreditSpecifications = creditSpecifications;
            DefaultVersion = defaultVersion;
            Description = description;
            DisableApiTermination = disableApiTermination;
            EbsOptimized = ebsOptimized;
            ElasticGpuSpecifications = elasticGpuSpecifications;
            Filters = filters;
            IamInstanceProfiles = iamInstanceProfiles;
            ImageId = imageId;
            InstanceInitiatedShutdownBehavior = instanceInitiatedShutdownBehavior;
            InstanceMarketOptions = instanceMarketOptions;
            InstanceType = instanceType;
            KernelId = kernelId;
            KeyName = keyName;
            LatestVersion = latestVersion;
            MetadataOptions = metadataOptions;
            Monitorings = monitorings;
            Name = name;
            NetworkInterfaces = networkInterfaces;
            Placements = placements;
            RamDiskId = ramDiskId;
            SecurityGroupNames = securityGroupNames;
            TagSpecifications = tagSpecifications;
            Tags = tags;
            UserData = userData;
            VpcSecurityGroupIds = vpcSecurityGroupIds;
            Id = id;
        }
    }

    namespace Inputs
    {

    public sealed class GetLaunchTemplateFiltersArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the filter field. Valid values can be found in the [EC2 DescribeLaunchTemplates API Reference](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeLaunchTemplates.html).
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        [Input("values", required: true)]
        private List<string>? _values;

        /// <summary>
        /// Set of values that are accepted for the given filter field. Results will be selected if any given value matches.
        /// </summary>
        public List<string> Values
        {
            get => _values ?? (_values = new List<string>());
            set => _values = value;
        }

        public GetLaunchTemplateFiltersArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetLaunchTemplateBlockDeviceMappingsEbsResult
    {
        public readonly string DeleteOnTermination;
        public readonly string Encrypted;
        public readonly int Iops;
        public readonly string KmsKeyId;
        public readonly string SnapshotId;
        public readonly int VolumeSize;
        public readonly string VolumeType;

        [OutputConstructor]
        private GetLaunchTemplateBlockDeviceMappingsEbsResult(
            string deleteOnTermination,
            string encrypted,
            int iops,
            string kmsKeyId,
            string snapshotId,
            int volumeSize,
            string volumeType)
        {
            DeleteOnTermination = deleteOnTermination;
            Encrypted = encrypted;
            Iops = iops;
            KmsKeyId = kmsKeyId;
            SnapshotId = snapshotId;
            VolumeSize = volumeSize;
            VolumeType = volumeType;
        }
    }

    [OutputType]
    public sealed class GetLaunchTemplateBlockDeviceMappingsResult
    {
        public readonly string DeviceName;
        public readonly ImmutableArray<GetLaunchTemplateBlockDeviceMappingsEbsResult> Ebs;
        public readonly string NoDevice;
        public readonly string VirtualName;

        [OutputConstructor]
        private GetLaunchTemplateBlockDeviceMappingsResult(
            string deviceName,
            ImmutableArray<GetLaunchTemplateBlockDeviceMappingsEbsResult> ebs,
            string noDevice,
            string virtualName)
        {
            DeviceName = deviceName;
            Ebs = ebs;
            NoDevice = noDevice;
            VirtualName = virtualName;
        }
    }

    [OutputType]
    public sealed class GetLaunchTemplateCreditSpecificationsResult
    {
        public readonly string CpuCredits;

        [OutputConstructor]
        private GetLaunchTemplateCreditSpecificationsResult(string cpuCredits)
        {
            CpuCredits = cpuCredits;
        }
    }

    [OutputType]
    public sealed class GetLaunchTemplateElasticGpuSpecificationsResult
    {
        public readonly string Type;

        [OutputConstructor]
        private GetLaunchTemplateElasticGpuSpecificationsResult(string type)
        {
            Type = type;
        }
    }

    [OutputType]
    public sealed class GetLaunchTemplateFiltersResult
    {
        /// <summary>
        /// The name of the filter field. Valid values can be found in the [EC2 DescribeLaunchTemplates API Reference](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeLaunchTemplates.html).
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Set of values that are accepted for the given filter field. Results will be selected if any given value matches.
        /// </summary>
        public readonly ImmutableArray<string> Values;

        [OutputConstructor]
        private GetLaunchTemplateFiltersResult(
            string name,
            ImmutableArray<string> values)
        {
            Name = name;
            Values = values;
        }
    }

    [OutputType]
    public sealed class GetLaunchTemplateIamInstanceProfilesResult
    {
        /// <summary>
        /// Amazon Resource Name (ARN) of the launch template.
        /// </summary>
        public readonly string Arn;
        /// <summary>
        /// The name of the filter field. Valid values can be found in the [EC2 DescribeLaunchTemplates API Reference](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeLaunchTemplates.html).
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private GetLaunchTemplateIamInstanceProfilesResult(
            string arn,
            string name)
        {
            Arn = arn;
            Name = name;
        }
    }

    [OutputType]
    public sealed class GetLaunchTemplateInstanceMarketOptionsResult
    {
        public readonly string MarketType;
        public readonly ImmutableArray<GetLaunchTemplateInstanceMarketOptionsSpotOptionsResult> SpotOptions;

        [OutputConstructor]
        private GetLaunchTemplateInstanceMarketOptionsResult(
            string marketType,
            ImmutableArray<GetLaunchTemplateInstanceMarketOptionsSpotOptionsResult> spotOptions)
        {
            MarketType = marketType;
            SpotOptions = spotOptions;
        }
    }

    [OutputType]
    public sealed class GetLaunchTemplateInstanceMarketOptionsSpotOptionsResult
    {
        public readonly int BlockDurationMinutes;
        public readonly string InstanceInterruptionBehavior;
        public readonly string MaxPrice;
        public readonly string SpotInstanceType;
        public readonly string ValidUntil;

        [OutputConstructor]
        private GetLaunchTemplateInstanceMarketOptionsSpotOptionsResult(
            int blockDurationMinutes,
            string instanceInterruptionBehavior,
            string maxPrice,
            string spotInstanceType,
            string validUntil)
        {
            BlockDurationMinutes = blockDurationMinutes;
            InstanceInterruptionBehavior = instanceInterruptionBehavior;
            MaxPrice = maxPrice;
            SpotInstanceType = spotInstanceType;
            ValidUntil = validUntil;
        }
    }

    [OutputType]
    public sealed class GetLaunchTemplateMetadataOptionsResult
    {
        /// <summary>
        /// The state of the metadata service: `enabled`, `disabled`.
        /// </summary>
        public readonly string HttpEndpoint;
        /// <summary>
        /// The desired HTTP PUT response hop limit for instance metadata requests.
        /// </summary>
        public readonly int HttpPutResponseHopLimit;
        /// <summary>
        /// If session tokens are required: `optional`, `required`.
        /// </summary>
        public readonly string HttpTokens;

        [OutputConstructor]
        private GetLaunchTemplateMetadataOptionsResult(
            string httpEndpoint,
            int httpPutResponseHopLimit,
            string httpTokens)
        {
            HttpEndpoint = httpEndpoint;
            HttpPutResponseHopLimit = httpPutResponseHopLimit;
            HttpTokens = httpTokens;
        }
    }

    [OutputType]
    public sealed class GetLaunchTemplateMonitoringsResult
    {
        public readonly bool Enabled;

        [OutputConstructor]
        private GetLaunchTemplateMonitoringsResult(bool enabled)
        {
            Enabled = enabled;
        }
    }

    [OutputType]
    public sealed class GetLaunchTemplateNetworkInterfacesResult
    {
        public readonly bool AssociatePublicIpAddress;
        public readonly bool DeleteOnTermination;
        /// <summary>
        /// Description of the launch template.
        /// </summary>
        public readonly string Description;
        public readonly int DeviceIndex;
        public readonly int Ipv4AddressCount;
        public readonly ImmutableArray<string> Ipv4Addresses;
        public readonly int Ipv6AddressCount;
        public readonly ImmutableArray<string> Ipv6Addresses;
        public readonly string NetworkInterfaceId;
        public readonly string PrivateIpAddress;
        public readonly ImmutableArray<string> SecurityGroups;
        public readonly string SubnetId;

        [OutputConstructor]
        private GetLaunchTemplateNetworkInterfacesResult(
            bool associatePublicIpAddress,
            bool deleteOnTermination,
            string description,
            int deviceIndex,
            int ipv4AddressCount,
            ImmutableArray<string> ipv4Addresses,
            int ipv6AddressCount,
            ImmutableArray<string> ipv6Addresses,
            string networkInterfaceId,
            string privateIpAddress,
            ImmutableArray<string> securityGroups,
            string subnetId)
        {
            AssociatePublicIpAddress = associatePublicIpAddress;
            DeleteOnTermination = deleteOnTermination;
            Description = description;
            DeviceIndex = deviceIndex;
            Ipv4AddressCount = ipv4AddressCount;
            Ipv4Addresses = ipv4Addresses;
            Ipv6AddressCount = ipv6AddressCount;
            Ipv6Addresses = ipv6Addresses;
            NetworkInterfaceId = networkInterfaceId;
            PrivateIpAddress = privateIpAddress;
            SecurityGroups = securityGroups;
            SubnetId = subnetId;
        }
    }

    [OutputType]
    public sealed class GetLaunchTemplatePlacementsResult
    {
        public readonly string Affinity;
        public readonly string AvailabilityZone;
        public readonly string GroupName;
        public readonly string HostId;
        public readonly string SpreadDomain;
        public readonly string Tenancy;

        [OutputConstructor]
        private GetLaunchTemplatePlacementsResult(
            string affinity,
            string availabilityZone,
            string groupName,
            string hostId,
            string spreadDomain,
            string tenancy)
        {
            Affinity = affinity;
            AvailabilityZone = availabilityZone;
            GroupName = groupName;
            HostId = hostId;
            SpreadDomain = spreadDomain;
            Tenancy = tenancy;
        }
    }

    [OutputType]
    public sealed class GetLaunchTemplateTagSpecificationsResult
    {
        public readonly string ResourceType;
        /// <summary>
        /// A mapping of tags, each pair of which must exactly match a pair on the desired Launch Template.
        /// </summary>
        public readonly ImmutableDictionary<string, object> Tags;

        [OutputConstructor]
        private GetLaunchTemplateTagSpecificationsResult(
            string resourceType,
            ImmutableDictionary<string, object> tags)
        {
            ResourceType = resourceType;
            Tags = tags;
        }
    }
    }
}
