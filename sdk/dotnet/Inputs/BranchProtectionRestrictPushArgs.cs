// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github.Inputs
{

    public sealed class BranchProtectionRestrictPushArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Boolean, setting this to `false` allows people, teams, or apps to create new branches matching this rule. Defaults to `true`.
        /// </summary>
        [Input("blocksCreations")]
        public Input<bool>? BlocksCreations { get; set; }

        [Input("pushAllowances")]
        private InputList<string>? _pushAllowances;

        /// <summary>
        /// A list of actor Names/IDs that may push to the branch. Actor names must either begin with a "/" for users or the organization name followed by a "/" for teams. Organization administrators, repository administrators, and users with the Maintain role on the repository can always push when all other requirements have passed.
        /// </summary>
        public InputList<string> PushAllowances
        {
            get => _pushAllowances ?? (_pushAllowances = new InputList<string>());
            set => _pushAllowances = value;
        }

        public BranchProtectionRestrictPushArgs()
        {
        }
        public static new BranchProtectionRestrictPushArgs Empty => new BranchProtectionRestrictPushArgs();
    }
}
