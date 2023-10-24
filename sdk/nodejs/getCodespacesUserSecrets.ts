// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Use this data source to retrieve the list of codespaces secrets of the user.
 */
export function getCodespacesUserSecrets(opts?: pulumi.InvokeOptions): Promise<GetCodespacesUserSecretsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("github:index/getCodespacesUserSecrets:getCodespacesUserSecrets", {
    }, opts);
}

/**
 * A collection of values returned by getCodespacesUserSecrets.
 */
export interface GetCodespacesUserSecretsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * list of secrets for the repository
     */
    readonly secrets: outputs.GetCodespacesUserSecretsSecret[];
}
/**
 * Use this data source to retrieve the list of codespaces secrets of the user.
 */
export function getCodespacesUserSecretsOutput(opts?: pulumi.InvokeOptions): pulumi.Output<GetCodespacesUserSecretsResult> {
    return pulumi.output(getCodespacesUserSecrets(opts))
}
