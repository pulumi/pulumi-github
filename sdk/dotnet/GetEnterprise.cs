// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github
{
    public static class GetEnterprise
    {
        /// <summary>
        /// Use this data source to retrieve basic information about a GitHub enterprise.
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Github = Pulumi.Github;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Github.GetEnterprise.Invoke(new()
        ///     {
        ///         Slug = "example-co",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetEnterpriseResult> InvokeAsync(GetEnterpriseArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetEnterpriseResult>("github:index/getEnterprise:getEnterprise", args ?? new GetEnterpriseArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve basic information about a GitHub enterprise.
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Github = Pulumi.Github;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Github.GetEnterprise.Invoke(new()
        ///     {
        ///         Slug = "example-co",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetEnterpriseResult> Invoke(GetEnterpriseInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetEnterpriseResult>("github:index/getEnterprise:getEnterprise", args ?? new GetEnterpriseInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetEnterpriseArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The URL slug identifying the enterprise.
        /// </summary>
        [Input("slug", required: true)]
        public string Slug { get; set; } = null!;

        public GetEnterpriseArgs()
        {
        }
        public static new GetEnterpriseArgs Empty => new GetEnterpriseArgs();
    }

    public sealed class GetEnterpriseInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The URL slug identifying the enterprise.
        /// </summary>
        [Input("slug", required: true)]
        public Input<string> Slug { get; set; } = null!;

        public GetEnterpriseInvokeArgs()
        {
        }
        public static new GetEnterpriseInvokeArgs Empty => new GetEnterpriseInvokeArgs();
    }


    [OutputType]
    public sealed class GetEnterpriseResult
    {
        /// <summary>
        /// The time the enterprise was created.
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// The description of the enterprise.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the enterprise.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The URL slug identifying the enterprise.
        /// </summary>
        public readonly string Slug;
        /// <summary>
        /// The url for the enterprise.
        /// </summary>
        public readonly string Url;

        [OutputConstructor]
        private GetEnterpriseResult(
            string createdAt,

            string description,

            string id,

            string name,

            string slug,

            string url)
        {
            CreatedAt = createdAt;
            Description = description;
            Id = id;
            Name = name;
            Slug = slug;
            Url = url;
        }
    }
}
