// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github
{
    public static class GetRef
    {
        /// <summary>
        /// Use this data source to retrieve information about a repository ref.
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
        ///     var development = Github.GetRef.Invoke(new()
        ///     {
        ///         Owner = "example",
        ///         Repository = "example",
        ///         Ref = "heads/development",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetRefResult> InvokeAsync(GetRefArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRefResult>("github:index/getRef:getRef", args ?? new GetRefArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve information about a repository ref.
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
        ///     var development = Github.GetRef.Invoke(new()
        ///     {
        ///         Owner = "example",
        ///         Repository = "example",
        ///         Ref = "heads/development",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetRefResult> Invoke(GetRefInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRefResult>("github:index/getRef:getRef", args ?? new GetRefInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve information about a repository ref.
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
        ///     var development = Github.GetRef.Invoke(new()
        ///     {
        ///         Owner = "example",
        ///         Repository = "example",
        ///         Ref = "heads/development",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetRefResult> Invoke(GetRefInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetRefResult>("github:index/getRef:getRef", args ?? new GetRefInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRefArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Owner of the repository.
        /// </summary>
        [Input("owner")]
        public string? Owner { get; set; }

        /// <summary>
        /// The repository ref to look up. Must be formatted `heads/&lt;ref&gt;` for branches, and `tags/&lt;ref&gt;` for tags.
        /// </summary>
        [Input("ref", required: true)]
        public string Ref { get; set; } = null!;

        /// <summary>
        /// The GitHub repository name.
        /// </summary>
        [Input("repository", required: true)]
        public string Repository { get; set; } = null!;

        public GetRefArgs()
        {
        }
        public static new GetRefArgs Empty => new GetRefArgs();
    }

    public sealed class GetRefInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Owner of the repository.
        /// </summary>
        [Input("owner")]
        public Input<string>? Owner { get; set; }

        /// <summary>
        /// The repository ref to look up. Must be formatted `heads/&lt;ref&gt;` for branches, and `tags/&lt;ref&gt;` for tags.
        /// </summary>
        [Input("ref", required: true)]
        public Input<string> Ref { get; set; } = null!;

        /// <summary>
        /// The GitHub repository name.
        /// </summary>
        [Input("repository", required: true)]
        public Input<string> Repository { get; set; } = null!;

        public GetRefInvokeArgs()
        {
        }
        public static new GetRefInvokeArgs Empty => new GetRefInvokeArgs();
    }


    [OutputType]
    public sealed class GetRefResult
    {
        /// <summary>
        /// An etag representing the ref.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? Owner;
        public readonly string Ref;
        public readonly string Repository;
        /// <summary>
        /// A string storing the reference's `HEAD` commit's SHA1.
        /// </summary>
        public readonly string Sha;

        [OutputConstructor]
        private GetRefResult(
            string etag,

            string id,

            string? owner,

            string @ref,

            string repository,

            string sha)
        {
            Etag = etag;
            Id = id;
            Owner = owner;
            Ref = @ref;
            Repository = repository;
            Sha = sha;
        }
    }
}
