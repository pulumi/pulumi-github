// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This resource allows you to create and manage GitHub Actions variables within your GitHub repository environments.
 * You must have write access to a repository to use this resource.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as github from "@pulumi/github";
 *
 * const exampleVariable = new github.ActionsEnvironmentVariable("example_variable", {
 *     environment: "example_environment",
 *     variableName: "example_variable_name",
 *     value: "example_variable_value",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as github from "@pulumi/github";
 *
 * const repo = github.getRepository({
 *     fullName: "my-org/repo",
 * });
 * const repoEnvironment = new github.RepositoryEnvironment("repo_environment", {
 *     repository: repo.then(repo => repo.name),
 *     environment: "example_environment",
 * });
 * const exampleVariable = new github.ActionsEnvironmentVariable("example_variable", {
 *     repository: repo.then(repo => repo.name),
 *     environment: repoEnvironment.environment,
 *     variableName: "example_variable_name",
 *     value: "example_variable_value",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * This resource can be imported using an ID made up of the repository name, environment name, and variable name:
 *
 * ```sh
 * $ pulumi import github:index/actionsEnvironmentVariable:ActionsEnvironmentVariable test_variable myrepo:myenv:myvariable
 * ```
 */
export class ActionsEnvironmentVariable extends pulumi.CustomResource {
    /**
     * Get an existing ActionsEnvironmentVariable resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ActionsEnvironmentVariableState, opts?: pulumi.CustomResourceOptions): ActionsEnvironmentVariable {
        return new ActionsEnvironmentVariable(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'github:index/actionsEnvironmentVariable:ActionsEnvironmentVariable';

    /**
     * Returns true if the given object is an instance of ActionsEnvironmentVariable.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ActionsEnvironmentVariable {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ActionsEnvironmentVariable.__pulumiType;
    }

    /**
     * Date of actionsEnvironmentSecret creation.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * Name of the environment.
     */
    public readonly environment!: pulumi.Output<string>;
    /**
     * Name of the repository.
     */
    public readonly repository!: pulumi.Output<string>;
    /**
     * Date of actionsEnvironmentSecret update.
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;
    /**
     * Value of the variable
     */
    public readonly value!: pulumi.Output<string>;
    /**
     * Name of the variable.
     */
    public readonly variableName!: pulumi.Output<string>;

    /**
     * Create a ActionsEnvironmentVariable resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ActionsEnvironmentVariableArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ActionsEnvironmentVariableArgs | ActionsEnvironmentVariableState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ActionsEnvironmentVariableState | undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["environment"] = state ? state.environment : undefined;
            resourceInputs["repository"] = state ? state.repository : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
            resourceInputs["value"] = state ? state.value : undefined;
            resourceInputs["variableName"] = state ? state.variableName : undefined;
        } else {
            const args = argsOrState as ActionsEnvironmentVariableArgs | undefined;
            if ((!args || args.environment === undefined) && !opts.urn) {
                throw new Error("Missing required property 'environment'");
            }
            if ((!args || args.repository === undefined) && !opts.urn) {
                throw new Error("Missing required property 'repository'");
            }
            if ((!args || args.value === undefined) && !opts.urn) {
                throw new Error("Missing required property 'value'");
            }
            if ((!args || args.variableName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'variableName'");
            }
            resourceInputs["environment"] = args ? args.environment : undefined;
            resourceInputs["repository"] = args ? args.repository : undefined;
            resourceInputs["value"] = args ? args.value : undefined;
            resourceInputs["variableName"] = args ? args.variableName : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ActionsEnvironmentVariable.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ActionsEnvironmentVariable resources.
 */
export interface ActionsEnvironmentVariableState {
    /**
     * Date of actionsEnvironmentSecret creation.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * Name of the environment.
     */
    environment?: pulumi.Input<string>;
    /**
     * Name of the repository.
     */
    repository?: pulumi.Input<string>;
    /**
     * Date of actionsEnvironmentSecret update.
     */
    updatedAt?: pulumi.Input<string>;
    /**
     * Value of the variable
     */
    value?: pulumi.Input<string>;
    /**
     * Name of the variable.
     */
    variableName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ActionsEnvironmentVariable resource.
 */
export interface ActionsEnvironmentVariableArgs {
    /**
     * Name of the environment.
     */
    environment: pulumi.Input<string>;
    /**
     * Name of the repository.
     */
    repository: pulumi.Input<string>;
    /**
     * Value of the variable
     */
    value: pulumi.Input<string>;
    /**
     * Name of the variable.
     */
    variableName: pulumi.Input<string>;
}
