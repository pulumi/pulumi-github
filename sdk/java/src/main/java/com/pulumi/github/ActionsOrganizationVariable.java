// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.github.ActionsOrganizationVariableArgs;
import com.pulumi.github.Utilities;
import com.pulumi.github.inputs.ActionsOrganizationVariableState;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * This resource allows you to create and manage GitHub Actions variables within your GitHub organization.
 * You must have write access to a repository to use this resource.
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.github.ActionsOrganizationVariable;
 * import com.pulumi.github.ActionsOrganizationVariableArgs;
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
 *         var exampleVariable = new ActionsOrganizationVariable(&#34;exampleVariable&#34;, ActionsOrganizationVariableArgs.builder()        
 *             .variableName(&#34;example_variable_name&#34;)
 *             .visibility(&#34;private&#34;)
 *             .value(&#34;example_variable_value&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.github.GithubFunctions;
 * import com.pulumi.github.inputs.GetRepositoryArgs;
 * import com.pulumi.github.ActionsOrganizationVariable;
 * import com.pulumi.github.ActionsOrganizationVariableArgs;
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
 *         final var repo = GithubFunctions.getRepository(GetRepositoryArgs.builder()
 *             .fullName(&#34;my-org/repo&#34;)
 *             .build());
 * 
 *         var exampleVariable = new ActionsOrganizationVariable(&#34;exampleVariable&#34;, ActionsOrganizationVariableArgs.builder()        
 *             .variableName(&#34;example_variable_name&#34;)
 *             .visibility(&#34;selected&#34;)
 *             .value(&#34;example_variable_value&#34;)
 *             .selectedRepositoryIds(repo.applyValue(getRepositoryResult -&gt; getRepositoryResult.repoId()))
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * This resource can be imported using an ID made up of the variable name:
 * 
 * ```sh
 * $ pulumi import github:index/actionsOrganizationVariable:ActionsOrganizationVariable test_variable test_variable_name
 * ```
 * 
 */
@ResourceType(type="github:index/actionsOrganizationVariable:ActionsOrganizationVariable")
public class ActionsOrganizationVariable extends com.pulumi.resources.CustomResource {
    /**
     * Date of actions_variable creation.
     * 
     */
    @Export(name="createdAt", refs={String.class}, tree="[0]")
    private Output<String> createdAt;

    /**
     * @return Date of actions_variable creation.
     * 
     */
    public Output<String> createdAt() {
        return this.createdAt;
    }
    /**
     * An array of repository ids that can access the organization variable.
     * 
     */
    @Export(name="selectedRepositoryIds", refs={List.class,Integer.class}, tree="[0,1]")
    private Output</* @Nullable */ List<Integer>> selectedRepositoryIds;

    /**
     * @return An array of repository ids that can access the organization variable.
     * 
     */
    public Output<Optional<List<Integer>>> selectedRepositoryIds() {
        return Codegen.optional(this.selectedRepositoryIds);
    }
    /**
     * Date of actions_variable update.
     * 
     */
    @Export(name="updatedAt", refs={String.class}, tree="[0]")
    private Output<String> updatedAt;

    /**
     * @return Date of actions_variable update.
     * 
     */
    public Output<String> updatedAt() {
        return this.updatedAt;
    }
    /**
     * Value of the variable
     * 
     */
    @Export(name="value", refs={String.class}, tree="[0]")
    private Output<String> value;

    /**
     * @return Value of the variable
     * 
     */
    public Output<String> value() {
        return this.value;
    }
    /**
     * Name of the variable
     * 
     */
    @Export(name="variableName", refs={String.class}, tree="[0]")
    private Output<String> variableName;

    /**
     * @return Name of the variable
     * 
     */
    public Output<String> variableName() {
        return this.variableName;
    }
    /**
     * Configures the access that repositories have to the organization variable.
     * Must be one of `all`, `private`, `selected`. `selected_repository_ids` is required if set to `selected`.
     * 
     */
    @Export(name="visibility", refs={String.class}, tree="[0]")
    private Output<String> visibility;

    /**
     * @return Configures the access that repositories have to the organization variable.
     * Must be one of `all`, `private`, `selected`. `selected_repository_ids` is required if set to `selected`.
     * 
     */
    public Output<String> visibility() {
        return this.visibility;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ActionsOrganizationVariable(String name) {
        this(name, ActionsOrganizationVariableArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ActionsOrganizationVariable(String name, ActionsOrganizationVariableArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ActionsOrganizationVariable(String name, ActionsOrganizationVariableArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/actionsOrganizationVariable:ActionsOrganizationVariable", name, args == null ? ActionsOrganizationVariableArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ActionsOrganizationVariable(String name, Output<String> id, @Nullable ActionsOrganizationVariableState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/actionsOrganizationVariable:ActionsOrganizationVariable", name, state, makeResourceOptions(options, id));
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
    public static ActionsOrganizationVariable get(String name, Output<String> id, @Nullable ActionsOrganizationVariableState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ActionsOrganizationVariable(name, id, state, options);
    }
}
