# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['MembershipArgs', 'Membership']

@pulumi.input_type
class MembershipArgs:
    def __init__(__self__, *,
                 username: pulumi.Input[str],
                 downgrade_on_destroy: Optional[pulumi.Input[bool]] = None,
                 role: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Membership resource.
        :param pulumi.Input[str] username: The user to add to the organization.
        :param pulumi.Input[bool] downgrade_on_destroy: Defaults to `false`. If set to true,
               when this resource is destroyed, the member will not be removed
               from the organization. Instead, the member's role will be
               downgraded to 'member'.
        :param pulumi.Input[str] role: The role of the user within the organization.
               Must be one of `member` or `admin`. Defaults to `member`.
               `admin` role represents the `owner` role available via GitHub UI.
        """
        pulumi.set(__self__, "username", username)
        if downgrade_on_destroy is not None:
            pulumi.set(__self__, "downgrade_on_destroy", downgrade_on_destroy)
        if role is not None:
            pulumi.set(__self__, "role", role)

    @property
    @pulumi.getter
    def username(self) -> pulumi.Input[str]:
        """
        The user to add to the organization.
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: pulumi.Input[str]):
        pulumi.set(self, "username", value)

    @property
    @pulumi.getter(name="downgradeOnDestroy")
    def downgrade_on_destroy(self) -> Optional[pulumi.Input[bool]]:
        """
        Defaults to `false`. If set to true,
        when this resource is destroyed, the member will not be removed
        from the organization. Instead, the member's role will be
        downgraded to 'member'.
        """
        return pulumi.get(self, "downgrade_on_destroy")

    @downgrade_on_destroy.setter
    def downgrade_on_destroy(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "downgrade_on_destroy", value)

    @property
    @pulumi.getter
    def role(self) -> Optional[pulumi.Input[str]]:
        """
        The role of the user within the organization.
        Must be one of `member` or `admin`. Defaults to `member`.
        `admin` role represents the `owner` role available via GitHub UI.
        """
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role", value)


