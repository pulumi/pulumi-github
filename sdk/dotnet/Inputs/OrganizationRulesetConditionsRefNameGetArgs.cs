// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github.Inputs
{

    public sealed class OrganizationRulesetConditionsRefNameGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("excludes", required: true)]
        private InputList<string>? _excludes;

        /// <summary>
        /// (List of String) Array of repository names or patterns to exclude. The condition will not pass if any of these patterns match.
        /// </summary>
        public InputList<string> Excludes
        {
            get => _excludes ?? (_excludes = new InputList<string>());
            set => _excludes = value;
        }

        [Input("includes", required: true)]
        private InputList<string>? _includes;

        /// <summary>
        /// (List of String) Array of repository names or patterns to include. One of these patterns must match for the condition to pass. Also accepts `~ALL` to include all repositories.
        /// </summary>
        public InputList<string> Includes
        {
            get => _includes ?? (_includes = new InputList<string>());
            set => _includes = value;
        }

        public OrganizationRulesetConditionsRefNameGetArgs()
        {
        }
        public static new OrganizationRulesetConditionsRefNameGetArgs Empty => new OrganizationRulesetConditionsRefNameGetArgs();
    }
}
