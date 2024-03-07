# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['RepositoryTagProtectionArgs', 'RepositoryTagProtection']

@pulumi.input_type
class RepositoryTagProtectionArgs:
    def __init__(__self__, *,
                 pattern: pulumi.Input[str],
                 repository: pulumi.Input[str]):
        """
        The set of arguments for constructing a RepositoryTagProtection resource.
        :param pulumi.Input[str] pattern: The pattern of the tag to protect.
        :param pulumi.Input[str] repository: Name of the repository to add the tag protection to.
        """
        pulumi.set(__self__, "pattern", pattern)
        pulumi.set(__self__, "repository", repository)

    @property
    @pulumi.getter
    def pattern(self) -> pulumi.Input[str]:
        """
        The pattern of the tag to protect.
        """
        return pulumi.get(self, "pattern")

    @pattern.setter
    def pattern(self, value: pulumi.Input[str]):
        pulumi.set(self, "pattern", value)

    @property
    @pulumi.getter
    def repository(self) -> pulumi.Input[str]:
        """
        Name of the repository to add the tag protection to.
        """
        return pulumi.get(self, "repository")

    @repository.setter
    def repository(self, value: pulumi.Input[str]):
        pulumi.set(self, "repository", value)


@pulumi.input_type
class _RepositoryTagProtectionState:
    def __init__(__self__, *,
                 pattern: Optional[pulumi.Input[str]] = None,
                 repository: Optional[pulumi.Input[str]] = None,
                 tag_protection_id: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering RepositoryTagProtection resources.
        :param pulumi.Input[str] pattern: The pattern of the tag to protect.
        :param pulumi.Input[str] repository: Name of the repository to add the tag protection to.
        :param pulumi.Input[int] tag_protection_id: The ID of the tag protection.
        """
        if pattern is not None:
            pulumi.set(__self__, "pattern", pattern)
        if repository is not None:
            pulumi.set(__self__, "repository", repository)
        if tag_protection_id is not None:
            pulumi.set(__self__, "tag_protection_id", tag_protection_id)

    @property
    @pulumi.getter
    def pattern(self) -> Optional[pulumi.Input[str]]:
        """
        The pattern of the tag to protect.
        """
        return pulumi.get(self, "pattern")

    @pattern.setter
    def pattern(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "pattern", value)

    @property
    @pulumi.getter
    def repository(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the repository to add the tag protection to.
        """
        return pulumi.get(self, "repository")

    @repository.setter
    def repository(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "repository", value)

    @property
    @pulumi.getter(name="tagProtectionId")
    def tag_protection_id(self) -> Optional[pulumi.Input[int]]:
        """
        The ID of the tag protection.
        """
        return pulumi.get(self, "tag_protection_id")

    @tag_protection_id.setter
    def tag_protection_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "tag_protection_id", value)


class RepositoryTagProtection(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 pattern: Optional[pulumi.Input[str]] = None,
                 repository: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        This resource allows you to create and manage a repository tag protection for repositories within your GitHub organization or personal account.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_github as github

        example = github.RepositoryTagProtection("example",
            pattern="v*",
            repository="example-repository")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Repository tag protections can be imported using the `name` of the repository, combined with the `id` of the tag protection, separated by a `/` character.
        The `id` of the tag protection can be found using the [GitHub API](https://docs.github.com/en/rest/repos/tags#list-tag-protection-states-for-a-repository).

        Importing uses the name of the repository, as well as the ID of the tag protection, e.g.

        ```sh
        $ pulumi import github:index/repositoryTagProtection:RepositoryTagProtection terraform my-repo/31077
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] pattern: The pattern of the tag to protect.
        :param pulumi.Input[str] repository: Name of the repository to add the tag protection to.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RepositoryTagProtectionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This resource allows you to create and manage a repository tag protection for repositories within your GitHub organization or personal account.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_github as github

        example = github.RepositoryTagProtection("example",
            pattern="v*",
            repository="example-repository")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Repository tag protections can be imported using the `name` of the repository, combined with the `id` of the tag protection, separated by a `/` character.
        The `id` of the tag protection can be found using the [GitHub API](https://docs.github.com/en/rest/repos/tags#list-tag-protection-states-for-a-repository).

        Importing uses the name of the repository, as well as the ID of the tag protection, e.g.

        ```sh
        $ pulumi import github:index/repositoryTagProtection:RepositoryTagProtection terraform my-repo/31077
        ```

        :param str resource_name: The name of the resource.
        :param RepositoryTagProtectionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RepositoryTagProtectionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 pattern: Optional[pulumi.Input[str]] = None,
                 repository: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RepositoryTagProtectionArgs.__new__(RepositoryTagProtectionArgs)

            if pattern is None and not opts.urn:
                raise TypeError("Missing required property 'pattern'")
            __props__.__dict__["pattern"] = pattern
            if repository is None and not opts.urn:
                raise TypeError("Missing required property 'repository'")
            __props__.__dict__["repository"] = repository
            __props__.__dict__["tag_protection_id"] = None
        super(RepositoryTagProtection, __self__).__init__(
            'github:index/repositoryTagProtection:RepositoryTagProtection',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            pattern: Optional[pulumi.Input[str]] = None,
            repository: Optional[pulumi.Input[str]] = None,
            tag_protection_id: Optional[pulumi.Input[int]] = None) -> 'RepositoryTagProtection':
        """
        Get an existing RepositoryTagProtection resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] pattern: The pattern of the tag to protect.
        :param pulumi.Input[str] repository: Name of the repository to add the tag protection to.
        :param pulumi.Input[int] tag_protection_id: The ID of the tag protection.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RepositoryTagProtectionState.__new__(_RepositoryTagProtectionState)

        __props__.__dict__["pattern"] = pattern
        __props__.__dict__["repository"] = repository
        __props__.__dict__["tag_protection_id"] = tag_protection_id
        return RepositoryTagProtection(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def pattern(self) -> pulumi.Output[str]:
        """
        The pattern of the tag to protect.
        """
        return pulumi.get(self, "pattern")

    @property
    @pulumi.getter
    def repository(self) -> pulumi.Output[str]:
        """
        Name of the repository to add the tag protection to.
        """
        return pulumi.get(self, "repository")

    @property
    @pulumi.getter(name="tagProtectionId")
    def tag_protection_id(self) -> pulumi.Output[int]:
        """
        The ID of the tag protection.
        """
        return pulumi.get(self, "tag_protection_id")

