// Copyright 2017 Pulumi, Inc. All rights reserved.

import * as config from "../config";
import {API} from "@lumi/aws/serverless";

// Make sure the username and repo are configured, and use that information to set up a prefix.
let user: string = config.requireUser();
let repo: string = config.requireRepo();
export let prefix: string = "webhook-" + user + "-" + repo;

// Set up one global per-user / per-repo API gateway to handle all incoming webhooks for that combination.
export let gateway = new API(prefix + "-api");
// TODO: figure out how and when to publish the stage.

