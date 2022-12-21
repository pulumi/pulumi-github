// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getRelease(args: GetReleaseArgs, opts?: pulumi.InvokeOptions): Promise<GetReleaseResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("github:index/getRelease:getRelease", {
        "owner": args.owner,
        "releaseId": args.releaseId,
        "releaseTag": args.releaseTag,
        "repository": args.repository,
        "retrieveBy": args.retrieveBy,
    }, opts);
}

/**
 * A collection of arguments for invoking getRelease.
 */
export interface GetReleaseArgs {
    owner: string;
    releaseId?: number;
    releaseTag?: string;
    repository: string;
    retrieveBy: string;
}

/**
 * A collection of values returned by getRelease.
 */
export interface GetReleaseResult {
    /**
     * @deprecated use assets_url instead
     */
    readonly assertsUrl: string;
    readonly assets: outputs.GetReleaseAsset[];
    readonly assetsUrl: string;
    readonly body: string;
    readonly createdAt: string;
    readonly draft: boolean;
    readonly htmlUrl: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    readonly owner: string;
    readonly prerelease: boolean;
    readonly publishedAt: string;
    readonly releaseId?: number;
    readonly releaseTag?: string;
    readonly repository: string;
    readonly retrieveBy: string;
    readonly tarballUrl: string;
    readonly targetCommitish: string;
    readonly uploadUrl: string;
    readonly url: string;
    readonly zipballUrl: string;
}

export function getReleaseOutput(args: GetReleaseOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetReleaseResult> {
    return pulumi.output(args).apply(a => getRelease(a, opts))
}

/**
 * A collection of arguments for invoking getRelease.
 */
export interface GetReleaseOutputArgs {
    owner: pulumi.Input<string>;
    releaseId?: pulumi.Input<number>;
    releaseTag?: pulumi.Input<string>;
    repository: pulumi.Input<string>;
    retrieveBy: pulumi.Input<string>;
}
