// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github.Inputs
{

    public sealed class OrganizationRulesetRulesCommitterEmailPatternArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// (String) The name of the ruleset.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// If true, the rule will fail if the pattern matches.
        /// </summary>
        [Input("negate")]
        public Input<bool>? Negate { get; set; }

        /// <summary>
        /// The operator to use for matching. Can be one of: `starts_with`, `ends_with`, `contains`, `regex`.
        /// </summary>
        [Input("operator", required: true)]
        public Input<string> Operator { get; set; } = null!;

        /// <summary>
        /// The pattern to match with.
        /// </summary>
        [Input("pattern", required: true)]
        public Input<string> Pattern { get; set; } = null!;

        public OrganizationRulesetRulesCommitterEmailPatternArgs()
        {
        }
        public static new OrganizationRulesetRulesCommitterEmailPatternArgs Empty => new OrganizationRulesetRulesCommitterEmailPatternArgs();
    }
}