# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

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
    'GetRepositoryPullRequestsResult',
    'AwaitableGetRepositoryPullRequestsResult',
    'get_repository_pull_requests',
    'get_repository_pull_requests_output',
]

@pulumi.output_type
class GetRepositoryPullRequestsResult:
    """
    A collection of values returned by getRepositoryPullRequests.
    """
    def __init__(__self__, base_ref=None, base_repository=None, head_ref=None, id=None, owner=None, results=None, sort_by=None, sort_direction=None, state=None):
        if base_ref and not isinstance(base_ref, str):
            raise TypeError("Expected argument 'base_ref' to be a str")
        pulumi.set(__self__, "base_ref", base_ref)
        if base_repository and not isinstance(base_repository, str):
            raise TypeError("Expected argument 'base_repository' to be a str")
        pulumi.set(__self__, "base_repository", base_repository)
        if head_ref and not isinstance(head_ref, str):
            raise TypeError("Expected argument 'head_ref' to be a str")
        pulumi.set(__self__, "head_ref", head_ref)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if owner and not isinstance(owner, str):
            raise TypeError("Expected argument 'owner' to be a str")
        pulumi.set(__self__, "owner", owner)
        if results and not isinstance(results, list):
            raise TypeError("Expected argument 'results' to be a list")
        pulumi.set(__self__, "results", results)
        if sort_by and not isinstance(sort_by, str):
            raise TypeError("Expected argument 'sort_by' to be a str")
        pulumi.set(__self__, "sort_by", sort_by)
        if sort_direction and not isinstance(sort_direction, str):
            raise TypeError("Expected argument 'sort_direction' to be a str")
        pulumi.set(__self__, "sort_direction", sort_direction)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)

    @property
    @pulumi.getter(name="baseRef")
    def base_ref(self) -> Optional[str]:
        """
        Name of the ref (branch) of the Pull Request base.
        """
        return pulumi.get(self, "base_ref")

    @property
    @pulumi.getter(name="baseRepository")
    def base_repository(self) -> str:
        return pulumi.get(self, "base_repository")

    @property
    @pulumi.getter(name="headRef")
    def head_ref(self) -> Optional[str]:
        """
        Value of the Pull Request `HEAD` reference.
        """
        return pulumi.get(self, "head_ref")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def owner(self) -> Optional[str]:
        return pulumi.get(self, "owner")

    @property
    @pulumi.getter
    def results(self) -> Sequence['outputs.GetRepositoryPullRequestsResultResult']:
        """
        Collection of Pull Requests matching the filters. Each of the results conforms to the following scheme:
        """
        return pulumi.get(self, "results")

    @property
    @pulumi.getter(name="sortBy")
    def sort_by(self) -> Optional[str]:
        return pulumi.get(self, "sort_by")

    @property
    @pulumi.getter(name="sortDirection")
    def sort_direction(self) -> Optional[str]:
        return pulumi.get(self, "sort_direction")

    @property
    @pulumi.getter
    def state(self) -> Optional[str]:
        """
        the current Pull Request state - can be "open", "closed" or "merged".
        """
        return pulumi.get(self, "state")


