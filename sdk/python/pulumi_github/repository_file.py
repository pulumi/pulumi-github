# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['RepositoryFileArgs', 'RepositoryFile']

@pulumi.input_type
class RepositoryFileArgs:
    def __init__(__self__, *,
                 content: pulumi.Input[str],
                 file: pulumi.Input[str],
                 repository: pulumi.Input[str],
                 branch: Optional[pulumi.Input[str]] = None,
                 commit_author: Optional[pulumi.Input[str]] = None,
                 commit_email: Optional[pulumi.Input[str]] = None,
                 commit_message: Optional[pulumi.Input[str]] = None,
                 overwrite_on_create: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a RepositoryFile resource.
        :param pulumi.Input[str] content: The file content.
        :param pulumi.Input[str] file: The path of the file to manage.
        :param pulumi.Input[str] repository: The repository to create the file in.
        :param pulumi.Input[str] branch: Git branch (defaults to the repository's default branch).
               The branch must already exist, it will not be created if it does not already exist.
        :param pulumi.Input[str] commit_author: Committer author name to use. **NOTE:** GitHub app users may omit author and email information so GitHub can verify commits as the GitHub App. This maybe useful when a branch protection rule requires signed commits.
        :param pulumi.Input[str] commit_email: Committer email address to use. **NOTE:** GitHub app users may omit author and email information so GitHub can verify commits as the GitHub App. This may be useful when a branch protection rule requires signed commits.
        :param pulumi.Input[str] commit_message: Commit message when adding or updating the managed file.
        :param pulumi.Input[bool] overwrite_on_create: Enable overwriting existing files
        """
        RepositoryFileArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            content=content,
            file=file,
            repository=repository,
            branch=branch,
            commit_author=commit_author,
            commit_email=commit_email,
            commit_message=commit_message,
            overwrite_on_create=overwrite_on_create,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             content: Optional[pulumi.Input[str]] = None,
             file: Optional[pulumi.Input[str]] = None,
             repository: Optional[pulumi.Input[str]] = None,
             branch: Optional[pulumi.Input[str]] = None,
             commit_author: Optional[pulumi.Input[str]] = None,
             commit_email: Optional[pulumi.Input[str]] = None,
             commit_message: Optional[pulumi.Input[str]] = None,
             overwrite_on_create: Optional[pulumi.Input[bool]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if content is None:
            raise TypeError("Missing 'content' argument")
        if file is None:
            raise TypeError("Missing 'file' argument")
        if repository is None:
            raise TypeError("Missing 'repository' argument")
        if commit_author is None and 'commitAuthor' in kwargs:
            commit_author = kwargs['commitAuthor']
        if commit_email is None and 'commitEmail' in kwargs:
            commit_email = kwargs['commitEmail']
        if commit_message is None and 'commitMessage' in kwargs:
            commit_message = kwargs['commitMessage']
        if overwrite_on_create is None and 'overwriteOnCreate' in kwargs:
            overwrite_on_create = kwargs['overwriteOnCreate']

        _setter("content", content)
        _setter("file", file)
        _setter("repository", repository)
        if branch is not None:
            _setter("branch", branch)
        if commit_author is not None:
            _setter("commit_author", commit_author)
        if commit_email is not None:
            _setter("commit_email", commit_email)
        if commit_message is not None:
            _setter("commit_message", commit_message)
        if overwrite_on_create is not None:
            _setter("overwrite_on_create", overwrite_on_create)

    @property
    @pulumi.getter
    def content(self) -> pulumi.Input[str]:
        """
        The file content.
        """
        return pulumi.get(self, "content")

    @content.setter
    def content(self, value: pulumi.Input[str]):
        pulumi.set(self, "content", value)

    @property
    @pulumi.getter
    def file(self) -> pulumi.Input[str]:
        """
        The path of the file to manage.
        """
        return pulumi.get(self, "file")

    @file.setter
    def file(self, value: pulumi.Input[str]):
        pulumi.set(self, "file", value)

    @property
    @pulumi.getter
    def repository(self) -> pulumi.Input[str]:
        """
        The repository to create the file in.
        """
        return pulumi.get(self, "repository")

    @repository.setter
    def repository(self, value: pulumi.Input[str]):
        pulumi.set(self, "repository", value)

    @property
    @pulumi.getter
    def branch(self) -> Optional[pulumi.Input[str]]:
        """
        Git branch (defaults to the repository's default branch).
        The branch must already exist, it will not be created if it does not already exist.
        """
        return pulumi.get(self, "branch")

    @branch.setter
    def branch(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "branch", value)

    @property
    @pulumi.getter(name="commitAuthor")
    def commit_author(self) -> Optional[pulumi.Input[str]]:
        """
        Committer author name to use. **NOTE:** GitHub app users may omit author and email information so GitHub can verify commits as the GitHub App. This maybe useful when a branch protection rule requires signed commits.
        """
        return pulumi.get(self, "commit_author")

    @commit_author.setter
    def commit_author(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "commit_author", value)

    @property
    @pulumi.getter(name="commitEmail")
    def commit_email(self) -> Optional[pulumi.Input[str]]:
        """
        Committer email address to use. **NOTE:** GitHub app users may omit author and email information so GitHub can verify commits as the GitHub App. This may be useful when a branch protection rule requires signed commits.
        """
        return pulumi.get(self, "commit_email")

    @commit_email.setter
    def commit_email(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "commit_email", value)

    @property
    @pulumi.getter(name="commitMessage")
    def commit_message(self) -> Optional[pulumi.Input[str]]:
        """
        Commit message when adding or updating the managed file.
        """
        return pulumi.get(self, "commit_message")

    @commit_message.setter
    def commit_message(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "commit_message", value)

    @property
    @pulumi.getter(name="overwriteOnCreate")
    def overwrite_on_create(self) -> Optional[pulumi.Input[bool]]:
        """
        Enable overwriting existing files
        """
        return pulumi.get(self, "overwrite_on_create")

    @overwrite_on_create.setter
    def overwrite_on_create(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "overwrite_on_create", value)


@pulumi.input_type
class _RepositoryFileState:
    def __init__(__self__, *,
                 branch: Optional[pulumi.Input[str]] = None,
                 commit_author: Optional[pulumi.Input[str]] = None,
                 commit_email: Optional[pulumi.Input[str]] = None,
                 commit_message: Optional[pulumi.Input[str]] = None,
                 commit_sha: Optional[pulumi.Input[str]] = None,
                 content: Optional[pulumi.Input[str]] = None,
                 file: Optional[pulumi.Input[str]] = None,
                 overwrite_on_create: Optional[pulumi.Input[bool]] = None,
                 ref: Optional[pulumi.Input[str]] = None,
                 repository: Optional[pulumi.Input[str]] = None,
                 sha: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering RepositoryFile resources.
        :param pulumi.Input[str] branch: Git branch (defaults to the repository's default branch).
               The branch must already exist, it will not be created if it does not already exist.
        :param pulumi.Input[str] commit_author: Committer author name to use. **NOTE:** GitHub app users may omit author and email information so GitHub can verify commits as the GitHub App. This maybe useful when a branch protection rule requires signed commits.
        :param pulumi.Input[str] commit_email: Committer email address to use. **NOTE:** GitHub app users may omit author and email information so GitHub can verify commits as the GitHub App. This may be useful when a branch protection rule requires signed commits.
        :param pulumi.Input[str] commit_message: Commit message when adding or updating the managed file.
        :param pulumi.Input[str] commit_sha: The SHA of the commit that modified the file.
        :param pulumi.Input[str] content: The file content.
        :param pulumi.Input[str] file: The path of the file to manage.
        :param pulumi.Input[bool] overwrite_on_create: Enable overwriting existing files
        :param pulumi.Input[str] ref: The name of the commit/branch/tag.
        :param pulumi.Input[str] repository: The repository to create the file in.
        :param pulumi.Input[str] sha: The SHA blob of the file.
        """
        _RepositoryFileState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            branch=branch,
            commit_author=commit_author,
            commit_email=commit_email,
            commit_message=commit_message,
            commit_sha=commit_sha,
            content=content,
            file=file,
            overwrite_on_create=overwrite_on_create,
            ref=ref,
            repository=repository,
            sha=sha,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             branch: Optional[pulumi.Input[str]] = None,
             commit_author: Optional[pulumi.Input[str]] = None,
             commit_email: Optional[pulumi.Input[str]] = None,
             commit_message: Optional[pulumi.Input[str]] = None,
             commit_sha: Optional[pulumi.Input[str]] = None,
             content: Optional[pulumi.Input[str]] = None,
             file: Optional[pulumi.Input[str]] = None,
             overwrite_on_create: Optional[pulumi.Input[bool]] = None,
             ref: Optional[pulumi.Input[str]] = None,
             repository: Optional[pulumi.Input[str]] = None,
             sha: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if commit_author is None and 'commitAuthor' in kwargs:
            commit_author = kwargs['commitAuthor']
        if commit_email is None and 'commitEmail' in kwargs:
            commit_email = kwargs['commitEmail']
        if commit_message is None and 'commitMessage' in kwargs:
            commit_message = kwargs['commitMessage']
        if commit_sha is None and 'commitSha' in kwargs:
            commit_sha = kwargs['commitSha']
        if overwrite_on_create is None and 'overwriteOnCreate' in kwargs:
            overwrite_on_create = kwargs['overwriteOnCreate']

        if branch is not None:
            _setter("branch", branch)
        if commit_author is not None:
            _setter("commit_author", commit_author)
        if commit_email is not None:
            _setter("commit_email", commit_email)
        if commit_message is not None:
            _setter("commit_message", commit_message)
        if commit_sha is not None:
            _setter("commit_sha", commit_sha)
        if content is not None:
            _setter("content", content)
        if file is not None:
            _setter("file", file)
        if overwrite_on_create is not None:
            _setter("overwrite_on_create", overwrite_on_create)
        if ref is not None:
            _setter("ref", ref)
        if repository is not None:
            _setter("repository", repository)
        if sha is not None:
            _setter("sha", sha)

    @property
    @pulumi.getter
    def branch(self) -> Optional[pulumi.Input[str]]:
        """
        Git branch (defaults to the repository's default branch).
        The branch must already exist, it will not be created if it does not already exist.
        """
        return pulumi.get(self, "branch")

    @branch.setter
    def branch(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "branch", value)

    @property
    @pulumi.getter(name="commitAuthor")
    def commit_author(self) -> Optional[pulumi.Input[str]]:
        """
        Committer author name to use. **NOTE:** GitHub app users may omit author and email information so GitHub can verify commits as the GitHub App. This maybe useful when a branch protection rule requires signed commits.
        """
        return pulumi.get(self, "commit_author")

    @commit_author.setter
    def commit_author(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "commit_author", value)

    @property
    @pulumi.getter(name="commitEmail")
    def commit_email(self) -> Optional[pulumi.Input[str]]:
        """
        Committer email address to use. **NOTE:** GitHub app users may omit author and email information so GitHub can verify commits as the GitHub App. This may be useful when a branch protection rule requires signed commits.
        """
        return pulumi.get(self, "commit_email")

    @commit_email.setter
    def commit_email(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "commit_email", value)

    @property
    @pulumi.getter(name="commitMessage")
    def commit_message(self) -> Optional[pulumi.Input[str]]:
        """
        Commit message when adding or updating the managed file.
        """
        return pulumi.get(self, "commit_message")

    @commit_message.setter
    def commit_message(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "commit_message", value)

    @property
    @pulumi.getter(name="commitSha")
    def commit_sha(self) -> Optional[pulumi.Input[str]]:
        """
        The SHA of the commit that modified the file.
        """
        return pulumi.get(self, "commit_sha")

    @commit_sha.setter
    def commit_sha(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "commit_sha", value)

    @property
    @pulumi.getter
    def content(self) -> Optional[pulumi.Input[str]]:
        """
        The file content.
        """
        return pulumi.get(self, "content")

    @content.setter
    def content(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "content", value)

    @property
    @pulumi.getter
    def file(self) -> Optional[pulumi.Input[str]]:
        """
        The path of the file to manage.
        """
        return pulumi.get(self, "file")

    @file.setter
    def file(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "file", value)

    @property
    @pulumi.getter(name="overwriteOnCreate")
    def overwrite_on_create(self) -> Optional[pulumi.Input[bool]]:
        """
        Enable overwriting existing files
        """
        return pulumi.get(self, "overwrite_on_create")

    @overwrite_on_create.setter
    def overwrite_on_create(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "overwrite_on_create", value)

    @property
    @pulumi.getter
    def ref(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the commit/branch/tag.
        """
        return pulumi.get(self, "ref")

    @ref.setter
    def ref(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ref", value)

    @property
    @pulumi.getter
    def repository(self) -> Optional[pulumi.Input[str]]:
        """
        The repository to create the file in.
        """
        return pulumi.get(self, "repository")

    @repository.setter
    def repository(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "repository", value)

    @property
    @pulumi.getter
    def sha(self) -> Optional[pulumi.Input[str]]:
        """
        The SHA blob of the file.
        """
        return pulumi.get(self, "sha")

    @sha.setter
    def sha(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sha", value)


class RepositoryFile(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 branch: Optional[pulumi.Input[str]] = None,
                 commit_author: Optional[pulumi.Input[str]] = None,
                 commit_email: Optional[pulumi.Input[str]] = None,
                 commit_message: Optional[pulumi.Input[str]] = None,
                 content: Optional[pulumi.Input[str]] = None,
                 file: Optional[pulumi.Input[str]] = None,
                 overwrite_on_create: Optional[pulumi.Input[bool]] = None,
                 repository: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        This resource allows you to create and manage files within a
        GitHub repository.

        ## Import

        Repository files can be imported using a combination of the `repo` and `file`, e.g.

        ```sh
         $ pulumi import github:index/repositoryFile:RepositoryFile gitignore example/.gitignore
        ```
         To import a file from a branch other than the default branch, append `:` and the branch name, e.g.

        ```sh
         $ pulumi import github:index/repositoryFile:RepositoryFile gitignore example/.gitignore:dev
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] branch: Git branch (defaults to the repository's default branch).
               The branch must already exist, it will not be created if it does not already exist.
        :param pulumi.Input[str] commit_author: Committer author name to use. **NOTE:** GitHub app users may omit author and email information so GitHub can verify commits as the GitHub App. This maybe useful when a branch protection rule requires signed commits.
        :param pulumi.Input[str] commit_email: Committer email address to use. **NOTE:** GitHub app users may omit author and email information so GitHub can verify commits as the GitHub App. This may be useful when a branch protection rule requires signed commits.
        :param pulumi.Input[str] commit_message: Commit message when adding or updating the managed file.
        :param pulumi.Input[str] content: The file content.
        :param pulumi.Input[str] file: The path of the file to manage.
        :param pulumi.Input[bool] overwrite_on_create: Enable overwriting existing files
        :param pulumi.Input[str] repository: The repository to create the file in.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RepositoryFileArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This resource allows you to create and manage files within a
        GitHub repository.

        ## Import

        Repository files can be imported using a combination of the `repo` and `file`, e.g.

        ```sh
         $ pulumi import github:index/repositoryFile:RepositoryFile gitignore example/.gitignore
        ```
         To import a file from a branch other than the default branch, append `:` and the branch name, e.g.

        ```sh
         $ pulumi import github:index/repositoryFile:RepositoryFile gitignore example/.gitignore:dev
        ```

        :param str resource_name: The name of the resource.
        :param RepositoryFileArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RepositoryFileArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            RepositoryFileArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 branch: Optional[pulumi.Input[str]] = None,
                 commit_author: Optional[pulumi.Input[str]] = None,
                 commit_email: Optional[pulumi.Input[str]] = None,
                 commit_message: Optional[pulumi.Input[str]] = None,
                 content: Optional[pulumi.Input[str]] = None,
                 file: Optional[pulumi.Input[str]] = None,
                 overwrite_on_create: Optional[pulumi.Input[bool]] = None,
                 repository: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RepositoryFileArgs.__new__(RepositoryFileArgs)

            __props__.__dict__["branch"] = branch
            __props__.__dict__["commit_author"] = commit_author
            __props__.__dict__["commit_email"] = commit_email
            __props__.__dict__["commit_message"] = commit_message
            if content is None and not opts.urn:
                raise TypeError("Missing required property 'content'")
            __props__.__dict__["content"] = content
            if file is None and not opts.urn:
                raise TypeError("Missing required property 'file'")
            __props__.__dict__["file"] = file
            __props__.__dict__["overwrite_on_create"] = overwrite_on_create
            if repository is None and not opts.urn:
                raise TypeError("Missing required property 'repository'")
            __props__.__dict__["repository"] = repository
            __props__.__dict__["commit_sha"] = None
            __props__.__dict__["ref"] = None
            __props__.__dict__["sha"] = None
        super(RepositoryFile, __self__).__init__(
            'github:index/repositoryFile:RepositoryFile',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            branch: Optional[pulumi.Input[str]] = None,
            commit_author: Optional[pulumi.Input[str]] = None,
            commit_email: Optional[pulumi.Input[str]] = None,
            commit_message: Optional[pulumi.Input[str]] = None,
            commit_sha: Optional[pulumi.Input[str]] = None,
            content: Optional[pulumi.Input[str]] = None,
            file: Optional[pulumi.Input[str]] = None,
            overwrite_on_create: Optional[pulumi.Input[bool]] = None,
            ref: Optional[pulumi.Input[str]] = None,
            repository: Optional[pulumi.Input[str]] = None,
            sha: Optional[pulumi.Input[str]] = None) -> 'RepositoryFile':
        """
        Get an existing RepositoryFile resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] branch: Git branch (defaults to the repository's default branch).
               The branch must already exist, it will not be created if it does not already exist.
        :param pulumi.Input[str] commit_author: Committer author name to use. **NOTE:** GitHub app users may omit author and email information so GitHub can verify commits as the GitHub App. This maybe useful when a branch protection rule requires signed commits.
        :param pulumi.Input[str] commit_email: Committer email address to use. **NOTE:** GitHub app users may omit author and email information so GitHub can verify commits as the GitHub App. This may be useful when a branch protection rule requires signed commits.
        :param pulumi.Input[str] commit_message: Commit message when adding or updating the managed file.
        :param pulumi.Input[str] commit_sha: The SHA of the commit that modified the file.
        :param pulumi.Input[str] content: The file content.
        :param pulumi.Input[str] file: The path of the file to manage.
        :param pulumi.Input[bool] overwrite_on_create: Enable overwriting existing files
        :param pulumi.Input[str] ref: The name of the commit/branch/tag.
        :param pulumi.Input[str] repository: The repository to create the file in.
        :param pulumi.Input[str] sha: The SHA blob of the file.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RepositoryFileState.__new__(_RepositoryFileState)

        __props__.__dict__["branch"] = branch
        __props__.__dict__["commit_author"] = commit_author
        __props__.__dict__["commit_email"] = commit_email
        __props__.__dict__["commit_message"] = commit_message
        __props__.__dict__["commit_sha"] = commit_sha
        __props__.__dict__["content"] = content
        __props__.__dict__["file"] = file
        __props__.__dict__["overwrite_on_create"] = overwrite_on_create
        __props__.__dict__["ref"] = ref
        __props__.__dict__["repository"] = repository
        __props__.__dict__["sha"] = sha
        return RepositoryFile(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def branch(self) -> pulumi.Output[Optional[str]]:
        """
        Git branch (defaults to the repository's default branch).
        The branch must already exist, it will not be created if it does not already exist.
        """
        return pulumi.get(self, "branch")

    @property
    @pulumi.getter(name="commitAuthor")
    def commit_author(self) -> pulumi.Output[Optional[str]]:
        """
        Committer author name to use. **NOTE:** GitHub app users may omit author and email information so GitHub can verify commits as the GitHub App. This maybe useful when a branch protection rule requires signed commits.
        """
        return pulumi.get(self, "commit_author")

    @property
    @pulumi.getter(name="commitEmail")
    def commit_email(self) -> pulumi.Output[Optional[str]]:
        """
        Committer email address to use. **NOTE:** GitHub app users may omit author and email information so GitHub can verify commits as the GitHub App. This may be useful when a branch protection rule requires signed commits.
        """
        return pulumi.get(self, "commit_email")

    @property
    @pulumi.getter(name="commitMessage")
    def commit_message(self) -> pulumi.Output[str]:
        """
        Commit message when adding or updating the managed file.
        """
        return pulumi.get(self, "commit_message")

    @property
    @pulumi.getter(name="commitSha")
    def commit_sha(self) -> pulumi.Output[str]:
        """
        The SHA of the commit that modified the file.
        """
        return pulumi.get(self, "commit_sha")

    @property
    @pulumi.getter
    def content(self) -> pulumi.Output[str]:
        """
        The file content.
        """
        return pulumi.get(self, "content")

    @property
    @pulumi.getter
    def file(self) -> pulumi.Output[str]:
        """
        The path of the file to manage.
        """
        return pulumi.get(self, "file")

    @property
    @pulumi.getter(name="overwriteOnCreate")
    def overwrite_on_create(self) -> pulumi.Output[Optional[bool]]:
        """
        Enable overwriting existing files
        """
        return pulumi.get(self, "overwrite_on_create")

    @property
    @pulumi.getter
    def ref(self) -> pulumi.Output[str]:
        """
        The name of the commit/branch/tag.
        """
        return pulumi.get(self, "ref")

    @property
    @pulumi.getter
    def repository(self) -> pulumi.Output[str]:
        """
        The repository to create the file in.
        """
        return pulumi.get(self, "repository")

    @property
    @pulumi.getter
    def sha(self) -> pulumi.Output[str]:
        """
        The SHA blob of the file.
        """
        return pulumi.get(self, "sha")

