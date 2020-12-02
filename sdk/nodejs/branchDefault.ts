// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Provides a GitHub branch default resource.
 *
 * This resource allows you to set the default branch for a given repository.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as github from "@pulumi/github";
 *
 * const example = new github.Repository("example", {
 *     description: "My awesome codebase",
 *     "private": true,
 *     template: {
 *         owner: "github",
 *         repository: "terraform-module-template",
 *     },
 * });
 * const development = new github.Branch("development", {
 *     repository: example.name,
 *     branch: "development",
 * });
 * const _default = new github.BranchDefault("default", {
 *     repository: example.name,
 *     branch: development.branch,
 * });
 * ```
 */
export class BranchDefault extends pulumi.CustomResource {
    /**
     * Get an existing BranchDefault resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BranchDefaultState, opts?: pulumi.CustomResourceOptions): BranchDefault {
        return new BranchDefault(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'github:index/branchDefault:BranchDefault';

    /**
     * Returns true if the given object is an instance of BranchDefault.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BranchDefault {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BranchDefault.__pulumiType;
    }

    /**
     * The branch (e.g. `main`)
     */
    public readonly branch!: pulumi.Output<string>;
    /**
     * The GitHub repository
     */
    public readonly repository!: pulumi.Output<string>;

    /**
     * Create a BranchDefault resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BranchDefaultArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BranchDefaultArgs | BranchDefaultState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as BranchDefaultState | undefined;
            inputs["branch"] = state ? state.branch : undefined;
            inputs["repository"] = state ? state.repository : undefined;
        } else {
            const args = argsOrState as BranchDefaultArgs | undefined;
            if (!args || args.branch === undefined) {
                throw new Error("Missing required property 'branch'");
            }
            if (!args || args.repository === undefined) {
                throw new Error("Missing required property 'repository'");
            }
            inputs["branch"] = args ? args.branch : undefined;
            inputs["repository"] = args ? args.repository : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(BranchDefault.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BranchDefault resources.
 */
export interface BranchDefaultState {
    /**
     * The branch (e.g. `main`)
     */
    readonly branch?: pulumi.Input<string>;
    /**
     * The GitHub repository
     */
    readonly repository?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a BranchDefault resource.
 */
export interface BranchDefaultArgs {
    /**
     * The branch (e.g. `main`)
     */
    readonly branch: pulumi.Input<string>;
    /**
     * The GitHub repository
     */
    readonly repository: pulumi.Input<string>;
}
