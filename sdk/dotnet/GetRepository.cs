// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github
{
    public static class GetRepository
    {
        /// <summary>
        /// Use this data source to retrieve information about a GitHub repository.
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
        ///     var example = Github.GetRepository.Invoke(new()
        ///     {
        ///         FullName = "hashicorp/terraform",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetRepositoryResult> InvokeAsync(GetRepositoryArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRepositoryResult>("github:index/getRepository:getRepository", args ?? new GetRepositoryArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve information about a GitHub repository.
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
        ///     var example = Github.GetRepository.Invoke(new()
        ///     {
        ///         FullName = "hashicorp/terraform",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetRepositoryResult> Invoke(GetRepositoryInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRepositoryResult>("github:index/getRepository:getRepository", args ?? new GetRepositoryInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRepositoryArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A description of the license.
        /// </summary>
        [Input("description")]
        public string? Description { get; set; }

        /// <summary>
        /// Full name of the repository (in `org/name` format).
        /// </summary>
        [Input("fullName")]
        public string? FullName { get; set; }

        /// <summary>
        /// URL of a page describing the project.
        /// </summary>
        [Input("homepageUrl")]
        public string? HomepageUrl { get; set; }

        /// <summary>
        /// The name of the repository.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        public GetRepositoryArgs()
        {
        }
        public static new GetRepositoryArgs Empty => new GetRepositoryArgs();
    }

    public sealed class GetRepositoryInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A description of the license.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Full name of the repository (in `org/name` format).
        /// </summary>
        [Input("fullName")]
        public Input<string>? FullName { get; set; }

        /// <summary>
        /// URL of a page describing the project.
        /// </summary>
        [Input("homepageUrl")]
        public Input<string>? HomepageUrl { get; set; }

        /// <summary>
        /// The name of the repository.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public GetRepositoryInvokeArgs()
        {
        }
        public static new GetRepositoryInvokeArgs Empty => new GetRepositoryInvokeArgs();
    }


    [OutputType]
    public sealed class GetRepositoryResult
    {
        /// <summary>
        /// Whether the repository allows auto-merging pull requests.
        /// </summary>
        public readonly bool AllowAutoMerge;
        /// <summary>
        /// Whether the repository allows merge commits.
        /// </summary>
        public readonly bool AllowMergeCommit;
        /// <summary>
        /// Whether the repository allows rebase merges.
        /// </summary>
        public readonly bool AllowRebaseMerge;
        /// <summary>
        /// Whether the repository allows squash merges.
        /// </summary>
        public readonly bool AllowSquashMerge;
        public readonly bool AllowUpdateBranch;
        /// <summary>
        /// Whether the repository is archived.
        /// </summary>
        public readonly bool Archived;
        /// <summary>
        /// The name of the default branch of the repository.
        /// </summary>
        public readonly string DefaultBranch;
        public readonly bool DeleteBranchOnMerge;
        /// <summary>
        /// A description of the license.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Whether the repository is a fork.
        /// </summary>
        public readonly bool Fork;
        public readonly string FullName;
        /// <summary>
        /// URL that can be provided to `git clone` to clone the repository anonymously via the git protocol.
        /// </summary>
        public readonly string GitCloneUrl;
        /// <summary>
        /// Whether the repository has GitHub Discussions enabled.
        /// </summary>
        public readonly bool HasDiscussions;
        /// <summary>
        /// Whether the repository has Downloads feature enabled.
        /// </summary>
        public readonly bool HasDownloads;
        /// <summary>
        /// Whether the repository has GitHub Issues enabled.
        /// </summary>
        public readonly bool HasIssues;
        /// <summary>
        /// Whether the repository has the GitHub Projects enabled.
        /// </summary>
        public readonly bool HasProjects;
        /// <summary>
        /// Whether the repository has the GitHub Wiki enabled.
        /// </summary>
        public readonly bool HasWiki;
        /// <summary>
        /// URL of a page describing the project.
        /// </summary>
        public readonly string? HomepageUrl;
        /// <summary>
        /// The URL to view the license details on GitHub.
        /// </summary>
        public readonly string HtmlUrl;
        /// <summary>
        /// URL that can be provided to `git clone` to clone the repository via HTTPS.
        /// </summary>
        public readonly string HttpCloneUrl;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Whether the repository is a template repository.
        /// </summary>
        public readonly bool IsTemplate;
        /// <summary>
        /// The default value for a merge commit message.
        /// </summary>
        public readonly string MergeCommitMessage;
        /// <summary>
        /// The default value for a merge commit title.
        /// </summary>
        public readonly string MergeCommitTitle;
        /// <summary>
        /// The name of the license (e.g., "Apache License 2.0").
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// GraphQL global node id for use with v4 API
        /// </summary>
        public readonly string NodeId;
        /// <summary>
        /// The repository's GitHub Pages configuration.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetRepositoryPageResult> Pages;
        /// <summary>
        /// The primary language used in the repository.
        /// </summary>
        public readonly string PrimaryLanguage;
        /// <summary>
        /// Whether the repository is private.
        /// </summary>
        public readonly bool Private;
        /// <summary>
        /// GitHub ID for the repository
        /// </summary>
        public readonly int RepoId;
        /// <summary>
        /// An Array of GitHub repository licenses. Each `repository_license` block consists of the fields documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetRepositoryRepositoryLicenseResult> RepositoryLicenses;
        /// <summary>
        /// The default value for a squash merge commit message.
        /// </summary>
        public readonly string SquashMergeCommitMessage;
        /// <summary>
        /// The default value for a squash merge commit title.
        /// </summary>
        public readonly string SquashMergeCommitTitle;
        /// <summary>
        /// URL that can be provided to `git clone` to clone the repository via SSH.
        /// </summary>
        public readonly string SshCloneUrl;
        /// <summary>
        /// URL that can be provided to `svn checkout` to check out the repository via GitHub's Subversion protocol emulation.
        /// </summary>
        public readonly string SvnUrl;
        /// <summary>
        /// The repository source template configuration.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetRepositoryTemplateResult> Templates;
        /// <summary>
        /// The list of topics of the repository.
        /// </summary>
        public readonly ImmutableArray<string> Topics;
        /// <summary>
        /// Whether the repository is public, private or internal.
        /// </summary>
        public readonly string Visibility;

        [OutputConstructor]
        private GetRepositoryResult(
            bool allowAutoMerge,

            bool allowMergeCommit,

            bool allowRebaseMerge,

            bool allowSquashMerge,

            bool allowUpdateBranch,

            bool archived,

            string defaultBranch,

            bool deleteBranchOnMerge,

            string? description,

            bool fork,

            string fullName,

            string gitCloneUrl,

            bool hasDiscussions,

            bool hasDownloads,

            bool hasIssues,

            bool hasProjects,

            bool hasWiki,

            string? homepageUrl,

            string htmlUrl,

            string httpCloneUrl,

            string id,

            bool isTemplate,

            string mergeCommitMessage,

            string mergeCommitTitle,

            string name,

            string nodeId,

            ImmutableArray<Outputs.GetRepositoryPageResult> pages,

            string primaryLanguage,

            bool @private,

            int repoId,

            ImmutableArray<Outputs.GetRepositoryRepositoryLicenseResult> repositoryLicenses,

            string squashMergeCommitMessage,

            string squashMergeCommitTitle,

            string sshCloneUrl,

            string svnUrl,

            ImmutableArray<Outputs.GetRepositoryTemplateResult> templates,

            ImmutableArray<string> topics,

            string visibility)
        {
            AllowAutoMerge = allowAutoMerge;
            AllowMergeCommit = allowMergeCommit;
            AllowRebaseMerge = allowRebaseMerge;
            AllowSquashMerge = allowSquashMerge;
            AllowUpdateBranch = allowUpdateBranch;
            Archived = archived;
            DefaultBranch = defaultBranch;
            DeleteBranchOnMerge = deleteBranchOnMerge;
            Description = description;
            Fork = fork;
            FullName = fullName;
            GitCloneUrl = gitCloneUrl;
            HasDiscussions = hasDiscussions;
            HasDownloads = hasDownloads;
            HasIssues = hasIssues;
            HasProjects = hasProjects;
            HasWiki = hasWiki;
            HomepageUrl = homepageUrl;
            HtmlUrl = htmlUrl;
            HttpCloneUrl = httpCloneUrl;
            Id = id;
            IsTemplate = isTemplate;
            MergeCommitMessage = mergeCommitMessage;
            MergeCommitTitle = mergeCommitTitle;
            Name = name;
            NodeId = nodeId;
            Pages = pages;
            PrimaryLanguage = primaryLanguage;
            Private = @private;
            RepoId = repoId;
            RepositoryLicenses = repositoryLicenses;
            SquashMergeCommitMessage = squashMergeCommitMessage;
            SquashMergeCommitTitle = squashMergeCommitTitle;
            SshCloneUrl = sshCloneUrl;
            SvnUrl = svnUrl;
            Templates = templates;
            Topics = topics;
            Visibility = visibility;
        }
    }
}
