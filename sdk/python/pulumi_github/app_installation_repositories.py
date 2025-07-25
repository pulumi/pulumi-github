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

__all__ = ['AppInstallationRepositoriesArgs', 'AppInstallationRepositories']

@pulumi.input_type
class AppInstallationRepositoriesArgs:
    def __init__(__self__, *,
                 installation_id: pulumi.Input[_builtins.str],
                 selected_repositories: pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]):
        """
        The set of arguments for constructing a AppInstallationRepositories resource.
        :param pulumi.Input[_builtins.str] installation_id: The GitHub app installation id.
        :param pulumi.Input[Sequence[pulumi.Input[_builtins.str]]] selected_repositories: A list of repository names to install the app on.
               
               > **Note**: Due to how GitHub implements app installations, apps cannot be installed with no repositories selected. Therefore deleting this resource will leave one repository with the app installed. Manually uninstall the app or set the installation to all repositories via the GUI as after deleting this resource.
        """
        pulumi.set(__self__, "installation_id", installation_id)
        pulumi.set(__self__, "selected_repositories", selected_repositories)

    @_builtins.property
    @pulumi.getter(name="installationId")
    def installation_id(self) -> pulumi.Input[_builtins.str]:
        """
        The GitHub app installation id.
        """
        return pulumi.get(self, "installation_id")

    @installation_id.setter
    def installation_id(self, value: pulumi.Input[_builtins.str]):
        pulumi.set(self, "installation_id", value)

    @_builtins.property
    @pulumi.getter(name="selectedRepositories")
    def selected_repositories(self) -> pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]:
        """
        A list of repository names to install the app on.

        > **Note**: Due to how GitHub implements app installations, apps cannot be installed with no repositories selected. Therefore deleting this resource will leave one repository with the app installed. Manually uninstall the app or set the installation to all repositories via the GUI as after deleting this resource.
        """
        return pulumi.get(self, "selected_repositories")

    @selected_repositories.setter
    def selected_repositories(self, value: pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]):
        pulumi.set(self, "selected_repositories", value)


