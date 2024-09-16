// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github
{
    public static class GetOrganization
    {
        /// <summary>
        /// Use this data source to retrieve basic information about a GitHub Organization.
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
        ///     var example = Github.GetOrganization.Invoke(new()
        ///     {
        ///         Name = "github",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetOrganizationResult> InvokeAsync(GetOrganizationArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetOrganizationResult>("github:index/getOrganization:getOrganization", args ?? new GetOrganizationArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve basic information about a GitHub Organization.
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
        ///     var example = Github.GetOrganization.Invoke(new()
        ///     {
        ///         Name = "github",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetOrganizationResult> Invoke(GetOrganizationInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetOrganizationResult>("github:index/getOrganization:getOrganization", args ?? new GetOrganizationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetOrganizationArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Whether or not to include archived repos in the `repositories` list. Defaults to `false`.
        /// </summary>
        [Input("ignoreArchivedRepos")]
        public bool? IgnoreArchivedRepos { get; set; }

        /// <summary>
        /// The name of the organization.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Exclude the repos, members and other attributes from the returned result. Defaults to `false`.
        /// </summary>
        [Input("summaryOnly")]
        public bool? SummaryOnly { get; set; }

        public GetOrganizationArgs()
        {
        }
        public static new GetOrganizationArgs Empty => new GetOrganizationArgs();
    }

    public sealed class GetOrganizationInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Whether or not to include archived repos in the `repositories` list. Defaults to `false`.
        /// </summary>
        [Input("ignoreArchivedRepos")]
        public Input<bool>? IgnoreArchivedRepos { get; set; }

        /// <summary>
        /// The name of the organization.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Exclude the repos, members and other attributes from the returned result. Defaults to `false`.
        /// </summary>
        [Input("summaryOnly")]
        public Input<bool>? SummaryOnly { get; set; }

        public GetOrganizationInvokeArgs()
        {
        }
        public static new GetOrganizationInvokeArgs Empty => new GetOrganizationInvokeArgs();
    }


    [OutputType]
    public sealed class GetOrganizationResult
    {
        /// <summary>
        /// Whether advanced security is enabled for new repositories.
        /// </summary>
        public readonly bool AdvancedSecurityEnabledForNewRepositories;
        /// <summary>
        /// Default permission level members have for organization repositories.
        /// </summary>
        public readonly string DefaultRepositoryPermission;
        /// <summary>
        /// Whether Dependabot alerts is automatically enabled for new repositories.
        /// </summary>
        public readonly bool DependabotAlertsEnabledForNewRepositories;
        /// <summary>
        /// Whether Dependabot security updates is automatically enabled for new repositories.
        /// </summary>
        public readonly bool DependabotSecurityUpdatesEnabledForNewRepositories;
        /// <summary>
        /// Whether dependency graph is automatically enabled for new repositories.
        /// </summary>
        public readonly bool DependencyGraphEnabledForNewRepositories;
        /// <summary>
        /// The organization account description
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly bool? IgnoreArchivedRepos;
        /// <summary>
        /// The members login
        /// </summary>
        public readonly string Login;
        /// <summary>
        /// **Deprecated**: use `users` instead by replacing `github_organization.example.members` to `github_organization.example.users[*].login` which will give you the same value, expect this field to be removed in next major version
        /// </summary>
        public readonly ImmutableArray<string> Members;
        /// <summary>
        /// The type of repository allowed to be created by members of the organization. Can be one of `ALL`, `PUBLIC`, `PRIVATE`, `NONE`.
        /// </summary>
        public readonly string MembersAllowedRepositoryCreationType;
        /// <summary>
        /// Whether organization members can create internal repositories.
        /// </summary>
        public readonly bool MembersCanCreateInternalRepositories;
        /// <summary>
        /// Whether organization members can create pages sites.
        /// </summary>
        public readonly bool MembersCanCreatePages;
        /// <summary>
        /// Whether organization members can create private pages sites.
        /// </summary>
        public readonly bool MembersCanCreatePrivatePages;
        /// <summary>
        /// Whether organization members can create private repositories.
        /// </summary>
        public readonly bool MembersCanCreatePrivateRepositories;
        /// <summary>
        /// Whether organization members can create public pages sites.
        /// </summary>
        public readonly bool MembersCanCreatePublicPages;
        /// <summary>
        /// Whether organization members can create public repositories.
        /// </summary>
        public readonly bool MembersCanCreatePublicRepositories;
        /// <summary>
        /// Whether non-admin organization members can create repositories.
        /// </summary>
        public readonly bool MembersCanCreateRepositories;
        /// <summary>
        /// Whether organization members can create private repository forks.
        /// </summary>
        public readonly bool MembersCanForkPrivateRepositories;
        /// <summary>
        /// The organization's public profile name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// GraphQL global node ID for use with the v4 API
        /// </summary>
        public readonly string NodeId;
        /// <summary>
        /// The organization's name as used in URLs and the API
        /// </summary>
        public readonly string Orgname;
        /// <summary>
        /// The organization account plan name
        /// </summary>
        public readonly string Plan;
        /// <summary>
        /// (`list`) A list of the full names of the repositories in the organization formatted as `owner/name` strings
        /// </summary>
        public readonly ImmutableArray<string> Repositories;
        /// <summary>
        /// Whether secret scanning is automatically enabled for new repositories.
        /// </summary>
        public readonly bool SecretScanningEnabledForNewRepositories;
        /// <summary>
        /// Whether secret scanning push protection is automatically enabled for new repositories.
        /// </summary>
        public readonly bool SecretScanningPushProtectionEnabledForNewRepositories;
        public readonly bool? SummaryOnly;
        /// <summary>
        /// Whether two-factor authentication is required for all members of the organization.
        /// </summary>
        public readonly bool TwoFactorRequirementEnabled;
        /// <summary>
        /// (`list`) A list with the members of the organization with following fields:
        /// </summary>
        public readonly ImmutableArray<ImmutableDictionary<string, string>> Users;
        /// <summary>
        /// Whether organization members must sign all commits.
        /// </summary>
        public readonly bool WebCommitSignoffRequired;

        [OutputConstructor]
        private GetOrganizationResult(
            bool advancedSecurityEnabledForNewRepositories,

            string defaultRepositoryPermission,

            bool dependabotAlertsEnabledForNewRepositories,

            bool dependabotSecurityUpdatesEnabledForNewRepositories,

            bool dependencyGraphEnabledForNewRepositories,

            string description,

            string id,

            bool? ignoreArchivedRepos,

            string login,

            ImmutableArray<string> members,

            string membersAllowedRepositoryCreationType,

            bool membersCanCreateInternalRepositories,

            bool membersCanCreatePages,

            bool membersCanCreatePrivatePages,

            bool membersCanCreatePrivateRepositories,

            bool membersCanCreatePublicPages,

            bool membersCanCreatePublicRepositories,

            bool membersCanCreateRepositories,

            bool membersCanForkPrivateRepositories,

            string name,

            string nodeId,

            string orgname,

            string plan,

            ImmutableArray<string> repositories,

            bool secretScanningEnabledForNewRepositories,

            bool secretScanningPushProtectionEnabledForNewRepositories,

            bool? summaryOnly,

            bool twoFactorRequirementEnabled,

            ImmutableArray<ImmutableDictionary<string, string>> users,

            bool webCommitSignoffRequired)
        {
            AdvancedSecurityEnabledForNewRepositories = advancedSecurityEnabledForNewRepositories;
            DefaultRepositoryPermission = defaultRepositoryPermission;
            DependabotAlertsEnabledForNewRepositories = dependabotAlertsEnabledForNewRepositories;
            DependabotSecurityUpdatesEnabledForNewRepositories = dependabotSecurityUpdatesEnabledForNewRepositories;
            DependencyGraphEnabledForNewRepositories = dependencyGraphEnabledForNewRepositories;
            Description = description;
            Id = id;
            IgnoreArchivedRepos = ignoreArchivedRepos;
            Login = login;
            Members = members;
            MembersAllowedRepositoryCreationType = membersAllowedRepositoryCreationType;
            MembersCanCreateInternalRepositories = membersCanCreateInternalRepositories;
            MembersCanCreatePages = membersCanCreatePages;
            MembersCanCreatePrivatePages = membersCanCreatePrivatePages;
            MembersCanCreatePrivateRepositories = membersCanCreatePrivateRepositories;
            MembersCanCreatePublicPages = membersCanCreatePublicPages;
            MembersCanCreatePublicRepositories = membersCanCreatePublicRepositories;
            MembersCanCreateRepositories = membersCanCreateRepositories;
            MembersCanForkPrivateRepositories = membersCanForkPrivateRepositories;
            Name = name;
            NodeId = nodeId;
            Orgname = orgname;
            Plan = plan;
            Repositories = repositories;
            SecretScanningEnabledForNewRepositories = secretScanningEnabledForNewRepositories;
            SecretScanningPushProtectionEnabledForNewRepositories = secretScanningPushProtectionEnabledForNewRepositories;
            SummaryOnly = summaryOnly;
            TwoFactorRequirementEnabled = twoFactorRequirementEnabled;
            Users = users;
            WebCommitSignoffRequired = webCommitSignoffRequired;
        }
    }
}
