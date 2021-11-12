// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Provides a GitHub user's GPG key resource.
 *
 * This resource allows you to add/remove GPG keys from your user account.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as github from "@pulumi/github";
 *
 * const example = new github.UserGpgKey("example", {
 *     armoredPublicKey: `-----BEGIN PGP PUBLIC KEY BLOCK-----
 * ...
 * -----END PGP PUBLIC KEY BLOCK-----`,
 * });
 * ```
 *
 * ## Import
 *
 * GPG keys are not importable due to the fact that [API](https://developer.github.com/v3/users/gpg_keys/#gpg-keys) does not return previously uploaded GPG key.
 */
export class UserGpgKey extends pulumi.CustomResource {
    /**
     * Get an existing UserGpgKey resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UserGpgKeyState, opts?: pulumi.CustomResourceOptions): UserGpgKey {
        return new UserGpgKey(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'github:index/userGpgKey:UserGpgKey';

    /**
     * Returns true if the given object is an instance of UserGpgKey.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is UserGpgKey {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === UserGpgKey.__pulumiType;
    }

    /**
     * Your public GPG key, generated in ASCII-armored format.
     * See [Generating a new GPG key](https://help.github.com/articles/generating-a-new-gpg-key/) for help on creating a GPG key.
     */
    public readonly armoredPublicKey!: pulumi.Output<string>;
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * The key ID of the GPG key, e.g. `3262EFF25BA0D270`
     */
    public /*out*/ readonly keyId!: pulumi.Output<string>;

    /**
     * Create a UserGpgKey resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UserGpgKeyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UserGpgKeyArgs | UserGpgKeyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as UserGpgKeyState | undefined;
            resourceInputs["armoredPublicKey"] = state ? state.armoredPublicKey : undefined;
            resourceInputs["etag"] = state ? state.etag : undefined;
            resourceInputs["keyId"] = state ? state.keyId : undefined;
        } else {
            const args = argsOrState as UserGpgKeyArgs | undefined;
            if ((!args || args.armoredPublicKey === undefined) && !opts.urn) {
                throw new Error("Missing required property 'armoredPublicKey'");
            }
            resourceInputs["armoredPublicKey"] = args ? args.armoredPublicKey : undefined;
            resourceInputs["etag"] = undefined /*out*/;
            resourceInputs["keyId"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(UserGpgKey.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering UserGpgKey resources.
 */
export interface UserGpgKeyState {
    /**
     * Your public GPG key, generated in ASCII-armored format.
     * See [Generating a new GPG key](https://help.github.com/articles/generating-a-new-gpg-key/) for help on creating a GPG key.
     */
    armoredPublicKey?: pulumi.Input<string>;
    etag?: pulumi.Input<string>;
    /**
     * The key ID of the GPG key, e.g. `3262EFF25BA0D270`
     */
    keyId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a UserGpgKey resource.
 */
export interface UserGpgKeyArgs {
    /**
     * Your public GPG key, generated in ASCII-armored format.
     * See [Generating a new GPG key](https://help.github.com/articles/generating-a-new-gpg-key/) for help on creating a GPG key.
     */
    armoredPublicKey: pulumi.Input<string>;
}
