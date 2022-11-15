// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github
{
    public static class GetRepositoryBranches
    {
        public static Task<GetRepositoryBranchesResult> InvokeAsync(GetRepositoryBranchesArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRepositoryBranchesResult>("github:index/getRepositoryBranches:getRepositoryBranches", args ?? new GetRepositoryBranchesArgs(), options.WithDefaults());

        public static Output<GetRepositoryBranchesResult> Invoke(GetRepositoryBranchesInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetRepositoryBranchesResult>("github:index/getRepositoryBranches:getRepositoryBranches", args ?? new GetRepositoryBranchesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRepositoryBranchesArgs : global::Pulumi.InvokeArgs
    {
        [Input("repository", required: true)]
        public string Repository { get; set; } = null!;

        public GetRepositoryBranchesArgs()
        {
        }
        public static new GetRepositoryBranchesArgs Empty => new GetRepositoryBranchesArgs();
    }

    public sealed class GetRepositoryBranchesInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("repository", required: true)]
        public Input<string> Repository { get; set; } = null!;

        public GetRepositoryBranchesInvokeArgs()
        {
        }
        public static new GetRepositoryBranchesInvokeArgs Empty => new GetRepositoryBranchesInvokeArgs();
    }


    [OutputType]
    public sealed class GetRepositoryBranchesResult
    {
        public readonly ImmutableArray<Outputs.GetRepositoryBranchesBranchResult> Branches;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Repository;

        [OutputConstructor]
        private GetRepositoryBranchesResult(
            ImmutableArray<Outputs.GetRepositoryBranchesBranchResult> branches,

            string id,

            string repository)
        {
            Branches = branches;
            Id = id;
            Repository = repository;
        }
    }
}
