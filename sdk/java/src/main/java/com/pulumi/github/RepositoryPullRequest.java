// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.github.RepositoryPullRequestArgs;
import com.pulumi.github.Utilities;
import com.pulumi.github.inputs.RepositoryPullRequestState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * This resource allows you to create and manage PullRequests for repositories within your GitHub organization or personal account.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.github.RepositoryPullRequest;
 * import com.pulumi.github.RepositoryPullRequestArgs;
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
 *         var example = new RepositoryPullRequest(&#34;example&#34;, RepositoryPullRequestArgs.builder()        
 *             .baseRef(&#34;main&#34;)
 *             .baseRepository(&#34;example-repository&#34;)
 *             .body(&#34;This will change everything&#34;)
 *             .headRef(&#34;feature-branch&#34;)
 *             .title(&#34;My newest feature&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 */
@ResourceType(type="github:index/repositoryPullRequest:RepositoryPullRequest")
public class RepositoryPullRequest extends com.pulumi.resources.CustomResource {
    /**
     * Name of the branch serving as the base of the Pull Request.
     * 
     */
    @Export(name="baseRef", type=String.class, parameters={})
    private Output<String> baseRef;

    /**
     * @return Name of the branch serving as the base of the Pull Request.
     * 
     */
    public Output<String> baseRef() {
        return this.baseRef;
    }
    /**
     * Name of the base repository to retrieve the Pull Requests from.
     * 
     */
    @Export(name="baseRepository", type=String.class, parameters={})
    private Output<String> baseRepository;

    /**
     * @return Name of the base repository to retrieve the Pull Requests from.
     * 
     */
    public Output<String> baseRepository() {
        return this.baseRepository;
    }
    /**
     * Head commit SHA of the Pull Request base.
     * 
     */
    @Export(name="baseSha", type=String.class, parameters={})
    private Output<String> baseSha;

    /**
     * @return Head commit SHA of the Pull Request base.
     * 
     */
    public Output<String> baseSha() {
        return this.baseSha;
    }
    /**
     * Body of the Pull Request.
     * 
     */
    @Export(name="body", type=String.class, parameters={})
    private Output</* @Nullable */ String> body;

    /**
     * @return Body of the Pull Request.
     * 
     */
    public Output<Optional<String>> body() {
        return Codegen.optional(this.body);
    }
    /**
     * Indicates Whether this Pull Request is a draft.
     * 
     */
    @Export(name="draft", type=Boolean.class, parameters={})
    private Output<Boolean> draft;

    /**
     * @return Indicates Whether this Pull Request is a draft.
     * 
     */
    public Output<Boolean> draft() {
        return this.draft;
    }
    /**
     * Name of the branch serving as the head of the Pull Request.
     * 
     */
    @Export(name="headRef", type=String.class, parameters={})
    private Output<String> headRef;

    /**
     * @return Name of the branch serving as the head of the Pull Request.
     * 
     */
    public Output<String> headRef() {
        return this.headRef;
    }
    /**
     * Head commit SHA of the Pull Request head.
     * 
     */
    @Export(name="headSha", type=String.class, parameters={})
    private Output<String> headSha;

    /**
     * @return Head commit SHA of the Pull Request head.
     * 
     */
    public Output<String> headSha() {
        return this.headSha;
    }
    /**
     * List of label names set on the Pull Request.
     * 
     */
    @Export(name="labels", type=List.class, parameters={String.class})
    private Output<List<String>> labels;

    /**
     * @return List of label names set on the Pull Request.
     * 
     */
    public Output<List<String>> labels() {
        return this.labels;
    }
    /**
     * Controls whether the base repository maintainers can modify the Pull Request. Default: false.
     * 
     */
    @Export(name="maintainerCanModify", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> maintainerCanModify;

    /**
     * @return Controls whether the base repository maintainers can modify the Pull Request. Default: false.
     * 
     */
    public Output<Optional<Boolean>> maintainerCanModify() {
        return Codegen.optional(this.maintainerCanModify);
    }
    /**
     * The number of the Pull Request within the repository.
     * 
     */
    @Export(name="number", type=Integer.class, parameters={})
    private Output<Integer> number;

    /**
     * @return The number of the Pull Request within the repository.
     * 
     */
    public Output<Integer> number() {
        return this.number;
    }
    /**
     * Unix timestamp indicating the Pull Request creation time.
     * 
     */
    @Export(name="openedAt", type=Integer.class, parameters={})
    private Output<Integer> openedAt;

    /**
     * @return Unix timestamp indicating the Pull Request creation time.
     * 
     */
    public Output<Integer> openedAt() {
        return this.openedAt;
    }
    /**
     * GitHub login of the user who opened the Pull Request.
     * 
     */
    @Export(name="openedBy", type=String.class, parameters={})
    private Output<String> openedBy;

    /**
     * @return GitHub login of the user who opened the Pull Request.
     * 
     */
    public Output<String> openedBy() {
        return this.openedBy;
    }
    /**
     * Owner of the repository. If not provided, the provider&#39;s default owner is used.
     * 
     */
    @Export(name="owner", type=String.class, parameters={})
    private Output</* @Nullable */ String> owner;

    /**
     * @return Owner of the repository. If not provided, the provider&#39;s default owner is used.
     * 
     */
    public Output<Optional<String>> owner() {
        return Codegen.optional(this.owner);
    }
    /**
     * the current Pull Request state - can be &#34;open&#34;, &#34;closed&#34; or &#34;merged&#34;.
     * 
     */
    @Export(name="state", type=String.class, parameters={})
    private Output<String> state;

    /**
     * @return the current Pull Request state - can be &#34;open&#34;, &#34;closed&#34; or &#34;merged&#34;.
     * 
     */
    public Output<String> state() {
        return this.state;
    }
    /**
     * The title of the Pull Request.
     * 
     */
    @Export(name="title", type=String.class, parameters={})
    private Output<String> title;

    /**
     * @return The title of the Pull Request.
     * 
     */
    public Output<String> title() {
        return this.title;
    }
    /**
     * The timestamp of the last Pull Request update.
     * 
     */
    @Export(name="updatedAt", type=Integer.class, parameters={})
    private Output<Integer> updatedAt;

    /**
     * @return The timestamp of the last Pull Request update.
     * 
     */
    public Output<Integer> updatedAt() {
        return this.updatedAt;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public RepositoryPullRequest(String name) {
        this(name, RepositoryPullRequestArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public RepositoryPullRequest(String name, RepositoryPullRequestArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public RepositoryPullRequest(String name, RepositoryPullRequestArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/repositoryPullRequest:RepositoryPullRequest", name, args == null ? RepositoryPullRequestArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private RepositoryPullRequest(String name, Output<String> id, @Nullable RepositoryPullRequestState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/repositoryPullRequest:RepositoryPullRequest", name, state, makeResourceOptions(options, id));
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
    public static RepositoryPullRequest get(String name, Output<String> id, @Nullable RepositoryPullRequestState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new RepositoryPullRequest(name, id, state, options);
    }
}