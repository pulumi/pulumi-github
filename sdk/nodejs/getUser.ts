// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Use this data source to retrieve information about a GitHub user.
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as github from "@pulumi/github";
 * 
 * const example = pulumi.output(github.getUser({
 *     username: "example",
 * }, { async: true }));
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-github/blob/master/website/docs/d/user.html.markdown.
 */
export function getUser(args: GetUserArgs, opts?: pulumi.InvokeOptions): Promise<GetUserResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("github:index/getUser:getUser", {
        "username": args.username,
    }, opts);
}

/**
 * A collection of arguments for invoking getUser.
 */
export interface GetUserArgs {
    /**
     * The username.
     */
    readonly username: string;
}

/**
 * A collection of values returned by getUser.
 */
export interface GetUserResult {
    /**
     * the user's avatar URL.
     */
    readonly avatarUrl: string;
    /**
     * the user's bio.
     */
    readonly bio: string;
    /**
     * the user's blog location.
     */
    readonly blog: string;
    /**
     * the user's company name.
     */
    readonly company: string;
    /**
     * the creation date.
     */
    readonly createdAt: string;
    /**
     * the user's email.
     */
    readonly email: string;
    /**
     * the number of followers.
     */
    readonly followers: number;
    /**
     * the number of following users.
     */
    readonly following: number;
    /**
     * list of user's GPG keys.
     */
    readonly gpgKeys: string[];
    /**
     * the user's gravatar ID.
     */
    readonly gravatarId: string;
    /**
     * the user's location.
     */
    readonly location: string;
    /**
     * the user's login.
     */
    readonly login: string;
    /**
     * the user's full name.
     */
    readonly name: string;
    readonly nodeId: string;
    /**
     * the number of public gists.
     */
    readonly publicGists: number;
    /**
     * the number of public repositories.
     */
    readonly publicRepos: number;
    /**
     * whether the user is a GitHub admin.
     */
    readonly siteAdmin: boolean;
    /**
     * list of user's SSH keys.
     */
    readonly sshKeys: string[];
    /**
     * the update date.
     */
    readonly updatedAt: string;
    readonly username: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
