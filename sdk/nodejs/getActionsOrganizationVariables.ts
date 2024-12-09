// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Use this data source to retrieve the list of variables of the organization.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as github from "@pulumi/github";
 *
 * const example = github.getActionsOrganizationVariables({});
 * ```
 */
export function getActionsOrganizationVariables(opts?: pulumi.InvokeOptions): Promise<GetActionsOrganizationVariablesResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("github:index/getActionsOrganizationVariables:getActionsOrganizationVariables", {
    }, opts);
}

/**
 * A collection of values returned by getActionsOrganizationVariables.
 */
export interface GetActionsOrganizationVariablesResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * list of variables for the repository
     */
    readonly variables: outputs.GetActionsOrganizationVariablesVariable[];
}
/**
 * Use this data source to retrieve the list of variables of the organization.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as github from "@pulumi/github";
 *
 * const example = github.getActionsOrganizationVariables({});
 * ```
 */
export function getActionsOrganizationVariablesOutput(opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetActionsOrganizationVariablesResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("github:index/getActionsOrganizationVariables:getActionsOrganizationVariables", {
    }, opts);
}
