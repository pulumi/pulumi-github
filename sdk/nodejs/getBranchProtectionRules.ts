// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Use this data source to retrieve a list of repository branch protection rules.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as github from "@pulumi/github";
 *
 * const example = github.getBranchProtectionRules({
 *     repository: "example",
 * });
 * ```
 */
export function getBranchProtectionRules(args: GetBranchProtectionRulesArgs, opts?: pulumi.InvokeOptions): Promise<GetBranchProtectionRulesResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("github:index/getBranchProtectionRules:getBranchProtectionRules", {
        "repository": args.repository,
    }, opts);
}

/**
 * A collection of arguments for invoking getBranchProtectionRules.
 */
export interface GetBranchProtectionRulesArgs {
    /**
     * The GitHub repository name.
     */
    repository: string;
}

/**
 * A collection of values returned by getBranchProtectionRules.
 */
export interface GetBranchProtectionRulesResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly repository: string;
    /**
     * Collection of Branch Protection Rules. Each of the results conforms to the following scheme:
     */
    readonly rules: outputs.GetBranchProtectionRulesRule[];
}
/**
 * Use this data source to retrieve a list of repository branch protection rules.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as github from "@pulumi/github";
 *
 * const example = github.getBranchProtectionRules({
 *     repository: "example",
 * });
 * ```
 */
export function getBranchProtectionRulesOutput(args: GetBranchProtectionRulesOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetBranchProtectionRulesResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("github:index/getBranchProtectionRules:getBranchProtectionRules", {
        "repository": args.repository,
    }, opts);
}

/**
 * A collection of arguments for invoking getBranchProtectionRules.
 */
export interface GetBranchProtectionRulesOutputArgs {
    /**
     * The GitHub repository name.
     */
    repository: pulumi.Input<string>;
}
