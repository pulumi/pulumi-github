// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Use this data source to retrieve the list of secrets of the repository environment.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as github from "@pulumi/github";
 *
 * const example = github.getActionsEnvironmentSecrets({
 *     name: "exampleRepo",
 *     environment: "exampleEnvironment",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getActionsEnvironmentSecrets(args: GetActionsEnvironmentSecretsArgs, opts?: pulumi.InvokeOptions): Promise<GetActionsEnvironmentSecretsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("github:index/getActionsEnvironmentSecrets:getActionsEnvironmentSecrets", {
        "environment": args.environment,
        "fullName": args.fullName,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getActionsEnvironmentSecrets.
 */
export interface GetActionsEnvironmentSecretsArgs {
    environment: string;
    fullName?: string;
    /**
     * Name of the secret
     */
    name?: string;
}

/**
 * A collection of values returned by getActionsEnvironmentSecrets.
 */
export interface GetActionsEnvironmentSecretsResult {
    readonly environment: string;
    readonly fullName: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Name of the secret
     */
    readonly name: string;
    /**
     * list of secrets for the environment
     */
    readonly secrets: outputs.GetActionsEnvironmentSecretsSecret[];
}
/**
 * Use this data source to retrieve the list of secrets of the repository environment.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as github from "@pulumi/github";
 *
 * const example = github.getActionsEnvironmentSecrets({
 *     name: "exampleRepo",
 *     environment: "exampleEnvironment",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getActionsEnvironmentSecretsOutput(args: GetActionsEnvironmentSecretsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetActionsEnvironmentSecretsResult> {
    return pulumi.output(args).apply((a: any) => getActionsEnvironmentSecrets(a, opts))
}

/**
 * A collection of arguments for invoking getActionsEnvironmentSecrets.
 */
export interface GetActionsEnvironmentSecretsOutputArgs {
    environment: pulumi.Input<string>;
    fullName?: pulumi.Input<string>;
    /**
     * Name of the secret
     */
    name?: pulumi.Input<string>;
}
