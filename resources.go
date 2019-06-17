// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

package github

import (
	"unicode"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pulumi/pulumi-terraform/pkg/tfbridge"
	"github.com/pulumi/pulumi/pkg/tokens"
	"github.com/terraform-providers/terraform-provider-github/github"
)

// all of the GitHub token components used below.
const (
	// packages:
	githubPkg = "github"
	// modules:
	githubMod = "index"
	reposMod  = "repos"
	orgsMod   = "orgs"
	teamsMod  = "teams"
	usersMod  = "users"
)

// githubMember manufactures a type token for the GitHub package and the given module and type.
func githubMember(mod string, mem string) tokens.ModuleMember {
	return tokens.ModuleMember(githubPkg + ":" + mod + ":" + mem)
}

// githubType manufactures a type token for the GitHub package and the given module and type.
func githubType(mod string, typ string) tokens.Type {
	return tokens.Type(githubMember(mod, typ))
}

// githubDataSource manufactures a standard resource token given a module and resource name.
// It automatically uses the GitHub
// package and names the file by simply lower casing the data source's first character.
func githubDataSource(mod string, res string) tokens.ModuleMember {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return githubMember(mod+"/"+fn, res)
}

// githubResource manufactures a standard resource token given a module and resource name.
// It automatically uses the GitHub
// package and names the file by simply lower casing the resource's first character.
func githubResource(mod string, res string) tokens.Type {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return githubType(mod+"/"+fn, res)
}

// Provider returns additional overlaid schema and metadata associated with the github package.
// Modules and resource names are based on the GitHub v3 API (https://developer.github.com/v3/).
func Provider() tfbridge.ProviderInfo {
	p := github.Provider().(*schema.Provider)
	prov := tfbridge.ProviderInfo{
		P:           p,
		Name:        "github",
		Description: "A Pulumi package for creating and managing GitHub resources.",
		Keywords:    []string{"pulumi", "github"},
		Homepage:    "https://pulumi.io/github",
		Repository:  "https://github.com/pulumi/pulumi-github",
		Config:      map[string]*tfbridge.SchemaInfo{},
		Resources: map[string]*tfbridge.ResourceInfo{
			// Repository based Resources
			"github_branch_protection": {Tok: githubResource(reposMod, "BranchProtection")},
			"github_issue_label":       {Tok: githubResource(reposMod, "Label")},
			"github_repository_collaborator": {
				Tok: githubResource(reposMod, "Collaborator"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"permission": {
						Type: githubType(githubMod, "Permission"),
					},
				},
			},
			"github_repository_deploy_key": {Tok: githubResource(reposMod, "DeployKey")},
			"github_repository_webhook": {
				Tok: githubResource(reposMod, "Webhook"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"configuration": {
						Type: githubType(githubMod, "WebhookConfiguration"),
					},
				},
			},
			"github_repository":         {Tok: githubResource(reposMod, "Repository")},
			"github_repository_project": {Tok: githubResource(reposMod, "Project")},
			// Organization based resources
			"github_organization_block":   {Tok: githubResource(orgsMod, "Block")},
			"github_organization_project": {Tok: githubResource(orgsMod, "Project")},
			"github_organization_webhook": {
				Tok: githubResource(orgsMod, "Webhook"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"configuration": {
						Type: githubType(githubMod, "WebhookConfiguration"),
					},
				},
			},
			"github_membership": {
				Tok: githubResource(orgsMod, "Membership"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"role": {
						Type: githubType(orgsMod, "MembershipRole"),
					},
				},
			},
			"github_project_column": {Tok: githubResource(orgsMod, "ProjectColumn")},
			// Teams based resources
			"github_team_membership": {
				Tok: githubResource(teamsMod, "Membership"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"role": {
						Type: githubType(teamsMod, "MembershipRole"),
					},
				},
			},
			"github_team_repository": {
				Tok: githubResource(teamsMod, "Repository"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"permission": {
						Type: githubType(githubMod, "Permission"),
					},
				},
			},
			"github_team": {
				Tok: githubResource(teamsMod, "Team"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"privacy": {
						Type: githubType(teamsMod, "Privacy"),
					},
				},
			},
			// User based resources
			"github_user_gpg_key":             {Tok: githubResource(usersMod, "GpgKey")},
			"github_user_ssh_key":             {Tok: githubResource(usersMod, "SshKey")},
			"github_user_invitation_accepter": {Tok: githubResource(usersMod, "InvitationAcceptor")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"github_team":         {Tok: githubDataSource(teamsMod, "getTeam")},
			"github_user":         {Tok: githubDataSource(usersMod, "getUser")},
			"github_ip_ranges":    {Tok: githubDataSource(githubMod, "getIpRanges")},
			"github_repository":   {Tok: githubDataSource(reposMod, "getRepository")},
			"github_repositories": {Tok: githubDataSource(reposMod, "getRepositories")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			Dependencies: map[string]string{
				"@pulumi/pulumi":    "^0.17.1",
				"builtin-modules":   "3.0.0",
				"read-package-tree": "^5.2.1",
				"resolve":           "^1.7.1",
			},
			DevDependencies: map[string]string{
				"@types/node": "^8.0.25", // so we can access strongly typed node definitions.
			},
			Overlay: &tfbridge.OverlayInfo{
				DestFiles: []string{
					"permission.ts",
					"webhookConfiguration.ts",
				},
				Modules: map[string]*tfbridge.OverlayInfo{
					"orgs": {
						DestFiles: []string{
							"membershipRole.ts",
						},
					},
					"teams": {
						DestFiles: []string{
							"membershipRole.ts",
							"privacy.ts",
						},
					},
				},
			},
		},
		Python: &tfbridge.PythonInfo{
			Requires: map[string]string{
				"pulumi": ">=0.17.1,<0.18.0",
			},
		},
	}

	// For all resources with name properties, we will add an auto-name property.  Make sure to skip those that
	// already have a name mapping entry, since those may have custom overrides set above (e.g., for length).
	const nameField = "name"
	for resname, res := range prov.Resources {
		if schema := p.ResourcesMap[resname]; schema != nil {
			// Only apply auto-name to input properties (Optional || Required) named `name`
			if tfs, has := schema.Schema[nameField]; has && (tfs.Optional || tfs.Required) {
				if _, hasfield := res.Fields[nameField]; !hasfield {
					if res.Fields == nil {
						res.Fields = make(map[string]*tfbridge.SchemaInfo)
					}
					res.Fields[nameField] = tfbridge.AutoName(nameField, 255)
				}
			}
		}
	}

	return prov
}
