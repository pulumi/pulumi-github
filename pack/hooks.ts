// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.

import * as config from "./config";
import * as webhooks from "./webhooks";
import * as crypto from "crypto";
import {HttpAPI} from "@pulumi/pulumi";

// WebHooks represents a shared webhooks API endpoint and associated handlers.
export class WebHooks {
    public readonly prefix: string;                // the globally unique resource name prefix.
    public readonly gateway: HttpAPI;                  // the shared API gateway for all handlers.
    public readonly hooks: webhooks.WebHookBase[]; // the list of registered webhooks for this endpoint.

    constructor() {
        // Make sure the username and repo are configured, and use that information to set up a prefix.
        let userRepoHash: crypto.Hash = crypto.createHash("sha1");
        userRepoHash.update(config.user + "-" + config.repo);
        this.prefix  = "webhooks-" + userRepoHash.digest("hex");
        this.gateway = new HttpAPI(this.prefix + "-api");
        this.hooks   = [];
    }

    // listen is called when you are finished registering APIs, to make the hooks available.
    public listen(): void {
        let url = this.gateway.publish();
        for (let i = 0; i < (<any>this.hooks).length; i++) {
            this.hooks[i].activate(url);
        }
    }

    // onIssueComment registers a handler that runs anytime a comment on an issue is created, edited, or deleted.
    public onIssueComment(handler: webhooks.issueComment.Handler): void {
        (<any>this.hooks).push(new webhooks.issueComment.WebHook(this, handler));
    }

    // onPush registers a handler that runs on any push to a repository, including editing tags or branches.
    public onPush(handler: webhooks.push.Handler): void {
        (<any>this.hooks).push(new webhooks.push.WebHook(this, handler));
    }
}

