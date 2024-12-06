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
    'GetActionsOrganizationRegistrationTokenResult',
    'AwaitableGetActionsOrganizationRegistrationTokenResult',
    'get_actions_organization_registration_token',
    'get_actions_organization_registration_token_output',
]

@pulumi.output_type
class GetActionsOrganizationRegistrationTokenResult:
    """
    A collection of values returned by getActionsOrganizationRegistrationToken.
    """
    def __init__(__self__, expires_at=None, id=None, token=None):
        if expires_at and not isinstance(expires_at, int):
            raise TypeError("Expected argument 'expires_at' to be a int")
        pulumi.set(__self__, "expires_at", expires_at)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if token and not isinstance(token, str):
            raise TypeError("Expected argument 'token' to be a str")
        pulumi.set(__self__, "token", token)

    @property
    @pulumi.getter(name="expiresAt")
    def expires_at(self) -> int:
        """
        The token expiration date.
        """
        return pulumi.get(self, "expires_at")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def token(self) -> str:
        """
        The token that has been retrieved.
        """
        return pulumi.get(self, "token")


class AwaitableGetActionsOrganizationRegistrationTokenResult(GetActionsOrganizationRegistrationTokenResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetActionsOrganizationRegistrationTokenResult(
            expires_at=self.expires_at,
            id=self.id,
            token=self.token)


def get_actions_organization_registration_token(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetActionsOrganizationRegistrationTokenResult:
    """
    Use this data source to retrieve a GitHub Actions organization registration token. This token can then be used to register a self-hosted runner.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_github as github

    example = github.get_actions_organization_registration_token()
    ```
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('github:index/getActionsOrganizationRegistrationToken:getActionsOrganizationRegistrationToken', __args__, opts=opts, typ=GetActionsOrganizationRegistrationTokenResult).value

    return AwaitableGetActionsOrganizationRegistrationTokenResult(
        expires_at=pulumi.get(__ret__, 'expires_at'),
        id=pulumi.get(__ret__, 'id'),
        token=pulumi.get(__ret__, 'token'))
def get_actions_organization_registration_token_output(opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetActionsOrganizationRegistrationTokenResult]:
    """
    Use this data source to retrieve a GitHub Actions organization registration token. This token can then be used to register a self-hosted runner.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_github as github

    example = github.get_actions_organization_registration_token()
    ```
    """
    __args__ = dict()
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('github:index/getActionsOrganizationRegistrationToken:getActionsOrganizationRegistrationToken', __args__, opts=opts, typ=GetActionsOrganizationRegistrationTokenResult)
    return __ret__.apply(lambda __response__: GetActionsOrganizationRegistrationTokenResult(
        expires_at=pulumi.get(__response__, 'expires_at'),
        id=pulumi.get(__response__, 'id'),
        token=pulumi.get(__response__, 'token')))
