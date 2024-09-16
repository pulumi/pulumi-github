# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetOrganizationResult',
    'AwaitableGetOrganizationResult',
    'get_organization',
    'get_organization_output',
]

@pulumi.output_type
class GetOrganizationResult:
    """
    A collection of values returned by getOrganization.
    """
    def __init__(__self__, advanced_security_enabled_for_new_repositories=None, default_repository_permission=None, dependabot_alerts_enabled_for_new_repositories=None, dependabot_security_updates_enabled_for_new_repositories=None, dependency_graph_enabled_for_new_repositories=None, description=None, id=None, ignore_archived_repos=None, login=None, members=None, members_allowed_repository_creation_type=None, members_can_create_internal_repositories=None, members_can_create_pages=None, members_can_create_private_pages=None, members_can_create_private_repositories=None, members_can_create_public_pages=None, members_can_create_public_repositories=None, members_can_create_repositories=None, members_can_fork_private_repositories=None, name=None, node_id=None, orgname=None, plan=None, repositories=None, secret_scanning_enabled_for_new_repositories=None, secret_scanning_push_protection_enabled_for_new_repositories=None, summary_only=None, two_factor_requirement_enabled=None, users=None, web_commit_signoff_required=None):
        if advanced_security_enabled_for_new_repositories and not isinstance(advanced_security_enabled_for_new_repositories, bool):
            raise TypeError("Expected argument 'advanced_security_enabled_for_new_repositories' to be a bool")
        pulumi.set(__self__, "advanced_security_enabled_for_new_repositories", advanced_security_enabled_for_new_repositories)
        if default_repository_permission and not isinstance(default_repository_permission, str):
            raise TypeError("Expected argument 'default_repository_permission' to be a str")
        pulumi.set(__self__, "default_repository_permission", default_repository_permission)
        if dependabot_alerts_enabled_for_new_repositories and not isinstance(dependabot_alerts_enabled_for_new_repositories, bool):
            raise TypeError("Expected argument 'dependabot_alerts_enabled_for_new_repositories' to be a bool")
        pulumi.set(__self__, "dependabot_alerts_enabled_for_new_repositories", dependabot_alerts_enabled_for_new_repositories)
        if dependabot_security_updates_enabled_for_new_repositories and not isinstance(dependabot_security_updates_enabled_for_new_repositories, bool):
            raise TypeError("Expected argument 'dependabot_security_updates_enabled_for_new_repositories' to be a bool")
        pulumi.set(__self__, "dependabot_security_updates_enabled_for_new_repositories", dependabot_security_updates_enabled_for_new_repositories)
        if dependency_graph_enabled_for_new_repositories and not isinstance(dependency_graph_enabled_for_new_repositories, bool):
            raise TypeError("Expected argument 'dependency_graph_enabled_for_new_repositories' to be a bool")
        pulumi.set(__self__, "dependency_graph_enabled_for_new_repositories", dependency_graph_enabled_for_new_repositories)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ignore_archived_repos and not isinstance(ignore_archived_repos, bool):
            raise TypeError("Expected argument 'ignore_archived_repos' to be a bool")
        pulumi.set(__self__, "ignore_archived_repos", ignore_archived_repos)
        if login and not isinstance(login, str):
            raise TypeError("Expected argument 'login' to be a str")
        pulumi.set(__self__, "login", login)
        if members and not isinstance(members, list):
            raise TypeError("Expected argument 'members' to be a list")
        pulumi.set(__self__, "members", members)
        if members_allowed_repository_creation_type and not isinstance(members_allowed_repository_creation_type, str):
            raise TypeError("Expected argument 'members_allowed_repository_creation_type' to be a str")
        pulumi.set(__self__, "members_allowed_repository_creation_type", members_allowed_repository_creation_type)
        if members_can_create_internal_repositories and not isinstance(members_can_create_internal_repositories, bool):
            raise TypeError("Expected argument 'members_can_create_internal_repositories' to be a bool")
        pulumi.set(__self__, "members_can_create_internal_repositories", members_can_create_internal_repositories)
        if members_can_create_pages and not isinstance(members_can_create_pages, bool):
            raise TypeError("Expected argument 'members_can_create_pages' to be a bool")
        pulumi.set(__self__, "members_can_create_pages", members_can_create_pages)
        if members_can_create_private_pages and not isinstance(members_can_create_private_pages, bool):
            raise TypeError("Expected argument 'members_can_create_private_pages' to be a bool")
        pulumi.set(__self__, "members_can_create_private_pages", members_can_create_private_pages)
        if members_can_create_private_repositories and not isinstance(members_can_create_private_repositories, bool):
            raise TypeError("Expected argument 'members_can_create_private_repositories' to be a bool")
        pulumi.set(__self__, "members_can_create_private_repositories", members_can_create_private_repositories)
        if members_can_create_public_pages and not isinstance(members_can_create_public_pages, bool):
            raise TypeError("Expected argument 'members_can_create_public_pages' to be a bool")
        pulumi.set(__self__, "members_can_create_public_pages", members_can_create_public_pages)
        if members_can_create_public_repositories and not isinstance(members_can_create_public_repositories, bool):
            raise TypeError("Expected argument 'members_can_create_public_repositories' to be a bool")
        pulumi.set(__self__, "members_can_create_public_repositories", members_can_create_public_repositories)
        if members_can_create_repositories and not isinstance(members_can_create_repositories, bool):
            raise TypeError("Expected argument 'members_can_create_repositories' to be a bool")
        pulumi.set(__self__, "members_can_create_repositories", members_can_create_repositories)
        if members_can_fork_private_repositories and not isinstance(members_can_fork_private_repositories, bool):
            raise TypeError("Expected argument 'members_can_fork_private_repositories' to be a bool")
        pulumi.set(__self__, "members_can_fork_private_repositories", members_can_fork_private_repositories)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if node_id and not isinstance(node_id, str):
            raise TypeError("Expected argument 'node_id' to be a str")
        pulumi.set(__self__, "node_id", node_id)
        if orgname and not isinstance(orgname, str):
            raise TypeError("Expected argument 'orgname' to be a str")
        pulumi.set(__self__, "orgname", orgname)
        if plan and not isinstance(plan, str):
            raise TypeError("Expected argument 'plan' to be a str")
        pulumi.set(__self__, "plan", plan)
        if repositories and not isinstance(repositories, list):
            raise TypeError("Expected argument 'repositories' to be a list")
        pulumi.set(__self__, "repositories", repositories)
        if secret_scanning_enabled_for_new_repositories and not isinstance(secret_scanning_enabled_for_new_repositories, bool):
            raise TypeError("Expected argument 'secret_scanning_enabled_for_new_repositories' to be a bool")
        pulumi.set(__self__, "secret_scanning_enabled_for_new_repositories", secret_scanning_enabled_for_new_repositories)
        if secret_scanning_push_protection_enabled_for_new_repositories and not isinstance(secret_scanning_push_protection_enabled_for_new_repositories, bool):
            raise TypeError("Expected argument 'secret_scanning_push_protection_enabled_for_new_repositories' to be a bool")
        pulumi.set(__self__, "secret_scanning_push_protection_enabled_for_new_repositories", secret_scanning_push_protection_enabled_for_new_repositories)
        if summary_only and not isinstance(summary_only, bool):
            raise TypeError("Expected argument 'summary_only' to be a bool")
        pulumi.set(__self__, "summary_only", summary_only)
        if two_factor_requirement_enabled and not isinstance(two_factor_requirement_enabled, bool):
            raise TypeError("Expected argument 'two_factor_requirement_enabled' to be a bool")
        pulumi.set(__self__, "two_factor_requirement_enabled", two_factor_requirement_enabled)
        if users and not isinstance(users, list):
            raise TypeError("Expected argument 'users' to be a list")
        pulumi.set(__self__, "users", users)
        if web_commit_signoff_required and not isinstance(web_commit_signoff_required, bool):
            raise TypeError("Expected argument 'web_commit_signoff_required' to be a bool")
        pulumi.set(__self__, "web_commit_signoff_required", web_commit_signoff_required)

    @property
    @pulumi.getter(name="advancedSecurityEnabledForNewRepositories")
    def advanced_security_enabled_for_new_repositories(self) -> bool:
        """
        Whether advanced security is enabled for new repositories.
        """
        return pulumi.get(self, "advanced_security_enabled_for_new_repositories")

    @property
    @pulumi.getter(name="defaultRepositoryPermission")
    def default_repository_permission(self) -> str:
        """
        Default permission level members have for organization repositories.
        """
        return pulumi.get(self, "default_repository_permission")

    @property
    @pulumi.getter(name="dependabotAlertsEnabledForNewRepositories")
    def dependabot_alerts_enabled_for_new_repositories(self) -> bool:
        """
        Whether Dependabot alerts is automatically enabled for new repositories.
        """
        return pulumi.get(self, "dependabot_alerts_enabled_for_new_repositories")

    @property
    @pulumi.getter(name="dependabotSecurityUpdatesEnabledForNewRepositories")
    def dependabot_security_updates_enabled_for_new_repositories(self) -> bool:
        """
        Whether Dependabot security updates is automatically enabled for new repositories.
        """
        return pulumi.get(self, "dependabot_security_updates_enabled_for_new_repositories")

    @property
    @pulumi.getter(name="dependencyGraphEnabledForNewRepositories")
    def dependency_graph_enabled_for_new_repositories(self) -> bool:
        """
        Whether dependency graph is automatically enabled for new repositories.
        """
        return pulumi.get(self, "dependency_graph_enabled_for_new_repositories")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        The organization account description
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="ignoreArchivedRepos")
    def ignore_archived_repos(self) -> Optional[bool]:
        return pulumi.get(self, "ignore_archived_repos")

    @property
    @pulumi.getter
    def login(self) -> str:
        """
        The members login
        """
        return pulumi.get(self, "login")

    @property
    @pulumi.getter
    @_utilities.deprecated("""Use `users` instead by replacing `github_organization.example.members` to `github_organization.example.users[*].login`. Expect this field to be removed in next major version.""")
    def members(self) -> Sequence[str]:
        """
        **Deprecated**: use `users` instead by replacing `github_organization.example.members` to `github_organization.example.users[*].login` which will give you the same value, expect this field to be removed in next major version
        """
        return pulumi.get(self, "members")

    @property
    @pulumi.getter(name="membersAllowedRepositoryCreationType")
    def members_allowed_repository_creation_type(self) -> str:
        """
        The type of repository allowed to be created by members of the organization. Can be one of `ALL`, `PUBLIC`, `PRIVATE`, `NONE`.
        """
        return pulumi.get(self, "members_allowed_repository_creation_type")

    @property
    @pulumi.getter(name="membersCanCreateInternalRepositories")
    def members_can_create_internal_repositories(self) -> bool:
        """
        Whether organization members can create internal repositories.
        """
        return pulumi.get(self, "members_can_create_internal_repositories")

    @property
    @pulumi.getter(name="membersCanCreatePages")
    def members_can_create_pages(self) -> bool:
        """
        Whether organization members can create pages sites.
        """
        return pulumi.get(self, "members_can_create_pages")

    @property
    @pulumi.getter(name="membersCanCreatePrivatePages")
    def members_can_create_private_pages(self) -> bool:
        """
        Whether organization members can create private pages sites.
        """
        return pulumi.get(self, "members_can_create_private_pages")

    @property
    @pulumi.getter(name="membersCanCreatePrivateRepositories")
    def members_can_create_private_repositories(self) -> bool:
        """
        Whether organization members can create private repositories.
        """
        return pulumi.get(self, "members_can_create_private_repositories")

    @property
    @pulumi.getter(name="membersCanCreatePublicPages")
    def members_can_create_public_pages(self) -> bool:
        """
        Whether organization members can create public pages sites.
        """
        return pulumi.get(self, "members_can_create_public_pages")

    @property
    @pulumi.getter(name="membersCanCreatePublicRepositories")
    def members_can_create_public_repositories(self) -> bool:
        """
        Whether organization members can create public repositories.
        """
        return pulumi.get(self, "members_can_create_public_repositories")

    @property
    @pulumi.getter(name="membersCanCreateRepositories")
    def members_can_create_repositories(self) -> bool:
        """
        Whether non-admin organization members can create repositories.
        """
        return pulumi.get(self, "members_can_create_repositories")

    @property
    @pulumi.getter(name="membersCanForkPrivateRepositories")
    def members_can_fork_private_repositories(self) -> bool:
        """
        Whether organization members can create private repository forks.
        """
        return pulumi.get(self, "members_can_fork_private_repositories")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The organization's public profile name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nodeId")
    def node_id(self) -> str:
        """
        GraphQL global node ID for use with the v4 API
        """
        return pulumi.get(self, "node_id")

    @property
    @pulumi.getter
    def orgname(self) -> str:
        """
        The organization's name as used in URLs and the API
        """
        return pulumi.get(self, "orgname")

    @property
    @pulumi.getter
    def plan(self) -> str:
        """
        The organization account plan name
        """
        return pulumi.get(self, "plan")

    @property
    @pulumi.getter
    def repositories(self) -> Sequence[str]:
        """
        (`list`) A list of the full names of the repositories in the organization formatted as `owner/name` strings
        """
        return pulumi.get(self, "repositories")

    @property
    @pulumi.getter(name="secretScanningEnabledForNewRepositories")
    def secret_scanning_enabled_for_new_repositories(self) -> bool:
        """
        Whether secret scanning is automatically enabled for new repositories.
        """
        return pulumi.get(self, "secret_scanning_enabled_for_new_repositories")

    @property
    @pulumi.getter(name="secretScanningPushProtectionEnabledForNewRepositories")
    def secret_scanning_push_protection_enabled_for_new_repositories(self) -> bool:
        """
        Whether secret scanning push protection is automatically enabled for new repositories.
        """
        return pulumi.get(self, "secret_scanning_push_protection_enabled_for_new_repositories")

    @property
    @pulumi.getter(name="summaryOnly")
    def summary_only(self) -> Optional[bool]:
        return pulumi.get(self, "summary_only")

    @property
    @pulumi.getter(name="twoFactorRequirementEnabled")
    def two_factor_requirement_enabled(self) -> bool:
        """
        Whether two-factor authentication is required for all members of the organization.
        """
        return pulumi.get(self, "two_factor_requirement_enabled")

    @property
    @pulumi.getter
    def users(self) -> Sequence[Mapping[str, str]]:
        """
        (`list`) A list with the members of the organization with following fields:
        """
        return pulumi.get(self, "users")

    @property
    @pulumi.getter(name="webCommitSignoffRequired")
    def web_commit_signoff_required(self) -> bool:
        """
        Whether organization members must sign all commits.
        """
        return pulumi.get(self, "web_commit_signoff_required")


