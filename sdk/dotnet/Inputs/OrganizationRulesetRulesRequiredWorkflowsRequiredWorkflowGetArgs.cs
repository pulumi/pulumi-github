// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github.Inputs
{

    public sealed class OrganizationRulesetRulesRequiredWorkflowsRequiredWorkflowGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The path to the workflow YAML definition file.
        /// </summary>
        [Input("path", required: true)]
        public Input<string> Path { get; set; } = null!;

        /// <summary>
        /// The ref (branch or tag) of the workflow file to use.
        /// </summary>
        [Input("ref")]
        public Input<string>? Ref { get; set; }

        /// <summary>
        /// The repository in which the workflow is defined.
        /// </summary>
        [Input("repositoryId", required: true)]
        public Input<int> RepositoryId { get; set; } = null!;

        public OrganizationRulesetRulesRequiredWorkflowsRequiredWorkflowGetArgs()
        {
        }
        public static new OrganizationRulesetRulesRequiredWorkflowsRequiredWorkflowGetArgs Empty => new OrganizationRulesetRulesRequiredWorkflowsRequiredWorkflowGetArgs();
    }
}
