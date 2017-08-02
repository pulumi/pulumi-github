// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.

package ghctx

import (
	"github.com/pulumi/pulumi-fabric/pkg/resource/provider"
)

// BaseURL returns the base URL to use for all GitHub API calls.
func BaseURL(host *provider.HostClient) (string, error) {
	user, err := User(host)
	if err != nil {
		return "", err
	}
	token, err := Token(host)
	if err != nil {
		return "", err
	}
	return "https://" + user + ":" + token + "@api.github.com", nil
}
