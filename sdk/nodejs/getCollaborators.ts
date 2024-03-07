// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Use this data source to retrieve the collaborators for a given repository.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as github from "@pulumi/github";
 *
 * const test = github.getCollaborators({
 *     owner: "example_owner",
 *     repository: "example_repository",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getCollaborators(args: GetCollaboratorsArgs, opts?: pulumi.InvokeOptions): Promise<GetCollaboratorsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("github:index/getCollaborators:getCollaborators", {
        "affiliation": args.affiliation,
        "owner": args.owner,
        "repository": args.repository,
    }, opts);
}

/**
 * A collection of arguments for invoking getCollaborators.
 */
export interface GetCollaboratorsArgs {
    /**
     * Filter collaborators returned by their affiliation. Can be one of: `outside`, `direct`, `all`.  Defaults to `all`.
     */
    affiliation?: string;
    /**
     * The organization that owns the repository.
     */
    owner: string;
    /**
     * The name of the repository.
     */
    repository: string;
}

/**
 * A collection of values returned by getCollaborators.
 */
export interface GetCollaboratorsResult {
    readonly affiliation?: string;
    /**
     * An Array of GitHub collaborators.  Each `collaborator` block consists of the fields documented below.
     */
    readonly collaborators: outputs.GetCollaboratorsCollaborator[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly owner: string;
    readonly repository: string;
}
/**
 * Use this data source to retrieve the collaborators for a given repository.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as github from "@pulumi/github";
 *
 * const test = github.getCollaborators({
 *     owner: "example_owner",
 *     repository: "example_repository",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getCollaboratorsOutput(args: GetCollaboratorsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetCollaboratorsResult> {
    return pulumi.output(args).apply((a: any) => getCollaborators(a, opts))
}

/**
 * A collection of arguments for invoking getCollaborators.
 */
export interface GetCollaboratorsOutputArgs {
    /**
     * Filter collaborators returned by their affiliation. Can be one of: `outside`, `direct`, `all`.  Defaults to `all`.
     */
    affiliation?: pulumi.Input<string>;
    /**
     * The organization that owns the repository.
     */
    owner: pulumi.Input<string>;
    /**
     * The name of the repository.
     */
    repository: pulumi.Input<string>;
}
