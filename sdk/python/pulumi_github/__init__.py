# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .actions_environment_secret import *
from .actions_environment_variable import *
from .actions_organization_oidc_subject_claim_customization_template import *
from .actions_organization_permissions import *
from .actions_organization_secret import *
from .actions_organization_secret_repositories import *
from .actions_organization_variable import *
from .actions_repository_access_level import *
from .actions_repository_oidc_subject_claim_customization_template import *
from .actions_repository_permissions import *
from .actions_runner_group import *
from .actions_secret import *
from .actions_variable import *
from .app_installation_repositories import *
from .app_installation_repository import *
from .branch import *
from .branch_default import *
from .branch_protection import *
from .branch_protection_v3 import *
from .codespaces_organization_secret import *
from .codespaces_organization_secret_repositories import *
from .codespaces_secret import *
from .codespaces_user_secret import *
from .dependabot_organization_secret import *
from .dependabot_organization_secret_repositories import *
from .dependabot_secret import *
from .emu_group_mapping import *
from .enterprise_actions_permissions import *
from .enterprise_actions_runner_group import *
from .enterprise_organization import *
from .get_actions_environment_secrets import *
from .get_actions_environment_variables import *
from .get_actions_organization_oidc_subject_claim_customization_template import *
from .get_actions_organization_public_key import *
from .get_actions_organization_registration_token import *
from .get_actions_organization_secrets import *
from .get_actions_organization_variables import *
from .get_actions_public_key import *
from .get_actions_registration_token import *
from .get_actions_repository_oidc_subject_claim_customization_template import *
from .get_actions_secrets import *
from .get_actions_variables import *
from .get_app_token import *
from .get_branch import *
from .get_branch_protection_rules import *
from .get_codespaces_organization_public_key import *
from .get_codespaces_organization_secrets import *
from .get_codespaces_public_key import *
from .get_codespaces_secrets import *
from .get_codespaces_user_public_key import *
from .get_codespaces_user_secrets import *
from .get_collaborators import *
from .get_dependabot_organization_public_key import *
from .get_dependabot_organization_secrets import *
from .get_dependabot_public_key import *
from .get_dependabot_secrets import *
from .get_enterprise import *
from .get_external_groups import *
from .get_github_app import *
from .get_ip_ranges import *
from .get_issue_labels import *
from .get_membership import *
from .get_organization import *
from .get_organization_custom_role import *
from .get_organization_external_identities import *
from .get_organization_ip_allow_list import *
from .get_organization_team_sync_groups import *
from .get_organization_teams import *
from .get_organization_webhooks import *
from .get_ref import *
from .get_release import *
from .get_repositories import *
from .get_repository import *
from .get_repository_autolink_references import *
from .get_repository_branches import *
from .get_repository_deploy_keys import *
from .get_repository_deployment_branch_policies import *
from .get_repository_environments import *
from .get_repository_file import *
from .get_repository_milestone import *
from .get_repository_pull_request import *
from .get_repository_pull_requests import *
from .get_repository_teams import *
from .get_repository_webhooks import *
from .get_rest_api import *
from .get_ssh_keys import *
from .get_team import *
from .get_tree import *
from .get_user import *
from .get_user_external_identity import *
from .get_users import *
from .issue import *
from .issue_label import *
from .issue_labels import *
from .membership import *
from .organization_block import *
from .organization_custom_role import *
from .organization_project import *
from .organization_ruleset import *
from .organization_security_manager import *
from .organization_settings import *
from .organization_webhook import *
from .project_card import *
from .project_column import *
from .provider import *
from .release import *
from .repository import *
from .repository_autolink_reference import *
from .repository_collaborator import *
from .repository_collaborators import *
from .repository_dependabot_security_updates import *
from .repository_deploy_key import *
from .repository_deployment_branch_policy import *
from .repository_environment import *
from .repository_environment_deployment_policy import *
from .repository_file import *
from .repository_milestone import *
from .repository_project import *
from .repository_pull_request import *
from .repository_ruleset import *
from .repository_tag_protection import *
from .repository_topics import *
from .repository_webhook import *
from .team import *
from .team_members import *
from .team_membership import *
from .team_repository import *
from .team_settings import *
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
  "mod": "index/actionsEnvironmentVariable",
  "fqn": "pulumi_github",
  "classes": {
   "github:index/actionsEnvironmentVariable:ActionsEnvironmentVariable": "ActionsEnvironmentVariable"
  }
 },
 {
  "pkg": "github",
  "mod": "index/actionsOrganizationOidcSubjectClaimCustomizationTemplate",
  "fqn": "pulumi_github",
  "classes": {
   "github:index/actionsOrganizationOidcSubjectClaimCustomizationTemplate:ActionsOrganizationOidcSubjectClaimCustomizationTemplate": "ActionsOrganizationOidcSubjectClaimCustomizationTemplate"
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
  "mod": "index/actionsOrganizationVariable",
  "fqn": "pulumi_github",
  "classes": {
   "github:index/actionsOrganizationVariable:ActionsOrganizationVariable": "ActionsOrganizationVariable"
  }
 },
 {
  "pkg": "github",
  "mod": "index/actionsRepositoryAccessLevel",
  "fqn": "pulumi_github",
  "classes": {
   "github:index/actionsRepositoryAccessLevel:ActionsRepositoryAccessLevel": "ActionsRepositoryAccessLevel"
  }
 },
 {
  "pkg": "github",
  "mod": "index/actionsRepositoryOidcSubjectClaimCustomizationTemplate",
  "fqn": "pulumi_github",
  "classes": {
   "github:index/actionsRepositoryOidcSubjectClaimCustomizationTemplate:ActionsRepositoryOidcSubjectClaimCustomizationTemplate": "ActionsRepositoryOidcSubjectClaimCustomizationTemplate"
  }
 },
 {
  "pkg": "github",
  "mod": "index/actionsRepositoryPermissions",
  "fqn": "pulumi_github",
  "classes": {
   "github:index/actionsRepositoryPermissions:ActionsRepositoryPermissions": "ActionsRepositoryPermissions"
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
  "mod": "index/actionsVariable",
  "fqn": "pulumi_github",
  "classes": {
   "github:index/actionsVariable:ActionsVariable": "ActionsVariable"
  }
 },
 {
  "pkg": "github",
  "mod": "index/appInstallationRepositories",
  "fqn": "pulumi_github",
  "classes": {
   "github:index/appInstallationRepositories:AppInstallationRepositories": "AppInstallationRepositories"
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
  "mod": "index/codespacesOrganizationSecret",
  "fqn": "pulumi_github",
  "classes": {
   "github:index/codespacesOrganizationSecret:CodespacesOrganizationSecret": "CodespacesOrganizationSecret"
  }
 },
 {
  "pkg": "github",
  "mod": "index/codespacesOrganizationSecretRepositories",
  "fqn": "pulumi_github",
  "classes": {
   "github:index/codespacesOrganizationSecretRepositories:CodespacesOrganizationSecretRepositories": "CodespacesOrganizationSecretRepositories"
  }
 },
 {
  "pkg": "github",
  "mod": "index/codespacesSecret",
  "fqn": "pulumi_github",
  "classes": {
   "github:index/codespacesSecret:CodespacesSecret": "CodespacesSecret"
  }
 },
 {
  "pkg": "github",
  "mod": "index/codespacesUserSecret",
  "fqn": "pulumi_github",
  "classes": {
   "github:index/codespacesUserSecret:CodespacesUserSecret": "CodespacesUserSecret"
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
  "mod": "index/enterpriseActionsPermissions",
  "fqn": "pulumi_github",
  "classes": {
   "github:index/enterpriseActionsPermissions:EnterpriseActionsPermissions": "EnterpriseActionsPermissions"
  }
 },
 {
  "pkg": "github",
  "mod": "index/enterpriseActionsRunnerGroup",
  "fqn": "pulumi_github",
  "classes": {
   "github:index/enterpriseActionsRunnerGroup:EnterpriseActionsRunnerGroup": "EnterpriseActionsRunnerGroup"
  }
 },
 {
  "pkg": "github",
  "mod": "index/enterpriseOrganization",
  "fqn": "pulumi_github",
  "classes": {
   "github:index/enterpriseOrganization:EnterpriseOrganization": "EnterpriseOrganization"
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
  "mod": "index/issueLabels",
  "fqn": "pulumi_github",
  "classes": {
   "github:index/issueLabels:IssueLabels": "IssueLabels"
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
  "mod": "index/organizationCustomRole",
  "fqn": "pulumi_github",
  "classes": {
   "github:index/organizationCustomRole:OrganizationCustomRole": "OrganizationCustomRole"
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
  "mod": "index/organizationRuleset",
  "fqn": "pulumi_github",
  "classes": {
   "github:index/organizationRuleset:OrganizationRuleset": "OrganizationRuleset"
  }
 },
 {
  "pkg": "github",
  "mod": "index/organizationSecurityManager",
  "fqn": "pulumi_github",
  "classes": {
   "github:index/organizationSecurityManager:OrganizationSecurityManager": "OrganizationSecurityManager"
  }
 },
 {
  "pkg": "github",
  "mod": "index/organizationSettings",
  "fqn": "pulumi_github",
  "classes": {
   "github:index/organizationSettings:OrganizationSettings": "OrganizationSettings"
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
  "mod": "index/release",
  "fqn": "pulumi_github",
  "classes": {
   "github:index/release:Release": "Release"
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
  "mod": "index/repositoryCollaborators",
  "fqn": "pulumi_github",
  "classes": {
   "github:index/repositoryCollaborators:RepositoryCollaborators": "RepositoryCollaborators"
  }
 },
 {
  "pkg": "github",
  "mod": "index/repositoryDependabotSecurityUpdates",
  "fqn": "pulumi_github",
  "classes": {
   "github:index/repositoryDependabotSecurityUpdates:RepositoryDependabotSecurityUpdates": "RepositoryDependabotSecurityUpdates"
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
  "mod": "index/repositoryDeploymentBranchPolicy",
  "fqn": "pulumi_github",
  "classes": {
   "github:index/repositoryDeploymentBranchPolicy:RepositoryDeploymentBranchPolicy": "RepositoryDeploymentBranchPolicy"
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
  "mod": "index/repositoryEnvironmentDeploymentPolicy",
  "fqn": "pulumi_github",
  "classes": {
   "github:index/repositoryEnvironmentDeploymentPolicy:RepositoryEnvironmentDeploymentPolicy": "RepositoryEnvironmentDeploymentPolicy"
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
  "mod": "index/repositoryRuleset",
  "fqn": "pulumi_github",
  "classes": {
   "github:index/repositoryRuleset:RepositoryRuleset": "RepositoryRuleset"
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
  "mod": "index/repositoryTopics",
  "fqn": "pulumi_github",
  "classes": {
   "github:index/repositoryTopics:RepositoryTopics": "RepositoryTopics"
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
  "mod": "index/teamSettings",
  "fqn": "pulumi_github",
  "classes": {
   "github:index/teamSettings:TeamSettings": "TeamSettings"
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
