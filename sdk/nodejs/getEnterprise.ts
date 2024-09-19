// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use this data source to retrieve basic information about a GitHub enterprise.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as github from "@pulumi/github";
 *
 * const example = github.getEnterprise({
 *     slug: "example-co",
 * });
 * ```
 */
export function getEnterprise(args: GetEnterpriseArgs, opts?: pulumi.InvokeOptions): Promise<GetEnterpriseResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("github:index/getEnterprise:getEnterprise", {
        "slug": args.slug,
    }, opts);
}

/**
 * A collection of arguments for invoking getEnterprise.
 */
export interface GetEnterpriseArgs {
    /**
     * The URL slug identifying the enterprise.
     */
    slug: string;
}

/**
 * A collection of values returned by getEnterprise.
 */
export interface GetEnterpriseResult {
    /**
     * The time the enterprise was created.
     */
    readonly createdAt: string;
    /**
     * The database ID of the enterprise.
     */
    readonly databaseId: number;
    /**
     * The description of the enterprise.
     */
    readonly description: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The name of the enterprise.
     */
    readonly name: string;
    /**
     * The URL slug identifying the enterprise.
     */
    readonly slug: string;
    /**
     * The url for the enterprise.
     */
    readonly url: string;
}
/**
 * Use this data source to retrieve basic information about a GitHub enterprise.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as github from "@pulumi/github";
 *
 * const example = github.getEnterprise({
 *     slug: "example-co",
 * });
 * ```
 */
export function getEnterpriseOutput(args: GetEnterpriseOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetEnterpriseResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("github:index/getEnterprise:getEnterprise", {
        "slug": args.slug,
    }, opts);
}

/**
 * A collection of arguments for invoking getEnterprise.
 */
export interface GetEnterpriseOutputArgs {
    /**
     * The URL slug identifying the enterprise.
     */
    slug: pulumi.Input<string>;
}
