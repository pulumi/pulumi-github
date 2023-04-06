// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github
{
    public static class GetSshKeys
    {
        /// <summary>
        /// Use this data source to retrieve information about GitHub's SSH keys.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Github = Pulumi.Github;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var test = Github.GetSshKeys.Invoke();
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetSshKeysResult> InvokeAsync(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSshKeysResult>("github:index/getSshKeys:getSshKeys", InvokeArgs.Empty, options.WithDefaults());
    }


    [OutputType]
    public sealed class GetSshKeysResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// An array of GitHub's SSH public keys.
        /// </summary>
        public readonly ImmutableArray<string> Keys;

        [OutputConstructor]
        private GetSshKeysResult(
            string id,

            ImmutableArray<string> keys)
        {
            Id = id;
            Keys = keys;
        }
    }
}
