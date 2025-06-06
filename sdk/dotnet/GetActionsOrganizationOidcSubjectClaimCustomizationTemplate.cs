// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github
{
    public static class GetActionsOrganizationOidcSubjectClaimCustomizationTemplate
    {
        /// <summary>
        /// Use this data source to retrieve the OpenID Connect subject claim customization template for an organization
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
        ///     var example = Github.GetActionsOrganizationOidcSubjectClaimCustomizationTemplate.Invoke();
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetActionsOrganizationOidcSubjectClaimCustomizationTemplateResult> InvokeAsync(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetActionsOrganizationOidcSubjectClaimCustomizationTemplateResult>("github:index/getActionsOrganizationOidcSubjectClaimCustomizationTemplate:getActionsOrganizationOidcSubjectClaimCustomizationTemplate", InvokeArgs.Empty, options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve the OpenID Connect subject claim customization template for an organization
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
        ///     var example = Github.GetActionsOrganizationOidcSubjectClaimCustomizationTemplate.Invoke();
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetActionsOrganizationOidcSubjectClaimCustomizationTemplateResult> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetActionsOrganizationOidcSubjectClaimCustomizationTemplateResult>("github:index/getActionsOrganizationOidcSubjectClaimCustomizationTemplate:getActionsOrganizationOidcSubjectClaimCustomizationTemplate", InvokeArgs.Empty, options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve the OpenID Connect subject claim customization template for an organization
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
        ///     var example = Github.GetActionsOrganizationOidcSubjectClaimCustomizationTemplate.Invoke();
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetActionsOrganizationOidcSubjectClaimCustomizationTemplateResult> Invoke(InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetActionsOrganizationOidcSubjectClaimCustomizationTemplateResult>("github:index/getActionsOrganizationOidcSubjectClaimCustomizationTemplate:getActionsOrganizationOidcSubjectClaimCustomizationTemplate", InvokeArgs.Empty, options.WithDefaults());
    }


    [OutputType]
    public sealed class GetActionsOrganizationOidcSubjectClaimCustomizationTemplateResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The list of OpenID Connect claim keys.
        /// </summary>
        public readonly ImmutableArray<string> IncludeClaimKeys;

        [OutputConstructor]
        private GetActionsOrganizationOidcSubjectClaimCustomizationTemplateResult(
            string id,

            ImmutableArray<string> includeClaimKeys)
        {
            Id = id;
            IncludeClaimKeys = includeClaimKeys;
        }
    }
}
