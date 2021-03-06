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
    public sealed class BranchProtectionRequiredStatusCheck
    {
        public readonly ImmutableArray<string> Contexts;
        public readonly bool? Strict;

        [OutputConstructor]
        private BranchProtectionRequiredStatusCheck(
            ImmutableArray<string> contexts,

            bool? strict)
        {
            Contexts = contexts;
            Strict = strict;
        }
    }
}
