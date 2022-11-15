# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .actions_environment_secret import *
from .actions_organization_permissions import *
from .actions_organization_secret import *
from .actions_organization_secret_repositories import *
from .actions_runner_group import *
from .actions_secret import *
from .app_installation_repository import *
from .branch import *
from .branch_default import *
from .branch_protection import *
from .branch_protection_v3 import *
from .dependabot_organization_secret import *
from .dependabot_organization_secret_repositories import *
from .dependabot_secret import *
from .emu_group_mapping import *
from .get_actions_organization_secrets import *
from .get_actions_public_key import *
from .get_actions_secrets import *
from .get_branch import *
from .get_collaborators import *
from .get_dependabot_organization_secrets import *
from .get_dependabot_public_key import *
from .get_dependabot_secrets import *
from .get_external_groups import *
from .get_ip_ranges import *
from .get_membership import *
from .get_organization import *
from .get_organization_team_sync_groups import *
from .get_organization_teams import *
from .get_ref import *
from .get_release import *
from .get_repositories import *
from .get_repository import *
from .get_repository_branches import *
from .get_repository_file import *
from .get_repository_milestone import *
from .get_repository_pull_request import *
from .get_repository_pull_requests import *
from .get_repository_teams import *
from .get_team import *
from .get_tree import *
from .get_user import *
from .get_users import *
from .issue import *
from .issue_label import *
from .membership import *
from .organization_block import *
from .organization_ip_allow_list import *
from .organization_project import *
from .organization_webhook import *
from .project_card import *
from .project_column import *
from .provider import *
from .repository import *
from .repository_autolink_reference import *
from .repository_collaborator import *
from .repository_deploy_key import *
from .repository_environment import *
from .repository_file import *
from .repository_milestone import *
from .repository_project import *
from .repository_pull_request import *
from .repository_tag_protection import *
from .repository_webhook import *
from .team import *
from .team_members import *
from .team_membership import *
from .team_repository import *
from .team_sync_group_mapping import *
from .user_gpg_key import *
from .user_invitation_accepter import *
from .user_ssh_key import *
from ._inputs import *
from . import outputs

# Make subpackages available:
if typing.TYPE_CHECKING:
    import pulumi_github.config as __config
    config = __config
else:
    config = _utilities.lazy_import('pulumi_github.config')

