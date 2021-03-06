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
        /// using Pulumi;
        /// using Github = Pulumi.Github;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(Github.GetRepositories.InvokeAsync(new Github.GetRepositoriesArgs
        ///         {
        ///             Query = "org:hashicorp language:Go",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetRepositoriesResult> InvokeAsync(GetRepositoriesArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRepositoriesResult>("github:index/getRepositories:getRepositories", args ?? new GetRepositoriesArgs(), options.WithVersion());
    }


    public sealed class GetRepositoriesArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Search query. See [documentation for the search syntax](https://help.github.com/articles/understanding-the-search-syntax/).
        /// </summary>
        [Input("query", required: true)]
        public string Query { get; set; } = null!;

        /// <summary>
        /// Sorts the repositories returned by the specified attribute. Valid values include `stars`, `fork`, and `updated`. Defaults to `updated`.
        /// </summary>
        [Input("sort")]
        public string? Sort { get; set; }

        public GetRepositoriesArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetRepositoriesResult
    {
        public readonly ImmutableArray<string> FullNames;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Names;
        public readonly string Query;
        public readonly string? Sort;

        [OutputConstructor]
        private GetRepositoriesResult(
            ImmutableArray<string> fullNames,

            string id,

            ImmutableArray<string> names,

            string query,

            string? sort)
        {
            FullNames = fullNames;
            Id = id;
            Names = names;
            Query = query;
            Sort = sort;
        }
    }
}
