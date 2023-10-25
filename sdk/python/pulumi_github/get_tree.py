# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs

__all__ = [
    'GetTreeResult',
    'AwaitableGetTreeResult',
    'get_tree',
    'get_tree_output',
]

@pulumi.output_type
class GetTreeResult:
    """
    A collection of values returned by getTree.
    """
    def __init__(__self__, entries=None, id=None, recursive=None, repository=None, tree_sha=None):
        if entries and not isinstance(entries, list):
            raise TypeError("Expected argument 'entries' to be a list")
        pulumi.set(__self__, "entries", entries)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if recursive and not isinstance(recursive, bool):
            raise TypeError("Expected argument 'recursive' to be a bool")
        pulumi.set(__self__, "recursive", recursive)
        if repository and not isinstance(repository, str):
            raise TypeError("Expected argument 'repository' to be a str")
        pulumi.set(__self__, "repository", repository)
        if tree_sha and not isinstance(tree_sha, str):
            raise TypeError("Expected argument 'tree_sha' to be a str")
        pulumi.set(__self__, "tree_sha", tree_sha)

    @property
    @pulumi.getter
    def entries(self) -> Sequence['outputs.GetTreeEntryResult']:
        """
        Objects (of `path`, `mode`, `type`, `size`, and `sha`) specifying a tree structure.
        """
        return pulumi.get(self, "entries")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def recursive(self) -> Optional[bool]:
        return pulumi.get(self, "recursive")

    @property
    @pulumi.getter
    def repository(self) -> str:
        return pulumi.get(self, "repository")

    @property
    @pulumi.getter(name="treeSha")
    def tree_sha(self) -> str:
        return pulumi.get(self, "tree_sha")


class AwaitableGetTreeResult(GetTreeResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTreeResult(
            entries=self.entries,
            id=self.id,
            recursive=self.recursive,
            repository=self.repository,
            tree_sha=self.tree_sha)


def get_tree(recursive: Optional[bool] = None,
             repository: Optional[str] = None,
             tree_sha: Optional[str] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTreeResult:
    """
    Use this data source to retrieve information about a single tree.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_github as github

    this_repository = github.get_repository(name="example")
    this_branch = github.get_branch(branch=this_repository.default_branch,
        repository=this_repository.name)
    this_tree = github.get_tree(recursive=False,
        repository=this_repository.name,
        tree_sha=this_branch.sha)
    pulumi.export("entries", this_tree.entries)
    ```


    :param bool recursive: Setting this parameter to `true` returns the objects or subtrees referenced by the tree specified in `tree_sha`.
    :param str repository: The name of the repository.
    :param str tree_sha: The SHA1 value for the tree.
    """
    __args__ = dict()
    __args__['recursive'] = recursive
    __args__['repository'] = repository
    __args__['treeSha'] = tree_sha
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('github:index/getTree:getTree', __args__, opts=opts, typ=GetTreeResult).value

    return AwaitableGetTreeResult(
        entries=pulumi.get(__ret__, 'entries'),
        id=pulumi.get(__ret__, 'id'),
        recursive=pulumi.get(__ret__, 'recursive'),
        repository=pulumi.get(__ret__, 'repository'),
        tree_sha=pulumi.get(__ret__, 'tree_sha'))


@_utilities.lift_output_func(get_tree)
def get_tree_output(recursive: Optional[pulumi.Input[Optional[bool]]] = None,
                    repository: Optional[pulumi.Input[str]] = None,
                    tree_sha: Optional[pulumi.Input[str]] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetTreeResult]:
    """
    Use this data source to retrieve information about a single tree.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_github as github

    this_repository = github.get_repository(name="example")
    this_branch = github.get_branch(branch=this_repository.default_branch,
        repository=this_repository.name)
    this_tree = github.get_tree(recursive=False,
        repository=this_repository.name,
        tree_sha=this_branch.sha)
    pulumi.export("entries", this_tree.entries)
    ```


    :param bool recursive: Setting this parameter to `true` returns the objects or subtrees referenced by the tree specified in `tree_sha`.
    :param str repository: The name of the repository.
    :param str tree_sha: The SHA1 value for the tree.
    """
    ...
