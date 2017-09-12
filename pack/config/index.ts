// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.

import {Config} from "@pulumi/pulumi-fabric";

let _config = new Config("github:config");

export let user: string = _config.require("user");   // the GitHub identity to assume.
export let repo: string = _config.require("repo");;  // the repo to use for all webhooks, API calls, etc.
export let token: string = _config.require("token"); // the token to use for API calls.

