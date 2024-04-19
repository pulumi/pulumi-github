# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['AppInstallationRepositoryArgs', 'AppInstallationRepository']

@pulumi.input_type
class AppInstallationRepositoryArgs:
    def __init__(__self__, *,
                 installation_id: pulumi.Input[str],
                 repository: pulumi.Input[str]):
        """
        The set of arguments for constructing a AppInstallationRepository resource.
        :param pulumi.Input[str] installation_id: The GitHub app installation id.
        :param pulumi.Input[str] repository: The repository to install the app on.
        """
        pulumi.set(__self__, "installation_id", installation_id)
        pulumi.set(__self__, "repository", repository)

    @property
    @pulumi.getter(name="installationId")
    def installation_id(self) -> pulumi.Input[str]:
        """
        The GitHub app installation id.
        """
        return pulumi.get(self, "installation_id")

    @installation_id.setter
    def installation_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "installation_id", value)

    @property
    @pulumi.getter
    def repository(self) -> pulumi.Input[str]:
        """
        The repository to install the app on.
        """
        return pulumi.get(self, "repository")

    @repository.setter
    def repository(self, value: pulumi.Input[str]):
        pulumi.set(self, "repository", value)


@pulumi.input_type
class _AppInstallationRepositoryState:
    def __init__(__self__, *,
                 installation_id: Optional[pulumi.Input[str]] = None,
                 repo_id: Optional[pulumi.Input[int]] = None,
                 repository: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering AppInstallationRepository resources.
        :param pulumi.Input[str] installation_id: The GitHub app installation id.
        :param pulumi.Input[str] repository: The repository to install the app on.
        """
        if installation_id is not None:
            pulumi.set(__self__, "installation_id", installation_id)
        if repo_id is not None:
            pulumi.set(__self__, "repo_id", repo_id)
        if repository is not None:
            pulumi.set(__self__, "repository", repository)

    @property
    @pulumi.getter(name="installationId")
    def installation_id(self) -> Optional[pulumi.Input[str]]:
        """
        The GitHub app installation id.
        """
        return pulumi.get(self, "installation_id")

    @installation_id.setter
    def installation_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "installation_id", value)

    @property
    @pulumi.getter(name="repoId")
    def repo_id(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "repo_id")

    @repo_id.setter
    def repo_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "repo_id", value)

    @property
    @pulumi.getter
    def repository(self) -> Optional[pulumi.Input[str]]:
        """
        The repository to install the app on.
        """
        return pulumi.get(self, "repository")

    @repository.setter
    def repository(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "repository", value)


class AppInstallationRepository(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 installation_id: Optional[pulumi.Input[str]] = None,
                 repository: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        > **Note**: This resource is not compatible with the GitHub App Installation authentication method.

        This resource manages relationships between app installations and repositories
        in your GitHub organization.

        Creating this resource installs a particular app on a particular repository.

        The app installation and the repository must both belong to the same
        organization on GitHub. Note: you can review your organization's installations
        by the following the instructions at this
        [link](https://docs.github.com/en/github/setting-up-and-managing-organizations-and-teams/reviewing-your-organizations-installed-integrations).

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_github as github

        # Create a repository.
        some_repo = github.Repository("some_repo", name="some-repo")
        some_app_repo = github.AppInstallationRepository("some_app_repo",
            installation_id="1234567",
            repository=some_repo.name)
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        GitHub App Installation Repository can be imported
        using an ID made up of `installation_id:repository`, e.g.

        ```sh
        $ pulumi import github:index/appInstallationRepository:AppInstallationRepository terraform_repo 1234567:terraform
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] installation_id: The GitHub app installation id.
        :param pulumi.Input[str] repository: The repository to install the app on.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AppInstallationRepositoryArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        > **Note**: This resource is not compatible with the GitHub App Installation authentication method.

        This resource manages relationships between app installations and repositories
        in your GitHub organization.

        Creating this resource installs a particular app on a particular repository.

        The app installation and the repository must both belong to the same
        organization on GitHub. Note: you can review your organization's installations
        by the following the instructions at this
        [link](https://docs.github.com/en/github/setting-up-and-managing-organizations-and-teams/reviewing-your-organizations-installed-integrations).

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_github as github

        # Create a repository.
        some_repo = github.Repository("some_repo", name="some-repo")
        some_app_repo = github.AppInstallationRepository("some_app_repo",
            installation_id="1234567",
            repository=some_repo.name)
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        GitHub App Installation Repository can be imported
        using an ID made up of `installation_id:repository`, e.g.

        ```sh
        $ pulumi import github:index/appInstallationRepository:AppInstallationRepository terraform_repo 1234567:terraform
        ```

        :param str resource_name: The name of the resource.
        :param AppInstallationRepositoryArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AppInstallationRepositoryArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 installation_id: Optional[pulumi.Input[str]] = None,
                 repository: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AppInstallationRepositoryArgs.__new__(AppInstallationRepositoryArgs)

            if installation_id is None and not opts.urn:
                raise TypeError("Missing required property 'installation_id'")
            __props__.__dict__["installation_id"] = installation_id
            if repository is None and not opts.urn:
                raise TypeError("Missing required property 'repository'")
            __props__.__dict__["repository"] = repository
            __props__.__dict__["repo_id"] = None
        super(AppInstallationRepository, __self__).__init__(
            'github:index/appInstallationRepository:AppInstallationRepository',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            installation_id: Optional[pulumi.Input[str]] = None,
            repo_id: Optional[pulumi.Input[int]] = None,
            repository: Optional[pulumi.Input[str]] = None) -> 'AppInstallationRepository':
        """
        Get an existing AppInstallationRepository resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] installation_id: The GitHub app installation id.
        :param pulumi.Input[str] repository: The repository to install the app on.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AppInstallationRepositoryState.__new__(_AppInstallationRepositoryState)

        __props__.__dict__["installation_id"] = installation_id
        __props__.__dict__["repo_id"] = repo_id
        __props__.__dict__["repository"] = repository
        return AppInstallationRepository(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="installationId")
    def installation_id(self) -> pulumi.Output[str]:
        """
        The GitHub app installation id.
        """
        return pulumi.get(self, "installation_id")

    @property
    @pulumi.getter(name="repoId")
    def repo_id(self) -> pulumi.Output[int]:
        return pulumi.get(self, "repo_id")

    @property
    @pulumi.getter
    def repository(self) -> pulumi.Output[str]:
        """
        The repository to install the app on.
        """
        return pulumi.get(self, "repository")

