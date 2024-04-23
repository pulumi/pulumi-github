// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This resource allows you to create and manage GitHub Actions runner groups within your GitHub enterprise organizations.
 * You must have admin access to an organization to use this resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as github from "@pulumi/github";
 *
 * const example = new github.Repository("example", {name: "my-repository"});
 * const exampleActionsRunnerGroup = new github.ActionsRunnerGroup("example", {
 *     name: example.name,
 *     visibility: "selected",
 *     selectedRepositoryIds: [example.repoId],
 * });
 * ```
 *
 * ## Import
 *
 * This resource can be imported using the ID of the runner group:
 *
 * ```sh
 * $ pulumi import github:index/actionsRunnerGroup:ActionsRunnerGroup test 7
 * ```
 */
export class ActionsRunnerGroup extends pulumi.CustomResource {
    /**
     * Get an existing ActionsRunnerGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ActionsRunnerGroupState, opts?: pulumi.CustomResourceOptions): ActionsRunnerGroup {
        return new ActionsRunnerGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'github:index/actionsRunnerGroup:ActionsRunnerGroup';

    /**
     * Returns true if the given object is an instance of ActionsRunnerGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ActionsRunnerGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ActionsRunnerGroup.__pulumiType;
    }

    /**
     * Whether public repositories can be added to the runner group. Defaults to false.
     */
    public readonly allowsPublicRepositories!: pulumi.Output<boolean | undefined>;
    /**
     * Whether this is the default runner group
     */
    public /*out*/ readonly default!: pulumi.Output<boolean>;
    /**
     * An etag representing the runner group object
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * Whether the runner group is inherited from the enterprise level
     */
    public /*out*/ readonly inherited!: pulumi.Output<boolean>;
    /**
     * Name of the runner group
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * If true, the runner group will be restricted to running only the workflows specified in the selectedWorkflows array. Defaults to false.
     */
    public readonly restrictedToWorkflows!: pulumi.Output<boolean | undefined>;
    /**
     * The GitHub API URL for the runner group's runners
     */
    public /*out*/ readonly runnersUrl!: pulumi.Output<string>;
    /**
     * GitHub API URL for the runner group's repositories
     */
    public /*out*/ readonly selectedRepositoriesUrl!: pulumi.Output<string>;
    /**
     * IDs of the repositories which should be added to the runner group
     */
    public readonly selectedRepositoryIds!: pulumi.Output<number[] | undefined>;
    /**
     * List of workflows the runner group should be allowed to run. This setting will be ignored unless restrictedToWorkflows is set to true.
     */
    public readonly selectedWorkflows!: pulumi.Output<string[] | undefined>;
    /**
     * Visibility of a runner group. Whether the runner group can include `all`, `selected`, or `private` repositories. A value of `private` is not currently supported due to limitations in the GitHub API.
     */
    public readonly visibility!: pulumi.Output<string>;

    /**
     * Create a ActionsRunnerGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ActionsRunnerGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ActionsRunnerGroupArgs | ActionsRunnerGroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ActionsRunnerGroupState | undefined;
            resourceInputs["allowsPublicRepositories"] = state ? state.allowsPublicRepositories : undefined;
            resourceInputs["default"] = state ? state.default : undefined;
            resourceInputs["etag"] = state ? state.etag : undefined;
            resourceInputs["inherited"] = state ? state.inherited : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["restrictedToWorkflows"] = state ? state.restrictedToWorkflows : undefined;
            resourceInputs["runnersUrl"] = state ? state.runnersUrl : undefined;
            resourceInputs["selectedRepositoriesUrl"] = state ? state.selectedRepositoriesUrl : undefined;
            resourceInputs["selectedRepositoryIds"] = state ? state.selectedRepositoryIds : undefined;
            resourceInputs["selectedWorkflows"] = state ? state.selectedWorkflows : undefined;
            resourceInputs["visibility"] = state ? state.visibility : undefined;
        } else {
            const args = argsOrState as ActionsRunnerGroupArgs | undefined;
            if ((!args || args.visibility === undefined) && !opts.urn) {
                throw new Error("Missing required property 'visibility'");
            }
            resourceInputs["allowsPublicRepositories"] = args ? args.allowsPublicRepositories : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["restrictedToWorkflows"] = args ? args.restrictedToWorkflows : undefined;
            resourceInputs["selectedRepositoryIds"] = args ? args.selectedRepositoryIds : undefined;
            resourceInputs["selectedWorkflows"] = args ? args.selectedWorkflows : undefined;
            resourceInputs["visibility"] = args ? args.visibility : undefined;
            resourceInputs["default"] = undefined /*out*/;
            resourceInputs["etag"] = undefined /*out*/;
            resourceInputs["inherited"] = undefined /*out*/;
            resourceInputs["runnersUrl"] = undefined /*out*/;
            resourceInputs["selectedRepositoriesUrl"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ActionsRunnerGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ActionsRunnerGroup resources.
 */
export interface ActionsRunnerGroupState {
    /**
     * Whether public repositories can be added to the runner group. Defaults to false.
     */
    allowsPublicRepositories?: pulumi.Input<boolean>;
    /**
     * Whether this is the default runner group
     */
    default?: pulumi.Input<boolean>;
    /**
     * An etag representing the runner group object
     */
    etag?: pulumi.Input<string>;
    /**
     * Whether the runner group is inherited from the enterprise level
     */
    inherited?: pulumi.Input<boolean>;
    /**
     * Name of the runner group
     */
    name?: pulumi.Input<string>;
    /**
     * If true, the runner group will be restricted to running only the workflows specified in the selectedWorkflows array. Defaults to false.
     */
    restrictedToWorkflows?: pulumi.Input<boolean>;
    /**
     * The GitHub API URL for the runner group's runners
     */
    runnersUrl?: pulumi.Input<string>;
    /**
     * GitHub API URL for the runner group's repositories
     */
    selectedRepositoriesUrl?: pulumi.Input<string>;
    /**
     * IDs of the repositories which should be added to the runner group
     */
    selectedRepositoryIds?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * List of workflows the runner group should be allowed to run. This setting will be ignored unless restrictedToWorkflows is set to true.
     */
    selectedWorkflows?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Visibility of a runner group. Whether the runner group can include `all`, `selected`, or `private` repositories. A value of `private` is not currently supported due to limitations in the GitHub API.
     */
    visibility?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ActionsRunnerGroup resource.
 */
export interface ActionsRunnerGroupArgs {
    /**
     * Whether public repositories can be added to the runner group. Defaults to false.
     */
    allowsPublicRepositories?: pulumi.Input<boolean>;
    /**
     * Name of the runner group
     */
    name?: pulumi.Input<string>;
    /**
     * If true, the runner group will be restricted to running only the workflows specified in the selectedWorkflows array. Defaults to false.
     */
    restrictedToWorkflows?: pulumi.Input<boolean>;
    /**
     * IDs of the repositories which should be added to the runner group
     */
    selectedRepositoryIds?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * List of workflows the runner group should be allowed to run. This setting will be ignored unless restrictedToWorkflows is set to true.
     */
    selectedWorkflows?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Visibility of a runner group. Whether the runner group can include `all`, `selected`, or `private` repositories. A value of `private` is not currently supported due to limitations in the GitHub API.
     */
    visibility: pulumi.Input<string>;
}
