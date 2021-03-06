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
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Github = Pulumi.Github;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(Github.GetRepository.InvokeAsync(new Github.GetRepositoryArgs
        ///         {
        ///             FullName = "hashicorp/terraform",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetRepositoryResult> InvokeAsync(GetRepositoryArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRepositoryResult>("github:index/getRepository:getRepository", args ?? new GetRepositoryArgs(), options.WithVersion());
    }


    public sealed class GetRepositoryArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Full name of the repository (in `org/name` format).
        /// </summary>
        [Input("fullName")]
        public string? FullName { get; set; }

        /// <summary>
        /// The name of the repository.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        public GetRepositoryArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetRepositoryResult
    {
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
        /// <summary>
        /// Whether the repository is archived.
        /// </summary>
        public readonly bool Archived;
        /// <summary>
        /// The name of the default branch of the repository.
        /// </summary>
        public readonly string DefaultBranch;
        /// <summary>
        /// A description of the repository.
        /// </summary>
        public readonly string Description;
        public readonly string? FullName;
        /// <summary>
        /// URL that can be provided to `git clone` to clone the repository anonymously via the git protocol.
        /// </summary>
        public readonly string GitCloneUrl;
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
        public readonly string HomepageUrl;
        /// <summary>
        /// URL to the repository on the web.
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
        public readonly string? Name;
        /// <summary>
        /// GraphQL global node id for use with v4 API
        /// </summary>
        public readonly string NodeId;
        /// <summary>
        /// The repository's GitHub Pages configuration.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetRepositoryPageResult> Pages;
        /// <summary>
        /// Whether the repository is private.
        /// </summary>
        public readonly bool Private;
        /// <summary>
        /// GitHub ID for the repository
        /// </summary>
        public readonly int RepoId;
        /// <summary>
        /// URL that can be provided to `git clone` to clone the repository via SSH.
        /// </summary>
        public readonly string SshCloneUrl;
        /// <summary>
        /// URL that can be provided to `svn checkout` to check out the repository via GitHub's Subversion protocol emulation.
        /// </summary>
        public readonly string SvnUrl;
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
            bool allowMergeCommit,

            bool allowRebaseMerge,

            bool allowSquashMerge,

            bool archived,

            string defaultBranch,

            string description,

            string? fullName,

            string gitCloneUrl,

            bool hasDownloads,

            bool hasIssues,

            bool hasProjects,

            bool hasWiki,

            string homepageUrl,

            string htmlUrl,

            string httpCloneUrl,

            string id,

            string? name,

            string nodeId,

            ImmutableArray<Outputs.GetRepositoryPageResult> pages,

            bool @private,

            int repoId,

            string sshCloneUrl,

            string svnUrl,

            ImmutableArray<string> topics,

            string visibility)
        {
            AllowMergeCommit = allowMergeCommit;
            AllowRebaseMerge = allowRebaseMerge;
            AllowSquashMerge = allowSquashMerge;
            Archived = archived;
            DefaultBranch = defaultBranch;
            Description = description;
            FullName = fullName;
            GitCloneUrl = gitCloneUrl;
            HasDownloads = hasDownloads;
            HasIssues = hasIssues;
            HasProjects = hasProjects;
            HasWiki = hasWiki;
            HomepageUrl = homepageUrl;
            HtmlUrl = htmlUrl;
            HttpCloneUrl = httpCloneUrl;
            Id = id;
            Name = name;
            NodeId = nodeId;
            Pages = pages;
            Private = @private;
            RepoId = repoId;
            SshCloneUrl = sshCloneUrl;
            SvnUrl = svnUrl;
            Topics = topics;
            Visibility = visibility;
        }
    }
}
