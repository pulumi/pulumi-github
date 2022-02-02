// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * > **Note:** The data source will return a maximum of `1000` repositories
 * 	[as documented in official API docs](https://developer.github.com/v3/search/#about-the-search-api).
 *
 * Use this data source to retrieve a list of GitHub repositories using a search query.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as github from "@pulumi/github";
 *
 * const example = pulumi.output(github.getRepositories({
 *     query: "org:hashicorp language:Go",
 * }));
 * ```
 */
export function getRepositories(args: GetRepositoriesArgs, opts?: pulumi.InvokeOptions): Promise<GetRepositoriesResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("github:index/getRepositories:getRepositories", {
        "query": args.query,
        "sort": args.sort,
    }, opts);
}

/**
 * A collection of arguments for invoking getRepositories.
 */
export interface GetRepositoriesArgs {
    /**
     * Search query. See [documentation for the search syntax](https://help.github.com/articles/understanding-the-search-syntax/).
     */
    query: string;
    /**
     * Sorts the repositories returned by the specified attribute. Valid values include `stars`, `fork`, and `updated`. Defaults to `updated`.
     */
    sort?: string;
}

/**
 * A collection of values returned by getRepositories.
 */
export interface GetRepositoriesResult {
    readonly fullNames: string[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly names: string[];
    readonly query: string;
    readonly sort?: string;
}

export function getRepositoriesOutput(args: GetRepositoriesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRepositoriesResult> {
    return pulumi.output(args).apply(a => getRepositories(a, opts))
}

/**
 * A collection of arguments for invoking getRepositories.
 */
export interface GetRepositoriesOutputArgs {
    /**
     * Search query. See [documentation for the search syntax](https://help.github.com/articles/understanding-the-search-syntax/).
     */
    query: pulumi.Input<string>;
    /**
     * Sorts the repositories returned by the specified attribute. Valid values include `stars`, `fork`, and `updated`. Defaults to `updated`.
     */
    sort?: pulumi.Input<string>;
}
