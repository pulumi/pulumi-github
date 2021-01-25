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
    public sealed class GetRepositoryPageSourceResult
    {
        public readonly string Branch;
        public readonly string Path;

        [OutputConstructor]
        private GetRepositoryPageSourceResult(
            string branch,

            string path)
        {
            Branch = branch;
            Path = path;
        }
    }
}