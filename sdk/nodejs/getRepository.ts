// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Use this data source to retrieve information about a GitHub repository.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as github from "@pulumi/github";
 *
 * const example = github.getRepository({
 *     fullName: "hashicorp/terraform",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getRepository(args?: GetRepositoryArgs, opts?: pulumi.InvokeOptions): Promise<GetRepositoryResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
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
     * A description of the license.
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
     * The name of the default branch of the repository.
     */
    readonly defaultBranch: string;
    /**
     * A description of the license.
     */
    readonly description?: string;
    /**
     * Whether the repository is a fork.
     */
    readonly fork: boolean;
    readonly fullName: string;
    /**
     * URL that can be provided to `git clone` to clone the repository anonymously via the git protocol.
     */
    readonly gitCloneUrl: string;
    /**
     * Whether the repository has GitHub Discussions enabled.
     */
    readonly hasDiscussions: boolean;
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
     * The URL to view the license details on GitHub.
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
     * Whether the repository is a template repository.
     */
    readonly isTemplate: boolean;
    /**
     * The default value for a merge commit message.
     */
    readonly mergeCommitMessage: string;
    /**
     * The default value for a merge commit title.
     */
    readonly mergeCommitTitle: string;
    /**
     * The name of the license (e.g., "Apache License 2.0").
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
     * The primary language used in the repository.
     */
    readonly primaryLanguage: string;
    /**
     * Whether the repository is private.
     */
    readonly private: boolean;
    /**
     * GitHub ID for the repository
     */
    readonly repoId: number;
    /**
     * An Array of GitHub repository licenses. Each `repositoryLicense` block consists of the fields documented below.
     */
    readonly repositoryLicenses: outputs.GetRepositoryRepositoryLicense[];
    /**
     * The default value for a squash merge commit message.
     */
    readonly squashMergeCommitMessage: string;
    /**
     * The default value for a squash merge commit title.
     */
    readonly squashMergeCommitTitle: string;
    /**
     * URL that can be provided to `git clone` to clone the repository via SSH.
     */
    readonly sshCloneUrl: string;
    /**
     * URL that can be provided to `svn checkout` to check out the repository via GitHub's Subversion protocol emulation.
     */
    readonly svnUrl: string;
    /**
     * The repository source template configuration.
     */
    readonly templates: outputs.GetRepositoryTemplate[];
    /**
     * The list of topics of the repository.
     */
    readonly topics: string[];
    /**
     * Whether the repository is public, private or internal.
     */
    readonly visibility: string;
}
/**
 * Use this data source to retrieve information about a GitHub repository.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as github from "@pulumi/github";
 *
 * const example = github.getRepository({
 *     fullName: "hashicorp/terraform",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getRepositoryOutput(args?: GetRepositoryOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRepositoryResult> {
    return pulumi.output(args).apply((a: any) => getRepository(a, opts))
}

/**
 * A collection of arguments for invoking getRepository.
 */
export interface GetRepositoryOutputArgs {
    /**
     * A description of the license.
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
