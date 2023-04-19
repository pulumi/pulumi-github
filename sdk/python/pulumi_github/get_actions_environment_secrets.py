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
    'GetActionsEnvironmentSecretsResult',
    'AwaitableGetActionsEnvironmentSecretsResult',
    'get_actions_environment_secrets',
    'get_actions_environment_secrets_output',
]

@pulumi.output_type
class GetActionsEnvironmentSecretsResult:
    """
    A collection of values returned by getActionsEnvironmentSecrets.
    """
    def __init__(__self__, environment=None, full_name=None, id=None, name=None, secrets=None):
        if environment and not isinstance(environment, str):
            raise TypeError("Expected argument 'environment' to be a str")
        pulumi.set(__self__, "environment", environment)
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
    @pulumi.getter
    def environment(self) -> str:
        return pulumi.get(self, "environment")

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
        Name of the variable
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def secrets(self) -> Sequence['outputs.GetActionsEnvironmentSecretsSecretResult']:
        return pulumi.get(self, "secrets")


class AwaitableGetActionsEnvironmentSecretsResult(GetActionsEnvironmentSecretsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetActionsEnvironmentSecretsResult(
            environment=self.environment,
            full_name=self.full_name,
            id=self.id,
            name=self.name,
            secrets=self.secrets)


def get_actions_environment_secrets(environment: Optional[str] = None,
                                    full_name: Optional[str] = None,
                                    name: Optional[str] = None,
                                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetActionsEnvironmentSecretsResult:
    """
    Use this data source to retrieve the list of secrets of the repository environment.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_github as github

    example = github.get_actions_environment_secrets(environment="exampleEnvironment",
        name="exampleRepo")
    ```


    :param str name: Name of the variable
    """
    __args__ = dict()
    __args__['environment'] = environment
    __args__['fullName'] = full_name
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('github:index/getActionsEnvironmentSecrets:getActionsEnvironmentSecrets', __args__, opts=opts, typ=GetActionsEnvironmentSecretsResult).value

    return AwaitableGetActionsEnvironmentSecretsResult(
        environment=__ret__.environment,
        full_name=__ret__.full_name,
        id=__ret__.id,
        name=__ret__.name,
        secrets=__ret__.secrets)


@_utilities.lift_output_func(get_actions_environment_secrets)
def get_actions_environment_secrets_output(environment: Optional[pulumi.Input[str]] = None,
                                           full_name: Optional[pulumi.Input[Optional[str]]] = None,
                                           name: Optional[pulumi.Input[Optional[str]]] = None,
                                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetActionsEnvironmentSecretsResult]:
    """
    Use this data source to retrieve the list of secrets of the repository environment.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_github as github

    example = github.get_actions_environment_secrets(environment="exampleEnvironment",
        name="exampleRepo")
    ```


    :param str name: Name of the variable
    """
    ...