// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github.Inputs
{

    public sealed class BranchProtectionRequiredPullRequestReviewArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Dismiss approved reviews automatically when a new commit is pushed. Defaults to `false`.
        /// </summary>
        [Input("dismissStaleReviews")]
        public Input<bool>? DismissStaleReviews { get; set; }

        [Input("dismissalRestrictions")]
        private InputList<string>? _dismissalRestrictions;

        /// <summary>
        /// The list of actor Names/IDs with dismissal access. If not empty, `restrict_dismissals` is ignored. Actor names must either begin with a "/" for users or the organization name followed by a "/" for teams.
        /// </summary>
        public InputList<string> DismissalRestrictions
        {
            get => _dismissalRestrictions ?? (_dismissalRestrictions = new InputList<string>());
            set => _dismissalRestrictions = value;
        }

        [Input("pullRequestBypassers")]
        private InputList<string>? _pullRequestBypassers;

        /// <summary>
        /// The list of actor Names/IDs that are allowed to bypass pull request requirements. Actor names must either begin with a "/" for users or the organization name followed by a "/" for teams.
        /// </summary>
        public InputList<string> PullRequestBypassers
        {
            get => _pullRequestBypassers ?? (_pullRequestBypassers = new InputList<string>());
            set => _pullRequestBypassers = value;
        }

        /// <summary>
        /// Require an approved review in pull requests including files with a designated code owner. Defaults to `false`.
        /// </summary>
        [Input("requireCodeOwnerReviews")]
        public Input<bool>? RequireCodeOwnerReviews { get; set; }

        /// <summary>
        /// Require that The most recent push must be approved by someone other than the last pusher.  Defaults to `false`
        /// </summary>
        [Input("requireLastPushApproval")]
        public Input<bool>? RequireLastPushApproval { get; set; }

        /// <summary>
        /// Require x number of approvals to satisfy branch protection requirements. If this is specified it must be a number between 0-6. This requirement matches GitHub's API, see the upstream [documentation](https://developer.github.com/v3/repos/branches/#parameters-1) for more information.
        /// (https://developer.github.com/v3/repos/branches/#parameters-1) for more information.
        /// </summary>
        [Input("requiredApprovingReviewCount")]
        public Input<int>? RequiredApprovingReviewCount { get; set; }

        /// <summary>
        /// Restrict pull request review dismissals.
        /// </summary>
        [Input("restrictDismissals")]
        public Input<bool>? RestrictDismissals { get; set; }

        public BranchProtectionRequiredPullRequestReviewArgs()
        {
        }
        public static new BranchProtectionRequiredPullRequestReviewArgs Empty => new BranchProtectionRequiredPullRequestReviewArgs();
    }
}
