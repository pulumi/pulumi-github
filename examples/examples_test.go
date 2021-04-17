package examples

import (
	"os"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

func getBaseOptions(t *testing.T) integration.ProgramTestOptions {
	token := checkGitHubTestingToken(t)
	org := getGitHubTestingOrg(t)
	return integration.ProgramTestOptions{
		ExpectRefreshChanges: true,
		Secrets: map[string]string{
			"github:token":        token,
			"github:organization": org,
		},
	}
}

func getGitHubTestingOrg(t *testing.T) string {
	org := os.Getenv("GH_ORGANIZATION")
	if org == "" {
		t.Skipf("Skipping test due to missing GH_ORGANIZATION environment variable")
	}

	return org
}

func checkGitHubTestingToken(t *testing.T) string {
	token := os.Getenv("GH_TESTING_TOKEN")
	if token == "" {
		t.Skipf("Skipping test due to missing GH_TESTING_TOKEN environment variable")
	}

	return token
}

func getCwd(t *testing.T) string {
	cwd, err := os.Getwd()
	if err != nil {
		t.FailNow()
	}

	return cwd
}
