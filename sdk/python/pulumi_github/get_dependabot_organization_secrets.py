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
from . import outputs

__all__ = [
    'GetDependabotOrganizationSecretsResult',
    'AwaitableGetDependabotOrganizationSecretsResult',
    'get_dependabot_organization_secrets',
    'get_dependabot_organization_secrets_output',
]

@pulumi.output_type
class GetDependabotOrganizationSecretsResult:
    """
    A collection of values returned by getDependabotOrganizationSecrets.
    """
    def __init__(__self__, id=None, secrets=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if secrets and not isinstance(secrets, list):
            raise TypeError("Expected argument 'secrets' to be a list")
        pulumi.set(__self__, "secrets", secrets)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def secrets(self) -> Sequence['outputs.GetDependabotOrganizationSecretsSecretResult']:
        """
        list of secrets for the repository
        """
        return pulumi.get(self, "secrets")


class AwaitableGetDependabotOrganizationSecretsResult(GetDependabotOrganizationSecretsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDependabotOrganizationSecretsResult(
            id=self.id,
            secrets=self.secrets)


def get_dependabot_organization_secrets(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDependabotOrganizationSecretsResult:
    """
    Use this data source to retrieve the list of dependabot secrets of the organization.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_github as github

    example = github.get_dependabot_organization_secrets()
    ```
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('github:index/getDependabotOrganizationSecrets:getDependabotOrganizationSecrets', __args__, opts=opts, typ=GetDependabotOrganizationSecretsResult).value

    return AwaitableGetDependabotOrganizationSecretsResult(
        id=pulumi.get(__ret__, 'id'),
        secrets=pulumi.get(__ret__, 'secrets'))
def get_dependabot_organization_secrets_output(opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDependabotOrganizationSecretsResult]:
    """
    Use this data source to retrieve the list of dependabot secrets of the organization.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_github as github

    example = github.get_dependabot_organization_secrets()
    ```
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('github:index/getDependabotOrganizationSecrets:getDependabotOrganizationSecrets', __args__, opts=opts, typ=GetDependabotOrganizationSecretsResult)
    return __ret__.apply(lambda __response__: GetDependabotOrganizationSecretsResult(
        id=pulumi.get(__response__, 'id'),
        secrets=pulumi.get(__response__, 'secrets')))
