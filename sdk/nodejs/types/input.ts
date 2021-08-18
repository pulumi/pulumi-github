// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";

export interface BranchProtectionRequiredPullRequestReview {
    dismissStaleReviews?: pulumi.Input<boolean>;
    dismissalRestrictions?: pulumi.Input<pulumi.Input<string>[]>;
    requireCodeOwnerReviews?: pulumi.Input<boolean>;
    requiredApprovingReviewCount?: pulumi.Input<number>;
    restrictDismissals?: pulumi.Input<boolean>;
}

export interface BranchProtectionRequiredStatusCheck {
    contexts?: pulumi.Input<pulumi.Input<string>[]>;
    strict?: pulumi.Input<boolean>;
}

export interface BranchProtectionV3RequiredPullRequestReviews {
    dismissStaleReviews?: pulumi.Input<boolean>;
    dismissalTeams?: pulumi.Input<pulumi.Input<string>[]>;
    dismissalUsers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * @deprecated Use enforce_admins instead
     */
    includeAdmins?: pulumi.Input<boolean>;
    requireCodeOwnerReviews?: pulumi.Input<boolean>;
    requiredApprovingReviewCount?: pulumi.Input<number>;
}

export interface BranchProtectionV3RequiredStatusChecks {
    contexts?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * @deprecated Use enforce_admins instead
     */
    includeAdmins?: pulumi.Input<boolean>;
    strict?: pulumi.Input<boolean>;
}

export interface BranchProtectionV3Restrictions {
    apps?: pulumi.Input<pulumi.Input<string>[]>;
    teams?: pulumi.Input<pulumi.Input<string>[]>;
    users?: pulumi.Input<pulumi.Input<string>[]>;
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
    source: pulumi.Input<inputs.RepositoryPagesSource>;
    /**
     * The GitHub Pages site's build status e.g. `building` or `built`.
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

export interface RepositoryTemplate {
    owner: pulumi.Input<string>;
    repository: pulumi.Input<string>;
}

export interface RepositoryWebhookConfiguration {
    contentType?: pulumi.Input<string>;
    insecureSsl?: pulumi.Input<boolean>;
    secret?: pulumi.Input<string>;
    /**
     * URL of the webhook.  This is a sensitive attribute because it may include basic auth credentials.
     */
    url: pulumi.Input<string>;
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
