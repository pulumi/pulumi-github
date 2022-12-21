// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github.Outputs
{

    [OutputType]
    public sealed class BranchProtectionRequiredPullRequestReview
    {
        public readonly bool? DismissStaleReviews;
        public readonly ImmutableArray<string> DismissalRestrictions;
        public readonly ImmutableArray<string> PullRequestBypassers;
        public readonly bool? RequireCodeOwnerReviews;
        public readonly bool? RequireLastPushApproval;
        public readonly int? RequiredApprovingReviewCount;
        public readonly bool? RestrictDismissals;

        [OutputConstructor]
        private BranchProtectionRequiredPullRequestReview(
            bool? dismissStaleReviews,

            ImmutableArray<string> dismissalRestrictions,

            ImmutableArray<string> pullRequestBypassers,

            bool? requireCodeOwnerReviews,

            bool? requireLastPushApproval,

            int? requiredApprovingReviewCount,

            bool? restrictDismissals)
        {
            DismissStaleReviews = dismissStaleReviews;
            DismissalRestrictions = dismissalRestrictions;
            PullRequestBypassers = pullRequestBypassers;
            RequireCodeOwnerReviews = requireCodeOwnerReviews;
            RequireLastPushApproval = requireLastPushApproval;
            RequiredApprovingReviewCount = requiredApprovingReviewCount;
            RestrictDismissals = restrictDismissals;
        }
    }
}
