// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This resource allows you to create and manage deployment branch policies.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as github from "@pulumi/github";
 *
 * const env = new github.RepositoryEnvironment("env", {
 *     repository: "my_repo",
 *     environment: "my_env",
 *     deploymentBranchPolicy: {
 *         protectedBranches: false,
 *         customBranchPolicies: true,
 *     },
 * });
 * const foo = new github.RepositoryDeploymentBranchPolicy("foo", {
 *     repository: "my_repo",
 *     environmentName: "my_env",
 * }, {
 *     dependsOn: [env],
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 * $ pulumi import github:index/repositoryDeploymentBranchPolicy:RepositoryDeploymentBranchPolicy foo repo:env:id
 * ```
 */
export class RepositoryDeploymentBranchPolicy extends pulumi.CustomResource {
    /**
     * Get an existing RepositoryDeploymentBranchPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RepositoryDeploymentBranchPolicyState, opts?: pulumi.CustomResourceOptions): RepositoryDeploymentBranchPolicy {
        return new RepositoryDeploymentBranchPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'github:index/repositoryDeploymentBranchPolicy:RepositoryDeploymentBranchPolicy';

    /**
     * Returns true if the given object is an instance of RepositoryDeploymentBranchPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RepositoryDeploymentBranchPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RepositoryDeploymentBranchPolicy.__pulumiType;
    }

    /**
     * The name of the environment. This environment must have `deployment_branch_policy.custom_branch_policies` set to true or a 404 error will be thrown.
     */
    public readonly environmentName!: pulumi.Output<string>;
    /**
     * An etag representing the Branch object.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * The name pattern that branches must match in order to deploy to the environment.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The repository to create the policy in.
     */
    public readonly repository!: pulumi.Output<string>;

    /**
     * Create a RepositoryDeploymentBranchPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RepositoryDeploymentBranchPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RepositoryDeploymentBranchPolicyArgs | RepositoryDeploymentBranchPolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RepositoryDeploymentBranchPolicyState | undefined;
            resourceInputs["environmentName"] = state ? state.environmentName : undefined;
            resourceInputs["etag"] = state ? state.etag : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["repository"] = state ? state.repository : undefined;
        } else {
            const args = argsOrState as RepositoryDeploymentBranchPolicyArgs | undefined;
            if ((!args || args.environmentName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'environmentName'");
            }
            if ((!args || args.repository === undefined) && !opts.urn) {
                throw new Error("Missing required property 'repository'");
            }
            resourceInputs["environmentName"] = args ? args.environmentName : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["repository"] = args ? args.repository : undefined;
            resourceInputs["etag"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RepositoryDeploymentBranchPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RepositoryDeploymentBranchPolicy resources.
 */
export interface RepositoryDeploymentBranchPolicyState {
    /**
     * The name of the environment. This environment must have `deployment_branch_policy.custom_branch_policies` set to true or a 404 error will be thrown.
     */
    environmentName?: pulumi.Input<string>;
    /**
     * An etag representing the Branch object.
     */
    etag?: pulumi.Input<string>;
    /**
     * The name pattern that branches must match in order to deploy to the environment.
     */
    name?: pulumi.Input<string>;
    /**
     * The repository to create the policy in.
     */
    repository?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RepositoryDeploymentBranchPolicy resource.
 */
export interface RepositoryDeploymentBranchPolicyArgs {
    /**
     * The name of the environment. This environment must have `deployment_branch_policy.custom_branch_policies` set to true or a 404 error will be thrown.
     */
    environmentName: pulumi.Input<string>;
    /**
     * The name pattern that branches must match in order to deploy to the environment.
     */
    name?: pulumi.Input<string>;
    /**
     * The repository to create the policy in.
     */
    repository: pulumi.Input<string>;
}
