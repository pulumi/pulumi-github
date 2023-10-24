// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use this data source to generate a [GitHub App JWT](https://docs.github.com/en/apps/creating-github-apps/authenticating-with-a-github-app/generating-a-json-web-token-jwt-for-a-github-app).
 */
export function getAppToken(args: GetAppTokenArgs, opts?: pulumi.InvokeOptions): Promise<GetAppTokenResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("github:index/getAppToken:getAppToken", {
        "appId": args.appId,
        "installationId": args.installationId,
        "pemFile": args.pemFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getAppToken.
 */
export interface GetAppTokenArgs {
    /**
     * This is the ID of the GitHub App.
     */
    appId: string;
    /**
     * This is the ID of the GitHub App installation.
     */
    installationId: string;
    /**
     * This is the contents of the GitHub App private key PEM file.
     */
    pemFile: string;
}

/**
 * A collection of values returned by getAppToken.
 */
export interface GetAppTokenResult {
    readonly appId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly installationId: string;
    readonly pemFile: string;
    /**
     * The generated GitHub APP JWT.
     */
    readonly token: string;
}
/**
 * Use this data source to generate a [GitHub App JWT](https://docs.github.com/en/apps/creating-github-apps/authenticating-with-a-github-app/generating-a-json-web-token-jwt-for-a-github-app).
 */
export function getAppTokenOutput(args: GetAppTokenOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAppTokenResult> {
    return pulumi.output(args).apply((a: any) => getAppToken(a, opts))
}

/**
 * A collection of arguments for invoking getAppToken.
 */
export interface GetAppTokenOutputArgs {
    /**
     * This is the ID of the GitHub App.
     */
    appId: pulumi.Input<string>;
    /**
     * This is the ID of the GitHub App installation.
     */
    installationId: pulumi.Input<string>;
    /**
     * This is the contents of the GitHub App private key PEM file.
     */
    pemFile: pulumi.Input<string>;
}
