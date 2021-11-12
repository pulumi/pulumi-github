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
     * ID of the invitation to accept
     */
    public readonly invitationId!: pulumi.Output<string>;

    /**
     * Create a UserInvitationAccepter resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UserInvitationAccepterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UserInvitationAccepterArgs | UserInvitationAccepterState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as UserInvitationAccepterState | undefined;
            resourceInputs["invitationId"] = state ? state.invitationId : undefined;
        } else {
            const args = argsOrState as UserInvitationAccepterArgs | undefined;
            if ((!args || args.invitationId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'invitationId'");
            }
            resourceInputs["invitationId"] = args ? args.invitationId : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(UserInvitationAccepter.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering UserInvitationAccepter resources.
 */
export interface UserInvitationAccepterState {
    /**
     * ID of the invitation to accept
     */
    invitationId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a UserInvitationAccepter resource.
 */
export interface UserInvitationAccepterArgs {
    /**
     * ID of the invitation to accept
     */
    invitationId: pulumi.Input<string>;
}
