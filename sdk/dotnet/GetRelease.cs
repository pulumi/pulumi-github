// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github
{
    public static class GetRelease
    {
        public static Task<GetReleaseResult> InvokeAsync(GetReleaseArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetReleaseResult>("github:index/getRelease:getRelease", args ?? new GetReleaseArgs(), options.WithDefaults());

        public static Output<GetReleaseResult> Invoke(GetReleaseInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetReleaseResult>("github:index/getRelease:getRelease", args ?? new GetReleaseInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetReleaseArgs : global::Pulumi.InvokeArgs
    {
        [Input("owner", required: true)]
        public string Owner { get; set; } = null!;

        [Input("releaseId")]
        public int? ReleaseId { get; set; }

        [Input("releaseTag")]
        public string? ReleaseTag { get; set; }

        [Input("repository", required: true)]
        public string Repository { get; set; } = null!;

        [Input("retrieveBy", required: true)]
        public string RetrieveBy { get; set; } = null!;

        public GetReleaseArgs()
        {
        }
        public static new GetReleaseArgs Empty => new GetReleaseArgs();
    }

    public sealed class GetReleaseInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("owner", required: true)]
        public Input<string> Owner { get; set; } = null!;

        [Input("releaseId")]
        public Input<int>? ReleaseId { get; set; }

        [Input("releaseTag")]
        public Input<string>? ReleaseTag { get; set; }

        [Input("repository", required: true)]
        public Input<string> Repository { get; set; } = null!;

        [Input("retrieveBy", required: true)]
        public Input<string> RetrieveBy { get; set; } = null!;

        public GetReleaseInvokeArgs()
        {
        }
        public static new GetReleaseInvokeArgs Empty => new GetReleaseInvokeArgs();
    }


    [OutputType]
    public sealed class GetReleaseResult
    {
        public readonly string AssertsUrl;
        public readonly ImmutableArray<Outputs.GetReleaseAssetResult> Assets;
        public readonly string AssetsUrl;
        public readonly string Body;
        public readonly string CreatedAt;
        public readonly bool Draft;
        public readonly string HtmlUrl;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string Owner;
        public readonly bool Prerelease;
        public readonly string PublishedAt;
        public readonly int? ReleaseId;
        public readonly string? ReleaseTag;
        public readonly string Repository;
        public readonly string RetrieveBy;
        public readonly string TarballUrl;
        public readonly string TargetCommitish;
        public readonly string UploadUrl;
        public readonly string Url;
        public readonly string ZipballUrl;

        [OutputConstructor]
        private GetReleaseResult(
            string assertsUrl,

            ImmutableArray<Outputs.GetReleaseAssetResult> assets,

            string assetsUrl,

            string body,

            string createdAt,

            bool draft,

            string htmlUrl,

            string id,

            string name,

            string owner,

            bool prerelease,

            string publishedAt,

            int? releaseId,

            string? releaseTag,

            string repository,

            string retrieveBy,

            string tarballUrl,

            string targetCommitish,

            string uploadUrl,

            string url,

            string zipballUrl)
        {
            AssertsUrl = assertsUrl;
            Assets = assets;
            AssetsUrl = assetsUrl;
            Body = body;
            CreatedAt = createdAt;
            Draft = draft;
            HtmlUrl = htmlUrl;
            Id = id;
            Name = name;
            Owner = owner;
            Prerelease = prerelease;
            PublishedAt = publishedAt;
            ReleaseId = releaseId;
            ReleaseTag = releaseTag;
            Repository = repository;
            RetrieveBy = retrieveBy;
            TarballUrl = tarballUrl;
            TargetCommitish = targetCommitish;
            UploadUrl = uploadUrl;
            Url = url;
            ZipballUrl = zipballUrl;
        }
    }
}
