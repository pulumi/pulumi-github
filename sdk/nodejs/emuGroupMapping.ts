// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This resource manages mappings between external groups for enterprise managed users and GitHub teams. It wraps the API detailed [here](https://docs.github.com/en/rest/reference/teams#external-groups). Note that this is a distinct resource from `github.TeamSyncGroupMapping`. `github.EmuGroupMapping` is special to the Enterprise Managed User (EMU) external group feature, whereas `github.TeamSyncGroupMapping` is specific to Identity Provider Groups.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as github from "@pulumi/github";
 *
 * const exampleEmuGroupMapping = new github.EmuGroupMapping("example_emu_group_mapping", {
 *     teamSlug: "emu-test-team",
 *     groupId: 28836,
 * });
 * ```
 *
 * ## Import
 *
 * GitHub EMU External Group Mappings can be imported using the external `group_id`, e.g.
 *
 * ```sh
 * $ pulumi import github:index/emuGroupMapping:EmuGroupMapping example_emu_group_mapping 28836
 * ```
 */
export class EmuGroupMapping extends pulumi.CustomResource {
    /**
     * Get an existing EmuGroupMapping resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EmuGroupMappingState, opts?: pulumi.CustomResourceOptions): EmuGroupMapping {
        return new EmuGroupMapping(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'github:index/emuGroupMapping:EmuGroupMapping';

    /**
     * Returns true if the given object is an instance of EmuGroupMapping.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EmuGroupMapping {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EmuGroupMapping.__pulumiType;
    }

    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * Integer corresponding to the external group ID to be linked
     */
    public readonly groupId!: pulumi.Output<number>;
    /**
     * Slug of the GitHub team
     */
    public readonly teamSlug!: pulumi.Output<string>;

    /**
     * Create a EmuGroupMapping resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EmuGroupMappingArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EmuGroupMappingArgs | EmuGroupMappingState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EmuGroupMappingState | undefined;
            resourceInputs["etag"] = state ? state.etag : undefined;
            resourceInputs["groupId"] = state ? state.groupId : undefined;
            resourceInputs["teamSlug"] = state ? state.teamSlug : undefined;
        } else {
            const args = argsOrState as EmuGroupMappingArgs | undefined;
            if ((!args || args.groupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'groupId'");
            }
            if ((!args || args.teamSlug === undefined) && !opts.urn) {
                throw new Error("Missing required property 'teamSlug'");
            }
            resourceInputs["groupId"] = args ? args.groupId : undefined;
            resourceInputs["teamSlug"] = args ? args.teamSlug : undefined;
            resourceInputs["etag"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(EmuGroupMapping.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EmuGroupMapping resources.
 */
export interface EmuGroupMappingState {
    etag?: pulumi.Input<string>;
    /**
     * Integer corresponding to the external group ID to be linked
     */
    groupId?: pulumi.Input<number>;
    /**
     * Slug of the GitHub team
     */
    teamSlug?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a EmuGroupMapping resource.
 */
export interface EmuGroupMappingArgs {
    /**
     * Integer corresponding to the external group ID to be linked
     */
    groupId: pulumi.Input<number>;
    /**
     * Slug of the GitHub team
     */
    teamSlug: pulumi.Input<string>;
}
