// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github
{
    public static class GetUser
    {
        /// <summary>
        /// Use this data source to retrieve information about a GitHub user.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Github = Pulumi.Github;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Github.GetUser.Invoke(new()
        ///     {
        ///         Username = "example",
        ///     });
        /// 
        ///     var current = Github.GetUser.Invoke(new()
        ///     {
        ///         Username = "",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["currentGithubLogin"] = current.Apply(getUserResult =&gt; getUserResult.Login),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetUserResult> InvokeAsync(GetUserArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetUserResult>("github:index/getUser:getUser", args ?? new GetUserArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve information about a GitHub user.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Github = Pulumi.Github;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Github.GetUser.Invoke(new()
        ///     {
        ///         Username = "example",
        ///     });
        /// 
        ///     var current = Github.GetUser.Invoke(new()
        ///     {
        ///         Username = "",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["currentGithubLogin"] = current.Apply(getUserResult =&gt; getUserResult.Login),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetUserResult> Invoke(GetUserInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetUserResult>("github:index/getUser:getUser", args ?? new GetUserInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetUserArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The username. Use an empty string `""` to retrieve information about the currently authenticated user.
        /// </summary>
        [Input("username", required: true)]
        public string Username { get; set; } = null!;

        public GetUserArgs()
        {
        }
        public static new GetUserArgs Empty => new GetUserArgs();
    }

    public sealed class GetUserInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The username. Use an empty string `""` to retrieve information about the currently authenticated user.
        /// </summary>
        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public GetUserInvokeArgs()
        {
        }
        public static new GetUserInvokeArgs Empty => new GetUserInvokeArgs();
    }


    [OutputType]
    public sealed class GetUserResult
    {
        /// <summary>
        /// the user's avatar URL.
        /// </summary>
        public readonly string AvatarUrl;
        /// <summary>
        /// the user's bio.
        /// </summary>
        public readonly string Bio;
        /// <summary>
        /// the user's blog location.
        /// </summary>
        public readonly string Blog;
        /// <summary>
        /// the user's company name.
        /// </summary>
        public readonly string Company;
        /// <summary>
        /// the creation date.
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// the user's email.
        /// </summary>
        public readonly string Email;
        /// <summary>
        /// the number of followers.
        /// </summary>
        public readonly int Followers;
        /// <summary>
        /// the number of following users.
        /// </summary>
        public readonly int Following;
        /// <summary>
        /// list of user's GPG keys.
        /// </summary>
        public readonly ImmutableArray<string> GpgKeys;
        /// <summary>
        /// the user's gravatar ID.
        /// </summary>
        public readonly string GravatarId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// the user's location.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// the user's login.
        /// </summary>
        public readonly string Login;
        /// <summary>
        /// the user's full name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// the Node ID of the user.
        /// </summary>
        public readonly string NodeId;
        /// <summary>
        /// the number of public gists.
        /// </summary>
        public readonly int PublicGists;
        /// <summary>
        /// the number of public repositories.
        /// </summary>
        public readonly int PublicRepos;
        /// <summary>
        /// whether the user is a GitHub admin.
        /// </summary>
        public readonly bool SiteAdmin;
        /// <summary>
        /// list of user's SSH keys.
        /// </summary>
        public readonly ImmutableArray<string> SshKeys;
        /// <summary>
        /// the suspended date if the user is suspended.
        /// </summary>
        public readonly string SuspendedAt;
        /// <summary>
        /// the update date.
        /// </summary>
        public readonly string UpdatedAt;
        public readonly string Username;

        [OutputConstructor]
        private GetUserResult(
            string avatarUrl,

            string bio,

            string blog,

            string company,

            string createdAt,

            string email,

            int followers,

            int following,

            ImmutableArray<string> gpgKeys,

            string gravatarId,

            string id,

            string location,

            string login,

            string name,

            string nodeId,

            int publicGists,

            int publicRepos,

            bool siteAdmin,

            ImmutableArray<string> sshKeys,

            string suspendedAt,

            string updatedAt,

            string username)
        {
            AvatarUrl = avatarUrl;
            Bio = bio;
            Blog = blog;
            Company = company;
            CreatedAt = createdAt;
            Email = email;
            Followers = followers;
            Following = following;
            GpgKeys = gpgKeys;
            GravatarId = gravatarId;
            Id = id;
            Location = location;
            Login = login;
            Name = name;
            NodeId = nodeId;
            PublicGists = publicGists;
            PublicRepos = publicRepos;
            SiteAdmin = siteAdmin;
            SshKeys = sshKeys;
            SuspendedAt = suspendedAt;
            UpdatedAt = updatedAt;
            Username = username;
        }
    }
}
