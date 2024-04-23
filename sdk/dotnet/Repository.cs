// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github
{
    /// <summary>
    /// This resource allows you to create and manage repositories within your
    /// GitHub organization or personal account.
    /// 
    /// &gt; Note: When used with GitHub App authentication, even GET requests must have
    /// the `contents:write` permission or else the `allow_merge_commit`, `allow_rebase_merge`,
    /// and `allow_squash_merge` attributes will be ignored, causing confusing diffs.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Github = Pulumi.Github;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Github.Repository("example", new()
    ///     {
    ///         Name = "example",
    ///         Description = "My awesome codebase",
    ///         Visibility = "public",
    ///         Template = new Github.Inputs.RepositoryTemplateArgs
    ///         {
    ///             Owner = "github",
    ///             Repository = "terraform-template-module",
    ///             IncludeAllBranches = true,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ### With GitHub Pages Enabled
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Github = Pulumi.Github;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Github.Repository("example", new()
    ///     {
    ///         Name = "example",
    ///         Description = "My awesome web page",
    ///         Private = false,
    ///         Pages = new Github.Inputs.RepositoryPagesArgs
    ///         {
    ///             Source = new Github.Inputs.RepositoryPagesSourceArgs
    ///             {
    ///                 Branch = "master",
    ///                 Path = "/docs",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Repositories can be imported using the `name`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import github:index/repository:Repository terraform terraform
    /// ```
    /// </summary>
    [GithubResourceType("github:index/repository:Repository")]
    public partial class Repository : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Set to `true` to allow auto-merging pull requests on the repository.
        /// </summary>
        [Output("allowAutoMerge")]
        public Output<bool?> AllowAutoMerge { get; private set; } = null!;

        /// <summary>
        /// Set to `false` to disable merge commits on the repository.
        /// </summary>
        [Output("allowMergeCommit")]
        public Output<bool?> AllowMergeCommit { get; private set; } = null!;

        /// <summary>
        /// Set to `false` to disable rebase merges on the repository.
        /// </summary>
        [Output("allowRebaseMerge")]
        public Output<bool?> AllowRebaseMerge { get; private set; } = null!;

        /// <summary>
        /// Set to `false` to disable squash merges on the repository.
        /// </summary>
        [Output("allowSquashMerge")]
        public Output<bool?> AllowSquashMerge { get; private set; } = null!;

        /// <summary>
        /// Set to `true` to always suggest updating pull request branches.
        /// </summary>
        [Output("allowUpdateBranch")]
        public Output<bool?> AllowUpdateBranch { get; private set; } = null!;

        /// <summary>
        /// Set to `true` to archive the repository instead of deleting on destroy.
        /// </summary>
        [Output("archiveOnDestroy")]
        public Output<bool?> ArchiveOnDestroy { get; private set; } = null!;

        /// <summary>
        /// Specifies if the repository should be archived. Defaults to `false`. **NOTE** Currently, the API does not support unarchiving.
        /// </summary>
        [Output("archived")]
        public Output<bool?> Archived { get; private set; } = null!;

        /// <summary>
        /// Set to `true` to produce an initial commit in the repository.
        /// </summary>
        [Output("autoInit")]
        public Output<bool?> AutoInit { get; private set; } = null!;

        /// <summary>
        /// (Deprecated: Use `github.BranchDefault` resource instead) The name of the default branch of the repository. **NOTE:** This can only be set after a repository has already been created,
        /// and after a correct reference has been created for the target branch inside the repository. This means a user will have to omit this parameter from the
        /// initial repository creation and create the target branch inside of the repository prior to setting this attribute.
        /// </summary>
        [Output("defaultBranch")]
        public Output<string> DefaultBranch { get; private set; } = null!;

        /// <summary>
        /// Automatically delete head branch after a pull request is merged. Defaults to `false`.
        /// </summary>
        [Output("deleteBranchOnMerge")]
        public Output<bool?> DeleteBranchOnMerge { get; private set; } = null!;

        /// <summary>
        /// A description of the repository.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// A string of the form "orgname/reponame".
        /// </summary>
        [Output("fullName")]
        public Output<string> FullName { get; private set; } = null!;

        /// <summary>
        /// URL that can be provided to `git clone` to clone the repository anonymously via the git protocol.
        /// </summary>
        [Output("gitCloneUrl")]
        public Output<string> GitCloneUrl { get; private set; } = null!;

        /// <summary>
        /// Use the [name of the template](https://github.com/github/gitignore) without the extension. For example, "Haskell".
        /// </summary>
        [Output("gitignoreTemplate")]
        public Output<string?> GitignoreTemplate { get; private set; } = null!;

        /// <summary>
        /// Set to `true` to enable GitHub Discussions on the repository. Defaults to `false`.
        /// </summary>
        [Output("hasDiscussions")]
        public Output<bool?> HasDiscussions { get; private set; } = null!;

        /// <summary>
        /// Set to `true` to enable the (deprecated) downloads features on the repository.
        /// </summary>
        [Output("hasDownloads")]
        public Output<bool?> HasDownloads { get; private set; } = null!;

        /// <summary>
        /// Set to `true` to enable the GitHub Issues features
        /// on the repository.
        /// </summary>
        [Output("hasIssues")]
        public Output<bool?> HasIssues { get; private set; } = null!;

        /// <summary>
        /// Set to `true` to enable the GitHub Projects features on the repository. Per the GitHub [documentation](https://developer.github.com/v3/repos/#create) when in an organization that has disabled repository projects it will default to `false` and will otherwise default to `true`. If you specify `true` when it has been disabled it will return an error.
        /// </summary>
        [Output("hasProjects")]
        public Output<bool?> HasProjects { get; private set; } = null!;

        /// <summary>
        /// Set to `true` to enable the GitHub Wiki features on
        /// the repository.
        /// </summary>
        [Output("hasWiki")]
        public Output<bool?> HasWiki { get; private set; } = null!;

        /// <summary>
        /// URL of a page describing the project.
        /// </summary>
        [Output("homepageUrl")]
        public Output<string?> HomepageUrl { get; private set; } = null!;

        /// <summary>
        /// The absolute URL (including scheme) of the rendered GitHub Pages site e.g. `https://username.github.io`.
        /// </summary>
        [Output("htmlUrl")]
        public Output<string> HtmlUrl { get; private set; } = null!;

        /// <summary>
        /// URL that can be provided to `git clone` to clone the repository via HTTPS.
        /// </summary>
        [Output("httpCloneUrl")]
        public Output<string> HttpCloneUrl { get; private set; } = null!;

        /// <summary>
        /// Set to `true` to not call the vulnerability alerts endpoint so the resource can also be used without admin permissions during read.
        /// </summary>
        [Output("ignoreVulnerabilityAlertsDuringRead")]
        public Output<bool?> IgnoreVulnerabilityAlertsDuringRead { get; private set; } = null!;

        /// <summary>
        /// Set to `true` to tell GitHub that this is a template repository.
        /// </summary>
        [Output("isTemplate")]
        public Output<bool?> IsTemplate { get; private set; } = null!;

        /// <summary>
        /// Use the [name of the template](https://github.com/github/choosealicense.com/tree/gh-pages/_licenses) without the extension. For example, "mit" or "mpl-2.0".
        /// </summary>
        [Output("licenseTemplate")]
        public Output<string?> LicenseTemplate { get; private set; } = null!;

        /// <summary>
        /// Can be `PR_BODY`, `PR_TITLE`, or `BLANK` for a default merge commit message. Applicable only if `allow_merge_commit` is `true`.
        /// </summary>
        [Output("mergeCommitMessage")]
        public Output<string?> MergeCommitMessage { get; private set; } = null!;

        /// <summary>
        /// Can be `PR_TITLE` or `MERGE_MESSAGE` for a default merge commit title. Applicable only if `allow_merge_commit` is `true`.
        /// </summary>
        [Output("mergeCommitTitle")]
        public Output<string?> MergeCommitTitle { get; private set; } = null!;

        /// <summary>
        /// The name of the repository.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// GraphQL global node id for use with v4 API
        /// </summary>
        [Output("nodeId")]
        public Output<string> NodeId { get; private set; } = null!;

        /// <summary>
        /// The repository's GitHub Pages configuration. See GitHub Pages Configuration below for details.
        /// </summary>
        [Output("pages")]
        public Output<Outputs.RepositoryPages?> Pages { get; private set; } = null!;

        /// <summary>
        /// The primary language used in the repository.
        /// </summary>
        [Output("primaryLanguage")]
        public Output<string> PrimaryLanguage { get; private set; } = null!;

        /// <summary>
        /// Set to `true` to create a private repository.
        /// Repositories are created as public (e.g. open source) by default.
        /// </summary>
        [Output("private")]
        public Output<bool> Private { get; private set; } = null!;

        /// <summary>
        /// GitHub ID for the repository
        /// </summary>
        [Output("repoId")]
        public Output<int> RepoId { get; private set; } = null!;

        /// <summary>
        /// The repository's [security and analysis](https://docs.github.com/en/repositories/managing-your-repositorys-settings-and-features/enabling-features-for-your-repository/managing-security-and-analysis-settings-for-your-repository) configuration. See Security and Analysis Configuration below for details.
        /// </summary>
        [Output("securityAndAnalysis")]
        public Output<Outputs.RepositorySecurityAndAnalysis> SecurityAndAnalysis { get; private set; } = null!;

        /// <summary>
        /// Can be `PR_BODY`, `COMMIT_MESSAGES`, or `BLANK` for a default squash merge commit message. Applicable only if `allow_squash_merge` is `true`.
        /// </summary>
        [Output("squashMergeCommitMessage")]
        public Output<string?> SquashMergeCommitMessage { get; private set; } = null!;

        /// <summary>
        /// Can be `PR_TITLE` or `COMMIT_OR_PR_TITLE` for a default squash merge commit title. Applicable only if `allow_squash_merge` is `true`.
        /// </summary>
        [Output("squashMergeCommitTitle")]
        public Output<string?> SquashMergeCommitTitle { get; private set; } = null!;

        /// <summary>
        /// URL that can be provided to `git clone` to clone the repository via SSH.
        /// </summary>
        [Output("sshCloneUrl")]
        public Output<string> SshCloneUrl { get; private set; } = null!;

        /// <summary>
        /// URL that can be provided to `svn checkout` to check out the repository via GitHub's Subversion protocol emulation.
        /// </summary>
        [Output("svnUrl")]
        public Output<string> SvnUrl { get; private set; } = null!;

        /// <summary>
        /// Use a template repository to create this resource. See Template Repositories below for details.
        /// </summary>
        [Output("template")]
        public Output<Outputs.RepositoryTemplate?> Template { get; private set; } = null!;

        /// <summary>
        /// The list of topics of the repository.
        /// </summary>
        [Output("topics")]
        public Output<ImmutableArray<string>> Topics { get; private set; } = null!;

        /// <summary>
        /// Can be `public` or `private`. If your organization is associated with an enterprise account using GitHub Enterprise Cloud or GitHub Enterprise Server 2.20+, visibility can also be `internal`. The `visibility` parameter overrides the `private` parameter.
        /// </summary>
        [Output("visibility")]
        public Output<string> Visibility { get; private set; } = null!;

        /// <summary>
        /// Set to `true` to enable security alerts for vulnerable dependencies. Enabling requires alerts to be enabled on the owner level. (Note for importing: GitHub enables the alerts on public repos but disables them on private repos by default.) See [GitHub Documentation](https://help.github.com/en/github/managing-security-vulnerabilities/about-security-alerts-for-vulnerable-dependencies) for details. Note that vulnerability alerts have not been successfully tested on any GitHub Enterprise instance and may be unavailable in those settings.
        /// </summary>
        [Output("vulnerabilityAlerts")]
        public Output<bool?> VulnerabilityAlerts { get; private set; } = null!;

        /// <summary>
        /// Require contributors to sign off on web-based commits. See more [here](https://docs.github.com/en/repositories/managing-your-repositorys-settings-and-features/managing-repository-settings/managing-the-commit-signoff-policy-for-your-repository). Defaults to `false`.
        /// </summary>
        [Output("webCommitSignoffRequired")]
        public Output<bool?> WebCommitSignoffRequired { get; private set; } = null!;


        /// <summary>
        /// Create a Repository resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Repository(string name, RepositoryArgs? args = null, CustomResourceOptions? options = null)
            : base("github:index/repository:Repository", name, args ?? new RepositoryArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Repository(string name, Input<string> id, RepositoryState? state = null, CustomResourceOptions? options = null)
            : base("github:index/repository:Repository", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Repository resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Repository Get(string name, Input<string> id, RepositoryState? state = null, CustomResourceOptions? options = null)
        {
            return new Repository(name, id, state, options);
        }
    }

    public sealed class RepositoryArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Set to `true` to allow auto-merging pull requests on the repository.
        /// </summary>
        [Input("allowAutoMerge")]
        public Input<bool>? AllowAutoMerge { get; set; }

        /// <summary>
        /// Set to `false` to disable merge commits on the repository.
        /// </summary>
        [Input("allowMergeCommit")]
        public Input<bool>? AllowMergeCommit { get; set; }

        /// <summary>
        /// Set to `false` to disable rebase merges on the repository.
        /// </summary>
        [Input("allowRebaseMerge")]
        public Input<bool>? AllowRebaseMerge { get; set; }

        /// <summary>
        /// Set to `false` to disable squash merges on the repository.
        /// </summary>
        [Input("allowSquashMerge")]
        public Input<bool>? AllowSquashMerge { get; set; }

        /// <summary>
        /// Set to `true` to always suggest updating pull request branches.
        /// </summary>
        [Input("allowUpdateBranch")]
        public Input<bool>? AllowUpdateBranch { get; set; }

        /// <summary>
        /// Set to `true` to archive the repository instead of deleting on destroy.
        /// </summary>
        [Input("archiveOnDestroy")]
        public Input<bool>? ArchiveOnDestroy { get; set; }

        /// <summary>
        /// Specifies if the repository should be archived. Defaults to `false`. **NOTE** Currently, the API does not support unarchiving.
        /// </summary>
        [Input("archived")]
        public Input<bool>? Archived { get; set; }

        /// <summary>
        /// Set to `true` to produce an initial commit in the repository.
        /// </summary>
        [Input("autoInit")]
        public Input<bool>? AutoInit { get; set; }

        /// <summary>
        /// (Deprecated: Use `github.BranchDefault` resource instead) The name of the default branch of the repository. **NOTE:** This can only be set after a repository has already been created,
        /// and after a correct reference has been created for the target branch inside the repository. This means a user will have to omit this parameter from the
        /// initial repository creation and create the target branch inside of the repository prior to setting this attribute.
        /// </summary>
        [Input("defaultBranch")]
        public Input<string>? DefaultBranch { get; set; }

        /// <summary>
        /// Automatically delete head branch after a pull request is merged. Defaults to `false`.
        /// </summary>
        [Input("deleteBranchOnMerge")]
        public Input<bool>? DeleteBranchOnMerge { get; set; }

        /// <summary>
        /// A description of the repository.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Use the [name of the template](https://github.com/github/gitignore) without the extension. For example, "Haskell".
        /// </summary>
        [Input("gitignoreTemplate")]
        public Input<string>? GitignoreTemplate { get; set; }

        /// <summary>
        /// Set to `true` to enable GitHub Discussions on the repository. Defaults to `false`.
        /// </summary>
        [Input("hasDiscussions")]
        public Input<bool>? HasDiscussions { get; set; }

        /// <summary>
        /// Set to `true` to enable the (deprecated) downloads features on the repository.
        /// </summary>
        [Input("hasDownloads")]
        public Input<bool>? HasDownloads { get; set; }

        /// <summary>
        /// Set to `true` to enable the GitHub Issues features
        /// on the repository.
        /// </summary>
        [Input("hasIssues")]
        public Input<bool>? HasIssues { get; set; }

        /// <summary>
        /// Set to `true` to enable the GitHub Projects features on the repository. Per the GitHub [documentation](https://developer.github.com/v3/repos/#create) when in an organization that has disabled repository projects it will default to `false` and will otherwise default to `true`. If you specify `true` when it has been disabled it will return an error.
        /// </summary>
        [Input("hasProjects")]
        public Input<bool>? HasProjects { get; set; }

        /// <summary>
        /// Set to `true` to enable the GitHub Wiki features on
        /// the repository.
        /// </summary>
        [Input("hasWiki")]
        public Input<bool>? HasWiki { get; set; }

        /// <summary>
        /// URL of a page describing the project.
        /// </summary>
        [Input("homepageUrl")]
        public Input<string>? HomepageUrl { get; set; }

        /// <summary>
        /// Set to `true` to not call the vulnerability alerts endpoint so the resource can also be used without admin permissions during read.
        /// </summary>
        [Input("ignoreVulnerabilityAlertsDuringRead")]
        public Input<bool>? IgnoreVulnerabilityAlertsDuringRead { get; set; }

        /// <summary>
        /// Set to `true` to tell GitHub that this is a template repository.
        /// </summary>
        [Input("isTemplate")]
        public Input<bool>? IsTemplate { get; set; }

        /// <summary>
        /// Use the [name of the template](https://github.com/github/choosealicense.com/tree/gh-pages/_licenses) without the extension. For example, "mit" or "mpl-2.0".
        /// </summary>
        [Input("licenseTemplate")]
        public Input<string>? LicenseTemplate { get; set; }

        /// <summary>
        /// Can be `PR_BODY`, `PR_TITLE`, or `BLANK` for a default merge commit message. Applicable only if `allow_merge_commit` is `true`.
        /// </summary>
        [Input("mergeCommitMessage")]
        public Input<string>? MergeCommitMessage { get; set; }

        /// <summary>
        /// Can be `PR_TITLE` or `MERGE_MESSAGE` for a default merge commit title. Applicable only if `allow_merge_commit` is `true`.
        /// </summary>
        [Input("mergeCommitTitle")]
        public Input<string>? MergeCommitTitle { get; set; }

        /// <summary>
        /// The name of the repository.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The repository's GitHub Pages configuration. See GitHub Pages Configuration below for details.
        /// </summary>
        [Input("pages")]
        public Input<Inputs.RepositoryPagesArgs>? Pages { get; set; }

        /// <summary>
        /// Set to `true` to create a private repository.
        /// Repositories are created as public (e.g. open source) by default.
        /// </summary>
        [Input("private")]
        public Input<bool>? Private { get; set; }

        /// <summary>
        /// The repository's [security and analysis](https://docs.github.com/en/repositories/managing-your-repositorys-settings-and-features/enabling-features-for-your-repository/managing-security-and-analysis-settings-for-your-repository) configuration. See Security and Analysis Configuration below for details.
        /// </summary>
        [Input("securityAndAnalysis")]
        public Input<Inputs.RepositorySecurityAndAnalysisArgs>? SecurityAndAnalysis { get; set; }

        /// <summary>
        /// Can be `PR_BODY`, `COMMIT_MESSAGES`, or `BLANK` for a default squash merge commit message. Applicable only if `allow_squash_merge` is `true`.
        /// </summary>
        [Input("squashMergeCommitMessage")]
        public Input<string>? SquashMergeCommitMessage { get; set; }

        /// <summary>
        /// Can be `PR_TITLE` or `COMMIT_OR_PR_TITLE` for a default squash merge commit title. Applicable only if `allow_squash_merge` is `true`.
        /// </summary>
        [Input("squashMergeCommitTitle")]
        public Input<string>? SquashMergeCommitTitle { get; set; }

        /// <summary>
        /// Use a template repository to create this resource. See Template Repositories below for details.
        /// </summary>
        [Input("template")]
        public Input<Inputs.RepositoryTemplateArgs>? Template { get; set; }

        [Input("topics")]
        private InputList<string>? _topics;

        /// <summary>
        /// The list of topics of the repository.
        /// </summary>
        public InputList<string> Topics
        {
            get => _topics ?? (_topics = new InputList<string>());
            set => _topics = value;
        }

        /// <summary>
        /// Can be `public` or `private`. If your organization is associated with an enterprise account using GitHub Enterprise Cloud or GitHub Enterprise Server 2.20+, visibility can also be `internal`. The `visibility` parameter overrides the `private` parameter.
        /// </summary>
        [Input("visibility")]
        public Input<string>? Visibility { get; set; }

        /// <summary>
        /// Set to `true` to enable security alerts for vulnerable dependencies. Enabling requires alerts to be enabled on the owner level. (Note for importing: GitHub enables the alerts on public repos but disables them on private repos by default.) See [GitHub Documentation](https://help.github.com/en/github/managing-security-vulnerabilities/about-security-alerts-for-vulnerable-dependencies) for details. Note that vulnerability alerts have not been successfully tested on any GitHub Enterprise instance and may be unavailable in those settings.
        /// </summary>
        [Input("vulnerabilityAlerts")]
        public Input<bool>? VulnerabilityAlerts { get; set; }

        /// <summary>
        /// Require contributors to sign off on web-based commits. See more [here](https://docs.github.com/en/repositories/managing-your-repositorys-settings-and-features/managing-repository-settings/managing-the-commit-signoff-policy-for-your-repository). Defaults to `false`.
        /// </summary>
        [Input("webCommitSignoffRequired")]
        public Input<bool>? WebCommitSignoffRequired { get; set; }

        public RepositoryArgs()
        {
        }
        public static new RepositoryArgs Empty => new RepositoryArgs();
    }

    public sealed class RepositoryState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Set to `true` to allow auto-merging pull requests on the repository.
        /// </summary>
        [Input("allowAutoMerge")]
        public Input<bool>? AllowAutoMerge { get; set; }

        /// <summary>
        /// Set to `false` to disable merge commits on the repository.
        /// </summary>
        [Input("allowMergeCommit")]
        public Input<bool>? AllowMergeCommit { get; set; }

        /// <summary>
        /// Set to `false` to disable rebase merges on the repository.
        /// </summary>
        [Input("allowRebaseMerge")]
        public Input<bool>? AllowRebaseMerge { get; set; }

        /// <summary>
        /// Set to `false` to disable squash merges on the repository.
        /// </summary>
        [Input("allowSquashMerge")]
        public Input<bool>? AllowSquashMerge { get; set; }

        /// <summary>
        /// Set to `true` to always suggest updating pull request branches.
        /// </summary>
        [Input("allowUpdateBranch")]
        public Input<bool>? AllowUpdateBranch { get; set; }

        /// <summary>
        /// Set to `true` to archive the repository instead of deleting on destroy.
        /// </summary>
        [Input("archiveOnDestroy")]
        public Input<bool>? ArchiveOnDestroy { get; set; }

        /// <summary>
        /// Specifies if the repository should be archived. Defaults to `false`. **NOTE** Currently, the API does not support unarchiving.
        /// </summary>
        [Input("archived")]
        public Input<bool>? Archived { get; set; }

        /// <summary>
        /// Set to `true` to produce an initial commit in the repository.
        /// </summary>
        [Input("autoInit")]
        public Input<bool>? AutoInit { get; set; }

        /// <summary>
        /// (Deprecated: Use `github.BranchDefault` resource instead) The name of the default branch of the repository. **NOTE:** This can only be set after a repository has already been created,
        /// and after a correct reference has been created for the target branch inside the repository. This means a user will have to omit this parameter from the
        /// initial repository creation and create the target branch inside of the repository prior to setting this attribute.
        /// </summary>
        [Input("defaultBranch")]
        public Input<string>? DefaultBranch { get; set; }

        /// <summary>
        /// Automatically delete head branch after a pull request is merged. Defaults to `false`.
        /// </summary>
        [Input("deleteBranchOnMerge")]
        public Input<bool>? DeleteBranchOnMerge { get; set; }

        /// <summary>
        /// A description of the repository.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// A string of the form "orgname/reponame".
        /// </summary>
        [Input("fullName")]
        public Input<string>? FullName { get; set; }

        /// <summary>
        /// URL that can be provided to `git clone` to clone the repository anonymously via the git protocol.
        /// </summary>
        [Input("gitCloneUrl")]
        public Input<string>? GitCloneUrl { get; set; }

        /// <summary>
        /// Use the [name of the template](https://github.com/github/gitignore) without the extension. For example, "Haskell".
        /// </summary>
        [Input("gitignoreTemplate")]
        public Input<string>? GitignoreTemplate { get; set; }

        /// <summary>
        /// Set to `true` to enable GitHub Discussions on the repository. Defaults to `false`.
        /// </summary>
        [Input("hasDiscussions")]
        public Input<bool>? HasDiscussions { get; set; }

        /// <summary>
        /// Set to `true` to enable the (deprecated) downloads features on the repository.
        /// </summary>
        [Input("hasDownloads")]
        public Input<bool>? HasDownloads { get; set; }

        /// <summary>
        /// Set to `true` to enable the GitHub Issues features
        /// on the repository.
        /// </summary>
        [Input("hasIssues")]
        public Input<bool>? HasIssues { get; set; }

        /// <summary>
        /// Set to `true` to enable the GitHub Projects features on the repository. Per the GitHub [documentation](https://developer.github.com/v3/repos/#create) when in an organization that has disabled repository projects it will default to `false` and will otherwise default to `true`. If you specify `true` when it has been disabled it will return an error.
        /// </summary>
        [Input("hasProjects")]
        public Input<bool>? HasProjects { get; set; }

        /// <summary>
        /// Set to `true` to enable the GitHub Wiki features on
        /// the repository.
        /// </summary>
        [Input("hasWiki")]
        public Input<bool>? HasWiki { get; set; }

        /// <summary>
        /// URL of a page describing the project.
        /// </summary>
        [Input("homepageUrl")]
        public Input<string>? HomepageUrl { get; set; }

        /// <summary>
        /// The absolute URL (including scheme) of the rendered GitHub Pages site e.g. `https://username.github.io`.
        /// </summary>
        [Input("htmlUrl")]
        public Input<string>? HtmlUrl { get; set; }

        /// <summary>
        /// URL that can be provided to `git clone` to clone the repository via HTTPS.
        /// </summary>
        [Input("httpCloneUrl")]
        public Input<string>? HttpCloneUrl { get; set; }

        /// <summary>
        /// Set to `true` to not call the vulnerability alerts endpoint so the resource can also be used without admin permissions during read.
        /// </summary>
        [Input("ignoreVulnerabilityAlertsDuringRead")]
        public Input<bool>? IgnoreVulnerabilityAlertsDuringRead { get; set; }

        /// <summary>
        /// Set to `true` to tell GitHub that this is a template repository.
        /// </summary>
        [Input("isTemplate")]
        public Input<bool>? IsTemplate { get; set; }

        /// <summary>
        /// Use the [name of the template](https://github.com/github/choosealicense.com/tree/gh-pages/_licenses) without the extension. For example, "mit" or "mpl-2.0".
        /// </summary>
        [Input("licenseTemplate")]
        public Input<string>? LicenseTemplate { get; set; }

        /// <summary>
        /// Can be `PR_BODY`, `PR_TITLE`, or `BLANK` for a default merge commit message. Applicable only if `allow_merge_commit` is `true`.
        /// </summary>
        [Input("mergeCommitMessage")]
        public Input<string>? MergeCommitMessage { get; set; }

        /// <summary>
        /// Can be `PR_TITLE` or `MERGE_MESSAGE` for a default merge commit title. Applicable only if `allow_merge_commit` is `true`.
        /// </summary>
        [Input("mergeCommitTitle")]
        public Input<string>? MergeCommitTitle { get; set; }

        /// <summary>
        /// The name of the repository.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// GraphQL global node id for use with v4 API
        /// </summary>
        [Input("nodeId")]
        public Input<string>? NodeId { get; set; }

        /// <summary>
        /// The repository's GitHub Pages configuration. See GitHub Pages Configuration below for details.
        /// </summary>
        [Input("pages")]
        public Input<Inputs.RepositoryPagesGetArgs>? Pages { get; set; }

        /// <summary>
        /// The primary language used in the repository.
        /// </summary>
        [Input("primaryLanguage")]
        public Input<string>? PrimaryLanguage { get; set; }

        /// <summary>
        /// Set to `true` to create a private repository.
        /// Repositories are created as public (e.g. open source) by default.
        /// </summary>
        [Input("private")]
        public Input<bool>? Private { get; set; }

        /// <summary>
        /// GitHub ID for the repository
        /// </summary>
        [Input("repoId")]
        public Input<int>? RepoId { get; set; }

        /// <summary>
        /// The repository's [security and analysis](https://docs.github.com/en/repositories/managing-your-repositorys-settings-and-features/enabling-features-for-your-repository/managing-security-and-analysis-settings-for-your-repository) configuration. See Security and Analysis Configuration below for details.
        /// </summary>
        [Input("securityAndAnalysis")]
        public Input<Inputs.RepositorySecurityAndAnalysisGetArgs>? SecurityAndAnalysis { get; set; }

        /// <summary>
        /// Can be `PR_BODY`, `COMMIT_MESSAGES`, or `BLANK` for a default squash merge commit message. Applicable only if `allow_squash_merge` is `true`.
        /// </summary>
        [Input("squashMergeCommitMessage")]
        public Input<string>? SquashMergeCommitMessage { get; set; }

        /// <summary>
        /// Can be `PR_TITLE` or `COMMIT_OR_PR_TITLE` for a default squash merge commit title. Applicable only if `allow_squash_merge` is `true`.
        /// </summary>
        [Input("squashMergeCommitTitle")]
        public Input<string>? SquashMergeCommitTitle { get; set; }

        /// <summary>
        /// URL that can be provided to `git clone` to clone the repository via SSH.
        /// </summary>
        [Input("sshCloneUrl")]
        public Input<string>? SshCloneUrl { get; set; }

        /// <summary>
        /// URL that can be provided to `svn checkout` to check out the repository via GitHub's Subversion protocol emulation.
        /// </summary>
        [Input("svnUrl")]
        public Input<string>? SvnUrl { get; set; }

        /// <summary>
        /// Use a template repository to create this resource. See Template Repositories below for details.
        /// </summary>
        [Input("template")]
        public Input<Inputs.RepositoryTemplateGetArgs>? Template { get; set; }

        [Input("topics")]
        private InputList<string>? _topics;

        /// <summary>
        /// The list of topics of the repository.
        /// </summary>
        public InputList<string> Topics
        {
            get => _topics ?? (_topics = new InputList<string>());
            set => _topics = value;
        }

        /// <summary>
        /// Can be `public` or `private`. If your organization is associated with an enterprise account using GitHub Enterprise Cloud or GitHub Enterprise Server 2.20+, visibility can also be `internal`. The `visibility` parameter overrides the `private` parameter.
        /// </summary>
        [Input("visibility")]
        public Input<string>? Visibility { get; set; }

        /// <summary>
        /// Set to `true` to enable security alerts for vulnerable dependencies. Enabling requires alerts to be enabled on the owner level. (Note for importing: GitHub enables the alerts on public repos but disables them on private repos by default.) See [GitHub Documentation](https://help.github.com/en/github/managing-security-vulnerabilities/about-security-alerts-for-vulnerable-dependencies) for details. Note that vulnerability alerts have not been successfully tested on any GitHub Enterprise instance and may be unavailable in those settings.
        /// </summary>
        [Input("vulnerabilityAlerts")]
        public Input<bool>? VulnerabilityAlerts { get; set; }

        /// <summary>
        /// Require contributors to sign off on web-based commits. See more [here](https://docs.github.com/en/repositories/managing-your-repositorys-settings-and-features/managing-repository-settings/managing-the-commit-signoff-policy-for-your-repository). Defaults to `false`.
        /// </summary>
        [Input("webCommitSignoffRequired")]
        public Input<bool>? WebCommitSignoffRequired { get; set; }

        public RepositoryState()
        {
        }
        public static new RepositoryState Empty => new RepositoryState();
    }
}
