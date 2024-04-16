# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['RepositoryTopicsArgs', 'RepositoryTopics']

@pulumi.input_type
class RepositoryTopicsArgs:
    def __init__(__self__, *,
                 repository: pulumi.Input[str],
                 topics: pulumi.Input[Sequence[pulumi.Input[str]]]):
        """
        The set of arguments for constructing a RepositoryTopics resource.
        :param pulumi.Input[str] repository: The repository name.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] topics: A list of topics to add to the repository.
        """
        pulumi.set(__self__, "repository", repository)
        pulumi.set(__self__, "topics", topics)

    @property
    @pulumi.getter
    def repository(self) -> pulumi.Input[str]:
        """
        The repository name.
        """
        return pulumi.get(self, "repository")

    @repository.setter
    def repository(self, value: pulumi.Input[str]):
        pulumi.set(self, "repository", value)

    @property
    @pulumi.getter
    def topics(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        A list of topics to add to the repository.
        """
        return pulumi.get(self, "topics")

    @topics.setter
    def topics(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "topics", value)


@pulumi.input_type
class _RepositoryTopicsState:
    def __init__(__self__, *,
                 repository: Optional[pulumi.Input[str]] = None,
                 topics: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering RepositoryTopics resources.
        :param pulumi.Input[str] repository: The repository name.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] topics: A list of topics to add to the repository.
        """
        if repository is not None:
            pulumi.set(__self__, "repository", repository)
        if topics is not None:
            pulumi.set(__self__, "topics", topics)

    @property
    @pulumi.getter
    def repository(self) -> Optional[pulumi.Input[str]]:
        """
        The repository name.
        """
        return pulumi.get(self, "repository")

    @repository.setter
    def repository(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "repository", value)

    @property
    @pulumi.getter
    def topics(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of topics to add to the repository.
        """
        return pulumi.get(self, "topics")

    @topics.setter
    def topics(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "topics", value)


class RepositoryTopics(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 repository: Optional[pulumi.Input[str]] = None,
                 topics: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_github as github

        test = github.get_repository(name="test")
        test_repository_topics = github.RepositoryTopics("test",
            repository=test_github_repository["name"],
            topics=[
                "topic-1",
                "topic-2",
            ])
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Repository topics can be imported using the `name` of the repository.

        ```sh
        $ pulumi import github:index/repositoryTopics:RepositoryTopics terraform terraform
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] repository: The repository name.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] topics: A list of topics to add to the repository.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RepositoryTopicsArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_github as github

        test = github.get_repository(name="test")
        test_repository_topics = github.RepositoryTopics("test",
            repository=test_github_repository["name"],
            topics=[
                "topic-1",
                "topic-2",
            ])
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Repository topics can be imported using the `name` of the repository.

        ```sh
        $ pulumi import github:index/repositoryTopics:RepositoryTopics terraform terraform
        ```

        :param str resource_name: The name of the resource.
        :param RepositoryTopicsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RepositoryTopicsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 repository: Optional[pulumi.Input[str]] = None,
                 topics: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RepositoryTopicsArgs.__new__(RepositoryTopicsArgs)

            if repository is None and not opts.urn:
                raise TypeError("Missing required property 'repository'")
            __props__.__dict__["repository"] = repository
            if topics is None and not opts.urn:
                raise TypeError("Missing required property 'topics'")
            __props__.__dict__["topics"] = topics
        super(RepositoryTopics, __self__).__init__(
            'github:index/repositoryTopics:RepositoryTopics',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            repository: Optional[pulumi.Input[str]] = None,
            topics: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'RepositoryTopics':
        """
        Get an existing RepositoryTopics resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] repository: The repository name.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] topics: A list of topics to add to the repository.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RepositoryTopicsState.__new__(_RepositoryTopicsState)

        __props__.__dict__["repository"] = repository
        __props__.__dict__["topics"] = topics
        return RepositoryTopics(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def repository(self) -> pulumi.Output[str]:
        """
        The repository name.
        """
        return pulumi.get(self, "repository")

    @property
    @pulumi.getter
    def topics(self) -> pulumi.Output[Sequence[str]]:
        """
        A list of topics to add to the repository.
        """
        return pulumi.get(self, "topics")

