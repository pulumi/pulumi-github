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

__all__ = ['RepositoryAutolinkReferenceArgs', 'RepositoryAutolinkReference']

@pulumi.input_type
class RepositoryAutolinkReferenceArgs:
    def __init__(__self__, *,
                 key_prefix: pulumi.Input[str],
                 repository: pulumi.Input[str],
                 target_url_template: pulumi.Input[str],
                 is_alphanumeric: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a RepositoryAutolinkReference resource.
        :param pulumi.Input[str] key_prefix: This prefix appended by a number will generate a link any time it is found in an issue, pull request, or commit.
        :param pulumi.Input[str] repository: The repository of the autolink reference.
        :param pulumi.Input[str] target_url_template: The template of the target URL used for the links; must be a valid URL and contain `<num>` for the reference number
        :param pulumi.Input[bool] is_alphanumeric: Whether this autolink reference matches alphanumeric characters. If false, this autolink reference only matches numeric characters. Default is true.
        """
        pulumi.set(__self__, "key_prefix", key_prefix)
        pulumi.set(__self__, "repository", repository)
        pulumi.set(__self__, "target_url_template", target_url_template)
        if is_alphanumeric is not None:
            pulumi.set(__self__, "is_alphanumeric", is_alphanumeric)

    @property
    @pulumi.getter(name="keyPrefix")
    def key_prefix(self) -> pulumi.Input[str]:
        """
        This prefix appended by a number will generate a link any time it is found in an issue, pull request, or commit.
        """
        return pulumi.get(self, "key_prefix")

    @key_prefix.setter
    def key_prefix(self, value: pulumi.Input[str]):
        pulumi.set(self, "key_prefix", value)

    @property
    @pulumi.getter
    def repository(self) -> pulumi.Input[str]:
        """
        The repository of the autolink reference.
        """
        return pulumi.get(self, "repository")

    @repository.setter
    def repository(self, value: pulumi.Input[str]):
        pulumi.set(self, "repository", value)

    @property
    @pulumi.getter(name="targetUrlTemplate")
    def target_url_template(self) -> pulumi.Input[str]:
        """
        The template of the target URL used for the links; must be a valid URL and contain `<num>` for the reference number
        """
        return pulumi.get(self, "target_url_template")

    @target_url_template.setter
    def target_url_template(self, value: pulumi.Input[str]):
        pulumi.set(self, "target_url_template", value)

    @property
    @pulumi.getter(name="isAlphanumeric")
    def is_alphanumeric(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether this autolink reference matches alphanumeric characters. If false, this autolink reference only matches numeric characters. Default is true.
        """
        return pulumi.get(self, "is_alphanumeric")

    @is_alphanumeric.setter
    def is_alphanumeric(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_alphanumeric", value)


@pulumi.input_type
class _RepositoryAutolinkReferenceState:
    def __init__(__self__, *,
                 etag: Optional[pulumi.Input[str]] = None,
                 is_alphanumeric: Optional[pulumi.Input[bool]] = None,
                 key_prefix: Optional[pulumi.Input[str]] = None,
                 repository: Optional[pulumi.Input[str]] = None,
                 target_url_template: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering RepositoryAutolinkReference resources.
        :param pulumi.Input[str] etag: An etag representing the autolink reference object.
        :param pulumi.Input[bool] is_alphanumeric: Whether this autolink reference matches alphanumeric characters. If false, this autolink reference only matches numeric characters. Default is true.
        :param pulumi.Input[str] key_prefix: This prefix appended by a number will generate a link any time it is found in an issue, pull request, or commit.
        :param pulumi.Input[str] repository: The repository of the autolink reference.
        :param pulumi.Input[str] target_url_template: The template of the target URL used for the links; must be a valid URL and contain `<num>` for the reference number
        """
        if etag is not None:
            pulumi.set(__self__, "etag", etag)
        if is_alphanumeric is not None:
            pulumi.set(__self__, "is_alphanumeric", is_alphanumeric)
        if key_prefix is not None:
            pulumi.set(__self__, "key_prefix", key_prefix)
        if repository is not None:
            pulumi.set(__self__, "repository", repository)
        if target_url_template is not None:
            pulumi.set(__self__, "target_url_template", target_url_template)

    @property
    @pulumi.getter
    def etag(self) -> Optional[pulumi.Input[str]]:
        """
        An etag representing the autolink reference object.
        """
        return pulumi.get(self, "etag")

    @etag.setter
    def etag(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "etag", value)

    @property
    @pulumi.getter(name="isAlphanumeric")
    def is_alphanumeric(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether this autolink reference matches alphanumeric characters. If false, this autolink reference only matches numeric characters. Default is true.
        """
        return pulumi.get(self, "is_alphanumeric")

    @is_alphanumeric.setter
    def is_alphanumeric(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_alphanumeric", value)

    @property
    @pulumi.getter(name="keyPrefix")
    def key_prefix(self) -> Optional[pulumi.Input[str]]:
        """
        This prefix appended by a number will generate a link any time it is found in an issue, pull request, or commit.
        """
        return pulumi.get(self, "key_prefix")

    @key_prefix.setter
    def key_prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_prefix", value)

    @property
    @pulumi.getter
    def repository(self) -> Optional[pulumi.Input[str]]:
        """
        The repository of the autolink reference.
        """
        return pulumi.get(self, "repository")

    @repository.setter
    def repository(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "repository", value)

    @property
    @pulumi.getter(name="targetUrlTemplate")
    def target_url_template(self) -> Optional[pulumi.Input[str]]:
        """
        The template of the target URL used for the links; must be a valid URL and contain `<num>` for the reference number
        """
        return pulumi.get(self, "target_url_template")

    @target_url_template.setter
    def target_url_template(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "target_url_template", value)


class RepositoryAutolinkReference(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 is_alphanumeric: Optional[pulumi.Input[bool]] = None,
                 key_prefix: Optional[pulumi.Input[str]] = None,
                 repository: Optional[pulumi.Input[str]] = None,
                 target_url_template: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        This resource allows you to create and manage an autolink reference for a single repository.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_github as github

        repo = github.Repository("repo",
            name="my-repo",
            description="GitHub repo managed by Terraform",
            private=False)
        autolink = github.RepositoryAutolinkReference("autolink",
            repository=repo.name,
            key_prefix="TICKET-",
            target_url_template="https://example.com/TICKET?query=<num>")
        ```

        ## Import

        ### Import by key prefix

        ```sh
        $ pulumi import github:index/repositoryAutolinkReference:RepositoryAutolinkReference auto oof/OOF-
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] is_alphanumeric: Whether this autolink reference matches alphanumeric characters. If false, this autolink reference only matches numeric characters. Default is true.
        :param pulumi.Input[str] key_prefix: This prefix appended by a number will generate a link any time it is found in an issue, pull request, or commit.
        :param pulumi.Input[str] repository: The repository of the autolink reference.
        :param pulumi.Input[str] target_url_template: The template of the target URL used for the links; must be a valid URL and contain `<num>` for the reference number
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RepositoryAutolinkReferenceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This resource allows you to create and manage an autolink reference for a single repository.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_github as github

        repo = github.Repository("repo",
            name="my-repo",
            description="GitHub repo managed by Terraform",
            private=False)
        autolink = github.RepositoryAutolinkReference("autolink",
            repository=repo.name,
            key_prefix="TICKET-",
            target_url_template="https://example.com/TICKET?query=<num>")
        ```

        ## Import

        ### Import by key prefix

        ```sh
        $ pulumi import github:index/repositoryAutolinkReference:RepositoryAutolinkReference auto oof/OOF-
        ```

        :param str resource_name: The name of the resource.
        :param RepositoryAutolinkReferenceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RepositoryAutolinkReferenceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 is_alphanumeric: Optional[pulumi.Input[bool]] = None,
                 key_prefix: Optional[pulumi.Input[str]] = None,
                 repository: Optional[pulumi.Input[str]] = None,
                 target_url_template: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RepositoryAutolinkReferenceArgs.__new__(RepositoryAutolinkReferenceArgs)

            __props__.__dict__["is_alphanumeric"] = is_alphanumeric
            if key_prefix is None and not opts.urn:
                raise TypeError("Missing required property 'key_prefix'")
            __props__.__dict__["key_prefix"] = key_prefix
            if repository is None and not opts.urn:
                raise TypeError("Missing required property 'repository'")
            __props__.__dict__["repository"] = repository
            if target_url_template is None and not opts.urn:
                raise TypeError("Missing required property 'target_url_template'")
            __props__.__dict__["target_url_template"] = target_url_template
            __props__.__dict__["etag"] = None
        super(RepositoryAutolinkReference, __self__).__init__(
            'github:index/repositoryAutolinkReference:RepositoryAutolinkReference',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            etag: Optional[pulumi.Input[str]] = None,
            is_alphanumeric: Optional[pulumi.Input[bool]] = None,
            key_prefix: Optional[pulumi.Input[str]] = None,
            repository: Optional[pulumi.Input[str]] = None,
            target_url_template: Optional[pulumi.Input[str]] = None) -> 'RepositoryAutolinkReference':
        """
        Get an existing RepositoryAutolinkReference resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] etag: An etag representing the autolink reference object.
        :param pulumi.Input[bool] is_alphanumeric: Whether this autolink reference matches alphanumeric characters. If false, this autolink reference only matches numeric characters. Default is true.
        :param pulumi.Input[str] key_prefix: This prefix appended by a number will generate a link any time it is found in an issue, pull request, or commit.
        :param pulumi.Input[str] repository: The repository of the autolink reference.
        :param pulumi.Input[str] target_url_template: The template of the target URL used for the links; must be a valid URL and contain `<num>` for the reference number
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RepositoryAutolinkReferenceState.__new__(_RepositoryAutolinkReferenceState)

        __props__.__dict__["etag"] = etag
        __props__.__dict__["is_alphanumeric"] = is_alphanumeric
        __props__.__dict__["key_prefix"] = key_prefix
        __props__.__dict__["repository"] = repository
        __props__.__dict__["target_url_template"] = target_url_template
        return RepositoryAutolinkReference(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        """
        An etag representing the autolink reference object.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="isAlphanumeric")
    def is_alphanumeric(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether this autolink reference matches alphanumeric characters. If false, this autolink reference only matches numeric characters. Default is true.
        """
        return pulumi.get(self, "is_alphanumeric")

    @property
    @pulumi.getter(name="keyPrefix")
    def key_prefix(self) -> pulumi.Output[str]:
        """
        This prefix appended by a number will generate a link any time it is found in an issue, pull request, or commit.
        """
        return pulumi.get(self, "key_prefix")

    @property
    @pulumi.getter
    def repository(self) -> pulumi.Output[str]:
        """
        The repository of the autolink reference.
        """
        return pulumi.get(self, "repository")

    @property
    @pulumi.getter(name="targetUrlTemplate")
    def target_url_template(self) -> pulumi.Output[str]:
        """
        The template of the target URL used for the links; must be a valid URL and contain `<num>` for the reference number
        """
        return pulumi.get(self, "target_url_template")

