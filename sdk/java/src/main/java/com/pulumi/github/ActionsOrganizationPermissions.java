// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.github.ActionsOrganizationPermissionsArgs;
import com.pulumi.github.Utilities;
import com.pulumi.github.inputs.ActionsOrganizationPermissionsState;
import com.pulumi.github.outputs.ActionsOrganizationPermissionsAllowedActionsConfig;
import com.pulumi.github.outputs.ActionsOrganizationPermissionsEnabledRepositoriesConfig;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * This resource allows you to create and manage GitHub Actions permissions within your GitHub enterprise organizations.
 * You must have admin access to an organization to use this resource.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.github.Repository;
 * import com.pulumi.github.ActionsOrganizationPermissions;
 * import com.pulumi.github.ActionsOrganizationPermissionsArgs;
 * import com.pulumi.github.inputs.ActionsOrganizationPermissionsAllowedActionsConfigArgs;
 * import com.pulumi.github.inputs.ActionsOrganizationPermissionsEnabledRepositoriesConfigArgs;
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
 *         var test = new ActionsOrganizationPermissions(&#34;test&#34;, ActionsOrganizationPermissionsArgs.builder()        
 *             .allowedActions(&#34;selected&#34;)
 *             .enabledRepositories(&#34;selected&#34;)
 *             .allowedActionsConfig(ActionsOrganizationPermissionsAllowedActionsConfigArgs.builder()
 *                 .githubOwnedAllowed(true)
 *                 .patternsAlloweds(                
 *                     &#34;actions/cache@*&#34;,
 *                     &#34;actions/checkout@*&#34;)
 *                 .verifiedAllowed(true)
 *                 .build())
 *             .enabledRepositoriesConfig(ActionsOrganizationPermissionsEnabledRepositoriesConfigArgs.builder()
 *                 .repositoryIds(example.repoId())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * This resource can be imported using the ID of the GitHub organization
 * 
 * ```sh
 *  $ pulumi import github:index/actionsOrganizationPermissions:ActionsOrganizationPermissions test &lt;github_organization_name&gt;
 * ```
 * 
 */
@ResourceType(type="github:index/actionsOrganizationPermissions:ActionsOrganizationPermissions")
public class ActionsOrganizationPermissions extends com.pulumi.resources.CustomResource {
    /**
     * The permissions policy that controls the actions that are allowed to run. Can be one of: `all`, `local_only`, or `selected`.
     * 
     */
    @Export(name="allowedActions", type=String.class, parameters={})
    private Output</* @Nullable */ String> allowedActions;

    /**
     * @return The permissions policy that controls the actions that are allowed to run. Can be one of: `all`, `local_only`, or `selected`.
     * 
     */
    public Output<Optional<String>> allowedActions() {
        return Codegen.optional(this.allowedActions);
    }
    /**
     * Sets the actions that are allowed in an organization. Only available when `allowed_actions` = `selected`. See Allowed Actions Config below for details.
     * 
     */
    @Export(name="allowedActionsConfig", type=ActionsOrganizationPermissionsAllowedActionsConfig.class, parameters={})
    private Output</* @Nullable */ ActionsOrganizationPermissionsAllowedActionsConfig> allowedActionsConfig;

    /**
     * @return Sets the actions that are allowed in an organization. Only available when `allowed_actions` = `selected`. See Allowed Actions Config below for details.
     * 
     */
    public Output<Optional<ActionsOrganizationPermissionsAllowedActionsConfig>> allowedActionsConfig() {
        return Codegen.optional(this.allowedActionsConfig);
    }
    /**
     * The policy that controls the repositories in the organization that are allowed to run GitHub Actions. Can be one of: `all`, `none`, or `selected`.
     * 
     */
    @Export(name="enabledRepositories", type=String.class, parameters={})
    private Output<String> enabledRepositories;

    /**
     * @return The policy that controls the repositories in the organization that are allowed to run GitHub Actions. Can be one of: `all`, `none`, or `selected`.
     * 
     */
    public Output<String> enabledRepositories() {
        return this.enabledRepositories;
    }
    /**
     * Sets the list of selected repositories that are enabled for GitHub Actions in an organization. Only available when `enabled_repositories` = `selected`. See Enabled Repositories Config below for details.
     * 
     */
    @Export(name="enabledRepositoriesConfig", type=ActionsOrganizationPermissionsEnabledRepositoriesConfig.class, parameters={})
    private Output</* @Nullable */ ActionsOrganizationPermissionsEnabledRepositoriesConfig> enabledRepositoriesConfig;

    /**
     * @return Sets the list of selected repositories that are enabled for GitHub Actions in an organization. Only available when `enabled_repositories` = `selected`. See Enabled Repositories Config below for details.
     * 
     */
    public Output<Optional<ActionsOrganizationPermissionsEnabledRepositoriesConfig>> enabledRepositoriesConfig() {
        return Codegen.optional(this.enabledRepositoriesConfig);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ActionsOrganizationPermissions(String name) {
        this(name, ActionsOrganizationPermissionsArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ActionsOrganizationPermissions(String name, ActionsOrganizationPermissionsArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ActionsOrganizationPermissions(String name, ActionsOrganizationPermissionsArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/actionsOrganizationPermissions:ActionsOrganizationPermissions", name, args == null ? ActionsOrganizationPermissionsArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ActionsOrganizationPermissions(String name, Output<String> id, @Nullable ActionsOrganizationPermissionsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/actionsOrganizationPermissions:ActionsOrganizationPermissions", name, state, makeResourceOptions(options, id));
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
    public static ActionsOrganizationPermissions get(String name, Output<String> id, @Nullable ActionsOrganizationPermissionsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ActionsOrganizationPermissions(name, id, state, options);
    }
}