// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * This resource allows you to create and manage repositories within your
 * GitHub organization or personal account.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as github from "@pulumi/github";
 *
 * const example = new github.Repository("example", {
 *     description: "My awesome codebase",
 *     template: {
 *         owner: "github",
 *         repository: "terraform-module-template",
 *     },
 *     visibility: "public",
 * });
 * ```
 * ### With GitHub Pages Enabled
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as github from "@pulumi/github";
 *
 * const example = new github.Repository("example", {
 *     description: "My awesome web page",
 *     pages: {
 *         source: {
 *             branch: "master",
 *             path: "/docs",
 *         },
 *     },
 *     private: false,
 * });
 * ```
 *
 * ## Import
 *
 * Repositories can be imported using the `name`, e.g.
 *
 * ```sh
 *  $ pulumi import github:index/repository:Repository terraform terraform
 * ```
 */
export class Repository extends pulumi.CustomResource {
    /**
     * Get an existing Repository resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RepositoryState, opts?: pulumi.CustomResourceOptions): Repository {
        return new Repository(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'github:index/repository:Repository';

    /**
     * Returns true if the given object is an instance of Repository.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Repository {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Repository.__pulumiType;
    }

    /**
     * Set to `false` to disable merge commits on the repository.
     */
    public readonly allowMergeCommit!: pulumi.Output<boolean | undefined>;
    /**
     * Set to `false` to disable rebase merges on the repository.
     */
    public readonly allowRebaseMerge!: pulumi.Output<boolean | undefined>;
    /**
     * Set to `false` to disable squash merges on the repository.
     */
    public readonly allowSquashMerge!: pulumi.Output<boolean | undefined>;
    /**
     * Set to `true` to archive the repository instead of deleting on destroy.
     */
    public readonly archiveOnDestroy!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies if the repository should be archived. Defaults to `false`. **NOTE** Currently, the API does not support unarchiving.
     */
    public readonly archived!: pulumi.Output<boolean | undefined>;
    /**
     * Set to `true` to produce an initial commit in the repository.
     */
    public readonly autoInit!: pulumi.Output<boolean | undefined>;
    /**
     * (Deprecated: Use `github.BranchDefault` resource instead) The name of the default branch of the repository. **NOTE:** This can only be set after a repository has already been created,
     * and after a correct reference has been created for the target branch inside the repository. This means a user will have to omit this parameter from the
     * initial repository creation and create the target branch inside of the repository prior to setting this attribute.
     *
     * @deprecated Use the github_branch_default resource instead
     */
    public readonly defaultBranch!: pulumi.Output<string>;
    /**
     * Automatically delete head branch after a pull request is merged. Defaults to `false`.
     */
    public readonly deleteBranchOnMerge!: pulumi.Output<boolean | undefined>;
    /**
     * A description of the repository.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * A string of the form "orgname/reponame".
     */
    public /*out*/ readonly fullName!: pulumi.Output<string>;
    /**
     * URL that can be provided to `git clone` to clone the repository anonymously via the git protocol.
     */
    public /*out*/ readonly gitCloneUrl!: pulumi.Output<string>;
    /**
     * Use the [name of the template](https://github.com/github/gitignore) without the extension. For example, "Haskell".
     */
    public readonly gitignoreTemplate!: pulumi.Output<string | undefined>;
    /**
     * Set to `true` to enable the (deprecated) downloads features on the repository.
     */
    public readonly hasDownloads!: pulumi.Output<boolean | undefined>;
    /**
     * Set to `true` to enable the GitHub Issues features
     * on the repository.
     */
    public readonly hasIssues!: pulumi.Output<boolean | undefined>;
    /**
     * Set to `true` to enable the GitHub Projects features on the repository. Per the GitHub [documentation](https://developer.github.com/v3/repos/#create) when in an organization that has disabled repository projects it will default to `false` and will otherwise default to `true`. If you specify `true` when it has been disabled it will return an error.
     */
    public readonly hasProjects!: pulumi.Output<boolean | undefined>;
    /**
     * Set to `true` to enable the GitHub Wiki features on
     * the repository.
     */
    public readonly hasWiki!: pulumi.Output<boolean | undefined>;
    /**
     * URL of a page describing the project.
     */
    public readonly homepageUrl!: pulumi.Output<string | undefined>;
    /**
     * The absolute URL (including scheme) of the rendered GitHub Pages site e.g. `https://username.github.io`.
     */
    public /*out*/ readonly htmlUrl!: pulumi.Output<string>;
    /**
     * URL that can be provided to `git clone` to clone the repository via HTTPS.
     */
    public /*out*/ readonly httpCloneUrl!: pulumi.Output<string>;
    /**
     * Set to `true` to tell GitHub that this is a template repository.
     */
    public readonly isTemplate!: pulumi.Output<boolean | undefined>;
    /**
     * Use the [name of the template](https://github.com/github/choosealicense.com/tree/gh-pages/_licenses) without the extension. For example, "mit" or "mpl-2.0".
     */
    public readonly licenseTemplate!: pulumi.Output<string | undefined>;
    /**
     * The name of the repository.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * GraphQL global node id for use with v4 API
     */
    public /*out*/ readonly nodeId!: pulumi.Output<string>;
    /**
     * The repository's GitHub Pages configuration. See GitHub Pages Configuration below for details.
     */
    public readonly pages!: pulumi.Output<outputs.RepositoryPages | undefined>;
    /**
     * Set to `true` to create a private repository.
     * Repositories are created as public (e.g. open source) by default.
     *
     * @deprecated use visibility instead
     */
    public readonly private!: pulumi.Output<boolean>;
    /**
     * GitHub ID for the repository
     */
    public /*out*/ readonly repoId!: pulumi.Output<number>;
    /**
     * URL that can be provided to `git clone` to clone the repository via SSH.
     */
    public /*out*/ readonly sshCloneUrl!: pulumi.Output<string>;
    /**
     * URL that can be provided to `svn checkout` to check out the repository via GitHub's Subversion protocol emulation.
     */
    public /*out*/ readonly svnUrl!: pulumi.Output<string>;
    /**
     * Use a template repository to create this resource. See Template Repositories below for details.
     */
    public readonly template!: pulumi.Output<outputs.RepositoryTemplate | undefined>;
    /**
     * The list of topics of the repository.
     */
    public readonly topics!: pulumi.Output<string[] | undefined>;
    /**
     * Can be `public` or `private`. If your organization is associated with an enterprise account using GitHub Enterprise Cloud or GitHub Enterprise Server 2.20+, visibility can also be `internal`. The `visibility` parameter overrides the `private` parameter.
     */
    public readonly visibility!: pulumi.Output<string>;
    /**
     * Set to `true` to enable security alerts for vulnerable dependencies. Enabling requires alerts to be enabled on the owner level. (Note for importing: GitHub enables the alerts on public repos but disables them on private repos by default.) See [GitHub Documentation](https://help.github.com/en/github/managing-security-vulnerabilities/about-security-alerts-for-vulnerable-dependencies) for details.
     */
    public readonly vulnerabilityAlerts!: pulumi.Output<boolean | undefined>;

