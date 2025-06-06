// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github
{
    public static class GetOrganizationExternalIdentities
    {
        /// <summary>
        /// Use this data source to retrieve each organization member's SAML or SCIM user
        /// attributes.
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
        ///     var all = Github.GetOrganizationExternalIdentities.Invoke();
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetOrganizationExternalIdentitiesResult> InvokeAsync(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetOrganizationExternalIdentitiesResult>("github:index/getOrganizationExternalIdentities:getOrganizationExternalIdentities", InvokeArgs.Empty, options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve each organization member's SAML or SCIM user
        /// attributes.
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
        ///     var all = Github.GetOrganizationExternalIdentities.Invoke();
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetOrganizationExternalIdentitiesResult> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetOrganizationExternalIdentitiesResult>("github:index/getOrganizationExternalIdentities:getOrganizationExternalIdentities", InvokeArgs.Empty, options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve each organization member's SAML or SCIM user
        /// attributes.
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
        ///     var all = Github.GetOrganizationExternalIdentities.Invoke();
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetOrganizationExternalIdentitiesResult> Invoke(InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetOrganizationExternalIdentitiesResult>("github:index/getOrganizationExternalIdentities:getOrganizationExternalIdentities", InvokeArgs.Empty, options.WithDefaults());
    }


    [OutputType]
    public sealed class GetOrganizationExternalIdentitiesResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// An Array of identities returned from GitHub
        /// </summary>
        public readonly ImmutableArray<Outputs.GetOrganizationExternalIdentitiesIdentityResult> Identities;

        [OutputConstructor]
        private GetOrganizationExternalIdentitiesResult(
            string id,

            ImmutableArray<Outputs.GetOrganizationExternalIdentitiesIdentityResult> identities)
        {
            Id = id;
            Identities = identities;
        }
    }
}
