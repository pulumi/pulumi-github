// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

export interface WebhookConfiguration {
    readonly url: pulumi.Input<string>;
    readonly contentType?: pulumi.Input<ContentType>;
    readonly secret?: pulumi.Input<string>;
    readonly insecureSsl?: pulumi.Input<string>;
}

export type ContentType = "form" | "json";
