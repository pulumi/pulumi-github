// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github.Inputs
{

    public sealed class RepositoryRulesetRulesRequiredDeploymentsArgs : global::Pulumi.ResourceArgs
    {
        [Input("requiredDeploymentEnvironments", required: true)]
        private InputList<string>? _requiredDeploymentEnvironments;

        /// <summary>
        /// The environments that must be successfully deployed to before branches can be merged.
        /// </summary>
        public InputList<string> RequiredDeploymentEnvironments
        {
            get => _requiredDeploymentEnvironments ?? (_requiredDeploymentEnvironments = new InputList<string>());
            set => _requiredDeploymentEnvironments = value;
        }

        public RepositoryRulesetRulesRequiredDeploymentsArgs()
        {
        }
        public static new RepositoryRulesetRulesRequiredDeploymentsArgs Empty => new RepositoryRulesetRulesRequiredDeploymentsArgs();
    }
}
