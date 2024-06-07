// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github.Inputs
{

    public sealed class OrganizationRulesetRulesRequiredWorkflowsGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("requiredWorkflows", required: true)]
        private InputList<Inputs.OrganizationRulesetRulesRequiredWorkflowsRequiredWorkflowGetArgs>? _requiredWorkflows;

        /// <summary>
        /// Actions workflows that are required. Several can be defined.
        /// </summary>
        public InputList<Inputs.OrganizationRulesetRulesRequiredWorkflowsRequiredWorkflowGetArgs> RequiredWorkflows
        {
            get => _requiredWorkflows ?? (_requiredWorkflows = new InputList<Inputs.OrganizationRulesetRulesRequiredWorkflowsRequiredWorkflowGetArgs>());
            set => _requiredWorkflows = value;
        }

        public OrganizationRulesetRulesRequiredWorkflowsGetArgs()
        {
        }
        public static new OrganizationRulesetRulesRequiredWorkflowsGetArgs Empty => new OrganizationRulesetRulesRequiredWorkflowsGetArgs();
    }
}
