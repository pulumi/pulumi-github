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
    ///     var testRepository = Github.GetRepository.Invoke(new()
    ///     {
    ///         Name = "test",
    ///     });
    /// 
    ///     var testRepositoryTopics = new Github.RepositoryTopics("testRepositoryTopics", new()
    ///     {
    ///         Repository = github_repository.Test.Name,
    ///         Topics = new[]
    ///         {
    ///             "topic-1",
    ///             "topic-2",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Repository topics can be imported using the `name` of the repository.
    /// 
    /// ```sh
    ///  $ pulumi import github:index/repositoryTopics:RepositoryTopics terraform terraform
    /// ```
    /// </summary>
    [GithubResourceType("github:index/repositoryTopics:RepositoryTopics")]
    public partial class RepositoryTopics : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The repository name.
        /// </summary>
        [Output("repository")]
        public Output<string> Repository { get; private set; } = null!;

        /// <summary>
        /// A list of topics to add to the repository.
        /// </summary>
        [Output("topics")]
        public Output<ImmutableArray<string>> Topics { get; private set; } = null!;


        /// <summary>
        /// Create a RepositoryTopics resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RepositoryTopics(string name, RepositoryTopicsArgs args, CustomResourceOptions? options = null)
            : base("github:index/repositoryTopics:RepositoryTopics", name, args ?? new RepositoryTopicsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RepositoryTopics(string name, Input<string> id, RepositoryTopicsState? state = null, CustomResourceOptions? options = null)
            : base("github:index/repositoryTopics:RepositoryTopics", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RepositoryTopics resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RepositoryTopics Get(string name, Input<string> id, RepositoryTopicsState? state = null, CustomResourceOptions? options = null)
        {
            return new RepositoryTopics(name, id, state, options);
        }
    }

    public sealed class RepositoryTopicsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The repository name.
        /// </summary>
        [Input("repository", required: true)]
        public Input<string> Repository { get; set; } = null!;

        [Input("topics", required: true)]
        private InputList<string>? _topics;

        /// <summary>
        /// A list of topics to add to the repository.
        /// </summary>
        public InputList<string> Topics
        {
            get => _topics ?? (_topics = new InputList<string>());
            set => _topics = value;
        }

        public RepositoryTopicsArgs()
        {
        }
        public static new RepositoryTopicsArgs Empty => new RepositoryTopicsArgs();
    }

    public sealed class RepositoryTopicsState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The repository name.
        /// </summary>
        [Input("repository")]
        public Input<string>? Repository { get; set; }

        [Input("topics")]
        private InputList<string>? _topics;

        /// <summary>
        /// A list of topics to add to the repository.
        /// </summary>
        public InputList<string> Topics
        {
            get => _topics ?? (_topics = new InputList<string>());
            set => _topics = value;
        }

        public RepositoryTopicsState()
        {
        }
        public static new RepositoryTopicsState Empty => new RepositoryTopicsState();
    }
}
