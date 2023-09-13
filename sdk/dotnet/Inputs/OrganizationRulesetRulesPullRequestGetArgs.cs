// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github.Inputs
{

    public sealed class OrganizationRulesetRulesPullRequestGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// (Boolean) New, reviewable commits pushed will dismiss previous pull request review approvals. Defaults to `false`.
        /// </summary>
        [Input("dismissStaleReviewsOnPush")]
        public Input<bool>? DismissStaleReviewsOnPush { get; set; }

        /// <summary>
        /// (Boolean) Require an approving review in pull requests that modify files that have a designated code owner. Defaults to `false`.
        /// </summary>
        [Input("requireCodeOwnerReview")]
        public Input<bool>? RequireCodeOwnerReview { get; set; }

        /// <summary>
        /// (Boolean) Whether the most recent reviewable push must be approved by someone other than the person who pushed it. Defaults to `false`.
        /// </summary>
        [Input("requireLastPushApproval")]
        public Input<bool>? RequireLastPushApproval { get; set; }

        /// <summary>
        /// (Number) The number of approving reviews that are required before a pull request can be merged. Defaults to `0`.
        /// </summary>
        [Input("requiredApprovingReviewCount")]
        public Input<int>? RequiredApprovingReviewCount { get; set; }

        /// <summary>
        /// (Boolean) All conversations on code must be resolved before a pull request can be merged. Defaults to `false`.
        /// </summary>
        [Input("requiredReviewThreadResolution")]
        public Input<bool>? RequiredReviewThreadResolution { get; set; }

        public OrganizationRulesetRulesPullRequestGetArgs()
        {
        }
        public static new OrganizationRulesetRulesPullRequestGetArgs Empty => new OrganizationRulesetRulesPullRequestGetArgs();
    }
}
