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
    public sealed class RepositoryBranch
    {
        /// <summary>
        /// The name of the repository.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Whether the branch is protected.
        /// </summary>
        public readonly bool? Protected;

        [OutputConstructor]
        private RepositoryBranch(
            string? name,

            bool? @protected)
        {
            Name = name;
            Protected = @protected;
        }
    }
}