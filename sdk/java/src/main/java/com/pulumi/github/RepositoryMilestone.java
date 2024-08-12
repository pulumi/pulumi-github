// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.github.RepositoryMilestoneArgs;
import com.pulumi.github.Utilities;
import com.pulumi.github.inputs.RepositoryMilestoneState;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a GitHub repository milestone resource.
 * 
 * This resource allows you to create and manage milestones for a GitHub Repository within an organization or user account.
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
 * import com.pulumi.github.RepositoryMilestone;
 * import com.pulumi.github.RepositoryMilestoneArgs;
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
 *         // Create a milestone for a repository
 *         var example = new RepositoryMilestone("example", RepositoryMilestoneArgs.builder()
 *             .owner("example-owner")
 *             .repository("example-repository")
 *             .title("v1.1.0")
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
 * A GitHub Repository Milestone can be imported using an ID made up of `owner/repository/number`, e.g.
 * 
 * ```sh
 * $ pulumi import github:index/repositoryMilestone:RepositoryMilestone example example-owner/example-repository/1
 * ```
 * 
 */
@ResourceType(type="github:index/repositoryMilestone:RepositoryMilestone")
public class RepositoryMilestone extends com.pulumi.resources.CustomResource {
    /**
     * A description of the milestone.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return A description of the milestone.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The milestone due date. In `yyyy-mm-dd` format.
     * 
     */
    @Export(name="dueDate", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> dueDate;

    /**
     * @return The milestone due date. In `yyyy-mm-dd` format.
     * 
     */
    public Output<Optional<String>> dueDate() {
        return Codegen.optional(this.dueDate);
    }
    /**
     * The number of the milestone.
     * 
     */
    @Export(name="number", refs={Integer.class}, tree="[0]")
    private Output<Integer> number;

    /**
     * @return The number of the milestone.
     * 
     */
    public Output<Integer> number() {
        return this.number;
    }
    /**
     * The owner of the GitHub Repository.
     * 
     */
    @Export(name="owner", refs={String.class}, tree="[0]")
    private Output<String> owner;

    /**
     * @return The owner of the GitHub Repository.
     * 
     */
    public Output<String> owner() {
        return this.owner;
    }
    /**
     * The name of the GitHub Repository.
     * 
     */
    @Export(name="repository", refs={String.class}, tree="[0]")
    private Output<String> repository;

    /**
     * @return The name of the GitHub Repository.
     * 
     */
    public Output<String> repository() {
        return this.repository;
    }
    /**
     * The state of the milestone. Either `open` or `closed`. Default: `open`
     * 
     */
    @Export(name="state", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> state;

    /**
     * @return The state of the milestone. Either `open` or `closed`. Default: `open`
     * 
     */
    public Output<Optional<String>> state() {
        return Codegen.optional(this.state);
    }
    /**
     * The title of the milestone.
     * 
     */
    @Export(name="title", refs={String.class}, tree="[0]")
    private Output<String> title;

    /**
     * @return The title of the milestone.
     * 
     */
    public Output<String> title() {
        return this.title;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public RepositoryMilestone(java.lang.String name) {
        this(name, RepositoryMilestoneArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public RepositoryMilestone(java.lang.String name, RepositoryMilestoneArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public RepositoryMilestone(java.lang.String name, RepositoryMilestoneArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/repositoryMilestone:RepositoryMilestone", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private RepositoryMilestone(java.lang.String name, Output<java.lang.String> id, @Nullable RepositoryMilestoneState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/repositoryMilestone:RepositoryMilestone", name, state, makeResourceOptions(options, id), false);
    }

    private static RepositoryMilestoneArgs makeArgs(RepositoryMilestoneArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? RepositoryMilestoneArgs.Empty : args;
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
    public static RepositoryMilestone get(java.lang.String name, Output<java.lang.String> id, @Nullable RepositoryMilestoneState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new RepositoryMilestone(name, id, state, options);
    }
}
