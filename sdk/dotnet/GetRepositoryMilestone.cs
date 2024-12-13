// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github
{
    public static class GetRepositoryMilestone
    {
        /// <summary>
        /// Use this data source to retrieve information about a specific GitHub milestone in a repository.
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
        ///     var example = Github.GetRepositoryMilestone.Invoke(new()
        ///     {
        ///         Owner = "example-owner",
        ///         Repository = "example-repository",
        ///         Number = 1,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetRepositoryMilestoneResult> InvokeAsync(GetRepositoryMilestoneArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRepositoryMilestoneResult>("github:index/getRepositoryMilestone:getRepositoryMilestone", args ?? new GetRepositoryMilestoneArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve information about a specific GitHub milestone in a repository.
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
        ///     var example = Github.GetRepositoryMilestone.Invoke(new()
        ///     {
        ///         Owner = "example-owner",
        ///         Repository = "example-repository",
        ///         Number = 1,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetRepositoryMilestoneResult> Invoke(GetRepositoryMilestoneInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRepositoryMilestoneResult>("github:index/getRepositoryMilestone:getRepositoryMilestone", args ?? new GetRepositoryMilestoneInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve information about a specific GitHub milestone in a repository.
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
        ///     var example = Github.GetRepositoryMilestone.Invoke(new()
        ///     {
        ///         Owner = "example-owner",
        ///         Repository = "example-repository",
        ///         Number = 1,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetRepositoryMilestoneResult> Invoke(GetRepositoryMilestoneInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetRepositoryMilestoneResult>("github:index/getRepositoryMilestone:getRepositoryMilestone", args ?? new GetRepositoryMilestoneInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRepositoryMilestoneArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The number of the milestone.
        /// </summary>
        [Input("number", required: true)]
        public int Number { get; set; }

        /// <summary>
        /// Owner of the repository.
        /// </summary>
        [Input("owner", required: true)]
        public string Owner { get; set; } = null!;

        /// <summary>
        /// Name of the repository to retrieve the milestone from.
        /// </summary>
        [Input("repository", required: true)]
        public string Repository { get; set; } = null!;

        public GetRepositoryMilestoneArgs()
        {
        }
        public static new GetRepositoryMilestoneArgs Empty => new GetRepositoryMilestoneArgs();
    }

    public sealed class GetRepositoryMilestoneInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The number of the milestone.
        /// </summary>
        [Input("number", required: true)]
        public Input<int> Number { get; set; } = null!;

        /// <summary>
        /// Owner of the repository.
        /// </summary>
        [Input("owner", required: true)]
        public Input<string> Owner { get; set; } = null!;

        /// <summary>
        /// Name of the repository to retrieve the milestone from.
        /// </summary>
        [Input("repository", required: true)]
        public Input<string> Repository { get; set; } = null!;

        public GetRepositoryMilestoneInvokeArgs()
        {
        }
        public static new GetRepositoryMilestoneInvokeArgs Empty => new GetRepositoryMilestoneInvokeArgs();
    }


    [OutputType]
    public sealed class GetRepositoryMilestoneResult
    {
        /// <summary>
        /// Description of the milestone.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The milestone due date (in ISO-8601 `yyyy-mm-dd` format).
        /// </summary>
        public readonly string DueDate;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly int Number;
        public readonly string Owner;
        public readonly string Repository;
        /// <summary>
        /// State of the milestone.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// Title of the milestone.
        /// </summary>
        public readonly string Title;

        [OutputConstructor]
        private GetRepositoryMilestoneResult(
            string description,

            string dueDate,

            string id,

            int number,

            string owner,

            string repository,

            string state,

            string title)
        {
            Description = description;
            DueDate = dueDate;
            Id = id;
            Number = number;
            Owner = owner;
            Repository = repository;
            State = state;
            Title = title;
        }
    }
}
