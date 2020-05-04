// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * This resource manages relationships between teams and repositories
 * in your GitHub organization.
 * 
 * Creating this resource grants a particular team permissions on a
 * particular repository.
 * 
 * The repository and the team must both belong to the same organization
 * on GitHub. This resource does not actually *create* any repositories;
 * to do that, see `github..Repository`.
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as github from "@pulumi/github";
 * 
 * // Add a repository to the team
 * const someTeam = new github.Team("someTeam", {
 *     description: "Some cool team",
 * });
 * const someRepo = new github.Repository("someRepo", {});
 * const someTeamRepo = new github.TeamRepository("someTeamRepo", {
 *     permission: "pull",
 *     repository: someRepo.name,
 *     teamId: someTeam.id,
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-github/blob/master/website/docs/r/team_repository.html.markdown.
 */
export class TeamRepository extends pulumi.CustomResource {
    /**
     * Get an existing TeamRepository resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TeamRepositoryState, opts?: pulumi.CustomResourceOptions): TeamRepository {
        return new TeamRepository(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'github:index/teamRepository:TeamRepository';

    /**
     * Returns true if the given object is an instance of TeamRepository.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TeamRepository {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TeamRepository.__pulumiType;
    }

    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * The permissions of team members regarding the repository.
     * Must be one of `pull`, `triage`, `push`, `maintain`, or `admin`. Defaults to `pull`.
     */
    public readonly permission!: pulumi.Output<string | undefined>;
    /**
     * The repository to add to the team.
     */
    public readonly repository!: pulumi.Output<string>;
    /**
     * The GitHub team id
     */
    public readonly teamId!: pulumi.Output<string>;

    /**
     * Create a TeamRepository resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TeamRepositoryArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TeamRepositoryArgs | TeamRepositoryState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as TeamRepositoryState | undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["permission"] = state ? state.permission : undefined;
            inputs["repository"] = state ? state.repository : undefined;
            inputs["teamId"] = state ? state.teamId : undefined;
        } else {
            const args = argsOrState as TeamRepositoryArgs | undefined;
            if (!args || args.repository === undefined) {
                throw new Error("Missing required property 'repository'");
            }
            if (!args || args.teamId === undefined) {
                throw new Error("Missing required property 'teamId'");
            }
            inputs["permission"] = args ? args.permission : undefined;
            inputs["repository"] = args ? args.repository : undefined;
            inputs["teamId"] = args ? args.teamId : undefined;
            inputs["etag"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(TeamRepository.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TeamRepository resources.
 */
export interface TeamRepositoryState {
    readonly etag?: pulumi.Input<string>;
    /**
     * The permissions of team members regarding the repository.
     * Must be one of `pull`, `triage`, `push`, `maintain`, or `admin`. Defaults to `pull`.
     */
    readonly permission?: pulumi.Input<string>;
    /**
     * The repository to add to the team.
     */
    readonly repository?: pulumi.Input<string>;
    /**
     * The GitHub team id
     */
    readonly teamId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a TeamRepository resource.
 */
export interface TeamRepositoryArgs {
    /**
     * The permissions of team members regarding the repository.
     * Must be one of `pull`, `triage`, `push`, `maintain`, or `admin`. Defaults to `pull`.
     */
    readonly permission?: pulumi.Input<string>;
    /**
     * The repository to add to the team.
     */
    readonly repository: pulumi.Input<string>;
    /**
     * The GitHub team id
     */
    readonly teamId: pulumi.Input<string>;
}