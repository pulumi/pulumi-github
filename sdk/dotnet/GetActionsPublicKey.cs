// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github
{
    public static class GetActionsPublicKey
    {
        /// <summary>
        /// Use this data source to retrieve information about a GitHub Actions public key. This data source is required to be used with other GitHub secrets interactions.
        /// Note that the provider `token` must have admin rights to a repository to retrieve it's action public key.
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
        ///     var example = Github.GetActionsPublicKey.Invoke(new()
        ///     {
        ///         Repository = "example_repo",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetActionsPublicKeyResult> InvokeAsync(GetActionsPublicKeyArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetActionsPublicKeyResult>("github:index/getActionsPublicKey:getActionsPublicKey", args ?? new GetActionsPublicKeyArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve information about a GitHub Actions public key. This data source is required to be used with other GitHub secrets interactions.
        /// Note that the provider `token` must have admin rights to a repository to retrieve it's action public key.
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
        ///     var example = Github.GetActionsPublicKey.Invoke(new()
        ///     {
        ///         Repository = "example_repo",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetActionsPublicKeyResult> Invoke(GetActionsPublicKeyInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetActionsPublicKeyResult>("github:index/getActionsPublicKey:getActionsPublicKey", args ?? new GetActionsPublicKeyInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve information about a GitHub Actions public key. This data source is required to be used with other GitHub secrets interactions.
        /// Note that the provider `token` must have admin rights to a repository to retrieve it's action public key.
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
        ///     var example = Github.GetActionsPublicKey.Invoke(new()
        ///     {
        ///         Repository = "example_repo",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetActionsPublicKeyResult> Invoke(GetActionsPublicKeyInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetActionsPublicKeyResult>("github:index/getActionsPublicKey:getActionsPublicKey", args ?? new GetActionsPublicKeyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetActionsPublicKeyArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the repository to get public key from.
        /// </summary>
        [Input("repository", required: true)]
        public string Repository { get; set; } = null!;

        public GetActionsPublicKeyArgs()
        {
        }
        public static new GetActionsPublicKeyArgs Empty => new GetActionsPublicKeyArgs();
    }

    public sealed class GetActionsPublicKeyInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the repository to get public key from.
        /// </summary>
        [Input("repository", required: true)]
        public Input<string> Repository { get; set; } = null!;

        public GetActionsPublicKeyInvokeArgs()
        {
        }
        public static new GetActionsPublicKeyInvokeArgs Empty => new GetActionsPublicKeyInvokeArgs();
    }


    [OutputType]
    public sealed class GetActionsPublicKeyResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Actual key retrieved.
        /// </summary>
        public readonly string Key;
        /// <summary>
        /// ID of the key that has been retrieved.
        /// </summary>
        public readonly string KeyId;
        public readonly string Repository;

        [OutputConstructor]
        private GetActionsPublicKeyResult(
            string id,

            string key,

            string keyId,

            string repository)
        {
            Id = id;
            Key = key;
            KeyId = keyId;
            Repository = repository;
        }
    }
}
