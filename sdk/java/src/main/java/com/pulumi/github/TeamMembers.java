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
 *         var membershipForSomeUser = new Membership(&#34;membershipForSomeUser&#34;, MembershipArgs.builder()        
 *             .username(&#34;SomeUser&#34;)
 *             .role(&#34;member&#34;)
 *             .build());
 * 
 *         var membershipForAnotherUser = new Membership(&#34;membershipForAnotherUser&#34;, MembershipArgs.builder()        
 *             .username(&#34;AnotherUser&#34;)
 *             .role(&#34;member&#34;)
 *             .build());
 * 
 *         var someTeam = new Team(&#34;someTeam&#34;, TeamArgs.builder()        
 *             .description(&#34;Some cool team&#34;)
 *             .build());
 * 
 *         var someTeamMembers = new TeamMembers(&#34;someTeamMembers&#34;, TeamMembersArgs.builder()        
 *             .teamId(someTeam.id())
 *             .members(            
 *                 TeamMembersMemberArgs.builder()
 *                     .username(&#34;SomeUser&#34;)
 *                     .role(&#34;maintainer&#34;)
 *                     .build(),
 *                 TeamMembersMemberArgs.builder()
 *                     .username(&#34;AnotherUser&#34;)
 *                     .role(&#34;member&#34;)
 *                     .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * GitHub Team Membership can be imported using the team ID `teamid` or team name, e.g.
 * 
 * ```sh
 *  $ pulumi import github:index/teamMembers:TeamMembers some_team 1234567
 * ```
 * 
 * ```sh
 *  $ pulumi import github:index/teamMembers:TeamMembers some_team Administrators
 * ```
 * 
 */
@ResourceType(type="github:index/teamMembers:TeamMembers")
public class TeamMembers extends com.pulumi.resources.CustomResource {
    @Export(name="etag", type=String.class, parameters={})
    private Output<String> etag;

    public Output<String> etag() {
        return this.etag;
    }
    /**
     * List of team members. See Members below for details.
     * 
     */
    @Export(name="members", type=List.class, parameters={TeamMembersMember.class})
    private Output<List<TeamMembersMember>> members;

    /**
     * @return List of team members. See Members below for details.
     * 
     */
    public Output<List<TeamMembersMember>> members() {
        return this.members;
    }
    /**
     * The GitHub team id or the GitHub team slug
     * 
     */
    @Export(name="teamId", type=String.class, parameters={})
    private Output<String> teamId;

    /**
     * @return The GitHub team id or the GitHub team slug
     * 
     */
    public Output<String> teamId() {
        return this.teamId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public TeamMembers(String name) {
        this(name, TeamMembersArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public TeamMembers(String name, TeamMembersArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public TeamMembers(String name, TeamMembersArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/teamMembers:TeamMembers", name, args == null ? TeamMembersArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private TeamMembers(String name, Output<String> id, @Nullable TeamMembersState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/teamMembers:TeamMembers", name, state, makeResourceOptions(options, id));
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
    public static TeamMembers get(String name, Output<String> id, @Nullable TeamMembersState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new TeamMembers(name, id, state, options);
    }
}
