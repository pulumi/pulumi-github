// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

import {WebhookConfiguration} from "../index";

/**
 * This resource allows you to create and manage webhooks for repositories within your
 * GitHub organization.
 * 
 * This resource cannot currently be used to manage webhooks for *personal* repositories,
 * outside of organizations.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as github from "@pulumi/github";
 * 
 * const repo = new github.repos.Repository("repo", {
 *     description: "Terraform acceptance tests",
 *     homepageUrl: "http://example.com/",
 *     private: false,
 * });
 * const foo = new github.repos.Webhook("foo", {
 *     active: false,
 *     configuration: {
 *         contentType: "form",
 *         insecureSsl: "false",
 *         url: "https://google.de/",
 *     },
 *     events: ["issues"],
 *     repository: repo.name,
 * });
 * ```
 */
export class Webhook extends pulumi.CustomResource {
    /**
     * Get an existing Webhook resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WebhookState, opts?: pulumi.CustomResourceOptions): Webhook {
        return new Webhook(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'github:repos/webhook:Webhook';

    /**
     * Returns true if the given object is an instance of Webhook.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Webhook {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Webhook.__pulumiType;
    }

    /**
     * Indicate of the webhook should receive events. Defaults to `true`.
     */
    public readonly active!: pulumi.Output<boolean | undefined>;
    /**
     * key/value pair of configuration for this webhook. Available keys are `url`, `content_type`, `secret` and `insecure_ssl`. `secret` is [the shared secret, see API documentation](https://developer.github.com/v3/repos/hooks/#create-a-hook).
     */
    public readonly configuration!: pulumi.Output<WebhookConfiguration | undefined>;
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * A list of events which should trigger the webhook. See a list of [available events](https://developer.github.com/v3/activity/events/types/).
     */
    public readonly events!: pulumi.Output<string[]>;
    /**
     * The repository of the webhook.
     */
    public readonly repository!: pulumi.Output<string>;
    /**
     * URL of the webhook
     */
    public /*out*/ readonly url!: pulumi.Output<string>;

    /**
     * Create a Webhook resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WebhookArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WebhookArgs | WebhookState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as WebhookState | undefined;
            inputs["active"] = state ? state.active : undefined;
            inputs["configuration"] = state ? state.configuration : undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["events"] = state ? state.events : undefined;
            inputs["repository"] = state ? state.repository : undefined;
            inputs["url"] = state ? state.url : undefined;
        } else {
            const args = argsOrState as WebhookArgs | undefined;
            if (!args || args.events === undefined) {
                throw new Error("Missing required property 'events'");
            }
            if (!args || args.repository === undefined) {
                throw new Error("Missing required property 'repository'");
            }
            inputs["active"] = args ? args.active : undefined;
            inputs["configuration"] = args ? args.configuration : undefined;
            inputs["events"] = args ? args.events : undefined;
            inputs["repository"] = args ? args.repository : undefined;
            inputs["etag"] = undefined /*out*/;
            inputs["url"] = undefined /*out*/;
        }
        super(Webhook.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Webhook resources.
 */
export interface WebhookState {
    /**
     * Indicate of the webhook should receive events. Defaults to `true`.
     */
    readonly active?: pulumi.Input<boolean>;
    /**
     * key/value pair of configuration for this webhook. Available keys are `url`, `content_type`, `secret` and `insecure_ssl`. `secret` is [the shared secret, see API documentation](https://developer.github.com/v3/repos/hooks/#create-a-hook).
     */
    readonly configuration?: pulumi.Input<WebhookConfiguration>;
    readonly etag?: pulumi.Input<string>;
    /**
     * A list of events which should trigger the webhook. See a list of [available events](https://developer.github.com/v3/activity/events/types/).
     */
    readonly events?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The repository of the webhook.
     */
    readonly repository?: pulumi.Input<string>;
    /**
     * URL of the webhook
     */
    readonly url?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Webhook resource.
 */
export interface WebhookArgs {
    /**
     * Indicate of the webhook should receive events. Defaults to `true`.
     */
    readonly active?: pulumi.Input<boolean>;
    /**
     * key/value pair of configuration for this webhook. Available keys are `url`, `content_type`, `secret` and `insecure_ssl`. `secret` is [the shared secret, see API documentation](https://developer.github.com/v3/repos/hooks/#create-a-hook).
     */
    readonly configuration?: pulumi.Input<WebhookConfiguration>;
    /**
     * A list of events which should trigger the webhook. See a list of [available events](https://developer.github.com/v3/activity/events/types/).
     */
    readonly events: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The repository of the webhook.
     */
    readonly repository: pulumi.Input<string>;
}
