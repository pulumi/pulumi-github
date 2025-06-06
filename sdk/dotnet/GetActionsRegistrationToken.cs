// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github
{
    public static class GetActionsRegistrationToken
    {
        /// <summary>
        /// Use this data source to retrieve a GitHub Actions repository registration token. This token can then be used to register a self-hosted runner.
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
        ///     var example = Github.GetActionsRegistrationToken.Invoke(new()
        ///     {
        ///         Repository = "example_repo",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetActionsRegistrationTokenResult> InvokeAsync(GetActionsRegistrationTokenArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetActionsRegistrationTokenResult>("github:index/getActionsRegistrationToken:getActionsRegistrationToken", args ?? new GetActionsRegistrationTokenArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve a GitHub Actions repository registration token. This token can then be used to register a self-hosted runner.
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
        ///     var example = Github.GetActionsRegistrationToken.Invoke(new()
        ///     {
        ///         Repository = "example_repo",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetActionsRegistrationTokenResult> Invoke(GetActionsRegistrationTokenInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetActionsRegistrationTokenResult>("github:index/getActionsRegistrationToken:getActionsRegistrationToken", args ?? new GetActionsRegistrationTokenInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve a GitHub Actions repository registration token. This token can then be used to register a self-hosted runner.
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
        ///     var example = Github.GetActionsRegistrationToken.Invoke(new()
        ///     {
        ///         Repository = "example_repo",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetActionsRegistrationTokenResult> Invoke(GetActionsRegistrationTokenInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetActionsRegistrationTokenResult>("github:index/getActionsRegistrationToken:getActionsRegistrationToken", args ?? new GetActionsRegistrationTokenInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetActionsRegistrationTokenArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the repository to get a GitHub Actions registration token for.
        /// </summary>
        [Input("repository", required: true)]
        public string Repository { get; set; } = null!;

        public GetActionsRegistrationTokenArgs()
        {
        }
        public static new GetActionsRegistrationTokenArgs Empty => new GetActionsRegistrationTokenArgs();
    }

    public sealed class GetActionsRegistrationTokenInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the repository to get a GitHub Actions registration token for.
        /// </summary>
        [Input("repository", required: true)]
        public Input<string> Repository { get; set; } = null!;

        public GetActionsRegistrationTokenInvokeArgs()
        {
        }
        public static new GetActionsRegistrationTokenInvokeArgs Empty => new GetActionsRegistrationTokenInvokeArgs();
    }


    [OutputType]
    public sealed class GetActionsRegistrationTokenResult
    {
        /// <summary>
        /// The token expiration date.
        /// </summary>
        public readonly int ExpiresAt;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Repository;
        /// <summary>
        /// The token that has been retrieved.
        /// </summary>
        public readonly string Token;

        [OutputConstructor]
        private GetActionsRegistrationTokenResult(
            int expiresAt,

            string id,

            string repository,

            string token)
        {
            ExpiresAt = expiresAt;
            Id = id;
            Repository = repository;
            Token = token;
        }
    }
}
