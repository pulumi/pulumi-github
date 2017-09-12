// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.

import {Config} from "@pulumi/pulumi-fabric";

let _config = new Config("gh-serverless-ftw:config");
export let slackToken: string = _config.require("slackToken"); // a required configuration variable.

