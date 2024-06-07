# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['TeamArgs', 'Team']

@pulumi.input_type
class TeamArgs:
    def __init__(__self__, *,
                 create_default_maintainer: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 ldap_dn: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parent_team_id: Optional[pulumi.Input[str]] = None,
                 parent_team_read_id: Optional[pulumi.Input[str]] = None,
                 parent_team_read_slug: Optional[pulumi.Input[str]] = None,
                 privacy: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Team resource.
        :param pulumi.Input[bool] create_default_maintainer: Adds a default maintainer to the team. Defaults to `false` and adds the creating user to the team when `true`.
        :param pulumi.Input[str] description: A description of the team.
        :param pulumi.Input[str] ldap_dn: The LDAP Distinguished Name of the group where membership will be synchronized. Only available in GitHub Enterprise Server.
        :param pulumi.Input[str] name: The name of the team.
        :param pulumi.Input[str] parent_team_id: The ID or slug of the parent team, if this is a nested team.
        :param pulumi.Input[str] parent_team_read_id: The id of the parent team read in Github.
        :param pulumi.Input[str] parent_team_read_slug: The id of the parent team read in Github.
        :param pulumi.Input[str] privacy: The level of privacy for the team. Must be one of `secret` or `closed`.
               Defaults to `secret`.
        """
        if create_default_maintainer is not None:
            pulumi.set(__self__, "create_default_maintainer", create_default_maintainer)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if ldap_dn is not None:
            pulumi.set(__self__, "ldap_dn", ldap_dn)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if parent_team_id is not None:
            pulumi.set(__self__, "parent_team_id", parent_team_id)
        if parent_team_read_id is not None:
            pulumi.set(__self__, "parent_team_read_id", parent_team_read_id)
        if parent_team_read_slug is not None:
            pulumi.set(__self__, "parent_team_read_slug", parent_team_read_slug)
        if privacy is not None:
            pulumi.set(__self__, "privacy", privacy)

    @property
    @pulumi.getter(name="createDefaultMaintainer")
    def create_default_maintainer(self) -> Optional[pulumi.Input[bool]]:
        """
        Adds a default maintainer to the team. Defaults to `false` and adds the creating user to the team when `true`.
        """
        return pulumi.get(self, "create_default_maintainer")

    @create_default_maintainer.setter
    def create_default_maintainer(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "create_default_maintainer", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A description of the team.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="ldapDn")
    def ldap_dn(self) -> Optional[pulumi.Input[str]]:
        """
        The LDAP Distinguished Name of the group where membership will be synchronized. Only available in GitHub Enterprise Server.
        """
        return pulumi.get(self, "ldap_dn")

    @ldap_dn.setter
    def ldap_dn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ldap_dn", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the team.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="parentTeamId")
    def parent_team_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID or slug of the parent team, if this is a nested team.
        """
        return pulumi.get(self, "parent_team_id")

    @parent_team_id.setter
    def parent_team_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "parent_team_id", value)

    @property
    @pulumi.getter(name="parentTeamReadId")
    def parent_team_read_id(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the parent team read in Github.
        """
        return pulumi.get(self, "parent_team_read_id")

    @parent_team_read_id.setter
    def parent_team_read_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "parent_team_read_id", value)

    @property
    @pulumi.getter(name="parentTeamReadSlug")
    def parent_team_read_slug(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the parent team read in Github.
        """
        return pulumi.get(self, "parent_team_read_slug")

    @parent_team_read_slug.setter
    def parent_team_read_slug(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "parent_team_read_slug", value)

    @property
    @pulumi.getter
    def privacy(self) -> Optional[pulumi.Input[str]]:
        """
        The level of privacy for the team. Must be one of `secret` or `closed`.
        Defaults to `secret`.
        """
        return pulumi.get(self, "privacy")

    @privacy.setter
    def privacy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "privacy", value)


