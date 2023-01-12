// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.
//go:build go || all
// +build go all

package examples

import (
	"path"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

func TestAccRepoGo(t *testing.T) {
	test := getGoBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: path.Join(getCwd(t), "repo", "go"),
		})

	integration.ProgramTest(t, &test)
}

func getGoBaseOptions(t *testing.T) integration.ProgramTestOptions {
	base := getBaseOptions(t)
	baseGo := base.With(integration.ProgramTestOptions{
		Dependencies: []string{
			"github.com/pulumi/pulumi-github/sdk/v5",
		},
	})

	return baseGo
}
