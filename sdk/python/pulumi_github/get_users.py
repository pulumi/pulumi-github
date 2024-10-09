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
    'GetUsersResult',
    'AwaitableGetUsersResult',
    'get_users',
    'get_users_output',
]

@pulumi.output_type
class GetUsersResult:
    """
    A collection of values returned by getUsers.
    """
    def __init__(__self__, emails=None, id=None, logins=None, node_ids=None, unknown_logins=None, usernames=None):
        if emails and not isinstance(emails, list):
            raise TypeError("Expected argument 'emails' to be a list")
        pulumi.set(__self__, "emails", emails)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if logins and not isinstance(logins, list):
            raise TypeError("Expected argument 'logins' to be a list")
        pulumi.set(__self__, "logins", logins)
        if node_ids and not isinstance(node_ids, list):
            raise TypeError("Expected argument 'node_ids' to be a list")
        pulumi.set(__self__, "node_ids", node_ids)
        if unknown_logins and not isinstance(unknown_logins, list):
            raise TypeError("Expected argument 'unknown_logins' to be a list")
        pulumi.set(__self__, "unknown_logins", unknown_logins)
        if usernames and not isinstance(usernames, list):
            raise TypeError("Expected argument 'usernames' to be a list")
        pulumi.set(__self__, "usernames", usernames)

    @property
    @pulumi.getter
    def emails(self) -> Sequence[str]:
        """
        list of the user's publicly visible profile email (will be empty string in case if user decided not to show it).
        """
        return pulumi.get(self, "emails")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def logins(self) -> Sequence[str]:
        """
        list of logins of users that could be found.
        """
        return pulumi.get(self, "logins")

    @property
    @pulumi.getter(name="nodeIds")
    def node_ids(self) -> Sequence[str]:
        """
        list of Node IDs of users that could be found.
        """
        return pulumi.get(self, "node_ids")

    @property
    @pulumi.getter(name="unknownLogins")
    def unknown_logins(self) -> Sequence[str]:
        """
        list of logins without matching user.
        """
        return pulumi.get(self, "unknown_logins")

    @property
    @pulumi.getter
    def usernames(self) -> Sequence[str]:
        return pulumi.get(self, "usernames")


class AwaitableGetUsersResult(GetUsersResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetUsersResult(
            emails=self.emails,
            id=self.id,
            logins=self.logins,
            node_ids=self.node_ids,
            unknown_logins=self.unknown_logins,
            usernames=self.usernames)


def get_users(usernames: Optional[Sequence[str]] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetUsersResult:
    """
    Use this data source to retrieve information about multiple GitHub users at once.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_github as github

    # Retrieve information about multiple GitHub users.
    example = github.get_users(usernames=[
        "example1",
        "example2",
        "example3",
    ])
    pulumi.export("validUsers", example.logins)
    pulumi.export("invalidUsers", example.unknown_logins)
    ```


    :param Sequence[str] usernames: List of usernames.
    """
    __args__ = dict()
    __args__['usernames'] = usernames
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('github:index/getUsers:getUsers', __args__, opts=opts, typ=GetUsersResult).value

    return AwaitableGetUsersResult(
        emails=pulumi.get(__ret__, 'emails'),
        id=pulumi.get(__ret__, 'id'),
        logins=pulumi.get(__ret__, 'logins'),
        node_ids=pulumi.get(__ret__, 'node_ids'),
        unknown_logins=pulumi.get(__ret__, 'unknown_logins'),
        usernames=pulumi.get(__ret__, 'usernames'))
def get_users_output(usernames: Optional[pulumi.Input[Sequence[str]]] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetUsersResult]:
    """
    Use this data source to retrieve information about multiple GitHub users at once.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_github as github

    # Retrieve information about multiple GitHub users.
    example = github.get_users(usernames=[
        "example1",
        "example2",
        "example3",
    ])
    pulumi.export("validUsers", example.logins)
    pulumi.export("invalidUsers", example.unknown_logins)
    ```


    :param Sequence[str] usernames: List of usernames.
    """
    __args__ = dict()
    __args__['usernames'] = usernames
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('github:index/getUsers:getUsers', __args__, opts=opts, typ=GetUsersResult)
    return __ret__.apply(lambda __response__: GetUsersResult(
        emails=pulumi.get(__response__, 'emails'),
        id=pulumi.get(__response__, 'id'),
        logins=pulumi.get(__response__, 'logins'),
        node_ids=pulumi.get(__response__, 'node_ids'),
        unknown_logins=pulumi.get(__response__, 'unknown_logins'),
        usernames=pulumi.get(__response__, 'usernames')))
