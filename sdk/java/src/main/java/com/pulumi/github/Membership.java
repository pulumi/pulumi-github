// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.github.MembershipArgs;
import com.pulumi.github.Utilities;
import com.pulumi.github.inputs.MembershipState;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a GitHub membership resource.
 * 
 * This resource allows you to add/remove users from your organization. When applied,
 * an invitation will be sent to the user to become part of the organization. When
 * destroyed, either the invitation will be cancelled or the user will be removed.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.github.Membership;
 * import com.pulumi.github.MembershipArgs;
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
 *         var membershipForSomeUser = new Membership(&#34;membershipForSomeUser&#34;, MembershipArgs.builder()        
 *             .role(&#34;member&#34;)
 *             .username(&#34;SomeUser&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * GitHub Membership can be imported using an ID made up of `organization:username`, e.g.
 * 
 * ```sh
 *  $ pulumi import github:index/membership:Membership member hashicorp:someuser
 * ```
 * 
 */
@ResourceType(type="github:index/membership:Membership")
public class Membership extends com.pulumi.resources.CustomResource {
    /**
     * Defaults to `false`. If set to true,
     * when this resource is destroyed, the member will not be removed
     * from the organization. Instead, the member&#39;s role will be
     * downgraded to &#39;member&#39;.
     * 
     */
    @Export(name="downgradeOnDestroy", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> downgradeOnDestroy;

    /**
     * @return Defaults to `false`. If set to true,
     * when this resource is destroyed, the member will not be removed
     * from the organization. Instead, the member&#39;s role will be
     * downgraded to &#39;member&#39;.
     * 
     */
    public Output<Optional<Boolean>> downgradeOnDestroy() {
        return Codegen.optional(this.downgradeOnDestroy);
    }
    @Export(name="etag", type=String.class, parameters={})
    private Output<String> etag;

    public Output<String> etag() {
        return this.etag;
    }
    /**
     * The role of the user within the organization.
     * Must be one of `member` or `admin`. Defaults to `member`.
     * 
     */
    @Export(name="role", type=String.class, parameters={})
    private Output</* @Nullable */ String> role;

    /**
     * @return The role of the user within the organization.
     * Must be one of `member` or `admin`. Defaults to `member`.
     * 
     */
    public Output<Optional<String>> role() {
        return Codegen.optional(this.role);
    }
    /**
     * The user to add to the organization.
     * 
     */
    @Export(name="username", type=String.class, parameters={})
    private Output<String> username;

    /**
     * @return The user to add to the organization.
     * 
     */
    public Output<String> username() {
        return this.username;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Membership(String name) {
        this(name, MembershipArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Membership(String name, MembershipArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Membership(String name, MembershipArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/membership:Membership", name, args == null ? MembershipArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Membership(String name, Output<String> id, @Nullable MembershipState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/membership:Membership", name, state, makeResourceOptions(options, id));
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
    public static Membership get(String name, Output<String> id, @Nullable MembershipState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Membership(name, id, state, options);
    }
}
