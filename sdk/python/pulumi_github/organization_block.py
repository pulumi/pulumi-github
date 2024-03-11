# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['OrganizationBlockArgs', 'OrganizationBlock']

@pulumi.input_type
class OrganizationBlockArgs:
    def __init__(__self__, *,
                 username: pulumi.Input[str]):
        """
        The set of arguments for constructing a OrganizationBlock resource.
        :param pulumi.Input[str] username: The name of the user to block.
        """
        pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter
    def username(self) -> pulumi.Input[str]:
        """
        The name of the user to block.
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: pulumi.Input[str]):
        pulumi.set(self, "username", value)


@pulumi.input_type
class _OrganizationBlockState:
    def __init__(__self__, *,
                 etag: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering OrganizationBlock resources.
        :param pulumi.Input[str] username: The name of the user to block.
        """
        if etag is not None:
            pulumi.set(__self__, "etag", etag)
        if username is not None:
            pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter
    def etag(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "etag")

    @etag.setter
    def etag(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "etag", value)

    @property
    @pulumi.getter
    def username(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the user to block.
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username", value)


class OrganizationBlock(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        This resource allows you to create and manage blocks for GitHub organizations.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_github as github

        example = github.OrganizationBlock("example", username="paultyng")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        GitHub organization block can be imported using a username, e.g.

        ```sh
        $ pulumi import github:index/organizationBlock:OrganizationBlock example someuser
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] username: The name of the user to block.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: OrganizationBlockArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This resource allows you to create and manage blocks for GitHub organizations.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_github as github

        example = github.OrganizationBlock("example", username="paultyng")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        GitHub organization block can be imported using a username, e.g.

        ```sh
        $ pulumi import github:index/organizationBlock:OrganizationBlock example someuser
        ```

        :param str resource_name: The name of the resource.
        :param OrganizationBlockArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(OrganizationBlockArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = OrganizationBlockArgs.__new__(OrganizationBlockArgs)

            if username is None and not opts.urn:
                raise TypeError("Missing required property 'username'")
            __props__.__dict__["username"] = username
            __props__.__dict__["etag"] = None
        super(OrganizationBlock, __self__).__init__(
            'github:index/organizationBlock:OrganizationBlock',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            etag: Optional[pulumi.Input[str]] = None,
            username: Optional[pulumi.Input[str]] = None) -> 'OrganizationBlock':
        """
        Get an existing OrganizationBlock resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] username: The name of the user to block.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _OrganizationBlockState.__new__(_OrganizationBlockState)

        __props__.__dict__["etag"] = etag
        __props__.__dict__["username"] = username
        return OrganizationBlock(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def username(self) -> pulumi.Output[str]:
        """
        The name of the user to block.
        """
        return pulumi.get(self, "username")

