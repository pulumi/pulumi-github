// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "pulumi";

// TODO: https://github.com/pulumi/pulumi-terraform/issues/111
// Rename `content_type` and `insecure_ssl` to `contentType` and `insecureSsl`
// to match the naming scheme we use everywhere else once we're able to
// indicate a map's names should be transformed. Until then, we need to keep
// the names as `content_type` and `insecure_ssl` as those are the names that
// GitHub expects.

export interface WebhookConfiguration {
    readonly url: pulumi.Input<string>;
    readonly content_type?: pulumi.Input<ContentType>;
    readonly secret?: pulumi.Input<string>;
    readonly insecure_ssl?: pulumi.Input<string>;
}

export type ContentType = "form" | "json";
