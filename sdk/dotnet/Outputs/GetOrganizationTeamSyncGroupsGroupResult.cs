// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github.Outputs
{

    [OutputType]
    public sealed class GetOrganizationTeamSyncGroupsGroupResult
    {
        /// <summary>
        /// The description of the IdP group.
        /// </summary>
        public readonly string GroupDescription;
        /// <summary>
        /// The ID of the IdP group.
        /// </summary>
        public readonly string GroupId;
        /// <summary>
        /// The name of the IdP group.
        /// </summary>
        public readonly string GroupName;

        [OutputConstructor]
        private GetOrganizationTeamSyncGroupsGroupResult(
            string groupDescription,

            string groupId,

            string groupName)
        {
            GroupDescription = groupDescription;
            GroupId = groupId;
            GroupName = groupName;
        }
    }
}
