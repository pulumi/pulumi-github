# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs

__all__ = [
    'GetOrganizationTeamsResult',
    'AwaitableGetOrganizationTeamsResult',
    'get_organization_teams',
]

@pulumi.output_type
class GetOrganizationTeamsResult:
    """
    A collection of values returned by getOrganizationTeams.
    """
    def __init__(__self__, id=None, root_teams_only=None, teams=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if root_teams_only and not isinstance(root_teams_only, bool):
            raise TypeError("Expected argument 'root_teams_only' to be a bool")
        pulumi.set(__self__, "root_teams_only", root_teams_only)
        if teams and not isinstance(teams, list):
            raise TypeError("Expected argument 'teams' to be a list")
        pulumi.set(__self__, "teams", teams)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="rootTeamsOnly")
    def root_teams_only(self) -> Optional[bool]:
        """
        Only return teams that are at the organization's root, i.e. no nested teams. Defaults to `false`.
        """
        return pulumi.get(self, "root_teams_only")

    @property
    @pulumi.getter
    def teams(self) -> Sequence['outputs.GetOrganizationTeamsTeamResult']:
        """
        An Array of GitHub Teams.  Each `team` block consists of the fields documented below.
        ___
        """
        return pulumi.get(self, "teams")


class AwaitableGetOrganizationTeamsResult(GetOrganizationTeamsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetOrganizationTeamsResult(
            id=self.id,
            root_teams_only=self.root_teams_only,
            teams=self.teams)


def get_organization_teams(root_teams_only: Optional[bool] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetOrganizationTeamsResult:
    """
    Use this data source to retrieve information about all GitHub teams in an organization.

    ## Example Usage

    To retrieve *all* teams of the organization:

    ```python
    import pulumi
    import pulumi_github as github

    all = github.get_organization_teams()
    ```

    To retrieve only the team's at the root of the organization:

    ```python
    import pulumi
    import pulumi_github as github

    root_teams = github.get_organization_teams(root_teams_only=True)
    ```


    :param bool root_teams_only: Only return teams that are at the organization's root, i.e. no nested teams. Defaults to `false`.
    """
    __args__ = dict()
    __args__['rootTeamsOnly'] = root_teams_only
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('github:index/getOrganizationTeams:getOrganizationTeams', __args__, opts=opts, typ=GetOrganizationTeamsResult).value

    return AwaitableGetOrganizationTeamsResult(
        id=__ret__.id,
        root_teams_only=__ret__.root_teams_only,
        teams=__ret__.teams)
