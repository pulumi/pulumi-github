package examples

import (
	"path"
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
)

func TestAccRepoTs(t *testing.T) {
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: path.Join(getCwd(t), "repo", "ts"),
		})

	integration.ProgramTest(t, &test)
}

func getJSBaseOptions(t *testing.T) integration.ProgramTestOptions {
	base := getBaseOptions(t)
	baseJS := base.With(integration.ProgramTestOptions{
		Dependencies: []string{
			"@pulumi/github",
		},
	})

	return baseJS
}
