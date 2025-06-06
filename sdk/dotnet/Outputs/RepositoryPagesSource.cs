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
    public sealed class RepositoryPagesSource
    {
        /// <summary>
        /// The repository branch used to publish the site's source files. (i.e. `main` or `gh-pages`.
        /// </summary>
        public readonly string Branch;
        /// <summary>
        /// The repository directory from which the site publishes (Default: `/`).
        /// </summary>
        public readonly string? Path;

        [OutputConstructor]
        private RepositoryPagesSource(
            string branch,

            string? path)
        {
            Branch = branch;
            Path = path;
        }
    }
}
