// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This resource allows you to create and manage an OpenID Connect subject claim customization template within a GitHub
 * organization.
 *
 * More information on integrating GitHub with cloud providers using OpenID Connect and a list of available claims is
 * available in the [Actions documentation](https://docs.github.com/en/actions/deployment/security-hardening-your-deployments/about-security-hardening-with-openid-connect).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as github from "@pulumi/github";
 *
 * const exampleTemplate = new github.ActionsOrganizationOidcSubjectClaimCustomizationTemplate("example_template", {includeClaimKeys: [
 *     "actor",
 *     "context",
 *     "repository_owner",
 * ]});
 * ```
 *
 * ## Import
 *
 * This resource can be imported using the organization's name.
 *
 * ```sh
 * $ pulumi import github:index/actionsOrganizationOidcSubjectClaimCustomizationTemplate:ActionsOrganizationOidcSubjectClaimCustomizationTemplate test example_organization
 * ```
 */
export class ActionsOrganizationOidcSubjectClaimCustomizationTemplate extends pulumi.CustomResource {
    /**
     * Get an existing ActionsOrganizationOidcSubjectClaimCustomizationTemplate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ActionsOrganizationOidcSubjectClaimCustomizationTemplateState, opts?: pulumi.CustomResourceOptions): ActionsOrganizationOidcSubjectClaimCustomizationTemplate {
        return new ActionsOrganizationOidcSubjectClaimCustomizationTemplate(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'github:index/actionsOrganizationOidcSubjectClaimCustomizationTemplate:ActionsOrganizationOidcSubjectClaimCustomizationTemplate';

    /**
     * Returns true if the given object is an instance of ActionsOrganizationOidcSubjectClaimCustomizationTemplate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ActionsOrganizationOidcSubjectClaimCustomizationTemplate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ActionsOrganizationOidcSubjectClaimCustomizationTemplate.__pulumiType;
    }

    /**
     * A list of OpenID Connect claims.
     */
    public readonly includeClaimKeys!: pulumi.Output<string[]>;

    /**
     * Create a ActionsOrganizationOidcSubjectClaimCustomizationTemplate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ActionsOrganizationOidcSubjectClaimCustomizationTemplateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ActionsOrganizationOidcSubjectClaimCustomizationTemplateArgs | ActionsOrganizationOidcSubjectClaimCustomizationTemplateState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ActionsOrganizationOidcSubjectClaimCustomizationTemplateState | undefined;
            resourceInputs["includeClaimKeys"] = state ? state.includeClaimKeys : undefined;
        } else {
            const args = argsOrState as ActionsOrganizationOidcSubjectClaimCustomizationTemplateArgs | undefined;
            if ((!args || args.includeClaimKeys === undefined) && !opts.urn) {
                throw new Error("Missing required property 'includeClaimKeys'");
            }
            resourceInputs["includeClaimKeys"] = args ? args.includeClaimKeys : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ActionsOrganizationOidcSubjectClaimCustomizationTemplate.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ActionsOrganizationOidcSubjectClaimCustomizationTemplate resources.
 */
export interface ActionsOrganizationOidcSubjectClaimCustomizationTemplateState {
    /**
     * A list of OpenID Connect claims.
     */
    includeClaimKeys?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a ActionsOrganizationOidcSubjectClaimCustomizationTemplate resource.
 */
export interface ActionsOrganizationOidcSubjectClaimCustomizationTemplateArgs {
    /**
     * A list of OpenID Connect claims.
     */
    includeClaimKeys: pulumi.Input<pulumi.Input<string>[]>;
}
