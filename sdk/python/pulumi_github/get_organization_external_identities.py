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
    'GetOrganizationExternalIdentitiesResult',
    'AwaitableGetOrganizationExternalIdentitiesResult',
    'get_organization_external_identities',
    'get_organization_external_identities_output',
]

@pulumi.output_type
class GetOrganizationExternalIdentitiesResult:
    """
    A collection of values returned by getOrganizationExternalIdentities.
    """
    def __init__(__self__, id=None, identities=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if identities and not isinstance(identities, list):
            raise TypeError("Expected argument 'identities' to be a list")
        pulumi.set(__self__, "identities", identities)

    @_builtins.property
    @pulumi.getter
    def id(self) -> _builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @_builtins.property
    @pulumi.getter
    def identities(self) -> Sequence['outputs.GetOrganizationExternalIdentitiesIdentityResult']:
        """
        An Array of identities returned from GitHub
        """
        return pulumi.get(self, "identities")


class AwaitableGetOrganizationExternalIdentitiesResult(GetOrganizationExternalIdentitiesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetOrganizationExternalIdentitiesResult(
            id=self.id,
            identities=self.identities)


def get_organization_external_identities(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetOrganizationExternalIdentitiesResult:
    """
    Use this data source to retrieve each organization member's SAML or SCIM user
    attributes.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_github as github

    all = github.get_organization_external_identities()
    ```
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('github:index/getOrganizationExternalIdentities:getOrganizationExternalIdentities', __args__, opts=opts, typ=GetOrganizationExternalIdentitiesResult).value

    return AwaitableGetOrganizationExternalIdentitiesResult(
        id=pulumi.get(__ret__, 'id'),
        identities=pulumi.get(__ret__, 'identities'))
def get_organization_external_identities_output(opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetOrganizationExternalIdentitiesResult]:
    """
    Use this data source to retrieve each organization member's SAML or SCIM user
    attributes.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_github as github

    all = github.get_organization_external_identities()
    ```
    """
    __args__ = dict()
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('github:index/getOrganizationExternalIdentities:getOrganizationExternalIdentities', __args__, opts=opts, typ=GetOrganizationExternalIdentitiesResult)
    return __ret__.apply(lambda __response__: GetOrganizationExternalIdentitiesResult(
        id=pulumi.get(__response__, 'id'),
        identities=pulumi.get(__response__, 'identities')))
