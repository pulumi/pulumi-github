// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use this data source to retrieve information about a repository ref.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as github from "@pulumi/github";
 *
 * const development = github.getRef({
 *     owner: "example",
 *     ref: "heads/development",
 *     repository: "example",
 * });
 * ```
 */
export function getRef(args: GetRefArgs, opts?: pulumi.InvokeOptions): Promise<GetRefResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("github:index/getRef:getRef", {
        "owner": args.owner,
        "ref": args.ref,
        "repository": args.repository,
    }, opts);
}

/**
 * A collection of arguments for invoking getRef.
 */
export interface GetRefArgs {
    /**
     * Owner of the repository.
     */
    owner?: string;
    /**
     * The repository ref to look up. Must be formatted `heads/<ref>` for branches, and `tags/<ref>` for tags.
     */
    ref: string;
    /**
     * The GitHub repository name.
     */
    repository: string;
}

/**
 * A collection of values returned by getRef.
 */
export interface GetRefResult {
    /**
     * An etag representing the ref.
     */
    readonly etag: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly owner?: string;
    readonly ref: string;
    readonly repository: string;
    /**
     * A string storing the reference's `HEAD` commit's SHA1.
     */
    readonly sha: string;
}
/**
 * Use this data source to retrieve information about a repository ref.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as github from "@pulumi/github";
 *
 * const development = github.getRef({
 *     owner: "example",
 *     ref: "heads/development",
 *     repository: "example",
 * });
 * ```
 */
export function getRefOutput(args: GetRefOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRefResult> {
    return pulumi.output(args).apply((a: any) => getRef(a, opts))
}

/**
 * A collection of arguments for invoking getRef.
 */
export interface GetRefOutputArgs {
    /**
     * Owner of the repository.
     */
    owner?: pulumi.Input<string>;
    /**
     * The repository ref to look up. Must be formatted `heads/<ref>` for branches, and `tags/<ref>` for tags.
     */
    ref: pulumi.Input<string>;
    /**
     * The GitHub repository name.
     */
    repository: pulumi.Input<string>;
}
