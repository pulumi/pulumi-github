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
    public sealed class RepositoryWebhookConfiguration
    {
        public readonly string? ContentType;
        public readonly bool? InsecureSsl;
        public readonly string? Secret;
        public readonly string Url;

        [OutputConstructor]
        private RepositoryWebhookConfiguration(
            string? contentType,

            bool? insecureSsl,

            string? secret,

            string url)
        {
            ContentType = contentType;
            InsecureSsl = insecureSsl;
            Secret = secret;
            Url = url;
        }
    }
}
