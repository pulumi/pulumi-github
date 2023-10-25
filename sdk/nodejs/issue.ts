// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Provides a GitHub issue resource.
 *
 * This resource allows you to create and manage issue within your
 * GitHub repository.
 *
 * ## Import
 *
 * GitHub Issues can be imported using an ID made up of `repository:number`, e.g.
 *
 * ```sh
 *  $ pulumi import github:index/issue:Issue issue_15 myrepo:15
 * ```
 */
export class Issue extends pulumi.CustomResource {
    /**
     * Get an existing Issue resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IssueState, opts?: pulumi.CustomResourceOptions): Issue {
        return new Issue(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'github:index/issue:Issue';

    /**
     * Returns true if the given object is an instance of Issue.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Issue {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Issue.__pulumiType;
    }

    /**
     * List of Logins to assign the to the issue
     */
    public readonly assignees!: pulumi.Output<string[] | undefined>;
    /**
     * Body of the issue
     */
    public readonly body!: pulumi.Output<string | undefined>;
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * (Computed) - The issue id
     */
    public /*out*/ readonly issueId!: pulumi.Output<number>;
    /**
     * List of labels to attach to the issue
     */
    public readonly labels!: pulumi.Output<string[] | undefined>;
    /**
     * Milestone number to assign to the issue
     */
    public readonly milestoneNumber!: pulumi.Output<number | undefined>;
    /**
     * (Computed) - The issue number
     */
    public /*out*/ readonly number!: pulumi.Output<number>;
    /**
     * The GitHub repository name
     */
    public readonly repository!: pulumi.Output<string>;
    /**
     * Title of the issue
     */
    public readonly title!: pulumi.Output<string>;

    /**
     * Create a Issue resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IssueArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IssueArgs | IssueState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as IssueState | undefined;
            resourceInputs["assignees"] = state ? state.assignees : undefined;
            resourceInputs["body"] = state ? state.body : undefined;
            resourceInputs["etag"] = state ? state.etag : undefined;
            resourceInputs["issueId"] = state ? state.issueId : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["milestoneNumber"] = state ? state.milestoneNumber : undefined;
            resourceInputs["number"] = state ? state.number : undefined;
            resourceInputs["repository"] = state ? state.repository : undefined;
            resourceInputs["title"] = state ? state.title : undefined;
        } else {
            const args = argsOrState as IssueArgs | undefined;
            if ((!args || args.repository === undefined) && !opts.urn) {
                throw new Error("Missing required property 'repository'");
            }
            if ((!args || args.title === undefined) && !opts.urn) {
                throw new Error("Missing required property 'title'");
            }
            resourceInputs["assignees"] = args ? args.assignees : undefined;
            resourceInputs["body"] = args ? args.body : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["milestoneNumber"] = args ? args.milestoneNumber : undefined;
            resourceInputs["repository"] = args ? args.repository : undefined;
            resourceInputs["title"] = args ? args.title : undefined;
            resourceInputs["etag"] = undefined /*out*/;
            resourceInputs["issueId"] = undefined /*out*/;
            resourceInputs["number"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Issue.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Issue resources.
 */
export interface IssueState {
    /**
     * List of Logins to assign the to the issue
     */
    assignees?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Body of the issue
     */
    body?: pulumi.Input<string>;
    etag?: pulumi.Input<string>;
    /**
     * (Computed) - The issue id
     */
    issueId?: pulumi.Input<number>;
    /**
     * List of labels to attach to the issue
     */
    labels?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Milestone number to assign to the issue
     */
    milestoneNumber?: pulumi.Input<number>;
    /**
     * (Computed) - The issue number
     */
    number?: pulumi.Input<number>;
    /**
     * The GitHub repository name
     */
    repository?: pulumi.Input<string>;
    /**
     * Title of the issue
     */
    title?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Issue resource.
 */
export interface IssueArgs {
    /**
     * List of Logins to assign the to the issue
     */
    assignees?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Body of the issue
     */
    body?: pulumi.Input<string>;
    /**
     * List of labels to attach to the issue
     */
    labels?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Milestone number to assign to the issue
     */
    milestoneNumber?: pulumi.Input<number>;
    /**
     * The GitHub repository name
     */
    repository: pulumi.Input<string>;
    /**
     * Title of the issue
     */
    title: pulumi.Input<string>;
}
