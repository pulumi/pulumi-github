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
    'GetRepositoryAutolinkReferencesResult',
    'AwaitableGetRepositoryAutolinkReferencesResult',
    'get_repository_autolink_references',
    'get_repository_autolink_references_output',
]

@pulumi.output_type
class GetRepositoryAutolinkReferencesResult:
    """
    A collection of values returned by getRepositoryAutolinkReferences.
    """
    def __init__(__self__, autolink_references=None, id=None, repository=None):
        if autolink_references and not isinstance(autolink_references, list):
            raise TypeError("Expected argument 'autolink_references' to be a list")
        pulumi.set(__self__, "autolink_references", autolink_references)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if repository and not isinstance(repository, str):
            raise TypeError("Expected argument 'repository' to be a str")
        pulumi.set(__self__, "repository", repository)

    @_builtins.property
    @pulumi.getter(name="autolinkReferences")
    def autolink_references(self) -> Sequence['outputs.GetRepositoryAutolinkReferencesAutolinkReferenceResult']:
        """
        The list of this repository's autolink references. Each element of `autolink_references` has the following attributes:
        """
        return pulumi.get(self, "autolink_references")

    @_builtins.property
    @pulumi.getter
    def id(self) -> _builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @_builtins.property
    @pulumi.getter
    def repository(self) -> _builtins.str:
        return pulumi.get(self, "repository")


class AwaitableGetRepositoryAutolinkReferencesResult(GetRepositoryAutolinkReferencesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRepositoryAutolinkReferencesResult(
            autolink_references=self.autolink_references,
            id=self.id,
            repository=self.repository)


def get_repository_autolink_references(repository: Optional[_builtins.str] = None,
                                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRepositoryAutolinkReferencesResult:
    """
    Use this data source to retrieve autolink references for a repository.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_github as github

    example = github.get_repository_autolink_references(repository="example-repository")
    ```


    :param _builtins.str repository: Name of the repository to retrieve the autolink references from.
    """
    __args__ = dict()
    __args__['repository'] = repository
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('github:index/getRepositoryAutolinkReferences:getRepositoryAutolinkReferences', __args__, opts=opts, typ=GetRepositoryAutolinkReferencesResult).value

    return AwaitableGetRepositoryAutolinkReferencesResult(
        autolink_references=pulumi.get(__ret__, 'autolink_references'),
        id=pulumi.get(__ret__, 'id'),
        repository=pulumi.get(__ret__, 'repository'))
def get_repository_autolink_references_output(repository: Optional[pulumi.Input[_builtins.str]] = None,
                                              opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetRepositoryAutolinkReferencesResult]:
    """
    Use this data source to retrieve autolink references for a repository.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_github as github

    example = github.get_repository_autolink_references(repository="example-repository")
    ```


    :param _builtins.str repository: Name of the repository to retrieve the autolink references from.
    """
    __args__ = dict()
    __args__['repository'] = repository
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('github:index/getRepositoryAutolinkReferences:getRepositoryAutolinkReferences', __args__, opts=opts, typ=GetRepositoryAutolinkReferencesResult)
    return __ret__.apply(lambda __response__: GetRepositoryAutolinkReferencesResult(
        autolink_references=pulumi.get(__response__, 'autolink_references'),
        id=pulumi.get(__response__, 'id'),
        repository=pulumi.get(__response__, 'repository')))
