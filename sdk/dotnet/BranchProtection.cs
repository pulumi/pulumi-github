// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Github
{
    /// <summary>
    /// Protects a GitHub branch.
    /// 
    /// This resource allows you to configure branch protection for repositories in your organization. When applied, the branch will be protected from forced pushes and deletion. Additional constraints, such as required status checks or restrictions on users, teams, and apps, can also be configured.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Github = Pulumi.Github;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var exampleRepository = new Github.Repository("exampleRepository");
    /// 
    ///     var exampleUser = Github.GetUser.Invoke(new()
    ///     {
    ///         Username = "example",
    ///     });
    /// 
    ///     var exampleTeam = new Github.Team("exampleTeam");
    /// 
    ///     // Protect the main branch of the foo repository. Additionally, require that
    ///     // the "ci/travis" context to be passing and only allow the engineers team merge
    ///     // to the branch.
    ///     var exampleBranchProtection = new Github.BranchProtection("exampleBranchProtection", new()
    ///     {
    ///         RepositoryId = exampleRepository.NodeId,
    ///         Pattern = "main",
    ///         EnforceAdmins = true,
    ///         AllowsDeletions = true,
    ///         RequiredStatusChecks = new[]
    ///         {
    ///             new Github.Inputs.BranchProtectionRequiredStatusCheckArgs
    ///             {
    ///                 Strict = false,
    ///                 Contexts = new[]
    ///                 {
    ///                     "ci/travis",
    ///                 },
    ///             },
    ///         },
    ///         RequiredPullRequestReviews = new[]
    ///         {
    ///             new Github.Inputs.BranchProtectionRequiredPullRequestReviewArgs
    ///             {
    ///                 DismissStaleReviews = true,
    ///                 RestrictDismissals = true,
    ///                 DismissalRestrictions = new[]
    ///                 {
    ///                     exampleUser.Apply(getUserResult =&gt; getUserResult.NodeId),
    ///                     exampleTeam.NodeId,
    ///                     "/exampleuser",
    ///                     "exampleorganization/exampleteam",
    ///                 },
    ///             },
    ///         },
    ///         PushRestrictions = new[]
    ///         {
    ///             exampleUser.Apply(getUserResult =&gt; getUserResult.NodeId),
    ///             "/exampleuser",
    ///             "exampleorganization/exampleteam",
    ///         },
    ///     });
    /// 
    ///     var exampleTeamRepository = new Github.TeamRepository("exampleTeamRepository", new()
    ///     {
    ///         TeamId = exampleTeam.Id,
    ///         Repository = exampleRepository.Name,
    ///         Permission = "pull",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// GitHub Branch Protection can be imported using an ID made up of `repository:pattern`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import github:index/branchProtection:BranchProtection terraform terraform:main
    /// ```
    /// </summary>
    [GithubResourceType("github:index/branchProtection:BranchProtection")]
    public partial class BranchProtection : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Boolean, setting this to `true` to allow the branch to be deleted.
        /// </summary>
        [Output("allowsDeletions")]
        public Output<bool?> AllowsDeletions { get; private set; } = null!;

        /// <summary>
        /// Boolean, setting this to `true` to allow force pushes on the branch.
        /// </summary>
        [Output("allowsForcePushes")]
        public Output<bool?> AllowsForcePushes { get; private set; } = null!;

        /// <summary>
        /// Boolean, setting this to `true` to block creating the branch.
        /// </summary>
        [Output("blocksCreations")]
        public Output<bool?> BlocksCreations { get; private set; } = null!;

        /// <summary>
        /// Boolean, setting this to `true` enforces status checks for repository administrators.
        /// </summary>
        [Output("enforceAdmins")]
        public Output<bool?> EnforceAdmins { get; private set; } = null!;

        /// <summary>
        /// Boolean, Setting this to `true` will make the branch read-only and preventing any pushes to it. Defaults to `false`
        /// </summary>
        [Output("lockBranch")]
        public Output<bool?> LockBranch { get; private set; } = null!;

        /// <summary>
        /// Identifies the protection rule pattern.
        /// </summary>
        [Output("pattern")]
        public Output<string> Pattern { get; private set; } = null!;

        /// <summary>
        /// The list of actor Names/IDs that may push to the branch. Actor names must either begin with a "/" for users or the organization name followed by a "/" for teams.
        /// </summary>
        [Output("pushRestrictions")]
        public Output<ImmutableArray<string>> PushRestrictions { get; private set; } = null!;

        /// <summary>
        /// The name or node ID of the repository associated with this branch protection rule.
        /// </summary>
        [Output("repositoryId")]
        public Output<string> RepositoryId { get; private set; } = null!;

        /// <summary>
        /// Boolean, setting this to `true` requires all conversations on code must be resolved before a pull request can be merged.
        /// </summary>
        [Output("requireConversationResolution")]
        public Output<bool?> RequireConversationResolution { get; private set; } = null!;

        /// <summary>
        /// Boolean, setting this to `true` requires all commits to be signed with GPG.
        /// </summary>
        [Output("requireSignedCommits")]
        public Output<bool?> RequireSignedCommits { get; private set; } = null!;

        /// <summary>
        /// Boolean, setting this to `true` enforces a linear commit Git history, which prevents anyone from pushing merge commits to a branch
        /// </summary>
        [Output("requiredLinearHistory")]
        public Output<bool?> RequiredLinearHistory { get; private set; } = null!;

        /// <summary>
        /// Enforce restrictions for pull request reviews. See Required Pull Request Reviews below for details.
        /// </summary>
        [Output("requiredPullRequestReviews")]
        public Output<ImmutableArray<Outputs.BranchProtectionRequiredPullRequestReview>> RequiredPullRequestReviews { get; private set; } = null!;

        /// <summary>
        /// Enforce restrictions for required status checks. See Required Status Checks below for details.
        /// </summary>
        [Output("requiredStatusChecks")]
        public Output<ImmutableArray<Outputs.BranchProtectionRequiredStatusCheck>> RequiredStatusChecks { get; private set; } = null!;


        /// <summary>
        /// Create a BranchProtection resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BranchProtection(string name, BranchProtectionArgs args, CustomResourceOptions? options = null)
            : base("github:index/branchProtection:BranchProtection", name, args ?? new BranchProtectionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BranchProtection(string name, Input<string> id, BranchProtectionState? state = null, CustomResourceOptions? options = null)
            : base("github:index/branchProtection:BranchProtection", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing BranchProtection resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BranchProtection Get(string name, Input<string> id, BranchProtectionState? state = null, CustomResourceOptions? options = null)
        {
            return new BranchProtection(name, id, state, options);
        }
    }

    public sealed class BranchProtectionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Boolean, setting this to `true` to allow the branch to be deleted.
        /// </summary>
        [Input("allowsDeletions")]
        public Input<bool>? AllowsDeletions { get; set; }

        /// <summary>
        /// Boolean, setting this to `true` to allow force pushes on the branch.
        /// </summary>
        [Input("allowsForcePushes")]
        public Input<bool>? AllowsForcePushes { get; set; }

        /// <summary>
        /// Boolean, setting this to `true` to block creating the branch.
        /// </summary>
        [Input("blocksCreations")]
        public Input<bool>? BlocksCreations { get; set; }

        /// <summary>
        /// Boolean, setting this to `true` enforces status checks for repository administrators.
        /// </summary>
        [Input("enforceAdmins")]
        public Input<bool>? EnforceAdmins { get; set; }

        /// <summary>
        /// Boolean, Setting this to `true` will make the branch read-only and preventing any pushes to it. Defaults to `false`
        /// </summary>
        [Input("lockBranch")]
        public Input<bool>? LockBranch { get; set; }

        /// <summary>
        /// Identifies the protection rule pattern.
        /// </summary>
        [Input("pattern", required: true)]
        public Input<string> Pattern { get; set; } = null!;

        [Input("pushRestrictions")]
        private InputList<string>? _pushRestrictions;

        /// <summary>
        /// The list of actor Names/IDs that may push to the branch. Actor names must either begin with a "/" for users or the organization name followed by a "/" for teams.
        /// </summary>
        public InputList<string> PushRestrictions
        {
            get => _pushRestrictions ?? (_pushRestrictions = new InputList<string>());
            set => _pushRestrictions = value;
        }

        /// <summary>
        /// The name or node ID of the repository associated with this branch protection rule.
        /// </summary>
        [Input("repositoryId", required: true)]
        public Input<string> RepositoryId { get; set; } = null!;

        /// <summary>
        /// Boolean, setting this to `true` requires all conversations on code must be resolved before a pull request can be merged.
        /// </summary>
        [Input("requireConversationResolution")]
        public Input<bool>? RequireConversationResolution { get; set; }

        /// <summary>
        /// Boolean, setting this to `true` requires all commits to be signed with GPG.
        /// </summary>
        [Input("requireSignedCommits")]
        public Input<bool>? RequireSignedCommits { get; set; }

        /// <summary>
        /// Boolean, setting this to `true` enforces a linear commit Git history, which prevents anyone from pushing merge commits to a branch
        /// </summary>
        [Input("requiredLinearHistory")]
        public Input<bool>? RequiredLinearHistory { get; set; }

        [Input("requiredPullRequestReviews")]
        private InputList<Inputs.BranchProtectionRequiredPullRequestReviewArgs>? _requiredPullRequestReviews;

        /// <summary>
        /// Enforce restrictions for pull request reviews. See Required Pull Request Reviews below for details.
        /// </summary>
        public InputList<Inputs.BranchProtectionRequiredPullRequestReviewArgs> RequiredPullRequestReviews
        {
            get => _requiredPullRequestReviews ?? (_requiredPullRequestReviews = new InputList<Inputs.BranchProtectionRequiredPullRequestReviewArgs>());
            set => _requiredPullRequestReviews = value;
        }

        [Input("requiredStatusChecks")]
        private InputList<Inputs.BranchProtectionRequiredStatusCheckArgs>? _requiredStatusChecks;

        /// <summary>
        /// Enforce restrictions for required status checks. See Required Status Checks below for details.
        /// </summary>
        public InputList<Inputs.BranchProtectionRequiredStatusCheckArgs> RequiredStatusChecks
        {
            get => _requiredStatusChecks ?? (_requiredStatusChecks = new InputList<Inputs.BranchProtectionRequiredStatusCheckArgs>());
            set => _requiredStatusChecks = value;
        }

        public BranchProtectionArgs()
        {
        }
        public static new BranchProtectionArgs Empty => new BranchProtectionArgs();
    }

    public sealed class BranchProtectionState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Boolean, setting this to `true` to allow the branch to be deleted.
        /// </summary>
        [Input("allowsDeletions")]
        public Input<bool>? AllowsDeletions { get; set; }

        /// <summary>
        /// Boolean, setting this to `true` to allow force pushes on the branch.
        /// </summary>
        [Input("allowsForcePushes")]
        public Input<bool>? AllowsForcePushes { get; set; }

        /// <summary>
        /// Boolean, setting this to `true` to block creating the branch.
        /// </summary>
        [Input("blocksCreations")]
        public Input<bool>? BlocksCreations { get; set; }

        /// <summary>
        /// Boolean, setting this to `true` enforces status checks for repository administrators.
        /// </summary>
        [Input("enforceAdmins")]
        public Input<bool>? EnforceAdmins { get; set; }

        /// <summary>
        /// Boolean, Setting this to `true` will make the branch read-only and preventing any pushes to it. Defaults to `false`
        /// </summary>
        [Input("lockBranch")]
        public Input<bool>? LockBranch { get; set; }

        /// <summary>
        /// Identifies the protection rule pattern.
        /// </summary>
        [Input("pattern")]
        public Input<string>? Pattern { get; set; }

        [Input("pushRestrictions")]
        private InputList<string>? _pushRestrictions;

        /// <summary>
        /// The list of actor Names/IDs that may push to the branch. Actor names must either begin with a "/" for users or the organization name followed by a "/" for teams.
        /// </summary>
        public InputList<string> PushRestrictions
        {
            get => _pushRestrictions ?? (_pushRestrictions = new InputList<string>());
            set => _pushRestrictions = value;
        }

        /// <summary>
        /// The name or node ID of the repository associated with this branch protection rule.
        /// </summary>
        [Input("repositoryId")]
        public Input<string>? RepositoryId { get; set; }

        /// <summary>
        /// Boolean, setting this to `true` requires all conversations on code must be resolved before a pull request can be merged.
        /// </summary>
        [Input("requireConversationResolution")]
        public Input<bool>? RequireConversationResolution { get; set; }

        /// <summary>
        /// Boolean, setting this to `true` requires all commits to be signed with GPG.
        /// </summary>
        [Input("requireSignedCommits")]
        public Input<bool>? RequireSignedCommits { get; set; }

        /// <summary>
        /// Boolean, setting this to `true` enforces a linear commit Git history, which prevents anyone from pushing merge commits to a branch
        /// </summary>
        [Input("requiredLinearHistory")]
        public Input<bool>? RequiredLinearHistory { get; set; }

        [Input("requiredPullRequestReviews")]
        private InputList<Inputs.BranchProtectionRequiredPullRequestReviewGetArgs>? _requiredPullRequestReviews;

        /// <summary>
        /// Enforce restrictions for pull request reviews. See Required Pull Request Reviews below for details.
        /// </summary>
        public InputList<Inputs.BranchProtectionRequiredPullRequestReviewGetArgs> RequiredPullRequestReviews
        {
            get => _requiredPullRequestReviews ?? (_requiredPullRequestReviews = new InputList<Inputs.BranchProtectionRequiredPullRequestReviewGetArgs>());
            set => _requiredPullRequestReviews = value;
        }

        [Input("requiredStatusChecks")]
        private InputList<Inputs.BranchProtectionRequiredStatusCheckGetArgs>? _requiredStatusChecks;

        /// <summary>
        /// Enforce restrictions for required status checks. See Required Status Checks below for details.
        /// </summary>
        public InputList<Inputs.BranchProtectionRequiredStatusCheckGetArgs> RequiredStatusChecks
        {
            get => _requiredStatusChecks ?? (_requiredStatusChecks = new InputList<Inputs.BranchProtectionRequiredStatusCheckGetArgs>());
            set => _requiredStatusChecks = value;
        }

        public BranchProtectionState()
        {
        }
        public static new BranchProtectionState Empty => new BranchProtectionState();
    }
}
