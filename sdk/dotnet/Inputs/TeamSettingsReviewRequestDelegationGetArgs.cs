// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github.Inputs
{

    public sealed class TeamSettingsReviewRequestDelegationGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The algorithm to use when assigning pull requests to team members. Supported values are 'ROUND_ROBIN' and 'LOAD_BALANCE'.
        /// </summary>
        [Input("algorithm")]
        public Input<string>? Algorithm { get; set; }

        /// <summary>
        /// The number of team members to assign to a pull request.
        /// </summary>
        [Input("memberCount")]
        public Input<int>? MemberCount { get; set; }

        /// <summary>
        /// whether to notify the entire team when at least one member is also assigned to the pull request.
        /// </summary>
        [Input("notify")]
        public Input<bool>? Notify { get; set; }

        public TeamSettingsReviewRequestDelegationGetArgs()
        {
        }
        public static new TeamSettingsReviewRequestDelegationGetArgs Empty => new TeamSettingsReviewRequestDelegationGetArgs();
    }
}
