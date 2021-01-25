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
    public sealed class RepositoryPages
    {
        /// <summary>
        /// The custom domain for the repository. This can only be set after the repository has been created.
        /// </summary>
        public readonly string? Cname;
        /// <summary>
        /// Whether the rendered Github Pages site has a custom 404 page.
        /// </summary>
        public readonly bool? Custom404;
        /// <summary>
        /// The absolute URL (including scheme) of the rendered Github Pages site e.g. `https://username.github.io`.
        /// </summary>
        public readonly string? HtmlUrl;
        /// <summary>
        /// The source branch and directory for the rendered Pages site. See Github Pages Source below for details.
        /// </summary>
        public readonly Outputs.RepositoryPagesSource Source;
        /// <summary>
        /// The Github Pages site's build status e.g. `building` or `built`.
        /// </summary>
        public readonly string? Status;
        public readonly string? Url;

        [OutputConstructor]
        private RepositoryPages(
            string? cname,

            bool? custom404,

            string? htmlUrl,

            Outputs.RepositoryPagesSource source,

            string? status,

            string? url)
        {
            Cname = cname;
            Custom404 = custom404;
            HtmlUrl = htmlUrl;
            Source = source;
            Status = status;
            Url = url;
        }
    }
}