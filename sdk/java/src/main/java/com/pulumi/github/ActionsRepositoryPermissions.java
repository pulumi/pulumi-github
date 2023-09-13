// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.github.ActionsRepositoryPermissionsArgs;
import com.pulumi.github.Utilities;
import com.pulumi.github.inputs.ActionsRepositoryPermissionsState;
import com.pulumi.github.outputs.ActionsRepositoryPermissionsAllowedActionsConfig;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * This resource allows you to enable and manage GitHub Actions permissions for a given repository.
 * You must have admin access to an repository to use this resource.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.github.Repository;
 * import com.pulumi.github.ActionsRepositoryPermissions;
 * import com.pulumi.github.ActionsRepositoryPermissionsArgs;
 * import com.pulumi.github.inputs.ActionsRepositoryPermissionsAllowedActionsConfigArgs;
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
 *         var example = new Repository(&#34;example&#34;);
 * 
 *         var test = new ActionsRepositoryPermissions(&#34;test&#34;, ActionsRepositoryPermissionsArgs.builder()        
 *             .allowedActions(&#34;selected&#34;)
 *             .allowedActionsConfig(ActionsRepositoryPermissionsAllowedActionsConfigArgs.builder()
 *                 .githubOwnedAllowed(true)
 *                 .patternsAlloweds(                
 *                     &#34;actions/cache@*&#34;,
 *                     &#34;actions/checkout@*&#34;)
 *                 .verifiedAllowed(true)
 *                 .build())
 *             .repository(example.name())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * This resource can be imported using the name of the GitHub repository:
 * 
 * ```sh
 *  $ pulumi import github:index/actionsRepositoryPermissions:ActionsRepositoryPermissions test my-repository
 * ```
 * 
 */
@ResourceType(type="github:index/actionsRepositoryPermissions:ActionsRepositoryPermissions")
public class ActionsRepositoryPermissions extends com.pulumi.resources.CustomResource {
    /**
     * The permissions policy that controls the actions that are allowed to run. Can be one of: `all`, `local_only`, or `selected`.
     * 
     */
    @Export(name="allowedActions", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> allowedActions;

    /**
     * @return The permissions policy that controls the actions that are allowed to run. Can be one of: `all`, `local_only`, or `selected`.
     * 
     */
    public Output<Optional<String>> allowedActions() {
        return Codegen.optional(this.allowedActions);
    }
    /**
     * Sets the actions that are allowed in an repository. Only available when `allowed_actions` = `selected`. See Allowed Actions Config below for details.
     * 
     */
    @Export(name="allowedActionsConfig", refs={ActionsRepositoryPermissionsAllowedActionsConfig.class}, tree="[0]")
    private Output</* @Nullable */ ActionsRepositoryPermissionsAllowedActionsConfig> allowedActionsConfig;

    /**
     * @return Sets the actions that are allowed in an repository. Only available when `allowed_actions` = `selected`. See Allowed Actions Config below for details.
     * 
     */
    public Output<Optional<ActionsRepositoryPermissionsAllowedActionsConfig>> allowedActionsConfig() {
        return Codegen.optional(this.allowedActionsConfig);
    }
    /**
     * Should GitHub actions be enabled on this repository?
     * 
     */
    @Export(name="enabled", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> enabled;

    /**
     * @return Should GitHub actions be enabled on this repository?
     * 
     */
    public Output<Optional<Boolean>> enabled() {
        return Codegen.optional(this.enabled);
    }
    /**
     * The GitHub repository
     * 
     */
    @Export(name="repository", refs={String.class}, tree="[0]")
    private Output<String> repository;

    /**
     * @return The GitHub repository
     * 
     */
    public Output<String> repository() {
        return this.repository;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ActionsRepositoryPermissions(String name) {
        this(name, ActionsRepositoryPermissionsArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ActionsRepositoryPermissions(String name, ActionsRepositoryPermissionsArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ActionsRepositoryPermissions(String name, ActionsRepositoryPermissionsArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/actionsRepositoryPermissions:ActionsRepositoryPermissions", name, args == null ? ActionsRepositoryPermissionsArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ActionsRepositoryPermissions(String name, Output<String> id, @Nullable ActionsRepositoryPermissionsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/actionsRepositoryPermissions:ActionsRepositoryPermissions", name, state, makeResourceOptions(options, id));
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
    public static ActionsRepositoryPermissions get(String name, Output<String> id, @Nullable ActionsRepositoryPermissionsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ActionsRepositoryPermissions(name, id, state, options);
    }
}
