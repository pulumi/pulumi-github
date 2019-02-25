// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

package examples

import (
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pulumi/pulumi/pkg/testing/integration"
)

func TestExamples(t *testing.T) {
	required := []string{"GITHUB_ORGANIZATION", "GITHUB_TOKEN"}
	for _, v := range required {
		if os.Getenv(v) == "" {
			t.Skipf("Skipping test due to missing %q environment variable. (%v are required to run this test.)", v, required)
		}
	}

	cwd, err := os.Getwd()
	if !assert.NoError(t, err, "expected a valid working directory: %v", err) {
		return
	}

	// base options shared amongst all tests.
	base := integration.ProgramTestOptions{
		Config: map[string]string{
			"github:organization": os.Getenv("GITHUB_ORGANIZATION"),
			"github:token":        os.Getenv("GITHUB_TOKEN"),
		},
		Dependencies: []string{
			"@pulumi/github",
			"@pulumi/pulumi",
		},
	}

	shortTests := []integration.ProgramTestOptions{
		base.With(integration.ProgramTestOptions{Dir: path.Join(cwd, "minimal")}),
	}

	longTests := []integration.ProgramTestOptions{
		base.With(integration.ProgramTestOptions{Dir: path.Join(cwd, "orgs")}),
		base.With(integration.ProgramTestOptions{Dir: path.Join(cwd, "repos")}),
		base.With(integration.ProgramTestOptions{Dir: path.Join(cwd, "teams")}),
	}

	if !testing.Short() {
		tests = append(examples, longTests...)
	}

	for _, ex := range tests {
		example := ex
		t.Run(example.Dir, func(t *testing.T) {
			integration.ProgramTest(t, &example)
		})
	}
}
