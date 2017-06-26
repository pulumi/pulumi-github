// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.

import * as config from "../config";
import {WebHooks} from "../hooks";
import {Subscription} from "./subscription";
import * as aws from "@lumi/aws";

// TODO: everything below is currently AWS-specific.  Eventually we want to make this cloud agnostic.

// WebHookBase combines a serverless function with a GitHub webhook subscription.
export class WebHookBase {
    private parent: WebHooks;                       // the parent webhooks object.
    private event: string;                          // the GitHub webhook event name.
    private readonly func: aws.serverless.Function; // the function to invoke in response to the event.
    private sub: Subscription;                      // the subscription, once it is initialized.

    constructor(parent: WebHooks, event: string, handler: aws.serverless.Handler) {
        this.parent = parent;
        this.event = event;

        // New up a lambda that will invoke the handler code.
        this.func = new aws.serverless.Function(
            this.parent.prefix + "-" + event + "-handler",
            [ aws.iam.AWSLambdaFullAccess ], // TODO: consider making this configurable.
            handler,
        );

        // Associate an endpoint with the handler.  Note that the hooks object ensures that the user/repo pair are
        // uniquely isolated from all others, such that the URL needn't actually contain them.
        this.parent.gateway.route("POST", `/${event}`, this.func);
    }

    // activate tells GitHub to call the specified URL in response to the right kind of events.
    public activate(url: string) {
        this.sub = new Subscription(this.parent.prefix + "-sub", {
            service: "web",
            active: true,
            events: [ this.event ],
            config: {
                contentType: "json",
                url: url + `/${this.event}`,
            },
        });
    }
}

