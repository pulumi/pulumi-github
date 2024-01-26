// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Provides a resource to manage GitHub repository collaborator invitations.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as github from "@pulumi/github";
 *
 * const exampleRepository = new github.Repository("exampleRepository", {});
 * const exampleRepositoryCollaborator = new github.RepositoryCollaborator("exampleRepositoryCollaborator", {
 *     repository: exampleRepository.name,
 *     username: "example-username",
 *     permission: "push",
 * });
 * const invitee = new github.Provider("invitee", {token: _var.invitee_token});
 * const exampleUserInvitationAccepter = new github.UserInvitationAccepter("exampleUserInvitationAccepter", {invitationId: exampleRepositoryCollaborator.invitationId}, {
 *     provider: "github.invitee",
 * });
 * ```
 * ## Allowing empty invitation IDs
 *
 * Set `allowEmptyId` when using `forEach` over a list of `github_repository_collaborator.invitation_id`'s.
 *
 * This allows applying a module again when a new `github.RepositoryCollaborator` resource is added to the `forEach` loop.
 * This is needed as the `github_repository_collaborator.invitation_id` will be empty after a state refresh when the invitation has been accepted.
 *
 * Note that when an invitation is accepted manually or by another tool between a state refresh and a `pulumi up` using that refreshed state,
 * the plan will contain the invitation ID, but the apply will receive an HTTP 404 from the API since the invitation has already been accepted.
 *
 * This is tracked in #1157.
 */
export class UserInvitationAccepter extends pulumi.CustomResource {
    /**
     * Get an existing UserInvitationAccepter resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UserInvitationAccepterState, opts?: pulumi.CustomResourceOptions): UserInvitationAccepter {
        return new UserInvitationAccepter(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'github:index/userInvitationAccepter:UserInvitationAccepter';

    /**
     * Returns true if the given object is an instance of UserInvitationAccepter.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is UserInvitationAccepter {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === UserInvitationAccepter.__pulumiType;
    }

    /**
     * Allow the ID to be unset. This will result in the resource being skipped when the ID is not set instead of returning an error.
     */
    public readonly allowEmptyId!: pulumi.Output<boolean | undefined>;
    /**
     * ID of the invitation to accept. Must be set when `allowEmptyId` is `false`.
     */
    public readonly invitationId!: pulumi.Output<string | undefined>;

    /**
     * Create a UserInvitationAccepter resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: UserInvitationAccepterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UserInvitationAccepterArgs | UserInvitationAccepterState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as UserInvitationAccepterState | undefined;
            resourceInputs["allowEmptyId"] = state ? state.allowEmptyId : undefined;
            resourceInputs["invitationId"] = state ? state.invitationId : undefined;
        } else {
            const args = argsOrState as UserInvitationAccepterArgs | undefined;
            resourceInputs["allowEmptyId"] = args ? args.allowEmptyId : undefined;
            resourceInputs["invitationId"] = args ? args.invitationId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(UserInvitationAccepter.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering UserInvitationAccepter resources.
 */
export interface UserInvitationAccepterState {
    /**
     * Allow the ID to be unset. This will result in the resource being skipped when the ID is not set instead of returning an error.
     */
    allowEmptyId?: pulumi.Input<boolean>;
    /**
     * ID of the invitation to accept. Must be set when `allowEmptyId` is `false`.
     */
    invitationId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a UserInvitationAccepter resource.
 */
export interface UserInvitationAccepterArgs {
    /**
     * Allow the ID to be unset. This will result in the resource being skipped when the ID is not set instead of returning an error.
     */
    allowEmptyId?: pulumi.Input<boolean>;
    /**
     * ID of the invitation to accept. Must be set when `allowEmptyId` is `false`.
     */
    invitationId?: pulumi.Input<string>;
}
