// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This resource allows you to create and manage GitHub Actions variables within your GitHub organization.
 * You must have write access to a repository to use this resource.
 *
 * ## Import
 *
 * This resource can be imported using an ID made up of the variable name:
 *
 * ```sh
 *  $ pulumi import github:index/actionsOrganizationVariable:ActionsOrganizationVariable test_variable test_variable_name
 * ```
 */
export class ActionsOrganizationVariable extends pulumi.CustomResource {
    /**
     * Get an existing ActionsOrganizationVariable resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ActionsOrganizationVariableState, opts?: pulumi.CustomResourceOptions): ActionsOrganizationVariable {
        return new ActionsOrganizationVariable(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'github:index/actionsOrganizationVariable:ActionsOrganizationVariable';

    /**
     * Returns true if the given object is an instance of ActionsOrganizationVariable.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ActionsOrganizationVariable {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ActionsOrganizationVariable.__pulumiType;
    }

    /**
     * Date of actionsVariable creation.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * An array of repository ids that can access the organization variable.
     */
    public readonly selectedRepositoryIds!: pulumi.Output<number[] | undefined>;
    /**
     * Date of actionsVariable update.
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;
    /**
     * Value of the variable
     */
    public readonly value!: pulumi.Output<string>;
    /**
     * Name of the variable
     */
    public readonly variableName!: pulumi.Output<string>;
    /**
     * Configures the access that repositories have to the organization variable.
     * Must be one of `all`, `private`, `selected`. `selectedRepositoryIds` is required if set to `selected`.
     */
    public readonly visibility!: pulumi.Output<string>;

    /**
     * Create a ActionsOrganizationVariable resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ActionsOrganizationVariableArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ActionsOrganizationVariableArgs | ActionsOrganizationVariableState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ActionsOrganizationVariableState | undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["selectedRepositoryIds"] = state ? state.selectedRepositoryIds : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
            resourceInputs["value"] = state ? state.value : undefined;
            resourceInputs["variableName"] = state ? state.variableName : undefined;
            resourceInputs["visibility"] = state ? state.visibility : undefined;
        } else {
            const args = argsOrState as ActionsOrganizationVariableArgs | undefined;
            if ((!args || args.value === undefined) && !opts.urn) {
                throw new Error("Missing required property 'value'");
            }
            if ((!args || args.variableName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'variableName'");
            }
            if ((!args || args.visibility === undefined) && !opts.urn) {
                throw new Error("Missing required property 'visibility'");
            }
            resourceInputs["selectedRepositoryIds"] = args ? args.selectedRepositoryIds : undefined;
            resourceInputs["value"] = args ? args.value : undefined;
            resourceInputs["variableName"] = args ? args.variableName : undefined;
            resourceInputs["visibility"] = args ? args.visibility : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ActionsOrganizationVariable.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ActionsOrganizationVariable resources.
 */
export interface ActionsOrganizationVariableState {
    /**
     * Date of actionsVariable creation.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * An array of repository ids that can access the organization variable.
     */
    selectedRepositoryIds?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * Date of actionsVariable update.
     */
    updatedAt?: pulumi.Input<string>;
    /**
     * Value of the variable
     */
    value?: pulumi.Input<string>;
    /**
     * Name of the variable
     */
    variableName?: pulumi.Input<string>;
    /**
     * Configures the access that repositories have to the organization variable.
     * Must be one of `all`, `private`, `selected`. `selectedRepositoryIds` is required if set to `selected`.
     */
    visibility?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ActionsOrganizationVariable resource.
 */
export interface ActionsOrganizationVariableArgs {
    /**
     * An array of repository ids that can access the organization variable.
     */
    selectedRepositoryIds?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * Value of the variable
     */
    value: pulumi.Input<string>;
    /**
     * Name of the variable
     */
    variableName: pulumi.Input<string>;
    /**
     * Configures the access that repositories have to the organization variable.
     * Must be one of `all`, `private`, `selected`. `selectedRepositoryIds` is required if set to `selected`.
     */
    visibility: pulumi.Input<string>;
}
