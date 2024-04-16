// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * This resource allows you to create and manage environments for a GitHub repository.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as github from "@pulumi/github";
 *
 * const current = github.getUser({
 *     username: "",
 * });
 * const example = new github.Repository("example", {
 *     name: "A Repository Project",
 *     description: "My awesome codebase",
 * });
 * const exampleRepositoryEnvironment = new github.RepositoryEnvironment("example", {
 *     environment: "example",
 *     repository: example.name,
 *     preventSelfReview: true,
 *     reviewers: [{
 *         users: [current.then(current => current.id)],
 *     }],
 *     deploymentBranchPolicy: {
 *         protectedBranches: true,
 *         customBranchPolicies: false,
 *     },
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * GitHub Repository Environment can be imported using an ID made up of `name` of the repository combined with the `environment` name of the environment, separated by a `:` character, e.g.
 *
 * ```sh
 * $ pulumi import github:index/repositoryEnvironment:RepositoryEnvironment daily terraform:daily
 * ```
 */
export class RepositoryEnvironment extends pulumi.CustomResource {
    /**
     * Get an existing RepositoryEnvironment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RepositoryEnvironmentState, opts?: pulumi.CustomResourceOptions): RepositoryEnvironment {
        return new RepositoryEnvironment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'github:index/repositoryEnvironment:RepositoryEnvironment';

    /**
     * Returns true if the given object is an instance of RepositoryEnvironment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RepositoryEnvironment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RepositoryEnvironment.__pulumiType;
    }

    /**
     * Can repository admins bypass the environment protections.  Defaults to `true`.
     */
    public readonly canAdminsBypass!: pulumi.Output<boolean | undefined>;
    /**
     * The deployment branch policy configuration
     */
    public readonly deploymentBranchPolicy!: pulumi.Output<outputs.RepositoryEnvironmentDeploymentBranchPolicy | undefined>;
    /**
     * The name of the environment.
     */
    public readonly environment!: pulumi.Output<string>;
    /**
     * Whether or not a user who created the job is prevented from approving their own job. Defaults to `false`.
     */
    public readonly preventSelfReview!: pulumi.Output<boolean | undefined>;
    /**
     * The repository of the environment.
     */
    public readonly repository!: pulumi.Output<string>;
    /**
     * The environment reviewers configuration.
     */
    public readonly reviewers!: pulumi.Output<outputs.RepositoryEnvironmentReviewer[] | undefined>;
    /**
     * Amount of time to delay a job after the job is initially triggered.
     */
    public readonly waitTimer!: pulumi.Output<number | undefined>;

    /**
     * Create a RepositoryEnvironment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RepositoryEnvironmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RepositoryEnvironmentArgs | RepositoryEnvironmentState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RepositoryEnvironmentState | undefined;
            resourceInputs["canAdminsBypass"] = state ? state.canAdminsBypass : undefined;
            resourceInputs["deploymentBranchPolicy"] = state ? state.deploymentBranchPolicy : undefined;
            resourceInputs["environment"] = state ? state.environment : undefined;
            resourceInputs["preventSelfReview"] = state ? state.preventSelfReview : undefined;
            resourceInputs["repository"] = state ? state.repository : undefined;
            resourceInputs["reviewers"] = state ? state.reviewers : undefined;
            resourceInputs["waitTimer"] = state ? state.waitTimer : undefined;
        } else {
            const args = argsOrState as RepositoryEnvironmentArgs | undefined;
            if ((!args || args.environment === undefined) && !opts.urn) {
                throw new Error("Missing required property 'environment'");
            }
            if ((!args || args.repository === undefined) && !opts.urn) {
                throw new Error("Missing required property 'repository'");
            }
            resourceInputs["canAdminsBypass"] = args ? args.canAdminsBypass : undefined;
            resourceInputs["deploymentBranchPolicy"] = args ? args.deploymentBranchPolicy : undefined;
            resourceInputs["environment"] = args ? args.environment : undefined;
            resourceInputs["preventSelfReview"] = args ? args.preventSelfReview : undefined;
            resourceInputs["repository"] = args ? args.repository : undefined;
            resourceInputs["reviewers"] = args ? args.reviewers : undefined;
            resourceInputs["waitTimer"] = args ? args.waitTimer : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RepositoryEnvironment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RepositoryEnvironment resources.
 */
export interface RepositoryEnvironmentState {
    /**
     * Can repository admins bypass the environment protections.  Defaults to `true`.
     */
    canAdminsBypass?: pulumi.Input<boolean>;
    /**
     * The deployment branch policy configuration
     */
    deploymentBranchPolicy?: pulumi.Input<inputs.RepositoryEnvironmentDeploymentBranchPolicy>;
    /**
     * The name of the environment.
     */
    environment?: pulumi.Input<string>;
    /**
     * Whether or not a user who created the job is prevented from approving their own job. Defaults to `false`.
     */
    preventSelfReview?: pulumi.Input<boolean>;
    /**
     * The repository of the environment.
     */
    repository?: pulumi.Input<string>;
    /**
     * The environment reviewers configuration.
     */
    reviewers?: pulumi.Input<pulumi.Input<inputs.RepositoryEnvironmentReviewer>[]>;
    /**
     * Amount of time to delay a job after the job is initially triggered.
     */
    waitTimer?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a RepositoryEnvironment resource.
 */
export interface RepositoryEnvironmentArgs {
    /**
     * Can repository admins bypass the environment protections.  Defaults to `true`.
     */
    canAdminsBypass?: pulumi.Input<boolean>;
    /**
     * The deployment branch policy configuration
     */
    deploymentBranchPolicy?: pulumi.Input<inputs.RepositoryEnvironmentDeploymentBranchPolicy>;
    /**
     * The name of the environment.
     */
    environment: pulumi.Input<string>;
    /**
     * Whether or not a user who created the job is prevented from approving their own job. Defaults to `false`.
     */
    preventSelfReview?: pulumi.Input<boolean>;
    /**
     * The repository of the environment.
     */
    repository: pulumi.Input<string>;
    /**
     * The environment reviewers configuration.
     */
    reviewers?: pulumi.Input<pulumi.Input<inputs.RepositoryEnvironmentReviewer>[]>;
    /**
     * Amount of time to delay a job after the job is initially triggered.
     */
    waitTimer?: pulumi.Input<number>;
}
