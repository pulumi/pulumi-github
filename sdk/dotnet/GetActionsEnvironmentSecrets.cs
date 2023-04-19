// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github
{
    public static class GetActionsEnvironmentSecrets
    {
        /// <summary>
        /// Use this data source to retrieve the list of secrets of the repository environment.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Github = Pulumi.Github;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Github.GetActionsEnvironmentSecrets.Invoke(new()
        ///     {
        ///         Environment = "exampleEnvironment",
        ///         Name = "exampleRepo",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetActionsEnvironmentSecretsResult> InvokeAsync(GetActionsEnvironmentSecretsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetActionsEnvironmentSecretsResult>("github:index/getActionsEnvironmentSecrets:getActionsEnvironmentSecrets", args ?? new GetActionsEnvironmentSecretsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve the list of secrets of the repository environment.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Github = Pulumi.Github;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Github.GetActionsEnvironmentSecrets.Invoke(new()
        ///     {
        ///         Environment = "exampleEnvironment",
        ///         Name = "exampleRepo",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetActionsEnvironmentSecretsResult> Invoke(GetActionsEnvironmentSecretsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetActionsEnvironmentSecretsResult>("github:index/getActionsEnvironmentSecrets:getActionsEnvironmentSecrets", args ?? new GetActionsEnvironmentSecretsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetActionsEnvironmentSecretsArgs : global::Pulumi.InvokeArgs
    {
        [Input("environment", required: true)]
        public string Environment { get; set; } = null!;

        [Input("fullName")]
        public string? FullName { get; set; }

        /// <summary>
        /// Name of the variable
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        public GetActionsEnvironmentSecretsArgs()
        {
        }
        public static new GetActionsEnvironmentSecretsArgs Empty => new GetActionsEnvironmentSecretsArgs();
    }

    public sealed class GetActionsEnvironmentSecretsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("environment", required: true)]
        public Input<string> Environment { get; set; } = null!;

        [Input("fullName")]
        public Input<string>? FullName { get; set; }

        /// <summary>
        /// Name of the variable
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public GetActionsEnvironmentSecretsInvokeArgs()
        {
        }
        public static new GetActionsEnvironmentSecretsInvokeArgs Empty => new GetActionsEnvironmentSecretsInvokeArgs();
    }


    [OutputType]
    public sealed class GetActionsEnvironmentSecretsResult
    {
        public readonly string Environment;
        public readonly string FullName;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Name of the variable
        /// </summary>
        public readonly string Name;
        public readonly ImmutableArray<Outputs.GetActionsEnvironmentSecretsSecretResult> Secrets;

        [OutputConstructor]
        private GetActionsEnvironmentSecretsResult(
            string environment,

            string fullName,

            string id,

            string name,

            ImmutableArray<Outputs.GetActionsEnvironmentSecretsSecretResult> secrets)
        {
            Environment = environment;
            FullName = fullName;
            Id = id;
            Name = name;
            Secrets = secrets;
        }
    }
}
