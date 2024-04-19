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
    /// &gt; Note: github.TeamRepository cannot be used in conjunction with github.RepositoryCollaborators or
    /// they will fight over what your policy should be.
    /// 
    /// This resource manages relationships between teams and repositories
    /// in your GitHub organization.
    /// 
    /// Creating this resource grants a particular team permissions on a
    /// particular repository.
    /// 
    /// The repository and the team must both belong to the same organization
    /// on GitHub. This resource does not actually *create* any repositories;
    /// to do that, see `github.Repository`.
    /// 
    /// This resource is non-authoritative, for managing ALL collaborators of a repo, use github.RepositoryCollaborators
    /// instead.
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
    ///     // Add a repository to the team
    ///     var someTeam = new Github.Team("some_team", new()
    ///     {
    ///         Name = "SomeTeam",
    ///         Description = "Some cool team",
    ///     });
    /// 
    ///     var someRepo = new Github.Repository("some_repo", new()
    ///     {
    ///         Name = "some-repo",
    ///     });
    /// 
    ///     var someTeamRepo = new Github.TeamRepository("some_team_repo", new()
    ///     {
    ///         TeamId = someTeam.Id,
    ///         Repository = someRepo.Name,
    ///         Permission = "pull",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// GitHub Team Repository can be imported using an ID made up of `team_id:repository` or `team_name:repository`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import github:index/teamRepository:TeamRepository terraform_repo 1234567:terraform
    /// ```
    /// 
    /// ```sh
    /// $ pulumi import github:index/teamRepository:TeamRepository terraform_repo Administrators:terraform
    /// ```
    /// </summary>
    [GithubResourceType("github:index/teamRepository:TeamRepository")]
    public partial class TeamRepository : global::Pulumi.CustomResource
    {
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// The permissions of team members regarding the repository.
        /// Must be one of `pull`, `triage`, `push`, `maintain`, `admin` or the name of an existing [custom repository role](https://docs.github.com/en/enterprise-cloud@latest/organizations/managing-peoples-access-to-your-organization-with-roles/managing-custom-repository-roles-for-an-organization) within the organisation. Defaults to `pull`.
        /// </summary>
        [Output("permission")]
        public Output<string?> Permission { get; private set; } = null!;

        /// <summary>
        /// The repository to add to the team.
        /// </summary>
        [Output("repository")]
        public Output<string> Repository { get; private set; } = null!;

        /// <summary>
        /// The GitHub team id or the GitHub team slug
        /// </summary>
        [Output("teamId")]
        public Output<string> TeamId { get; private set; } = null!;


        /// <summary>
        /// Create a TeamRepository resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TeamRepository(string name, TeamRepositoryArgs args, CustomResourceOptions? options = null)
            : base("github:index/teamRepository:TeamRepository", name, args ?? new TeamRepositoryArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TeamRepository(string name, Input<string> id, TeamRepositoryState? state = null, CustomResourceOptions? options = null)
            : base("github:index/teamRepository:TeamRepository", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing TeamRepository resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TeamRepository Get(string name, Input<string> id, TeamRepositoryState? state = null, CustomResourceOptions? options = null)
        {
            return new TeamRepository(name, id, state, options);
        }
    }

    public sealed class TeamRepositoryArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The permissions of team members regarding the repository.
        /// Must be one of `pull`, `triage`, `push`, `maintain`, `admin` or the name of an existing [custom repository role](https://docs.github.com/en/enterprise-cloud@latest/organizations/managing-peoples-access-to-your-organization-with-roles/managing-custom-repository-roles-for-an-organization) within the organisation. Defaults to `pull`.
        /// </summary>
        [Input("permission")]
        public Input<string>? Permission { get; set; }

        /// <summary>
        /// The repository to add to the team.
        /// </summary>
        [Input("repository", required: true)]
        public Input<string> Repository { get; set; } = null!;

        /// <summary>
        /// The GitHub team id or the GitHub team slug
        /// </summary>
        [Input("teamId", required: true)]
        public Input<string> TeamId { get; set; } = null!;

        public TeamRepositoryArgs()
        {
        }
        public static new TeamRepositoryArgs Empty => new TeamRepositoryArgs();
    }

    public sealed class TeamRepositoryState : global::Pulumi.ResourceArgs
    {
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// The permissions of team members regarding the repository.
        /// Must be one of `pull`, `triage`, `push`, `maintain`, `admin` or the name of an existing [custom repository role](https://docs.github.com/en/enterprise-cloud@latest/organizations/managing-peoples-access-to-your-organization-with-roles/managing-custom-repository-roles-for-an-organization) within the organisation. Defaults to `pull`.
        /// </summary>
        [Input("permission")]
        public Input<string>? Permission { get; set; }

        /// <summary>
        /// The repository to add to the team.
        /// </summary>
        [Input("repository")]
        public Input<string>? Repository { get; set; }

        /// <summary>
        /// The GitHub team id or the GitHub team slug
        /// </summary>
        [Input("teamId")]
        public Input<string>? TeamId { get; set; }

        public TeamRepositoryState()
        {
        }
        public static new TeamRepositoryState Empty => new TeamRepositoryState();
    }
}
