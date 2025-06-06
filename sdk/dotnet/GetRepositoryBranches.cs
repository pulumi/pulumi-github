// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github
{
    public static class GetRepositoryBranches
    {
        /// <summary>
        /// Use this data source to retrieve information about branches in a repository.
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
        ///     var example = Github.GetRepositoryBranches.Invoke(new()
        ///     {
        ///         Repository = "example-repository",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetRepositoryBranchesResult> InvokeAsync(GetRepositoryBranchesArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRepositoryBranchesResult>("github:index/getRepositoryBranches:getRepositoryBranches", args ?? new GetRepositoryBranchesArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve information about branches in a repository.
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
        ///     var example = Github.GetRepositoryBranches.Invoke(new()
        ///     {
        ///         Repository = "example-repository",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetRepositoryBranchesResult> Invoke(GetRepositoryBranchesInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRepositoryBranchesResult>("github:index/getRepositoryBranches:getRepositoryBranches", args ?? new GetRepositoryBranchesInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve information about branches in a repository.
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
        ///     var example = Github.GetRepositoryBranches.Invoke(new()
        ///     {
        ///         Repository = "example-repository",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetRepositoryBranchesResult> Invoke(GetRepositoryBranchesInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetRepositoryBranchesResult>("github:index/getRepositoryBranches:getRepositoryBranches", args ?? new GetRepositoryBranchesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRepositoryBranchesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// . If true, the `branches` attributes will be populated only with non protected branches. Default: `false`.
        /// </summary>
        [Input("onlyNonProtectedBranches")]
        public bool? OnlyNonProtectedBranches { get; set; }

        /// <summary>
        /// . If true, the `branches` attributes will be populated only with protected branches. Default: `false`.
        /// </summary>
        [Input("onlyProtectedBranches")]
        public bool? OnlyProtectedBranches { get; set; }

        /// <summary>
        /// Name of the repository to retrieve the branches from.
        /// </summary>
        [Input("repository", required: true)]
        public string Repository { get; set; } = null!;

        public GetRepositoryBranchesArgs()
        {
        }
        public static new GetRepositoryBranchesArgs Empty => new GetRepositoryBranchesArgs();
    }

    public sealed class GetRepositoryBranchesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// . If true, the `branches` attributes will be populated only with non protected branches. Default: `false`.
        /// </summary>
        [Input("onlyNonProtectedBranches")]
        public Input<bool>? OnlyNonProtectedBranches { get; set; }

        /// <summary>
        /// . If true, the `branches` attributes will be populated only with protected branches. Default: `false`.
        /// </summary>
        [Input("onlyProtectedBranches")]
        public Input<bool>? OnlyProtectedBranches { get; set; }

        /// <summary>
        /// Name of the repository to retrieve the branches from.
        /// </summary>
        [Input("repository", required: true)]
        public Input<string> Repository { get; set; } = null!;

        public GetRepositoryBranchesInvokeArgs()
        {
        }
        public static new GetRepositoryBranchesInvokeArgs Empty => new GetRepositoryBranchesInvokeArgs();
    }


    [OutputType]
    public sealed class GetRepositoryBranchesResult
    {
        /// <summary>
        /// The list of this repository's branches. Each element of `branches` has the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetRepositoryBranchesBranchResult> Branches;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly bool? OnlyNonProtectedBranches;
        public readonly bool? OnlyProtectedBranches;
        public readonly string Repository;

        [OutputConstructor]
        private GetRepositoryBranchesResult(
            ImmutableArray<Outputs.GetRepositoryBranchesBranchResult> branches,

            string id,

            bool? onlyNonProtectedBranches,

            bool? onlyProtectedBranches,

            string repository)
        {
            Branches = branches;
            Id = id;
            OnlyNonProtectedBranches = onlyNonProtectedBranches;
            OnlyProtectedBranches = onlyProtectedBranches;
            Repository = repository;
        }
    }
}
