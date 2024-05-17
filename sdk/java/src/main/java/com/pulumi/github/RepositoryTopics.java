// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.github.RepositoryTopicsArgs;
import com.pulumi.github.Utilities;
import com.pulumi.github.inputs.RepositoryTopicsState;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

/**
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
 * import com.pulumi.github.inputs.GetRepositoryArgs;
 * import com.pulumi.github.RepositoryTopics;
 * import com.pulumi.github.RepositoryTopicsArgs;
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
 *         final var test = GithubFunctions.getRepository(GetRepositoryArgs.builder()
 *             .name("test")
 *             .build());
 * 
 *         var testRepositoryTopics = new RepositoryTopics("testRepositoryTopics", RepositoryTopicsArgs.builder()
 *             .repository(testGithubRepository.name())
 *             .topics(            
 *                 "topic-1",
 *                 "topic-2")
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
 * Repository topics can be imported using the `name` of the repository.
 * 
 * ```sh
 * $ pulumi import github:index/repositoryTopics:RepositoryTopics terraform terraform
 * ```
 * 
 */
@ResourceType(type="github:index/repositoryTopics:RepositoryTopics")
public class RepositoryTopics extends com.pulumi.resources.CustomResource {
    /**
     * The repository name.
     * 
     */
    @Export(name="repository", refs={String.class}, tree="[0]")
    private Output<String> repository;

    /**
     * @return The repository name.
     * 
     */
    public Output<String> repository() {
        return this.repository;
    }
    /**
     * A list of topics to add to the repository.
     * 
     */
    @Export(name="topics", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> topics;

    /**
     * @return A list of topics to add to the repository.
     * 
     */
    public Output<List<String>> topics() {
        return this.topics;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public RepositoryTopics(String name) {
        this(name, RepositoryTopicsArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public RepositoryTopics(String name, RepositoryTopicsArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public RepositoryTopics(String name, RepositoryTopicsArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/repositoryTopics:RepositoryTopics", name, args == null ? RepositoryTopicsArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private RepositoryTopics(String name, Output<String> id, @Nullable RepositoryTopicsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/repositoryTopics:RepositoryTopics", name, state, makeResourceOptions(options, id));
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
    public static RepositoryTopics get(String name, Output<String> id, @Nullable RepositoryTopicsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new RepositoryTopics(name, id, state, options);
    }
}
