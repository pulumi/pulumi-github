# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['UserInvitationAccepterArgs', 'UserInvitationAccepter']

@pulumi.input_type
class UserInvitationAccepterArgs:
    def __init__(__self__, *,
                 allow_empty_id: Optional[pulumi.Input[bool]] = None,
                 invitation_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a UserInvitationAccepter resource.
        :param pulumi.Input[bool] allow_empty_id: Allow the ID to be unset. This will result in the resource being skipped when the ID is not set instead of returning an error.
        :param pulumi.Input[str] invitation_id: ID of the invitation to accept. Must be set when `allow_empty_id` is `false`.
        """
        if allow_empty_id is not None:
            pulumi.set(__self__, "allow_empty_id", allow_empty_id)
        if invitation_id is not None:
            pulumi.set(__self__, "invitation_id", invitation_id)

    @property
    @pulumi.getter(name="allowEmptyId")
    def allow_empty_id(self) -> Optional[pulumi.Input[bool]]:
        """
        Allow the ID to be unset. This will result in the resource being skipped when the ID is not set instead of returning an error.
        """
        return pulumi.get(self, "allow_empty_id")

    @allow_empty_id.setter
    def allow_empty_id(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "allow_empty_id", value)

    @property
    @pulumi.getter(name="invitationId")
    def invitation_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the invitation to accept. Must be set when `allow_empty_id` is `false`.
        """
        return pulumi.get(self, "invitation_id")

    @invitation_id.setter
    def invitation_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "invitation_id", value)


@pulumi.input_type
class _UserInvitationAccepterState:
    def __init__(__self__, *,
                 allow_empty_id: Optional[pulumi.Input[bool]] = None,
                 invitation_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering UserInvitationAccepter resources.
        :param pulumi.Input[bool] allow_empty_id: Allow the ID to be unset. This will result in the resource being skipped when the ID is not set instead of returning an error.
        :param pulumi.Input[str] invitation_id: ID of the invitation to accept. Must be set when `allow_empty_id` is `false`.
        """
        if allow_empty_id is not None:
            pulumi.set(__self__, "allow_empty_id", allow_empty_id)
        if invitation_id is not None:
            pulumi.set(__self__, "invitation_id", invitation_id)

    @property
    @pulumi.getter(name="allowEmptyId")
    def allow_empty_id(self) -> Optional[pulumi.Input[bool]]:
        """
        Allow the ID to be unset. This will result in the resource being skipped when the ID is not set instead of returning an error.
        """
        return pulumi.get(self, "allow_empty_id")

    @allow_empty_id.setter
    def allow_empty_id(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "allow_empty_id", value)

    @property
    @pulumi.getter(name="invitationId")
    def invitation_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the invitation to accept. Must be set when `allow_empty_id` is `false`.
        """
        return pulumi.get(self, "invitation_id")

    @invitation_id.setter
    def invitation_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "invitation_id", value)


class UserInvitationAccepter(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allow_empty_id: Optional[pulumi.Input[bool]] = None,
                 invitation_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to manage GitHub repository collaborator invitations.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_github as github

        example_repository = github.Repository("exampleRepository")
        example_repository_collaborator = github.RepositoryCollaborator("exampleRepositoryCollaborator",
            repository=example_repository.name,
            username="example-username",
            permission="push")
        invitee = github.Provider("invitee", token=var["invitee_token"])
        example_user_invitation_accepter = github.UserInvitationAccepter("exampleUserInvitationAccepter", invitation_id=example_repository_collaborator.invitation_id,
        opts=pulumi.ResourceOptions(provider="github.invitee"))
        ```
        ## Allowing empty invitation IDs

        Set `allow_empty_id` when using `for_each` over a list of `github_repository_collaborator.invitation_id`'s.

        This allows applying a module again when a new `RepositoryCollaborator` resource is added to the `for_each` loop.
        This is needed as the `github_repository_collaborator.invitation_id` will be empty after a state refresh when the invitation has been accepted.

        Note that when an invitation is accepted manually or by another tool between a state refresh and a `pulumi up` using that refreshed state,
        the plan will contain the invitation ID, but the apply will receive an HTTP 404 from the API since the invitation has already been accepted.

        This is tracked in #1157.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] allow_empty_id: Allow the ID to be unset. This will result in the resource being skipped when the ID is not set instead of returning an error.
        :param pulumi.Input[str] invitation_id: ID of the invitation to accept. Must be set when `allow_empty_id` is `false`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[UserInvitationAccepterArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to manage GitHub repository collaborator invitations.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_github as github

        example_repository = github.Repository("exampleRepository")
        example_repository_collaborator = github.RepositoryCollaborator("exampleRepositoryCollaborator",
            repository=example_repository.name,
            username="example-username",
            permission="push")
        invitee = github.Provider("invitee", token=var["invitee_token"])
        example_user_invitation_accepter = github.UserInvitationAccepter("exampleUserInvitationAccepter", invitation_id=example_repository_collaborator.invitation_id,
        opts=pulumi.ResourceOptions(provider="github.invitee"))
        ```
        ## Allowing empty invitation IDs

        Set `allow_empty_id` when using `for_each` over a list of `github_repository_collaborator.invitation_id`'s.

        This allows applying a module again when a new `RepositoryCollaborator` resource is added to the `for_each` loop.
        This is needed as the `github_repository_collaborator.invitation_id` will be empty after a state refresh when the invitation has been accepted.

        Note that when an invitation is accepted manually or by another tool between a state refresh and a `pulumi up` using that refreshed state,
        the plan will contain the invitation ID, but the apply will receive an HTTP 404 from the API since the invitation has already been accepted.

        This is tracked in #1157.

        :param str resource_name: The name of the resource.
        :param UserInvitationAccepterArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(UserInvitationAccepterArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allow_empty_id: Optional[pulumi.Input[bool]] = None,
                 invitation_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = UserInvitationAccepterArgs.__new__(UserInvitationAccepterArgs)

            __props__.__dict__["allow_empty_id"] = allow_empty_id
            __props__.__dict__["invitation_id"] = invitation_id
        super(UserInvitationAccepter, __self__).__init__(
            'github:index/userInvitationAccepter:UserInvitationAccepter',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            allow_empty_id: Optional[pulumi.Input[bool]] = None,
            invitation_id: Optional[pulumi.Input[str]] = None) -> 'UserInvitationAccepter':
        """
        Get an existing UserInvitationAccepter resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] allow_empty_id: Allow the ID to be unset. This will result in the resource being skipped when the ID is not set instead of returning an error.
        :param pulumi.Input[str] invitation_id: ID of the invitation to accept. Must be set when `allow_empty_id` is `false`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _UserInvitationAccepterState.__new__(_UserInvitationAccepterState)

        __props__.__dict__["allow_empty_id"] = allow_empty_id
        __props__.__dict__["invitation_id"] = invitation_id
        return UserInvitationAccepter(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="allowEmptyId")
    def allow_empty_id(self) -> pulumi.Output[Optional[bool]]:
        """
        Allow the ID to be unset. This will result in the resource being skipped when the ID is not set instead of returning an error.
        """
        return pulumi.get(self, "allow_empty_id")

    @property
    @pulumi.getter(name="invitationId")
    def invitation_id(self) -> pulumi.Output[Optional[str]]:
        """
        ID of the invitation to accept. Must be set when `allow_empty_id` is `false`.
        """
        return pulumi.get(self, "invitation_id")

