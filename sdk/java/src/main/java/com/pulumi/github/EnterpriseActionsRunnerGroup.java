// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.github.EnterpriseActionsRunnerGroupArgs;
import com.pulumi.github.Utilities;
import com.pulumi.github.inputs.EnterpriseActionsRunnerGroupState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * This resource allows you to create and manage GitHub Actions runner groups within your GitHub enterprise.
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
 * import com.pulumi.github.inputs.GetEnterpriseArgs;
 * import com.pulumi.github.EnterpriseOrganization;
 * import com.pulumi.github.EnterpriseOrganizationArgs;
 * import com.pulumi.github.EnterpriseActionsRunnerGroup;
 * import com.pulumi.github.EnterpriseActionsRunnerGroupArgs;
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
 *         final var enterprise = GithubFunctions.getEnterprise(GetEnterpriseArgs.builder()
 *             .slug("my-enterprise")
 *             .build());
 * 
 *         var enterpriseOrganization = new EnterpriseOrganization("enterpriseOrganization", EnterpriseOrganizationArgs.builder()        
 *             .enterpriseId(enterprise.applyValue(getEnterpriseResult -> getEnterpriseResult.id()))
 *             .name("my-organization")
 *             .billingEmail("octocat{@literal @}octo.cat")
 *             .adminLogins("octocat")
 *             .build());
 * 
 *         var example = new EnterpriseActionsRunnerGroup("example", EnterpriseActionsRunnerGroupArgs.builder()        
 *             .name("my-awesome-runner-group")
 *             .enterpriseSlug(enterprise.applyValue(getEnterpriseResult -> getEnterpriseResult.slug()))
 *             .allowsPublicRepositories(true)
 *             .visibility("selected")
 *             .selectedOrganizationIds(enterpriseOrganization.databaseId())
 *             .restrictedToWorkflows(true)
 *             .selectedWorkflows("my-organization/my-repo/.github/workflows/cool-workflow.yaml{@literal @}refs/tags/v1")
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
 * This resource can be imported using the enterprise slug and the ID of the runner group:
 * 
 * ```sh
 * $ pulumi import github:index/enterpriseActionsRunnerGroup:EnterpriseActionsRunnerGroup test enterprise-slug/42
 * ```
 * 
 */
@ResourceType(type="github:index/enterpriseActionsRunnerGroup:EnterpriseActionsRunnerGroup")
public class EnterpriseActionsRunnerGroup extends com.pulumi.resources.CustomResource {
    /**
     * Whether public repositories can be added to the runner group. Defaults to false.
     * 
     */
    @Export(name="allowsPublicRepositories", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> allowsPublicRepositories;

    /**
     * @return Whether public repositories can be added to the runner group. Defaults to false.
     * 
     */
    public Output<Optional<Boolean>> allowsPublicRepositories() {
        return Codegen.optional(this.allowsPublicRepositories);
    }
    /**
     * Whether this is the default runner group
     * 
     */
    @Export(name="default", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> default_;

    /**
     * @return Whether this is the default runner group
     * 
     */
    public Output<Boolean> default_() {
        return this.default_;
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
     * An etag representing the runner group object
     * 
     */
    @Export(name="etag", refs={String.class}, tree="[0]")
    private Output<String> etag;

    /**
     * @return An etag representing the runner group object
     * 
     */
    public Output<String> etag() {
        return this.etag;
    }
    /**
     * Name of the runner group
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the runner group
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * If true, the runner group will be restricted to running only the workflows specified in the selected_workflows array. Defaults to false.
     * 
     */
    @Export(name="restrictedToWorkflows", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> restrictedToWorkflows;

    /**
     * @return If true, the runner group will be restricted to running only the workflows specified in the selected_workflows array. Defaults to false.
     * 
     */
    public Output<Optional<Boolean>> restrictedToWorkflows() {
        return Codegen.optional(this.restrictedToWorkflows);
    }
    /**
     * The GitHub API URL for the runner group&#39;s runners
     * 
     */
    @Export(name="runnersUrl", refs={String.class}, tree="[0]")
    private Output<String> runnersUrl;

    /**
     * @return The GitHub API URL for the runner group&#39;s runners
     * 
     */
    public Output<String> runnersUrl() {
        return this.runnersUrl;
    }
    /**
     * IDs of the organizations which should be added to the runner group
     * 
     */
    @Export(name="selectedOrganizationIds", refs={List.class,Integer.class}, tree="[0,1]")
    private Output</* @Nullable */ List<Integer>> selectedOrganizationIds;

    /**
     * @return IDs of the organizations which should be added to the runner group
     * 
     */
    public Output<Optional<List<Integer>>> selectedOrganizationIds() {
        return Codegen.optional(this.selectedOrganizationIds);
    }
    /**
     * The GitHub API URL for the runner group&#39;s selected organizations
     * 
     */
    @Export(name="selectedOrganizationsUrl", refs={String.class}, tree="[0]")
    private Output<String> selectedOrganizationsUrl;

    /**
     * @return The GitHub API URL for the runner group&#39;s selected organizations
     * 
     */
    public Output<String> selectedOrganizationsUrl() {
        return this.selectedOrganizationsUrl;
    }
    /**
     * List of workflows the runner group should be allowed to run. This setting will be ignored unless restricted_to_workflows is set to true.
     * 
     */
    @Export(name="selectedWorkflows", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> selectedWorkflows;

    /**
     * @return List of workflows the runner group should be allowed to run. This setting will be ignored unless restricted_to_workflows is set to true.
     * 
     */
    public Output<Optional<List<String>>> selectedWorkflows() {
        return Codegen.optional(this.selectedWorkflows);
    }
    /**
     * Visibility of a runner group to enterprise organizations. Whether the runner group can include `all` or `selected`
     * 
     */
    @Export(name="visibility", refs={String.class}, tree="[0]")
    private Output<String> visibility;

    /**
     * @return Visibility of a runner group to enterprise organizations. Whether the runner group can include `all` or `selected`
     * 
     */
    public Output<String> visibility() {
        return this.visibility;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public EnterpriseActionsRunnerGroup(String name) {
        this(name, EnterpriseActionsRunnerGroupArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public EnterpriseActionsRunnerGroup(String name, EnterpriseActionsRunnerGroupArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public EnterpriseActionsRunnerGroup(String name, EnterpriseActionsRunnerGroupArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/enterpriseActionsRunnerGroup:EnterpriseActionsRunnerGroup", name, args == null ? EnterpriseActionsRunnerGroupArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private EnterpriseActionsRunnerGroup(String name, Output<String> id, @Nullable EnterpriseActionsRunnerGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/enterpriseActionsRunnerGroup:EnterpriseActionsRunnerGroup", name, state, makeResourceOptions(options, id));
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
    public static EnterpriseActionsRunnerGroup get(String name, Output<String> id, @Nullable EnterpriseActionsRunnerGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new EnterpriseActionsRunnerGroup(name, id, state, options);
    }
}
