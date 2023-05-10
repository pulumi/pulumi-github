// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.github.TeamRepositoryArgs;
import com.pulumi.github.Utilities;
import com.pulumi.github.inputs.TeamRepositoryState;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * &gt; Note: github.TeamRepository cannot be used in conjunction with github.RepositoryCollaborators or
 * they will fight over what your policy should be.
 * 
 * This resource manages relationships between teams and repositories
 * in your GitHub organization.
 * 
 * Creating this resource grants a particular team permissions on a
 * particular repository.
 * 
 * The repository and the team must both belong to the same organization
 * on GitHub. This resource does not actually *create* any repositories;
 * to do that, see `github.Repository`.
 * 
 * This resource is non-authoritative, for managing ALL collaborators of a repo, use github.RepositoryCollaborators
 * instead.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.github.Team;
 * import com.pulumi.github.TeamArgs;
 * import com.pulumi.github.Repository;
 * import com.pulumi.github.TeamRepository;
 * import com.pulumi.github.TeamRepositoryArgs;
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
 *         var someTeam = new Team(&#34;someTeam&#34;, TeamArgs.builder()        
 *             .description(&#34;Some cool team&#34;)
 *             .build());
 * 
 *         var someRepo = new Repository(&#34;someRepo&#34;);
 * 
 *         var someTeamRepo = new TeamRepository(&#34;someTeamRepo&#34;, TeamRepositoryArgs.builder()        
 *             .teamId(someTeam.id())
 *             .repository(someRepo.name())
 *             .permission(&#34;pull&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * GitHub Team Repository can be imported using an ID made up of `team_id:repository` or `team_name:repository`, e.g.
 * 
 * ```sh
 *  $ pulumi import github:index/teamRepository:TeamRepository terraform_repo 1234567:terraform
 * ```
 * 
 * ```sh
 *  $ pulumi import github:index/teamRepository:TeamRepository terraform_repo Administrators:terraform
 * ```
 * 
 */
@ResourceType(type="github:index/teamRepository:TeamRepository")
public class TeamRepository extends com.pulumi.resources.CustomResource {
    @Export(name="etag", type=String.class, parameters={})
    private Output<String> etag;

    public Output<String> etag() {
        return this.etag;
    }
    /**
     * The permissions of team members regarding the repository.
     * Must be one of `pull`, `triage`, `push`, `maintain`, `admin` or the name of an existing [custom repository role](https://docs.github.com/en/enterprise-cloud@latest/organizations/managing-peoples-access-to-your-organization-with-roles/managing-custom-repository-roles-for-an-organization) within the organisation. Defaults to `pull`.
     * 
     */
    @Export(name="permission", type=String.class, parameters={})
    private Output</* @Nullable */ String> permission;

    /**
     * @return The permissions of team members regarding the repository.
     * Must be one of `pull`, `triage`, `push`, `maintain`, `admin` or the name of an existing [custom repository role](https://docs.github.com/en/enterprise-cloud@latest/organizations/managing-peoples-access-to-your-organization-with-roles/managing-custom-repository-roles-for-an-organization) within the organisation. Defaults to `pull`.
     * 
     */
    public Output<Optional<String>> permission() {
        return Codegen.optional(this.permission);
    }
    /**
     * The repository to add to the team.
     * 
     */
    @Export(name="repository", type=String.class, parameters={})
    private Output<String> repository;

    /**
     * @return The repository to add to the team.
     * 
     */
    public Output<String> repository() {
        return this.repository;
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
    public TeamRepository(String name) {
        this(name, TeamRepositoryArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public TeamRepository(String name, TeamRepositoryArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public TeamRepository(String name, TeamRepositoryArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/teamRepository:TeamRepository", name, args == null ? TeamRepositoryArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private TeamRepository(String name, Output<String> id, @Nullable TeamRepositoryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/teamRepository:TeamRepository", name, state, makeResourceOptions(options, id));
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
    public static TeamRepository get(String name, Output<String> id, @Nullable TeamRepositoryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new TeamRepository(name, id, state, options);
    }
}
