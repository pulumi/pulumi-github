// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.github.BranchArgs;
import com.pulumi.github.Utilities;
import com.pulumi.github.inputs.BranchState;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * This resource allows you to create and manage branches within your repository.
 * 
 * Additional constraints can be applied to ensure your branch is created from
 * another branch or commit.
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
 * import com.pulumi.github.Branch;
 * import com.pulumi.github.BranchArgs;
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
 *         var development = new Branch("development", BranchArgs.builder()
 *             .repository("example")
 *             .branch("development")
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
 * GitHub Branch can be imported using an ID made up of `repository:branch`, e.g.
 * 
 * ```sh
 * $ pulumi import github:index/branch:Branch terraform terraform:main
 * ```
 * Importing github branch into an instance object (when using a for each block to manage multiple branches)
 * 
 * ```sh
 * $ pulumi import github:index/branch:Branch terraform[&#34;terraform&#34;] terraform:main
 * ```
 * Optionally, a source branch may be specified using an ID of `repository:branch:source_branch`.
 * This is useful for importing branches that do not branch directly off main.
 * 
 * ```sh
 * $ pulumi import github:index/branch:Branch terraform terraform:feature-branch:dev
 * ```
 * 
 */
@ResourceType(type="github:index/branch:Branch")
public class Branch extends com.pulumi.resources.CustomResource {
    /**
     * The repository branch to create.
     * 
     */
    @Export(name="branch", refs={String.class}, tree="[0]")
    private Output<String> branch;

    /**
     * @return The repository branch to create.
     * 
     */
    public Output<String> branch() {
        return this.branch;
    }
    /**
     * An etag representing the Branch object.
     * 
     */
    @Export(name="etag", refs={String.class}, tree="[0]")
    private Output<String> etag;

    /**
     * @return An etag representing the Branch object.
     * 
     */
    public Output<String> etag() {
        return this.etag;
    }
    /**
     * A string representing a branch reference, in the form of `refs/heads/&lt;branch&gt;`.
     * 
     */
    @Export(name="ref", refs={String.class}, tree="[0]")
    private Output<String> ref;

    /**
     * @return A string representing a branch reference, in the form of `refs/heads/&lt;branch&gt;`.
     * 
     */
    public Output<String> ref() {
        return this.ref;
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
     * A string storing the reference&#39;s `HEAD` commit&#39;s SHA1.
     * 
     */
    @Export(name="sha", refs={String.class}, tree="[0]")
    private Output<String> sha;

    /**
     * @return A string storing the reference&#39;s `HEAD` commit&#39;s SHA1.
     * 
     */
    public Output<String> sha() {
        return this.sha;
    }
    /**
     * The branch name to start from. Defaults to `main`.
     * 
     */
    @Export(name="sourceBranch", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> sourceBranch;

    /**
     * @return The branch name to start from. Defaults to `main`.
     * 
     */
    public Output<Optional<String>> sourceBranch() {
        return Codegen.optional(this.sourceBranch);
    }
    /**
     * The commit hash to start from. Defaults to the tip of `source_branch`. If provided, `source_branch` is ignored.
     * 
     */
    @Export(name="sourceSha", refs={String.class}, tree="[0]")
    private Output<String> sourceSha;

    /**
     * @return The commit hash to start from. Defaults to the tip of `source_branch`. If provided, `source_branch` is ignored.
     * 
     */
    public Output<String> sourceSha() {
        return this.sourceSha;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Branch(java.lang.String name) {
        this(name, BranchArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Branch(java.lang.String name, BranchArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Branch(java.lang.String name, BranchArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/branch:Branch", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private Branch(java.lang.String name, Output<java.lang.String> id, @Nullable BranchState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/branch:Branch", name, state, makeResourceOptions(options, id), false);
    }

    private static BranchArgs makeArgs(BranchArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? BranchArgs.Empty : args;
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
    public static Branch get(java.lang.String name, Output<java.lang.String> id, @Nullable BranchState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Branch(name, id, state, options);
    }
}
