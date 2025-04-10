// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.github.RepositoryEnvironmentArgs;
import com.pulumi.github.Utilities;
import com.pulumi.github.inputs.RepositoryEnvironmentState;
import com.pulumi.github.outputs.RepositoryEnvironmentDeploymentBranchPolicy;
import com.pulumi.github.outputs.RepositoryEnvironmentReviewer;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * This resource allows you to create and manage environments for a GitHub repository.
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.github.GithubFunctions;
 * import com.pulumi.github.inputs.GetUserArgs;
 * import com.pulumi.github.Repository;
 * import com.pulumi.github.RepositoryArgs;
 * import com.pulumi.github.RepositoryEnvironment;
 * import com.pulumi.github.RepositoryEnvironmentArgs;
 * import com.pulumi.github.inputs.RepositoryEnvironmentReviewerArgs;
 * import com.pulumi.github.inputs.RepositoryEnvironmentDeploymentBranchPolicyArgs;
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
 *         final var current = GithubFunctions.getUser(GetUserArgs.builder()
 *             .username("")
 *             .build());
 * 
 *         var example = new Repository("example", RepositoryArgs.builder()
 *             .name("A Repository Project")
 *             .description("My awesome codebase")
 *             .build());
 * 
 *         var exampleRepositoryEnvironment = new RepositoryEnvironment("exampleRepositoryEnvironment", RepositoryEnvironmentArgs.builder()
 *             .environment("example")
 *             .repository(example.name())
 *             .preventSelfReview(true)
 *             .reviewers(RepositoryEnvironmentReviewerArgs.builder()
 *                 .users(current.id())
 *                 .build())
 *             .deploymentBranchPolicy(RepositoryEnvironmentDeploymentBranchPolicyArgs.builder()
 *                 .protectedBranches(true)
 *                 .customBranchPolicies(false)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * GitHub Repository Environment can be imported using an ID made up of `name` of the repository combined with the `environment` name of the environment, separated by a `:` character, e.g.
 * 
 * ```sh
 * $ pulumi import github:index/repositoryEnvironment:RepositoryEnvironment daily terraform:daily
 * ```
 * 
 */
@ResourceType(type="github:index/repositoryEnvironment:RepositoryEnvironment")
public class RepositoryEnvironment extends com.pulumi.resources.CustomResource {
    /**
     * Can repository admins bypass the environment protections.  Defaults to `true`.
     * 
     */
    @Export(name="canAdminsBypass", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> canAdminsBypass;

    /**
     * @return Can repository admins bypass the environment protections.  Defaults to `true`.
     * 
     */
    public Output<Optional<Boolean>> canAdminsBypass() {
        return Codegen.optional(this.canAdminsBypass);
    }
    /**
     * The deployment branch policy configuration
     * 
     */
    @Export(name="deploymentBranchPolicy", refs={RepositoryEnvironmentDeploymentBranchPolicy.class}, tree="[0]")
    private Output</* @Nullable */ RepositoryEnvironmentDeploymentBranchPolicy> deploymentBranchPolicy;

    /**
     * @return The deployment branch policy configuration
     * 
     */
    public Output<Optional<RepositoryEnvironmentDeploymentBranchPolicy>> deploymentBranchPolicy() {
        return Codegen.optional(this.deploymentBranchPolicy);
    }
    /**
     * The name of the environment.
     * 
     */
    @Export(name="environment", refs={String.class}, tree="[0]")
    private Output<String> environment;

    /**
     * @return The name of the environment.
     * 
     */
    public Output<String> environment() {
        return this.environment;
    }
    /**
     * Whether or not a user who created the job is prevented from approving their own job. Defaults to `false`.
     * 
     */
    @Export(name="preventSelfReview", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> preventSelfReview;

    /**
     * @return Whether or not a user who created the job is prevented from approving their own job. Defaults to `false`.
     * 
     */
    public Output<Optional<Boolean>> preventSelfReview() {
        return Codegen.optional(this.preventSelfReview);
    }
    /**
     * The repository of the environment.
     * 
     */
    @Export(name="repository", refs={String.class}, tree="[0]")
    private Output<String> repository;

    /**
     * @return The repository of the environment.
     * 
     */
    public Output<String> repository() {
        return this.repository;
    }
    /**
     * The environment reviewers configuration.
     * 
     */
    @Export(name="reviewers", refs={List.class,RepositoryEnvironmentReviewer.class}, tree="[0,1]")
    private Output</* @Nullable */ List<RepositoryEnvironmentReviewer>> reviewers;

    /**
     * @return The environment reviewers configuration.
     * 
     */
    public Output<Optional<List<RepositoryEnvironmentReviewer>>> reviewers() {
        return Codegen.optional(this.reviewers);
    }
    /**
     * Amount of time to delay a job after the job is initially triggered.
     * 
     */
    @Export(name="waitTimer", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> waitTimer;

    /**
     * @return Amount of time to delay a job after the job is initially triggered.
     * 
     */
    public Output<Optional<Integer>> waitTimer() {
        return Codegen.optional(this.waitTimer);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public RepositoryEnvironment(java.lang.String name) {
        this(name, RepositoryEnvironmentArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public RepositoryEnvironment(java.lang.String name, RepositoryEnvironmentArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public RepositoryEnvironment(java.lang.String name, RepositoryEnvironmentArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/repositoryEnvironment:RepositoryEnvironment", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private RepositoryEnvironment(java.lang.String name, Output<java.lang.String> id, @Nullable RepositoryEnvironmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/repositoryEnvironment:RepositoryEnvironment", name, state, makeResourceOptions(options, id), false);
    }

    private static RepositoryEnvironmentArgs makeArgs(RepositoryEnvironmentArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? RepositoryEnvironmentArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
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
    public static RepositoryEnvironment get(java.lang.String name, Output<java.lang.String> id, @Nullable RepositoryEnvironmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new RepositoryEnvironment(name, id, state, options);
    }
}
