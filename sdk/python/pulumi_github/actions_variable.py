# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ActionsVariableArgs', 'ActionsVariable']

@pulumi.input_type
class ActionsVariableArgs:
    def __init__(__self__, *,
                 repository: pulumi.Input[str],
                 value: pulumi.Input[str],
                 variable_name: pulumi.Input[str]):
        """
        The set of arguments for constructing a ActionsVariable resource.
        :param pulumi.Input[str] repository: Name of the repository
        :param pulumi.Input[str] value: Value of the variable
        :param pulumi.Input[str] variable_name: Name of the variable
        """
        ActionsVariableArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            repository=repository,
            value=value,
            variable_name=variable_name,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             repository: Optional[pulumi.Input[str]] = None,
             value: Optional[pulumi.Input[str]] = None,
             variable_name: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if repository is None:
            raise TypeError("Missing 'repository' argument")
        if value is None:
            raise TypeError("Missing 'value' argument")
        if variable_name is None and 'variableName' in kwargs:
            variable_name = kwargs['variableName']
        if variable_name is None:
            raise TypeError("Missing 'variable_name' argument")

        _setter("repository", repository)
        _setter("value", value)
        _setter("variable_name", variable_name)

    @property
    @pulumi.getter
    def repository(self) -> pulumi.Input[str]:
        """
        Name of the repository
        """
        return pulumi.get(self, "repository")

    @repository.setter
    def repository(self, value: pulumi.Input[str]):
        pulumi.set(self, "repository", value)

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


@pulumi.input_type
class _ActionsVariableState:
    def __init__(__self__, *,
                 created_at: Optional[pulumi.Input[str]] = None,
                 repository: Optional[pulumi.Input[str]] = None,
                 updated_at: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None,
                 variable_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ActionsVariable resources.
        :param pulumi.Input[str] created_at: Date of actions_variable creation.
        :param pulumi.Input[str] repository: Name of the repository
        :param pulumi.Input[str] updated_at: Date of actions_variable update.
        :param pulumi.Input[str] value: Value of the variable
        :param pulumi.Input[str] variable_name: Name of the variable
        """
        _ActionsVariableState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            created_at=created_at,
            repository=repository,
            updated_at=updated_at,
            value=value,
            variable_name=variable_name,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             created_at: Optional[pulumi.Input[str]] = None,
             repository: Optional[pulumi.Input[str]] = None,
             updated_at: Optional[pulumi.Input[str]] = None,
             value: Optional[pulumi.Input[str]] = None,
             variable_name: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if created_at is None and 'createdAt' in kwargs:
            created_at = kwargs['createdAt']
        if updated_at is None and 'updatedAt' in kwargs:
            updated_at = kwargs['updatedAt']
        if variable_name is None and 'variableName' in kwargs:
            variable_name = kwargs['variableName']

        if created_at is not None:
            _setter("created_at", created_at)
        if repository is not None:
            _setter("repository", repository)
        if updated_at is not None:
            _setter("updated_at", updated_at)
        if value is not None:
            _setter("value", value)
        if variable_name is not None:
            _setter("variable_name", variable_name)

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
    @pulumi.getter
    def repository(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the repository
        """
        return pulumi.get(self, "repository")

    @repository.setter
    def repository(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "repository", value)

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


class ActionsVariable(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 repository: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None,
                 variable_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        This resource allows you to create and manage GitHub Actions variables within your GitHub repositories.
        You must have write access to a repository to use this resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_github as github

        example_variable = github.ActionsVariable("exampleVariable",
            repository="example_repository",
            value="example_variable_value",
            variable_name="example_variable_name")
        ```

        ## Import

        GitHub Actions variables can be imported using an ID made up of `repository:variable_name`, e.g.

        ```sh
         $ pulumi import github:index/actionsVariable:ActionsVariable myvariable myrepo:myvariable
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] repository: Name of the repository
        :param pulumi.Input[str] value: Value of the variable
        :param pulumi.Input[str] variable_name: Name of the variable
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ActionsVariableArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This resource allows you to create and manage GitHub Actions variables within your GitHub repositories.
        You must have write access to a repository to use this resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_github as github

        example_variable = github.ActionsVariable("exampleVariable",
            repository="example_repository",
            value="example_variable_value",
            variable_name="example_variable_name")
        ```

        ## Import

        GitHub Actions variables can be imported using an ID made up of `repository:variable_name`, e.g.

        ```sh
         $ pulumi import github:index/actionsVariable:ActionsVariable myvariable myrepo:myvariable
        ```

        :param str resource_name: The name of the resource.
        :param ActionsVariableArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ActionsVariableArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            ActionsVariableArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 repository: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None,
                 variable_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ActionsVariableArgs.__new__(ActionsVariableArgs)

            if repository is None and not opts.urn:
                raise TypeError("Missing required property 'repository'")
            __props__.__dict__["repository"] = repository
            if value is None and not opts.urn:
                raise TypeError("Missing required property 'value'")
            __props__.__dict__["value"] = value
            if variable_name is None and not opts.urn:
                raise TypeError("Missing required property 'variable_name'")
            __props__.__dict__["variable_name"] = variable_name
            __props__.__dict__["created_at"] = None
            __props__.__dict__["updated_at"] = None
        super(ActionsVariable, __self__).__init__(
            'github:index/actionsVariable:ActionsVariable',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            created_at: Optional[pulumi.Input[str]] = None,
            repository: Optional[pulumi.Input[str]] = None,
            updated_at: Optional[pulumi.Input[str]] = None,
            value: Optional[pulumi.Input[str]] = None,
            variable_name: Optional[pulumi.Input[str]] = None) -> 'ActionsVariable':
        """
        Get an existing ActionsVariable resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] created_at: Date of actions_variable creation.
        :param pulumi.Input[str] repository: Name of the repository
        :param pulumi.Input[str] updated_at: Date of actions_variable update.
        :param pulumi.Input[str] value: Value of the variable
        :param pulumi.Input[str] variable_name: Name of the variable
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ActionsVariableState.__new__(_ActionsVariableState)

        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["repository"] = repository
        __props__.__dict__["updated_at"] = updated_at
        __props__.__dict__["value"] = value
        __props__.__dict__["variable_name"] = variable_name
        return ActionsVariable(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        """
        Date of actions_variable creation.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def repository(self) -> pulumi.Output[str]:
        """
        Name of the repository
        """
        return pulumi.get(self, "repository")

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

