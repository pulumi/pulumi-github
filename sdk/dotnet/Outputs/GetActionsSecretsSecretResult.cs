// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github.Outputs
{

    [OutputType]
    public sealed class GetActionsSecretsSecretResult
    {
        /// <summary>
        /// Timestamp of the secret creation
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// The name of the repository.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Timestamp of the secret last update
        /// </summary>
        public readonly string UpdatedAt;

        [OutputConstructor]
        private GetActionsSecretsSecretResult(
            string createdAt,

            string name,

            string updatedAt)
        {
            CreatedAt = createdAt;
            Name = name;
            UpdatedAt = updatedAt;
        }
    }
}
