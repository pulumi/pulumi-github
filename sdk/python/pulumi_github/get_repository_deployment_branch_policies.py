# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
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
    'GetRepositoryDeploymentBranchPoliciesResult',
    'AwaitableGetRepositoryDeploymentBranchPoliciesResult',
    'get_repository_deployment_branch_policies',
    'get_repository_deployment_branch_policies_output',
]

@pulumi.output_type
class GetRepositoryDeploymentBranchPoliciesResult:
    """
    A collection of values returned by getRepositoryDeploymentBranchPolicies.
    """
    def __init__(__self__, deployment_branch_policies=None, environment_name=None, id=None, repository=None):
        if deployment_branch_policies and not isinstance(deployment_branch_policies, list):
            raise TypeError("Expected argument 'deployment_branch_policies' to be a list")
        pulumi.set(__self__, "deployment_branch_policies", deployment_branch_policies)
        if environment_name and not isinstance(environment_name, str):
            raise TypeError("Expected argument 'environment_name' to be a str")
        pulumi.set(__self__, "environment_name", environment_name)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if repository and not isinstance(repository, str):
            raise TypeError("Expected argument 'repository' to be a str")
        pulumi.set(__self__, "repository", repository)

    @property
    @pulumi.getter(name="deploymentBranchPolicies")
    def deployment_branch_policies(self) -> Sequence['outputs.GetRepositoryDeploymentBranchPoliciesDeploymentBranchPolicyResult']:
        """
        The list of this repository / environment deployment policies. Each element of `deployment_branch_policies` has the following attributes:
        """
        return pulumi.get(self, "deployment_branch_policies")

    @property
    @pulumi.getter(name="environmentName")
    def environment_name(self) -> builtins.str:
        return pulumi.get(self, "environment_name")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def repository(self) -> builtins.str:
        return pulumi.get(self, "repository")


class AwaitableGetRepositoryDeploymentBranchPoliciesResult(GetRepositoryDeploymentBranchPoliciesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRepositoryDeploymentBranchPoliciesResult(
            deployment_branch_policies=self.deployment_branch_policies,
            environment_name=self.environment_name,
            id=self.id,
            repository=self.repository)


def get_repository_deployment_branch_policies(environment_name: Optional[builtins.str] = None,
                                              repository: Optional[builtins.str] = None,
                                              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRepositoryDeploymentBranchPoliciesResult:
    """
    Use this data source to retrieve deployment branch policies for a repository / environment.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_github as github

    example = github.get_repository_deployment_branch_policies(repository="example-repository",
        environment_name="env_name")
    ```


    :param builtins.str environment_name: Name of the environment to retrieve the deployment branch policies  from.
    :param builtins.str repository: Name of the repository to retrieve the deployment branch policies from.
    """
    __args__ = dict()
    __args__['environmentName'] = environment_name
    __args__['repository'] = repository
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('github:index/getRepositoryDeploymentBranchPolicies:getRepositoryDeploymentBranchPolicies', __args__, opts=opts, typ=GetRepositoryDeploymentBranchPoliciesResult).value

    return AwaitableGetRepositoryDeploymentBranchPoliciesResult(
        deployment_branch_policies=pulumi.get(__ret__, 'deployment_branch_policies'),
        environment_name=pulumi.get(__ret__, 'environment_name'),
        id=pulumi.get(__ret__, 'id'),
        repository=pulumi.get(__ret__, 'repository'))
def get_repository_deployment_branch_policies_output(environment_name: Optional[pulumi.Input[builtins.str]] = None,
                                                     repository: Optional[pulumi.Input[builtins.str]] = None,
                                                     opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetRepositoryDeploymentBranchPoliciesResult]:
    """
    Use this data source to retrieve deployment branch policies for a repository / environment.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_github as github

    example = github.get_repository_deployment_branch_policies(repository="example-repository",
        environment_name="env_name")
    ```


    :param builtins.str environment_name: Name of the environment to retrieve the deployment branch policies  from.
    :param builtins.str repository: Name of the repository to retrieve the deployment branch policies from.
    """
    __args__ = dict()
    __args__['environmentName'] = environment_name
    __args__['repository'] = repository
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('github:index/getRepositoryDeploymentBranchPolicies:getRepositoryDeploymentBranchPolicies', __args__, opts=opts, typ=GetRepositoryDeploymentBranchPoliciesResult)
    return __ret__.apply(lambda __response__: GetRepositoryDeploymentBranchPoliciesResult(
        deployment_branch_policies=pulumi.get(__response__, 'deployment_branch_policies'),
        environment_name=pulumi.get(__response__, 'environment_name'),
        id=pulumi.get(__response__, 'id'),
        repository=pulumi.get(__response__, 'repository')))
