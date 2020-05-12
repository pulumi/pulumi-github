# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class GetIpRangesResult:
    """
    A collection of values returned by getIpRanges.
    """
    def __init__(__self__, gits=None, hooks=None, id=None, importers=None, pages=None):
        if gits and not isinstance(gits, list):
            raise TypeError("Expected argument 'gits' to be a list")
        __self__.gits = gits
        """
        An Array of IP addresses in CIDR format specifying the Git servers.
        """
        if hooks and not isinstance(hooks, list):
            raise TypeError("Expected argument 'hooks' to be a list")
        __self__.hooks = hooks
        """
        An Array of IP addresses in CIDR format specifying the addresses that incoming service hooks will originate from.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if importers and not isinstance(importers, list):
            raise TypeError("Expected argument 'importers' to be a list")
        __self__.importers = importers
        """
        An Array of IP addresses in CIDR format specifying the A records for GitHub Importer.
        """
        if pages and not isinstance(pages, list):
            raise TypeError("Expected argument 'pages' to be a list")
        __self__.pages = pages
        """
        An Array of IP addresses in CIDR format specifying the A records for GitHub Pages.
        """
class AwaitableGetIpRangesResult(GetIpRangesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetIpRangesResult(
            gits=self.gits,
            hooks=self.hooks,
            id=self.id,
            importers=self.importers,
            pages=self.pages)

def get_ip_ranges(opts=None):
    """
    ## Example Usage



    ```python
    import pulumi
    import pulumi_github as github

    test = github.get_ip_ranges()
    ```
    """
    __args__ = dict()


    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('github:index/getIpRanges:getIpRanges', __args__, opts=opts).value

    return AwaitableGetIpRangesResult(
        gits=__ret__.get('gits'),
        hooks=__ret__.get('hooks'),
        id=__ret__.get('id'),
        importers=__ret__.get('importers'),
        pages=__ret__.get('pages'))
