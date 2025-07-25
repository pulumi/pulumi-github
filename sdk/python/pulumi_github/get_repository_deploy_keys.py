# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins as _builtins
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
    'GetRepositoryDeployKeysResult',
    'AwaitableGetRepositoryDeployKeysResult',
    'get_repository_deploy_keys',
    'get_repository_deploy_keys_output',
]

@pulumi.output_type
class GetRepositoryDeployKeysResult:
    """
    A collection of values returned by getRepositoryDeployKeys.
    """
    def __init__(__self__, id=None, keys=None, repository=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if keys and not isinstance(keys, list):
            raise TypeError("Expected argument 'keys' to be a list")
        pulumi.set(__self__, "keys", keys)
        if repository and not isinstance(repository, str):
            raise TypeError("Expected argument 'repository' to be a str")
        pulumi.set(__self__, "repository", repository)

    @_builtins.property
    @pulumi.getter
    def id(self) -> _builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @_builtins.property
    @pulumi.getter
    def keys(self) -> Sequence['outputs.GetRepositoryDeployKeysKeyResult']:
        """
        The list of this repository's deploy keys. Each element of `keys` has the following attributes:
        """
        return pulumi.get(self, "keys")

    @_builtins.property
    @pulumi.getter
    def repository(self) -> _builtins.str:
        return pulumi.get(self, "repository")


class AwaitableGetRepositoryDeployKeysResult(GetRepositoryDeployKeysResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRepositoryDeployKeysResult(
            id=self.id,
            keys=self.keys,
            repository=self.repository)


def get_repository_deploy_keys(repository: Optional[_builtins.str] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRepositoryDeployKeysResult:
    """
    Use this data source to retrieve all deploy keys of a repository.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_github as github

    example = github.get_repository_deploy_keys(repository="example-repository")
    ```


    :param _builtins.str repository: Name of the repository to retrieve the branches from.
    """
    __args__ = dict()
    __args__['repository'] = repository
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('github:index/getRepositoryDeployKeys:getRepositoryDeployKeys', __args__, opts=opts, typ=GetRepositoryDeployKeysResult).value

    return AwaitableGetRepositoryDeployKeysResult(
        id=pulumi.get(__ret__, 'id'),
        keys=pulumi.get(__ret__, 'keys'),
        repository=pulumi.get(__ret__, 'repository'))
def get_repository_deploy_keys_output(repository: Optional[pulumi.Input[_builtins.str]] = None,
                                      opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetRepositoryDeployKeysResult]:
    """
    Use this data source to retrieve all deploy keys of a repository.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_github as github

    example = github.get_repository_deploy_keys(repository="example-repository")
    ```


    :param _builtins.str repository: Name of the repository to retrieve the branches from.
    """
    __args__ = dict()
    __args__['repository'] = repository
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('github:index/getRepositoryDeployKeys:getRepositoryDeployKeys', __args__, opts=opts, typ=GetRepositoryDeployKeysResult)
    return __ret__.apply(lambda __response__: GetRepositoryDeployKeysResult(
        id=pulumi.get(__response__, 'id'),
        keys=pulumi.get(__response__, 'keys'),
        repository=pulumi.get(__response__, 'repository')))
