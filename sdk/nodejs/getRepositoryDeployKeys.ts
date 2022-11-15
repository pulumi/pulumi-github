// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export function getRepositoryDeployKeys(args: GetRepositoryDeployKeysArgs, opts?: pulumi.InvokeOptions): Promise<GetRepositoryDeployKeysResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("github:index/getRepositoryDeployKeys:getRepositoryDeployKeys", {
        "repository": args.repository,
    }, opts);
}

/**
 * A collection of arguments for invoking getRepositoryDeployKeys.
 */
export interface GetRepositoryDeployKeysArgs {
    repository: string;
}

/**
 * A collection of values returned by getRepositoryDeployKeys.
 */
export interface GetRepositoryDeployKeysResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly keys: outputs.GetRepositoryDeployKeysKey[];
    readonly repository: string;
}

export function getRepositoryDeployKeysOutput(args: GetRepositoryDeployKeysOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRepositoryDeployKeysResult> {
    return pulumi.output(args).apply(a => getRepositoryDeployKeys(a, opts))
}

/**
 * A collection of arguments for invoking getRepositoryDeployKeys.
 */
export interface GetRepositoryDeployKeysOutputArgs {
    repository: pulumi.Input<string>;
}
