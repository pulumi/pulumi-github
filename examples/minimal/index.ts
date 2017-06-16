// Licensed to Pulumi Corporation ("Pulumi") under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// Pulumi licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
            `*Issue #${<any>e.issue.number} ${e.issue.title}* by ${e.issue.user.login}\n` +
                `${e.comment.body}\n` +
            `(with :love_letter: from Lumi)`,
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

