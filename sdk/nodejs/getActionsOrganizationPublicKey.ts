// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use this data source to retrieve information about a GitHub Actions Organization public key. This data source is required to be used with other GitHub secrets interactions.
 * Note that the provider `token` must have admin rights to an organization to retrieve it's action public key.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as github from "@pulumi/github";
 *
 * const example = github.getActionsOrganizationPublicKey({});
 * ```
 */
export function getActionsOrganizationPublicKey(opts?: pulumi.InvokeOptions): Promise<GetActionsOrganizationPublicKeyResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("github:index/getActionsOrganizationPublicKey:getActionsOrganizationPublicKey", {
    }, opts);
}

/**
 * A collection of values returned by getActionsOrganizationPublicKey.
 */
export interface GetActionsOrganizationPublicKeyResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Actual key retrieved.
     */
    readonly key: string;
    /**
     * ID of the key that has been retrieved.
     */
    readonly keyId: string;
}
