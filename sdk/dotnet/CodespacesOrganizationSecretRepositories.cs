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
    /// This resource allows you to manage repository allow list for existing GitHub Codespaces secrets within your GitHub organization.
    /// 
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
    ///     var orgSecretRepos = new Github.CodespacesOrganizationSecretRepositories("orgSecretRepos", new()
    ///     {
    ///         SecretName = "existing_secret_name",
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
    ///  $ pulumi import github:index/codespacesOrganizationSecretRepositories:CodespacesOrganizationSecretRepositories org_secret_repos existing_secret_name
    /// ```
    /// </summary>
    [GithubResourceType("github:index/codespacesOrganizationSecretRepositories:CodespacesOrganizationSecretRepositories")]
    public partial class CodespacesOrganizationSecretRepositories : global::Pulumi.CustomResource
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
        /// Create a CodespacesOrganizationSecretRepositories resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CodespacesOrganizationSecretRepositories(string name, CodespacesOrganizationSecretRepositoriesArgs args, CustomResourceOptions? options = null)
            : base("github:index/codespacesOrganizationSecretRepositories:CodespacesOrganizationSecretRepositories", name, args ?? new CodespacesOrganizationSecretRepositoriesArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CodespacesOrganizationSecretRepositories(string name, Input<string> id, CodespacesOrganizationSecretRepositoriesState? state = null, CustomResourceOptions? options = null)
            : base("github:index/codespacesOrganizationSecretRepositories:CodespacesOrganizationSecretRepositories", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing CodespacesOrganizationSecretRepositories resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CodespacesOrganizationSecretRepositories Get(string name, Input<string> id, CodespacesOrganizationSecretRepositoriesState? state = null, CustomResourceOptions? options = null)
        {
            return new CodespacesOrganizationSecretRepositories(name, id, state, options);
        }
    }

    public sealed class CodespacesOrganizationSecretRepositoriesArgs : global::Pulumi.ResourceArgs
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

        public CodespacesOrganizationSecretRepositoriesArgs()
        {
        }
        public static new CodespacesOrganizationSecretRepositoriesArgs Empty => new CodespacesOrganizationSecretRepositoriesArgs();
    }

    public sealed class CodespacesOrganizationSecretRepositoriesState : global::Pulumi.ResourceArgs
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

        public CodespacesOrganizationSecretRepositoriesState()
        {
        }
        public static new CodespacesOrganizationSecretRepositoriesState Empty => new CodespacesOrganizationSecretRepositoriesState();
    }
}
