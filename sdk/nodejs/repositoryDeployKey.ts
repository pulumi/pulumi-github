// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

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
 * import * as tls from "@pulumi/tls";
 *
 * // Generate an ssh key using provider "hashicorp/tls"
 * const exampleRepositoryDeployKeyPrivateKey = new tls.PrivateKey("exampleRepositoryDeployKeyPrivateKey", {algorithm: "ED25519"});
 * // Add the ssh key as a deploy key
 * const exampleRepositoryDeployKeyRepositoryDeployKey = new github.RepositoryDeployKey("exampleRepositoryDeployKeyRepositoryDeployKey", {
 *     title: "Repository test key",
 *     repository: "test-repo",
 *     key: exampleRepositoryDeployKeyPrivateKey.publicKeyOpenssh,
 *     readOnly: true,
 * });
 * ```
 *
 * ## Import
 *
 * Repository deploy keys can be imported using a colon-separated pair of repository name
 * and GitHub's key id. The latter can be obtained by GitHub's SDKs and API.
 *
 * ```sh
 * $ pulumi import github:index/repositoryDeployKey:RepositoryDeployKey foo test-repo:23824728
 * ```
 */
export class RepositoryDeployKey extends pulumi.CustomResource {
    /**
     * Get an existing RepositoryDeployKey resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RepositoryDeployKeyState, opts?: pulumi.CustomResourceOptions): RepositoryDeployKey {
        return new RepositoryDeployKey(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'github:index/repositoryDeployKey:RepositoryDeployKey';

    /**
     * Returns true if the given object is an instance of RepositoryDeployKey.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RepositoryDeployKey {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RepositoryDeployKey.__pulumiType;
    }

    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * A SSH key.
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
     *
     * Changing any of the fields forces re-creating the resource.
     */
    public readonly title!: pulumi.Output<string>;

    /**
     * Create a RepositoryDeployKey resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RepositoryDeployKeyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RepositoryDeployKeyArgs | RepositoryDeployKeyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RepositoryDeployKeyState | undefined;
            resourceInputs["etag"] = state ? state.etag : undefined;
            resourceInputs["key"] = state ? state.key : undefined;
            resourceInputs["readOnly"] = state ? state.readOnly : undefined;
            resourceInputs["repository"] = state ? state.repository : undefined;
            resourceInputs["title"] = state ? state.title : undefined;
        } else {
            const args = argsOrState as RepositoryDeployKeyArgs | undefined;
            if ((!args || args.key === undefined) && !opts.urn) {
                throw new Error("Missing required property 'key'");
            }
            if ((!args || args.repository === undefined) && !opts.urn) {
                throw new Error("Missing required property 'repository'");
            }
            if ((!args || args.title === undefined) && !opts.urn) {
                throw new Error("Missing required property 'title'");
            }
            resourceInputs["key"] = args ? args.key : undefined;
            resourceInputs["readOnly"] = args ? args.readOnly : undefined;
            resourceInputs["repository"] = args ? args.repository : undefined;
            resourceInputs["title"] = args ? args.title : undefined;
            resourceInputs["etag"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RepositoryDeployKey.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RepositoryDeployKey resources.
 */
export interface RepositoryDeployKeyState {
    etag?: pulumi.Input<string>;
    /**
     * A SSH key.
     */
    key?: pulumi.Input<string>;
    /**
     * A boolean qualifying the key to be either read only or read/write.
     */
    readOnly?: pulumi.Input<boolean>;
    /**
     * Name of the GitHub repository.
     */
    repository?: pulumi.Input<string>;
    /**
     * A title.
     *
     * Changing any of the fields forces re-creating the resource.
     */
    title?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RepositoryDeployKey resource.
 */
export interface RepositoryDeployKeyArgs {
    /**
     * A SSH key.
     */
    key: pulumi.Input<string>;
    /**
     * A boolean qualifying the key to be either read only or read/write.
     */
    readOnly?: pulumi.Input<boolean>;
    /**
     * Name of the GitHub repository.
     */
    repository: pulumi.Input<string>;
    /**
     * A title.
     *
     * Changing any of the fields forces re-creating the resource.
     */
    title: pulumi.Input<string>;
}
