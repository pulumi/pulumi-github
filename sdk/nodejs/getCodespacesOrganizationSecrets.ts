// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Use this data source to retrieve the list of codespaces secrets of the organization.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as github from "@pulumi/github";
 *
 * const example = github.getCodespacesOrganizationSecrets({});
 * ```
 */
export function getCodespacesOrganizationSecrets(opts?: pulumi.InvokeOptions): Promise<GetCodespacesOrganizationSecretsResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("github:index/getCodespacesOrganizationSecrets:getCodespacesOrganizationSecrets", {
    }, opts);
}

/**
 * A collection of values returned by getCodespacesOrganizationSecrets.
 */
export interface GetCodespacesOrganizationSecretsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * list of secrets for the repository
     */
    readonly secrets: outputs.GetCodespacesOrganizationSecretsSecret[];
}
/**
 * Use this data source to retrieve the list of codespaces secrets of the organization.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as github from "@pulumi/github";
 *
 * const example = github.getCodespacesOrganizationSecrets({});
 * ```
 */
export function getCodespacesOrganizationSecretsOutput(opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetCodespacesOrganizationSecretsResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("github:index/getCodespacesOrganizationSecrets:getCodespacesOrganizationSecrets", {
    }, opts);
}
