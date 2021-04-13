// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

// Export members:
export * from "./actionsOrganizationSecret";
export * from "./actionsSecret";
export * from "./appInstallationRepository";
export * from "./branch";
export * from "./branchDefault";
export * from "./branchProtection";
export * from "./branchProtectionV3";
export * from "./getActionsPublicKey";
export * from "./getBranch";
export * from "./getCollaborators";
export * from "./getIpRanges";
export * from "./getMembership";
export * from "./getOrganization";
export * from "./getOrganizationTeamSyncGroups";
export * from "./getOrganizationTeams";
export * from "./getRelease";
export * from "./getRepositories";
export * from "./getRepository";
export * from "./getRepositoryMilestone";
export * from "./getTeam";
export * from "./getUser";
export * from "./issueLabel";
export * from "./membership";
export * from "./organizationBlock";
export * from "./organizationProject";
export * from "./organizationWebhook";
export * from "./projectCard";
export * from "./projectColumn";
export * from "./provider";
export * from "./repository";
export * from "./repositoryCollaborator";
export * from "./repositoryDeployKey";
export * from "./repositoryFile";
export * from "./repositoryMilestone";
export * from "./repositoryProject";
export * from "./repositoryWebhook";
export * from "./team";
export * from "./teamMembership";
export * from "./teamRepository";
export * from "./teamSyncGroupMapping";
export * from "./userGpgKey";
export * from "./userInvitationAccepter";
export * from "./userSshKey";

// Export sub-modules:
import * as config from "./config";
import * as types from "./types";

export {
    config,
    types,
};

