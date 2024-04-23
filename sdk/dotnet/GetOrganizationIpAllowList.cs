// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github
{
    public static class GetOrganizationIpAllowList
    {
        /// <summary>
        /// Use this data source to retrieve information about the IP allow list of an organization.
        /// The allow list for IP addresses will block access to private resources via the web, API,
        /// and Git from any IP addresses that are not on the allow list.
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
        ///     var all = Github.GetOrganizationIpAllowList.Invoke();
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetOrganizationIpAllowListResult> InvokeAsync(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetOrganizationIpAllowListResult>("github:index/getOrganizationIpAllowList:getOrganizationIpAllowList", InvokeArgs.Empty, options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve information about the IP allow list of an organization.
        /// The allow list for IP addresses will block access to private resources via the web, API,
        /// and Git from any IP addresses that are not on the allow list.
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
        ///     var all = Github.GetOrganizationIpAllowList.Invoke();
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetOrganizationIpAllowListResult> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetOrganizationIpAllowListResult>("github:index/getOrganizationIpAllowList:getOrganizationIpAllowList", InvokeArgs.Empty, options.WithDefaults());
    }


    [OutputType]
    public sealed class GetOrganizationIpAllowListResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// An Array of allowed IP addresses.
        /// ___
        /// </summary>
        public readonly ImmutableArray<Outputs.GetOrganizationIpAllowListIpAllowListResult> IpAllowLists;

        [OutputConstructor]
        private GetOrganizationIpAllowListResult(
            string id,

            ImmutableArray<Outputs.GetOrganizationIpAllowListIpAllowListResult> ipAllowLists)
        {
            Id = id;
            IpAllowLists = ipAllowLists;
        }
    }
}
