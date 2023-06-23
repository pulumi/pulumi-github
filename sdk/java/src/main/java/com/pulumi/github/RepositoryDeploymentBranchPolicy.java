// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.github.RepositoryDeploymentBranchPolicyArgs;
import com.pulumi.github.Utilities;
import com.pulumi.github.inputs.RepositoryDeploymentBranchPolicyState;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * This resource allows you to create and manage deployment branch policies.
 * 
 * ## Example Usage
 * 
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.github.RepositoryEnvironment;
 * import com.pulumi.github.RepositoryEnvironmentArgs;
 * import com.pulumi.github.inputs.RepositoryEnvironmentDeploymentBranchPolicyArgs;
 * import com.pulumi.github.RepositoryDeploymentBranchPolicy;
 * import com.pulumi.github.RepositoryDeploymentBranchPolicyArgs;
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
 *         var env = new RepositoryEnvironment(&#34;env&#34;, RepositoryEnvironmentArgs.builder()        
 *             .deploymentBranchPolicy(RepositoryEnvironmentDeploymentBranchPolicyArgs.builder()
 *                 .customBranchPolicies(true)
 *                 .protectedBranches(false)
 *                 .build())
 *             .environment(&#34;my_env&#34;)
 *             .repository(&#34;my_repo&#34;)
 *             .build());
 * 
 *         var foo = new RepositoryDeploymentBranchPolicy(&#34;foo&#34;, RepositoryDeploymentBranchPolicyArgs.builder()        
 *             .environmentName(&#34;my_env&#34;)
 *             .reposistory(&#34;my_repo&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * ```sh
 *  $ pulumi import github:index/repositoryDeploymentBranchPolicy:RepositoryDeploymentBranchPolicy foo repo:env:id
 * ```
 * 
 */
@ResourceType(type="github:index/repositoryDeploymentBranchPolicy:RepositoryDeploymentBranchPolicy")
public class RepositoryDeploymentBranchPolicy extends com.pulumi.resources.CustomResource {
    /**
     * The name of the environment. This environment must have `deployment_branch_policy.custom_branch_policies` set to true.
     * 
     */
    @Export(name="environmentName", type=String.class, parameters={})
    private Output<String> environmentName;

    /**
     * @return The name of the environment. This environment must have `deployment_branch_policy.custom_branch_policies` set to true.
     * 
     */
    public Output<String> environmentName() {
        return this.environmentName;
    }
    /**
     * An etag representing the Branch object.
     * 
     */
    @Export(name="etag", type=String.class, parameters={})
    private Output<String> etag;

    /**
     * @return An etag representing the Branch object.
     * 
     */
    public Output<String> etag() {
        return this.etag;
    }
    /**
     * The name pattern that branches must match in order to deploy to the environment.
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return The name pattern that branches must match in order to deploy to the environment.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The repository to create the policy in.
     * 
     */
    @Export(name="repository", type=String.class, parameters={})
    private Output<String> repository;

    /**
     * @return The repository to create the policy in.
     * 
     */
    public Output<String> repository() {
        return this.repository;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public RepositoryDeploymentBranchPolicy(String name) {
        this(name, RepositoryDeploymentBranchPolicyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public RepositoryDeploymentBranchPolicy(String name, RepositoryDeploymentBranchPolicyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public RepositoryDeploymentBranchPolicy(String name, RepositoryDeploymentBranchPolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/repositoryDeploymentBranchPolicy:RepositoryDeploymentBranchPolicy", name, args == null ? RepositoryDeploymentBranchPolicyArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private RepositoryDeploymentBranchPolicy(String name, Output<String> id, @Nullable RepositoryDeploymentBranchPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/repositoryDeploymentBranchPolicy:RepositoryDeploymentBranchPolicy", name, state, makeResourceOptions(options, id));
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
    public static RepositoryDeploymentBranchPolicy get(String name, Output<String> id, @Nullable RepositoryDeploymentBranchPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new RepositoryDeploymentBranchPolicy(name, id, state, options);
    }
}
