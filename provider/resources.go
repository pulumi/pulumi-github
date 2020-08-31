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
	"unicode"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/pulumi/pulumi-terraform-bridge/v2/pkg/tfbridge"
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v2/go/common/tokens"
	"github.com/terraform-providers/terraform-provider-github/github"
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
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return makeType(mod+"/"+fn, res)
}

// preConfigureCallback is called before the providerConfigure function of the underlying provider.
// It should validate that the provider can be configured, and provide actionable errors in the case
// it cannot be. Configuration variables can be read from `vars` using the `stringValue` function -
// for example `stringValue(vars, "accessKey")`.
func preConfigureCallback(vars resource.PropertyMap, c *terraform.ResourceConfig) error {
	return nil
}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := github.Provider().(*schema.Provider)

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:           p,
		Name:        "github",
		Description: "A Pulumi package for creating and managing github cloud resources.",
		Keywords:    []string{"pulumi", "github"},
		License:     "Apache-2.0",
		Homepage:    "https://pulumi.io",
		Repository:  "https://github.com/pulumi/pulumi-github",
		Config: map[string]*tfbridge.SchemaInfo{
			"token": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"GITHUB_TOKEN"},
				},
			},
			"organization": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"GITHUB_ORGANIZATION"},
				},
			},
			"base_url": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"GITHUB_BASE_URL"},
					Value:   "https://api.github.com/",
				},
			},
		},
		PreConfigureCallback: preConfigureCallback,
		Resources: map[string]*tfbridge.ResourceInfo{
			"github_actions_secret": {Tok: makeResource(mainMod, "ActionsSecret")},
			"github_branch": {
				Tok: makeResource(mainMod, "Branch"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"branch": {
						CSharpName: "BranchName",
					},
				},
			},
			"github_branch_protection":        {Tok: makeResource(mainMod, "BranchProtection")},
			"github_issue_label":              {Tok: makeResource(mainMod, "IssueLabel")},
			"github_membership":               {Tok: makeResource(mainMod, "Membership")},
			"github_organization_block":       {Tok: makeResource(mainMod, "OrganizationBlock")},
			"github_organization_project":     {Tok: makeResource(mainMod, "OrganizationProject")},
			"github_organization_webhook":     {Tok: makeResource(mainMod, "OrganizationWebhook")},
			"github_project_column":           {Tok: makeResource(mainMod, "ProjectColumn")},
			"github_repository":               {Tok: makeResource(mainMod, "Repository")},
			"github_repository_collaborator":  {Tok: makeResource(mainMod, "RepositoryCollaborator")},
			"github_repository_deploy_key":    {Tok: makeResource(mainMod, "RepositoryDeployKey")},
			"github_repository_file":          {Tok: makeResource(mainMod, "RepositoryFile")},
			"github_repository_project":       {Tok: makeResource(mainMod, "RepositoryProject")},
			"github_repository_webhook":       {Tok: makeResource(mainMod, "RepositoryWebhook")},
			"github_team":                     {Tok: makeResource(mainMod, "Team")},
			"github_team_membership":          {Tok: makeResource(mainMod, "TeamMembership")},
			"github_team_repository":          {Tok: makeResource(mainMod, "TeamRepository")},
			"github_user_gpg_key":             {Tok: makeResource(mainMod, "UserGpgKey")},
			"github_user_invitation_accepter": {Tok: makeResource(mainMod, "UserInvitationAccepter")},
			"github_user_ssh_key":             {Tok: makeResource(mainMod, "UserSshKey")},
			"github_team_sync_group_mapping":  {Tok: makeResource(mainMod, "TeamSyncGroupMapping")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"github_actions_public_key":            {Tok: makeDataSource(mainMod, "getActionsPublicKey")},
			"github_branch":                        {Tok: makeDataSource(mainMod, "getBranch")},
			"github_collaborators":                 {Tok: makeDataSource(mainMod, "getCollaborators")},
			"github_ip_ranges":                     {Tok: makeDataSource(mainMod, "getIpRanges")},
			"github_membership":                    {Tok: makeDataSource(mainMod, "getMembership")},
			"github_release":                       {Tok: makeDataSource(mainMod, "getRelease")},
			"github_repositories":                  {Tok: makeDataSource(mainMod, "getRepositories")},
			"github_repository":                    {Tok: makeDataSource(mainMod, "getRepository")},
			"github_team":                          {Tok: makeDataSource(mainMod, "getTeam")},
			"github_user":                          {Tok: makeDataSource(mainMod, "getUser")},
			"github_organization_team_sync_groups": {Tok: makeDataSource(mainMod, "getOrganizationTeamSyncGroups")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^2.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^8.0.25", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
		},
		Python: &tfbridge.PythonInfo{
			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=2.9.0,<3.0.0",
			},
			UsesIOClasses: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi":                       "2.*",
				"System.Collections.Immutable": "1.6.0",
			},
		},
	}

	// For all resources with name properties, we will add an auto-name property.  Make sure to skip those that
	// already have a name mapping entry, since those may have custom overrides set above (e.g., for length).
	const nameProperty = "name"
	for resname, res := range prov.Resources {
		if schema := p.ResourcesMap[resname]; schema != nil {
			// Only apply auto-name to input properties (Optional || Required) named `name`
			if tfs, has := schema.Schema[nameProperty]; has && (tfs.Optional || tfs.Required) {
				if _, hasfield := res.Fields[nameProperty]; !hasfield {
					if res.Fields == nil {
						res.Fields = make(map[string]*tfbridge.SchemaInfo)
					}
					res.Fields[nameProperty] = tfbridge.AutoName(nameProperty, 255)
				}
			}
		}
	}

	return prov
}
