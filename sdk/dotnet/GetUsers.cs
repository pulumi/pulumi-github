// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github
{
    public static class GetUsers
    {
        /// <summary>
        /// Use this data source to retrieve information about multiple GitHub users at once.
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
        ///     var example = Github.GetUsers.Invoke(new()
        ///     {
        ///         Usernames = new[]
        ///         {
        ///             "example1",
        ///             "example2",
        ///             "example3",
        ///         },
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["validUsers"] = example.Apply(getUsersResult =&gt; getUsersResult.Logins),
        ///         ["invalidUsers"] = example.Apply(getUsersResult =&gt; getUsersResult.UnknownLogins),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetUsersResult> InvokeAsync(GetUsersArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetUsersResult>("github:index/getUsers:getUsers", args ?? new GetUsersArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve information about multiple GitHub users at once.
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
        ///     var example = Github.GetUsers.Invoke(new()
        ///     {
        ///         Usernames = new[]
        ///         {
        ///             "example1",
        ///             "example2",
        ///             "example3",
        ///         },
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["validUsers"] = example.Apply(getUsersResult =&gt; getUsersResult.Logins),
        ///         ["invalidUsers"] = example.Apply(getUsersResult =&gt; getUsersResult.UnknownLogins),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetUsersResult> Invoke(GetUsersInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetUsersResult>("github:index/getUsers:getUsers", args ?? new GetUsersInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetUsersArgs : global::Pulumi.InvokeArgs
    {
        [Input("usernames", required: true)]
        private List<string>? _usernames;

        /// <summary>
        /// List of usernames.
        /// </summary>
        public List<string> Usernames
        {
            get => _usernames ?? (_usernames = new List<string>());
            set => _usernames = value;
        }

        public GetUsersArgs()
        {
        }
        public static new GetUsersArgs Empty => new GetUsersArgs();
    }

    public sealed class GetUsersInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("usernames", required: true)]
        private InputList<string>? _usernames;

        /// <summary>
        /// List of usernames.
        /// </summary>
        public InputList<string> Usernames
        {
            get => _usernames ?? (_usernames = new InputList<string>());
            set => _usernames = value;
        }

        public GetUsersInvokeArgs()
        {
        }
        public static new GetUsersInvokeArgs Empty => new GetUsersInvokeArgs();
    }


    [OutputType]
    public sealed class GetUsersResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// list of logins of users that could be found.
        /// </summary>
        public readonly ImmutableArray<string> Logins;
        /// <summary>
        /// list of Node IDs of users that could be found.
        /// </summary>
        public readonly ImmutableArray<string> NodeIds;
        /// <summary>
        /// list of logins without matching user.
        /// </summary>
        public readonly ImmutableArray<string> UnknownLogins;
        public readonly ImmutableArray<string> Usernames;

        [OutputConstructor]
        private GetUsersResult(
            string id,

            ImmutableArray<string> logins,

            ImmutableArray<string> nodeIds,

            ImmutableArray<string> unknownLogins,

            ImmutableArray<string> usernames)
        {
            Id = id;
            Logins = logins;
            NodeIds = nodeIds;
            UnknownLogins = unknownLogins;
            Usernames = usernames;
        }
    }
}
