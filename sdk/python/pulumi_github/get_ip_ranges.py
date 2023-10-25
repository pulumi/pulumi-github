# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetIpRangesResult',
    'AwaitableGetIpRangesResult',
    'get_ip_ranges',
    'get_ip_ranges_output',
]

@pulumi.output_type
class GetIpRangesResult:
    """
    A collection of values returned by getIpRanges.
    """
    def __init__(__self__, actions=None, actions_ipv4s=None, actions_ipv6s=None, api_ipv4s=None, api_ipv6s=None, apis=None, dependabot_ipv4s=None, dependabot_ipv6s=None, dependabots=None, git_ipv4s=None, git_ipv6s=None, gits=None, hooks=None, hooks_ipv4s=None, hooks_ipv6s=None, id=None, importer_ipv4s=None, importer_ipv6s=None, importers=None, pages=None, pages_ipv4s=None, pages_ipv6s=None, web_ipv4s=None, web_ipv6s=None, webs=None):
        if actions and not isinstance(actions, list):
            raise TypeError("Expected argument 'actions' to be a list")
        pulumi.set(__self__, "actions", actions)
        if actions_ipv4s and not isinstance(actions_ipv4s, list):
            raise TypeError("Expected argument 'actions_ipv4s' to be a list")
        pulumi.set(__self__, "actions_ipv4s", actions_ipv4s)
        if actions_ipv6s and not isinstance(actions_ipv6s, list):
            raise TypeError("Expected argument 'actions_ipv6s' to be a list")
        pulumi.set(__self__, "actions_ipv6s", actions_ipv6s)
        if api_ipv4s and not isinstance(api_ipv4s, list):
            raise TypeError("Expected argument 'api_ipv4s' to be a list")
        pulumi.set(__self__, "api_ipv4s", api_ipv4s)
        if api_ipv6s and not isinstance(api_ipv6s, list):
            raise TypeError("Expected argument 'api_ipv6s' to be a list")
        pulumi.set(__self__, "api_ipv6s", api_ipv6s)
        if apis and not isinstance(apis, list):
            raise TypeError("Expected argument 'apis' to be a list")
        pulumi.set(__self__, "apis", apis)
        if dependabot_ipv4s and not isinstance(dependabot_ipv4s, list):
            raise TypeError("Expected argument 'dependabot_ipv4s' to be a list")
        pulumi.set(__self__, "dependabot_ipv4s", dependabot_ipv4s)
        if dependabot_ipv6s and not isinstance(dependabot_ipv6s, list):
            raise TypeError("Expected argument 'dependabot_ipv6s' to be a list")
        pulumi.set(__self__, "dependabot_ipv6s", dependabot_ipv6s)
        if dependabots and not isinstance(dependabots, list):
            raise TypeError("Expected argument 'dependabots' to be a list")
        pulumi.set(__self__, "dependabots", dependabots)
        if git_ipv4s and not isinstance(git_ipv4s, list):
            raise TypeError("Expected argument 'git_ipv4s' to be a list")
        pulumi.set(__self__, "git_ipv4s", git_ipv4s)
        if git_ipv6s and not isinstance(git_ipv6s, list):
            raise TypeError("Expected argument 'git_ipv6s' to be a list")
        pulumi.set(__self__, "git_ipv6s", git_ipv6s)
        if gits and not isinstance(gits, list):
            raise TypeError("Expected argument 'gits' to be a list")
        pulumi.set(__self__, "gits", gits)
        if hooks and not isinstance(hooks, list):
            raise TypeError("Expected argument 'hooks' to be a list")
        pulumi.set(__self__, "hooks", hooks)
        if hooks_ipv4s and not isinstance(hooks_ipv4s, list):
            raise TypeError("Expected argument 'hooks_ipv4s' to be a list")
        pulumi.set(__self__, "hooks_ipv4s", hooks_ipv4s)
        if hooks_ipv6s and not isinstance(hooks_ipv6s, list):
            raise TypeError("Expected argument 'hooks_ipv6s' to be a list")
        pulumi.set(__self__, "hooks_ipv6s", hooks_ipv6s)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if importer_ipv4s and not isinstance(importer_ipv4s, list):
            raise TypeError("Expected argument 'importer_ipv4s' to be a list")
        pulumi.set(__self__, "importer_ipv4s", importer_ipv4s)
        if importer_ipv6s and not isinstance(importer_ipv6s, list):
            raise TypeError("Expected argument 'importer_ipv6s' to be a list")
        pulumi.set(__self__, "importer_ipv6s", importer_ipv6s)
        if importers and not isinstance(importers, list):
            raise TypeError("Expected argument 'importers' to be a list")
        pulumi.set(__self__, "importers", importers)
        if pages and not isinstance(pages, list):
            raise TypeError("Expected argument 'pages' to be a list")
        pulumi.set(__self__, "pages", pages)
        if pages_ipv4s and not isinstance(pages_ipv4s, list):
            raise TypeError("Expected argument 'pages_ipv4s' to be a list")
        pulumi.set(__self__, "pages_ipv4s", pages_ipv4s)
        if pages_ipv6s and not isinstance(pages_ipv6s, list):
            raise TypeError("Expected argument 'pages_ipv6s' to be a list")
        pulumi.set(__self__, "pages_ipv6s", pages_ipv6s)
        if web_ipv4s and not isinstance(web_ipv4s, list):
            raise TypeError("Expected argument 'web_ipv4s' to be a list")
        pulumi.set(__self__, "web_ipv4s", web_ipv4s)
        if web_ipv6s and not isinstance(web_ipv6s, list):
            raise TypeError("Expected argument 'web_ipv6s' to be a list")
        pulumi.set(__self__, "web_ipv6s", web_ipv6s)
        if webs and not isinstance(webs, list):
            raise TypeError("Expected argument 'webs' to be a list")
        pulumi.set(__self__, "webs", webs)

    @property
    @pulumi.getter
    def actions(self) -> Sequence[str]:
        """
        An array of IP addresses in CIDR format specifying the addresses that incoming requests from GitHub actions will originate from.
        """
        return pulumi.get(self, "actions")

    @property
    @pulumi.getter(name="actionsIpv4s")
    def actions_ipv4s(self) -> Sequence[str]:
        """
        A subset of the `actions` array that contains IP addresses in IPv4 CIDR format.
        """
        return pulumi.get(self, "actions_ipv4s")

    @property
    @pulumi.getter(name="actionsIpv6s")
    def actions_ipv6s(self) -> Sequence[str]:
        """
        A subset of the `actions` array that contains IP addresses in IPv6 CIDR format.
        """
        return pulumi.get(self, "actions_ipv6s")

    @property
    @pulumi.getter(name="apiIpv4s")
    def api_ipv4s(self) -> Sequence[str]:
        """
        A subset of the `api` array that contains IP addresses in IPv4 CIDR format.
        """
        return pulumi.get(self, "api_ipv4s")

    @property
    @pulumi.getter(name="apiIpv6s")
    def api_ipv6s(self) -> Sequence[str]:
        """
        A subset of the `api` array that contains IP addresses in IPv6 CIDR format.
        """
        return pulumi.get(self, "api_ipv6s")

    @property
    @pulumi.getter
    def apis(self) -> Sequence[str]:
        """
        An Array of IP addresses in CIDR format for the GitHub API.
        """
        return pulumi.get(self, "apis")

    @property
    @pulumi.getter(name="dependabotIpv4s")
    def dependabot_ipv4s(self) -> Sequence[str]:
        """
        A subset of the `dependabot` array that contains IP addresses in IPv4 CIDR format.
        """
        return pulumi.get(self, "dependabot_ipv4s")

    @property
    @pulumi.getter(name="dependabotIpv6s")
    def dependabot_ipv6s(self) -> Sequence[str]:
        """
        A subset of the `dependabot` array that contains IP addresses in IPv6 CIDR format.
        """
        return pulumi.get(self, "dependabot_ipv6s")

    @property
    @pulumi.getter
    def dependabots(self) -> Sequence[str]:
        """
        An array of IP addresses in CIDR format specifying the A records for dependabot.
        """
        return pulumi.get(self, "dependabots")

    @property
    @pulumi.getter(name="gitIpv4s")
    def git_ipv4s(self) -> Sequence[str]:
        """
        A subset of the `git` array that contains IP addresses in IPv4 CIDR format.
        """
        return pulumi.get(self, "git_ipv4s")

    @property
    @pulumi.getter(name="gitIpv6s")
    def git_ipv6s(self) -> Sequence[str]:
        """
        A subset of the `git` array that contains IP addresses in IPv6 CIDR format.
        """
        return pulumi.get(self, "git_ipv6s")

    @property
    @pulumi.getter
    def gits(self) -> Sequence[str]:
        """
        An Array of IP addresses in CIDR format specifying the Git servers.
        """
        return pulumi.get(self, "gits")

    @property
    @pulumi.getter
    def hooks(self) -> Sequence[str]:
        """
        An Array of IP addresses in CIDR format specifying the addresses that incoming service hooks will originate from.
        """
        return pulumi.get(self, "hooks")

    @property
    @pulumi.getter(name="hooksIpv4s")
    def hooks_ipv4s(self) -> Sequence[str]:
        """
        A subset of the `hooks` array that contains IP addresses in IPv4 CIDR format.
        """
        return pulumi.get(self, "hooks_ipv4s")

    @property
    @pulumi.getter(name="hooksIpv6s")
    def hooks_ipv6s(self) -> Sequence[str]:
        """
        A subset of the `hooks` array that contains IP addresses in IPv6 CIDR format.
        """
        return pulumi.get(self, "hooks_ipv6s")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="importerIpv4s")
    def importer_ipv4s(self) -> Sequence[str]:
        """
        A subset of the `importer` array that contains IP addresses in IPv4 CIDR format.
        """
        return pulumi.get(self, "importer_ipv4s")

    @property
    @pulumi.getter(name="importerIpv6s")
    def importer_ipv6s(self) -> Sequence[str]:
        """
        A subset of the `importer` array that contains IP addresses in IPv6 CIDR format.
        """
        return pulumi.get(self, "importer_ipv6s")

    @property
    @pulumi.getter
    def importers(self) -> Sequence[str]:
        """
        An Array of IP addresses in CIDR format specifying the A records for GitHub Importer.
        """
        return pulumi.get(self, "importers")

    @property
    @pulumi.getter
    def pages(self) -> Sequence[str]:
        """
        An Array of IP addresses in CIDR format specifying the A records for GitHub Pages.
        """
        return pulumi.get(self, "pages")

    @property
    @pulumi.getter(name="pagesIpv4s")
    def pages_ipv4s(self) -> Sequence[str]:
        """
        A subset of the `pages` array that contains IP addresses in IPv4 CIDR format.
        """
        return pulumi.get(self, "pages_ipv4s")

    @property
    @pulumi.getter(name="pagesIpv6s")
    def pages_ipv6s(self) -> Sequence[str]:
        """
        A subset of the `pages` array that contains IP addresses in IPv6 CIDR format.
        """
        return pulumi.get(self, "pages_ipv6s")

    @property
    @pulumi.getter(name="webIpv4s")
    def web_ipv4s(self) -> Sequence[str]:
        """
        A subset of the `web` array that contains IP addresses in IPv4 CIDR format.
        """
        return pulumi.get(self, "web_ipv4s")

    @property
    @pulumi.getter(name="webIpv6s")
    def web_ipv6s(self) -> Sequence[str]:
        """
        A subset of the `web` array that contains IP addresses in IPv6 CIDR format.
        """
        return pulumi.get(self, "web_ipv6s")

    @property
    @pulumi.getter
    def webs(self) -> Sequence[str]:
        """
        An Array of IP addresses in CIDR format for GitHub Web.
        """
        return pulumi.get(self, "webs")


