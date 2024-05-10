// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.github.OrganizationProjectArgs;
import com.pulumi.github.Utilities;
import com.pulumi.github.inputs.OrganizationProjectState;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * This resource allows you to create and manage projects for GitHub organization.
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
 * import com.pulumi.github.OrganizationProject;
 * import com.pulumi.github.OrganizationProjectArgs;
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
 *         var project = new OrganizationProject("project", OrganizationProjectArgs.builder()        
 *             .name("A Organization Project")
 *             .body("This is a organization project.")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="github:index/organizationProject:OrganizationProject")
public class OrganizationProject extends com.pulumi.resources.CustomResource {
    /**
     * The body of the project.
     * 
     */
    @Export(name="body", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> body;

    /**
     * @return The body of the project.
     * 
     */
    public Output<Optional<String>> body() {
        return Codegen.optional(this.body);
    }
    @Export(name="etag", refs={String.class}, tree="[0]")
    private Output<String> etag;

    public Output<String> etag() {
        return this.etag;
    }
    /**
     * The name of the project.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the project.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * URL of the project
     * 
     */
    @Export(name="url", refs={String.class}, tree="[0]")
    private Output<String> url;

    /**
     * @return URL of the project
     * 
     */
    public Output<String> url() {
        return this.url;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public OrganizationProject(String name) {
        this(name, OrganizationProjectArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public OrganizationProject(String name, @Nullable OrganizationProjectArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public OrganizationProject(String name, @Nullable OrganizationProjectArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/organizationProject:OrganizationProject", name, args == null ? OrganizationProjectArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private OrganizationProject(String name, Output<String> id, @Nullable OrganizationProjectState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/organizationProject:OrganizationProject", name, state, makeResourceOptions(options, id));
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
    public static OrganizationProject get(String name, Output<String> id, @Nullable OrganizationProjectState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new OrganizationProject(name, id, state, options);
    }
}
