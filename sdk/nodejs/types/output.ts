// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";

export interface ActionsOrganizationPermissionsAllowedActionsConfig {
    /**
     * Whether GitHub-owned actions are allowed in the organization.
     */
    githubOwnedAllowed: boolean;
    /**
     * Specifies a list of string-matching patterns to allow specific action(s). Wildcards, tags, and SHAs are allowed. For example, monalisa/octocat@*, monalisa/octocat@v2, monalisa/*."
     */
    patternsAlloweds?: string[];
    /**
     * Whether actions in GitHub Marketplace from verified creators are allowed. Set to true to allow all GitHub Marketplace actions by verified creators.
     */
    verifiedAllowed?: boolean;
}

export interface ActionsOrganizationPermissionsEnabledRepositoriesConfig {
    /**
     * List of repository IDs to enable for GitHub Actions.
     */
    repositoryIds: number[];
}

export interface ActionsRepositoryPermissionsAllowedActionsConfig {
    /**
     * Whether GitHub-owned actions are allowed in the repository.
     */
    githubOwnedAllowed: boolean;
    /**
     * Specifies a list of string-matching patterns to allow specific action(s). Wildcards, tags, and SHAs are allowed. For example, monalisa/octocat@*, monalisa/octocat@v2, monalisa/*."
     */
    patternsAlloweds?: string[];
    /**
     * Whether actions in GitHub Marketplace from verified creators are allowed. Set to true to allow all GitHub Marketplace actions by verified creators.
     */
    verifiedAllowed?: boolean;
}

export interface BranchProtectionRequiredPullRequestReview {
    /**
     * Dismiss approved reviews automatically when a new commit is pushed. Defaults to `false`.
     */
    dismissStaleReviews?: boolean;
    /**
     * The list of actor Names/IDs with dismissal access. If not empty, `restrictDismissals` is ignored. Actor names must either begin with a "/" for users or the organization name followed by a "/" for teams.
     */
    dismissalRestrictions?: string[];
    /**
     * The list of actor Names/IDs that are allowed to bypass pull request requirements. Actor names must either begin with a "/" for users or the organization name followed by a "/" for teams.
     */
    pullRequestBypassers?: string[];
    /**
     * Require an approved review in pull requests including files with a designated code owner. Defaults to `false`.
     */
    requireCodeOwnerReviews?: boolean;
    /**
     * Require that The most recent push must be approved by someone other than the last pusher.  Defaults to `false`
     */
    requireLastPushApproval?: boolean;
    /**
     * Require x number of approvals to satisfy branch protection requirements. If this is specified it must be a number between 0-6. This requirement matches GitHub's API, see the upstream [documentation](https://developer.github.com/v3/repos/branches/#parameters-1) for more information.
     * (https://developer.github.com/v3/repos/branches/#parameters-1) for more information.
     */
    requiredApprovingReviewCount?: number;
    /**
     * Restrict pull request review dismissals.
     */
    restrictDismissals?: boolean;
}

export interface BranchProtectionRequiredStatusCheck {
    /**
     * The list of status checks to require in order to merge into this branch. No status checks are required by default.
     */
    contexts?: string[];
    /**
     * Require branches to be up to date before merging. Defaults to `false`.
     */
    strict?: boolean;
}

export interface BranchProtectionV3RequiredPullRequestReviews {
    /**
     * Dismiss approved reviews automatically when a new commit is pushed. Defaults to `false`.
     */
    dismissStaleReviews?: boolean;
    /**
     * The list of team slugs with dismissal access.
     * Always use `slug` of the team, **not** its name. Each team already **has** to have access to the repository.
     */
    dismissalTeams?: string[];
    /**
     * The list of user logins with dismissal access
     */
    dismissalUsers?: string[];
    /**
     * @deprecated Use enforce_admins instead
     */
    includeAdmins?: boolean;
    /**
     * Require an approved review in pull requests including files with a designated code owner. Defaults to `false`.
     */
    requireCodeOwnerReviews?: boolean;
    /**
     * Require x number of approvals to satisfy branch protection requirements. If this is specified it must be a number between 0-6. This requirement matches GitHub's API, see the upstream [documentation](https://developer.github.com/v3/repos/branches/#parameters-1) for more information.
     */
    requiredApprovingReviewCount?: number;
}

export interface BranchProtectionV3RequiredStatusChecks {
    /**
     * The list of status checks to require in order to merge into this branch. No status checks are required by default.
     */
    contexts?: string[];
    /**
     * @deprecated Use enforce_admins instead
     */
    includeAdmins?: boolean;
    /**
     * Require branches to be up to date before merging. Defaults to `false`.
     */
    strict?: boolean;
}

