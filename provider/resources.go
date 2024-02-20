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
	"fmt"
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
	p := shimv2.NewProvider(github.Provider())

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
		Config: map[string]*tfbridge.SchemaInfo{
			"base_url": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"GITHUB_BASE_URL"},
					Value:   "https://api.github.com/",
				},
			},
		},
		Resources: map[string]*tfbridge.ResourceInfo{
			"github_actions_environment_secret": {
				Tok:                 makeResource(mainMod, "ActionsEnvironmentSecret"),
				DeleteBeforeReplace: true,
			},
			"github_actions_organization_secret":      {Tok: makeResource(mainMod, "ActionsOrganizationSecret")},
			"github_actions_organization_permissions": {Tok: makeResource(mainMod, "ActionsOrganizationPermissions")},
			"github_actions_organization_secret_repositories": {
				Tok: makeResource(mainMod, "ActionsOrganizationSecretRepositories"),
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
			"github_branch_default":       {Tok: makeResource(mainMod, "BranchDefault")},
			"github_branch_protection":    {Tok: makeResource(mainMod, "BranchProtection")},
			"github_branch_protection_v3": {Tok: makeResource(mainMod, "BranchProtectionV3")},
			"github_dependabot_organization_secret": {
				Tok:  makeResource(mainMod, "DependabotOrganizationSecret"),
				Docs: &tfbridge.DocInfo{AllowMissing: true},
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

			"github_repository":                             {Tok: makeResource(mainMod, "Repository")},
			"github_repository_collaborator":                {Tok: makeResource(mainMod, "RepositoryCollaborator")},
			"github_repository_dependabot_security_updates": {Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"github_repository_deploy_key": {
				Tok:                 makeResource(mainMod, "RepositoryDeployKey"),
				DeleteBeforeReplace: true},
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
			"github_team":                     {Tok: makeResource(mainMod, "Team")},
			"github_team_members":             {Tok: makeResource(mainMod, "TeamMembers")},
			"github_team_membership":          {Tok: makeResource(mainMod, "TeamMembership")},
			"github_team_repository":          {Tok: makeResource(mainMod, "TeamRepository")},
			"github_team_settings":            {Tok: makeResource(mainMod, "TeamSettings")},
			"github_team_sync_group_mapping":  {Tok: makeResource(mainMod, "TeamSyncGroupMapping")},
			"github_user_gpg_key":             {Tok: makeResource(mainMod, "UserGpgKey")},
			"github_user_invitation_accepter": {Tok: makeResource(mainMod, "UserInvitationAccepter")},
			"github_user_ssh_key":             {Tok: makeResource(mainMod, "UserSshKey")},
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
			"github_team":                          {Tok: makeDataSource(mainMod, "getTeam")},
			"github_user":                          {Tok: makeDataSource(mainMod, "getUser")},
			"github_users":                         {Tok: makeDataSource(mainMod, "getUsers")},
			"github_ref":                           {Tok: makeDataSource(mainMod, "getRef")},
			"github_tree":                          {Tok: makeDataSource(mainMod, "getTree")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
		},
		Python: &tfbridge.PythonInfo{
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
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
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
	}

	prov.MustComputeTokens(tfbridgetokens.SingleModule("github_", mainMod,
		tfbridgetokens.MakeStandard(mainPkg)))
	prov.MustApplyAutoAliases()

	prov.SetAutonaming(255, "-")

	return prov
}

//go:embed cmd/pulumi-resource-github/bridge-metadata.json
var metadata []byte
