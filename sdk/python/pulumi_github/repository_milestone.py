# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['RepositoryMilestoneArgs', 'RepositoryMilestone']

@pulumi.input_type
class RepositoryMilestoneArgs:
    def __init__(__self__, *,
                 owner: pulumi.Input[str],
                 repository: pulumi.Input[str],
                 title: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 due_date: Optional[pulumi.Input[str]] = None,
                 state: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a RepositoryMilestone resource.
        :param pulumi.Input[str] owner: The owner of the GitHub Repository.
        :param pulumi.Input[str] repository: The name of the GitHub Repository.
        :param pulumi.Input[str] title: The title of the milestone.
        :param pulumi.Input[str] description: A description of the milestone.
        :param pulumi.Input[str] due_date: The milestone due date. In `yyyy-mm-dd` format.
        :param pulumi.Input[str] state: The state of the milestone. Either `open` or `closed`. Default: `open`
        """
        pulumi.set(__self__, "owner", owner)
        pulumi.set(__self__, "repository", repository)
        pulumi.set(__self__, "title", title)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if due_date is not None:
            pulumi.set(__self__, "due_date", due_date)
        if state is not None:
            pulumi.set(__self__, "state", state)

    @property
    @pulumi.getter
    def owner(self) -> pulumi.Input[str]:
        """
        The owner of the GitHub Repository.
        """
        return pulumi.get(self, "owner")

    @owner.setter
    def owner(self, value: pulumi.Input[str]):
        pulumi.set(self, "owner", value)

    @property
    @pulumi.getter
    def repository(self) -> pulumi.Input[str]:
        """
        The name of the GitHub Repository.
        """
        return pulumi.get(self, "repository")

    @repository.setter
    def repository(self, value: pulumi.Input[str]):
        pulumi.set(self, "repository", value)

    @property
    @pulumi.getter
    def title(self) -> pulumi.Input[str]:
        """
        The title of the milestone.
        """
        return pulumi.get(self, "title")

    @title.setter
    def title(self, value: pulumi.Input[str]):
        pulumi.set(self, "title", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A description of the milestone.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="dueDate")
    def due_date(self) -> Optional[pulumi.Input[str]]:
        """
        The milestone due date. In `yyyy-mm-dd` format.
        """
        return pulumi.get(self, "due_date")

    @due_date.setter
    def due_date(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "due_date", value)

    @property
    @pulumi.getter
    def state(self) -> Optional[pulumi.Input[str]]:
        """
        The state of the milestone. Either `open` or `closed`. Default: `open`
        """
        return pulumi.get(self, "state")

    @state.setter
    def state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "state", value)


@pulumi.input_type
class _RepositoryMilestoneState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 due_date: Optional[pulumi.Input[str]] = None,
                 number: Optional[pulumi.Input[int]] = None,
                 owner: Optional[pulumi.Input[str]] = None,
                 repository: Optional[pulumi.Input[str]] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 title: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering RepositoryMilestone resources.
        :param pulumi.Input[str] description: A description of the milestone.
        :param pulumi.Input[str] due_date: The milestone due date. In `yyyy-mm-dd` format.
        :param pulumi.Input[int] number: The number of the milestone.
        :param pulumi.Input[str] owner: The owner of the GitHub Repository.
        :param pulumi.Input[str] repository: The name of the GitHub Repository.
        :param pulumi.Input[str] state: The state of the milestone. Either `open` or `closed`. Default: `open`
        :param pulumi.Input[str] title: The title of the milestone.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if due_date is not None:
            pulumi.set(__self__, "due_date", due_date)
        if number is not None:
            pulumi.set(__self__, "number", number)
        if owner is not None:
            pulumi.set(__self__, "owner", owner)
        if repository is not None:
            pulumi.set(__self__, "repository", repository)
        if state is not None:
            pulumi.set(__self__, "state", state)
        if title is not None:
            pulumi.set(__self__, "title", title)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A description of the milestone.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="dueDate")
    def due_date(self) -> Optional[pulumi.Input[str]]:
        """
        The milestone due date. In `yyyy-mm-dd` format.
        """
        return pulumi.get(self, "due_date")

    @due_date.setter
    def due_date(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "due_date", value)

    @property
    @pulumi.getter
    def number(self) -> Optional[pulumi.Input[int]]:
        """
        The number of the milestone.
        """
        return pulumi.get(self, "number")

    @number.setter
    def number(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "number", value)

    @property
    @pulumi.getter
    def owner(self) -> Optional[pulumi.Input[str]]:
        """
        The owner of the GitHub Repository.
        """
        return pulumi.get(self, "owner")

    @owner.setter
    def owner(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "owner", value)

    @property
    @pulumi.getter
    def repository(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the GitHub Repository.
        """
        return pulumi.get(self, "repository")

    @repository.setter
    def repository(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "repository", value)

    @property
    @pulumi.getter
    def state(self) -> Optional[pulumi.Input[str]]:
        """
        The state of the milestone. Either `open` or `closed`. Default: `open`
        """
        return pulumi.get(self, "state")

    @state.setter
    def state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "state", value)

    @property
    @pulumi.getter
    def title(self) -> Optional[pulumi.Input[str]]:
        """
        The title of the milestone.
        """
        return pulumi.get(self, "title")

    @title.setter
    def title(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "title", value)


