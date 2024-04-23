# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ProjectCardArgs', 'ProjectCard']

@pulumi.input_type
class ProjectCardArgs:
    def __init__(__self__, *,
                 column_id: pulumi.Input[str],
                 content_id: Optional[pulumi.Input[int]] = None,
                 content_type: Optional[pulumi.Input[str]] = None,
                 note: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ProjectCard resource.
        :param pulumi.Input[str] column_id: The ID of the card.
        :param pulumi.Input[int] content_id: `github_issue.issue_id`.
        :param pulumi.Input[str] content_type: Must be either `Issue` or `PullRequest`
               
               **Remarks:** You must either set the `note` attribute or both `content_id` and `content_type`.
               See note example or issue example for more information.
        :param pulumi.Input[str] note: The note contents of the card. Markdown supported.
        """
        pulumi.set(__self__, "column_id", column_id)
        if content_id is not None:
            pulumi.set(__self__, "content_id", content_id)
        if content_type is not None:
            pulumi.set(__self__, "content_type", content_type)
        if note is not None:
            pulumi.set(__self__, "note", note)

    @property
    @pulumi.getter(name="columnId")
    def column_id(self) -> pulumi.Input[str]:
        """
        The ID of the card.
        """
        return pulumi.get(self, "column_id")

    @column_id.setter
    def column_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "column_id", value)

    @property
    @pulumi.getter(name="contentId")
    def content_id(self) -> Optional[pulumi.Input[int]]:
        """
        `github_issue.issue_id`.
        """
        return pulumi.get(self, "content_id")

    @content_id.setter
    def content_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "content_id", value)

    @property
    @pulumi.getter(name="contentType")
    def content_type(self) -> Optional[pulumi.Input[str]]:
        """
        Must be either `Issue` or `PullRequest`

        **Remarks:** You must either set the `note` attribute or both `content_id` and `content_type`.
        See note example or issue example for more information.
        """
        return pulumi.get(self, "content_type")

    @content_type.setter
    def content_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "content_type", value)

    @property
    @pulumi.getter
    def note(self) -> Optional[pulumi.Input[str]]:
        """
        The note contents of the card. Markdown supported.
        """
        return pulumi.get(self, "note")

    @note.setter
    def note(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "note", value)


@pulumi.input_type
class _ProjectCardState:
    def __init__(__self__, *,
                 card_id: Optional[pulumi.Input[int]] = None,
                 column_id: Optional[pulumi.Input[str]] = None,
                 content_id: Optional[pulumi.Input[int]] = None,
                 content_type: Optional[pulumi.Input[str]] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 note: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ProjectCard resources.
        :param pulumi.Input[int] card_id: The ID of the card.
        :param pulumi.Input[str] column_id: The ID of the card.
        :param pulumi.Input[int] content_id: `github_issue.issue_id`.
        :param pulumi.Input[str] content_type: Must be either `Issue` or `PullRequest`
               
               **Remarks:** You must either set the `note` attribute or both `content_id` and `content_type`.
               See note example or issue example for more information.
        :param pulumi.Input[str] note: The note contents of the card. Markdown supported.
        """
        if card_id is not None:
            pulumi.set(__self__, "card_id", card_id)
        if column_id is not None:
            pulumi.set(__self__, "column_id", column_id)
        if content_id is not None:
            pulumi.set(__self__, "content_id", content_id)
        if content_type is not None:
            pulumi.set(__self__, "content_type", content_type)
        if etag is not None:
            pulumi.set(__self__, "etag", etag)
        if note is not None:
            pulumi.set(__self__, "note", note)

    @property
    @pulumi.getter(name="cardId")
    def card_id(self) -> Optional[pulumi.Input[int]]:
        """
        The ID of the card.
        """
        return pulumi.get(self, "card_id")

    @card_id.setter
    def card_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "card_id", value)

    @property
    @pulumi.getter(name="columnId")
    def column_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the card.
        """
        return pulumi.get(self, "column_id")

    @column_id.setter
    def column_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "column_id", value)

    @property
    @pulumi.getter(name="contentId")
    def content_id(self) -> Optional[pulumi.Input[int]]:
        """
        `github_issue.issue_id`.
        """
        return pulumi.get(self, "content_id")

    @content_id.setter
    def content_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "content_id", value)

    @property
    @pulumi.getter(name="contentType")
    def content_type(self) -> Optional[pulumi.Input[str]]:
        """
        Must be either `Issue` or `PullRequest`

        **Remarks:** You must either set the `note` attribute or both `content_id` and `content_type`.
        See note example or issue example for more information.
        """
        return pulumi.get(self, "content_type")

    @content_type.setter
    def content_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "content_type", value)

    @property
    @pulumi.getter
    def etag(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "etag")

    @etag.setter
    def etag(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "etag", value)

    @property
    @pulumi.getter
    def note(self) -> Optional[pulumi.Input[str]]:
        """
        The note contents of the card. Markdown supported.
        """
        return pulumi.get(self, "note")

    @note.setter
    def note(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "note", value)


class ProjectCard(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 column_id: Optional[pulumi.Input[str]] = None,
                 content_id: Optional[pulumi.Input[int]] = None,
                 content_type: Optional[pulumi.Input[str]] = None,
                 note: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        This resource allows you to create and manage cards for GitHub projects.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_github as github

        project = github.OrganizationProject("project",
            name="An Organization Project",
            body="This is an organization project.")
        column = github.ProjectColumn("column",
            project_id=project.id,
            name="Backlog")
        card = github.ProjectCard("card",
            column_id=column.column_id,
            note="## Unaccepted 👇")
        ```

        ### Adding An Issue To A Project

        ```python
        import pulumi
        import pulumi_github as github

        test = github.Repository("test",
            name="myrepo",
            has_projects=True,
            has_issues=True)
        test_issue = github.Issue("test",
            repository=test.id,
            title="Test issue title",
            body="Test issue body")
        test_repository_project = github.RepositoryProject("test",
            name="test",
            repository=test.name,
            body="this is a test project")
        test_project_column = github.ProjectColumn("test",
            project_id=test_repository_project.id,
            name="Backlog")
        test_project_card = github.ProjectCard("test",
            column_id=test_project_column.column_id,
            content_id=test_issue.issue_id,
            content_type="Issue")
        ```

        ## Import

        A GitHub Project Card can be imported using its [Card ID](https://developer.github.com/v3/projects/cards/#get-a-project-card):

        ```sh
        $ pulumi import github:index/projectCard:ProjectCard card 01234567
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] column_id: The ID of the card.
        :param pulumi.Input[int] content_id: `github_issue.issue_id`.
        :param pulumi.Input[str] content_type: Must be either `Issue` or `PullRequest`
               
               **Remarks:** You must either set the `note` attribute or both `content_id` and `content_type`.
               See note example or issue example for more information.
        :param pulumi.Input[str] note: The note contents of the card. Markdown supported.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ProjectCardArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This resource allows you to create and manage cards for GitHub projects.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_github as github

        project = github.OrganizationProject("project",
            name="An Organization Project",
            body="This is an organization project.")
        column = github.ProjectColumn("column",
            project_id=project.id,
            name="Backlog")
        card = github.ProjectCard("card",
            column_id=column.column_id,
            note="## Unaccepted 👇")
        ```

        ### Adding An Issue To A Project

        ```python
        import pulumi
        import pulumi_github as github

        test = github.Repository("test",
            name="myrepo",
            has_projects=True,
            has_issues=True)
        test_issue = github.Issue("test",
            repository=test.id,
            title="Test issue title",
            body="Test issue body")
        test_repository_project = github.RepositoryProject("test",
            name="test",
            repository=test.name,
            body="this is a test project")
        test_project_column = github.ProjectColumn("test",
            project_id=test_repository_project.id,
            name="Backlog")
        test_project_card = github.ProjectCard("test",
            column_id=test_project_column.column_id,
            content_id=test_issue.issue_id,
            content_type="Issue")
        ```

        ## Import

        A GitHub Project Card can be imported using its [Card ID](https://developer.github.com/v3/projects/cards/#get-a-project-card):

        ```sh
        $ pulumi import github:index/projectCard:ProjectCard card 01234567
        ```

        :param str resource_name: The name of the resource.
        :param ProjectCardArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProjectCardArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 column_id: Optional[pulumi.Input[str]] = None,
                 content_id: Optional[pulumi.Input[int]] = None,
                 content_type: Optional[pulumi.Input[str]] = None,
                 note: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProjectCardArgs.__new__(ProjectCardArgs)

            if column_id is None and not opts.urn:
                raise TypeError("Missing required property 'column_id'")
            __props__.__dict__["column_id"] = column_id
            __props__.__dict__["content_id"] = content_id
            __props__.__dict__["content_type"] = content_type
            __props__.__dict__["note"] = note
            __props__.__dict__["card_id"] = None
            __props__.__dict__["etag"] = None
        super(ProjectCard, __self__).__init__(
            'github:index/projectCard:ProjectCard',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            card_id: Optional[pulumi.Input[int]] = None,
            column_id: Optional[pulumi.Input[str]] = None,
            content_id: Optional[pulumi.Input[int]] = None,
            content_type: Optional[pulumi.Input[str]] = None,
            etag: Optional[pulumi.Input[str]] = None,
            note: Optional[pulumi.Input[str]] = None) -> 'ProjectCard':
        """
        Get an existing ProjectCard resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] card_id: The ID of the card.
        :param pulumi.Input[str] column_id: The ID of the card.
        :param pulumi.Input[int] content_id: `github_issue.issue_id`.
        :param pulumi.Input[str] content_type: Must be either `Issue` or `PullRequest`
               
               **Remarks:** You must either set the `note` attribute or both `content_id` and `content_type`.
               See note example or issue example for more information.
        :param pulumi.Input[str] note: The note contents of the card. Markdown supported.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ProjectCardState.__new__(_ProjectCardState)

        __props__.__dict__["card_id"] = card_id
        __props__.__dict__["column_id"] = column_id
        __props__.__dict__["content_id"] = content_id
        __props__.__dict__["content_type"] = content_type
        __props__.__dict__["etag"] = etag
        __props__.__dict__["note"] = note
        return ProjectCard(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="cardId")
    def card_id(self) -> pulumi.Output[int]:
        """
        The ID of the card.
        """
        return pulumi.get(self, "card_id")

    @property
    @pulumi.getter(name="columnId")
    def column_id(self) -> pulumi.Output[str]:
        """
        The ID of the card.
        """
        return pulumi.get(self, "column_id")

    @property
    @pulumi.getter(name="contentId")
    def content_id(self) -> pulumi.Output[Optional[int]]:
        """
        `github_issue.issue_id`.
        """
        return pulumi.get(self, "content_id")

    @property
    @pulumi.getter(name="contentType")
    def content_type(self) -> pulumi.Output[Optional[str]]:
        """
        Must be either `Issue` or `PullRequest`

        **Remarks:** You must either set the `note` attribute or both `content_id` and `content_type`.
        See note example or issue example for more information.
        """
        return pulumi.get(self, "content_type")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def note(self) -> pulumi.Output[Optional[str]]:
        """
        The note contents of the card. Markdown supported.
        """
        return pulumi.get(self, "note")

