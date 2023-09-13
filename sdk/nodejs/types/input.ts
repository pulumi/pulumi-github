// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";

export interface ActionsOrganizationPermissionsAllowedActionsConfig {
    /**
     * Whether GitHub-owned actions are allowed in the organization.
     */
    githubOwnedAllowed: pulumi.Input<boolean>;
    /**
     * Specifies a list of string-matching patterns to allow specific action(s). Wildcards, tags, and SHAs are allowed. For example, monalisa/octocat@*, monalisa/octocat@v2, monalisa/*."
     */
    patternsAlloweds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Whether actions in GitHub Marketplace from verified creators are allowed. Set to true to allow all GitHub Marketplace actions by verified creators.
     */
    verifiedAllowed?: pulumi.Input<boolean>;
}

export interface ActionsOrganizationPermissionsEnabledRepositoriesConfig {
    /**
     * List of repository IDs to enable for GitHub Actions.
     */
    repositoryIds: pulumi.Input<pulumi.Input<number>[]>;
}

export interface ActionsRepositoryPermissionsAllowedActionsConfig {
    /**
     * Whether GitHub-owned actions are allowed in the repository.
     */
    githubOwnedAllowed: pulumi.Input<boolean>;
    /**
     * Specifies a list of string-matching patterns to allow specific action(s). Wildcards, tags, and SHAs are allowed. For example, monalisa/octocat@*, monalisa/octocat@v2, monalisa/*."
     */
    patternsAlloweds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Whether actions in GitHub Marketplace from verified creators are allowed. Set to true to allow all GitHub Marketplace actions by verified creators.
     */
    verifiedAllowed?: pulumi.Input<boolean>;
}

export interface BranchProtectionRequiredPullRequestReview {
    /**
     * Dismiss approved reviews automatically when a new commit is pushed. Defaults to `false`.
     */
    dismissStaleReviews?: pulumi.Input<boolean>;
    /**
     * The list of actor Names/IDs with dismissal access. If not empty, `restrictDismissals` is ignored. Actor names must either begin with a "/" for users or the organization name followed by a "/" for teams.
     */
    dismissalRestrictions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The list of actor Names/IDs that are allowed to bypass pull request requirements. Actor names must either begin with a "/" for users or the organization name followed by a "/" for teams.
     */
    pullRequestBypassers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Require an approved review in pull requests including files with a designated code owner. Defaults to `false`.
     */
    requireCodeOwnerReviews?: pulumi.Input<boolean>;
    /**
     * Require that The most recent push must be approved by someone other than the last pusher.  Defaults to `false`
     */
    requireLastPushApproval?: pulumi.Input<boolean>;
    /**
     * Require x number of approvals to satisfy branch protection requirements. If this is specified it must be a number between 0-6. This requirement matches GitHub's API, see the upstream [documentation](https://developer.github.com/v3/repos/branches/#parameters-1) for more information.
     * (https://developer.github.com/v3/repos/branches/#parameters-1) for more information.
     */
    requiredApprovingReviewCount?: pulumi.Input<number>;
    /**
     * Restrict pull request review dismissals.
     */
    restrictDismissals?: pulumi.Input<boolean>;
}

export interface BranchProtectionRequiredStatusCheck {
    /**
     * The list of status checks to require in order to merge into this branch. No status checks are required by default.
     */
    contexts?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Require branches to be up to date before merging. Defaults to `false`.
     */
    strict?: pulumi.Input<boolean>;
}

export interface BranchProtectionV3RequiredPullRequestReviews {
    /**
     * Allow specific users, teams, or apps to bypass pull request requirements. See Bypass Pull Request Allowances below for details.
     */
    bypassPullRequestAllowances?: pulumi.Input<inputs.BranchProtectionV3RequiredPullRequestReviewsBypassPullRequestAllowances>;
    /**
     * Dismiss approved reviews automatically when a new commit is pushed. Defaults to `false`.
     */
    dismissStaleReviews?: pulumi.Input<boolean>;
    /**
     * The list of team slugs with dismissal access.
     * Always use `slug` of the team, **not** its name. Each team already **has** to have access to the repository.
     */
    dismissalTeams?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The list of user logins with dismissal access
     */
    dismissalUsers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * @deprecated Use enforce_admins instead
     */
    includeAdmins?: pulumi.Input<boolean>;
    /**
     * Require an approved review in pull requests including files with a designated code owner. Defaults to `false`.
     */
    requireCodeOwnerReviews?: pulumi.Input<boolean>;
    /**
     * Require x number of approvals to satisfy branch protection requirements. If this is specified it must be a number between 0-6. This requirement matches GitHub's API, see the upstream [documentation](https://developer.github.com/v3/repos/branches/#parameters-1) for more information.
     */
    requiredApprovingReviewCount?: pulumi.Input<number>;
}

