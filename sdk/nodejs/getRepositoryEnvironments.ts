// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Use this data source to retrieve information about environments for a repository.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as github from "@pulumi/github";
 *
 * const example = github.getRepositoryEnvironments({
 *     repository: "example-repository",
 * });
 * ```
 */
export function getRepositoryEnvironments(args: GetRepositoryEnvironmentsArgs, opts?: pulumi.InvokeOptions): Promise<GetRepositoryEnvironmentsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("github:index/getRepositoryEnvironments:getRepositoryEnvironments", {
        "repository": args.repository,
    }, opts);
}

/**
 * A collection of arguments for invoking getRepositoryEnvironments.
 */
export interface GetRepositoryEnvironmentsArgs {
    /**
     * Name of the repository to retrieve the environments from.
     */
    repository: string;
}

/**
 * A collection of values returned by getRepositoryEnvironments.
 */
export interface GetRepositoryEnvironmentsResult {
    /**
     * The list of this repository's environments. Each element of `environments` has the following attributes:
     */
    readonly environments: outputs.GetRepositoryEnvironmentsEnvironment[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly repository: string;
}
/**
 * Use this data source to retrieve information about environments for a repository.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as github from "@pulumi/github";
 *
 * const example = github.getRepositoryEnvironments({
 *     repository: "example-repository",
 * });
 * ```
 */
export function getRepositoryEnvironmentsOutput(args: GetRepositoryEnvironmentsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRepositoryEnvironmentsResult> {
    return pulumi.output(args).apply((a: any) => getRepositoryEnvironments(a, opts))
}

/**
 * A collection of arguments for invoking getRepositoryEnvironments.
 */
export interface GetRepositoryEnvironmentsOutputArgs {
    /**
     * Name of the repository to retrieve the environments from.
     */
    repository: pulumi.Input<string>;
}