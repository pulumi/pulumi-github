// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.github.IssueLabelsArgs;
import com.pulumi.github.Utilities;
import com.pulumi.github.inputs.IssueLabelsState;
import com.pulumi.github.outputs.IssueLabelsLabel;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides GitHub issue labels resource.
 * 
 * This resource allows you to create and manage issue labels within your
 * GitHub organization.
 * 
 * &gt; Note: github.IssueLabels cannot be used in conjunction with github.IssueLabel or they will fight over what your policy should be.
 * 
 * This resource is authoritative. For adding a label to a repo in a non-authoritative manner, use github.IssueLabel instead.
 * 
 * If you change the case of a label&#39;s name, its&#39; color, or description, this resource will edit the existing label to match the new values. However, if you change the name of a label, this resource will create a new label with the new name and delete the old label. Beware that this will remove the label from any issues it was previously attached to.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.github.IssueLabels;
 * import com.pulumi.github.IssueLabelsArgs;
 * import com.pulumi.github.inputs.IssueLabelsLabelArgs;
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
 *         var testRepo = new IssueLabels(&#34;testRepo&#34;, IssueLabelsArgs.builder()        
 *             .labels(            
 *                 IssueLabelsLabelArgs.builder()
 *                     .color(&#34;FF0000&#34;)
 *                     .name(&#34;Urgent&#34;)
 *                     .build(),
 *                 IssueLabelsLabelArgs.builder()
 *                     .color(&#34;FF0000&#34;)
 *                     .name(&#34;Critical&#34;)
 *                     .build())
 *             .repository(&#34;test-repo&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * GitHub Issue Labels can be imported using the repository `name`, e.g.
 * 
 * ```sh
 *  $ pulumi import github:index/issueLabels:IssueLabels test_repo test_repo
 * ```
 * 
 */
@ResourceType(type="github:index/issueLabels:IssueLabels")
public class IssueLabels extends com.pulumi.resources.CustomResource {
    /**
     * List of labels
     * 
     */
    @Export(name="labels", refs={List.class,IssueLabelsLabel.class}, tree="[0,1]")
    private Output</* @Nullable */ List<IssueLabelsLabel>> labels;

    /**
     * @return List of labels
     * 
     */
    public Output<Optional<List<IssueLabelsLabel>>> labels() {
        return Codegen.optional(this.labels);
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
    public IssueLabels(String name) {
        this(name, IssueLabelsArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public IssueLabels(String name, IssueLabelsArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public IssueLabels(String name, IssueLabelsArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/issueLabels:IssueLabels", name, args == null ? IssueLabelsArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private IssueLabels(String name, Output<String> id, @Nullable IssueLabelsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/issueLabels:IssueLabels", name, state, makeResourceOptions(options, id));
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
    public static IssueLabels get(String name, Output<String> id, @Nullable IssueLabelsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new IssueLabels(name, id, state, options);
    }
}
