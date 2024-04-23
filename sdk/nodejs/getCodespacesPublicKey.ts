// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use this data source to retrieve information about a GitHub Codespaces public key. This data source is required to be used with other GitHub secrets interactions.
 * Note that the provider `token` must have admin rights to a repository to retrieve it's Codespaces public key.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as github from "@pulumi/github";
 *
 * const example = github.getCodespacesPublicKey({
 *     repository: "example_repo",
 * });
 * ```
 */
export function getCodespacesPublicKey(args: GetCodespacesPublicKeyArgs, opts?: pulumi.InvokeOptions): Promise<GetCodespacesPublicKeyResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("github:index/getCodespacesPublicKey:getCodespacesPublicKey", {
        "repository": args.repository,
    }, opts);
}

/**
 * A collection of arguments for invoking getCodespacesPublicKey.
 */
export interface GetCodespacesPublicKeyArgs {
    /**
     * Name of the repository to get public key from.
     */
    repository: string;
}

/**
 * A collection of values returned by getCodespacesPublicKey.
 */
export interface GetCodespacesPublicKeyResult {
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
    readonly repository: string;
}
/**
 * Use this data source to retrieve information about a GitHub Codespaces public key. This data source is required to be used with other GitHub secrets interactions.
 * Note that the provider `token` must have admin rights to a repository to retrieve it's Codespaces public key.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as github from "@pulumi/github";
 *
 * const example = github.getCodespacesPublicKey({
 *     repository: "example_repo",
 * });
 * ```
 */
export function getCodespacesPublicKeyOutput(args: GetCodespacesPublicKeyOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetCodespacesPublicKeyResult> {
    return pulumi.output(args).apply((a: any) => getCodespacesPublicKey(a, opts))
}

/**
 * A collection of arguments for invoking getCodespacesPublicKey.
 */
export interface GetCodespacesPublicKeyOutputArgs {
    /**
     * Name of the repository to get public key from.
     */
    repository: pulumi.Input<string>;
}
