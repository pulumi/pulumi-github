// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.github.TeamMembersArgs;
import com.pulumi.github.Utilities;
import com.pulumi.github.inputs.TeamMembersState;
import com.pulumi.github.outputs.TeamMembersMember;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

/**
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
 * import com.pulumi.github.Membership;
 * import com.pulumi.github.MembershipArgs;
 * import com.pulumi.github.Team;
 * import com.pulumi.github.TeamArgs;
 * import com.pulumi.github.TeamMembers;
 * import com.pulumi.github.TeamMembersArgs;
 * import com.pulumi.github.inputs.TeamMembersMemberArgs;
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
 *         // Add a user to the organization
 *         var membershipForSomeUser = new Membership("membershipForSomeUser", MembershipArgs.builder()
 *             .username("SomeUser")
 *             .role("member")
 *             .build());
 * 
 *         var membershipForAnotherUser = new Membership("membershipForAnotherUser", MembershipArgs.builder()
 *             .username("AnotherUser")
 *             .role("member")
 *             .build());
 * 
 *         var someTeam = new Team("someTeam", TeamArgs.builder()
 *             .name("SomeTeam")
 *             .description("Some cool team")
 *             .build());
 * 
 *         var someTeamMembers = new TeamMembers("someTeamMembers", TeamMembersArgs.builder()
 *             .teamId(someTeam.id())
 *             .members(            
 *                 TeamMembersMemberArgs.builder()
 *                     .username("SomeUser")
 *                     .role("maintainer")
 *                     .build(),
 *                 TeamMembersMemberArgs.builder()
 *                     .username("AnotherUser")
 *                     .role("member")
 *                     .build())
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
 * ~&gt; **Note** Although the team id or team slug can be used it is recommended to use the team id.  Using the team slug will result in terraform doing conversions between the team slug and team id.  This will cause team members associations to the team to be destroyed and recreated on import.
 * 
 * GitHub Team Membership can be imported using the team ID team id or team slug, e.g.
 * 
 * ```sh
 * $ pulumi import github:index/teamMembers:TeamMembers some_team 1234567
 * ```
 * 
 * ```sh
 * $ pulumi import github:index/teamMembers:TeamMembers some_team Administrators
 * ```
 * 
 */
@ResourceType(type="github:index/teamMembers:TeamMembers")
public class TeamMembers extends com.pulumi.resources.CustomResource {
    /**
     * List of team members. See Members below for details.
     * 
     */
    @Export(name="members", refs={List.class,TeamMembersMember.class}, tree="[0,1]")
    private Output<List<TeamMembersMember>> members;

    /**
     * @return List of team members. See Members below for details.
     * 
     */
    public Output<List<TeamMembersMember>> members() {
        return this.members;
    }
    /**
     * The team id or the team slug
     * 
     * &gt; **Note** Although the team id or team slug can be used it is recommended to use the team id.  Using the team slug will cause the team members associations to the team to be destroyed and recreated if the team name is updated.
     * 
     */
    @Export(name="teamId", refs={String.class}, tree="[0]")
    private Output<String> teamId;

    /**
     * @return The team id or the team slug
     * 
     * &gt; **Note** Although the team id or team slug can be used it is recommended to use the team id.  Using the team slug will cause the team members associations to the team to be destroyed and recreated if the team name is updated.
     * 
     */
    public Output<String> teamId() {
        return this.teamId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public TeamMembers(java.lang.String name) {
        this(name, TeamMembersArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public TeamMembers(java.lang.String name, TeamMembersArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public TeamMembers(java.lang.String name, TeamMembersArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/teamMembers:TeamMembers", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private TeamMembers(java.lang.String name, Output<java.lang.String> id, @Nullable TeamMembersState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/teamMembers:TeamMembers", name, state, makeResourceOptions(options, id), false);
    }

    private static TeamMembersArgs makeArgs(TeamMembersArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? TeamMembersArgs.Empty : args;
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
    public static TeamMembers get(java.lang.String name, Output<java.lang.String> id, @Nullable TeamMembersState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new TeamMembers(name, id, state, options);
    }
}
