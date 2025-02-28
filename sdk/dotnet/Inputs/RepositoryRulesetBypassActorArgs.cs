// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github.Inputs
{

    public sealed class RepositoryRulesetBypassActorArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// (Number) The ID of the actor that can bypass a ruleset. If `actor_type` is `Integration`, `actor_id` is a GitHub App ID. App ID can be obtained by following instructions from the [Get an App API docs](https://docs.github.com/en/rest/apps/apps?apiVersion=2022-11-28#get-an-app)
        /// </summary>
        [Input("actorId", required: true)]
        public Input<int> ActorId { get; set; } = null!;

        /// <summary>
        /// The type of actor that can bypass a ruleset. Can be one of: `RepositoryRole`, `Team`, `Integration`, `OrganizationAdmin`.
        /// </summary>
        [Input("actorType", required: true)]
        public Input<string> ActorType { get; set; } = null!;

        /// <summary>
        /// (String) When the specified actor can bypass the ruleset. pull_request means that an actor can only bypass rules on pull requests. Can be one of: `always`, `pull_request`.
        /// 
        /// &gt; Note: at the time of writing this, the following actor types correspond to the following actor IDs:
        /// 
        /// * `OrganizationAdmin` &gt; `1`
        /// * `RepositoryRole` (This is the actor type, the following are the base repository roles and their associated IDs.)
        /// </summary>
        [Input("bypassMode", required: true)]
        public Input<string> BypassMode { get; set; } = null!;

        public RepositoryRulesetBypassActorArgs()
        {
        }
        public static new RepositoryRulesetBypassActorArgs Empty => new RepositoryRulesetBypassActorArgs();
    }
}