export interface BranchProtectionV3RequiredPullRequestReviewsBypassPullRequestAllowances {
    /**
     * The list of app slugs allowed to bypass pull request requirements.
     */
    apps?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The list of team slugs allowed to bypass pull request requirements.
     */
    teams?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The list of user logins allowed to bypass pull request requirements.
     */
    users?: pulumi.Input<pulumi.Input<string>[]>;
}

export interface BranchProtectionV3RequiredStatusChecks {
    /**
     * The list of status checks to require in order to merge into this branch. No status checks are required by default. Checks should be strings containing the context and appId like so "context:app_id".
     */
    checks?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * [**DEPRECATED**] (Optional) The list of status checks to require in order to merge into this branch. No status checks are required by default.
     *
     * @deprecated GitHub is deprecating the use of `contexts`. Use a `checks` array instead.
     */
    contexts?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * @deprecated Use enforce_admins instead
     */
    includeAdmins?: pulumi.Input<boolean>;
    /**
     * Require branches to be up to date before merging. Defaults to `false`.
     */
    strict?: pulumi.Input<boolean>;
}

export interface BranchProtectionV3Restrictions {
    /**
     * The list of app slugs with push access.
     *
     * `restrictions` is only available for organization-owned repositories.
     */
    apps?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The list of team slugs with push access.
     * Always use `slug` of the team, **not** its name. Each team already **has** to have access to the repository.
     */
    teams?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The list of user logins with push access.
     */
    users?: pulumi.Input<pulumi.Input<string>[]>;
}

export interface OrganizationRulesetBypassActor {
    /**
     * (Number) The ID of the actor that can bypass a ruleset
     */
    actorId: pulumi.Input<number>;
    /**
     * The type of actor that can bypass a ruleset. Can be one of: `RepositoryRole`, `Team`, `Integration`, `OrganizationAdmin`.
     */
    actorType: pulumi.Input<string>;
    /**
     * (String) When the specified actor can bypass the ruleset. pullRequest means that an actor can only bypass rules on pull requests. Can be one of: `always`, `pullRequest`.
     */
    bypassMode?: pulumi.Input<string>;
}

export interface OrganizationRulesetConditions {
    /**
     * (Block List, Min: 1, Max: 1) (see below for nested schema)
     */
    refName: pulumi.Input<inputs.OrganizationRulesetConditionsRefName>;
    /**
     * The repository IDs that the ruleset applies to. One of these IDs must match for the condition to pass. Conflicts with `repositoryName`.
     */
    repositoryId?: pulumi.Input<number>;
    /**
     * Conflicts with `repositoryId`. (see below for nested schema)
     *
     * One of `repositoryId` and `repositoryName` must be set for the rule to target any repositories.
     */
    repositoryName?: pulumi.Input<inputs.OrganizationRulesetConditionsRepositoryName>;
}

export interface OrganizationRulesetConditionsRefName {
    /**
     * (List of String) Array of ref names or patterns to exclude. The condition will not pass if any of these patterns match.
     */
    excludes: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * (List of String) Array of ref names or patterns to include. One of these patterns must match for the condition to pass. Also accepts `~DEFAULT_BRANCH` to include the default branch or `~ALL` to include all branches.
     */
    includes: pulumi.Input<pulumi.Input<string>[]>;
}

export interface OrganizationRulesetConditionsRepositoryName {
    /**
     * (List of String) Array of ref names or patterns to exclude. The condition will not pass if any of these patterns match.
     */
    excludes: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * (List of String) Array of repository names or patterns to include. One of these patterns must match for the condition to pass. Also accepts `~ALL` to include all repositories.
     */
    inlcudes: pulumi.Input<pulumi.Input<string>[]>;
    protected?: pulumi.Input<boolean>;
}

