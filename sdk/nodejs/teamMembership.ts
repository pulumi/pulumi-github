// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Provides a GitHub team membership resource.
 *
 * This resource allows you to add/remove users from teams in your organization. When applied,
 * the user will be added to the team. If the user hasn't accepted their invitation to the
 * organization, they won't be part of the team until they do. When
 * destroyed, the user will be removed from the team.
 *
 * > **Note**: This resource is not compatible with `github.TeamMembers`. Use either `github.TeamMembers` or `github.TeamMembership`.
 *
 * ## Import
 *
 * GitHub Team Membership can be imported using an ID made up of `teamid:username`, e.g.
 *
 * ```sh
 *  $ pulumi import github:index/teamMembership:TeamMembership member 1234567:someuser
 * ```
 */
export class TeamMembership extends pulumi.CustomResource {
    /**
     * Get an existing TeamMembership resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TeamMembershipState, opts?: pulumi.CustomResourceOptions): TeamMembership {
        return new TeamMembership(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'github:index/teamMembership:TeamMembership';

    /**
     * Returns true if the given object is an instance of TeamMembership.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TeamMembership {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TeamMembership.__pulumiType;
    }

    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * The role of the user within the team.
     * Must be one of `member` or `maintainer`. Defaults to `member`.
     */
    public readonly role!: pulumi.Output<string | undefined>;
    /**
     * The GitHub team id
     */
    public readonly teamId!: pulumi.Output<string>;
    /**
     * The user to add to the team.
     */
    public readonly username!: pulumi.Output<string>;

    /**
     * Create a TeamMembership resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TeamMembershipArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TeamMembershipArgs | TeamMembershipState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TeamMembershipState | undefined;
            resourceInputs["etag"] = state ? state.etag : undefined;
            resourceInputs["role"] = state ? state.role : undefined;
            resourceInputs["teamId"] = state ? state.teamId : undefined;
            resourceInputs["username"] = state ? state.username : undefined;
        } else {
            const args = argsOrState as TeamMembershipArgs | undefined;
            if ((!args || args.teamId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'teamId'");
            }
            if ((!args || args.username === undefined) && !opts.urn) {
                throw new Error("Missing required property 'username'");
            }
            resourceInputs["role"] = args ? args.role : undefined;
            resourceInputs["teamId"] = args ? args.teamId : undefined;
            resourceInputs["username"] = args ? args.username : undefined;
            resourceInputs["etag"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(TeamMembership.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TeamMembership resources.
 */
export interface TeamMembershipState {
    etag?: pulumi.Input<string>;
    /**
     * The role of the user within the team.
     * Must be one of `member` or `maintainer`. Defaults to `member`.
     */
    role?: pulumi.Input<string>;
    /**
     * The GitHub team id
     */
    teamId?: pulumi.Input<string>;
    /**
     * The user to add to the team.
     */
    username?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a TeamMembership resource.
 */
export interface TeamMembershipArgs {
    /**
     * The role of the user within the team.
     * Must be one of `member` or `maintainer`. Defaults to `member`.
     */
    role?: pulumi.Input<string>;
    /**
     * The GitHub team id
     */
    teamId: pulumi.Input<string>;
    /**
     * The user to add to the team.
     */
    username: pulumi.Input<string>;
}
