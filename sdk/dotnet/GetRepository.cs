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
        public static Task<GetRepositoryResult> InvokeAsync(GetRepositoryArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRepositoryResult>("github:index/getRepository:getRepository", args ?? new GetRepositoryArgs(), options.WithDefaults());

        public static Output<GetRepositoryResult> Invoke(GetRepositoryInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetRepositoryResult>("github:index/getRepository:getRepository", args ?? new GetRepositoryInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRepositoryArgs : global::Pulumi.InvokeArgs
    {
        [Input("description")]
        public string? Description { get; set; }

        [Input("fullName")]
        public string? FullName { get; set; }

        [Input("homepageUrl")]
        public string? HomepageUrl { get; set; }

        [Input("name")]
        public string? Name { get; set; }

        [Input("onlyProtectedBranches")]
        public bool? OnlyProtectedBranches { get; set; }

        public GetRepositoryArgs()
        {
        }
        public static new GetRepositoryArgs Empty => new GetRepositoryArgs();
    }

    public sealed class GetRepositoryInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("fullName")]
        public Input<string>? FullName { get; set; }

        [Input("homepageUrl")]
        public Input<string>? HomepageUrl { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("onlyProtectedBranches")]
        public Input<bool>? OnlyProtectedBranches { get; set; }

        public GetRepositoryInvokeArgs()
        {
        }
        public static new GetRepositoryInvokeArgs Empty => new GetRepositoryInvokeArgs();
    }


    [OutputType]
    public sealed class GetRepositoryResult
    {
        public readonly bool AllowAutoMerge;
        public readonly bool AllowMergeCommit;
        public readonly bool AllowRebaseMerge;
        public readonly bool AllowSquashMerge;
        public readonly bool Archived;
        public readonly string DefaultBranch;
        public readonly string? Description;
        public readonly string FullName;
        public readonly string GitCloneUrl;
        public readonly bool HasDownloads;
        public readonly bool HasIssues;
        public readonly bool HasProjects;
        public readonly bool HasWiki;
        public readonly string? HomepageUrl;
        public readonly string HtmlUrl;
        public readonly string HttpCloneUrl;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string MergeCommitMessage;
        public readonly string MergeCommitTitle;
        public readonly string Name;
        public readonly string NodeId;
        public readonly bool? OnlyProtectedBranches;
        public readonly ImmutableArray<Outputs.GetRepositoryPageResult> Pages;
        public readonly bool Private;
        public readonly int RepoId;
        public readonly string SquashMergeCommitMessage;
        public readonly string SquashMergeCommitTitle;
        public readonly string SshCloneUrl;
        public readonly string SvnUrl;
        public readonly ImmutableArray<string> Topics;
        public readonly string Visibility;

        [OutputConstructor]
        private GetRepositoryResult(
            bool allowAutoMerge,

            bool allowMergeCommit,

            bool allowRebaseMerge,

            bool allowSquashMerge,

            bool archived,

            string defaultBranch,

            string? description,

            string fullName,

            string gitCloneUrl,

            bool hasDownloads,

            bool hasIssues,

            bool hasProjects,

            bool hasWiki,

            string? homepageUrl,

            string htmlUrl,

            string httpCloneUrl,

            string id,

            string mergeCommitMessage,

            string mergeCommitTitle,

            string name,

            string nodeId,

            bool? onlyProtectedBranches,

            ImmutableArray<Outputs.GetRepositoryPageResult> pages,

            bool @private,

            int repoId,

            string squashMergeCommitMessage,

            string squashMergeCommitTitle,

            string sshCloneUrl,

            string svnUrl,

            ImmutableArray<string> topics,

            string visibility)
        {
            AllowAutoMerge = allowAutoMerge;
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
            MergeCommitMessage = mergeCommitMessage;
            MergeCommitTitle = mergeCommitTitle;
            Name = name;
            NodeId = nodeId;
            OnlyProtectedBranches = onlyProtectedBranches;
            Pages = pages;
            Private = @private;
            RepoId = repoId;
            SquashMergeCommitMessage = squashMergeCommitMessage;
            SquashMergeCommitTitle = squashMergeCommitTitle;
            SshCloneUrl = sshCloneUrl;
            SvnUrl = svnUrl;
            Topics = topics;
            Visibility = visibility;
        }
    }
}
