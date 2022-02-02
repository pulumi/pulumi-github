// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Provides a GitHub user's SSH key resource.
 *
 * This resource allows you to add/remove SSH keys from your user account.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as github from "@pulumi/github";
 * import * from "fs";
 *
 * const example = new github.UserSshKey("example", {
 *     title: "example title",
 *     key: fs.readFileSync("~/.ssh/id_rsa.pub"),
 * });
 * ```
 *
 * ## Import
 *
 * SSH keys can be imported using their ID e.g.
 *
 * ```sh
 *  $ pulumi import github:index/userSshKey:UserSshKey example 1234567
 * ```
 */
export class UserSshKey extends pulumi.CustomResource {
    /**
     * Get an existing UserSshKey resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UserSshKeyState, opts?: pulumi.CustomResourceOptions): UserSshKey {
        return new UserSshKey(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'github:index/userSshKey:UserSshKey';

    /**
     * Returns true if the given object is an instance of UserSshKey.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is UserSshKey {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === UserSshKey.__pulumiType;
    }

    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * The public SSH key to add to your GitHub account.
     */
    public readonly key!: pulumi.Output<string>;
    /**
     * A descriptive name for the new key. e.g. `Personal MacBook Air`
     */
    public readonly title!: pulumi.Output<string>;
    /**
     * The URL of the SSH key
     */
    public /*out*/ readonly url!: pulumi.Output<string>;

    /**
     * Create a UserSshKey resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UserSshKeyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UserSshKeyArgs | UserSshKeyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as UserSshKeyState | undefined;
            resourceInputs["etag"] = state ? state.etag : undefined;
            resourceInputs["key"] = state ? state.key : undefined;
            resourceInputs["title"] = state ? state.title : undefined;
            resourceInputs["url"] = state ? state.url : undefined;
        } else {
            const args = argsOrState as UserSshKeyArgs | undefined;
            if ((!args || args.key === undefined) && !opts.urn) {
                throw new Error("Missing required property 'key'");
            }
            if ((!args || args.title === undefined) && !opts.urn) {
                throw new Error("Missing required property 'title'");
            }
            resourceInputs["key"] = args ? args.key : undefined;
            resourceInputs["title"] = args ? args.title : undefined;
            resourceInputs["etag"] = undefined /*out*/;
            resourceInputs["url"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(UserSshKey.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering UserSshKey resources.
 */
export interface UserSshKeyState {
    etag?: pulumi.Input<string>;
    /**
     * The public SSH key to add to your GitHub account.
     */
    key?: pulumi.Input<string>;
    /**
     * A descriptive name for the new key. e.g. `Personal MacBook Air`
     */
    title?: pulumi.Input<string>;
    /**
     * The URL of the SSH key
     */
    url?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a UserSshKey resource.
 */
export interface UserSshKeyArgs {
    /**
     * The public SSH key to add to your GitHub account.
     */
    key: pulumi.Input<string>;
    /**
     * A descriptive name for the new key. e.g. `Personal MacBook Air`
     */
    title: pulumi.Input<string>;
}
