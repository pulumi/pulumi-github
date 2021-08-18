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
    public sealed class RepositoryEnvironmentReviewer
    {
        /// <summary>
        /// Up to 6 IDs for teams who may review jobs that reference the environment. Reviewers must have at least read access to the repository. Only one of the required reviewers needs to approve the job for it to proceed.
        /// </summary>
        public readonly ImmutableArray<int> Teams;
        /// <summary>
        /// Up to 6 IDs for users who may review jobs that reference the environment. Reviewers must have at least read access to the repository. Only one of the required reviewers needs to approve the job for it to proceed.
        /// </summary>
        public readonly ImmutableArray<int> Users;

        [OutputConstructor]
        private RepositoryEnvironmentReviewer(
            ImmutableArray<int> teams,

            ImmutableArray<int> users)
        {
            Teams = teams;
            Users = users;
        }
    }
}
