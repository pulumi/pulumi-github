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
    /// &gt; **Note**: This resource is not compatible with the GitHub App Installation authentication method.
    /// 
    /// This resource manages relationships between app installations and repositories
    /// in your GitHub organization.
    /// 
    /// Creating this resource installs a particular app on a particular repository.
    /// 
    /// The app installation and the repository must both belong to the same
    /// organization on GitHub. Note: you can review your organization's installations
    /// by the following the instructions at this
    /// [link](https://docs.github.com/en/github/setting-up-and-managing-organizations-and-teams/reviewing-your-organizations-installed-integrations).
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
    ///     // Create a repository.
    ///     var someRepo = new Github.Repository("some_repo", new()
    ///     {
    ///         Name = "some-repo",
    ///     });
    /// 
    ///     var someAppRepo = new Github.AppInstallationRepository("some_app_repo", new()
    ///     {
    ///         InstallationId = "1234567",
    ///         Repository = someRepo.Name,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// GitHub App Installation Repository can be imported
    /// using an ID made up of `installation_id:repository`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import github:index/appInstallationRepository:AppInstallationRepository terraform_repo 1234567:terraform
    /// ```
    /// </summary>
    [GithubResourceType("github:index/appInstallationRepository:AppInstallationRepository")]
    public partial class AppInstallationRepository : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The GitHub app installation id.
        /// </summary>
        [Output("installationId")]
        public Output<string> InstallationId { get; private set; } = null!;

        [Output("repoId")]
        public Output<int> RepoId { get; private set; } = null!;

        /// <summary>
        /// The repository to install the app on.
        /// </summary>
        [Output("repository")]
        public Output<string> Repository { get; private set; } = null!;


        /// <summary>
        /// Create a AppInstallationRepository resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AppInstallationRepository(string name, AppInstallationRepositoryArgs args, CustomResourceOptions? options = null)
            : base("github:index/appInstallationRepository:AppInstallationRepository", name, args ?? new AppInstallationRepositoryArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AppInstallationRepository(string name, Input<string> id, AppInstallationRepositoryState? state = null, CustomResourceOptions? options = null)
            : base("github:index/appInstallationRepository:AppInstallationRepository", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AppInstallationRepository resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AppInstallationRepository Get(string name, Input<string> id, AppInstallationRepositoryState? state = null, CustomResourceOptions? options = null)
        {
            return new AppInstallationRepository(name, id, state, options);
        }
    }

    public sealed class AppInstallationRepositoryArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The GitHub app installation id.
        /// </summary>
        [Input("installationId", required: true)]
        public Input<string> InstallationId { get; set; } = null!;

        /// <summary>
        /// The repository to install the app on.
        /// </summary>
        [Input("repository", required: true)]
        public Input<string> Repository { get; set; } = null!;

        public AppInstallationRepositoryArgs()
        {
        }
        public static new AppInstallationRepositoryArgs Empty => new AppInstallationRepositoryArgs();
    }

    public sealed class AppInstallationRepositoryState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The GitHub app installation id.
        /// </summary>
        [Input("installationId")]
        public Input<string>? InstallationId { get; set; }

        [Input("repoId")]
        public Input<int>? RepoId { get; set; }

        /// <summary>
        /// The repository to install the app on.
        /// </summary>
        [Input("repository")]
        public Input<string>? Repository { get; set; }

        public AppInstallationRepositoryState()
        {
        }
        public static new AppInstallationRepositoryState Empty => new AppInstallationRepositoryState();
    }
}
