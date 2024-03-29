// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class BranchProtectionV3RequiredPullRequestReviewsBypassPullRequestAllowancesArgs extends com.pulumi.resources.ResourceArgs {

    public static final BranchProtectionV3RequiredPullRequestReviewsBypassPullRequestAllowancesArgs Empty = new BranchProtectionV3RequiredPullRequestReviewsBypassPullRequestAllowancesArgs();

    /**
     * The list of app slugs allowed to bypass pull request requirements.
     * 
     */
    @Import(name="apps")
    private @Nullable Output<List<String>> apps;

    /**
     * @return The list of app slugs allowed to bypass pull request requirements.
     * 
     */
    public Optional<Output<List<String>>> apps() {
        return Optional.ofNullable(this.apps);
    }

    /**
     * The list of team slugs allowed to bypass pull request requirements.
     * 
     */
    @Import(name="teams")
    private @Nullable Output<List<String>> teams;

    /**
     * @return The list of team slugs allowed to bypass pull request requirements.
     * 
     */
    public Optional<Output<List<String>>> teams() {
        return Optional.ofNullable(this.teams);
    }

    /**
     * The list of user logins allowed to bypass pull request requirements.
     * 
     */
    @Import(name="users")
    private @Nullable Output<List<String>> users;

    /**
     * @return The list of user logins allowed to bypass pull request requirements.
     * 
     */
    public Optional<Output<List<String>>> users() {
        return Optional.ofNullable(this.users);
    }

    private BranchProtectionV3RequiredPullRequestReviewsBypassPullRequestAllowancesArgs() {}

    private BranchProtectionV3RequiredPullRequestReviewsBypassPullRequestAllowancesArgs(BranchProtectionV3RequiredPullRequestReviewsBypassPullRequestAllowancesArgs $) {
        this.apps = $.apps;
        this.teams = $.teams;
        this.users = $.users;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(BranchProtectionV3RequiredPullRequestReviewsBypassPullRequestAllowancesArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private BranchProtectionV3RequiredPullRequestReviewsBypassPullRequestAllowancesArgs $;

        public Builder() {
            $ = new BranchProtectionV3RequiredPullRequestReviewsBypassPullRequestAllowancesArgs();
        }

        public Builder(BranchProtectionV3RequiredPullRequestReviewsBypassPullRequestAllowancesArgs defaults) {
            $ = new BranchProtectionV3RequiredPullRequestReviewsBypassPullRequestAllowancesArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param apps The list of app slugs allowed to bypass pull request requirements.
         * 
         * @return builder
         * 
         */
        public Builder apps(@Nullable Output<List<String>> apps) {
            $.apps = apps;
            return this;
        }

        /**
         * @param apps The list of app slugs allowed to bypass pull request requirements.
         * 
         * @return builder
         * 
         */
        public Builder apps(List<String> apps) {
            return apps(Output.of(apps));
        }

        /**
         * @param apps The list of app slugs allowed to bypass pull request requirements.
         * 
         * @return builder
         * 
         */
        public Builder apps(String... apps) {
            return apps(List.of(apps));
        }

        /**
         * @param teams The list of team slugs allowed to bypass pull request requirements.
         * 
         * @return builder
         * 
         */
        public Builder teams(@Nullable Output<List<String>> teams) {
            $.teams = teams;
            return this;
        }

        /**
         * @param teams The list of team slugs allowed to bypass pull request requirements.
         * 
         * @return builder
         * 
         */
        public Builder teams(List<String> teams) {
            return teams(Output.of(teams));
        }

        /**
         * @param teams The list of team slugs allowed to bypass pull request requirements.
         * 
         * @return builder
         * 
         */
        public Builder teams(String... teams) {
            return teams(List.of(teams));
        }

        /**
         * @param users The list of user logins allowed to bypass pull request requirements.
         * 
         * @return builder
         * 
         */
        public Builder users(@Nullable Output<List<String>> users) {
            $.users = users;
            return this;
        }

        /**
         * @param users The list of user logins allowed to bypass pull request requirements.
         * 
         * @return builder
         * 
         */
        public Builder users(List<String> users) {
            return users(Output.of(users));
        }

        /**
         * @param users The list of user logins allowed to bypass pull request requirements.
         * 
         * @return builder
         * 
         */
        public Builder users(String... users) {
            return users(List.of(users));
        }

        public BranchProtectionV3RequiredPullRequestReviewsBypassPullRequestAllowancesArgs build() {
            return $;
        }
    }

}
