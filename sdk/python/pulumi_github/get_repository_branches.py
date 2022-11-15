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
    'GetRepositoryBranchesResult',
    'AwaitableGetRepositoryBranchesResult',
    'get_repository_branches',
    'get_repository_branches_output',
]

@pulumi.output_type
class GetRepositoryBranchesResult:
    """
    A collection of values returned by getRepositoryBranches.
    """
    def __init__(__self__, branches=None, id=None, repository=None):
        if branches and not isinstance(branches, list):
            raise TypeError("Expected argument 'branches' to be a list")
        pulumi.set(__self__, "branches", branches)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if repository and not isinstance(repository, str):
            raise TypeError("Expected argument 'repository' to be a str")
        pulumi.set(__self__, "repository", repository)

    @property
    @pulumi.getter
    def branches(self) -> Sequence['outputs.GetRepositoryBranchesBranchResult']:
        return pulumi.get(self, "branches")

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


class AwaitableGetRepositoryBranchesResult(GetRepositoryBranchesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRepositoryBranchesResult(
            branches=self.branches,
            id=self.id,
            repository=self.repository)


def get_repository_branches(repository: Optional[str] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRepositoryBranchesResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['repository'] = repository
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('github:index/getRepositoryBranches:getRepositoryBranches', __args__, opts=opts, typ=GetRepositoryBranchesResult).value

    return AwaitableGetRepositoryBranchesResult(
        branches=__ret__.branches,
        id=__ret__.id,
        repository=__ret__.repository)


@_utilities.lift_output_func(get_repository_branches)
def get_repository_branches_output(repository: Optional[pulumi.Input[str]] = None,
                                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetRepositoryBranchesResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
