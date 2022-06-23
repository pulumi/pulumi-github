// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This data source allows you to read files within a
 * GitHub repository.
 */
export function getRepositoryFile(args: GetRepositoryFileArgs, opts?: pulumi.InvokeOptions): Promise<GetRepositoryFileResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
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
     * Git branch (defaults to `main`).
     * The branch must already exist, it will not be created if it does not already exist.
     */
    branch?: string;
    /**
     * The path of the file to manage.
     */
    file: string;
    /**
     * The repository to create the file in.
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
    readonly repository: string;
    /**
     * The SHA blob of the file.
     */
    readonly sha: string;
}

export function getRepositoryFileOutput(args: GetRepositoryFileOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRepositoryFileResult> {
    return pulumi.output(args).apply(a => getRepositoryFile(a, opts))
}

/**
 * A collection of arguments for invoking getRepositoryFile.
 */
export interface GetRepositoryFileOutputArgs {
    /**
     * Git branch (defaults to `main`).
     * The branch must already exist, it will not be created if it does not already exist.
     */
    branch?: pulumi.Input<string>;
    /**
     * The path of the file to manage.
     */
    file: pulumi.Input<string>;
    /**
     * The repository to create the file in.
     */
    repository: pulumi.Input<string>;
}
