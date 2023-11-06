# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ActionsOrganizationSecretRepositoriesArgs', 'ActionsOrganizationSecretRepositories']

@pulumi.input_type
class ActionsOrganizationSecretRepositoriesArgs:
    def __init__(__self__, *,
                 secret_name: pulumi.Input[str],
                 selected_repository_ids: pulumi.Input[Sequence[pulumi.Input[int]]]):
        """
        The set of arguments for constructing a ActionsOrganizationSecretRepositories resource.
        :param pulumi.Input[str] secret_name: Name of the existing secret
        :param pulumi.Input[Sequence[pulumi.Input[int]]] selected_repository_ids: An array of repository ids that can access the organization secret.
        """
        ActionsOrganizationSecretRepositoriesArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            secret_name=secret_name,
            selected_repository_ids=selected_repository_ids,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             secret_name: Optional[pulumi.Input[str]] = None,
             selected_repository_ids: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if secret_name is None and 'secretName' in kwargs:
            secret_name = kwargs['secretName']
        if secret_name is None:
            raise TypeError("Missing 'secret_name' argument")
        if selected_repository_ids is None and 'selectedRepositoryIds' in kwargs:
            selected_repository_ids = kwargs['selectedRepositoryIds']
        if selected_repository_ids is None:
            raise TypeError("Missing 'selected_repository_ids' argument")

        _setter("secret_name", secret_name)
        _setter("selected_repository_ids", selected_repository_ids)

    @property
    @pulumi.getter(name="secretName")
    def secret_name(self) -> pulumi.Input[str]:
        """
        Name of the existing secret
        """
        return pulumi.get(self, "secret_name")

    @secret_name.setter
    def secret_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "secret_name", value)

    @property
    @pulumi.getter(name="selectedRepositoryIds")
    def selected_repository_ids(self) -> pulumi.Input[Sequence[pulumi.Input[int]]]:
        """
        An array of repository ids that can access the organization secret.
        """
        return pulumi.get(self, "selected_repository_ids")

    @selected_repository_ids.setter
    def selected_repository_ids(self, value: pulumi.Input[Sequence[pulumi.Input[int]]]):
        pulumi.set(self, "selected_repository_ids", value)


