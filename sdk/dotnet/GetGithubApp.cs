// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github
{
    public static class GetGithubApp
    {
        /// <summary>
        /// Use this data source to retrieve information about an app.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Github = Pulumi.Github;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var foobar = Github.GetGithubApp.Invoke(new()
        ///     {
        ///         Slug = "foobar",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetGithubAppResult> InvokeAsync(GetGithubAppArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetGithubAppResult>("github:index/getGithubApp:getGithubApp", args ?? new GetGithubAppArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve information about an app.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Github = Pulumi.Github;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var foobar = Github.GetGithubApp.Invoke(new()
        ///     {
        ///         Slug = "foobar",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetGithubAppResult> Invoke(GetGithubAppInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetGithubAppResult>("github:index/getGithubApp:getGithubApp", args ?? new GetGithubAppInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve information about an app.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Github = Pulumi.Github;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var foobar = Github.GetGithubApp.Invoke(new()
        ///     {
        ///         Slug = "foobar",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetGithubAppResult> Invoke(GetGithubAppInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetGithubAppResult>("github:index/getGithubApp:getGithubApp", args ?? new GetGithubAppInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGithubAppArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The URL-friendly name of your GitHub App.
        /// </summary>
        [Input("slug", required: true)]
        public string Slug { get; set; } = null!;

        public GetGithubAppArgs()
        {
        }
        public static new GetGithubAppArgs Empty => new GetGithubAppArgs();
    }

    public sealed class GetGithubAppInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The URL-friendly name of your GitHub App.
        /// </summary>
        [Input("slug", required: true)]
        public Input<string> Slug { get; set; } = null!;

        public GetGithubAppInvokeArgs()
        {
        }
        public static new GetGithubAppInvokeArgs Empty => new GetGithubAppInvokeArgs();
    }


    [OutputType]
    public sealed class GetGithubAppResult
    {
        /// <summary>
        /// The app's description.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The app's full name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The Node ID of the app.
        /// </summary>
        public readonly string NodeId;
        public readonly string Slug;

        [OutputConstructor]
        private GetGithubAppResult(
            string description,

            string id,

            string name,

            string nodeId,

            string slug)
        {
            Description = description;
            Id = id;
            Name = name;
            NodeId = nodeId;
            Slug = slug;
        }
    }
}
