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
    'GetBranchProtectionRulesResult',
    'AwaitableGetBranchProtectionRulesResult',
    'get_branch_protection_rules',
    'get_branch_protection_rules_output',
]

@pulumi.output_type
class GetBranchProtectionRulesResult:
    """
    A collection of values returned by getBranchProtectionRules.
    """
    def __init__(__self__, id=None, repository=None, rules=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if repository and not isinstance(repository, str):
            raise TypeError("Expected argument 'repository' to be a str")
        pulumi.set(__self__, "repository", repository)
        if rules and not isinstance(rules, list):
            raise TypeError("Expected argument 'rules' to be a list")
        pulumi.set(__self__, "rules", rules)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def repository(self) -> str:
        return pulumi.get(self, "repository")

    @property
    @pulumi.getter
    def rules(self) -> Sequence['outputs.GetBranchProtectionRulesRuleResult']:
        """
        Collection of Branch Protection Rules. Each of the results conforms to the following scheme:
        """
        return pulumi.get(self, "rules")


class AwaitableGetBranchProtectionRulesResult(GetBranchProtectionRulesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetBranchProtectionRulesResult(
            id=self.id,
            repository=self.repository,
            rules=self.rules)


def get_branch_protection_rules(repository: Optional[str] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetBranchProtectionRulesResult:
    """
    Use this data source to retrieve a list of repository branch protection rules.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_github as github

    example = github.get_branch_protection_rules(repository="example")
    ```


    :param str repository: The GitHub repository name.
    """
    __args__ = dict()
    __args__['repository'] = repository
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('github:index/getBranchProtectionRules:getBranchProtectionRules', __args__, opts=opts, typ=GetBranchProtectionRulesResult).value

    return AwaitableGetBranchProtectionRulesResult(
        id=__ret__.id,
        repository=__ret__.repository,
        rules=__ret__.rules)


@_utilities.lift_output_func(get_branch_protection_rules)
def get_branch_protection_rules_output(repository: Optional[pulumi.Input[str]] = None,
                                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetBranchProtectionRulesResult]:
    """
    Use this data source to retrieve a list of repository branch protection rules.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_github as github

    example = github.get_branch_protection_rules(repository="example")
    ```


    :param str repository: The GitHub repository name.
    """
    ...