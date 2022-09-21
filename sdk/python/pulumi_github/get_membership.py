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
    'GetMembershipResult',
    'AwaitableGetMembershipResult',
    'get_membership',
    'get_membership_output',
]

@pulumi.output_type
class GetMembershipResult:
    """
    A collection of values returned by getMembership.
    """
    def __init__(__self__, etag=None, id=None, organization=None, role=None, username=None):
        if etag and not isinstance(etag, str):
            raise TypeError("Expected argument 'etag' to be a str")
        pulumi.set(__self__, "etag", etag)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if organization and not isinstance(organization, str):
            raise TypeError("Expected argument 'organization' to be a str")
        pulumi.set(__self__, "organization", organization)
        if role and not isinstance(role, str):
            raise TypeError("Expected argument 'role' to be a str")
        pulumi.set(__self__, "role", role)
        if username and not isinstance(username, str):
            raise TypeError("Expected argument 'username' to be a str")
        pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter
    def etag(self) -> str:
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def organization(self) -> Optional[str]:
        return pulumi.get(self, "organization")

    @property
    @pulumi.getter
    def role(self) -> str:
        return pulumi.get(self, "role")

    @property
    @pulumi.getter
    def username(self) -> str:
        return pulumi.get(self, "username")


class AwaitableGetMembershipResult(GetMembershipResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetMembershipResult(
            etag=self.etag,
            id=self.id,
            organization=self.organization,
            role=self.role,
            username=self.username)


def get_membership(organization: Optional[str] = None,
                   username: Optional[str] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetMembershipResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['organization'] = organization
    __args__['username'] = username
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('github:index/getMembership:getMembership', __args__, opts=opts, typ=GetMembershipResult).value

    return AwaitableGetMembershipResult(
        etag=__ret__.etag,
        id=__ret__.id,
        organization=__ret__.organization,
        role=__ret__.role,
        username=__ret__.username)


@_utilities.lift_output_func(get_membership)
def get_membership_output(organization: Optional[pulumi.Input[Optional[str]]] = None,
                          username: Optional[pulumi.Input[str]] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetMembershipResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