export interface OrganizationRulesetRules {
    /**
     * (Block List, Max: 1) Parameters to be used for the branchNamePattern rule. This rule only applies to repositories within an enterprise, it cannot be applied to repositories owned by individuals or regular organizations. Conflicts with `tagNamePattern` as it only applies to rulesets with target `branch`. (see below for nested schema)
     */
    branchNamePattern?: pulumi.Input<inputs.OrganizationRulesetRulesBranchNamePattern>;
    /**
     * (Block List, Max: 1) Parameters to be used for the commitAuthorEmailPattern rule. This rule only applies to repositories within an enterprise, it cannot be applied to repositories owned by individuals or regular organizations. (see below for nested schema)
     */
    commitAuthorEmailPattern?: pulumi.Input<inputs.OrganizationRulesetRulesCommitAuthorEmailPattern>;
    /**
     * (Block List, Max: 1) Parameters to be used for the commitMessagePattern rule. This rule only applies to repositories within an enterprise, it cannot be applied to repositories owned by individuals or regular organizations. (see below for nested schema)
     */
    commitMessagePattern?: pulumi.Input<inputs.OrganizationRulesetRulesCommitMessagePattern>;
    /**
     * (Block List, Max: 1) Parameters to be used for the committerEmailPattern rule. This rule only applies to repositories within an enterprise, it cannot be applied to repositories owned by individuals or regular organizations. (see below for nested schema)
     */
    committerEmailPattern?: pulumi.Input<inputs.OrganizationRulesetRulesCommitterEmailPattern>;
    /**
     * (Boolean) Only allow users with bypass permission to create matching refs.
     */
    creation?: pulumi.Input<boolean>;
    /**
     * (Boolean) Only allow users with bypass permissions to delete matching refs.
     */
    deletion?: pulumi.Input<boolean>;
    /**
     * (Boolean) Prevent users with push access from force pushing to branches.
     */
    nonFastForward?: pulumi.Input<boolean>;
    /**
     * (Block List, Max: 1) Require all commits be made to a non-target branch and submitted via a pull request before they can be merged. (see below for nested schema)
     */
    pullRequest?: pulumi.Input<inputs.OrganizationRulesetRulesPullRequest>;
    /**
     * (Boolean) Prevent merge commits from being pushed to matching branches.
     */
    requiredLinearHistory?: pulumi.Input<boolean>;
    /**
     * (Boolean) Commits pushed to matching branches must have verified signatures.
     */
    requiredSignatures?: pulumi.Input<boolean>;
    /**
     * (Block List, Max: 1) Choose which status checks must pass before branches can be merged into a branch that matches this rule. When enabled, commits must first be pushed to another branch, then merged or pushed directly to a branch that matches this rule after status checks have passed. (see below for nested schema)
     */
    requiredStatusChecks?: pulumi.Input<inputs.OrganizationRulesetRulesRequiredStatusChecks>;
    /**
     * (Block List, Max: 1) Parameters to be used for the tagNamePattern rule. This rule only applies to repositories within an enterprise, it cannot be applied to repositories owned by individuals or regular organizations. Conflicts with `branchNamePattern` as it only applies to rulesets with target `tag`. (see below for nested schema)
     */
    tagNamePattern?: pulumi.Input<inputs.OrganizationRulesetRulesTagNamePattern>;
    /**
     * (Boolean) Only allow users with bypass permission to update matching refs.
     */
    update?: pulumi.Input<boolean>;
}

export interface OrganizationRulesetRulesBranchNamePattern {
    /**
     * (String) The name of the ruleset.
     */
    name?: pulumi.Input<string>;
    /**
     * (Boolean) If true, the rule will fail if the pattern matches.
     */
    negate?: pulumi.Input<boolean>;
    /**
     * (String) The operator to use for matching. Can be one of: `startsWith`, `endsWith`, `contains`, `regex`.
     */
    operator: pulumi.Input<string>;
    /**
     * (String) The pattern to match with.
     */
    pattern: pulumi.Input<string>;
}

export interface OrganizationRulesetRulesCommitAuthorEmailPattern {
    /**
     * (String) The name of the ruleset.
     */
    name?: pulumi.Input<string>;
    /**
     * (Boolean) If true, the rule will fail if the pattern matches.
     */
    negate?: pulumi.Input<boolean>;
    /**
     * (String) The operator to use for matching. Can be one of: `startsWith`, `endsWith`, `contains`, `regex`.
     */
    operator: pulumi.Input<string>;
    /**
     * (String) The pattern to match with.
     */
    pattern: pulumi.Input<string>;
}

