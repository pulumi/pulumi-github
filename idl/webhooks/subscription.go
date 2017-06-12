// Copyright 2017 Pulumi, Inc. All rights reserved.

package push

import (
	"github.com/pulumi/lumi/pkg/resource/idl"
)

// Subscription represents an active webhook subscription for a given GitHub repository.
type Subscription struct {
	idl.NamedResource

	// A repo that this subscription pertains to.  TODO: take this from config.
	Repo string `lumi:"repo,replaces"`
	// An API token to use for API calls.  TODO: take this from config.
	Token string `lumi:"token"`

	// The events array determines what events the hook is triggered for. Default: ["push"].
	Events *[]string `lumi:"events,optional"`

	// A required string defining the URL to which the payloads will be delivered.
	URL string `lumi:"url,replaces"`
	// An optional string defining the media type used to serialize the payloads.  The default is `form`.
	ContentType *ContentType `lumi:"contentType,optional"`
	// An optional string that's passed with the HTTP requests as an `X-Hub-Signature` header.  The value of this
	// header is computed as the HMAC hex digest of the body, using the secret as the key.
	Secret *string `lumi:"secret,optional"`
	// An optional setting that determines whether the SSL certificate of the host for url will be verified when
	// delivering payloads.  The default is off.a
	InsecureSSL *bool `lumi:"insecureSSL,optional"`
}

type ContentType string

const (
	FormContentType ContentType = "form"
	JSONContentType ContentType = "json"
)
