// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a GitHub user's SSH key resource.
 * 
 * This resource allows you to add/remove SSH keys from your user account.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fs from "fs";
 * import * as github from "@pulumi/github";
 * 
 * const example = new github.users.SshKey("example", {
 *     key: fs.readFileSync("~/.ssh/id_rsa.pub", "utf-8"),
 *     title: "example title",
 * });
 * ```
 */
export class SshKey extends pulumi.CustomResource {
    /**
     * Get an existing SshKey resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SshKeyState, opts?: pulumi.CustomResourceOptions): SshKey {
        return new SshKey(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'github:users/sshKey:SshKey';

    /**
     * Returns true if the given object is an instance of SshKey.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SshKey {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SshKey.__pulumiType;
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
     * Create a SshKey resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SshKeyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SshKeyArgs | SshKeyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as SshKeyState | undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["key"] = state ? state.key : undefined;
            inputs["title"] = state ? state.title : undefined;
            inputs["url"] = state ? state.url : undefined;
        } else {
            const args = argsOrState as SshKeyArgs | undefined;
            if (!args || args.key === undefined) {
                throw new Error("Missing required property 'key'");
            }
            if (!args || args.title === undefined) {
                throw new Error("Missing required property 'title'");
            }
            inputs["key"] = args ? args.key : undefined;
            inputs["title"] = args ? args.title : undefined;
            inputs["etag"] = undefined /*out*/;
            inputs["url"] = undefined /*out*/;
        }
        super(SshKey.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SshKey resources.
 */
export interface SshKeyState {
    readonly etag?: pulumi.Input<string>;
    /**
     * The public SSH key to add to your GitHub account.
     */
    readonly key?: pulumi.Input<string>;
    /**
     * A descriptive name for the new key. e.g. `Personal MacBook Air`
     */
    readonly title?: pulumi.Input<string>;
    /**
     * The URL of the SSH key
     */
    readonly url?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SshKey resource.
 */
export interface SshKeyArgs {
    /**
     * The public SSH key to add to your GitHub account.
     */
    readonly key: pulumi.Input<string>;
    /**
     * A descriptive name for the new key. e.g. `Personal MacBook Air`
     */
    readonly title: pulumi.Input<string>;
}
