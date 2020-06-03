package examples

import (
	"os"
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
)

func getBaseOptions(t *testing.T) integration.ProgramTestOptions {
	checkGitHubToken(t)
	return integration.ProgramTestOptions{
		ExpectRefreshChanges: true,
	}
}

func checkGitHubToken(t *testing.T) {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		t.Skipf("Skipping test due to missing GITHUB_TOKEN environment variable")
	}
}

func getCwd(t *testing.T) string {
	cwd, err := os.Getwd()
	if err != nil {
		t.FailNow()
	}

	return cwd
}
