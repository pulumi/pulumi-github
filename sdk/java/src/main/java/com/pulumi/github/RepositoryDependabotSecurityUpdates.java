// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.github.RepositoryDependabotSecurityUpdatesArgs;
import com.pulumi.github.Utilities;
import com.pulumi.github.inputs.RepositoryDependabotSecurityUpdatesState;
import java.lang.Boolean;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * This resource allows you to manage dependabot automated security fixes for a single repository. See the
 * [documentation](https://docs.github.com/en/code-security/dependabot/dependabot-security-updates/about-dependabot-security-updates)
 * for details of usage and how this will impact your repository
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.github.Repository;
 * import com.pulumi.github.RepositoryArgs;
 * import com.pulumi.github.RepositoryDependabotSecurityUpdates;
 * import com.pulumi.github.RepositoryDependabotSecurityUpdatesArgs;
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
 *         var repo = new Repository(&#34;repo&#34;, RepositoryArgs.builder()        
 *             .description(&#34;GitHub repo managed by Terraform&#34;)
 *             .private_(false)
 *             .vulnerabilityAlerts(true)
 *             .build());
 * 
 *         var example = new RepositoryDependabotSecurityUpdates(&#34;example&#34;, RepositoryDependabotSecurityUpdatesArgs.builder()        
 *             .repository(github_repository.test().id())
 *             .enabled(true)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * ### Import by name
 * 
 * ```sh
 * $ pulumi import github:index/repositoryDependabotSecurityUpdates:RepositoryDependabotSecurityUpdates example my-repo
 * ```
 * 
 */
@ResourceType(type="github:index/repositoryDependabotSecurityUpdates:RepositoryDependabotSecurityUpdates")
public class RepositoryDependabotSecurityUpdates extends com.pulumi.resources.CustomResource {
    /**
     * The state of the automated security fixes.
     * 
     */
    @Export(name="enabled", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> enabled;

    /**
     * @return The state of the automated security fixes.
     * 
     */
    public Output<Boolean> enabled() {
        return this.enabled;
    }
    /**
     * The repository to manage.
     * 
     */
    @Export(name="repository", refs={String.class}, tree="[0]")
    private Output<String> repository;

    /**
     * @return The repository to manage.
     * 
     */
    public Output<String> repository() {
        return this.repository;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public RepositoryDependabotSecurityUpdates(String name) {
        this(name, RepositoryDependabotSecurityUpdatesArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public RepositoryDependabotSecurityUpdates(String name, RepositoryDependabotSecurityUpdatesArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public RepositoryDependabotSecurityUpdates(String name, RepositoryDependabotSecurityUpdatesArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/repositoryDependabotSecurityUpdates:RepositoryDependabotSecurityUpdates", name, args == null ? RepositoryDependabotSecurityUpdatesArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private RepositoryDependabotSecurityUpdates(String name, Output<String> id, @Nullable RepositoryDependabotSecurityUpdatesState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/repositoryDependabotSecurityUpdates:RepositoryDependabotSecurityUpdates", name, state, makeResourceOptions(options, id));
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
    public static RepositoryDependabotSecurityUpdates get(String name, Output<String> id, @Nullable RepositoryDependabotSecurityUpdatesState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new RepositoryDependabotSecurityUpdates(name, id, state, options);
    }
}
