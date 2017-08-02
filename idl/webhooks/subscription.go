// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.

package webhooks

import (
	"github.com/pulumi/pulumi-fabric/pkg/resource/idl"
)

// Subscription represents an active webhook subscription for a given GitHub repository.
type Subscription struct {
	idl.NamedResource

	// Service is the name of the webhook service.  Use `web` for a webhook.
	Service string `lumi:"service"`
	// Config contains the various webhook configuration options.
	Config Config `lumi:"config"`

	// The events array determines what events the hook is triggered for. Default: ["push"].
	Events *[]string `lumi:"events,optional"`
	// Active indicates whether this hook is actually triggered on the given events.
	Active *bool `lumi:"active,optional"`
}

// Config represents a webhook's subscription configuration.
type Config struct {
	// A required string defining the URL to which the payloads will be delivered.
	URL string `lumi:"url"`
	// An optional string defining the media type used to serialize the payloads.  The default is `form`.
	ContentType *ContentType `lumi:"contentType,optional"`
	// An optional string that's passed with the HTTP requests as an `X-Hub-Signature` header.  The value of this
	// header is computed as the HMAC hex digest of the body, using the secret as the key.
	Secret *string `lumi:"secret,optional"`
	// An optional setting that determines whether the SSL certificate of the host for url will be verified when
	// delivering payloads.  The default is off.
	InsecureSSL *bool `lumi:"insecureSSL,optional"`
}

type ContentType string

const (
	FormContentType ContentType = "form"
	JSONContentType ContentType = "json"
)
