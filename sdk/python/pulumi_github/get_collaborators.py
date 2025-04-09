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
from . import outputs

__all__ = [
    'GetCollaboratorsResult',
    'AwaitableGetCollaboratorsResult',
    'get_collaborators',
    'get_collaborators_output',
]

@pulumi.output_type
class GetCollaboratorsResult:
    """
    A collection of values returned by getCollaborators.
    """
    def __init__(__self__, affiliation=None, collaborators=None, id=None, owner=None, permission=None, repository=None):
        if affiliation and not isinstance(affiliation, str):
            raise TypeError("Expected argument 'affiliation' to be a str")
        pulumi.set(__self__, "affiliation", affiliation)
        if collaborators and not isinstance(collaborators, list):
            raise TypeError("Expected argument 'collaborators' to be a list")
        pulumi.set(__self__, "collaborators", collaborators)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if owner and not isinstance(owner, str):
            raise TypeError("Expected argument 'owner' to be a str")
        pulumi.set(__self__, "owner", owner)
        if permission and not isinstance(permission, str):
            raise TypeError("Expected argument 'permission' to be a str")
        pulumi.set(__self__, "permission", permission)
        if repository and not isinstance(repository, str):
            raise TypeError("Expected argument 'repository' to be a str")
        pulumi.set(__self__, "repository", repository)

    @property
    @pulumi.getter
    def affiliation(self) -> Optional[builtins.str]:
        return pulumi.get(self, "affiliation")

    @property
    @pulumi.getter
    def collaborators(self) -> Sequence['outputs.GetCollaboratorsCollaboratorResult']:
        """
        An Array of GitHub collaborators.  Each `collaborator` block consists of the fields documented below.
        """
        return pulumi.get(self, "collaborators")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def owner(self) -> builtins.str:
        return pulumi.get(self, "owner")

    @property
    @pulumi.getter
    def permission(self) -> Optional[builtins.str]:
        """
        The permission of the collaborator.
        """
        return pulumi.get(self, "permission")

    @property
    @pulumi.getter
    def repository(self) -> builtins.str:
        return pulumi.get(self, "repository")


class AwaitableGetCollaboratorsResult(GetCollaboratorsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCollaboratorsResult(
            affiliation=self.affiliation,
            collaborators=self.collaborators,
            id=self.id,
            owner=self.owner,
            permission=self.permission,
            repository=self.repository)


def get_collaborators(affiliation: Optional[builtins.str] = None,
                      owner: Optional[builtins.str] = None,
                      permission: Optional[builtins.str] = None,
                      repository: Optional[builtins.str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetCollaboratorsResult:
    """
    Use this data source to retrieve the collaborators for a given repository.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_github as github

    test = github.get_collaborators(owner="example_owner",
        repository="example_repository")
    ```


    :param builtins.str affiliation: Filter collaborators returned by their affiliation. Can be one of: `outside`, `direct`, `all`.  Defaults to `all`.
    :param builtins.str owner: The organization that owns the repository.
    :param builtins.str permission: Filter collaborators returned by their permission. Can be one of: `pull`, `triage`, `push`, `maintain`, `admin`.  Defaults to not doing any filtering on permission.
    :param builtins.str repository: The name of the repository.
    """
    __args__ = dict()
    __args__['affiliation'] = affiliation
    __args__['owner'] = owner
    __args__['permission'] = permission
    __args__['repository'] = repository
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('github:index/getCollaborators:getCollaborators', __args__, opts=opts, typ=GetCollaboratorsResult).value

    return AwaitableGetCollaboratorsResult(
        affiliation=pulumi.get(__ret__, 'affiliation'),
        collaborators=pulumi.get(__ret__, 'collaborators'),
        id=pulumi.get(__ret__, 'id'),
        owner=pulumi.get(__ret__, 'owner'),
        permission=pulumi.get(__ret__, 'permission'),
        repository=pulumi.get(__ret__, 'repository'))
def get_collaborators_output(affiliation: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                             owner: Optional[pulumi.Input[builtins.str]] = None,
                             permission: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                             repository: Optional[pulumi.Input[builtins.str]] = None,
                             opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetCollaboratorsResult]:
    """
    Use this data source to retrieve the collaborators for a given repository.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_github as github

    test = github.get_collaborators(owner="example_owner",
        repository="example_repository")
    ```


    :param builtins.str affiliation: Filter collaborators returned by their affiliation. Can be one of: `outside`, `direct`, `all`.  Defaults to `all`.
    :param builtins.str owner: The organization that owns the repository.
    :param builtins.str permission: Filter collaborators returned by their permission. Can be one of: `pull`, `triage`, `push`, `maintain`, `admin`.  Defaults to not doing any filtering on permission.
    :param builtins.str repository: The name of the repository.
    """
    __args__ = dict()
    __args__['affiliation'] = affiliation
    __args__['owner'] = owner
    __args__['permission'] = permission
    __args__['repository'] = repository
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('github:index/getCollaborators:getCollaborators', __args__, opts=opts, typ=GetCollaboratorsResult)
    return __ret__.apply(lambda __response__: GetCollaboratorsResult(
        affiliation=pulumi.get(__response__, 'affiliation'),
        collaborators=pulumi.get(__response__, 'collaborators'),
        id=pulumi.get(__response__, 'id'),
        owner=pulumi.get(__response__, 'owner'),
        permission=pulumi.get(__response__, 'permission'),
        repository=pulumi.get(__response__, 'repository')))
