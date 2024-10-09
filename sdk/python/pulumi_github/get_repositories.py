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

__all__ = [
    'GetRepositoriesResult',
    'AwaitableGetRepositoriesResult',
    'get_repositories',
    'get_repositories_output',
]

@pulumi.output_type
class GetRepositoriesResult:
    """
    A collection of values returned by getRepositories.
    """
    def __init__(__self__, full_names=None, id=None, include_repo_id=None, names=None, query=None, repo_ids=None, results_per_page=None, sort=None):
        if full_names and not isinstance(full_names, list):
            raise TypeError("Expected argument 'full_names' to be a list")
        pulumi.set(__self__, "full_names", full_names)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if include_repo_id and not isinstance(include_repo_id, bool):
            raise TypeError("Expected argument 'include_repo_id' to be a bool")
        pulumi.set(__self__, "include_repo_id", include_repo_id)
        if names and not isinstance(names, list):
            raise TypeError("Expected argument 'names' to be a list")
        pulumi.set(__self__, "names", names)
        if query and not isinstance(query, str):
            raise TypeError("Expected argument 'query' to be a str")
        pulumi.set(__self__, "query", query)
        if repo_ids and not isinstance(repo_ids, list):
            raise TypeError("Expected argument 'repo_ids' to be a list")
        pulumi.set(__self__, "repo_ids", repo_ids)
        if results_per_page and not isinstance(results_per_page, int):
            raise TypeError("Expected argument 'results_per_page' to be a int")
        pulumi.set(__self__, "results_per_page", results_per_page)
        if sort and not isinstance(sort, str):
            raise TypeError("Expected argument 'sort' to be a str")
        pulumi.set(__self__, "sort", sort)

    @property
    @pulumi.getter(name="fullNames")
    def full_names(self) -> Sequence[str]:
        return pulumi.get(self, "full_names")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="includeRepoId")
    def include_repo_id(self) -> Optional[bool]:
        return pulumi.get(self, "include_repo_id")

    @property
    @pulumi.getter
    def names(self) -> Sequence[str]:
        return pulumi.get(self, "names")

    @property
    @pulumi.getter
    def query(self) -> str:
        return pulumi.get(self, "query")

    @property
    @pulumi.getter(name="repoIds")
    def repo_ids(self) -> Sequence[int]:
        """
        (Optional) A list of found repository IDs (e.g. `449898861`)
        """
        return pulumi.get(self, "repo_ids")

    @property
    @pulumi.getter(name="resultsPerPage")
    def results_per_page(self) -> Optional[int]:
        return pulumi.get(self, "results_per_page")

    @property
    @pulumi.getter
    def sort(self) -> Optional[str]:
        return pulumi.get(self, "sort")


class AwaitableGetRepositoriesResult(GetRepositoriesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRepositoriesResult(
            full_names=self.full_names,
            id=self.id,
            include_repo_id=self.include_repo_id,
            names=self.names,
            query=self.query,
            repo_ids=self.repo_ids,
            results_per_page=self.results_per_page,
            sort=self.sort)


def get_repositories(include_repo_id: Optional[bool] = None,
                     query: Optional[str] = None,
                     results_per_page: Optional[int] = None,
                     sort: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRepositoriesResult:
    """
    > **Note:** The data source will return a maximum of `1000` repositories
    	[as documented in official API docs](https://developer.github.com/v3/search/#about-the-search-api).

    Use this data source to retrieve a list of GitHub repositories using a search query.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_github as github

    example = github.get_repositories(query="org:hashicorp language:Go",
        include_repo_id=True)
    ```


    :param bool include_repo_id: Returns a list of found repository IDs
    :param str query: Search query. See [documentation for the search syntax](https://help.github.com/articles/understanding-the-search-syntax/).
    :param int results_per_page: Set the number of repositories requested per API call. Can be useful to decrease if requests are timing out or to increase to reduce the number of API calls. Defaults to 100.
    :param str sort: Sorts the repositories returned by the specified attribute. Valid values include `stars`, `fork`, and `updated`. Defaults to `updated`.
    """
    __args__ = dict()
    __args__['includeRepoId'] = include_repo_id
    __args__['query'] = query
    __args__['resultsPerPage'] = results_per_page
    __args__['sort'] = sort
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('github:index/getRepositories:getRepositories', __args__, opts=opts, typ=GetRepositoriesResult).value

    return AwaitableGetRepositoriesResult(
        full_names=pulumi.get(__ret__, 'full_names'),
        id=pulumi.get(__ret__, 'id'),
        include_repo_id=pulumi.get(__ret__, 'include_repo_id'),
        names=pulumi.get(__ret__, 'names'),
        query=pulumi.get(__ret__, 'query'),
        repo_ids=pulumi.get(__ret__, 'repo_ids'),
        results_per_page=pulumi.get(__ret__, 'results_per_page'),
        sort=pulumi.get(__ret__, 'sort'))
def get_repositories_output(include_repo_id: Optional[pulumi.Input[Optional[bool]]] = None,
                            query: Optional[pulumi.Input[str]] = None,
                            results_per_page: Optional[pulumi.Input[Optional[int]]] = None,
                            sort: Optional[pulumi.Input[Optional[str]]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetRepositoriesResult]:
    """
    > **Note:** The data source will return a maximum of `1000` repositories
    	[as documented in official API docs](https://developer.github.com/v3/search/#about-the-search-api).

    Use this data source to retrieve a list of GitHub repositories using a search query.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_github as github

    example = github.get_repositories(query="org:hashicorp language:Go",
        include_repo_id=True)
    ```


    :param bool include_repo_id: Returns a list of found repository IDs
    :param str query: Search query. See [documentation for the search syntax](https://help.github.com/articles/understanding-the-search-syntax/).
    :param int results_per_page: Set the number of repositories requested per API call. Can be useful to decrease if requests are timing out or to increase to reduce the number of API calls. Defaults to 100.
    :param str sort: Sorts the repositories returned by the specified attribute. Valid values include `stars`, `fork`, and `updated`. Defaults to `updated`.
    """
    __args__ = dict()
    __args__['includeRepoId'] = include_repo_id
    __args__['query'] = query
    __args__['resultsPerPage'] = results_per_page
    __args__['sort'] = sort
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('github:index/getRepositories:getRepositories', __args__, opts=opts, typ=GetRepositoriesResult)
    return __ret__.apply(lambda __response__: GetRepositoriesResult(
        full_names=pulumi.get(__response__, 'full_names'),
        id=pulumi.get(__response__, 'id'),
        include_repo_id=pulumi.get(__response__, 'include_repo_id'),
        names=pulumi.get(__response__, 'names'),
        query=pulumi.get(__response__, 'query'),
        repo_ids=pulumi.get(__response__, 'repo_ids'),
        results_per_page=pulumi.get(__response__, 'results_per_page'),
        sort=pulumi.get(__response__, 'sort')))
