// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.

import {requireSlackToken} from "./config";
import * as github from "@lumi/github";

let hooks = new github.WebHooks();
let slackToken = requireSlackToken();

declare let require;
hooks.onIssueComment((e, callback) => {
    if (e.action == "created") {
        let slack = require("@slack/client");
        let client = new slack.WebClient(slackToken);
        client.chat.postMessage(
            "#issue-spam",
            `*Issue #${<any>e.issue.number} ${e.issue.title}* by ${e.comment.user.login}\n` +
                `${e.comment.body}\n` +
            `(with :love_letter: from Pulumi)`,
            {
                as_user: true,
            },
            (err, res) => {
                if (err) {
                    console.log(err);
                    callback(null, { statusCode: 500 });
                } else {
                    callback(null, { statusCode: 200 });
                }
            },
        );
    } else {
        callback(null, { statusCode: 200 });
    }
});

hooks.listen();

