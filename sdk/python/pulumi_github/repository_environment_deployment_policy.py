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

__all__ = ['RepositoryEnvironmentDeploymentPolicyArgs', 'RepositoryEnvironmentDeploymentPolicy']

@pulumi.input_type
class RepositoryEnvironmentDeploymentPolicyArgs:
    def __init__(__self__, *,
                 environment: pulumi.Input[_builtins.str],
                 repository: pulumi.Input[_builtins.str],
                 branch_pattern: Optional[pulumi.Input[_builtins.str]] = None,
                 tag_pattern: Optional[pulumi.Input[_builtins.str]] = None):
        """
        The set of arguments for constructing a RepositoryEnvironmentDeploymentPolicy resource.
        :param pulumi.Input[_builtins.str] environment: The name of the environment.
        :param pulumi.Input[_builtins.str] repository: The repository of the environment.
        :param pulumi.Input[_builtins.str] branch_pattern: The name pattern that branches must match in order to deploy to the environment. If not specified, `tag_pattern` must be specified.
        :param pulumi.Input[_builtins.str] tag_pattern: The name pattern that tags must match in order to deploy to the environment. If not specified, `branch_pattern` must be specified.
        """
        pulumi.set(__self__, "environment", environment)
        pulumi.set(__self__, "repository", repository)
        if branch_pattern is not None:
            pulumi.set(__self__, "branch_pattern", branch_pattern)
        if tag_pattern is not None:
            pulumi.set(__self__, "tag_pattern", tag_pattern)

    @_builtins.property
    @pulumi.getter
    def environment(self) -> pulumi.Input[_builtins.str]:
        """
        The name of the environment.
        """
        return pulumi.get(self, "environment")

    @environment.setter
    def environment(self, value: pulumi.Input[_builtins.str]):
        pulumi.set(self, "environment", value)

    @_builtins.property
    @pulumi.getter
    def repository(self) -> pulumi.Input[_builtins.str]:
        """
        The repository of the environment.
        """
        return pulumi.get(self, "repository")

    @repository.setter
    def repository(self, value: pulumi.Input[_builtins.str]):
        pulumi.set(self, "repository", value)

    @_builtins.property
    @pulumi.getter(name="branchPattern")
    def branch_pattern(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The name pattern that branches must match in order to deploy to the environment. If not specified, `tag_pattern` must be specified.
        """
        return pulumi.get(self, "branch_pattern")

    @branch_pattern.setter
    def branch_pattern(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "branch_pattern", value)

    @_builtins.property
    @pulumi.getter(name="tagPattern")
    def tag_pattern(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The name pattern that tags must match in order to deploy to the environment. If not specified, `branch_pattern` must be specified.
        """
        return pulumi.get(self, "tag_pattern")

    @tag_pattern.setter
    def tag_pattern(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "tag_pattern", value)


@pulumi.input_type
class _RepositoryEnvironmentDeploymentPolicyState:
    def __init__(__self__, *,
                 branch_pattern: Optional[pulumi.Input[_builtins.str]] = None,
                 environment: Optional[pulumi.Input[_builtins.str]] = None,
                 repository: Optional[pulumi.Input[_builtins.str]] = None,
                 tag_pattern: Optional[pulumi.Input[_builtins.str]] = None):
        """
        Input properties used for looking up and filtering RepositoryEnvironmentDeploymentPolicy resources.
        :param pulumi.Input[_builtins.str] branch_pattern: The name pattern that branches must match in order to deploy to the environment. If not specified, `tag_pattern` must be specified.
        :param pulumi.Input[_builtins.str] environment: The name of the environment.
        :param pulumi.Input[_builtins.str] repository: The repository of the environment.
        :param pulumi.Input[_builtins.str] tag_pattern: The name pattern that tags must match in order to deploy to the environment. If not specified, `branch_pattern` must be specified.
        """
        if branch_pattern is not None:
            pulumi.set(__self__, "branch_pattern", branch_pattern)
        if environment is not None:
            pulumi.set(__self__, "environment", environment)
        if repository is not None:
            pulumi.set(__self__, "repository", repository)
        if tag_pattern is not None:
            pulumi.set(__self__, "tag_pattern", tag_pattern)

    @_builtins.property
    @pulumi.getter(name="branchPattern")
    def branch_pattern(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The name pattern that branches must match in order to deploy to the environment. If not specified, `tag_pattern` must be specified.
        """
        return pulumi.get(self, "branch_pattern")

    @branch_pattern.setter
    def branch_pattern(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "branch_pattern", value)

    @_builtins.property
    @pulumi.getter
    def environment(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The name of the environment.
        """
        return pulumi.get(self, "environment")

    @environment.setter
    def environment(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "environment", value)

    @_builtins.property
    @pulumi.getter
    def repository(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The repository of the environment.
        """
        return pulumi.get(self, "repository")

    @repository.setter
    def repository(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "repository", value)

    @_builtins.property
    @pulumi.getter(name="tagPattern")
    def tag_pattern(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The name pattern that tags must match in order to deploy to the environment. If not specified, `branch_pattern` must be specified.
        """
        return pulumi.get(self, "tag_pattern")

    @tag_pattern.setter
    def tag_pattern(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "tag_pattern", value)


@pulumi.type_token("github:index/repositoryEnvironmentDeploymentPolicy:RepositoryEnvironmentDeploymentPolicy")
class RepositoryEnvironmentDeploymentPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 branch_pattern: Optional[pulumi.Input[_builtins.str]] = None,
                 environment: Optional[pulumi.Input[_builtins.str]] = None,
                 repository: Optional[pulumi.Input[_builtins.str]] = None,
                 tag_pattern: Optional[pulumi.Input[_builtins.str]] = None,
                 __props__=None):
        """
        This resource allows you to create and manage environment deployment branch policies for a GitHub repository.

        ## Example Usage

        Create a branch-based deployment policy:

        ```python
        import pulumi
        import pulumi_github as github

        current = github.get_user(username="")
        test = github.Repository("test", name="tf-acc-test-%s")
        test_repository_environment = github.RepositoryEnvironment("test",
            repository=test.name,
            environment="environment/test",
            wait_timer=10000,
            reviewers=[{
                "users": [current.id],
            }],
            deployment_branch_policy={
                "protected_branches": False,
                "custom_branch_policies": True,
            })
        test_repository_environment_deployment_policy = github.RepositoryEnvironmentDeploymentPolicy("test",
            repository=test.name,
            environment=test_repository_environment.environment,
            branch_pattern="releases/*")
        ```

        Create a tag-based deployment policy:

        ```python
        import pulumi
        import pulumi_github as github

        current = github.get_user(username="")
        test = github.Repository("test", name="tf-acc-test-%s")
        test_repository_environment = github.RepositoryEnvironment("test",
            repository=test.name,
            environment="environment/test",
            wait_timer=10000,
            reviewers=[{
                "users": [current.id],
            }],
            deployment_branch_policy={
                "protected_branches": False,
                "custom_branch_policies": True,
            })
        test_repository_environment_deployment_policy = github.RepositoryEnvironmentDeploymentPolicy("test",
            repository=test.name,
            environment=test_repository_environment.environment,
            tag_pattern="v*")
        ```

        ## Import

        GitHub Repository Environment Deployment Policy can be imported using an ID made up of `name` of the repository combined with the `environment` name of the environment with the `Id` of the deployment policy, separated by a `:` character, e.g.

        ```sh
        $ pulumi import github:index/repositoryEnvironmentDeploymentPolicy:RepositoryEnvironmentDeploymentPolicy daily terraform:daily:123456
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[_builtins.str] branch_pattern: The name pattern that branches must match in order to deploy to the environment. If not specified, `tag_pattern` must be specified.
        :param pulumi.Input[_builtins.str] environment: The name of the environment.
        :param pulumi.Input[_builtins.str] repository: The repository of the environment.
        :param pulumi.Input[_builtins.str] tag_pattern: The name pattern that tags must match in order to deploy to the environment. If not specified, `branch_pattern` must be specified.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RepositoryEnvironmentDeploymentPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This resource allows you to create and manage environment deployment branch policies for a GitHub repository.

        ## Example Usage

        Create a branch-based deployment policy:

        ```python
        import pulumi
        import pulumi_github as github

        current = github.get_user(username="")
        test = github.Repository("test", name="tf-acc-test-%s")
        test_repository_environment = github.RepositoryEnvironment("test",
            repository=test.name,
            environment="environment/test",
            wait_timer=10000,
            reviewers=[{
                "users": [current.id],
            }],
            deployment_branch_policy={
                "protected_branches": False,
                "custom_branch_policies": True,
            })
        test_repository_environment_deployment_policy = github.RepositoryEnvironmentDeploymentPolicy("test",
            repository=test.name,
            environment=test_repository_environment.environment,
            branch_pattern="releases/*")
        ```

        Create a tag-based deployment policy:

        ```python
        import pulumi
        import pulumi_github as github

        current = github.get_user(username="")
        test = github.Repository("test", name="tf-acc-test-%s")
        test_repository_environment = github.RepositoryEnvironment("test",
            repository=test.name,
            environment="environment/test",
            wait_timer=10000,
            reviewers=[{
                "users": [current.id],
            }],
            deployment_branch_policy={
                "protected_branches": False,
                "custom_branch_policies": True,
            })
        test_repository_environment_deployment_policy = github.RepositoryEnvironmentDeploymentPolicy("test",
            repository=test.name,
            environment=test_repository_environment.environment,
            tag_pattern="v*")
        ```

        ## Import

        GitHub Repository Environment Deployment Policy can be imported using an ID made up of `name` of the repository combined with the `environment` name of the environment with the `Id` of the deployment policy, separated by a `:` character, e.g.

        ```sh
        $ pulumi import github:index/repositoryEnvironmentDeploymentPolicy:RepositoryEnvironmentDeploymentPolicy daily terraform:daily:123456
        ```

        :param str resource_name: The name of the resource.
        :param RepositoryEnvironmentDeploymentPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RepositoryEnvironmentDeploymentPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 branch_pattern: Optional[pulumi.Input[_builtins.str]] = None,
                 environment: Optional[pulumi.Input[_builtins.str]] = None,
                 repository: Optional[pulumi.Input[_builtins.str]] = None,
                 tag_pattern: Optional[pulumi.Input[_builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RepositoryEnvironmentDeploymentPolicyArgs.__new__(RepositoryEnvironmentDeploymentPolicyArgs)

            __props__.__dict__["branch_pattern"] = branch_pattern
            if environment is None and not opts.urn:
                raise TypeError("Missing required property 'environment'")
            __props__.__dict__["environment"] = environment
            if repository is None and not opts.urn:
                raise TypeError("Missing required property 'repository'")
            __props__.__dict__["repository"] = repository
            __props__.__dict__["tag_pattern"] = tag_pattern
        super(RepositoryEnvironmentDeploymentPolicy, __self__).__init__(
            'github:index/repositoryEnvironmentDeploymentPolicy:RepositoryEnvironmentDeploymentPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            branch_pattern: Optional[pulumi.Input[_builtins.str]] = None,
            environment: Optional[pulumi.Input[_builtins.str]] = None,
            repository: Optional[pulumi.Input[_builtins.str]] = None,
            tag_pattern: Optional[pulumi.Input[_builtins.str]] = None) -> 'RepositoryEnvironmentDeploymentPolicy':
        """
        Get an existing RepositoryEnvironmentDeploymentPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[_builtins.str] branch_pattern: The name pattern that branches must match in order to deploy to the environment. If not specified, `tag_pattern` must be specified.
        :param pulumi.Input[_builtins.str] environment: The name of the environment.
        :param pulumi.Input[_builtins.str] repository: The repository of the environment.
        :param pulumi.Input[_builtins.str] tag_pattern: The name pattern that tags must match in order to deploy to the environment. If not specified, `branch_pattern` must be specified.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RepositoryEnvironmentDeploymentPolicyState.__new__(_RepositoryEnvironmentDeploymentPolicyState)

        __props__.__dict__["branch_pattern"] = branch_pattern
        __props__.__dict__["environment"] = environment
        __props__.__dict__["repository"] = repository
        __props__.__dict__["tag_pattern"] = tag_pattern
        return RepositoryEnvironmentDeploymentPolicy(resource_name, opts=opts, __props__=__props__)

    @_builtins.property
    @pulumi.getter(name="branchPattern")
    def branch_pattern(self) -> pulumi.Output[Optional[_builtins.str]]:
        """
        The name pattern that branches must match in order to deploy to the environment. If not specified, `tag_pattern` must be specified.
        """
        return pulumi.get(self, "branch_pattern")

    @_builtins.property
    @pulumi.getter
    def environment(self) -> pulumi.Output[_builtins.str]:
        """
        The name of the environment.
        """
        return pulumi.get(self, "environment")

    @_builtins.property
    @pulumi.getter
    def repository(self) -> pulumi.Output[_builtins.str]:
        """
        The repository of the environment.
        """
        return pulumi.get(self, "repository")

    @_builtins.property
    @pulumi.getter(name="tagPattern")
    def tag_pattern(self) -> pulumi.Output[Optional[_builtins.str]]:
        """
        The name pattern that tags must match in order to deploy to the environment. If not specified, `branch_pattern` must be specified.
        """
        return pulumi.get(self, "tag_pattern")

