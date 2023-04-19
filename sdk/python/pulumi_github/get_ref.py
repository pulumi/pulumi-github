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
    'GetRefResult',
    'AwaitableGetRefResult',
    'get_ref',
    'get_ref_output',
]

@pulumi.output_type
class GetRefResult:
    """
    A collection of values returned by getRef.
    """
    def __init__(__self__, etag=None, id=None, owner=None, ref=None, repository=None, sha=None):
        if etag and not isinstance(etag, str):
            raise TypeError("Expected argument 'etag' to be a str")
        pulumi.set(__self__, "etag", etag)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if owner and not isinstance(owner, str):
            raise TypeError("Expected argument 'owner' to be a str")
        pulumi.set(__self__, "owner", owner)
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
    def etag(self) -> str:
        """
        An etag representing the ref.
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
    def owner(self) -> Optional[str]:
        return pulumi.get(self, "owner")

    @property
    @pulumi.getter
    def ref(self) -> str:
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


class AwaitableGetRefResult(GetRefResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRefResult(
            etag=self.etag,
            id=self.id,
            owner=self.owner,
            ref=self.ref,
            repository=self.repository,
            sha=self.sha)


def get_ref(owner: Optional[str] = None,
            ref: Optional[str] = None,
            repository: Optional[str] = None,
            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRefResult:
    """
    Use this data source to retrieve information about a repository ref.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_github as github

    development = github.get_ref(owner="example",
        ref="heads/development",
        repository="example")
    ```


    :param str owner: Owner of the repository.
    :param str ref: The repository ref to look up. Must be formatted `heads/<ref>` for branches, and `tags/<ref>` for tags.
    :param str repository: The GitHub repository name.
    """
    __args__ = dict()
    __args__['owner'] = owner
    __args__['ref'] = ref
    __args__['repository'] = repository
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('github:index/getRef:getRef', __args__, opts=opts, typ=GetRefResult).value

    return AwaitableGetRefResult(
        etag=__ret__.etag,
        id=__ret__.id,
        owner=__ret__.owner,
        ref=__ret__.ref,
        repository=__ret__.repository,
        sha=__ret__.sha)


@_utilities.lift_output_func(get_ref)
def get_ref_output(owner: Optional[pulumi.Input[Optional[str]]] = None,
                   ref: Optional[pulumi.Input[str]] = None,
                   repository: Optional[pulumi.Input[str]] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetRefResult]:
    """
    Use this data source to retrieve information about a repository ref.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_github as github

    development = github.get_ref(owner="example",
        ref="heads/development",
        repository="example")
    ```


    :param str owner: Owner of the repository.
    :param str ref: The repository ref to look up. Must be formatted `heads/<ref>` for branches, and `tags/<ref>` for tags.
    :param str repository: The GitHub repository name.
    """
    ...
