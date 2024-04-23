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
    'GetBranchResult',
    'AwaitableGetBranchResult',
    'get_branch',
    'get_branch_output',
]

@pulumi.output_type
class GetBranchResult:
    """
    A collection of values returned by getBranch.
    """
    def __init__(__self__, branch=None, etag=None, id=None, ref=None, repository=None, sha=None):
        if branch and not isinstance(branch, str):
            raise TypeError("Expected argument 'branch' to be a str")
        pulumi.set(__self__, "branch", branch)
        if etag and not isinstance(etag, str):
            raise TypeError("Expected argument 'etag' to be a str")
        pulumi.set(__self__, "etag", etag)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ref and not isinstance(ref, str):
            raise TypeError("Expected argument 'ref' to be a str")
        pulumi.set(__self__, "ref", ref)
        if repository and not isinstance(repository, str):
            raise TypeError("Expected argument 'repository' to be a str")
        pulumi.set(__self__, "repository", repository)
        if sha and not isinstance(sha, str):
            raise TypeError("Expected argument 'sha' to be a str")
        pulumi.set(__self__, "sha", sha)

    @property
    @pulumi.getter
    def branch(self) -> str:
        return pulumi.get(self, "branch")

    @property
    @pulumi.getter
    def etag(self) -> str:
        """
        An etag representing the Branch object.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def ref(self) -> str:
        """
        A string representing a branch reference, in the form of `refs/heads/<branch>`.
        """
        return pulumi.get(self, "ref")

    @property
    @pulumi.getter
    def repository(self) -> str:
        return pulumi.get(self, "repository")

    @property
    @pulumi.getter
    def sha(self) -> str:
        """
        A string storing the reference's `HEAD` commit's SHA1.
        """
        return pulumi.get(self, "sha")


class AwaitableGetBranchResult(GetBranchResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetBranchResult(
            branch=self.branch,
            etag=self.etag,
            id=self.id,
            ref=self.ref,
            repository=self.repository,
            sha=self.sha)


def get_branch(branch: Optional[str] = None,
               repository: Optional[str] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetBranchResult:
    """
    Use this data source to retrieve information about a repository branch.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_github as github

    development = github.get_branch(repository="example",
        branch="development")
    ```


    :param str branch: The repository branch to retrieve.
    :param str repository: The GitHub repository name.
    """
    __args__ = dict()
    __args__['branch'] = branch
    __args__['repository'] = repository
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('github:index/getBranch:getBranch', __args__, opts=opts, typ=GetBranchResult).value

    return AwaitableGetBranchResult(
        branch=pulumi.get(__ret__, 'branch'),
        etag=pulumi.get(__ret__, 'etag'),
        id=pulumi.get(__ret__, 'id'),
        ref=pulumi.get(__ret__, 'ref'),
        repository=pulumi.get(__ret__, 'repository'),
        sha=pulumi.get(__ret__, 'sha'))


@_utilities.lift_output_func(get_branch)
def get_branch_output(branch: Optional[pulumi.Input[str]] = None,
                      repository: Optional[pulumi.Input[str]] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetBranchResult]:
    """
    Use this data source to retrieve information about a repository branch.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_github as github

    development = github.get_branch(repository="example",
        branch="development")
    ```


    :param str branch: The repository branch to retrieve.
    :param str repository: The GitHub repository name.
    """
    ...
