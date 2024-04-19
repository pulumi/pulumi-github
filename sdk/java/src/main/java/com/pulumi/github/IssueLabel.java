// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.github.IssueLabelArgs;
import com.pulumi.github.Utilities;
import com.pulumi.github.inputs.IssueLabelState;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.github.IssueLabel;
 * import com.pulumi.github.IssueLabelArgs;
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
 *         // Create a new, red colored label
 *         var testRepo = new IssueLabel(&#34;testRepo&#34;, IssueLabelArgs.builder()        
 *             .repository(&#34;test-repo&#34;)
 *             .name(&#34;Urgent&#34;)
 *             .color(&#34;FF0000&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * GitHub Issue Labels can be imported using an ID made up of `repository:name`, e.g.
 * 
 * ```sh
 * $ pulumi import github:index/issueLabel:IssueLabel panic_label terraform:panic
 * ```
 * 
 */
@ResourceType(type="github:index/issueLabel:IssueLabel")
public class IssueLabel extends com.pulumi.resources.CustomResource {
    /**
     * A 6 character hex code, **without the leading #**, identifying the color of the label.
     * 
     */
    @Export(name="color", refs={String.class}, tree="[0]")
    private Output<String> color;

    /**
     * @return A 6 character hex code, **without the leading #**, identifying the color of the label.
     * 
     */
    public Output<String> color() {
        return this.color;
    }
    /**
     * A short description of the label.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return A short description of the label.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    @Export(name="etag", refs={String.class}, tree="[0]")
    private Output<String> etag;

    public Output<String> etag() {
        return this.etag;
    }
    /**
     * The name of the label.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the label.
     * 
     */
    public Output<String> name() {
        return this.name;
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
     * The URL to the issue label
     * 
     */
    @Export(name="url", refs={String.class}, tree="[0]")
    private Output<String> url;

    /**
     * @return The URL to the issue label
     * 
     */
    public Output<String> url() {
        return this.url;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public IssueLabel(String name) {
        this(name, IssueLabelArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public IssueLabel(String name, IssueLabelArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public IssueLabel(String name, IssueLabelArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/issueLabel:IssueLabel", name, args == null ? IssueLabelArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private IssueLabel(String name, Output<String> id, @Nullable IssueLabelState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/issueLabel:IssueLabel", name, state, makeResourceOptions(options, id));
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
    public static IssueLabel get(String name, Output<String> id, @Nullable IssueLabelState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new IssueLabel(name, id, state, options);
    }
}
