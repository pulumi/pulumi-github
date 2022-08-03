// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.github.inputs.BranchProtectionV3RequiredPullRequestReviewsArgs;
import com.pulumi.github.inputs.BranchProtectionV3RequiredStatusChecksArgs;
import com.pulumi.github.inputs.BranchProtectionV3RestrictionsArgs;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class BranchProtectionV3State extends com.pulumi.resources.ResourceArgs {

    public static final BranchProtectionV3State Empty = new BranchProtectionV3State();

    /**
     * The Git branch to protect.
     * 
     */
    @Import(name="branch")
    private @Nullable Output<String> branch;

    /**
     * @return The Git branch to protect.
     * 
     */
    public Optional<Output<String>> branch() {
        return Optional.ofNullable(this.branch);
    }

    /**
     * Boolean, setting this to `true` enforces status checks for repository administrators.
     * 
     */
    @Import(name="enforceAdmins")
    private @Nullable Output<Boolean> enforceAdmins;

    /**
     * @return Boolean, setting this to `true` enforces status checks for repository administrators.
     * 
     */
    public Optional<Output<Boolean>> enforceAdmins() {
        return Optional.ofNullable(this.enforceAdmins);
    }

    @Import(name="etag")
    private @Nullable Output<String> etag;

    public Optional<Output<String>> etag() {
        return Optional.ofNullable(this.etag);
    }

    /**
     * The GitHub repository name.
     * 
     */
    @Import(name="repository")
    private @Nullable Output<String> repository;

    /**
     * @return The GitHub repository name.
     * 
     */
    public Optional<Output<String>> repository() {
        return Optional.ofNullable(this.repository);
    }

    /**
     * Boolean, setting this to `true` requires all conversations on code must be resolved before a pull request can be merged.
     * 
     */
    @Import(name="requireConversationResolution")
    private @Nullable Output<Boolean> requireConversationResolution;

    /**
     * @return Boolean, setting this to `true` requires all conversations on code must be resolved before a pull request can be merged.
     * 
     */
    public Optional<Output<Boolean>> requireConversationResolution() {
        return Optional.ofNullable(this.requireConversationResolution);
    }

    /**
     * Boolean, setting this to `true` requires all commits to be signed with GPG.
     * 
     */
    @Import(name="requireSignedCommits")
    private @Nullable Output<Boolean> requireSignedCommits;

    /**
     * @return Boolean, setting this to `true` requires all commits to be signed with GPG.
     * 
     */
    public Optional<Output<Boolean>> requireSignedCommits() {
        return Optional.ofNullable(this.requireSignedCommits);
    }

    /**
     * Enforce restrictions for pull request reviews. See Required Pull Request Reviews below for details.
     * 
     */
    @Import(name="requiredPullRequestReviews")
    private @Nullable Output<BranchProtectionV3RequiredPullRequestReviewsArgs> requiredPullRequestReviews;

    /**
     * @return Enforce restrictions for pull request reviews. See Required Pull Request Reviews below for details.
     * 
     */
    public Optional<Output<BranchProtectionV3RequiredPullRequestReviewsArgs>> requiredPullRequestReviews() {
        return Optional.ofNullable(this.requiredPullRequestReviews);
    }

    /**
     * Enforce restrictions for required status checks. See Required Status Checks below for details.
     * 
     */
    @Import(name="requiredStatusChecks")
    private @Nullable Output<BranchProtectionV3RequiredStatusChecksArgs> requiredStatusChecks;

    /**
     * @return Enforce restrictions for required status checks. See Required Status Checks below for details.
     * 
     */
    public Optional<Output<BranchProtectionV3RequiredStatusChecksArgs>> requiredStatusChecks() {
        return Optional.ofNullable(this.requiredStatusChecks);
    }

    /**
     * Enforce restrictions for the users and teams that may push to the branch. See Restrictions below for details.
     * 
     */
    @Import(name="restrictions")
    private @Nullable Output<BranchProtectionV3RestrictionsArgs> restrictions;

    /**
     * @return Enforce restrictions for the users and teams that may push to the branch. See Restrictions below for details.
     * 
     */
    public Optional<Output<BranchProtectionV3RestrictionsArgs>> restrictions() {
        return Optional.ofNullable(this.restrictions);
    }

    private BranchProtectionV3State() {}

    private BranchProtectionV3State(BranchProtectionV3State $) {
        this.branch = $.branch;
        this.enforceAdmins = $.enforceAdmins;
        this.etag = $.etag;
        this.repository = $.repository;
        this.requireConversationResolution = $.requireConversationResolution;
        this.requireSignedCommits = $.requireSignedCommits;
        this.requiredPullRequestReviews = $.requiredPullRequestReviews;
        this.requiredStatusChecks = $.requiredStatusChecks;
        this.restrictions = $.restrictions;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(BranchProtectionV3State defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private BranchProtectionV3State $;

        public Builder() {
            $ = new BranchProtectionV3State();
        }

        public Builder(BranchProtectionV3State defaults) {
            $ = new BranchProtectionV3State(Objects.requireNonNull(defaults));
        }

        /**
         * @param branch The Git branch to protect.
         * 
         * @return builder
         * 
         */
        public Builder branch(@Nullable Output<String> branch) {
            $.branch = branch;
            return this;
        }

        /**
         * @param branch The Git branch to protect.
         * 
         * @return builder
         * 
         */
        public Builder branch(String branch) {
            return branch(Output.of(branch));
        }

        /**
         * @param enforceAdmins Boolean, setting this to `true` enforces status checks for repository administrators.
         * 
         * @return builder
         * 
         */
        public Builder enforceAdmins(@Nullable Output<Boolean> enforceAdmins) {
            $.enforceAdmins = enforceAdmins;
            return this;
        }

        /**
         * @param enforceAdmins Boolean, setting this to `true` enforces status checks for repository administrators.
         * 
         * @return builder
         * 
         */
        public Builder enforceAdmins(Boolean enforceAdmins) {
            return enforceAdmins(Output.of(enforceAdmins));
        }

        public Builder etag(@Nullable Output<String> etag) {
            $.etag = etag;
            return this;
        }

        public Builder etag(String etag) {
            return etag(Output.of(etag));
        }

        /**
         * @param repository The GitHub repository name.
         * 
         * @return builder
         * 
         */
        public Builder repository(@Nullable Output<String> repository) {
            $.repository = repository;
            return this;
        }

        /**
         * @param repository The GitHub repository name.
         * 
         * @return builder
         * 
         */
        public Builder repository(String repository) {
            return repository(Output.of(repository));
        }

        /**
         * @param requireConversationResolution Boolean, setting this to `true` requires all conversations on code must be resolved before a pull request can be merged.
         * 
         * @return builder
         * 
         */
        public Builder requireConversationResolution(@Nullable Output<Boolean> requireConversationResolution) {
            $.requireConversationResolution = requireConversationResolution;
            return this;
        }

        /**
         * @param requireConversationResolution Boolean, setting this to `true` requires all conversations on code must be resolved before a pull request can be merged.
         * 
         * @return builder
         * 
         */
        public Builder requireConversationResolution(Boolean requireConversationResolution) {
            return requireConversationResolution(Output.of(requireConversationResolution));
        }

        /**
         * @param requireSignedCommits Boolean, setting this to `true` requires all commits to be signed with GPG.
         * 
         * @return builder
         * 
         */
        public Builder requireSignedCommits(@Nullable Output<Boolean> requireSignedCommits) {
            $.requireSignedCommits = requireSignedCommits;
            return this;
        }

        /**
         * @param requireSignedCommits Boolean, setting this to `true` requires all commits to be signed with GPG.
         * 
         * @return builder
         * 
         */
        public Builder requireSignedCommits(Boolean requireSignedCommits) {
            return requireSignedCommits(Output.of(requireSignedCommits));
        }

        /**
         * @param requiredPullRequestReviews Enforce restrictions for pull request reviews. See Required Pull Request Reviews below for details.
         * 
         * @return builder
         * 
         */
        public Builder requiredPullRequestReviews(@Nullable Output<BranchProtectionV3RequiredPullRequestReviewsArgs> requiredPullRequestReviews) {
            $.requiredPullRequestReviews = requiredPullRequestReviews;
            return this;
        }

        /**
         * @param requiredPullRequestReviews Enforce restrictions for pull request reviews. See Required Pull Request Reviews below for details.
         * 
         * @return builder
         * 
         */
        public Builder requiredPullRequestReviews(BranchProtectionV3RequiredPullRequestReviewsArgs requiredPullRequestReviews) {
            return requiredPullRequestReviews(Output.of(requiredPullRequestReviews));
        }

        /**
         * @param requiredStatusChecks Enforce restrictions for required status checks. See Required Status Checks below for details.
         * 
         * @return builder
         * 
         */
        public Builder requiredStatusChecks(@Nullable Output<BranchProtectionV3RequiredStatusChecksArgs> requiredStatusChecks) {
            $.requiredStatusChecks = requiredStatusChecks;
            return this;
        }

        /**
         * @param requiredStatusChecks Enforce restrictions for required status checks. See Required Status Checks below for details.
         * 
         * @return builder
         * 
         */
        public Builder requiredStatusChecks(BranchProtectionV3RequiredStatusChecksArgs requiredStatusChecks) {
            return requiredStatusChecks(Output.of(requiredStatusChecks));
        }

        /**
         * @param restrictions Enforce restrictions for the users and teams that may push to the branch. See Restrictions below for details.
         * 
         * @return builder
         * 
         */
        public Builder restrictions(@Nullable Output<BranchProtectionV3RestrictionsArgs> restrictions) {
            $.restrictions = restrictions;
            return this;
        }

        /**
         * @param restrictions Enforce restrictions for the users and teams that may push to the branch. See Restrictions below for details.
         * 
         * @return builder
         * 
         */
        public Builder restrictions(BranchProtectionV3RestrictionsArgs restrictions) {
            return restrictions(Output.of(restrictions));
        }

        public BranchProtectionV3State build() {
            return $;
        }
    }

}
