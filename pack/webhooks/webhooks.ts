// Copyright 2017 Pulumi, Inc. All rights reserved.

import * as issueComment from "./issueComment";
import * as push from "./push";

// onIssueComment registers a handler that runs anytime a comment on an issue is created, edited, or deleted.
export function onIssueComment(handler: issueComment.Handler) {
    new issueComment.WebHook(handler);
}

// onPush registers a handler that runs on any push to a repository, including editing tags or branches.
export function onPush(handler: push.Handler) {
    new push.WebHook(handler);
}

