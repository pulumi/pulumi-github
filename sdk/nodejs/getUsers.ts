// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use this data source to retrieve information about multiple GitHub users at once.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as github from "@pulumi/github";
 *
 * // Retrieve information about multiple GitHub users.
 * const example = github.getUsers({
 *     usernames: [
 *         "example1",
 *         "example2",
 *         "example3",
 *     ],
 * });
 * export const validUsers = example.then(example => example.logins);
 * export const invalidUsers = example.then(example => example.unknownLogins);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getUsers(args: GetUsersArgs, opts?: pulumi.InvokeOptions): Promise<GetUsersResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("github:index/getUsers:getUsers", {
        "usernames": args.usernames,
    }, opts);
}

/**
 * A collection of arguments for invoking getUsers.
 */
export interface GetUsersArgs {
    /**
     * List of usernames.
     */
    usernames: string[];
}

/**
 * A collection of values returned by getUsers.
 */
export interface GetUsersResult {
    /**
     * list of the user's publicly visible profile email (will be empty string in case if user decided not to show it).
     */
    readonly emails: string[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * list of logins of users that could be found.
     */
    readonly logins: string[];
    /**
     * list of Node IDs of users that could be found.
     */
    readonly nodeIds: string[];
    /**
     * list of logins without matching user.
     */
    readonly unknownLogins: string[];
    readonly usernames: string[];
}
/**
 * Use this data source to retrieve information about multiple GitHub users at once.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as github from "@pulumi/github";
 *
 * // Retrieve information about multiple GitHub users.
 * const example = github.getUsers({
 *     usernames: [
 *         "example1",
 *         "example2",
 *         "example3",
 *     ],
 * });
 * export const validUsers = example.then(example => example.logins);
 * export const invalidUsers = example.then(example => example.unknownLogins);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getUsersOutput(args: GetUsersOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetUsersResult> {
    return pulumi.output(args).apply((a: any) => getUsers(a, opts))
}

/**
 * A collection of arguments for invoking getUsers.
 */
export interface GetUsersOutputArgs {
    /**
     * List of usernames.
     */
    usernames: pulumi.Input<pulumi.Input<string>[]>;
}
