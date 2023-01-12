// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Use this data source to retrieve the list of dependabot secrets for a GitHub repository.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as github from "@pulumi/github";
 *
 * const example = github.getDependabotSecrets({
 *     name: "example",
 * });
 * ```
 */
export function getDependabotSecrets(args?: GetDependabotSecretsArgs, opts?: pulumi.InvokeOptions): Promise<GetDependabotSecretsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("github:index/getDependabotSecrets:getDependabotSecrets", {
        "fullName": args.fullName,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getDependabotSecrets.
 */
export interface GetDependabotSecretsArgs {
    /**
     * Full name of the repository (in `org/name` format).
     */
    fullName?: string;
    /**
     * The name of the repository.
     */
    name?: string;
}

/**
 * A collection of values returned by getDependabotSecrets.
 */
export interface GetDependabotSecretsResult {
    readonly fullName: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Secret name
     */
    readonly name: string;
    /**
     * list of dependabot secrets for the repository
     */
    readonly secrets: outputs.GetDependabotSecretsSecret[];
}
/**
 * Use this data source to retrieve the list of dependabot secrets for a GitHub repository.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as github from "@pulumi/github";
 *
 * const example = github.getDependabotSecrets({
 *     name: "example",
 * });
 * ```
 */
export function getDependabotSecretsOutput(args?: GetDependabotSecretsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDependabotSecretsResult> {
    return pulumi.output(args).apply((a: any) => getDependabotSecrets(a, opts))
}

/**
 * A collection of arguments for invoking getDependabotSecrets.
 */
export interface GetDependabotSecretsOutputArgs {
    /**
     * Full name of the repository (in `org/name` format).
     */
    fullName?: pulumi.Input<string>;
    /**
     * The name of the repository.
     */
    name?: pulumi.Input<string>;
}
