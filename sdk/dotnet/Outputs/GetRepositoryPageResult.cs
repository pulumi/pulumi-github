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
    public sealed class GetRepositoryPageResult
    {
        public readonly string BuildType;
        public readonly string Cname;
        public readonly bool Custom404;
        /// <summary>
        /// URL to the repository on the web.
        /// </summary>
        public readonly string HtmlUrl;
        public readonly ImmutableArray<Outputs.GetRepositoryPageSourceResult> Sources;
        public readonly string Status;
        public readonly string Url;

        [OutputConstructor]
        private GetRepositoryPageResult(
            string buildType,

            string cname,

            bool custom404,

            string htmlUrl,

            ImmutableArray<Outputs.GetRepositoryPageSourceResult> sources,

            string status,

            string url)
        {
            BuildType = buildType;
            Cname = cname;
            Custom404 = custom404;
            HtmlUrl = htmlUrl;
            Sources = sources;
            Status = status;
            Url = url;
        }
    }
}
