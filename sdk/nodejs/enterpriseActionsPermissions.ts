// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * This resource allows you to create and manage GitHub Actions permissions within your GitHub enterprise.
 * You must have admin access to an enterprise to use this resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as github from "@pulumi/github";
 *
 * const example-enterprise = github.getEnterprise({
 *     slug: "my-enterprise",
 * });
 * const example-org = github.getOrganization({
 *     name: "my-org",
 * });
 * const test = new github.EnterpriseActionsPermissions("test", {
 *     enterpriseId: example_enterprise.then(example_enterprise => example_enterprise.slug),
 *     allowedActions: "selected",
 *     enabledOrganizations: "selected",
 *     allowedActionsConfig: {
 *         githubOwnedAllowed: true,
 *         patternsAlloweds: [
 *             "actions/cache@*",
 *             "actions/checkout@*",
 *         ],
 *         verifiedAllowed: true,
 *     },
 *     enabledOrganizationsConfig: {
 *         organizationIds: [example_org.then(example_org => example_org.id)],
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * This resource can be imported using the name of the GitHub enterprise:
 *
 * ```sh
 * $ pulumi import github:index/enterpriseActionsPermissions:EnterpriseActionsPermissions test github_enterprise_name
 * ```
 */
export class EnterpriseActionsPermissions extends pulumi.CustomResource {
    /**
     * Get an existing EnterpriseActionsPermissions resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EnterpriseActionsPermissionsState, opts?: pulumi.CustomResourceOptions): EnterpriseActionsPermissions {
        return new EnterpriseActionsPermissions(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'github:index/enterpriseActionsPermissions:EnterpriseActionsPermissions';

    /**
     * Returns true if the given object is an instance of EnterpriseActionsPermissions.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EnterpriseActionsPermissions {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EnterpriseActionsPermissions.__pulumiType;
    }

    /**
     * The permissions policy that controls the actions that are allowed to run. Can be one of: `all`, `localOnly`, or `selected`.
     */
    public readonly allowedActions!: pulumi.Output<string | undefined>;
    /**
     * Sets the actions that are allowed in an enterprise. Only available when `allowedActions` = `selected`. See Allowed Actions Config below for details.
     */
    public readonly allowedActionsConfig!: pulumi.Output<outputs.EnterpriseActionsPermissionsAllowedActionsConfig | undefined>;
    /**
     * The policy that controls the organizations in the enterprise that are allowed to run GitHub Actions. Can be one of: `all`, `none`, or `selected`.
     */
    public readonly enabledOrganizations!: pulumi.Output<string>;
    /**
     * Sets the list of selected organizations that are enabled for GitHub Actions in an enterprise. Only available when `enabledOrganizations` = `selected`. See Enabled Organizations Config below for details.
     */
    public readonly enabledOrganizationsConfig!: pulumi.Output<outputs.EnterpriseActionsPermissionsEnabledOrganizationsConfig | undefined>;
    /**
     * The ID of the enterprise.
     */
    public readonly enterpriseId!: pulumi.Output<string>;

    /**
     * Create a EnterpriseActionsPermissions resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EnterpriseActionsPermissionsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EnterpriseActionsPermissionsArgs | EnterpriseActionsPermissionsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EnterpriseActionsPermissionsState | undefined;
            resourceInputs["allowedActions"] = state ? state.allowedActions : undefined;
            resourceInputs["allowedActionsConfig"] = state ? state.allowedActionsConfig : undefined;
            resourceInputs["enabledOrganizations"] = state ? state.enabledOrganizations : undefined;
            resourceInputs["enabledOrganizationsConfig"] = state ? state.enabledOrganizationsConfig : undefined;
            resourceInputs["enterpriseId"] = state ? state.enterpriseId : undefined;
        } else {
            const args = argsOrState as EnterpriseActionsPermissionsArgs | undefined;
            if ((!args || args.enabledOrganizations === undefined) && !opts.urn) {
                throw new Error("Missing required property 'enabledOrganizations'");
            }
            if ((!args || args.enterpriseId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'enterpriseId'");
            }
            resourceInputs["allowedActions"] = args ? args.allowedActions : undefined;
            resourceInputs["allowedActionsConfig"] = args ? args.allowedActionsConfig : undefined;
            resourceInputs["enabledOrganizations"] = args ? args.enabledOrganizations : undefined;
            resourceInputs["enabledOrganizationsConfig"] = args ? args.enabledOrganizationsConfig : undefined;
            resourceInputs["enterpriseId"] = args ? args.enterpriseId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(EnterpriseActionsPermissions.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EnterpriseActionsPermissions resources.
 */
export interface EnterpriseActionsPermissionsState {
    /**
     * The permissions policy that controls the actions that are allowed to run. Can be one of: `all`, `localOnly`, or `selected`.
     */
    allowedActions?: pulumi.Input<string>;
    /**
     * Sets the actions that are allowed in an enterprise. Only available when `allowedActions` = `selected`. See Allowed Actions Config below for details.
     */
    allowedActionsConfig?: pulumi.Input<inputs.EnterpriseActionsPermissionsAllowedActionsConfig>;
    /**
     * The policy that controls the organizations in the enterprise that are allowed to run GitHub Actions. Can be one of: `all`, `none`, or `selected`.
     */
    enabledOrganizations?: pulumi.Input<string>;
    /**
     * Sets the list of selected organizations that are enabled for GitHub Actions in an enterprise. Only available when `enabledOrganizations` = `selected`. See Enabled Organizations Config below for details.
     */
    enabledOrganizationsConfig?: pulumi.Input<inputs.EnterpriseActionsPermissionsEnabledOrganizationsConfig>;
    /**
     * The ID of the enterprise.
     */
    enterpriseId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a EnterpriseActionsPermissions resource.
 */
export interface EnterpriseActionsPermissionsArgs {
    /**
     * The permissions policy that controls the actions that are allowed to run. Can be one of: `all`, `localOnly`, or `selected`.
     */
    allowedActions?: pulumi.Input<string>;
    /**
     * Sets the actions that are allowed in an enterprise. Only available when `allowedActions` = `selected`. See Allowed Actions Config below for details.
     */
    allowedActionsConfig?: pulumi.Input<inputs.EnterpriseActionsPermissionsAllowedActionsConfig>;
    /**
     * The policy that controls the organizations in the enterprise that are allowed to run GitHub Actions. Can be one of: `all`, `none`, or `selected`.
     */
    enabledOrganizations: pulumi.Input<string>;
    /**
     * Sets the list of selected organizations that are enabled for GitHub Actions in an enterprise. Only available when `enabledOrganizations` = `selected`. See Enabled Organizations Config below for details.
     */
    enabledOrganizationsConfig?: pulumi.Input<inputs.EnterpriseActionsPermissionsEnabledOrganizationsConfig>;
    /**
     * The ID of the enterprise.
     */
    enterpriseId: pulumi.Input<string>;
}
