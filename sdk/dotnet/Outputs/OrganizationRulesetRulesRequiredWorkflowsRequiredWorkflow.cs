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
    public sealed class OrganizationRulesetRulesRequiredWorkflowsRequiredWorkflow
    {
        /// <summary>
        /// (String) The path to the YAML definition file of the workflow.
        /// </summary>
        public readonly string Path;
        /// <summary>
        /// (String) The optional ref from which to fetch the workflow. Defaults to `master`.
        /// </summary>
        public readonly string? Ref;
        /// <summary>
        /// The repository IDs that the ruleset applies to. One of these IDs must match for the condition to pass. Conflicts with `repository_name`.
        /// </summary>
        public readonly int RepositoryId;

        [OutputConstructor]
        private OrganizationRulesetRulesRequiredWorkflowsRequiredWorkflow(
            string path,

            string? @ref,

            int repositoryId)
        {
            Path = path;
            Ref = @ref;
            RepositoryId = repositoryId;
        }
    }
}
