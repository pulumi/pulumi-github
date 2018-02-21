// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

package main

import (
	"github.com/pulumi/pulumi-terraform/pkg/tfbridge"

	"github.com/pulumi/pulumi-github"
	"github.com/pulumi/pulumi-github/pkg/version"
)

func main() {
	tfbridge.Main("github", version.Version, github.Provider())
}
