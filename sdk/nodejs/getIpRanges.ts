// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getIpRanges(opts?: pulumi.InvokeOptions): Promise<GetIpRangesResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("github:index/getIpRanges:getIpRanges", {
    }, opts);
}

/**
 * A collection of values returned by getIpRanges.
 */
export interface GetIpRangesResult {
    readonly actions: string[];
    readonly actionsIpv4s: string[];
    readonly actionsIpv6s: string[];
    readonly apiIpv4s: string[];
    readonly apiIpv6s: string[];
    readonly apis: string[];
    readonly dependabotIpv4s: string[];
    readonly dependabotIpv6s: string[];
    readonly dependabots: string[];
    readonly gitIpv4s: string[];
    readonly gitIpv6s: string[];
    readonly gits: string[];
    readonly hooks: string[];
    readonly hooksIpv4s: string[];
    readonly hooksIpv6s: string[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly importerIpv4s: string[];
    readonly importerIpv6s: string[];
    readonly importers: string[];
    readonly pages: string[];
    readonly pagesIpv4s: string[];
    readonly pagesIpv6s: string[];
    readonly webIpv4s: string[];
    readonly webIpv6s: string[];
    readonly webs: string[];
}
