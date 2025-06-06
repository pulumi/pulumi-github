// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This resource allows you to manage repository allow list for existing GitHub Actions secrets within your GitHub organization.
 * You must have write access to an organization secret to use this resource.
 *
 * This resource is only applicable when `visibility` of the existing organization secret has been set to `selected`.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as github from "@pulumi/github";
 *
 * const repo = github.getRepository({
 *     fullName: "my-org/repo",
 * });
 * const orgSecretRepos = new github.ActionsOrganizationSecretRepositories("org_secret_repos", {
 *     secretName: "existing_secret_name",
 *     selectedRepositoryIds: [repo.then(repo => repo.repoId)],
 * });
 * ```
 *
 * ## Import
 *
 * This resource can be imported using an ID made up of the secret name:
 *
 * ```sh
 * $ pulumi import github:index/actionsOrganizationSecretRepositories:ActionsOrganizationSecretRepositories test_secret_repos test_secret_name
 * ```
 */
export class ActionsOrganizationSecretRepositories extends pulumi.CustomResource {
    /**
     * Get an existing ActionsOrganizationSecretRepositories resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ActionsOrganizationSecretRepositoriesState, opts?: pulumi.CustomResourceOptions): ActionsOrganizationSecretRepositories {
        return new ActionsOrganizationSecretRepositories(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'github:index/actionsOrganizationSecretRepositories:ActionsOrganizationSecretRepositories';

    /**
     * Returns true if the given object is an instance of ActionsOrganizationSecretRepositories.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ActionsOrganizationSecretRepositories {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ActionsOrganizationSecretRepositories.__pulumiType;
    }

    /**
     * Name of the existing secret
     */
    public readonly secretName!: pulumi.Output<string>;
    /**
     * An array of repository ids that can access the organization secret.
     */
    public readonly selectedRepositoryIds!: pulumi.Output<number[]>;

    /**
     * Create a ActionsOrganizationSecretRepositories resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ActionsOrganizationSecretRepositoriesArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ActionsOrganizationSecretRepositoriesArgs | ActionsOrganizationSecretRepositoriesState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ActionsOrganizationSecretRepositoriesState | undefined;
            resourceInputs["secretName"] = state ? state.secretName : undefined;
            resourceInputs["selectedRepositoryIds"] = state ? state.selectedRepositoryIds : undefined;
        } else {
            const args = argsOrState as ActionsOrganizationSecretRepositoriesArgs | undefined;
            if ((!args || args.secretName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'secretName'");
            }
            if ((!args || args.selectedRepositoryIds === undefined) && !opts.urn) {
                throw new Error("Missing required property 'selectedRepositoryIds'");
            }
            resourceInputs["secretName"] = args ? args.secretName : undefined;
            resourceInputs["selectedRepositoryIds"] = args ? args.selectedRepositoryIds : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ActionsOrganizationSecretRepositories.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ActionsOrganizationSecretRepositories resources.
 */
export interface ActionsOrganizationSecretRepositoriesState {
    /**
     * Name of the existing secret
     */
    secretName?: pulumi.Input<string>;
    /**
     * An array of repository ids that can access the organization secret.
     */
    selectedRepositoryIds?: pulumi.Input<pulumi.Input<number>[]>;
}

/**
 * The set of arguments for constructing a ActionsOrganizationSecretRepositories resource.
 */
export interface ActionsOrganizationSecretRepositoriesArgs {
    /**
     * Name of the existing secret
     */
    secretName: pulumi.Input<string>;
    /**
     * An array of repository ids that can access the organization secret.
     */
    selectedRepositoryIds: pulumi.Input<pulumi.Input<number>[]>;
}
