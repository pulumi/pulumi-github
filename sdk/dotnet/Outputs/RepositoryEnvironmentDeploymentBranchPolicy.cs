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
    public sealed class RepositoryEnvironmentDeploymentBranchPolicy
    {
        /// <summary>
        /// Whether only branches that match the specified name patterns can deploy to this environment.
        /// </summary>
        public readonly bool CustomBranchPolicies;
        /// <summary>
        /// Whether only branches with branch protection rules can deploy to this environment.
        /// </summary>
        public readonly bool ProtectedBranches;

        [OutputConstructor]
        private RepositoryEnvironmentDeploymentBranchPolicy(
            bool customBranchPolicies,

            bool protectedBranches)
        {
            CustomBranchPolicies = customBranchPolicies;
            ProtectedBranches = protectedBranches;
        }
    }
}
