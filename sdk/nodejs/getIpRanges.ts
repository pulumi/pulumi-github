// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use this data source to retrieve information about GitHub's IP addresses.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as github from "@pulumi/github";
 *
 * const test = github.getIpRanges({});
 * ```
 */
export function getIpRanges(opts?: pulumi.InvokeOptions): Promise<GetIpRangesResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("github:index/getIpRanges:getIpRanges", {
    }, opts);
}

/**
 * A collection of values returned by getIpRanges.
 */
export interface GetIpRangesResult {
    /**
     * An array of IP addresses in CIDR format specifying the addresses that incoming requests from GitHub actions will originate from.
     */
    readonly actions: string[];
    /**
     * A subset of the `actions` array that contains IP addresses in IPv4 CIDR format.
     */
    readonly actionsIpv4s: string[];
    /**
     * A subset of the `actions` array that contains IP addresses in IPv6 CIDR format.
     */
    readonly actionsIpv6s: string[];
    /**
     * A subset of the `api` array that contains IP addresses in IPv4 CIDR format.
     */
    readonly apiIpv4s: string[];
    /**
     * A subset of the `api` array that contains IP addresses in IPv6 CIDR format.
     */
    readonly apiIpv6s: string[];
    /**
     * An Array of IP addresses in CIDR format for the GitHub API.
     */
    readonly apis: string[];
    /**
     * A subset of the `dependabot` array that contains IP addresses in IPv4 CIDR format.
     */
    readonly dependabotIpv4s: string[];
    /**
     * A subset of the `dependabot` array that contains IP addresses in IPv6 CIDR format.
     */
    readonly dependabotIpv6s: string[];
    /**
     * An array of IP addresses in CIDR format specifying the A records for dependabot.
     */
    readonly dependabots: string[];
    /**
     * A subset of the `git` array that contains IP addresses in IPv4 CIDR format.
     */
    readonly gitIpv4s: string[];
    /**
     * A subset of the `git` array that contains IP addresses in IPv6 CIDR format.
     */
    readonly gitIpv6s: string[];
    /**
     * An Array of IP addresses in CIDR format specifying the Git servers.
     */
    readonly gits: string[];
    /**
     * An Array of IP addresses in CIDR format specifying the addresses that incoming service hooks will originate from.
     */
    readonly hooks: string[];
    /**
     * A subset of the `hooks` array that contains IP addresses in IPv4 CIDR format.
     */
    readonly hooksIpv4s: string[];
    /**
     * A subset of the `hooks` array that contains IP addresses in IPv6 CIDR format.
     */
    readonly hooksIpv6s: string[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A subset of the `importer` array that contains IP addresses in IPv4 CIDR format.
     */
    readonly importerIpv4s: string[];
    /**
     * A subset of the `importer` array that contains IP addresses in IPv6 CIDR format.
     */
    readonly importerIpv6s: string[];
    /**
     * An Array of IP addresses in CIDR format specifying the A records for GitHub Importer.
     */
    readonly importers: string[];
    /**
     * An Array of IP addresses in CIDR format specifying the A records for GitHub Pages.
     */
    readonly pages: string[];
    /**
     * A subset of the `pages` array that contains IP addresses in IPv4 CIDR format.
     */
    readonly pagesIpv4s: string[];
    /**
     * A subset of the `pages` array that contains IP addresses in IPv6 CIDR format.
     */
    readonly pagesIpv6s: string[];
    /**
     * A subset of the `web` array that contains IP addresses in IPv4 CIDR format.
     */
    readonly webIpv4s: string[];
    /**
     * A subset of the `web` array that contains IP addresses in IPv6 CIDR format.
     */
    readonly webIpv6s: string[];
    /**
     * An Array of IP addresses in CIDR format for GitHub Web.
     */
    readonly webs: string[];
}
/**
 * Use this data source to retrieve information about GitHub's IP addresses.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as github from "@pulumi/github";
 *
 * const test = github.getIpRanges({});
 * ```
 */
export function getIpRangesOutput(opts?: pulumi.InvokeOptions): pulumi.Output<GetIpRangesResult> {
    return pulumi.output(getIpRanges(opts))
}
