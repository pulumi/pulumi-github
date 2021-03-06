// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Use this data source to retrieve information about multiple GitHub Pull Requests in a repository.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as github from "@pulumi/github";
 *
 * const example = pulumi.output(github.getRepositoryPullRequests({
 *     baseRef: "main",
 *     baseRepository: "example-repository",
 *     sortBy: "updated",
 *     sortDirection: "desc",
 *     state: "open",
 * }, { async: true }));
 * ```
 */
export function getRepositoryPullRequests(args: GetRepositoryPullRequestsArgs, opts?: pulumi.InvokeOptions): Promise<GetRepositoryPullRequestsResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("github:index/getRepositoryPullRequests:getRepositoryPullRequests", {
        "baseRef": args.baseRef,
        "baseRepository": args.baseRepository,
        "headRef": args.headRef,
        "owner": args.owner,
        "sortBy": args.sortBy,
        "sortDirection": args.sortDirection,
        "state": args.state,
    }, opts);
}

/**
 * A collection of arguments for invoking getRepositoryPullRequests.
 */
export interface GetRepositoryPullRequestsArgs {
    /**
     * If set, filters Pull Requests by base branch name.
     */
    readonly baseRef?: string;
    /**
     * Name of the base repository to retrieve the Pull Requests from.
     */
    readonly baseRepository: string;
    /**
     * If set, filters Pull Requests by head user or head organization and branch name in the format of "user:ref-name" or "organization:ref-name". For example: "github:new-script-format" or "octocat:test-branch".
     */
    readonly headRef?: string;
    /**
     * Owner of the repository. If not provided, the provider's default owner is used.
     */
    readonly owner?: string;
    /**
     * If set, indicates what to sort results by. Can be either "created", "updated", "popularity" (comment count) or "long-running" (age, filtering by pulls updated in the last month). Default: "created".
     */
    readonly sortBy?: string;
    /**
     * If set, controls the direction of the sort. Can be either "asc" or "desc". Default: "asc".
     */
    readonly sortDirection?: string;
    /**
     * If set, filters Pull Requests by state. Can be "open", "closed", or "all". Default: "open".
     */
    readonly state?: string;
}

/**
 * A collection of values returned by getRepositoryPullRequests.
 */
export interface GetRepositoryPullRequestsResult {
    /**
     * Name of the ref (branch) of the Pull Request base.
     */
    readonly baseRef?: string;
    readonly baseRepository: string;
    readonly headRef?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly owner?: string;
    /**
     * Collection of Pull Requests matching the filters. Each of the results conforms to the following scheme:
     */
    readonly results: outputs.GetRepositoryPullRequestsResult[];
    readonly sortBy?: string;
    readonly sortDirection?: string;
    /**
     * the current Pull Request state - can be "open", "closed" or "merged".
     */
    readonly state?: string;
}
