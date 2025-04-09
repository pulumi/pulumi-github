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
    def id(self) -> builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="resultsPerPage")
    def results_per_page(self) -> Optional[builtins.int]:
        """
        (Optional) Set the number of results per graphql query. Reducing this number can alleviate timeout errors. Accepts a value between 0 - 100. Defaults to `100`.
        """
        return pulumi.get(self, "results_per_page")

    @property
    @pulumi.getter(name="rootTeamsOnly")
    def root_teams_only(self) -> Optional[builtins.bool]:
        """
        (Optional) Only return teams that are at the organization's root, i.e. no nested teams. Defaults to `false`.
        """
        return pulumi.get(self, "root_teams_only")

    @property
    @pulumi.getter(name="summaryOnly")
    def summary_only(self) -> Optional[builtins.bool]:
        """
        (Optional) Exclude the members and repositories of the team from the returned result. Defaults to `false`.
        """
        return pulumi.get(self, "summary_only")

    @property
    @pulumi.getter
    def teams(self) -> Sequence['outputs.GetOrganizationTeamsTeamResult']:
        """
        (Required) An Array of GitHub Teams.  Each `team` block consists of the fields documented below.
        """
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


def get_organization_teams(results_per_page: Optional[builtins.int] = None,
                           root_teams_only: Optional[builtins.bool] = None,
                           summary_only: Optional[builtins.bool] = None,
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


    :param builtins.int results_per_page: (Optional) Set the number of results per graphql query. Reducing this number can alleviate timeout errors. Accepts a value between 0 - 100. Defaults to `100`.
    :param builtins.bool root_teams_only: (Optional) Only return teams that are at the organization's root, i.e. no nested teams. Defaults to `false`.
    :param builtins.bool summary_only: (Optional) Exclude the members and repositories of the team from the returned result. Defaults to `false`.
    """
    __args__ = dict()
    __args__['resultsPerPage'] = results_per_page
    __args__['rootTeamsOnly'] = root_teams_only
    __args__['summaryOnly'] = summary_only
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('github:index/getOrganizationTeams:getOrganizationTeams', __args__, opts=opts, typ=GetOrganizationTeamsResult).value

    return AwaitableGetOrganizationTeamsResult(
        id=pulumi.get(__ret__, 'id'),
        results_per_page=pulumi.get(__ret__, 'results_per_page'),
        root_teams_only=pulumi.get(__ret__, 'root_teams_only'),
        summary_only=pulumi.get(__ret__, 'summary_only'),
        teams=pulumi.get(__ret__, 'teams'))
def get_organization_teams_output(results_per_page: Optional[pulumi.Input[Optional[builtins.int]]] = None,
                                  root_teams_only: Optional[pulumi.Input[Optional[builtins.bool]]] = None,
                                  summary_only: Optional[pulumi.Input[Optional[builtins.bool]]] = None,
                                  opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetOrganizationTeamsResult]:
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


    :param builtins.int results_per_page: (Optional) Set the number of results per graphql query. Reducing this number can alleviate timeout errors. Accepts a value between 0 - 100. Defaults to `100`.
    :param builtins.bool root_teams_only: (Optional) Only return teams that are at the organization's root, i.e. no nested teams. Defaults to `false`.
    :param builtins.bool summary_only: (Optional) Exclude the members and repositories of the team from the returned result. Defaults to `false`.
    """
    __args__ = dict()
    __args__['resultsPerPage'] = results_per_page
    __args__['rootTeamsOnly'] = root_teams_only
    __args__['summaryOnly'] = summary_only
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('github:index/getOrganizationTeams:getOrganizationTeams', __args__, opts=opts, typ=GetOrganizationTeamsResult)
    return __ret__.apply(lambda __response__: GetOrganizationTeamsResult(
        id=pulumi.get(__response__, 'id'),
        results_per_page=pulumi.get(__response__, 'results_per_page'),
        root_teams_only=pulumi.get(__response__, 'root_teams_only'),
        summary_only=pulumi.get(__response__, 'summary_only'),
        teams=pulumi.get(__response__, 'teams')))
