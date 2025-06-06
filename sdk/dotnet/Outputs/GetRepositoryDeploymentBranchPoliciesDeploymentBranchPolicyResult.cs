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
    public sealed class GetRepositoryDeploymentBranchPoliciesDeploymentBranchPolicyResult
    {
        /// <summary>
        /// Id of the policy.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name pattern that branches must match in order to deploy to the environment.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private GetRepositoryDeploymentBranchPoliciesDeploymentBranchPolicyResult(
            string id,

            string name)
        {
            Id = id;
            Name = name;
        }
    }
}
