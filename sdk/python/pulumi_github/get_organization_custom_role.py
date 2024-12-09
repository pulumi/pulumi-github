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
    'GetOrganizationCustomRoleResult',
    'AwaitableGetOrganizationCustomRoleResult',
    'get_organization_custom_role',
    'get_organization_custom_role_output',
]

@pulumi.output_type
class GetOrganizationCustomRoleResult:
    """
    A collection of values returned by getOrganizationCustomRole.
    """
    def __init__(__self__, base_role=None, description=None, id=None, name=None, permissions=None):
        if base_role and not isinstance(base_role, str):
            raise TypeError("Expected argument 'base_role' to be a str")
        pulumi.set(__self__, "base_role", base_role)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if permissions and not isinstance(permissions, list):
            raise TypeError("Expected argument 'permissions' to be a list")
        pulumi.set(__self__, "permissions", permissions)

    @property
    @pulumi.getter(name="baseRole")
    def base_role(self) -> str:
        """
        The system role from which the role inherits permissions.
        """
        return pulumi.get(self, "base_role")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        The description for the custom role.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def permissions(self) -> Sequence[str]:
        """
        A list of additional permissions included in this role.
        """
        return pulumi.get(self, "permissions")


class AwaitableGetOrganizationCustomRoleResult(GetOrganizationCustomRoleResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetOrganizationCustomRoleResult(
            base_role=self.base_role,
            description=self.description,
            id=self.id,
            name=self.name,
            permissions=self.permissions)


def get_organization_custom_role(name: Optional[str] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetOrganizationCustomRoleResult:
    """
    Use this data source to retrieve information about a custom role in a GitHub Organization.

    > Note: Custom roles are currently only available in GitHub Enterprise Cloud.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_github as github

    example = github.get_organization_custom_role(name="example")
    ```


    :param str name: The name of the custom role.
    """
    __args__ = dict()
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('github:index/getOrganizationCustomRole:getOrganizationCustomRole', __args__, opts=opts, typ=GetOrganizationCustomRoleResult).value

    return AwaitableGetOrganizationCustomRoleResult(
        base_role=pulumi.get(__ret__, 'base_role'),
        description=pulumi.get(__ret__, 'description'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        permissions=pulumi.get(__ret__, 'permissions'))
def get_organization_custom_role_output(name: Optional[pulumi.Input[str]] = None,
                                        opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetOrganizationCustomRoleResult]:
    """
    Use this data source to retrieve information about a custom role in a GitHub Organization.

    > Note: Custom roles are currently only available in GitHub Enterprise Cloud.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_github as github

    example = github.get_organization_custom_role(name="example")
    ```


    :param str name: The name of the custom role.
    """
    __args__ = dict()
    __args__['name'] = name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('github:index/getOrganizationCustomRole:getOrganizationCustomRole', __args__, opts=opts, typ=GetOrganizationCustomRoleResult)
    return __ret__.apply(lambda __response__: GetOrganizationCustomRoleResult(
        base_role=pulumi.get(__response__, 'base_role'),
        description=pulumi.get(__response__, 'description'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        permissions=pulumi.get(__response__, 'permissions')))
