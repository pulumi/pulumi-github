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
    public sealed class RepositoryRulesetConditionsRefName
    {
        /// <summary>
        /// (List of String) Array of ref names or patterns to exclude. The condition will not pass if any of these patterns match.
        /// </summary>
        public readonly ImmutableArray<string> Excludes;
        /// <summary>
        /// (List of String) Array of ref names or patterns to include. One of these patterns must match for the condition to pass. Also accepts `~DEFAULT_BRANCH` to include the default branch or `~ALL` to include all branches.
        /// </summary>
        public readonly ImmutableArray<string> Includes;

        [OutputConstructor]
        private RepositoryRulesetConditionsRefName(
            ImmutableArray<string> excludes,

            ImmutableArray<string> includes)
        {
            Excludes = excludes;
            Includes = includes;
        }
    }
}
