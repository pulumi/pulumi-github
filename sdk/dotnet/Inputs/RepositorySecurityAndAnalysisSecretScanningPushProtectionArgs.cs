// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github.Inputs
{

    public sealed class RepositorySecurityAndAnalysisSecretScanningPushProtectionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The GitHub Pages site's build status e.g. `building` or `built`.
        /// </summary>
        [Input("status", required: true)]
        public Input<string> Status { get; set; } = null!;

        public RepositorySecurityAndAnalysisSecretScanningPushProtectionArgs()
        {
        }
        public static new RepositorySecurityAndAnalysisSecretScanningPushProtectionArgs Empty => new RepositorySecurityAndAnalysisSecretScanningPushProtectionArgs();
    }
}