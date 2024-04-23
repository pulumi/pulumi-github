// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github
{
    public static class GetRepositoryDeploymentBranchPolicies
    {
        /// <summary>
        /// Use this data source to retrieve deployment branch policies for a repository / environment.
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
        ///     var example = Github.GetRepositoryDeploymentBranchPolicies.Invoke(new()
        ///     {
        ///         Repository = "example-repository",
        ///         EnvironmentName = "env_name",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetRepositoryDeploymentBranchPoliciesResult> InvokeAsync(GetRepositoryDeploymentBranchPoliciesArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRepositoryDeploymentBranchPoliciesResult>("github:index/getRepositoryDeploymentBranchPolicies:getRepositoryDeploymentBranchPolicies", args ?? new GetRepositoryDeploymentBranchPoliciesArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve deployment branch policies for a repository / environment.
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
        ///     var example = Github.GetRepositoryDeploymentBranchPolicies.Invoke(new()
        ///     {
        ///         Repository = "example-repository",
        ///         EnvironmentName = "env_name",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetRepositoryDeploymentBranchPoliciesResult> Invoke(GetRepositoryDeploymentBranchPoliciesInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRepositoryDeploymentBranchPoliciesResult>("github:index/getRepositoryDeploymentBranchPolicies:getRepositoryDeploymentBranchPolicies", args ?? new GetRepositoryDeploymentBranchPoliciesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRepositoryDeploymentBranchPoliciesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the environment to retrieve the deployment branch policies  from.
        /// </summary>
        [Input("environmentName", required: true)]
        public string EnvironmentName { get; set; } = null!;

        /// <summary>
        /// Name of the repository to retrieve the deployment branch policies from.
        /// </summary>
        [Input("repository", required: true)]
        public string Repository { get; set; } = null!;

        public GetRepositoryDeploymentBranchPoliciesArgs()
        {
        }
        public static new GetRepositoryDeploymentBranchPoliciesArgs Empty => new GetRepositoryDeploymentBranchPoliciesArgs();
    }

    public sealed class GetRepositoryDeploymentBranchPoliciesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the environment to retrieve the deployment branch policies  from.
        /// </summary>
        [Input("environmentName", required: true)]
        public Input<string> EnvironmentName { get; set; } = null!;

        /// <summary>
        /// Name of the repository to retrieve the deployment branch policies from.
        /// </summary>
        [Input("repository", required: true)]
        public Input<string> Repository { get; set; } = null!;

        public GetRepositoryDeploymentBranchPoliciesInvokeArgs()
        {
        }
        public static new GetRepositoryDeploymentBranchPoliciesInvokeArgs Empty => new GetRepositoryDeploymentBranchPoliciesInvokeArgs();
    }


    [OutputType]
    public sealed class GetRepositoryDeploymentBranchPoliciesResult
    {
        /// <summary>
        /// The list of this repository / environment deployment policies. Each element of `deployment_branch_policies` has the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetRepositoryDeploymentBranchPoliciesDeploymentBranchPolicyResult> DeploymentBranchPolicies;
        public readonly string EnvironmentName;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Repository;

        [OutputConstructor]
        private GetRepositoryDeploymentBranchPoliciesResult(
            ImmutableArray<Outputs.GetRepositoryDeploymentBranchPoliciesDeploymentBranchPolicyResult> deploymentBranchPolicies,

            string environmentName,

            string id,

            string repository)
        {
            DeploymentBranchPolicies = deploymentBranchPolicies;
            EnvironmentName = environmentName;
            Id = id;
            Repository = repository;
        }
    }
}
