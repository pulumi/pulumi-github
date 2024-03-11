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
    'GetDependabotSecretsResult',
    'AwaitableGetDependabotSecretsResult',
    'get_dependabot_secrets',
    'get_dependabot_secrets_output',
]

@pulumi.output_type
class GetDependabotSecretsResult:
    """
    A collection of values returned by getDependabotSecrets.
    """
    def __init__(__self__, full_name=None, id=None, name=None, secrets=None):
        if full_name and not isinstance(full_name, str):
            raise TypeError("Expected argument 'full_name' to be a str")
        pulumi.set(__self__, "full_name", full_name)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if secrets and not isinstance(secrets, list):
            raise TypeError("Expected argument 'secrets' to be a list")
        pulumi.set(__self__, "secrets", secrets)

    @property
    @pulumi.getter(name="fullName")
    def full_name(self) -> str:
        return pulumi.get(self, "full_name")

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
        """
        Secret name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def secrets(self) -> Sequence['outputs.GetDependabotSecretsSecretResult']:
        """
        list of dependabot secrets for the repository
        """
        return pulumi.get(self, "secrets")


class AwaitableGetDependabotSecretsResult(GetDependabotSecretsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDependabotSecretsResult(
            full_name=self.full_name,
            id=self.id,
            name=self.name,
            secrets=self.secrets)


def get_dependabot_secrets(full_name: Optional[str] = None,
                           name: Optional[str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDependabotSecretsResult:
    """
    Use this data source to retrieve the list of dependabot secrets for a GitHub repository.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_github as github

    example = github.get_dependabot_secrets(name="example")
    ```
    <!--End PulumiCodeChooser -->


    :param str full_name: Full name of the repository (in `org/name` format).
    :param str name: The name of the repository.
    """
    __args__ = dict()
    __args__['fullName'] = full_name
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('github:index/getDependabotSecrets:getDependabotSecrets', __args__, opts=opts, typ=GetDependabotSecretsResult).value

    return AwaitableGetDependabotSecretsResult(
        full_name=pulumi.get(__ret__, 'full_name'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        secrets=pulumi.get(__ret__, 'secrets'))


@_utilities.lift_output_func(get_dependabot_secrets)
def get_dependabot_secrets_output(full_name: Optional[pulumi.Input[Optional[str]]] = None,
                                  name: Optional[pulumi.Input[Optional[str]]] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDependabotSecretsResult]:
    """
    Use this data source to retrieve the list of dependabot secrets for a GitHub repository.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_github as github

    example = github.get_dependabot_secrets(name="example")
    ```
    <!--End PulumiCodeChooser -->


    :param str full_name: Full name of the repository (in `org/name` format).
    :param str name: The name of the repository.
    """
    ...
