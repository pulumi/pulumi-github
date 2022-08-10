// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.github.BranchProtectionArgs;
import com.pulumi.github.Utilities;
import com.pulumi.github.inputs.BranchProtectionState;
import com.pulumi.github.outputs.BranchProtectionRequiredPullRequestReview;
import com.pulumi.github.outputs.BranchProtectionRequiredStatusCheck;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Protects a GitHub branch.
 * 
 * This resource allows you to configure branch protection for repositories in your organization. When applied, the branch will be protected from forced pushes and deletion. Additional constraints, such as required status checks or restrictions on users, teams, and apps, can also be configured.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.github.Repository;
 * import com.pulumi.github.GithubFunctions;
 * import com.pulumi.github.inputs.GetUserArgs;
 * import com.pulumi.github.Team;
 * import com.pulumi.github.BranchProtection;
 * import com.pulumi.github.BranchProtectionArgs;
 * import com.pulumi.github.inputs.BranchProtectionRequiredStatusCheckArgs;
 * import com.pulumi.github.inputs.BranchProtectionRequiredPullRequestReviewArgs;
 * import com.pulumi.github.TeamRepository;
 * import com.pulumi.github.TeamRepositoryArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var exampleRepository = new Repository(&#34;exampleRepository&#34;);
 * 
 *         final var exampleUser = GithubFunctions.getUser(GetUserArgs.builder()
 *             .username(&#34;example&#34;)
 *             .build());
 * 
 *         var exampleTeam = new Team(&#34;exampleTeam&#34;);
 * 
 *         var exampleBranchProtection = new BranchProtection(&#34;exampleBranchProtection&#34;, BranchProtectionArgs.builder()        
 *             .repositoryId(exampleRepository.nodeId())
 *             .pattern(&#34;main&#34;)
 *             .enforceAdmins(true)
 *             .allowsDeletions(true)
 *             .requiredStatusChecks(BranchProtectionRequiredStatusCheckArgs.builder()
 *                 .strict(false)
 *                 .contexts(&#34;ci/travis&#34;)
 *                 .build())
 *             .requiredPullRequestReviews(BranchProtectionRequiredPullRequestReviewArgs.builder()
 *                 .dismissStaleReviews(true)
 *                 .restrictDismissals(true)
 *                 .dismissalRestrictions(                
 *                     exampleUser.applyValue(getUserResult -&gt; getUserResult.nodeId()),
 *                     exampleTeam.nodeId())
 *                 .build())
 *             .pushRestrictions(exampleUser.applyValue(getUserResult -&gt; getUserResult.nodeId()))
 *             .build());
 * 
 *         var exampleTeamRepository = new TeamRepository(&#34;exampleTeamRepository&#34;, TeamRepositoryArgs.builder()        
 *             .teamId(exampleTeam.id())
 *             .repository(exampleRepository.name())
 *             .permission(&#34;pull&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * GitHub Branch Protection can be imported using an ID made up of `repository:pattern`, e.g.
 * 
 * ```sh
 *  $ pulumi import github:index/branchProtection:BranchProtection terraform terraform:main
 * ```
 * 
 */
@ResourceType(type="github:index/branchProtection:BranchProtection")
public class BranchProtection extends com.pulumi.resources.CustomResource {
    /**
     * Boolean, setting this to `true` to allow the branch to be deleted.
     * 
     */
    @Export(name="allowsDeletions", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> allowsDeletions;

    /**
     * @return Boolean, setting this to `true` to allow the branch to be deleted.
     * 
     */
    public Output<Optional<Boolean>> allowsDeletions() {
        return Codegen.optional(this.allowsDeletions);
    }
    /**
     * Boolean, setting this to `true` to allow force pushes on the branch.
     * 
     */
    @Export(name="allowsForcePushes", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> allowsForcePushes;

    /**
     * @return Boolean, setting this to `true` to allow force pushes on the branch.
     * 
     */
    public Output<Optional<Boolean>> allowsForcePushes() {
        return Codegen.optional(this.allowsForcePushes);
    }
    /**
     * Boolean, setting this to `true` to block creating the branch.
     * 
     */
    @Export(name="blocksCreations", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> blocksCreations;

    /**
     * @return Boolean, setting this to `true` to block creating the branch.
     * 
     */
    public Output<Optional<Boolean>> blocksCreations() {
        return Codegen.optional(this.blocksCreations);
    }
    /**
     * Boolean, setting this to `true` enforces status checks for repository administrators.
     * 
     */
    @Export(name="enforceAdmins", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> enforceAdmins;

    /**
     * @return Boolean, setting this to `true` enforces status checks for repository administrators.
     * 
     */
    public Output<Optional<Boolean>> enforceAdmins() {
        return Codegen.optional(this.enforceAdmins);
    }
    /**
     * Identifies the protection rule pattern.
     * 
     */
    @Export(name="pattern", type=String.class, parameters={})
    private Output<String> pattern;

    /**
     * @return Identifies the protection rule pattern.
     * 
     */
    public Output<String> pattern() {
        return this.pattern;
    }
    /**
     * The list of actor IDs that may push to the branch.
     * 
     */
    @Export(name="pushRestrictions", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> pushRestrictions;

    /**
     * @return The list of actor IDs that may push to the branch.
     * 
     */
    public Output<Optional<List<String>>> pushRestrictions() {
        return Codegen.optional(this.pushRestrictions);
    }
    /**
     * The name or node ID of the repository associated with this branch protection rule.
     * 
     */
    @Export(name="repositoryId", type=String.class, parameters={})
    private Output<String> repositoryId;

    /**
     * @return The name or node ID of the repository associated with this branch protection rule.
     * 
     */
    public Output<String> repositoryId() {
        return this.repositoryId;
    }
    /**
     * Boolean, setting this to `true` requires all conversations on code must be resolved before a pull request can be merged.
     * 
     */
    @Export(name="requireConversationResolution", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> requireConversationResolution;

    /**
     * @return Boolean, setting this to `true` requires all conversations on code must be resolved before a pull request can be merged.
     * 
     */
    public Output<Optional<Boolean>> requireConversationResolution() {
        return Codegen.optional(this.requireConversationResolution);
    }
    /**
     * Boolean, setting this to `true` requires all commits to be signed with GPG.
     * 
     */
    @Export(name="requireSignedCommits", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> requireSignedCommits;

    /**
     * @return Boolean, setting this to `true` requires all commits to be signed with GPG.
     * 
     */
    public Output<Optional<Boolean>> requireSignedCommits() {
        return Codegen.optional(this.requireSignedCommits);
    }
    /**
     * Boolean, setting this to `true` enforces a linear commit Git history, which prevents anyone from pushing merge commits to a branch
     * 
     */
    @Export(name="requiredLinearHistory", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> requiredLinearHistory;

    /**
     * @return Boolean, setting this to `true` enforces a linear commit Git history, which prevents anyone from pushing merge commits to a branch
     * 
     */
    public Output<Optional<Boolean>> requiredLinearHistory() {
        return Codegen.optional(this.requiredLinearHistory);
    }
    /**
     * Enforce restrictions for pull request reviews. See Required Pull Request Reviews below for details.
     * 
     */
    @Export(name="requiredPullRequestReviews", type=List.class, parameters={BranchProtectionRequiredPullRequestReview.class})
    private Output</* @Nullable */ List<BranchProtectionRequiredPullRequestReview>> requiredPullRequestReviews;

    /**
     * @return Enforce restrictions for pull request reviews. See Required Pull Request Reviews below for details.
     * 
     */
    public Output<Optional<List<BranchProtectionRequiredPullRequestReview>>> requiredPullRequestReviews() {
        return Codegen.optional(this.requiredPullRequestReviews);
    }
    /**
     * Enforce restrictions for required status checks. See Required Status Checks below for details.
     * 
     */
    @Export(name="requiredStatusChecks", type=List.class, parameters={BranchProtectionRequiredStatusCheck.class})
    private Output</* @Nullable */ List<BranchProtectionRequiredStatusCheck>> requiredStatusChecks;

    /**
     * @return Enforce restrictions for required status checks. See Required Status Checks below for details.
     * 
     */
    public Output<Optional<List<BranchProtectionRequiredStatusCheck>>> requiredStatusChecks() {
        return Codegen.optional(this.requiredStatusChecks);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public BranchProtection(String name) {
        this(name, BranchProtectionArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public BranchProtection(String name, BranchProtectionArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public BranchProtection(String name, BranchProtectionArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/branchProtection:BranchProtection", name, args == null ? BranchProtectionArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private BranchProtection(String name, Output<String> id, @Nullable BranchProtectionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/branchProtection:BranchProtection", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static BranchProtection get(String name, Output<String> id, @Nullable BranchProtectionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new BranchProtection(name, id, state, options);
    }
}