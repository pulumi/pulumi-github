// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github
{
    public static class GetDependabotOrganizationPublicKey
    {
        /// <summary>
        /// Use this data source to retrieve information about a GitHub Dependabot Organization public key. This data source is required to be used with other GitHub secrets interactions.
        /// Note that the provider `token` must have admin rights to an organization to retrieve it's Dependabot public key.
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
        ///     var example = Github.GetDependabotOrganizationPublicKey.Invoke();
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetDependabotOrganizationPublicKeyResult> InvokeAsync(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDependabotOrganizationPublicKeyResult>("github:index/getDependabotOrganizationPublicKey:getDependabotOrganizationPublicKey", InvokeArgs.Empty, options.WithDefaults());
    }


    [OutputType]
    public sealed class GetDependabotOrganizationPublicKeyResult
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
        private GetDependabotOrganizationPublicKeyResult(
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
