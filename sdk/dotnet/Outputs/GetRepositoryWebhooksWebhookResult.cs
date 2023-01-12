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
    public sealed class GetRepositoryWebhooksWebhookResult
    {
        /// <summary>
        /// `true` if the webhook is active.
        /// </summary>
        public readonly bool Active;
        /// <summary>
        /// the ID of the webhook.
        /// </summary>
        public readonly int Id;
        /// <summary>
        /// the name of the webhook.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// the type of the webhook.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// the url of the webhook.
        /// </summary>
        public readonly string Url;

        [OutputConstructor]
        private GetRepositoryWebhooksWebhookResult(
            bool active,

            int id,

            string name,

            string type,

            string url)
        {
            Active = active;
            Id = id;
            Name = name;
            Type = type;
            Url = url;
        }
    }
}
