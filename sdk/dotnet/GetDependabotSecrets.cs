// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github
{
    public static class GetDependabotSecrets
    {
        /// <summary>
        /// Use this data source to retrieve the list of dependabot secrets for a GitHub repository.
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
        ///     var example = Github.GetDependabotSecrets.Invoke(new()
        ///     {
        ///         Name = "example",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetDependabotSecretsResult> InvokeAsync(GetDependabotSecretsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDependabotSecretsResult>("github:index/getDependabotSecrets:getDependabotSecrets", args ?? new GetDependabotSecretsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve the list of dependabot secrets for a GitHub repository.
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
        ///     var example = Github.GetDependabotSecrets.Invoke(new()
        ///     {
        ///         Name = "example",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetDependabotSecretsResult> Invoke(GetDependabotSecretsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDependabotSecretsResult>("github:index/getDependabotSecrets:getDependabotSecrets", args ?? new GetDependabotSecretsInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve the list of dependabot secrets for a GitHub repository.
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
        ///     var example = Github.GetDependabotSecrets.Invoke(new()
        ///     {
        ///         Name = "example",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetDependabotSecretsResult> Invoke(GetDependabotSecretsInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetDependabotSecretsResult>("github:index/getDependabotSecrets:getDependabotSecrets", args ?? new GetDependabotSecretsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDependabotSecretsArgs : global::Pulumi.InvokeArgs
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

        public GetDependabotSecretsArgs()
        {
        }
        public static new GetDependabotSecretsArgs Empty => new GetDependabotSecretsArgs();
    }

    public sealed class GetDependabotSecretsInvokeArgs : global::Pulumi.InvokeArgs
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

        public GetDependabotSecretsInvokeArgs()
        {
        }
        public static new GetDependabotSecretsInvokeArgs Empty => new GetDependabotSecretsInvokeArgs();
    }


    [OutputType]
    public sealed class GetDependabotSecretsResult
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
        /// list of dependabot secrets for the repository
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDependabotSecretsSecretResult> Secrets;

        [OutputConstructor]
        private GetDependabotSecretsResult(
            string fullName,

            string id,

            string name,

            ImmutableArray<Outputs.GetDependabotSecretsSecretResult> secrets)
        {
            FullName = fullName;
            Id = id;
            Name = name;
            Secrets = secrets;
        }
    }
}
