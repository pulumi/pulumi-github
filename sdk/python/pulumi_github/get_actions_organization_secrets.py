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
    'GetActionsOrganizationSecretsResult',
    'AwaitableGetActionsOrganizationSecretsResult',
    'get_actions_organization_secrets',
    'get_actions_organization_secrets_output',
]

@pulumi.output_type
class GetActionsOrganizationSecretsResult:
    """
    A collection of values returned by getActionsOrganizationSecrets.
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
    def secrets(self) -> Sequence['outputs.GetActionsOrganizationSecretsSecretResult']:
        """
        list of secrets for the repository
        """
        return pulumi.get(self, "secrets")


class AwaitableGetActionsOrganizationSecretsResult(GetActionsOrganizationSecretsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetActionsOrganizationSecretsResult(
            id=self.id,
            secrets=self.secrets)


def get_actions_organization_secrets(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetActionsOrganizationSecretsResult:
    """
    Use this data source to retrieve the list of secrets of the organization.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_github as github

    example = github.get_actions_organization_secrets()
    ```
    <!--End PulumiCodeChooser -->
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('github:index/getActionsOrganizationSecrets:getActionsOrganizationSecrets', __args__, opts=opts, typ=GetActionsOrganizationSecretsResult).value

    return AwaitableGetActionsOrganizationSecretsResult(
        id=pulumi.get(__ret__, 'id'),
        secrets=pulumi.get(__ret__, 'secrets'))


@_utilities.lift_output_func(get_actions_organization_secrets)
def get_actions_organization_secrets_output(opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetActionsOrganizationSecretsResult]:
    """
    Use this data source to retrieve the list of secrets of the organization.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_github as github

    example = github.get_actions_organization_secrets()
    ```
    <!--End PulumiCodeChooser -->
    """
    ...
