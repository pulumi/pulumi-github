// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package github

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"unicode"

	// embed package blank import
	_ "embed"

	"github.com/integrations/terraform-provider-github/v6/github"

	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	tfbridgetokens "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/tokens"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"

	"github.com/pulumi/pulumi-github/provider/v6/pkg/version"
)

// all of the token components used below.
const (
	// packages:
	mainPkg = "github"
	// modules:
	mainMod = "index" // the y module
)

// makeMember manufactures a type token for the package and the given module and type.
func makeMember(mod string, mem string) tokens.ModuleMember {
	return tokens.ModuleMember(mainPkg + ":" + mod + ":" + mem)
}

// makeType manufactures a type token for the package and the given module and type.
func makeType(mod string, typ string) tokens.Type {
	return tokens.Type(makeMember(mod, typ))
}

// makeDataSource manufactures a standard resource token given a module and resource name.  It
// automatically uses the main package and names the file by simply lower casing the data source's
// first character.
func makeDataSource(mod string, res string) tokens.ModuleMember {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return makeMember(mod+"/"+fn, res)
}

// makeResource manufactures a standard resource token given a module and resource name.  It
// automatically uses the main package and names the file by simply lower casing the resource's
// first character.
func makeResource(mod string, res string) tokens.Type {
	lower := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return makeType(mod+"/"+lower, res)
}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := shimv2.NewProvider(github.Provider(),
		shimv2.WithPlanResourceChange(func(string) bool { return true }))

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:                       p,
		Name:                    "github",
		Description:             "A Pulumi package for creating and managing github cloud resources.",
		Keywords:                []string{"pulumi", "github"},
		License:                 "Apache-2.0",
		Homepage:                "https://pulumi.io",
		Repository:              "https://github.com/pulumi/pulumi-github",
		TFProviderModuleVersion: "v6",
		GitHubOrg:               "integrations",
		Version:                 version.Version,
		MetadataInfo:            tfbridge.NewProviderMetadata(metadata),
		UpstreamRepoPath:        "./upstream",
		DocRules:                &tfbridge.DocRuleInfo{EditRules: editRules},
		Config: map[string]*tfbridge.SchemaInfo{
			"base_url": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"GITHUB_BASE_URL"},
					Value:   "https://api.github.com/",
				},
			},
			"token": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"GITHUB_TOKEN"},
				},
				Secret: tfbridge.True(),
			},
		},
		Resources: map[string]*tfbridge.ResourceInfo{
			"github_actions_environment_secret": {DeleteBeforeReplace: true},
			"github_actions_organization_secret": {
				Fields: map[string]*tfbridge.SchemaInfo{
					"selected_repository_ids": {
						ForceNew: tfbridge.True(),
					},
				},
			},

			"github_actions_repository_permissions": {Tok: makeResource(mainMod, "ActionsRepositoryPermissions")},
			"github_actions_runner_group":           {Tok: makeResource(mainMod, "ActionsRunnerGroup")},
			"github_actions_secret":                 {Tok: makeResource(mainMod, "ActionsSecret")},
			"github_app_installation_repository":    {Tok: makeResource(mainMod, "AppInstallationRepository")},
			"github_app_installation_repositories":  {Tok: makeResource(mainMod, "AppInstallationRepositories")},

			"github_branch": {
				Tok: makeResource(mainMod, "Branch"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"branch": {
						CSharpName: "BranchName",
					},
				},
			},
			"github_branch_default": {Tok: makeResource(mainMod, "BranchDefault")},
			"github_branch_protection": {
				Tok:                makeResource(mainMod, "BranchProtection"),
				TransformFromState: nil,
			},
			"github_branch_protection_v3": {Tok: makeResource(mainMod, "BranchProtectionV3")},
			"github_codespaces_organization_secret": {
				Fields: map[string]*tfbridge.SchemaInfo{
					"selected_repository_ids": {
						ForceNew: tfbridge.True(),
					},
				},
			},
			"github_dependabot_organization_secret": {
				Tok:  makeResource(mainMod, "DependabotOrganizationSecret"),
				Docs: &tfbridge.DocInfo{AllowMissing: true},
				Fields: map[string]*tfbridge.SchemaInfo{
					"selected_repository_ids": {
						ForceNew: tfbridge.True(),
					},
				},
			},
			"github_dependabot_organization_secret_repositories": {
				Tok:  makeResource(mainMod, "DependabotOrganizationSecretRepositories"),
				Docs: &tfbridge.DocInfo{AllowMissing: true},
			},
			"github_dependabot_secret": {Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"github_emu_group_mapping":             {Tok: makeResource(mainMod, "EmuGroupMapping")},
			"github_issue":                         {Tok: makeResource(mainMod, "Issue")},
			"github_issue_label":                   {Tok: makeResource(mainMod, "IssueLabel")},
			"github_membership":                    {Tok: makeResource(mainMod, "Membership")},
			"github_organization_block":            {Tok: makeResource(mainMod, "OrganizationBlock")},
			"github_organization_project":          {Tok: makeResource(mainMod, "OrganizationProject")},
			"github_organization_security_manager": {Tok: makeResource(mainMod, "OrganizationSecurityManager")},
			"github_organization_settings":         {Tok: makeResource(mainMod, "OrganizationSettings")},
			"github_organization_webhook":          {Tok: makeResource(mainMod, "OrganizationWebhook")},
			"github_project_card":                  {Tok: makeResource(mainMod, "ProjectCard")},
			"github_project_column":                {Tok: makeResource(mainMod, "ProjectColumn")},
			"github_release":                       {Tok: makeResource(mainMod, "Release")},

			"github_repository": {Tok: makeResource(mainMod, "Repository")},
			"github_repository_collaborator": {
				Tok:                 makeResource(mainMod, "RepositoryCollaborator"),
				DeleteBeforeReplace: true,
			},
			"github_repository_dependabot_security_updates": {Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"github_repository_deploy_key": {
				Tok:                 makeResource(mainMod, "RepositoryDeployKey"),
				DeleteBeforeReplace: true,
			},
			"github_repository_environment":    {Tok: makeResource(mainMod, "RepositoryEnvironment")},
			"github_repository_file":           {Tok: makeResource(mainMod, "RepositoryFile")},
			"github_repository_pull_request":   {Tok: makeResource(mainMod, "RepositoryPullRequest")},
			"github_repository_project":        {Tok: makeResource(mainMod, "RepositoryProject")},
			"github_repository_tag_protection": {Tok: makeResource(mainMod, "RepositoryTagProtection")},
			"github_repository_webhook":        {Tok: makeResource(mainMod, "RepositoryWebhook")},
			"github_repository_milestone":      {Tok: makeResource(mainMod, "RepositoryMilestone")},
			"github_repository_autolink_reference": {
				Tok:                 makeResource(mainMod, "RepositoryAutolinkReference"),
				DeleteBeforeReplace: true,
			},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"github_actions_public_key": {Tok: makeDataSource(mainMod, "getActionsPublicKey")},
			"github_actions_organization_registration_token": {
				Tok: makeDataSource(mainMod, "getActionsOrganizationRegistrationToken"),
			},
			"github_actions_organization_secrets": {
				Tok: makeDataSource(mainMod, "getActionsOrganizationSecrets"),
			},
			"github_actions_registration_token": {
				Tok: makeDataSource(mainMod, "getActionsRegistrationToken"),
			},
			"github_actions_secrets":                 {Tok: makeDataSource(mainMod, "getActionsSecrets")},
			"github_app":                             {Tok: makeDataSource(mainMod, "getGithubApp")},
			"github_branch":                          {Tok: makeDataSource(mainMod, "getBranch")},
			"github_collaborators":                   {Tok: makeDataSource(mainMod, "getCollaborators")},
			"github_dependabot_organization_secrets": {Tok: makeDataSource(mainMod, "getDependabotOrganizationSecrets")},
			"github_dependabot_public_key": {
				Tok:  makeDataSource(mainMod, "getDependabotPublicKey"),
				Docs: &tfbridge.DocInfo{AllowMissing: true},
			},
			"github_dependabot_secrets":            {Tok: makeDataSource(mainMod, "getDependabotSecrets")},
			"github_external_groups":               {Tok: makeDataSource(mainMod, "getExternalGroups")},
			"github_ip_ranges":                     {Tok: makeDataSource(mainMod, "getIpRanges")},
			"github_membership":                    {Tok: makeDataSource(mainMod, "getMembership")},
			"github_organization":                  {Tok: makeDataSource(mainMod, "getOrganization")},
			"github_organization_ip_allow_list":    {Tok: makeDataSource(mainMod, "getOrganizationIpAllowList")},
			"github_organization_team_sync_groups": {Tok: makeDataSource(mainMod, "getOrganizationTeamSyncGroups")},
			"github_organization_teams":            {Tok: makeDataSource(mainMod, "getOrganizationTeams")},
			"github_organization_webhooks":         {Tok: makeDataSource(mainMod, "getOrganizationWebhooks")},
			"github_release":                       {Tok: makeDataSource(mainMod, "getRelease")},
			"github_repositories":                  {Tok: makeDataSource(mainMod, "getRepositories")},
			"github_repository":                    {Tok: makeDataSource(mainMod, "getRepository")},
			"github_repository_branches":           {Tok: makeDataSource(mainMod, "getRepositoryBranches")},
			"github_repository_deploy_keys":        {Tok: makeDataSource(mainMod, "getRepositoryDeployKeys")},
			"github_repository_file":               {Tok: makeDataSource(mainMod, "getRepositoryFile")},
			"github_repository_milestone":          {Tok: makeDataSource(mainMod, "getRepositoryMilestone")},
			"github_repository_pull_request":       {Tok: makeDataSource(mainMod, "getRepositoryPullRequest")},
			"github_repository_pull_requests":      {Tok: makeDataSource(mainMod, "getRepositoryPullRequests")},
			"github_repository_teams":              {Tok: makeDataSource(mainMod, "getRepositoryTeams")},
			"github_repository_webhooks":           {Tok: makeDataSource(mainMod, "getRepositoryWebhooks")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			// List any npm dependencies and their versions

			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
			RespectSchemaVersion: true,
		},
		Python: &tfbridge.PythonInfo{
			RespectSchemaVersion: true,

			PyProject: struct{ Enabled bool }{true},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/pulumi/pulumi-%[1]s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
			RespectSchemaVersion:           true,
		},
		CSharp: &tfbridge.CSharpInfo{
			RespectSchemaVersion: true,
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
		EnableZeroDefaultSchemaVersion: true,
	}

	prov.MustComputeTokens(tfbridgetokens.SingleModule("github_", mainMod,
		tfbridgetokens.MakeStandard(mainPkg)))
	prov.MustApplyAutoAliases()

	prov.SetAutonaming(255, "-")

	return prov
}

// Docs Edits

// Ensure the text of IssueLabel makes it into the API documentation.
// It contains important information on a courtesy labels import-before-create.
var ensureIssueLabelsContent = tfbridge.DocsEdit{
	Path: "*issue_label.html.markdown",
	Edit: func(_ string, content []byte) ([]byte, error) {
		content = bytes.ReplaceAll(content,
			[]byte("Terraform"), []byte("Pulumi"))
		return content, nil
	},
}

func editRules(defaults []tfbridge.DocsEdit) []tfbridge.DocsEdit {
	// This rule strips a section that would otherwise be changed by default edits. It should run before.
	edits := []tfbridge.DocsEdit{
		{
			Path: "index.html.markdown",
			Edit: func(_ string, content []byte) ([]byte, error) {
				input, err := os.ReadFile("provider/installation-replaces/required-providers-input.md")
				if err != nil {
					return nil, err
				}
				content = bytes.ReplaceAll(
					content,
					input,
					nil)
				return content, nil
			},
		},
	}
	// Append default edits
	edits = append(edits, defaults...)
	// Append additional rules
	return append(edits,
		ensureIssueLabelsContent,
		tfbridge.DocsEdit{
			Path: "index.html.markdown",
			Edit: func(_ string, content []byte) ([]byte, error) {
				files := []string{
					"note1-input.md",
					"note2-input.md",
				}
				for _, file := range files {
					input, err := os.ReadFile("provider/installation-replaces/" + file)
					if err != nil {
						return nil, err
					}
					content = bytes.ReplaceAll(
						content,
						input,
						nil)

				}
				return content, nil
			},
		})
}

//go:embed cmd/pulumi-resource-github/bridge-metadata.json
var metadata []byte