class RepositoryMilestone(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 due_date: Optional[pulumi.Input[str]] = None,
                 owner: Optional[pulumi.Input[str]] = None,
                 repository: Optional[pulumi.Input[str]] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 title: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a GitHub repository milestone resource.

        This resource allows you to create and manage milestones for a GitHub Repository within an organization or user account.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_github as github

        # Create a milestone for a repository
        example = github.RepositoryMilestone("example",
            owner="example-owner",
            repository="example-repository",
            title="v1.1.0")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        A GitHub Repository Milestone can be imported using an ID made up of `owner/repository/number`, e.g.

        ```sh
        $ pulumi import github:index/repositoryMilestone:RepositoryMilestone example example-owner/example-repository/1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A description of the milestone.
        :param pulumi.Input[str] due_date: The milestone due date. In `yyyy-mm-dd` format.
        :param pulumi.Input[str] owner: The owner of the GitHub Repository.
        :param pulumi.Input[str] repository: The name of the GitHub Repository.
        :param pulumi.Input[str] state: The state of the milestone. Either `open` or `closed`. Default: `open`
        :param pulumi.Input[str] title: The title of the milestone.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RepositoryMilestoneArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a GitHub repository milestone resource.

        This resource allows you to create and manage milestones for a GitHub Repository within an organization or user account.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_github as github

        # Create a milestone for a repository
        example = github.RepositoryMilestone("example",
            owner="example-owner",
            repository="example-repository",
            title="v1.1.0")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        A GitHub Repository Milestone can be imported using an ID made up of `owner/repository/number`, e.g.

        ```sh
        $ pulumi import github:index/repositoryMilestone:RepositoryMilestone example example-owner/example-repository/1
        ```

        :param str resource_name: The name of the resource.
        :param RepositoryMilestoneArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RepositoryMilestoneArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 due_date: Optional[pulumi.Input[str]] = None,
                 owner: Optional[pulumi.Input[str]] = None,
                 repository: Optional[pulumi.Input[str]] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 title: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RepositoryMilestoneArgs.__new__(RepositoryMilestoneArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["due_date"] = due_date
            if owner is None and not opts.urn:
                raise TypeError("Missing required property 'owner'")
            __props__.__dict__["owner"] = owner
            if repository is None and not opts.urn:
                raise TypeError("Missing required property 'repository'")
            __props__.__dict__["repository"] = repository
            __props__.__dict__["state"] = state
            if title is None and not opts.urn:
                raise TypeError("Missing required property 'title'")
            __props__.__dict__["title"] = title
            __props__.__dict__["number"] = None
        super(RepositoryMilestone, __self__).__init__(
            'github:index/repositoryMilestone:RepositoryMilestone',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            due_date: Optional[pulumi.Input[str]] = None,
            number: Optional[pulumi.Input[int]] = None,
            owner: Optional[pulumi.Input[str]] = None,
            repository: Optional[pulumi.Input[str]] = None,
            state: Optional[pulumi.Input[str]] = None,
            title: Optional[pulumi.Input[str]] = None) -> 'RepositoryMilestone':
        """
        Get an existing RepositoryMilestone resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A description of the milestone.
        :param pulumi.Input[str] due_date: The milestone due date. In `yyyy-mm-dd` format.
        :param pulumi.Input[int] number: The number of the milestone.
        :param pulumi.Input[str] owner: The owner of the GitHub Repository.
        :param pulumi.Input[str] repository: The name of the GitHub Repository.
        :param pulumi.Input[str] state: The state of the milestone. Either `open` or `closed`. Default: `open`
        :param pulumi.Input[str] title: The title of the milestone.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RepositoryMilestoneState.__new__(_RepositoryMilestoneState)

        __props__.__dict__["description"] = description
        __props__.__dict__["due_date"] = due_date
        __props__.__dict__["number"] = number
        __props__.__dict__["owner"] = owner
        __props__.__dict__["repository"] = repository
        __props__.__dict__["state"] = state
        __props__.__dict__["title"] = title
        return RepositoryMilestone(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A description of the milestone.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="dueDate")
    def due_date(self) -> pulumi.Output[Optional[str]]:
        """
        The milestone due date. In `yyyy-mm-dd` format.
        """
        return pulumi.get(self, "due_date")

    @property
    @pulumi.getter
    def number(self) -> pulumi.Output[int]:
        """
        The number of the milestone.
        """
        return pulumi.get(self, "number")

    @property
    @pulumi.getter
    def owner(self) -> pulumi.Output[str]:
        """
        The owner of the GitHub Repository.
        """
        return pulumi.get(self, "owner")

    @property
    @pulumi.getter
    def repository(self) -> pulumi.Output[str]:
        """
        The name of the GitHub Repository.
        """
        return pulumi.get(self, "repository")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[Optional[str]]:
        """
        The state of the milestone. Either `open` or `closed`. Default: `open`
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def title(self) -> pulumi.Output[str]:
        """
        The title of the milestone.
        """
        return pulumi.get(self, "title")

