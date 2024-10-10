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
    'GetActionsVariablesResult',
    'AwaitableGetActionsVariablesResult',
    'get_actions_variables',
    'get_actions_variables_output',
]

@pulumi.output_type
class GetActionsVariablesResult:
    """
    A collection of values returned by getActionsVariables.
    """
    def __init__(__self__, full_name=None, id=None, name=None, variables=None):
        if full_name and not isinstance(full_name, str):
            raise TypeError("Expected argument 'full_name' to be a str")
        pulumi.set(__self__, "full_name", full_name)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if variables and not isinstance(variables, list):
            raise TypeError("Expected argument 'variables' to be a list")
        pulumi.set(__self__, "variables", variables)

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
    def variables(self) -> Sequence['outputs.GetActionsVariablesVariableResult']:
        """
        list of variables for the repository
        """
        return pulumi.get(self, "variables")


class AwaitableGetActionsVariablesResult(GetActionsVariablesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetActionsVariablesResult(
            full_name=self.full_name,
            id=self.id,
            name=self.name,
            variables=self.variables)


def get_actions_variables(full_name: Optional[str] = None,
                          name: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetActionsVariablesResult:
    """
    Use this data source to retrieve the list of variables for a GitHub repository.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_github as github

    example = github.get_actions_variables(name="example")
    ```


    :param str full_name: Full name of the repository (in `org/name` format).
    :param str name: The name of the repository.
    """
    __args__ = dict()
    __args__['fullName'] = full_name
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('github:index/getActionsVariables:getActionsVariables', __args__, opts=opts, typ=GetActionsVariablesResult).value

    return AwaitableGetActionsVariablesResult(
        full_name=pulumi.get(__ret__, 'full_name'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        variables=pulumi.get(__ret__, 'variables'))
def get_actions_variables_output(full_name: Optional[pulumi.Input[Optional[str]]] = None,
                                 name: Optional[pulumi.Input[Optional[str]]] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetActionsVariablesResult]:
    """
    Use this data source to retrieve the list of variables for a GitHub repository.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_github as github

    example = github.get_actions_variables(name="example")
    ```


    :param str full_name: Full name of the repository (in `org/name` format).
    :param str name: The name of the repository.
    """
    __args__ = dict()
    __args__['fullName'] = full_name
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('github:index/getActionsVariables:getActionsVariables', __args__, opts=opts, typ=GetActionsVariablesResult)
    return __ret__.apply(lambda __response__: GetActionsVariablesResult(
        full_name=pulumi.get(__response__, 'full_name'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        variables=pulumi.get(__response__, 'variables')))
