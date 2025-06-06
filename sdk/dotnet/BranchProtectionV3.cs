// *** WARNING: this file was generated by pulumi-language-dotnet. ***
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
    /// The `github.BranchProtection` resource has moved to the GraphQL API, while this resource will continue to leverage the REST API.
    /// 
    /// This resource allows you to configure branch protection for repositories in your organization. When applied, the branch will be protected from forced pushes and deletion. Additional constraints, such as required status checks or restrictions on users, teams, and apps, can also be configured.
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
    ///     // Protect the main branch of the foo repository. Only allow a specific user to merge to the branch.
    ///     var example = new Github.BranchProtectionV3("example", new()
    ///     {
    ///         Repository = exampleGithubRepository.Name,
    ///         Branch = "main",
    ///         Restrictions = new Github.Inputs.BranchProtectionV3RestrictionsArgs
    ///         {
    ///             Users = new[]
    ///             {
    ///                 "foo-user",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// GitHub Branch Protection can be imported using an ID made up of `repository:branch`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import github:index/branchProtectionV3:BranchProtectionV3 terraform terraform:main
    /// ```
    /// </summary>
    [GithubResourceType("github:index/branchProtectionV3:BranchProtectionV3")]
    public partial class BranchProtectionV3 : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Git branch to protect.
        /// </summary>
        [Output("branch")]
        public Output<string> Branch { get; private set; } = null!;

        /// <summary>
        /// Boolean, setting this to `true` enforces status checks for repository administrators.
        /// </summary>
        [Output("enforceAdmins")]
        public Output<bool?> EnforceAdmins { get; private set; } = null!;

        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// The GitHub repository name.
        /// </summary>
        [Output("repository")]
        public Output<string> Repository { get; private set; } = null!;

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
        /// Enforce restrictions for pull request reviews. See Required Pull Request Reviews below for details.
        /// </summary>
        [Output("requiredPullRequestReviews")]
        public Output<Outputs.BranchProtectionV3RequiredPullRequestReviews?> RequiredPullRequestReviews { get; private set; } = null!;

        /// <summary>
        /// Enforce restrictions for required status checks. See Required Status Checks below for details.
        /// </summary>
        [Output("requiredStatusChecks")]
        public Output<Outputs.BranchProtectionV3RequiredStatusChecks?> RequiredStatusChecks { get; private set; } = null!;

        /// <summary>
        /// Enforce restrictions for the users and teams that may push to the branch. See Restrictions below for details.
        /// </summary>
        [Output("restrictions")]
        public Output<Outputs.BranchProtectionV3Restrictions?> Restrictions { get; private set; } = null!;


        /// <summary>
        /// Create a BranchProtectionV3 resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BranchProtectionV3(string name, BranchProtectionV3Args args, CustomResourceOptions? options = null)
            : base("github:index/branchProtectionV3:BranchProtectionV3", name, args ?? new BranchProtectionV3Args(), MakeResourceOptions(options, ""))
        {
        }

        private BranchProtectionV3(string name, Input<string> id, BranchProtectionV3State? state = null, CustomResourceOptions? options = null)
            : base("github:index/branchProtectionV3:BranchProtectionV3", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing BranchProtectionV3 resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BranchProtectionV3 Get(string name, Input<string> id, BranchProtectionV3State? state = null, CustomResourceOptions? options = null)
        {
            return new BranchProtectionV3(name, id, state, options);
        }
    }

    public sealed class BranchProtectionV3Args : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Git branch to protect.
        /// </summary>
        [Input("branch", required: true)]
        public Input<string> Branch { get; set; } = null!;

        /// <summary>
        /// Boolean, setting this to `true` enforces status checks for repository administrators.
        /// </summary>
        [Input("enforceAdmins")]
        public Input<bool>? EnforceAdmins { get; set; }

        /// <summary>
        /// The GitHub repository name.
        /// </summary>
        [Input("repository", required: true)]
        public Input<string> Repository { get; set; } = null!;

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
        /// Enforce restrictions for pull request reviews. See Required Pull Request Reviews below for details.
        /// </summary>
        [Input("requiredPullRequestReviews")]
        public Input<Inputs.BranchProtectionV3RequiredPullRequestReviewsArgs>? RequiredPullRequestReviews { get; set; }

        /// <summary>
        /// Enforce restrictions for required status checks. See Required Status Checks below for details.
        /// </summary>
        [Input("requiredStatusChecks")]
        public Input<Inputs.BranchProtectionV3RequiredStatusChecksArgs>? RequiredStatusChecks { get; set; }

        /// <summary>
        /// Enforce restrictions for the users and teams that may push to the branch. See Restrictions below for details.
        /// </summary>
        [Input("restrictions")]
        public Input<Inputs.BranchProtectionV3RestrictionsArgs>? Restrictions { get; set; }

        public BranchProtectionV3Args()
        {
        }
        public static new BranchProtectionV3Args Empty => new BranchProtectionV3Args();
    }

    public sealed class BranchProtectionV3State : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Git branch to protect.
        /// </summary>
        [Input("branch")]
        public Input<string>? Branch { get; set; }

        /// <summary>
        /// Boolean, setting this to `true` enforces status checks for repository administrators.
        /// </summary>
        [Input("enforceAdmins")]
        public Input<bool>? EnforceAdmins { get; set; }

        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// The GitHub repository name.
        /// </summary>
        [Input("repository")]
        public Input<string>? Repository { get; set; }

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
        /// Enforce restrictions for pull request reviews. See Required Pull Request Reviews below for details.
        /// </summary>
        [Input("requiredPullRequestReviews")]
        public Input<Inputs.BranchProtectionV3RequiredPullRequestReviewsGetArgs>? RequiredPullRequestReviews { get; set; }

        /// <summary>
        /// Enforce restrictions for required status checks. See Required Status Checks below for details.
        /// </summary>
        [Input("requiredStatusChecks")]
        public Input<Inputs.BranchProtectionV3RequiredStatusChecksGetArgs>? RequiredStatusChecks { get; set; }

        /// <summary>
        /// Enforce restrictions for the users and teams that may push to the branch. See Restrictions below for details.
        /// </summary>
        [Input("restrictions")]
        public Input<Inputs.BranchProtectionV3RestrictionsGetArgs>? Restrictions { get; set; }

        public BranchProtectionV3State()
        {
        }
        public static new BranchProtectionV3State Empty => new BranchProtectionV3State();
    }
}
