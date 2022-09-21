// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github
{
    [GithubResourceType("github:index/repositoryPullRequest:RepositoryPullRequest")]
    public partial class RepositoryPullRequest : global::Pulumi.CustomResource
    {
        [Output("baseRef")]
        public Output<string> BaseRef { get; private set; } = null!;

        [Output("baseRepository")]
        public Output<string> BaseRepository { get; private set; } = null!;

        [Output("baseSha")]
        public Output<string> BaseSha { get; private set; } = null!;

        [Output("body")]
        public Output<string?> Body { get; private set; } = null!;

        [Output("draft")]
        public Output<bool> Draft { get; private set; } = null!;

        [Output("headRef")]
        public Output<string> HeadRef { get; private set; } = null!;

        [Output("headSha")]
        public Output<string> HeadSha { get; private set; } = null!;

        /// <summary>
        /// List of names of labels on the PR
        /// </summary>
        [Output("labels")]
        public Output<ImmutableArray<string>> Labels { get; private set; } = null!;

        [Output("maintainerCanModify")]
        public Output<bool?> MaintainerCanModify { get; private set; } = null!;

        [Output("number")]
        public Output<int> Number { get; private set; } = null!;

        [Output("openedAt")]
        public Output<int> OpenedAt { get; private set; } = null!;

        /// <summary>
        /// Username of the PR creator
        /// </summary>
        [Output("openedBy")]
        public Output<string> OpenedBy { get; private set; } = null!;

        [Output("owner")]
        public Output<string?> Owner { get; private set; } = null!;

        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        [Output("title")]
        public Output<string> Title { get; private set; } = null!;

        [Output("updatedAt")]
        public Output<int> UpdatedAt { get; private set; } = null!;


        /// <summary>
        /// Create a RepositoryPullRequest resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RepositoryPullRequest(string name, RepositoryPullRequestArgs args, CustomResourceOptions? options = null)
            : base("github:index/repositoryPullRequest:RepositoryPullRequest", name, args ?? new RepositoryPullRequestArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RepositoryPullRequest(string name, Input<string> id, RepositoryPullRequestState? state = null, CustomResourceOptions? options = null)
            : base("github:index/repositoryPullRequest:RepositoryPullRequest", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing RepositoryPullRequest resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RepositoryPullRequest Get(string name, Input<string> id, RepositoryPullRequestState? state = null, CustomResourceOptions? options = null)
        {
            return new RepositoryPullRequest(name, id, state, options);
        }
    }

    public sealed class RepositoryPullRequestArgs : global::Pulumi.ResourceArgs
    {
        [Input("baseRef", required: true)]
        public Input<string> BaseRef { get; set; } = null!;

        [Input("baseRepository", required: true)]
        public Input<string> BaseRepository { get; set; } = null!;

        [Input("body")]
        public Input<string>? Body { get; set; }

        [Input("headRef", required: true)]
        public Input<string> HeadRef { get; set; } = null!;

        [Input("maintainerCanModify")]
        public Input<bool>? MaintainerCanModify { get; set; }

        [Input("owner")]
        public Input<string>? Owner { get; set; }

        [Input("title", required: true)]
        public Input<string> Title { get; set; } = null!;

        public RepositoryPullRequestArgs()
        {
        }
        public static new RepositoryPullRequestArgs Empty => new RepositoryPullRequestArgs();
    }

    public sealed class RepositoryPullRequestState : global::Pulumi.ResourceArgs
    {
        [Input("baseRef")]
        public Input<string>? BaseRef { get; set; }

        [Input("baseRepository")]
        public Input<string>? BaseRepository { get; set; }

        [Input("baseSha")]
        public Input<string>? BaseSha { get; set; }

        [Input("body")]
        public Input<string>? Body { get; set; }

        [Input("draft")]
        public Input<bool>? Draft { get; set; }

        [Input("headRef")]
        public Input<string>? HeadRef { get; set; }

        [Input("headSha")]
        public Input<string>? HeadSha { get; set; }

        [Input("labels")]
        private InputList<string>? _labels;

        /// <summary>
        /// List of names of labels on the PR
        /// </summary>
        public InputList<string> Labels
        {
            get => _labels ?? (_labels = new InputList<string>());
            set => _labels = value;
        }

        [Input("maintainerCanModify")]
        public Input<bool>? MaintainerCanModify { get; set; }

        [Input("number")]
        public Input<int>? Number { get; set; }

        [Input("openedAt")]
        public Input<int>? OpenedAt { get; set; }

        /// <summary>
        /// Username of the PR creator
        /// </summary>
        [Input("openedBy")]
        public Input<string>? OpenedBy { get; set; }

        [Input("owner")]
        public Input<string>? Owner { get; set; }

        [Input("state")]
        public Input<string>? State { get; set; }

        [Input("title")]
        public Input<string>? Title { get; set; }

        [Input("updatedAt")]
        public Input<int>? UpdatedAt { get; set; }

        public RepositoryPullRequestState()
        {
        }
        public static new RepositoryPullRequestState Empty => new RepositoryPullRequestState();
    }
}
