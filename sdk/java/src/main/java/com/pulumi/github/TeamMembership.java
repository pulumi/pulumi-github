// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.github.TeamMembershipArgs;
import com.pulumi.github.Utilities;
import com.pulumi.github.inputs.TeamMembershipState;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a GitHub team membership resource.
 * 
 * This resource allows you to add/remove users from teams in your organization. When applied,
 * the user will be added to the team. If the user hasn&#39;t accepted their invitation to the
 * organization, they won&#39;t be part of the team until they do. When
 * destroyed, the user will be removed from the team.
 * 
 * &gt; **Note** This resource is not compatible with `github.TeamMembers`. Use either `github.TeamMembers` or `github.TeamMembership`.
 * 
 * &gt; **Note** Organization owners may not be set as &#34;members&#34; of a team; they may only be set as &#34;maintainers&#34;. Attempting to set organization an owner to &#34;member&#34; of a may result in a `pulumi preview` diff that changes their status back to &#34;maintainer&#34;.
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
 * import com.pulumi.github.Membership;
 * import com.pulumi.github.MembershipArgs;
 * import com.pulumi.github.Team;
 * import com.pulumi.github.TeamArgs;
 * import com.pulumi.github.TeamMembership;
 * import com.pulumi.github.TeamMembershipArgs;
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
 *             .username(&#34;SomeUser&#34;)
 *             .role(&#34;member&#34;)
 *             .build());
 * 
 *         var someTeam = new Team(&#34;someTeam&#34;, TeamArgs.builder()        
 *             .description(&#34;Some cool team&#34;)
 *             .build());
 * 
 *         var someTeamMembership = new TeamMembership(&#34;someTeamMembership&#34;, TeamMembershipArgs.builder()        
 *             .teamId(someTeam.id())
 *             .username(&#34;SomeUser&#34;)
 *             .role(&#34;member&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * GitHub Team Membership can be imported using an ID made up of `teamid:username` or `teamname:username`, e.g.
 * 
 * ```sh
 * $ pulumi import github:index/teamMembership:TeamMembership member 1234567:someuser
 * ```
 * 
 * ```sh
 * $ pulumi import github:index/teamMembership:TeamMembership member Administrators:someuser
 * ```
 * 
 */
@ResourceType(type="github:index/teamMembership:TeamMembership")
public class TeamMembership extends com.pulumi.resources.CustomResource {
    @Export(name="etag", refs={String.class}, tree="[0]")
    private Output<String> etag;

    public Output<String> etag() {
        return this.etag;
    }
    /**
     * The role of the user within the team.
     * Must be one of `member` or `maintainer`. Defaults to `member`.
     * 
     */
    @Export(name="role", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> role;

    /**
     * @return The role of the user within the team.
     * Must be one of `member` or `maintainer`. Defaults to `member`.
     * 
     */
    public Output<Optional<String>> role() {
        return Codegen.optional(this.role);
    }
    /**
     * The GitHub team id or the GitHub team slug
     * 
     */
    @Export(name="teamId", refs={String.class}, tree="[0]")
    private Output<String> teamId;

    /**
     * @return The GitHub team id or the GitHub team slug
     * 
     */
    public Output<String> teamId() {
        return this.teamId;
    }
    /**
     * The user to add to the team.
     * 
     */
    @Export(name="username", refs={String.class}, tree="[0]")
    private Output<String> username;

    /**
     * @return The user to add to the team.
     * 
     */
    public Output<String> username() {
        return this.username;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public TeamMembership(String name) {
        this(name, TeamMembershipArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public TeamMembership(String name, TeamMembershipArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public TeamMembership(String name, TeamMembershipArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/teamMembership:TeamMembership", name, args == null ? TeamMembershipArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private TeamMembership(String name, Output<String> id, @Nullable TeamMembershipState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/teamMembership:TeamMembership", name, state, makeResourceOptions(options, id));
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
    public static TeamMembership get(String name, Output<String> id, @Nullable TeamMembershipState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new TeamMembership(name, id, state, options);
    }
}
