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
    public sealed class RepositoryCollaboratorsUser
    {
        /// <summary>
        /// The permission of the outside collaborators for the repository.
        /// Must be one of `pull`, `push`, `maintain`, `triage` or `admin` or the name of an existing [custom repository role](https://docs.github.com/en/enterprise-cloud@latest/organizations/managing-peoples-access-to-your-organization-with-roles/managing-custom-repository-roles-for-an-organization) within the organization for organization-owned repositories.
        /// Must be `push` for personal repositories. Defaults to `push`.
        /// </summary>
        public readonly string? Permission;
        public readonly string Username;

        [OutputConstructor]
        private RepositoryCollaboratorsUser(
            string? permission,

            string username)
        {
            Permission = permission;
            Username = username;
        }
    }
}