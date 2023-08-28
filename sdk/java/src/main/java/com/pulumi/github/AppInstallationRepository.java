// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.github.AppInstallationRepositoryArgs;
import com.pulumi.github.Utilities;
import com.pulumi.github.inputs.AppInstallationRepositoryState;
import java.lang.Integer;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * &gt; **Note**: This resource is not compatible with the GitHub App Installation authentication method.
 * 
 * This resource manages relationships between app installations and repositories
 * in your GitHub organization.
 * 
 * Creating this resource installs a particular app on a particular repository.
 * 
 * The app installation and the repository must both belong to the same
 * organization on GitHub. Note: you can review your organization&#39;s installations
 * by the following the instructions at this
 * [link](https://docs.github.com/en/github/setting-up-and-managing-organizations-and-teams/reviewing-your-organizations-installed-integrations).
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.github.Repository;
 * import com.pulumi.github.AppInstallationRepository;
 * import com.pulumi.github.AppInstallationRepositoryArgs;
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
 *         var someRepo = new Repository(&#34;someRepo&#34;);
 * 
 *         var someAppRepo = new AppInstallationRepository(&#34;someAppRepo&#34;, AppInstallationRepositoryArgs.builder()        
 *             .installationId(&#34;1234567&#34;)
 *             .repository(someRepo.name())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * GitHub App Installation Repository can be imported using an ID made up of `installation_id:repository`, e.g.
 * 
 * ```sh
 *  $ pulumi import github:index/appInstallationRepository:AppInstallationRepository terraform_repo 1234567:terraform
 * ```
 * 
 */
@ResourceType(type="github:index/appInstallationRepository:AppInstallationRepository")
public class AppInstallationRepository extends com.pulumi.resources.CustomResource {
    /**
     * The GitHub app installation id.
     * 
     */
    @Export(name="installationId", refs={String.class}, tree="[0]")
    private Output<String> installationId;

    /**
     * @return The GitHub app installation id.
     * 
     */
    public Output<String> installationId() {
        return this.installationId;
    }
    @Export(name="repoId", refs={Integer.class}, tree="[0]")
    private Output<Integer> repoId;

    public Output<Integer> repoId() {
        return this.repoId;
    }
    /**
     * The repository to install the app on.
     * 
     */
    @Export(name="repository", refs={String.class}, tree="[0]")
    private Output<String> repository;

    /**
     * @return The repository to install the app on.
     * 
     */
    public Output<String> repository() {
        return this.repository;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AppInstallationRepository(String name) {
        this(name, AppInstallationRepositoryArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AppInstallationRepository(String name, AppInstallationRepositoryArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AppInstallationRepository(String name, AppInstallationRepositoryArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/appInstallationRepository:AppInstallationRepository", name, args == null ? AppInstallationRepositoryArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AppInstallationRepository(String name, Output<String> id, @Nullable AppInstallationRepositoryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/appInstallationRepository:AppInstallationRepository", name, state, makeResourceOptions(options, id));
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
    public static AppInstallationRepository get(String name, Output<String> id, @Nullable AppInstallationRepositoryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AppInstallationRepository(name, id, state, options);
    }
}
