// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This data source allows you to read files within a
 * GitHub repository.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as github from "@pulumi/github";
 *
 * const foo = github.getRepositoryFile({
 *     repository: fooGithubRepository.name,
 *     branch: "main",
 *     file: ".gitignore",
 * });
 * ```
 */
export function getRepositoryFile(args: GetRepositoryFileArgs, opts?: pulumi.InvokeOptions): Promise<GetRepositoryFileResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("github:index/getRepositoryFile:getRepositoryFile", {
        "branch": args.branch,
        "file": args.file,
        "repository": args.repository,
    }, opts);
}

/**
 * A collection of arguments for invoking getRepositoryFile.
 */
export interface GetRepositoryFileArgs {
    /**
     * Git branch. Defaults to the repository's default branch.
     */
    branch?: string;
    /**
     * The path of the file to read.
     */
    file: string;
    /**
     * The repository to read the file from. If an unqualified repo name (without an owner) is passed, the owner will be inferred from the owner of the token used to execute the plan. If a name of the type "owner/repo" (with a slash in the middle) is passed, the owner will be as specified and not the owner of the token.
     */
    repository: string;
}

/**
 * A collection of values returned by getRepositoryFile.
 */
export interface GetRepositoryFileResult {
    readonly branch?: string;
    /**
     * Committer author name.
     */
    readonly commitAuthor: string;
    /**
     * Committer email address.
     */
    readonly commitEmail: string;
    /**
     * Commit message when file was last updated.
     */
    readonly commitMessage: string;
    /**
     * The SHA of the commit that modified the file.
     */
    readonly commitSha: string;
    /**
     * The file content.
     */
    readonly content: string;
    readonly file: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The name of the commit/branch/tag.
     */
    readonly ref: string;
    readonly repository: string;
    /**
     * The SHA blob of the file.
     */
    readonly sha: string;
}
/**
 * This data source allows you to read files within a
 * GitHub repository.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as github from "@pulumi/github";
 *
 * const foo = github.getRepositoryFile({
 *     repository: fooGithubRepository.name,
 *     branch: "main",
 *     file: ".gitignore",
 * });
 * ```
 */
export function getRepositoryFileOutput(args: GetRepositoryFileOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetRepositoryFileResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("github:index/getRepositoryFile:getRepositoryFile", {
        "branch": args.branch,
        "file": args.file,
        "repository": args.repository,
    }, opts);
}

/**
 * A collection of arguments for invoking getRepositoryFile.
 */
export interface GetRepositoryFileOutputArgs {
    /**
     * Git branch. Defaults to the repository's default branch.
     */
    branch?: pulumi.Input<string>;
    /**
     * The path of the file to read.
     */
    file: pulumi.Input<string>;
    /**
     * The repository to read the file from. If an unqualified repo name (without an owner) is passed, the owner will be inferred from the owner of the token used to execute the plan. If a name of the type "owner/repo" (with a slash in the middle) is passed, the owner will be as specified and not the owner of the token.
     */
    repository: pulumi.Input<string>;
}
