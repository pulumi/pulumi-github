// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.

import * as config from "./config";
import * as github from "@pulumi/github";

let hooks = new github.WebHooks();

declare let require;
hooks.onIssueComment(async (e) => {
    console.log(e);
    if (e.action === "created") {
        let slack = require("@slack/client");
        let client = new slack.WebClient(config.slackToken);
        try {
            console.log("posting message...")
            await client.chat.postMessage(
                "#issue-spam",
                `*Issue #${<any>e.issue.number} ${e.issue.title}* by ${e.comment.user.login}\n` +
                    `${e.comment.body}\n` +
                `(with :love_letter: from Pulumi)`,
                {
                    as_user: true,
                }
            );    
        } catch (err) {
            console.log("error posting message")
            throw new Error(err);
        }
    }
});

hooks.listen();

