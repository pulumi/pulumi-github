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

__all__ = [
    'GetActionsOrganizationOidcSubjectClaimCustomizationTemplateResult',
    'AwaitableGetActionsOrganizationOidcSubjectClaimCustomizationTemplateResult',
    'get_actions_organization_oidc_subject_claim_customization_template',
    'get_actions_organization_oidc_subject_claim_customization_template_output',
]

@pulumi.output_type
class GetActionsOrganizationOidcSubjectClaimCustomizationTemplateResult:
    """
    A collection of values returned by getActionsOrganizationOidcSubjectClaimCustomizationTemplate.
    """
    def __init__(__self__, id=None, include_claim_keys=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if include_claim_keys and not isinstance(include_claim_keys, list):
            raise TypeError("Expected argument 'include_claim_keys' to be a list")
        pulumi.set(__self__, "include_claim_keys", include_claim_keys)

    @_builtins.property
    @pulumi.getter
    def id(self) -> _builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @_builtins.property
    @pulumi.getter(name="includeClaimKeys")
    def include_claim_keys(self) -> Sequence[_builtins.str]:
        """
        The list of OpenID Connect claim keys.
        """
        return pulumi.get(self, "include_claim_keys")


class AwaitableGetActionsOrganizationOidcSubjectClaimCustomizationTemplateResult(GetActionsOrganizationOidcSubjectClaimCustomizationTemplateResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetActionsOrganizationOidcSubjectClaimCustomizationTemplateResult(
            id=self.id,
            include_claim_keys=self.include_claim_keys)


def get_actions_organization_oidc_subject_claim_customization_template(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetActionsOrganizationOidcSubjectClaimCustomizationTemplateResult:
    """
    Use this data source to retrieve the OpenID Connect subject claim customization template for an organization

    ## Example Usage

    ```python
    import pulumi
    import pulumi_github as github

    example = github.get_actions_organization_oidc_subject_claim_customization_template()
    ```
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('github:index/getActionsOrganizationOidcSubjectClaimCustomizationTemplate:getActionsOrganizationOidcSubjectClaimCustomizationTemplate', __args__, opts=opts, typ=GetActionsOrganizationOidcSubjectClaimCustomizationTemplateResult).value

    return AwaitableGetActionsOrganizationOidcSubjectClaimCustomizationTemplateResult(
        id=pulumi.get(__ret__, 'id'),
        include_claim_keys=pulumi.get(__ret__, 'include_claim_keys'))
def get_actions_organization_oidc_subject_claim_customization_template_output(opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetActionsOrganizationOidcSubjectClaimCustomizationTemplateResult]:
    """
    Use this data source to retrieve the OpenID Connect subject claim customization template for an organization

    ## Example Usage

    ```python
    import pulumi
    import pulumi_github as github

    example = github.get_actions_organization_oidc_subject_claim_customization_template()
    ```
    """
    __args__ = dict()
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('github:index/getActionsOrganizationOidcSubjectClaimCustomizationTemplate:getActionsOrganizationOidcSubjectClaimCustomizationTemplate', __args__, opts=opts, typ=GetActionsOrganizationOidcSubjectClaimCustomizationTemplateResult)
    return __ret__.apply(lambda __response__: GetActionsOrganizationOidcSubjectClaimCustomizationTemplateResult(
        id=pulumi.get(__response__, 'id'),
        include_claim_keys=pulumi.get(__response__, 'include_claim_keys')))
