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
        public static Task<GetMembershipResult> InvokeAsync(GetMembershipArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetMembershipResult>("github:index/getMembership:getMembership", args ?? new GetMembershipArgs(), options.WithDefaults());

        public static Output<GetMembershipResult> Invoke(GetMembershipInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetMembershipResult>("github:index/getMembership:getMembership", args ?? new GetMembershipInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetMembershipArgs : global::Pulumi.InvokeArgs
    {
        [Input("organization")]
        public string? Organization { get; set; }

        [Input("username", required: true)]
        public string Username { get; set; } = null!;

        public GetMembershipArgs()
        {
        }
        public static new GetMembershipArgs Empty => new GetMembershipArgs();
    }

    public sealed class GetMembershipInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("organization")]
        public Input<string>? Organization { get; set; }

        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public GetMembershipInvokeArgs()
        {
        }
        public static new GetMembershipInvokeArgs Empty => new GetMembershipInvokeArgs();
    }


    [OutputType]
    public sealed class GetMembershipResult
    {
        public readonly string Etag;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? Organization;
        public readonly string Role;
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
