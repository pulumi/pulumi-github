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
    /// This resource allows you to create and manage deployment branch policies.
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
    ///     var env = new Github.RepositoryEnvironment("env", new()
    ///     {
    ///         Repository = "my_repo",
    ///         Environment = "my_env",
    ///         DeploymentBranchPolicy = new Github.Inputs.RepositoryEnvironmentDeploymentBranchPolicyArgs
    ///         {
    ///             ProtectedBranches = false,
    ///             CustomBranchPolicies = true,
    ///         },
    ///     });
    /// 
    ///     var foo = new Github.RepositoryDeploymentBranchPolicy("foo", new()
    ///     {
    ///         Repository = "my_repo",
    ///         EnvironmentName = "my_env",
    ///     }, new CustomResourceOptions
    ///     {
    ///         DependsOn = new[]
    ///         {
    ///             env,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ```sh
    ///  $ pulumi import github:index/repositoryDeploymentBranchPolicy:RepositoryDeploymentBranchPolicy foo repo:env:id
    /// ```
    /// </summary>
    [GithubResourceType("github:index/repositoryDeploymentBranchPolicy:RepositoryDeploymentBranchPolicy")]
    public partial class RepositoryDeploymentBranchPolicy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the environment. This environment must have `deployment_branch_policy.custom_branch_policies` set to true or a 404 error will be thrown.
        /// </summary>
        [Output("environmentName")]
        public Output<string> EnvironmentName { get; private set; } = null!;

        /// <summary>
        /// An etag representing the Branch object.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// The name pattern that branches must match in order to deploy to the environment.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The repository to create the policy in.
        /// </summary>
        [Output("repository")]
        public Output<string> Repository { get; private set; } = null!;


        /// <summary>
        /// Create a RepositoryDeploymentBranchPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RepositoryDeploymentBranchPolicy(string name, RepositoryDeploymentBranchPolicyArgs args, CustomResourceOptions? options = null)
            : base("github:index/repositoryDeploymentBranchPolicy:RepositoryDeploymentBranchPolicy", name, args ?? new RepositoryDeploymentBranchPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RepositoryDeploymentBranchPolicy(string name, Input<string> id, RepositoryDeploymentBranchPolicyState? state = null, CustomResourceOptions? options = null)
            : base("github:index/repositoryDeploymentBranchPolicy:RepositoryDeploymentBranchPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RepositoryDeploymentBranchPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RepositoryDeploymentBranchPolicy Get(string name, Input<string> id, RepositoryDeploymentBranchPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new RepositoryDeploymentBranchPolicy(name, id, state, options);
        }
    }

    public sealed class RepositoryDeploymentBranchPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the environment. This environment must have `deployment_branch_policy.custom_branch_policies` set to true or a 404 error will be thrown.
        /// </summary>
        [Input("environmentName", required: true)]
        public Input<string> EnvironmentName { get; set; } = null!;

        /// <summary>
        /// The name pattern that branches must match in order to deploy to the environment.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The repository to create the policy in.
        /// </summary>
        [Input("repository", required: true)]
        public Input<string> Repository { get; set; } = null!;

        public RepositoryDeploymentBranchPolicyArgs()
        {
        }
        public static new RepositoryDeploymentBranchPolicyArgs Empty => new RepositoryDeploymentBranchPolicyArgs();
    }

    public sealed class RepositoryDeploymentBranchPolicyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the environment. This environment must have `deployment_branch_policy.custom_branch_policies` set to true or a 404 error will be thrown.
        /// </summary>
        [Input("environmentName")]
        public Input<string>? EnvironmentName { get; set; }

        /// <summary>
        /// An etag representing the Branch object.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// The name pattern that branches must match in order to deploy to the environment.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The repository to create the policy in.
        /// </summary>
        [Input("repository")]
        public Input<string>? Repository { get; set; }

        public RepositoryDeploymentBranchPolicyState()
        {
        }
        public static new RepositoryDeploymentBranchPolicyState Empty => new RepositoryDeploymentBranchPolicyState();
    }
}
