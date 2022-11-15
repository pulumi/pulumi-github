// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class DependabotSecret extends pulumi.CustomResource {
    /**
     * Get an existing DependabotSecret resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DependabotSecretState, opts?: pulumi.CustomResourceOptions): DependabotSecret {
        return new DependabotSecret(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'github:index/dependabotSecret:DependabotSecret';

    /**
     * Returns true if the given object is an instance of DependabotSecret.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DependabotSecret {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DependabotSecret.__pulumiType;
    }

    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    public readonly encryptedValue!: pulumi.Output<string | undefined>;
    public readonly plaintextValue!: pulumi.Output<string | undefined>;
    public readonly repository!: pulumi.Output<string>;
    public readonly secretName!: pulumi.Output<string>;
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;

    /**
     * Create a DependabotSecret resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DependabotSecretArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DependabotSecretArgs | DependabotSecretState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DependabotSecretState | undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["encryptedValue"] = state ? state.encryptedValue : undefined;
            resourceInputs["plaintextValue"] = state ? state.plaintextValue : undefined;
            resourceInputs["repository"] = state ? state.repository : undefined;
            resourceInputs["secretName"] = state ? state.secretName : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
        } else {
            const args = argsOrState as DependabotSecretArgs | undefined;
            if ((!args || args.repository === undefined) && !opts.urn) {
                throw new Error("Missing required property 'repository'");
            }
            if ((!args || args.secretName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'secretName'");
            }
            resourceInputs["encryptedValue"] = args?.encryptedValue ? pulumi.secret(args.encryptedValue) : undefined;
            resourceInputs["plaintextValue"] = args?.plaintextValue ? pulumi.secret(args.plaintextValue) : undefined;
            resourceInputs["repository"] = args ? args.repository : undefined;
            resourceInputs["secretName"] = args ? args.secretName : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["encryptedValue", "plaintextValue"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(DependabotSecret.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DependabotSecret resources.
 */
export interface DependabotSecretState {
    createdAt?: pulumi.Input<string>;
    encryptedValue?: pulumi.Input<string>;
    plaintextValue?: pulumi.Input<string>;
    repository?: pulumi.Input<string>;
    secretName?: pulumi.Input<string>;
    updatedAt?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DependabotSecret resource.
 */
export interface DependabotSecretArgs {
    encryptedValue?: pulumi.Input<string>;
    plaintextValue?: pulumi.Input<string>;
    repository: pulumi.Input<string>;
    secretName: pulumi.Input<string>;
}
