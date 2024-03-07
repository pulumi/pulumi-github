// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github
{
    public static class GetOrganizationTeams
    {
        /// <summary>
        /// Use this data source to retrieve information about all GitHub teams in an organization.
        /// 
        /// ## Example Usage
        /// 
        /// To retrieve *all* teams of the organization:
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
        ///     var all = Github.GetOrganizationTeams.Invoke();
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// 
        /// To retrieve only the team's at the root of the organization:
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
        ///     var rootTeams = Github.GetOrganizationTeams.Invoke(new()
        ///     {
        ///         RootTeamsOnly = true,
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetOrganizationTeamsResult> InvokeAsync(GetOrganizationTeamsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetOrganizationTeamsResult>("github:index/getOrganizationTeams:getOrganizationTeams", args ?? new GetOrganizationTeamsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve information about all GitHub teams in an organization.
        /// 
        /// ## Example Usage
        /// 
        /// To retrieve *all* teams of the organization:
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
        ///     var all = Github.GetOrganizationTeams.Invoke();
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// 
        /// To retrieve only the team's at the root of the organization:
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
        ///     var rootTeams = Github.GetOrganizationTeams.Invoke(new()
        ///     {
        ///         RootTeamsOnly = true,
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetOrganizationTeamsResult> Invoke(GetOrganizationTeamsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetOrganizationTeamsResult>("github:index/getOrganizationTeams:getOrganizationTeams", args ?? new GetOrganizationTeamsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetOrganizationTeamsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// (Optional) Set the number of results per graphql query. Reducing this number can alleviate timeout errors. Accepts a value between 0 - 100. Defaults to `100`.
        /// </summary>
        [Input("resultsPerPage")]
        public int? ResultsPerPage { get; set; }

        /// <summary>
        /// (Optional) Only return teams that are at the organization's root, i.e. no nested teams. Defaults to `false`.
        /// </summary>
        [Input("rootTeamsOnly")]
        public bool? RootTeamsOnly { get; set; }

        /// <summary>
        /// (Optional) Exclude the members and repositories of the team from the returned result. Defaults to `false`.
        /// </summary>
        [Input("summaryOnly")]
        public bool? SummaryOnly { get; set; }

        public GetOrganizationTeamsArgs()
        {
        }
        public static new GetOrganizationTeamsArgs Empty => new GetOrganizationTeamsArgs();
    }

    public sealed class GetOrganizationTeamsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// (Optional) Set the number of results per graphql query. Reducing this number can alleviate timeout errors. Accepts a value between 0 - 100. Defaults to `100`.
        /// </summary>
        [Input("resultsPerPage")]
        public Input<int>? ResultsPerPage { get; set; }

        /// <summary>
        /// (Optional) Only return teams that are at the organization's root, i.e. no nested teams. Defaults to `false`.
        /// </summary>
        [Input("rootTeamsOnly")]
        public Input<bool>? RootTeamsOnly { get; set; }

        /// <summary>
        /// (Optional) Exclude the members and repositories of the team from the returned result. Defaults to `false`.
        /// </summary>
        [Input("summaryOnly")]
        public Input<bool>? SummaryOnly { get; set; }

        public GetOrganizationTeamsInvokeArgs()
        {
        }
        public static new GetOrganizationTeamsInvokeArgs Empty => new GetOrganizationTeamsInvokeArgs();
    }


    [OutputType]
    public sealed class GetOrganizationTeamsResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// (Optional) Set the number of results per graphql query. Reducing this number can alleviate timeout errors. Accepts a value between 0 - 100. Defaults to `100`.
        /// </summary>
        public readonly int? ResultsPerPage;
        /// <summary>
        /// (Optional) Only return teams that are at the organization's root, i.e. no nested teams. Defaults to `false`.
        /// </summary>
        public readonly bool? RootTeamsOnly;
        /// <summary>
        /// (Optional) Exclude the members and repositories of the team from the returned result. Defaults to `false`.
        /// </summary>
        public readonly bool? SummaryOnly;
        /// <summary>
        /// (Required) An Array of GitHub Teams.  Each `team` block consists of the fields documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetOrganizationTeamsTeamResult> Teams;

        [OutputConstructor]
        private GetOrganizationTeamsResult(
            string id,

            int? resultsPerPage,

            bool? rootTeamsOnly,

            bool? summaryOnly,

            ImmutableArray<Outputs.GetOrganizationTeamsTeamResult> teams)
        {
            Id = id;
            ResultsPerPage = resultsPerPage;
            RootTeamsOnly = rootTeamsOnly;
            SummaryOnly = summaryOnly;
            Teams = teams;
        }
    }
}
