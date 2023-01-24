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
    /// This resource allows you to set the access level of a non-public repositories actions and reusable workflows for use in other repositories.
    /// You must have admin access to a repository to use this resource.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Github = Pulumi.Github;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Github.Repository("example", new()
    ///     {
    ///         Visibility = "private",
    ///     });
    /// 
    ///     var test = new Github.ActionsRepositoryAccessLevel("test", new()
    ///     {
    ///         AccessLevel = "user",
    ///         Repository = example.Name,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// This resource can be imported using the name of the GitHub repository
    /// 
    /// ```sh
    ///  $ pulumi import github:index/actionsRepositoryAccessLevel:ActionsRepositoryAccessLevel test &lt;github_repository_name&gt;
    /// ```
    /// </summary>
    [GithubResourceType("github:index/actionsRepositoryAccessLevel:ActionsRepositoryAccessLevel")]
    public partial class ActionsRepositoryAccessLevel : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Where the actions or reusable workflows of the repository may be used. Possible values are `none`, `user`, `organization`, or `enterprise`.
        /// </summary>
        [Output("accessLevel")]
        public Output<string> AccessLevel { get; private set; } = null!;

        /// <summary>
        /// The GitHub repository
        /// </summary>
        [Output("repository")]
        public Output<string> Repository { get; private set; } = null!;


        /// <summary>
        /// Create a ActionsRepositoryAccessLevel resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ActionsRepositoryAccessLevel(string name, ActionsRepositoryAccessLevelArgs args, CustomResourceOptions? options = null)
            : base("github:index/actionsRepositoryAccessLevel:ActionsRepositoryAccessLevel", name, args ?? new ActionsRepositoryAccessLevelArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ActionsRepositoryAccessLevel(string name, Input<string> id, ActionsRepositoryAccessLevelState? state = null, CustomResourceOptions? options = null)
            : base("github:index/actionsRepositoryAccessLevel:ActionsRepositoryAccessLevel", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ActionsRepositoryAccessLevel resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ActionsRepositoryAccessLevel Get(string name, Input<string> id, ActionsRepositoryAccessLevelState? state = null, CustomResourceOptions? options = null)
        {
            return new ActionsRepositoryAccessLevel(name, id, state, options);
        }
    }

    public sealed class ActionsRepositoryAccessLevelArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Where the actions or reusable workflows of the repository may be used. Possible values are `none`, `user`, `organization`, or `enterprise`.
        /// </summary>
        [Input("accessLevel", required: true)]
        public Input<string> AccessLevel { get; set; } = null!;

        /// <summary>
        /// The GitHub repository
        /// </summary>
        [Input("repository", required: true)]
        public Input<string> Repository { get; set; } = null!;

        public ActionsRepositoryAccessLevelArgs()
        {
        }
        public static new ActionsRepositoryAccessLevelArgs Empty => new ActionsRepositoryAccessLevelArgs();
    }

    public sealed class ActionsRepositoryAccessLevelState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Where the actions or reusable workflows of the repository may be used. Possible values are `none`, `user`, `organization`, or `enterprise`.
        /// </summary>
        [Input("accessLevel")]
        public Input<string>? AccessLevel { get; set; }

        /// <summary>
        /// The GitHub repository
        /// </summary>
        [Input("repository")]
        public Input<string>? Repository { get; set; }

        public ActionsRepositoryAccessLevelState()
        {
        }
        public static new ActionsRepositoryAccessLevelState Empty => new ActionsRepositoryAccessLevelState();
    }
}
