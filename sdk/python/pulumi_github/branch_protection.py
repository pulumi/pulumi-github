# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['BranchProtectionArgs', 'BranchProtection']

@pulumi.input_type
class BranchProtectionArgs:
    def __init__(__self__, *,
                 pattern: pulumi.Input[str],
                 repository_id: pulumi.Input[str],
                 allows_deletions: Optional[pulumi.Input[bool]] = None,
                 allows_force_pushes: Optional[pulumi.Input[bool]] = None,
                 blocks_creations: Optional[pulumi.Input[bool]] = None,
                 enforce_admins: Optional[pulumi.Input[bool]] = None,
                 lock_branch: Optional[pulumi.Input[bool]] = None,
                 push_restrictions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 require_conversation_resolution: Optional[pulumi.Input[bool]] = None,
                 require_signed_commits: Optional[pulumi.Input[bool]] = None,
                 required_linear_history: Optional[pulumi.Input[bool]] = None,
                 required_pull_request_reviews: Optional[pulumi.Input[Sequence[pulumi.Input['BranchProtectionRequiredPullRequestReviewArgs']]]] = None,
                 required_status_checks: Optional[pulumi.Input[Sequence[pulumi.Input['BranchProtectionRequiredStatusCheckArgs']]]] = None):
        """
        The set of arguments for constructing a BranchProtection resource.
        :param pulumi.Input[str] repository_id: Node ID or name of repository
        """
        pulumi.set(__self__, "pattern", pattern)
        pulumi.set(__self__, "repository_id", repository_id)
        if allows_deletions is not None:
            pulumi.set(__self__, "allows_deletions", allows_deletions)
        if allows_force_pushes is not None:
            pulumi.set(__self__, "allows_force_pushes", allows_force_pushes)
        if blocks_creations is not None:
            pulumi.set(__self__, "blocks_creations", blocks_creations)
        if enforce_admins is not None:
            pulumi.set(__self__, "enforce_admins", enforce_admins)
        if lock_branch is not None:
            pulumi.set(__self__, "lock_branch", lock_branch)
        if push_restrictions is not None:
            pulumi.set(__self__, "push_restrictions", push_restrictions)
        if require_conversation_resolution is not None:
            pulumi.set(__self__, "require_conversation_resolution", require_conversation_resolution)
        if require_signed_commits is not None:
            pulumi.set(__self__, "require_signed_commits", require_signed_commits)
        if required_linear_history is not None:
            pulumi.set(__self__, "required_linear_history", required_linear_history)
        if required_pull_request_reviews is not None:
            pulumi.set(__self__, "required_pull_request_reviews", required_pull_request_reviews)
        if required_status_checks is not None:
            pulumi.set(__self__, "required_status_checks", required_status_checks)

    @property
    @pulumi.getter
    def pattern(self) -> pulumi.Input[str]:
        return pulumi.get(self, "pattern")

    @pattern.setter
    def pattern(self, value: pulumi.Input[str]):
        pulumi.set(self, "pattern", value)

    @property
    @pulumi.getter(name="repositoryId")
    def repository_id(self) -> pulumi.Input[str]:
        """
        Node ID or name of repository
        """
        return pulumi.get(self, "repository_id")

    @repository_id.setter
    def repository_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "repository_id", value)

    @property
    @pulumi.getter(name="allowsDeletions")
    def allows_deletions(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "allows_deletions")

    @allows_deletions.setter
    def allows_deletions(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "allows_deletions", value)

    @property
    @pulumi.getter(name="allowsForcePushes")
    def allows_force_pushes(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "allows_force_pushes")

    @allows_force_pushes.setter
    def allows_force_pushes(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "allows_force_pushes", value)

    @property
    @pulumi.getter(name="blocksCreations")
    def blocks_creations(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "blocks_creations")

    @blocks_creations.setter
    def blocks_creations(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "blocks_creations", value)

    @property
    @pulumi.getter(name="enforceAdmins")
    def enforce_admins(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "enforce_admins")

    @enforce_admins.setter
    def enforce_admins(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enforce_admins", value)

    @property
    @pulumi.getter(name="lockBranch")
    def lock_branch(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "lock_branch")

    @lock_branch.setter
    def lock_branch(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "lock_branch", value)

    @property
    @pulumi.getter(name="pushRestrictions")
    def push_restrictions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "push_restrictions")

    @push_restrictions.setter
    def push_restrictions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "push_restrictions", value)

    @property
    @pulumi.getter(name="requireConversationResolution")
    def require_conversation_resolution(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "require_conversation_resolution")

    @require_conversation_resolution.setter
    def require_conversation_resolution(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "require_conversation_resolution", value)

    @property
    @pulumi.getter(name="requireSignedCommits")
    def require_signed_commits(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "require_signed_commits")

    @require_signed_commits.setter
    def require_signed_commits(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "require_signed_commits", value)

    @property
    @pulumi.getter(name="requiredLinearHistory")
    def required_linear_history(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "required_linear_history")

    @required_linear_history.setter
    def required_linear_history(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "required_linear_history", value)

    @property
    @pulumi.getter(name="requiredPullRequestReviews")
    def required_pull_request_reviews(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['BranchProtectionRequiredPullRequestReviewArgs']]]]:
        return pulumi.get(self, "required_pull_request_reviews")

    @required_pull_request_reviews.setter
    def required_pull_request_reviews(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['BranchProtectionRequiredPullRequestReviewArgs']]]]):
        pulumi.set(self, "required_pull_request_reviews", value)

    @property
    @pulumi.getter(name="requiredStatusChecks")
    def required_status_checks(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['BranchProtectionRequiredStatusCheckArgs']]]]:
        return pulumi.get(self, "required_status_checks")

    @required_status_checks.setter
    def required_status_checks(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['BranchProtectionRequiredStatusCheckArgs']]]]):
        pulumi.set(self, "required_status_checks", value)


