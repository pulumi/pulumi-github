// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
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
 * $ pulumi import github:index/codespacesUserSecret:CodespacesUserSecret test_secret test_secret_name
 * ```
 *
 * NOTE: the implementation is limited in that it won't fetch the value of the
 *
 * `plaintext_value` or `encrypted_value` fields when importing. You may need to ignore changes for these as a workaround.
 */
export class CodespacesUserSecret extends pulumi.CustomResource {
    /**
     * Get an existing CodespacesUserSecret resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CodespacesUserSecretState, opts?: pulumi.CustomResourceOptions): CodespacesUserSecret {
        return new CodespacesUserSecret(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'github:index/codespacesUserSecret:CodespacesUserSecret';

    /**
     * Returns true if the given object is an instance of CodespacesUserSecret.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CodespacesUserSecret {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CodespacesUserSecret.__pulumiType;
    }

    /**
     * Date of codespacesSecret creation.
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
     * An array of repository ids that can access the user secret.
     */
    public readonly selectedRepositoryIds!: pulumi.Output<number[] | undefined>;
    /**
     * Date of codespacesSecret update.
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;

    /**
     * Create a CodespacesUserSecret resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CodespacesUserSecretArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CodespacesUserSecretArgs | CodespacesUserSecretState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CodespacesUserSecretState | undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["encryptedValue"] = state ? state.encryptedValue : undefined;
            resourceInputs["plaintextValue"] = state ? state.plaintextValue : undefined;
            resourceInputs["secretName"] = state ? state.secretName : undefined;
            resourceInputs["selectedRepositoryIds"] = state ? state.selectedRepositoryIds : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
        } else {
            const args = argsOrState as CodespacesUserSecretArgs | undefined;
            if ((!args || args.secretName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'secretName'");
            }
            resourceInputs["encryptedValue"] = args?.encryptedValue ? pulumi.secret(args.encryptedValue) : undefined;
            resourceInputs["plaintextValue"] = args?.plaintextValue ? pulumi.secret(args.plaintextValue) : undefined;
            resourceInputs["secretName"] = args ? args.secretName : undefined;
            resourceInputs["selectedRepositoryIds"] = args ? args.selectedRepositoryIds : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["encryptedValue", "plaintextValue"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(CodespacesUserSecret.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CodespacesUserSecret resources.
 */
export interface CodespacesUserSecretState {
    /**
     * Date of codespacesSecret creation.
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
     * An array of repository ids that can access the user secret.
     */
    selectedRepositoryIds?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * Date of codespacesSecret update.
     */
    updatedAt?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a CodespacesUserSecret resource.
 */
export interface CodespacesUserSecretArgs {
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
     * An array of repository ids that can access the user secret.
     */
    selectedRepositoryIds?: pulumi.Input<pulumi.Input<number>[]>;
}
