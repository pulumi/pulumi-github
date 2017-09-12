// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.

import * as config from "../config";
import {WebHooks} from "../hooks";
import {Subscription} from "./subscription";
import {Computed} from "@pulumi/pulumi-fabric";

export type WebHookHandler = (body: any) => Promise<void>;

// WebHookBase combines a serverless function with a GitHub webhook subscription.
export class WebHookBase {
    private parent: WebHooks;                       // the parent webhooks object.
    private event: string;                          // the GitHub webhook event name.
    private sub: Subscription;                      // the subscription, once it is initialized.

    constructor(parent: WebHooks, event: string, handler: WebHookHandler) {
        this.parent = parent;
        this.event = event;

        // Associate an endpoint with the handler.  Note that the hooks object ensures that the user/repo pair are
        // uniquely isolated from all others, such that the URL needn't actually contain them.
        this.parent.gateway.post(`/${event}`, [], async(req, res) => {
            try {
                let body = (<any>JSON).parse(req.body);
                await handler(body);
                console.log("Returning 200 success.");
                res.status(200).end("");
            } catch(err) {
                console.log("Returning 500 error: " + err);
                res.status(500).end("");
            }
        });
    }

    // activate tells GitHub to call the specified URL in response to the right kind of events.
    public activate(url: Computed<string>) {
        this.sub = new Subscription(this.parent.prefix + "-sub", {
            service: "web",
            active: true,
            events: [ this.event ],
            config: {
                contentType: "json",
                url: url.mapValue((u: string) => u + `/${this.event}`),
            },
        });
    }
}

