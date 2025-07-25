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

__all__ = ['ActionsVariableArgs', 'ActionsVariable']

@pulumi.input_type
class ActionsVariableArgs:
    def __init__(__self__, *,
                 repository: pulumi.Input[_builtins.str],
                 value: pulumi.Input[_builtins.str],
                 variable_name: pulumi.Input[_builtins.str]):
        """
        The set of arguments for constructing a ActionsVariable resource.
        :param pulumi.Input[_builtins.str] repository: Name of the repository
        :param pulumi.Input[_builtins.str] value: Value of the variable
        :param pulumi.Input[_builtins.str] variable_name: Name of the variable
        """
        pulumi.set(__self__, "repository", repository)
        pulumi.set(__self__, "value", value)
        pulumi.set(__self__, "variable_name", variable_name)

    @_builtins.property
    @pulumi.getter
    def repository(self) -> pulumi.Input[_builtins.str]:
        """
        Name of the repository
        """
        return pulumi.get(self, "repository")

    @repository.setter
    def repository(self, value: pulumi.Input[_builtins.str]):
        pulumi.set(self, "repository", value)

    @_builtins.property
    @pulumi.getter
    def value(self) -> pulumi.Input[_builtins.str]:
        """
        Value of the variable
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[_builtins.str]):
        pulumi.set(self, "value", value)

    @_builtins.property
    @pulumi.getter(name="variableName")
    def variable_name(self) -> pulumi.Input[_builtins.str]:
        """
        Name of the variable
        """
        return pulumi.get(self, "variable_name")

    @variable_name.setter
    def variable_name(self, value: pulumi.Input[_builtins.str]):
        pulumi.set(self, "variable_name", value)


@pulumi.input_type
class _ActionsVariableState:
    def __init__(__self__, *,
                 created_at: Optional[pulumi.Input[_builtins.str]] = None,
                 repository: Optional[pulumi.Input[_builtins.str]] = None,
                 updated_at: Optional[pulumi.Input[_builtins.str]] = None,
                 value: Optional[pulumi.Input[_builtins.str]] = None,
                 variable_name: Optional[pulumi.Input[_builtins.str]] = None):
        """
        Input properties used for looking up and filtering ActionsVariable resources.
        :param pulumi.Input[_builtins.str] created_at: Date of actions_variable creation.
        :param pulumi.Input[_builtins.str] repository: Name of the repository
        :param pulumi.Input[_builtins.str] updated_at: Date of actions_variable update.
        :param pulumi.Input[_builtins.str] value: Value of the variable
        :param pulumi.Input[_builtins.str] variable_name: Name of the variable
        """
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if repository is not None:
            pulumi.set(__self__, "repository", repository)
        if updated_at is not None:
            pulumi.set(__self__, "updated_at", updated_at)
        if value is not None:
            pulumi.set(__self__, "value", value)
        if variable_name is not None:
            pulumi.set(__self__, "variable_name", variable_name)

    @_builtins.property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Date of actions_variable creation.
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "created_at", value)

    @_builtins.property
    @pulumi.getter
    def repository(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Name of the repository
        """
        return pulumi.get(self, "repository")

    @repository.setter
    def repository(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "repository", value)

    @_builtins.property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Date of actions_variable update.
        """
        return pulumi.get(self, "updated_at")

    @updated_at.setter
    def updated_at(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "updated_at", value)

    @_builtins.property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Value of the variable
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "value", value)

    @_builtins.property
    @pulumi.getter(name="variableName")
    def variable_name(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Name of the variable
        """
        return pulumi.get(self, "variable_name")

    @variable_name.setter
    def variable_name(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "variable_name", value)


@pulumi.type_token("github:index/actionsVariable:ActionsVariable")
class ActionsVariable(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 repository: Optional[pulumi.Input[_builtins.str]] = None,
                 value: Optional[pulumi.Input[_builtins.str]] = None,
                 variable_name: Optional[pulumi.Input[_builtins.str]] = None,
                 __props__=None):
        """
        This resource allows you to create and manage GitHub Actions variables within your GitHub repositories.
        You must have write access to a repository to use this resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_github as github

        example_variable = github.ActionsVariable("example_variable",
            repository="example_repository",
            variable_name="example_variable_name",
            value="example_variable_value")
        ```

        ## Import

        GitHub Actions variables can be imported using an ID made up of `repository:variable_name`, e.g.

        ```sh
        $ pulumi import github:index/actionsVariable:ActionsVariable myvariable myrepo:myvariable
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[_builtins.str] repository: Name of the repository
        :param pulumi.Input[_builtins.str] value: Value of the variable
        :param pulumi.Input[_builtins.str] variable_name: Name of the variable
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

        example_variable = github.ActionsVariable("example_variable",
            repository="example_repository",
            variable_name="example_variable_name",
            value="example_variable_value")
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
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 repository: Optional[pulumi.Input[_builtins.str]] = None,
                 value: Optional[pulumi.Input[_builtins.str]] = None,
                 variable_name: Optional[pulumi.Input[_builtins.str]] = None,
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
            created_at: Optional[pulumi.Input[_builtins.str]] = None,
            repository: Optional[pulumi.Input[_builtins.str]] = None,
            updated_at: Optional[pulumi.Input[_builtins.str]] = None,
            value: Optional[pulumi.Input[_builtins.str]] = None,
            variable_name: Optional[pulumi.Input[_builtins.str]] = None) -> 'ActionsVariable':
        """
        Get an existing ActionsVariable resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[_builtins.str] created_at: Date of actions_variable creation.
        :param pulumi.Input[_builtins.str] repository: Name of the repository
        :param pulumi.Input[_builtins.str] updated_at: Date of actions_variable update.
        :param pulumi.Input[_builtins.str] value: Value of the variable
        :param pulumi.Input[_builtins.str] variable_name: Name of the variable
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ActionsVariableState.__new__(_ActionsVariableState)

        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["repository"] = repository
        __props__.__dict__["updated_at"] = updated_at
        __props__.__dict__["value"] = value
        __props__.__dict__["variable_name"] = variable_name
        return ActionsVariable(resource_name, opts=opts, __props__=__props__)

    @_builtins.property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[_builtins.str]:
        """
        Date of actions_variable creation.
        """
        return pulumi.get(self, "created_at")

    @_builtins.property
    @pulumi.getter
    def repository(self) -> pulumi.Output[_builtins.str]:
        """
        Name of the repository
        """
        return pulumi.get(self, "repository")

    @_builtins.property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> pulumi.Output[_builtins.str]:
        """
        Date of actions_variable update.
        """
        return pulumi.get(self, "updated_at")

    @_builtins.property
    @pulumi.getter
    def value(self) -> pulumi.Output[_builtins.str]:
        """
        Value of the variable
        """
        return pulumi.get(self, "value")

    @_builtins.property
    @pulumi.getter(name="variableName")
    def variable_name(self) -> pulumi.Output[_builtins.str]:
        """
        Name of the variable
        """
        return pulumi.get(self, "variable_name")

