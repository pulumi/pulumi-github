// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.github.outputs.RepositoryRulesetRulesBranchNamePattern;
import com.pulumi.github.outputs.RepositoryRulesetRulesCommitAuthorEmailPattern;
import com.pulumi.github.outputs.RepositoryRulesetRulesCommitMessagePattern;
import com.pulumi.github.outputs.RepositoryRulesetRulesCommitterEmailPattern;
import com.pulumi.github.outputs.RepositoryRulesetRulesPullRequest;
import com.pulumi.github.outputs.RepositoryRulesetRulesRequiredDeployments;
import com.pulumi.github.outputs.RepositoryRulesetRulesRequiredStatusChecks;
import com.pulumi.github.outputs.RepositoryRulesetRulesTagNamePattern;
import java.lang.Boolean;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class RepositoryRulesetRules {
    /**
     * @return (Block List, Max: 1) Parameters to be used for the branch_name_pattern rule. This rule only applies to repositories within an enterprise, it cannot be applied to repositories owned by individuals or regular organizations. Conflicts with `tag_name_pattern` as it only applied to rulesets with target `branch`. (see below for nested schema)
     * 
     */
    private @Nullable RepositoryRulesetRulesBranchNamePattern branchNamePattern;
    /**
     * @return (Block List, Max: 1) Parameters to be used for the commit_author_email_pattern rule. This rule only applies to repositories within an enterprise, it cannot be applied to repositories owned by individuals or regular organizations. (see below for nested schema)
     * 
     */
    private @Nullable RepositoryRulesetRulesCommitAuthorEmailPattern commitAuthorEmailPattern;
    /**
     * @return (Block List, Max: 1) Parameters to be used for the commit_message_pattern rule. This rule only applies to repositories within an enterprise, it cannot be applied to repositories owned by individuals or regular organizations. (see below for nested schema)
     * 
     */
    private @Nullable RepositoryRulesetRulesCommitMessagePattern commitMessagePattern;
    /**
     * @return (Block List, Max: 1) Parameters to be used for the committer_email_pattern rule. This rule only applies to repositories within an enterprise, it cannot be applied to repositories owned by individuals or regular organizations. (see below for nested schema)
     * 
     */
    private @Nullable RepositoryRulesetRulesCommitterEmailPattern committerEmailPattern;
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
    private @Nullable RepositoryRulesetRulesPullRequest pullRequest;
    /**
     * @return (Block List, Max: 1) Choose which environments must be successfully deployed to before branches can be merged into a branch that matches this rule. (see below for nested schema)
     * 
     */
    private @Nullable RepositoryRulesetRulesRequiredDeployments requiredDeployments;
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
    private @Nullable RepositoryRulesetRulesRequiredStatusChecks requiredStatusChecks;
    /**
     * @return (Block List, Max: 1) Parameters to be used for the tag_name_pattern rule. This rule only applies to repositories within an enterprise, it cannot be applied to repositories owned by individuals or regular organizations. Conflicts with `branch_name_pattern` as it only applied to rulesets with target `tag`. (see below for nested schema)
     * 
     */
    private @Nullable RepositoryRulesetRulesTagNamePattern tagNamePattern;
    /**
     * @return (Boolean) Only allow users with bypass permission to update matching refs.
     * 
     */
    private @Nullable Boolean update;
    /**
     * @return (Boolean) Branch can pull changes from its upstream repository. This is only applicable to forked repositories. Requires `update` to be set to `true`. Note: behaviour is affected by a known bug on the GitHub side which may cause issues when using this parameter.
     * 
     */
    private @Nullable Boolean updateAllowsFetchAndMerge;

    private RepositoryRulesetRules() {}
    /**
     * @return (Block List, Max: 1) Parameters to be used for the branch_name_pattern rule. This rule only applies to repositories within an enterprise, it cannot be applied to repositories owned by individuals or regular organizations. Conflicts with `tag_name_pattern` as it only applied to rulesets with target `branch`. (see below for nested schema)
     * 
     */
    public Optional<RepositoryRulesetRulesBranchNamePattern> branchNamePattern() {
        return Optional.ofNullable(this.branchNamePattern);
    }
    /**
     * @return (Block List, Max: 1) Parameters to be used for the commit_author_email_pattern rule. This rule only applies to repositories within an enterprise, it cannot be applied to repositories owned by individuals or regular organizations. (see below for nested schema)
     * 
     */
    public Optional<RepositoryRulesetRulesCommitAuthorEmailPattern> commitAuthorEmailPattern() {
        return Optional.ofNullable(this.commitAuthorEmailPattern);
    }
    /**
     * @return (Block List, Max: 1) Parameters to be used for the commit_message_pattern rule. This rule only applies to repositories within an enterprise, it cannot be applied to repositories owned by individuals or regular organizations. (see below for nested schema)
     * 
     */
    public Optional<RepositoryRulesetRulesCommitMessagePattern> commitMessagePattern() {
        return Optional.ofNullable(this.commitMessagePattern);
    }
    /**
     * @return (Block List, Max: 1) Parameters to be used for the committer_email_pattern rule. This rule only applies to repositories within an enterprise, it cannot be applied to repositories owned by individuals or regular organizations. (see below for nested schema)
     * 
     */
    public Optional<RepositoryRulesetRulesCommitterEmailPattern> committerEmailPattern() {
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
    public Optional<RepositoryRulesetRulesPullRequest> pullRequest() {
        return Optional.ofNullable(this.pullRequest);
    }
    /**
     * @return (Block List, Max: 1) Choose which environments must be successfully deployed to before branches can be merged into a branch that matches this rule. (see below for nested schema)
     * 
     */
    public Optional<RepositoryRulesetRulesRequiredDeployments> requiredDeployments() {
        return Optional.ofNullable(this.requiredDeployments);
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
    public Optional<RepositoryRulesetRulesRequiredStatusChecks> requiredStatusChecks() {
        return Optional.ofNullable(this.requiredStatusChecks);
    }
    /**
     * @return (Block List, Max: 1) Parameters to be used for the tag_name_pattern rule. This rule only applies to repositories within an enterprise, it cannot be applied to repositories owned by individuals or regular organizations. Conflicts with `branch_name_pattern` as it only applied to rulesets with target `tag`. (see below for nested schema)
     * 
     */
    public Optional<RepositoryRulesetRulesTagNamePattern> tagNamePattern() {
        return Optional.ofNullable(this.tagNamePattern);
    }
    /**
     * @return (Boolean) Only allow users with bypass permission to update matching refs.
     * 
     */
    public Optional<Boolean> update() {
        return Optional.ofNullable(this.update);
    }
    /**
     * @return (Boolean) Branch can pull changes from its upstream repository. This is only applicable to forked repositories. Requires `update` to be set to `true`. Note: behaviour is affected by a known bug on the GitHub side which may cause issues when using this parameter.
     * 
     */
    public Optional<Boolean> updateAllowsFetchAndMerge() {
        return Optional.ofNullable(this.updateAllowsFetchAndMerge);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(RepositoryRulesetRules defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable RepositoryRulesetRulesBranchNamePattern branchNamePattern;
        private @Nullable RepositoryRulesetRulesCommitAuthorEmailPattern commitAuthorEmailPattern;
        private @Nullable RepositoryRulesetRulesCommitMessagePattern commitMessagePattern;
        private @Nullable RepositoryRulesetRulesCommitterEmailPattern committerEmailPattern;
        private @Nullable Boolean creation;
        private @Nullable Boolean deletion;
        private @Nullable Boolean nonFastForward;
        private @Nullable RepositoryRulesetRulesPullRequest pullRequest;
        private @Nullable RepositoryRulesetRulesRequiredDeployments requiredDeployments;
        private @Nullable Boolean requiredLinearHistory;
        private @Nullable Boolean requiredSignatures;
        private @Nullable RepositoryRulesetRulesRequiredStatusChecks requiredStatusChecks;
        private @Nullable RepositoryRulesetRulesTagNamePattern tagNamePattern;
        private @Nullable Boolean update;
        private @Nullable Boolean updateAllowsFetchAndMerge;
        public Builder() {}
        public Builder(RepositoryRulesetRules defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.branchNamePattern = defaults.branchNamePattern;
    	      this.commitAuthorEmailPattern = defaults.commitAuthorEmailPattern;
    	      this.commitMessagePattern = defaults.commitMessagePattern;
    	      this.committerEmailPattern = defaults.committerEmailPattern;
    	      this.creation = defaults.creation;
    	      this.deletion = defaults.deletion;
    	      this.nonFastForward = defaults.nonFastForward;
    	      this.pullRequest = defaults.pullRequest;
    	      this.requiredDeployments = defaults.requiredDeployments;
    	      this.requiredLinearHistory = defaults.requiredLinearHistory;
    	      this.requiredSignatures = defaults.requiredSignatures;
    	      this.requiredStatusChecks = defaults.requiredStatusChecks;
    	      this.tagNamePattern = defaults.tagNamePattern;
    	      this.update = defaults.update;
    	      this.updateAllowsFetchAndMerge = defaults.updateAllowsFetchAndMerge;
        }

        @CustomType.Setter
        public Builder branchNamePattern(@Nullable RepositoryRulesetRulesBranchNamePattern branchNamePattern) {
            this.branchNamePattern = branchNamePattern;
            return this;
        }
        @CustomType.Setter
        public Builder commitAuthorEmailPattern(@Nullable RepositoryRulesetRulesCommitAuthorEmailPattern commitAuthorEmailPattern) {
            this.commitAuthorEmailPattern = commitAuthorEmailPattern;
            return this;
        }
        @CustomType.Setter
        public Builder commitMessagePattern(@Nullable RepositoryRulesetRulesCommitMessagePattern commitMessagePattern) {
            this.commitMessagePattern = commitMessagePattern;
            return this;
        }
        @CustomType.Setter
        public Builder committerEmailPattern(@Nullable RepositoryRulesetRulesCommitterEmailPattern committerEmailPattern) {
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
        public Builder pullRequest(@Nullable RepositoryRulesetRulesPullRequest pullRequest) {
            this.pullRequest = pullRequest;
            return this;
        }
        @CustomType.Setter
        public Builder requiredDeployments(@Nullable RepositoryRulesetRulesRequiredDeployments requiredDeployments) {
            this.requiredDeployments = requiredDeployments;
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
        public Builder requiredStatusChecks(@Nullable RepositoryRulesetRulesRequiredStatusChecks requiredStatusChecks) {
            this.requiredStatusChecks = requiredStatusChecks;
            return this;
        }
        @CustomType.Setter
        public Builder tagNamePattern(@Nullable RepositoryRulesetRulesTagNamePattern tagNamePattern) {
            this.tagNamePattern = tagNamePattern;
            return this;
        }
        @CustomType.Setter
        public Builder update(@Nullable Boolean update) {
            this.update = update;
            return this;
        }
        @CustomType.Setter
        public Builder updateAllowsFetchAndMerge(@Nullable Boolean updateAllowsFetchAndMerge) {
            this.updateAllowsFetchAndMerge = updateAllowsFetchAndMerge;
            return this;
        }
        public RepositoryRulesetRules build() {
            final var _resultValue = new RepositoryRulesetRules();
            _resultValue.branchNamePattern = branchNamePattern;
            _resultValue.commitAuthorEmailPattern = commitAuthorEmailPattern;
            _resultValue.commitMessagePattern = commitMessagePattern;
            _resultValue.committerEmailPattern = committerEmailPattern;
            _resultValue.creation = creation;
            _resultValue.deletion = deletion;
            _resultValue.nonFastForward = nonFastForward;
            _resultValue.pullRequest = pullRequest;
            _resultValue.requiredDeployments = requiredDeployments;
            _resultValue.requiredLinearHistory = requiredLinearHistory;
            _resultValue.requiredSignatures = requiredSignatures;
            _resultValue.requiredStatusChecks = requiredStatusChecks;
            _resultValue.tagNamePattern = tagNamePattern;
            _resultValue.update = update;
            _resultValue.updateAllowsFetchAndMerge = updateAllowsFetchAndMerge;
            return _resultValue;
        }
    }
}
