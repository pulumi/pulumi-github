// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use this data source to retrieve basic information about a GitHub Organization.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as github from "@pulumi/github";
 *
 * const example = github.getOrganization({
 *     name: "github",
 * });
 * ```
 */
export function getOrganization(args: GetOrganizationArgs, opts?: pulumi.InvokeOptions): Promise<GetOrganizationResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("github:index/getOrganization:getOrganization", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getOrganization.
 */
export interface GetOrganizationArgs {
    /**
     * The organization's public profile name
     */
    name: string;
}

/**
 * A collection of values returned by getOrganization.
 */
export interface GetOrganizationResult {
    /**
     * Whether advanced security is enabled for new repositories.
     */
    readonly advancedSecurityEnabledForNewRepositories: boolean;
    /**
     * Default permission level members have for organization repositories.
     */
    readonly defaultRepositoryPermission: string;
    /**
     * Whether Dependabot alerts is automatically enabled for new repositories.
     */
    readonly dependabotAlertsEnabledForNewRepositories: boolean;
    /**
     * Whether Dependabot security updates is automatically enabled for new repositories.
     */
    readonly dependabotSecurityUpdatesEnabledForNewRepositories: boolean;
    /**
     * Whether dependency graph is automatically enabled for new repositories.
     */
    readonly dependencyGraphEnabledForNewRepositories: boolean;
    /**
     * The organization account description
     */
    readonly description: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The members login
     */
    readonly login: string;
    /**
     * **Deprecated**: use `users` instead by replacing `github_organization.example.members` to `github_organization.example.users[*].login` which will give you the same value, expect this field to be removed in next major version
     *
     * @deprecated Use `users` instead by replacing `github_organization.example.members` to `github_organization.example.users[*].login`. Expect this field to be removed in next major version.
     */
    readonly members: string[];
    /**
     * The type of repository allowed to be created by members of the organization. Can be one of `ALL`, `PUBLIC`, `PRIVATE`, `NONE`.
     */
    readonly membersAllowedRepositoryCreationType: string;
    /**
     * Whether organization members can create internal repositories.
     */
    readonly membersCanCreateInternalRepositories: boolean;
    /**
     * Whether organization members can create pages sites.
     */
    readonly membersCanCreatePages: boolean;
    /**
     * Whether organization members can create private pages sites.
     */
    readonly membersCanCreatePrivatePages: boolean;
    /**
     * Whether organization members can create private repositories.
     */
    readonly membersCanCreatePrivateRepositories: boolean;
    /**
     * Whether organization members can create public pages sites.
     */
    readonly membersCanCreatePublicPages: boolean;
    /**
     * Whether organization members can create public repositories.
     */
    readonly membersCanCreatePublicRepositories: boolean;
    /**
     * Whether non-admin organization members can create repositories.
     */
    readonly membersCanCreateRepositories: boolean;
    /**
     * Whether organization members can create private repository forks.
     */
    readonly membersCanForkPrivateRepositories: boolean;
    /**
     * The organization's public profile name
     */
    readonly name: string;
    /**
     * GraphQL global node ID for use with the v4 API
     */
    readonly nodeId: string;
    /**
     * The organization's name as used in URLs and the API
     */
    readonly orgname: string;
    /**
     * The organization account plan name
     */
    readonly plan: string;
    /**
     * (`list`) A list of the full names of the repositories in the organization formatted as `owner/name` strings
     */
    readonly repositories: string[];
    /**
     * Whether secret scanning is automatically enabled for new repositories.
     */
    readonly secretScanningEnabledForNewRepositories: boolean;
    /**
     * Whether secret scanning push protection is automatically enabled for new repositories.
     */
    readonly secretScanningPushProtectionEnabledForNewRepositories: boolean;
    /**
     * Whether two-factor authentication is required for all members of the organization.
     */
    readonly twoFactorRequirementEnabled: boolean;
    /**
     * (`list`) A list with the members of the organization with following fields:
     */
    readonly users: {[key: string]: string}[];
    /**
     * Whether organization members must sign all commits.
     */
    readonly webCommitSignoffRequired: boolean;
}
/**
 * Use this data source to retrieve basic information about a GitHub Organization.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as github from "@pulumi/github";
 *
 * const example = github.getOrganization({
 *     name: "github",
 * });
 * ```
 */
export function getOrganizationOutput(args: GetOrganizationOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetOrganizationResult> {
    return pulumi.output(args).apply((a: any) => getOrganization(a, opts))
}

/**
 * A collection of arguments for invoking getOrganization.
 */
export interface GetOrganizationOutputArgs {
    /**
     * The organization's public profile name
     */
    name: pulumi.Input<string>;
}
