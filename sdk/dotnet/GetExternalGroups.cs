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
        public static Task<GetExternalGroupsResult> InvokeAsync(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetExternalGroupsResult>("github:index/getExternalGroups:getExternalGroups", InvokeArgs.Empty, options.WithDefaults());
    }


    [OutputType]
    public sealed class GetExternalGroupsResult
    {
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
