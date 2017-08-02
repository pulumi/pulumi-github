// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.

package issues

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-fabric/pkg/resource/provider"

	"github.com/pulumi/pulumi-github/provider/ghctx"
)

// defaultRepo fetches either the repo provided by the repop argument, or gets it from host configuration.
func defaultRepo(host *provider.HostClient, repop *string) (string, error) {
	if repop != nil {
		return *repop, nil
	}
	// If repo is nil, fetch the default one configured for this package.
	r, err := ghctx.Repo(host)
	if err != nil {
		return "", err
	}
	return r, nil
}

// defaultBaseURL fetches the base GitHub API endpoint for all repo API management.
func defaultBaseURL(host *provider.HostClient, repo string) (string, error) {
	base, err := ghctx.BaseURL(host)
	if err != nil {
		return "", err
	}
	return base + "/repos/" + repo, nil
}

// checkCommon performs check that are common to all repo resources.
func checkCommon(host *provider.HostClient, repop *string) error {
	// Ensure that we can create a base URL.
	if _, err := ghctx.BaseURL(host); err != nil {
		return err
	}
	// Ensure that, if there is no repo provided, a value has been configured.
	repo, err := defaultRepo(host, repop)
	if err != nil {
		return errors.Errorf("No repo provided on this label, and configuration lookup failed: %v", err)
	} else if repo == "" {
		return errors.Errorf("Invalid blank repo configuration")
	}
	return nil
}
