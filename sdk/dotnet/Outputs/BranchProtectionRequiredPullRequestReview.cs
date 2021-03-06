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
        public readonly bool? RequireCodeOwnerReviews;
        public readonly int? RequiredApprovingReviewCount;

        [OutputConstructor]
        private BranchProtectionRequiredPullRequestReview(
            bool? dismissStaleReviews,

            ImmutableArray<string> dismissalRestrictions,

            bool? requireCodeOwnerReviews,

            int? requiredApprovingReviewCount)
        {
            DismissStaleReviews = dismissStaleReviews;
            DismissalRestrictions = dismissalRestrictions;
            RequireCodeOwnerReviews = requireCodeOwnerReviews;
            RequiredApprovingReviewCount = requiredApprovingReviewCount;
        }
    }
}
