# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities, _tables

__all__ = ['RepositoryCollaboratorArgs', 'RepositoryCollaborator']

@pulumi.input_type
class RepositoryCollaboratorArgs:
    def __init__(__self__, *,
                 repository: pulumi.Input[str],
                 username: pulumi.Input[str],
                 permission: Optional[pulumi.Input[str]] = None,
                 permission_diff_suppression: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a RepositoryCollaborator resource.
        :param pulumi.Input[str] repository: The GitHub repository
        :param pulumi.Input[str] username: The user to add to the repository as a collaborator.
        :param pulumi.Input[str] permission: The permission of the outside collaborator for the repository.
               Must be one of `pull`, `push`, `maintain`, `triage` or `admin` for organization-owned repositories.
               Must be `push` for personal repositories. Defaults to `push`.
        :param pulumi.Input[bool] permission_diff_suppression: Suppress plan diffs for `triage` and `maintain`.  Defaults to `false`.
        """
        pulumi.set(__self__, "repository", repository)
        pulumi.set(__self__, "username", username)
        if permission is not None:
            pulumi.set(__self__, "permission", permission)
        if permission_diff_suppression is not None:
            pulumi.set(__self__, "permission_diff_suppression", permission_diff_suppression)

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
    def username(self) -> pulumi.Input[str]:
        """
        The user to add to the repository as a collaborator.
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: pulumi.Input[str]):
        pulumi.set(self, "username", value)

    @property
    @pulumi.getter
    def permission(self) -> Optional[pulumi.Input[str]]:
        """
        The permission of the outside collaborator for the repository.
        Must be one of `pull`, `push`, `maintain`, `triage` or `admin` for organization-owned repositories.
        Must be `push` for personal repositories. Defaults to `push`.
        """
        return pulumi.get(self, "permission")

    @permission.setter
    def permission(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "permission", value)

    @property
    @pulumi.getter(name="permissionDiffSuppression")
    def permission_diff_suppression(self) -> Optional[pulumi.Input[bool]]:
        """
        Suppress plan diffs for `triage` and `maintain`.  Defaults to `false`.
        """
        return pulumi.get(self, "permission_diff_suppression")

    @permission_diff_suppression.setter
    def permission_diff_suppression(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "permission_diff_suppression", value)


