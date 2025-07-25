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
from . import outputs
from ._inputs import *

__all__ = ['ActionsOrganizationPermissionsArgs', 'ActionsOrganizationPermissions']

@pulumi.input_type
class ActionsOrganizationPermissionsArgs:
    def __init__(__self__, *,
                 enabled_repositories: pulumi.Input[_builtins.str],
                 allowed_actions: Optional[pulumi.Input[_builtins.str]] = None,
                 allowed_actions_config: Optional[pulumi.Input['ActionsOrganizationPermissionsAllowedActionsConfigArgs']] = None,
                 enabled_repositories_config: Optional[pulumi.Input['ActionsOrganizationPermissionsEnabledRepositoriesConfigArgs']] = None):
        """
        The set of arguments for constructing a ActionsOrganizationPermissions resource.
        :param pulumi.Input[_builtins.str] enabled_repositories: The policy that controls the repositories in the organization that are allowed to run GitHub Actions. Can be one of: `all`, `none`, or `selected`.
        :param pulumi.Input[_builtins.str] allowed_actions: The permissions policy that controls the actions that are allowed to run. Can be one of: `all`, `local_only`, or `selected`.
        :param pulumi.Input['ActionsOrganizationPermissionsAllowedActionsConfigArgs'] allowed_actions_config: Sets the actions that are allowed in an organization. Only available when `allowed_actions` = `selected`. See Allowed Actions Config below for details.
        :param pulumi.Input['ActionsOrganizationPermissionsEnabledRepositoriesConfigArgs'] enabled_repositories_config: Sets the list of selected repositories that are enabled for GitHub Actions in an organization. Only available when `enabled_repositories` = `selected`. See Enabled Repositories Config below for details.
        """
        pulumi.set(__self__, "enabled_repositories", enabled_repositories)
        if allowed_actions is not None:
            pulumi.set(__self__, "allowed_actions", allowed_actions)
        if allowed_actions_config is not None:
            pulumi.set(__self__, "allowed_actions_config", allowed_actions_config)
        if enabled_repositories_config is not None:
            pulumi.set(__self__, "enabled_repositories_config", enabled_repositories_config)

    @_builtins.property
    @pulumi.getter(name="enabledRepositories")
    def enabled_repositories(self) -> pulumi.Input[_builtins.str]:
        """
        The policy that controls the repositories in the organization that are allowed to run GitHub Actions. Can be one of: `all`, `none`, or `selected`.
        """
        return pulumi.get(self, "enabled_repositories")

    @enabled_repositories.setter
    def enabled_repositories(self, value: pulumi.Input[_builtins.str]):
        pulumi.set(self, "enabled_repositories", value)

    @_builtins.property
    @pulumi.getter(name="allowedActions")
    def allowed_actions(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The permissions policy that controls the actions that are allowed to run. Can be one of: `all`, `local_only`, or `selected`.
        """
        return pulumi.get(self, "allowed_actions")

    @allowed_actions.setter
    def allowed_actions(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "allowed_actions", value)

    @_builtins.property
    @pulumi.getter(name="allowedActionsConfig")
    def allowed_actions_config(self) -> Optional[pulumi.Input['ActionsOrganizationPermissionsAllowedActionsConfigArgs']]:
        """
        Sets the actions that are allowed in an organization. Only available when `allowed_actions` = `selected`. See Allowed Actions Config below for details.
        """
        return pulumi.get(self, "allowed_actions_config")

    @allowed_actions_config.setter
    def allowed_actions_config(self, value: Optional[pulumi.Input['ActionsOrganizationPermissionsAllowedActionsConfigArgs']]):
        pulumi.set(self, "allowed_actions_config", value)

    @_builtins.property
    @pulumi.getter(name="enabledRepositoriesConfig")
    def enabled_repositories_config(self) -> Optional[pulumi.Input['ActionsOrganizationPermissionsEnabledRepositoriesConfigArgs']]:
        """
        Sets the list of selected repositories that are enabled for GitHub Actions in an organization. Only available when `enabled_repositories` = `selected`. See Enabled Repositories Config below for details.
        """
        return pulumi.get(self, "enabled_repositories_config")

    @enabled_repositories_config.setter
    def enabled_repositories_config(self, value: Optional[pulumi.Input['ActionsOrganizationPermissionsEnabledRepositoriesConfigArgs']]):
        pulumi.set(self, "enabled_repositories_config", value)


@pulumi.input_type
class _ActionsOrganizationPermissionsState:
    def __init__(__self__, *,
                 allowed_actions: Optional[pulumi.Input[_builtins.str]] = None,
                 allowed_actions_config: Optional[pulumi.Input['ActionsOrganizationPermissionsAllowedActionsConfigArgs']] = None,
                 enabled_repositories: Optional[pulumi.Input[_builtins.str]] = None,
                 enabled_repositories_config: Optional[pulumi.Input['ActionsOrganizationPermissionsEnabledRepositoriesConfigArgs']] = None):
        """
        Input properties used for looking up and filtering ActionsOrganizationPermissions resources.
        :param pulumi.Input[_builtins.str] allowed_actions: The permissions policy that controls the actions that are allowed to run. Can be one of: `all`, `local_only`, or `selected`.
        :param pulumi.Input['ActionsOrganizationPermissionsAllowedActionsConfigArgs'] allowed_actions_config: Sets the actions that are allowed in an organization. Only available when `allowed_actions` = `selected`. See Allowed Actions Config below for details.
        :param pulumi.Input[_builtins.str] enabled_repositories: The policy that controls the repositories in the organization that are allowed to run GitHub Actions. Can be one of: `all`, `none`, or `selected`.
        :param pulumi.Input['ActionsOrganizationPermissionsEnabledRepositoriesConfigArgs'] enabled_repositories_config: Sets the list of selected repositories that are enabled for GitHub Actions in an organization. Only available when `enabled_repositories` = `selected`. See Enabled Repositories Config below for details.
        """
        if allowed_actions is not None:
            pulumi.set(__self__, "allowed_actions", allowed_actions)
        if allowed_actions_config is not None:
            pulumi.set(__self__, "allowed_actions_config", allowed_actions_config)
        if enabled_repositories is not None:
            pulumi.set(__self__, "enabled_repositories", enabled_repositories)
        if enabled_repositories_config is not None:
            pulumi.set(__self__, "enabled_repositories_config", enabled_repositories_config)

    @_builtins.property
    @pulumi.getter(name="allowedActions")
    def allowed_actions(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The permissions policy that controls the actions that are allowed to run. Can be one of: `all`, `local_only`, or `selected`.
        """
        return pulumi.get(self, "allowed_actions")

    @allowed_actions.setter
    def allowed_actions(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "allowed_actions", value)

    @_builtins.property
    @pulumi.getter(name="allowedActionsConfig")
    def allowed_actions_config(self) -> Optional[pulumi.Input['ActionsOrganizationPermissionsAllowedActionsConfigArgs']]:
        """
        Sets the actions that are allowed in an organization. Only available when `allowed_actions` = `selected`. See Allowed Actions Config below for details.
        """
        return pulumi.get(self, "allowed_actions_config")

    @allowed_actions_config.setter
    def allowed_actions_config(self, value: Optional[pulumi.Input['ActionsOrganizationPermissionsAllowedActionsConfigArgs']]):
        pulumi.set(self, "allowed_actions_config", value)

    @_builtins.property
    @pulumi.getter(name="enabledRepositories")
    def enabled_repositories(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The policy that controls the repositories in the organization that are allowed to run GitHub Actions. Can be one of: `all`, `none`, or `selected`.
        """
        return pulumi.get(self, "enabled_repositories")

    @enabled_repositories.setter
    def enabled_repositories(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "enabled_repositories", value)

    @_builtins.property
    @pulumi.getter(name="enabledRepositoriesConfig")
    def enabled_repositories_config(self) -> Optional[pulumi.Input['ActionsOrganizationPermissionsEnabledRepositoriesConfigArgs']]:
        """
        Sets the list of selected repositories that are enabled for GitHub Actions in an organization. Only available when `enabled_repositories` = `selected`. See Enabled Repositories Config below for details.
        """
        return pulumi.get(self, "enabled_repositories_config")

    @enabled_repositories_config.setter
    def enabled_repositories_config(self, value: Optional[pulumi.Input['ActionsOrganizationPermissionsEnabledRepositoriesConfigArgs']]):
        pulumi.set(self, "enabled_repositories_config", value)


@pulumi.type_token("github:index/actionsOrganizationPermissions:ActionsOrganizationPermissions")
class ActionsOrganizationPermissions(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allowed_actions: Optional[pulumi.Input[_builtins.str]] = None,
                 allowed_actions_config: Optional[pulumi.Input[Union['ActionsOrganizationPermissionsAllowedActionsConfigArgs', 'ActionsOrganizationPermissionsAllowedActionsConfigArgsDict']]] = None,
                 enabled_repositories: Optional[pulumi.Input[_builtins.str]] = None,
                 enabled_repositories_config: Optional[pulumi.Input[Union['ActionsOrganizationPermissionsEnabledRepositoriesConfigArgs', 'ActionsOrganizationPermissionsEnabledRepositoriesConfigArgsDict']]] = None,
                 __props__=None):
        """
        This resource allows you to create and manage GitHub Actions permissions within your GitHub enterprise organizations.
        You must have admin access to an organization to use this resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_github as github

        example = github.Repository("example", name="my-repository")
        test = github.ActionsOrganizationPermissions("test",
            allowed_actions="selected",
            enabled_repositories="selected",
            allowed_actions_config={
                "github_owned_allowed": True,
                "patterns_alloweds": [
                    "actions/cache@*",
                    "actions/checkout@*",
                ],
                "verified_allowed": True,
            },
            enabled_repositories_config={
                "repository_ids": [example.repo_id],
            })
        ```

        ## Import

        This resource can be imported using the name of the GitHub organization:

        ```sh
        $ pulumi import github:index/actionsOrganizationPermissions:ActionsOrganizationPermissions test github_organization_name
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[_builtins.str] allowed_actions: The permissions policy that controls the actions that are allowed to run. Can be one of: `all`, `local_only`, or `selected`.
        :param pulumi.Input[Union['ActionsOrganizationPermissionsAllowedActionsConfigArgs', 'ActionsOrganizationPermissionsAllowedActionsConfigArgsDict']] allowed_actions_config: Sets the actions that are allowed in an organization. Only available when `allowed_actions` = `selected`. See Allowed Actions Config below for details.
        :param pulumi.Input[_builtins.str] enabled_repositories: The policy that controls the repositories in the organization that are allowed to run GitHub Actions. Can be one of: `all`, `none`, or `selected`.
        :param pulumi.Input[Union['ActionsOrganizationPermissionsEnabledRepositoriesConfigArgs', 'ActionsOrganizationPermissionsEnabledRepositoriesConfigArgsDict']] enabled_repositories_config: Sets the list of selected repositories that are enabled for GitHub Actions in an organization. Only available when `enabled_repositories` = `selected`. See Enabled Repositories Config below for details.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ActionsOrganizationPermissionsArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This resource allows you to create and manage GitHub Actions permissions within your GitHub enterprise organizations.
        You must have admin access to an organization to use this resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_github as github

        example = github.Repository("example", name="my-repository")
        test = github.ActionsOrganizationPermissions("test",
            allowed_actions="selected",
            enabled_repositories="selected",
            allowed_actions_config={
                "github_owned_allowed": True,
                "patterns_alloweds": [
                    "actions/cache@*",
                    "actions/checkout@*",
                ],
                "verified_allowed": True,
            },
            enabled_repositories_config={
                "repository_ids": [example.repo_id],
            })
        ```

        ## Import

        This resource can be imported using the name of the GitHub organization:

        ```sh
        $ pulumi import github:index/actionsOrganizationPermissions:ActionsOrganizationPermissions test github_organization_name
        ```

        :param str resource_name: The name of the resource.
        :param ActionsOrganizationPermissionsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ActionsOrganizationPermissionsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allowed_actions: Optional[pulumi.Input[_builtins.str]] = None,
                 allowed_actions_config: Optional[pulumi.Input[Union['ActionsOrganizationPermissionsAllowedActionsConfigArgs', 'ActionsOrganizationPermissionsAllowedActionsConfigArgsDict']]] = None,
                 enabled_repositories: Optional[pulumi.Input[_builtins.str]] = None,
                 enabled_repositories_config: Optional[pulumi.Input[Union['ActionsOrganizationPermissionsEnabledRepositoriesConfigArgs', 'ActionsOrganizationPermissionsEnabledRepositoriesConfigArgsDict']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ActionsOrganizationPermissionsArgs.__new__(ActionsOrganizationPermissionsArgs)

            __props__.__dict__["allowed_actions"] = allowed_actions
            __props__.__dict__["allowed_actions_config"] = allowed_actions_config
            if enabled_repositories is None and not opts.urn:
                raise TypeError("Missing required property 'enabled_repositories'")
            __props__.__dict__["enabled_repositories"] = enabled_repositories
            __props__.__dict__["enabled_repositories_config"] = enabled_repositories_config
        super(ActionsOrganizationPermissions, __self__).__init__(
            'github:index/actionsOrganizationPermissions:ActionsOrganizationPermissions',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            allowed_actions: Optional[pulumi.Input[_builtins.str]] = None,
            allowed_actions_config: Optional[pulumi.Input[Union['ActionsOrganizationPermissionsAllowedActionsConfigArgs', 'ActionsOrganizationPermissionsAllowedActionsConfigArgsDict']]] = None,
            enabled_repositories: Optional[pulumi.Input[_builtins.str]] = None,
            enabled_repositories_config: Optional[pulumi.Input[Union['ActionsOrganizationPermissionsEnabledRepositoriesConfigArgs', 'ActionsOrganizationPermissionsEnabledRepositoriesConfigArgsDict']]] = None) -> 'ActionsOrganizationPermissions':
        """
        Get an existing ActionsOrganizationPermissions resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[_builtins.str] allowed_actions: The permissions policy that controls the actions that are allowed to run. Can be one of: `all`, `local_only`, or `selected`.
        :param pulumi.Input[Union['ActionsOrganizationPermissionsAllowedActionsConfigArgs', 'ActionsOrganizationPermissionsAllowedActionsConfigArgsDict']] allowed_actions_config: Sets the actions that are allowed in an organization. Only available when `allowed_actions` = `selected`. See Allowed Actions Config below for details.
        :param pulumi.Input[_builtins.str] enabled_repositories: The policy that controls the repositories in the organization that are allowed to run GitHub Actions. Can be one of: `all`, `none`, or `selected`.
        :param pulumi.Input[Union['ActionsOrganizationPermissionsEnabledRepositoriesConfigArgs', 'ActionsOrganizationPermissionsEnabledRepositoriesConfigArgsDict']] enabled_repositories_config: Sets the list of selected repositories that are enabled for GitHub Actions in an organization. Only available when `enabled_repositories` = `selected`. See Enabled Repositories Config below for details.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ActionsOrganizationPermissionsState.__new__(_ActionsOrganizationPermissionsState)

        __props__.__dict__["allowed_actions"] = allowed_actions
        __props__.__dict__["allowed_actions_config"] = allowed_actions_config
        __props__.__dict__["enabled_repositories"] = enabled_repositories
        __props__.__dict__["enabled_repositories_config"] = enabled_repositories_config
        return ActionsOrganizationPermissions(resource_name, opts=opts, __props__=__props__)

    @_builtins.property
    @pulumi.getter(name="allowedActions")
    def allowed_actions(self) -> pulumi.Output[Optional[_builtins.str]]:
        """
        The permissions policy that controls the actions that are allowed to run. Can be one of: `all`, `local_only`, or `selected`.
        """
        return pulumi.get(self, "allowed_actions")

    @_builtins.property
    @pulumi.getter(name="allowedActionsConfig")
    def allowed_actions_config(self) -> pulumi.Output[Optional['outputs.ActionsOrganizationPermissionsAllowedActionsConfig']]:
        """
        Sets the actions that are allowed in an organization. Only available when `allowed_actions` = `selected`. See Allowed Actions Config below for details.
        """
        return pulumi.get(self, "allowed_actions_config")

    @_builtins.property
    @pulumi.getter(name="enabledRepositories")
    def enabled_repositories(self) -> pulumi.Output[_builtins.str]:
        """
        The policy that controls the repositories in the organization that are allowed to run GitHub Actions. Can be one of: `all`, `none`, or `selected`.
        """
        return pulumi.get(self, "enabled_repositories")

    @_builtins.property
    @pulumi.getter(name="enabledRepositoriesConfig")
    def enabled_repositories_config(self) -> pulumi.Output[Optional['outputs.ActionsOrganizationPermissionsEnabledRepositoriesConfig']]:
        """
        Sets the list of selected repositories that are enabled for GitHub Actions in an organization. Only available when `enabled_repositories` = `selected`. See Enabled Repositories Config below for details.
        """
        return pulumi.get(self, "enabled_repositories_config")

