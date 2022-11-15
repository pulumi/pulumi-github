// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github
{
    public static class GetBranch
    {
        public static Task<GetBranchResult> InvokeAsync(GetBranchArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetBranchResult>("github:index/getBranch:getBranch", args ?? new GetBranchArgs(), options.WithDefaults());

        public static Output<GetBranchResult> Invoke(GetBranchInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetBranchResult>("github:index/getBranch:getBranch", args ?? new GetBranchInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetBranchArgs : global::Pulumi.InvokeArgs
    {
        [Input("branch", required: true)]
        public string Branch { get; set; } = null!;

        [Input("repository", required: true)]
        public string Repository { get; set; } = null!;

        public GetBranchArgs()
        {
        }
        public static new GetBranchArgs Empty => new GetBranchArgs();
    }

    public sealed class GetBranchInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("branch", required: true)]
        public Input<string> Branch { get; set; } = null!;

        [Input("repository", required: true)]
        public Input<string> Repository { get; set; } = null!;

        public GetBranchInvokeArgs()
        {
        }
        public static new GetBranchInvokeArgs Empty => new GetBranchInvokeArgs();
    }


    [OutputType]
    public sealed class GetBranchResult
    {
        public readonly string Branch;
        public readonly string Etag;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Ref;
        public readonly string Repository;
        public readonly string Sha;

        [OutputConstructor]
        private GetBranchResult(
            string branch,

            string etag,

            string id,

            string @ref,

            string repository,

            string sha)
        {
            Branch = branch;
            Etag = etag;
            Id = id;
            Ref = @ref;
            Repository = repository;
            Sha = sha;
        }
    }
}
