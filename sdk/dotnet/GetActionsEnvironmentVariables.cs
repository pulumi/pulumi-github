// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github
{
    public static class GetActionsEnvironmentVariables
    {
        /// <summary>
        /// Use this data source to retrieve the list of variables of the repository environment.
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
        ///     var example = Github.GetActionsEnvironmentVariables.Invoke(new()
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
        public static Task<GetActionsEnvironmentVariablesResult> InvokeAsync(GetActionsEnvironmentVariablesArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetActionsEnvironmentVariablesResult>("github:index/getActionsEnvironmentVariables:getActionsEnvironmentVariables", args ?? new GetActionsEnvironmentVariablesArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve the list of variables of the repository environment.
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
        ///     var example = Github.GetActionsEnvironmentVariables.Invoke(new()
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
        public static Output<GetActionsEnvironmentVariablesResult> Invoke(GetActionsEnvironmentVariablesInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetActionsEnvironmentVariablesResult>("github:index/getActionsEnvironmentVariables:getActionsEnvironmentVariables", args ?? new GetActionsEnvironmentVariablesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetActionsEnvironmentVariablesArgs : global::Pulumi.InvokeArgs
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

        public GetActionsEnvironmentVariablesArgs()
        {
        }
        public static new GetActionsEnvironmentVariablesArgs Empty => new GetActionsEnvironmentVariablesArgs();
    }

    public sealed class GetActionsEnvironmentVariablesInvokeArgs : global::Pulumi.InvokeArgs
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

        public GetActionsEnvironmentVariablesInvokeArgs()
        {
        }
        public static new GetActionsEnvironmentVariablesInvokeArgs Empty => new GetActionsEnvironmentVariablesInvokeArgs();
    }


    [OutputType]
    public sealed class GetActionsEnvironmentVariablesResult
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
        /// <summary>
        /// list of variables for the environment
        /// </summary>
        public readonly ImmutableArray<Outputs.GetActionsEnvironmentVariablesVariableResult> Variables;

        [OutputConstructor]
        private GetActionsEnvironmentVariablesResult(
            string environment,

            string fullName,

            string id,

            string name,

            ImmutableArray<Outputs.GetActionsEnvironmentVariablesVariableResult> variables)
        {
            Environment = environment;
            FullName = fullName;
            Id = id;
            Name = name;
            Variables = variables;
        }
    }
}
