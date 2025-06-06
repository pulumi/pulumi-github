// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github.Inputs
{

    public sealed class ProviderAppAuthArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The GitHub App ID.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        /// <summary>
        /// The GitHub App installation instance ID.
        /// </summary>
        [Input("installationId", required: true)]
        public Input<string> InstallationId { get; set; } = null!;

        [Input("pemFile", required: true)]
        private Input<string>? _pemFile;

        /// <summary>
        /// The GitHub App PEM file contents.
        /// </summary>
        public Input<string>? PemFile
        {
            get => _pemFile;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _pemFile = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        public ProviderAppAuthArgs()
        {
        }
        public static new ProviderAppAuthArgs Empty => new ProviderAppAuthArgs();
    }
}
