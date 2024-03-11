// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github
{
    public static class GetUserExternalIdentity
    {
        /// <summary>
        /// Use this data source to retrieve a specific organization member's SAML or SCIM user
        /// attributes.
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
        ///     var exampleUser = Github.GetUserExternalIdentity.Invoke(new()
        ///     {
        ///         Username = "example-user",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetUserExternalIdentityResult> InvokeAsync(GetUserExternalIdentityArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetUserExternalIdentityResult>("github:index/getUserExternalIdentity:getUserExternalIdentity", args ?? new GetUserExternalIdentityArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve a specific organization member's SAML or SCIM user
        /// attributes.
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
        ///     var exampleUser = Github.GetUserExternalIdentity.Invoke(new()
        ///     {
        ///         Username = "example-user",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetUserExternalIdentityResult> Invoke(GetUserExternalIdentityInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetUserExternalIdentityResult>("github:index/getUserExternalIdentity:getUserExternalIdentity", args ?? new GetUserExternalIdentityInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetUserExternalIdentityArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The username of the member to fetch external identity for.
        /// </summary>
        [Input("username", required: true)]
        public string Username { get; set; } = null!;

        public GetUserExternalIdentityArgs()
        {
        }
        public static new GetUserExternalIdentityArgs Empty => new GetUserExternalIdentityArgs();
    }

    public sealed class GetUserExternalIdentityInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The username of the member to fetch external identity for.
        /// </summary>
        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public GetUserExternalIdentityInvokeArgs()
        {
        }
        public static new GetUserExternalIdentityInvokeArgs Empty => new GetUserExternalIdentityInvokeArgs();
    }


    [OutputType]
    public sealed class GetUserExternalIdentityResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The username of the GitHub user
        /// </summary>
        public readonly string Login;
        /// <summary>
        /// An Object containing the user's SAML data. This object will
        /// be empty if the user is not managed by SAML.
        /// </summary>
        public readonly ImmutableDictionary<string, string> SamlIdentity;
        /// <summary>
        /// An Object contining the user's SCIM data. This object will
        /// be empty if the user is not managed by SCIM.
        /// </summary>
        public readonly ImmutableDictionary<string, string> ScimIdentity;
        /// <summary>
        /// The member's SAML Username
        /// </summary>
        public readonly string Username;

        [OutputConstructor]
        private GetUserExternalIdentityResult(
            string id,

            string login,

            ImmutableDictionary<string, string> samlIdentity,

            ImmutableDictionary<string, string> scimIdentity,

            string username)
        {
            Id = id;
            Login = login;
            SamlIdentity = samlIdentity;
            ScimIdentity = scimIdentity;
            Username = username;
        }
    }
}
