// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.github.UserSshKeyArgs;
import com.pulumi.github.Utilities;
import com.pulumi.github.inputs.UserSshKeyState;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Provides a GitHub user&#39;s SSH key resource.
 * 
 * This resource allows you to add/remove SSH keys from your user account.
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
 * import com.pulumi.github.UserSshKey;
 * import com.pulumi.github.UserSshKeyArgs;
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
 *         var example = new UserSshKey(&#34;example&#34;, UserSshKeyArgs.builder()        
 *             .title(&#34;example title&#34;)
 *             .key(Files.readString(Paths.get(&#34;~/.ssh/id_rsa.pub&#34;)))
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * SSH keys can be imported using their ID e.g.
 * 
 * ```sh
 * $ pulumi import github:index/userSshKey:UserSshKey example 1234567
 * ```
 * 
 */
@ResourceType(type="github:index/userSshKey:UserSshKey")
public class UserSshKey extends com.pulumi.resources.CustomResource {
    @Export(name="etag", refs={String.class}, tree="[0]")
    private Output<String> etag;

    public Output<String> etag() {
        return this.etag;
    }
    /**
     * The public SSH key to add to your GitHub account.
     * 
     */
    @Export(name="key", refs={String.class}, tree="[0]")
    private Output<String> key;

    /**
     * @return The public SSH key to add to your GitHub account.
     * 
     */
    public Output<String> key() {
        return this.key;
    }
    /**
     * A descriptive name for the new key. e.g. `Personal MacBook Air`
     * 
     */
    @Export(name="title", refs={String.class}, tree="[0]")
    private Output<String> title;

    /**
     * @return A descriptive name for the new key. e.g. `Personal MacBook Air`
     * 
     */
    public Output<String> title() {
        return this.title;
    }
    /**
     * The URL of the SSH key
     * 
     */
    @Export(name="url", refs={String.class}, tree="[0]")
    private Output<String> url;

    /**
     * @return The URL of the SSH key
     * 
     */
    public Output<String> url() {
        return this.url;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public UserSshKey(String name) {
        this(name, UserSshKeyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public UserSshKey(String name, UserSshKeyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public UserSshKey(String name, UserSshKeyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/userSshKey:UserSshKey", name, args == null ? UserSshKeyArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private UserSshKey(String name, Output<String> id, @Nullable UserSshKeyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/userSshKey:UserSshKey", name, state, makeResourceOptions(options, id));
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
    public static UserSshKey get(String name, Output<String> id, @Nullable UserSshKeyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new UserSshKey(name, id, state, options);
    }
}
