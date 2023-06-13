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
    'GetRepositoryEnvironmentsResult',
    'AwaitableGetRepositoryEnvironmentsResult',
    'get_repository_environments',
    'get_repository_environments_output',
]

@pulumi.output_type
class GetRepositoryEnvironmentsResult:
    """
    A collection of values returned by getRepositoryEnvironments.
    """
    def __init__(__self__, environments=None, id=None, repository=None):
        if environments and not isinstance(environments, list):
            raise TypeError("Expected argument 'environments' to be a list")
        pulumi.set(__self__, "environments", environments)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if repository and not isinstance(repository, str):
            raise TypeError("Expected argument 'repository' to be a str")
        pulumi.set(__self__, "repository", repository)

    @property
    @pulumi.getter
    def environments(self) -> Sequence['outputs.GetRepositoryEnvironmentsEnvironmentResult']:
        """
        The list of this repository's environments. Each element of `environments` has the following attributes:
        """
        return pulumi.get(self, "environments")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def repository(self) -> str:
        return pulumi.get(self, "repository")


class AwaitableGetRepositoryEnvironmentsResult(GetRepositoryEnvironmentsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRepositoryEnvironmentsResult(
            environments=self.environments,
            id=self.id,
            repository=self.repository)


def get_repository_environments(repository: Optional[str] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRepositoryEnvironmentsResult:
    """
    Use this data source to retrieve information about environments for a repository.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_github as github

    example = github.get_repository_environments(repository="example-repository")
    ```


    :param str repository: Name of the repository to retrieve the environments from.
    """
    __args__ = dict()
    __args__['repository'] = repository
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('github:index/getRepositoryEnvironments:getRepositoryEnvironments', __args__, opts=opts, typ=GetRepositoryEnvironmentsResult).value

    return AwaitableGetRepositoryEnvironmentsResult(
        environments=__ret__.environments,
        id=__ret__.id,
        repository=__ret__.repository)


@_utilities.lift_output_func(get_repository_environments)
def get_repository_environments_output(repository: Optional[pulumi.Input[str]] = None,
                                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetRepositoryEnvironmentsResult]:
    """
    Use this data source to retrieve information about environments for a repository.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_github as github

    example = github.get_repository_environments(repository="example-repository")
    ```


    :param str repository: Name of the repository to retrieve the environments from.
    """
    ...