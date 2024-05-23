// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.github.RepositoryCollaboratorsArgs;
import com.pulumi.github.Utilities;
import com.pulumi.github.inputs.RepositoryCollaboratorsState;
import com.pulumi.github.outputs.RepositoryCollaboratorsTeam;
import com.pulumi.github.outputs.RepositoryCollaboratorsUser;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a GitHub repository collaborators resource.
 * 
 * &gt; Note: github.RepositoryCollaborators cannot be used in conjunction with github.RepositoryCollaborator and
 * github.TeamRepository or they will fight over what your policy should be.
 * 
 * This resource allows you to manage all collaborators for repositories in your
 * organization or personal account. For organization repositories, collaborators can
 * have explicit (and differing levels of) read, write, or administrator access to
 * specific repositories, without giving the user full organization membership.
 * For personal repositories, collaborators can only be granted write
 * (implicitly includes read) permission.
 * 
 * When applied, an invitation will be sent to the user to become a collaborators
 * on a repository. When destroyed, either the invitation will be cancelled or the
 * collaborators will be removed from the repository.
 * 
 * This resource is authoritative. For adding a collaborator to a repo in a non-authoritative manner, use
 * github.RepositoryCollaborator instead.
 * 
 * Further documentation on GitHub collaborators:
 * 
 * - [Adding outside collaborators to your personal repositories](https://help.github.com/en/github/setting-up-and-managing-your-github-user-account/managing-access-to-your-personal-repositories)
 * - [Adding outside collaborators to repositories in your organization](https://help.github.com/articles/adding-outside-collaborators-to-repositories-in-your-organization/)
 * - [Converting an organization member to an outside collaborators](https://help.github.com/articles/converting-an-organization-member-to-an-outside-collaborator/)
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
 * import com.pulumi.github.Repository;
 * import com.pulumi.github.RepositoryArgs;
 * import com.pulumi.github.RepositoryCollaborators;
 * import com.pulumi.github.RepositoryCollaboratorsArgs;
 * import com.pulumi.github.inputs.RepositoryCollaboratorsUserArgs;
 * import com.pulumi.github.inputs.RepositoryCollaboratorsTeamArgs;
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
 *         // Add collaborators to a repository
 *         var someTeam = new Team("someTeam", TeamArgs.builder()
 *             .name("SomeTeam")
 *             .description("Some cool team")
 *             .build());
 * 
 *         var someRepo = new Repository("someRepo", RepositoryArgs.builder()
 *             .name("some-repo")
 *             .build());
 * 
 *         var someRepoCollaborators = new RepositoryCollaborators("someRepoCollaborators", RepositoryCollaboratorsArgs.builder()
 *             .repository(someRepo.name())
 *             .users(RepositoryCollaboratorsUserArgs.builder()
 *                 .permission("admin")
 *                 .username("SomeUser")
 *                 .build())
 *             .teams(RepositoryCollaboratorsTeamArgs.builder()
 *                 .permission("pull")
 *                 .teamId(someTeam.slug())
 *                 .build())
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
 * GitHub Repository Collaborators can be imported using the name `name`, e.g.
 * 
 * ```sh
 * $ pulumi import github:index/repositoryCollaborators:RepositoryCollaborators collaborators terraform
 * ```
 * 
 */
@ResourceType(type="github:index/repositoryCollaborators:RepositoryCollaborators")
public class RepositoryCollaborators extends com.pulumi.resources.CustomResource {
    /**
     * Map of usernames to invitation ID for any users added as part of creation of this resource to
     * be used in `github.UserInvitationAccepter`.
     * 
     */
    @Export(name="invitationIds", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> invitationIds;

    /**
     * @return Map of usernames to invitation ID for any users added as part of creation of this resource to
     * be used in `github.UserInvitationAccepter`.
     * 
     */
    public Output<Map<String,String>> invitationIds() {
        return this.invitationIds;
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
     * List of teams
     * 
     */
    @Export(name="teams", refs={List.class,RepositoryCollaboratorsTeam.class}, tree="[0,1]")
    private Output</* @Nullable */ List<RepositoryCollaboratorsTeam>> teams;

    /**
     * @return List of teams
     * 
     */
    public Output<Optional<List<RepositoryCollaboratorsTeam>>> teams() {
        return Codegen.optional(this.teams);
    }
    /**
     * List of users
     * 
     */
    @Export(name="users", refs={List.class,RepositoryCollaboratorsUser.class}, tree="[0,1]")
    private Output</* @Nullable */ List<RepositoryCollaboratorsUser>> users;

    /**
     * @return List of users
     * 
     */
    public Output<Optional<List<RepositoryCollaboratorsUser>>> users() {
        return Codegen.optional(this.users);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public RepositoryCollaborators(String name) {
        this(name, RepositoryCollaboratorsArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public RepositoryCollaborators(String name, RepositoryCollaboratorsArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public RepositoryCollaborators(String name, RepositoryCollaboratorsArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/repositoryCollaborators:RepositoryCollaborators", name, args == null ? RepositoryCollaboratorsArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private RepositoryCollaborators(String name, Output<String> id, @Nullable RepositoryCollaboratorsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/repositoryCollaborators:RepositoryCollaborators", name, state, makeResourceOptions(options, id));
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
    public static RepositoryCollaborators get(String name, Output<String> id, @Nullable RepositoryCollaboratorsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new RepositoryCollaborators(name, id, state, options);
    }
}
