// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github
{
    public static class GetCodespacesUserPublicKey
    {
        /// <summary>
        /// Use this data source to retrieve information about a GitHub Codespaces User public key. This data source is required to be used with other GitHub secrets interactions.
        /// Note that the provider `token` must have admin rights to an user to retrieve it's Codespaces public key.
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Github = Pulumi.Github;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Github.GetCodespacesUserPublicKey.Invoke();
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetCodespacesUserPublicKeyResult> InvokeAsync(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetCodespacesUserPublicKeyResult>("github:index/getCodespacesUserPublicKey:getCodespacesUserPublicKey", InvokeArgs.Empty, options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve information about a GitHub Codespaces User public key. This data source is required to be used with other GitHub secrets interactions.
        /// Note that the provider `token` must have admin rights to an user to retrieve it's Codespaces public key.
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Github = Pulumi.Github;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Github.GetCodespacesUserPublicKey.Invoke();
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetCodespacesUserPublicKeyResult> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetCodespacesUserPublicKeyResult>("github:index/getCodespacesUserPublicKey:getCodespacesUserPublicKey", InvokeArgs.Empty, options.WithDefaults());
    }


    [OutputType]
    public sealed class GetCodespacesUserPublicKeyResult
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

        [OutputConstructor]
        private GetCodespacesUserPublicKeyResult(
            string id,

            string key,

            string keyId)
        {
            Id = id;
            Key = key;
            KeyId = keyId;
        }
    }
}
