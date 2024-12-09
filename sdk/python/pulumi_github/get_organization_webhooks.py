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
    'GetOrganizationWebhooksResult',
    'AwaitableGetOrganizationWebhooksResult',
    'get_organization_webhooks',
    'get_organization_webhooks_output',
]

@pulumi.output_type
class GetOrganizationWebhooksResult:
    """
    A collection of values returned by getOrganizationWebhooks.
    """
    def __init__(__self__, id=None, webhooks=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if webhooks and not isinstance(webhooks, list):
            raise TypeError("Expected argument 'webhooks' to be a list")
        pulumi.set(__self__, "webhooks", webhooks)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def webhooks(self) -> Sequence['outputs.GetOrganizationWebhooksWebhookResult']:
        """
        An Array of GitHub Webhooks.  Each `webhook` block consists of the fields documented below.
        ___
        """
        return pulumi.get(self, "webhooks")


class AwaitableGetOrganizationWebhooksResult(GetOrganizationWebhooksResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetOrganizationWebhooksResult(
            id=self.id,
            webhooks=self.webhooks)


def get_organization_webhooks(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetOrganizationWebhooksResult:
    """
    Use this data source to retrieve all webhooks of the organization.

    ## Example Usage

    To retrieve *all* webhooks of the organization:

    ```python
    import pulumi
    import pulumi_github as github

    all = github.get_organization_webhooks()
    ```
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('github:index/getOrganizationWebhooks:getOrganizationWebhooks', __args__, opts=opts, typ=GetOrganizationWebhooksResult).value

    return AwaitableGetOrganizationWebhooksResult(
        id=pulumi.get(__ret__, 'id'),
        webhooks=pulumi.get(__ret__, 'webhooks'))
def get_organization_webhooks_output(opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetOrganizationWebhooksResult]:
    """
    Use this data source to retrieve all webhooks of the organization.

    ## Example Usage

    To retrieve *all* webhooks of the organization:

    ```python
    import pulumi
    import pulumi_github as github

    all = github.get_organization_webhooks()
    ```
    """
    __args__ = dict()
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('github:index/getOrganizationWebhooks:getOrganizationWebhooks', __args__, opts=opts, typ=GetOrganizationWebhooksResult)
    return __ret__.apply(lambda __response__: GetOrganizationWebhooksResult(
        id=pulumi.get(__response__, 'id'),
        webhooks=pulumi.get(__response__, 'webhooks')))
