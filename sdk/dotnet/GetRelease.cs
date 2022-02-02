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
        /// <summary>
        /// Use this data source to retrieve information about a GitHub release in a specific repository.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// To retrieve the latest release that is present in a repository:
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Github = Pulumi.Github;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(Github.GetRelease.InvokeAsync(new Github.GetReleaseArgs
        ///         {
        ///             Owner = "example-owner",
        ///             Repository = "example-repository",
        ///             RetrieveBy = "latest",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// 
        /// To retrieve a specific release from a repository based on it's ID:
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Github = Pulumi.Github;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(Github.GetRelease.InvokeAsync(new Github.GetReleaseArgs
        ///         {
        ///             Id = 12345,
        ///             Owner = "example-owner",
        ///             Repository = "example-repository",
        ///             RetrieveBy = "id",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// 
        /// Finally, to retrieve a release based on it's tag:
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Github = Pulumi.Github;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(Github.GetRelease.InvokeAsync(new Github.GetReleaseArgs
        ///         {
        ///             Owner = "example-owner",
        ///             ReleaseTag = "v1.0.0",
        ///             Repository = "example-repository",
        ///             RetrieveBy = "tag",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetReleaseResult> InvokeAsync(GetReleaseArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetReleaseResult>("github:index/getRelease:getRelease", args ?? new GetReleaseArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve information about a GitHub release in a specific repository.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// To retrieve the latest release that is present in a repository:
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Github = Pulumi.Github;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(Github.GetRelease.InvokeAsync(new Github.GetReleaseArgs
        ///         {
        ///             Owner = "example-owner",
        ///             Repository = "example-repository",
        ///             RetrieveBy = "latest",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// 
        /// To retrieve a specific release from a repository based on it's ID:
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Github = Pulumi.Github;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(Github.GetRelease.InvokeAsync(new Github.GetReleaseArgs
        ///         {
        ///             Id = 12345,
        ///             Owner = "example-owner",
        ///             Repository = "example-repository",
        ///             RetrieveBy = "id",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// 
        /// Finally, to retrieve a release based on it's tag:
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Github = Pulumi.Github;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(Github.GetRelease.InvokeAsync(new Github.GetReleaseArgs
        ///         {
        ///             Owner = "example-owner",
        ///             ReleaseTag = "v1.0.0",
        ///             Repository = "example-repository",
        ///             RetrieveBy = "tag",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetReleaseResult> Invoke(GetReleaseInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetReleaseResult>("github:index/getRelease:getRelease", args ?? new GetReleaseInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetReleaseArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Owner of the repository.
        /// </summary>
        [Input("owner", required: true)]
        public string Owner { get; set; } = null!;

        /// <summary>
        /// ID of the release to retrieve. Must be specified when `retrieve_by` = `id`.
        /// </summary>
        [Input("releaseId")]
        public int? ReleaseId { get; set; }

        /// <summary>
        /// Tag of the release to retrieve. Must be specified when `retrieve_by` = `tag`.
        /// </summary>
        [Input("releaseTag")]
        public string? ReleaseTag { get; set; }

        /// <summary>
        /// Name of the repository to retrieve the release from.
        /// </summary>
        [Input("repository", required: true)]
        public string Repository { get; set; } = null!;

        /// <summary>
        /// Describes how to fetch the release. Valid values are `id`, `tag`, `latest`.
        /// </summary>
        [Input("retrieveBy", required: true)]
        public string RetrieveBy { get; set; } = null!;

        public GetReleaseArgs()
        {
        }
    }

    public sealed class GetReleaseInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Owner of the repository.
        /// </summary>
        [Input("owner", required: true)]
        public Input<string> Owner { get; set; } = null!;

        /// <summary>
        /// ID of the release to retrieve. Must be specified when `retrieve_by` = `id`.
        /// </summary>
        [Input("releaseId")]
        public Input<int>? ReleaseId { get; set; }

        /// <summary>
        /// Tag of the release to retrieve. Must be specified when `retrieve_by` = `tag`.
        /// </summary>
        [Input("releaseTag")]
        public Input<string>? ReleaseTag { get; set; }

        /// <summary>
        /// Name of the repository to retrieve the release from.
        /// </summary>
        [Input("repository", required: true)]
        public Input<string> Repository { get; set; } = null!;

        /// <summary>
        /// Describes how to fetch the release. Valid values are `id`, `tag`, `latest`.
        /// </summary>
        [Input("retrieveBy", required: true)]
        public Input<string> RetrieveBy { get; set; } = null!;

        public GetReleaseInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetReleaseResult
    {
        /// <summary>
        /// URL of any associated assets with the release
        /// </summary>
        public readonly string AssertsUrl;
        /// <summary>
        /// Contents of the description (body) of a release
        /// </summary>
        public readonly string Body;
        /// <summary>
        /// Date of release creation
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// (`Boolean`) indicates whether the release is a draft
        /// </summary>
        public readonly bool Draft;
        /// <summary>
        /// URL directing to detailed information on the release
        /// </summary>
        public readonly string HtmlUrl;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Name of release
        /// </summary>
        public readonly string Name;
        public readonly string Owner;
        /// <summary>
        /// (`Boolean`) indicates whether the release is a prerelease
        /// </summary>
        public readonly bool Prerelease;
        /// <summary>
        /// Date of release publishing
        /// </summary>
        public readonly string PublishedAt;
        /// <summary>
        /// ID of release
        /// </summary>
        public readonly int? ReleaseId;
        /// <summary>
        /// Tag of release
        /// </summary>
        public readonly string? ReleaseTag;
        public readonly string Repository;
        public readonly string RetrieveBy;
        /// <summary>
        /// Download URL of a specific release in `tar.gz` format
        /// </summary>
        public readonly string TarballUrl;
        /// <summary>
        /// Commitish value that determines where the Git release is created from
        /// </summary>
        public readonly string TargetCommitish;
        /// <summary>
        /// URL that can be used to upload Assets to the release
        /// </summary>
        public readonly string UploadUrl;
        /// <summary>
        /// Base URL of the release
        /// </summary>
        public readonly string Url;
        /// <summary>
        /// Download URL of a specific release in `zip` format
        /// </summary>
        public readonly string ZipballUrl;

        [OutputConstructor]
        private GetReleaseResult(
            string assertsUrl,

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
