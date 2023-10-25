# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetRepositoryPullRequestResult',
    'AwaitableGetRepositoryPullRequestResult',
    'get_repository_pull_request',
    'get_repository_pull_request_output',
]

@pulumi.output_type
class GetRepositoryPullRequestResult:
    """
    A collection of values returned by getRepositoryPullRequest.
    """
    def __init__(__self__, base_ref=None, base_repository=None, base_sha=None, body=None, draft=None, head_owner=None, head_ref=None, head_repository=None, head_sha=None, id=None, labels=None, maintainer_can_modify=None, number=None, opened_at=None, opened_by=None, owner=None, state=None, title=None, updated_at=None):
        if base_ref and not isinstance(base_ref, str):
            raise TypeError("Expected argument 'base_ref' to be a str")
        pulumi.set(__self__, "base_ref", base_ref)
        if base_repository and not isinstance(base_repository, str):
            raise TypeError("Expected argument 'base_repository' to be a str")
        pulumi.set(__self__, "base_repository", base_repository)
        if base_sha and not isinstance(base_sha, str):
            raise TypeError("Expected argument 'base_sha' to be a str")
        pulumi.set(__self__, "base_sha", base_sha)
        if body and not isinstance(body, str):
            raise TypeError("Expected argument 'body' to be a str")
        pulumi.set(__self__, "body", body)
        if draft and not isinstance(draft, bool):
            raise TypeError("Expected argument 'draft' to be a bool")
        pulumi.set(__self__, "draft", draft)
        if head_owner and not isinstance(head_owner, str):
            raise TypeError("Expected argument 'head_owner' to be a str")
        pulumi.set(__self__, "head_owner", head_owner)
        if head_ref and not isinstance(head_ref, str):
            raise TypeError("Expected argument 'head_ref' to be a str")
        pulumi.set(__self__, "head_ref", head_ref)
        if head_repository and not isinstance(head_repository, str):
            raise TypeError("Expected argument 'head_repository' to be a str")
        pulumi.set(__self__, "head_repository", head_repository)
        if head_sha and not isinstance(head_sha, str):
            raise TypeError("Expected argument 'head_sha' to be a str")
        pulumi.set(__self__, "head_sha", head_sha)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if labels and not isinstance(labels, list):
            raise TypeError("Expected argument 'labels' to be a list")
        pulumi.set(__self__, "labels", labels)
        if maintainer_can_modify and not isinstance(maintainer_can_modify, bool):
            raise TypeError("Expected argument 'maintainer_can_modify' to be a bool")
        pulumi.set(__self__, "maintainer_can_modify", maintainer_can_modify)
        if number and not isinstance(number, int):
            raise TypeError("Expected argument 'number' to be a int")
        pulumi.set(__self__, "number", number)
        if opened_at and not isinstance(opened_at, int):
            raise TypeError("Expected argument 'opened_at' to be a int")
        pulumi.set(__self__, "opened_at", opened_at)
        if opened_by and not isinstance(opened_by, str):
            raise TypeError("Expected argument 'opened_by' to be a str")
        pulumi.set(__self__, "opened_by", opened_by)
        if owner and not isinstance(owner, str):
            raise TypeError("Expected argument 'owner' to be a str")
        pulumi.set(__self__, "owner", owner)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if title and not isinstance(title, str):
            raise TypeError("Expected argument 'title' to be a str")
        pulumi.set(__self__, "title", title)
        if updated_at and not isinstance(updated_at, int):
            raise TypeError("Expected argument 'updated_at' to be a int")
        pulumi.set(__self__, "updated_at", updated_at)

    @property
    @pulumi.getter(name="baseRef")
    def base_ref(self) -> str:
        """
        Name of the ref (branch) of the Pull Request base.
        """
        return pulumi.get(self, "base_ref")

    @property
    @pulumi.getter(name="baseRepository")
    def base_repository(self) -> str:
        return pulumi.get(self, "base_repository")

    @property
    @pulumi.getter(name="baseSha")
    def base_sha(self) -> str:
        """
        Head commit SHA of the Pull Request base.
        """
        return pulumi.get(self, "base_sha")

    @property
    @pulumi.getter
    def body(self) -> str:
        """
        Body of the Pull Request.
        """
        return pulumi.get(self, "body")

    @property
    @pulumi.getter
    def draft(self) -> bool:
        """
        Indicates Whether this Pull Request is a draft.
        """
        return pulumi.get(self, "draft")

    @property
    @pulumi.getter(name="headOwner")
    def head_owner(self) -> str:
        """
        Owner of the Pull Request head repository.
        """
        return pulumi.get(self, "head_owner")

    @property
    @pulumi.getter(name="headRef")
    def head_ref(self) -> str:
        return pulumi.get(self, "head_ref")

    @property
    @pulumi.getter(name="headRepository")
    def head_repository(self) -> str:
        """
        Name of the Pull Request head repository.
        """
        return pulumi.get(self, "head_repository")

    @property
    @pulumi.getter(name="headSha")
    def head_sha(self) -> str:
        """
        Head commit SHA of the Pull Request head.
        """
        return pulumi.get(self, "head_sha")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def labels(self) -> Sequence[str]:
        """
        List of label names set on the Pull Request.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter(name="maintainerCanModify")
    def maintainer_can_modify(self) -> bool:
        """
        Indicates whether the base repository maintainers can modify the Pull Request.
        """
        return pulumi.get(self, "maintainer_can_modify")

    @property
    @pulumi.getter
    def number(self) -> int:
        return pulumi.get(self, "number")

    @property
    @pulumi.getter(name="openedAt")
    def opened_at(self) -> int:
        """
        Unix timestamp indicating the Pull Request creation time.
        """
        return pulumi.get(self, "opened_at")

    @property
    @pulumi.getter(name="openedBy")
    def opened_by(self) -> str:
        """
        GitHub login of the user who opened the Pull Request.
        """
        return pulumi.get(self, "opened_by")

    @property
    @pulumi.getter
    def owner(self) -> Optional[str]:
        return pulumi.get(self, "owner")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        the current Pull Request state - can be "open", "closed" or "merged".
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def title(self) -> str:
        """
        The title of the Pull Request.
        """
        return pulumi.get(self, "title")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> int:
        """
        The timestamp of the last Pull Request update.
        """
        return pulumi.get(self, "updated_at")


class AwaitableGetRepositoryPullRequestResult(GetRepositoryPullRequestResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRepositoryPullRequestResult(
            base_ref=self.base_ref,
            base_repository=self.base_repository,
            base_sha=self.base_sha,
            body=self.body,
            draft=self.draft,
            head_owner=self.head_owner,
            head_ref=self.head_ref,
            head_repository=self.head_repository,
            head_sha=self.head_sha,
            id=self.id,
            labels=self.labels,
            maintainer_can_modify=self.maintainer_can_modify,
            number=self.number,
            opened_at=self.opened_at,
            opened_by=self.opened_by,
            owner=self.owner,
            state=self.state,
            title=self.title,
            updated_at=self.updated_at)


def get_repository_pull_request(base_repository: Optional[str] = None,
                                number: Optional[int] = None,
                                owner: Optional[str] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRepositoryPullRequestResult:
    """
    Use this data source to retrieve information about a specific GitHub Pull Request in a repository.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_github as github

    example = github.get_repository_pull_request(base_repository="example_repository",
        number=1)
    ```


    :param str base_repository: Name of the base repository to retrieve the Pull Request from.
    :param int number: The number of the Pull Request within the repository.
    :param str owner: Owner of the repository. If not provided, the provider's default owner is used.
    """
    __args__ = dict()
    __args__['baseRepository'] = base_repository
    __args__['number'] = number
    __args__['owner'] = owner
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('github:index/getRepositoryPullRequest:getRepositoryPullRequest', __args__, opts=opts, typ=GetRepositoryPullRequestResult).value

    return AwaitableGetRepositoryPullRequestResult(
        base_ref=pulumi.get(__ret__, 'base_ref'),
        base_repository=pulumi.get(__ret__, 'base_repository'),
        base_sha=pulumi.get(__ret__, 'base_sha'),
        body=pulumi.get(__ret__, 'body'),
        draft=pulumi.get(__ret__, 'draft'),
        head_owner=pulumi.get(__ret__, 'head_owner'),
        head_ref=pulumi.get(__ret__, 'head_ref'),
        head_repository=pulumi.get(__ret__, 'head_repository'),
        head_sha=pulumi.get(__ret__, 'head_sha'),
        id=pulumi.get(__ret__, 'id'),
        labels=pulumi.get(__ret__, 'labels'),
        maintainer_can_modify=pulumi.get(__ret__, 'maintainer_can_modify'),
        number=pulumi.get(__ret__, 'number'),
        opened_at=pulumi.get(__ret__, 'opened_at'),
        opened_by=pulumi.get(__ret__, 'opened_by'),
        owner=pulumi.get(__ret__, 'owner'),
        state=pulumi.get(__ret__, 'state'),
        title=pulumi.get(__ret__, 'title'),
        updated_at=pulumi.get(__ret__, 'updated_at'))


@_utilities.lift_output_func(get_repository_pull_request)
def get_repository_pull_request_output(base_repository: Optional[pulumi.Input[str]] = None,
                                       number: Optional[pulumi.Input[int]] = None,
                                       owner: Optional[pulumi.Input[Optional[str]]] = None,
                                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetRepositoryPullRequestResult]:
    """
    Use this data source to retrieve information about a specific GitHub Pull Request in a repository.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_github as github

    example = github.get_repository_pull_request(base_repository="example_repository",
        number=1)
    ```


    :param str base_repository: Name of the base repository to retrieve the Pull Request from.
    :param int number: The number of the Pull Request within the repository.
    :param str owner: Owner of the repository. If not provided, the provider's default owner is used.
    """
    ...
