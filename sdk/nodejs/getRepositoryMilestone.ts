// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use this data source to retrieve information about a specific GitHub milestone in a repository.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as github from "@pulumi/github";
 *
 * const example = github.getRepositoryMilestone({
 *     owner: "example-owner",
 *     repository: "example-repository",
 *     number: 1,
 * });
 * ```
 */
export function getRepositoryMilestone(args: GetRepositoryMilestoneArgs, opts?: pulumi.InvokeOptions): Promise<GetRepositoryMilestoneResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("github:index/getRepositoryMilestone:getRepositoryMilestone", {
        "number": args.number,
        "owner": args.owner,
        "repository": args.repository,
    }, opts);
}

/**
 * A collection of arguments for invoking getRepositoryMilestone.
 */
export interface GetRepositoryMilestoneArgs {
    /**
     * The number of the milestone.
     */
    number: number;
    /**
     * Owner of the repository.
     */
    owner: string;
    /**
     * Name of the repository to retrieve the milestone from.
     */
    repository: string;
}

/**
 * A collection of values returned by getRepositoryMilestone.
 */
export interface GetRepositoryMilestoneResult {
    /**
     * Description of the milestone.
     */
    readonly description: string;
    /**
     * The milestone due date (in ISO-8601 `yyyy-mm-dd` format).
     */
    readonly dueDate: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly number: number;
    readonly owner: string;
    readonly repository: string;
    /**
     * State of the milestone.
     */
    readonly state: string;
    /**
     * Title of the milestone.
     */
    readonly title: string;
}
/**
 * Use this data source to retrieve information about a specific GitHub milestone in a repository.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as github from "@pulumi/github";
 *
 * const example = github.getRepositoryMilestone({
 *     owner: "example-owner",
 *     repository: "example-repository",
 *     number: 1,
 * });
 * ```
 */
export function getRepositoryMilestoneOutput(args: GetRepositoryMilestoneOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetRepositoryMilestoneResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("github:index/getRepositoryMilestone:getRepositoryMilestone", {
        "number": args.number,
        "owner": args.owner,
        "repository": args.repository,
    }, opts);
}

/**
 * A collection of arguments for invoking getRepositoryMilestone.
 */
export interface GetRepositoryMilestoneOutputArgs {
    /**
     * The number of the milestone.
     */
    number: pulumi.Input<number>;
    /**
     * Owner of the repository.
     */
    owner: pulumi.Input<string>;
    /**
     * Name of the repository to retrieve the milestone from.
     */
    repository: pulumi.Input<string>;
}
