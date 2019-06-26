// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

import {Permission} from "../index";

/**
 * Provides a GitHub repository collaborator resource.
 * 
 * This resource allows you to add/remove collaborators from repositories in your
 * organization. Collaborators can have explicit (and differing levels of) read,
 * write, or administrator access to specific repositories in your organization,
 * without giving the user full organization membership.
 * 
 * When applied, an invitation will be sent to the user to become a collaborator
 * on a repository. When destroyed, either the invitation will be cancelled or the
 * collaborator will be removed from the repository.
 * 
 * Further documentation on GitHub collaborators:
 * 
 * - [Adding outside collaborators to repositories in your organization](https://help.github.com/articles/adding-outside-collaborators-to-repositories-in-your-organization/)
 * - [Converting an organization member to an outside collaborator](https://help.github.com/articles/converting-an-organization-member-to-an-outside-collaborator/)
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as github from "@pulumi/github";
 * 
 * // Add a collaborator to a repository
 * const aRepoCollaborator = new github.repos.Collaborator("a_repo_collaborator", {
 *     permission: "admin",
 *     repository: "our-cool-repo",
 *     username: "SomeUser",
 * });
 * ```
 */
export class Collaborator extends pulumi.CustomResource {
    /**
     * Get an existing Collaborator resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CollaboratorState, opts?: pulumi.CustomResourceOptions): Collaborator {
        return new Collaborator(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'github:repos/collaborator:Collaborator';

    /**
     * Returns true if the given object is an instance of Collaborator.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Collaborator {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Collaborator.__pulumiType;
    }

    /**
     * ID of the invitation to be used in `github_user_invitation_accepter`
     */
    public /*out*/ readonly invitationId!: pulumi.Output<string>;
    /**
     * The permission of the outside collaborator for the repository.
     * Must be one of `pull`, `push`, or `admin`. Defaults to `push`.
     */
    public readonly permission!: pulumi.Output<Permission | undefined>;
    /**
     * The GitHub repository
     */
    public readonly repository!: pulumi.Output<string>;
    /**
     * The user to add to the repository as a collaborator.
     */
    public readonly username!: pulumi.Output<string>;

    /**
     * Create a Collaborator resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CollaboratorArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CollaboratorArgs | CollaboratorState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as CollaboratorState | undefined;
            inputs["invitationId"] = state ? state.invitationId : undefined;
            inputs["permission"] = state ? state.permission : undefined;
            inputs["repository"] = state ? state.repository : undefined;
            inputs["username"] = state ? state.username : undefined;
        } else {
            const args = argsOrState as CollaboratorArgs | undefined;
            if (!args || args.repository === undefined) {
                throw new Error("Missing required property 'repository'");
            }
            if (!args || args.username === undefined) {
                throw new Error("Missing required property 'username'");
            }
            inputs["permission"] = args ? args.permission : undefined;
            inputs["repository"] = args ? args.repository : undefined;
            inputs["username"] = args ? args.username : undefined;
            inputs["invitationId"] = undefined /*out*/;
        }
        super(Collaborator.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Collaborator resources.
 */
export interface CollaboratorState {
    /**
     * ID of the invitation to be used in `github_user_invitation_accepter`
     */
    readonly invitationId?: pulumi.Input<string>;
    /**
     * The permission of the outside collaborator for the repository.
     * Must be one of `pull`, `push`, or `admin`. Defaults to `push`.
     */
    readonly permission?: pulumi.Input<Permission>;
    /**
     * The GitHub repository
     */
    readonly repository?: pulumi.Input<string>;
    /**
     * The user to add to the repository as a collaborator.
     */
    readonly username?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Collaborator resource.
 */
export interface CollaboratorArgs {
    /**
     * The permission of the outside collaborator for the repository.
     * Must be one of `pull`, `push`, or `admin`. Defaults to `push`.
     */
    readonly permission?: pulumi.Input<Permission>;
    /**
     * The GitHub repository
     */
    readonly repository: pulumi.Input<string>;
    /**
     * The user to add to the repository as a collaborator.
     */
    readonly username: pulumi.Input<string>;
}
