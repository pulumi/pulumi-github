// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class AppInstallationRepositories extends pulumi.CustomResource {
    /**
     * Get an existing AppInstallationRepositories resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AppInstallationRepositoriesState, opts?: pulumi.CustomResourceOptions): AppInstallationRepositories {
        return new AppInstallationRepositories(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'github:index/appInstallationRepositories:AppInstallationRepositories';

    /**
     * Returns true if the given object is an instance of AppInstallationRepositories.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AppInstallationRepositories {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AppInstallationRepositories.__pulumiType;
    }

    public readonly installationId!: pulumi.Output<string>;
    public readonly selectedRepositories!: pulumi.Output<string[]>;

    /**
     * Create a AppInstallationRepositories resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AppInstallationRepositoriesArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AppInstallationRepositoriesArgs | AppInstallationRepositoriesState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AppInstallationRepositoriesState | undefined;
            resourceInputs["installationId"] = state ? state.installationId : undefined;
            resourceInputs["selectedRepositories"] = state ? state.selectedRepositories : undefined;
        } else {
            const args = argsOrState as AppInstallationRepositoriesArgs | undefined;
            if ((!args || args.installationId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'installationId'");
            }
            if ((!args || args.selectedRepositories === undefined) && !opts.urn) {
                throw new Error("Missing required property 'selectedRepositories'");
            }
            resourceInputs["installationId"] = args ? args.installationId : undefined;
            resourceInputs["selectedRepositories"] = args ? args.selectedRepositories : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AppInstallationRepositories.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AppInstallationRepositories resources.
 */
export interface AppInstallationRepositoriesState {
    installationId?: pulumi.Input<string>;
    selectedRepositories?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a AppInstallationRepositories resource.
 */
export interface AppInstallationRepositoriesArgs {
    installationId: pulumi.Input<string>;
    selectedRepositories: pulumi.Input<pulumi.Input<string>[]>;
}
