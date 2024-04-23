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
    public sealed class RepositoryRulesetRulesRequiredStatusChecksRequiredCheck
    {
        /// <summary>
        /// The status check context name that must be present on the commit.
        /// </summary>
        public readonly string Context;
        /// <summary>
        /// The optional integration ID that this status check must originate from.
        /// </summary>
        public readonly int? IntegrationId;

        [OutputConstructor]
        private RepositoryRulesetRulesRequiredStatusChecksRequiredCheck(
            string context,

            int? integrationId)
        {
            Context = context;
            IntegrationId = integrationId;
        }
    }
}
