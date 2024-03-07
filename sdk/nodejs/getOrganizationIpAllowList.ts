// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Use this data source to retrieve information about the IP allow list of an organization.
 * The allow list for IP addresses will block access to private resources via the web, API,
 * and Git from any IP addresses that are not on the allow list.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as github from "@pulumi/github";
 *
 * const all = github.getOrganizationIpAllowList({});
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getOrganizationIpAllowList(opts?: pulumi.InvokeOptions): Promise<GetOrganizationIpAllowListResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("github:index/getOrganizationIpAllowList:getOrganizationIpAllowList", {
    }, opts);
}

/**
 * A collection of values returned by getOrganizationIpAllowList.
 */
export interface GetOrganizationIpAllowListResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * An Array of allowed IP addresses.
     * ___
     */
    readonly ipAllowLists: outputs.GetOrganizationIpAllowListIpAllowList[];
}
/**
 * Use this data source to retrieve information about the IP allow list of an organization.
 * The allow list for IP addresses will block access to private resources via the web, API,
 * and Git from any IP addresses that are not on the allow list.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as github from "@pulumi/github";
 *
 * const all = github.getOrganizationIpAllowList({});
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getOrganizationIpAllowListOutput(opts?: pulumi.InvokeOptions): pulumi.Output<GetOrganizationIpAllowListResult> {
    return pulumi.output(getOrganizationIpAllowList(opts))
}
