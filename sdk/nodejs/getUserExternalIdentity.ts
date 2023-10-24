// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use this data source to retrieve a specific organization member's SAML or SCIM user
 * attributes.
 */
export function getUserExternalIdentity(args: GetUserExternalIdentityArgs, opts?: pulumi.InvokeOptions): Promise<GetUserExternalIdentityResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("github:index/getUserExternalIdentity:getUserExternalIdentity", {
        "username": args.username,
    }, opts);
}

/**
 * A collection of arguments for invoking getUserExternalIdentity.
 */
export interface GetUserExternalIdentityArgs {
    /**
     * The username of the member to fetch external identity for.
     */
    username: string;
}

/**
 * A collection of values returned by getUserExternalIdentity.
 */
export interface GetUserExternalIdentityResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The username of the GitHub user
     */
    readonly login: string;
    /**
     * An Object containing the user's SAML data. This object will
     * be empty if the user is not managed by SAML.
     */
    readonly samlIdentity: {[key: string]: string};
    /**
     * An Object contining the user's SCIM data. This object will
     * be empty if the user is not managed by SCIM.
     */
    readonly scimIdentity: {[key: string]: string};
    /**
     * The member's SAML Username
     */
    readonly username: string;
}
/**
 * Use this data source to retrieve a specific organization member's SAML or SCIM user
 * attributes.
 */
export function getUserExternalIdentityOutput(args: GetUserExternalIdentityOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetUserExternalIdentityResult> {
    return pulumi.output(args).apply((a: any) => getUserExternalIdentity(a, opts))
}

/**
 * A collection of arguments for invoking getUserExternalIdentity.
 */
export interface GetUserExternalIdentityOutputArgs {
    /**
     * The username of the member to fetch external identity for.
     */
    username: pulumi.Input<string>;
}
