# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['TeamSyncGroupMappingArgs', 'TeamSyncGroupMapping']

@pulumi.input_type
class TeamSyncGroupMappingArgs:
    def __init__(__self__, *,
                 team_slug: pulumi.Input[str],
                 groups: Optional[pulumi.Input[Sequence[pulumi.Input['TeamSyncGroupMappingGroupArgs']]]] = None):
        """
        The set of arguments for constructing a TeamSyncGroupMapping resource.
        :param pulumi.Input[str] team_slug: Slug of the team
        :param pulumi.Input[Sequence[pulumi.Input['TeamSyncGroupMappingGroupArgs']]] groups: An Array of GitHub Identity Provider Groups (or empty []).  Each `group` block consists of the fields documented below.
               ___
        """
        pulumi.set(__self__, "team_slug", team_slug)
        if groups is not None:
            pulumi.set(__self__, "groups", groups)

    @property
    @pulumi.getter(name="teamSlug")
    def team_slug(self) -> pulumi.Input[str]:
        """
        Slug of the team
        """
        return pulumi.get(self, "team_slug")

    @team_slug.setter
    def team_slug(self, value: pulumi.Input[str]):
        pulumi.set(self, "team_slug", value)

    @property
    @pulumi.getter
    def groups(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['TeamSyncGroupMappingGroupArgs']]]]:
        """
        An Array of GitHub Identity Provider Groups (or empty []).  Each `group` block consists of the fields documented below.
        ___
        """
        return pulumi.get(self, "groups")

    @groups.setter
    def groups(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['TeamSyncGroupMappingGroupArgs']]]]):
        pulumi.set(self, "groups", value)


@pulumi.input_type
class _TeamSyncGroupMappingState:
    def __init__(__self__, *,
                 etag: Optional[pulumi.Input[str]] = None,
                 groups: Optional[pulumi.Input[Sequence[pulumi.Input['TeamSyncGroupMappingGroupArgs']]]] = None,
                 team_slug: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering TeamSyncGroupMapping resources.
        :param pulumi.Input[Sequence[pulumi.Input['TeamSyncGroupMappingGroupArgs']]] groups: An Array of GitHub Identity Provider Groups (or empty []).  Each `group` block consists of the fields documented below.
               ___
        :param pulumi.Input[str] team_slug: Slug of the team
        """
        if etag is not None:
            pulumi.set(__self__, "etag", etag)
        if groups is not None:
            pulumi.set(__self__, "groups", groups)
        if team_slug is not None:
            pulumi.set(__self__, "team_slug", team_slug)

    @property
    @pulumi.getter
    def etag(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "etag")

    @etag.setter
    def etag(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "etag", value)

    @property
    @pulumi.getter
    def groups(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['TeamSyncGroupMappingGroupArgs']]]]:
        """
        An Array of GitHub Identity Provider Groups (or empty []).  Each `group` block consists of the fields documented below.
        ___
        """
        return pulumi.get(self, "groups")

    @groups.setter
    def groups(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['TeamSyncGroupMappingGroupArgs']]]]):
        pulumi.set(self, "groups", value)

    @property
    @pulumi.getter(name="teamSlug")
    def team_slug(self) -> Optional[pulumi.Input[str]]:
        """
        Slug of the team
        """
        return pulumi.get(self, "team_slug")

    @team_slug.setter
    def team_slug(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "team_slug", value)


class TeamSyncGroupMapping(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 groups: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TeamSyncGroupMappingGroupArgs']]]]] = None,
                 team_slug: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        This resource allows you to create and manage Identity Provider (IdP) group connections within your GitHub teams.
        You must have team synchronization enabled for organizations owned by enterprise accounts.

        To learn more about team synchronization between IdPs and GitHub, please refer to:
        https://help.github.com/en/github/setting-up-and-managing-organizations-and-teams/synchronizing-teams-between-your-identity-provider-and-github

        ## Import

        GitHub Team Sync Group Mappings can be imported using the GitHub team `slug` e.g.

        ```sh
         $ pulumi import github:index/teamSyncGroupMapping:TeamSyncGroupMapping example some_team
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TeamSyncGroupMappingGroupArgs']]]] groups: An Array of GitHub Identity Provider Groups (or empty []).  Each `group` block consists of the fields documented below.
               ___
        :param pulumi.Input[str] team_slug: Slug of the team
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TeamSyncGroupMappingArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This resource allows you to create and manage Identity Provider (IdP) group connections within your GitHub teams.
        You must have team synchronization enabled for organizations owned by enterprise accounts.

        To learn more about team synchronization between IdPs and GitHub, please refer to:
        https://help.github.com/en/github/setting-up-and-managing-organizations-and-teams/synchronizing-teams-between-your-identity-provider-and-github

        ## Import

        GitHub Team Sync Group Mappings can be imported using the GitHub team `slug` e.g.

        ```sh
         $ pulumi import github:index/teamSyncGroupMapping:TeamSyncGroupMapping example some_team
        ```

        :param str resource_name: The name of the resource.
        :param TeamSyncGroupMappingArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TeamSyncGroupMappingArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 groups: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TeamSyncGroupMappingGroupArgs']]]]] = None,
                 team_slug: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TeamSyncGroupMappingArgs.__new__(TeamSyncGroupMappingArgs)

            __props__.__dict__["groups"] = groups
            if team_slug is None and not opts.urn:
                raise TypeError("Missing required property 'team_slug'")
            __props__.__dict__["team_slug"] = team_slug
            __props__.__dict__["etag"] = None
        super(TeamSyncGroupMapping, __self__).__init__(
            'github:index/teamSyncGroupMapping:TeamSyncGroupMapping',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            etag: Optional[pulumi.Input[str]] = None,
            groups: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TeamSyncGroupMappingGroupArgs']]]]] = None,
            team_slug: Optional[pulumi.Input[str]] = None) -> 'TeamSyncGroupMapping':
        """
        Get an existing TeamSyncGroupMapping resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TeamSyncGroupMappingGroupArgs']]]] groups: An Array of GitHub Identity Provider Groups (or empty []).  Each `group` block consists of the fields documented below.
               ___
        :param pulumi.Input[str] team_slug: Slug of the team
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TeamSyncGroupMappingState.__new__(_TeamSyncGroupMappingState)

        __props__.__dict__["etag"] = etag
        __props__.__dict__["groups"] = groups
        __props__.__dict__["team_slug"] = team_slug
        return TeamSyncGroupMapping(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def groups(self) -> pulumi.Output[Optional[Sequence['outputs.TeamSyncGroupMappingGroup']]]:
        """
        An Array of GitHub Identity Provider Groups (or empty []).  Each `group` block consists of the fields documented below.
        ___
        """
        return pulumi.get(self, "groups")

    @property
    @pulumi.getter(name="teamSlug")
    def team_slug(self) -> pulumi.Output[str]:
        """
        Slug of the team
        """
        return pulumi.get(self, "team_slug")

