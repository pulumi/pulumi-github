// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github
{
    public static class GetExternalGroups
    {
        /// <summary>
        /// Use this data source to retrieve external groups belonging to an organization.
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
        ///     var exampleExternalGroups = Github.GetExternalGroups.Invoke();
        /// 
        ///     var localGroups = exampleExternalGroups.Apply(getExternalGroupsResult =&gt; getExternalGroupsResult);
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["groups"] = localGroups,
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetExternalGroupsResult> InvokeAsync(InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetExternalGroupsResult>("github:index/getExternalGroups:getExternalGroups", InvokeArgs.Empty, options.WithDefaults());
    }


    [OutputType]
    public sealed class GetExternalGroupsResult
    {
        /// <summary>
        /// an array of external groups belonging to the organization. Each group consists of the fields documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetExternalGroupsExternalGroupResult> ExternalGroups;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetExternalGroupsResult(
            ImmutableArray<Outputs.GetExternalGroupsExternalGroupResult> externalGroups,

            string id)
        {
            ExternalGroups = externalGroups;
            Id = id;
        }
    }
}
