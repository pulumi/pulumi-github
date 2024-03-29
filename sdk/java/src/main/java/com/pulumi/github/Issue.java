// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.github.IssueArgs;
import com.pulumi.github.Utilities;
import com.pulumi.github.inputs.IssueState;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a GitHub issue resource.
 * 
 * This resource allows you to create and manage issue within your
 * GitHub repository.
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
 * import com.pulumi.github.Repository;
 * import com.pulumi.github.RepositoryArgs;
 * import com.pulumi.github.Issue;
 * import com.pulumi.github.IssueArgs;
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
 *         // Create a simple issue
 *         var testRepository = new Repository(&#34;testRepository&#34;, RepositoryArgs.builder()        
 *             .autoInit(true)
 *             .hasIssues(true)
 *             .build());
 * 
 *         var testIssue = new Issue(&#34;testIssue&#34;, IssueArgs.builder()        
 *             .repository(testRepository.name())
 *             .title(&#34;My issue title&#34;)
 *             .body(&#34;The body of my issue&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ### With Milestone And Project Assignment
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.github.Repository;
 * import com.pulumi.github.RepositoryArgs;
 * import com.pulumi.github.RepositoryMilestone;
 * import com.pulumi.github.RepositoryMilestoneArgs;
 * import com.pulumi.github.Issue;
 * import com.pulumi.github.IssueArgs;
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
 *         // Create an issue with milestone and project assignment
 *         var testRepository = new Repository(&#34;testRepository&#34;, RepositoryArgs.builder()        
 *             .autoInit(true)
 *             .hasIssues(true)
 *             .build());
 * 
 *         var testRepositoryMilestone = new RepositoryMilestone(&#34;testRepositoryMilestone&#34;, RepositoryMilestoneArgs.builder()        
 *             .owner(testRepository.fullName().applyValue(fullName -&gt; fullName.split(&#34;/&#34;)).applyValue(split -&gt; split[0]))
 *             .repository(testRepository.name())
 *             .title(&#34;v1.0.0&#34;)
 *             .description(&#34;General Availability&#34;)
 *             .dueDate(&#34;2022-11-22&#34;)
 *             .state(&#34;open&#34;)
 *             .build());
 * 
 *         var testIssue = new Issue(&#34;testIssue&#34;, IssueArgs.builder()        
 *             .repository(testRepository.name())
 *             .title(&#34;My issue&#34;)
 *             .body(&#34;My issue body&#34;)
 *             .labels(            
 *                 &#34;bug&#34;,
 *                 &#34;documentation&#34;)
 *             .assignees(&#34;bob-github&#34;)
 *             .milestoneNumber(testRepositoryMilestone.number())
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * GitHub Issues can be imported using an ID made up of `repository:number`, e.g.
 * 
 * ```sh
 * $ pulumi import github:index/issue:Issue issue_15 myrepo:15
 * ```
 * 
 */
@ResourceType(type="github:index/issue:Issue")
public class Issue extends com.pulumi.resources.CustomResource {
    /**
     * List of Logins to assign the to the issue
     * 
     */
    @Export(name="assignees", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> assignees;

    /**
     * @return List of Logins to assign the to the issue
     * 
     */
    public Output<Optional<List<String>>> assignees() {
        return Codegen.optional(this.assignees);
    }
    /**
     * Body of the issue
     * 
     */
    @Export(name="body", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> body;

    /**
     * @return Body of the issue
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
     * (Computed) - The issue id
     * 
     */
    @Export(name="issueId", refs={Integer.class}, tree="[0]")
    private Output<Integer> issueId;

    /**
     * @return (Computed) - The issue id
     * 
     */
    public Output<Integer> issueId() {
        return this.issueId;
    }
    /**
     * List of labels to attach to the issue
     * 
     */
    @Export(name="labels", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> labels;

    /**
     * @return List of labels to attach to the issue
     * 
     */
    public Output<Optional<List<String>>> labels() {
        return Codegen.optional(this.labels);
    }
    /**
     * Milestone number to assign to the issue
     * 
     */
    @Export(name="milestoneNumber", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> milestoneNumber;

    /**
     * @return Milestone number to assign to the issue
     * 
     */
    public Output<Optional<Integer>> milestoneNumber() {
        return Codegen.optional(this.milestoneNumber);
    }
    /**
     * (Computed) - The issue number
     * 
     */
    @Export(name="number", refs={Integer.class}, tree="[0]")
    private Output<Integer> number;

    /**
     * @return (Computed) - The issue number
     * 
     */
    public Output<Integer> number() {
        return this.number;
    }
    /**
     * The GitHub repository name
     * 
     */
    @Export(name="repository", refs={String.class}, tree="[0]")
    private Output<String> repository;

    /**
     * @return The GitHub repository name
     * 
     */
    public Output<String> repository() {
        return this.repository;
    }
    /**
     * Title of the issue
     * 
     */
    @Export(name="title", refs={String.class}, tree="[0]")
    private Output<String> title;

    /**
     * @return Title of the issue
     * 
     */
    public Output<String> title() {
        return this.title;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Issue(String name) {
        this(name, IssueArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Issue(String name, IssueArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Issue(String name, IssueArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/issue:Issue", name, args == null ? IssueArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Issue(String name, Output<String> id, @Nullable IssueState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/issue:Issue", name, state, makeResourceOptions(options, id));
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
    public static Issue get(String name, Output<String> id, @Nullable IssueState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Issue(name, id, state, options);
    }
}
