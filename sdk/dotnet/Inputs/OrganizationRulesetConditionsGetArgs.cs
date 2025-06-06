// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github.Inputs
{

    public sealed class OrganizationRulesetConditionsGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// (Block List, Min: 1, Max: 1) (see below for nested schema)
        /// </summary>
        [Input("refName", required: true)]
        public Input<Inputs.OrganizationRulesetConditionsRefNameGetArgs> RefName { get; set; } = null!;

        [Input("repositoryIds")]
        private InputList<int>? _repositoryIds;

        /// <summary>
        /// The repository IDs that the ruleset applies to. One of these IDs must match for the condition to pass. Conflicts with `repository_name`.
        /// </summary>
        public InputList<int> RepositoryIds
        {
            get => _repositoryIds ?? (_repositoryIds = new InputList<int>());
            set => _repositoryIds = value;
        }

        /// <summary>
        /// Conflicts with `repository_id`. (see below for nested schema)
        /// 
        /// One of `repository_id` and `repository_name` must be set for the rule to target any repositories.
        /// </summary>
        [Input("repositoryName")]
        public Input<Inputs.OrganizationRulesetConditionsRepositoryNameGetArgs>? RepositoryName { get; set; }

        public OrganizationRulesetConditionsGetArgs()
        {
        }
        public static new OrganizationRulesetConditionsGetArgs Empty => new OrganizationRulesetConditionsGetArgs();
    }
}
