# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
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

__all__ = ['RepositoryEnvironmentArgs', 'RepositoryEnvironment']

@pulumi.input_type
class RepositoryEnvironmentArgs:
    def __init__(__self__, *,
                 environment: pulumi.Input[str],
                 repository: pulumi.Input[str],
                 can_admins_bypass: Optional[pulumi.Input[bool]] = None,
                 deployment_branch_policy: Optional[pulumi.Input['RepositoryEnvironmentDeploymentBranchPolicyArgs']] = None,
                 prevent_self_review: Optional[pulumi.Input[bool]] = None,
                 reviewers: Optional[pulumi.Input[Sequence[pulumi.Input['RepositoryEnvironmentReviewerArgs']]]] = None,
                 wait_timer: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a RepositoryEnvironment resource.
        :param pulumi.Input[str] environment: The name of the environment.
        :param pulumi.Input[str] repository: The repository of the environment.
        :param pulumi.Input[bool] can_admins_bypass: Can repository admins bypass the environment protections.  Defaults to `true`.
        :param pulumi.Input['RepositoryEnvironmentDeploymentBranchPolicyArgs'] deployment_branch_policy: The deployment branch policy configuration
        :param pulumi.Input[bool] prevent_self_review: Whether or not a user who created the job is prevented from approving their own job. Defaults to `false`.
        :param pulumi.Input[Sequence[pulumi.Input['RepositoryEnvironmentReviewerArgs']]] reviewers: The environment reviewers configuration.
        :param pulumi.Input[int] wait_timer: Amount of time to delay a job after the job is initially triggered.
        """
        pulumi.set(__self__, "environment", environment)
        pulumi.set(__self__, "repository", repository)
        if can_admins_bypass is not None:
            pulumi.set(__self__, "can_admins_bypass", can_admins_bypass)
        if deployment_branch_policy is not None:
            pulumi.set(__self__, "deployment_branch_policy", deployment_branch_policy)
        if prevent_self_review is not None:
            pulumi.set(__self__, "prevent_self_review", prevent_self_review)
        if reviewers is not None:
            pulumi.set(__self__, "reviewers", reviewers)
        if wait_timer is not None:
            pulumi.set(__self__, "wait_timer", wait_timer)

    @property
    @pulumi.getter
    def environment(self) -> pulumi.Input[str]:
        """
        The name of the environment.
        """
        return pulumi.get(self, "environment")

    @environment.setter
    def environment(self, value: pulumi.Input[str]):
        pulumi.set(self, "environment", value)

    @property
    @pulumi.getter
    def repository(self) -> pulumi.Input[str]:
        """
        The repository of the environment.
        """
        return pulumi.get(self, "repository")

    @repository.setter
    def repository(self, value: pulumi.Input[str]):
        pulumi.set(self, "repository", value)

    @property
    @pulumi.getter(name="canAdminsBypass")
    def can_admins_bypass(self) -> Optional[pulumi.Input[bool]]:
        """
        Can repository admins bypass the environment protections.  Defaults to `true`.
        """
        return pulumi.get(self, "can_admins_bypass")

    @can_admins_bypass.setter
    def can_admins_bypass(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "can_admins_bypass", value)

    @property
    @pulumi.getter(name="deploymentBranchPolicy")
    def deployment_branch_policy(self) -> Optional[pulumi.Input['RepositoryEnvironmentDeploymentBranchPolicyArgs']]:
        """
        The deployment branch policy configuration
        """
        return pulumi.get(self, "deployment_branch_policy")

    @deployment_branch_policy.setter
    def deployment_branch_policy(self, value: Optional[pulumi.Input['RepositoryEnvironmentDeploymentBranchPolicyArgs']]):
        pulumi.set(self, "deployment_branch_policy", value)

    @property
    @pulumi.getter(name="preventSelfReview")
    def prevent_self_review(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether or not a user who created the job is prevented from approving their own job. Defaults to `false`.
        """
        return pulumi.get(self, "prevent_self_review")

    @prevent_self_review.setter
    def prevent_self_review(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "prevent_self_review", value)

    @property
    @pulumi.getter
    def reviewers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RepositoryEnvironmentReviewerArgs']]]]:
        """
        The environment reviewers configuration.
        """
        return pulumi.get(self, "reviewers")

    @reviewers.setter
    def reviewers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RepositoryEnvironmentReviewerArgs']]]]):
        pulumi.set(self, "reviewers", value)

    @property
    @pulumi.getter(name="waitTimer")
    def wait_timer(self) -> Optional[pulumi.Input[int]]:
        """
        Amount of time to delay a job after the job is initially triggered.
        """
        return pulumi.get(self, "wait_timer")

    @wait_timer.setter
    def wait_timer(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "wait_timer", value)


@pulumi.input_type
class _RepositoryEnvironmentState:
    def __init__(__self__, *,
                 can_admins_bypass: Optional[pulumi.Input[bool]] = None,
                 deployment_branch_policy: Optional[pulumi.Input['RepositoryEnvironmentDeploymentBranchPolicyArgs']] = None,
                 environment: Optional[pulumi.Input[str]] = None,
                 prevent_self_review: Optional[pulumi.Input[bool]] = None,
                 repository: Optional[pulumi.Input[str]] = None,
                 reviewers: Optional[pulumi.Input[Sequence[pulumi.Input['RepositoryEnvironmentReviewerArgs']]]] = None,
                 wait_timer: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering RepositoryEnvironment resources.
        :param pulumi.Input[bool] can_admins_bypass: Can repository admins bypass the environment protections.  Defaults to `true`.
        :param pulumi.Input['RepositoryEnvironmentDeploymentBranchPolicyArgs'] deployment_branch_policy: The deployment branch policy configuration
        :param pulumi.Input[str] environment: The name of the environment.
        :param pulumi.Input[bool] prevent_self_review: Whether or not a user who created the job is prevented from approving their own job. Defaults to `false`.
        :param pulumi.Input[str] repository: The repository of the environment.
        :param pulumi.Input[Sequence[pulumi.Input['RepositoryEnvironmentReviewerArgs']]] reviewers: The environment reviewers configuration.
        :param pulumi.Input[int] wait_timer: Amount of time to delay a job after the job is initially triggered.
        """
        if can_admins_bypass is not None:
            pulumi.set(__self__, "can_admins_bypass", can_admins_bypass)
        if deployment_branch_policy is not None:
            pulumi.set(__self__, "deployment_branch_policy", deployment_branch_policy)
        if environment is not None:
            pulumi.set(__self__, "environment", environment)
        if prevent_self_review is not None:
            pulumi.set(__self__, "prevent_self_review", prevent_self_review)
        if repository is not None:
            pulumi.set(__self__, "repository", repository)
        if reviewers is not None:
            pulumi.set(__self__, "reviewers", reviewers)
        if wait_timer is not None:
            pulumi.set(__self__, "wait_timer", wait_timer)

    @property
    @pulumi.getter(name="canAdminsBypass")
    def can_admins_bypass(self) -> Optional[pulumi.Input[bool]]:
        """
        Can repository admins bypass the environment protections.  Defaults to `true`.
        """
        return pulumi.get(self, "can_admins_bypass")

    @can_admins_bypass.setter
    def can_admins_bypass(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "can_admins_bypass", value)

    @property
    @pulumi.getter(name="deploymentBranchPolicy")
    def deployment_branch_policy(self) -> Optional[pulumi.Input['RepositoryEnvironmentDeploymentBranchPolicyArgs']]:
        """
        The deployment branch policy configuration
        """
        return pulumi.get(self, "deployment_branch_policy")

    @deployment_branch_policy.setter
    def deployment_branch_policy(self, value: Optional[pulumi.Input['RepositoryEnvironmentDeploymentBranchPolicyArgs']]):
        pulumi.set(self, "deployment_branch_policy", value)

    @property
    @pulumi.getter
    def environment(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the environment.
        """
        return pulumi.get(self, "environment")

    @environment.setter
    def environment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "environment", value)

    @property
    @pulumi.getter(name="preventSelfReview")
    def prevent_self_review(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether or not a user who created the job is prevented from approving their own job. Defaults to `false`.
        """
        return pulumi.get(self, "prevent_self_review")

    @prevent_self_review.setter
    def prevent_self_review(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "prevent_self_review", value)

    @property
    @pulumi.getter
    def repository(self) -> Optional[pulumi.Input[str]]:
        """
        The repository of the environment.
        """
        return pulumi.get(self, "repository")

    @repository.setter
    def repository(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "repository", value)

    @property
    @pulumi.getter
    def reviewers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RepositoryEnvironmentReviewerArgs']]]]:
        """
        The environment reviewers configuration.
        """
        return pulumi.get(self, "reviewers")

    @reviewers.setter
    def reviewers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RepositoryEnvironmentReviewerArgs']]]]):
        pulumi.set(self, "reviewers", value)

    @property
    @pulumi.getter(name="waitTimer")
    def wait_timer(self) -> Optional[pulumi.Input[int]]:
        """
        Amount of time to delay a job after the job is initially triggered.
        """
        return pulumi.get(self, "wait_timer")

    @wait_timer.setter
    def wait_timer(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "wait_timer", value)


class RepositoryEnvironment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 can_admins_bypass: Optional[pulumi.Input[bool]] = None,
                 deployment_branch_policy: Optional[pulumi.Input[Union['RepositoryEnvironmentDeploymentBranchPolicyArgs', 'RepositoryEnvironmentDeploymentBranchPolicyArgsDict']]] = None,
                 environment: Optional[pulumi.Input[str]] = None,
                 prevent_self_review: Optional[pulumi.Input[bool]] = None,
                 repository: Optional[pulumi.Input[str]] = None,
                 reviewers: Optional[pulumi.Input[Sequence[pulumi.Input[Union['RepositoryEnvironmentReviewerArgs', 'RepositoryEnvironmentReviewerArgsDict']]]]] = None,
                 wait_timer: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        This resource allows you to create and manage environments for a GitHub repository.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_github as github

        current = github.get_user(username="")
        example = github.Repository("example",
            name="A Repository Project",
            description="My awesome codebase")
        example_repository_environment = github.RepositoryEnvironment("example",
            environment="example",
            repository=example.name,
            prevent_self_review=True,
            reviewers=[{
                "users": [current.id],
            }],
            deployment_branch_policy={
                "protected_branches": True,
                "custom_branch_policies": False,
            })
        ```

        ## Import

        GitHub Repository Environment can be imported using an ID made up of `name` of the repository combined with the `environment` name of the environment, separated by a `:` character, e.g.

        ```sh
        $ pulumi import github:index/repositoryEnvironment:RepositoryEnvironment daily terraform:daily
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] can_admins_bypass: Can repository admins bypass the environment protections.  Defaults to `true`.
        :param pulumi.Input[Union['RepositoryEnvironmentDeploymentBranchPolicyArgs', 'RepositoryEnvironmentDeploymentBranchPolicyArgsDict']] deployment_branch_policy: The deployment branch policy configuration
        :param pulumi.Input[str] environment: The name of the environment.
        :param pulumi.Input[bool] prevent_self_review: Whether or not a user who created the job is prevented from approving their own job. Defaults to `false`.
        :param pulumi.Input[str] repository: The repository of the environment.
        :param pulumi.Input[Sequence[pulumi.Input[Union['RepositoryEnvironmentReviewerArgs', 'RepositoryEnvironmentReviewerArgsDict']]]] reviewers: The environment reviewers configuration.
        :param pulumi.Input[int] wait_timer: Amount of time to delay a job after the job is initially triggered.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RepositoryEnvironmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This resource allows you to create and manage environments for a GitHub repository.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_github as github

        current = github.get_user(username="")
        example = github.Repository("example",
            name="A Repository Project",
            description="My awesome codebase")
        example_repository_environment = github.RepositoryEnvironment("example",
            environment="example",
            repository=example.name,
            prevent_self_review=True,
            reviewers=[{
                "users": [current.id],
            }],
            deployment_branch_policy={
                "protected_branches": True,
                "custom_branch_policies": False,
            })
        ```

        ## Import

        GitHub Repository Environment can be imported using an ID made up of `name` of the repository combined with the `environment` name of the environment, separated by a `:` character, e.g.

        ```sh
        $ pulumi import github:index/repositoryEnvironment:RepositoryEnvironment daily terraform:daily
        ```

        :param str resource_name: The name of the resource.
        :param RepositoryEnvironmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RepositoryEnvironmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 can_admins_bypass: Optional[pulumi.Input[bool]] = None,
                 deployment_branch_policy: Optional[pulumi.Input[Union['RepositoryEnvironmentDeploymentBranchPolicyArgs', 'RepositoryEnvironmentDeploymentBranchPolicyArgsDict']]] = None,
                 environment: Optional[pulumi.Input[str]] = None,
                 prevent_self_review: Optional[pulumi.Input[bool]] = None,
                 repository: Optional[pulumi.Input[str]] = None,
                 reviewers: Optional[pulumi.Input[Sequence[pulumi.Input[Union['RepositoryEnvironmentReviewerArgs', 'RepositoryEnvironmentReviewerArgsDict']]]]] = None,
                 wait_timer: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RepositoryEnvironmentArgs.__new__(RepositoryEnvironmentArgs)

            __props__.__dict__["can_admins_bypass"] = can_admins_bypass
            __props__.__dict__["deployment_branch_policy"] = deployment_branch_policy
            if environment is None and not opts.urn:
                raise TypeError("Missing required property 'environment'")
            __props__.__dict__["environment"] = environment
            __props__.__dict__["prevent_self_review"] = prevent_self_review
            if repository is None and not opts.urn:
                raise TypeError("Missing required property 'repository'")
            __props__.__dict__["repository"] = repository
            __props__.__dict__["reviewers"] = reviewers
            __props__.__dict__["wait_timer"] = wait_timer
        super(RepositoryEnvironment, __self__).__init__(
            'github:index/repositoryEnvironment:RepositoryEnvironment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            can_admins_bypass: Optional[pulumi.Input[bool]] = None,
            deployment_branch_policy: Optional[pulumi.Input[Union['RepositoryEnvironmentDeploymentBranchPolicyArgs', 'RepositoryEnvironmentDeploymentBranchPolicyArgsDict']]] = None,
            environment: Optional[pulumi.Input[str]] = None,
            prevent_self_review: Optional[pulumi.Input[bool]] = None,
            repository: Optional[pulumi.Input[str]] = None,
            reviewers: Optional[pulumi.Input[Sequence[pulumi.Input[Union['RepositoryEnvironmentReviewerArgs', 'RepositoryEnvironmentReviewerArgsDict']]]]] = None,
            wait_timer: Optional[pulumi.Input[int]] = None) -> 'RepositoryEnvironment':
        """
        Get an existing RepositoryEnvironment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] can_admins_bypass: Can repository admins bypass the environment protections.  Defaults to `true`.
        :param pulumi.Input[Union['RepositoryEnvironmentDeploymentBranchPolicyArgs', 'RepositoryEnvironmentDeploymentBranchPolicyArgsDict']] deployment_branch_policy: The deployment branch policy configuration
        :param pulumi.Input[str] environment: The name of the environment.
        :param pulumi.Input[bool] prevent_self_review: Whether or not a user who created the job is prevented from approving their own job. Defaults to `false`.
        :param pulumi.Input[str] repository: The repository of the environment.
        :param pulumi.Input[Sequence[pulumi.Input[Union['RepositoryEnvironmentReviewerArgs', 'RepositoryEnvironmentReviewerArgsDict']]]] reviewers: The environment reviewers configuration.
        :param pulumi.Input[int] wait_timer: Amount of time to delay a job after the job is initially triggered.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RepositoryEnvironmentState.__new__(_RepositoryEnvironmentState)

        __props__.__dict__["can_admins_bypass"] = can_admins_bypass
        __props__.__dict__["deployment_branch_policy"] = deployment_branch_policy
        __props__.__dict__["environment"] = environment
        __props__.__dict__["prevent_self_review"] = prevent_self_review
        __props__.__dict__["repository"] = repository
        __props__.__dict__["reviewers"] = reviewers
        __props__.__dict__["wait_timer"] = wait_timer
        return RepositoryEnvironment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="canAdminsBypass")
    def can_admins_bypass(self) -> pulumi.Output[Optional[bool]]:
        """
        Can repository admins bypass the environment protections.  Defaults to `true`.
        """
        return pulumi.get(self, "can_admins_bypass")

    @property
    @pulumi.getter(name="deploymentBranchPolicy")
    def deployment_branch_policy(self) -> pulumi.Output[Optional['outputs.RepositoryEnvironmentDeploymentBranchPolicy']]:
        """
        The deployment branch policy configuration
        """
        return pulumi.get(self, "deployment_branch_policy")

    @property
    @pulumi.getter
    def environment(self) -> pulumi.Output[str]:
        """
        The name of the environment.
        """
        return pulumi.get(self, "environment")

    @property
    @pulumi.getter(name="preventSelfReview")
    def prevent_self_review(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether or not a user who created the job is prevented from approving their own job. Defaults to `false`.
        """
        return pulumi.get(self, "prevent_self_review")

    @property
    @pulumi.getter
    def repository(self) -> pulumi.Output[str]:
        """
        The repository of the environment.
        """
        return pulumi.get(self, "repository")

    @property
    @pulumi.getter
    def reviewers(self) -> pulumi.Output[Optional[Sequence['outputs.RepositoryEnvironmentReviewer']]]:
        """
        The environment reviewers configuration.
        """
        return pulumi.get(self, "reviewers")

    @property
    @pulumi.getter(name="waitTimer")
    def wait_timer(self) -> pulumi.Output[Optional[int]]:
        """
        Amount of time to delay a job after the job is initially triggered.
        """
        return pulumi.get(self, "wait_timer")

