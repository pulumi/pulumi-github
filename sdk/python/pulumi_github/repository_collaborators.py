# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['RepositoryCollaboratorsArgs', 'RepositoryCollaborators']

@pulumi.input_type
class RepositoryCollaboratorsArgs:
    def __init__(__self__, *,
                 repository: pulumi.Input[builtins.str],
                 ignore_teams: Optional[pulumi.Input[Sequence[pulumi.Input['RepositoryCollaboratorsIgnoreTeamArgs']]]] = None,
                 teams: Optional[pulumi.Input[Sequence[pulumi.Input['RepositoryCollaboratorsTeamArgs']]]] = None,
                 users: Optional[pulumi.Input[Sequence[pulumi.Input['RepositoryCollaboratorsUserArgs']]]] = None):
        """
        The set of arguments for constructing a RepositoryCollaborators resource.
        :param pulumi.Input[builtins.str] repository: The GitHub repository.
        :param pulumi.Input[Sequence[pulumi.Input['RepositoryCollaboratorsIgnoreTeamArgs']]] ignore_teams: List of teams to ignore when checking for repository access. This supports ignoring teams granted access at an organizational level.
        :param pulumi.Input[Sequence[pulumi.Input['RepositoryCollaboratorsTeamArgs']]] teams: List of teams to grant access to the repository.
        :param pulumi.Input[Sequence[pulumi.Input['RepositoryCollaboratorsUserArgs']]] users: List of users to grant access to the repository.
        """
        pulumi.set(__self__, "repository", repository)
        if ignore_teams is not None:
            pulumi.set(__self__, "ignore_teams", ignore_teams)
        if teams is not None:
            pulumi.set(__self__, "teams", teams)
        if users is not None:
            pulumi.set(__self__, "users", users)

    @property
    @pulumi.getter
    def repository(self) -> pulumi.Input[builtins.str]:
        """
        The GitHub repository.
        """
        return pulumi.get(self, "repository")

    @repository.setter
    def repository(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "repository", value)

    @property
    @pulumi.getter(name="ignoreTeams")
    def ignore_teams(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RepositoryCollaboratorsIgnoreTeamArgs']]]]:
        """
        List of teams to ignore when checking for repository access. This supports ignoring teams granted access at an organizational level.
        """
        return pulumi.get(self, "ignore_teams")

    @ignore_teams.setter
    def ignore_teams(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RepositoryCollaboratorsIgnoreTeamArgs']]]]):
        pulumi.set(self, "ignore_teams", value)

    @property
    @pulumi.getter
    def teams(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RepositoryCollaboratorsTeamArgs']]]]:
        """
        List of teams to grant access to the repository.
        """
        return pulumi.get(self, "teams")

    @teams.setter
    def teams(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RepositoryCollaboratorsTeamArgs']]]]):
        pulumi.set(self, "teams", value)

    @property
    @pulumi.getter
    def users(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RepositoryCollaboratorsUserArgs']]]]:
        """
        List of users to grant access to the repository.
        """
        return pulumi.get(self, "users")

    @users.setter
    def users(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RepositoryCollaboratorsUserArgs']]]]):
        pulumi.set(self, "users", value)


@pulumi.input_type
class _RepositoryCollaboratorsState:
    def __init__(__self__, *,
                 ignore_teams: Optional[pulumi.Input[Sequence[pulumi.Input['RepositoryCollaboratorsIgnoreTeamArgs']]]] = None,
                 invitation_ids: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 repository: Optional[pulumi.Input[builtins.str]] = None,
                 teams: Optional[pulumi.Input[Sequence[pulumi.Input['RepositoryCollaboratorsTeamArgs']]]] = None,
                 users: Optional[pulumi.Input[Sequence[pulumi.Input['RepositoryCollaboratorsUserArgs']]]] = None):
        """
        Input properties used for looking up and filtering RepositoryCollaborators resources.
        :param pulumi.Input[Sequence[pulumi.Input['RepositoryCollaboratorsIgnoreTeamArgs']]] ignore_teams: List of teams to ignore when checking for repository access. This supports ignoring teams granted access at an organizational level.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] invitation_ids: Map of usernames to invitation ID for any users added as part of creation of this resource to
               be used in `UserInvitationAccepter`.
        :param pulumi.Input[builtins.str] repository: The GitHub repository.
        :param pulumi.Input[Sequence[pulumi.Input['RepositoryCollaboratorsTeamArgs']]] teams: List of teams to grant access to the repository.
        :param pulumi.Input[Sequence[pulumi.Input['RepositoryCollaboratorsUserArgs']]] users: List of users to grant access to the repository.
        """
        if ignore_teams is not None:
            pulumi.set(__self__, "ignore_teams", ignore_teams)
        if invitation_ids is not None:
            pulumi.set(__self__, "invitation_ids", invitation_ids)
        if repository is not None:
            pulumi.set(__self__, "repository", repository)
        if teams is not None:
            pulumi.set(__self__, "teams", teams)
        if users is not None:
            pulumi.set(__self__, "users", users)

    @property
    @pulumi.getter(name="ignoreTeams")
    def ignore_teams(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RepositoryCollaboratorsIgnoreTeamArgs']]]]:
        """
        List of teams to ignore when checking for repository access. This supports ignoring teams granted access at an organizational level.
        """
        return pulumi.get(self, "ignore_teams")

    @ignore_teams.setter
    def ignore_teams(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RepositoryCollaboratorsIgnoreTeamArgs']]]]):
        pulumi.set(self, "ignore_teams", value)

    @property
    @pulumi.getter(name="invitationIds")
    def invitation_ids(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]:
        """
        Map of usernames to invitation ID for any users added as part of creation of this resource to
        be used in `UserInvitationAccepter`.
        """
        return pulumi.get(self, "invitation_ids")

    @invitation_ids.setter
    def invitation_ids(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "invitation_ids", value)

    @property
    @pulumi.getter
    def repository(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The GitHub repository.
        """
        return pulumi.get(self, "repository")

    @repository.setter
    def repository(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "repository", value)

    @property
    @pulumi.getter
    def teams(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RepositoryCollaboratorsTeamArgs']]]]:
        """
        List of teams to grant access to the repository.
        """
        return pulumi.get(self, "teams")

    @teams.setter
    def teams(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RepositoryCollaboratorsTeamArgs']]]]):
        pulumi.set(self, "teams", value)

    @property
    @pulumi.getter
    def users(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RepositoryCollaboratorsUserArgs']]]]:
        """
        List of users to grant access to the repository.
        """
        return pulumi.get(self, "users")

    @users.setter
    def users(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RepositoryCollaboratorsUserArgs']]]]):
        pulumi.set(self, "users", value)


class RepositoryCollaborators(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ignore_teams: Optional[pulumi.Input[Sequence[pulumi.Input[Union['RepositoryCollaboratorsIgnoreTeamArgs', 'RepositoryCollaboratorsIgnoreTeamArgsDict']]]]] = None,
                 repository: Optional[pulumi.Input[builtins.str]] = None,
                 teams: Optional[pulumi.Input[Sequence[pulumi.Input[Union['RepositoryCollaboratorsTeamArgs', 'RepositoryCollaboratorsTeamArgsDict']]]]] = None,
                 users: Optional[pulumi.Input[Sequence[pulumi.Input[Union['RepositoryCollaboratorsUserArgs', 'RepositoryCollaboratorsUserArgsDict']]]]] = None,
                 __props__=None):
        """
        Provides a GitHub repository collaborators resource.

        > Note: RepositoryCollaborators cannot be used in conjunction with RepositoryCollaborator and
        TeamRepository or they will fight over what your policy should be.

        This resource allows you to manage all collaborators for repositories in your
        organization or personal account. For organization repositories, collaborators can
        have explicit (and differing levels of) read, write, or administrator access to
        specific repositories, without giving the user full organization membership.
        For personal repositories, collaborators can only be granted write
        (implicitly includes read) permission.

        When applied, an invitation will be sent to the user to become a collaborators
        on a repository. When destroyed, either the invitation will be cancelled or the
        collaborators will be removed from the repository.

        This resource is authoritative. For adding a collaborator to a repo in a non-authoritative manner, use
        RepositoryCollaborator instead.

        Further documentation on GitHub collaborators:

        - [Adding outside collaborators to your personal repositories](https://help.github.com/en/github/setting-up-and-managing-your-github-user-account/managing-access-to-your-personal-repositories)
        - [Adding outside collaborators to repositories in your organization](https://help.github.com/articles/adding-outside-collaborators-to-repositories-in-your-organization/)
        - [Converting an organization member to an outside collaborators](https://help.github.com/articles/converting-an-organization-member-to-an-outside-collaborator/)

        ## Example Usage

        ```python
        import pulumi
        import pulumi_github as github

        # Add collaborators to a repository
        some_team = github.Team("some_team",
            name="SomeTeam",
            description="Some cool team")
        some_repo = github.Repository("some_repo", name="some-repo")
        some_repo_collaborators = github.RepositoryCollaborators("some_repo_collaborators",
            repository=some_repo.name,
            users=[{
                "permission": "admin",
                "username": "SomeUser",
            }],
            teams=[{
                "permission": "pull",
                "team_id": some_team.slug,
            }])
        ```

        ## Import

        GitHub Repository Collaborators can be imported using the name `name`, e.g.

        ```sh
        $ pulumi import github:index/repositoryCollaborators:RepositoryCollaborators collaborators terraform
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[Union['RepositoryCollaboratorsIgnoreTeamArgs', 'RepositoryCollaboratorsIgnoreTeamArgsDict']]]] ignore_teams: List of teams to ignore when checking for repository access. This supports ignoring teams granted access at an organizational level.
        :param pulumi.Input[builtins.str] repository: The GitHub repository.
        :param pulumi.Input[Sequence[pulumi.Input[Union['RepositoryCollaboratorsTeamArgs', 'RepositoryCollaboratorsTeamArgsDict']]]] teams: List of teams to grant access to the repository.
        :param pulumi.Input[Sequence[pulumi.Input[Union['RepositoryCollaboratorsUserArgs', 'RepositoryCollaboratorsUserArgsDict']]]] users: List of users to grant access to the repository.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RepositoryCollaboratorsArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a GitHub repository collaborators resource.

        > Note: RepositoryCollaborators cannot be used in conjunction with RepositoryCollaborator and
        TeamRepository or they will fight over what your policy should be.

        This resource allows you to manage all collaborators for repositories in your
        organization or personal account. For organization repositories, collaborators can
        have explicit (and differing levels of) read, write, or administrator access to
        specific repositories, without giving the user full organization membership.
        For personal repositories, collaborators can only be granted write
        (implicitly includes read) permission.

        When applied, an invitation will be sent to the user to become a collaborators
        on a repository. When destroyed, either the invitation will be cancelled or the
        collaborators will be removed from the repository.

        This resource is authoritative. For adding a collaborator to a repo in a non-authoritative manner, use
        RepositoryCollaborator instead.

        Further documentation on GitHub collaborators:

        - [Adding outside collaborators to your personal repositories](https://help.github.com/en/github/setting-up-and-managing-your-github-user-account/managing-access-to-your-personal-repositories)
        - [Adding outside collaborators to repositories in your organization](https://help.github.com/articles/adding-outside-collaborators-to-repositories-in-your-organization/)
        - [Converting an organization member to an outside collaborators](https://help.github.com/articles/converting-an-organization-member-to-an-outside-collaborator/)

        ## Example Usage

        ```python
        import pulumi
        import pulumi_github as github

        # Add collaborators to a repository
        some_team = github.Team("some_team",
            name="SomeTeam",
            description="Some cool team")
        some_repo = github.Repository("some_repo", name="some-repo")
        some_repo_collaborators = github.RepositoryCollaborators("some_repo_collaborators",
            repository=some_repo.name,
            users=[{
                "permission": "admin",
                "username": "SomeUser",
            }],
            teams=[{
                "permission": "pull",
                "team_id": some_team.slug,
            }])
        ```

        ## Import

        GitHub Repository Collaborators can be imported using the name `name`, e.g.

        ```sh
        $ pulumi import github:index/repositoryCollaborators:RepositoryCollaborators collaborators terraform
        ```

        :param str resource_name: The name of the resource.
        :param RepositoryCollaboratorsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RepositoryCollaboratorsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ignore_teams: Optional[pulumi.Input[Sequence[pulumi.Input[Union['RepositoryCollaboratorsIgnoreTeamArgs', 'RepositoryCollaboratorsIgnoreTeamArgsDict']]]]] = None,
                 repository: Optional[pulumi.Input[builtins.str]] = None,
                 teams: Optional[pulumi.Input[Sequence[pulumi.Input[Union['RepositoryCollaboratorsTeamArgs', 'RepositoryCollaboratorsTeamArgsDict']]]]] = None,
                 users: Optional[pulumi.Input[Sequence[pulumi.Input[Union['RepositoryCollaboratorsUserArgs', 'RepositoryCollaboratorsUserArgsDict']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RepositoryCollaboratorsArgs.__new__(RepositoryCollaboratorsArgs)

            __props__.__dict__["ignore_teams"] = ignore_teams
            if repository is None and not opts.urn:
                raise TypeError("Missing required property 'repository'")
            __props__.__dict__["repository"] = repository
            __props__.__dict__["teams"] = teams
            __props__.__dict__["users"] = users
            __props__.__dict__["invitation_ids"] = None
        super(RepositoryCollaborators, __self__).__init__(
            'github:index/repositoryCollaborators:RepositoryCollaborators',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            ignore_teams: Optional[pulumi.Input[Sequence[pulumi.Input[Union['RepositoryCollaboratorsIgnoreTeamArgs', 'RepositoryCollaboratorsIgnoreTeamArgsDict']]]]] = None,
            invitation_ids: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
            repository: Optional[pulumi.Input[builtins.str]] = None,
            teams: Optional[pulumi.Input[Sequence[pulumi.Input[Union['RepositoryCollaboratorsTeamArgs', 'RepositoryCollaboratorsTeamArgsDict']]]]] = None,
            users: Optional[pulumi.Input[Sequence[pulumi.Input[Union['RepositoryCollaboratorsUserArgs', 'RepositoryCollaboratorsUserArgsDict']]]]] = None) -> 'RepositoryCollaborators':
        """
        Get an existing RepositoryCollaborators resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[Union['RepositoryCollaboratorsIgnoreTeamArgs', 'RepositoryCollaboratorsIgnoreTeamArgsDict']]]] ignore_teams: List of teams to ignore when checking for repository access. This supports ignoring teams granted access at an organizational level.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] invitation_ids: Map of usernames to invitation ID for any users added as part of creation of this resource to
               be used in `UserInvitationAccepter`.
        :param pulumi.Input[builtins.str] repository: The GitHub repository.
        :param pulumi.Input[Sequence[pulumi.Input[Union['RepositoryCollaboratorsTeamArgs', 'RepositoryCollaboratorsTeamArgsDict']]]] teams: List of teams to grant access to the repository.
        :param pulumi.Input[Sequence[pulumi.Input[Union['RepositoryCollaboratorsUserArgs', 'RepositoryCollaboratorsUserArgsDict']]]] users: List of users to grant access to the repository.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RepositoryCollaboratorsState.__new__(_RepositoryCollaboratorsState)

        __props__.__dict__["ignore_teams"] = ignore_teams
        __props__.__dict__["invitation_ids"] = invitation_ids
        __props__.__dict__["repository"] = repository
        __props__.__dict__["teams"] = teams
        __props__.__dict__["users"] = users
        return RepositoryCollaborators(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="ignoreTeams")
    def ignore_teams(self) -> pulumi.Output[Optional[Sequence['outputs.RepositoryCollaboratorsIgnoreTeam']]]:
        """
        List of teams to ignore when checking for repository access. This supports ignoring teams granted access at an organizational level.
        """
        return pulumi.get(self, "ignore_teams")

    @property
    @pulumi.getter(name="invitationIds")
    def invitation_ids(self) -> pulumi.Output[Mapping[str, builtins.str]]:
        """
        Map of usernames to invitation ID for any users added as part of creation of this resource to
        be used in `UserInvitationAccepter`.
        """
        return pulumi.get(self, "invitation_ids")

    @property
    @pulumi.getter
    def repository(self) -> pulumi.Output[builtins.str]:
        """
        The GitHub repository.
        """
        return pulumi.get(self, "repository")

    @property
    @pulumi.getter
    def teams(self) -> pulumi.Output[Optional[Sequence['outputs.RepositoryCollaboratorsTeam']]]:
        """
        List of teams to grant access to the repository.
        """
        return pulumi.get(self, "teams")

    @property
    @pulumi.getter
    def users(self) -> pulumi.Output[Optional[Sequence['outputs.RepositoryCollaboratorsUser']]]:
        """
        List of users to grant access to the repository.
        """
        return pulumi.get(self, "users")

