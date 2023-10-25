# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ActionsOrganizationVariableArgs', 'ActionsOrganizationVariable']

@pulumi.input_type
class ActionsOrganizationVariableArgs:
    def __init__(__self__, *,
                 value: pulumi.Input[str],
                 variable_name: pulumi.Input[str],
                 visibility: pulumi.Input[str],
                 selected_repository_ids: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None):
        """
        The set of arguments for constructing a ActionsOrganizationVariable resource.
        :param pulumi.Input[str] value: Value of the variable
        :param pulumi.Input[str] variable_name: Name of the variable
        :param pulumi.Input[str] visibility: Configures the access that repositories have to the organization variable.
               Must be one of `all`, `private`, `selected`. `selected_repository_ids` is required if set to `selected`.
        :param pulumi.Input[Sequence[pulumi.Input[int]]] selected_repository_ids: An array of repository ids that can access the organization variable.
        """
        ActionsOrganizationVariableArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            value=value,
            variable_name=variable_name,
            visibility=visibility,
            selected_repository_ids=selected_repository_ids,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             value: Optional[pulumi.Input[str]] = None,
             variable_name: Optional[pulumi.Input[str]] = None,
             visibility: Optional[pulumi.Input[str]] = None,
             selected_repository_ids: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if value is None:
            raise TypeError("Missing 'value' argument")
        if variable_name is None and 'variableName' in kwargs:
            variable_name = kwargs['variableName']
        if variable_name is None:
            raise TypeError("Missing 'variable_name' argument")
        if visibility is None:
            raise TypeError("Missing 'visibility' argument")
        if selected_repository_ids is None and 'selectedRepositoryIds' in kwargs:
            selected_repository_ids = kwargs['selectedRepositoryIds']

        _setter("value", value)
        _setter("variable_name", variable_name)
        _setter("visibility", visibility)
        if selected_repository_ids is not None:
            _setter("selected_repository_ids", selected_repository_ids)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        """
        Value of the variable
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)

    @property
    @pulumi.getter(name="variableName")
    def variable_name(self) -> pulumi.Input[str]:
        """
        Name of the variable
        """
        return pulumi.get(self, "variable_name")

    @variable_name.setter
    def variable_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "variable_name", value)

    @property
    @pulumi.getter
    def visibility(self) -> pulumi.Input[str]:
        """
        Configures the access that repositories have to the organization variable.
        Must be one of `all`, `private`, `selected`. `selected_repository_ids` is required if set to `selected`.
        """
        return pulumi.get(self, "visibility")

    @visibility.setter
    def visibility(self, value: pulumi.Input[str]):
        pulumi.set(self, "visibility", value)

    @property
    @pulumi.getter(name="selectedRepositoryIds")
    def selected_repository_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[int]]]]:
        """
        An array of repository ids that can access the organization variable.
        """
        return pulumi.get(self, "selected_repository_ids")

    @selected_repository_ids.setter
    def selected_repository_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]]):
        pulumi.set(self, "selected_repository_ids", value)


@pulumi.input_type
class _ActionsOrganizationVariableState:
    def __init__(__self__, *,
                 created_at: Optional[pulumi.Input[str]] = None,
                 selected_repository_ids: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
                 updated_at: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None,
                 variable_name: Optional[pulumi.Input[str]] = None,
                 visibility: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ActionsOrganizationVariable resources.
        :param pulumi.Input[str] created_at: Date of actions_variable creation.
        :param pulumi.Input[Sequence[pulumi.Input[int]]] selected_repository_ids: An array of repository ids that can access the organization variable.
        :param pulumi.Input[str] updated_at: Date of actions_variable update.
        :param pulumi.Input[str] value: Value of the variable
        :param pulumi.Input[str] variable_name: Name of the variable
        :param pulumi.Input[str] visibility: Configures the access that repositories have to the organization variable.
               Must be one of `all`, `private`, `selected`. `selected_repository_ids` is required if set to `selected`.
        """
        _ActionsOrganizationVariableState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            created_at=created_at,
            selected_repository_ids=selected_repository_ids,
            updated_at=updated_at,
            value=value,
            variable_name=variable_name,
            visibility=visibility,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             created_at: Optional[pulumi.Input[str]] = None,
             selected_repository_ids: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
             updated_at: Optional[pulumi.Input[str]] = None,
             value: Optional[pulumi.Input[str]] = None,
             variable_name: Optional[pulumi.Input[str]] = None,
             visibility: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if created_at is None and 'createdAt' in kwargs:
            created_at = kwargs['createdAt']
        if selected_repository_ids is None and 'selectedRepositoryIds' in kwargs:
            selected_repository_ids = kwargs['selectedRepositoryIds']
        if updated_at is None and 'updatedAt' in kwargs:
            updated_at = kwargs['updatedAt']
        if variable_name is None and 'variableName' in kwargs:
            variable_name = kwargs['variableName']

        if created_at is not None:
            _setter("created_at", created_at)
        if selected_repository_ids is not None:
            _setter("selected_repository_ids", selected_repository_ids)
        if updated_at is not None:
            _setter("updated_at", updated_at)
        if value is not None:
            _setter("value", value)
        if variable_name is not None:
            _setter("variable_name", variable_name)
        if visibility is not None:
            _setter("visibility", visibility)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[str]]:
        """
        Date of actions_variable creation.
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter(name="selectedRepositoryIds")
    def selected_repository_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[int]]]]:
        """
        An array of repository ids that can access the organization variable.
        """
        return pulumi.get(self, "selected_repository_ids")

    @selected_repository_ids.setter
    def selected_repository_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]]):
        pulumi.set(self, "selected_repository_ids", value)

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> Optional[pulumi.Input[str]]:
        """
        Date of actions_variable update.
        """
        return pulumi.get(self, "updated_at")

    @updated_at.setter
    def updated_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "updated_at", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[str]]:
        """
        Value of the variable
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value", value)

    @property
    @pulumi.getter(name="variableName")
    def variable_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the variable
        """
        return pulumi.get(self, "variable_name")

    @variable_name.setter
    def variable_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "variable_name", value)

    @property
    @pulumi.getter
    def visibility(self) -> Optional[pulumi.Input[str]]:
        """
        Configures the access that repositories have to the organization variable.
        Must be one of `all`, `private`, `selected`. `selected_repository_ids` is required if set to `selected`.
        """
        return pulumi.get(self, "visibility")

    @visibility.setter
    def visibility(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "visibility", value)


class ActionsOrganizationVariable(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 selected_repository_ids: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
                 value: Optional[pulumi.Input[str]] = None,
                 variable_name: Optional[pulumi.Input[str]] = None,
                 visibility: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        This resource allows you to create and manage GitHub Actions variables within your GitHub organization.
        You must have write access to a repository to use this resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_github as github

        example_variable = github.ActionsOrganizationVariable("exampleVariable",
            value="example_variable_value",
            variable_name="example_variable_name",
            visibility="private")
        ```

        ```python
        import pulumi
        import pulumi_github as github

        repo = github.get_repository(full_name="my-org/repo")
        example_variable = github.ActionsOrganizationVariable("exampleVariable",
            variable_name="example_variable_name",
            visibility="selected",
            value="example_variable_value",
            selected_repository_ids=[repo.repo_id])
        ```

        ## Import

        This resource can be imported using an ID made up of the variable name:

        ```sh
         $ pulumi import github:index/actionsOrganizationVariable:ActionsOrganizationVariable test_variable test_variable_name
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[int]]] selected_repository_ids: An array of repository ids that can access the organization variable.
        :param pulumi.Input[str] value: Value of the variable
        :param pulumi.Input[str] variable_name: Name of the variable
        :param pulumi.Input[str] visibility: Configures the access that repositories have to the organization variable.
               Must be one of `all`, `private`, `selected`. `selected_repository_ids` is required if set to `selected`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ActionsOrganizationVariableArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This resource allows you to create and manage GitHub Actions variables within your GitHub organization.
        You must have write access to a repository to use this resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_github as github

        example_variable = github.ActionsOrganizationVariable("exampleVariable",
            value="example_variable_value",
            variable_name="example_variable_name",
            visibility="private")
        ```

        ```python
        import pulumi
        import pulumi_github as github

        repo = github.get_repository(full_name="my-org/repo")
        example_variable = github.ActionsOrganizationVariable("exampleVariable",
            variable_name="example_variable_name",
            visibility="selected",
            value="example_variable_value",
            selected_repository_ids=[repo.repo_id])
        ```

        ## Import

        This resource can be imported using an ID made up of the variable name:

        ```sh
         $ pulumi import github:index/actionsOrganizationVariable:ActionsOrganizationVariable test_variable test_variable_name
        ```

        :param str resource_name: The name of the resource.
        :param ActionsOrganizationVariableArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ActionsOrganizationVariableArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            ActionsOrganizationVariableArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 selected_repository_ids: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
                 value: Optional[pulumi.Input[str]] = None,
                 variable_name: Optional[pulumi.Input[str]] = None,
                 visibility: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ActionsOrganizationVariableArgs.__new__(ActionsOrganizationVariableArgs)

            __props__.__dict__["selected_repository_ids"] = selected_repository_ids
            if value is None and not opts.urn:
                raise TypeError("Missing required property 'value'")
            __props__.__dict__["value"] = value
            if variable_name is None and not opts.urn:
                raise TypeError("Missing required property 'variable_name'")
            __props__.__dict__["variable_name"] = variable_name
            if visibility is None and not opts.urn:
                raise TypeError("Missing required property 'visibility'")
            __props__.__dict__["visibility"] = visibility
            __props__.__dict__["created_at"] = None
            __props__.__dict__["updated_at"] = None
        super(ActionsOrganizationVariable, __self__).__init__(
            'github:index/actionsOrganizationVariable:ActionsOrganizationVariable',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            created_at: Optional[pulumi.Input[str]] = None,
            selected_repository_ids: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
            updated_at: Optional[pulumi.Input[str]] = None,
            value: Optional[pulumi.Input[str]] = None,
            variable_name: Optional[pulumi.Input[str]] = None,
            visibility: Optional[pulumi.Input[str]] = None) -> 'ActionsOrganizationVariable':
        """
        Get an existing ActionsOrganizationVariable resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] created_at: Date of actions_variable creation.
        :param pulumi.Input[Sequence[pulumi.Input[int]]] selected_repository_ids: An array of repository ids that can access the organization variable.
        :param pulumi.Input[str] updated_at: Date of actions_variable update.
        :param pulumi.Input[str] value: Value of the variable
        :param pulumi.Input[str] variable_name: Name of the variable
        :param pulumi.Input[str] visibility: Configures the access that repositories have to the organization variable.
               Must be one of `all`, `private`, `selected`. `selected_repository_ids` is required if set to `selected`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ActionsOrganizationVariableState.__new__(_ActionsOrganizationVariableState)

        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["selected_repository_ids"] = selected_repository_ids
        __props__.__dict__["updated_at"] = updated_at
        __props__.__dict__["value"] = value
        __props__.__dict__["variable_name"] = variable_name
        __props__.__dict__["visibility"] = visibility
        return ActionsOrganizationVariable(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        """
        Date of actions_variable creation.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="selectedRepositoryIds")
    def selected_repository_ids(self) -> pulumi.Output[Optional[Sequence[int]]]:
        """
        An array of repository ids that can access the organization variable.
        """
        return pulumi.get(self, "selected_repository_ids")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> pulumi.Output[str]:
        """
        Date of actions_variable update.
        """
        return pulumi.get(self, "updated_at")

    @property
    @pulumi.getter
    def value(self) -> pulumi.Output[str]:
        """
        Value of the variable
        """
        return pulumi.get(self, "value")

    @property
    @pulumi.getter(name="variableName")
    def variable_name(self) -> pulumi.Output[str]:
        """
        Name of the variable
        """
        return pulumi.get(self, "variable_name")

    @property
    @pulumi.getter
    def visibility(self) -> pulumi.Output[str]:
        """
        Configures the access that repositories have to the organization variable.
        Must be one of `all`, `private`, `selected`. `selected_repository_ids` is required if set to `selected`.
        """
        return pulumi.get(self, "visibility")

