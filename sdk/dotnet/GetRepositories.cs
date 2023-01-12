// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github
{
    public static class GetRepositories
    {
        /// <summary>
        /// &gt; **Note:** The data source will return a maximum of `1000` repositories
        /// 	[as documented in official API docs](https://developer.github.com/v3/search/#about-the-search-api).
        /// 
        /// Use this data source to retrieve a list of GitHub repositories using a search query.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Github = Pulumi.Github;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Github.GetRepositories.Invoke(new()
        ///     {
        ///         IncludeRepoId = true,
        ///         Query = "org:hashicorp language:Go",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetRepositoriesResult> InvokeAsync(GetRepositoriesArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRepositoriesResult>("github:index/getRepositories:getRepositories", args ?? new GetRepositoriesArgs(), options.WithDefaults());

        /// <summary>
        /// &gt; **Note:** The data source will return a maximum of `1000` repositories
        /// 	[as documented in official API docs](https://developer.github.com/v3/search/#about-the-search-api).
        /// 
        /// Use this data source to retrieve a list of GitHub repositories using a search query.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Github = Pulumi.Github;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Github.GetRepositories.Invoke(new()
        ///     {
        ///         IncludeRepoId = true,
        ///         Query = "org:hashicorp language:Go",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetRepositoriesResult> Invoke(GetRepositoriesInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRepositoriesResult>("github:index/getRepositories:getRepositories", args ?? new GetRepositoriesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRepositoriesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Returns a list of found repository IDs
        /// </summary>
        [Input("includeRepoId")]
        public bool? IncludeRepoId { get; set; }

        /// <summary>
        /// Search query. See [documentation for the search syntax](https://help.github.com/articles/understanding-the-search-syntax/).
        /// </summary>
        [Input("query", required: true)]
        public string Query { get; set; } = null!;

        /// <summary>
        /// Set the number of repositories requested per API call. Can be useful to decrease if requests are timing out or to increase to reduce the number of API calls. Defaults to 100.
        /// </summary>
        [Input("resultsPerPage")]
        public int? ResultsPerPage { get; set; }

        /// <summary>
        /// Sorts the repositories returned by the specified attribute. Valid values include `stars`, `fork`, and `updated`. Defaults to `updated`.
        /// </summary>
        [Input("sort")]
        public string? Sort { get; set; }

        public GetRepositoriesArgs()
        {
        }
        public static new GetRepositoriesArgs Empty => new GetRepositoriesArgs();
    }

    public sealed class GetRepositoriesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Returns a list of found repository IDs
        /// </summary>
        [Input("includeRepoId")]
        public Input<bool>? IncludeRepoId { get; set; }

        /// <summary>
        /// Search query. See [documentation for the search syntax](https://help.github.com/articles/understanding-the-search-syntax/).
        /// </summary>
        [Input("query", required: true)]
        public Input<string> Query { get; set; } = null!;

        /// <summary>
        /// Set the number of repositories requested per API call. Can be useful to decrease if requests are timing out or to increase to reduce the number of API calls. Defaults to 100.
        /// </summary>
        [Input("resultsPerPage")]
        public Input<int>? ResultsPerPage { get; set; }

        /// <summary>
        /// Sorts the repositories returned by the specified attribute. Valid values include `stars`, `fork`, and `updated`. Defaults to `updated`.
        /// </summary>
        [Input("sort")]
        public Input<string>? Sort { get; set; }

        public GetRepositoriesInvokeArgs()
        {
        }
        public static new GetRepositoriesInvokeArgs Empty => new GetRepositoriesInvokeArgs();
    }


    [OutputType]
    public sealed class GetRepositoriesResult
    {
        public readonly ImmutableArray<string> FullNames;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly bool? IncludeRepoId;
        public readonly ImmutableArray<string> Names;
        public readonly string Query;
        /// <summary>
        /// (Optional) A list of found repository IDs (e.g. `449898861`)
        /// </summary>
        public readonly ImmutableArray<int> RepoIds;
        public readonly int? ResultsPerPage;
        public readonly string? Sort;

        [OutputConstructor]
        private GetRepositoriesResult(
            ImmutableArray<string> fullNames,

            string id,

            bool? includeRepoId,

            ImmutableArray<string> names,

            string query,

            ImmutableArray<int> repoIds,

            int? resultsPerPage,

            string? sort)
        {
            FullNames = fullNames;
            Id = id;
            IncludeRepoId = includeRepoId;
            Names = names;
            Query = query;
            RepoIds = repoIds;
            ResultsPerPage = resultsPerPage;
            Sort = sort;
        }
    }
}
