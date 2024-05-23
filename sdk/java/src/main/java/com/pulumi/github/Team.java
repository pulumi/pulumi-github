// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.github.TeamArgs;
import com.pulumi.github.Utilities;
import com.pulumi.github.inputs.TeamState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a GitHub team resource.
 * 
 * This resource allows you to add/remove teams from your organization. When applied,
 * a new team will be created. When destroyed, that team will be removed.
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
 * import com.pulumi.github.Team;
 * import com.pulumi.github.TeamArgs;
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
 *         // Add a team to the organization
 *         var someTeam = new Team("someTeam", TeamArgs.builder()
 *             .name("some-team")
 *             .description("Some cool team")
 *             .privacy("closed")
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
 * GitHub Teams can be imported using the GitHub team ID or name e.g.
 * 
 * ```sh
 * $ pulumi import github:index/team:Team core 1234567
 * ```
 * 
 * ```sh
 * $ pulumi import github:index/team:Team core Administrators
 * ```
 * 
 */
@ResourceType(type="github:index/team:Team")
public class Team extends com.pulumi.resources.CustomResource {
    /**
     * Adds a default maintainer to the team. Defaults to `false` and adds the creating user to the team when `true`.
     * 
     */
    @Export(name="createDefaultMaintainer", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> createDefaultMaintainer;

    /**
     * @return Adds a default maintainer to the team. Defaults to `false` and adds the creating user to the team when `true`.
     * 
     */
    public Output<Optional<Boolean>> createDefaultMaintainer() {
        return Codegen.optional(this.createDefaultMaintainer);
    }
    /**
     * A description of the team.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return A description of the team.
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
     * The LDAP Distinguished Name of the group where membership will be synchronized. Only available in GitHub Enterprise Server.
     * 
     */
    @Export(name="ldapDn", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> ldapDn;

    /**
     * @return The LDAP Distinguished Name of the group where membership will be synchronized. Only available in GitHub Enterprise Server.
     * 
     */
    public Output<Optional<String>> ldapDn() {
        return Codegen.optional(this.ldapDn);
    }
    @Export(name="membersCount", refs={Integer.class}, tree="[0]")
    private Output<Integer> membersCount;

    public Output<Integer> membersCount() {
        return this.membersCount;
    }
    /**
     * The name of the team.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the team.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The Node ID of the created team.
     * 
     */
    @Export(name="nodeId", refs={String.class}, tree="[0]")
    private Output<String> nodeId;

    /**
     * @return The Node ID of the created team.
     * 
     */
    public Output<String> nodeId() {
        return this.nodeId;
    }
    /**
     * The ID or slug of the parent team, if this is a nested team.
     * 
     */
    @Export(name="parentTeamId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> parentTeamId;

    /**
     * @return The ID or slug of the parent team, if this is a nested team.
     * 
     */
    public Output<Optional<String>> parentTeamId() {
        return Codegen.optional(this.parentTeamId);
    }
    /**
     * The id of the parent team read in Github.
     * 
     */
    @Export(name="parentTeamReadId", refs={String.class}, tree="[0]")
    private Output<String> parentTeamReadId;

    /**
     * @return The id of the parent team read in Github.
     * 
     */
    public Output<String> parentTeamReadId() {
        return this.parentTeamReadId;
    }
    /**
     * The id of the parent team read in Github.
     * 
     */
    @Export(name="parentTeamReadSlug", refs={String.class}, tree="[0]")
    private Output<String> parentTeamReadSlug;

    /**
     * @return The id of the parent team read in Github.
     * 
     */
    public Output<String> parentTeamReadSlug() {
        return this.parentTeamReadSlug;
    }
    /**
     * The level of privacy for the team. Must be one of `secret` or `closed`.
     * Defaults to `secret`.
     * 
     */
    @Export(name="privacy", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> privacy;

    /**
     * @return The level of privacy for the team. Must be one of `secret` or `closed`.
     * Defaults to `secret`.
     * 
     */
    public Output<Optional<String>> privacy() {
        return Codegen.optional(this.privacy);
    }
    /**
     * The slug of the created team, which may or may not differ from `name`,
     * depending on whether `name` contains &#34;URL-unsafe&#34; characters.
     * Useful when referencing the team in [`github.BranchProtection`](https://www.terraform.io/docs/providers/github/r/branch_protection.html).
     * 
     */
    @Export(name="slug", refs={String.class}, tree="[0]")
    private Output<String> slug;

    /**
     * @return The slug of the created team, which may or may not differ from `name`,
     * depending on whether `name` contains &#34;URL-unsafe&#34; characters.
     * Useful when referencing the team in [`github.BranchProtection`](https://www.terraform.io/docs/providers/github/r/branch_protection.html).
     * 
     */
    public Output<String> slug() {
        return this.slug;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Team(String name) {
        this(name, TeamArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Team(String name, @Nullable TeamArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Team(String name, @Nullable TeamArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/team:Team", name, args == null ? TeamArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Team(String name, Output<String> id, @Nullable TeamState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/team:Team", name, state, makeResourceOptions(options, id));
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
    public static Team get(String name, Output<String> id, @Nullable TeamState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Team(name, id, state, options);
    }
}
