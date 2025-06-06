// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github
{
    public static class GetRepositoryAutolinkReferences
    {
        /// <summary>
        /// Use this data source to retrieve autolink references for a repository.
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
        ///     var example = Github.GetRepositoryAutolinkReferences.Invoke(new()
        ///     {
        ///         Repository = "example-repository",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetRepositoryAutolinkReferencesResult> InvokeAsync(GetRepositoryAutolinkReferencesArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRepositoryAutolinkReferencesResult>("github:index/getRepositoryAutolinkReferences:getRepositoryAutolinkReferences", args ?? new GetRepositoryAutolinkReferencesArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve autolink references for a repository.
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
        ///     var example = Github.GetRepositoryAutolinkReferences.Invoke(new()
        ///     {
        ///         Repository = "example-repository",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetRepositoryAutolinkReferencesResult> Invoke(GetRepositoryAutolinkReferencesInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRepositoryAutolinkReferencesResult>("github:index/getRepositoryAutolinkReferences:getRepositoryAutolinkReferences", args ?? new GetRepositoryAutolinkReferencesInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve autolink references for a repository.
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
        ///     var example = Github.GetRepositoryAutolinkReferences.Invoke(new()
        ///     {
        ///         Repository = "example-repository",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetRepositoryAutolinkReferencesResult> Invoke(GetRepositoryAutolinkReferencesInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetRepositoryAutolinkReferencesResult>("github:index/getRepositoryAutolinkReferences:getRepositoryAutolinkReferences", args ?? new GetRepositoryAutolinkReferencesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRepositoryAutolinkReferencesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the repository to retrieve the autolink references from.
        /// </summary>
        [Input("repository", required: true)]
        public string Repository { get; set; } = null!;

        public GetRepositoryAutolinkReferencesArgs()
        {
        }
        public static new GetRepositoryAutolinkReferencesArgs Empty => new GetRepositoryAutolinkReferencesArgs();
    }

    public sealed class GetRepositoryAutolinkReferencesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the repository to retrieve the autolink references from.
        /// </summary>
        [Input("repository", required: true)]
        public Input<string> Repository { get; set; } = null!;

        public GetRepositoryAutolinkReferencesInvokeArgs()
        {
        }
        public static new GetRepositoryAutolinkReferencesInvokeArgs Empty => new GetRepositoryAutolinkReferencesInvokeArgs();
    }


    [OutputType]
    public sealed class GetRepositoryAutolinkReferencesResult
    {
        /// <summary>
        /// The list of this repository's autolink references. Each element of `autolink_references` has the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetRepositoryAutolinkReferencesAutolinkReferenceResult> AutolinkReferences;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Repository;

        [OutputConstructor]
        private GetRepositoryAutolinkReferencesResult(
            ImmutableArray<Outputs.GetRepositoryAutolinkReferencesAutolinkReferenceResult> autolinkReferences,

            string id,

            string repository)
        {
            AutolinkReferences = autolinkReferences;
            Id = id;
            Repository = repository;
        }
    }
}
