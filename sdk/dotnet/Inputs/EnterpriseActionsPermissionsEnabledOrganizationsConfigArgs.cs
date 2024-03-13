// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github.Inputs
{

    public sealed class EnterpriseActionsPermissionsEnabledOrganizationsConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("organizationIds", required: true)]
        private InputList<int>? _organizationIds;

        /// <summary>
        /// List of organization IDs to enable for GitHub Actions.
        /// </summary>
        public InputList<int> OrganizationIds
        {
            get => _organizationIds ?? (_organizationIds = new InputList<int>());
            set => _organizationIds = value;
        }

        public EnterpriseActionsPermissionsEnabledOrganizationsConfigArgs()
        {
        }
        public static new EnterpriseActionsPermissionsEnabledOrganizationsConfigArgs Empty => new EnterpriseActionsPermissionsEnabledOrganizationsConfigArgs();
    }
}