// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This resource allows you to create and manage cards for GitHub projects.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as github from "@pulumi/github";
 *
 * const project = new github.OrganizationProject("project", {
 *     name: "An Organization Project",
 *     body: "This is an organization project.",
 * });
 * const column = new github.ProjectColumn("column", {
 *     projectId: project.id,
 *     name: "Backlog",
 * });
 * const card = new github.ProjectCard("card", {
 *     columnId: column.columnId,
 *     note: "## Unaccepted 👇",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ### Adding An Issue To A Project
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as github from "@pulumi/github";
 *
 * const test = new github.Repository("test", {
 *     name: "myrepo",
 *     hasProjects: true,
 *     hasIssues: true,
 * });
 * const testIssue = new github.Issue("test", {
 *     repository: test.id,
 *     title: "Test issue title",
 *     body: "Test issue body",
 * });
 * const testRepositoryProject = new github.RepositoryProject("test", {
 *     name: "test",
 *     repository: test.name,
 *     body: "this is a test project",
 * });
 * const testProjectColumn = new github.ProjectColumn("test", {
 *     projectId: testRepositoryProject.id,
 *     name: "Backlog",
 * });
 * const testProjectCard = new github.ProjectCard("test", {
 *     columnId: testProjectColumn.columnId,
 *     contentId: testIssue.issueId,
 *     contentType: "Issue",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * A GitHub Project Card can be imported using its [Card ID](https://developer.github.com/v3/projects/cards/#get-a-project-card):
 *
 * ```sh
 * $ pulumi import github:index/projectCard:ProjectCard card 01234567
 * ```
 */
export class ProjectCard extends pulumi.CustomResource {
    /**
     * Get an existing ProjectCard resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProjectCardState, opts?: pulumi.CustomResourceOptions): ProjectCard {
        return new ProjectCard(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'github:index/projectCard:ProjectCard';

    /**
     * Returns true if the given object is an instance of ProjectCard.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ProjectCard {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ProjectCard.__pulumiType;
    }

    /**
     * The ID of the card.
     */
    public /*out*/ readonly cardId!: pulumi.Output<number>;
    /**
     * The ID of the card.
     */
    public readonly columnId!: pulumi.Output<string>;
    /**
     * `github_issue.issue_id`.
     */
    public readonly contentId!: pulumi.Output<number | undefined>;
    /**
     * Must be either `Issue` or `PullRequest`
     *
     * **Remarks:** You must either set the `note` attribute or both `contentId` and `contentType`.
     * See note example or issue example for more information.
     */
    public readonly contentType!: pulumi.Output<string | undefined>;
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * The note contents of the card. Markdown supported.
     */
    public readonly note!: pulumi.Output<string | undefined>;

    /**
     * Create a ProjectCard resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProjectCardArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProjectCardArgs | ProjectCardState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProjectCardState | undefined;
            resourceInputs["cardId"] = state ? state.cardId : undefined;
            resourceInputs["columnId"] = state ? state.columnId : undefined;
            resourceInputs["contentId"] = state ? state.contentId : undefined;
            resourceInputs["contentType"] = state ? state.contentType : undefined;
            resourceInputs["etag"] = state ? state.etag : undefined;
            resourceInputs["note"] = state ? state.note : undefined;
        } else {
            const args = argsOrState as ProjectCardArgs | undefined;
            if ((!args || args.columnId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'columnId'");
            }
            resourceInputs["columnId"] = args ? args.columnId : undefined;
            resourceInputs["contentId"] = args ? args.contentId : undefined;
            resourceInputs["contentType"] = args ? args.contentType : undefined;
            resourceInputs["note"] = args ? args.note : undefined;
            resourceInputs["cardId"] = undefined /*out*/;
            resourceInputs["etag"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ProjectCard.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ProjectCard resources.
 */
export interface ProjectCardState {
    /**
     * The ID of the card.
     */
    cardId?: pulumi.Input<number>;
    /**
     * The ID of the card.
     */
    columnId?: pulumi.Input<string>;
    /**
     * `github_issue.issue_id`.
     */
    contentId?: pulumi.Input<number>;
    /**
     * Must be either `Issue` or `PullRequest`
     *
     * **Remarks:** You must either set the `note` attribute or both `contentId` and `contentType`.
     * See note example or issue example for more information.
     */
    contentType?: pulumi.Input<string>;
    etag?: pulumi.Input<string>;
    /**
     * The note contents of the card. Markdown supported.
     */
    note?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ProjectCard resource.
 */
export interface ProjectCardArgs {
    /**
     * The ID of the card.
     */
    columnId: pulumi.Input<string>;
    /**
     * `github_issue.issue_id`.
     */
    contentId?: pulumi.Input<number>;
    /**
     * Must be either `Issue` or `PullRequest`
     *
     * **Remarks:** You must either set the `note` attribute or both `contentId` and `contentType`.
     * See note example or issue example for more information.
     */
    contentType?: pulumi.Input<string>;
    /**
     * The note contents of the card. Markdown supported.
     */
    note?: pulumi.Input<string>;
}
