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
    public sealed class OrganizationRulesetRulesRequiredWorkflows
    {
        /// <summary>
        /// Actions workflows that are required. Several can be defined.
        /// </summary>
        public readonly ImmutableArray<Outputs.OrganizationRulesetRulesRequiredWorkflowsRequiredWorkflow> RequiredWorkflows;

        [OutputConstructor]
        private OrganizationRulesetRulesRequiredWorkflows(ImmutableArray<Outputs.OrganizationRulesetRulesRequiredWorkflowsRequiredWorkflow> requiredWorkflows)
        {
            RequiredWorkflows = requiredWorkflows;
        }
    }
}
