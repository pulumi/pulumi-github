// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to retrieve information about a GitHub repository.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as github from "@pulumi/github";
 * 
 * const example = pulumi.output(github.repos.getRepository({
 *     fullName: "hashicorp/terraform",
 * }));
 * ```
 */
export function getRepository(args?: GetRepositoryArgs, opts?: pulumi.InvokeOptions): Promise<GetRepositoryResult> {
    args = args || {};
    return pulumi.runtime.invoke("github:repos/getRepository:getRepository", {
        "fullName": args.fullName,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getRepository.
 */
export interface GetRepositoryArgs {
    /**
     * Full name of the repository (in `org/name` format).
     */
    readonly fullName?: string;
    /**
     * The name of the repository.
     */
    readonly name?: string;
}

/**
 * A collection of values returned by getRepository.
 */
export interface GetRepositoryResult {
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
     * The name of the default branch of the repository.
     */
    readonly defaultBranch: string;
    /**
     * A description of the repository.
     */
    readonly description: string;
    readonly fullName?: string;
    /**
     * URL that can be provided to `git clone` to clone the
     * repository anonymously via the git protocol.
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
    readonly homepageUrl: string;
    /**
     * URL to the repository on the web.
     */
    readonly htmlUrl: string;
    /**
     * URL that can be provided to `git clone` to clone the
     * repository via HTTPS.
     */
    readonly httpCloneUrl: string;
    readonly name?: string;
    /**
     * Whether the repository is private.
     */
    readonly private: boolean;
    /**
     * URL that can be provided to `git clone` to clone the
     * repository via SSH.
     */
    readonly sshCloneUrl: string;
    /**
     * URL that can be provided to `svn checkout` to check out
     * the repository via GitHub's Subversion protocol emulation.
     */
    readonly svnUrl: string;
    /**
     * The list of topics of the repository.
     */
    readonly topics: string[];
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
