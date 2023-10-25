# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['BranchDefaultArgs', 'BranchDefault']

@pulumi.input_type
class BranchDefaultArgs:
    def __init__(__self__, *,
                 branch: pulumi.Input[str],
                 repository: pulumi.Input[str],
                 rename: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a BranchDefault resource.
        :param pulumi.Input[str] branch: The branch (e.g. `main`)
        :param pulumi.Input[str] repository: The GitHub repository
        :param pulumi.Input[bool] rename: Indicate if it should rename the branch rather than use an existing branch. Defaults to `false`.
        """
        BranchDefaultArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            branch=branch,
            repository=repository,
            rename=rename,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             branch: Optional[pulumi.Input[str]] = None,
             repository: Optional[pulumi.Input[str]] = None,
             rename: Optional[pulumi.Input[bool]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if branch is None:
            raise TypeError("Missing 'branch' argument")
        if repository is None:
            raise TypeError("Missing 'repository' argument")

        _setter("branch", branch)
        _setter("repository", repository)
        if rename is not None:
            _setter("rename", rename)

    @property
    @pulumi.getter
    def branch(self) -> pulumi.Input[str]:
        """
        The branch (e.g. `main`)
        """
        return pulumi.get(self, "branch")

    @branch.setter
    def branch(self, value: pulumi.Input[str]):
        pulumi.set(self, "branch", value)

    @property
    @pulumi.getter
    def repository(self) -> pulumi.Input[str]:
        """
        The GitHub repository
        """
        return pulumi.get(self, "repository")

    @repository.setter
    def repository(self, value: pulumi.Input[str]):
        pulumi.set(self, "repository", value)

    @property
    @pulumi.getter
    def rename(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicate if it should rename the branch rather than use an existing branch. Defaults to `false`.
        """
        return pulumi.get(self, "rename")

    @rename.setter
    def rename(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "rename", value)


@pulumi.input_type
class _BranchDefaultState:
    def __init__(__self__, *,
                 branch: Optional[pulumi.Input[str]] = None,
                 rename: Optional[pulumi.Input[bool]] = None,
                 repository: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering BranchDefault resources.
        :param pulumi.Input[str] branch: The branch (e.g. `main`)
        :param pulumi.Input[bool] rename: Indicate if it should rename the branch rather than use an existing branch. Defaults to `false`.
        :param pulumi.Input[str] repository: The GitHub repository
        """
        _BranchDefaultState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            branch=branch,
            rename=rename,
            repository=repository,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             branch: Optional[pulumi.Input[str]] = None,
             rename: Optional[pulumi.Input[bool]] = None,
             repository: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):

        if branch is not None:
            _setter("branch", branch)
        if rename is not None:
            _setter("rename", rename)
        if repository is not None:
            _setter("repository", repository)

    @property
    @pulumi.getter
    def branch(self) -> Optional[pulumi.Input[str]]:
        """
        The branch (e.g. `main`)
        """
        return pulumi.get(self, "branch")

    @branch.setter
    def branch(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "branch", value)

    @property
    @pulumi.getter
    def rename(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicate if it should rename the branch rather than use an existing branch. Defaults to `false`.
        """
        return pulumi.get(self, "rename")

    @rename.setter
    def rename(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "rename", value)

    @property
    @pulumi.getter
    def repository(self) -> Optional[pulumi.Input[str]]:
        """
        The GitHub repository
        """
        return pulumi.get(self, "repository")

    @repository.setter
    def repository(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "repository", value)


class BranchDefault(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 branch: Optional[pulumi.Input[str]] = None,
                 rename: Optional[pulumi.Input[bool]] = None,
                 repository: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a GitHub branch default resource.

        This resource allows you to set the default branch for a given repository.

        Note that use of this resource is incompatible with the `default_branch` option of the `Repository` resource.  Using both will result in plans always showing a diff.

        ## Example Usage

        Basic usage:

        ```python
        import pulumi
        import pulumi_github as github

        example = github.Repository("example",
            description="My awesome codebase",
            auto_init=True)
        development = github.Branch("development",
            repository=example.name,
            branch="development")
        default = github.BranchDefault("default",
            repository=example.name,
            branch=development.branch)
        ```

        Renaming to a branch that doesn't exist:

        ```python
        import pulumi
        import pulumi_github as github

        example = github.Repository("example",
            description="My awesome codebase",
            auto_init=True)
        default = github.BranchDefault("default",
            repository=example.name,
            branch="development",
            rename=True)
        ```

        ## Import

        GitHub Branch Defaults can be imported using an ID made up of `repository`, e.g.

        ```sh
         $ pulumi import github:index/branchDefault:BranchDefault branch_default my-repo
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] branch: The branch (e.g. `main`)
        :param pulumi.Input[bool] rename: Indicate if it should rename the branch rather than use an existing branch. Defaults to `false`.
        :param pulumi.Input[str] repository: The GitHub repository
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: BranchDefaultArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a GitHub branch default resource.

        This resource allows you to set the default branch for a given repository.

        Note that use of this resource is incompatible with the `default_branch` option of the `Repository` resource.  Using both will result in plans always showing a diff.

        ## Example Usage

        Basic usage:

        ```python
        import pulumi
        import pulumi_github as github

        example = github.Repository("example",
            description="My awesome codebase",
            auto_init=True)
        development = github.Branch("development",
            repository=example.name,
            branch="development")
        default = github.BranchDefault("default",
            repository=example.name,
            branch=development.branch)
        ```

        Renaming to a branch that doesn't exist:

        ```python
        import pulumi
        import pulumi_github as github

        example = github.Repository("example",
            description="My awesome codebase",
            auto_init=True)
        default = github.BranchDefault("default",
            repository=example.name,
            branch="development",
            rename=True)
        ```

        ## Import

        GitHub Branch Defaults can be imported using an ID made up of `repository`, e.g.

        ```sh
         $ pulumi import github:index/branchDefault:BranchDefault branch_default my-repo
        ```

        :param str resource_name: The name of the resource.
        :param BranchDefaultArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BranchDefaultArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            BranchDefaultArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 branch: Optional[pulumi.Input[str]] = None,
                 rename: Optional[pulumi.Input[bool]] = None,
                 repository: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = BranchDefaultArgs.__new__(BranchDefaultArgs)

            if branch is None and not opts.urn:
                raise TypeError("Missing required property 'branch'")
            __props__.__dict__["branch"] = branch
            __props__.__dict__["rename"] = rename
            if repository is None and not opts.urn:
                raise TypeError("Missing required property 'repository'")
            __props__.__dict__["repository"] = repository
        super(BranchDefault, __self__).__init__(
            'github:index/branchDefault:BranchDefault',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            branch: Optional[pulumi.Input[str]] = None,
            rename: Optional[pulumi.Input[bool]] = None,
            repository: Optional[pulumi.Input[str]] = None) -> 'BranchDefault':
        """
        Get an existing BranchDefault resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] branch: The branch (e.g. `main`)
        :param pulumi.Input[bool] rename: Indicate if it should rename the branch rather than use an existing branch. Defaults to `false`.
        :param pulumi.Input[str] repository: The GitHub repository
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _BranchDefaultState.__new__(_BranchDefaultState)

        __props__.__dict__["branch"] = branch
        __props__.__dict__["rename"] = rename
        __props__.__dict__["repository"] = repository
        return BranchDefault(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def branch(self) -> pulumi.Output[str]:
        """
        The branch (e.g. `main`)
        """
        return pulumi.get(self, "branch")

    @property
    @pulumi.getter
    def rename(self) -> pulumi.Output[Optional[bool]]:
        """
        Indicate if it should rename the branch rather than use an existing branch. Defaults to `false`.
        """
        return pulumi.get(self, "rename")

    @property
    @pulumi.getter
    def repository(self) -> pulumi.Output[str]:
        """
        The GitHub repository
        """
        return pulumi.get(self, "repository")

