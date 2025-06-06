// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github.Outputs
{

    [OutputType]
    public sealed class GetRepositoryRepositoryLicenseLicenseResult
    {
        /// <summary>
        /// The text of the license.
        /// </summary>
        public readonly string Body;
        /// <summary>
        /// Conditions associated with the license.
        /// </summary>
        public readonly ImmutableArray<string> Conditions;
        /// <summary>
        /// A description of the license.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Indicates if the license is featured.
        /// </summary>
        public readonly bool Featured;
        /// <summary>
        /// The URL to view the license details on GitHub.
        /// </summary>
        public readonly string HtmlUrl;
        /// <summary>
        /// Details about the implementation of the license.
        /// </summary>
        public readonly string Implementation;
        /// <summary>
        /// A key representing the license type (e.g., "apache-2.0").
        /// </summary>
        public readonly string Key;
        /// <summary>
        /// Limitations associated with the license.
        /// </summary>
        public readonly ImmutableArray<string> Limitations;
        /// <summary>
        /// The name of the repository.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Permissions associated with the license.
        /// </summary>
        public readonly ImmutableArray<string> Permissions;
        /// <summary>
        /// The SPDX identifier for the license (e.g., "Apache-2.0").
        /// </summary>
        public readonly string SpdxId;
        /// <summary>
        /// The URL to access information about the license on GitHub.
        /// </summary>
        public readonly string Url;

        [OutputConstructor]
        private GetRepositoryRepositoryLicenseLicenseResult(
            string body,

            ImmutableArray<string> conditions,

            string description,

            bool featured,

            string htmlUrl,

            string implementation,

            string key,

            ImmutableArray<string> limitations,

            string name,

            ImmutableArray<string> permissions,

            string spdxId,

            string url)
        {
            Body = body;
            Conditions = conditions;
            Description = description;
            Featured = featured;
            HtmlUrl = htmlUrl;
            Implementation = implementation;
            Key = key;
            Limitations = limitations;
            Name = name;
            Permissions = permissions;
            SpdxId = spdxId;
            Url = url;
        }
    }
}
