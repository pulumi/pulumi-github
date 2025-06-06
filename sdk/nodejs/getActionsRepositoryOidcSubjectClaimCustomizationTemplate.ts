// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Use this data source to retrieve the OpenID Connect subject claim customization template for a repository
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as github from "@pulumi/github";
 *
 * const example = github.getActionsRepositoryOidcSubjectClaimCustomizationTemplate({
 *     name: "example_repository",
 * });
 * ```
 */
export function getActionsRepositoryOidcSubjectClaimCustomizationTemplate(args: GetActionsRepositoryOidcSubjectClaimCustomizationTemplateArgs, opts?: pulumi.InvokeOptions): Promise<GetActionsRepositoryOidcSubjectClaimCustomizationTemplateResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("github:index/getActionsRepositoryOidcSubjectClaimCustomizationTemplate:getActionsRepositoryOidcSubjectClaimCustomizationTemplate", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getActionsRepositoryOidcSubjectClaimCustomizationTemplate.
 */
export interface GetActionsRepositoryOidcSubjectClaimCustomizationTemplateArgs {
    /**
     * Name of the repository to get the OpenID Connect subject claim customization template for.
     */
    name: string;
}

/**
 * A collection of values returned by getActionsRepositoryOidcSubjectClaimCustomizationTemplate.
 */
export interface GetActionsRepositoryOidcSubjectClaimCustomizationTemplateResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The list of OpenID Connect claim keys.
     */
    readonly includeClaimKeys: string[];
    readonly name: string;
    /**
     * Whether the repository uses the default template.
     */
    readonly useDefault: boolean;
}
/**
 * Use this data source to retrieve the OpenID Connect subject claim customization template for a repository
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as github from "@pulumi/github";
 *
 * const example = github.getActionsRepositoryOidcSubjectClaimCustomizationTemplate({
 *     name: "example_repository",
 * });
 * ```
 */
export function getActionsRepositoryOidcSubjectClaimCustomizationTemplateOutput(args: GetActionsRepositoryOidcSubjectClaimCustomizationTemplateOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetActionsRepositoryOidcSubjectClaimCustomizationTemplateResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("github:index/getActionsRepositoryOidcSubjectClaimCustomizationTemplate:getActionsRepositoryOidcSubjectClaimCustomizationTemplate", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getActionsRepositoryOidcSubjectClaimCustomizationTemplate.
 */
export interface GetActionsRepositoryOidcSubjectClaimCustomizationTemplateOutputArgs {
    /**
     * Name of the repository to get the OpenID Connect subject claim customization template for.
     */
    name: pulumi.Input<string>;
}
