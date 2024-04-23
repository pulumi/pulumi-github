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
    public sealed class OrganizationRulesetConditionsRepositoryName
    {
        /// <summary>
        /// Array of repository names or patterns to exclude. The condition will not pass if any of these patterns match.
        /// </summary>
        public readonly ImmutableArray<string> Excludes;
        /// <summary>
        /// Array of repository names or patterns to include. One of these patterns must match for the condition to pass. Also accepts `~ALL` to include all repositories.
        /// </summary>
        public readonly ImmutableArray<string> Includes;
        /// <summary>
        /// Whether renaming of target repositories is prevented.
        /// </summary>
        public readonly bool? Protected;

        [OutputConstructor]
        private OrganizationRulesetConditionsRepositoryName(
            ImmutableArray<string> excludes,

            ImmutableArray<string> includes,

            bool? @protected)
        {
            Excludes = excludes;
            Includes = includes;
            Protected = @protected;
        }
    }
}