    /**
     * Create a Repository resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: RepositoryArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RepositoryArgs | RepositoryState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RepositoryState | undefined;
            inputs["allowMergeCommit"] = state ? state.allowMergeCommit : undefined;
            inputs["allowRebaseMerge"] = state ? state.allowRebaseMerge : undefined;
            inputs["allowSquashMerge"] = state ? state.allowSquashMerge : undefined;
            inputs["archiveOnDestroy"] = state ? state.archiveOnDestroy : undefined;
            inputs["archived"] = state ? state.archived : undefined;
            inputs["autoInit"] = state ? state.autoInit : undefined;
            inputs["defaultBranch"] = state ? state.defaultBranch : undefined;
            inputs["deleteBranchOnMerge"] = state ? state.deleteBranchOnMerge : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["fullName"] = state ? state.fullName : undefined;
            inputs["gitCloneUrl"] = state ? state.gitCloneUrl : undefined;
            inputs["gitignoreTemplate"] = state ? state.gitignoreTemplate : undefined;
            inputs["hasDownloads"] = state ? state.hasDownloads : undefined;
            inputs["hasIssues"] = state ? state.hasIssues : undefined;
            inputs["hasProjects"] = state ? state.hasProjects : undefined;
            inputs["hasWiki"] = state ? state.hasWiki : undefined;
            inputs["homepageUrl"] = state ? state.homepageUrl : undefined;
            inputs["htmlUrl"] = state ? state.htmlUrl : undefined;
            inputs["httpCloneUrl"] = state ? state.httpCloneUrl : undefined;
            inputs["isTemplate"] = state ? state.isTemplate : undefined;
            inputs["licenseTemplate"] = state ? state.licenseTemplate : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["nodeId"] = state ? state.nodeId : undefined;
            inputs["pages"] = state ? state.pages : undefined;
            inputs["private"] = state ? state.private : undefined;
            inputs["repoId"] = state ? state.repoId : undefined;
            inputs["sshCloneUrl"] = state ? state.sshCloneUrl : undefined;
            inputs["svnUrl"] = state ? state.svnUrl : undefined;
            inputs["template"] = state ? state.template : undefined;
            inputs["topics"] = state ? state.topics : undefined;
            inputs["visibility"] = state ? state.visibility : undefined;
            inputs["vulnerabilityAlerts"] = state ? state.vulnerabilityAlerts : undefined;
        } else {
            const args = argsOrState as RepositoryArgs | undefined;
            inputs["allowMergeCommit"] = args ? args.allowMergeCommit : undefined;
            inputs["allowRebaseMerge"] = args ? args.allowRebaseMerge : undefined;
            inputs["allowSquashMerge"] = args ? args.allowSquashMerge : undefined;
            inputs["archiveOnDestroy"] = args ? args.archiveOnDestroy : undefined;
            inputs["archived"] = args ? args.archived : undefined;
            inputs["autoInit"] = args ? args.autoInit : undefined;
            inputs["defaultBranch"] = args ? args.defaultBranch : undefined;
            inputs["deleteBranchOnMerge"] = args ? args.deleteBranchOnMerge : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["gitignoreTemplate"] = args ? args.gitignoreTemplate : undefined;
            inputs["hasDownloads"] = args ? args.hasDownloads : undefined;
            inputs["hasIssues"] = args ? args.hasIssues : undefined;
            inputs["hasProjects"] = args ? args.hasProjects : undefined;
            inputs["hasWiki"] = args ? args.hasWiki : undefined;
            inputs["homepageUrl"] = args ? args.homepageUrl : undefined;
            inputs["isTemplate"] = args ? args.isTemplate : undefined;
            inputs["licenseTemplate"] = args ? args.licenseTemplate : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["pages"] = args ? args.pages : undefined;
            inputs["private"] = args ? args.private : undefined;
            inputs["template"] = args ? args.template : undefined;
            inputs["topics"] = args ? args.topics : undefined;
            inputs["visibility"] = args ? args.visibility : undefined;
            inputs["vulnerabilityAlerts"] = args ? args.vulnerabilityAlerts : undefined;
            inputs["etag"] = undefined /*out*/;
            inputs["fullName"] = undefined /*out*/;
            inputs["gitCloneUrl"] = undefined /*out*/;
            inputs["htmlUrl"] = undefined /*out*/;
            inputs["httpCloneUrl"] = undefined /*out*/;
            inputs["nodeId"] = undefined /*out*/;
            inputs["repoId"] = undefined /*out*/;
            inputs["sshCloneUrl"] = undefined /*out*/;
            inputs["svnUrl"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Repository.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Repository resources.
 */
export interface RepositoryState {
    /**
     * Set to `false` to disable merge commits on the repository.
     */
    readonly allowMergeCommit?: pulumi.Input<boolean>;
    /**
     * Set to `false` to disable rebase merges on the repository.
     */
    readonly allowRebaseMerge?: pulumi.Input<boolean>;
    /**
     * Set to `false` to disable squash merges on the repository.
     */
    readonly allowSquashMerge?: pulumi.Input<boolean>;
    /**
     * Set to `true` to archive the repository instead of deleting on destroy.
     */
    readonly archiveOnDestroy?: pulumi.Input<boolean>;
    /**
     * Specifies if the repository should be archived. Defaults to `false`. **NOTE** Currently, the API does not support unarchiving.
     */
    readonly archived?: pulumi.Input<boolean>;
    /**
     * Set to `true` to produce an initial commit in the repository.
     */
    readonly autoInit?: pulumi.Input<boolean>;
    /**
     * (Deprecated: Use `github.BranchDefault` resource instead) The name of the default branch of the repository. **NOTE:** This can only be set after a repository has already been created,
     * and after a correct reference has been created for the target branch inside the repository. This means a user will have to omit this parameter from the
     * initial repository creation and create the target branch inside of the repository prior to setting this attribute.
     *
     * @deprecated Use the github_branch_default resource instead
     */
    readonly defaultBranch?: pulumi.Input<string>;
    /**
     * Automatically delete head branch after a pull request is merged. Defaults to `false`.
     */
    readonly deleteBranchOnMerge?: pulumi.Input<boolean>;
    /**
     * A description of the repository.
     */
    readonly description?: pulumi.Input<string>;
    readonly etag?: pulumi.Input<string>;
    /**
     * A string of the form "orgname/reponame".
     */
    readonly fullName?: pulumi.Input<string>;
    /**
     * URL that can be provided to `git clone` to clone the repository anonymously via the git protocol.
     */
    readonly gitCloneUrl?: pulumi.Input<string>;
    /**
     * Use the [name of the template](https://github.com/github/gitignore) without the extension. For example, "Haskell".
     */
    readonly gitignoreTemplate?: pulumi.Input<string>;
    /**
     * Set to `true` to enable the (deprecated) downloads features on the repository.
     */
    readonly hasDownloads?: pulumi.Input<boolean>;
    /**
     * Set to `true` to enable the GitHub Issues features
     * on the repository.
     */
    readonly hasIssues?: pulumi.Input<boolean>;
    /**
     * Set to `true` to enable the GitHub Projects features on the repository. Per the GitHub [documentation](https://developer.github.com/v3/repos/#create) when in an organization that has disabled repository projects it will default to `false` and will otherwise default to `true`. If you specify `true` when it has been disabled it will return an error.
     */
    readonly hasProjects?: pulumi.Input<boolean>;
    /**
     * Set to `true` to enable the GitHub Wiki features on
     * the repository.
     */
    readonly hasWiki?: pulumi.Input<boolean>;
    /**
     * URL of a page describing the project.
     */
    readonly homepageUrl?: pulumi.Input<string>;
    /**
     * The absolute URL (including scheme) of the rendered GitHub Pages site e.g. `https://username.github.io`.
     */
    readonly htmlUrl?: pulumi.Input<string>;
    /**
     * URL that can be provided to `git clone` to clone the repository via HTTPS.
     */
    readonly httpCloneUrl?: pulumi.Input<string>;
    /**
     * Set to `true` to tell GitHub that this is a template repository.
     */
    readonly isTemplate?: pulumi.Input<boolean>;
    /**
     * Use the [name of the template](https://github.com/github/choosealicense.com/tree/gh-pages/_licenses) without the extension. For example, "mit" or "mpl-2.0".
     */
    readonly licenseTemplate?: pulumi.Input<string>;
    /**
     * The name of the repository.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * GraphQL global node id for use with v4 API
     */
    readonly nodeId?: pulumi.Input<string>;
    /**
     * The repository's GitHub Pages configuration. See GitHub Pages Configuration below for details.
     */
    readonly pages?: pulumi.Input<inputs.RepositoryPages>;
    /**
     * Set to `true` to create a private repository.
     * Repositories are created as public (e.g. open source) by default.
     *
     * @deprecated use visibility instead
     */
    readonly private?: pulumi.Input<boolean>;
    /**
     * GitHub ID for the repository
     */
    readonly repoId?: pulumi.Input<number>;
    /**
     * URL that can be provided to `git clone` to clone the repository via SSH.
     */
    readonly sshCloneUrl?: pulumi.Input<string>;
    /**
     * URL that can be provided to `svn checkout` to check out the repository via GitHub's Subversion protocol emulation.
     */
    readonly svnUrl?: pulumi.Input<string>;
    /**
     * Use a template repository to create this resource. See Template Repositories below for details.
     */
    readonly template?: pulumi.Input<inputs.RepositoryTemplate>;
    /**
     * The list of topics of the repository.
     */
    readonly topics?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Can be `public` or `private`. If your organization is associated with an enterprise account using GitHub Enterprise Cloud or GitHub Enterprise Server 2.20+, visibility can also be `internal`. The `visibility` parameter overrides the `private` parameter.
     */
    readonly visibility?: pulumi.Input<string>;
    /**
     * Set to `true` to enable security alerts for vulnerable dependencies. Enabling requires alerts to be enabled on the owner level. (Note for importing: GitHub enables the alerts on public repos but disables them on private repos by default.) See [GitHub Documentation](https://help.github.com/en/github/managing-security-vulnerabilities/about-security-alerts-for-vulnerable-dependencies) for details.
     */
    readonly vulnerabilityAlerts?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a Repository resource.
 */
export interface RepositoryArgs {
    /**
     * Set to `false` to disable merge commits on the repository.
     */
    readonly allowMergeCommit?: pulumi.Input<boolean>;
    /**
     * Set to `false` to disable rebase merges on the repository.
     */
    readonly allowRebaseMerge?: pulumi.Input<boolean>;
    /**
     * Set to `false` to disable squash merges on the repository.
     */
    readonly allowSquashMerge?: pulumi.Input<boolean>;
    /**
     * Set to `true` to archive the repository instead of deleting on destroy.
     */
    readonly archiveOnDestroy?: pulumi.Input<boolean>;
    /**
     * Specifies if the repository should be archived. Defaults to `false`. **NOTE** Currently, the API does not support unarchiving.
     */
    readonly archived?: pulumi.Input<boolean>;
    /**
     * Set to `true` to produce an initial commit in the repository.
     */
    readonly autoInit?: pulumi.Input<boolean>;
    /**
     * (Deprecated: Use `github.BranchDefault` resource instead) The name of the default branch of the repository. **NOTE:** This can only be set after a repository has already been created,
     * and after a correct reference has been created for the target branch inside the repository. This means a user will have to omit this parameter from the
     * initial repository creation and create the target branch inside of the repository prior to setting this attribute.
     *
     * @deprecated Use the github_branch_default resource instead
     */
    readonly defaultBranch?: pulumi.Input<string>;
    /**
     * Automatically delete head branch after a pull request is merged. Defaults to `false`.
     */
    readonly deleteBranchOnMerge?: pulumi.Input<boolean>;
    /**
     * A description of the repository.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Use the [name of the template](https://github.com/github/gitignore) without the extension. For example, "Haskell".
     */
    readonly gitignoreTemplate?: pulumi.Input<string>;
    /**
     * Set to `true` to enable the (deprecated) downloads features on the repository.
     */
    readonly hasDownloads?: pulumi.Input<boolean>;
    /**
     * Set to `true` to enable the GitHub Issues features
     * on the repository.
     */
    readonly hasIssues?: pulumi.Input<boolean>;
    /**
     * Set to `true` to enable the GitHub Projects features on the repository. Per the GitHub [documentation](https://developer.github.com/v3/repos/#create) when in an organization that has disabled repository projects it will default to `false` and will otherwise default to `true`. If you specify `true` when it has been disabled it will return an error.
     */
    readonly hasProjects?: pulumi.Input<boolean>;
    /**
     * Set to `true` to enable the GitHub Wiki features on
     * the repository.
     */
    readonly hasWiki?: pulumi.Input<boolean>;
    /**
     * URL of a page describing the project.
     */
    readonly homepageUrl?: pulumi.Input<string>;
    /**
     * Set to `true` to tell GitHub that this is a template repository.
     */
    readonly isTemplate?: pulumi.Input<boolean>;
    /**
     * Use the [name of the template](https://github.com/github/choosealicense.com/tree/gh-pages/_licenses) without the extension. For example, "mit" or "mpl-2.0".
     */
    readonly licenseTemplate?: pulumi.Input<string>;
    /**
     * The name of the repository.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The repository's GitHub Pages configuration. See GitHub Pages Configuration below for details.
     */
    readonly pages?: pulumi.Input<inputs.RepositoryPages>;
    /**
     * Set to `true` to create a private repository.
     * Repositories are created as public (e.g. open source) by default.
     *
     * @deprecated use visibility instead
     */
    readonly private?: pulumi.Input<boolean>;
    /**
     * Use a template repository to create this resource. See Template Repositories below for details.
     */
    readonly template?: pulumi.Input<inputs.RepositoryTemplate>;
    /**
     * The list of topics of the repository.
     */
    readonly topics?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Can be `public` or `private`. If your organization is associated with an enterprise account using GitHub Enterprise Cloud or GitHub Enterprise Server 2.20+, visibility can also be `internal`. The `visibility` parameter overrides the `private` parameter.
     */
    readonly visibility?: pulumi.Input<string>;
    /**
     * Set to `true` to enable security alerts for vulnerable dependencies. Enabling requires alerts to be enabled on the owner level. (Note for importing: GitHub enables the alerts on public repos but disables them on private repos by default.) See [GitHub Documentation](https://help.github.com/en/github/managing-security-vulnerabilities/about-security-alerts-for-vulnerable-dependencies) for details.
     */
    readonly vulnerabilityAlerts?: pulumi.Input<boolean>;
}
