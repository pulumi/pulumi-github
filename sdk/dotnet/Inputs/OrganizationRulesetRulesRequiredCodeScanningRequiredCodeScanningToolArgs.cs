// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github.Inputs
{

    public sealed class OrganizationRulesetRulesRequiredCodeScanningRequiredCodeScanningToolArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The severity level at which code scanning results that raise alerts block a reference update. Can be one of: `none`, `errors`, `errors_and_warnings`, `all`.
        /// </summary>
        [Input("alertsThreshold", required: true)]
        public Input<string> AlertsThreshold { get; set; } = null!;

        /// <summary>
        /// The severity level at which code scanning results that raise security alerts block a reference update. Can be one of: `none`, `critical`, `high_or_higher`, `medium_or_higher`, `all`.
        /// </summary>
        [Input("securityAlertsThreshold", required: true)]
        public Input<string> SecurityAlertsThreshold { get; set; } = null!;

        /// <summary>
        /// The name of a code scanning tool.
        /// </summary>
        [Input("tool", required: true)]
        public Input<string> Tool { get; set; } = null!;

        public OrganizationRulesetRulesRequiredCodeScanningRequiredCodeScanningToolArgs()
        {
        }
        public static new OrganizationRulesetRulesRequiredCodeScanningRequiredCodeScanningToolArgs Empty => new OrganizationRulesetRulesRequiredCodeScanningRequiredCodeScanningToolArgs();
    }
}
