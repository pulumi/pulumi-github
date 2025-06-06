// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github
{
    /// <summary>
    /// This resource allows you to create and manage a specific custom property for a GitHub repository.
    /// 
    /// ## Example Usage
    /// 
    /// &gt; Note that this assumes there already is a custom property defined on the org level called `my-cool-property` of type `string`
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
    ///     });
    /// 
    ///     var @string = new Github.RepositoryCustomProperty("string", new()
    ///     {
    ///         Repository = example.Name,
    ///         PropertyName = "my-cool-property",
    ///         PropertyType = "string",
    ///         PropertyValues = new[]
    ///         {
    ///             "test",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// GitHub Repository Custom Property can be imported using an ID made up of a comibnation of the names of the organization, repository, custom property separated by a `:` character, e.g.
    /// 
    /// ```sh
    /// $ pulumi import github:index/repositoryCustomProperty:RepositoryCustomProperty example &lt;organization-name&gt;:&lt;repo-name&gt;:&lt;custom-property-name&gt;
    /// ```
    /// </summary>
    [GithubResourceType("github:index/repositoryCustomProperty:RepositoryCustomProperty")]
    public partial class RepositoryCustomProperty : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Name of the custom property. Note that a pre-requisiste for this resource is that a custom property of this name has already been defined on the organization level
        /// </summary>
        [Output("propertyName")]
        public Output<string> PropertyName { get; private set; } = null!;

        /// <summary>
        /// Type of the custom property. Can be one of `single_select`, `multi_select`, `string`, or `true_false`
        /// </summary>
        [Output("propertyType")]
        public Output<string> PropertyType { get; private set; } = null!;

        /// <summary>
        /// Value of the custom property in the form of an array. Properties of type `single_select`, `string`, and `true_false` are represented as a string array of length 1
        /// </summary>
        [Output("propertyValues")]
        public Output<ImmutableArray<string>> PropertyValues { get; private set; } = null!;

        /// <summary>
        /// The repository of the environment.
        /// </summary>
        [Output("repository")]
        public Output<string> Repository { get; private set; } = null!;


        /// <summary>
        /// Create a RepositoryCustomProperty resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RepositoryCustomProperty(string name, RepositoryCustomPropertyArgs args, CustomResourceOptions? options = null)
            : base("github:index/repositoryCustomProperty:RepositoryCustomProperty", name, args ?? new RepositoryCustomPropertyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RepositoryCustomProperty(string name, Input<string> id, RepositoryCustomPropertyState? state = null, CustomResourceOptions? options = null)
            : base("github:index/repositoryCustomProperty:RepositoryCustomProperty", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RepositoryCustomProperty resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RepositoryCustomProperty Get(string name, Input<string> id, RepositoryCustomPropertyState? state = null, CustomResourceOptions? options = null)
        {
            return new RepositoryCustomProperty(name, id, state, options);
        }
    }

    public sealed class RepositoryCustomPropertyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the custom property. Note that a pre-requisiste for this resource is that a custom property of this name has already been defined on the organization level
        /// </summary>
        [Input("propertyName", required: true)]
        public Input<string> PropertyName { get; set; } = null!;

        /// <summary>
        /// Type of the custom property. Can be one of `single_select`, `multi_select`, `string`, or `true_false`
        /// </summary>
        [Input("propertyType", required: true)]
        public Input<string> PropertyType { get; set; } = null!;

        [Input("propertyValues", required: true)]
        private InputList<string>? _propertyValues;

        /// <summary>
        /// Value of the custom property in the form of an array. Properties of type `single_select`, `string`, and `true_false` are represented as a string array of length 1
        /// </summary>
        public InputList<string> PropertyValues
        {
            get => _propertyValues ?? (_propertyValues = new InputList<string>());
            set => _propertyValues = value;
        }

        /// <summary>
        /// The repository of the environment.
        /// </summary>
        [Input("repository", required: true)]
        public Input<string> Repository { get; set; } = null!;

        public RepositoryCustomPropertyArgs()
        {
        }
        public static new RepositoryCustomPropertyArgs Empty => new RepositoryCustomPropertyArgs();
    }

    public sealed class RepositoryCustomPropertyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the custom property. Note that a pre-requisiste for this resource is that a custom property of this name has already been defined on the organization level
        /// </summary>
        [Input("propertyName")]
        public Input<string>? PropertyName { get; set; }

        /// <summary>
        /// Type of the custom property. Can be one of `single_select`, `multi_select`, `string`, or `true_false`
        /// </summary>
        [Input("propertyType")]
        public Input<string>? PropertyType { get; set; }

        [Input("propertyValues")]
        private InputList<string>? _propertyValues;

        /// <summary>
        /// Value of the custom property in the form of an array. Properties of type `single_select`, `string`, and `true_false` are represented as a string array of length 1
        /// </summary>
        public InputList<string> PropertyValues
        {
            get => _propertyValues ?? (_propertyValues = new InputList<string>());
            set => _propertyValues = value;
        }

        /// <summary>
        /// The repository of the environment.
        /// </summary>
        [Input("repository")]
        public Input<string>? Repository { get; set; }

        public RepositoryCustomPropertyState()
        {
        }
        public static new RepositoryCustomPropertyState Empty => new RepositoryCustomPropertyState();
    }
}
