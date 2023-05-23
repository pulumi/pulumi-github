// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.
// +build nodejs all

package examples

import (
	"path"
	"testing"
	"encoding/json"
	"github.com/stretchr/testify/assert"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

func TestAccRepoTs(t *testing.T) {
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: path.Join(getCwd(t), "repo", "ts"),
			// Temporary profilactic check until pulumi/pulumi#12981 is resolved.
			ExtraRuntimeValidation: func(t *testing.T, stack integration.RuntimeValidationStackInfo) {
			for _, e := range stack.Events {
				eventsJSON, err := json.MarshalIndent(e, "", "  ")
				assert.NoError(t, err)
				assert.NotContainsf(t, string(eventsJSON), "panic",
					"Unexpected panic recorded in engine events")
			}
		},
		})

	integration.ProgramTest(t, &test)
}

func getJSBaseOptions(t *testing.T) integration.ProgramTestOptions {
	base := getBaseOptions(t)
	baseJS := base.With(integration.ProgramTestOptions{
		Dependencies: []string{
			"@pulumi/github",
		},
		// Temporary profilactic check until pulumi/pulumi#12981 is resolved.
		ExtraRuntimeValidation: func(t *testing.T, stack integration.RuntimeValidationStackInfo) {
			for _, e := range stack.Events {
				eventsJSON, err := json.MarshalIndent(e, "", "  ")
				assert.NoError(t, err)
				assert.NotContainsf(t, string(eventsJSON), "panic",
					"Unexpected panic recorded in engine events")
			}
		},
	})

	return baseJS
}
