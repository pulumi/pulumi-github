// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This resource allows you to create and manage files within a
 * GitHub repository.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as github from "@pulumi/github";
 *
 * const fooRepository = new github.Repository("fooRepository", {autoInit: true});
 * const fooRepositoryFile = new github.RepositoryFile("fooRepositoryFile", {
 *     repository: fooRepository.name,
 *     branch: "main",
 *     file: ".gitignore",
 *     content: "**&#47;*.tfstate",
 *     commitMessage: "Managed by Terraform",
 *     commitAuthor: "Terraform User",
 *     commitEmail: "terraform@example.com",
 *     overwriteOnCreate: true,
 * });
 * ```
 *
 * ## Import
 *
 * Repository files can be imported using a combination of the `repo` and `file`, e.g.
 *
 * ```sh
 *  $ pulumi import github:index/repositoryFile:RepositoryFile gitignore example/.gitignore
 * ```
 *  To import a file from a branch other than the default branch, append `:` and the branch name, e.g.
 *
 * ```sh
 *  $ pulumi import github:index/repositoryFile:RepositoryFile gitignore example/.gitignore:dev
 * ```
 */
export class RepositoryFile extends pulumi.CustomResource {
    /**
     * Get an existing RepositoryFile resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RepositoryFileState, opts?: pulumi.CustomResourceOptions): RepositoryFile {
        return new RepositoryFile(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'github:index/repositoryFile:RepositoryFile';

    /**
     * Returns true if the given object is an instance of RepositoryFile.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RepositoryFile {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RepositoryFile.__pulumiType;
    }

    /**
     * Git branch (defaults to the repository's default branch).
     * The branch must already exist, it will not be created if it does not already exist.
     */
    public readonly branch!: pulumi.Output<string | undefined>;
    /**
     * Committer author name to use. **NOTE:** GitHub app users may omit author and email information so GitHub can verify commits as the GitHub App. This maybe useful when a branch protection rule requires signed commits.
     */
    public readonly commitAuthor!: pulumi.Output<string | undefined>;
    /**
     * Committer email address to use. **NOTE:** GitHub app users may omit author and email information so GitHub can verify commits as the GitHub App. This may be useful when a branch protection rule requires signed commits.
     */
    public readonly commitEmail!: pulumi.Output<string | undefined>;
    /**
     * Commit message when adding or updating the managed file.
     */
    public readonly commitMessage!: pulumi.Output<string>;
    /**
     * The SHA of the commit that modified the file.
     */
    public /*out*/ readonly commitSha!: pulumi.Output<string>;
    /**
     * The file content.
     */
    public readonly content!: pulumi.Output<string>;
    /**
     * The path of the file to manage.
     */
    public readonly file!: pulumi.Output<string>;
    /**
     * Enable overwriting existing files
     */
    public readonly overwriteOnCreate!: pulumi.Output<boolean | undefined>;
    /**
     * The name of the commit/branch/tag.
     */
    public /*out*/ readonly ref!: pulumi.Output<string>;
    /**
     * The repository to create the file in.
     */
    public readonly repository!: pulumi.Output<string>;
    /**
     * The SHA blob of the file.
     */
    public /*out*/ readonly sha!: pulumi.Output<string>;

    /**
     * Create a RepositoryFile resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RepositoryFileArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RepositoryFileArgs | RepositoryFileState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RepositoryFileState | undefined;
            resourceInputs["branch"] = state ? state.branch : undefined;
            resourceInputs["commitAuthor"] = state ? state.commitAuthor : undefined;
            resourceInputs["commitEmail"] = state ? state.commitEmail : undefined;
            resourceInputs["commitMessage"] = state ? state.commitMessage : undefined;
            resourceInputs["commitSha"] = state ? state.commitSha : undefined;
            resourceInputs["content"] = state ? state.content : undefined;
            resourceInputs["file"] = state ? state.file : undefined;
            resourceInputs["overwriteOnCreate"] = state ? state.overwriteOnCreate : undefined;
            resourceInputs["ref"] = state ? state.ref : undefined;
            resourceInputs["repository"] = state ? state.repository : undefined;
            resourceInputs["sha"] = state ? state.sha : undefined;
        } else {
            const args = argsOrState as RepositoryFileArgs | undefined;
            if ((!args || args.content === undefined) && !opts.urn) {
                throw new Error("Missing required property 'content'");
            }
            if ((!args || args.file === undefined) && !opts.urn) {
                throw new Error("Missing required property 'file'");
            }
            if ((!args || args.repository === undefined) && !opts.urn) {
                throw new Error("Missing required property 'repository'");
            }
            resourceInputs["branch"] = args ? args.branch : undefined;
            resourceInputs["commitAuthor"] = args ? args.commitAuthor : undefined;
            resourceInputs["commitEmail"] = args ? args.commitEmail : undefined;
            resourceInputs["commitMessage"] = args ? args.commitMessage : undefined;
            resourceInputs["content"] = args ? args.content : undefined;
            resourceInputs["file"] = args ? args.file : undefined;
            resourceInputs["overwriteOnCreate"] = args ? args.overwriteOnCreate : undefined;
            resourceInputs["repository"] = args ? args.repository : undefined;
            resourceInputs["commitSha"] = undefined /*out*/;
            resourceInputs["ref"] = undefined /*out*/;
            resourceInputs["sha"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RepositoryFile.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RepositoryFile resources.
 */
export interface RepositoryFileState {
    /**
     * Git branch (defaults to the repository's default branch).
     * The branch must already exist, it will not be created if it does not already exist.
     */
    branch?: pulumi.Input<string>;
    /**
     * Committer author name to use. **NOTE:** GitHub app users may omit author and email information so GitHub can verify commits as the GitHub App. This maybe useful when a branch protection rule requires signed commits.
     */
    commitAuthor?: pulumi.Input<string>;
    /**
     * Committer email address to use. **NOTE:** GitHub app users may omit author and email information so GitHub can verify commits as the GitHub App. This may be useful when a branch protection rule requires signed commits.
     */
    commitEmail?: pulumi.Input<string>;
    /**
     * Commit message when adding or updating the managed file.
     */
    commitMessage?: pulumi.Input<string>;
    /**
     * The SHA of the commit that modified the file.
     */
    commitSha?: pulumi.Input<string>;
    /**
     * The file content.
     */
    content?: pulumi.Input<string>;
    /**
     * The path of the file to manage.
     */
    file?: pulumi.Input<string>;
    /**
     * Enable overwriting existing files
     */
    overwriteOnCreate?: pulumi.Input<boolean>;
    /**
     * The name of the commit/branch/tag.
     */
    ref?: pulumi.Input<string>;
    /**
     * The repository to create the file in.
     */
    repository?: pulumi.Input<string>;
    /**
     * The SHA blob of the file.
     */
    sha?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RepositoryFile resource.
 */
export interface RepositoryFileArgs {
    /**
     * Git branch (defaults to the repository's default branch).
     * The branch must already exist, it will not be created if it does not already exist.
     */
    branch?: pulumi.Input<string>;
    /**
     * Committer author name to use. **NOTE:** GitHub app users may omit author and email information so GitHub can verify commits as the GitHub App. This maybe useful when a branch protection rule requires signed commits.
     */
    commitAuthor?: pulumi.Input<string>;
    /**
     * Committer email address to use. **NOTE:** GitHub app users may omit author and email information so GitHub can verify commits as the GitHub App. This may be useful when a branch protection rule requires signed commits.
     */
    commitEmail?: pulumi.Input<string>;
    /**
     * Commit message when adding or updating the managed file.
     */
    commitMessage?: pulumi.Input<string>;
    /**
     * The file content.
     */
    content: pulumi.Input<string>;
    /**
     * The path of the file to manage.
     */
    file: pulumi.Input<string>;
    /**
     * Enable overwriting existing files
     */
    overwriteOnCreate?: pulumi.Input<boolean>;
    /**
     * The repository to create the file in.
     */
    repository: pulumi.Input<string>;
}
