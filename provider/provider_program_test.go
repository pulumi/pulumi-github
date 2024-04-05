//go:build !go && !nodejs && !python && !dotnet
// +build !go,!nodejs,!python,!dotnet

// Copyright 2016-2024, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License")
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package github

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/pulumi/providertest/pulumitest"
	"github.com/pulumi/providertest/pulumitest/opttest"
)

const providerName = "github"

func testProgram(t *testing.T, dir string) {
	if testing.Short() {
		t.Skipf("Skipping in testing.Short() mode, assuming this is a CI run without credentials")
	}
	cwd, err := os.Getwd()
	require.NoError(t, err)
	test := pulumitest.NewPulumiTest(t, dir,
		opttest.LocalProviderPath(providerName, filepath.Join(cwd, "..", "bin")),
		opttest.SkipInstall(),
	)
	test.Up()
}

func TestPrograms(t *testing.T) {
	programs := []string{
		"test-programs/index_repository",
		"test-programs/index_actionssecret",
		"test-programs/index_team",
		"test-programs/index_branchprotection",
		"test-programs/index_teamrepository",
		"test-programs/index_branchdefault",
		"test-programs/index_membership",
		"test-programs/index_actionsorganizationsecret",
		"test-programs/index_repositoryenvironment",
		"test-programs/index_teammembership",
		"test-programs/index_actionsenvironmentsecret",
		"test-programs/index_branch",
		"test-programs/index_repositorydeploykey",
		"test-programs/index_actionsenvironmentvariable",
		"test-programs/index_actionsvariable",
		"test-programs/index_repositoryfile",
		"test-programs/index_repositorywebhook",
		"test-programs/index_organizationsettings",
		"test-programs/index_branchprotectionv3",
		"test-programs/index_repositorycollaborator",
	}
	for _, p := range programs {
		t.Run(p, func(t *testing.T) {
			testProgram(t, p)
		})
	}
}
