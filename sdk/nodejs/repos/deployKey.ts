// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a GitHub repository deploy key resource.
 * 
 * A deploy key is an SSH key that is stored on your server and grants
 * access to a single GitHub repository. This key is attached directly to the repository instead of to a personal user
 * account.
 * 
 * This resource allows you to add/remove repository deploy keys.
 * 
 * Further documentation on GitHub repository deploy keys:
 * - [About deploy keys](https://developer.github.com/guides/managing-deploy-keys/#deploy-keys)
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as github from "@pulumi/github";
 * 
 * // Add a deploy key
 * const exampleRepositoryDeployKey = new github.repos.DeployKey("example_repository_deploy_key", {
 *     key: "ssh-rsa AAA...",
 *     readOnly: false,
 *     repository: "test-repo",
 *     title: "Repository test key",
 * });
 * ```
 */
export class DeployKey extends pulumi.CustomResource {
    /**
     * Get an existing DeployKey resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DeployKeyState, opts?: pulumi.CustomResourceOptions): DeployKey {
        return new DeployKey(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'github:repos/deployKey:DeployKey';

    /**
     * Returns true if the given object is an instance of DeployKey.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DeployKey {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DeployKey.__pulumiType;
    }

    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * A ssh key.
     */
    public readonly key!: pulumi.Output<string>;
    /**
     * A boolean qualifying the key to be either read only or read/write.
     */
    public readonly readOnly!: pulumi.Output<boolean | undefined>;
    /**
     * Name of the GitHub repository.
     */
    public readonly repository!: pulumi.Output<string>;
    /**
     * A title.
     */
    public readonly title!: pulumi.Output<string>;

    /**
     * Create a DeployKey resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DeployKeyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DeployKeyArgs | DeployKeyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as DeployKeyState | undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["key"] = state ? state.key : undefined;
            inputs["readOnly"] = state ? state.readOnly : undefined;
            inputs["repository"] = state ? state.repository : undefined;
            inputs["title"] = state ? state.title : undefined;
        } else {
            const args = argsOrState as DeployKeyArgs | undefined;
            if (!args || args.key === undefined) {
                throw new Error("Missing required property 'key'");
            }
            if (!args || args.repository === undefined) {
                throw new Error("Missing required property 'repository'");
            }
            if (!args || args.title === undefined) {
                throw new Error("Missing required property 'title'");
            }
            inputs["key"] = args ? args.key : undefined;
            inputs["readOnly"] = args ? args.readOnly : undefined;
            inputs["repository"] = args ? args.repository : undefined;
            inputs["title"] = args ? args.title : undefined;
            inputs["etag"] = undefined /*out*/;
        }
        super(DeployKey.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DeployKey resources.
 */
export interface DeployKeyState {
    readonly etag?: pulumi.Input<string>;
    /**
     * A ssh key.
     */
    readonly key?: pulumi.Input<string>;
    /**
     * A boolean qualifying the key to be either read only or read/write.
     */
    readonly readOnly?: pulumi.Input<boolean>;
    /**
     * Name of the GitHub repository.
     */
    readonly repository?: pulumi.Input<string>;
    /**
     * A title.
     */
    readonly title?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DeployKey resource.
 */
export interface DeployKeyArgs {
    /**
     * A ssh key.
     */
    readonly key: pulumi.Input<string>;
    /**
     * A boolean qualifying the key to be either read only or read/write.
     */
    readonly readOnly?: pulumi.Input<boolean>;
    /**
     * Name of the GitHub repository.
     */
    readonly repository: pulumi.Input<string>;
    /**
     * A title.
     */
    readonly title: pulumi.Input<string>;
}
