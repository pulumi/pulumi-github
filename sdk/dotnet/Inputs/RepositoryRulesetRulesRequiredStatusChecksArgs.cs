// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github.Inputs
{

    public sealed class RepositoryRulesetRulesRequiredStatusChecksArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Allow repositories and branches to be created if a check would otherwise prohibit it.
        /// </summary>
        [Input("doNotEnforceOnCreate")]
        public Input<bool>? DoNotEnforceOnCreate { get; set; }

        [Input("requiredChecks", required: true)]
        private InputList<Inputs.RepositoryRulesetRulesRequiredStatusChecksRequiredCheckArgs>? _requiredChecks;

        /// <summary>
        /// Status checks that are required. Several can be defined.
        /// </summary>
        public InputList<Inputs.RepositoryRulesetRulesRequiredStatusChecksRequiredCheckArgs> RequiredChecks
        {
            get => _requiredChecks ?? (_requiredChecks = new InputList<Inputs.RepositoryRulesetRulesRequiredStatusChecksRequiredCheckArgs>());
            set => _requiredChecks = value;
        }

        /// <summary>
        /// Whether pull requests targeting a matching branch must be tested with the latest code. This setting will not take effect unless at least one status check is enabled. Defaults to `false`.
        /// </summary>
        [Input("strictRequiredStatusChecksPolicy")]
        public Input<bool>? StrictRequiredStatusChecksPolicy { get; set; }

        public RepositoryRulesetRulesRequiredStatusChecksArgs()
        {
        }
        public static new RepositoryRulesetRulesRequiredStatusChecksArgs Empty => new RepositoryRulesetRulesRequiredStatusChecksArgs();
    }
}
