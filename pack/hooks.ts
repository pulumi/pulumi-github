// Copyright 2017 Pulumi, Inc. All rights reserved.

import * as config from "./config";
import * as webhooks from "./webhooks";
import {API} from "@lumi/aws/serverless";

// Make sure the username and repo are configured, and use that information to set up a prefix.
let user: string = config.requireUser();
let repo: string = config.requireRepo();

// WebHooks represents a shared webhooks API endpoint and associated handlers.
export class WebHooks {
    public readonly prefix: string; // the globally unique resource name prefix.
    public readonly gateway: API;   // the shared API gateway for all handlers.

    constructor() {
        this.prefix = "webhooks-" + user + "-" + repo;
    }

    // publish is called when you are finished registering APIs, to make the hooks available.
    // TODO: ideally we wouldn't need this.
    public publish(): void {
        this.gateway.publish(this.prefix + "-stage");
    }

    // onIssueComment registers a handler that runs anytime a comment on an issue is created, edited, or deleted.
    public onIssueComment(handler: webhooks.issueComment.Handler): void {
        new webhooks.issueComment.WebHook(this, handler);
    }

    // onPush registers a handler that runs on any push to a repository, including editing tags or branches.
    public onPush(handler: webhooks.push.Handler): void {
        new webhooks.push.WebHook(this, handler);
    }
}

