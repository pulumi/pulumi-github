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
    /// Provides a GitHub repository collaborator resource.
    /// 
    /// &gt; Note: github.RepositoryCollaborator cannot be used in conjunction with github.RepositoryCollaborators or
    /// they will fight over what your policy should be.
    /// 
    /// This resource allows you to add/remove collaborators from repositories in your
    /// organization or personal account. For organization repositories, collaborators can
    /// have explicit (and differing levels of) read, write, or administrator access to
    /// specific repositories, without giving the user full organization membership.
    /// For personal repositories, collaborators can only be granted write
    /// (implicitly includes read) permission.
    /// 
    /// When applied, an invitation will be sent to the user to become a collaborator
    /// on a repository. When destroyed, either the invitation will be cancelled or the
    /// collaborator will be removed from the repository.
    /// 
    /// This resource is non-authoritative, for managing ALL collaborators of a repo, use github.RepositoryCollaborators
    /// instead.
    /// 
    /// Further documentation on GitHub collaborators:
    /// 
    /// - [Adding outside collaborators to your personal repositories](https://help.github.com/en/github/setting-up-and-managing-your-github-user-account/managing-access-to-your-personal-repositories)
    /// - [Adding outside collaborators to repositories in your organization](https://help.github.com/articles/adding-outside-collaborators-to-repositories-in-your-organization/)
    /// - [Converting an organization member to an outside collaborator](https://help.github.com/articles/converting-an-organization-member-to-an-outside-collaborator/)
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
    ///     // Add a collaborator to a repository
    ///     var aRepoCollaborator = new Github.RepositoryCollaborator("aRepoCollaborator", new()
    ///     {
    ///         Permission = "admin",
    ///         Repository = "our-cool-repo",
    ///         Username = "SomeUser",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// GitHub Repository Collaborators can be imported using an ID made up of `repository:username`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import github:index/repositoryCollaborator:RepositoryCollaborator collaborator terraform:someuser
    /// ```
    /// </summary>
    [GithubResourceType("github:index/repositoryCollaborator:RepositoryCollaborator")]
    public partial class RepositoryCollaborator : global::Pulumi.CustomResource
    {
        /// <summary>
        /// ID of the invitation to be used in `github.UserInvitationAccepter`
        /// </summary>
        [Output("invitationId")]
        public Output<string> InvitationId { get; private set; } = null!;

        /// <summary>
        /// The permission of the outside collaborator for the repository.
        /// Must be one of `pull`, `push`, `maintain`, `triage` or `admin` or the name of an existing [custom repository role](https://docs.github.com/en/enterprise-cloud@latest/organizations/managing-peoples-access-to-your-organization-with-roles/managing-custom-repository-roles-for-an-organization) within the organization for organization-owned repositories.
        /// Must be `push` for personal repositories. Defaults to `push`.
        /// </summary>
        [Output("permission")]
        public Output<string?> Permission { get; private set; } = null!;

        /// <summary>
        /// Suppress plan diffs for `triage` and `maintain`.  Defaults to `false`.
        /// </summary>
        [Output("permissionDiffSuppression")]
        public Output<bool?> PermissionDiffSuppression { get; private set; } = null!;

        /// <summary>
        /// The GitHub repository
        /// </summary>
        [Output("repository")]
        public Output<string> Repository { get; private set; } = null!;

        /// <summary>
        /// The user to add to the repository as a collaborator.
        /// </summary>
        [Output("username")]
        public Output<string> Username { get; private set; } = null!;


        /// <summary>
        /// Create a RepositoryCollaborator resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RepositoryCollaborator(string name, RepositoryCollaboratorArgs args, CustomResourceOptions? options = null)
            : base("github:index/repositoryCollaborator:RepositoryCollaborator", name, args ?? new RepositoryCollaboratorArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RepositoryCollaborator(string name, Input<string> id, RepositoryCollaboratorState? state = null, CustomResourceOptions? options = null)
            : base("github:index/repositoryCollaborator:RepositoryCollaborator", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RepositoryCollaborator resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RepositoryCollaborator Get(string name, Input<string> id, RepositoryCollaboratorState? state = null, CustomResourceOptions? options = null)
        {
            return new RepositoryCollaborator(name, id, state, options);
        }
    }

    public sealed class RepositoryCollaboratorArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The permission of the outside collaborator for the repository.
        /// Must be one of `pull`, `push`, `maintain`, `triage` or `admin` or the name of an existing [custom repository role](https://docs.github.com/en/enterprise-cloud@latest/organizations/managing-peoples-access-to-your-organization-with-roles/managing-custom-repository-roles-for-an-organization) within the organization for organization-owned repositories.
        /// Must be `push` for personal repositories. Defaults to `push`.
        /// </summary>
        [Input("permission")]
        public Input<string>? Permission { get; set; }

        /// <summary>
        /// Suppress plan diffs for `triage` and `maintain`.  Defaults to `false`.
        /// </summary>
        [Input("permissionDiffSuppression")]
        public Input<bool>? PermissionDiffSuppression { get; set; }

        /// <summary>
        /// The GitHub repository
        /// </summary>
        [Input("repository", required: true)]
        public Input<string> Repository { get; set; } = null!;

        /// <summary>
        /// The user to add to the repository as a collaborator.
        /// </summary>
        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public RepositoryCollaboratorArgs()
        {
        }
        public static new RepositoryCollaboratorArgs Empty => new RepositoryCollaboratorArgs();
    }

    public sealed class RepositoryCollaboratorState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the invitation to be used in `github.UserInvitationAccepter`
        /// </summary>
        [Input("invitationId")]
        public Input<string>? InvitationId { get; set; }

        /// <summary>
        /// The permission of the outside collaborator for the repository.
        /// Must be one of `pull`, `push`, `maintain`, `triage` or `admin` or the name of an existing [custom repository role](https://docs.github.com/en/enterprise-cloud@latest/organizations/managing-peoples-access-to-your-organization-with-roles/managing-custom-repository-roles-for-an-organization) within the organization for organization-owned repositories.
        /// Must be `push` for personal repositories. Defaults to `push`.
        /// </summary>
        [Input("permission")]
        public Input<string>? Permission { get; set; }

        /// <summary>
        /// Suppress plan diffs for `triage` and `maintain`.  Defaults to `false`.
        /// </summary>
        [Input("permissionDiffSuppression")]
        public Input<bool>? PermissionDiffSuppression { get; set; }

        /// <summary>
        /// The GitHub repository
        /// </summary>
        [Input("repository")]
        public Input<string>? Repository { get; set; }

        /// <summary>
        /// The user to add to the repository as a collaborator.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public RepositoryCollaboratorState()
        {
        }
        public static new RepositoryCollaboratorState Empty => new RepositoryCollaboratorState();
    }
}
