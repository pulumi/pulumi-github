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
    public sealed class RepositorySecurityAndAnalysis
    {
        /// <summary>
        /// The advanced security configuration for the repository. See Advanced Security Configuration below for details. If a repository's visibility is `public`, advanced security is always enabled and cannot be changed, so this setting cannot be supplied.
        /// </summary>
        public readonly Outputs.RepositorySecurityAndAnalysisAdvancedSecurity? AdvancedSecurity;
        /// <summary>
        /// The secret scanning configuration for the repository. See Secret Scanning Configuration below for details.
        /// </summary>
        public readonly Outputs.RepositorySecurityAndAnalysisSecretScanning? SecretScanning;
        /// <summary>
        /// The secret scanning push protection configuration for the repository. See Secret Scanning Push Protection Configuration below for details.
        /// </summary>
        public readonly Outputs.RepositorySecurityAndAnalysisSecretScanningPushProtection? SecretScanningPushProtection;

        [OutputConstructor]
        private RepositorySecurityAndAnalysis(
            Outputs.RepositorySecurityAndAnalysisAdvancedSecurity? advancedSecurity,

            Outputs.RepositorySecurityAndAnalysisSecretScanning? secretScanning,

            Outputs.RepositorySecurityAndAnalysisSecretScanningPushProtection? secretScanningPushProtection)
        {
            AdvancedSecurity = advancedSecurity;
            SecretScanning = secretScanning;
            SecretScanningPushProtection = secretScanningPushProtection;
        }
    }
}
