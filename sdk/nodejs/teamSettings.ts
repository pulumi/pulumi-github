// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * This resource manages the team settings (in particular the request review delegation settings) within the organization
 *
 * Creating this resource will alter the team Code Review settings.
 *
 * The team must both belong to the same organization configured in the provider on GitHub.
 *
 * > **Note**: This resource relies on the v4 GraphQl GitHub API. If this API is not available, or the Stone Crop schema preview is not available, then this resource will not work as intended.
 *
 * ## Import
 *
 * GitHub Teams can be imported using the GitHub team ID, or the team slug e.g.
 *
 * ```sh
 *  $ pulumi import github:index/teamSettings:TeamSettings code_review_settings 1234567
 * ```
 *  or,
 *
 * ```sh
 *  $ pulumi import github:index/teamSettings:TeamSettings code_review_settings SomeTeam
 * ```
 */
export class TeamSettings extends pulumi.CustomResource {
    /**
     * Get an existing TeamSettings resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TeamSettingsState, opts?: pulumi.CustomResourceOptions): TeamSettings {
        return new TeamSettings(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'github:index/teamSettings:TeamSettings';

    /**
     * Returns true if the given object is an instance of TeamSettings.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TeamSettings {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TeamSettings.__pulumiType;
    }

    /**
     * The settings for delegating code reviews to individuals on behalf of the team. If this block is present, even without any fields, then review request delegation will be enabled for the team. See GitHub Review Request Delegation below for details. See [GitHub's documentation](https://docs.github.com/en/organizations/organizing-members-into-teams/managing-code-review-settings-for-your-team#configuring-team-notifications) for more configuration details.
     */
    public readonly reviewRequestDelegation!: pulumi.Output<outputs.TeamSettingsReviewRequestDelegation | undefined>;
    /**
     * The GitHub team id or the GitHub team slug
     */
    public readonly teamId!: pulumi.Output<string>;
    /**
     * The slug of the Team within the Organization.
     */
    public /*out*/ readonly teamSlug!: pulumi.Output<string>;
    /**
     * The unique ID of the Team on GitHub. Corresponds to the ID of the 'github_team_settings' resource.
     */
    public /*out*/ readonly teamUid!: pulumi.Output<string>;

    /**
     * Create a TeamSettings resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TeamSettingsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TeamSettingsArgs | TeamSettingsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TeamSettingsState | undefined;
            resourceInputs["reviewRequestDelegation"] = state ? state.reviewRequestDelegation : undefined;
            resourceInputs["teamId"] = state ? state.teamId : undefined;
            resourceInputs["teamSlug"] = state ? state.teamSlug : undefined;
            resourceInputs["teamUid"] = state ? state.teamUid : undefined;
        } else {
            const args = argsOrState as TeamSettingsArgs | undefined;
            if ((!args || args.teamId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'teamId'");
            }
            resourceInputs["reviewRequestDelegation"] = args ? args.reviewRequestDelegation : undefined;
            resourceInputs["teamId"] = args ? args.teamId : undefined;
            resourceInputs["teamSlug"] = undefined /*out*/;
            resourceInputs["teamUid"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(TeamSettings.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TeamSettings resources.
 */
export interface TeamSettingsState {
    /**
     * The settings for delegating code reviews to individuals on behalf of the team. If this block is present, even without any fields, then review request delegation will be enabled for the team. See GitHub Review Request Delegation below for details. See [GitHub's documentation](https://docs.github.com/en/organizations/organizing-members-into-teams/managing-code-review-settings-for-your-team#configuring-team-notifications) for more configuration details.
     */
    reviewRequestDelegation?: pulumi.Input<inputs.TeamSettingsReviewRequestDelegation>;
    /**
     * The GitHub team id or the GitHub team slug
     */
    teamId?: pulumi.Input<string>;
    /**
     * The slug of the Team within the Organization.
     */
    teamSlug?: pulumi.Input<string>;
    /**
     * The unique ID of the Team on GitHub. Corresponds to the ID of the 'github_team_settings' resource.
     */
    teamUid?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a TeamSettings resource.
 */
export interface TeamSettingsArgs {
    /**
     * The settings for delegating code reviews to individuals on behalf of the team. If this block is present, even without any fields, then review request delegation will be enabled for the team. See GitHub Review Request Delegation below for details. See [GitHub's documentation](https://docs.github.com/en/organizations/organizing-members-into-teams/managing-code-review-settings-for-your-team#configuring-team-notifications) for more configuration details.
     */
    reviewRequestDelegation?: pulumi.Input<inputs.TeamSettingsReviewRequestDelegation>;
    /**
     * The GitHub team id or the GitHub team slug
     */
    teamId: pulumi.Input<string>;
}
