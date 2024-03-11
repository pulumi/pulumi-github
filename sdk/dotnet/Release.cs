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
    /// This resource allows you to create and manage a release in a specific
    /// GitHub repository.
    /// 
    /// ## Example Usage
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Github = Pulumi.Github;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var repo = new Github.Repository("repo", new()
    ///     {
    ///         Description = "GitHub repo managed by Terraform",
    ///         Private = false,
    ///     });
    /// 
    ///     var example = new Github.Release("example", new()
    ///     {
    ///         Repository = repo.Name,
    ///         TagName = "v1.0.0",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ### On Non-Default Branch
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Github = Pulumi.Github;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var exampleRepository = new Github.Repository("exampleRepository", new()
    ///     {
    ///         AutoInit = true,
    ///     });
    /// 
    ///     var exampleBranch = new Github.Branch("exampleBranch", new()
    ///     {
    ///         Repository = exampleRepository.Name,
    ///         BranchName = "branch_name",
    ///         SourceBranch = exampleRepository.DefaultBranch,
    ///     });
    /// 
    ///     var exampleRelease = new Github.Release("exampleRelease", new()
    ///     {
    ///         Repository = exampleRepository.Name,
    ///         TagName = "v1.0.0",
    ///         TargetCommitish = exampleBranch.BranchName,
    ///         Draft = false,
    ///         Prerelease = false,
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// This resource can be imported using the `name` of the repository, combined with the `id` of the release, and a `:` character for separating components, e.g.
    /// 
    /// ```sh
    /// $ pulumi import github:index/release:Release example repo:12345678
    /// ```
    /// </summary>
    [GithubResourceType("github:index/release:Release")]
    public partial class Release : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Text describing the contents of the tag.
        /// </summary>
        [Output("body")]
        public Output<string?> Body { get; private set; } = null!;

        /// <summary>
        /// If specified, a discussion of the specified category is created and linked to the release. The value must be a category that already exists in the repository. For more information, see [Managing categories for discussions in your repository](https://docs.github.com/discussions/managing-discussions-for-your-community/managing-categories-for-discussions-in-your-repository).
        /// </summary>
        [Output("discussionCategoryName")]
        public Output<string?> DiscussionCategoryName { get; private set; } = null!;

        /// <summary>
        /// Set to `false` to create a published release.
        /// </summary>
        [Output("draft")]
        public Output<bool?> Draft { get; private set; } = null!;

        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// Set to `true` to automatically generate the name and body for this release. If `name` is specified, the specified `name` will be used; otherwise, a name will be automatically generated. If `body` is specified, the `body` will be pre-pended to the automatically generated notes.
        /// </summary>
        [Output("generateReleaseNotes")]
        public Output<bool?> GenerateReleaseNotes { get; private set; } = null!;

        /// <summary>
        /// The name of the release.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Set to `false` to identify the release as a full release.
        /// </summary>
        [Output("prerelease")]
        public Output<bool?> Prerelease { get; private set; } = null!;

        /// <summary>
        /// The name of the repository.
        /// </summary>
        [Output("repository")]
        public Output<string> Repository { get; private set; } = null!;

        /// <summary>
        /// The name of the tag.
        /// </summary>
        [Output("tagName")]
        public Output<string> TagName { get; private set; } = null!;

        /// <summary>
        /// The branch name or commit SHA the tag is created from. Defaults to the default branch of the repository.
        /// </summary>
        [Output("targetCommitish")]
        public Output<string?> TargetCommitish { get; private set; } = null!;


        /// <summary>
        /// Create a Release resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Release(string name, ReleaseArgs args, CustomResourceOptions? options = null)
            : base("github:index/release:Release", name, args ?? new ReleaseArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Release(string name, Input<string> id, ReleaseState? state = null, CustomResourceOptions? options = null)
            : base("github:index/release:Release", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Release resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Release Get(string name, Input<string> id, ReleaseState? state = null, CustomResourceOptions? options = null)
        {
            return new Release(name, id, state, options);
        }
    }

    public sealed class ReleaseArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Text describing the contents of the tag.
        /// </summary>
        [Input("body")]
        public Input<string>? Body { get; set; }

        /// <summary>
        /// If specified, a discussion of the specified category is created and linked to the release. The value must be a category that already exists in the repository. For more information, see [Managing categories for discussions in your repository](https://docs.github.com/discussions/managing-discussions-for-your-community/managing-categories-for-discussions-in-your-repository).
        /// </summary>
        [Input("discussionCategoryName")]
        public Input<string>? DiscussionCategoryName { get; set; }

        /// <summary>
        /// Set to `false` to create a published release.
        /// </summary>
        [Input("draft")]
        public Input<bool>? Draft { get; set; }

        /// <summary>
        /// Set to `true` to automatically generate the name and body for this release. If `name` is specified, the specified `name` will be used; otherwise, a name will be automatically generated. If `body` is specified, the `body` will be pre-pended to the automatically generated notes.
        /// </summary>
        [Input("generateReleaseNotes")]
        public Input<bool>? GenerateReleaseNotes { get; set; }

        /// <summary>
        /// The name of the release.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Set to `false` to identify the release as a full release.
        /// </summary>
        [Input("prerelease")]
        public Input<bool>? Prerelease { get; set; }

        /// <summary>
        /// The name of the repository.
        /// </summary>
        [Input("repository", required: true)]
        public Input<string> Repository { get; set; } = null!;

        /// <summary>
        /// The name of the tag.
        /// </summary>
        [Input("tagName", required: true)]
        public Input<string> TagName { get; set; } = null!;

        /// <summary>
        /// The branch name or commit SHA the tag is created from. Defaults to the default branch of the repository.
        /// </summary>
        [Input("targetCommitish")]
        public Input<string>? TargetCommitish { get; set; }

        public ReleaseArgs()
        {
        }
        public static new ReleaseArgs Empty => new ReleaseArgs();
    }

    public sealed class ReleaseState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Text describing the contents of the tag.
        /// </summary>
        [Input("body")]
        public Input<string>? Body { get; set; }

        /// <summary>
        /// If specified, a discussion of the specified category is created and linked to the release. The value must be a category that already exists in the repository. For more information, see [Managing categories for discussions in your repository](https://docs.github.com/discussions/managing-discussions-for-your-community/managing-categories-for-discussions-in-your-repository).
        /// </summary>
        [Input("discussionCategoryName")]
        public Input<string>? DiscussionCategoryName { get; set; }

        /// <summary>
        /// Set to `false` to create a published release.
        /// </summary>
        [Input("draft")]
        public Input<bool>? Draft { get; set; }

        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// Set to `true` to automatically generate the name and body for this release. If `name` is specified, the specified `name` will be used; otherwise, a name will be automatically generated. If `body` is specified, the `body` will be pre-pended to the automatically generated notes.
        /// </summary>
        [Input("generateReleaseNotes")]
        public Input<bool>? GenerateReleaseNotes { get; set; }

        /// <summary>
        /// The name of the release.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Set to `false` to identify the release as a full release.
        /// </summary>
        [Input("prerelease")]
        public Input<bool>? Prerelease { get; set; }

        /// <summary>
        /// The name of the repository.
        /// </summary>
        [Input("repository")]
        public Input<string>? Repository { get; set; }

        /// <summary>
        /// The name of the tag.
        /// </summary>
        [Input("tagName")]
        public Input<string>? TagName { get; set; }

        /// <summary>
        /// The branch name or commit SHA the tag is created from. Defaults to the default branch of the repository.
        /// </summary>
        [Input("targetCommitish")]
        public Input<string>? TargetCommitish { get; set; }

        public ReleaseState()
        {
        }
        public static new ReleaseState Empty => new ReleaseState();
    }
}
