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
    'GetSshKeysResult',
    'AwaitableGetSshKeysResult',
    'get_ssh_keys',
    'get_ssh_keys_output',
]

@pulumi.output_type
class GetSshKeysResult:
    """
    A collection of values returned by getSshKeys.
    """
    def __init__(__self__, id=None, keys=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if keys and not isinstance(keys, list):
            raise TypeError("Expected argument 'keys' to be a list")
        pulumi.set(__self__, "keys", keys)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def keys(self) -> Sequence[str]:
        """
        An array of GitHub's SSH public keys.
        """
        return pulumi.get(self, "keys")


class AwaitableGetSshKeysResult(GetSshKeysResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSshKeysResult(
            id=self.id,
            keys=self.keys)


def get_ssh_keys(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSshKeysResult:
    """
    Use this data source to retrieve information about GitHub's SSH keys.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_github as github

    test = github.get_ssh_keys()
    ```
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('github:index/getSshKeys:getSshKeys', __args__, opts=opts, typ=GetSshKeysResult).value

    return AwaitableGetSshKeysResult(
        id=pulumi.get(__ret__, 'id'),
        keys=pulumi.get(__ret__, 'keys'))
def get_ssh_keys_output(opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSshKeysResult]:
    """
    Use this data source to retrieve information about GitHub's SSH keys.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_github as github

    test = github.get_ssh_keys()
    ```
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('github:index/getSshKeys:getSshKeys', __args__, opts=opts, typ=GetSshKeysResult)
    return __ret__.apply(lambda __response__: GetSshKeysResult(
        id=pulumi.get(__response__, 'id'),
        keys=pulumi.get(__response__, 'keys')))