class AwaitableGetOrganizationResult(GetOrganizationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetOrganizationResult(
            advanced_security_enabled_for_new_repositories=self.advanced_security_enabled_for_new_repositories,
            default_repository_permission=self.default_repository_permission,
            dependabot_alerts_enabled_for_new_repositories=self.dependabot_alerts_enabled_for_new_repositories,
            dependabot_security_updates_enabled_for_new_repositories=self.dependabot_security_updates_enabled_for_new_repositories,
            dependency_graph_enabled_for_new_repositories=self.dependency_graph_enabled_for_new_repositories,
            description=self.description,
            id=self.id,
            ignore_archived_repos=self.ignore_archived_repos,
            login=self.login,
            members=self.members,
            members_allowed_repository_creation_type=self.members_allowed_repository_creation_type,
            members_can_create_internal_repositories=self.members_can_create_internal_repositories,
            members_can_create_pages=self.members_can_create_pages,
            members_can_create_private_pages=self.members_can_create_private_pages,
            members_can_create_private_repositories=self.members_can_create_private_repositories,
            members_can_create_public_pages=self.members_can_create_public_pages,
            members_can_create_public_repositories=self.members_can_create_public_repositories,
            members_can_create_repositories=self.members_can_create_repositories,
            members_can_fork_private_repositories=self.members_can_fork_private_repositories,
            name=self.name,
            node_id=self.node_id,
            orgname=self.orgname,
            plan=self.plan,
            repositories=self.repositories,
            secret_scanning_enabled_for_new_repositories=self.secret_scanning_enabled_for_new_repositories,
            secret_scanning_push_protection_enabled_for_new_repositories=self.secret_scanning_push_protection_enabled_for_new_repositories,
            summary_only=self.summary_only,
            two_factor_requirement_enabled=self.two_factor_requirement_enabled,
            users=self.users,
            web_commit_signoff_required=self.web_commit_signoff_required)


def get_organization(ignore_archived_repos: Optional[bool] = None,
                     name: Optional[str] = None,
                     summary_only: Optional[bool] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetOrganizationResult:
    """
    Use this data source to retrieve basic information about a GitHub Organization.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_github as github

    example = github.get_organization(name="github")
    ```


    :param bool ignore_archived_repos: Whether or not to include archived repos in the `repositories` list. Defaults to `false`.
    :param str name: The name of the organization.
    :param bool summary_only: Exclude the repos, members and other attributes from the returned result. Defaults to `false`.
    """
    __args__ = dict()
    __args__['ignoreArchivedRepos'] = ignore_archived_repos
    __args__['name'] = name
    __args__['summaryOnly'] = summary_only
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('github:index/getOrganization:getOrganization', __args__, opts=opts, typ=GetOrganizationResult).value

    return AwaitableGetOrganizationResult(
        advanced_security_enabled_for_new_repositories=pulumi.get(__ret__, 'advanced_security_enabled_for_new_repositories'),
        default_repository_permission=pulumi.get(__ret__, 'default_repository_permission'),
        dependabot_alerts_enabled_for_new_repositories=pulumi.get(__ret__, 'dependabot_alerts_enabled_for_new_repositories'),
        dependabot_security_updates_enabled_for_new_repositories=pulumi.get(__ret__, 'dependabot_security_updates_enabled_for_new_repositories'),
        dependency_graph_enabled_for_new_repositories=pulumi.get(__ret__, 'dependency_graph_enabled_for_new_repositories'),
        description=pulumi.get(__ret__, 'description'),
        id=pulumi.get(__ret__, 'id'),
        ignore_archived_repos=pulumi.get(__ret__, 'ignore_archived_repos'),
        login=pulumi.get(__ret__, 'login'),
        members=pulumi.get(__ret__, 'members'),
        members_allowed_repository_creation_type=pulumi.get(__ret__, 'members_allowed_repository_creation_type'),
        members_can_create_internal_repositories=pulumi.get(__ret__, 'members_can_create_internal_repositories'),
        members_can_create_pages=pulumi.get(__ret__, 'members_can_create_pages'),
        members_can_create_private_pages=pulumi.get(__ret__, 'members_can_create_private_pages'),
        members_can_create_private_repositories=pulumi.get(__ret__, 'members_can_create_private_repositories'),
        members_can_create_public_pages=pulumi.get(__ret__, 'members_can_create_public_pages'),
        members_can_create_public_repositories=pulumi.get(__ret__, 'members_can_create_public_repositories'),
        members_can_create_repositories=pulumi.get(__ret__, 'members_can_create_repositories'),
        members_can_fork_private_repositories=pulumi.get(__ret__, 'members_can_fork_private_repositories'),
        name=pulumi.get(__ret__, 'name'),
        node_id=pulumi.get(__ret__, 'node_id'),
        orgname=pulumi.get(__ret__, 'orgname'),
        plan=pulumi.get(__ret__, 'plan'),
        repositories=pulumi.get(__ret__, 'repositories'),
        secret_scanning_enabled_for_new_repositories=pulumi.get(__ret__, 'secret_scanning_enabled_for_new_repositories'),
        secret_scanning_push_protection_enabled_for_new_repositories=pulumi.get(__ret__, 'secret_scanning_push_protection_enabled_for_new_repositories'),
        summary_only=pulumi.get(__ret__, 'summary_only'),
        two_factor_requirement_enabled=pulumi.get(__ret__, 'two_factor_requirement_enabled'),
        users=pulumi.get(__ret__, 'users'),
        web_commit_signoff_required=pulumi.get(__ret__, 'web_commit_signoff_required'))


@_utilities.lift_output_func(get_organization)
def get_organization_output(ignore_archived_repos: Optional[pulumi.Input[Optional[bool]]] = None,
                            name: Optional[pulumi.Input[str]] = None,
                            summary_only: Optional[pulumi.Input[Optional[bool]]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetOrganizationResult]:
    """
    Use this data source to retrieve basic information about a GitHub Organization.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_github as github

    example = github.get_organization(name="github")
    ```


    :param bool ignore_archived_repos: Whether or not to include archived repos in the `repositories` list. Defaults to `false`.
    :param str name: The name of the organization.
    :param bool summary_only: Exclude the repos, members and other attributes from the returned result. Defaults to `false`.
    """
    ...