_utilities.register(
    resource_modules="""
[
 {
  "pkg": "github",
  "mod": "index/actionsEnvironmentSecret",
  "fqn": "pulumi_github",
  "classes": {
   "github:index/actionsEnvironmentSecret:ActionsEnvironmentSecret": "ActionsEnvironmentSecret"
  }
 },
 {
  "pkg": "github",
  "mod": "index/actionsOrganizationPermissions",
  "fqn": "pulumi_github",
  "classes": {
   "github:index/actionsOrganizationPermissions:ActionsOrganizationPermissions": "ActionsOrganizationPermissions"
  }
 },
 {
  "pkg": "github",
  "mod": "index/actionsOrganizationSecret",
  "fqn": "pulumi_github",
  "classes": {
   "github:index/actionsOrganizationSecret:ActionsOrganizationSecret": "ActionsOrganizationSecret"
  }
 },
 {
  "pkg": "github",
  "mod": "index/actionsOrganizationSecretRepositories",
  "fqn": "pulumi_github",
  "classes": {
   "github:index/actionsOrganizationSecretRepositories:ActionsOrganizationSecretRepositories": "ActionsOrganizationSecretRepositories"
  }
 },
 {
  "pkg": "github",
  "mod": "index/actionsRunnerGroup",
  "fqn": "pulumi_github",
  "classes": {
   "github:index/actionsRunnerGroup:ActionsRunnerGroup": "ActionsRunnerGroup"
  }
 },
 {
  "pkg": "github",
  "mod": "index/actionsSecret",
  "fqn": "pulumi_github",
  "classes": {
   "github:index/actionsSecret:ActionsSecret": "ActionsSecret"
  }
 },
 {
  "pkg": "github",
  "mod": "index/appInstallationRepository",
  "fqn": "pulumi_github",
  "classes": {
   "github:index/appInstallationRepository:AppInstallationRepository": "AppInstallationRepository"
  }
 },
 {
  "pkg": "github",
  "mod": "index/branch",
  "fqn": "pulumi_github",
  "classes": {
   "github:index/branch:Branch": "Branch"
  }
 },
 {
  "pkg": "github",
  "mod": "index/branchDefault",
  "fqn": "pulumi_github",
  "classes": {
   "github:index/branchDefault:BranchDefault": "BranchDefault"
  }
 },
 {
  "pkg": "github",
  "mod": "index/branchProtection",
  "fqn": "pulumi_github",
  "classes": {
   "github:index/branchProtection:BranchProtection": "BranchProtection"
  }
 },
 {
  "pkg": "github",
  "mod": "index/branchProtectionV3",
  "fqn": "pulumi_github",
  "classes": {
   "github:index/branchProtectionV3:BranchProtectionV3": "BranchProtectionV3"
  }
 },
 {
  "pkg": "github",
  "mod": "index/dependabotOrganizationSecret",
  "fqn": "pulumi_github",
  "classes": {
   "github:index/dependabotOrganizationSecret:DependabotOrganizationSecret": "DependabotOrganizationSecret"
  }
 },
 {
  "pkg": "github",
  "mod": "index/dependabotOrganizationSecretRepositories",
  "fqn": "pulumi_github",
  "classes": {
   "github:index/dependabotOrganizationSecretRepositories:DependabotOrganizationSecretRepositories": "DependabotOrganizationSecretRepositories"
  }
 },
 {
  "pkg": "github",
  "mod": "index/dependabotSecret",
  "fqn": "pulumi_github",
  "classes": {
   "github:index/dependabotSecret:DependabotSecret": "DependabotSecret"
  }
 },
 {
  "pkg": "github",
  "mod": "index/emuGroupMapping",
  "fqn": "pulumi_github",
  "classes": {
   "github:index/emuGroupMapping:EmuGroupMapping": "EmuGroupMapping"
  }
 },
 {
  "pkg": "github",
  "mod": "index/issue",
  "fqn": "pulumi_github",
  "classes": {
   "github:index/issue:Issue": "Issue"
  }
 },
 {
  "pkg": "github",
  "mod": "index/issueLabel",
  "fqn": "pulumi_github",
  "classes": {
   "github:index/issueLabel:IssueLabel": "IssueLabel"
  }
 },
 {
  "pkg": "github",
  "mod": "index/membership",
  "fqn": "pulumi_github",
  "classes": {
   "github:index/membership:Membership": "Membership"
  }
 },
 {
  "pkg": "github",
  "mod": "index/organizationBlock",
  "fqn": "pulumi_github",
  "classes": {
   "github:index/organizationBlock:OrganizationBlock": "OrganizationBlock"
  }
 },
 {
  "pkg": "github",
  "mod": "index/organizationProject",
  "fqn": "pulumi_github",
  "classes": {
   "github:index/organizationProject:OrganizationProject": "OrganizationProject"
  }
 },
 {
  "pkg": "github",
  "mod": "index/organizationWebhook",
  "fqn": "pulumi_github",
  "classes": {
   "github:index/organizationWebhook:OrganizationWebhook": "OrganizationWebhook"
  }
 },
 {
  "pkg": "github",
  "mod": "index/projectCard",
  "fqn": "pulumi_github",
  "classes": {
   "github:index/projectCard:ProjectCard": "ProjectCard"
  }
 },
 {
  "pkg": "github",
  "mod": "index/projectColumn",
  "fqn": "pulumi_github",
  "classes": {
   "github:index/projectColumn:ProjectColumn": "ProjectColumn"
  }
 },
 {
  "pkg": "github",
  "mod": "index/repository",
  "fqn": "pulumi_github",
  "classes": {
   "github:index/repository:Repository": "Repository"
  }
 },
 {
  "pkg": "github",
  "mod": "index/repositoryAutolinkReference",
  "fqn": "pulumi_github",
  "classes": {
   "github:index/repositoryAutolinkReference:RepositoryAutolinkReference": "RepositoryAutolinkReference"
  }
 },
 {
  "pkg": "github",
  "mod": "index/repositoryCollaborator",
  "fqn": "pulumi_github",
  "classes": {
   "github:index/repositoryCollaborator:RepositoryCollaborator": "RepositoryCollaborator"
  }
 },
 {
  "pkg": "github",
  "mod": "index/repositoryDeployKey",
  "fqn": "pulumi_github",
  "classes": {
   "github:index/repositoryDeployKey:RepositoryDeployKey": "RepositoryDeployKey"
  }
 },
 {
  "pkg": "github",
  "mod": "index/repositoryEnvironment",
  "fqn": "pulumi_github",
  "classes": {
   "github:index/repositoryEnvironment:RepositoryEnvironment": "RepositoryEnvironment"
  }
 },
 {
  "pkg": "github",
  "mod": "index/repositoryFile",
  "fqn": "pulumi_github",
  "classes": {
   "github:index/repositoryFile:RepositoryFile": "RepositoryFile"
  }
 },
 {
  "pkg": "github",
  "mod": "index/repositoryMilestone",
  "fqn": "pulumi_github",
  "classes": {
   "github:index/repositoryMilestone:RepositoryMilestone": "RepositoryMilestone"
  }
 },
 {
  "pkg": "github",
  "mod": "index/repositoryProject",
  "fqn": "pulumi_github",
  "classes": {
   "github:index/repositoryProject:RepositoryProject": "RepositoryProject"
  }
 },
 {
  "pkg": "github",
  "mod": "index/repositoryPullRequest",
  "fqn": "pulumi_github",
  "classes": {
   "github:index/repositoryPullRequest:RepositoryPullRequest": "RepositoryPullRequest"
  }
 },
 {
  "pkg": "github",
  "mod": "index/repositoryTagProtection",
  "fqn": "pulumi_github",
  "classes": {
   "github:index/repositoryTagProtection:RepositoryTagProtection": "RepositoryTagProtection"
  }
 },
 {
  "pkg": "github",
  "mod": "index/repositoryWebhook",
  "fqn": "pulumi_github",
  "classes": {
   "github:index/repositoryWebhook:RepositoryWebhook": "RepositoryWebhook"
  }
 },
 {
  "pkg": "github",
  "mod": "index/team",
  "fqn": "pulumi_github",
  "classes": {
   "github:index/team:Team": "Team"
  }
 },
 {
  "pkg": "github",
  "mod": "index/teamMembers",
  "fqn": "pulumi_github",
  "classes": {
   "github:index/teamMembers:TeamMembers": "TeamMembers"
  }
 },
 {
  "pkg": "github",
  "mod": "index/teamMembership",
  "fqn": "pulumi_github",
  "classes": {
   "github:index/teamMembership:TeamMembership": "TeamMembership"
  }
 },
 {
  "pkg": "github",
  "mod": "index/teamRepository",
  "fqn": "pulumi_github",
  "classes": {
   "github:index/teamRepository:TeamRepository": "TeamRepository"
  }
 },
 {
  "pkg": "github",
  "mod": "index/teamSyncGroupMapping",
  "fqn": "pulumi_github",
  "classes": {
   "github:index/teamSyncGroupMapping:TeamSyncGroupMapping": "TeamSyncGroupMapping"
  }
 },
 {
  "pkg": "github",
  "mod": "index/userGpgKey",
  "fqn": "pulumi_github",
  "classes": {
   "github:index/userGpgKey:UserGpgKey": "UserGpgKey"
  }
 },
 {
  "pkg": "github",
  "mod": "index/userInvitationAccepter",
  "fqn": "pulumi_github",
  "classes": {
   "github:index/userInvitationAccepter:UserInvitationAccepter": "UserInvitationAccepter"
  }
 },
 {
  "pkg": "github",
  "mod": "index/userSshKey",
  "fqn": "pulumi_github",
  "classes": {
   "github:index/userSshKey:UserSshKey": "UserSshKey"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "github",
  "token": "pulumi:providers:github",
  "fqn": "pulumi_github",
  "class": "Provider"
 }
]
"""
)
