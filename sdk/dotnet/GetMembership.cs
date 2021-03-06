// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github
{
    public static class GetMembership
    {
        /// <summary>
        /// Use this data source to find out if a user is a member of your organization, as well
        /// as what role they have within it.
        /// If the user's membership in the organization is pending their acceptance of an invite,
        /// the role they would have once they accept will be returned.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Github = Pulumi.Github;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var membershipForSomeUser = Output.Create(Github.GetMembership.InvokeAsync(new Github.GetMembershipArgs
        ///         {
        ///             Username = "SomeUser",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetMembershipResult> InvokeAsync(GetMembershipArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetMembershipResult>("github:index/getMembership:getMembership", args ?? new GetMembershipArgs(), options.WithVersion());
    }


    public sealed class GetMembershipArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The organization to check for the above username.
        /// </summary>
        [Input("organization")]
        public string? Organization { get; set; }

        /// <summary>
        /// The username to lookup in the organization.
        /// </summary>
        [Input("username", required: true)]
        public string Username { get; set; } = null!;

        public GetMembershipArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetMembershipResult
    {
        /// <summary>
        /// An etag representing the membership object.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? Organization;
        /// <summary>
        /// `admin` or `member` -- the role the user has within the organization.
        /// </summary>
        public readonly string Role;
        /// <summary>
        /// The username.
        /// </summary>
        public readonly string Username;

        [OutputConstructor]
        private GetMembershipResult(
            string etag,

            string id,

            string? organization,

            string role,

            string username)
        {
            Etag = etag;
            Id = id;
            Organization = organization;
            Role = role;
            Username = username;
        }
    }
}
