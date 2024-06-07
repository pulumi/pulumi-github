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
    /// This resource allows you to create and manage GitHub Actions variables within your GitHub organization.
    /// You must have write access to a repository to use this resource.
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
    ///     var exampleVariable = new Github.ActionsOrganizationVariable("example_variable", new()
    ///     {
    ///         VariableName = "example_variable_name",
    ///         Visibility = "private",
    ///         Value = "example_variable_value",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Github = Pulumi.Github;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var repo = Github.GetRepository.Invoke(new()
    ///     {
    ///         FullName = "my-org/repo",
    ///     });
    /// 
    ///     var exampleVariable = new Github.ActionsOrganizationVariable("example_variable", new()
    ///     {
    ///         VariableName = "example_variable_name",
    ///         Visibility = "selected",
    ///         Value = "example_variable_value",
    ///         SelectedRepositoryIds = new[]
    ///         {
    ///             repo.Apply(getRepositoryResult =&gt; getRepositoryResult.RepoId),
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// This resource can be imported using an ID made up of the variable name:
    /// 
    /// ```sh
    /// $ pulumi import github:index/actionsOrganizationVariable:ActionsOrganizationVariable test_variable test_variable_name
    /// ```
    /// </summary>
    [GithubResourceType("github:index/actionsOrganizationVariable:ActionsOrganizationVariable")]
    public partial class ActionsOrganizationVariable : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Date of actions_variable creation.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// An array of repository ids that can access the organization variable.
        /// </summary>
        [Output("selectedRepositoryIds")]
        public Output<ImmutableArray<int>> SelectedRepositoryIds { get; private set; } = null!;

        /// <summary>
        /// Date of actions_variable update.
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;

        /// <summary>
        /// Value of the variable
        /// </summary>
        [Output("value")]
        public Output<string> Value { get; private set; } = null!;

        /// <summary>
        /// Name of the variable
        /// </summary>
        [Output("variableName")]
        public Output<string> VariableName { get; private set; } = null!;

        /// <summary>
        /// Configures the access that repositories have to the organization variable.
        /// Must be one of `all`, `private`, `selected`. `selected_repository_ids` is required if set to `selected`.
        /// </summary>
        [Output("visibility")]
        public Output<string> Visibility { get; private set; } = null!;


        /// <summary>
        /// Create a ActionsOrganizationVariable resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ActionsOrganizationVariable(string name, ActionsOrganizationVariableArgs args, CustomResourceOptions? options = null)
            : base("github:index/actionsOrganizationVariable:ActionsOrganizationVariable", name, args ?? new ActionsOrganizationVariableArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ActionsOrganizationVariable(string name, Input<string> id, ActionsOrganizationVariableState? state = null, CustomResourceOptions? options = null)
            : base("github:index/actionsOrganizationVariable:ActionsOrganizationVariable", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ActionsOrganizationVariable resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ActionsOrganizationVariable Get(string name, Input<string> id, ActionsOrganizationVariableState? state = null, CustomResourceOptions? options = null)
        {
            return new ActionsOrganizationVariable(name, id, state, options);
        }
    }

    public sealed class ActionsOrganizationVariableArgs : global::Pulumi.ResourceArgs
    {
        [Input("selectedRepositoryIds")]
        private InputList<int>? _selectedRepositoryIds;

        /// <summary>
        /// An array of repository ids that can access the organization variable.
        /// </summary>
        public InputList<int> SelectedRepositoryIds
        {
            get => _selectedRepositoryIds ?? (_selectedRepositoryIds = new InputList<int>());
            set => _selectedRepositoryIds = value;
        }

        /// <summary>
        /// Value of the variable
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        /// <summary>
        /// Name of the variable
        /// </summary>
        [Input("variableName", required: true)]
        public Input<string> VariableName { get; set; } = null!;

        /// <summary>
        /// Configures the access that repositories have to the organization variable.
        /// Must be one of `all`, `private`, `selected`. `selected_repository_ids` is required if set to `selected`.
        /// </summary>
        [Input("visibility", required: true)]
        public Input<string> Visibility { get; set; } = null!;

        public ActionsOrganizationVariableArgs()
        {
        }
        public static new ActionsOrganizationVariableArgs Empty => new ActionsOrganizationVariableArgs();
    }

    public sealed class ActionsOrganizationVariableState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Date of actions_variable creation.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        [Input("selectedRepositoryIds")]
        private InputList<int>? _selectedRepositoryIds;

        /// <summary>
        /// An array of repository ids that can access the organization variable.
        /// </summary>
        public InputList<int> SelectedRepositoryIds
        {
            get => _selectedRepositoryIds ?? (_selectedRepositoryIds = new InputList<int>());
            set => _selectedRepositoryIds = value;
        }

        /// <summary>
        /// Date of actions_variable update.
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        /// <summary>
        /// Value of the variable
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        /// <summary>
        /// Name of the variable
        /// </summary>
        [Input("variableName")]
        public Input<string>? VariableName { get; set; }

        /// <summary>
        /// Configures the access that repositories have to the organization variable.
        /// Must be one of `all`, `private`, `selected`. `selected_repository_ids` is required if set to `selected`.
        /// </summary>
        [Input("visibility")]
        public Input<string>? Visibility { get; set; }

        public ActionsOrganizationVariableState()
        {
        }
        public static new ActionsOrganizationVariableState Empty => new ActionsOrganizationVariableState();
    }
}
