// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.

export let user: string;  // the GitHub identity to assume.
export let token: string; // the token to use for API calls.
export let repo: string;  // the repo to use for all webhooks, API calls, etc.

export function requireUser(): string {
    if (user === undefined) {
        throw new Error("GitHub user was not configured properly");
    }
    return user;
}

export function requireToken(): string {
    if (token === undefined) {
        throw new Error("GitHub token was not configured properly");
    }
    return token;
}

export function requireRepo(): string {
    if (repo === undefined) {
        throw new Error("GitHub repo was not configured properly");
    }
    return repo;
}

