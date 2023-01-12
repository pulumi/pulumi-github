// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github
{
    public static class GetRepositoryTeams
    {
        /// <summary>
        /// Use this data source to retrieve the list of teams which have access to a GitHub repository.
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
        ///     var example = Github.GetRepositoryTeams.Invoke(new()
        ///     {
        ///         Name = "example",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetRepositoryTeamsResult> InvokeAsync(GetRepositoryTeamsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRepositoryTeamsResult>("github:index/getRepositoryTeams:getRepositoryTeams", args ?? new GetRepositoryTeamsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve the list of teams which have access to a GitHub repository.
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
        ///     var example = Github.GetRepositoryTeams.Invoke(new()
        ///     {
        ///         Name = "example",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetRepositoryTeamsResult> Invoke(GetRepositoryTeamsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRepositoryTeamsResult>("github:index/getRepositoryTeams:getRepositoryTeams", args ?? new GetRepositoryTeamsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRepositoryTeamsArgs : global::Pulumi.InvokeArgs
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

        public GetRepositoryTeamsArgs()
        {
        }
        public static new GetRepositoryTeamsArgs Empty => new GetRepositoryTeamsArgs();
    }

    public sealed class GetRepositoryTeamsInvokeArgs : global::Pulumi.InvokeArgs
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

        public GetRepositoryTeamsInvokeArgs()
        {
        }
        public static new GetRepositoryTeamsInvokeArgs Empty => new GetRepositoryTeamsInvokeArgs();
    }


    [OutputType]
    public sealed class GetRepositoryTeamsResult
    {
        public readonly string FullName;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Team name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// List of teams which have access to the repository
        /// </summary>
        public readonly ImmutableArray<Outputs.GetRepositoryTeamsTeamResult> Teams;

        [OutputConstructor]
        private GetRepositoryTeamsResult(
            string fullName,

            string id,

            string name,

            ImmutableArray<Outputs.GetRepositoryTeamsTeamResult> teams)
        {
            FullName = fullName;
            Id = id;
            Name = name;
            Teams = teams;
        }
    }
}
