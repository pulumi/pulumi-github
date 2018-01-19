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
	issuesMod = "issues"
	orgsMod   = "orgs"
	reposMod  = "repos"
)

// githubMember manufactures a type token for the GitHub package and the given module and type.
func githubMember(mod string, mem string) tokens.ModuleMember {
	return tokens.ModuleMember(githubPkg + ":" + mod + ":" + mem)
}

// githubType manufactures a type token for the GitHub package and the given module and type.
func githubType(mod string, typ string) tokens.Type {
	return tokens.Type(githubMember(mod, typ))
}

// githubDataSource manufactures a standard resource token given a module and resource name.  It automatically uses the GitHub
// package and names the file by simply lower casing the data source's first character.
func githubDataSource(mod string, res string) tokens.ModuleMember {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return githubMember(mod+"/"+fn, res)
}

// githubResource manufactures a standard resource token given a module and resource name.  It automatically uses the GitHub
// package and names the file by simply lower casing the resource's first character.
func githubResource(mod string, res string) tokens.Type {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return githubType(mod+"/"+fn, res)
}

// Provider returns additional overlaid schema and metadata associated with the github package.
func Provider() tfbridge.ProviderInfo {
	p := github.Provider().(*schema.Provider)
	prov := tfbridge.ProviderInfo{
		P:      p,
		Name:   "github",
		Config: map[string]*tfbridge.SchemaInfo{},
		Resources: map[string]*tfbridge.ResourceInfo{
			"github_branch_protection": {Tok: githubResource(reposMod, "BranchProtection")},
			"github_issue_label":       {Tok: githubResource(issuesMod, "Label")},
			"github_membership": {
				Tok: githubResource(orgsMod, "Membership"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"role": {
						Type: githubType(orgsMod, "MembershipRole"),
					},
				},
			},
			"github_organization_webhook": {Tok: githubResource(orgsMod, "WebHook")},
			"github_repository_collaborator": {
				Tok: githubResource(reposMod, "Collaborator"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"permission": {
						Type: githubType(githubMod, "Permission"),
					},
				},
			},
			"github_repository_deploy_key": {Tok: githubResource(reposMod, "DeployKey")},
			"github_repository_webhook":    {Tok: githubResource(reposMod, "WebHook")},
			"github_repository":            {Tok: githubResource(reposMod, "Repository")},
			"github_team_membership": {
				Tok: githubResource(orgsMod, "TeamMembership"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"role": {
						Type: githubType(orgsMod, "TeamMembershipRole"),
					},
				},
			},
			"github_team_repository": {
				Tok: githubResource(orgsMod, "TeamRepository"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"permission": {
						Type: githubType(githubMod, "Permission"),
					},
				},
			},
			"github_team": {
				Tok: githubResource(orgsMod, "Team"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"privacy": {
						Type: githubType(orgsMod, "Privacy"),
					},
				},
			},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"github_team": {Tok: githubDataSource(orgsMod, "getTeam")},
			"github_user": {Tok: githubDataSource(githubMod, "getUser")},
		},
		Overlay: &tfbridge.OverlayInfo{
			Files: []string{
				"permission.ts",
			},
			Modules: map[string]*tfbridge.OverlayInfo{
				"orgs": {
					Files: []string{
						"privacy.ts",
						"membershipRole.ts",
						"teamMembershipRole.ts",
					},
				},
			},
			Dependencies: map[string]string{
				"@types/node": "^8.0.25", // so we can access strongly typed node definitions.
			},
		},
	}

	return prov
}
