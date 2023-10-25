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
    'GetCodespacesUserSecretsResult',
    'AwaitableGetCodespacesUserSecretsResult',
    'get_codespaces_user_secrets',
    'get_codespaces_user_secrets_output',
]

@pulumi.output_type
class GetCodespacesUserSecretsResult:
    """
    A collection of values returned by getCodespacesUserSecrets.
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
    def secrets(self) -> Sequence['outputs.GetCodespacesUserSecretsSecretResult']:
        """
        list of secrets for the repository
        """
        return pulumi.get(self, "secrets")


class AwaitableGetCodespacesUserSecretsResult(GetCodespacesUserSecretsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCodespacesUserSecretsResult(
            id=self.id,
            secrets=self.secrets)


def get_codespaces_user_secrets(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetCodespacesUserSecretsResult:
    """
    Use this data source to retrieve the list of codespaces secrets of the user.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_github as github

    example = github.get_codespaces_user_secrets()
    ```
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('github:index/getCodespacesUserSecrets:getCodespacesUserSecrets', __args__, opts=opts, typ=GetCodespacesUserSecretsResult).value

    return AwaitableGetCodespacesUserSecretsResult(
        id=pulumi.get(__ret__, 'id'),
        secrets=pulumi.get(__ret__, 'secrets'))


@_utilities.lift_output_func(get_codespaces_user_secrets)
def get_codespaces_user_secrets_output(opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetCodespacesUserSecretsResult]:
    """
    Use this data source to retrieve the list of codespaces secrets of the user.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_github as github

    example = github.get_codespaces_user_secrets()
    ```
    """
    ...