export interface OrganizationRulesetRulesCommitMessagePattern {
    /**
     * (String) The name of the ruleset.
     */
    name?: pulumi.Input<string>;
    /**
     * (Boolean) If true, the rule will fail if the pattern matches.
     */
    negate?: pulumi.Input<boolean>;
    /**
     * (String) The operator to use for matching. Can be one of: `startsWith`, `endsWith`, `contains`, `regex`.
     */
    operator: pulumi.Input<string>;
    /**
     * (String) The pattern to match with.
     */
    pattern: pulumi.Input<string>;
}

export interface OrganizationRulesetRulesCommitterEmailPattern {
    /**
     * (String) The name of the ruleset.
     */
    name?: pulumi.Input<string>;
    /**
     * (Boolean) If true, the rule will fail if the pattern matches.
     */
    negate?: pulumi.Input<boolean>;
    /**
     * (String) The operator to use for matching. Can be one of: `startsWith`, `endsWith`, `contains`, `regex`.
     */
    operator: pulumi.Input<string>;
    /**
     * (String) The pattern to match with.
     */
    pattern: pulumi.Input<string>;
}

export interface OrganizationRulesetRulesPullRequest {
    /**
     * (Boolean) New, reviewable commits pushed will dismiss previous pull request review approvals. Defaults to `false`.
     */
    dismissStaleReviewsOnPush?: pulumi.Input<boolean>;
    /**
     * (Boolean) Require an approving review in pull requests that modify files that have a designated code owner. Defaults to `false`.
     */
    requireCodeOwnerReview?: pulumi.Input<boolean>;
    /**
     * (Boolean) Whether the most recent reviewable push must be approved by someone other than the person who pushed it. Defaults to `false`.
     */
    requireLastPushApproval?: pulumi.Input<boolean>;
    /**
     * (Number) The number of approving reviews that are required before a pull request can be merged. Defaults to `0`.
     */
    requiredApprovingReviewCount?: pulumi.Input<number>;
    /**
     * (Boolean) All conversations on code must be resolved before a pull request can be merged. Defaults to `false`.
     */
    requiredReviewThreadResolution?: pulumi.Input<boolean>;
}

export interface OrganizationRulesetRulesRequiredStatusChecks {
    /**
     * (Block Set, Min: 1) Status checks that are required. Several can be defined. (see below for nested schema)
     */
    requiredChecks: pulumi.Input<pulumi.Input<inputs.OrganizationRulesetRulesRequiredStatusChecksRequiredCheck>[]>;
    /**
     * (Boolean) Whether pull requests targeting a matching branch must be tested with the latest code. This setting will not take effect unless at least one status check is enabled. Defaults to `false`.
     */
    strictRequiredStatusChecksPolicy?: pulumi.Input<boolean>;
}

export interface OrganizationRulesetRulesRequiredStatusChecksRequiredCheck {
    /**
     * (String) The status check context name that must be present on the commit.
     */
    context: pulumi.Input<string>;
    /**
     * (Number) The optional integration ID that this status check must originate from.
     */
    integrationId?: pulumi.Input<number>;
}

export interface OrganizationRulesetRulesTagNamePattern {
    /**
     * (String) The name of the ruleset.
     */
    name?: pulumi.Input<string>;
    /**
     * (Boolean) If true, the rule will fail if the pattern matches.
     */
    negate?: pulumi.Input<boolean>;
    /**
     * (String) The operator to use for matching. Can be one of: `startsWith`, `endsWith`, `contains`, `regex`.
     */
    operator: pulumi.Input<string>;
    /**
     * (String) The pattern to match with.
     */
    pattern: pulumi.Input<string>;
}

export interface OrganizationWebhookConfiguration {
    contentType?: pulumi.Input<string>;
    insecureSsl?: pulumi.Input<boolean>;
    secret?: pulumi.Input<string>;
    /**
     * URL of the webhook
     */
    url: pulumi.Input<string>;
}

export interface ProviderAppAuth {
    id: pulumi.Input<string>;
    installationId: pulumi.Input<string>;
    pemFile: pulumi.Input<string>;
}