class AwaitableGetIpRangesResult(GetIpRangesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetIpRangesResult(
            actions=self.actions,
            actions_ipv4s=self.actions_ipv4s,
            actions_ipv6s=self.actions_ipv6s,
            api_ipv4s=self.api_ipv4s,
            api_ipv6s=self.api_ipv6s,
            apis=self.apis,
            dependabot_ipv4s=self.dependabot_ipv4s,
            dependabot_ipv6s=self.dependabot_ipv6s,
            dependabots=self.dependabots,
            git_ipv4s=self.git_ipv4s,
            git_ipv6s=self.git_ipv6s,
            gits=self.gits,
            hooks=self.hooks,
            hooks_ipv4s=self.hooks_ipv4s,
            hooks_ipv6s=self.hooks_ipv6s,
            id=self.id,
            importer_ipv4s=self.importer_ipv4s,
            importer_ipv6s=self.importer_ipv6s,
            importers=self.importers,
            pages=self.pages,
            pages_ipv4s=self.pages_ipv4s,
            pages_ipv6s=self.pages_ipv6s,
            web_ipv4s=self.web_ipv4s,
            web_ipv6s=self.web_ipv6s,
            webs=self.webs)


def get_ip_ranges(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetIpRangesResult:
    """
    Use this data source to retrieve information about GitHub's IP addresses.
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('github:index/getIpRanges:getIpRanges', __args__, opts=opts, typ=GetIpRangesResult).value

    return AwaitableGetIpRangesResult(
        actions=pulumi.get(__ret__, 'actions'),
        actions_ipv4s=pulumi.get(__ret__, 'actions_ipv4s'),
        actions_ipv6s=pulumi.get(__ret__, 'actions_ipv6s'),
        api_ipv4s=pulumi.get(__ret__, 'api_ipv4s'),
        api_ipv6s=pulumi.get(__ret__, 'api_ipv6s'),
        apis=pulumi.get(__ret__, 'apis'),
        dependabot_ipv4s=pulumi.get(__ret__, 'dependabot_ipv4s'),
        dependabot_ipv6s=pulumi.get(__ret__, 'dependabot_ipv6s'),
        dependabots=pulumi.get(__ret__, 'dependabots'),
        git_ipv4s=pulumi.get(__ret__, 'git_ipv4s'),
        git_ipv6s=pulumi.get(__ret__, 'git_ipv6s'),
        gits=pulumi.get(__ret__, 'gits'),
        hooks=pulumi.get(__ret__, 'hooks'),
        hooks_ipv4s=pulumi.get(__ret__, 'hooks_ipv4s'),
        hooks_ipv6s=pulumi.get(__ret__, 'hooks_ipv6s'),
        id=pulumi.get(__ret__, 'id'),
        importer_ipv4s=pulumi.get(__ret__, 'importer_ipv4s'),
        importer_ipv6s=pulumi.get(__ret__, 'importer_ipv6s'),
        importers=pulumi.get(__ret__, 'importers'),
        pages=pulumi.get(__ret__, 'pages'),
        pages_ipv4s=pulumi.get(__ret__, 'pages_ipv4s'),
        pages_ipv6s=pulumi.get(__ret__, 'pages_ipv6s'),
        web_ipv4s=pulumi.get(__ret__, 'web_ipv4s'),
        web_ipv6s=pulumi.get(__ret__, 'web_ipv6s'),
        webs=pulumi.get(__ret__, 'webs'))


@_utilities.lift_output_func(get_ip_ranges)
def get_ip_ranges_output(opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetIpRangesResult]:
    """
    Use this data source to retrieve information about GitHub's IP addresses.
    """
    ...
