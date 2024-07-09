// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.github.EnterpriseActionsPermissionsArgs;
import com.pulumi.github.Utilities;
import com.pulumi.github.inputs.EnterpriseActionsPermissionsState;
import com.pulumi.github.outputs.EnterpriseActionsPermissionsAllowedActionsConfig;
import com.pulumi.github.outputs.EnterpriseActionsPermissionsEnabledOrganizationsConfig;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * This resource allows you to create and manage GitHub Actions permissions within your GitHub enterprise.
 * You must have admin access to an enterprise to use this resource.
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
 * import com.pulumi.github.inputs.GetOrganizationArgs;
 * import com.pulumi.github.EnterpriseActionsPermissions;
 * import com.pulumi.github.EnterpriseActionsPermissionsArgs;
 * import com.pulumi.github.inputs.EnterpriseActionsPermissionsAllowedActionsConfigArgs;
 * import com.pulumi.github.inputs.EnterpriseActionsPermissionsEnabledOrganizationsConfigArgs;
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
 *         final var example-org = GithubFunctions.getOrganization(GetOrganizationArgs.builder()
 *             .name("my-org")
 *             .build());
 * 
 *         var test = new EnterpriseActionsPermissions("test", EnterpriseActionsPermissionsArgs.builder()
 *             .enterpriseSlug("my-enterprise")
 *             .allowedActions("selected")
 *             .enabledOrganizations("selected")
 *             .allowedActionsConfig(EnterpriseActionsPermissionsAllowedActionsConfigArgs.builder()
 *                 .githubOwnedAllowed(true)
 *                 .patternsAlloweds(                
 *                     "actions/cache{@literal @}*",
 *                     "actions/checkout{@literal @}*")
 *                 .verifiedAllowed(true)
 *                 .build())
 *             .enabledOrganizationsConfig(EnterpriseActionsPermissionsEnabledOrganizationsConfigArgs.builder()
 *                 .organizationIds(example_org.id())
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
 * This resource can be imported using the name of the GitHub enterprise:
 * 
 * ```sh
 * $ pulumi import github:index/enterpriseActionsPermissions:EnterpriseActionsPermissions test github_enterprise_name
 * ```
 * 
 */
@ResourceType(type="github:index/enterpriseActionsPermissions:EnterpriseActionsPermissions")
public class EnterpriseActionsPermissions extends com.pulumi.resources.CustomResource {
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
     * Sets the actions that are allowed in an enterprise. Only available when `allowed_actions` = `selected`. See Allowed Actions Config below for details.
     * 
     */
    @Export(name="allowedActionsConfig", refs={EnterpriseActionsPermissionsAllowedActionsConfig.class}, tree="[0]")
    private Output</* @Nullable */ EnterpriseActionsPermissionsAllowedActionsConfig> allowedActionsConfig;

    /**
     * @return Sets the actions that are allowed in an enterprise. Only available when `allowed_actions` = `selected`. See Allowed Actions Config below for details.
     * 
     */
    public Output<Optional<EnterpriseActionsPermissionsAllowedActionsConfig>> allowedActionsConfig() {
        return Codegen.optional(this.allowedActionsConfig);
    }
    /**
     * The policy that controls the organizations in the enterprise that are allowed to run GitHub Actions. Can be one of: `all`, `none`, or `selected`.
     * 
     */
    @Export(name="enabledOrganizations", refs={String.class}, tree="[0]")
    private Output<String> enabledOrganizations;

    /**
     * @return The policy that controls the organizations in the enterprise that are allowed to run GitHub Actions. Can be one of: `all`, `none`, or `selected`.
     * 
     */
    public Output<String> enabledOrganizations() {
        return this.enabledOrganizations;
    }
    /**
     * Sets the list of selected organizations that are enabled for GitHub Actions in an enterprise. Only available when `enabled_organizations` = `selected`. See Enabled Organizations Config below for details.
     * 
     */
    @Export(name="enabledOrganizationsConfig", refs={EnterpriseActionsPermissionsEnabledOrganizationsConfig.class}, tree="[0]")
    private Output</* @Nullable */ EnterpriseActionsPermissionsEnabledOrganizationsConfig> enabledOrganizationsConfig;

    /**
     * @return Sets the list of selected organizations that are enabled for GitHub Actions in an enterprise. Only available when `enabled_organizations` = `selected`. See Enabled Organizations Config below for details.
     * 
     */
    public Output<Optional<EnterpriseActionsPermissionsEnabledOrganizationsConfig>> enabledOrganizationsConfig() {
        return Codegen.optional(this.enabledOrganizationsConfig);
    }
    /**
     * The slug of the enterprise.
     * 
     */
    @Export(name="enterpriseSlug", refs={String.class}, tree="[0]")
    private Output<String> enterpriseSlug;

    /**
     * @return The slug of the enterprise.
     * 
     */
    public Output<String> enterpriseSlug() {
        return this.enterpriseSlug;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public EnterpriseActionsPermissions(String name) {
        this(name, EnterpriseActionsPermissionsArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public EnterpriseActionsPermissions(String name, EnterpriseActionsPermissionsArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public EnterpriseActionsPermissions(String name, EnterpriseActionsPermissionsArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/enterpriseActionsPermissions:EnterpriseActionsPermissions", name, args == null ? EnterpriseActionsPermissionsArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private EnterpriseActionsPermissions(String name, Output<String> id, @Nullable EnterpriseActionsPermissionsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/enterpriseActionsPermissions:EnterpriseActionsPermissions", name, state, makeResourceOptions(options, id));
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
    public static EnterpriseActionsPermissions get(String name, Output<String> id, @Nullable EnterpriseActionsPermissionsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new EnterpriseActionsPermissions(name, id, state, options);
    }
}