export interface RepositoryCollaboratorsTeam {
    /**
     * The permission of the outside collaborators for the repository.
     * Must be one of `pull`, `triage`, `push`, `maintain`, `admin` or the name of an existing [custom repository role](https://docs.github.com/en/enterprise-cloud@latest/organizations/managing-peoples-access-to-your-organization-with-roles/managing-custom-repository-roles-for-an-organization) within the organisation. Defaults to `pull`.
     * Must be `push` for personal repositories. Defaults to `push`.
     */
    permission?: pulumi.Input<string>;
    /**
     * The GitHub team id or the GitHub team slug
     */
    teamId: pulumi.Input<string>;
}

export interface RepositoryCollaboratorsUser {
    /**
     * The permission of the outside collaborators for the repository.
     * Must be one of `pull`, `push`, `maintain`, `triage` or `admin` or the name of an existing [custom repository role](https://docs.github.com/en/enterprise-cloud@latest/organizations/managing-peoples-access-to-your-organization-with-roles/managing-custom-repository-roles-for-an-organization) within the organization for organization-owned repositories.
     * Must be `push` for personal repositories. Defaults to `push`.
     */
    permission?: pulumi.Input<string>;
    username: pulumi.Input<string>;
}

export interface RepositoryEnvironmentDeploymentBranchPolicy {
    /**
     * Whether only branches that match the specified name patterns can deploy to this environment.
     */
    customBranchPolicies: pulumi.Input<boolean>;
    /**
     * Whether only branches with branch protection rules can deploy to this environment.
     */
    protectedBranches: pulumi.Input<boolean>;
}

export interface RepositoryEnvironmentReviewer {
    /**
     * Up to 6 IDs for teams who may review jobs that reference the environment. Reviewers must have at least read access to the repository. Only one of the required reviewers needs to approve the job for it to proceed.
     */
    teams?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * Up to 6 IDs for users who may review jobs that reference the environment. Reviewers must have at least read access to the repository. Only one of the required reviewers needs to approve the job for it to proceed.
     */
    users?: pulumi.Input<pulumi.Input<number>[]>;
}

export interface RepositoryPages {
    /**
     * The type of GitHub Pages site to build. Can be `legacy` or `workflow`. If you use `legacy` as build type you need to set the option `source`.
     */
    buildType?: pulumi.Input<string>;
    /**
     * The custom domain for the repository. This can only be set after the repository has been created.
     */
    cname?: pulumi.Input<string>;
    /**
     * Whether the rendered GitHub Pages site has a custom 404 page.
     */
    custom404?: pulumi.Input<boolean>;
    /**
     * The absolute URL (including scheme) of the rendered GitHub Pages site e.g. `https://username.github.io`.
     */
    htmlUrl?: pulumi.Input<string>;
    /**
     * The source branch and directory for the rendered Pages site. See GitHub Pages Source below for details.
     */
    source?: pulumi.Input<inputs.RepositoryPagesSource>;
    /**
     * Set to `enabled` to enable advanced security features on the repository. Can be `enabled` or `disabled`.
     */
    status?: pulumi.Input<string>;
    url?: pulumi.Input<string>;
}

export interface RepositoryPagesSource {
    /**
     * The repository branch used to publish the site's source files. (i.e. `main` or `gh-pages`.
     */
    branch: pulumi.Input<string>;
    /**
     * The repository directory from which the site publishes (Default: `/`).
     */
    path?: pulumi.Input<string>;
}

export interface RepositoryRulesetBypassActor {
    /**
     * (Number) The ID of the actor that can bypass a ruleset
     */
    actorId: pulumi.Input<number>;
    /**
     * The type of actor that can bypass a ruleset. Can be one of: `RepositoryRole`, `Team`, `Integration`, `OrganizationAdmin`.
     */
    actorType: pulumi.Input<string>;
    /**
     * (String) When the specified actor can bypass the ruleset. pullRequest means that an actor can only bypass rules on pull requests. Can be one of: `always`, `pullRequest`.
     */
    bypassMode?: pulumi.Input<string>;
}

export interface RepositoryRulesetConditions {
    /**
     * (Block List, Min: 1, Max: 1) (see below for nested schema)
     */
    refName: pulumi.Input<inputs.RepositoryRulesetConditionsRefName>;
}

