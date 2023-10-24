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
    /// This resource allows you to create and manage branches within your repository.
    /// 
    /// Additional constraints can be applied to ensure your branch is created from
    /// another branch or commit.
    /// 
    /// ## Import
    /// 
    /// GitHub Branch can be imported using an ID made up of `repository:branch`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import github:index/branch:Branch terraform terraform:main
    /// ```
    ///  Importing github branch into an instance object (when using a for each block to manage multiple branches)
    /// 
    /// ```sh
    ///  $ pulumi import github:index/branch:Branch terraform["terraform"] terraform:main
    /// ```
    ///  Optionally, a source branch may be specified using an ID of `repository:branch:source_branch`. This is useful for importing branches that do not branch directly off main.
    /// 
    /// ```sh
    ///  $ pulumi import github:index/branch:Branch terraform terraform:feature-branch:dev
    /// ```
    /// </summary>
    [GithubResourceType("github:index/branch:Branch")]
    public partial class Branch : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The repository branch to create.
        /// </summary>
        [Output("branch")]
        public Output<string> BranchName { get; private set; } = null!;

        /// <summary>
        /// An etag representing the Branch object.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// A string representing a branch reference, in the form of `refs/heads/&lt;branch&gt;`.
        /// </summary>
        [Output("ref")]
        public Output<string> Ref { get; private set; } = null!;

        /// <summary>
        /// The GitHub repository name.
        /// </summary>
        [Output("repository")]
        public Output<string> Repository { get; private set; } = null!;

        /// <summary>
        /// A string storing the reference's `HEAD` commit's SHA1.
        /// </summary>
        [Output("sha")]
        public Output<string> Sha { get; private set; } = null!;

        /// <summary>
        /// The branch name to start from. Defaults to `main`.
        /// </summary>
        [Output("sourceBranch")]
        public Output<string?> SourceBranch { get; private set; } = null!;

        /// <summary>
        /// The commit hash to start from. Defaults to the tip of `source_branch`. If provided, `source_branch` is ignored.
        /// </summary>
        [Output("sourceSha")]
        public Output<string> SourceSha { get; private set; } = null!;


        /// <summary>
        /// Create a Branch resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Branch(string name, BranchArgs args, CustomResourceOptions? options = null)
            : base("github:index/branch:Branch", name, args ?? new BranchArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Branch(string name, Input<string> id, BranchState? state = null, CustomResourceOptions? options = null)
            : base("github:index/branch:Branch", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Branch resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Branch Get(string name, Input<string> id, BranchState? state = null, CustomResourceOptions? options = null)
        {
            return new Branch(name, id, state, options);
        }
    }

    public sealed class BranchArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The repository branch to create.
        /// </summary>
        [Input("branch", required: true)]
        public Input<string> BranchName { get; set; } = null!;

        /// <summary>
        /// The GitHub repository name.
        /// </summary>
        [Input("repository", required: true)]
        public Input<string> Repository { get; set; } = null!;

        /// <summary>
        /// The branch name to start from. Defaults to `main`.
        /// </summary>
        [Input("sourceBranch")]
        public Input<string>? SourceBranch { get; set; }

        /// <summary>
        /// The commit hash to start from. Defaults to the tip of `source_branch`. If provided, `source_branch` is ignored.
        /// </summary>
        [Input("sourceSha")]
        public Input<string>? SourceSha { get; set; }

        public BranchArgs()
        {
        }
        public static new BranchArgs Empty => new BranchArgs();
    }

    public sealed class BranchState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The repository branch to create.
        /// </summary>
        [Input("branch")]
        public Input<string>? BranchName { get; set; }

        /// <summary>
        /// An etag representing the Branch object.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// A string representing a branch reference, in the form of `refs/heads/&lt;branch&gt;`.
        /// </summary>
        [Input("ref")]
        public Input<string>? Ref { get; set; }

        /// <summary>
        /// The GitHub repository name.
        /// </summary>
        [Input("repository")]
        public Input<string>? Repository { get; set; }

        /// <summary>
        /// A string storing the reference's `HEAD` commit's SHA1.
        /// </summary>
        [Input("sha")]
        public Input<string>? Sha { get; set; }

        /// <summary>
        /// The branch name to start from. Defaults to `main`.
        /// </summary>
        [Input("sourceBranch")]
        public Input<string>? SourceBranch { get; set; }

        /// <summary>
        /// The commit hash to start from. Defaults to the tip of `source_branch`. If provided, `source_branch` is ignored.
        /// </summary>
        [Input("sourceSha")]
        public Input<string>? SourceSha { get; set; }

        public BranchState()
        {
        }
        public static new BranchState Empty => new BranchState();
    }
}
