// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Use this data source to retrieve information about a GitHub repository.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as github from "@pulumi/github";
 *
 * const example = pulumi.output(github.getRepository({
 *     fullName: "hashicorp/terraform",
 * }));
 * ```
 */
export function getRepository(args?: GetRepositoryArgs, opts?: pulumi.InvokeOptions): Promise<GetRepositoryResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("github:index/getRepository:getRepository", {
        "description": args.description,
        "fullName": args.fullName,
        "homepageUrl": args.homepageUrl,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getRepository.
 */
export interface GetRepositoryArgs {
    /**
     * A description of the repository.
     */
    description?: string;
    /**
     * Full name of the repository (in `org/name` format).
     */
    fullName?: string;
    /**
     * URL of a page describing the project.
     */
    homepageUrl?: string;
    /**
     * The name of the repository.
     */
    name?: string;
}

/**
 * A collection of values returned by getRepository.
 */
export interface GetRepositoryResult {
    /**
     * Whether the repository allows auto-merging pull requests.
     */
    readonly allowAutoMerge: boolean;
    /**
     * Whether the repository allows merge commits.
     */
    readonly allowMergeCommit: boolean;
    /**
     * Whether the repository allows rebase merges.
     */
    readonly allowRebaseMerge: boolean;
    /**
     * Whether the repository allows squash merges.
     */
    readonly allowSquashMerge: boolean;
    /**
     * Whether the repository is archived.
     */
    readonly archived: boolean;
    /**
     * The list of this repository's branches. Each element of `branches` has the following attributes:
     */
    readonly branches: outputs.GetRepositoryBranch[];
    /**
     * The name of the default branch of the repository.
     */
    readonly defaultBranch: string;
    /**
     * A description of the repository.
     */
    readonly description?: string;
    readonly fullName: string;
    /**
     * URL that can be provided to `git clone` to clone the repository anonymously via the git protocol.
     */
    readonly gitCloneUrl: string;
    /**
     * Whether the repository has Downloads feature enabled.
     */
    readonly hasDownloads: boolean;
    /**
     * Whether the repository has GitHub Issues enabled.
     */
    readonly hasIssues: boolean;
    /**
     * Whether the repository has the GitHub Projects enabled.
     */
    readonly hasProjects: boolean;
    /**
     * Whether the repository has the GitHub Wiki enabled.
     */
    readonly hasWiki: boolean;
    /**
     * URL of a page describing the project.
     */
    readonly homepageUrl?: string;
    /**
     * URL to the repository on the web.
     */
    readonly htmlUrl: string;
    /**
     * URL that can be provided to `git clone` to clone the repository via HTTPS.
     */
    readonly httpCloneUrl: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Name of the branch.
     */
    readonly name: string;
    /**
     * GraphQL global node id for use with v4 API
     */
    readonly nodeId: string;
    /**
     * The repository's GitHub Pages configuration.
     */
    readonly pages: outputs.GetRepositoryPage[];
    /**
     * Whether the repository is private.
     */
    readonly private: boolean;
    /**
     * GitHub ID for the repository
     */
    readonly repoId: number;
    /**
     * URL that can be provided to `git clone` to clone the repository via SSH.
     */
    readonly sshCloneUrl: string;
    /**
     * URL that can be provided to `svn checkout` to check out the repository via GitHub's Subversion protocol emulation.
     */
    readonly svnUrl: string;
    /**
     * The list of topics of the repository.
     */
    readonly topics: string[];
    /**
     * Whether the repository is public, private or internal.
     */
    readonly visibility: string;
}

export function getRepositoryOutput(args?: GetRepositoryOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRepositoryResult> {
    return pulumi.output(args).apply(a => getRepository(a, opts))
}

/**
 * A collection of arguments for invoking getRepository.
 */
export interface GetRepositoryOutputArgs {
    /**
     * A description of the repository.
     */
    description?: pulumi.Input<string>;
    /**
     * Full name of the repository (in `org/name` format).
     */
    fullName?: pulumi.Input<string>;
    /**
     * URL of a page describing the project.
     */
    homepageUrl?: pulumi.Input<string>;
    /**
     * The name of the repository.
     */
    name?: pulumi.Input<string>;
}
