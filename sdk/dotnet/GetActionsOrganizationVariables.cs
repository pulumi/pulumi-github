// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github
{
    public static class GetActionsOrganizationVariables
    {
        /// <summary>
        /// Use this data source to retrieve the list of variables of the organization.
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
        ///     var example = Github.GetActionsOrganizationVariables.Invoke();
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetActionsOrganizationVariablesResult> InvokeAsync(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetActionsOrganizationVariablesResult>("github:index/getActionsOrganizationVariables:getActionsOrganizationVariables", InvokeArgs.Empty, options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve the list of variables of the organization.
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
        ///     var example = Github.GetActionsOrganizationVariables.Invoke();
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetActionsOrganizationVariablesResult> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetActionsOrganizationVariablesResult>("github:index/getActionsOrganizationVariables:getActionsOrganizationVariables", InvokeArgs.Empty, options.WithDefaults());
    }


    [OutputType]
    public sealed class GetActionsOrganizationVariablesResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// list of variables for the repository
        /// </summary>
        public readonly ImmutableArray<Outputs.GetActionsOrganizationVariablesVariableResult> Variables;

        [OutputConstructor]
        private GetActionsOrganizationVariablesResult(
            string id,

            ImmutableArray<Outputs.GetActionsOrganizationVariablesVariableResult> variables)
        {
            Id = id;
            Variables = variables;
        }
    }
}
