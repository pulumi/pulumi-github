// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

declare var exports: any;
const __config = new pulumi.Config("github");

/**
 * The GitHub App credentials used to connect to GitHub. Conflicts with `token`. Anonymous mode is enabled if both `token`
 * and `appAuth` are not set.
 */
export declare const appAuth: outputs.config.AppAuth | undefined;
Object.defineProperty(exports, "appAuth", {
    get() {
        return __config.getObject<outputs.config.AppAuth>("appAuth");
    },
    enumerable: true,
});

/**
 * The GitHub Base API URL
 */
export declare const baseUrl: string;
Object.defineProperty(exports, "baseUrl", {
    get() {
        return __config.get("baseUrl") ?? (utilities.getEnv("GITHUB_BASE_URL") || "https://api.github.com/");
    },
    enumerable: true,
});

/**
 * Enable `insecure` mode for testing purposes
 */
export declare const insecure: boolean | undefined;
Object.defineProperty(exports, "insecure", {
    get() {
        return __config.getObject<boolean>("insecure");
    },
    enumerable: true,
});

/**
 * Number of times to retry a request after receiving an error status codeDefaults to 3
 */
export declare const maxRetries: number | undefined;
Object.defineProperty(exports, "maxRetries", {
    get() {
        return __config.getObject<number>("maxRetries");
    },
    enumerable: true,
});

/**
 * The GitHub organization name to manage. Use this field instead of `owner` when managing organization accounts.
 */
export declare const organization: string | undefined;
Object.defineProperty(exports, "organization", {
    get() {
        return __config.get("organization");
    },
    enumerable: true,
});

/**
 * The GitHub owner name to manage. Use this field instead of `organization` when managing individual accounts.
 */
export declare const owner: string | undefined;
Object.defineProperty(exports, "owner", {
    get() {
        return __config.get("owner");
    },
    enumerable: true,
});

/**
 * Allow the provider to make parallel API calls to GitHub. You may want to set it to true when you have a private Github
 * Enterprise without strict rate limits. Although, it is not possible to enable this setting on github.com because we
 * enforce the respect of github.com's best practices to avoid hitting abuse rate limitsDefaults to false if not set
 */
export declare const parallelRequests: boolean | undefined;
Object.defineProperty(exports, "parallelRequests", {
    get() {
        return __config.getObject<boolean>("parallelRequests");
    },
    enumerable: true,
});

/**
 * Amount of time in milliseconds to sleep in between non-write requests to GitHub API. Defaults to 0ms if not set.
 */
export declare const readDelayMs: number | undefined;
Object.defineProperty(exports, "readDelayMs", {
    get() {
        return __config.getObject<number>("readDelayMs");
    },
    enumerable: true,
});

/**
 * Amount of time in milliseconds to sleep in between requests to GitHub API after an error response. Defaults to 1000ms or
 * 1s if not set, the maxRetries must be set to greater than zero.
 */
export declare const retryDelayMs: number | undefined;
Object.defineProperty(exports, "retryDelayMs", {
    get() {
        return __config.getObject<number>("retryDelayMs");
    },
    enumerable: true,
});

/**
 * Allow the provider to retry after receiving an error status code, the maxRetries should be set for this to workDefaults
 * to [500, 502, 503, 504]
 */
export declare const retryableErrors: number[] | undefined;
Object.defineProperty(exports, "retryableErrors", {
    get() {
        return __config.getObject<number[]>("retryableErrors");
    },
    enumerable: true,
});

/**
 * The OAuth token used to connect to GitHub. Anonymous mode is enabled if both `token` and `appAuth` are not set.
 */
export declare const token: string | undefined;
Object.defineProperty(exports, "token", {
    get() {
        return __config.get("token") ?? utilities.getEnv("GITHUB_TOKEN");
    },
    enumerable: true,
});

/**
 * Amount of time in milliseconds to sleep in between writes to GitHub API. Defaults to 1000ms or 1s if not set.
 */
export declare const writeDelayMs: number | undefined;
Object.defineProperty(exports, "writeDelayMs", {
    get() {
        return __config.getObject<number>("writeDelayMs");
    },
    enumerable: true,
});