export interface RepositoryRulesetConditionsRefName {
    /**
     * (List of String) Array of ref names or patterns to exclude. The condition will not pass if any of these patterns match.
     */
    excludes: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * (List of String) Array of ref names or patterns to include. One of these patterns must match for the condition to pass. Also accepts `~DEFAULT_BRANCH` to include the default branch or `~ALL` to include all branches.
     */
    includes: pulumi.Input<pulumi.Input<string>[]>;
}

export interface RepositoryRulesetRules {
    /**
     * (Block List, Max: 1) Parameters to be used for the branchNamePattern rule. This rule only applies to repositories within an enterprise, it cannot be applied to repositories owned by individuals or regular organizations. Conflicts with `tagNamePattern` as it only applied to rulesets with target `branch`. (see below for nested schema)
     */
    branchNamePattern?: pulumi.Input<inputs.RepositoryRulesetRulesBranchNamePattern>;
    /**
     * (Block List, Max: 1) Parameters to be used for the commitAuthorEmailPattern rule. This rule only applies to repositories within an enterprise, it cannot be applied to repositories owned by individuals or regular organizations. (see below for nested schema)
     */
    commitAuthorEmailPattern?: pulumi.Input<inputs.RepositoryRulesetRulesCommitAuthorEmailPattern>;
    /**
     * (Block List, Max: 1) Parameters to be used for the commitMessagePattern rule. This rule only applies to repositories within an enterprise, it cannot be applied to repositories owned by individuals or regular organizations. (see below for nested schema)
     */
    commitMessagePattern?: pulumi.Input<inputs.RepositoryRulesetRulesCommitMessagePattern>;
    /**
     * (Block List, Max: 1) Parameters to be used for the committerEmailPattern rule. This rule only applies to repositories within an enterprise, it cannot be applied to repositories owned by individuals or regular organizations. (see below for nested schema)
     */
    committerEmailPattern?: pulumi.Input<inputs.RepositoryRulesetRulesCommitterEmailPattern>;
    /**
     * (Boolean) Only allow users with bypass permission to create matching refs.
     */
    creation?: pulumi.Input<boolean>;
    /**
     * (Boolean) Only allow users with bypass permissions to delete matching refs.
     */
    deletion?: pulumi.Input<boolean>;
    /**
     * (Boolean) Prevent users with push access from force pushing to branches.
     */
    nonFastForward?: pulumi.Input<boolean>;
    /**
     * (Block List, Max: 1) Require all commits be made to a non-target branch and submitted via a pull request before they can be merged. (see below for nested schema)
     */
    pullRequest?: pulumi.Input<inputs.RepositoryRulesetRulesPullRequest>;
    /**
     * (Block List, Max: 1) Choose which environments must be successfully deployed to before branches can be merged into a branch that matches this rule. (see below for nested schema)
     */
    requiredDeployments?: pulumi.Input<inputs.RepositoryRulesetRulesRequiredDeployments>;
    /**
     * (Boolean) Prevent merge commits from being pushed to matching branches.
     */
    requiredLinearHistory?: pulumi.Input<boolean>;
    /**
     * (Boolean) Commits pushed to matching branches must have verified signatures.
     */
    requiredSignatures?: pulumi.Input<boolean>;
    /**
     * (Block List, Max: 1) Choose which status checks must pass before branches can be merged into a branch that matches this rule. When enabled, commits must first be pushed to another branch, then merged or pushed directly to a branch that matches this rule after status checks have passed. (see below for nested schema)
     */
    requiredStatusChecks?: pulumi.Input<inputs.RepositoryRulesetRulesRequiredStatusChecks>;
    /**
     * (Block List, Max: 1) Parameters to be used for the tagNamePattern rule. This rule only applies to repositories within an enterprise, it cannot be applied to repositories owned by individuals or regular organizations. Conflicts with `branchNamePattern` as it only applied to rulesets with target `tag`. (see below for nested schema)
     */
    tagNamePattern?: pulumi.Input<inputs.RepositoryRulesetRulesTagNamePattern>;
    /**
     * (Boolean) Only allow users with bypass permission to update matching refs.
     */
    update?: pulumi.Input<boolean>;
    /**
     * (Boolean) Branch can pull changes from its upstream repository. This is only applicable to forked repositories. Requires `update` to be set to `true`. Note: behaviour is affected by a known bug on the GitHub side which may cause issues when using this parameter.
     */
    updateAllowsFetchAndMerge?: pulumi.Input<boolean>;
}

