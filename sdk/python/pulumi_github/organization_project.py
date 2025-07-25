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

__all__ = ['OrganizationProjectArgs', 'OrganizationProject']

@pulumi.input_type
class OrganizationProjectArgs:
    def __init__(__self__, *,
                 body: Optional[pulumi.Input[_builtins.str]] = None,
                 name: Optional[pulumi.Input[_builtins.str]] = None):
        """
        The set of arguments for constructing a OrganizationProject resource.
        :param pulumi.Input[_builtins.str] body: The body of the project.
        :param pulumi.Input[_builtins.str] name: The name of the project.
        """
        if body is not None:
            pulumi.set(__self__, "body", body)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @_builtins.property
    @pulumi.getter
    def body(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The body of the project.
        """
        return pulumi.get(self, "body")

    @body.setter
    def body(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "body", value)

    @_builtins.property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The name of the project.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _OrganizationProjectState:
    def __init__(__self__, *,
                 body: Optional[pulumi.Input[_builtins.str]] = None,
                 etag: Optional[pulumi.Input[_builtins.str]] = None,
                 name: Optional[pulumi.Input[_builtins.str]] = None,
                 url: Optional[pulumi.Input[_builtins.str]] = None):
        """
        Input properties used for looking up and filtering OrganizationProject resources.
        :param pulumi.Input[_builtins.str] body: The body of the project.
        :param pulumi.Input[_builtins.str] name: The name of the project.
        :param pulumi.Input[_builtins.str] url: URL of the project
        """
        if body is not None:
            pulumi.set(__self__, "body", body)
        if etag is not None:
            pulumi.set(__self__, "etag", etag)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if url is not None:
            pulumi.set(__self__, "url", url)

    @_builtins.property
    @pulumi.getter
    def body(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The body of the project.
        """
        return pulumi.get(self, "body")

    @body.setter
    def body(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "body", value)

    @_builtins.property
    @pulumi.getter
    def etag(self) -> Optional[pulumi.Input[_builtins.str]]:
        return pulumi.get(self, "etag")

    @etag.setter
    def etag(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "etag", value)

    @_builtins.property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The name of the project.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "name", value)

    @_builtins.property
    @pulumi.getter
    def url(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        URL of the project
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "url", value)


@pulumi.type_token("github:index/organizationProject:OrganizationProject")
class OrganizationProject(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 body: Optional[pulumi.Input[_builtins.str]] = None,
                 name: Optional[pulumi.Input[_builtins.str]] = None,
                 __props__=None):
        """
        This resource allows you to create and manage projects for GitHub organization.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_github as github

        project = github.OrganizationProject("project",
            name="A Organization Project",
            body="This is a organization project.")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[_builtins.str] body: The body of the project.
        :param pulumi.Input[_builtins.str] name: The name of the project.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[OrganizationProjectArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This resource allows you to create and manage projects for GitHub organization.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_github as github

        project = github.OrganizationProject("project",
            name="A Organization Project",
            body="This is a organization project.")
        ```

        :param str resource_name: The name of the resource.
        :param OrganizationProjectArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(OrganizationProjectArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 body: Optional[pulumi.Input[_builtins.str]] = None,
                 name: Optional[pulumi.Input[_builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = OrganizationProjectArgs.__new__(OrganizationProjectArgs)

            __props__.__dict__["body"] = body
            __props__.__dict__["name"] = name
            __props__.__dict__["etag"] = None
            __props__.__dict__["url"] = None
        super(OrganizationProject, __self__).__init__(
            'github:index/organizationProject:OrganizationProject',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            body: Optional[pulumi.Input[_builtins.str]] = None,
            etag: Optional[pulumi.Input[_builtins.str]] = None,
            name: Optional[pulumi.Input[_builtins.str]] = None,
            url: Optional[pulumi.Input[_builtins.str]] = None) -> 'OrganizationProject':
        """
        Get an existing OrganizationProject resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[_builtins.str] body: The body of the project.
        :param pulumi.Input[_builtins.str] name: The name of the project.
        :param pulumi.Input[_builtins.str] url: URL of the project
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _OrganizationProjectState.__new__(_OrganizationProjectState)

        __props__.__dict__["body"] = body
        __props__.__dict__["etag"] = etag
        __props__.__dict__["name"] = name
        __props__.__dict__["url"] = url
        return OrganizationProject(resource_name, opts=opts, __props__=__props__)

    @_builtins.property
    @pulumi.getter
    def body(self) -> pulumi.Output[Optional[_builtins.str]]:
        """
        The body of the project.
        """
        return pulumi.get(self, "body")

    @_builtins.property
    @pulumi.getter
    def etag(self) -> pulumi.Output[_builtins.str]:
        return pulumi.get(self, "etag")

    @_builtins.property
    @pulumi.getter
    def name(self) -> pulumi.Output[_builtins.str]:
        """
        The name of the project.
        """
        return pulumi.get(self, "name")

    @_builtins.property
    @pulumi.getter
    def url(self) -> pulumi.Output[_builtins.str]:
        """
        URL of the project
        """
        return pulumi.get(self, "url")

