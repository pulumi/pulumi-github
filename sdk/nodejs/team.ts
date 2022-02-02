// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Provides a GitHub team resource.
 *
 * This resource allows you to add/remove teams from your organization. When applied,
 * a new team will be created. When destroyed, that team will be removed.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as github from "@pulumi/github";
 *
 * // Add a team to the organization
 * const someTeam = new github.Team("some_team", {
 *     description: "Some cool team",
 *     privacy: "closed",
 * });
 * ```
 *
 * ## Import
 *
 * GitHub Teams can be imported using the GitHub team ID e.g.
 *
 * ```sh
 *  $ pulumi import github:index/team:Team core 1234567
 * ```
 */
export class Team extends pulumi.CustomResource {
    /**
     * Get an existing Team resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TeamState, opts?: pulumi.CustomResourceOptions): Team {
        return new Team(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'github:index/team:Team';

    /**
     * Returns true if the given object is an instance of Team.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Team {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Team.__pulumiType;
    }

    /**
     * Adds a default maintainer to the team. Defaults to `false` and adds the creating user to the team when `true`.
     */
    public readonly createDefaultMaintainer!: pulumi.Output<boolean | undefined>;
    /**
     * A description of the team.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * The LDAP Distinguished Name of the group where membership will be synchronized. Only available in GitHub Enterprise Server.
     */
    public readonly ldapDn!: pulumi.Output<string | undefined>;
    public /*out*/ readonly membersCount!: pulumi.Output<number>;
    /**
     * The name of the team.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The Node ID of the created team.
     */
    public /*out*/ readonly nodeId!: pulumi.Output<string>;
    /**
     * The ID of the parent team, if this is a nested team.
     */
    public readonly parentTeamId!: pulumi.Output<number | undefined>;
    /**
     * The level of privacy for the team. Must be one of `secret` or `closed`.
     * Defaults to `secret`.
     */
    public readonly privacy!: pulumi.Output<string | undefined>;
    /**
     * The slug of the created team, which may or may not differ from `name`,
     * depending on whether `name` contains "URL-unsafe" characters.
     * Useful when referencing the team in [`github.BranchProtection`](https://www.terraform.io/docs/providers/github/r/branch_protection.html).
     */
    public /*out*/ readonly slug!: pulumi.Output<string>;

    /**
     * Create a Team resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: TeamArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TeamArgs | TeamState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TeamState | undefined;
            resourceInputs["createDefaultMaintainer"] = state ? state.createDefaultMaintainer : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["etag"] = state ? state.etag : undefined;
            resourceInputs["ldapDn"] = state ? state.ldapDn : undefined;
            resourceInputs["membersCount"] = state ? state.membersCount : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["nodeId"] = state ? state.nodeId : undefined;
            resourceInputs["parentTeamId"] = state ? state.parentTeamId : undefined;
            resourceInputs["privacy"] = state ? state.privacy : undefined;
            resourceInputs["slug"] = state ? state.slug : undefined;
        } else {
            const args = argsOrState as TeamArgs | undefined;
            resourceInputs["createDefaultMaintainer"] = args ? args.createDefaultMaintainer : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["ldapDn"] = args ? args.ldapDn : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["parentTeamId"] = args ? args.parentTeamId : undefined;
            resourceInputs["privacy"] = args ? args.privacy : undefined;
            resourceInputs["etag"] = undefined /*out*/;
            resourceInputs["membersCount"] = undefined /*out*/;
            resourceInputs["nodeId"] = undefined /*out*/;
            resourceInputs["slug"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Team.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Team resources.
 */
export interface TeamState {
    /**
     * Adds a default maintainer to the team. Defaults to `false` and adds the creating user to the team when `true`.
     */
    createDefaultMaintainer?: pulumi.Input<boolean>;
    /**
     * A description of the team.
     */
    description?: pulumi.Input<string>;
    etag?: pulumi.Input<string>;
    /**
     * The LDAP Distinguished Name of the group where membership will be synchronized. Only available in GitHub Enterprise Server.
     */
    ldapDn?: pulumi.Input<string>;
    membersCount?: pulumi.Input<number>;
    /**
     * The name of the team.
     */
    name?: pulumi.Input<string>;
    /**
     * The Node ID of the created team.
     */
    nodeId?: pulumi.Input<string>;
    /**
     * The ID of the parent team, if this is a nested team.
     */
    parentTeamId?: pulumi.Input<number>;
    /**
     * The level of privacy for the team. Must be one of `secret` or `closed`.
     * Defaults to `secret`.
     */
    privacy?: pulumi.Input<string>;
    /**
     * The slug of the created team, which may or may not differ from `name`,
     * depending on whether `name` contains "URL-unsafe" characters.
     * Useful when referencing the team in [`github.BranchProtection`](https://www.terraform.io/docs/providers/github/r/branch_protection.html).
     */
    slug?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Team resource.
 */
export interface TeamArgs {
    /**
     * Adds a default maintainer to the team. Defaults to `false` and adds the creating user to the team when `true`.
     */
    createDefaultMaintainer?: pulumi.Input<boolean>;
    /**
     * A description of the team.
     */
    description?: pulumi.Input<string>;
    /**
     * The LDAP Distinguished Name of the group where membership will be synchronized. Only available in GitHub Enterprise Server.
     */
    ldapDn?: pulumi.Input<string>;
    /**
     * The name of the team.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the parent team, if this is a nested team.
     */
    parentTeamId?: pulumi.Input<number>;
    /**
     * The level of privacy for the team. Must be one of `secret` or `closed`.
     * Defaults to `secret`.
     */
    privacy?: pulumi.Input<string>;
}
