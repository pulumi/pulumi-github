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
    public sealed class RepositoryCollaboratorsIgnoreTeam
    {
        /// <summary>
        /// ID or slug of the team to ignore.
        /// </summary>
        public readonly string TeamId;

        [OutputConstructor]
        private RepositoryCollaboratorsIgnoreTeam(string teamId)
        {
            TeamId = teamId;
        }
    }
}