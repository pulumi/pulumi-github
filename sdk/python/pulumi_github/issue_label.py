# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['IssueLabelArgs', 'IssueLabel']

@pulumi.input_type
class IssueLabelArgs:
    def __init__(__self__, *,
                 color: pulumi.Input[str],
                 repository: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a IssueLabel resource.
        :param pulumi.Input[str] color: A 6 character hex code, **without the leading #**, identifying the color of the label.
        :param pulumi.Input[str] repository: The GitHub repository
        :param pulumi.Input[str] description: A short description of the label.
        :param pulumi.Input[str] name: The name of the label.
        """
        pulumi.set(__self__, "color", color)
        pulumi.set(__self__, "repository", repository)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def color(self) -> pulumi.Input[str]:
        """
        A 6 character hex code, **without the leading #**, identifying the color of the label.
        """
        return pulumi.get(self, "color")

    @color.setter
    def color(self, value: pulumi.Input[str]):
        pulumi.set(self, "color", value)

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
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A short description of the label.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the label.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _IssueLabelState:
    def __init__(__self__, *,
                 color: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 repository: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering IssueLabel resources.
        :param pulumi.Input[str] color: A 6 character hex code, **without the leading #**, identifying the color of the label.
        :param pulumi.Input[str] description: A short description of the label.
        :param pulumi.Input[str] name: The name of the label.
        :param pulumi.Input[str] repository: The GitHub repository
        :param pulumi.Input[str] url: The URL to the issue label
        """
        if color is not None:
            pulumi.set(__self__, "color", color)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if etag is not None:
            pulumi.set(__self__, "etag", etag)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if repository is not None:
            pulumi.set(__self__, "repository", repository)
        if url is not None:
            pulumi.set(__self__, "url", url)

    @property
    @pulumi.getter
    def color(self) -> Optional[pulumi.Input[str]]:
        """
        A 6 character hex code, **without the leading #**, identifying the color of the label.
        """
        return pulumi.get(self, "color")

    @color.setter
    def color(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "color", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A short description of the label.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def etag(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "etag")

    @etag.setter
    def etag(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "etag", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the label.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

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

    @property
    @pulumi.getter
    def url(self) -> Optional[pulumi.Input[str]]:
        """
        The URL to the issue label
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url", value)


class IssueLabel(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 color: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 repository: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_github as github

        # Create a new, red colored label
        test_repo = github.IssueLabel("testRepo",
            color="FF0000",
            repository="test-repo")
        ```

        ## Import

        GitHub Issue Labels can be imported using an ID made up of `repository:name`, e.g.

        ```sh
         $ pulumi import github:index/issueLabel:IssueLabel panic_label terraform:panic
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] color: A 6 character hex code, **without the leading #**, identifying the color of the label.
        :param pulumi.Input[str] description: A short description of the label.
        :param pulumi.Input[str] name: The name of the label.
        :param pulumi.Input[str] repository: The GitHub repository
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: IssueLabelArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_github as github

        # Create a new, red colored label
        test_repo = github.IssueLabel("testRepo",
            color="FF0000",
            repository="test-repo")
        ```

        ## Import

        GitHub Issue Labels can be imported using an ID made up of `repository:name`, e.g.

        ```sh
         $ pulumi import github:index/issueLabel:IssueLabel panic_label terraform:panic
        ```

        :param str resource_name: The name of the resource.
        :param IssueLabelArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IssueLabelArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 color: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 repository: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = IssueLabelArgs.__new__(IssueLabelArgs)

            if color is None and not opts.urn:
                raise TypeError("Missing required property 'color'")
            __props__.__dict__["color"] = color
            __props__.__dict__["description"] = description
            __props__.__dict__["name"] = name
            if repository is None and not opts.urn:
                raise TypeError("Missing required property 'repository'")
            __props__.__dict__["repository"] = repository
            __props__.__dict__["etag"] = None
            __props__.__dict__["url"] = None
        super(IssueLabel, __self__).__init__(
            'github:index/issueLabel:IssueLabel',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            color: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            etag: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            repository: Optional[pulumi.Input[str]] = None,
            url: Optional[pulumi.Input[str]] = None) -> 'IssueLabel':
        """
        Get an existing IssueLabel resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] color: A 6 character hex code, **without the leading #**, identifying the color of the label.
        :param pulumi.Input[str] description: A short description of the label.
        :param pulumi.Input[str] name: The name of the label.
        :param pulumi.Input[str] repository: The GitHub repository
        :param pulumi.Input[str] url: The URL to the issue label
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _IssueLabelState.__new__(_IssueLabelState)

        __props__.__dict__["color"] = color
        __props__.__dict__["description"] = description
        __props__.__dict__["etag"] = etag
        __props__.__dict__["name"] = name
        __props__.__dict__["repository"] = repository
        __props__.__dict__["url"] = url
        return IssueLabel(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def color(self) -> pulumi.Output[str]:
        """
        A 6 character hex code, **without the leading #**, identifying the color of the label.
        """
        return pulumi.get(self, "color")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A short description of the label.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the label.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def repository(self) -> pulumi.Output[str]:
        """
        The GitHub repository
        """
        return pulumi.get(self, "repository")

    @property
    @pulumi.getter
    def url(self) -> pulumi.Output[str]:
        """
        The URL to the issue label
        """
        return pulumi.get(self, "url")

