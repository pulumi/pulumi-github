# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
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

__all__ = ['BranchArgs', 'Branch']

@pulumi.input_type
class BranchArgs:
    def __init__(__self__, *,
                 branch: pulumi.Input[builtins.str],
                 repository: pulumi.Input[builtins.str],
                 source_branch: Optional[pulumi.Input[builtins.str]] = None,
                 source_sha: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a Branch resource.
        :param pulumi.Input[builtins.str] branch: The repository branch to create.
        :param pulumi.Input[builtins.str] repository: The GitHub repository name.
        :param pulumi.Input[builtins.str] source_branch: The branch name to start from. Defaults to `main`.
        :param pulumi.Input[builtins.str] source_sha: The commit hash to start from. Defaults to the tip of `source_branch`. If provided, `source_branch` is ignored.
        """
        pulumi.set(__self__, "branch", branch)
        pulumi.set(__self__, "repository", repository)
        if source_branch is not None:
            pulumi.set(__self__, "source_branch", source_branch)
        if source_sha is not None:
            pulumi.set(__self__, "source_sha", source_sha)

    @property
    @pulumi.getter
    def branch(self) -> pulumi.Input[builtins.str]:
        """
        The repository branch to create.
        """
        return pulumi.get(self, "branch")

    @branch.setter
    def branch(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "branch", value)

    @property
    @pulumi.getter
    def repository(self) -> pulumi.Input[builtins.str]:
        """
        The GitHub repository name.
        """
        return pulumi.get(self, "repository")

    @repository.setter
    def repository(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "repository", value)

    @property
    @pulumi.getter(name="sourceBranch")
    def source_branch(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The branch name to start from. Defaults to `main`.
        """
        return pulumi.get(self, "source_branch")

    @source_branch.setter
    def source_branch(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "source_branch", value)

    @property
    @pulumi.getter(name="sourceSha")
    def source_sha(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The commit hash to start from. Defaults to the tip of `source_branch`. If provided, `source_branch` is ignored.
        """
        return pulumi.get(self, "source_sha")

    @source_sha.setter
    def source_sha(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "source_sha", value)


@pulumi.input_type
class _BranchState:
    def __init__(__self__, *,
                 branch: Optional[pulumi.Input[builtins.str]] = None,
                 etag: Optional[pulumi.Input[builtins.str]] = None,
                 ref: Optional[pulumi.Input[builtins.str]] = None,
                 repository: Optional[pulumi.Input[builtins.str]] = None,
                 sha: Optional[pulumi.Input[builtins.str]] = None,
                 source_branch: Optional[pulumi.Input[builtins.str]] = None,
                 source_sha: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering Branch resources.
        :param pulumi.Input[builtins.str] branch: The repository branch to create.
        :param pulumi.Input[builtins.str] etag: An etag representing the Branch object.
        :param pulumi.Input[builtins.str] ref: A string representing a branch reference, in the form of `refs/heads/<branch>`.
        :param pulumi.Input[builtins.str] repository: The GitHub repository name.
        :param pulumi.Input[builtins.str] sha: A string storing the reference's `HEAD` commit's SHA1.
        :param pulumi.Input[builtins.str] source_branch: The branch name to start from. Defaults to `main`.
        :param pulumi.Input[builtins.str] source_sha: The commit hash to start from. Defaults to the tip of `source_branch`. If provided, `source_branch` is ignored.
        """
        if branch is not None:
            pulumi.set(__self__, "branch", branch)
        if etag is not None:
            pulumi.set(__self__, "etag", etag)
        if ref is not None:
            pulumi.set(__self__, "ref", ref)
        if repository is not None:
            pulumi.set(__self__, "repository", repository)
        if sha is not None:
            pulumi.set(__self__, "sha", sha)
        if source_branch is not None:
            pulumi.set(__self__, "source_branch", source_branch)
        if source_sha is not None:
            pulumi.set(__self__, "source_sha", source_sha)

    @property
    @pulumi.getter
    def branch(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The repository branch to create.
        """
        return pulumi.get(self, "branch")

    @branch.setter
    def branch(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "branch", value)

    @property
    @pulumi.getter
    def etag(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        An etag representing the Branch object.
        """
        return pulumi.get(self, "etag")

    @etag.setter
    def etag(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "etag", value)

    @property
    @pulumi.getter
    def ref(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        A string representing a branch reference, in the form of `refs/heads/<branch>`.
        """
        return pulumi.get(self, "ref")

    @ref.setter
    def ref(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "ref", value)

    @property
    @pulumi.getter
    def repository(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The GitHub repository name.
        """
        return pulumi.get(self, "repository")

    @repository.setter
    def repository(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "repository", value)

    @property
    @pulumi.getter
    def sha(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        A string storing the reference's `HEAD` commit's SHA1.
        """
        return pulumi.get(self, "sha")

    @sha.setter
    def sha(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "sha", value)

    @property
    @pulumi.getter(name="sourceBranch")
    def source_branch(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The branch name to start from. Defaults to `main`.
        """
        return pulumi.get(self, "source_branch")

    @source_branch.setter
    def source_branch(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "source_branch", value)

    @property
    @pulumi.getter(name="sourceSha")
    def source_sha(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The commit hash to start from. Defaults to the tip of `source_branch`. If provided, `source_branch` is ignored.
        """
        return pulumi.get(self, "source_sha")

    @source_sha.setter
    def source_sha(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "source_sha", value)


class Branch(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 branch: Optional[pulumi.Input[builtins.str]] = None,
                 repository: Optional[pulumi.Input[builtins.str]] = None,
                 source_branch: Optional[pulumi.Input[builtins.str]] = None,
                 source_sha: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        This resource allows you to create and manage branches within your repository.

        Additional constraints can be applied to ensure your branch is created from
        another branch or commit.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_github as github

        development = github.Branch("development",
            repository="example",
            branch="development")
        ```

        ## Import

        GitHub Branch can be imported using an ID made up of `repository:branch`, e.g.

        ```sh
        $ pulumi import github:index/branch:Branch terraform terraform:main
        ```
        Importing github branch into an instance object (when using a for each block to manage multiple branches)

        ```sh
        $ pulumi import github:index/branch:Branch terraform["terraform"] terraform:main
        ```
        Optionally, a source branch may be specified using an ID of `repository:branch:source_branch`.
        This is useful for importing branches that do not branch directly off main.

        ```sh
        $ pulumi import github:index/branch:Branch terraform terraform:feature-branch:dev
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] branch: The repository branch to create.
        :param pulumi.Input[builtins.str] repository: The GitHub repository name.
        :param pulumi.Input[builtins.str] source_branch: The branch name to start from. Defaults to `main`.
        :param pulumi.Input[builtins.str] source_sha: The commit hash to start from. Defaults to the tip of `source_branch`. If provided, `source_branch` is ignored.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: BranchArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This resource allows you to create and manage branches within your repository.

        Additional constraints can be applied to ensure your branch is created from
        another branch or commit.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_github as github

        development = github.Branch("development",
            repository="example",
            branch="development")
        ```

        ## Import

        GitHub Branch can be imported using an ID made up of `repository:branch`, e.g.

        ```sh
        $ pulumi import github:index/branch:Branch terraform terraform:main
        ```
        Importing github branch into an instance object (when using a for each block to manage multiple branches)

        ```sh
        $ pulumi import github:index/branch:Branch terraform["terraform"] terraform:main
        ```
        Optionally, a source branch may be specified using an ID of `repository:branch:source_branch`.
        This is useful for importing branches that do not branch directly off main.

        ```sh
        $ pulumi import github:index/branch:Branch terraform terraform:feature-branch:dev
        ```

        :param str resource_name: The name of the resource.
        :param BranchArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BranchArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 branch: Optional[pulumi.Input[builtins.str]] = None,
                 repository: Optional[pulumi.Input[builtins.str]] = None,
                 source_branch: Optional[pulumi.Input[builtins.str]] = None,
                 source_sha: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = BranchArgs.__new__(BranchArgs)

            if branch is None and not opts.urn:
                raise TypeError("Missing required property 'branch'")
            __props__.__dict__["branch"] = branch
            if repository is None and not opts.urn:
                raise TypeError("Missing required property 'repository'")
            __props__.__dict__["repository"] = repository
            __props__.__dict__["source_branch"] = source_branch
            __props__.__dict__["source_sha"] = source_sha
            __props__.__dict__["etag"] = None
            __props__.__dict__["ref"] = None
            __props__.__dict__["sha"] = None
        super(Branch, __self__).__init__(
            'github:index/branch:Branch',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            branch: Optional[pulumi.Input[builtins.str]] = None,
            etag: Optional[pulumi.Input[builtins.str]] = None,
            ref: Optional[pulumi.Input[builtins.str]] = None,
            repository: Optional[pulumi.Input[builtins.str]] = None,
            sha: Optional[pulumi.Input[builtins.str]] = None,
            source_branch: Optional[pulumi.Input[builtins.str]] = None,
            source_sha: Optional[pulumi.Input[builtins.str]] = None) -> 'Branch':
        """
        Get an existing Branch resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] branch: The repository branch to create.
        :param pulumi.Input[builtins.str] etag: An etag representing the Branch object.
        :param pulumi.Input[builtins.str] ref: A string representing a branch reference, in the form of `refs/heads/<branch>`.
        :param pulumi.Input[builtins.str] repository: The GitHub repository name.
        :param pulumi.Input[builtins.str] sha: A string storing the reference's `HEAD` commit's SHA1.
        :param pulumi.Input[builtins.str] source_branch: The branch name to start from. Defaults to `main`.
        :param pulumi.Input[builtins.str] source_sha: The commit hash to start from. Defaults to the tip of `source_branch`. If provided, `source_branch` is ignored.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _BranchState.__new__(_BranchState)

        __props__.__dict__["branch"] = branch
        __props__.__dict__["etag"] = etag
        __props__.__dict__["ref"] = ref
        __props__.__dict__["repository"] = repository
        __props__.__dict__["sha"] = sha
        __props__.__dict__["source_branch"] = source_branch
        __props__.__dict__["source_sha"] = source_sha
        return Branch(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def branch(self) -> pulumi.Output[builtins.str]:
        """
        The repository branch to create.
        """
        return pulumi.get(self, "branch")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[builtins.str]:
        """
        An etag representing the Branch object.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def ref(self) -> pulumi.Output[builtins.str]:
        """
        A string representing a branch reference, in the form of `refs/heads/<branch>`.
        """
        return pulumi.get(self, "ref")

    @property
    @pulumi.getter
    def repository(self) -> pulumi.Output[builtins.str]:
        """
        The GitHub repository name.
        """
        return pulumi.get(self, "repository")

    @property
    @pulumi.getter
    def sha(self) -> pulumi.Output[builtins.str]:
        """
        A string storing the reference's `HEAD` commit's SHA1.
        """
        return pulumi.get(self, "sha")

    @property
    @pulumi.getter(name="sourceBranch")
    def source_branch(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The branch name to start from. Defaults to `main`.
        """
        return pulumi.get(self, "source_branch")

    @property
    @pulumi.getter(name="sourceSha")
    def source_sha(self) -> pulumi.Output[builtins.str]:
        """
        The commit hash to start from. Defaults to the tip of `source_branch`. If provided, `source_branch` is ignored.
        """
        return pulumi.get(self, "source_sha")

