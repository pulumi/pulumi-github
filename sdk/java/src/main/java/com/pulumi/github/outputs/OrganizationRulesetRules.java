// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.github.outputs.OrganizationRulesetRulesBranchNamePattern;
import com.pulumi.github.outputs.OrganizationRulesetRulesCommitAuthorEmailPattern;
import com.pulumi.github.outputs.OrganizationRulesetRulesCommitMessagePattern;
import com.pulumi.github.outputs.OrganizationRulesetRulesCommitterEmailPattern;
import com.pulumi.github.outputs.OrganizationRulesetRulesPullRequest;
import com.pulumi.github.outputs.OrganizationRulesetRulesRequiredStatusChecks;
import com.pulumi.github.outputs.OrganizationRulesetRulesRequiredWorkflows;
import com.pulumi.github.outputs.OrganizationRulesetRulesTagNamePattern;
import java.lang.Boolean;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class OrganizationRulesetRules {
    /**
     * @return (Block List, Max: 1) Parameters to be used for the branch_name_pattern rule. This rule only applies to repositories within an enterprise, it cannot be applied to repositories owned by individuals or regular organizations. Conflicts with `tag_name_pattern` as it only applies to rulesets with target `branch`. (see below for nested schema)
     * 
     */
    private @Nullable OrganizationRulesetRulesBranchNamePattern branchNamePattern;
    /**
     * @return (Block List, Max: 1) Parameters to be used for the commit_author_email_pattern rule. This rule only applies to repositories within an enterprise, it cannot be applied to repositories owned by individuals or regular organizations. (see below for nested schema)
     * 
     */
    private @Nullable OrganizationRulesetRulesCommitAuthorEmailPattern commitAuthorEmailPattern;
    /**
     * @return (Block List, Max: 1) Parameters to be used for the commit_message_pattern rule. This rule only applies to repositories within an enterprise, it cannot be applied to repositories owned by individuals or regular organizations. (see below for nested schema)
     * 
     */
    private @Nullable OrganizationRulesetRulesCommitMessagePattern commitMessagePattern;
    /**
     * @return (Block List, Max: 1) Parameters to be used for the committer_email_pattern rule. This rule only applies to repositories within an enterprise, it cannot be applied to repositories owned by individuals or regular organizations. (see below for nested schema)
     * 
     */
    private @Nullable OrganizationRulesetRulesCommitterEmailPattern committerEmailPattern;
    /**
     * @return (Boolean) Only allow users with bypass permission to create matching refs.
     * 
     */
    private @Nullable Boolean creation;
    /**
     * @return (Boolean) Only allow users with bypass permissions to delete matching refs.
     * 
     */
    private @Nullable Boolean deletion;
    /**
     * @return (Boolean) Prevent users with push access from force pushing to branches.
     * 
     */
    private @Nullable Boolean nonFastForward;
    /**
     * @return (Block List, Max: 1) Require all commits be made to a non-target branch and submitted via a pull request before they can be merged. (see below for nested schema)
     * 
     */
    private @Nullable OrganizationRulesetRulesPullRequest pullRequest;
    /**
     * @return (Boolean) Prevent merge commits from being pushed to matching branches.
     * 
     */
    private @Nullable Boolean requiredLinearHistory;
    /**
     * @return (Boolean) Commits pushed to matching branches must have verified signatures.
     * 
     */
    private @Nullable Boolean requiredSignatures;
    /**
     * @return (Block List, Max: 1) Choose which status checks must pass before branches can be merged into a branch that matches this rule. When enabled, commits must first be pushed to another branch, then merged or pushed directly to a branch that matches this rule after status checks have passed. (see below for nested schema)
     * 
     */
    private @Nullable OrganizationRulesetRulesRequiredStatusChecks requiredStatusChecks;
    /**
     * @return (Block List, Max: 1) Define which Actions workflows must pass before changes can be merged into a branch matching the rule. Multiple workflows can be specified. (see below for nested schema)
     * 
     */
    private @Nullable OrganizationRulesetRulesRequiredWorkflows requiredWorkflows;
    /**
     * @return (Block List, Max: 1) Parameters to be used for the tag_name_pattern rule. This rule only applies to repositories within an enterprise, it cannot be applied to repositories owned by individuals or regular organizations. Conflicts with `branch_name_pattern` as it only applies to rulesets with target `tag`. (see below for nested schema)
     * 
     */
    private @Nullable OrganizationRulesetRulesTagNamePattern tagNamePattern;
    /**
     * @return (Boolean) Only allow users with bypass permission to update matching refs.
     * 
     */
    private @Nullable Boolean update;

    private OrganizationRulesetRules() {}
    /**
     * @return (Block List, Max: 1) Parameters to be used for the branch_name_pattern rule. This rule only applies to repositories within an enterprise, it cannot be applied to repositories owned by individuals or regular organizations. Conflicts with `tag_name_pattern` as it only applies to rulesets with target `branch`. (see below for nested schema)
     * 
     */
    public Optional<OrganizationRulesetRulesBranchNamePattern> branchNamePattern() {
        return Optional.ofNullable(this.branchNamePattern);
    }
    /**
     * @return (Block List, Max: 1) Parameters to be used for the commit_author_email_pattern rule. This rule only applies to repositories within an enterprise, it cannot be applied to repositories owned by individuals or regular organizations. (see below for nested schema)
     * 
     */
    public Optional<OrganizationRulesetRulesCommitAuthorEmailPattern> commitAuthorEmailPattern() {
        return Optional.ofNullable(this.commitAuthorEmailPattern);
    }
    /**
     * @return (Block List, Max: 1) Parameters to be used for the commit_message_pattern rule. This rule only applies to repositories within an enterprise, it cannot be applied to repositories owned by individuals or regular organizations. (see below for nested schema)
     * 
     */
    public Optional<OrganizationRulesetRulesCommitMessagePattern> commitMessagePattern() {
        return Optional.ofNullable(this.commitMessagePattern);
    }
    /**
     * @return (Block List, Max: 1) Parameters to be used for the committer_email_pattern rule. This rule only applies to repositories within an enterprise, it cannot be applied to repositories owned by individuals or regular organizations. (see below for nested schema)
     * 
     */
    public Optional<OrganizationRulesetRulesCommitterEmailPattern> committerEmailPattern() {
        return Optional.ofNullable(this.committerEmailPattern);
    }
    /**
     * @return (Boolean) Only allow users with bypass permission to create matching refs.
     * 
     */
    public Optional<Boolean> creation() {
        return Optional.ofNullable(this.creation);
    }
    /**
     * @return (Boolean) Only allow users with bypass permissions to delete matching refs.
     * 
     */
    public Optional<Boolean> deletion() {
        return Optional.ofNullable(this.deletion);
    }
    /**
     * @return (Boolean) Prevent users with push access from force pushing to branches.
     * 
     */
    public Optional<Boolean> nonFastForward() {
        return Optional.ofNullable(this.nonFastForward);
    }
    /**
     * @return (Block List, Max: 1) Require all commits be made to a non-target branch and submitted via a pull request before they can be merged. (see below for nested schema)
     * 
     */
    public Optional<OrganizationRulesetRulesPullRequest> pullRequest() {
        return Optional.ofNullable(this.pullRequest);
    }
    /**
     * @return (Boolean) Prevent merge commits from being pushed to matching branches.
     * 
     */
    public Optional<Boolean> requiredLinearHistory() {
        return Optional.ofNullable(this.requiredLinearHistory);
    }
    /**
     * @return (Boolean) Commits pushed to matching branches must have verified signatures.
     * 
     */
    public Optional<Boolean> requiredSignatures() {
        return Optional.ofNullable(this.requiredSignatures);
    }
    /**
     * @return (Block List, Max: 1) Choose which status checks must pass before branches can be merged into a branch that matches this rule. When enabled, commits must first be pushed to another branch, then merged or pushed directly to a branch that matches this rule after status checks have passed. (see below for nested schema)
     * 
     */
    public Optional<OrganizationRulesetRulesRequiredStatusChecks> requiredStatusChecks() {
        return Optional.ofNullable(this.requiredStatusChecks);
    }
    /**
     * @return (Block List, Max: 1) Define which Actions workflows must pass before changes can be merged into a branch matching the rule. Multiple workflows can be specified. (see below for nested schema)
     * 
     */
    public Optional<OrganizationRulesetRulesRequiredWorkflows> requiredWorkflows() {
        return Optional.ofNullable(this.requiredWorkflows);
    }
    /**
     * @return (Block List, Max: 1) Parameters to be used for the tag_name_pattern rule. This rule only applies to repositories within an enterprise, it cannot be applied to repositories owned by individuals or regular organizations. Conflicts with `branch_name_pattern` as it only applies to rulesets with target `tag`. (see below for nested schema)
     * 
     */
    public Optional<OrganizationRulesetRulesTagNamePattern> tagNamePattern() {
        return Optional.ofNullable(this.tagNamePattern);
    }
    /**
     * @return (Boolean) Only allow users with bypass permission to update matching refs.
     * 
     */
    public Optional<Boolean> update() {
        return Optional.ofNullable(this.update);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(OrganizationRulesetRules defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable OrganizationRulesetRulesBranchNamePattern branchNamePattern;
        private @Nullable OrganizationRulesetRulesCommitAuthorEmailPattern commitAuthorEmailPattern;
        private @Nullable OrganizationRulesetRulesCommitMessagePattern commitMessagePattern;
        private @Nullable OrganizationRulesetRulesCommitterEmailPattern committerEmailPattern;
        private @Nullable Boolean creation;
        private @Nullable Boolean deletion;
        private @Nullable Boolean nonFastForward;
        private @Nullable OrganizationRulesetRulesPullRequest pullRequest;
        private @Nullable Boolean requiredLinearHistory;
        private @Nullable Boolean requiredSignatures;
        private @Nullable OrganizationRulesetRulesRequiredStatusChecks requiredStatusChecks;
        private @Nullable OrganizationRulesetRulesRequiredWorkflows requiredWorkflows;
        private @Nullable OrganizationRulesetRulesTagNamePattern tagNamePattern;
        private @Nullable Boolean update;
        public Builder() {}
        public Builder(OrganizationRulesetRules defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.branchNamePattern = defaults.branchNamePattern;
    	      this.commitAuthorEmailPattern = defaults.commitAuthorEmailPattern;
    	      this.commitMessagePattern = defaults.commitMessagePattern;
    	      this.committerEmailPattern = defaults.committerEmailPattern;
    	      this.creation = defaults.creation;
    	      this.deletion = defaults.deletion;
    	      this.nonFastForward = defaults.nonFastForward;
    	      this.pullRequest = defaults.pullRequest;
    	      this.requiredLinearHistory = defaults.requiredLinearHistory;
    	      this.requiredSignatures = defaults.requiredSignatures;
    	      this.requiredStatusChecks = defaults.requiredStatusChecks;
    	      this.requiredWorkflows = defaults.requiredWorkflows;
    	      this.tagNamePattern = defaults.tagNamePattern;
    	      this.update = defaults.update;
        }

        @CustomType.Setter
        public Builder branchNamePattern(@Nullable OrganizationRulesetRulesBranchNamePattern branchNamePattern) {

            this.branchNamePattern = branchNamePattern;
            return this;
        }
        @CustomType.Setter
        public Builder commitAuthorEmailPattern(@Nullable OrganizationRulesetRulesCommitAuthorEmailPattern commitAuthorEmailPattern) {

            this.commitAuthorEmailPattern = commitAuthorEmailPattern;
            return this;
        }
        @CustomType.Setter
        public Builder commitMessagePattern(@Nullable OrganizationRulesetRulesCommitMessagePattern commitMessagePattern) {

            this.commitMessagePattern = commitMessagePattern;
            return this;
        }
        @CustomType.Setter
        public Builder committerEmailPattern(@Nullable OrganizationRulesetRulesCommitterEmailPattern committerEmailPattern) {

            this.committerEmailPattern = committerEmailPattern;
            return this;
        }
        @CustomType.Setter
        public Builder creation(@Nullable Boolean creation) {

            this.creation = creation;
            return this;
        }
        @CustomType.Setter
        public Builder deletion(@Nullable Boolean deletion) {

            this.deletion = deletion;
            return this;
        }
        @CustomType.Setter
        public Builder nonFastForward(@Nullable Boolean nonFastForward) {

            this.nonFastForward = nonFastForward;
            return this;
        }
        @CustomType.Setter
        public Builder pullRequest(@Nullable OrganizationRulesetRulesPullRequest pullRequest) {

            this.pullRequest = pullRequest;
            return this;
        }
        @CustomType.Setter
        public Builder requiredLinearHistory(@Nullable Boolean requiredLinearHistory) {

            this.requiredLinearHistory = requiredLinearHistory;
            return this;
        }
        @CustomType.Setter
        public Builder requiredSignatures(@Nullable Boolean requiredSignatures) {

            this.requiredSignatures = requiredSignatures;
            return this;
        }
        @CustomType.Setter
        public Builder requiredStatusChecks(@Nullable OrganizationRulesetRulesRequiredStatusChecks requiredStatusChecks) {

            this.requiredStatusChecks = requiredStatusChecks;
            return this;
        }
        @CustomType.Setter
        public Builder requiredWorkflows(@Nullable OrganizationRulesetRulesRequiredWorkflows requiredWorkflows) {

            this.requiredWorkflows = requiredWorkflows;
            return this;
        }
        @CustomType.Setter
        public Builder tagNamePattern(@Nullable OrganizationRulesetRulesTagNamePattern tagNamePattern) {

            this.tagNamePattern = tagNamePattern;
            return this;
        }
        @CustomType.Setter
        public Builder update(@Nullable Boolean update) {

            this.update = update;
            return this;
        }
        public OrganizationRulesetRules build() {
            final var _resultValue = new OrganizationRulesetRules();
            _resultValue.branchNamePattern = branchNamePattern;
            _resultValue.commitAuthorEmailPattern = commitAuthorEmailPattern;
            _resultValue.commitMessagePattern = commitMessagePattern;
            _resultValue.committerEmailPattern = committerEmailPattern;
            _resultValue.creation = creation;
            _resultValue.deletion = deletion;
            _resultValue.nonFastForward = nonFastForward;
            _resultValue.pullRequest = pullRequest;
            _resultValue.requiredLinearHistory = requiredLinearHistory;
            _resultValue.requiredSignatures = requiredSignatures;
            _resultValue.requiredStatusChecks = requiredStatusChecks;
            _resultValue.requiredWorkflows = requiredWorkflows;
            _resultValue.tagNamePattern = tagNamePattern;
            _resultValue.update = update;
            return _resultValue;
        }
    }
}