// Import resources to register:
import { ActionsOrganizationSecret } from "./actionsOrganizationSecret";
import { ActionsSecret } from "./actionsSecret";
import { AppInstallationRepository } from "./appInstallationRepository";
import { Branch } from "./branch";
import { BranchDefault } from "./branchDefault";
import { BranchProtection } from "./branchProtection";
import { BranchProtectionV3 } from "./branchProtectionV3";
import { IssueLabel } from "./issueLabel";
import { Membership } from "./membership";
import { OrganizationBlock } from "./organizationBlock";
import { OrganizationProject } from "./organizationProject";
import { OrganizationWebhook } from "./organizationWebhook";
import { ProjectCard } from "./projectCard";
import { ProjectColumn } from "./projectColumn";
import { Repository } from "./repository";
import { RepositoryCollaborator } from "./repositoryCollaborator";
import { RepositoryDeployKey } from "./repositoryDeployKey";
import { RepositoryFile } from "./repositoryFile";
import { RepositoryMilestone } from "./repositoryMilestone";
import { RepositoryProject } from "./repositoryProject";
import { RepositoryWebhook } from "./repositoryWebhook";
import { Team } from "./team";
import { TeamMembership } from "./teamMembership";
import { TeamRepository } from "./teamRepository";
import { TeamSyncGroupMapping } from "./teamSyncGroupMapping";
import { UserGpgKey } from "./userGpgKey";
import { UserInvitationAccepter } from "./userInvitationAccepter";
import { UserSshKey } from "./userSshKey";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "github:index/actionsOrganizationSecret:ActionsOrganizationSecret":
                return new ActionsOrganizationSecret(name, <any>undefined, { urn })
            case "github:index/actionsSecret:ActionsSecret":
                return new ActionsSecret(name, <any>undefined, { urn })
            case "github:index/appInstallationRepository:AppInstallationRepository":
                return new AppInstallationRepository(name, <any>undefined, { urn })
            case "github:index/branch:Branch":
                return new Branch(name, <any>undefined, { urn })
            case "github:index/branchDefault:BranchDefault":
                return new BranchDefault(name, <any>undefined, { urn })
            case "github:index/branchProtection:BranchProtection":
                return new BranchProtection(name, <any>undefined, { urn })
            case "github:index/branchProtectionV3:BranchProtectionV3":
                return new BranchProtectionV3(name, <any>undefined, { urn })
            case "github:index/issueLabel:IssueLabel":
                return new IssueLabel(name, <any>undefined, { urn })
            case "github:index/membership:Membership":
                return new Membership(name, <any>undefined, { urn })
            case "github:index/organizationBlock:OrganizationBlock":
                return new OrganizationBlock(name, <any>undefined, { urn })
            case "github:index/organizationProject:OrganizationProject":
                return new OrganizationProject(name, <any>undefined, { urn })
            case "github:index/organizationWebhook:OrganizationWebhook":
                return new OrganizationWebhook(name, <any>undefined, { urn })
            case "github:index/projectCard:ProjectCard":
                return new ProjectCard(name, <any>undefined, { urn })
            case "github:index/projectColumn:ProjectColumn":
                return new ProjectColumn(name, <any>undefined, { urn })
            case "github:index/repository:Repository":
                return new Repository(name, <any>undefined, { urn })
            case "github:index/repositoryCollaborator:RepositoryCollaborator":
                return new RepositoryCollaborator(name, <any>undefined, { urn })
            case "github:index/repositoryDeployKey:RepositoryDeployKey":
                return new RepositoryDeployKey(name, <any>undefined, { urn })
            case "github:index/repositoryFile:RepositoryFile":
                return new RepositoryFile(name, <any>undefined, { urn })
            case "github:index/repositoryMilestone:RepositoryMilestone":
                return new RepositoryMilestone(name, <any>undefined, { urn })
            case "github:index/repositoryProject:RepositoryProject":
                return new RepositoryProject(name, <any>undefined, { urn })
            case "github:index/repositoryWebhook:RepositoryWebhook":
                return new RepositoryWebhook(name, <any>undefined, { urn })
            case "github:index/team:Team":
                return new Team(name, <any>undefined, { urn })
            case "github:index/teamMembership:TeamMembership":
                return new TeamMembership(name, <any>undefined, { urn })
            case "github:index/teamRepository:TeamRepository":
                return new TeamRepository(name, <any>undefined, { urn })
            case "github:index/teamSyncGroupMapping:TeamSyncGroupMapping":
                return new TeamSyncGroupMapping(name, <any>undefined, { urn })
            case "github:index/userGpgKey:UserGpgKey":
                return new UserGpgKey(name, <any>undefined, { urn })
            case "github:index/userInvitationAccepter:UserInvitationAccepter":
                return new UserInvitationAccepter(name, <any>undefined, { urn })
            case "github:index/userSshKey:UserSshKey":
                return new UserSshKey(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("github", "index/actionsOrganizationSecret", _module)
pulumi.runtime.registerResourceModule("github", "index/actionsSecret", _module)
pulumi.runtime.registerResourceModule("github", "index/appInstallationRepository", _module)
pulumi.runtime.registerResourceModule("github", "index/branch", _module)
pulumi.runtime.registerResourceModule("github", "index/branchDefault", _module)
pulumi.runtime.registerResourceModule("github", "index/branchProtection", _module)
pulumi.runtime.registerResourceModule("github", "index/branchProtectionV3", _module)
pulumi.runtime.registerResourceModule("github", "index/issueLabel", _module)
pulumi.runtime.registerResourceModule("github", "index/membership", _module)
pulumi.runtime.registerResourceModule("github", "index/organizationBlock", _module)
pulumi.runtime.registerResourceModule("github", "index/organizationProject", _module)
pulumi.runtime.registerResourceModule("github", "index/organizationWebhook", _module)
pulumi.runtime.registerResourceModule("github", "index/projectCard", _module)
pulumi.runtime.registerResourceModule("github", "index/projectColumn", _module)
pulumi.runtime.registerResourceModule("github", "index/repository", _module)
pulumi.runtime.registerResourceModule("github", "index/repositoryCollaborator", _module)
pulumi.runtime.registerResourceModule("github", "index/repositoryDeployKey", _module)
pulumi.runtime.registerResourceModule("github", "index/repositoryFile", _module)
pulumi.runtime.registerResourceModule("github", "index/repositoryMilestone", _module)
pulumi.runtime.registerResourceModule("github", "index/repositoryProject", _module)
pulumi.runtime.registerResourceModule("github", "index/repositoryWebhook", _module)
pulumi.runtime.registerResourceModule("github", "index/team", _module)
pulumi.runtime.registerResourceModule("github", "index/teamMembership", _module)
pulumi.runtime.registerResourceModule("github", "index/teamRepository", _module)
pulumi.runtime.registerResourceModule("github", "index/teamSyncGroupMapping", _module)
pulumi.runtime.registerResourceModule("github", "index/userGpgKey", _module)
pulumi.runtime.registerResourceModule("github", "index/userInvitationAccepter", _module)
pulumi.runtime.registerResourceModule("github", "index/userSshKey", _module)

import { Provider } from "./provider";

pulumi.runtime.registerResourcePackage("github", {
    version: utilities.getVersion(),
    constructProvider: (name: string, type: string, urn: string): pulumi.ProviderResource => {
        if (type !== "pulumi:providers:github") {
            throw new Error(`unknown provider type ${type}`);
        }
        return new Provider(name, <any>undefined, { urn });
    },
});
