// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.github.inputs.TeamSettingsReviewRequestDelegationArgs;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class TeamSettingsArgs extends com.pulumi.resources.ResourceArgs {

    public static final TeamSettingsArgs Empty = new TeamSettingsArgs();

    /**
     * The settings for delegating code reviews to individuals on behalf of the team. If this block is present, even without any fields, then review request delegation will be enabled for the team. See GitHub Review Request Delegation below for details. See [GitHub&#39;s documentation](https://docs.github.com/en/organizations/organizing-members-into-teams/managing-code-review-settings-for-your-team#configuring-team-notifications) for more configuration details.
     * 
     */
    @Import(name="reviewRequestDelegation")
    private @Nullable Output<TeamSettingsReviewRequestDelegationArgs> reviewRequestDelegation;

    /**
     * @return The settings for delegating code reviews to individuals on behalf of the team. If this block is present, even without any fields, then review request delegation will be enabled for the team. See GitHub Review Request Delegation below for details. See [GitHub&#39;s documentation](https://docs.github.com/en/organizations/organizing-members-into-teams/managing-code-review-settings-for-your-team#configuring-team-notifications) for more configuration details.
     * 
     */
    public Optional<Output<TeamSettingsReviewRequestDelegationArgs>> reviewRequestDelegation() {
        return Optional.ofNullable(this.reviewRequestDelegation);
    }

    /**
     * The GitHub team id or the GitHub team slug
     * 
     */
    @Import(name="teamId", required=true)
    private Output<String> teamId;

    /**
     * @return The GitHub team id or the GitHub team slug
     * 
     */
    public Output<String> teamId() {
        return this.teamId;
    }

    private TeamSettingsArgs() {}

    private TeamSettingsArgs(TeamSettingsArgs $) {
        this.reviewRequestDelegation = $.reviewRequestDelegation;
        this.teamId = $.teamId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(TeamSettingsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private TeamSettingsArgs $;

        public Builder() {
            $ = new TeamSettingsArgs();
        }

        public Builder(TeamSettingsArgs defaults) {
            $ = new TeamSettingsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param reviewRequestDelegation The settings for delegating code reviews to individuals on behalf of the team. If this block is present, even without any fields, then review request delegation will be enabled for the team. See GitHub Review Request Delegation below for details. See [GitHub&#39;s documentation](https://docs.github.com/en/organizations/organizing-members-into-teams/managing-code-review-settings-for-your-team#configuring-team-notifications) for more configuration details.
         * 
         * @return builder
         * 
         */
        public Builder reviewRequestDelegation(@Nullable Output<TeamSettingsReviewRequestDelegationArgs> reviewRequestDelegation) {
            $.reviewRequestDelegation = reviewRequestDelegation;
            return this;
        }

        /**
         * @param reviewRequestDelegation The settings for delegating code reviews to individuals on behalf of the team. If this block is present, even without any fields, then review request delegation will be enabled for the team. See GitHub Review Request Delegation below for details. See [GitHub&#39;s documentation](https://docs.github.com/en/organizations/organizing-members-into-teams/managing-code-review-settings-for-your-team#configuring-team-notifications) for more configuration details.
         * 
         * @return builder
         * 
         */
        public Builder reviewRequestDelegation(TeamSettingsReviewRequestDelegationArgs reviewRequestDelegation) {
            return reviewRequestDelegation(Output.of(reviewRequestDelegation));
        }

        /**
         * @param teamId The GitHub team id or the GitHub team slug
         * 
         * @return builder
         * 
         */
        public Builder teamId(Output<String> teamId) {
            $.teamId = teamId;
            return this;
        }

        /**
         * @param teamId The GitHub team id or the GitHub team slug
         * 
         * @return builder
         * 
         */
        public Builder teamId(String teamId) {
            return teamId(Output.of(teamId));
        }

        public TeamSettingsArgs build() {
            $.teamId = Objects.requireNonNull($.teamId, "expected parameter 'teamId' to be non-null");
            return $;
        }
    }

}
