// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.

import {Comment, Issue, Repository, User} from "../../types";
import {WebHooks} from "../../hooks";
import {WebHookBase} from "../webhook";

// Handler is the function signature for serverless event handlers.
export type Handler = (e: Event) => Promise<void>;

// WebHook represents an active push webhook; that is, both a subscription and a handler.
export class WebHook extends WebHookBase {
    constructor(hooks: WebHooks, handler: Handler) {
        super(hooks, "issue_comment", async (body) => {
            await handler(<Event>body);
        });
    }
}

// See https://developer.github.com/v3/activity/events/types/#issuecommentevent.
export interface Event {
    action: IssueCommentAction; // The action that was performed on the comment.
    issue: Issue; // The issue the comment belongs to.
    comment: Comment; // The comment itself.
    repository: Repository; // The repository hosting the issue.
    sender: User; // The user triggering the comment event.
}

export type IssueCommentAction = "created" | "edited" | "deleted";

