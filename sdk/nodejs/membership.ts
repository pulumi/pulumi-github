// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Provides a GitHub membership resource.
 *
 * This resource allows you to add/remove users from your organization. When applied,
 * an invitation will be sent to the user to become part of the organization. When
 * destroyed, either the invitation will be cancelled or the user will be removed.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as github from "@pulumi/github";
 *
 * // Add a user to the organization
 * const membershipForSomeUser = new github.Membership("membership_for_some_user", {
 *     role: "member",
 *     username: "SomeUser",
 * });
 * ```
 *
 * ## Import
 *
 * GitHub Membership can be imported using an ID made up of `organization:username`, e.g.
 *
 * ```sh
 *  $ pulumi import github:index/membership:Membership member hashicorp:someuser
 * ```
 */
export class Membership extends pulumi.CustomResource {
    /**
     * Get an existing Membership resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MembershipState, opts?: pulumi.CustomResourceOptions): Membership {
        return new Membership(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'github:index/membership:Membership';

    /**
     * Returns true if the given object is an instance of Membership.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Membership {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Membership.__pulumiType;
    }

    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * The role of the user within the organization.
     * Must be one of `member` or `admin`. Defaults to `member`.
     */
    public readonly role!: pulumi.Output<string | undefined>;
    /**
     * The user to add to the organization.
     */
    public readonly username!: pulumi.Output<string>;

    /**
     * Create a Membership resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MembershipArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MembershipArgs | MembershipState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as MembershipState | undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["role"] = state ? state.role : undefined;
            inputs["username"] = state ? state.username : undefined;
        } else {
            const args = argsOrState as MembershipArgs | undefined;
            if ((!args || args.username === undefined) && !opts.urn) {
                throw new Error("Missing required property 'username'");
            }
            inputs["role"] = args ? args.role : undefined;
            inputs["username"] = args ? args.username : undefined;
            inputs["etag"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Membership.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Membership resources.
 */
export interface MembershipState {
    etag?: pulumi.Input<string>;
    /**
     * The role of the user within the organization.
     * Must be one of `member` or `admin`. Defaults to `member`.
     */
    role?: pulumi.Input<string>;
    /**
     * The user to add to the organization.
     */
    username?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Membership resource.
 */
export interface MembershipArgs {
    /**
     * The role of the user within the organization.
     * Must be one of `member` or `admin`. Defaults to `member`.
     */
    role?: pulumi.Input<string>;
    /**
     * The user to add to the organization.
     */
    username: pulumi.Input<string>;
}
