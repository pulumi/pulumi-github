// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.

package ghctx

import (
	"github.com/pulumi/pulumi-fabric/pkg/resource/provider"
)

const (
	configUser  = "github:config:user"
	configRepo  = "github:config:repo"
	configToken = "github:config:token"
)

// User fetches the current confiured GitHub user.
func User(host *provider.HostClient) (string, error) {
	return host.ReadStringLocation(configUser)
}

// Repo fetches the current confiured GitHub repo.
func Repo(host *provider.HostClient) (string, error) {
	return host.ReadStringLocation(configRepo)
}

// Token fetches the current confiured GitHub auth token.
func Token(host *provider.HostClient) (string, error) {
	return host.ReadStringLocation(configToken)
}
