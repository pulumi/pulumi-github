// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This resource allows you to create and manage projects for GitHub organization.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as github from "@pulumi/github";
 *
 * const project = new github.OrganizationProject("project", {
 *     name: "A Organization Project",
 *     body: "This is a organization project.",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export class OrganizationProject extends pulumi.CustomResource {
    /**
     * Get an existing OrganizationProject resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: OrganizationProjectState, opts?: pulumi.CustomResourceOptions): OrganizationProject {
        return new OrganizationProject(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'github:index/organizationProject:OrganizationProject';

    /**
     * Returns true if the given object is an instance of OrganizationProject.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OrganizationProject {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OrganizationProject.__pulumiType;
    }

    /**
     * The body of the project.
     */
    public readonly body!: pulumi.Output<string | undefined>;
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * The name of the project.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * URL of the project
     */
    public /*out*/ readonly url!: pulumi.Output<string>;

    /**
     * Create a OrganizationProject resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: OrganizationProjectArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: OrganizationProjectArgs | OrganizationProjectState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as OrganizationProjectState | undefined;
            resourceInputs["body"] = state ? state.body : undefined;
            resourceInputs["etag"] = state ? state.etag : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["url"] = state ? state.url : undefined;
        } else {
            const args = argsOrState as OrganizationProjectArgs | undefined;
            resourceInputs["body"] = args ? args.body : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["etag"] = undefined /*out*/;
            resourceInputs["url"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(OrganizationProject.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering OrganizationProject resources.
 */
export interface OrganizationProjectState {
    /**
     * The body of the project.
     */
    body?: pulumi.Input<string>;
    etag?: pulumi.Input<string>;
    /**
     * The name of the project.
     */
    name?: pulumi.Input<string>;
    /**
     * URL of the project
     */
    url?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a OrganizationProject resource.
 */
export interface OrganizationProjectArgs {
    /**
     * The body of the project.
     */
    body?: pulumi.Input<string>;
    /**
     * The name of the project.
     */
    name?: pulumi.Input<string>;
}
