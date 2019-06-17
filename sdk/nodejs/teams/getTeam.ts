// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to retrieve information about a GitHub team.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as github from "@pulumi/github";
 * 
 * const example = pulumi.output(github.teams.getTeam({
 *     slug: "example",
 * }));
 * ```
 */
export function getTeam(args: GetTeamArgs, opts?: pulumi.InvokeOptions): Promise<GetTeamResult> {
    return pulumi.runtime.invoke("github:teams/getTeam:getTeam", {
        "slug": args.slug,
    }, opts);
}

/**
 * A collection of arguments for invoking getTeam.
 */
export interface GetTeamArgs {
    /**
     * The team slug.
     */
    readonly slug: string;
}

/**
 * A collection of values returned by getTeam.
 */
export interface GetTeamResult {
    /**
     * the team's description.
     */
    readonly description: string;
    /**
     * List of team members
     */
    readonly members: string[];
    /**
     * the team's full name.
     */
    readonly name: string;
    /**
     * the team's permission level.
     */
    readonly permission: string;
    /**
     * the team's privacy type.
     */
    readonly privacy: string;
    readonly slug: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
