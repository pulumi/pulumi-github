// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.

package ghctx

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-fabric/pkg/tokens"
)

var config map[tokens.ModuleMember]string

const (
	configUser  = tokens.ModuleMember("github:config:user")
	configRepo  = tokens.ModuleMember("github:config:repo")
	configToken = tokens.ModuleMember("github:config:token")
)

// Configure configures the user, repo, and token.
func Configure(vars map[tokens.ModuleMember]string) error {
	config = vars
	return nil
}

// User fetches the current confiured GitHub user.
func User() (string, error) {
	if user, has := config[configUser]; has {
		return user, nil
	}
	return "", errors.Errorf("missing required user configuration '%v'", configUser)
}

// Repo fetches the current confiured GitHub repo.
func Repo() (string, error) {
	if repo, has := config[configRepo]; has {
		return repo, nil
	}
	return "", errors.Errorf("missing required repo configuration '%v'", configRepo)
}

// Token fetches the current confiured GitHub auth token.
func Token() (string, error) {
	if token, has := config[configToken]; has {
		return token, nil
	}
	return "", errors.Errorf("missing required token configuration '%v'", configToken)
}
