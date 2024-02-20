// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as github from "@pulumi/github";
 *
 * const exampleRepository = new github.Repository("exampleRepository", {});
 * const exampleUser = github.getUser({
 *     username: "example",
 * });
 * const exampleTeam = new github.Team("exampleTeam", {});
 * // Protect the main branch of the foo repository. Additionally, require that
 * // the "ci/travis" context to be passing and only allow the engineers team merge
 * // to the branch.
 * const exampleBranchProtection = new github.BranchProtection("exampleBranchProtection", {
 *     repositoryId: exampleRepository.nodeId,
 *     pattern: "main",
 *     enforceAdmins: true,
 *     allowsDeletions: true,
 *     requiredStatusChecks: [{
 *         strict: false,
 *         contexts: ["ci/travis"],
 *     }],
 *     requiredPullRequestReviews: [{
 *         dismissStaleReviews: true,
 *         restrictDismissals: true,
 *         dismissalRestrictions: [
 *             exampleUser.then(exampleUser => exampleUser.nodeId),
 *             exampleTeam.nodeId,
 *             "/exampleuser",
 *             "exampleorganization/exampleteam",
 *         ],
 *     }],
 *     restrictPushes: [{
 *         pushAllowances: [
 *             exampleUser.then(exampleUser => exampleUser.nodeId),
 *             "/exampleuser",
 *             "exampleorganization/exampleteam",
 *         ],
 *     }],
 *     forcePushBypassers: [
 *         exampleUser.then(exampleUser => exampleUser.nodeId),
 *         "/exampleuser",
 *         "exampleorganization/exampleteam",
 *     ],
 * });
 * const exampleTeamRepository = new github.TeamRepository("exampleTeamRepository", {
 *     teamId: exampleTeam.id,
 *     repository: exampleRepository.name,
 *     permission: "pull",
 * });
 * ```
 *
 * ## Import
 *
 * GitHub Branch Protection can be imported using an ID made up of `repository:pattern`, e.g.
 *
 * ```sh
 *  $ pulumi import github:index/branchProtection:BranchProtection terraform terraform:main
 * ```
 */
export class BranchProtection extends pulumi.CustomResource {
    /**
     * Get an existing BranchProtection resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BranchProtectionState, opts?: pulumi.CustomResourceOptions): BranchProtection {
        return new BranchProtection(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'github:index/branchProtection:BranchProtection';

    /**
     * Returns true if the given object is an instance of BranchProtection.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BranchProtection {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BranchProtection.__pulumiType;
    }

    /**
     * Boolean, setting this to `true` to allow the branch to be deleted.
     */
    public readonly allowsDeletions!: pulumi.Output<boolean | undefined>;
    /**
     * Boolean, setting this to `true` to allow force pushes on the branch to everyone. Set it to `false` if you specify `forcePushBypassers`.
     */
    public readonly allowsForcePushes!: pulumi.Output<boolean | undefined>;
    /**
     * Boolean, setting this to `true` enforces status checks for repository administrators.
     */
    public readonly enforceAdmins!: pulumi.Output<boolean | undefined>;
    /**
     * The list of actor Names/IDs that are allowed to bypass force push restrictions. Actor names must either begin with a "/" for users or the organization name followed by a "/" for teams. If the list is not empty, `allowsForcePushes` should be set to `false`.
     */
    public readonly forcePushBypassers!: pulumi.Output<string[] | undefined>;
    /**
     * Boolean, Setting this to `true` will make the branch read-only and preventing any pushes to it. Defaults to `false`
     */
    public readonly lockBranch!: pulumi.Output<boolean | undefined>;
    /**
     * Identifies the protection rule pattern.
     */
    public readonly pattern!: pulumi.Output<string>;
    /**
     * The name or node ID of the repository associated with this branch protection rule.
     */
    public readonly repositoryId!: pulumi.Output<string>;
    /**
     * Boolean, setting this to `true` requires all conversations on code must be resolved before a pull request can be merged.
     */
    public readonly requireConversationResolution!: pulumi.Output<boolean | undefined>;
    /**
     * Boolean, setting this to `true` requires all commits to be signed with GPG.
     */
    public readonly requireSignedCommits!: pulumi.Output<boolean | undefined>;
    /**
     * Boolean, setting this to `true` enforces a linear commit Git history, which prevents anyone from pushing merge commits to a branch
     */
    public readonly requiredLinearHistory!: pulumi.Output<boolean | undefined>;
    /**
     * Enforce restrictions for pull request reviews. See Required Pull Request Reviews below for details.
     */
    public readonly requiredPullRequestReviews!: pulumi.Output<outputs.BranchProtectionRequiredPullRequestReview[] | undefined>;
    /**
     * Enforce restrictions for required status checks. See Required Status Checks below for details.
     */
    public readonly requiredStatusChecks!: pulumi.Output<outputs.BranchProtectionRequiredStatusCheck[] | undefined>;
    /**
     * Restrict pushes to matching branches. See Restrict Pushes below for details.
     */
    public readonly restrictPushes!: pulumi.Output<outputs.BranchProtectionRestrictPush[] | undefined>;

    /**
     * Create a BranchProtection resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BranchProtectionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BranchProtectionArgs | BranchProtectionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BranchProtectionState | undefined;
            resourceInputs["allowsDeletions"] = state ? state.allowsDeletions : undefined;
            resourceInputs["allowsForcePushes"] = state ? state.allowsForcePushes : undefined;
            resourceInputs["enforceAdmins"] = state ? state.enforceAdmins : undefined;
            resourceInputs["forcePushBypassers"] = state ? state.forcePushBypassers : undefined;
            resourceInputs["lockBranch"] = state ? state.lockBranch : undefined;
            resourceInputs["pattern"] = state ? state.pattern : undefined;
            resourceInputs["repositoryId"] = state ? state.repositoryId : undefined;
            resourceInputs["requireConversationResolution"] = state ? state.requireConversationResolution : undefined;
            resourceInputs["requireSignedCommits"] = state ? state.requireSignedCommits : undefined;
            resourceInputs["requiredLinearHistory"] = state ? state.requiredLinearHistory : undefined;
            resourceInputs["requiredPullRequestReviews"] = state ? state.requiredPullRequestReviews : undefined;
            resourceInputs["requiredStatusChecks"] = state ? state.requiredStatusChecks : undefined;
            resourceInputs["restrictPushes"] = state ? state.restrictPushes : undefined;
        } else {
            const args = argsOrState as BranchProtectionArgs | undefined;
            if ((!args || args.pattern === undefined) && !opts.urn) {
                throw new Error("Missing required property 'pattern'");
            }
            if ((!args || args.repositoryId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'repositoryId'");
            }
            resourceInputs["allowsDeletions"] = args ? args.allowsDeletions : undefined;
            resourceInputs["allowsForcePushes"] = args ? args.allowsForcePushes : undefined;
            resourceInputs["enforceAdmins"] = args ? args.enforceAdmins : undefined;
            resourceInputs["forcePushBypassers"] = args ? args.forcePushBypassers : undefined;
            resourceInputs["lockBranch"] = args ? args.lockBranch : undefined;
            resourceInputs["pattern"] = args ? args.pattern : undefined;
            resourceInputs["repositoryId"] = args ? args.repositoryId : undefined;
            resourceInputs["requireConversationResolution"] = args ? args.requireConversationResolution : undefined;
            resourceInputs["requireSignedCommits"] = args ? args.requireSignedCommits : undefined;
            resourceInputs["requiredLinearHistory"] = args ? args.requiredLinearHistory : undefined;
            resourceInputs["requiredPullRequestReviews"] = args ? args.requiredPullRequestReviews : undefined;
            resourceInputs["requiredStatusChecks"] = args ? args.requiredStatusChecks : undefined;
            resourceInputs["restrictPushes"] = args ? args.restrictPushes : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(BranchProtection.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BranchProtection resources.
 */
export interface BranchProtectionState {
    /**
     * Boolean, setting this to `true` to allow the branch to be deleted.
     */
    allowsDeletions?: pulumi.Input<boolean>;
    /**
     * Boolean, setting this to `true` to allow force pushes on the branch to everyone. Set it to `false` if you specify `forcePushBypassers`.
     */
    allowsForcePushes?: pulumi.Input<boolean>;
    /**
     * Boolean, setting this to `true` enforces status checks for repository administrators.
     */
    enforceAdmins?: pulumi.Input<boolean>;
    /**
     * The list of actor Names/IDs that are allowed to bypass force push restrictions. Actor names must either begin with a "/" for users or the organization name followed by a "/" for teams. If the list is not empty, `allowsForcePushes` should be set to `false`.
     */
    forcePushBypassers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Boolean, Setting this to `true` will make the branch read-only and preventing any pushes to it. Defaults to `false`
     */
    lockBranch?: pulumi.Input<boolean>;
    /**
     * Identifies the protection rule pattern.
     */
    pattern?: pulumi.Input<string>;
    /**
     * The name or node ID of the repository associated with this branch protection rule.
     */
    repositoryId?: pulumi.Input<string>;
    /**
     * Boolean, setting this to `true` requires all conversations on code must be resolved before a pull request can be merged.
     */
    requireConversationResolution?: pulumi.Input<boolean>;
    /**
     * Boolean, setting this to `true` requires all commits to be signed with GPG.
     */
    requireSignedCommits?: pulumi.Input<boolean>;
    /**
     * Boolean, setting this to `true` enforces a linear commit Git history, which prevents anyone from pushing merge commits to a branch
     */
    requiredLinearHistory?: pulumi.Input<boolean>;
    /**
     * Enforce restrictions for pull request reviews. See Required Pull Request Reviews below for details.
     */
    requiredPullRequestReviews?: pulumi.Input<pulumi.Input<inputs.BranchProtectionRequiredPullRequestReview>[]>;
    /**
     * Enforce restrictions for required status checks. See Required Status Checks below for details.
     */
    requiredStatusChecks?: pulumi.Input<pulumi.Input<inputs.BranchProtectionRequiredStatusCheck>[]>;
    /**
     * Restrict pushes to matching branches. See Restrict Pushes below for details.
     */
    restrictPushes?: pulumi.Input<pulumi.Input<inputs.BranchProtectionRestrictPush>[]>;
}

/**
 * The set of arguments for constructing a BranchProtection resource.
 */
export interface BranchProtectionArgs {
    /**
     * Boolean, setting this to `true` to allow the branch to be deleted.
     */
    allowsDeletions?: pulumi.Input<boolean>;
    /**
     * Boolean, setting this to `true` to allow force pushes on the branch to everyone. Set it to `false` if you specify `forcePushBypassers`.
     */
    allowsForcePushes?: pulumi.Input<boolean>;
    /**
     * Boolean, setting this to `true` enforces status checks for repository administrators.
     */
    enforceAdmins?: pulumi.Input<boolean>;
    /**
     * The list of actor Names/IDs that are allowed to bypass force push restrictions. Actor names must either begin with a "/" for users or the organization name followed by a "/" for teams. If the list is not empty, `allowsForcePushes` should be set to `false`.
     */
    forcePushBypassers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Boolean, Setting this to `true` will make the branch read-only and preventing any pushes to it. Defaults to `false`
     */
    lockBranch?: pulumi.Input<boolean>;
    /**
     * Identifies the protection rule pattern.
     */
    pattern: pulumi.Input<string>;
    /**
     * The name or node ID of the repository associated with this branch protection rule.
     */
    repositoryId: pulumi.Input<string>;
    /**
     * Boolean, setting this to `true` requires all conversations on code must be resolved before a pull request can be merged.
     */
    requireConversationResolution?: pulumi.Input<boolean>;
    /**
     * Boolean, setting this to `true` requires all commits to be signed with GPG.
     */
    requireSignedCommits?: pulumi.Input<boolean>;
    /**
     * Boolean, setting this to `true` enforces a linear commit Git history, which prevents anyone from pushing merge commits to a branch
     */
    requiredLinearHistory?: pulumi.Input<boolean>;
    /**
     * Enforce restrictions for pull request reviews. See Required Pull Request Reviews below for details.
     */
    requiredPullRequestReviews?: pulumi.Input<pulumi.Input<inputs.BranchProtectionRequiredPullRequestReview>[]>;
    /**
     * Enforce restrictions for required status checks. See Required Status Checks below for details.
     */
    requiredStatusChecks?: pulumi.Input<pulumi.Input<inputs.BranchProtectionRequiredStatusCheck>[]>;
    /**
     * Restrict pushes to matching branches. See Restrict Pushes below for details.
     */
    restrictPushes?: pulumi.Input<pulumi.Input<inputs.BranchProtectionRestrictPush>[]>;
}