export interface BranchProtectionV3Restrictions {
    /**
     * The list of app slugs with push access.
     */
    apps?: string[];
    /**
     * The list of team slugs with push access.
     * Always use `slug` of the team, **not** its name. Each team already **has** to have access to the repository.
     */
    teams?: string[];
    /**
     * The list of user logins with push access.
     */
    users?: string[];
}

export interface GetActionsOrganizationSecretsSecret {
    /**
     * Timestamp of the secret creation
     */
    createdAt: string;
    /**
     * Secret name
     */
    name: string;
    /**
     * Timestamp of the secret last update
     */
    updatedAt: string;
    /**
     * Secret visibility
     */
    visibility: string;
}

export interface GetActionsSecretsSecret {
    /**
     * Timestamp of the secret creation
     */
    createdAt: string;
    /**
     * The name of the repository.
     */
    name: string;
    /**
     * Timestamp of the secret last update
     */
    updatedAt: string;
}

export interface GetCollaboratorsCollaborator {
    /**
     * The GitHub API URL for the collaborator's events.
     */
    eventsUrl: string;
    /**
     * The GitHub API URL for the collaborator's followers.
     */
    followersUrl: string;
    /**
     * The GitHub API URL for those following the collaborator.
     */
    followingUrl: string;
    /**
     * The GitHub API URL for the collaborator's gists.
     */
    gistsUrl: string;
    /**
     * The GitHub HTML URL for the collaborator.
     */
    htmlUrl: string;
    /**
     * The ID of the collaborator.
     */
    id: number;
    /**
     * The collaborator's login.
     */
    login: string;
    /**
     * The GitHub API URL for the collaborator's organizations.
     */
    organizationsUrl: string;
    /**
     * The permission of the collaborator.
     */
    permission: string;
    /**
     * The GitHub API URL for the collaborator's received events.
     */
    receivedEventsUrl: string;
    /**
     * The GitHub API URL for the collaborator's repositories.
     */
    reposUrl: string;
    /**
     * Whether the user is a GitHub admin.
     */
    siteAdmin: boolean;
    /**
     * The GitHub API URL for the collaborator's starred repositories.
     */
    starredUrl: string;
    /**
     * The GitHub API URL for the collaborator's subscribed repositories.
     */
    subscriptionsUrl: string;
    /**
     * The type of the collaborator (ex. `user`).
     */
    type: string;
    /**
     * The GitHub API URL for the collaborator.
     */
    url: string;
}

export interface GetDependabotOrganizationSecretsSecret {
    /**
     * Timestamp of the secret creation
     */
    createdAt: string;
    /**
     * Secret name
     */
    name: string;
    /**
     * Timestamp of the secret last update
     */
    updatedAt: string;
    /**
     * Secret visibility
     */
    visibility: string;
}

export interface GetDependabotSecretsSecret {
    /**
     * Timestamp of the secret creation
     */
    createdAt: string;
    /**
     * The name of the repository.
     */
    name: string;
    /**
     * Timestamp of the secret last update
     */
    updatedAt: string;
}

export interface GetExternalGroupsExternalGroup {
    /**
     * the ID of the group.
     */
    groupId: number;
    /**
     * the name of the group.
     */
    groupName: string;
    /**
     * the date the group was last updated.
     */
    updatedAt: string;
}

export interface GetOrganizationIpAllowListIpAllowList {
    /**
     * A single IP address or range of IP addresses in CIDR notation.
     */
    allowListValue: string;
    /**
     * Identifies the date and time when the object was created.
     */
    createdAt: string;
    /**
     * The ID of the IP allow list entry.
     */
    id: string;
    /**
     * Whether the entry is currently active.
     */
    isActive: boolean;
    /**
     * The name of the IP allow list entry.
     */
    name: string;
    /**
     * Identifies the date and time when the object was last updated.
     */
    updatedAt: string;
}

export interface GetOrganizationTeamSyncGroupsGroup {
    /**
     * The description of the IdP group.
     */
    groupDescription: string;
    /**
     * The ID of the IdP group.
     */
    groupId: string;
    /**
     * The name of the IdP group.
     */
    groupName: string;
}

export interface GetOrganizationTeamsTeam {
    /**
     * the team's description.
     */
    description: string;
    /**
     * the ID of the team.
     */
    id: number;
    /**
     * List of team members. Not returned if `summaryOnly = true`
     */
    members: string[];
    /**
     * the team's full name.
     */
    name: string;
    /**
     * the Node ID of the team.
     */
    nodeId: string;
    /**
     * the team's privacy type.
     */
    privacy: string;
    /**
     * List of team repositories. Not returned if `summaryOnly = true`
     */
    repositories: string[];
    /**
     * the slug of the team.
     */
    slug: string;
}

