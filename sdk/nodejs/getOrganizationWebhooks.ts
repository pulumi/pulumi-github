// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getOrganizationWebhooks(opts?: pulumi.InvokeOptions): Promise<GetOrganizationWebhooksResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("github:index/getOrganizationWebhooks:getOrganizationWebhooks", {
    }, opts);
}

/**
 * A collection of values returned by getOrganizationWebhooks.
 */
export interface GetOrganizationWebhooksResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly webhooks: outputs.GetOrganizationWebhooksWebhook[];
}
