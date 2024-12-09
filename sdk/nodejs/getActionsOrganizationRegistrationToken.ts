// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use this data source to retrieve a GitHub Actions organization registration token. This token can then be used to register a self-hosted runner.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as github from "@pulumi/github";
 *
 * const example = github.getActionsOrganizationRegistrationToken({});
 * ```
 */
export function getActionsOrganizationRegistrationToken(opts?: pulumi.InvokeOptions): Promise<GetActionsOrganizationRegistrationTokenResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("github:index/getActionsOrganizationRegistrationToken:getActionsOrganizationRegistrationToken", {
    }, opts);
}

/**
 * A collection of values returned by getActionsOrganizationRegistrationToken.
 */
export interface GetActionsOrganizationRegistrationTokenResult {
    /**
     * The token expiration date.
     */
    readonly expiresAt: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The token that has been retrieved.
     */
    readonly token: string;
}
/**
 * Use this data source to retrieve a GitHub Actions organization registration token. This token can then be used to register a self-hosted runner.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as github from "@pulumi/github";
 *
 * const example = github.getActionsOrganizationRegistrationToken({});
 * ```
 */
export function getActionsOrganizationRegistrationTokenOutput(opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetActionsOrganizationRegistrationTokenResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("github:index/getActionsOrganizationRegistrationToken:getActionsOrganizationRegistrationToken", {
    }, opts);
}
