// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "github:index/actionsOrganizationSecret:ActionsOrganizationSecret":
		r, err = NewActionsOrganizationSecret(ctx, name, nil, pulumi.URN_(urn))
	case "github:index/actionsSecret:ActionsSecret":
		r, err = NewActionsSecret(ctx, name, nil, pulumi.URN_(urn))
	case "github:index/branch:Branch":
		r, err = NewBranch(ctx, name, nil, pulumi.URN_(urn))
	case "github:index/branchDefault:BranchDefault":
		r, err = NewBranchDefault(ctx, name, nil, pulumi.URN_(urn))
	case "github:index/branchProtection:BranchProtection":
		r, err = NewBranchProtection(ctx, name, nil, pulumi.URN_(urn))
	case "github:index/branchProtectionV3:BranchProtectionV3":
		r, err = NewBranchProtectionV3(ctx, name, nil, pulumi.URN_(urn))
	case "github:index/issueLabel:IssueLabel":
		r, err = NewIssueLabel(ctx, name, nil, pulumi.URN_(urn))
	case "github:index/membership:Membership":
		r, err = NewMembership(ctx, name, nil, pulumi.URN_(urn))
	case "github:index/organizationBlock:OrganizationBlock":
		r, err = NewOrganizationBlock(ctx, name, nil, pulumi.URN_(urn))
	case "github:index/organizationProject:OrganizationProject":
		r, err = NewOrganizationProject(ctx, name, nil, pulumi.URN_(urn))
	case "github:index/organizationWebhook:OrganizationWebhook":
		r, err = NewOrganizationWebhook(ctx, name, nil, pulumi.URN_(urn))
	case "github:index/projectCard:ProjectCard":
		r, err = NewProjectCard(ctx, name, nil, pulumi.URN_(urn))
	case "github:index/projectColumn:ProjectColumn":
		r, err = NewProjectColumn(ctx, name, nil, pulumi.URN_(urn))
	case "github:index/repository:Repository":
		r, err = NewRepository(ctx, name, nil, pulumi.URN_(urn))
	case "github:index/repositoryCollaborator:RepositoryCollaborator":
		r, err = NewRepositoryCollaborator(ctx, name, nil, pulumi.URN_(urn))
	case "github:index/repositoryDeployKey:RepositoryDeployKey":
		r, err = NewRepositoryDeployKey(ctx, name, nil, pulumi.URN_(urn))
	case "github:index/repositoryFile:RepositoryFile":
		r, err = NewRepositoryFile(ctx, name, nil, pulumi.URN_(urn))
	case "github:index/repositoryMilestone:RepositoryMilestone":
		r, err = NewRepositoryMilestone(ctx, name, nil, pulumi.URN_(urn))
	case "github:index/repositoryProject:RepositoryProject":
		r, err = NewRepositoryProject(ctx, name, nil, pulumi.URN_(urn))
	case "github:index/repositoryWebhook:RepositoryWebhook":
		r, err = NewRepositoryWebhook(ctx, name, nil, pulumi.URN_(urn))
	case "github:index/team:Team":
		r, err = NewTeam(ctx, name, nil, pulumi.URN_(urn))
	case "github:index/teamMembership:TeamMembership":
		r, err = NewTeamMembership(ctx, name, nil, pulumi.URN_(urn))
	case "github:index/teamRepository:TeamRepository":
		r, err = NewTeamRepository(ctx, name, nil, pulumi.URN_(urn))
	case "github:index/teamSyncGroupMapping:TeamSyncGroupMapping":
		r, err = NewTeamSyncGroupMapping(ctx, name, nil, pulumi.URN_(urn))
	case "github:index/userGpgKey:UserGpgKey":
		r, err = NewUserGpgKey(ctx, name, nil, pulumi.URN_(urn))
	case "github:index/userInvitationAccepter:UserInvitationAccepter":
		r, err = NewUserInvitationAccepter(ctx, name, nil, pulumi.URN_(urn))
	case "github:index/userSshKey:UserSshKey":
		r, err = NewUserSshKey(ctx, name, nil, pulumi.URN_(urn))
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	return
}

type pkg struct {
	version semver.Version
}

func (p *pkg) Version() semver.Version {
	return p.version
}

func (p *pkg) ConstructProvider(ctx *pulumi.Context, name, typ, urn string) (pulumi.ProviderResource, error) {
	if typ != "pulumi:providers:github" {
		return nil, fmt.Errorf("unknown provider type: %s", typ)
	}

	return NewProvider(ctx, name, nil, pulumi.URN_(urn))
}

func init() {
	version, err := PkgVersion()
	if err != nil {
		fmt.Println("failed to determine package version. defaulting to v1: %v", err)
	}
	pulumi.RegisterResourceModule(
		"github",
		"index/actionsOrganizationSecret",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"github",
		"index/actionsSecret",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"github",
		"index/branch",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"github",
		"index/branchDefault",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"github",
		"index/branchProtection",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"github",
		"index/branchProtectionV3",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"github",
		"index/issueLabel",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"github",
		"index/membership",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"github",
		"index/organizationBlock",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"github",
		"index/organizationProject",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"github",
		"index/organizationWebhook",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"github",
		"index/projectCard",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"github",
		"index/projectColumn",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"github",
		"index/repository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"github",
		"index/repositoryCollaborator",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"github",
		"index/repositoryDeployKey",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"github",
		"index/repositoryFile",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"github",
		"index/repositoryMilestone",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"github",
		"index/repositoryProject",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"github",
		"index/repositoryWebhook",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"github",
		"index/team",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"github",
		"index/teamMembership",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"github",
		"index/teamRepository",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"github",
		"index/teamSyncGroupMapping",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"github",
		"index/userGpgKey",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"github",
		"index/userInvitationAccepter",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"github",
		"index/userSshKey",
		&module{version},
	)
	pulumi.RegisterResourcePackage(
		"github",
		&pkg{version},
	)
}