export interface RepositoryRulesetRulesBranchNamePattern {
    /**
     * (String) The name of the ruleset.
     */
    name?: pulumi.Input<string>;
    /**
     * (Boolean) If true, the rule will fail if the pattern matches.
     */
    negate?: pulumi.Input<boolean>;
    /**
     * (String) The operator to use for matching. Can be one of: `startsWith`, `endsWith`, `contains`, `regex`.
     */
    operator: pulumi.Input<string>;
    /**
     * (String) The pattern to match with.
     */
    pattern: pulumi.Input<string>;
}

export interface RepositoryRulesetRulesCommitAuthorEmailPattern {
    /**
     * (String) The name of the ruleset.
     */
    name?: pulumi.Input<string>;
    /**
     * (Boolean) If true, the rule will fail if the pattern matches.
     */
    negate?: pulumi.Input<boolean>;
    /**
     * (String) The operator to use for matching. Can be one of: `startsWith`, `endsWith`, `contains`, `regex`.
     */
    operator: pulumi.Input<string>;
    /**
     * (String) The pattern to match with.
     */
    pattern: pulumi.Input<string>;
}

export interface RepositoryRulesetRulesCommitMessagePattern {
    /**
     * (String) The name of the ruleset.
     */
    name?: pulumi.Input<string>;
    /**
     * (Boolean) If true, the rule will fail if the pattern matches.
     */
    negate?: pulumi.Input<boolean>;
    /**
     * (String) The operator to use for matching. Can be one of: `startsWith`, `endsWith`, `contains`, `regex`.
     */
    operator: pulumi.Input<string>;
    /**
     * (String) The pattern to match with.
     */
    pattern: pulumi.Input<string>;
}

export interface RepositoryRulesetRulesCommitterEmailPattern {
    /**
     * (String) The name of the ruleset.
     */
    name?: pulumi.Input<string>;
    /**
     * (Boolean) If true, the rule will fail if the pattern matches.
     */
    negate?: pulumi.Input<boolean>;
    /**
     * (String) The operator to use for matching. Can be one of: `startsWith`, `endsWith`, `contains`, `regex`.
     */
    operator: pulumi.Input<string>;
    /**
     * (String) The pattern to match with.
     */
    pattern: pulumi.Input<string>;
}

export interface RepositoryRulesetRulesPullRequest {
    /**
     * (Boolean) New, reviewable commits pushed will dismiss previous pull request review approvals. Defaults to `false`.
     */
    dismissStaleReviewsOnPush?: pulumi.Input<boolean>;
    /**
     * (Boolean) Require an approving review in pull requests that modify files that have a designated code owner. Defaults to `false`.
     */
    requireCodeOwnerReview?: pulumi.Input<boolean>;
    /**
     * (Boolean) Whether the most recent reviewable push must be approved by someone other than the person who pushed it. Defaults to `false`.
     */
    requireLastPushApproval?: pulumi.Input<boolean>;
    /**
     * (Number) The number of approving reviews that are required before a pull request can be merged. Defaults to `0`.
     */
    requiredApprovingReviewCount?: pulumi.Input<number>;
    /**
     * (Boolean) All conversations on code must be resolved before a pull request can be merged. Defaults to `false`.
     */
    requiredReviewThreadResolution?: pulumi.Input<boolean>;
}

export interface RepositoryRulesetRulesRequiredDeployments {
    /**
     * (List of String) The environments that must be successfully deployed to before branches can be merged.
     */
    requiredDeploymentEnvironments: pulumi.Input<pulumi.Input<string>[]>;
}

export interface RepositoryRulesetRulesRequiredStatusChecks {
    /**
     * (Block Set, Min: 1) Status checks that are required. Several can be defined. (see below for nested schema)
     */
    requiredChecks: pulumi.Input<pulumi.Input<inputs.RepositoryRulesetRulesRequiredStatusChecksRequiredCheck>[]>;
    /**
     * (Boolean) Whether pull requests targeting a matching branch must be tested with the latest code. This setting will not take effect unless at least one status check is enabled. Defaults to `false`.
     */
    strictRequiredStatusChecksPolicy?: pulumi.Input<boolean>;
}

export interface RepositoryRulesetRulesRequiredStatusChecksRequiredCheck {
    /**
     * (String) The status check context name that must be present on the commit.
     */
    context: pulumi.Input<string>;
    /**
     * (Number) The optional integration ID that this status check must originate from.
     */
    integrationId?: pulumi.Input<number>;
}

