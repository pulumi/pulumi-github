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

__all__ = ['RepositoryDeploymentBranchPolicyArgs', 'RepositoryDeploymentBranchPolicy']

@pulumi.input_type
class RepositoryDeploymentBranchPolicyArgs:
    def __init__(__self__, *,
                 environment_name: pulumi.Input[_builtins.str],
                 repository: pulumi.Input[_builtins.str],
                 name: Optional[pulumi.Input[_builtins.str]] = None):
        """
        The set of arguments for constructing a RepositoryDeploymentBranchPolicy resource.
        :param pulumi.Input[_builtins.str] environment_name: The name of the environment. This environment must have `deployment_branch_policy.custom_branch_policies` set to true or a 404 error will be thrown.
        :param pulumi.Input[_builtins.str] repository: The repository to create the policy in.
        :param pulumi.Input[_builtins.str] name: The name pattern that branches must match in order to deploy to the environment.
        """
        pulumi.set(__self__, "environment_name", environment_name)
        pulumi.set(__self__, "repository", repository)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @_builtins.property
    @pulumi.getter(name="environmentName")
    def environment_name(self) -> pulumi.Input[_builtins.str]:
        """
        The name of the environment. This environment must have `deployment_branch_policy.custom_branch_policies` set to true or a 404 error will be thrown.
        """
        return pulumi.get(self, "environment_name")

    @environment_name.setter
    def environment_name(self, value: pulumi.Input[_builtins.str]):
        pulumi.set(self, "environment_name", value)

    @_builtins.property
    @pulumi.getter
    def repository(self) -> pulumi.Input[_builtins.str]:
        """
        The repository to create the policy in.
        """
        return pulumi.get(self, "repository")

    @repository.setter
    def repository(self, value: pulumi.Input[_builtins.str]):
        pulumi.set(self, "repository", value)

    @_builtins.property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The name pattern that branches must match in order to deploy to the environment.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _RepositoryDeploymentBranchPolicyState:
    def __init__(__self__, *,
                 environment_name: Optional[pulumi.Input[_builtins.str]] = None,
                 etag: Optional[pulumi.Input[_builtins.str]] = None,
                 name: Optional[pulumi.Input[_builtins.str]] = None,
                 repository: Optional[pulumi.Input[_builtins.str]] = None):
        """
        Input properties used for looking up and filtering RepositoryDeploymentBranchPolicy resources.
        :param pulumi.Input[_builtins.str] environment_name: The name of the environment. This environment must have `deployment_branch_policy.custom_branch_policies` set to true or a 404 error will be thrown.
        :param pulumi.Input[_builtins.str] etag: An etag representing the Branch object.
        :param pulumi.Input[_builtins.str] name: The name pattern that branches must match in order to deploy to the environment.
        :param pulumi.Input[_builtins.str] repository: The repository to create the policy in.
        """
        if environment_name is not None:
            pulumi.set(__self__, "environment_name", environment_name)
        if etag is not None:
            pulumi.set(__self__, "etag", etag)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if repository is not None:
            pulumi.set(__self__, "repository", repository)

    @_builtins.property
    @pulumi.getter(name="environmentName")
    def environment_name(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The name of the environment. This environment must have `deployment_branch_policy.custom_branch_policies` set to true or a 404 error will be thrown.
        """
        return pulumi.get(self, "environment_name")

    @environment_name.setter
    def environment_name(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "environment_name", value)

    @_builtins.property
    @pulumi.getter
    def etag(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        An etag representing the Branch object.
        """
        return pulumi.get(self, "etag")

    @etag.setter
    def etag(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "etag", value)

    @_builtins.property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The name pattern that branches must match in order to deploy to the environment.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "name", value)

    @_builtins.property
    @pulumi.getter
    def repository(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The repository to create the policy in.
        """
        return pulumi.get(self, "repository")

    @repository.setter
    def repository(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "repository", value)


@pulumi.type_token("github:index/repositoryDeploymentBranchPolicy:RepositoryDeploymentBranchPolicy")
class RepositoryDeploymentBranchPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 environment_name: Optional[pulumi.Input[_builtins.str]] = None,
                 name: Optional[pulumi.Input[_builtins.str]] = None,
                 repository: Optional[pulumi.Input[_builtins.str]] = None,
                 __props__=None):
        """
        This resource allows you to create and manage deployment branch policies.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_github as github

        env = github.RepositoryEnvironment("env",
            repository="my_repo",
            environment="my_env",
            deployment_branch_policy={
                "protected_branches": False,
                "custom_branch_policies": True,
            })
        foo = github.RepositoryDeploymentBranchPolicy("foo",
            repository="my_repo",
            environment_name="my_env",
            name="foo",
            opts = pulumi.ResourceOptions(depends_on=[env]))
        ```

        ## Import

        ```sh
        $ pulumi import github:index/repositoryDeploymentBranchPolicy:RepositoryDeploymentBranchPolicy foo repo:env:id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[_builtins.str] environment_name: The name of the environment. This environment must have `deployment_branch_policy.custom_branch_policies` set to true or a 404 error will be thrown.
        :param pulumi.Input[_builtins.str] name: The name pattern that branches must match in order to deploy to the environment.
        :param pulumi.Input[_builtins.str] repository: The repository to create the policy in.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RepositoryDeploymentBranchPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This resource allows you to create and manage deployment branch policies.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_github as github

        env = github.RepositoryEnvironment("env",
            repository="my_repo",
            environment="my_env",
            deployment_branch_policy={
                "protected_branches": False,
                "custom_branch_policies": True,
            })
        foo = github.RepositoryDeploymentBranchPolicy("foo",
            repository="my_repo",
            environment_name="my_env",
            name="foo",
            opts = pulumi.ResourceOptions(depends_on=[env]))
        ```

        ## Import

        ```sh
        $ pulumi import github:index/repositoryDeploymentBranchPolicy:RepositoryDeploymentBranchPolicy foo repo:env:id
        ```

        :param str resource_name: The name of the resource.
        :param RepositoryDeploymentBranchPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RepositoryDeploymentBranchPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 environment_name: Optional[pulumi.Input[_builtins.str]] = None,
                 name: Optional[pulumi.Input[_builtins.str]] = None,
                 repository: Optional[pulumi.Input[_builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RepositoryDeploymentBranchPolicyArgs.__new__(RepositoryDeploymentBranchPolicyArgs)

            if environment_name is None and not opts.urn:
                raise TypeError("Missing required property 'environment_name'")
            __props__.__dict__["environment_name"] = environment_name
            __props__.__dict__["name"] = name
            if repository is None and not opts.urn:
                raise TypeError("Missing required property 'repository'")
            __props__.__dict__["repository"] = repository
            __props__.__dict__["etag"] = None
        super(RepositoryDeploymentBranchPolicy, __self__).__init__(
            'github:index/repositoryDeploymentBranchPolicy:RepositoryDeploymentBranchPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            environment_name: Optional[pulumi.Input[_builtins.str]] = None,
            etag: Optional[pulumi.Input[_builtins.str]] = None,
            name: Optional[pulumi.Input[_builtins.str]] = None,
            repository: Optional[pulumi.Input[_builtins.str]] = None) -> 'RepositoryDeploymentBranchPolicy':
        """
        Get an existing RepositoryDeploymentBranchPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[_builtins.str] environment_name: The name of the environment. This environment must have `deployment_branch_policy.custom_branch_policies` set to true or a 404 error will be thrown.
        :param pulumi.Input[_builtins.str] etag: An etag representing the Branch object.
        :param pulumi.Input[_builtins.str] name: The name pattern that branches must match in order to deploy to the environment.
        :param pulumi.Input[_builtins.str] repository: The repository to create the policy in.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RepositoryDeploymentBranchPolicyState.__new__(_RepositoryDeploymentBranchPolicyState)

        __props__.__dict__["environment_name"] = environment_name
        __props__.__dict__["etag"] = etag
        __props__.__dict__["name"] = name
        __props__.__dict__["repository"] = repository
        return RepositoryDeploymentBranchPolicy(resource_name, opts=opts, __props__=__props__)

    @_builtins.property
    @pulumi.getter(name="environmentName")
    def environment_name(self) -> pulumi.Output[_builtins.str]:
        """
        The name of the environment. This environment must have `deployment_branch_policy.custom_branch_policies` set to true or a 404 error will be thrown.
        """
        return pulumi.get(self, "environment_name")

    @_builtins.property
    @pulumi.getter
    def etag(self) -> pulumi.Output[_builtins.str]:
        """
        An etag representing the Branch object.
        """
        return pulumi.get(self, "etag")

    @_builtins.property
    @pulumi.getter
    def name(self) -> pulumi.Output[_builtins.str]:
        """
        The name pattern that branches must match in order to deploy to the environment.
        """
        return pulumi.get(self, "name")

    @_builtins.property
    @pulumi.getter
    def repository(self) -> pulumi.Output[_builtins.str]:
        """
        The repository to create the policy in.
        """
        return pulumi.get(self, "repository")

