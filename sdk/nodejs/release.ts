// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This resource allows you to create and manage a release in a specific
 * GitHub repository.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as github from "@pulumi/github";
 *
 * const repo = new github.Repository("repo", {
 *     description: "GitHub repo managed by Terraform",
 *     "private": false,
 * });
 * const example = new github.Release("example", {
 *     repository: repo.name,
 *     tagName: "v1.0.0",
 * });
 * ```
 *
 * ### On Non-Default Branch
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as github from "@pulumi/github";
 *
 * const exampleRepository = new github.Repository("exampleRepository", {autoInit: true});
 * const exampleBranch = new github.Branch("exampleBranch", {
 *     repository: exampleRepository.name,
 *     branch: "branch_name",
 *     sourceBranch: exampleRepository.defaultBranch,
 * });
 * const exampleRelease = new github.Release("exampleRelease", {
 *     repository: exampleRepository.name,
 *     tagName: "v1.0.0",
 *     targetCommitish: exampleBranch.branch,
 *     draft: false,
 *     prerelease: false,
 * });
 * ```
 *
 * ## Import
 *
 * This resource can be imported using the `name` of the repository, combined with the `id` of the release, and a `:` character for separating components, e.g.
 *
 * ```sh
 * $ pulumi import github:index/release:Release example repo:12345678
 * ```
 */
export class Release extends pulumi.CustomResource {
    /**
     * Get an existing Release resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ReleaseState, opts?: pulumi.CustomResourceOptions): Release {
        return new Release(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'github:index/release:Release';

    /**
     * Returns true if the given object is an instance of Release.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Release {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Release.__pulumiType;
    }

    /**
     * Text describing the contents of the tag.
     */
    public readonly body!: pulumi.Output<string | undefined>;
    /**
     * If specified, a discussion of the specified category is created and linked to the release. The value must be a category that already exists in the repository. For more information, see [Managing categories for discussions in your repository](https://docs.github.com/discussions/managing-discussions-for-your-community/managing-categories-for-discussions-in-your-repository).
     */
    public readonly discussionCategoryName!: pulumi.Output<string | undefined>;
    /**
     * Set to `false` to create a published release.
     */
    public readonly draft!: pulumi.Output<boolean | undefined>;
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * Set to `true` to automatically generate the name and body for this release. If `name` is specified, the specified `name` will be used; otherwise, a name will be automatically generated. If `body` is specified, the `body` will be pre-pended to the automatically generated notes.
     */
    public readonly generateReleaseNotes!: pulumi.Output<boolean | undefined>;
    /**
     * The name of the release.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Set to `false` to identify the release as a full release.
     */
    public readonly prerelease!: pulumi.Output<boolean | undefined>;
    /**
     * The name of the repository.
     */
    public readonly repository!: pulumi.Output<string>;
    /**
     * The name of the tag.
     */
    public readonly tagName!: pulumi.Output<string>;
    /**
     * The branch name or commit SHA the tag is created from. Defaults to the default branch of the repository.
     */
    public readonly targetCommitish!: pulumi.Output<string | undefined>;

    /**
     * Create a Release resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ReleaseArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ReleaseArgs | ReleaseState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ReleaseState | undefined;
            resourceInputs["body"] = state ? state.body : undefined;
            resourceInputs["discussionCategoryName"] = state ? state.discussionCategoryName : undefined;
            resourceInputs["draft"] = state ? state.draft : undefined;
            resourceInputs["etag"] = state ? state.etag : undefined;
            resourceInputs["generateReleaseNotes"] = state ? state.generateReleaseNotes : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["prerelease"] = state ? state.prerelease : undefined;
            resourceInputs["repository"] = state ? state.repository : undefined;
            resourceInputs["tagName"] = state ? state.tagName : undefined;
            resourceInputs["targetCommitish"] = state ? state.targetCommitish : undefined;
        } else {
            const args = argsOrState as ReleaseArgs | undefined;
            if ((!args || args.repository === undefined) && !opts.urn) {
                throw new Error("Missing required property 'repository'");
            }
            if ((!args || args.tagName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tagName'");
            }
            resourceInputs["body"] = args ? args.body : undefined;
            resourceInputs["discussionCategoryName"] = args ? args.discussionCategoryName : undefined;
            resourceInputs["draft"] = args ? args.draft : undefined;
            resourceInputs["generateReleaseNotes"] = args ? args.generateReleaseNotes : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["prerelease"] = args ? args.prerelease : undefined;
            resourceInputs["repository"] = args ? args.repository : undefined;
            resourceInputs["tagName"] = args ? args.tagName : undefined;
            resourceInputs["targetCommitish"] = args ? args.targetCommitish : undefined;
            resourceInputs["etag"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Release.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Release resources.
 */
export interface ReleaseState {
    /**
     * Text describing the contents of the tag.
     */
    body?: pulumi.Input<string>;
    /**
     * If specified, a discussion of the specified category is created and linked to the release. The value must be a category that already exists in the repository. For more information, see [Managing categories for discussions in your repository](https://docs.github.com/discussions/managing-discussions-for-your-community/managing-categories-for-discussions-in-your-repository).
     */
    discussionCategoryName?: pulumi.Input<string>;
    /**
     * Set to `false` to create a published release.
     */
    draft?: pulumi.Input<boolean>;
    etag?: pulumi.Input<string>;
    /**
     * Set to `true` to automatically generate the name and body for this release. If `name` is specified, the specified `name` will be used; otherwise, a name will be automatically generated. If `body` is specified, the `body` will be pre-pended to the automatically generated notes.
     */
    generateReleaseNotes?: pulumi.Input<boolean>;
    /**
     * The name of the release.
     */
    name?: pulumi.Input<string>;
    /**
     * Set to `false` to identify the release as a full release.
     */
    prerelease?: pulumi.Input<boolean>;
    /**
     * The name of the repository.
     */
    repository?: pulumi.Input<string>;
    /**
     * The name of the tag.
     */
    tagName?: pulumi.Input<string>;
    /**
     * The branch name or commit SHA the tag is created from. Defaults to the default branch of the repository.
     */
    targetCommitish?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Release resource.
 */
export interface ReleaseArgs {
    /**
     * Text describing the contents of the tag.
     */
    body?: pulumi.Input<string>;
    /**
     * If specified, a discussion of the specified category is created and linked to the release. The value must be a category that already exists in the repository. For more information, see [Managing categories for discussions in your repository](https://docs.github.com/discussions/managing-discussions-for-your-community/managing-categories-for-discussions-in-your-repository).
     */
    discussionCategoryName?: pulumi.Input<string>;
    /**
     * Set to `false` to create a published release.
     */
    draft?: pulumi.Input<boolean>;
    /**
     * Set to `true` to automatically generate the name and body for this release. If `name` is specified, the specified `name` will be used; otherwise, a name will be automatically generated. If `body` is specified, the `body` will be pre-pended to the automatically generated notes.
     */
    generateReleaseNotes?: pulumi.Input<boolean>;
    /**
     * The name of the release.
     */
    name?: pulumi.Input<string>;
    /**
     * Set to `false` to identify the release as a full release.
     */
    prerelease?: pulumi.Input<boolean>;
    /**
     * The name of the repository.
     */
    repository: pulumi.Input<string>;
    /**
     * The name of the tag.
     */
    tagName: pulumi.Input<string>;
    /**
     * The branch name or commit SHA the tag is created from. Defaults to the default branch of the repository.
     */
    targetCommitish?: pulumi.Input<string>;
}
