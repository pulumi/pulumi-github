// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This resource allows you to create and manage GitHub Actions variables within your GitHub repositories.
 * You must have write access to a repository to use this resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as github from "@pulumi/github";
 *
 * const exampleVariable = new github.ActionsVariable("example_variable", {
 *     repository: "example_repository",
 *     variableName: "example_variable_name",
 *     value: "example_variable_value",
 * });
 * ```
 *
 * ## Import
 *
 * GitHub Actions variables can be imported using an ID made up of `repository:variable_name`, e.g.
 *
 * ```sh
 * $ pulumi import github:index/actionsVariable:ActionsVariable myvariable myrepo:myvariable
 * ```
 */
export class ActionsVariable extends pulumi.CustomResource {
    /**
     * Get an existing ActionsVariable resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ActionsVariableState, opts?: pulumi.CustomResourceOptions): ActionsVariable {
        return new ActionsVariable(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'github:index/actionsVariable:ActionsVariable';

    /**
     * Returns true if the given object is an instance of ActionsVariable.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ActionsVariable {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ActionsVariable.__pulumiType;
    }

    /**
     * Date of actionsVariable creation.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * Name of the repository
     */
    public readonly repository!: pulumi.Output<string>;
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
     * Create a ActionsVariable resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ActionsVariableArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ActionsVariableArgs | ActionsVariableState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ActionsVariableState | undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["repository"] = state ? state.repository : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
            resourceInputs["value"] = state ? state.value : undefined;
            resourceInputs["variableName"] = state ? state.variableName : undefined;
        } else {
            const args = argsOrState as ActionsVariableArgs | undefined;
            if ((!args || args.repository === undefined) && !opts.urn) {
                throw new Error("Missing required property 'repository'");
            }
            if ((!args || args.value === undefined) && !opts.urn) {
                throw new Error("Missing required property 'value'");
            }
            if ((!args || args.variableName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'variableName'");
            }
            resourceInputs["repository"] = args ? args.repository : undefined;
            resourceInputs["value"] = args ? args.value : undefined;
            resourceInputs["variableName"] = args ? args.variableName : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ActionsVariable.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ActionsVariable resources.
 */
export interface ActionsVariableState {
    /**
     * Date of actionsVariable creation.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * Name of the repository
     */
    repository?: pulumi.Input<string>;
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
}

/**
 * The set of arguments for constructing a ActionsVariable resource.
 */
export interface ActionsVariableArgs {
    /**
     * Name of the repository
     */
    repository: pulumi.Input<string>;
    /**
     * Value of the variable
     */
    value: pulumi.Input<string>;
    /**
     * Name of the variable
     */
    variableName: pulumi.Input<string>;
}
