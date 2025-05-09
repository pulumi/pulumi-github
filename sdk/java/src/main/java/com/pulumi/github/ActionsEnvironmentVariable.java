// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.github.ActionsEnvironmentVariableArgs;
import com.pulumi.github.Utilities;
import com.pulumi.github.inputs.ActionsEnvironmentVariableState;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * This resource allows you to create and manage GitHub Actions variables within your GitHub repository environments.
 * You must have write access to a repository to use this resource.
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
 * import com.pulumi.github.ActionsEnvironmentVariable;
 * import com.pulumi.github.ActionsEnvironmentVariableArgs;
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
 *         var exampleVariable = new ActionsEnvironmentVariable("exampleVariable", ActionsEnvironmentVariableArgs.builder()
 *             .environment("example_environment")
 *             .variableName("example_variable_name")
 *             .value("example_variable_value")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
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
 * import com.pulumi.github.RepositoryEnvironment;
 * import com.pulumi.github.RepositoryEnvironmentArgs;
 * import com.pulumi.github.ActionsEnvironmentVariable;
 * import com.pulumi.github.ActionsEnvironmentVariableArgs;
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
 *             .fullName("my-org/repo")
 *             .build());
 * 
 *         var repoEnvironment = new RepositoryEnvironment("repoEnvironment", RepositoryEnvironmentArgs.builder()
 *             .repository(repo.name())
 *             .environment("example_environment")
 *             .build());
 * 
 *         var exampleVariable = new ActionsEnvironmentVariable("exampleVariable", ActionsEnvironmentVariableArgs.builder()
 *             .repository(repo.name())
 *             .environment(repoEnvironment.environment())
 *             .variableName("example_variable_name")
 *             .value("example_variable_value")
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
 * This resource can be imported using an ID made up of the repository name, environment name, and variable name:
 * 
 * ```sh
 * $ pulumi import github:index/actionsEnvironmentVariable:ActionsEnvironmentVariable test_variable myrepo:myenv:myvariable
 * ```
 * 
 */
@ResourceType(type="github:index/actionsEnvironmentVariable:ActionsEnvironmentVariable")
public class ActionsEnvironmentVariable extends com.pulumi.resources.CustomResource {
    /**
     * Date of actions_environment_secret creation.
     * 
     */
    @Export(name="createdAt", refs={String.class}, tree="[0]")
    private Output<String> createdAt;

    /**
     * @return Date of actions_environment_secret creation.
     * 
     */
    public Output<String> createdAt() {
        return this.createdAt;
    }
    /**
     * Name of the environment.
     * 
     */
    @Export(name="environment", refs={String.class}, tree="[0]")
    private Output<String> environment;

    /**
     * @return Name of the environment.
     * 
     */
    public Output<String> environment() {
        return this.environment;
    }
    /**
     * Name of the repository.
     * 
     */
    @Export(name="repository", refs={String.class}, tree="[0]")
    private Output<String> repository;

    /**
     * @return Name of the repository.
     * 
     */
    public Output<String> repository() {
        return this.repository;
    }
    /**
     * Date of actions_environment_secret update.
     * 
     */
    @Export(name="updatedAt", refs={String.class}, tree="[0]")
    private Output<String> updatedAt;

    /**
     * @return Date of actions_environment_secret update.
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
     * Name of the variable.
     * 
     */
    @Export(name="variableName", refs={String.class}, tree="[0]")
    private Output<String> variableName;

    /**
     * @return Name of the variable.
     * 
     */
    public Output<String> variableName() {
        return this.variableName;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ActionsEnvironmentVariable(java.lang.String name) {
        this(name, ActionsEnvironmentVariableArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ActionsEnvironmentVariable(java.lang.String name, ActionsEnvironmentVariableArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ActionsEnvironmentVariable(java.lang.String name, ActionsEnvironmentVariableArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/actionsEnvironmentVariable:ActionsEnvironmentVariable", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private ActionsEnvironmentVariable(java.lang.String name, Output<java.lang.String> id, @Nullable ActionsEnvironmentVariableState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/actionsEnvironmentVariable:ActionsEnvironmentVariable", name, state, makeResourceOptions(options, id), false);
    }

    private static ActionsEnvironmentVariableArgs makeArgs(ActionsEnvironmentVariableArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ActionsEnvironmentVariableArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
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
    public static ActionsEnvironmentVariable get(java.lang.String name, Output<java.lang.String> id, @Nullable ActionsEnvironmentVariableState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ActionsEnvironmentVariable(name, id, state, options);
    }
}