@pulumi.input_type
class _TeamState:
    def __init__(__self__, *,
                 create_default_maintainer: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 ldap_dn: Optional[pulumi.Input[str]] = None,
                 members_count: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 node_id: Optional[pulumi.Input[str]] = None,
                 parent_team_id: Optional[pulumi.Input[str]] = None,
                 parent_team_read_id: Optional[pulumi.Input[str]] = None,
                 parent_team_read_slug: Optional[pulumi.Input[str]] = None,
                 privacy: Optional[pulumi.Input[str]] = None,
                 slug: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Team resources.
        :param pulumi.Input[bool] create_default_maintainer: Adds a default maintainer to the team. Defaults to `false` and adds the creating user to the team when `true`.
        :param pulumi.Input[str] description: A description of the team.
        :param pulumi.Input[str] ldap_dn: The LDAP Distinguished Name of the group where membership will be synchronized. Only available in GitHub Enterprise Server.
        :param pulumi.Input[str] name: The name of the team.
        :param pulumi.Input[str] node_id: The Node ID of the created team.
        :param pulumi.Input[str] parent_team_id: The ID or slug of the parent team, if this is a nested team.
        :param pulumi.Input[str] parent_team_read_id: The id of the parent team read in Github.
        :param pulumi.Input[str] parent_team_read_slug: The id of the parent team read in Github.
        :param pulumi.Input[str] privacy: The level of privacy for the team. Must be one of `secret` or `closed`.
               Defaults to `secret`.
        :param pulumi.Input[str] slug: The slug of the created team, which may or may not differ from `name`,
               depending on whether `name` contains "URL-unsafe" characters.
               Useful when referencing the team in [`BranchProtection`](https://www.terraform.io/docs/providers/github/r/branch_protection.html).
        """
        if create_default_maintainer is not None:
            pulumi.set(__self__, "create_default_maintainer", create_default_maintainer)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if etag is not None:
            pulumi.set(__self__, "etag", etag)
        if ldap_dn is not None:
            pulumi.set(__self__, "ldap_dn", ldap_dn)
        if members_count is not None:
            pulumi.set(__self__, "members_count", members_count)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if node_id is not None:
            pulumi.set(__self__, "node_id", node_id)
        if parent_team_id is not None:
            pulumi.set(__self__, "parent_team_id", parent_team_id)
        if parent_team_read_id is not None:
            pulumi.set(__self__, "parent_team_read_id", parent_team_read_id)
        if parent_team_read_slug is not None:
            pulumi.set(__self__, "parent_team_read_slug", parent_team_read_slug)
        if privacy is not None:
            pulumi.set(__self__, "privacy", privacy)
        if slug is not None:
            pulumi.set(__self__, "slug", slug)

    @property
    @pulumi.getter(name="createDefaultMaintainer")
    def create_default_maintainer(self) -> Optional[pulumi.Input[bool]]:
        """
        Adds a default maintainer to the team. Defaults to `false` and adds the creating user to the team when `true`.
        """
        return pulumi.get(self, "create_default_maintainer")

    @create_default_maintainer.setter
    def create_default_maintainer(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "create_default_maintainer", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A description of the team.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def etag(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "etag")

    @etag.setter
    def etag(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "etag", value)

    @property
    @pulumi.getter(name="ldapDn")
    def ldap_dn(self) -> Optional[pulumi.Input[str]]:
        """
        The LDAP Distinguished Name of the group where membership will be synchronized. Only available in GitHub Enterprise Server.
        """
        return pulumi.get(self, "ldap_dn")

    @ldap_dn.setter
    def ldap_dn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ldap_dn", value)

    @property
    @pulumi.getter(name="membersCount")
    def members_count(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "members_count")

    @members_count.setter
    def members_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "members_count", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the team.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="nodeId")
    def node_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Node ID of the created team.
        """
        return pulumi.get(self, "node_id")

    @node_id.setter
    def node_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "node_id", value)

    @property
    @pulumi.getter(name="parentTeamId")
    def parent_team_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID or slug of the parent team, if this is a nested team.
        """
        return pulumi.get(self, "parent_team_id")

    @parent_team_id.setter
    def parent_team_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "parent_team_id", value)

    @property
    @pulumi.getter(name="parentTeamReadId")
    def parent_team_read_id(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the parent team read in Github.
        """
        return pulumi.get(self, "parent_team_read_id")

    @parent_team_read_id.setter
    def parent_team_read_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "parent_team_read_id", value)

    @property
    @pulumi.getter(name="parentTeamReadSlug")
    def parent_team_read_slug(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the parent team read in Github.
        """
        return pulumi.get(self, "parent_team_read_slug")

    @parent_team_read_slug.setter
    def parent_team_read_slug(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "parent_team_read_slug", value)

    @property
    @pulumi.getter
    def privacy(self) -> Optional[pulumi.Input[str]]:
        """
        The level of privacy for the team. Must be one of `secret` or `closed`.
        Defaults to `secret`.
        """
        return pulumi.get(self, "privacy")

    @privacy.setter
    def privacy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "privacy", value)

    @property
    @pulumi.getter
    def slug(self) -> Optional[pulumi.Input[str]]:
        """
        The slug of the created team, which may or may not differ from `name`,
        depending on whether `name` contains "URL-unsafe" characters.
        Useful when referencing the team in [`BranchProtection`](https://www.terraform.io/docs/providers/github/r/branch_protection.html).
        """
        return pulumi.get(self, "slug")

    @slug.setter
    def slug(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "slug", value)


class Team(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 create_default_maintainer: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 ldap_dn: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parent_team_id: Optional[pulumi.Input[str]] = None,
                 parent_team_read_id: Optional[pulumi.Input[str]] = None,
                 parent_team_read_slug: Optional[pulumi.Input[str]] = None,
                 privacy: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a GitHub team resource.

        This resource allows you to add/remove teams from your organization. When applied,
        a new team will be created. When destroyed, that team will be removed.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_github as github

        # Add a team to the organization
        some_team = github.Team("some_team",
            name="some-team",
            description="Some cool team",
            privacy="closed")
        ```

        ## Import

        GitHub Teams can be imported using the GitHub team ID or name e.g.

        ```sh
        $ pulumi import github:index/team:Team core 1234567
        ```

        ```sh
        $ pulumi import github:index/team:Team core Administrators
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] create_default_maintainer: Adds a default maintainer to the team. Defaults to `false` and adds the creating user to the team when `true`.
        :param pulumi.Input[str] description: A description of the team.
        :param pulumi.Input[str] ldap_dn: The LDAP Distinguished Name of the group where membership will be synchronized. Only available in GitHub Enterprise Server.
        :param pulumi.Input[str] name: The name of the team.
        :param pulumi.Input[str] parent_team_id: The ID or slug of the parent team, if this is a nested team.
        :param pulumi.Input[str] parent_team_read_id: The id of the parent team read in Github.
        :param pulumi.Input[str] parent_team_read_slug: The id of the parent team read in Github.
        :param pulumi.Input[str] privacy: The level of privacy for the team. Must be one of `secret` or `closed`.
               Defaults to `secret`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[TeamArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a GitHub team resource.

        This resource allows you to add/remove teams from your organization. When applied,
        a new team will be created. When destroyed, that team will be removed.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_github as github

        # Add a team to the organization
        some_team = github.Team("some_team",
            name="some-team",
            description="Some cool team",
            privacy="closed")
        ```

        ## Import

        GitHub Teams can be imported using the GitHub team ID or name e.g.

        ```sh
        $ pulumi import github:index/team:Team core 1234567
        ```

        ```sh
        $ pulumi import github:index/team:Team core Administrators
        ```

        :param str resource_name: The name of the resource.
        :param TeamArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TeamArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 create_default_maintainer: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 ldap_dn: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parent_team_id: Optional[pulumi.Input[str]] = None,
                 parent_team_read_id: Optional[pulumi.Input[str]] = None,
                 parent_team_read_slug: Optional[pulumi.Input[str]] = None,
                 privacy: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TeamArgs.__new__(TeamArgs)

            __props__.__dict__["create_default_maintainer"] = create_default_maintainer
            __props__.__dict__["description"] = description
            __props__.__dict__["ldap_dn"] = ldap_dn
            __props__.__dict__["name"] = name
            __props__.__dict__["parent_team_id"] = parent_team_id
            __props__.__dict__["parent_team_read_id"] = parent_team_read_id
            __props__.__dict__["parent_team_read_slug"] = parent_team_read_slug
            __props__.__dict__["privacy"] = privacy
            __props__.__dict__["etag"] = None
            __props__.__dict__["members_count"] = None
            __props__.__dict__["node_id"] = None
            __props__.__dict__["slug"] = None
        super(Team, __self__).__init__(
            'github:index/team:Team',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            create_default_maintainer: Optional[pulumi.Input[bool]] = None,
            description: Optional[pulumi.Input[str]] = None,
            etag: Optional[pulumi.Input[str]] = None,
            ldap_dn: Optional[pulumi.Input[str]] = None,
            members_count: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            node_id: Optional[pulumi.Input[str]] = None,
            parent_team_id: Optional[pulumi.Input[str]] = None,
            parent_team_read_id: Optional[pulumi.Input[str]] = None,
            parent_team_read_slug: Optional[pulumi.Input[str]] = None,
            privacy: Optional[pulumi.Input[str]] = None,
            slug: Optional[pulumi.Input[str]] = None) -> 'Team':
        """
        Get an existing Team resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] create_default_maintainer: Adds a default maintainer to the team. Defaults to `false` and adds the creating user to the team when `true`.
        :param pulumi.Input[str] description: A description of the team.
        :param pulumi.Input[str] ldap_dn: The LDAP Distinguished Name of the group where membership will be synchronized. Only available in GitHub Enterprise Server.
        :param pulumi.Input[str] name: The name of the team.
        :param pulumi.Input[str] node_id: The Node ID of the created team.
        :param pulumi.Input[str] parent_team_id: The ID or slug of the parent team, if this is a nested team.
        :param pulumi.Input[str] parent_team_read_id: The id of the parent team read in Github.
        :param pulumi.Input[str] parent_team_read_slug: The id of the parent team read in Github.
        :param pulumi.Input[str] privacy: The level of privacy for the team. Must be one of `secret` or `closed`.
               Defaults to `secret`.
        :param pulumi.Input[str] slug: The slug of the created team, which may or may not differ from `name`,
               depending on whether `name` contains "URL-unsafe" characters.
               Useful when referencing the team in [`BranchProtection`](https://www.terraform.io/docs/providers/github/r/branch_protection.html).
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TeamState.__new__(_TeamState)

        __props__.__dict__["create_default_maintainer"] = create_default_maintainer
        __props__.__dict__["description"] = description
        __props__.__dict__["etag"] = etag
        __props__.__dict__["ldap_dn"] = ldap_dn
        __props__.__dict__["members_count"] = members_count
        __props__.__dict__["name"] = name
        __props__.__dict__["node_id"] = node_id
        __props__.__dict__["parent_team_id"] = parent_team_id
        __props__.__dict__["parent_team_read_id"] = parent_team_read_id
        __props__.__dict__["parent_team_read_slug"] = parent_team_read_slug
        __props__.__dict__["privacy"] = privacy
        __props__.__dict__["slug"] = slug
        return Team(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createDefaultMaintainer")
    def create_default_maintainer(self) -> pulumi.Output[Optional[bool]]:
        """
        Adds a default maintainer to the team. Defaults to `false` and adds the creating user to the team when `true`.
        """
        return pulumi.get(self, "create_default_maintainer")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A description of the team.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="ldapDn")
    def ldap_dn(self) -> pulumi.Output[Optional[str]]:
        """
        The LDAP Distinguished Name of the group where membership will be synchronized. Only available in GitHub Enterprise Server.
        """
        return pulumi.get(self, "ldap_dn")

    @property
    @pulumi.getter(name="membersCount")
    def members_count(self) -> pulumi.Output[int]:
        return pulumi.get(self, "members_count")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the team.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nodeId")
    def node_id(self) -> pulumi.Output[str]:
        """
        The Node ID of the created team.
        """
        return pulumi.get(self, "node_id")

    @property
    @pulumi.getter(name="parentTeamId")
    def parent_team_id(self) -> pulumi.Output[Optional[str]]:
        """
        The ID or slug of the parent team, if this is a nested team.
        """
        return pulumi.get(self, "parent_team_id")

    @property
    @pulumi.getter(name="parentTeamReadId")
    def parent_team_read_id(self) -> pulumi.Output[str]:
        """
        The id of the parent team read in Github.
        """
        return pulumi.get(self, "parent_team_read_id")

    @property
    @pulumi.getter(name="parentTeamReadSlug")
    def parent_team_read_slug(self) -> pulumi.Output[str]:
        """
        The id of the parent team read in Github.
        """
        return pulumi.get(self, "parent_team_read_slug")

    @property
    @pulumi.getter
    def privacy(self) -> pulumi.Output[Optional[str]]:
        """
        The level of privacy for the team. Must be one of `secret` or `closed`.
        Defaults to `secret`.
        """
        return pulumi.get(self, "privacy")

    @property
    @pulumi.getter
    def slug(self) -> pulumi.Output[str]:
        """
        The slug of the created team, which may or may not differ from `name`,
        depending on whether `name` contains "URL-unsafe" characters.
        Useful when referencing the team in [`BranchProtection`](https://www.terraform.io/docs/providers/github/r/branch_protection.html).
        """
        return pulumi.get(self, "slug")