@pulumi.input_type
class _BranchProtectionState:
    def __init__(__self__, *,
                 allows_deletions: Optional[pulumi.Input[bool]] = None,
                 allows_force_pushes: Optional[pulumi.Input[bool]] = None,
                 blocks_creations: Optional[pulumi.Input[bool]] = None,
                 enforce_admins: Optional[pulumi.Input[bool]] = None,
                 lock_branch: Optional[pulumi.Input[bool]] = None,
                 pattern: Optional[pulumi.Input[str]] = None,
                 push_restrictions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 repository_id: Optional[pulumi.Input[str]] = None,
                 require_conversation_resolution: Optional[pulumi.Input[bool]] = None,
                 require_signed_commits: Optional[pulumi.Input[bool]] = None,
                 required_linear_history: Optional[pulumi.Input[bool]] = None,
                 required_pull_request_reviews: Optional[pulumi.Input[Sequence[pulumi.Input['BranchProtectionRequiredPullRequestReviewArgs']]]] = None,
                 required_status_checks: Optional[pulumi.Input[Sequence[pulumi.Input['BranchProtectionRequiredStatusCheckArgs']]]] = None):
        """
        Input properties used for looking up and filtering BranchProtection resources.
        :param pulumi.Input[str] repository_id: Node ID or name of repository
        """
        if allows_deletions is not None:
            pulumi.set(__self__, "allows_deletions", allows_deletions)
        if allows_force_pushes is not None:
            pulumi.set(__self__, "allows_force_pushes", allows_force_pushes)
        if blocks_creations is not None:
            pulumi.set(__self__, "blocks_creations", blocks_creations)
        if enforce_admins is not None:
            pulumi.set(__self__, "enforce_admins", enforce_admins)
        if lock_branch is not None:
            pulumi.set(__self__, "lock_branch", lock_branch)
        if pattern is not None:
            pulumi.set(__self__, "pattern", pattern)
        if push_restrictions is not None:
            pulumi.set(__self__, "push_restrictions", push_restrictions)
        if repository_id is not None:
            pulumi.set(__self__, "repository_id", repository_id)
        if require_conversation_resolution is not None:
            pulumi.set(__self__, "require_conversation_resolution", require_conversation_resolution)
        if require_signed_commits is not None:
            pulumi.set(__self__, "require_signed_commits", require_signed_commits)
        if required_linear_history is not None:
            pulumi.set(__self__, "required_linear_history", required_linear_history)
        if required_pull_request_reviews is not None:
            pulumi.set(__self__, "required_pull_request_reviews", required_pull_request_reviews)
        if required_status_checks is not None:
            pulumi.set(__self__, "required_status_checks", required_status_checks)

    @property
    @pulumi.getter(name="allowsDeletions")
    def allows_deletions(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "allows_deletions")

    @allows_deletions.setter
    def allows_deletions(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "allows_deletions", value)

    @property
    @pulumi.getter(name="allowsForcePushes")
    def allows_force_pushes(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "allows_force_pushes")

    @allows_force_pushes.setter
    def allows_force_pushes(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "allows_force_pushes", value)

    @property
    @pulumi.getter(name="blocksCreations")
    def blocks_creations(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "blocks_creations")

    @blocks_creations.setter
    def blocks_creations(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "blocks_creations", value)

    @property
    @pulumi.getter(name="enforceAdmins")
    def enforce_admins(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "enforce_admins")

    @enforce_admins.setter
    def enforce_admins(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enforce_admins", value)

    @property
    @pulumi.getter(name="lockBranch")
    def lock_branch(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "lock_branch")

    @lock_branch.setter
    def lock_branch(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "lock_branch", value)

    @property
    @pulumi.getter
    def pattern(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "pattern")

    @pattern.setter
    def pattern(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "pattern", value)

    @property
    @pulumi.getter(name="pushRestrictions")
    def push_restrictions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "push_restrictions")

    @push_restrictions.setter
    def push_restrictions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "push_restrictions", value)

    @property
    @pulumi.getter(name="repositoryId")
    def repository_id(self) -> Optional[pulumi.Input[str]]:
        """
        Node ID or name of repository
        """
        return pulumi.get(self, "repository_id")

    @repository_id.setter
    def repository_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "repository_id", value)

    @property
    @pulumi.getter(name="requireConversationResolution")
    def require_conversation_resolution(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "require_conversation_resolution")

    @require_conversation_resolution.setter
    def require_conversation_resolution(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "require_conversation_resolution", value)

    @property
    @pulumi.getter(name="requireSignedCommits")
    def require_signed_commits(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "require_signed_commits")

    @require_signed_commits.setter
    def require_signed_commits(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "require_signed_commits", value)

    @property
    @pulumi.getter(name="requiredLinearHistory")
    def required_linear_history(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "required_linear_history")

    @required_linear_history.setter
    def required_linear_history(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "required_linear_history", value)

    @property
    @pulumi.getter(name="requiredPullRequestReviews")
    def required_pull_request_reviews(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['BranchProtectionRequiredPullRequestReviewArgs']]]]:
        return pulumi.get(self, "required_pull_request_reviews")

    @required_pull_request_reviews.setter
    def required_pull_request_reviews(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['BranchProtectionRequiredPullRequestReviewArgs']]]]):
        pulumi.set(self, "required_pull_request_reviews", value)

    @property
    @pulumi.getter(name="requiredStatusChecks")
    def required_status_checks(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['BranchProtectionRequiredStatusCheckArgs']]]]:
        return pulumi.get(self, "required_status_checks")

    @required_status_checks.setter
    def required_status_checks(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['BranchProtectionRequiredStatusCheckArgs']]]]):
        pulumi.set(self, "required_status_checks", value)


