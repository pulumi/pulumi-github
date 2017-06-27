// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.

package issues

import (
	"github.com/pulumi/lumi/pkg/resource/idl"
)

// Labels can be applied to issues and pull requests to signify priority, category, or any other information that you
// find useful.
type Label struct {
	idl.Resource
	// The name of the label.  This must be unique within a repo.
	Name string `lumi:"name,replaces"`
	// A 6 character hex code, without the leading #, identifying the color.
	Color string `lumi:"color"`
	// The Repo in which this label exists; if empty, this is read from configuration.
	Repo *string `lumi:"repo,optional,replaces"`
}
