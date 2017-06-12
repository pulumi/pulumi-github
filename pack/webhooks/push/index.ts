// Copyright 2017 Pulumi, Inc. All rights reserved.

import {Repository, User} from "../../types";
import {WebHooks} from "../../hooks";
import {WebHookBase} from "../webhook";
import {Context} from "@lumi/aws/serverless";

// Handler is the function signature for serverless event handlers.
export type Handler = (e: Event, context: Context, callback: (error: any, result: any) => void) => any;

// WebHook represents an active push webhook; that is, both a subscription and a handler.
export class WebHook extends WebHookBase {
    constructor(hooks: WebHooks, handler: Handler) {
        super(hooks, "push", (e: any, context: Context, callback: (error: any, result: any) => void) => {
            handler(<Event>e, context, callback);
        });
    }
}

// See https://developer.github.com/v3/activity/events/types/#pushevent.
export interface Event {
    ref: string; // The full git ref that was pushed. Example: `"refs/heads/master"`.
    before: string; // The SHA of the most recent commit on `ref` before the push.
    after: string; // The SHA of the most recent commit on `ref` after the push.
    created: boolean;
    deleted: boolean;
    forced: boolean;
    baseRef?: string;
    compare: string;
    commits: Commit[]; // An array of commit objects describing the pushed commits.
    headCommit: Commit; // The most recent (head) commit.
    repository: Repository; // The repository hosting the pull request.
    pusher: CommitUser; // The user pushing the commit.
    sender: User; // The user triggering the pull request event.
}

export interface Commit {
    id: string; // The SHA of the commit.
    distinct: boolean; // Whether the commit is distinct from any that have been pushed before.
    message: string; // The commit message.
    timestamp: string;
    url: string; // Points to the commit API resource.
    author: CommitUser; // The git author of the commit.
    committer: CommitUser; // The git committer of the commit.
    added: string[]; // The files added by this commit.
    removed: string[]; // The files removed by this commit.
    modified: string[]; // The files modified by this commit.
}

export interface CommitUser {
    name: string; // The git author's name.
    email: string; // The git author's email address.
    username?: string; // The git author's username.
}

