// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Provides a GitHub repository collaborator resource.
 *
 * > Note: github.RepositoryCollaborator cannot be used in conjunction with github.RepositoryCollaborators or
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
 * ## Import
 *
 * GitHub Repository Collaborators can be imported using an ID made up of `repository:username`, e.g.
 *
 * ```sh
 *  $ pulumi import github:index/repositoryCollaborator:RepositoryCollaborator collaborator terraform:someuser
 * ```
 */
export class RepositoryCollaborator extends pulumi.CustomResource {
    /**
     * Get an existing RepositoryCollaborator resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RepositoryCollaboratorState, opts?: pulumi.CustomResourceOptions): RepositoryCollaborator {
        return new RepositoryCollaborator(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'github:index/repositoryCollaborator:RepositoryCollaborator';

    /**
     * Returns true if the given object is an instance of RepositoryCollaborator.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RepositoryCollaborator {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RepositoryCollaborator.__pulumiType;
    }

    /**
     * ID of the invitation to be used in `github.UserInvitationAccepter`
     */
    public /*out*/ readonly invitationId!: pulumi.Output<string>;
    /**
     * The permission of the outside collaborator for the repository.
     * Must be one of `pull`, `push`, `maintain`, `triage` or `admin` or the name of an existing [custom repository role](https://docs.github.com/en/enterprise-cloud@latest/organizations/managing-peoples-access-to-your-organization-with-roles/managing-custom-repository-roles-for-an-organization) within the organization for organization-owned repositories.
     * Must be `push` for personal repositories. Defaults to `push`.
     */
    public readonly permission!: pulumi.Output<string | undefined>;
    /**
     * Suppress plan diffs for `triage` and `maintain`.  Defaults to `false`.
     */
    public readonly permissionDiffSuppression!: pulumi.Output<boolean | undefined>;
    /**
     * The GitHub repository
     */
    public readonly repository!: pulumi.Output<string>;
    /**
     * The user to add to the repository as a collaborator.
     */
    public readonly username!: pulumi.Output<string>;

    /**
     * Create a RepositoryCollaborator resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RepositoryCollaboratorArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RepositoryCollaboratorArgs | RepositoryCollaboratorState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RepositoryCollaboratorState | undefined;
            resourceInputs["invitationId"] = state ? state.invitationId : undefined;
            resourceInputs["permission"] = state ? state.permission : undefined;
            resourceInputs["permissionDiffSuppression"] = state ? state.permissionDiffSuppression : undefined;
            resourceInputs["repository"] = state ? state.repository : undefined;
            resourceInputs["username"] = state ? state.username : undefined;
        } else {
            const args = argsOrState as RepositoryCollaboratorArgs | undefined;
            if ((!args || args.repository === undefined) && !opts.urn) {
                throw new Error("Missing required property 'repository'");
            }
            if ((!args || args.username === undefined) && !opts.urn) {
                throw new Error("Missing required property 'username'");
            }
            resourceInputs["permission"] = args ? args.permission : undefined;
            resourceInputs["permissionDiffSuppression"] = args ? args.permissionDiffSuppression : undefined;
            resourceInputs["repository"] = args ? args.repository : undefined;
            resourceInputs["username"] = args ? args.username : undefined;
            resourceInputs["invitationId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RepositoryCollaborator.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RepositoryCollaborator resources.
 */
export interface RepositoryCollaboratorState {
    /**
     * ID of the invitation to be used in `github.UserInvitationAccepter`
     */
    invitationId?: pulumi.Input<string>;
    /**
     * The permission of the outside collaborator for the repository.
     * Must be one of `pull`, `push`, `maintain`, `triage` or `admin` or the name of an existing [custom repository role](https://docs.github.com/en/enterprise-cloud@latest/organizations/managing-peoples-access-to-your-organization-with-roles/managing-custom-repository-roles-for-an-organization) within the organization for organization-owned repositories.
     * Must be `push` for personal repositories. Defaults to `push`.
     */
    permission?: pulumi.Input<string>;
    /**
     * Suppress plan diffs for `triage` and `maintain`.  Defaults to `false`.
     */
    permissionDiffSuppression?: pulumi.Input<boolean>;
    /**
     * The GitHub repository
     */
    repository?: pulumi.Input<string>;
    /**
     * The user to add to the repository as a collaborator.
     */
    username?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RepositoryCollaborator resource.
 */
export interface RepositoryCollaboratorArgs {
    /**
     * The permission of the outside collaborator for the repository.
     * Must be one of `pull`, `push`, `maintain`, `triage` or `admin` or the name of an existing [custom repository role](https://docs.github.com/en/enterprise-cloud@latest/organizations/managing-peoples-access-to-your-organization-with-roles/managing-custom-repository-roles-for-an-organization) within the organization for organization-owned repositories.
     * Must be `push` for personal repositories. Defaults to `push`.
     */
    permission?: pulumi.Input<string>;
    /**
     * Suppress plan diffs for `triage` and `maintain`.  Defaults to `false`.
     */
    permissionDiffSuppression?: pulumi.Input<boolean>;
    /**
     * The GitHub repository
     */
    repository: pulumi.Input<string>;
    /**
     * The user to add to the repository as a collaborator.
     */
    username: pulumi.Input<string>;
}
