// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github.Inputs
{

    public sealed class TeamMembersMemberArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The role of the user within the team.
        /// Must be one of `member` or `maintainer`. Defaults to `member`.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        /// <summary>
        /// The user to add to the team.
        /// </summary>
        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public TeamMembersMemberArgs()
        {
        }
    }
}
