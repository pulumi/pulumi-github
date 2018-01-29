// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

// TODO: https://github.com/pulumi/pulumi-terraform/issues/111
// Rename `content_type` and `insecure_ssl` to `contentType` and `insecureSsl`
// to match the naming scheme we use everywhere else once we're able to
// indicate a map's names should be transformed. Until then, we need to keep
// the names as `content_type` and `insecure_ssl` as those are the names that
// GitHub expects.

export interface WebhookConfiguration {
    readonly url: string;
    readonly content_type?: ContentType;
    readonly secret?: string;
    readonly insecure_ssl?: string;
}

export type ContentType = "form" | "json";