export interface RepositoryRulesetRulesTagNamePattern {
    /**
     * (String) The name of the ruleset.
     */
    name?: pulumi.Input<string>;
    /**
     * (Boolean) If true, the rule will fail if the pattern matches.
     */
    negate?: pulumi.Input<boolean>;
    /**
     * (String) The operator to use for matching. Can be one of: `startsWith`, `endsWith`, `contains`, `regex`.
     */
    operator: pulumi.Input<string>;
    /**
     * (String) The pattern to match with.
     */
    pattern: pulumi.Input<string>;
}

export interface RepositorySecurityAndAnalysis {
    /**
     * The advanced security configuration for the repository. See Advanced Security Configuration below for details. If a repository's visibility is `public`, advanced security is always enabled and cannot be changed, so this setting cannot be supplied.
     */
    advancedSecurity?: pulumi.Input<inputs.RepositorySecurityAndAnalysisAdvancedSecurity>;
    /**
     * The secret scanning configuration for the repository. See Secret Scanning Configuration below for details.
     */
    secretScanning?: pulumi.Input<inputs.RepositorySecurityAndAnalysisSecretScanning>;
    /**
     * The secret scanning push protection configuration for the repository. See Secret Scanning Push Protection Configuration below for details.
     */
    secretScanningPushProtection?: pulumi.Input<inputs.RepositorySecurityAndAnalysisSecretScanningPushProtection>;
}

export interface RepositorySecurityAndAnalysisAdvancedSecurity {
    /**
     * Set to `enabled` to enable advanced security features on the repository. Can be `enabled` or `disabled`.
     */
    status: pulumi.Input<string>;
}

export interface RepositorySecurityAndAnalysisSecretScanning {
    /**
     * Set to `enabled` to enable advanced security features on the repository. Can be `enabled` or `disabled`.
     */
    status: pulumi.Input<string>;
}

export interface RepositorySecurityAndAnalysisSecretScanningPushProtection {
    /**
     * Set to `enabled` to enable advanced security features on the repository. Can be `enabled` or `disabled`.
     */
    status: pulumi.Input<string>;
}

export interface RepositoryTemplate {
    /**
     * Whether the new repository should include all the branches from the template repository (defaults to false, which includes only the default branch from the template).
     */
    includeAllBranches?: pulumi.Input<boolean>;
    /**
     * The GitHub organization or user the template repository is owned by.
     */
    owner: pulumi.Input<string>;
    /**
     * The name of the template repository.
     */
    repository: pulumi.Input<string>;
}

export interface RepositoryWebhookConfiguration {
    /**
     * The content type for the payload. Valid values are either `form` or `json`.
     */
    contentType?: pulumi.Input<string>;
    /**
     * Insecure SSL boolean toggle. Defaults to `false`.
     */
    insecureSsl?: pulumi.Input<boolean>;
    /**
     * The shared secret for the webhook. [See API documentation](https://developer.github.com/v3/repos/hooks/#create-a-hook).
     */
    secret?: pulumi.Input<string>;
    /**
     * The URL of the webhook.
     */
    url: pulumi.Input<string>;
}

export interface TeamMembersMember {
    /**
     * The role of the user within the team.
     * Must be one of `member` or `maintainer`. Defaults to `member`.
     */
    role?: pulumi.Input<string>;
    /**
     * The user to add to the team.
     */
    username: pulumi.Input<string>;
}

export interface TeamSettingsReviewRequestDelegation {
    /**
     * The algorithm to use when assigning pull requests to team members. Supported values are `ROUND_ROBIN` and `LOAD_BALANCE`. Default value is `ROUND_ROBIN`
     */
    algorithm?: pulumi.Input<string>;
    /**
     * The number of team members to assign to a pull request
     */
    memberCount?: pulumi.Input<number>;
    /**
     * whether to notify the entire team when at least one member is also assigned to the pull request
     */
    notify?: pulumi.Input<boolean>;
}

export interface TeamSyncGroupMappingGroup {
    /**
     * The description of the IdP group.
     */
    groupDescription: pulumi.Input<string>;
    /**
     * The ID of the IdP group.
     */
    groupId: pulumi.Input<string>;
    /**
     * The name of the IdP group.
     */
    groupName: pulumi.Input<string>;
}
export namespace config {
}
