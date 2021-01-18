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
    public sealed class BranchProtectionV3RequiredStatusChecks
    {
        public readonly ImmutableArray<string> Contexts;
        public readonly bool? IncludeAdmins;
        public readonly bool? Strict;

        [OutputConstructor]
        private BranchProtectionV3RequiredStatusChecks(
            ImmutableArray<string> contexts,

            bool? includeAdmins,

            bool? strict)
        {
            Contexts = contexts;
            IncludeAdmins = includeAdmins;
            Strict = strict;
        }
    }
}
