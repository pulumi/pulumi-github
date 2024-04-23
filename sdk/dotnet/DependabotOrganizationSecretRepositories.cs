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
    /// This resource allows you to manage the repository allow list for existing GitHub Dependabot secrets within your GitHub organization.
    /// You must have write access to an organization secret to use this resource.
    /// 
    /// This resource is only applicable when `visibility` of the existing organization secret has been set to `selected`.
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
    ///     var repo = Github.GetRepository.Invoke(new()
    ///     {
    ///         FullName = "my-org/repo",
    ///     });
    /// 
    ///     var exampleSecret = new Github.DependabotOrganizationSecret("example_secret", new()
    ///     {
    ///         SecretName = "example_secret_name",
    ///         Visibility = "private",
    ///         PlaintextValue = someSecretString,
    ///     });
    /// 
    ///     var orgSecretRepos = new Github.DependabotOrganizationSecretRepositories("org_secret_repos", new()
    ///     {
    ///         SecretName = exampleSecret.SecretName,
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
    /// This resource can be imported using an ID made up of the secret name:
    /// 
    /// ```sh
    /// $ pulumi import github:index/dependabotOrganizationSecretRepositories:DependabotOrganizationSecretRepositories test_secret_repos test_secret_name
    /// ```
    /// </summary>
    [GithubResourceType("github:index/dependabotOrganizationSecretRepositories:DependabotOrganizationSecretRepositories")]
    public partial class DependabotOrganizationSecretRepositories : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Name of the existing secret
        /// </summary>
        [Output("secretName")]
        public Output<string> SecretName { get; private set; } = null!;

        /// <summary>
        /// An array of repository ids that can access the organization secret.
        /// </summary>
        [Output("selectedRepositoryIds")]
        public Output<ImmutableArray<int>> SelectedRepositoryIds { get; private set; } = null!;


        /// <summary>
        /// Create a DependabotOrganizationSecretRepositories resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DependabotOrganizationSecretRepositories(string name, DependabotOrganizationSecretRepositoriesArgs args, CustomResourceOptions? options = null)
            : base("github:index/dependabotOrganizationSecretRepositories:DependabotOrganizationSecretRepositories", name, args ?? new DependabotOrganizationSecretRepositoriesArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DependabotOrganizationSecretRepositories(string name, Input<string> id, DependabotOrganizationSecretRepositoriesState? state = null, CustomResourceOptions? options = null)
            : base("github:index/dependabotOrganizationSecretRepositories:DependabotOrganizationSecretRepositories", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DependabotOrganizationSecretRepositories resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DependabotOrganizationSecretRepositories Get(string name, Input<string> id, DependabotOrganizationSecretRepositoriesState? state = null, CustomResourceOptions? options = null)
        {
            return new DependabotOrganizationSecretRepositories(name, id, state, options);
        }
    }

    public sealed class DependabotOrganizationSecretRepositoriesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the existing secret
        /// </summary>
        [Input("secretName", required: true)]
        public Input<string> SecretName { get; set; } = null!;

        [Input("selectedRepositoryIds", required: true)]
        private InputList<int>? _selectedRepositoryIds;

        /// <summary>
        /// An array of repository ids that can access the organization secret.
        /// </summary>
        public InputList<int> SelectedRepositoryIds
        {
            get => _selectedRepositoryIds ?? (_selectedRepositoryIds = new InputList<int>());
            set => _selectedRepositoryIds = value;
        }

        public DependabotOrganizationSecretRepositoriesArgs()
        {
        }
        public static new DependabotOrganizationSecretRepositoriesArgs Empty => new DependabotOrganizationSecretRepositoriesArgs();
    }

    public sealed class DependabotOrganizationSecretRepositoriesState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the existing secret
        /// </summary>
        [Input("secretName")]
        public Input<string>? SecretName { get; set; }

        [Input("selectedRepositoryIds")]
        private InputList<int>? _selectedRepositoryIds;

        /// <summary>
        /// An array of repository ids that can access the organization secret.
        /// </summary>
        public InputList<int> SelectedRepositoryIds
        {
            get => _selectedRepositoryIds ?? (_selectedRepositoryIds = new InputList<int>());
            set => _selectedRepositoryIds = value;
        }

        public DependabotOrganizationSecretRepositoriesState()
        {
        }
        public static new DependabotOrganizationSecretRepositoriesState Empty => new DependabotOrganizationSecretRepositoriesState();
    }
}
