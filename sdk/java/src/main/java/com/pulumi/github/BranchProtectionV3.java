// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.github.BranchProtectionV3Args;
import com.pulumi.github.Utilities;
import com.pulumi.github.inputs.BranchProtectionV3State;
import com.pulumi.github.outputs.BranchProtectionV3RequiredPullRequestReviews;
import com.pulumi.github.outputs.BranchProtectionV3RequiredStatusChecks;
import com.pulumi.github.outputs.BranchProtectionV3Restrictions;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Protects a GitHub branch.
 * 
 * The `github.BranchProtection` resource has moved to the GraphQL API, while this resource will continue to leverage the REST API.
 * 
 * This resource allows you to configure branch protection for repositories in your organization. When applied, the branch will be protected from forced pushes and deletion. Additional constraints, such as required status checks or restrictions on users, teams, and apps, can also be configured.
 * 
 * ## Import
 * 
 * GitHub Branch Protection can be imported using an ID made up of `repository:branch`, e.g.
 * 
 * ```sh
 *  $ pulumi import github:index/branchProtectionV3:BranchProtectionV3 terraform terraform:main
 * ```
 * 
 */
@ResourceType(type="github:index/branchProtectionV3:BranchProtectionV3")
public class BranchProtectionV3 extends com.pulumi.resources.CustomResource {
    /**
     * The Git branch to protect.
     * 
     */
    @Export(name="branch", refs={String.class}, tree="[0]")
    private Output<String> branch;

    /**
     * @return The Git branch to protect.
     * 
     */
    public Output<String> branch() {
        return this.branch;
    }
    /**
     * Boolean, setting this to `true` enforces status checks for repository administrators.
     * 
     */
    @Export(name="enforceAdmins", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> enforceAdmins;

    /**
     * @return Boolean, setting this to `true` enforces status checks for repository administrators.
     * 
     */
    public Output<Optional<Boolean>> enforceAdmins() {
        return Codegen.optional(this.enforceAdmins);
    }
    @Export(name="etag", refs={String.class}, tree="[0]")
    private Output<String> etag;

    public Output<String> etag() {
        return this.etag;
    }
    /**
     * The GitHub repository name.
     * 
     */
    @Export(name="repository", refs={String.class}, tree="[0]")
    private Output<String> repository;

    /**
     * @return The GitHub repository name.
     * 
     */
    public Output<String> repository() {
        return this.repository;
    }
    /**
     * Boolean, setting this to `true` requires all conversations on code must be resolved before a pull request can be merged.
     * 
     */
    @Export(name="requireConversationResolution", refs={Boolean.class}, tree="[0]")
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
    @Export(name="requireSignedCommits", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> requireSignedCommits;

    /**
     * @return Boolean, setting this to `true` requires all commits to be signed with GPG.
     * 
     */
    public Output<Optional<Boolean>> requireSignedCommits() {
        return Codegen.optional(this.requireSignedCommits);
    }
    /**
     * Enforce restrictions for pull request reviews. See Required Pull Request Reviews below for details.
     * 
     */
    @Export(name="requiredPullRequestReviews", refs={BranchProtectionV3RequiredPullRequestReviews.class}, tree="[0]")
    private Output</* @Nullable */ BranchProtectionV3RequiredPullRequestReviews> requiredPullRequestReviews;

    /**
     * @return Enforce restrictions for pull request reviews. See Required Pull Request Reviews below for details.
     * 
     */
    public Output<Optional<BranchProtectionV3RequiredPullRequestReviews>> requiredPullRequestReviews() {
        return Codegen.optional(this.requiredPullRequestReviews);
    }
    /**
     * Enforce restrictions for required status checks. See Required Status Checks below for details.
     * 
     */
    @Export(name="requiredStatusChecks", refs={BranchProtectionV3RequiredStatusChecks.class}, tree="[0]")
    private Output</* @Nullable */ BranchProtectionV3RequiredStatusChecks> requiredStatusChecks;

    /**
     * @return Enforce restrictions for required status checks. See Required Status Checks below for details.
     * 
     */
    public Output<Optional<BranchProtectionV3RequiredStatusChecks>> requiredStatusChecks() {
        return Codegen.optional(this.requiredStatusChecks);
    }
    /**
     * Enforce restrictions for the users and teams that may push to the branch. See Restrictions below for details.
     * 
     */
    @Export(name="restrictions", refs={BranchProtectionV3Restrictions.class}, tree="[0]")
    private Output</* @Nullable */ BranchProtectionV3Restrictions> restrictions;

    /**
     * @return Enforce restrictions for the users and teams that may push to the branch. See Restrictions below for details.
     * 
     */
    public Output<Optional<BranchProtectionV3Restrictions>> restrictions() {
        return Codegen.optional(this.restrictions);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public BranchProtectionV3(String name) {
        this(name, BranchProtectionV3Args.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public BranchProtectionV3(String name, BranchProtectionV3Args args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public BranchProtectionV3(String name, BranchProtectionV3Args args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/branchProtectionV3:BranchProtectionV3", name, args == null ? BranchProtectionV3Args.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private BranchProtectionV3(String name, Output<String> id, @Nullable BranchProtectionV3State state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/branchProtectionV3:BranchProtectionV3", name, state, makeResourceOptions(options, id));
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
    public static BranchProtectionV3 get(String name, Output<String> id, @Nullable BranchProtectionV3State state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new BranchProtectionV3(name, id, state, options);
    }
}
