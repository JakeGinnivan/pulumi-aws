// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Macie.Outputs
{

    [OutputType]
    public sealed class S3BucketAssociationClassificationType
    {
        /// <summary>
        /// A string value indicating that Macie perform a one-time classification of all of the existing objects in the bucket.
        /// The only valid value is the default value, `FULL`.
        /// </summary>
        public readonly string? Continuous;
        /// <summary>
        /// A string value indicating whether or not Macie performs a one-time classification of all of the existing objects in the bucket.
        /// Valid values are `NONE` and `FULL`. Defaults to `NONE` indicating that Macie only classifies objects that are added after the association was created.
        /// </summary>
        public readonly string? OneTime;

        [OutputConstructor]
        private S3BucketAssociationClassificationType(
            string? continuous,

            string? oneTime)
        {
            Continuous = continuous;
            OneTime = oneTime;
        }
    }
}
