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
    /// Provides a resource to manage GitHub repository collaborator invitations.
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
    ///     var example = new Github.Repository("example", new()
    ///     {
    ///         Name = "example-repo",
    ///     });
    /// 
    ///     var exampleRepositoryCollaborator = new Github.RepositoryCollaborator("example", new()
    ///     {
    ///         Repository = example.Name,
    ///         Username = "example-username",
    ///         Permission = "push",
    ///     });
    /// 
    ///     var exampleUserInvitationAccepter = new Github.UserInvitationAccepter("example", new()
    ///     {
    ///         InvitationId = exampleRepositoryCollaborator.InvitationId,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Allowing empty invitation IDs
    /// 
    /// Set `allow_empty_id` when using `for_each` over a list of `github_repository_collaborator.invitation_id`'s.
    /// 
    /// This allows applying a module again when a new `github.RepositoryCollaborator` resource is added to the `for_each` loop.
    /// This is needed as the `github_repository_collaborator.invitation_id` will be empty after a state refresh when the invitation has been accepted.
    /// 
    /// Note that when an invitation is accepted manually or by another tool between a state refresh and a `pulumi up` using that refreshed state,
    /// the plan will contain the invitation ID, but the apply will receive an HTTP 404 from the API since the invitation has already been accepted.
    /// 
    /// This is tracked in #1157.
    /// </summary>
    [GithubResourceType("github:index/userInvitationAccepter:UserInvitationAccepter")]
    public partial class UserInvitationAccepter : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Allow the ID to be unset. This will result in the resource being skipped when the ID is not set instead of returning an error.
        /// </summary>
        [Output("allowEmptyId")]
        public Output<bool?> AllowEmptyId { get; private set; } = null!;

        /// <summary>
        /// ID of the invitation to accept. Must be set when `allow_empty_id` is `false`.
        /// </summary>
        [Output("invitationId")]
        public Output<string?> InvitationId { get; private set; } = null!;


        /// <summary>
        /// Create a UserInvitationAccepter resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public UserInvitationAccepter(string name, UserInvitationAccepterArgs? args = null, CustomResourceOptions? options = null)
            : base("github:index/userInvitationAccepter:UserInvitationAccepter", name, args ?? new UserInvitationAccepterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private UserInvitationAccepter(string name, Input<string> id, UserInvitationAccepterState? state = null, CustomResourceOptions? options = null)
            : base("github:index/userInvitationAccepter:UserInvitationAccepter", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing UserInvitationAccepter resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static UserInvitationAccepter Get(string name, Input<string> id, UserInvitationAccepterState? state = null, CustomResourceOptions? options = null)
        {
            return new UserInvitationAccepter(name, id, state, options);
        }
    }

    public sealed class UserInvitationAccepterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Allow the ID to be unset. This will result in the resource being skipped when the ID is not set instead of returning an error.
        /// </summary>
        [Input("allowEmptyId")]
        public Input<bool>? AllowEmptyId { get; set; }

        /// <summary>
        /// ID of the invitation to accept. Must be set when `allow_empty_id` is `false`.
        /// </summary>
        [Input("invitationId")]
        public Input<string>? InvitationId { get; set; }

        public UserInvitationAccepterArgs()
        {
        }
        public static new UserInvitationAccepterArgs Empty => new UserInvitationAccepterArgs();
    }

    public sealed class UserInvitationAccepterState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Allow the ID to be unset. This will result in the resource being skipped when the ID is not set instead of returning an error.
        /// </summary>
        [Input("allowEmptyId")]
        public Input<bool>? AllowEmptyId { get; set; }

        /// <summary>
        /// ID of the invitation to accept. Must be set when `allow_empty_id` is `false`.
        /// </summary>
        [Input("invitationId")]
        public Input<string>? InvitationId { get; set; }

        public UserInvitationAccepterState()
        {
        }
        public static new UserInvitationAccepterState Empty => new UserInvitationAccepterState();
    }
}
