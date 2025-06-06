// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github
{
    /// <summary>
    /// This resource allows you to create and manage settings for a GitHub Organization.
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
    ///     var test = new Github.OrganizationSettings("test", new()
    ///     {
    ///         BillingEmail = "test@example.com",
    ///         Company = "Test Company",
    ///         Blog = "https://example.com",
    ///         Email = "test@example.com",
    ///         TwitterUsername = "Test",
    ///         Location = "Test Location",
    ///         Name = "Test Name",
    ///         Description = "Test Description",
    ///         HasOrganizationProjects = true,
    ///         HasRepositoryProjects = true,
    ///         DefaultRepositoryPermission = "read",
    ///         MembersCanCreateRepositories = true,
    ///         MembersCanCreatePublicRepositories = true,
    ///         MembersCanCreatePrivateRepositories = true,
    ///         MembersCanCreateInternalRepositories = true,
    ///         MembersCanCreatePages = true,
    ///         MembersCanCreatePublicPages = true,
    ///         MembersCanCreatePrivatePages = true,
    ///         MembersCanForkPrivateRepositories = true,
    ///         WebCommitSignoffRequired = true,
    ///         AdvancedSecurityEnabledForNewRepositories = false,
    ///         DependabotAlertsEnabledForNewRepositories = false,
    ///         DependabotSecurityUpdatesEnabledForNewRepositories = false,
    ///         DependencyGraphEnabledForNewRepositories = false,
    ///         SecretScanningEnabledForNewRepositories = false,
    ///         SecretScanningPushProtectionEnabledForNewRepositories = false,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Organization settings can be imported using the `id` of the organization.
    /// The `id` of the organization can be found using the [get an organization](https://docs.github.com/en/rest/orgs/orgs#get-an-organization) API.
    /// 
    /// ```sh
    /// $ pulumi import github:index/organizationSettings:OrganizationSettings test 123456789
    /// ```
    /// </summary>
    [GithubResourceType("github:index/organizationSettings:OrganizationSettings")]
    public partial class OrganizationSettings : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Whether or not advanced security is enabled for new repositories. Defaults to `false`.
        /// </summary>
        [Output("advancedSecurityEnabledForNewRepositories")]
        public Output<bool?> AdvancedSecurityEnabledForNewRepositories { get; private set; } = null!;

        /// <summary>
        /// The billing email address for the organization.
        /// </summary>
        [Output("billingEmail")]
        public Output<string> BillingEmail { get; private set; } = null!;

        /// <summary>
        /// The blog URL for the organization.
        /// </summary>
        [Output("blog")]
        public Output<string?> Blog { get; private set; } = null!;

        /// <summary>
        /// The company name for the organization.
        /// </summary>
        [Output("company")]
        public Output<string?> Company { get; private set; } = null!;

        /// <summary>
        /// The default permission for organization members to create new repositories. Can be one of `read`, `write`, `admin`, or `none`. Defaults to `read`.
        /// </summary>
        [Output("defaultRepositoryPermission")]
        public Output<string?> DefaultRepositoryPermission { get; private set; } = null!;

        /// <summary>
        /// Whether or not dependabot alerts are enabled for new repositories. Defaults to `false`.
        /// </summary>
        [Output("dependabotAlertsEnabledForNewRepositories")]
        public Output<bool?> DependabotAlertsEnabledForNewRepositories { get; private set; } = null!;

        /// <summary>
        /// Whether or not dependabot security updates are enabled for new repositories. Defaults to `false`.
        /// </summary>
        [Output("dependabotSecurityUpdatesEnabledForNewRepositories")]
        public Output<bool?> DependabotSecurityUpdatesEnabledForNewRepositories { get; private set; } = null!;

        /// <summary>
        /// Whether or not dependency graph is enabled for new repositories. Defaults to `false`.
        /// </summary>
        [Output("dependencyGraphEnabledForNewRepositories")]
        public Output<bool?> DependencyGraphEnabledForNewRepositories { get; private set; } = null!;

        /// <summary>
        /// The description for the organization.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The email address for the organization.
        /// </summary>
        [Output("email")]
        public Output<string?> Email { get; private set; } = null!;

        /// <summary>
        /// Whether or not organization projects are enabled for the organization.
        /// </summary>
        [Output("hasOrganizationProjects")]
        public Output<bool?> HasOrganizationProjects { get; private set; } = null!;

        /// <summary>
        /// Whether or not repository projects are enabled for the organization.
        /// </summary>
        [Output("hasRepositoryProjects")]
        public Output<bool?> HasRepositoryProjects { get; private set; } = null!;

        /// <summary>
        /// The location for the organization.
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// Whether or not organization members can create new internal repositories. For Enterprise Organizations only.
        /// </summary>
        [Output("membersCanCreateInternalRepositories")]
        public Output<bool?> MembersCanCreateInternalRepositories { get; private set; } = null!;

        /// <summary>
        /// Whether or not organization members can create new pages. Defaults to `true`.
        /// </summary>
        [Output("membersCanCreatePages")]
        public Output<bool?> MembersCanCreatePages { get; private set; } = null!;

        /// <summary>
        /// Whether or not organization members can create new private pages. Defaults to `true`.
        /// </summary>
        [Output("membersCanCreatePrivatePages")]
        public Output<bool?> MembersCanCreatePrivatePages { get; private set; } = null!;

        /// <summary>
        /// Whether or not organization members can create new private repositories. Defaults to `true`.
        /// </summary>
        [Output("membersCanCreatePrivateRepositories")]
        public Output<bool?> MembersCanCreatePrivateRepositories { get; private set; } = null!;

        /// <summary>
        /// Whether or not organization members can create new public pages. Defaults to `true`.
        /// </summary>
        [Output("membersCanCreatePublicPages")]
        public Output<bool?> MembersCanCreatePublicPages { get; private set; } = null!;

        /// <summary>
        /// Whether or not organization members can create new public repositories. Defaults to `true`.
        /// </summary>
        [Output("membersCanCreatePublicRepositories")]
        public Output<bool?> MembersCanCreatePublicRepositories { get; private set; } = null!;

        /// <summary>
        /// Whether or not organization members can create new repositories. Defaults to `true`.
        /// </summary>
        [Output("membersCanCreateRepositories")]
        public Output<bool?> MembersCanCreateRepositories { get; private set; } = null!;

        /// <summary>
        /// Whether or not organization members can fork private repositories. Defaults to `false`.
        /// </summary>
        [Output("membersCanForkPrivateRepositories")]
        public Output<bool?> MembersCanForkPrivateRepositories { get; private set; } = null!;

        /// <summary>
        /// The name for the organization.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Whether or not secret scanning is enabled for new repositories. Defaults to `false`.
        /// </summary>
        [Output("secretScanningEnabledForNewRepositories")]
        public Output<bool?> SecretScanningEnabledForNewRepositories { get; private set; } = null!;

        /// <summary>
        /// Whether or not secret scanning push protection is enabled for new repositories. Defaults to `false`.
        /// </summary>
        [Output("secretScanningPushProtectionEnabledForNewRepositories")]
        public Output<bool?> SecretScanningPushProtectionEnabledForNewRepositories { get; private set; } = null!;

        /// <summary>
        /// The Twitter username for the organization.
        /// </summary>
        [Output("twitterUsername")]
        public Output<string?> TwitterUsername { get; private set; } = null!;

        /// <summary>
        /// Whether or not commit signatures are required for commits to the organization. Defaults to `false`.
        /// </summary>
        [Output("webCommitSignoffRequired")]
        public Output<bool?> WebCommitSignoffRequired { get; private set; } = null!;


        /// <summary>
        /// Create a OrganizationSettings resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public OrganizationSettings(string name, OrganizationSettingsArgs args, CustomResourceOptions? options = null)
            : base("github:index/organizationSettings:OrganizationSettings", name, args ?? new OrganizationSettingsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private OrganizationSettings(string name, Input<string> id, OrganizationSettingsState? state = null, CustomResourceOptions? options = null)
            : base("github:index/organizationSettings:OrganizationSettings", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing OrganizationSettings resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static OrganizationSettings Get(string name, Input<string> id, OrganizationSettingsState? state = null, CustomResourceOptions? options = null)
        {
            return new OrganizationSettings(name, id, state, options);
        }
    }

    public sealed class OrganizationSettingsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether or not advanced security is enabled for new repositories. Defaults to `false`.
        /// </summary>
        [Input("advancedSecurityEnabledForNewRepositories")]
        public Input<bool>? AdvancedSecurityEnabledForNewRepositories { get; set; }

        /// <summary>
        /// The billing email address for the organization.
        /// </summary>
        [Input("billingEmail", required: true)]
        public Input<string> BillingEmail { get; set; } = null!;

        /// <summary>
        /// The blog URL for the organization.
        /// </summary>
        [Input("blog")]
        public Input<string>? Blog { get; set; }

        /// <summary>
        /// The company name for the organization.
        /// </summary>
        [Input("company")]
        public Input<string>? Company { get; set; }

        /// <summary>
        /// The default permission for organization members to create new repositories. Can be one of `read`, `write`, `admin`, or `none`. Defaults to `read`.
        /// </summary>
        [Input("defaultRepositoryPermission")]
        public Input<string>? DefaultRepositoryPermission { get; set; }

        /// <summary>
        /// Whether or not dependabot alerts are enabled for new repositories. Defaults to `false`.
        /// </summary>
        [Input("dependabotAlertsEnabledForNewRepositories")]
        public Input<bool>? DependabotAlertsEnabledForNewRepositories { get; set; }

        /// <summary>
        /// Whether or not dependabot security updates are enabled for new repositories. Defaults to `false`.
        /// </summary>
        [Input("dependabotSecurityUpdatesEnabledForNewRepositories")]
        public Input<bool>? DependabotSecurityUpdatesEnabledForNewRepositories { get; set; }

        /// <summary>
        /// Whether or not dependency graph is enabled for new repositories. Defaults to `false`.
        /// </summary>
        [Input("dependencyGraphEnabledForNewRepositories")]
        public Input<bool>? DependencyGraphEnabledForNewRepositories { get; set; }

        /// <summary>
        /// The description for the organization.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The email address for the organization.
        /// </summary>
        [Input("email")]
        public Input<string>? Email { get; set; }

        /// <summary>
        /// Whether or not organization projects are enabled for the organization.
        /// </summary>
        [Input("hasOrganizationProjects")]
        public Input<bool>? HasOrganizationProjects { get; set; }

        /// <summary>
        /// Whether or not repository projects are enabled for the organization.
        /// </summary>
        [Input("hasRepositoryProjects")]
        public Input<bool>? HasRepositoryProjects { get; set; }

        /// <summary>
        /// The location for the organization.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Whether or not organization members can create new internal repositories. For Enterprise Organizations only.
        /// </summary>
        [Input("membersCanCreateInternalRepositories")]
        public Input<bool>? MembersCanCreateInternalRepositories { get; set; }

        /// <summary>
        /// Whether or not organization members can create new pages. Defaults to `true`.
        /// </summary>
        [Input("membersCanCreatePages")]
        public Input<bool>? MembersCanCreatePages { get; set; }

        /// <summary>
        /// Whether or not organization members can create new private pages. Defaults to `true`.
        /// </summary>
        [Input("membersCanCreatePrivatePages")]
        public Input<bool>? MembersCanCreatePrivatePages { get; set; }

        /// <summary>
        /// Whether or not organization members can create new private repositories. Defaults to `true`.
        /// </summary>
        [Input("membersCanCreatePrivateRepositories")]
        public Input<bool>? MembersCanCreatePrivateRepositories { get; set; }

        /// <summary>
        /// Whether or not organization members can create new public pages. Defaults to `true`.
        /// </summary>
        [Input("membersCanCreatePublicPages")]
        public Input<bool>? MembersCanCreatePublicPages { get; set; }

        /// <summary>
        /// Whether or not organization members can create new public repositories. Defaults to `true`.
        /// </summary>
        [Input("membersCanCreatePublicRepositories")]
        public Input<bool>? MembersCanCreatePublicRepositories { get; set; }

        /// <summary>
        /// Whether or not organization members can create new repositories. Defaults to `true`.
        /// </summary>
        [Input("membersCanCreateRepositories")]
        public Input<bool>? MembersCanCreateRepositories { get; set; }

        /// <summary>
        /// Whether or not organization members can fork private repositories. Defaults to `false`.
        /// </summary>
        [Input("membersCanForkPrivateRepositories")]
        public Input<bool>? MembersCanForkPrivateRepositories { get; set; }

        /// <summary>
        /// The name for the organization.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Whether or not secret scanning is enabled for new repositories. Defaults to `false`.
        /// </summary>
        [Input("secretScanningEnabledForNewRepositories")]
        public Input<bool>? SecretScanningEnabledForNewRepositories { get; set; }

        /// <summary>
        /// Whether or not secret scanning push protection is enabled for new repositories. Defaults to `false`.
        /// </summary>
        [Input("secretScanningPushProtectionEnabledForNewRepositories")]
        public Input<bool>? SecretScanningPushProtectionEnabledForNewRepositories { get; set; }

        /// <summary>
        /// The Twitter username for the organization.
        /// </summary>
        [Input("twitterUsername")]
        public Input<string>? TwitterUsername { get; set; }

        /// <summary>
        /// Whether or not commit signatures are required for commits to the organization. Defaults to `false`.
        /// </summary>
        [Input("webCommitSignoffRequired")]
        public Input<bool>? WebCommitSignoffRequired { get; set; }

        public OrganizationSettingsArgs()
        {
        }
        public static new OrganizationSettingsArgs Empty => new OrganizationSettingsArgs();
    }

    public sealed class OrganizationSettingsState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether or not advanced security is enabled for new repositories. Defaults to `false`.
        /// </summary>
        [Input("advancedSecurityEnabledForNewRepositories")]
        public Input<bool>? AdvancedSecurityEnabledForNewRepositories { get; set; }

        /// <summary>
        /// The billing email address for the organization.
        /// </summary>
        [Input("billingEmail")]
        public Input<string>? BillingEmail { get; set; }

        /// <summary>
        /// The blog URL for the organization.
        /// </summary>
        [Input("blog")]
        public Input<string>? Blog { get; set; }

        /// <summary>
        /// The company name for the organization.
        /// </summary>
        [Input("company")]
        public Input<string>? Company { get; set; }

        /// <summary>
        /// The default permission for organization members to create new repositories. Can be one of `read`, `write`, `admin`, or `none`. Defaults to `read`.
        /// </summary>
        [Input("defaultRepositoryPermission")]
        public Input<string>? DefaultRepositoryPermission { get; set; }

        /// <summary>
        /// Whether or not dependabot alerts are enabled for new repositories. Defaults to `false`.
        /// </summary>
        [Input("dependabotAlertsEnabledForNewRepositories")]
        public Input<bool>? DependabotAlertsEnabledForNewRepositories { get; set; }

        /// <summary>
        /// Whether or not dependabot security updates are enabled for new repositories. Defaults to `false`.
        /// </summary>
        [Input("dependabotSecurityUpdatesEnabledForNewRepositories")]
        public Input<bool>? DependabotSecurityUpdatesEnabledForNewRepositories { get; set; }

        /// <summary>
        /// Whether or not dependency graph is enabled for new repositories. Defaults to `false`.
        /// </summary>
        [Input("dependencyGraphEnabledForNewRepositories")]
        public Input<bool>? DependencyGraphEnabledForNewRepositories { get; set; }

        /// <summary>
        /// The description for the organization.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The email address for the organization.
        /// </summary>
        [Input("email")]
        public Input<string>? Email { get; set; }

        /// <summary>
        /// Whether or not organization projects are enabled for the organization.
        /// </summary>
        [Input("hasOrganizationProjects")]
        public Input<bool>? HasOrganizationProjects { get; set; }

        /// <summary>
        /// Whether or not repository projects are enabled for the organization.
        /// </summary>
        [Input("hasRepositoryProjects")]
        public Input<bool>? HasRepositoryProjects { get; set; }

        /// <summary>
        /// The location for the organization.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Whether or not organization members can create new internal repositories. For Enterprise Organizations only.
        /// </summary>
        [Input("membersCanCreateInternalRepositories")]
        public Input<bool>? MembersCanCreateInternalRepositories { get; set; }

        /// <summary>
        /// Whether or not organization members can create new pages. Defaults to `true`.
        /// </summary>
        [Input("membersCanCreatePages")]
        public Input<bool>? MembersCanCreatePages { get; set; }

        /// <summary>
        /// Whether or not organization members can create new private pages. Defaults to `true`.
        /// </summary>
        [Input("membersCanCreatePrivatePages")]
        public Input<bool>? MembersCanCreatePrivatePages { get; set; }

        /// <summary>
        /// Whether or not organization members can create new private repositories. Defaults to `true`.
        /// </summary>
        [Input("membersCanCreatePrivateRepositories")]
        public Input<bool>? MembersCanCreatePrivateRepositories { get; set; }

        /// <summary>
        /// Whether or not organization members can create new public pages. Defaults to `true`.
        /// </summary>
        [Input("membersCanCreatePublicPages")]
        public Input<bool>? MembersCanCreatePublicPages { get; set; }

        /// <summary>
        /// Whether or not organization members can create new public repositories. Defaults to `true`.
        /// </summary>
        [Input("membersCanCreatePublicRepositories")]
        public Input<bool>? MembersCanCreatePublicRepositories { get; set; }

        /// <summary>
        /// Whether or not organization members can create new repositories. Defaults to `true`.
        /// </summary>
        [Input("membersCanCreateRepositories")]
        public Input<bool>? MembersCanCreateRepositories { get; set; }

        /// <summary>
        /// Whether or not organization members can fork private repositories. Defaults to `false`.
        /// </summary>
        [Input("membersCanForkPrivateRepositories")]
        public Input<bool>? MembersCanForkPrivateRepositories { get; set; }

        /// <summary>
        /// The name for the organization.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Whether or not secret scanning is enabled for new repositories. Defaults to `false`.
        /// </summary>
        [Input("secretScanningEnabledForNewRepositories")]
        public Input<bool>? SecretScanningEnabledForNewRepositories { get; set; }

        /// <summary>
        /// Whether or not secret scanning push protection is enabled for new repositories. Defaults to `false`.
        /// </summary>
        [Input("secretScanningPushProtectionEnabledForNewRepositories")]
        public Input<bool>? SecretScanningPushProtectionEnabledForNewRepositories { get; set; }

        /// <summary>
        /// The Twitter username for the organization.
        /// </summary>
        [Input("twitterUsername")]
        public Input<string>? TwitterUsername { get; set; }

        /// <summary>
        /// Whether or not commit signatures are required for commits to the organization. Defaults to `false`.
        /// </summary>
        [Input("webCommitSignoffRequired")]
        public Input<bool>? WebCommitSignoffRequired { get; set; }

        public OrganizationSettingsState()
        {
        }
        public static new OrganizationSettingsState Empty => new OrganizationSettingsState();
    }
}
