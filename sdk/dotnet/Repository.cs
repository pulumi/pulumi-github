// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github
{
    [GithubResourceType("github:index/repository:Repository")]
    public partial class Repository : global::Pulumi.CustomResource
    {
        [Output("allowAutoMerge")]
        public Output<bool?> AllowAutoMerge { get; private set; } = null!;

        [Output("allowMergeCommit")]
        public Output<bool?> AllowMergeCommit { get; private set; } = null!;

        [Output("allowRebaseMerge")]
        public Output<bool?> AllowRebaseMerge { get; private set; } = null!;

        [Output("allowSquashMerge")]
        public Output<bool?> AllowSquashMerge { get; private set; } = null!;

        [Output("allowUpdateBranch")]
        public Output<bool?> AllowUpdateBranch { get; private set; } = null!;

        [Output("archiveOnDestroy")]
        public Output<bool?> ArchiveOnDestroy { get; private set; } = null!;

        [Output("archived")]
        public Output<bool?> Archived { get; private set; } = null!;

        [Output("autoInit")]
        public Output<bool?> AutoInit { get; private set; } = null!;

        /// <summary>
        /// Can only be set after initial repository creation, and only if the target branch exists
        /// </summary>
        [Output("defaultBranch")]
        public Output<string> DefaultBranch { get; private set; } = null!;

        [Output("deleteBranchOnMerge")]
        public Output<bool?> DeleteBranchOnMerge { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        [Output("fullName")]
        public Output<string> FullName { get; private set; } = null!;

        [Output("gitCloneUrl")]
        public Output<string> GitCloneUrl { get; private set; } = null!;

        [Output("gitignoreTemplate")]
        public Output<string?> GitignoreTemplate { get; private set; } = null!;

        [Output("hasDownloads")]
        public Output<bool?> HasDownloads { get; private set; } = null!;

        [Output("hasIssues")]
        public Output<bool?> HasIssues { get; private set; } = null!;

        [Output("hasProjects")]
        public Output<bool?> HasProjects { get; private set; } = null!;

        [Output("hasWiki")]
        public Output<bool?> HasWiki { get; private set; } = null!;

        [Output("homepageUrl")]
        public Output<string?> HomepageUrl { get; private set; } = null!;

        [Output("htmlUrl")]
        public Output<string> HtmlUrl { get; private set; } = null!;

        [Output("httpCloneUrl")]
        public Output<string> HttpCloneUrl { get; private set; } = null!;

        [Output("ignoreVulnerabilityAlertsDuringRead")]
        public Output<bool?> IgnoreVulnerabilityAlertsDuringRead { get; private set; } = null!;

        [Output("isTemplate")]
        public Output<bool?> IsTemplate { get; private set; } = null!;

        [Output("licenseTemplate")]
        public Output<string?> LicenseTemplate { get; private set; } = null!;

        [Output("mergeCommitMessage")]
        public Output<string?> MergeCommitMessage { get; private set; } = null!;

        [Output("mergeCommitTitle")]
        public Output<string?> MergeCommitTitle { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("nodeId")]
        public Output<string> NodeId { get; private set; } = null!;

        [Output("pages")]
        public Output<Outputs.RepositoryPages?> Pages { get; private set; } = null!;

        [Output("private")]
        public Output<bool> Private { get; private set; } = null!;

        [Output("repoId")]
        public Output<int> RepoId { get; private set; } = null!;

        [Output("squashMergeCommitMessage")]
        public Output<string?> SquashMergeCommitMessage { get; private set; } = null!;

        [Output("squashMergeCommitTitle")]
        public Output<string?> SquashMergeCommitTitle { get; private set; } = null!;

        [Output("sshCloneUrl")]
        public Output<string> SshCloneUrl { get; private set; } = null!;

        [Output("svnUrl")]
        public Output<string> SvnUrl { get; private set; } = null!;

        [Output("template")]
        public Output<Outputs.RepositoryTemplate?> Template { get; private set; } = null!;

        [Output("topics")]
        public Output<ImmutableArray<string>> Topics { get; private set; } = null!;

        [Output("visibility")]
        public Output<string> Visibility { get; private set; } = null!;

        [Output("vulnerabilityAlerts")]
        public Output<bool?> VulnerabilityAlerts { get; private set; } = null!;


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
        [Input("allowAutoMerge")]
        public Input<bool>? AllowAutoMerge { get; set; }

        [Input("allowMergeCommit")]
        public Input<bool>? AllowMergeCommit { get; set; }

        [Input("allowRebaseMerge")]
        public Input<bool>? AllowRebaseMerge { get; set; }

        [Input("allowSquashMerge")]
        public Input<bool>? AllowSquashMerge { get; set; }

        [Input("allowUpdateBranch")]
        public Input<bool>? AllowUpdateBranch { get; set; }

        [Input("archiveOnDestroy")]
        public Input<bool>? ArchiveOnDestroy { get; set; }

        [Input("archived")]
        public Input<bool>? Archived { get; set; }

        [Input("autoInit")]
        public Input<bool>? AutoInit { get; set; }

        /// <summary>
        /// Can only be set after initial repository creation, and only if the target branch exists
        /// </summary>
        [Input("defaultBranch")]
        public Input<string>? DefaultBranch { get; set; }

        [Input("deleteBranchOnMerge")]
        public Input<bool>? DeleteBranchOnMerge { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("gitignoreTemplate")]
        public Input<string>? GitignoreTemplate { get; set; }

        [Input("hasDownloads")]
        public Input<bool>? HasDownloads { get; set; }

        [Input("hasIssues")]
        public Input<bool>? HasIssues { get; set; }

        [Input("hasProjects")]
        public Input<bool>? HasProjects { get; set; }

        [Input("hasWiki")]
        public Input<bool>? HasWiki { get; set; }

        [Input("homepageUrl")]
        public Input<string>? HomepageUrl { get; set; }

        [Input("ignoreVulnerabilityAlertsDuringRead")]
        public Input<bool>? IgnoreVulnerabilityAlertsDuringRead { get; set; }

        [Input("isTemplate")]
        public Input<bool>? IsTemplate { get; set; }

        [Input("licenseTemplate")]
        public Input<string>? LicenseTemplate { get; set; }

        [Input("mergeCommitMessage")]
        public Input<string>? MergeCommitMessage { get; set; }

        [Input("mergeCommitTitle")]
        public Input<string>? MergeCommitTitle { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("pages")]
        public Input<Inputs.RepositoryPagesArgs>? Pages { get; set; }

        [Input("private")]
        public Input<bool>? Private { get; set; }

        [Input("squashMergeCommitMessage")]
        public Input<string>? SquashMergeCommitMessage { get; set; }

        [Input("squashMergeCommitTitle")]
        public Input<string>? SquashMergeCommitTitle { get; set; }

        [Input("template")]
        public Input<Inputs.RepositoryTemplateArgs>? Template { get; set; }

        [Input("topics")]
        private InputList<string>? _topics;
        public InputList<string> Topics
        {
            get => _topics ?? (_topics = new InputList<string>());
            set => _topics = value;
        }

        [Input("visibility")]
        public Input<string>? Visibility { get; set; }

        [Input("vulnerabilityAlerts")]
        public Input<bool>? VulnerabilityAlerts { get; set; }

        public RepositoryArgs()
        {
        }
        public static new RepositoryArgs Empty => new RepositoryArgs();
    }

    public sealed class RepositoryState : global::Pulumi.ResourceArgs
    {
        [Input("allowAutoMerge")]
        public Input<bool>? AllowAutoMerge { get; set; }

        [Input("allowMergeCommit")]
        public Input<bool>? AllowMergeCommit { get; set; }

        [Input("allowRebaseMerge")]
        public Input<bool>? AllowRebaseMerge { get; set; }

        [Input("allowSquashMerge")]
        public Input<bool>? AllowSquashMerge { get; set; }

        [Input("allowUpdateBranch")]
        public Input<bool>? AllowUpdateBranch { get; set; }

        [Input("archiveOnDestroy")]
        public Input<bool>? ArchiveOnDestroy { get; set; }

        [Input("archived")]
        public Input<bool>? Archived { get; set; }

        [Input("autoInit")]
        public Input<bool>? AutoInit { get; set; }

        /// <summary>
        /// Can only be set after initial repository creation, and only if the target branch exists
        /// </summary>
        [Input("defaultBranch")]
        public Input<string>? DefaultBranch { get; set; }

        [Input("deleteBranchOnMerge")]
        public Input<bool>? DeleteBranchOnMerge { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("etag")]
        public Input<string>? Etag { get; set; }

        [Input("fullName")]
        public Input<string>? FullName { get; set; }

        [Input("gitCloneUrl")]
        public Input<string>? GitCloneUrl { get; set; }

        [Input("gitignoreTemplate")]
        public Input<string>? GitignoreTemplate { get; set; }

        [Input("hasDownloads")]
        public Input<bool>? HasDownloads { get; set; }

        [Input("hasIssues")]
        public Input<bool>? HasIssues { get; set; }

        [Input("hasProjects")]
        public Input<bool>? HasProjects { get; set; }

        [Input("hasWiki")]
        public Input<bool>? HasWiki { get; set; }

        [Input("homepageUrl")]
        public Input<string>? HomepageUrl { get; set; }

        [Input("htmlUrl")]
        public Input<string>? HtmlUrl { get; set; }

        [Input("httpCloneUrl")]
        public Input<string>? HttpCloneUrl { get; set; }

        [Input("ignoreVulnerabilityAlertsDuringRead")]
        public Input<bool>? IgnoreVulnerabilityAlertsDuringRead { get; set; }

        [Input("isTemplate")]
        public Input<bool>? IsTemplate { get; set; }

        [Input("licenseTemplate")]
        public Input<string>? LicenseTemplate { get; set; }

        [Input("mergeCommitMessage")]
        public Input<string>? MergeCommitMessage { get; set; }

        [Input("mergeCommitTitle")]
        public Input<string>? MergeCommitTitle { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("nodeId")]
        public Input<string>? NodeId { get; set; }

        [Input("pages")]
        public Input<Inputs.RepositoryPagesGetArgs>? Pages { get; set; }

        [Input("private")]
        public Input<bool>? Private { get; set; }

        [Input("repoId")]
        public Input<int>? RepoId { get; set; }

        [Input("squashMergeCommitMessage")]
        public Input<string>? SquashMergeCommitMessage { get; set; }

        [Input("squashMergeCommitTitle")]
        public Input<string>? SquashMergeCommitTitle { get; set; }

        [Input("sshCloneUrl")]
        public Input<string>? SshCloneUrl { get; set; }

        [Input("svnUrl")]
        public Input<string>? SvnUrl { get; set; }

        [Input("template")]
        public Input<Inputs.RepositoryTemplateGetArgs>? Template { get; set; }

        [Input("topics")]
        private InputList<string>? _topics;
        public InputList<string> Topics
        {
            get => _topics ?? (_topics = new InputList<string>());
            set => _topics = value;
        }

        [Input("visibility")]
        public Input<string>? Visibility { get; set; }

        [Input("vulnerabilityAlerts")]
        public Input<bool>? VulnerabilityAlerts { get; set; }

        public RepositoryState()
        {
        }
        public static new RepositoryState Empty => new RepositoryState();
    }
}
