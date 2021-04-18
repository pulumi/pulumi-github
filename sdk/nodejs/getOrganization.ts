// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
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
 * const test = pulumi.output(github.getOrganization({
 *     name: "github",
 * }, { async: true }));
 * ```
 */
export function getOrganization(args: GetOrganizationArgs, opts?: pulumi.InvokeOptions): Promise<GetOrganizationResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("github:index/getOrganization:getOrganization", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getOrganization.
 */
export interface GetOrganizationArgs {
    readonly name: string;
}

/**
 * A collection of values returned by getOrganization.
 */
export interface GetOrganizationResult {
    readonly description: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly login: string;
    readonly name: string;
    readonly nodeId: string;
    /**
     * The plan name for the organization account
     */
    readonly plan: string;
    /**
     * (`list`) A list with the repositories on the organization
     */
    readonly repositories: string[];
}
