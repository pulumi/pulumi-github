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

__all__ = ['BranchProtectionV3Args', 'BranchProtectionV3']

@pulumi.input_type
class BranchProtectionV3Args:
    def __init__(__self__, *,
                 branch: pulumi.Input[str],
                 repository: pulumi.Input[str],
                 enforce_admins: Optional[pulumi.Input[bool]] = None,
                 require_conversation_resolution: Optional[pulumi.Input[bool]] = None,
                 require_signed_commits: Optional[pulumi.Input[bool]] = None,
                 required_pull_request_reviews: Optional[pulumi.Input['BranchProtectionV3RequiredPullRequestReviewsArgs']] = None,
                 required_status_checks: Optional[pulumi.Input['BranchProtectionV3RequiredStatusChecksArgs']] = None,
                 restrictions: Optional[pulumi.Input['BranchProtectionV3RestrictionsArgs']] = None):
        """
        The set of arguments for constructing a BranchProtectionV3 resource.
        :param pulumi.Input[str] branch: The Git branch to protect.
        :param pulumi.Input[str] repository: The GitHub repository name.
        :param pulumi.Input[bool] enforce_admins: Boolean, setting this to `true` enforces status checks for repository administrators.
        :param pulumi.Input[bool] require_conversation_resolution: Boolean, setting this to `true` requires all conversations on code must be resolved before a pull request can be merged.
        :param pulumi.Input[bool] require_signed_commits: Boolean, setting this to `true` requires all commits to be signed with GPG.
        :param pulumi.Input['BranchProtectionV3RequiredPullRequestReviewsArgs'] required_pull_request_reviews: Enforce restrictions for pull request reviews. See Required Pull Request Reviews below for details.
        :param pulumi.Input['BranchProtectionV3RequiredStatusChecksArgs'] required_status_checks: Enforce restrictions for required status checks. See Required Status Checks below for details.
        :param pulumi.Input['BranchProtectionV3RestrictionsArgs'] restrictions: Enforce restrictions for the users and teams that may push to the branch. See Restrictions below for details.
        """
        pulumi.set(__self__, "branch", branch)
        pulumi.set(__self__, "repository", repository)
        if enforce_admins is not None:
            pulumi.set(__self__, "enforce_admins", enforce_admins)
        if require_conversation_resolution is not None:
            pulumi.set(__self__, "require_conversation_resolution", require_conversation_resolution)
        if require_signed_commits is not None:
            pulumi.set(__self__, "require_signed_commits", require_signed_commits)
        if required_pull_request_reviews is not None:
            pulumi.set(__self__, "required_pull_request_reviews", required_pull_request_reviews)
        if required_status_checks is not None:
            pulumi.set(__self__, "required_status_checks", required_status_checks)
        if restrictions is not None:
            pulumi.set(__self__, "restrictions", restrictions)

    @property
    @pulumi.getter
    def branch(self) -> pulumi.Input[str]:
        """
        The Git branch to protect.
        """
        return pulumi.get(self, "branch")

    @branch.setter
    def branch(self, value: pulumi.Input[str]):
        pulumi.set(self, "branch", value)

    @property
    @pulumi.getter
    def repository(self) -> pulumi.Input[str]:
        """
        The GitHub repository name.
        """
        return pulumi.get(self, "repository")

    @repository.setter
    def repository(self, value: pulumi.Input[str]):
        pulumi.set(self, "repository", value)

    @property
    @pulumi.getter(name="enforceAdmins")
    def enforce_admins(self) -> Optional[pulumi.Input[bool]]:
        """
        Boolean, setting this to `true` enforces status checks for repository administrators.
        """
        return pulumi.get(self, "enforce_admins")

    @enforce_admins.setter
    def enforce_admins(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enforce_admins", value)

    @property
    @pulumi.getter(name="requireConversationResolution")
    def require_conversation_resolution(self) -> Optional[pulumi.Input[bool]]:
        """
        Boolean, setting this to `true` requires all conversations on code must be resolved before a pull request can be merged.
        """
        return pulumi.get(self, "require_conversation_resolution")

    @require_conversation_resolution.setter
    def require_conversation_resolution(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "require_conversation_resolution", value)

    @property
    @pulumi.getter(name="requireSignedCommits")
    def require_signed_commits(self) -> Optional[pulumi.Input[bool]]:
        """
        Boolean, setting this to `true` requires all commits to be signed with GPG.
        """
        return pulumi.get(self, "require_signed_commits")

    @require_signed_commits.setter
    def require_signed_commits(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "require_signed_commits", value)

    @property
    @pulumi.getter(name="requiredPullRequestReviews")
    def required_pull_request_reviews(self) -> Optional[pulumi.Input['BranchProtectionV3RequiredPullRequestReviewsArgs']]:
        """
        Enforce restrictions for pull request reviews. See Required Pull Request Reviews below for details.
        """
        return pulumi.get(self, "required_pull_request_reviews")

    @required_pull_request_reviews.setter
    def required_pull_request_reviews(self, value: Optional[pulumi.Input['BranchProtectionV3RequiredPullRequestReviewsArgs']]):
        pulumi.set(self, "required_pull_request_reviews", value)

    @property
    @pulumi.getter(name="requiredStatusChecks")
    def required_status_checks(self) -> Optional[pulumi.Input['BranchProtectionV3RequiredStatusChecksArgs']]:
        """
        Enforce restrictions for required status checks. See Required Status Checks below for details.
        """
        return pulumi.get(self, "required_status_checks")

    @required_status_checks.setter
    def required_status_checks(self, value: Optional[pulumi.Input['BranchProtectionV3RequiredStatusChecksArgs']]):
        pulumi.set(self, "required_status_checks", value)

    @property
    @pulumi.getter
    def restrictions(self) -> Optional[pulumi.Input['BranchProtectionV3RestrictionsArgs']]:
        """
        Enforce restrictions for the users and teams that may push to the branch. See Restrictions below for details.
        """
        return pulumi.get(self, "restrictions")

    @restrictions.setter
    def restrictions(self, value: Optional[pulumi.Input['BranchProtectionV3RestrictionsArgs']]):
        pulumi.set(self, "restrictions", value)


@pulumi.input_type
class _BranchProtectionV3State:
    def __init__(__self__, *,
                 branch: Optional[pulumi.Input[str]] = None,
                 enforce_admins: Optional[pulumi.Input[bool]] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 repository: Optional[pulumi.Input[str]] = None,
                 require_conversation_resolution: Optional[pulumi.Input[bool]] = None,
                 require_signed_commits: Optional[pulumi.Input[bool]] = None,
                 required_pull_request_reviews: Optional[pulumi.Input['BranchProtectionV3RequiredPullRequestReviewsArgs']] = None,
                 required_status_checks: Optional[pulumi.Input['BranchProtectionV3RequiredStatusChecksArgs']] = None,
                 restrictions: Optional[pulumi.Input['BranchProtectionV3RestrictionsArgs']] = None):
        """
        Input properties used for looking up and filtering BranchProtectionV3 resources.
        :param pulumi.Input[str] branch: The Git branch to protect.
        :param pulumi.Input[bool] enforce_admins: Boolean, setting this to `true` enforces status checks for repository administrators.
        :param pulumi.Input[str] repository: The GitHub repository name.
        :param pulumi.Input[bool] require_conversation_resolution: Boolean, setting this to `true` requires all conversations on code must be resolved before a pull request can be merged.
        :param pulumi.Input[bool] require_signed_commits: Boolean, setting this to `true` requires all commits to be signed with GPG.
        :param pulumi.Input['BranchProtectionV3RequiredPullRequestReviewsArgs'] required_pull_request_reviews: Enforce restrictions for pull request reviews. See Required Pull Request Reviews below for details.
        :param pulumi.Input['BranchProtectionV3RequiredStatusChecksArgs'] required_status_checks: Enforce restrictions for required status checks. See Required Status Checks below for details.
        :param pulumi.Input['BranchProtectionV3RestrictionsArgs'] restrictions: Enforce restrictions for the users and teams that may push to the branch. See Restrictions below for details.
        """
        if branch is not None:
            pulumi.set(__self__, "branch", branch)
        if enforce_admins is not None:
            pulumi.set(__self__, "enforce_admins", enforce_admins)
        if etag is not None:
            pulumi.set(__self__, "etag", etag)
        if repository is not None:
            pulumi.set(__self__, "repository", repository)
        if require_conversation_resolution is not None:
            pulumi.set(__self__, "require_conversation_resolution", require_conversation_resolution)
        if require_signed_commits is not None:
            pulumi.set(__self__, "require_signed_commits", require_signed_commits)
        if required_pull_request_reviews is not None:
            pulumi.set(__self__, "required_pull_request_reviews", required_pull_request_reviews)
        if required_status_checks is not None:
            pulumi.set(__self__, "required_status_checks", required_status_checks)
        if restrictions is not None:
            pulumi.set(__self__, "restrictions", restrictions)

    @property
    @pulumi.getter
    def branch(self) -> Optional[pulumi.Input[str]]:
        """
        The Git branch to protect.
        """
        return pulumi.get(self, "branch")

    @branch.setter
    def branch(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "branch", value)

    @property
    @pulumi.getter(name="enforceAdmins")
    def enforce_admins(self) -> Optional[pulumi.Input[bool]]:
        """
        Boolean, setting this to `true` enforces status checks for repository administrators.
        """
        return pulumi.get(self, "enforce_admins")

    @enforce_admins.setter
    def enforce_admins(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enforce_admins", value)

    @property
    @pulumi.getter
    def etag(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "etag")

    @etag.setter
    def etag(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "etag", value)

    @property
    @pulumi.getter
    def repository(self) -> Optional[pulumi.Input[str]]:
        """
        The GitHub repository name.
        """
        return pulumi.get(self, "repository")

    @repository.setter
    def repository(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "repository", value)

    @property
    @pulumi.getter(name="requireConversationResolution")
    def require_conversation_resolution(self) -> Optional[pulumi.Input[bool]]:
        """
        Boolean, setting this to `true` requires all conversations on code must be resolved before a pull request can be merged.
        """
        return pulumi.get(self, "require_conversation_resolution")

    @require_conversation_resolution.setter
    def require_conversation_resolution(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "require_conversation_resolution", value)

    @property
    @pulumi.getter(name="requireSignedCommits")
    def require_signed_commits(self) -> Optional[pulumi.Input[bool]]:
        """
        Boolean, setting this to `true` requires all commits to be signed with GPG.
        """
        return pulumi.get(self, "require_signed_commits")

    @require_signed_commits.setter
    def require_signed_commits(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "require_signed_commits", value)

    @property
    @pulumi.getter(name="requiredPullRequestReviews")
    def required_pull_request_reviews(self) -> Optional[pulumi.Input['BranchProtectionV3RequiredPullRequestReviewsArgs']]:
        """
        Enforce restrictions for pull request reviews. See Required Pull Request Reviews below for details.
        """
        return pulumi.get(self, "required_pull_request_reviews")

    @required_pull_request_reviews.setter
    def required_pull_request_reviews(self, value: Optional[pulumi.Input['BranchProtectionV3RequiredPullRequestReviewsArgs']]):
        pulumi.set(self, "required_pull_request_reviews", value)

    @property
    @pulumi.getter(name="requiredStatusChecks")
    def required_status_checks(self) -> Optional[pulumi.Input['BranchProtectionV3RequiredStatusChecksArgs']]:
        """
        Enforce restrictions for required status checks. See Required Status Checks below for details.
        """
        return pulumi.get(self, "required_status_checks")

    @required_status_checks.setter
    def required_status_checks(self, value: Optional[pulumi.Input['BranchProtectionV3RequiredStatusChecksArgs']]):
        pulumi.set(self, "required_status_checks", value)

    @property
    @pulumi.getter
    def restrictions(self) -> Optional[pulumi.Input['BranchProtectionV3RestrictionsArgs']]:
        """
        Enforce restrictions for the users and teams that may push to the branch. See Restrictions below for details.
        """
        return pulumi.get(self, "restrictions")

    @restrictions.setter
    def restrictions(self, value: Optional[pulumi.Input['BranchProtectionV3RestrictionsArgs']]):
        pulumi.set(self, "restrictions", value)


class BranchProtectionV3(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 branch: Optional[pulumi.Input[str]] = None,
                 enforce_admins: Optional[pulumi.Input[bool]] = None,
                 repository: Optional[pulumi.Input[str]] = None,
                 require_conversation_resolution: Optional[pulumi.Input[bool]] = None,
                 require_signed_commits: Optional[pulumi.Input[bool]] = None,
                 required_pull_request_reviews: Optional[pulumi.Input[pulumi.InputType['BranchProtectionV3RequiredPullRequestReviewsArgs']]] = None,
                 required_status_checks: Optional[pulumi.Input[pulumi.InputType['BranchProtectionV3RequiredStatusChecksArgs']]] = None,
                 restrictions: Optional[pulumi.Input[pulumi.InputType['BranchProtectionV3RestrictionsArgs']]] = None,
                 __props__=None):
        """
        Protects a GitHub branch.

        The `BranchProtection` resource has moved to the GraphQL API, while this resource will continue to leverage the REST API.

        This resource allows you to configure branch protection for repositories in your organization. When applied, the branch will be protected from forced pushes and deletion. Additional constraints, such as required status checks or restrictions on users, teams, and apps, can also be configured.

        ## Import

        GitHub Branch Protection can be imported using an ID made up of `repository:branch`, e.g.

        ```sh
         $ pulumi import github:index/branchProtectionV3:BranchProtectionV3 terraform terraform:main
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] branch: The Git branch to protect.
        :param pulumi.Input[bool] enforce_admins: Boolean, setting this to `true` enforces status checks for repository administrators.
        :param pulumi.Input[str] repository: The GitHub repository name.
        :param pulumi.Input[bool] require_conversation_resolution: Boolean, setting this to `true` requires all conversations on code must be resolved before a pull request can be merged.
        :param pulumi.Input[bool] require_signed_commits: Boolean, setting this to `true` requires all commits to be signed with GPG.
        :param pulumi.Input[pulumi.InputType['BranchProtectionV3RequiredPullRequestReviewsArgs']] required_pull_request_reviews: Enforce restrictions for pull request reviews. See Required Pull Request Reviews below for details.
        :param pulumi.Input[pulumi.InputType['BranchProtectionV3RequiredStatusChecksArgs']] required_status_checks: Enforce restrictions for required status checks. See Required Status Checks below for details.
        :param pulumi.Input[pulumi.InputType['BranchProtectionV3RestrictionsArgs']] restrictions: Enforce restrictions for the users and teams that may push to the branch. See Restrictions below for details.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: BranchProtectionV3Args,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Protects a GitHub branch.

        The `BranchProtection` resource has moved to the GraphQL API, while this resource will continue to leverage the REST API.

        This resource allows you to configure branch protection for repositories in your organization. When applied, the branch will be protected from forced pushes and deletion. Additional constraints, such as required status checks or restrictions on users, teams, and apps, can also be configured.

        ## Import

        GitHub Branch Protection can be imported using an ID made up of `repository:branch`, e.g.

        ```sh
         $ pulumi import github:index/branchProtectionV3:BranchProtectionV3 terraform terraform:main
        ```

        :param str resource_name: The name of the resource.
        :param BranchProtectionV3Args args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BranchProtectionV3Args, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 branch: Optional[pulumi.Input[str]] = None,
                 enforce_admins: Optional[pulumi.Input[bool]] = None,
                 repository: Optional[pulumi.Input[str]] = None,
                 require_conversation_resolution: Optional[pulumi.Input[bool]] = None,
                 require_signed_commits: Optional[pulumi.Input[bool]] = None,
                 required_pull_request_reviews: Optional[pulumi.Input[pulumi.InputType['BranchProtectionV3RequiredPullRequestReviewsArgs']]] = None,
                 required_status_checks: Optional[pulumi.Input[pulumi.InputType['BranchProtectionV3RequiredStatusChecksArgs']]] = None,
                 restrictions: Optional[pulumi.Input[pulumi.InputType['BranchProtectionV3RestrictionsArgs']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = BranchProtectionV3Args.__new__(BranchProtectionV3Args)

            if branch is None and not opts.urn:
                raise TypeError("Missing required property 'branch'")
            __props__.__dict__["branch"] = branch
            __props__.__dict__["enforce_admins"] = enforce_admins
            if repository is None and not opts.urn:
                raise TypeError("Missing required property 'repository'")
            __props__.__dict__["repository"] = repository
            __props__.__dict__["require_conversation_resolution"] = require_conversation_resolution
            __props__.__dict__["require_signed_commits"] = require_signed_commits
            __props__.__dict__["required_pull_request_reviews"] = required_pull_request_reviews
            __props__.__dict__["required_status_checks"] = required_status_checks
            __props__.__dict__["restrictions"] = restrictions
            __props__.__dict__["etag"] = None
        super(BranchProtectionV3, __self__).__init__(
            'github:index/branchProtectionV3:BranchProtectionV3',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            branch: Optional[pulumi.Input[str]] = None,
            enforce_admins: Optional[pulumi.Input[bool]] = None,
            etag: Optional[pulumi.Input[str]] = None,
            repository: Optional[pulumi.Input[str]] = None,
            require_conversation_resolution: Optional[pulumi.Input[bool]] = None,
            require_signed_commits: Optional[pulumi.Input[bool]] = None,
            required_pull_request_reviews: Optional[pulumi.Input[pulumi.InputType['BranchProtectionV3RequiredPullRequestReviewsArgs']]] = None,
            required_status_checks: Optional[pulumi.Input[pulumi.InputType['BranchProtectionV3RequiredStatusChecksArgs']]] = None,
            restrictions: Optional[pulumi.Input[pulumi.InputType['BranchProtectionV3RestrictionsArgs']]] = None) -> 'BranchProtectionV3':
        """
        Get an existing BranchProtectionV3 resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] branch: The Git branch to protect.
        :param pulumi.Input[bool] enforce_admins: Boolean, setting this to `true` enforces status checks for repository administrators.
        :param pulumi.Input[str] repository: The GitHub repository name.
        :param pulumi.Input[bool] require_conversation_resolution: Boolean, setting this to `true` requires all conversations on code must be resolved before a pull request can be merged.
        :param pulumi.Input[bool] require_signed_commits: Boolean, setting this to `true` requires all commits to be signed with GPG.
        :param pulumi.Input[pulumi.InputType['BranchProtectionV3RequiredPullRequestReviewsArgs']] required_pull_request_reviews: Enforce restrictions for pull request reviews. See Required Pull Request Reviews below for details.
        :param pulumi.Input[pulumi.InputType['BranchProtectionV3RequiredStatusChecksArgs']] required_status_checks: Enforce restrictions for required status checks. See Required Status Checks below for details.
        :param pulumi.Input[pulumi.InputType['BranchProtectionV3RestrictionsArgs']] restrictions: Enforce restrictions for the users and teams that may push to the branch. See Restrictions below for details.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _BranchProtectionV3State.__new__(_BranchProtectionV3State)

        __props__.__dict__["branch"] = branch
        __props__.__dict__["enforce_admins"] = enforce_admins
        __props__.__dict__["etag"] = etag
        __props__.__dict__["repository"] = repository
        __props__.__dict__["require_conversation_resolution"] = require_conversation_resolution
        __props__.__dict__["require_signed_commits"] = require_signed_commits
        __props__.__dict__["required_pull_request_reviews"] = required_pull_request_reviews
        __props__.__dict__["required_status_checks"] = required_status_checks
        __props__.__dict__["restrictions"] = restrictions
        return BranchProtectionV3(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def branch(self) -> pulumi.Output[str]:
        """
        The Git branch to protect.
        """
        return pulumi.get(self, "branch")

    @property
    @pulumi.getter(name="enforceAdmins")
    def enforce_admins(self) -> pulumi.Output[Optional[bool]]:
        """
        Boolean, setting this to `true` enforces status checks for repository administrators.
        """
        return pulumi.get(self, "enforce_admins")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def repository(self) -> pulumi.Output[str]:
        """
        The GitHub repository name.
        """
        return pulumi.get(self, "repository")

    @property
    @pulumi.getter(name="requireConversationResolution")
    def require_conversation_resolution(self) -> pulumi.Output[Optional[bool]]:
        """
        Boolean, setting this to `true` requires all conversations on code must be resolved before a pull request can be merged.
        """
        return pulumi.get(self, "require_conversation_resolution")

    @property
    @pulumi.getter(name="requireSignedCommits")
    def require_signed_commits(self) -> pulumi.Output[Optional[bool]]:
        """
        Boolean, setting this to `true` requires all commits to be signed with GPG.
        """
        return pulumi.get(self, "require_signed_commits")

    @property
    @pulumi.getter(name="requiredPullRequestReviews")
    def required_pull_request_reviews(self) -> pulumi.Output[Optional['outputs.BranchProtectionV3RequiredPullRequestReviews']]:
        """
        Enforce restrictions for pull request reviews. See Required Pull Request Reviews below for details.
        """
        return pulumi.get(self, "required_pull_request_reviews")

    @property
    @pulumi.getter(name="requiredStatusChecks")
    def required_status_checks(self) -> pulumi.Output[Optional['outputs.BranchProtectionV3RequiredStatusChecks']]:
        """
        Enforce restrictions for required status checks. See Required Status Checks below for details.
        """
        return pulumi.get(self, "required_status_checks")

    @property
    @pulumi.getter
    def restrictions(self) -> pulumi.Output[Optional['outputs.BranchProtectionV3Restrictions']]:
        """
        Enforce restrictions for the users and teams that may push to the branch. See Restrictions below for details.
        """
        return pulumi.get(self, "restrictions")

