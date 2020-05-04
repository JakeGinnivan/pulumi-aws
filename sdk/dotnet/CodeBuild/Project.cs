// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CodeBuild
{
    /// <summary>
    /// Provides a CodeBuild Project resource. See also the [`aws.codebuild.Webhook` resource](https://www.terraform.io/docs/providers/aws/r/codebuild_webhook.html), which manages the webhook to the source (e.g. the "rebuild every time a code change is pushed" option in the CodeBuild web console).
    /// </summary>
    public partial class Project : Pulumi.CustomResource
    {
        /// <summary>
        /// The ARN of the CodeBuild project.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Information about the project's build output artifacts. Artifact blocks are documented below.
        /// </summary>
        [Output("artifacts")]
        public Output<Outputs.ProjectArtifacts> Artifacts { get; private set; } = null!;

        /// <summary>
        /// Generates a publicly-accessible URL for the projects build badge. Available as `badge_url` attribute when enabled.
        /// </summary>
        [Output("badgeEnabled")]
        public Output<bool?> BadgeEnabled { get; private set; } = null!;

        /// <summary>
        /// The URL of the build badge when `badge_enabled` is enabled.
        /// </summary>
        [Output("badgeUrl")]
        public Output<string> BadgeUrl { get; private set; } = null!;

        /// <summary>
        /// How long in minutes, from 5 to 480 (8 hours), for AWS CodeBuild to wait until timing out any related build that does not get marked as completed. The default is 60 minutes.
        /// </summary>
        [Output("buildTimeout")]
        public Output<int?> BuildTimeout { get; private set; } = null!;

        /// <summary>
        /// Information about the cache storage for the project. Cache blocks are documented below.
        /// </summary>
        [Output("cache")]
        public Output<Outputs.ProjectCache?> Cache { get; private set; } = null!;

        /// <summary>
        /// A short description of the project.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// The AWS Key Management Service (AWS KMS) customer master key (CMK) to be used for encrypting the build project's build output artifacts.
        /// </summary>
        [Output("encryptionKey")]
        public Output<string> EncryptionKey { get; private set; } = null!;

        /// <summary>
        /// Information about the project's build environment. Environment blocks are documented below.
        /// </summary>
        [Output("environment")]
        public Output<Outputs.ProjectEnvironment> Environment { get; private set; } = null!;

        /// <summary>
        /// Configuration for the builds to store log data to CloudWatch or S3.
        /// </summary>
        [Output("logsConfig")]
        public Output<Outputs.ProjectLogsConfig?> LogsConfig { get; private set; } = null!;

        /// <summary>
        /// The projects name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// How long in minutes, from 5 to 480 (8 hours), a build is allowed to be queued before it times out. The default is 8 hours.
        /// </summary>
        [Output("queuedTimeout")]
        public Output<int?> QueuedTimeout { get; private set; } = null!;

        /// <summary>
        /// A set of secondary artifacts to be used inside the build. Secondary artifacts blocks are documented below.
        /// </summary>
        [Output("secondaryArtifacts")]
        public Output<ImmutableArray<Outputs.ProjectSecondaryArtifact>> SecondaryArtifacts { get; private set; } = null!;

        /// <summary>
        /// A set of secondary sources to be used inside the build. Secondary sources blocks are documented below.
        /// </summary>
        [Output("secondarySources")]
        public Output<ImmutableArray<Outputs.ProjectSecondarySource>> SecondarySources { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM) role that enables AWS CodeBuild to interact with dependent AWS services on behalf of the AWS account.
        /// </summary>
        [Output("serviceRole")]
        public Output<string> ServiceRole { get; private set; } = null!;

        /// <summary>
        /// Information about the project's input source code. Source blocks are documented below.
        /// </summary>
        [Output("source")]
        public Output<Outputs.ProjectSource> Source { get; private set; } = null!;

        /// <summary>
        /// A version of the build input to be built for this project. If not specified, the latest version is used.
        /// </summary>
        [Output("sourceVersion")]
        public Output<string?> SourceVersion { get; private set; } = null!;

        /// <summary>
        /// A map of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Configuration for the builds to run inside a VPC. VPC config blocks are documented below.
        /// </summary>
        [Output("vpcConfig")]
        public Output<Outputs.ProjectVpcConfig?> VpcConfig { get; private set; } = null!;


        /// <summary>
        /// Create a Project resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Project(string name, ProjectArgs args, CustomResourceOptions? options = null)
            : base("aws:codebuild/project:Project", name, args ?? new ProjectArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Project(string name, Input<string> id, ProjectState? state = null, CustomResourceOptions? options = null)
            : base("aws:codebuild/project:Project", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Project resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Project Get(string name, Input<string> id, ProjectState? state = null, CustomResourceOptions? options = null)
        {
            return new Project(name, id, state, options);
        }
    }

    public sealed class ProjectArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Information about the project's build output artifacts. Artifact blocks are documented below.
        /// </summary>
        [Input("artifacts", required: true)]
        public Input<Inputs.ProjectArtifactsArgs> Artifacts { get; set; } = null!;

        /// <summary>
        /// Generates a publicly-accessible URL for the projects build badge. Available as `badge_url` attribute when enabled.
        /// </summary>
        [Input("badgeEnabled")]
        public Input<bool>? BadgeEnabled { get; set; }

        /// <summary>
        /// How long in minutes, from 5 to 480 (8 hours), for AWS CodeBuild to wait until timing out any related build that does not get marked as completed. The default is 60 minutes.
        /// </summary>
        [Input("buildTimeout")]
        public Input<int>? BuildTimeout { get; set; }

        /// <summary>
        /// Information about the cache storage for the project. Cache blocks are documented below.
        /// </summary>
        [Input("cache")]
        public Input<Inputs.ProjectCacheArgs>? Cache { get; set; }

        /// <summary>
        /// A short description of the project.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The AWS Key Management Service (AWS KMS) customer master key (CMK) to be used for encrypting the build project's build output artifacts.
        /// </summary>
        [Input("encryptionKey")]
        public Input<string>? EncryptionKey { get; set; }

        /// <summary>
        /// Information about the project's build environment. Environment blocks are documented below.
        /// </summary>
        [Input("environment", required: true)]
        public Input<Inputs.ProjectEnvironmentArgs> Environment { get; set; } = null!;

        /// <summary>
        /// Configuration for the builds to store log data to CloudWatch or S3.
        /// </summary>
        [Input("logsConfig")]
        public Input<Inputs.ProjectLogsConfigArgs>? LogsConfig { get; set; }

        /// <summary>
        /// The projects name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// How long in minutes, from 5 to 480 (8 hours), a build is allowed to be queued before it times out. The default is 8 hours.
        /// </summary>
        [Input("queuedTimeout")]
        public Input<int>? QueuedTimeout { get; set; }

        [Input("secondaryArtifacts")]
        private InputList<Inputs.ProjectSecondaryArtifactArgs>? _secondaryArtifacts;

        /// <summary>
        /// A set of secondary artifacts to be used inside the build. Secondary artifacts blocks are documented below.
        /// </summary>
        public InputList<Inputs.ProjectSecondaryArtifactArgs> SecondaryArtifacts
        {
            get => _secondaryArtifacts ?? (_secondaryArtifacts = new InputList<Inputs.ProjectSecondaryArtifactArgs>());
            set => _secondaryArtifacts = value;
        }

        [Input("secondarySources")]
        private InputList<Inputs.ProjectSecondarySourceArgs>? _secondarySources;

        /// <summary>
        /// A set of secondary sources to be used inside the build. Secondary sources blocks are documented below.
        /// </summary>
        public InputList<Inputs.ProjectSecondarySourceArgs> SecondarySources
        {
            get => _secondarySources ?? (_secondarySources = new InputList<Inputs.ProjectSecondarySourceArgs>());
            set => _secondarySources = value;
        }

        /// <summary>
        /// The Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM) role that enables AWS CodeBuild to interact with dependent AWS services on behalf of the AWS account.
        /// </summary>
        [Input("serviceRole", required: true)]
        public Input<string> ServiceRole { get; set; } = null!;

        /// <summary>
        /// Information about the project's input source code. Source blocks are documented below.
        /// </summary>
        [Input("source", required: true)]
        public Input<Inputs.ProjectSourceArgs> Source { get; set; } = null!;

        /// <summary>
        /// A version of the build input to be built for this project. If not specified, the latest version is used.
        /// </summary>
        [Input("sourceVersion")]
        public Input<string>? SourceVersion { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A map of tags to assign to the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// Configuration for the builds to run inside a VPC. VPC config blocks are documented below.
        /// </summary>
        [Input("vpcConfig")]
        public Input<Inputs.ProjectVpcConfigArgs>? VpcConfig { get; set; }

        public ProjectArgs()
        {
        }
    }

    public sealed class ProjectState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN of the CodeBuild project.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// Information about the project's build output artifacts. Artifact blocks are documented below.
        /// </summary>
        [Input("artifacts")]
        public Input<Inputs.ProjectArtifactsGetArgs>? Artifacts { get; set; }

        /// <summary>
        /// Generates a publicly-accessible URL for the projects build badge. Available as `badge_url` attribute when enabled.
        /// </summary>
        [Input("badgeEnabled")]
        public Input<bool>? BadgeEnabled { get; set; }

        /// <summary>
        /// The URL of the build badge when `badge_enabled` is enabled.
        /// </summary>
        [Input("badgeUrl")]
        public Input<string>? BadgeUrl { get; set; }

        /// <summary>
        /// How long in minutes, from 5 to 480 (8 hours), for AWS CodeBuild to wait until timing out any related build that does not get marked as completed. The default is 60 minutes.
        /// </summary>
        [Input("buildTimeout")]
        public Input<int>? BuildTimeout { get; set; }

        /// <summary>
        /// Information about the cache storage for the project. Cache blocks are documented below.
        /// </summary>
        [Input("cache")]
        public Input<Inputs.ProjectCacheGetArgs>? Cache { get; set; }

        /// <summary>
        /// A short description of the project.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The AWS Key Management Service (AWS KMS) customer master key (CMK) to be used for encrypting the build project's build output artifacts.
        /// </summary>
        [Input("encryptionKey")]
        public Input<string>? EncryptionKey { get; set; }

        /// <summary>
        /// Information about the project's build environment. Environment blocks are documented below.
        /// </summary>
        [Input("environment")]
        public Input<Inputs.ProjectEnvironmentGetArgs>? Environment { get; set; }

        /// <summary>
        /// Configuration for the builds to store log data to CloudWatch or S3.
        /// </summary>
        [Input("logsConfig")]
        public Input<Inputs.ProjectLogsConfigGetArgs>? LogsConfig { get; set; }

        /// <summary>
        /// The projects name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// How long in minutes, from 5 to 480 (8 hours), a build is allowed to be queued before it times out. The default is 8 hours.
        /// </summary>
        [Input("queuedTimeout")]
        public Input<int>? QueuedTimeout { get; set; }

        [Input("secondaryArtifacts")]
        private InputList<Inputs.ProjectSecondaryArtifactGetArgs>? _secondaryArtifacts;

        /// <summary>
        /// A set of secondary artifacts to be used inside the build. Secondary artifacts blocks are documented below.
        /// </summary>
        public InputList<Inputs.ProjectSecondaryArtifactGetArgs> SecondaryArtifacts
        {
            get => _secondaryArtifacts ?? (_secondaryArtifacts = new InputList<Inputs.ProjectSecondaryArtifactGetArgs>());
            set => _secondaryArtifacts = value;
        }

        [Input("secondarySources")]
        private InputList<Inputs.ProjectSecondarySourceGetArgs>? _secondarySources;

        /// <summary>
        /// A set of secondary sources to be used inside the build. Secondary sources blocks are documented below.
        /// </summary>
        public InputList<Inputs.ProjectSecondarySourceGetArgs> SecondarySources
        {
            get => _secondarySources ?? (_secondarySources = new InputList<Inputs.ProjectSecondarySourceGetArgs>());
            set => _secondarySources = value;
        }

        /// <summary>
        /// The Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM) role that enables AWS CodeBuild to interact with dependent AWS services on behalf of the AWS account.
        /// </summary>
        [Input("serviceRole")]
        public Input<string>? ServiceRole { get; set; }

        /// <summary>
        /// Information about the project's input source code. Source blocks are documented below.
        /// </summary>
        [Input("source")]
        public Input<Inputs.ProjectSourceGetArgs>? Source { get; set; }

        /// <summary>
        /// A version of the build input to be built for this project. If not specified, the latest version is used.
        /// </summary>
        [Input("sourceVersion")]
        public Input<string>? SourceVersion { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A map of tags to assign to the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// Configuration for the builds to run inside a VPC. VPC config blocks are documented below.
        /// </summary>
        [Input("vpcConfig")]
        public Input<Inputs.ProjectVpcConfigGetArgs>? VpcConfig { get; set; }

        public ProjectState()
        {
        }
    }
}
