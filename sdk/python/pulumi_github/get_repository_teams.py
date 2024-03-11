# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs

__all__ = [
    'GetRepositoryTeamsResult',
    'AwaitableGetRepositoryTeamsResult',
    'get_repository_teams',
    'get_repository_teams_output',
]

@pulumi.output_type
class GetRepositoryTeamsResult:
    """
    A collection of values returned by getRepositoryTeams.
    """
    def __init__(__self__, full_name=None, id=None, name=None, teams=None):
        if full_name and not isinstance(full_name, str):
            raise TypeError("Expected argument 'full_name' to be a str")
        pulumi.set(__self__, "full_name", full_name)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if teams and not isinstance(teams, list):
            raise TypeError("Expected argument 'teams' to be a list")
        pulumi.set(__self__, "teams", teams)

    @property
    @pulumi.getter(name="fullName")
    def full_name(self) -> str:
        return pulumi.get(self, "full_name")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Team name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def teams(self) -> Sequence['outputs.GetRepositoryTeamsTeamResult']:
        """
        List of teams which have access to the repository
        """
        return pulumi.get(self, "teams")


class AwaitableGetRepositoryTeamsResult(GetRepositoryTeamsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRepositoryTeamsResult(
            full_name=self.full_name,
            id=self.id,
            name=self.name,
            teams=self.teams)


def get_repository_teams(full_name: Optional[str] = None,
                         name: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRepositoryTeamsResult:
    """
    Use this data source to retrieve the list of teams which have access to a GitHub repository.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_github as github

    example = github.get_repository_teams(name="example")
    ```
    <!--End PulumiCodeChooser -->


    :param str full_name: Full name of the repository (in `org/name` format).
    :param str name: The name of the repository.
    """
    __args__ = dict()
    __args__['fullName'] = full_name
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('github:index/getRepositoryTeams:getRepositoryTeams', __args__, opts=opts, typ=GetRepositoryTeamsResult).value

    return AwaitableGetRepositoryTeamsResult(
        full_name=pulumi.get(__ret__, 'full_name'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        teams=pulumi.get(__ret__, 'teams'))


@_utilities.lift_output_func(get_repository_teams)
def get_repository_teams_output(full_name: Optional[pulumi.Input[Optional[str]]] = None,
                                name: Optional[pulumi.Input[Optional[str]]] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetRepositoryTeamsResult]:
    """
    Use this data source to retrieve the list of teams which have access to a GitHub repository.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_github as github

    example = github.get_repository_teams(name="example")
    ```
    <!--End PulumiCodeChooser -->


    :param str full_name: Full name of the repository (in `org/name` format).
    :param str name: The name of the repository.
    """
    ...
