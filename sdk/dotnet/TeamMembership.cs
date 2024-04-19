// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github
{
    /// <summary>
    /// Provides a GitHub team membership resource.
    /// 
    /// This resource allows you to add/remove users from teams in your organization. When applied,
    /// the user will be added to the team. If the user hasn't accepted their invitation to the
    /// organization, they won't be part of the team until they do. When
    /// destroyed, the user will be removed from the team.
    /// 
    /// &gt; **Note** This resource is not compatible with `github.TeamMembers`. Use either `github.TeamMembers` or `github.TeamMembership`.
    /// 
    /// &gt; **Note** Organization owners may not be set as "members" of a team; they may only be set as "maintainers". Attempting to set organization an owner to "member" of a may result in a `pulumi preview` diff that changes their status back to "maintainer".
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
    ///     // Add a user to the organization
    ///     var membershipForSomeUser = new Github.Membership("membership_for_some_user", new()
    ///     {
    ///         Username = "SomeUser",
    ///         Role = "member",
    ///     });
    /// 
    ///     var someTeam = new Github.Team("some_team", new()
    ///     {
    ///         Name = "SomeTeam",
    ///         Description = "Some cool team",
    ///     });
    /// 
    ///     var someTeamMembership = new Github.TeamMembership("some_team_membership", new()
    ///     {
    ///         TeamId = someTeam.Id,
    ///         Username = "SomeUser",
    ///         Role = "member",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// GitHub Team Membership can be imported using an ID made up of `teamid:username` or `teamname:username`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import github:index/teamMembership:TeamMembership member 1234567:someuser
    /// ```
    /// 
    /// ```sh
    /// $ pulumi import github:index/teamMembership:TeamMembership member Administrators:someuser
    /// ```
    /// </summary>
    [GithubResourceType("github:index/teamMembership:TeamMembership")]
    public partial class TeamMembership : global::Pulumi.CustomResource
    {
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// The role of the user within the team.
        /// Must be one of `member` or `maintainer`. Defaults to `member`.
        /// </summary>
        [Output("role")]
        public Output<string?> Role { get; private set; } = null!;

        /// <summary>
        /// The GitHub team id or the GitHub team slug
        /// </summary>
        [Output("teamId")]
        public Output<string> TeamId { get; private set; } = null!;

        /// <summary>
        /// The user to add to the team.
        /// </summary>
        [Output("username")]
        public Output<string> Username { get; private set; } = null!;


        /// <summary>
        /// Create a TeamMembership resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TeamMembership(string name, TeamMembershipArgs args, CustomResourceOptions? options = null)
            : base("github:index/teamMembership:TeamMembership", name, args ?? new TeamMembershipArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TeamMembership(string name, Input<string> id, TeamMembershipState? state = null, CustomResourceOptions? options = null)
            : base("github:index/teamMembership:TeamMembership", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing TeamMembership resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TeamMembership Get(string name, Input<string> id, TeamMembershipState? state = null, CustomResourceOptions? options = null)
        {
            return new TeamMembership(name, id, state, options);
        }
    }

    public sealed class TeamMembershipArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The role of the user within the team.
        /// Must be one of `member` or `maintainer`. Defaults to `member`.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        /// <summary>
        /// The GitHub team id or the GitHub team slug
        /// </summary>
        [Input("teamId", required: true)]
        public Input<string> TeamId { get; set; } = null!;

        /// <summary>
        /// The user to add to the team.
        /// </summary>
        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public TeamMembershipArgs()
        {
        }
        public static new TeamMembershipArgs Empty => new TeamMembershipArgs();
    }

    public sealed class TeamMembershipState : global::Pulumi.ResourceArgs
    {
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// The role of the user within the team.
        /// Must be one of `member` or `maintainer`. Defaults to `member`.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        /// <summary>
        /// The GitHub team id or the GitHub team slug
        /// </summary>
        [Input("teamId")]
        public Input<string>? TeamId { get; set; }

        /// <summary>
        /// The user to add to the team.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public TeamMembershipState()
        {
        }
        public static new TeamMembershipState Empty => new TeamMembershipState();
    }
}
