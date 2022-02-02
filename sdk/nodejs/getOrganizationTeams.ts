// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Use this data source to retrieve information about all GitHub teams in an organization.
 *
 * ## Example Usage
 *
 * To retrieve *all* teams of the organization:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as github from "@pulumi/github";
 *
 * const all = pulumi.output(github.getOrganizationTeams());
 * ```
 *
 * To retrieve only the team's at the root of the organization:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as github from "@pulumi/github";
 *
 * const rootTeams = pulumi.output(github.getOrganizationTeams({
 *     rootTeamsOnly: true,
 * }));
 * ```
 */
export function getOrganizationTeams(args?: GetOrganizationTeamsArgs, opts?: pulumi.InvokeOptions): Promise<GetOrganizationTeamsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("github:index/getOrganizationTeams:getOrganizationTeams", {
        "rootTeamsOnly": args.rootTeamsOnly,
    }, opts);
}

/**
 * A collection of arguments for invoking getOrganizationTeams.
 */
export interface GetOrganizationTeamsArgs {
    /**
     * Only return teams that are at the organization's root, i.e. no nested teams. Defaults to `false`.
     */
    rootTeamsOnly?: boolean;
}

/**
 * A collection of values returned by getOrganizationTeams.
 */
export interface GetOrganizationTeamsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Only return teams that are at the organization's root, i.e. no nested teams. Defaults to `false`.
     */
    readonly rootTeamsOnly?: boolean;
    /**
     * An Array of GitHub Teams.  Each `team` block consists of the fields documented below.
     * ___
     */
    readonly teams: outputs.GetOrganizationTeamsTeam[];
}

export function getOrganizationTeamsOutput(args?: GetOrganizationTeamsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetOrganizationTeamsResult> {
    return pulumi.output(args).apply(a => getOrganizationTeams(a, opts))
}

/**
 * A collection of arguments for invoking getOrganizationTeams.
 */
export interface GetOrganizationTeamsOutputArgs {
    /**
     * Only return teams that are at the organization's root, i.e. no nested teams. Defaults to `false`.
     */
    rootTeamsOnly?: pulumi.Input<boolean>;
}
