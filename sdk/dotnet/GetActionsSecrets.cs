// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github
{
    public static class GetActionsSecrets
    {
        /// <summary>
        /// Use this data source to retrieve the list of secrets for a GitHub repository.
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
        ///     var example = Github.GetActionsSecrets.Invoke(new()
        ///     {
        ///         Name = "example",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetActionsSecretsResult> InvokeAsync(GetActionsSecretsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetActionsSecretsResult>("github:index/getActionsSecrets:getActionsSecrets", args ?? new GetActionsSecretsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve the list of secrets for a GitHub repository.
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
        ///     var example = Github.GetActionsSecrets.Invoke(new()
        ///     {
        ///         Name = "example",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetActionsSecretsResult> Invoke(GetActionsSecretsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetActionsSecretsResult>("github:index/getActionsSecrets:getActionsSecrets", args ?? new GetActionsSecretsInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve the list of secrets for a GitHub repository.
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
        ///     var example = Github.GetActionsSecrets.Invoke(new()
        ///     {
        ///         Name = "example",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetActionsSecretsResult> Invoke(GetActionsSecretsInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetActionsSecretsResult>("github:index/getActionsSecrets:getActionsSecrets", args ?? new GetActionsSecretsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetActionsSecretsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Full name of the repository (in `org/name` format).
        /// </summary>
        [Input("fullName")]
        public string? FullName { get; set; }

        /// <summary>
        /// The name of the repository.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        public GetActionsSecretsArgs()
        {
        }
        public static new GetActionsSecretsArgs Empty => new GetActionsSecretsArgs();
    }

    public sealed class GetActionsSecretsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Full name of the repository (in `org/name` format).
        /// </summary>
        [Input("fullName")]
        public Input<string>? FullName { get; set; }

        /// <summary>
        /// The name of the repository.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public GetActionsSecretsInvokeArgs()
        {
        }
        public static new GetActionsSecretsInvokeArgs Empty => new GetActionsSecretsInvokeArgs();
    }


    [OutputType]
    public sealed class GetActionsSecretsResult
    {
        public readonly string FullName;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Secret name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// list of secrets for the repository
        /// </summary>
        public readonly ImmutableArray<Outputs.GetActionsSecretsSecretResult> Secrets;

        [OutputConstructor]
        private GetActionsSecretsResult(
            string fullName,

            string id,

            string name,

            ImmutableArray<Outputs.GetActionsSecretsSecretResult> secrets)
        {
            FullName = fullName;
            Id = id;
            Name = name;
            Secrets = secrets;
        }
    }
}
