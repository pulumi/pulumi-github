# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['EnterpriseOrganizationArgs', 'EnterpriseOrganization']

@pulumi.input_type
class EnterpriseOrganizationArgs:
    def __init__(__self__, *,
                 admin_logins: pulumi.Input[Sequence[pulumi.Input[str]]],
                 billing_email: pulumi.Input[str],
                 enterprise_id: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a EnterpriseOrganization resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] admin_logins: List of organization owner usernames.
        :param pulumi.Input[str] billing_email: The billing email address.
        :param pulumi.Input[str] enterprise_id: The ID of the enterprise.
        :param pulumi.Input[str] description: The description of the organization.
        :param pulumi.Input[str] display_name: The display name of the organization.
        :param pulumi.Input[str] name: The name of the organization.
        """
        pulumi.set(__self__, "admin_logins", admin_logins)
        pulumi.set(__self__, "billing_email", billing_email)
        pulumi.set(__self__, "enterprise_id", enterprise_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="adminLogins")
    def admin_logins(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        List of organization owner usernames.
        """
        return pulumi.get(self, "admin_logins")

    @admin_logins.setter
    def admin_logins(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "admin_logins", value)

    @property
    @pulumi.getter(name="billingEmail")
    def billing_email(self) -> pulumi.Input[str]:
        """
        The billing email address.
        """
        return pulumi.get(self, "billing_email")

    @billing_email.setter
    def billing_email(self, value: pulumi.Input[str]):
        pulumi.set(self, "billing_email", value)

    @property
    @pulumi.getter(name="enterpriseId")
    def enterprise_id(self) -> pulumi.Input[str]:
        """
        The ID of the enterprise.
        """
        return pulumi.get(self, "enterprise_id")

    @enterprise_id.setter
    def enterprise_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "enterprise_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the organization.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        The display name of the organization.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the organization.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _EnterpriseOrganizationState:
    def __init__(__self__, *,
                 admin_logins: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 billing_email: Optional[pulumi.Input[str]] = None,
                 database_id: Optional[pulumi.Input[int]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 enterprise_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering EnterpriseOrganization resources.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] admin_logins: List of organization owner usernames.
        :param pulumi.Input[str] billing_email: The billing email address.
        :param pulumi.Input[int] database_id: The ID of the organization.
        :param pulumi.Input[str] description: The description of the organization.
        :param pulumi.Input[str] display_name: The display name of the organization.
        :param pulumi.Input[str] enterprise_id: The ID of the enterprise.
        :param pulumi.Input[str] name: The name of the organization.
        """
        if admin_logins is not None:
            pulumi.set(__self__, "admin_logins", admin_logins)
        if billing_email is not None:
            pulumi.set(__self__, "billing_email", billing_email)
        if database_id is not None:
            pulumi.set(__self__, "database_id", database_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if enterprise_id is not None:
            pulumi.set(__self__, "enterprise_id", enterprise_id)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="adminLogins")
    def admin_logins(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of organization owner usernames.
        """
        return pulumi.get(self, "admin_logins")

    @admin_logins.setter
    def admin_logins(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "admin_logins", value)

    @property
    @pulumi.getter(name="billingEmail")
    def billing_email(self) -> Optional[pulumi.Input[str]]:
        """
        The billing email address.
        """
        return pulumi.get(self, "billing_email")

    @billing_email.setter
    def billing_email(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "billing_email", value)

    @property
    @pulumi.getter(name="databaseId")
    def database_id(self) -> Optional[pulumi.Input[int]]:
        """
        The ID of the organization.
        """
        return pulumi.get(self, "database_id")

    @database_id.setter
    def database_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "database_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the organization.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        The display name of the organization.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="enterpriseId")
    def enterprise_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the enterprise.
        """
        return pulumi.get(self, "enterprise_id")

    @enterprise_id.setter
    def enterprise_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "enterprise_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the organization.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


class EnterpriseOrganization(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 admin_logins: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 billing_email: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 enterprise_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        This resource allows you to create and manage a GitHub enterprise organization.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_github as github

        org = github.EnterpriseOrganization("org",
            enterprise_id=data["github_enterprise"]["enterprise"]["id"],
            display_name="Some Awesome Org",
            description="Organization created with terraform",
            billing_email="jon@winteriscoming.com",
            admin_logins=["jon-snow"])
        ```

        ## Import

        GitHub Enterprise Organization can be imported using the `slug` of the enterprise, combined with the `orgname` of the organization, separated by a `/` character.

        ```sh
        $ pulumi import github:index/enterpriseOrganization:EnterpriseOrganization org enterp/some-awesome-org
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] admin_logins: List of organization owner usernames.
        :param pulumi.Input[str] billing_email: The billing email address.
        :param pulumi.Input[str] description: The description of the organization.
        :param pulumi.Input[str] display_name: The display name of the organization.
        :param pulumi.Input[str] enterprise_id: The ID of the enterprise.
        :param pulumi.Input[str] name: The name of the organization.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: EnterpriseOrganizationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This resource allows you to create and manage a GitHub enterprise organization.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_github as github

        org = github.EnterpriseOrganization("org",
            enterprise_id=data["github_enterprise"]["enterprise"]["id"],
            display_name="Some Awesome Org",
            description="Organization created with terraform",
            billing_email="jon@winteriscoming.com",
            admin_logins=["jon-snow"])
        ```

        ## Import

        GitHub Enterprise Organization can be imported using the `slug` of the enterprise, combined with the `orgname` of the organization, separated by a `/` character.

        ```sh
        $ pulumi import github:index/enterpriseOrganization:EnterpriseOrganization org enterp/some-awesome-org
        ```

        :param str resource_name: The name of the resource.
        :param EnterpriseOrganizationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EnterpriseOrganizationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 admin_logins: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 billing_email: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 enterprise_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = EnterpriseOrganizationArgs.__new__(EnterpriseOrganizationArgs)

            if admin_logins is None and not opts.urn:
                raise TypeError("Missing required property 'admin_logins'")
            __props__.__dict__["admin_logins"] = admin_logins
            if billing_email is None and not opts.urn:
                raise TypeError("Missing required property 'billing_email'")
            __props__.__dict__["billing_email"] = billing_email
            __props__.__dict__["description"] = description
            __props__.__dict__["display_name"] = display_name
            if enterprise_id is None and not opts.urn:
                raise TypeError("Missing required property 'enterprise_id'")
            __props__.__dict__["enterprise_id"] = enterprise_id
            __props__.__dict__["name"] = name
            __props__.__dict__["database_id"] = None
        super(EnterpriseOrganization, __self__).__init__(
            'github:index/enterpriseOrganization:EnterpriseOrganization',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            admin_logins: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            billing_email: Optional[pulumi.Input[str]] = None,
            database_id: Optional[pulumi.Input[int]] = None,
            description: Optional[pulumi.Input[str]] = None,
            display_name: Optional[pulumi.Input[str]] = None,
            enterprise_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None) -> 'EnterpriseOrganization':
        """
        Get an existing EnterpriseOrganization resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] admin_logins: List of organization owner usernames.
        :param pulumi.Input[str] billing_email: The billing email address.
        :param pulumi.Input[int] database_id: The ID of the organization.
        :param pulumi.Input[str] description: The description of the organization.
        :param pulumi.Input[str] display_name: The display name of the organization.
        :param pulumi.Input[str] enterprise_id: The ID of the enterprise.
        :param pulumi.Input[str] name: The name of the organization.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _EnterpriseOrganizationState.__new__(_EnterpriseOrganizationState)

        __props__.__dict__["admin_logins"] = admin_logins
        __props__.__dict__["billing_email"] = billing_email
        __props__.__dict__["database_id"] = database_id
        __props__.__dict__["description"] = description
        __props__.__dict__["display_name"] = display_name
        __props__.__dict__["enterprise_id"] = enterprise_id
        __props__.__dict__["name"] = name
        return EnterpriseOrganization(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="adminLogins")
    def admin_logins(self) -> pulumi.Output[Sequence[str]]:
        """
        List of organization owner usernames.
        """
        return pulumi.get(self, "admin_logins")

    @property
    @pulumi.getter(name="billingEmail")
    def billing_email(self) -> pulumi.Output[str]:
        """
        The billing email address.
        """
        return pulumi.get(self, "billing_email")

    @property
    @pulumi.getter(name="databaseId")
    def database_id(self) -> pulumi.Output[int]:
        """
        The ID of the organization.
        """
        return pulumi.get(self, "database_id")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the organization.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[Optional[str]]:
        """
        The display name of the organization.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="enterpriseId")
    def enterprise_id(self) -> pulumi.Output[str]:
        """
        The ID of the enterprise.
        """
        return pulumi.get(self, "enterprise_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the organization.
        """
        return pulumi.get(self, "name")

