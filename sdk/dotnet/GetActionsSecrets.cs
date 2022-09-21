// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
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
        public static Task<GetActionsSecretsResult> InvokeAsync(GetActionsSecretsArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetActionsSecretsResult>("github:index/getActionsSecrets:getActionsSecrets", args ?? new GetActionsSecretsArgs(), options.WithDefaults());

        public static Output<GetActionsSecretsResult> Invoke(GetActionsSecretsInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetActionsSecretsResult>("github:index/getActionsSecrets:getActionsSecrets", args ?? new GetActionsSecretsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetActionsSecretsArgs : global::Pulumi.InvokeArgs
    {
        [Input("fullName")]
        public string? FullName { get; set; }

        [Input("name")]
        public string? Name { get; set; }

        public GetActionsSecretsArgs()
        {
        }
        public static new GetActionsSecretsArgs Empty => new GetActionsSecretsArgs();
    }

    public sealed class GetActionsSecretsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("fullName")]
        public Input<string>? FullName { get; set; }

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
        public readonly string Name;
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
