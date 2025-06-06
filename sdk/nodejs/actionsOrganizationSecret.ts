// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 *
 * ## Import
 *
 * This resource can be imported using an ID made up of the secret name:
 *
 * ```sh
 * $ pulumi import github:index/actionsOrganizationSecret:ActionsOrganizationSecret test_secret test_secret_name
 * ```
 * NOTE: the implementation is limited in that it won't fetch the value of the
 * `plaintext_value` or `encrypted_value` fields when importing. You may need to ignore changes for these as a workaround.
 */
export class ActionsOrganizationSecret extends pulumi.CustomResource {
    /**
     * Get an existing ActionsOrganizationSecret resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ActionsOrganizationSecretState, opts?: pulumi.CustomResourceOptions): ActionsOrganizationSecret {
        return new ActionsOrganizationSecret(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'github:index/actionsOrganizationSecret:ActionsOrganizationSecret';

    /**
     * Returns true if the given object is an instance of ActionsOrganizationSecret.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ActionsOrganizationSecret {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ActionsOrganizationSecret.__pulumiType;
    }

    /**
     * Date of actionsSecret creation.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * Encrypted value of the secret using the GitHub public key in Base64 format.
     */
    public readonly encryptedValue!: pulumi.Output<string | undefined>;
    /**
     * Plaintext value of the secret to be encrypted
     */
    public readonly plaintextValue!: pulumi.Output<string | undefined>;
    /**
     * Name of the secret
     */
    public readonly secretName!: pulumi.Output<string>;
    /**
     * An array of repository ids that can access the organization secret.
     */
    public readonly selectedRepositoryIds!: pulumi.Output<number[] | undefined>;
    /**
     * Date of actionsSecret update.
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;
    /**
     * Configures the access that repositories have to the organization secret.
     * Must be one of `all`, `private`, `selected`. `selectedRepositoryIds` is required if set to `selected`.
     */
    public readonly visibility!: pulumi.Output<string>;

    /**
     * Create a ActionsOrganizationSecret resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ActionsOrganizationSecretArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ActionsOrganizationSecretArgs | ActionsOrganizationSecretState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ActionsOrganizationSecretState | undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["encryptedValue"] = state ? state.encryptedValue : undefined;
            resourceInputs["plaintextValue"] = state ? state.plaintextValue : undefined;
            resourceInputs["secretName"] = state ? state.secretName : undefined;
            resourceInputs["selectedRepositoryIds"] = state ? state.selectedRepositoryIds : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
            resourceInputs["visibility"] = state ? state.visibility : undefined;
        } else {
            const args = argsOrState as ActionsOrganizationSecretArgs | undefined;
            if ((!args || args.secretName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'secretName'");
            }
            if ((!args || args.visibility === undefined) && !opts.urn) {
                throw new Error("Missing required property 'visibility'");
            }
            resourceInputs["encryptedValue"] = args?.encryptedValue ? pulumi.secret(args.encryptedValue) : undefined;
            resourceInputs["plaintextValue"] = args?.plaintextValue ? pulumi.secret(args.plaintextValue) : undefined;
            resourceInputs["secretName"] = args ? args.secretName : undefined;
            resourceInputs["selectedRepositoryIds"] = args ? args.selectedRepositoryIds : undefined;
            resourceInputs["visibility"] = args ? args.visibility : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["encryptedValue", "plaintextValue"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(ActionsOrganizationSecret.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ActionsOrganizationSecret resources.
 */
export interface ActionsOrganizationSecretState {
    /**
     * Date of actionsSecret creation.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * Encrypted value of the secret using the GitHub public key in Base64 format.
     */
    encryptedValue?: pulumi.Input<string>;
    /**
     * Plaintext value of the secret to be encrypted
     */
    plaintextValue?: pulumi.Input<string>;
    /**
     * Name of the secret
     */
    secretName?: pulumi.Input<string>;
    /**
     * An array of repository ids that can access the organization secret.
     */
    selectedRepositoryIds?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * Date of actionsSecret update.
     */
    updatedAt?: pulumi.Input<string>;
    /**
     * Configures the access that repositories have to the organization secret.
     * Must be one of `all`, `private`, `selected`. `selectedRepositoryIds` is required if set to `selected`.
     */
    visibility?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ActionsOrganizationSecret resource.
 */
export interface ActionsOrganizationSecretArgs {
    /**
     * Encrypted value of the secret using the GitHub public key in Base64 format.
     */
    encryptedValue?: pulumi.Input<string>;
    /**
     * Plaintext value of the secret to be encrypted
     */
    plaintextValue?: pulumi.Input<string>;
    /**
     * Name of the secret
     */
    secretName: pulumi.Input<string>;
    /**
     * An array of repository ids that can access the organization secret.
     */
    selectedRepositoryIds?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * Configures the access that repositories have to the organization secret.
     * Must be one of `all`, `private`, `selected`. `selectedRepositoryIds` is required if set to `selected`.
     */
    visibility: pulumi.Input<string>;
}
