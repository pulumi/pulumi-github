// *** WARNING: this file was generated by the Lumi IDL Compiler (LUMIDL). ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

/* tslint:disable:ordered-imports variable-name */
import * as lumi from "@lumi/lumi";

export let FormContentType: ContentType = "form";
export let JSONContentType: ContentType = "json";

export interface Config {
    url: string;
    contentType?: ContentType;
    secret?: string;
    insecureSSL?: boolean;
}

export type ContentType =
    "form" |
    "json";

export class Subscription extends lumi.NamedResource implements SubscriptionArgs {
    public config: Config;
    public events?: string[];
    public active?: boolean;

    constructor(name: string, args: SubscriptionArgs) {
        super(name);
        if (args.config === undefined) {
            throw new Error("Missing required argument 'config'");
        }
        this.config = args.config;
        this.events = args.events;
        this.active = args.active;
    }
}

export interface SubscriptionArgs {
    config: Config;
    events?: string[];
    active?: boolean;
}


