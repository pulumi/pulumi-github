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
    'GetActionsPublicKeyResult',
    'AwaitableGetActionsPublicKeyResult',
    'get_actions_public_key',
    'get_actions_public_key_output',
]

@pulumi.output_type
class GetActionsPublicKeyResult:
    """
    A collection of values returned by getActionsPublicKey.
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
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def key(self) -> str:
        """
        Actual key retrieved.
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter(name="keyId")
    def key_id(self) -> str:
        """
        ID of the key that has been retrieved.
        """
        return pulumi.get(self, "key_id")

    @property
    @pulumi.getter
    def repository(self) -> str:
        return pulumi.get(self, "repository")


class AwaitableGetActionsPublicKeyResult(GetActionsPublicKeyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetActionsPublicKeyResult(
            id=self.id,
            key=self.key,
            key_id=self.key_id,
            repository=self.repository)


def get_actions_public_key(repository: Optional[str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetActionsPublicKeyResult:
    """
    Use this data source to retrieve information about a GitHub Actions public key. This data source is required to be used with other GitHub secrets interactions.
    Note that the provider `token` must have admin rights to a repository to retrieve it's action public key.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_github as github

    example = github.get_actions_public_key(repository="example_repo")
    ```
    <!--End PulumiCodeChooser -->


    :param str repository: Name of the repository to get public key from.
    """
    __args__ = dict()
    __args__['repository'] = repository
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('github:index/getActionsPublicKey:getActionsPublicKey', __args__, opts=opts, typ=GetActionsPublicKeyResult).value

    return AwaitableGetActionsPublicKeyResult(
        id=pulumi.get(__ret__, 'id'),
        key=pulumi.get(__ret__, 'key'),
        key_id=pulumi.get(__ret__, 'key_id'),
        repository=pulumi.get(__ret__, 'repository'))


@_utilities.lift_output_func(get_actions_public_key)
def get_actions_public_key_output(repository: Optional[pulumi.Input[str]] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetActionsPublicKeyResult]:
    """
    Use this data source to retrieve information about a GitHub Actions public key. This data source is required to be used with other GitHub secrets interactions.
    Note that the provider `token` must have admin rights to a repository to retrieve it's action public key.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_github as github

    example = github.get_actions_public_key(repository="example_repo")
    ```
    <!--End PulumiCodeChooser -->


    :param str repository: Name of the repository to get public key from.
    """
    ...
