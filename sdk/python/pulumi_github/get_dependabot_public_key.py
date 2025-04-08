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

__all__ = [
    'GetDependabotPublicKeyResult',
    'AwaitableGetDependabotPublicKeyResult',
    'get_dependabot_public_key',
    'get_dependabot_public_key_output',
]

@pulumi.output_type
class GetDependabotPublicKeyResult:
    """
    A collection of values returned by getDependabotPublicKey.
    """
    def __init__(__self__, id=None, key=None, key_id=None, repository=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if key and not isinstance(key, str):
            raise TypeError("Expected argument 'key' to be a str")
        pulumi.set(__self__, "key", key)
        if key_id and not isinstance(key_id, str):
            raise TypeError("Expected argument 'key_id' to be a str")
        pulumi.set(__self__, "key_id", key_id)
        if repository and not isinstance(repository, str):
            raise TypeError("Expected argument 'repository' to be a str")
        pulumi.set(__self__, "repository", repository)

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def key(self) -> builtins.str:
        """
        Actual key retrieved.
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter(name="keyId")
    def key_id(self) -> builtins.str:
        """
        ID of the key that has been retrieved.
        """
        return pulumi.get(self, "key_id")

    @property
    @pulumi.getter
    def repository(self) -> builtins.str:
        return pulumi.get(self, "repository")


class AwaitableGetDependabotPublicKeyResult(GetDependabotPublicKeyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDependabotPublicKeyResult(
            id=self.id,
            key=self.key,
            key_id=self.key_id,
            repository=self.repository)


def get_dependabot_public_key(repository: Optional[builtins.str] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDependabotPublicKeyResult:
    """
    Use this data source to retrieve information about a GitHub Dependabot public key. This data source is required to be used with other GitHub secrets interactions.
    Note that the provider `token` must have admin rights to a repository to retrieve it's Dependabot public key.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_github as github

    example = github.get_dependabot_public_key(repository="example_repo")
    ```


    :param builtins.str repository: Name of the repository to get public key from.
    """
    __args__ = dict()
    __args__['repository'] = repository
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('github:index/getDependabotPublicKey:getDependabotPublicKey', __args__, opts=opts, typ=GetDependabotPublicKeyResult).value

    return AwaitableGetDependabotPublicKeyResult(
        id=pulumi.get(__ret__, 'id'),
        key=pulumi.get(__ret__, 'key'),
        key_id=pulumi.get(__ret__, 'key_id'),
        repository=pulumi.get(__ret__, 'repository'))
def get_dependabot_public_key_output(repository: Optional[pulumi.Input[builtins.str]] = None,
                                     opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetDependabotPublicKeyResult]:
    """
    Use this data source to retrieve information about a GitHub Dependabot public key. This data source is required to be used with other GitHub secrets interactions.
    Note that the provider `token` must have admin rights to a repository to retrieve it's Dependabot public key.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_github as github

    example = github.get_dependabot_public_key(repository="example_repo")
    ```


    :param builtins.str repository: Name of the repository to get public key from.
    """
    __args__ = dict()
    __args__['repository'] = repository
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('github:index/getDependabotPublicKey:getDependabotPublicKey', __args__, opts=opts, typ=GetDependabotPublicKeyResult)
    return __ret__.apply(lambda __response__: GetDependabotPublicKeyResult(
        id=pulumi.get(__response__, 'id'),
        key=pulumi.get(__response__, 'key'),
        key_id=pulumi.get(__response__, 'key_id'),
        repository=pulumi.get(__response__, 'repository')))
