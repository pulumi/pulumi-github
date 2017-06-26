// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.

export let slackToken: string; // a required configuration variable.
export function requireSlackToken(): string {
    if (slackToken == undefined || slackToken == "") {
        throw new Error("Missing required gh:config:slackToken configuration variable");
    }
    return slackToken;
}