class RepositoryCollaborator(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 permission: Optional[pulumi.Input[str]] = None,
                 permission_diff_suppression: Optional[pulumi.Input[bool]] = None,
                 repository: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a GitHub repository collaborator resource.

        This resource allows you to add/remove collaborators from repositories in your
        organization or personal account. For organization repositories, collaborators can
        have explicit (and differing levels of) read, write, or administrator access to
        specific repositories, without giving the user full organization membership.
        For personal repositories, collaborators can only be granted write
        (implictly includes read) permission.

        When applied, an invitation will be sent to the user to become a collaborator
        on a repository. When destroyed, either the invitation will be cancelled or the
        collaborator will be removed from the repository.

        Further documentation on GitHub collaborators:

        - [Adding outside collaborators to your personal repositories](https://help.github.com/en/github/setting-up-and-managing-your-github-user-account/managing-access-to-your-personal-repositories)
        - [Adding outside collaborators to repositories in your organization](https://help.github.com/articles/adding-outside-collaborators-to-repositories-in-your-organization/)
        - [Converting an organization member to an outside collaborator](https://help.github.com/articles/converting-an-organization-member-to-an-outside-collaborator/)

        ## Example Usage

        ```python
        import pulumi
        import pulumi_github as github

        # Add a collaborator to a repository
        a_repo_collaborator = github.RepositoryCollaborator("aRepoCollaborator",
            permission="admin",
            repository="our-cool-repo",
            username="SomeUser")
        ```

        ## Import

        GitHub Repository Collaborators can be imported using an ID made up of `repository:username`, e.g.

        ```sh
         $ pulumi import github:index/repositoryCollaborator:RepositoryCollaborator collaborator terraform:someuser
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] permission: The permission of the outside collaborator for the repository.
               Must be one of `pull`, `push`, `maintain`, `triage` or `admin` for organization-owned repositories.
               Must be `push` for personal repositories. Defaults to `push`.
        :param pulumi.Input[bool] permission_diff_suppression: Suppress plan diffs for `triage` and `maintain`.  Defaults to `false`.
        :param pulumi.Input[str] repository: The GitHub repository
        :param pulumi.Input[str] username: The user to add to the repository as a collaborator.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RepositoryCollaboratorArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a GitHub repository collaborator resource.

        This resource allows you to add/remove collaborators from repositories in your
        organization or personal account. For organization repositories, collaborators can
        have explicit (and differing levels of) read, write, or administrator access to
        specific repositories, without giving the user full organization membership.
        For personal repositories, collaborators can only be granted write
        (implictly includes read) permission.

        When applied, an invitation will be sent to the user to become a collaborator
        on a repository. When destroyed, either the invitation will be cancelled or the
        collaborator will be removed from the repository.

        Further documentation on GitHub collaborators:

        - [Adding outside collaborators to your personal repositories](https://help.github.com/en/github/setting-up-and-managing-your-github-user-account/managing-access-to-your-personal-repositories)
        - [Adding outside collaborators to repositories in your organization](https://help.github.com/articles/adding-outside-collaborators-to-repositories-in-your-organization/)
        - [Converting an organization member to an outside collaborator](https://help.github.com/articles/converting-an-organization-member-to-an-outside-collaborator/)

        ## Example Usage

        ```python
        import pulumi
        import pulumi_github as github

        # Add a collaborator to a repository
        a_repo_collaborator = github.RepositoryCollaborator("aRepoCollaborator",
            permission="admin",
            repository="our-cool-repo",
            username="SomeUser")
        ```

        ## Import

        GitHub Repository Collaborators can be imported using an ID made up of `repository:username`, e.g.

        ```sh
         $ pulumi import github:index/repositoryCollaborator:RepositoryCollaborator collaborator terraform:someuser
        ```

        :param str resource_name: The name of the resource.
        :param RepositoryCollaboratorArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RepositoryCollaboratorArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 permission: Optional[pulumi.Input[str]] = None,
                 permission_diff_suppression: Optional[pulumi.Input[bool]] = None,
                 repository: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['permission'] = permission
            __props__['permission_diff_suppression'] = permission_diff_suppression
            if repository is None and not opts.urn:
                raise TypeError("Missing required property 'repository'")
            __props__['repository'] = repository
            if username is None and not opts.urn:
                raise TypeError("Missing required property 'username'")
            __props__['username'] = username
            __props__['invitation_id'] = None
        super(RepositoryCollaborator, __self__).__init__(
            'github:index/repositoryCollaborator:RepositoryCollaborator',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            invitation_id: Optional[pulumi.Input[str]] = None,
            permission: Optional[pulumi.Input[str]] = None,
            permission_diff_suppression: Optional[pulumi.Input[bool]] = None,
            repository: Optional[pulumi.Input[str]] = None,
            username: Optional[pulumi.Input[str]] = None) -> 'RepositoryCollaborator':
        """
        Get an existing RepositoryCollaborator resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] invitation_id: ID of the invitation to be used in `UserInvitationAccepter`
        :param pulumi.Input[str] permission: The permission of the outside collaborator for the repository.
               Must be one of `pull`, `push`, `maintain`, `triage` or `admin` for organization-owned repositories.
               Must be `push` for personal repositories. Defaults to `push`.
        :param pulumi.Input[bool] permission_diff_suppression: Suppress plan diffs for `triage` and `maintain`.  Defaults to `false`.
        :param pulumi.Input[str] repository: The GitHub repository
        :param pulumi.Input[str] username: The user to add to the repository as a collaborator.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["invitation_id"] = invitation_id
        __props__["permission"] = permission
        __props__["permission_diff_suppression"] = permission_diff_suppression
        __props__["repository"] = repository
        __props__["username"] = username
        return RepositoryCollaborator(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="invitationId")
    def invitation_id(self) -> pulumi.Output[str]:
        """
        ID of the invitation to be used in `UserInvitationAccepter`
        """
        return pulumi.get(self, "invitation_id")

    @property
    @pulumi.getter
    def permission(self) -> pulumi.Output[Optional[str]]:
        """
        The permission of the outside collaborator for the repository.
        Must be one of `pull`, `push`, `maintain`, `triage` or `admin` for organization-owned repositories.
        Must be `push` for personal repositories. Defaults to `push`.
        """
        return pulumi.get(self, "permission")

    @property
    @pulumi.getter(name="permissionDiffSuppression")
    def permission_diff_suppression(self) -> pulumi.Output[Optional[bool]]:
        """
        Suppress plan diffs for `triage` and `maintain`.  Defaults to `false`.
        """
        return pulumi.get(self, "permission_diff_suppression")

    @property
    @pulumi.getter
    def repository(self) -> pulumi.Output[str]:
        """
        The GitHub repository
        """
        return pulumi.get(self, "repository")

    @property
    @pulumi.getter
    def username(self) -> pulumi.Output[str]:
        """
        The user to add to the repository as a collaborator.
        """
        return pulumi.get(self, "username")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