export interface GetOrganizationWebhooksWebhook {
    /**
     * `true` if the webhook is active.
     */
    active: boolean;
    /**
     * the ID of the webhook.
     */
    id: number;
    /**
     * the name of the webhook.
     */
    name: string;
    /**
     * the type of the webhook.
     */
    type: string;
    /**
     * the url of the webhook.
     */
    url: string;
}

export interface GetReleaseAsset {
    /**
     * Browser download URL
     */
    browserDownloadUrl: string;
    /**
     * MIME type of the asset
     */
    contentType: string;
    /**
     * Date the asset was created
     */
    createdAt: string;
    /**
     * ID of the asset
     */
    id: number;
    /**
     * Label for the asset
     */
    label: string;
    /**
     * The file name of the asset
     */
    name: string;
    /**
     * Node ID of the asset
     */
    nodeId: string;
    /**
     * Size in byte
     */
    size: number;
    /**
     * Date the asset was last updated
     */
    updatedAt: string;
    /**
     * URL of the asset
     */
    url: string;
}

export interface GetRepositoryBranchesBranch {
    /**
     * Name of the branch.
     */
    name: string;
    /**
     * Whether the branch is protected.
     */
    protected: boolean;
}

export interface GetRepositoryDeployKeysKey {
    /**
     * Key id
     */
    id: number;
    /**
     * Key itself
     */
    key: string;
    /**
     * Key title
     */
    title: string;
    /**
     * `true` if the key was verified.
     */
    verified: boolean;
}

export interface GetRepositoryPage {
    cname: string;
    custom404: boolean;
    /**
     * URL to the repository on the web.
     */
    htmlUrl: string;
    sources: outputs.GetRepositoryPageSource[];
    status: string;
    url: string;
}

export interface GetRepositoryPageSource {
    branch: string;
    path: string;
}

export interface GetRepositoryPullRequestsResult {
    /**
     * If set, filters Pull Requests by base branch name.
     */
    baseRef: string;
    /**
     * Head commit SHA of the Pull Request base.
     */
    baseSha: string;
    /**
     * Body of the Pull Request.
     */
    body: string;
    /**
     * Indicates Whether this Pull Request is a draft.
     */
    draft: boolean;
    /**
     * Owner of the Pull Request head repository.
     */
    headOwner: string;
    /**
     * If set, filters Pull Requests by head user or head organization and branch name in the format of "user:ref-name" or "organization:ref-name". For example: "github:new-script-format" or "octocat:test-branch".
     */
    headRef: string;
    /**
     * Name of the Pull Request head repository.
     */
    headRepository: string;
    /**
     * Head commit SHA of the Pull Request head.
     */
    headSha: string;
    /**
     * List of label names set on the Pull Request.
     */
    labels: string[];
    /**
     * Indicates whether the base repository maintainers can modify the Pull Request.
     */
    maintainerCanModify: boolean;
    /**
     * The number of the Pull Request within the repository.
     */
    number: number;
    /**
     * Unix timestamp indicating the Pull Request creation time.
     */
    openedAt: number;
    /**
     * GitHub login of the user who opened the Pull Request.
     */
    openedBy: string;
    /**
     * If set, filters Pull Requests by state. Can be "open", "closed", or "all". Default: "open".
     */
    state: string;
    /**
     * The title of the Pull Request.
     */
    title: string;
    /**
     * The timestamp of the last Pull Request update.
     */
    updatedAt: number;
}

export interface GetRepositoryTeamsTeam {
    /**
     * The name of the repository.
     */
    name: string;
    /**
     * Team permission
     */
    permission: string;
    /**
     * Team slug
     */
    slug: string;
}

export interface GetRepositoryTemplate {
    owner: string;
    repository: string;
}

export interface GetRepositoryWebhooksWebhook {
    /**
     * `true` if the webhook is active.
     */
    active: boolean;
    /**
     * the ID of the webhook.
     */
    id: number;
    /**
     * the name of the webhook.
     */
    name: string;
    /**
     * the type of the webhook.
     */
    type: string;
    /**
     * the url of the webhook.
     */
    url: string;
}

export interface GetTreeEntry {
    mode: string;
    path: string;
    sha: string;
    size: number;
    type: string;
}

export interface OrganizationWebhookConfiguration {
    contentType?: string;
    insecureSsl?: boolean;
    secret?: string;
    /**
     * URL of the webhook
     */
    url: string;
}

export interface RepositoryEnvironmentDeploymentBranchPolicy {
    /**
     * Whether only branches that match the specified name patterns can deploy to this environment.
     */
    customBranchPolicies: boolean;
    /**
     * Whether only branches with branch protection rules can deploy to this environment.
     */
    protectedBranches: boolean;
}

