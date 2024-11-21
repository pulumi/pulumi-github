// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.github.inputs.OrganizationRulesetRulesBranchNamePatternArgs;
import com.pulumi.github.inputs.OrganizationRulesetRulesCommitAuthorEmailPatternArgs;
import com.pulumi.github.inputs.OrganizationRulesetRulesCommitMessagePatternArgs;
import com.pulumi.github.inputs.OrganizationRulesetRulesCommitterEmailPatternArgs;
import com.pulumi.github.inputs.OrganizationRulesetRulesPullRequestArgs;
import com.pulumi.github.inputs.OrganizationRulesetRulesRequiredCodeScanningArgs;
import com.pulumi.github.inputs.OrganizationRulesetRulesRequiredStatusChecksArgs;
import com.pulumi.github.inputs.OrganizationRulesetRulesRequiredWorkflowsArgs;
import com.pulumi.github.inputs.OrganizationRulesetRulesTagNamePatternArgs;
import java.lang.Boolean;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class OrganizationRulesetRulesArgs extends com.pulumi.resources.ResourceArgs {

    public static final OrganizationRulesetRulesArgs Empty = new OrganizationRulesetRulesArgs();

    /**
     * (Block List, Max: 1) Parameters to be used for the branch_name_pattern rule. This rule only applies to repositories within an enterprise, it cannot be applied to repositories owned by individuals or regular organizations. Conflicts with `tag_name_pattern` as it only applies to rulesets with target `branch`. (see below for nested schema)
     * 
     */
    @Import(name="branchNamePattern")
    private @Nullable Output<OrganizationRulesetRulesBranchNamePatternArgs> branchNamePattern;

    /**
     * @return (Block List, Max: 1) Parameters to be used for the branch_name_pattern rule. This rule only applies to repositories within an enterprise, it cannot be applied to repositories owned by individuals or regular organizations. Conflicts with `tag_name_pattern` as it only applies to rulesets with target `branch`. (see below for nested schema)
     * 
     */
    public Optional<Output<OrganizationRulesetRulesBranchNamePatternArgs>> branchNamePattern() {
        return Optional.ofNullable(this.branchNamePattern);
    }

    /**
     * (Block List, Max: 1) Parameters to be used for the commit_author_email_pattern rule. This rule only applies to repositories within an enterprise, it cannot be applied to repositories owned by individuals or regular organizations. (see below for nested schema)
     * 
     */
    @Import(name="commitAuthorEmailPattern")
    private @Nullable Output<OrganizationRulesetRulesCommitAuthorEmailPatternArgs> commitAuthorEmailPattern;

    /**
     * @return (Block List, Max: 1) Parameters to be used for the commit_author_email_pattern rule. This rule only applies to repositories within an enterprise, it cannot be applied to repositories owned by individuals or regular organizations. (see below for nested schema)
     * 
     */
    public Optional<Output<OrganizationRulesetRulesCommitAuthorEmailPatternArgs>> commitAuthorEmailPattern() {
        return Optional.ofNullable(this.commitAuthorEmailPattern);
    }

    /**
     * (Block List, Max: 1) Parameters to be used for the commit_message_pattern rule. This rule only applies to repositories within an enterprise, it cannot be applied to repositories owned by individuals or regular organizations. (see below for nested schema)
     * 
     */
    @Import(name="commitMessagePattern")
    private @Nullable Output<OrganizationRulesetRulesCommitMessagePatternArgs> commitMessagePattern;

    /**
     * @return (Block List, Max: 1) Parameters to be used for the commit_message_pattern rule. This rule only applies to repositories within an enterprise, it cannot be applied to repositories owned by individuals or regular organizations. (see below for nested schema)
     * 
     */
    public Optional<Output<OrganizationRulesetRulesCommitMessagePatternArgs>> commitMessagePattern() {
        return Optional.ofNullable(this.commitMessagePattern);
    }

    /**
     * (Block List, Max: 1) Parameters to be used for the committer_email_pattern rule. This rule only applies to repositories within an enterprise, it cannot be applied to repositories owned by individuals or regular organizations. (see below for nested schema)
     * 
     */
    @Import(name="committerEmailPattern")
    private @Nullable Output<OrganizationRulesetRulesCommitterEmailPatternArgs> committerEmailPattern;

    /**
     * @return (Block List, Max: 1) Parameters to be used for the committer_email_pattern rule. This rule only applies to repositories within an enterprise, it cannot be applied to repositories owned by individuals or regular organizations. (see below for nested schema)
     * 
     */
    public Optional<Output<OrganizationRulesetRulesCommitterEmailPatternArgs>> committerEmailPattern() {
        return Optional.ofNullable(this.committerEmailPattern);
    }

    /**
     * (Boolean) Only allow users with bypass permission to create matching refs.
     * 
     */
    @Import(name="creation")
    private @Nullable Output<Boolean> creation;

    /**
     * @return (Boolean) Only allow users with bypass permission to create matching refs.
     * 
     */
    public Optional<Output<Boolean>> creation() {
        return Optional.ofNullable(this.creation);
    }

    /**
     * (Boolean) Only allow users with bypass permissions to delete matching refs.
     * 
     */
    @Import(name="deletion")
    private @Nullable Output<Boolean> deletion;

    /**
     * @return (Boolean) Only allow users with bypass permissions to delete matching refs.
     * 
     */
    public Optional<Output<Boolean>> deletion() {
        return Optional.ofNullable(this.deletion);
    }

    /**
     * (Boolean) Prevent users with push access from force pushing to branches.
     * 
     */
    @Import(name="nonFastForward")
    private @Nullable Output<Boolean> nonFastForward;

    /**
     * @return (Boolean) Prevent users with push access from force pushing to branches.
     * 
     */
    public Optional<Output<Boolean>> nonFastForward() {
        return Optional.ofNullable(this.nonFastForward);
    }

    /**
     * (Block List, Max: 1) Require all commits be made to a non-target branch and submitted via a pull request before they can be merged. (see below for nested schema)
     * 
     */
    @Import(name="pullRequest")
    private @Nullable Output<OrganizationRulesetRulesPullRequestArgs> pullRequest;

    /**
     * @return (Block List, Max: 1) Require all commits be made to a non-target branch and submitted via a pull request before they can be merged. (see below for nested schema)
     * 
     */
    public Optional<Output<OrganizationRulesetRulesPullRequestArgs>> pullRequest() {
        return Optional.ofNullable(this.pullRequest);
    }

    /**
     * (Block List, Max: 1) Define which tools must provide code scanning results before the reference is updated. When configured, code scanning must be enabled and have results for both the commit and the reference being updated. Multiple code scanning tools can be specified. (see below for nested schema)
     * 
     */
    @Import(name="requiredCodeScanning")
    private @Nullable Output<OrganizationRulesetRulesRequiredCodeScanningArgs> requiredCodeScanning;

    /**
     * @return (Block List, Max: 1) Define which tools must provide code scanning results before the reference is updated. When configured, code scanning must be enabled and have results for both the commit and the reference being updated. Multiple code scanning tools can be specified. (see below for nested schema)
     * 
     */
    public Optional<Output<OrganizationRulesetRulesRequiredCodeScanningArgs>> requiredCodeScanning() {
        return Optional.ofNullable(this.requiredCodeScanning);
    }

    /**
     * (Boolean) Prevent merge commits from being pushed to matching branches.
     * 
     */
    @Import(name="requiredLinearHistory")
    private @Nullable Output<Boolean> requiredLinearHistory;

    /**
     * @return (Boolean) Prevent merge commits from being pushed to matching branches.
     * 
     */
    public Optional<Output<Boolean>> requiredLinearHistory() {
        return Optional.ofNullable(this.requiredLinearHistory);
    }

    /**
     * (Boolean) Commits pushed to matching branches must have verified signatures.
     * 
     */
    @Import(name="requiredSignatures")
    private @Nullable Output<Boolean> requiredSignatures;

    /**
     * @return (Boolean) Commits pushed to matching branches must have verified signatures.
     * 
     */
    public Optional<Output<Boolean>> requiredSignatures() {
        return Optional.ofNullable(this.requiredSignatures);
    }

    /**
     * (Block List, Max: 1) Choose which status checks must pass before branches can be merged into a branch that matches this rule. When enabled, commits must first be pushed to another branch, then merged or pushed directly to a branch that matches this rule after status checks have passed. (see below for nested schema)
     * 
     */
    @Import(name="requiredStatusChecks")
    private @Nullable Output<OrganizationRulesetRulesRequiredStatusChecksArgs> requiredStatusChecks;

    /**
     * @return (Block List, Max: 1) Choose which status checks must pass before branches can be merged into a branch that matches this rule. When enabled, commits must first be pushed to another branch, then merged or pushed directly to a branch that matches this rule after status checks have passed. (see below for nested schema)
     * 
     */
    public Optional<Output<OrganizationRulesetRulesRequiredStatusChecksArgs>> requiredStatusChecks() {
        return Optional.ofNullable(this.requiredStatusChecks);
    }

    /**
     * (Block List, Max: 1) Define which Actions workflows must pass before changes can be merged into a branch matching the rule. Multiple workflows can be specified. (see below for nested schema)
     * 
     */
    @Import(name="requiredWorkflows")
    private @Nullable Output<OrganizationRulesetRulesRequiredWorkflowsArgs> requiredWorkflows;

    /**
     * @return (Block List, Max: 1) Define which Actions workflows must pass before changes can be merged into a branch matching the rule. Multiple workflows can be specified. (see below for nested schema)
     * 
     */
    public Optional<Output<OrganizationRulesetRulesRequiredWorkflowsArgs>> requiredWorkflows() {
        return Optional.ofNullable(this.requiredWorkflows);
    }

    /**
     * (Block List, Max: 1) Parameters to be used for the tag_name_pattern rule. This rule only applies to repositories within an enterprise, it cannot be applied to repositories owned by individuals or regular organizations. Conflicts with `branch_name_pattern` as it only applies to rulesets with target `tag`. (see below for nested schema)
     * 
     */
    @Import(name="tagNamePattern")
    private @Nullable Output<OrganizationRulesetRulesTagNamePatternArgs> tagNamePattern;

    /**
     * @return (Block List, Max: 1) Parameters to be used for the tag_name_pattern rule. This rule only applies to repositories within an enterprise, it cannot be applied to repositories owned by individuals or regular organizations. Conflicts with `branch_name_pattern` as it only applies to rulesets with target `tag`. (see below for nested schema)
     * 
     */
    public Optional<Output<OrganizationRulesetRulesTagNamePatternArgs>> tagNamePattern() {
        return Optional.ofNullable(this.tagNamePattern);
    }

    /**
     * (Boolean) Only allow users with bypass permission to update matching refs.
     * 
     */
    @Import(name="update")
    private @Nullable Output<Boolean> update;

    /**
     * @return (Boolean) Only allow users with bypass permission to update matching refs.
     * 
     */
    public Optional<Output<Boolean>> update() {
        return Optional.ofNullable(this.update);
    }

    private OrganizationRulesetRulesArgs() {}

    private OrganizationRulesetRulesArgs(OrganizationRulesetRulesArgs $) {
        this.branchNamePattern = $.branchNamePattern;
        this.commitAuthorEmailPattern = $.commitAuthorEmailPattern;
        this.commitMessagePattern = $.commitMessagePattern;
        this.committerEmailPattern = $.committerEmailPattern;
        this.creation = $.creation;
        this.deletion = $.deletion;
        this.nonFastForward = $.nonFastForward;
        this.pullRequest = $.pullRequest;
        this.requiredCodeScanning = $.requiredCodeScanning;
        this.requiredLinearHistory = $.requiredLinearHistory;
        this.requiredSignatures = $.requiredSignatures;
        this.requiredStatusChecks = $.requiredStatusChecks;
        this.requiredWorkflows = $.requiredWorkflows;
        this.tagNamePattern = $.tagNamePattern;
        this.update = $.update;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(OrganizationRulesetRulesArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private OrganizationRulesetRulesArgs $;

        public Builder() {
            $ = new OrganizationRulesetRulesArgs();
        }

        public Builder(OrganizationRulesetRulesArgs defaults) {
            $ = new OrganizationRulesetRulesArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param branchNamePattern (Block List, Max: 1) Parameters to be used for the branch_name_pattern rule. This rule only applies to repositories within an enterprise, it cannot be applied to repositories owned by individuals or regular organizations. Conflicts with `tag_name_pattern` as it only applies to rulesets with target `branch`. (see below for nested schema)
         * 
         * @return builder
         * 
         */
        public Builder branchNamePattern(@Nullable Output<OrganizationRulesetRulesBranchNamePatternArgs> branchNamePattern) {
            $.branchNamePattern = branchNamePattern;
            return this;
        }

        /**
         * @param branchNamePattern (Block List, Max: 1) Parameters to be used for the branch_name_pattern rule. This rule only applies to repositories within an enterprise, it cannot be applied to repositories owned by individuals or regular organizations. Conflicts with `tag_name_pattern` as it only applies to rulesets with target `branch`. (see below for nested schema)
         * 
         * @return builder
         * 
         */
        public Builder branchNamePattern(OrganizationRulesetRulesBranchNamePatternArgs branchNamePattern) {
            return branchNamePattern(Output.of(branchNamePattern));
        }

        /**
         * @param commitAuthorEmailPattern (Block List, Max: 1) Parameters to be used for the commit_author_email_pattern rule. This rule only applies to repositories within an enterprise, it cannot be applied to repositories owned by individuals or regular organizations. (see below for nested schema)
         * 
         * @return builder
         * 
         */
        public Builder commitAuthorEmailPattern(@Nullable Output<OrganizationRulesetRulesCommitAuthorEmailPatternArgs> commitAuthorEmailPattern) {
            $.commitAuthorEmailPattern = commitAuthorEmailPattern;
            return this;
        }

        /**
         * @param commitAuthorEmailPattern (Block List, Max: 1) Parameters to be used for the commit_author_email_pattern rule. This rule only applies to repositories within an enterprise, it cannot be applied to repositories owned by individuals or regular organizations. (see below for nested schema)
         * 
         * @return builder
         * 
         */
        public Builder commitAuthorEmailPattern(OrganizationRulesetRulesCommitAuthorEmailPatternArgs commitAuthorEmailPattern) {
            return commitAuthorEmailPattern(Output.of(commitAuthorEmailPattern));
        }

        /**
         * @param commitMessagePattern (Block List, Max: 1) Parameters to be used for the commit_message_pattern rule. This rule only applies to repositories within an enterprise, it cannot be applied to repositories owned by individuals or regular organizations. (see below for nested schema)
         * 
         * @return builder
         * 
         */
        public Builder commitMessagePattern(@Nullable Output<OrganizationRulesetRulesCommitMessagePatternArgs> commitMessagePattern) {
            $.commitMessagePattern = commitMessagePattern;
            return this;
        }

        /**
         * @param commitMessagePattern (Block List, Max: 1) Parameters to be used for the commit_message_pattern rule. This rule only applies to repositories within an enterprise, it cannot be applied to repositories owned by individuals or regular organizations. (see below for nested schema)
         * 
         * @return builder
         * 
         */
        public Builder commitMessagePattern(OrganizationRulesetRulesCommitMessagePatternArgs commitMessagePattern) {
            return commitMessagePattern(Output.of(commitMessagePattern));
        }

        /**
         * @param committerEmailPattern (Block List, Max: 1) Parameters to be used for the committer_email_pattern rule. This rule only applies to repositories within an enterprise, it cannot be applied to repositories owned by individuals or regular organizations. (see below for nested schema)
         * 
         * @return builder
         * 
         */
        public Builder committerEmailPattern(@Nullable Output<OrganizationRulesetRulesCommitterEmailPatternArgs> committerEmailPattern) {
            $.committerEmailPattern = committerEmailPattern;
            return this;
        }

        /**
         * @param committerEmailPattern (Block List, Max: 1) Parameters to be used for the committer_email_pattern rule. This rule only applies to repositories within an enterprise, it cannot be applied to repositories owned by individuals or regular organizations. (see below for nested schema)
         * 
         * @return builder
         * 
         */
        public Builder committerEmailPattern(OrganizationRulesetRulesCommitterEmailPatternArgs committerEmailPattern) {
            return committerEmailPattern(Output.of(committerEmailPattern));
        }

        /**
         * @param creation (Boolean) Only allow users with bypass permission to create matching refs.
         * 
         * @return builder
         * 
         */
        public Builder creation(@Nullable Output<Boolean> creation) {
            $.creation = creation;
            return this;
        }

        /**
         * @param creation (Boolean) Only allow users with bypass permission to create matching refs.
         * 
         * @return builder
         * 
         */
        public Builder creation(Boolean creation) {
            return creation(Output.of(creation));
        }

        /**
         * @param deletion (Boolean) Only allow users with bypass permissions to delete matching refs.
         * 
         * @return builder
         * 
         */
        public Builder deletion(@Nullable Output<Boolean> deletion) {
            $.deletion = deletion;
            return this;
        }

        /**
         * @param deletion (Boolean) Only allow users with bypass permissions to delete matching refs.
         * 
         * @return builder
         * 
         */
        public Builder deletion(Boolean deletion) {
            return deletion(Output.of(deletion));
        }

        /**
         * @param nonFastForward (Boolean) Prevent users with push access from force pushing to branches.
         * 
         * @return builder
         * 
         */
        public Builder nonFastForward(@Nullable Output<Boolean> nonFastForward) {
            $.nonFastForward = nonFastForward;
            return this;
        }

        /**
         * @param nonFastForward (Boolean) Prevent users with push access from force pushing to branches.
         * 
         * @return builder
         * 
         */
        public Builder nonFastForward(Boolean nonFastForward) {
            return nonFastForward(Output.of(nonFastForward));
        }

        /**
         * @param pullRequest (Block List, Max: 1) Require all commits be made to a non-target branch and submitted via a pull request before they can be merged. (see below for nested schema)
         * 
         * @return builder
         * 
         */
        public Builder pullRequest(@Nullable Output<OrganizationRulesetRulesPullRequestArgs> pullRequest) {
            $.pullRequest = pullRequest;
            return this;
        }

        /**
         * @param pullRequest (Block List, Max: 1) Require all commits be made to a non-target branch and submitted via a pull request before they can be merged. (see below for nested schema)
         * 
         * @return builder
         * 
         */
        public Builder pullRequest(OrganizationRulesetRulesPullRequestArgs pullRequest) {
            return pullRequest(Output.of(pullRequest));
        }

        /**
         * @param requiredCodeScanning (Block List, Max: 1) Define which tools must provide code scanning results before the reference is updated. When configured, code scanning must be enabled and have results for both the commit and the reference being updated. Multiple code scanning tools can be specified. (see below for nested schema)
         * 
         * @return builder
         * 
         */
        public Builder requiredCodeScanning(@Nullable Output<OrganizationRulesetRulesRequiredCodeScanningArgs> requiredCodeScanning) {
            $.requiredCodeScanning = requiredCodeScanning;
            return this;
        }

        /**
         * @param requiredCodeScanning (Block List, Max: 1) Define which tools must provide code scanning results before the reference is updated. When configured, code scanning must be enabled and have results for both the commit and the reference being updated. Multiple code scanning tools can be specified. (see below for nested schema)
         * 
         * @return builder
         * 
         */
        public Builder requiredCodeScanning(OrganizationRulesetRulesRequiredCodeScanningArgs requiredCodeScanning) {
            return requiredCodeScanning(Output.of(requiredCodeScanning));
        }

        /**
         * @param requiredLinearHistory (Boolean) Prevent merge commits from being pushed to matching branches.
         * 
         * @return builder
         * 
         */
        public Builder requiredLinearHistory(@Nullable Output<Boolean> requiredLinearHistory) {
            $.requiredLinearHistory = requiredLinearHistory;
            return this;
        }

        /**
         * @param requiredLinearHistory (Boolean) Prevent merge commits from being pushed to matching branches.
         * 
         * @return builder
         * 
         */
        public Builder requiredLinearHistory(Boolean requiredLinearHistory) {
            return requiredLinearHistory(Output.of(requiredLinearHistory));
        }

        /**
         * @param requiredSignatures (Boolean) Commits pushed to matching branches must have verified signatures.
         * 
         * @return builder
         * 
         */
        public Builder requiredSignatures(@Nullable Output<Boolean> requiredSignatures) {
            $.requiredSignatures = requiredSignatures;
            return this;
        }

        /**
         * @param requiredSignatures (Boolean) Commits pushed to matching branches must have verified signatures.
         * 
         * @return builder
         * 
         */
        public Builder requiredSignatures(Boolean requiredSignatures) {
            return requiredSignatures(Output.of(requiredSignatures));
        }

        /**
         * @param requiredStatusChecks (Block List, Max: 1) Choose which status checks must pass before branches can be merged into a branch that matches this rule. When enabled, commits must first be pushed to another branch, then merged or pushed directly to a branch that matches this rule after status checks have passed. (see below for nested schema)
         * 
         * @return builder
         * 
         */
        public Builder requiredStatusChecks(@Nullable Output<OrganizationRulesetRulesRequiredStatusChecksArgs> requiredStatusChecks) {
            $.requiredStatusChecks = requiredStatusChecks;
            return this;
        }

        /**
         * @param requiredStatusChecks (Block List, Max: 1) Choose which status checks must pass before branches can be merged into a branch that matches this rule. When enabled, commits must first be pushed to another branch, then merged or pushed directly to a branch that matches this rule after status checks have passed. (see below for nested schema)
         * 
         * @return builder
         * 
         */
        public Builder requiredStatusChecks(OrganizationRulesetRulesRequiredStatusChecksArgs requiredStatusChecks) {
            return requiredStatusChecks(Output.of(requiredStatusChecks));
        }

        /**
         * @param requiredWorkflows (Block List, Max: 1) Define which Actions workflows must pass before changes can be merged into a branch matching the rule. Multiple workflows can be specified. (see below for nested schema)
         * 
         * @return builder
         * 
         */
        public Builder requiredWorkflows(@Nullable Output<OrganizationRulesetRulesRequiredWorkflowsArgs> requiredWorkflows) {
            $.requiredWorkflows = requiredWorkflows;
            return this;
        }

        /**
         * @param requiredWorkflows (Block List, Max: 1) Define which Actions workflows must pass before changes can be merged into a branch matching the rule. Multiple workflows can be specified. (see below for nested schema)
         * 
         * @return builder
         * 
         */
        public Builder requiredWorkflows(OrganizationRulesetRulesRequiredWorkflowsArgs requiredWorkflows) {
            return requiredWorkflows(Output.of(requiredWorkflows));
        }

        /**
         * @param tagNamePattern (Block List, Max: 1) Parameters to be used for the tag_name_pattern rule. This rule only applies to repositories within an enterprise, it cannot be applied to repositories owned by individuals or regular organizations. Conflicts with `branch_name_pattern` as it only applies to rulesets with target `tag`. (see below for nested schema)
         * 
         * @return builder
         * 
         */
        public Builder tagNamePattern(@Nullable Output<OrganizationRulesetRulesTagNamePatternArgs> tagNamePattern) {
            $.tagNamePattern = tagNamePattern;
            return this;
        }

        /**
         * @param tagNamePattern (Block List, Max: 1) Parameters to be used for the tag_name_pattern rule. This rule only applies to repositories within an enterprise, it cannot be applied to repositories owned by individuals or regular organizations. Conflicts with `branch_name_pattern` as it only applies to rulesets with target `tag`. (see below for nested schema)
         * 
         * @return builder
         * 
         */
        public Builder tagNamePattern(OrganizationRulesetRulesTagNamePatternArgs tagNamePattern) {
            return tagNamePattern(Output.of(tagNamePattern));
        }

        /**
         * @param update (Boolean) Only allow users with bypass permission to update matching refs.
         * 
         * @return builder
         * 
         */
        public Builder update(@Nullable Output<Boolean> update) {
            $.update = update;
            return this;
        }

        /**
         * @param update (Boolean) Only allow users with bypass permission to update matching refs.
         * 
         * @return builder
         * 
         */
        public Builder update(Boolean update) {
            return update(Output.of(update));
        }

        public OrganizationRulesetRulesArgs build() {
            return $;
        }
    }

}
