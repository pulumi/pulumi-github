// Copyright 2017 Pulumi, Inc. All rights reserved.

import * as config from "../config";
import {WebHooks} from "../hooks";
import {Subscription} from "./subscription";
import * as aws from "@lumi/aws";

// TODO: everything below is currently AWS-specific.  Eventually we want to make this cloud agnostic.

// WebHookBase combines a serverless function with a GitHub webhook subscription.
export class WebHookBase {
    private readonly func: aws.serverless.Function;
    private readonly sub: Subscription;

    constructor(hooks: WebHooks, event: string, handler: aws.serverless.Handler) {
        // Ensure the required GitHub configuration is available before even proceeding.
        let user: string = config.requireUser();
        let repo: string = config.requireRepo();
        let token: string = config.requireToken();
        let prefix: string = hooks.prefix + "-" + event;

        // New up a lambda that will invoke the handler code.
        this.func = new aws.serverless.Function(
            prefix + "-handler",
            [ aws.iam.AWSLambdaFullAccess ],
            handler,
        );

        // Associate an endpoint with the handler.
        hooks.gateway.route("POST", `/${user}/${repo}/${event}`, this.func);

        // Finally, tell GitHub to poke our endpoint in response to hook events.
        this.sub = new Subscription(prefix + "-sub", {
            repo: repo,
            token: token,
            events: [ event ],
            url: "TODO"/*hooks.api.url?*/,
        });
    }
}