export interface RepositoryEnvironmentReviewer {
    /**
     * Up to 6 IDs for teams who may review jobs that reference the environment. Reviewers must have at least read access to the repository. Only one of the required reviewers needs to approve the job for it to proceed.
     */
    teams?: number[];
    /**
     * Up to 6 IDs for users who may review jobs that reference the environment. Reviewers must have at least read access to the repository. Only one of the required reviewers needs to approve the job for it to proceed.
     */
    users?: number[];
}

export interface RepositoryPages {
    /**
     * The custom domain for the repository. This can only be set after the repository has been created.
     */
    cname?: string;
    /**
     * Whether the rendered GitHub Pages site has a custom 404 page.
     */
    custom404: boolean;
    /**
     * The absolute URL (including scheme) of the rendered GitHub Pages site e.g. `https://username.github.io`.
     */
    htmlUrl: string;
    /**
     * The source branch and directory for the rendered Pages site. See GitHub Pages Source below for details.
     */
    source: outputs.RepositoryPagesSource;
    /**
     * Set to `enabled` to enable advanced security features on the repository. Can be `enabled` or `disabled`.
     */
    status: string;
    url: string;
}

export interface RepositoryPagesSource {
    /**
     * The repository branch used to publish the site's source files. (i.e. `main` or `gh-pages`.
     */
    branch: string;
    /**
     * The repository directory from which the site publishes (Default: `/`).
     */
    path?: string;
}

export interface RepositorySecurityAndAnalysis {
    /**
     * The advanced security configuration for the repository. See Advanced Security Configuration below for details.
     */
    advancedSecurity: outputs.RepositorySecurityAndAnalysisAdvancedSecurity;
    /**
     * The secret scanning configuration for the repository. See Secret Scanning Configuration below for details.
     */
    secretScanning: outputs.RepositorySecurityAndAnalysisSecretScanning;
    /**
     * The secret scanning push protection configuration for the repository. See Secret Scanning Push Protection Configuration below for details.
     */
    secretScanningPushProtection: outputs.RepositorySecurityAndAnalysisSecretScanningPushProtection;
}

export interface RepositorySecurityAndAnalysisAdvancedSecurity {
    /**
     * Set to `enabled` to enable advanced security features on the repository. Can be `enabled` or `disabled`.
     */
    status: string;
}

export interface RepositorySecurityAndAnalysisSecretScanning {
    /**
     * Set to `enabled` to enable advanced security features on the repository. Can be `enabled` or `disabled`.
     */
    status: string;
}

export interface RepositorySecurityAndAnalysisSecretScanningPushProtection {
    /**
     * Set to `enabled` to enable advanced security features on the repository. Can be `enabled` or `disabled`.
     */
    status: string;
}

export interface RepositoryTemplate {
    /**
     * Whether the new repository should include all the branches from the template repository (defaults to false, which includes only the default branch from the template).
     */
    includeAllBranches?: boolean;
    /**
     * The GitHub organization or user the template repository is owned by.
     */
    owner: string;
    /**
     * The name of the template repository.
     */
    repository: string;
}

export interface RepositoryWebhookConfiguration {
    /**
     * The content type for the payload. Valid values are either `form` or `json`.
     */
    contentType?: string;
    /**
     * Insecure SSL boolean toggle. Defaults to `false`.
     */
    insecureSsl?: boolean;
    /**
     * The shared secret for the webhook. [See API documentation](https://developer.github.com/v3/repos/hooks/#create-a-hook).
     */
    secret?: string;
    /**
     * The URL of the webhook.
     */
    url: string;
}

export interface TeamMembersMember {
    /**
     * The role of the user within the team.
     * Must be one of `member` or `maintainer`. Defaults to `member`.
     */
    role?: string;
    /**
     * The user to add to the team.
     */
    username: string;
}

export interface TeamSettingsReviewRequestDelegation {
    /**
     * The algorithm to use when assigning pull requests to team members. Supported values are `ROUND_ROBIN` and `LOAD_BALANCE`. Default value is `ROUND_ROBIN`
     */
    algorithm?: string;
    /**
     * The number of team members to assign to a pull request
     */
    memberCount?: number;
    /**
     * whether to notify the entire team when at least one member is also assigned to the pull request
     */
    notify?: boolean;
}

export interface TeamSyncGroupMappingGroup {
    /**
     * The description of the IdP group.
     */
    groupDescription: string;
    /**
     * The ID of the IdP group.
     */
    groupId: string;
    /**
     * The name of the IdP group.
     */
    groupName: string;
}

export namespace config {
    export interface AppAuth {
        id: string;
        installationId: string;
        pemFile: string;
    }

}