class BranchProtection(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allows_deletions: Optional[pulumi.Input[bool]] = None,
                 allows_force_pushes: Optional[pulumi.Input[bool]] = None,
                 blocks_creations: Optional[pulumi.Input[bool]] = None,
                 enforce_admins: Optional[pulumi.Input[bool]] = None,
                 lock_branch: Optional[pulumi.Input[bool]] = None,
                 pattern: Optional[pulumi.Input[str]] = None,
                 push_restrictions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 repository_id: Optional[pulumi.Input[str]] = None,
                 require_conversation_resolution: Optional[pulumi.Input[bool]] = None,
                 require_signed_commits: Optional[pulumi.Input[bool]] = None,
                 required_linear_history: Optional[pulumi.Input[bool]] = None,
                 required_pull_request_reviews: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BranchProtectionRequiredPullRequestReviewArgs']]]]] = None,
                 required_status_checks: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BranchProtectionRequiredStatusCheckArgs']]]]] = None,
                 __props__=None):
        """
        Create a BranchProtection resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] repository_id: Node ID or name of repository
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: BranchProtectionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a BranchProtection resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param BranchProtectionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BranchProtectionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allows_deletions: Optional[pulumi.Input[bool]] = None,
                 allows_force_pushes: Optional[pulumi.Input[bool]] = None,
                 blocks_creations: Optional[pulumi.Input[bool]] = None,
                 enforce_admins: Optional[pulumi.Input[bool]] = None,
                 lock_branch: Optional[pulumi.Input[bool]] = None,
                 pattern: Optional[pulumi.Input[str]] = None,
                 push_restrictions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 repository_id: Optional[pulumi.Input[str]] = None,
                 require_conversation_resolution: Optional[pulumi.Input[bool]] = None,
                 require_signed_commits: Optional[pulumi.Input[bool]] = None,
                 required_linear_history: Optional[pulumi.Input[bool]] = None,
                 required_pull_request_reviews: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BranchProtectionRequiredPullRequestReviewArgs']]]]] = None,
                 required_status_checks: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BranchProtectionRequiredStatusCheckArgs']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = BranchProtectionArgs.__new__(BranchProtectionArgs)

            __props__.__dict__["allows_deletions"] = allows_deletions
            __props__.__dict__["allows_force_pushes"] = allows_force_pushes
            __props__.__dict__["blocks_creations"] = blocks_creations
            __props__.__dict__["enforce_admins"] = enforce_admins
            __props__.__dict__["lock_branch"] = lock_branch
            if pattern is None and not opts.urn:
                raise TypeError("Missing required property 'pattern'")
            __props__.__dict__["pattern"] = pattern
            __props__.__dict__["push_restrictions"] = push_restrictions
            if repository_id is None and not opts.urn:
                raise TypeError("Missing required property 'repository_id'")
            __props__.__dict__["repository_id"] = repository_id
            __props__.__dict__["require_conversation_resolution"] = require_conversation_resolution
            __props__.__dict__["require_signed_commits"] = require_signed_commits
            __props__.__dict__["required_linear_history"] = required_linear_history
            __props__.__dict__["required_pull_request_reviews"] = required_pull_request_reviews
            __props__.__dict__["required_status_checks"] = required_status_checks
        super(BranchProtection, __self__).__init__(
            'github:index/branchProtection:BranchProtection',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            allows_deletions: Optional[pulumi.Input[bool]] = None,
            allows_force_pushes: Optional[pulumi.Input[bool]] = None,
            blocks_creations: Optional[pulumi.Input[bool]] = None,
            enforce_admins: Optional[pulumi.Input[bool]] = None,
            lock_branch: Optional[pulumi.Input[bool]] = None,
            pattern: Optional[pulumi.Input[str]] = None,
            push_restrictions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            repository_id: Optional[pulumi.Input[str]] = None,
            require_conversation_resolution: Optional[pulumi.Input[bool]] = None,
            require_signed_commits: Optional[pulumi.Input[bool]] = None,
            required_linear_history: Optional[pulumi.Input[bool]] = None,
            required_pull_request_reviews: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BranchProtectionRequiredPullRequestReviewArgs']]]]] = None,
            required_status_checks: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BranchProtectionRequiredStatusCheckArgs']]]]] = None) -> 'BranchProtection':
        """
        Get an existing BranchProtection resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] repository_id: Node ID or name of repository
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _BranchProtectionState.__new__(_BranchProtectionState)

        __props__.__dict__["allows_deletions"] = allows_deletions
        __props__.__dict__["allows_force_pushes"] = allows_force_pushes
        __props__.__dict__["blocks_creations"] = blocks_creations
        __props__.__dict__["enforce_admins"] = enforce_admins
        __props__.__dict__["lock_branch"] = lock_branch
        __props__.__dict__["pattern"] = pattern
        __props__.__dict__["push_restrictions"] = push_restrictions
        __props__.__dict__["repository_id"] = repository_id
        __props__.__dict__["require_conversation_resolution"] = require_conversation_resolution
        __props__.__dict__["require_signed_commits"] = require_signed_commits
        __props__.__dict__["required_linear_history"] = required_linear_history
        __props__.__dict__["required_pull_request_reviews"] = required_pull_request_reviews
        __props__.__dict__["required_status_checks"] = required_status_checks
        return BranchProtection(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="allowsDeletions")
    def allows_deletions(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "allows_deletions")

    @property
    @pulumi.getter(name="allowsForcePushes")
    def allows_force_pushes(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "allows_force_pushes")

    @property
    @pulumi.getter(name="blocksCreations")
    def blocks_creations(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "blocks_creations")

    @property
    @pulumi.getter(name="enforceAdmins")
    def enforce_admins(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "enforce_admins")

    @property
    @pulumi.getter(name="lockBranch")
    def lock_branch(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "lock_branch")

    @property
    @pulumi.getter
    def pattern(self) -> pulumi.Output[str]:
        return pulumi.get(self, "pattern")

    @property
    @pulumi.getter(name="pushRestrictions")
    def push_restrictions(self) -> pulumi.Output[Optional[Sequence[str]]]:
        return pulumi.get(self, "push_restrictions")

    @property
    @pulumi.getter(name="repositoryId")
    def repository_id(self) -> pulumi.Output[str]:
        """
        Node ID or name of repository
        """
        return pulumi.get(self, "repository_id")

    @property
    @pulumi.getter(name="requireConversationResolution")
    def require_conversation_resolution(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "require_conversation_resolution")

    @property
    @pulumi.getter(name="requireSignedCommits")
    def require_signed_commits(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "require_signed_commits")

    @property
    @pulumi.getter(name="requiredLinearHistory")
    def required_linear_history(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "required_linear_history")

    @property
    @pulumi.getter(name="requiredPullRequestReviews")
    def required_pull_request_reviews(self) -> pulumi.Output[Optional[Sequence['outputs.BranchProtectionRequiredPullRequestReview']]]:
        return pulumi.get(self, "required_pull_request_reviews")

    @property
    @pulumi.getter(name="requiredStatusChecks")
    def required_status_checks(self) -> pulumi.Output[Optional[Sequence['outputs.BranchProtectionRequiredStatusCheck']]]:
        return pulumi.get(self, "required_status_checks")

