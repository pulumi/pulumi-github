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
    /// This resource allows you to create and manage projects for GitHub repository.
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
    ///         Description = "My awesome codebase",
    ///         HasProjects = true,
    ///     });
    /// 
    ///     var project = new Github.RepositoryProject("project", new()
    ///     {
    ///         Body = "This is a repository project.",
    ///         Repository = example.Name,
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [GithubResourceType("github:index/repositoryProject:RepositoryProject")]
    public partial class RepositoryProject : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The body of the project.
        /// </summary>
        [Output("body")]
        public Output<string?> Body { get; private set; } = null!;

        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// The name of the project.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The repository of the project.
        /// </summary>
        [Output("repository")]
        public Output<string> Repository { get; private set; } = null!;

        /// <summary>
        /// URL of the project
        /// </summary>
        [Output("url")]
        public Output<string> Url { get; private set; } = null!;


        /// <summary>
        /// Create a RepositoryProject resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RepositoryProject(string name, RepositoryProjectArgs args, CustomResourceOptions? options = null)
            : base("github:index/repositoryProject:RepositoryProject", name, args ?? new RepositoryProjectArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RepositoryProject(string name, Input<string> id, RepositoryProjectState? state = null, CustomResourceOptions? options = null)
            : base("github:index/repositoryProject:RepositoryProject", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RepositoryProject resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RepositoryProject Get(string name, Input<string> id, RepositoryProjectState? state = null, CustomResourceOptions? options = null)
        {
            return new RepositoryProject(name, id, state, options);
        }
    }

    public sealed class RepositoryProjectArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The body of the project.
        /// </summary>
        [Input("body")]
        public Input<string>? Body { get; set; }

        /// <summary>
        /// The name of the project.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The repository of the project.
        /// </summary>
        [Input("repository", required: true)]
        public Input<string> Repository { get; set; } = null!;

        public RepositoryProjectArgs()
        {
        }
        public static new RepositoryProjectArgs Empty => new RepositoryProjectArgs();
    }

    public sealed class RepositoryProjectState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The body of the project.
        /// </summary>
        [Input("body")]
        public Input<string>? Body { get; set; }

        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// The name of the project.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The repository of the project.
        /// </summary>
        [Input("repository")]
        public Input<string>? Repository { get; set; }

        /// <summary>
        /// URL of the project
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        public RepositoryProjectState()
        {
        }
        public static new RepositoryProjectState Empty => new RepositoryProjectState();
    }
}
