// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This resource allows you to create and manage PullRequests for repositories within your GitHub organization or personal account.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as github from "@pulumi/github";
 *
 * const example = new github.RepositoryPullRequest("example", {
 *     baseRef: "main",
 *     baseRepository: "example-repository",
 *     body: "This will change everything",
 *     headRef: "feature-branch",
 *     title: "My newest feature",
 * });
 * ```
 */
export class RepositoryPullRequest extends pulumi.CustomResource {
    /**
     * Get an existing RepositoryPullRequest resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RepositoryPullRequestState, opts?: pulumi.CustomResourceOptions): RepositoryPullRequest {
        return new RepositoryPullRequest(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'github:index/repositoryPullRequest:RepositoryPullRequest';

    /**
     * Returns true if the given object is an instance of RepositoryPullRequest.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RepositoryPullRequest {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RepositoryPullRequest.__pulumiType;
    }

    /**
     * Name of the branch serving as the base of the Pull Request.
     */
    public readonly baseRef!: pulumi.Output<string>;
    /**
     * Name of the base repository to retrieve the Pull Requests from.
     */
    public readonly baseRepository!: pulumi.Output<string>;
    /**
     * Head commit SHA of the Pull Request base.
     */
    public /*out*/ readonly baseSha!: pulumi.Output<string>;
    /**
     * Body of the Pull Request.
     */
    public readonly body!: pulumi.Output<string | undefined>;
    /**
     * Indicates Whether this Pull Request is a draft.
     */
    public /*out*/ readonly draft!: pulumi.Output<boolean>;
    /**
     * Name of the branch serving as the head of the Pull Request.
     */
    public readonly headRef!: pulumi.Output<string>;
    /**
     * Head commit SHA of the Pull Request head.
     */
    public /*out*/ readonly headSha!: pulumi.Output<string>;
    /**
     * List of label names set on the Pull Request.
     */
    public /*out*/ readonly labels!: pulumi.Output<string[]>;
    /**
     * Controls whether the base repository maintainers can modify the Pull Request. Default: false.
     */
    public readonly maintainerCanModify!: pulumi.Output<boolean | undefined>;
    /**
     * The number of the Pull Request within the repository.
     */
    public /*out*/ readonly number!: pulumi.Output<number>;
    /**
     * Unix timestamp indicating the Pull Request creation time.
     */
    public /*out*/ readonly openedAt!: pulumi.Output<number>;
    /**
     * GitHub login of the user who opened the Pull Request.
     */
    public /*out*/ readonly openedBy!: pulumi.Output<string>;
    /**
     * Owner of the repository. If not provided, the provider's default owner is used.
     */
    public readonly owner!: pulumi.Output<string | undefined>;
    /**
     * the current Pull Request state - can be "open", "closed" or "merged".
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * The title of the Pull Request.
     */
    public readonly title!: pulumi.Output<string>;
    /**
     * The timestamp of the last Pull Request update.
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<number>;

    /**
     * Create a RepositoryPullRequest resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RepositoryPullRequestArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RepositoryPullRequestArgs | RepositoryPullRequestState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RepositoryPullRequestState | undefined;
            inputs["baseRef"] = state ? state.baseRef : undefined;
            inputs["baseRepository"] = state ? state.baseRepository : undefined;
            inputs["baseSha"] = state ? state.baseSha : undefined;
            inputs["body"] = state ? state.body : undefined;
            inputs["draft"] = state ? state.draft : undefined;
            inputs["headRef"] = state ? state.headRef : undefined;
            inputs["headSha"] = state ? state.headSha : undefined;
            inputs["labels"] = state ? state.labels : undefined;
            inputs["maintainerCanModify"] = state ? state.maintainerCanModify : undefined;
            inputs["number"] = state ? state.number : undefined;
            inputs["openedAt"] = state ? state.openedAt : undefined;
            inputs["openedBy"] = state ? state.openedBy : undefined;
            inputs["owner"] = state ? state.owner : undefined;
            inputs["state"] = state ? state.state : undefined;
            inputs["title"] = state ? state.title : undefined;
            inputs["updatedAt"] = state ? state.updatedAt : undefined;
        } else {
            const args = argsOrState as RepositoryPullRequestArgs | undefined;
            if ((!args || args.baseRef === undefined) && !opts.urn) {
                throw new Error("Missing required property 'baseRef'");
            }
            if ((!args || args.baseRepository === undefined) && !opts.urn) {
                throw new Error("Missing required property 'baseRepository'");
            }
            if ((!args || args.headRef === undefined) && !opts.urn) {
                throw new Error("Missing required property 'headRef'");
            }
            if ((!args || args.title === undefined) && !opts.urn) {
                throw new Error("Missing required property 'title'");
            }
            inputs["baseRef"] = args ? args.baseRef : undefined;
            inputs["baseRepository"] = args ? args.baseRepository : undefined;
            inputs["body"] = args ? args.body : undefined;
            inputs["headRef"] = args ? args.headRef : undefined;
            inputs["maintainerCanModify"] = args ? args.maintainerCanModify : undefined;
            inputs["owner"] = args ? args.owner : undefined;
            inputs["title"] = args ? args.title : undefined;
            inputs["baseSha"] = undefined /*out*/;
            inputs["draft"] = undefined /*out*/;
            inputs["headSha"] = undefined /*out*/;
            inputs["labels"] = undefined /*out*/;
            inputs["number"] = undefined /*out*/;
            inputs["openedAt"] = undefined /*out*/;
            inputs["openedBy"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
            inputs["updatedAt"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(RepositoryPullRequest.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RepositoryPullRequest resources.
 */
export interface RepositoryPullRequestState {
    /**
     * Name of the branch serving as the base of the Pull Request.
     */
    baseRef?: pulumi.Input<string>;
    /**
     * Name of the base repository to retrieve the Pull Requests from.
     */
    baseRepository?: pulumi.Input<string>;
    /**
     * Head commit SHA of the Pull Request base.
     */
    baseSha?: pulumi.Input<string>;
    /**
     * Body of the Pull Request.
     */
    body?: pulumi.Input<string>;
    /**
     * Indicates Whether this Pull Request is a draft.
     */
    draft?: pulumi.Input<boolean>;
    /**
     * Name of the branch serving as the head of the Pull Request.
     */
    headRef?: pulumi.Input<string>;
    /**
     * Head commit SHA of the Pull Request head.
     */
    headSha?: pulumi.Input<string>;
    /**
     * List of label names set on the Pull Request.
     */
    labels?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Controls whether the base repository maintainers can modify the Pull Request. Default: false.
     */
    maintainerCanModify?: pulumi.Input<boolean>;
    /**
     * The number of the Pull Request within the repository.
     */
    number?: pulumi.Input<number>;
    /**
     * Unix timestamp indicating the Pull Request creation time.
     */
    openedAt?: pulumi.Input<number>;
    /**
     * GitHub login of the user who opened the Pull Request.
     */
    openedBy?: pulumi.Input<string>;
    /**
     * Owner of the repository. If not provided, the provider's default owner is used.
     */
    owner?: pulumi.Input<string>;
    /**
     * the current Pull Request state - can be "open", "closed" or "merged".
     */
    state?: pulumi.Input<string>;
    /**
     * The title of the Pull Request.
     */
    title?: pulumi.Input<string>;
    /**
     * The timestamp of the last Pull Request update.
     */
    updatedAt?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a RepositoryPullRequest resource.
 */
export interface RepositoryPullRequestArgs {
    /**
     * Name of the branch serving as the base of the Pull Request.
     */
    baseRef: pulumi.Input<string>;
    /**
     * Name of the base repository to retrieve the Pull Requests from.
     */
    baseRepository: pulumi.Input<string>;
    /**
     * Body of the Pull Request.
     */
    body?: pulumi.Input<string>;
    /**
     * Name of the branch serving as the head of the Pull Request.
     */
    headRef: pulumi.Input<string>;
    /**
     * Controls whether the base repository maintainers can modify the Pull Request. Default: false.
     */
    maintainerCanModify?: pulumi.Input<boolean>;
    /**
     * Owner of the repository. If not provided, the provider's default owner is used.
     */
    owner?: pulumi.Input<string>;
    /**
     * The title of the Pull Request.
     */
    title: pulumi.Input<string>;
}
