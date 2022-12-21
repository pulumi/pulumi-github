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
        [Input("dismissStaleReviews")]
        public Input<bool>? DismissStaleReviews { get; set; }

        [Input("dismissalRestrictions")]
        private InputList<string>? _dismissalRestrictions;
        public InputList<string> DismissalRestrictions
        {
            get => _dismissalRestrictions ?? (_dismissalRestrictions = new InputList<string>());
            set => _dismissalRestrictions = value;
        }

        [Input("pullRequestBypassers")]
        private InputList<string>? _pullRequestBypassers;
        public InputList<string> PullRequestBypassers
        {
            get => _pullRequestBypassers ?? (_pullRequestBypassers = new InputList<string>());
            set => _pullRequestBypassers = value;
        }

        [Input("requireCodeOwnerReviews")]
        public Input<bool>? RequireCodeOwnerReviews { get; set; }

        [Input("requireLastPushApproval")]
        public Input<bool>? RequireLastPushApproval { get; set; }

        [Input("requiredApprovingReviewCount")]
        public Input<int>? RequiredApprovingReviewCount { get; set; }

        [Input("restrictDismissals")]
        public Input<bool>? RestrictDismissals { get; set; }

        public BranchProtectionRequiredPullRequestReviewArgs()
        {
        }
        public static new BranchProtectionRequiredPullRequestReviewArgs Empty => new BranchProtectionRequiredPullRequestReviewArgs();
    }
}
