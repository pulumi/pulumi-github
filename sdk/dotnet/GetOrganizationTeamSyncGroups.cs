// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github
{
    public static class GetOrganizationTeamSyncGroups
    {
        /// <summary>
        /// Use this data source to retrieve the identity provider (IdP) groups for an organization.
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
        ///     var test = Github.GetOrganizationTeamSyncGroups.Invoke();
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetOrganizationTeamSyncGroupsResult> InvokeAsync(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetOrganizationTeamSyncGroupsResult>("github:index/getOrganizationTeamSyncGroups:getOrganizationTeamSyncGroups", InvokeArgs.Empty, options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve the identity provider (IdP) groups for an organization.
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
        ///     var test = Github.GetOrganizationTeamSyncGroups.Invoke();
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetOrganizationTeamSyncGroupsResult> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetOrganizationTeamSyncGroupsResult>("github:index/getOrganizationTeamSyncGroups:getOrganizationTeamSyncGroups", InvokeArgs.Empty, options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve the identity provider (IdP) groups for an organization.
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
        ///     var test = Github.GetOrganizationTeamSyncGroups.Invoke();
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetOrganizationTeamSyncGroupsResult> Invoke(InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetOrganizationTeamSyncGroupsResult>("github:index/getOrganizationTeamSyncGroups:getOrganizationTeamSyncGroups", InvokeArgs.Empty, options.WithDefaults());
    }


    [OutputType]
    public sealed class GetOrganizationTeamSyncGroupsResult
    {
        /// <summary>
        /// An Array of GitHub Identity Provider Groups.  Each `group` block consists of the fields documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetOrganizationTeamSyncGroupsGroupResult> Groups;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetOrganizationTeamSyncGroupsResult(
            ImmutableArray<Outputs.GetOrganizationTeamSyncGroupsGroupResult> groups,

            string id)
        {
            Groups = groups;
            Id = id;
        }
    }
}
