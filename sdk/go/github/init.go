// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "github:index/actionsEnvironmentSecret:ActionsEnvironmentSecret":
		r = &ActionsEnvironmentSecret{}
	case "github:index/actionsOrganizationPermissions:ActionsOrganizationPermissions":
		r = &ActionsOrganizationPermissions{}
	case "github:index/actionsOrganizationSecret:ActionsOrganizationSecret":
		r = &ActionsOrganizationSecret{}
	case "github:index/actionsOrganizationSecretRepositories:ActionsOrganizationSecretRepositories":
		r = &ActionsOrganizationSecretRepositories{}
	case "github:index/actionsRunnerGroup:ActionsRunnerGroup":
		r = &ActionsRunnerGroup{}
	case "github:index/actionsSecret:ActionsSecret":
		r = &ActionsSecret{}
	case "github:index/appInstallationRepository:AppInstallationRepository":
		r = &AppInstallationRepository{}
	case "github:index/branch:Branch":
		r = &Branch{}
	case "github:index/branchDefault:BranchDefault":
		r = &BranchDefault{}
	case "github:index/branchProtection:BranchProtection":
		r = &BranchProtection{}
	case "github:index/branchProtectionV3:BranchProtectionV3":
		r = &BranchProtectionV3{}
	case "github:index/issue:Issue":
		r = &Issue{}
	case "github:index/issueLabel:IssueLabel":
		r = &IssueLabel{}
	case "github:index/membership:Membership":
		r = &Membership{}
	case "github:index/organizationBlock:OrganizationBlock":
		r = &OrganizationBlock{}
	case "github:index/organizationProject:OrganizationProject":
		r = &OrganizationProject{}
	case "github:index/organizationWebhook:OrganizationWebhook":
		r = &OrganizationWebhook{}
	case "github:index/projectCard:ProjectCard":
		r = &ProjectCard{}
	case "github:index/projectColumn:ProjectColumn":
		r = &ProjectColumn{}
	case "github:index/repository:Repository":
		r = &Repository{}
	case "github:index/repositoryAutolinkReference:RepositoryAutolinkReference":
		r = &RepositoryAutolinkReference{}
	case "github:index/repositoryCollaborator:RepositoryCollaborator":
		r = &RepositoryCollaborator{}
	case "github:index/repositoryDeployKey:RepositoryDeployKey":
		r = &RepositoryDeployKey{}
	case "github:index/repositoryEnvironment:RepositoryEnvironment":
		r = &RepositoryEnvironment{}
	case "github:index/repositoryFile:RepositoryFile":
		r = &RepositoryFile{}
	case "github:index/repositoryMilestone:RepositoryMilestone":
		r = &RepositoryMilestone{}
	case "github:index/repositoryProject:RepositoryProject":
		r = &RepositoryProject{}
	case "github:index/repositoryPullRequest:RepositoryPullRequest":
		r = &RepositoryPullRequest{}
	case "github:index/repositoryWebhook:RepositoryWebhook":
		r = &RepositoryWebhook{}
	case "github:index/team:Team":
		r = &Team{}
	case "github:index/teamMembers:TeamMembers":
		r = &TeamMembers{}
	case "github:index/teamMembership:TeamMembership":
		r = &TeamMembership{}
	case "github:index/teamRepository:TeamRepository":
		r = &TeamRepository{}
	case "github:index/teamSyncGroupMapping:TeamSyncGroupMapping":
		r = &TeamSyncGroupMapping{}
	case "github:index/userGpgKey:UserGpgKey":
		r = &UserGpgKey{}
	case "github:index/userInvitationAccepter:UserInvitationAccepter":
		r = &UserInvitationAccepter{}
	case "github:index/userSshKey:UserSshKey":
		r = &UserSshKey{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
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

	r := &Provider{}
	err := ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return r, err
}

func init() {
	version, _ := PkgVersion()
	pulumi.RegisterResourceModule(
		"github",
		"index/actionsEnvironmentSecret",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"github",
		"index/actionsOrganizationPermissions",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"github",
		"index/actionsOrganizationSecret",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"github",
		"index/actionsOrganizationSecretRepositories",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"github",
		"index/actionsRunnerGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"github",
		"index/actionsSecret",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"github",
		"index/appInstallationRepository",
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
		"index/issue",
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
		"index/repositoryAutolinkReference",
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
		"index/repositoryEnvironment",
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
		"index/repositoryPullRequest",
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
		"index/teamMembers",
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
