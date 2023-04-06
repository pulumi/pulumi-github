// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.github.RepositoryCollaboratorArgs;
import com.pulumi.github.Utilities;
import com.pulumi.github.inputs.RepositoryCollaboratorState;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a GitHub repository collaborator resource.
 * 
 * &gt; Note: github.RepositoryCollaborator cannot be used in conjunction with github.RepositoryCollaborators or
 * they will fight over what your policy should be.
 * 
 * This resource allows you to add/remove collaborators from repositories in your
 * organization or personal account. For organization repositories, collaborators can
 * have explicit (and differing levels of) read, write, or administrator access to
 * specific repositories, without giving the user full organization membership.
 * For personal repositories, collaborators can only be granted write
 * (implicitly includes read) permission.
 * 
 * When applied, an invitation will be sent to the user to become a collaborator
 * on a repository. When destroyed, either the invitation will be cancelled or the
 * collaborator will be removed from the repository.
 * 
 * This resource is non-authoritative, for managing ALL collaborators of a repo, use github.RepositoryCollaborators
 * instead.
 * 
 * Further documentation on GitHub collaborators:
 * 
 * - [Adding outside collaborators to your personal repositories](https://help.github.com/en/github/setting-up-and-managing-your-github-user-account/managing-access-to-your-personal-repositories)
 * - [Adding outside collaborators to repositories in your organization](https://help.github.com/articles/adding-outside-collaborators-to-repositories-in-your-organization/)
 * - [Converting an organization member to an outside collaborator](https://help.github.com/articles/converting-an-organization-member-to-an-outside-collaborator/)
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.github.RepositoryCollaborator;
 * import com.pulumi.github.RepositoryCollaboratorArgs;
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
 *         var aRepoCollaborator = new RepositoryCollaborator(&#34;aRepoCollaborator&#34;, RepositoryCollaboratorArgs.builder()        
 *             .permission(&#34;admin&#34;)
 *             .repository(&#34;our-cool-repo&#34;)
 *             .username(&#34;SomeUser&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * GitHub Repository Collaborators can be imported using an ID made up of `repository:username`, e.g.
 * 
 * ```sh
 *  $ pulumi import github:index/repositoryCollaborator:RepositoryCollaborator collaborator terraform:someuser
 * ```
 * 
 */
@ResourceType(type="github:index/repositoryCollaborator:RepositoryCollaborator")
public class RepositoryCollaborator extends com.pulumi.resources.CustomResource {
    /**
     * ID of the invitation to be used in `github.UserInvitationAccepter`
     * 
     */
    @Export(name="invitationId", type=String.class, parameters={})
    private Output<String> invitationId;

    /**
     * @return ID of the invitation to be used in `github.UserInvitationAccepter`
     * 
     */
    public Output<String> invitationId() {
        return this.invitationId;
    }
    /**
     * The permission of the outside collaborator for the repository.
     * Must be one of `pull`, `push`, `maintain`, `triage` or `admin` or the name of an existing [custom repository role](https://docs.github.com/en/enterprise-cloud@latest/organizations/managing-peoples-access-to-your-organization-with-roles/managing-custom-repository-roles-for-an-organization) within the organization for organization-owned repositories.
     * Must be `push` for personal repositories. Defaults to `push`.
     * 
     */
    @Export(name="permission", type=String.class, parameters={})
    private Output</* @Nullable */ String> permission;

    /**
     * @return The permission of the outside collaborator for the repository.
     * Must be one of `pull`, `push`, `maintain`, `triage` or `admin` or the name of an existing [custom repository role](https://docs.github.com/en/enterprise-cloud@latest/organizations/managing-peoples-access-to-your-organization-with-roles/managing-custom-repository-roles-for-an-organization) within the organization for organization-owned repositories.
     * Must be `push` for personal repositories. Defaults to `push`.
     * 
     */
    public Output<Optional<String>> permission() {
        return Codegen.optional(this.permission);
    }
    /**
     * Suppress plan diffs for `triage` and `maintain`.  Defaults to `false`.
     * 
     */
    @Export(name="permissionDiffSuppression", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> permissionDiffSuppression;

    /**
     * @return Suppress plan diffs for `triage` and `maintain`.  Defaults to `false`.
     * 
     */
    public Output<Optional<Boolean>> permissionDiffSuppression() {
        return Codegen.optional(this.permissionDiffSuppression);
    }
    /**
     * The GitHub repository
     * 
     */
    @Export(name="repository", type=String.class, parameters={})
    private Output<String> repository;

    /**
     * @return The GitHub repository
     * 
     */
    public Output<String> repository() {
        return this.repository;
    }
    /**
     * The user to add to the repository as a collaborator.
     * 
     */
    @Export(name="username", type=String.class, parameters={})
    private Output<String> username;

    /**
     * @return The user to add to the repository as a collaborator.
     * 
     */
    public Output<String> username() {
        return this.username;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public RepositoryCollaborator(String name) {
        this(name, RepositoryCollaboratorArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public RepositoryCollaborator(String name, RepositoryCollaboratorArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public RepositoryCollaborator(String name, RepositoryCollaboratorArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/repositoryCollaborator:RepositoryCollaborator", name, args == null ? RepositoryCollaboratorArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private RepositoryCollaborator(String name, Output<String> id, @Nullable RepositoryCollaboratorState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/repositoryCollaborator:RepositoryCollaborator", name, state, makeResourceOptions(options, id));
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
    public static RepositoryCollaborator get(String name, Output<String> id, @Nullable RepositoryCollaboratorState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new RepositoryCollaborator(name, id, state, options);
    }
}
