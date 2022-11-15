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
    'GetOrganizationTeamsResult',
    'AwaitableGetOrganizationTeamsResult',
    'get_organization_teams',
    'get_organization_teams_output',
]

@pulumi.output_type
class GetOrganizationTeamsResult:
    """
    A collection of values returned by getOrganizationTeams.
    """
    def __init__(__self__, id=None, results_per_page=None, root_teams_only=None, summary_only=None, teams=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if results_per_page and not isinstance(results_per_page, int):
            raise TypeError("Expected argument 'results_per_page' to be a int")
        pulumi.set(__self__, "results_per_page", results_per_page)
        if root_teams_only and not isinstance(root_teams_only, bool):
            raise TypeError("Expected argument 'root_teams_only' to be a bool")
        pulumi.set(__self__, "root_teams_only", root_teams_only)
        if summary_only and not isinstance(summary_only, bool):
            raise TypeError("Expected argument 'summary_only' to be a bool")
        pulumi.set(__self__, "summary_only", summary_only)
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
    @pulumi.getter(name="resultsPerPage")
    def results_per_page(self) -> Optional[int]:
        return pulumi.get(self, "results_per_page")

    @property
    @pulumi.getter(name="rootTeamsOnly")
    def root_teams_only(self) -> Optional[bool]:
        return pulumi.get(self, "root_teams_only")

    @property
    @pulumi.getter(name="summaryOnly")
    def summary_only(self) -> Optional[bool]:
        return pulumi.get(self, "summary_only")

    @property
    @pulumi.getter
    def teams(self) -> Sequence['outputs.GetOrganizationTeamsTeamResult']:
        return pulumi.get(self, "teams")


class AwaitableGetOrganizationTeamsResult(GetOrganizationTeamsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetOrganizationTeamsResult(
            id=self.id,
            results_per_page=self.results_per_page,
            root_teams_only=self.root_teams_only,
            summary_only=self.summary_only,
            teams=self.teams)


def get_organization_teams(results_per_page: Optional[int] = None,
                           root_teams_only: Optional[bool] = None,
                           summary_only: Optional[bool] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetOrganizationTeamsResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['resultsPerPage'] = results_per_page
    __args__['rootTeamsOnly'] = root_teams_only
    __args__['summaryOnly'] = summary_only
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('github:index/getOrganizationTeams:getOrganizationTeams', __args__, opts=opts, typ=GetOrganizationTeamsResult).value

    return AwaitableGetOrganizationTeamsResult(
        id=__ret__.id,
        results_per_page=__ret__.results_per_page,
        root_teams_only=__ret__.root_teams_only,
        summary_only=__ret__.summary_only,
        teams=__ret__.teams)


@_utilities.lift_output_func(get_organization_teams)
def get_organization_teams_output(results_per_page: Optional[pulumi.Input[Optional[int]]] = None,
                                  root_teams_only: Optional[pulumi.Input[Optional[bool]]] = None,
                                  summary_only: Optional[pulumi.Input[Optional[bool]]] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetOrganizationTeamsResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
