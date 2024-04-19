# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetEnterpriseResult',
    'AwaitableGetEnterpriseResult',
    'get_enterprise',
    'get_enterprise_output',
]

@pulumi.output_type
class GetEnterpriseResult:
    """
    A collection of values returned by getEnterprise.
    """
    def __init__(__self__, created_at=None, database_id=None, description=None, id=None, name=None, slug=None, url=None):
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if database_id and not isinstance(database_id, int):
            raise TypeError("Expected argument 'database_id' to be a int")
        pulumi.set(__self__, "database_id", database_id)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if slug and not isinstance(slug, str):
            raise TypeError("Expected argument 'slug' to be a str")
        pulumi.set(__self__, "slug", slug)
        if url and not isinstance(url, str):
            raise TypeError("Expected argument 'url' to be a str")
        pulumi.set(__self__, "url", url)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        """
        The time the enterprise was created.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="databaseId")
    def database_id(self) -> int:
        """
        The database ID of the enterprise.
        """
        return pulumi.get(self, "database_id")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        The description of the enterprise.
        """
        return pulumi.get(self, "description")

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
        The name of the enterprise.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def slug(self) -> str:
        """
        The URL slug identifying the enterprise.
        """
        return pulumi.get(self, "slug")

    @property
    @pulumi.getter
    def url(self) -> str:
        """
        The url for the enterprise.
        """
        return pulumi.get(self, "url")


class AwaitableGetEnterpriseResult(GetEnterpriseResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetEnterpriseResult(
            created_at=self.created_at,
            database_id=self.database_id,
            description=self.description,
            id=self.id,
            name=self.name,
            slug=self.slug,
            url=self.url)


def get_enterprise(slug: Optional[str] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetEnterpriseResult:
    """
    Use this data source to retrieve basic information about a GitHub enterprise.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_github as github

    example = github.get_enterprise(slug="example-co")
    ```


    :param str slug: The URL slug identifying the enterprise.
    """
    __args__ = dict()
    __args__['slug'] = slug
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('github:index/getEnterprise:getEnterprise', __args__, opts=opts, typ=GetEnterpriseResult).value

    return AwaitableGetEnterpriseResult(
        created_at=pulumi.get(__ret__, 'created_at'),
        database_id=pulumi.get(__ret__, 'database_id'),
        description=pulumi.get(__ret__, 'description'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        slug=pulumi.get(__ret__, 'slug'),
        url=pulumi.get(__ret__, 'url'))


@_utilities.lift_output_func(get_enterprise)
def get_enterprise_output(slug: Optional[pulumi.Input[str]] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetEnterpriseResult]:
    """
    Use this data source to retrieve basic information about a GitHub enterprise.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_github as github

    example = github.get_enterprise(slug="example-co")
    ```


    :param str slug: The URL slug identifying the enterprise.
    """
    ...
