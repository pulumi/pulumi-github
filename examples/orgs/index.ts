// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as github from "@pulumi/github";

const name = "orgs";

// Add a membership
const membership = new github.orgs.Membership(name, {
    username: "e38ce0e-collaborator",
    role: "member"
});

// Add an organization webhook
const webhook = new github.orgs.Webhook(name, {
    configuration: github.WebhookConfiguration = {
        url: "https://google.com",
        content_type: github.ContentTypes.Form,
    },
    events: ["issues"],
    active: false
});

export const webhookId = webhook.id
