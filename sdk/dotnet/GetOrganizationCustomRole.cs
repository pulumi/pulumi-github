// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github
{
    public static class GetOrganizationCustomRole
    {
        /// <summary>
        /// Use this data source to retrieve information about a custom role in a GitHub Organization.
        /// 
        /// &gt; Note: Custom roles are currently only available in GitHub Enterprise Cloud.
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
        ///     var example = Github.GetOrganizationCustomRole.Invoke(new()
        ///     {
        ///         Name = "example",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetOrganizationCustomRoleResult> InvokeAsync(GetOrganizationCustomRoleArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetOrganizationCustomRoleResult>("github:index/getOrganizationCustomRole:getOrganizationCustomRole", args ?? new GetOrganizationCustomRoleArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve information about a custom role in a GitHub Organization.
        /// 
        /// &gt; Note: Custom roles are currently only available in GitHub Enterprise Cloud.
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
        ///     var example = Github.GetOrganizationCustomRole.Invoke(new()
        ///     {
        ///         Name = "example",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetOrganizationCustomRoleResult> Invoke(GetOrganizationCustomRoleInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetOrganizationCustomRoleResult>("github:index/getOrganizationCustomRole:getOrganizationCustomRole", args ?? new GetOrganizationCustomRoleInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetOrganizationCustomRoleArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the custom role.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetOrganizationCustomRoleArgs()
        {
        }
        public static new GetOrganizationCustomRoleArgs Empty => new GetOrganizationCustomRoleArgs();
    }

    public sealed class GetOrganizationCustomRoleInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the custom role.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetOrganizationCustomRoleInvokeArgs()
        {
        }
        public static new GetOrganizationCustomRoleInvokeArgs Empty => new GetOrganizationCustomRoleInvokeArgs();
    }


    [OutputType]
    public sealed class GetOrganizationCustomRoleResult
    {
        /// <summary>
        /// The system role from which the role inherits permissions.
        /// </summary>
        public readonly string BaseRole;
        /// <summary>
        /// The description for the custom role.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        /// <summary>
        /// A list of additional permissions included in this role.
        /// </summary>
        public readonly ImmutableArray<string> Permissions;

        [OutputConstructor]
        private GetOrganizationCustomRoleResult(
            string baseRole,

            string description,

            string id,

            string name,

            ImmutableArray<string> permissions)
        {
            BaseRole = baseRole;
            Description = description;
            Id = id;
            Name = name;
            Permissions = permissions;
        }
    }
}
