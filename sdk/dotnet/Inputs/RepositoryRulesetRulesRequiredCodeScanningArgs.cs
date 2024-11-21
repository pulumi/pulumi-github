// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github.Inputs
{

    public sealed class RepositoryRulesetRulesRequiredCodeScanningArgs : global::Pulumi.ResourceArgs
    {
        [Input("requiredCodeScanningTools", required: true)]
        private InputList<Inputs.RepositoryRulesetRulesRequiredCodeScanningRequiredCodeScanningToolArgs>? _requiredCodeScanningTools;

        /// <summary>
        /// Tools that must provide code scanning results for this rule to pass.
        /// </summary>
        public InputList<Inputs.RepositoryRulesetRulesRequiredCodeScanningRequiredCodeScanningToolArgs> RequiredCodeScanningTools
        {
            get => _requiredCodeScanningTools ?? (_requiredCodeScanningTools = new InputList<Inputs.RepositoryRulesetRulesRequiredCodeScanningRequiredCodeScanningToolArgs>());
            set => _requiredCodeScanningTools = value;
        }

        public RepositoryRulesetRulesRequiredCodeScanningArgs()
        {
        }
        public static new RepositoryRulesetRulesRequiredCodeScanningArgs Empty => new RepositoryRulesetRulesRequiredCodeScanningArgs();
    }
}