class AwaitableGetRepositoryPullRequestsResult(GetRepositoryPullRequestsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRepositoryPullRequestsResult(
            base_ref=self.base_ref,
            base_repository=self.base_repository,
            head_ref=self.head_ref,
            id=self.id,
            owner=self.owner,
            results=self.results,
            sort_by=self.sort_by,
            sort_direction=self.sort_direction,
            state=self.state)


def get_repository_pull_requests(base_ref: Optional[str] = None,
                                 base_repository: Optional[str] = None,
                                 head_ref: Optional[str] = None,
                                 owner: Optional[str] = None,
                                 sort_by: Optional[str] = None,
                                 sort_direction: Optional[str] = None,
                                 state: Optional[str] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRepositoryPullRequestsResult:
    """
    Use this data source to retrieve information about multiple GitHub Pull Requests in a repository.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_github as github

    example = github.get_repository_pull_requests(base_repository="example-repository",
        base_ref="main",
        sort_by="updated",
        sort_direction="desc",
        state="open")
    ```


    :param str base_ref: If set, filters Pull Requests by base branch name.
    :param str base_repository: Name of the base repository to retrieve the Pull Requests from.
    :param str head_ref: If set, filters Pull Requests by head user or head organization and branch name in the format of "user:ref-name" or "organization:ref-name". For example: "github:new-script-format" or "octocat:test-branch".
    :param str owner: Owner of the repository. If not provided, the provider's default owner is used.
    :param str sort_by: If set, indicates what to sort results by. Can be either "created", "updated", "popularity" (comment count) or "long-running" (age, filtering by pulls updated in the last month). Default: "created".
    :param str sort_direction: If set, controls the direction of the sort. Can be either "asc" or "desc". Default: "asc".
    :param str state: If set, filters Pull Requests by state. Can be "open", "closed", or "all". Default: "open".
    """
    __args__ = dict()
    __args__['baseRef'] = base_ref
    __args__['baseRepository'] = base_repository
    __args__['headRef'] = head_ref
    __args__['owner'] = owner
    __args__['sortBy'] = sort_by
    __args__['sortDirection'] = sort_direction
    __args__['state'] = state
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('github:index/getRepositoryPullRequests:getRepositoryPullRequests', __args__, opts=opts, typ=GetRepositoryPullRequestsResult).value

    return AwaitableGetRepositoryPullRequestsResult(
        base_ref=pulumi.get(__ret__, 'base_ref'),
        base_repository=pulumi.get(__ret__, 'base_repository'),
        head_ref=pulumi.get(__ret__, 'head_ref'),
        id=pulumi.get(__ret__, 'id'),
        owner=pulumi.get(__ret__, 'owner'),
        results=pulumi.get(__ret__, 'results'),
        sort_by=pulumi.get(__ret__, 'sort_by'),
        sort_direction=pulumi.get(__ret__, 'sort_direction'),
        state=pulumi.get(__ret__, 'state'))
def get_repository_pull_requests_output(base_ref: Optional[pulumi.Input[Optional[str]]] = None,
                                        base_repository: Optional[pulumi.Input[str]] = None,
                                        head_ref: Optional[pulumi.Input[Optional[str]]] = None,
                                        owner: Optional[pulumi.Input[Optional[str]]] = None,
                                        sort_by: Optional[pulumi.Input[Optional[str]]] = None,
                                        sort_direction: Optional[pulumi.Input[Optional[str]]] = None,
                                        state: Optional[pulumi.Input[Optional[str]]] = None,
                                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetRepositoryPullRequestsResult]:
    """
    Use this data source to retrieve information about multiple GitHub Pull Requests in a repository.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_github as github

    example = github.get_repository_pull_requests(base_repository="example-repository",
        base_ref="main",
        sort_by="updated",
        sort_direction="desc",
        state="open")
    ```


    :param str base_ref: If set, filters Pull Requests by base branch name.
    :param str base_repository: Name of the base repository to retrieve the Pull Requests from.
    :param str head_ref: If set, filters Pull Requests by head user or head organization and branch name in the format of "user:ref-name" or "organization:ref-name". For example: "github:new-script-format" or "octocat:test-branch".
    :param str owner: Owner of the repository. If not provided, the provider's default owner is used.
    :param str sort_by: If set, indicates what to sort results by. Can be either "created", "updated", "popularity" (comment count) or "long-running" (age, filtering by pulls updated in the last month). Default: "created".
    :param str sort_direction: If set, controls the direction of the sort. Can be either "asc" or "desc". Default: "asc".
    :param str state: If set, filters Pull Requests by state. Can be "open", "closed", or "all". Default: "open".
    """
    __args__ = dict()
    __args__['baseRef'] = base_ref
    __args__['baseRepository'] = base_repository
    __args__['headRef'] = head_ref
    __args__['owner'] = owner
    __args__['sortBy'] = sort_by
    __args__['sortDirection'] = sort_direction
    __args__['state'] = state
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('github:index/getRepositoryPullRequests:getRepositoryPullRequests', __args__, opts=opts, typ=GetRepositoryPullRequestsResult)
    return __ret__.apply(lambda __response__: GetRepositoryPullRequestsResult(
        base_ref=pulumi.get(__response__, 'base_ref'),
        base_repository=pulumi.get(__response__, 'base_repository'),
        head_ref=pulumi.get(__response__, 'head_ref'),
        id=pulumi.get(__response__, 'id'),
        owner=pulumi.get(__response__, 'owner'),
        results=pulumi.get(__response__, 'results'),
        sort_by=pulumi.get(__response__, 'sort_by'),
        sort_direction=pulumi.get(__response__, 'sort_direction'),
        state=pulumi.get(__response__, 'state')))
