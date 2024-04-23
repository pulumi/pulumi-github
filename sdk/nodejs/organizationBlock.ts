// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This resource allows you to create and manage blocks for GitHub organizations.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as github from "@pulumi/github";
 *
 * const example = new github.OrganizationBlock("example", {username: "paultyng"});
 * ```
 *
 * ## Import
 *
 * GitHub organization block can be imported using a username, e.g.
 *
 * ```sh
 * $ pulumi import github:index/organizationBlock:OrganizationBlock example someuser
 * ```
 */
export class OrganizationBlock extends pulumi.CustomResource {
    /**
     * Get an existing OrganizationBlock resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: OrganizationBlockState, opts?: pulumi.CustomResourceOptions): OrganizationBlock {
        return new OrganizationBlock(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'github:index/organizationBlock:OrganizationBlock';

    /**
     * Returns true if the given object is an instance of OrganizationBlock.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OrganizationBlock {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OrganizationBlock.__pulumiType;
    }

    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * The name of the user to block.
     */
    public readonly username!: pulumi.Output<string>;

    /**
     * Create a OrganizationBlock resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: OrganizationBlockArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: OrganizationBlockArgs | OrganizationBlockState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as OrganizationBlockState | undefined;
            resourceInputs["etag"] = state ? state.etag : undefined;
            resourceInputs["username"] = state ? state.username : undefined;
        } else {
            const args = argsOrState as OrganizationBlockArgs | undefined;
            if ((!args || args.username === undefined) && !opts.urn) {
                throw new Error("Missing required property 'username'");
            }
            resourceInputs["username"] = args ? args.username : undefined;
            resourceInputs["etag"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(OrganizationBlock.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering OrganizationBlock resources.
 */
export interface OrganizationBlockState {
    etag?: pulumi.Input<string>;
    /**
     * The name of the user to block.
     */
    username?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a OrganizationBlock resource.
 */
export interface OrganizationBlockArgs {
    /**
     * The name of the user to block.
     */
    username: pulumi.Input<string>;
}
