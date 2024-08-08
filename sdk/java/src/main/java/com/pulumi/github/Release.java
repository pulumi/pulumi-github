// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.github.ReleaseArgs;
import com.pulumi.github.Utilities;
import com.pulumi.github.inputs.ReleaseState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * This resource allows you to create and manage a release in a specific
 * GitHub repository.
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
 * import com.pulumi.github.Repository;
 * import com.pulumi.github.RepositoryArgs;
 * import com.pulumi.github.Release;
 * import com.pulumi.github.ReleaseArgs;
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
 *         var repo = new Repository("repo", RepositoryArgs.builder()
 *             .name("repo")
 *             .description("GitHub repo managed by Terraform")
 *             .private_(false)
 *             .build());
 * 
 *         var example = new Release("example", ReleaseArgs.builder()
 *             .repository(repo.name())
 *             .tagName("v1.0.0")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ### On Non-Default Branch
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.github.Repository;
 * import com.pulumi.github.RepositoryArgs;
 * import com.pulumi.github.Branch;
 * import com.pulumi.github.BranchArgs;
 * import com.pulumi.github.Release;
 * import com.pulumi.github.ReleaseArgs;
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
 *         var example = new Repository("example", RepositoryArgs.builder()
 *             .name("repo")
 *             .autoInit(true)
 *             .build());
 * 
 *         var exampleBranch = new Branch("exampleBranch", BranchArgs.builder()
 *             .repository(example.name())
 *             .branch("branch_name")
 *             .sourceBranch(example.defaultBranch())
 *             .build());
 * 
 *         var exampleRelease = new Release("exampleRelease", ReleaseArgs.builder()
 *             .repository(example.name())
 *             .tagName("v1.0.0")
 *             .targetCommitish(exampleBranch.branch())
 *             .draft(false)
 *             .prerelease(false)
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
 * This resource can be imported using the `name` of the repository, combined with the `id` of the release, and a `:` character for separating components, e.g.
 * 
 * ```sh
 * $ pulumi import github:index/release:Release example repo:12345678
 * ```
 * 
 */
@ResourceType(type="github:index/release:Release")
public class Release extends com.pulumi.resources.CustomResource {
    /**
     * URL that can be provided to API calls displaying the attached assets to this release.
     * 
     */
    @Export(name="assetsUrl", refs={String.class}, tree="[0]")
    private Output<String> assetsUrl;

    /**
     * @return URL that can be provided to API calls displaying the attached assets to this release.
     * 
     */
    public Output<String> assetsUrl() {
        return this.assetsUrl;
    }
    /**
     * Text describing the contents of the tag.
     * 
     */
    @Export(name="body", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> body;

    /**
     * @return Text describing the contents of the tag.
     * 
     */
    public Output<Optional<String>> body() {
        return Codegen.optional(this.body);
    }
    /**
     * This is the date of the commit used for the release, and not the date when the release was drafted or published.
     * 
     */
    @Export(name="createdAt", refs={String.class}, tree="[0]")
    private Output<String> createdAt;

    /**
     * @return This is the date of the commit used for the release, and not the date when the release was drafted or published.
     * 
     */
    public Output<String> createdAt() {
        return this.createdAt;
    }
    /**
     * If specified, a discussion of the specified category is created and linked to the release. The value must be a category that already exists in the repository. For more information, see [Managing categories for discussions in your repository](https://docs.github.com/discussions/managing-discussions-for-your-community/managing-categories-for-discussions-in-your-repository).
     * 
     */
    @Export(name="discussionCategoryName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> discussionCategoryName;

    /**
     * @return If specified, a discussion of the specified category is created and linked to the release. The value must be a category that already exists in the repository. For more information, see [Managing categories for discussions in your repository](https://docs.github.com/discussions/managing-discussions-for-your-community/managing-categories-for-discussions-in-your-repository).
     * 
     */
    public Output<Optional<String>> discussionCategoryName() {
        return Codegen.optional(this.discussionCategoryName);
    }
    /**
     * Set to `false` to create a published release.
     * 
     */
    @Export(name="draft", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> draft;

    /**
     * @return Set to `false` to create a published release.
     * 
     */
    public Output<Optional<Boolean>> draft() {
        return Codegen.optional(this.draft);
    }
    @Export(name="etag", refs={String.class}, tree="[0]")
    private Output<String> etag;

    public Output<String> etag() {
        return this.etag;
    }
    /**
     * Set to `true` to automatically generate the name and body for this release. If `name` is specified, the specified `name` will be used; otherwise, a name will be automatically generated. If `body` is specified, the `body` will be pre-pended to the automatically generated notes.
     * 
     */
    @Export(name="generateReleaseNotes", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> generateReleaseNotes;

    /**
     * @return Set to `true` to automatically generate the name and body for this release. If `name` is specified, the specified `name` will be used; otherwise, a name will be automatically generated. If `body` is specified, the `body` will be pre-pended to the automatically generated notes.
     * 
     */
    public Output<Optional<Boolean>> generateReleaseNotes() {
        return Codegen.optional(this.generateReleaseNotes);
    }
    /**
     * URL of the release in GitHub.
     * 
     */
    @Export(name="htmlUrl", refs={String.class}, tree="[0]")
    private Output<String> htmlUrl;

    /**
     * @return URL of the release in GitHub.
     * 
     */
    public Output<String> htmlUrl() {
        return this.htmlUrl;
    }
    /**
     * The name of the release.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the release.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * GraphQL global node id for use with v4 API
     * 
     */
    @Export(name="nodeId", refs={String.class}, tree="[0]")
    private Output<String> nodeId;

    /**
     * @return GraphQL global node id for use with v4 API
     * 
     */
    public Output<String> nodeId() {
        return this.nodeId;
    }
    /**
     * Set to `false` to identify the release as a full release.
     * 
     */
    @Export(name="prerelease", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> prerelease;

    /**
     * @return Set to `false` to identify the release as a full release.
     * 
     */
    public Output<Optional<Boolean>> prerelease() {
        return Codegen.optional(this.prerelease);
    }
    /**
     * This is the date when the release was published. This will be empty if the release is a draft.
     * 
     */
    @Export(name="publishedAt", refs={String.class}, tree="[0]")
    private Output<String> publishedAt;

    /**
     * @return This is the date when the release was published. This will be empty if the release is a draft.
     * 
     */
    public Output<String> publishedAt() {
        return this.publishedAt;
    }
    /**
     * The ID of the release.
     * 
     */
    @Export(name="releaseId", refs={Integer.class}, tree="[0]")
    private Output<Integer> releaseId;

    /**
     * @return The ID of the release.
     * 
     */
    public Output<Integer> releaseId() {
        return this.releaseId;
    }
    /**
     * The name of the repository.
     * 
     */
    @Export(name="repository", refs={String.class}, tree="[0]")
    private Output<String> repository;

    /**
     * @return The name of the repository.
     * 
     */
    public Output<String> repository() {
        return this.repository;
    }
    /**
     * The name of the tag.
     * 
     */
    @Export(name="tagName", refs={String.class}, tree="[0]")
    private Output<String> tagName;

    /**
     * @return The name of the tag.
     * 
     */
    public Output<String> tagName() {
        return this.tagName;
    }
    /**
     * URL that can be provided to API calls to fetch the release TAR archive.
     * 
     */
    @Export(name="tarballUrl", refs={String.class}, tree="[0]")
    private Output<String> tarballUrl;

    /**
     * @return URL that can be provided to API calls to fetch the release TAR archive.
     * 
     */
    public Output<String> tarballUrl() {
        return this.tarballUrl;
    }
    /**
     * The branch name or commit SHA the tag is created from. Defaults to the default branch of the repository.
     * 
     */
    @Export(name="targetCommitish", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> targetCommitish;

    /**
     * @return The branch name or commit SHA the tag is created from. Defaults to the default branch of the repository.
     * 
     */
    public Output<Optional<String>> targetCommitish() {
        return Codegen.optional(this.targetCommitish);
    }
    /**
     * URL that can be provided to API calls to upload assets.
     * 
     */
    @Export(name="uploadUrl", refs={String.class}, tree="[0]")
    private Output<String> uploadUrl;

    /**
     * @return URL that can be provided to API calls to upload assets.
     * 
     */
    public Output<String> uploadUrl() {
        return this.uploadUrl;
    }
    /**
     * URL that can be provided to API calls that reference this release.
     * 
     */
    @Export(name="url", refs={String.class}, tree="[0]")
    private Output<String> url;

    /**
     * @return URL that can be provided to API calls that reference this release.
     * 
     */
    public Output<String> url() {
        return this.url;
    }
    /**
     * URL that can be provided to API calls to fetch the release ZIP archive.
     * 
     */
    @Export(name="zipballUrl", refs={String.class}, tree="[0]")
    private Output<String> zipballUrl;

    /**
     * @return URL that can be provided to API calls to fetch the release ZIP archive.
     * 
     */
    public Output<String> zipballUrl() {
        return this.zipballUrl;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Release(java.lang.String name) {
        this(name, ReleaseArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Release(java.lang.String name, ReleaseArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Release(java.lang.String name, ReleaseArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/release:Release", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private Release(java.lang.String name, Output<java.lang.String> id, @Nullable ReleaseState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/release:Release", name, state, makeResourceOptions(options, id), false);
    }

    private static ReleaseArgs makeArgs(ReleaseArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ReleaseArgs.Empty : args;
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
    public static Release get(java.lang.String name, Output<java.lang.String> id, @Nullable ReleaseState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Release(name, id, state, options);
    }
}