@pulumi.input_type
class _AppInstallationRepositoriesState:
    def __init__(__self__, *,
                 installation_id: Optional[pulumi.Input[_builtins.str]] = None,
                 selected_repositories: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None):
        """
        Input properties used for looking up and filtering AppInstallationRepositories resources.
        :param pulumi.Input[_builtins.str] installation_id: The GitHub app installation id.
        :param pulumi.Input[Sequence[pulumi.Input[_builtins.str]]] selected_repositories: A list of repository names to install the app on.
               
               > **Note**: Due to how GitHub implements app installations, apps cannot be installed with no repositories selected. Therefore deleting this resource will leave one repository with the app installed. Manually uninstall the app or set the installation to all repositories via the GUI as after deleting this resource.
        """
        if installation_id is not None:
            pulumi.set(__self__, "installation_id", installation_id)
        if selected_repositories is not None:
            pulumi.set(__self__, "selected_repositories", selected_repositories)

    @_builtins.property
    @pulumi.getter(name="installationId")
    def installation_id(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The GitHub app installation id.
        """
        return pulumi.get(self, "installation_id")

    @installation_id.setter
    def installation_id(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "installation_id", value)

    @_builtins.property
    @pulumi.getter(name="selectedRepositories")
    def selected_repositories(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]]:
        """
        A list of repository names to install the app on.

        > **Note**: Due to how GitHub implements app installations, apps cannot be installed with no repositories selected. Therefore deleting this resource will leave one repository with the app installed. Manually uninstall the app or set the installation to all repositories via the GUI as after deleting this resource.
        """
        return pulumi.get(self, "selected_repositories")

    @selected_repositories.setter
    def selected_repositories(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]]):
        pulumi.set(self, "selected_repositories", value)


@pulumi.type_token("github:index/appInstallationRepositories:AppInstallationRepositories")
class AppInstallationRepositories(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 installation_id: Optional[pulumi.Input[_builtins.str]] = None,
                 selected_repositories: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None,
                 __props__=None):
        """
        > **Note**: This resource is not compatible with the GitHub App Installation authentication method.

        This resource manages relationships between app installations and repositories
        in your GitHub organization.

        Creating this resource installs a particular app on multiple repositories.

        The app installation and the repositories must all belong to the same
        organization on GitHub. Note: you can review your organization's installations
        by the following the instructions at this
        [link](https://docs.github.com/en/github/setting-up-and-managing-organizations-and-teams/reviewing-your-organizations-installed-integrations).

        ## Import

        GitHub App Installation Repositories can be imported
        using an ID made up of `installation_id`, e.g.

        ```sh
        $ pulumi import github:index/appInstallationRepositories:AppInstallationRepositories some_app_repos 1234567
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[_builtins.str] installation_id: The GitHub app installation id.
        :param pulumi.Input[Sequence[pulumi.Input[_builtins.str]]] selected_repositories: A list of repository names to install the app on.
               
               > **Note**: Due to how GitHub implements app installations, apps cannot be installed with no repositories selected. Therefore deleting this resource will leave one repository with the app installed. Manually uninstall the app or set the installation to all repositories via the GUI as after deleting this resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AppInstallationRepositoriesArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        > **Note**: This resource is not compatible with the GitHub App Installation authentication method.

        This resource manages relationships between app installations and repositories
        in your GitHub organization.

        Creating this resource installs a particular app on multiple repositories.

        The app installation and the repositories must all belong to the same
        organization on GitHub. Note: you can review your organization's installations
        by the following the instructions at this
        [link](https://docs.github.com/en/github/setting-up-and-managing-organizations-and-teams/reviewing-your-organizations-installed-integrations).

        ## Import

        GitHub App Installation Repositories can be imported
        using an ID made up of `installation_id`, e.g.

        ```sh
        $ pulumi import github:index/appInstallationRepositories:AppInstallationRepositories some_app_repos 1234567
        ```

        :param str resource_name: The name of the resource.
        :param AppInstallationRepositoriesArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AppInstallationRepositoriesArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 installation_id: Optional[pulumi.Input[_builtins.str]] = None,
                 selected_repositories: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AppInstallationRepositoriesArgs.__new__(AppInstallationRepositoriesArgs)

            if installation_id is None and not opts.urn:
                raise TypeError("Missing required property 'installation_id'")
            __props__.__dict__["installation_id"] = installation_id
            if selected_repositories is None and not opts.urn:
                raise TypeError("Missing required property 'selected_repositories'")
            __props__.__dict__["selected_repositories"] = selected_repositories
        super(AppInstallationRepositories, __self__).__init__(
            'github:index/appInstallationRepositories:AppInstallationRepositories',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            installation_id: Optional[pulumi.Input[_builtins.str]] = None,
            selected_repositories: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None) -> 'AppInstallationRepositories':
        """
        Get an existing AppInstallationRepositories resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[_builtins.str] installation_id: The GitHub app installation id.
        :param pulumi.Input[Sequence[pulumi.Input[_builtins.str]]] selected_repositories: A list of repository names to install the app on.
               
               > **Note**: Due to how GitHub implements app installations, apps cannot be installed with no repositories selected. Therefore deleting this resource will leave one repository with the app installed. Manually uninstall the app or set the installation to all repositories via the GUI as after deleting this resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AppInstallationRepositoriesState.__new__(_AppInstallationRepositoriesState)

        __props__.__dict__["installation_id"] = installation_id
        __props__.__dict__["selected_repositories"] = selected_repositories
        return AppInstallationRepositories(resource_name, opts=opts, __props__=__props__)

    @_builtins.property
    @pulumi.getter(name="installationId")
    def installation_id(self) -> pulumi.Output[_builtins.str]:
        """
        The GitHub app installation id.
        """
        return pulumi.get(self, "installation_id")

    @_builtins.property
    @pulumi.getter(name="selectedRepositories")
    def selected_repositories(self) -> pulumi.Output[Sequence[_builtins.str]]:
        """
        A list of repository names to install the app on.

        > **Note**: Due to how GitHub implements app installations, apps cannot be installed with no repositories selected. Therefore deleting this resource will leave one repository with the app installed. Manually uninstall the app or set the installation to all repositories via the GUI as after deleting this resource.
        """
        return pulumi.get(self, "selected_repositories")