@pulumi.input_type
class _ActionsOrganizationSecretRepositoriesState:
    def __init__(__self__, *,
                 secret_name: Optional[pulumi.Input[str]] = None,
                 selected_repository_ids: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None):
        """
        Input properties used for looking up and filtering ActionsOrganizationSecretRepositories resources.
        :param pulumi.Input[str] secret_name: Name of the existing secret
        :param pulumi.Input[Sequence[pulumi.Input[int]]] selected_repository_ids: An array of repository ids that can access the organization secret.
        """
        _ActionsOrganizationSecretRepositoriesState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            secret_name=secret_name,
            selected_repository_ids=selected_repository_ids,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             secret_name: Optional[pulumi.Input[str]] = None,
             selected_repository_ids: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if secret_name is None and 'secretName' in kwargs:
            secret_name = kwargs['secretName']
        if selected_repository_ids is None and 'selectedRepositoryIds' in kwargs:
            selected_repository_ids = kwargs['selectedRepositoryIds']

        if secret_name is not None:
            _setter("secret_name", secret_name)
        if selected_repository_ids is not None:
            _setter("selected_repository_ids", selected_repository_ids)

    @property
    @pulumi.getter(name="secretName")
    def secret_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the existing secret
        """
        return pulumi.get(self, "secret_name")

    @secret_name.setter
    def secret_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secret_name", value)

    @property
    @pulumi.getter(name="selectedRepositoryIds")
    def selected_repository_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[int]]]]:
        """
        An array of repository ids that can access the organization secret.
        """
        return pulumi.get(self, "selected_repository_ids")

    @selected_repository_ids.setter
    def selected_repository_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]]):
        pulumi.set(self, "selected_repository_ids", value)


class ActionsOrganizationSecretRepositories(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 secret_name: Optional[pulumi.Input[str]] = None,
                 selected_repository_ids: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
                 __props__=None):
        """
        This resource allows you to manage repository allow list for existing GitHub Actions secrets within your GitHub organization.
        You must have write access to an organization secret to use this resource.

        This resource is only applicable when `visibility` of the existing organization secret has been set to `selected`.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_github as github

        repo = github.get_repository(full_name="my-org/repo")
        org_secret_repos = github.ActionsOrganizationSecretRepositories("orgSecretRepos",
            secret_name="existing_secret_name",
            selected_repository_ids=[repo.repo_id])
        ```

        ## Import

        This resource can be imported using an ID made up of the secret name:

        ```sh
         $ pulumi import github:index/actionsOrganizationSecretRepositories:ActionsOrganizationSecretRepositories test_secret_repos test_secret_name
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] secret_name: Name of the existing secret
        :param pulumi.Input[Sequence[pulumi.Input[int]]] selected_repository_ids: An array of repository ids that can access the organization secret.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ActionsOrganizationSecretRepositoriesArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This resource allows you to manage repository allow list for existing GitHub Actions secrets within your GitHub organization.
        You must have write access to an organization secret to use this resource.

        This resource is only applicable when `visibility` of the existing organization secret has been set to `selected`.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_github as github

        repo = github.get_repository(full_name="my-org/repo")
        org_secret_repos = github.ActionsOrganizationSecretRepositories("orgSecretRepos",
            secret_name="existing_secret_name",
            selected_repository_ids=[repo.repo_id])
        ```

        ## Import

        This resource can be imported using an ID made up of the secret name:

        ```sh
         $ pulumi import github:index/actionsOrganizationSecretRepositories:ActionsOrganizationSecretRepositories test_secret_repos test_secret_name
        ```

        :param str resource_name: The name of the resource.
        :param ActionsOrganizationSecretRepositoriesArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ActionsOrganizationSecretRepositoriesArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            ActionsOrganizationSecretRepositoriesArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 secret_name: Optional[pulumi.Input[str]] = None,
                 selected_repository_ids: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ActionsOrganizationSecretRepositoriesArgs.__new__(ActionsOrganizationSecretRepositoriesArgs)

            if secret_name is None and not opts.urn:
                raise TypeError("Missing required property 'secret_name'")
            __props__.__dict__["secret_name"] = secret_name
            if selected_repository_ids is None and not opts.urn:
                raise TypeError("Missing required property 'selected_repository_ids'")
            __props__.__dict__["selected_repository_ids"] = selected_repository_ids
        super(ActionsOrganizationSecretRepositories, __self__).__init__(
            'github:index/actionsOrganizationSecretRepositories:ActionsOrganizationSecretRepositories',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            secret_name: Optional[pulumi.Input[str]] = None,
            selected_repository_ids: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None) -> 'ActionsOrganizationSecretRepositories':
        """
        Get an existing ActionsOrganizationSecretRepositories resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] secret_name: Name of the existing secret
        :param pulumi.Input[Sequence[pulumi.Input[int]]] selected_repository_ids: An array of repository ids that can access the organization secret.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ActionsOrganizationSecretRepositoriesState.__new__(_ActionsOrganizationSecretRepositoriesState)

        __props__.__dict__["secret_name"] = secret_name
        __props__.__dict__["selected_repository_ids"] = selected_repository_ids
        return ActionsOrganizationSecretRepositories(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="secretName")
    def secret_name(self) -> pulumi.Output[str]:
        """
        Name of the existing secret
        """
        return pulumi.get(self, "secret_name")

    @property
    @pulumi.getter(name="selectedRepositoryIds")
    def selected_repository_ids(self) -> pulumi.Output[Sequence[int]]:
        """
        An array of repository ids that can access the organization secret.
        """
        return pulumi.get(self, "selected_repository_ids")

