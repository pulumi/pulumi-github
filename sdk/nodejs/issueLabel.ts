// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## Import
 *
 * GitHub Issue Labels can be imported using an ID made up of `repository:name`, e.g.
 *
 * ```sh
 *  $ pulumi import github:index/issueLabel:IssueLabel panic_label terraform:panic
 * ```
 */
export class IssueLabel extends pulumi.CustomResource {
    /**
     * Get an existing IssueLabel resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IssueLabelState, opts?: pulumi.CustomResourceOptions): IssueLabel {
        return new IssueLabel(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'github:index/issueLabel:IssueLabel';

    /**
     * Returns true if the given object is an instance of IssueLabel.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IssueLabel {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IssueLabel.__pulumiType;
    }

    /**
     * A 6 character hex code, **without the leading #**, identifying the color of the label.
     */
    public readonly color!: pulumi.Output<string>;
    /**
     * A short description of the label.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * The name of the label.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The GitHub repository
     */
    public readonly repository!: pulumi.Output<string>;
    /**
     * The URL to the issue label
     */
    public /*out*/ readonly url!: pulumi.Output<string>;

    /**
     * Create a IssueLabel resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IssueLabelArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IssueLabelArgs | IssueLabelState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as IssueLabelState | undefined;
            inputs["color"] = state ? state.color : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["repository"] = state ? state.repository : undefined;
            inputs["url"] = state ? state.url : undefined;
        } else {
            const args = argsOrState as IssueLabelArgs | undefined;
            if ((!args || args.color === undefined) && !opts.urn) {
                throw new Error("Missing required property 'color'");
            }
            if ((!args || args.repository === undefined) && !opts.urn) {
                throw new Error("Missing required property 'repository'");
            }
            inputs["color"] = args ? args.color : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["repository"] = args ? args.repository : undefined;
            inputs["etag"] = undefined /*out*/;
            inputs["url"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(IssueLabel.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IssueLabel resources.
 */
export interface IssueLabelState {
    /**
     * A 6 character hex code, **without the leading #**, identifying the color of the label.
     */
    color?: pulumi.Input<string>;
    /**
     * A short description of the label.
     */
    description?: pulumi.Input<string>;
    etag?: pulumi.Input<string>;
    /**
     * The name of the label.
     */
    name?: pulumi.Input<string>;
    /**
     * The GitHub repository
     */
    repository?: pulumi.Input<string>;
    /**
     * The URL to the issue label
     */
    url?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a IssueLabel resource.
 */
export interface IssueLabelArgs {
    /**
     * A 6 character hex code, **without the leading #**, identifying the color of the label.
     */
    color: pulumi.Input<string>;
    /**
     * A short description of the label.
     */
    description?: pulumi.Input<string>;
    /**
     * The name of the label.
     */
    name?: pulumi.Input<string>;
    /**
     * The GitHub repository
     */
    repository: pulumi.Input<string>;
}
