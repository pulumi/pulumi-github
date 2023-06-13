// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github
{
    public static class GetRepositoryEnvironments
    {
        /// <summary>
        /// Use this data source to retrieve information about environments for a repository.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Github = Pulumi.Github;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Github.GetRepositoryEnvironments.Invoke(new()
        ///     {
        ///         Repository = "example-repository",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetRepositoryEnvironmentsResult> InvokeAsync(GetRepositoryEnvironmentsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRepositoryEnvironmentsResult>("github:index/getRepositoryEnvironments:getRepositoryEnvironments", args ?? new GetRepositoryEnvironmentsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve information about environments for a repository.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Github = Pulumi.Github;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Github.GetRepositoryEnvironments.Invoke(new()
        ///     {
        ///         Repository = "example-repository",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetRepositoryEnvironmentsResult> Invoke(GetRepositoryEnvironmentsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRepositoryEnvironmentsResult>("github:index/getRepositoryEnvironments:getRepositoryEnvironments", args ?? new GetRepositoryEnvironmentsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRepositoryEnvironmentsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the repository to retrieve the environments from.
        /// </summary>
        [Input("repository", required: true)]
        public string Repository { get; set; } = null!;

        public GetRepositoryEnvironmentsArgs()
        {
        }
        public static new GetRepositoryEnvironmentsArgs Empty => new GetRepositoryEnvironmentsArgs();
    }

    public sealed class GetRepositoryEnvironmentsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the repository to retrieve the environments from.
        /// </summary>
        [Input("repository", required: true)]
        public Input<string> Repository { get; set; } = null!;

        public GetRepositoryEnvironmentsInvokeArgs()
        {
        }
        public static new GetRepositoryEnvironmentsInvokeArgs Empty => new GetRepositoryEnvironmentsInvokeArgs();
    }


    [OutputType]
    public sealed class GetRepositoryEnvironmentsResult
    {
        /// <summary>
        /// The list of this repository's environments. Each element of `environments` has the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetRepositoryEnvironmentsEnvironmentResult> Environments;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Repository;

        [OutputConstructor]
        private GetRepositoryEnvironmentsResult(
            ImmutableArray<Outputs.GetRepositoryEnvironmentsEnvironmentResult> environments,

            string id,

            string repository)
        {
            Environments = environments;
            Id = id;
            Repository = repository;
        }
    }
}