@pulumi.input_type
class _MembershipState:
    def __init__(__self__, *,
                 downgrade_on_destroy: Optional[pulumi.Input[bool]] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Membership resources.
        :param pulumi.Input[bool] downgrade_on_destroy: Defaults to `false`. If set to true,
               when this resource is destroyed, the member will not be removed
               from the organization. Instead, the member's role will be
               downgraded to 'member'.
        :param pulumi.Input[str] role: The role of the user within the organization.
               Must be one of `member` or `admin`. Defaults to `member`.
               `admin` role represents the `owner` role available via GitHub UI.
        :param pulumi.Input[str] username: The user to add to the organization.
        """
        if downgrade_on_destroy is not None:
            pulumi.set(__self__, "downgrade_on_destroy", downgrade_on_destroy)
        if etag is not None:
            pulumi.set(__self__, "etag", etag)
        if role is not None:
            pulumi.set(__self__, "role", role)
        if username is not None:
            pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter(name="downgradeOnDestroy")
    def downgrade_on_destroy(self) -> Optional[pulumi.Input[bool]]:
        """
        Defaults to `false`. If set to true,
        when this resource is destroyed, the member will not be removed
        from the organization. Instead, the member's role will be
        downgraded to 'member'.
        """
        return pulumi.get(self, "downgrade_on_destroy")

    @downgrade_on_destroy.setter
    def downgrade_on_destroy(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "downgrade_on_destroy", value)

    @property
    @pulumi.getter
    def etag(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "etag")

    @etag.setter
    def etag(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "etag", value)

    @property
    @pulumi.getter
    def role(self) -> Optional[pulumi.Input[str]]:
        """
        The role of the user within the organization.
        Must be one of `member` or `admin`. Defaults to `member`.
        `admin` role represents the `owner` role available via GitHub UI.
        """
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role", value)

    @property
    @pulumi.getter
    def username(self) -> Optional[pulumi.Input[str]]:
        """
        The user to add to the organization.
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username", value)


class Membership(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 downgrade_on_destroy: Optional[pulumi.Input[bool]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a GitHub membership resource.

        This resource allows you to add/remove users from your organization. When applied,
        an invitation will be sent to the user to become part of the organization. When
        destroyed, either the invitation will be cancelled or the user will be removed.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_github as github

        # Add a user to the organization
        membership_for_some_user = github.Membership("membershipForSomeUser",
            role="member",
            username="SomeUser")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        GitHub Membership can be imported using an ID made up of `organization:username`, e.g.

        ```sh
        $ pulumi import github:index/membership:Membership member hashicorp:someuser
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] downgrade_on_destroy: Defaults to `false`. If set to true,
               when this resource is destroyed, the member will not be removed
               from the organization. Instead, the member's role will be
               downgraded to 'member'.
        :param pulumi.Input[str] role: The role of the user within the organization.
               Must be one of `member` or `admin`. Defaults to `member`.
               `admin` role represents the `owner` role available via GitHub UI.
        :param pulumi.Input[str] username: The user to add to the organization.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: MembershipArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a GitHub membership resource.

        This resource allows you to add/remove users from your organization. When applied,
        an invitation will be sent to the user to become part of the organization. When
        destroyed, either the invitation will be cancelled or the user will be removed.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_github as github

        # Add a user to the organization
        membership_for_some_user = github.Membership("membershipForSomeUser",
            role="member",
            username="SomeUser")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        GitHub Membership can be imported using an ID made up of `organization:username`, e.g.

        ```sh
        $ pulumi import github:index/membership:Membership member hashicorp:someuser
        ```

        :param str resource_name: The name of the resource.
        :param MembershipArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(MembershipArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 downgrade_on_destroy: Optional[pulumi.Input[bool]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = MembershipArgs.__new__(MembershipArgs)

            __props__.__dict__["downgrade_on_destroy"] = downgrade_on_destroy
            __props__.__dict__["role"] = role
            if username is None and not opts.urn:
                raise TypeError("Missing required property 'username'")
            __props__.__dict__["username"] = username
            __props__.__dict__["etag"] = None
        super(Membership, __self__).__init__(
            'github:index/membership:Membership',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            downgrade_on_destroy: Optional[pulumi.Input[bool]] = None,
            etag: Optional[pulumi.Input[str]] = None,
            role: Optional[pulumi.Input[str]] = None,
            username: Optional[pulumi.Input[str]] = None) -> 'Membership':
        """
        Get an existing Membership resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] downgrade_on_destroy: Defaults to `false`. If set to true,
               when this resource is destroyed, the member will not be removed
               from the organization. Instead, the member's role will be
               downgraded to 'member'.
        :param pulumi.Input[str] role: The role of the user within the organization.
               Must be one of `member` or `admin`. Defaults to `member`.
               `admin` role represents the `owner` role available via GitHub UI.
        :param pulumi.Input[str] username: The user to add to the organization.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _MembershipState.__new__(_MembershipState)

        __props__.__dict__["downgrade_on_destroy"] = downgrade_on_destroy
        __props__.__dict__["etag"] = etag
        __props__.__dict__["role"] = role
        __props__.__dict__["username"] = username
        return Membership(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="downgradeOnDestroy")
    def downgrade_on_destroy(self) -> pulumi.Output[Optional[bool]]:
        """
        Defaults to `false`. If set to true,
        when this resource is destroyed, the member will not be removed
        from the organization. Instead, the member's role will be
        downgraded to 'member'.
        """
        return pulumi.get(self, "downgrade_on_destroy")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def role(self) -> pulumi.Output[Optional[str]]:
        """
        The role of the user within the organization.
        Must be one of `member` or `admin`. Defaults to `member`.
        `admin` role represents the `owner` role available via GitHub UI.
        """
        return pulumi.get(self, "role")

    @property
    @pulumi.getter
    def username(self) -> pulumi.Output[str]:
        """
        The user to add to the organization.
        """
        return pulumi.get(self, "username")

