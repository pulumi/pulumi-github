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

    @property
    @pulumi.getter(name="autolinkReferences")
    def autolink_references(self) -> Sequence['outputs.GetRepositoryAutolinkReferencesAutolinkReferenceResult']:
        """
        The list of this repository's autolink references. Each element of `autolink_references` has the following attributes:
        """
        return pulumi.get(self, "autolink_references")

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


class AwaitableGetRepositoryAutolinkReferencesResult(GetRepositoryAutolinkReferencesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRepositoryAutolinkReferencesResult(
            autolink_references=self.autolink_references,
            id=self.id,
            repository=self.repository)


def get_repository_autolink_references(repository: Optional[str] = None,
                                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRepositoryAutolinkReferencesResult:
    """
    Use this data source to retrieve autolink references for a repository.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_github as github

    example = github.get_repository_autolink_references(repository="example-repository")
    ```
    <!--End PulumiCodeChooser -->


    :param str repository: Name of the repository to retrieve the autolink references from.
    """
    __args__ = dict()
    __args__['repository'] = repository
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('github:index/getRepositoryAutolinkReferences:getRepositoryAutolinkReferences', __args__, opts=opts, typ=GetRepositoryAutolinkReferencesResult).value

    return AwaitableGetRepositoryAutolinkReferencesResult(
        autolink_references=pulumi.get(__ret__, 'autolink_references'),
        id=pulumi.get(__ret__, 'id'),
        repository=pulumi.get(__ret__, 'repository'))


@_utilities.lift_output_func(get_repository_autolink_references)
def get_repository_autolink_references_output(repository: Optional[pulumi.Input[str]] = None,
                                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetRepositoryAutolinkReferencesResult]:
    """
    Use this data source to retrieve autolink references for a repository.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_github as github

    example = github.get_repository_autolink_references(repository="example-repository")
    ```
    <!--End PulumiCodeChooser -->


    :param str repository: Name of the repository to retrieve the autolink references from.
    """
    ...
