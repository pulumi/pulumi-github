// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * This resource allows you to create and manage repositories within your
 * GitHub organization or personal account.
 *
 * > Note: When used with GitHub App authentication, even GET requests must have
 * the `contents:write` permission or else the `allowMergeCommit`, `allowRebaseMerge`,
 * and `allowSquashMerge` attributes will be ignored, causing confusing diffs.
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
 *         includeAllBranches: true,
 *         owner: "github",
 *         repository: "terraform-template-module",
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
 *     "private": false,
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
     * Set to `true` to allow auto-merging pull requests on the repository.
     */
    public readonly allowAutoMerge!: pulumi.Output<boolean | undefined>;
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
     * Set to `true` to always suggest updating pull request branches.
     */
    public readonly allowUpdateBranch!: pulumi.Output<boolean | undefined>;
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
     * Set to `true` to enable GitHub Discussions on the repository. Defaults to `false`.
     */
    public readonly hasDiscussions!: pulumi.Output<boolean | undefined>;
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
     * Set to `true` to not call the vulnerability alerts endpoint so the resource can also be used without admin permissions during read.
     */
    public readonly ignoreVulnerabilityAlertsDuringRead!: pulumi.Output<boolean | undefined>;
    /**
     * Set to `true` to tell GitHub that this is a template repository.
     */
    public readonly isTemplate!: pulumi.Output<boolean | undefined>;
    /**
     * Use the [name of the template](https://github.com/github/choosealicense.com/tree/gh-pages/_licenses) without the extension. For example, "mit" or "mpl-2.0".
     */
    public readonly licenseTemplate!: pulumi.Output<string | undefined>;
    /**
     * Can be `PR_BODY`, `PR_TITLE`, or `BLANK` for a default merge commit message.
     */
    public readonly mergeCommitMessage!: pulumi.Output<string | undefined>;
    /**
     * Can be `PR_TITLE` or `MERGE_MESSAGE` for a default merge commit title.
     */
    public readonly mergeCommitTitle!: pulumi.Output<string | undefined>;
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
     * The repository's [security and analysis](https://docs.github.com/en/repositories/managing-your-repositorys-settings-and-features/enabling-features-for-your-repository/managing-security-and-analysis-settings-for-your-repository) configuration. See Security and Analysis Configuration below for details.
     */
    public readonly securityAndAnalysis!: pulumi.Output<outputs.RepositorySecurityAndAnalysis>;
    /**
     * Can be `PR_BODY`, `COMMIT_MESSAGES`, or `BLANK` for a default squash merge commit message.
     */
    public readonly squashMergeCommitMessage!: pulumi.Output<string | undefined>;
    /**
     * Can be `PR_TITLE` or `COMMIT_OR_PR_TITLE` for a default squash merge commit title.
     */
    public readonly squashMergeCommitTitle!: pulumi.Output<string | undefined>;
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
     * Set to `true` to enable security alerts for vulnerable dependencies. Enabling requires alerts to be enabled on the owner level. (Note for importing: GitHub enables the alerts on public repos but disables them on private repos by default.) See [GitHub Documentation](https://help.github.com/en/github/managing-security-vulnerabilities/about-security-alerts-for-vulnerable-dependencies) for details. Note that vulnerability alerts have not been successfully tested on any GitHub Enterprise instance and may be unavailable in those settings.
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
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RepositoryState | undefined;
            resourceInputs["allowAutoMerge"] = state ? state.allowAutoMerge : undefined;
            resourceInputs["allowMergeCommit"] = state ? state.allowMergeCommit : undefined;
            resourceInputs["allowRebaseMerge"] = state ? state.allowRebaseMerge : undefined;
            resourceInputs["allowSquashMerge"] = state ? state.allowSquashMerge : undefined;
            resourceInputs["allowUpdateBranch"] = state ? state.allowUpdateBranch : undefined;
            resourceInputs["archiveOnDestroy"] = state ? state.archiveOnDestroy : undefined;
            resourceInputs["archived"] = state ? state.archived : undefined;
            resourceInputs["autoInit"] = state ? state.autoInit : undefined;
            resourceInputs["defaultBranch"] = state ? state.defaultBranch : undefined;
            resourceInputs["deleteBranchOnMerge"] = state ? state.deleteBranchOnMerge : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["etag"] = state ? state.etag : undefined;
            resourceInputs["fullName"] = state ? state.fullName : undefined;
            resourceInputs["gitCloneUrl"] = state ? state.gitCloneUrl : undefined;
            resourceInputs["gitignoreTemplate"] = state ? state.gitignoreTemplate : undefined;
            resourceInputs["hasDiscussions"] = state ? state.hasDiscussions : undefined;
            resourceInputs["hasDownloads"] = state ? state.hasDownloads : undefined;
            resourceInputs["hasIssues"] = state ? state.hasIssues : undefined;
            resourceInputs["hasProjects"] = state ? state.hasProjects : undefined;
            resourceInputs["hasWiki"] = state ? state.hasWiki : undefined;
            resourceInputs["homepageUrl"] = state ? state.homepageUrl : undefined;
            resourceInputs["htmlUrl"] = state ? state.htmlUrl : undefined;
            resourceInputs["httpCloneUrl"] = state ? state.httpCloneUrl : undefined;
            resourceInputs["ignoreVulnerabilityAlertsDuringRead"] = state ? state.ignoreVulnerabilityAlertsDuringRead : undefined;
            resourceInputs["isTemplate"] = state ? state.isTemplate : undefined;
            resourceInputs["licenseTemplate"] = state ? state.licenseTemplate : undefined;
            resourceInputs["mergeCommitMessage"] = state ? state.mergeCommitMessage : undefined;
            resourceInputs["mergeCommitTitle"] = state ? state.mergeCommitTitle : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["nodeId"] = state ? state.nodeId : undefined;
            resourceInputs["pages"] = state ? state.pages : undefined;
            resourceInputs["private"] = state ? state.private : undefined;
            resourceInputs["repoId"] = state ? state.repoId : undefined;
            resourceInputs["securityAndAnalysis"] = state ? state.securityAndAnalysis : undefined;
            resourceInputs["squashMergeCommitMessage"] = state ? state.squashMergeCommitMessage : undefined;
            resourceInputs["squashMergeCommitTitle"] = state ? state.squashMergeCommitTitle : undefined;
            resourceInputs["sshCloneUrl"] = state ? state.sshCloneUrl : undefined;
            resourceInputs["svnUrl"] = state ? state.svnUrl : undefined;
            resourceInputs["template"] = state ? state.template : undefined;
            resourceInputs["topics"] = state ? state.topics : undefined;
            resourceInputs["visibility"] = state ? state.visibility : undefined;
            resourceInputs["vulnerabilityAlerts"] = state ? state.vulnerabilityAlerts : undefined;
        } else {
            const args = argsOrState as RepositoryArgs | undefined;
            resourceInputs["allowAutoMerge"] = args ? args.allowAutoMerge : undefined;
            resourceInputs["allowMergeCommit"] = args ? args.allowMergeCommit : undefined;
            resourceInputs["allowRebaseMerge"] = args ? args.allowRebaseMerge : undefined;
            resourceInputs["allowSquashMerge"] = args ? args.allowSquashMerge : undefined;
            resourceInputs["allowUpdateBranch"] = args ? args.allowUpdateBranch : undefined;
            resourceInputs["archiveOnDestroy"] = args ? args.archiveOnDestroy : undefined;
            resourceInputs["archived"] = args ? args.archived : undefined;
            resourceInputs["autoInit"] = args ? args.autoInit : undefined;
            resourceInputs["defaultBranch"] = args ? args.defaultBranch : undefined;
            resourceInputs["deleteBranchOnMerge"] = args ? args.deleteBranchOnMerge : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["gitignoreTemplate"] = args ? args.gitignoreTemplate : undefined;
            resourceInputs["hasDiscussions"] = args ? args.hasDiscussions : undefined;
            resourceInputs["hasDownloads"] = args ? args.hasDownloads : undefined;
            resourceInputs["hasIssues"] = args ? args.hasIssues : undefined;
            resourceInputs["hasProjects"] = args ? args.hasProjects : undefined;
            resourceInputs["hasWiki"] = args ? args.hasWiki : undefined;
            resourceInputs["homepageUrl"] = args ? args.homepageUrl : undefined;
            resourceInputs["ignoreVulnerabilityAlertsDuringRead"] = args ? args.ignoreVulnerabilityAlertsDuringRead : undefined;
            resourceInputs["isTemplate"] = args ? args.isTemplate : undefined;
            resourceInputs["licenseTemplate"] = args ? args.licenseTemplate : undefined;
            resourceInputs["mergeCommitMessage"] = args ? args.mergeCommitMessage : undefined;
            resourceInputs["mergeCommitTitle"] = args ? args.mergeCommitTitle : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["pages"] = args ? args.pages : undefined;
            resourceInputs["private"] = args ? args.private : undefined;
            resourceInputs["securityAndAnalysis"] = args ? args.securityAndAnalysis : undefined;
            resourceInputs["squashMergeCommitMessage"] = args ? args.squashMergeCommitMessage : undefined;
            resourceInputs["squashMergeCommitTitle"] = args ? args.squashMergeCommitTitle : undefined;
            resourceInputs["template"] = args ? args.template : undefined;
            resourceInputs["topics"] = args ? args.topics : undefined;
            resourceInputs["visibility"] = args ? args.visibility : undefined;
            resourceInputs["vulnerabilityAlerts"] = args ? args.vulnerabilityAlerts : undefined;
            resourceInputs["etag"] = undefined /*out*/;
            resourceInputs["fullName"] = undefined /*out*/;
            resourceInputs["gitCloneUrl"] = undefined /*out*/;
            resourceInputs["htmlUrl"] = undefined /*out*/;
            resourceInputs["httpCloneUrl"] = undefined /*out*/;
            resourceInputs["nodeId"] = undefined /*out*/;
            resourceInputs["repoId"] = undefined /*out*/;
            resourceInputs["sshCloneUrl"] = undefined /*out*/;
            resourceInputs["svnUrl"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Repository.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Repository resources.
 */
export interface RepositoryState {
    /**
     * Set to `true` to allow auto-merging pull requests on the repository.
     */
    allowAutoMerge?: pulumi.Input<boolean>;
    /**
     * Set to `false` to disable merge commits on the repository.
     */
    allowMergeCommit?: pulumi.Input<boolean>;
    /**
     * Set to `false` to disable rebase merges on the repository.
     */
    allowRebaseMerge?: pulumi.Input<boolean>;
    /**
     * Set to `false` to disable squash merges on the repository.
     */
    allowSquashMerge?: pulumi.Input<boolean>;
    /**
     * Set to `true` to always suggest updating pull request branches.
     */
    allowUpdateBranch?: pulumi.Input<boolean>;
    /**
     * Set to `true` to archive the repository instead of deleting on destroy.
     */
    archiveOnDestroy?: pulumi.Input<boolean>;
    /**
     * Specifies if the repository should be archived. Defaults to `false`. **NOTE** Currently, the API does not support unarchiving.
     */
    archived?: pulumi.Input<boolean>;
    /**
     * Set to `true` to produce an initial commit in the repository.
     */
    autoInit?: pulumi.Input<boolean>;
    /**
     * (Deprecated: Use `github.BranchDefault` resource instead) The name of the default branch of the repository. **NOTE:** This can only be set after a repository has already been created,
     * and after a correct reference has been created for the target branch inside the repository. This means a user will have to omit this parameter from the
     * initial repository creation and create the target branch inside of the repository prior to setting this attribute.
     *
     * @deprecated Use the github_branch_default resource instead
     */
    defaultBranch?: pulumi.Input<string>;
    /**
     * Automatically delete head branch after a pull request is merged. Defaults to `false`.
     */
    deleteBranchOnMerge?: pulumi.Input<boolean>;
    /**
     * A description of the repository.
     */
    description?: pulumi.Input<string>;
    etag?: pulumi.Input<string>;
    /**
     * A string of the form "orgname/reponame".
     */
    fullName?: pulumi.Input<string>;
    /**
     * URL that can be provided to `git clone` to clone the repository anonymously via the git protocol.
     */
    gitCloneUrl?: pulumi.Input<string>;
    /**
     * Use the [name of the template](https://github.com/github/gitignore) without the extension. For example, "Haskell".
     */
    gitignoreTemplate?: pulumi.Input<string>;
    /**
     * Set to `true` to enable GitHub Discussions on the repository. Defaults to `false`.
     */
    hasDiscussions?: pulumi.Input<boolean>;
    /**
     * Set to `true` to enable the (deprecated) downloads features on the repository.
     */
    hasDownloads?: pulumi.Input<boolean>;
    /**
     * Set to `true` to enable the GitHub Issues features
     * on the repository.
     */
    hasIssues?: pulumi.Input<boolean>;
    /**
     * Set to `true` to enable the GitHub Projects features on the repository. Per the GitHub [documentation](https://developer.github.com/v3/repos/#create) when in an organization that has disabled repository projects it will default to `false` and will otherwise default to `true`. If you specify `true` when it has been disabled it will return an error.
     */
    hasProjects?: pulumi.Input<boolean>;
    /**
     * Set to `true` to enable the GitHub Wiki features on
     * the repository.
     */
    hasWiki?: pulumi.Input<boolean>;
    /**
     * URL of a page describing the project.
     */
    homepageUrl?: pulumi.Input<string>;
    /**
     * The absolute URL (including scheme) of the rendered GitHub Pages site e.g. `https://username.github.io`.
     */
    htmlUrl?: pulumi.Input<string>;
    /**
     * URL that can be provided to `git clone` to clone the repository via HTTPS.
     */
    httpCloneUrl?: pulumi.Input<string>;
    /**
     * Set to `true` to not call the vulnerability alerts endpoint so the resource can also be used without admin permissions during read.
     */
    ignoreVulnerabilityAlertsDuringRead?: pulumi.Input<boolean>;
    /**
     * Set to `true` to tell GitHub that this is a template repository.
     */
    isTemplate?: pulumi.Input<boolean>;
    /**
     * Use the [name of the template](https://github.com/github/choosealicense.com/tree/gh-pages/_licenses) without the extension. For example, "mit" or "mpl-2.0".
     */
    licenseTemplate?: pulumi.Input<string>;
    /**
     * Can be `PR_BODY`, `PR_TITLE`, or `BLANK` for a default merge commit message.
     */
    mergeCommitMessage?: pulumi.Input<string>;
    /**
     * Can be `PR_TITLE` or `MERGE_MESSAGE` for a default merge commit title.
     */
    mergeCommitTitle?: pulumi.Input<string>;
    /**
     * The name of the repository.
     */
    name?: pulumi.Input<string>;
    /**
     * GraphQL global node id for use with v4 API
     */
    nodeId?: pulumi.Input<string>;
    /**
     * The repository's GitHub Pages configuration. See GitHub Pages Configuration below for details.
     */
    pages?: pulumi.Input<inputs.RepositoryPages>;
    /**
     * Set to `true` to create a private repository.
     * Repositories are created as public (e.g. open source) by default.
     *
     * @deprecated use visibility instead
     */
    private?: pulumi.Input<boolean>;
    /**
     * GitHub ID for the repository
     */
    repoId?: pulumi.Input<number>;
    /**
     * The repository's [security and analysis](https://docs.github.com/en/repositories/managing-your-repositorys-settings-and-features/enabling-features-for-your-repository/managing-security-and-analysis-settings-for-your-repository) configuration. See Security and Analysis Configuration below for details.
     */
    securityAndAnalysis?: pulumi.Input<inputs.RepositorySecurityAndAnalysis>;
    /**
     * Can be `PR_BODY`, `COMMIT_MESSAGES`, or `BLANK` for a default squash merge commit message.
     */
    squashMergeCommitMessage?: pulumi.Input<string>;
    /**
     * Can be `PR_TITLE` or `COMMIT_OR_PR_TITLE` for a default squash merge commit title.
     */
    squashMergeCommitTitle?: pulumi.Input<string>;
    /**
     * URL that can be provided to `git clone` to clone the repository via SSH.
     */
    sshCloneUrl?: pulumi.Input<string>;
    /**
     * URL that can be provided to `svn checkout` to check out the repository via GitHub's Subversion protocol emulation.
     */
    svnUrl?: pulumi.Input<string>;
    /**
     * Use a template repository to create this resource. See Template Repositories below for details.
     */
    template?: pulumi.Input<inputs.RepositoryTemplate>;
    /**
     * The list of topics of the repository.
     */
    topics?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Can be `public` or `private`. If your organization is associated with an enterprise account using GitHub Enterprise Cloud or GitHub Enterprise Server 2.20+, visibility can also be `internal`. The `visibility` parameter overrides the `private` parameter.
     */
    visibility?: pulumi.Input<string>;
    /**
     * Set to `true` to enable security alerts for vulnerable dependencies. Enabling requires alerts to be enabled on the owner level. (Note for importing: GitHub enables the alerts on public repos but disables them on private repos by default.) See [GitHub Documentation](https://help.github.com/en/github/managing-security-vulnerabilities/about-security-alerts-for-vulnerable-dependencies) for details. Note that vulnerability alerts have not been successfully tested on any GitHub Enterprise instance and may be unavailable in those settings.
     */
    vulnerabilityAlerts?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a Repository resource.
 */
export interface RepositoryArgs {
    /**
     * Set to `true` to allow auto-merging pull requests on the repository.
     */
    allowAutoMerge?: pulumi.Input<boolean>;
    /**
     * Set to `false` to disable merge commits on the repository.
     */
    allowMergeCommit?: pulumi.Input<boolean>;
    /**
     * Set to `false` to disable rebase merges on the repository.
     */
    allowRebaseMerge?: pulumi.Input<boolean>;
    /**
     * Set to `false` to disable squash merges on the repository.
     */
    allowSquashMerge?: pulumi.Input<boolean>;
    /**
     * Set to `true` to always suggest updating pull request branches.
     */
    allowUpdateBranch?: pulumi.Input<boolean>;
    /**
     * Set to `true` to archive the repository instead of deleting on destroy.
     */
    archiveOnDestroy?: pulumi.Input<boolean>;
    /**
     * Specifies if the repository should be archived. Defaults to `false`. **NOTE** Currently, the API does not support unarchiving.
     */
    archived?: pulumi.Input<boolean>;
    /**
     * Set to `true` to produce an initial commit in the repository.
     */
    autoInit?: pulumi.Input<boolean>;
    /**
     * (Deprecated: Use `github.BranchDefault` resource instead) The name of the default branch of the repository. **NOTE:** This can only be set after a repository has already been created,
     * and after a correct reference has been created for the target branch inside the repository. This means a user will have to omit this parameter from the
     * initial repository creation and create the target branch inside of the repository prior to setting this attribute.
     *
     * @deprecated Use the github_branch_default resource instead
     */
    defaultBranch?: pulumi.Input<string>;
    /**
     * Automatically delete head branch after a pull request is merged. Defaults to `false`.
     */
    deleteBranchOnMerge?: pulumi.Input<boolean>;
    /**
     * A description of the repository.
     */
    description?: pulumi.Input<string>;
    /**
     * Use the [name of the template](https://github.com/github/gitignore) without the extension. For example, "Haskell".
     */
    gitignoreTemplate?: pulumi.Input<string>;
    /**
     * Set to `true` to enable GitHub Discussions on the repository. Defaults to `false`.
     */
    hasDiscussions?: pulumi.Input<boolean>;
    /**
     * Set to `true` to enable the (deprecated) downloads features on the repository.
     */
    hasDownloads?: pulumi.Input<boolean>;
    /**
     * Set to `true` to enable the GitHub Issues features
     * on the repository.
     */
    hasIssues?: pulumi.Input<boolean>;
    /**
     * Set to `true` to enable the GitHub Projects features on the repository. Per the GitHub [documentation](https://developer.github.com/v3/repos/#create) when in an organization that has disabled repository projects it will default to `false` and will otherwise default to `true`. If you specify `true` when it has been disabled it will return an error.
     */
    hasProjects?: pulumi.Input<boolean>;
    /**
     * Set to `true` to enable the GitHub Wiki features on
     * the repository.
     */
    hasWiki?: pulumi.Input<boolean>;
    /**
     * URL of a page describing the project.
     */
    homepageUrl?: pulumi.Input<string>;
    /**
     * Set to `true` to not call the vulnerability alerts endpoint so the resource can also be used without admin permissions during read.
     */
    ignoreVulnerabilityAlertsDuringRead?: pulumi.Input<boolean>;
    /**
     * Set to `true` to tell GitHub that this is a template repository.
     */
    isTemplate?: pulumi.Input<boolean>;
    /**
     * Use the [name of the template](https://github.com/github/choosealicense.com/tree/gh-pages/_licenses) without the extension. For example, "mit" or "mpl-2.0".
     */
    licenseTemplate?: pulumi.Input<string>;
    /**
     * Can be `PR_BODY`, `PR_TITLE`, or `BLANK` for a default merge commit message.
     */
    mergeCommitMessage?: pulumi.Input<string>;
    /**
     * Can be `PR_TITLE` or `MERGE_MESSAGE` for a default merge commit title.
     */
    mergeCommitTitle?: pulumi.Input<string>;
    /**
     * The name of the repository.
     */
    name?: pulumi.Input<string>;
    /**
     * The repository's GitHub Pages configuration. See GitHub Pages Configuration below for details.
     */
    pages?: pulumi.Input<inputs.RepositoryPages>;
    /**
     * Set to `true` to create a private repository.
     * Repositories are created as public (e.g. open source) by default.
     *
     * @deprecated use visibility instead
     */
    private?: pulumi.Input<boolean>;
    /**
     * The repository's [security and analysis](https://docs.github.com/en/repositories/managing-your-repositorys-settings-and-features/enabling-features-for-your-repository/managing-security-and-analysis-settings-for-your-repository) configuration. See Security and Analysis Configuration below for details.
     */
    securityAndAnalysis?: pulumi.Input<inputs.RepositorySecurityAndAnalysis>;
    /**
     * Can be `PR_BODY`, `COMMIT_MESSAGES`, or `BLANK` for a default squash merge commit message.
     */
    squashMergeCommitMessage?: pulumi.Input<string>;
    /**
     * Can be `PR_TITLE` or `COMMIT_OR_PR_TITLE` for a default squash merge commit title.
     */
    squashMergeCommitTitle?: pulumi.Input<string>;
    /**
     * Use a template repository to create this resource. See Template Repositories below for details.
     */
    template?: pulumi.Input<inputs.RepositoryTemplate>;
    /**
     * The list of topics of the repository.
     */
    topics?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Can be `public` or `private`. If your organization is associated with an enterprise account using GitHub Enterprise Cloud or GitHub Enterprise Server 2.20+, visibility can also be `internal`. The `visibility` parameter overrides the `private` parameter.
     */
    visibility?: pulumi.Input<string>;
    /**
     * Set to `true` to enable security alerts for vulnerable dependencies. Enabling requires alerts to be enabled on the owner level. (Note for importing: GitHub enables the alerts on public repos but disables them on private repos by default.) See [GitHub Documentation](https://help.github.com/en/github/managing-security-vulnerabilities/about-security-alerts-for-vulnerable-dependencies) for details. Note that vulnerability alerts have not been successfully tested on any GitHub Enterprise instance and may be unavailable in those settings.
     */
    vulnerabilityAlerts?: pulumi.Input<boolean>;
}
