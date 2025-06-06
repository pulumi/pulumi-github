// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github.Outputs
{

    [OutputType]
    public sealed class GetDependabotOrganizationSecretsSecretResult
    {
        /// <summary>
        /// Timestamp of the secret creation
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// Secret name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Timestamp of the secret last update
        /// </summary>
        public readonly string UpdatedAt;
        /// <summary>
        /// Secret visibility
        /// </summary>
        public readonly string Visibility;

        [OutputConstructor]
        private GetDependabotOrganizationSecretsSecretResult(
            string createdAt,

            string name,

            string updatedAt,

            string visibility)
        {
            CreatedAt = createdAt;
            Name = name;
            UpdatedAt = updatedAt;
            Visibility = visibility;
        }
    }
}
