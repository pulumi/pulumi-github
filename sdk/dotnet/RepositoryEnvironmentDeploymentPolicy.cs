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
    /// This resource allows you to create and manage environment deployment branch policies for a GitHub repository.
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
    ///     var current = Github.GetUser.Invoke(new()
    ///     {
    ///         Username = "",
    ///     });
    /// 
    ///     var testRepository = new Github.Repository("testRepository");
    /// 
    ///     var testRepositoryEnvironment = new Github.RepositoryEnvironment("testRepositoryEnvironment", new()
    ///     {
    ///         Repository = testRepository.Name,
    ///         Environment = "environment/test",
    ///         WaitTimer = 10000,
    ///         Reviewers = new[]
    ///         {
    ///             new Github.Inputs.RepositoryEnvironmentReviewerArgs
    ///             {
    ///                 Users = new[]
    ///                 {
    ///                     current.Apply(getUserResult =&gt; getUserResult.Id),
    ///                 },
    ///             },
    ///         },
    ///         DeploymentBranchPolicy = new Github.Inputs.RepositoryEnvironmentDeploymentBranchPolicyArgs
    ///         {
    ///             ProtectedBranches = false,
    ///             CustomBranchPolicies = true,
    ///         },
    ///     });
    /// 
    ///     var testRepositoryEnvironmentDeploymentPolicy = new Github.RepositoryEnvironmentDeploymentPolicy("testRepositoryEnvironmentDeploymentPolicy", new()
    ///     {
    ///         Repository = testRepository.Name,
    ///         Environment = testRepositoryEnvironment.Environment,
    ///         BranchPattern = "releases/*",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// GitHub Repository Environment Deployment Policy can be imported using an ID made up of `name` of the repository combined with the `environment` name of the environment with the `Id` of the deployment policy, separated by a `:` character, e.g.
    /// 
    /// ```sh
    /// $ pulumi import github:index/repositoryEnvironmentDeploymentPolicy:RepositoryEnvironmentDeploymentPolicy daily terraform:daily:123456
    /// ```
    /// </summary>
    [GithubResourceType("github:index/repositoryEnvironmentDeploymentPolicy:RepositoryEnvironmentDeploymentPolicy")]
    public partial class RepositoryEnvironmentDeploymentPolicy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The name pattern that branches must match in order to deploy to the environment.
        /// </summary>
        [Output("branchPattern")]
        public Output<string> BranchPattern { get; private set; } = null!;

        /// <summary>
        /// The name of the environment.
        /// </summary>
        [Output("environment")]
        public Output<string> Environment { get; private set; } = null!;

        /// <summary>
        /// The repository of the environment.
        /// </summary>
        [Output("repository")]
        public Output<string> Repository { get; private set; } = null!;


        /// <summary>
        /// Create a RepositoryEnvironmentDeploymentPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RepositoryEnvironmentDeploymentPolicy(string name, RepositoryEnvironmentDeploymentPolicyArgs args, CustomResourceOptions? options = null)
            : base("github:index/repositoryEnvironmentDeploymentPolicy:RepositoryEnvironmentDeploymentPolicy", name, args ?? new RepositoryEnvironmentDeploymentPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RepositoryEnvironmentDeploymentPolicy(string name, Input<string> id, RepositoryEnvironmentDeploymentPolicyState? state = null, CustomResourceOptions? options = null)
            : base("github:index/repositoryEnvironmentDeploymentPolicy:RepositoryEnvironmentDeploymentPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RepositoryEnvironmentDeploymentPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RepositoryEnvironmentDeploymentPolicy Get(string name, Input<string> id, RepositoryEnvironmentDeploymentPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new RepositoryEnvironmentDeploymentPolicy(name, id, state, options);
        }
    }

    public sealed class RepositoryEnvironmentDeploymentPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name pattern that branches must match in order to deploy to the environment.
        /// </summary>
        [Input("branchPattern", required: true)]
        public Input<string> BranchPattern { get; set; } = null!;

        /// <summary>
        /// The name of the environment.
        /// </summary>
        [Input("environment", required: true)]
        public Input<string> Environment { get; set; } = null!;

        /// <summary>
        /// The repository of the environment.
        /// </summary>
        [Input("repository", required: true)]
        public Input<string> Repository { get; set; } = null!;

        public RepositoryEnvironmentDeploymentPolicyArgs()
        {
        }
        public static new RepositoryEnvironmentDeploymentPolicyArgs Empty => new RepositoryEnvironmentDeploymentPolicyArgs();
    }

    public sealed class RepositoryEnvironmentDeploymentPolicyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name pattern that branches must match in order to deploy to the environment.
        /// </summary>
        [Input("branchPattern")]
        public Input<string>? BranchPattern { get; set; }

        /// <summary>
        /// The name of the environment.
        /// </summary>
        [Input("environment")]
        public Input<string>? Environment { get; set; }

        /// <summary>
        /// The repository of the environment.
        /// </summary>
        [Input("repository")]
        public Input<string>? Repository { get; set; }

        public RepositoryEnvironmentDeploymentPolicyState()
        {
        }
        public static new RepositoryEnvironmentDeploymentPolicyState Empty => new RepositoryEnvironmentDeploymentPolicyState();
    }
}
