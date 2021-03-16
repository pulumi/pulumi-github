# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from . import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['Repository']


class Repository(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allow_merge_commit: Optional[pulumi.Input[bool]] = None,
                 allow_rebase_merge: Optional[pulumi.Input[bool]] = None,
                 allow_squash_merge: Optional[pulumi.Input[bool]] = None,
                 archive_on_destroy: Optional[pulumi.Input[bool]] = None,
                 archived: Optional[pulumi.Input[bool]] = None,
                 auto_init: Optional[pulumi.Input[bool]] = None,
                 default_branch: Optional[pulumi.Input[str]] = None,
                 delete_branch_on_merge: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 gitignore_template: Optional[pulumi.Input[str]] = None,
                 has_downloads: Optional[pulumi.Input[bool]] = None,
                 has_issues: Optional[pulumi.Input[bool]] = None,
                 has_projects: Optional[pulumi.Input[bool]] = None,
                 has_wiki: Optional[pulumi.Input[bool]] = None,
                 homepage_url: Optional[pulumi.Input[str]] = None,
                 is_template: Optional[pulumi.Input[bool]] = None,
                 license_template: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 pages: Optional[pulumi.Input[pulumi.InputType['RepositoryPagesArgs']]] = None,
                 private: Optional[pulumi.Input[bool]] = None,
                 template: Optional[pulumi.Input[pulumi.InputType['RepositoryTemplateArgs']]] = None,
                 topics: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 visibility: Optional[pulumi.Input[str]] = None,
                 vulnerability_alerts: Optional[pulumi.Input[bool]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        This resource allows you to create and manage repositories within your
        GitHub organization or personal account.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_github as github

        example = github.Repository("example",
            description="My awesome codebase",
            template=github.RepositoryTemplateArgs(
                owner="github",
                repository="terraform-module-template",
            ),
            visibility="public")
        ```
        ### With Github Pages Enabled

        ```python
        import pulumi
        import pulumi_github as github

        example = github.Repository("example",
            description="My awesome web page",
            pages=github.RepositoryPagesArgs(
                source=github.RepositoryPagesSourceArgs(
                    branch="master",
                    path="/docs",
                ),
            ),
            private=False)
        ```

        ## Import

        Repositories can be imported using the `name`, e.g.

        ```sh
         $ pulumi import github:index/repository:Repository terraform terraform
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] allow_merge_commit: Set to `false` to disable merge commits on the repository.
        :param pulumi.Input[bool] allow_rebase_merge: Set to `false` to disable rebase merges on the repository.
        :param pulumi.Input[bool] allow_squash_merge: Set to `false` to disable squash merges on the repository.
        :param pulumi.Input[bool] archive_on_destroy: Set to `true` to archive the repository instead of deleting on destroy.
        :param pulumi.Input[bool] archived: Specifies if the repository should be archived. Defaults to `false`. **NOTE** Currently, the API does not support unarchiving.
        :param pulumi.Input[bool] auto_init: Set to `true` to produce an initial commit in the repository.
        :param pulumi.Input[str] default_branch: (Deprecated: Use `BranchDefault` resource instead) The name of the default branch of the repository. **NOTE:** This can only be set after a repository has already been created,
               and after a correct reference has been created for the target branch inside the repository. This means a user will have to omit this parameter from the
               initial repository creation and create the target branch inside of the repository prior to setting this attribute.
        :param pulumi.Input[bool] delete_branch_on_merge: Automatically delete head branch after a pull request is merged. Defaults to `false`.
        :param pulumi.Input[str] description: A description of the repository.
        :param pulumi.Input[str] gitignore_template: Use the [name of the template](https://github.com/github/gitignore) without the extension. For example, "Haskell".
        :param pulumi.Input[bool] has_downloads: Set to `true` to enable the (deprecated) downloads features on the repository.
        :param pulumi.Input[bool] has_issues: Set to `true` to enable the GitHub Issues features
               on the repository.
        :param pulumi.Input[bool] has_projects: Set to `true` to enable the GitHub Projects features on the repository. Per the GitHub [documentation](https://developer.github.com/v3/repos/#create) when in an organization that has disabled repository projects it will default to `false` and will otherwise default to `true`. If you specify `true` when it has been disabled it will return an error.
        :param pulumi.Input[bool] has_wiki: Set to `true` to enable the GitHub Wiki features on
               the repository.
        :param pulumi.Input[str] homepage_url: URL of a page describing the project.
        :param pulumi.Input[bool] is_template: Set to `true` to tell GitHub that this is a template repository.
        :param pulumi.Input[str] license_template: Use the [name of the template](https://github.com/github/choosealicense.com/tree/gh-pages/_licenses) without the extension. For example, "mit" or "mpl-2.0".
        :param pulumi.Input[str] name: The name of the repository.
        :param pulumi.Input[pulumi.InputType['RepositoryPagesArgs']] pages: The repository's Github Pages configuration. See Github Pages Configuration below for details.
        :param pulumi.Input[bool] private: Set to `true` to create a private repository.
               Repositories are created as public (e.g. open source) by default.
        :param pulumi.Input[pulumi.InputType['RepositoryTemplateArgs']] template: Use a template repository to create this resource. See Template Repositories below for details.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] topics: The list of topics of the repository.
        :param pulumi.Input[str] visibility: Can be `public` or `private`. If your organization is associated with an enterprise account using GitHub Enterprise Cloud or GitHub Enterprise Server 2.20+, visibility can also be `internal`. The `visibility` parameter overrides the `private` parameter.
        :param pulumi.Input[bool] vulnerability_alerts: Set to `true` to enable security alerts for vulnerable dependencies. Enabling requires alerts to be enabled on the owner level. (Note for importing: GitHub enables the alerts on public repos but disables them on private repos by default.) See [GitHub Documentation](https://help.github.com/en/github/managing-security-vulnerabilities/about-security-alerts-for-vulnerable-dependencies) for details.
        """
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

            __props__['allow_merge_commit'] = allow_merge_commit
            __props__['allow_rebase_merge'] = allow_rebase_merge
            __props__['allow_squash_merge'] = allow_squash_merge
            __props__['archive_on_destroy'] = archive_on_destroy
            __props__['archived'] = archived
            __props__['auto_init'] = auto_init
            if default_branch is not None and not opts.urn:
                warnings.warn("""Use the github_branch_default resource instead""", DeprecationWarning)
                pulumi.log.warn("""default_branch is deprecated: Use the github_branch_default resource instead""")
            __props__['default_branch'] = default_branch
            __props__['delete_branch_on_merge'] = delete_branch_on_merge
            __props__['description'] = description
            __props__['gitignore_template'] = gitignore_template
            __props__['has_downloads'] = has_downloads
            __props__['has_issues'] = has_issues
            __props__['has_projects'] = has_projects
            __props__['has_wiki'] = has_wiki
            __props__['homepage_url'] = homepage_url
            __props__['is_template'] = is_template
            __props__['license_template'] = license_template
            __props__['name'] = name
            __props__['pages'] = pages
            if private is not None and not opts.urn:
                warnings.warn("""use visibility instead""", DeprecationWarning)
                pulumi.log.warn("""private is deprecated: use visibility instead""")
            __props__['private'] = private
            __props__['template'] = template
            __props__['topics'] = topics
            __props__['visibility'] = visibility
            __props__['vulnerability_alerts'] = vulnerability_alerts
            __props__['etag'] = None
            __props__['full_name'] = None
            __props__['git_clone_url'] = None
            __props__['html_url'] = None
            __props__['http_clone_url'] = None
            __props__['node_id'] = None
            __props__['repo_id'] = None
            __props__['ssh_clone_url'] = None
            __props__['svn_url'] = None
        super(Repository, __self__).__init__(
            'github:index/repository:Repository',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            allow_merge_commit: Optional[pulumi.Input[bool]] = None,
            allow_rebase_merge: Optional[pulumi.Input[bool]] = None,
            allow_squash_merge: Optional[pulumi.Input[bool]] = None,
            archive_on_destroy: Optional[pulumi.Input[bool]] = None,
            archived: Optional[pulumi.Input[bool]] = None,
            auto_init: Optional[pulumi.Input[bool]] = None,
            default_branch: Optional[pulumi.Input[str]] = None,
            delete_branch_on_merge: Optional[pulumi.Input[bool]] = None,
            description: Optional[pulumi.Input[str]] = None,
            etag: Optional[pulumi.Input[str]] = None,
            full_name: Optional[pulumi.Input[str]] = None,
            git_clone_url: Optional[pulumi.Input[str]] = None,
            gitignore_template: Optional[pulumi.Input[str]] = None,
            has_downloads: Optional[pulumi.Input[bool]] = None,
            has_issues: Optional[pulumi.Input[bool]] = None,
            has_projects: Optional[pulumi.Input[bool]] = None,
            has_wiki: Optional[pulumi.Input[bool]] = None,
            homepage_url: Optional[pulumi.Input[str]] = None,
            html_url: Optional[pulumi.Input[str]] = None,
            http_clone_url: Optional[pulumi.Input[str]] = None,
            is_template: Optional[pulumi.Input[bool]] = None,
            license_template: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            node_id: Optional[pulumi.Input[str]] = None,
            pages: Optional[pulumi.Input[pulumi.InputType['RepositoryPagesArgs']]] = None,
            private: Optional[pulumi.Input[bool]] = None,
            repo_id: Optional[pulumi.Input[int]] = None,
            ssh_clone_url: Optional[pulumi.Input[str]] = None,
            svn_url: Optional[pulumi.Input[str]] = None,
            template: Optional[pulumi.Input[pulumi.InputType['RepositoryTemplateArgs']]] = None,
            topics: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            visibility: Optional[pulumi.Input[str]] = None,
            vulnerability_alerts: Optional[pulumi.Input[bool]] = None) -> 'Repository':
        """
        Get an existing Repository resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] allow_merge_commit: Set to `false` to disable merge commits on the repository.
        :param pulumi.Input[bool] allow_rebase_merge: Set to `false` to disable rebase merges on the repository.
        :param pulumi.Input[bool] allow_squash_merge: Set to `false` to disable squash merges on the repository.
        :param pulumi.Input[bool] archive_on_destroy: Set to `true` to archive the repository instead of deleting on destroy.
        :param pulumi.Input[bool] archived: Specifies if the repository should be archived. Defaults to `false`. **NOTE** Currently, the API does not support unarchiving.
        :param pulumi.Input[bool] auto_init: Set to `true` to produce an initial commit in the repository.
        :param pulumi.Input[str] default_branch: (Deprecated: Use `BranchDefault` resource instead) The name of the default branch of the repository. **NOTE:** This can only be set after a repository has already been created,
               and after a correct reference has been created for the target branch inside the repository. This means a user will have to omit this parameter from the
               initial repository creation and create the target branch inside of the repository prior to setting this attribute.
        :param pulumi.Input[bool] delete_branch_on_merge: Automatically delete head branch after a pull request is merged. Defaults to `false`.
        :param pulumi.Input[str] description: A description of the repository.
        :param pulumi.Input[str] full_name: A string of the form "orgname/reponame".
        :param pulumi.Input[str] git_clone_url: URL that can be provided to `git clone` to clone the repository anonymously via the git protocol.
        :param pulumi.Input[str] gitignore_template: Use the [name of the template](https://github.com/github/gitignore) without the extension. For example, "Haskell".
        :param pulumi.Input[bool] has_downloads: Set to `true` to enable the (deprecated) downloads features on the repository.
        :param pulumi.Input[bool] has_issues: Set to `true` to enable the GitHub Issues features
               on the repository.
        :param pulumi.Input[bool] has_projects: Set to `true` to enable the GitHub Projects features on the repository. Per the GitHub [documentation](https://developer.github.com/v3/repos/#create) when in an organization that has disabled repository projects it will default to `false` and will otherwise default to `true`. If you specify `true` when it has been disabled it will return an error.
        :param pulumi.Input[bool] has_wiki: Set to `true` to enable the GitHub Wiki features on
               the repository.
        :param pulumi.Input[str] homepage_url: URL of a page describing the project.
        :param pulumi.Input[str] html_url: The absolute URL (including scheme) of the rendered Github Pages site e.g. `https://username.github.io`.
        :param pulumi.Input[str] http_clone_url: URL that can be provided to `git clone` to clone the repository via HTTPS.
        :param pulumi.Input[bool] is_template: Set to `true` to tell GitHub that this is a template repository.
        :param pulumi.Input[str] license_template: Use the [name of the template](https://github.com/github/choosealicense.com/tree/gh-pages/_licenses) without the extension. For example, "mit" or "mpl-2.0".
        :param pulumi.Input[str] name: The name of the repository.
        :param pulumi.Input[str] node_id: GraphQL global node id for use with v4 API
        :param pulumi.Input[pulumi.InputType['RepositoryPagesArgs']] pages: The repository's Github Pages configuration. See Github Pages Configuration below for details.
        :param pulumi.Input[bool] private: Set to `true` to create a private repository.
               Repositories are created as public (e.g. open source) by default.
        :param pulumi.Input[int] repo_id: Github ID for the repository
        :param pulumi.Input[str] ssh_clone_url: URL that can be provided to `git clone` to clone the repository via SSH.
        :param pulumi.Input[str] svn_url: URL that can be provided to `svn checkout` to check out the repository via GitHub's Subversion protocol emulation.
        :param pulumi.Input[pulumi.InputType['RepositoryTemplateArgs']] template: Use a template repository to create this resource. See Template Repositories below for details.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] topics: The list of topics of the repository.
        :param pulumi.Input[str] visibility: Can be `public` or `private`. If your organization is associated with an enterprise account using GitHub Enterprise Cloud or GitHub Enterprise Server 2.20+, visibility can also be `internal`. The `visibility` parameter overrides the `private` parameter.
        :param pulumi.Input[bool] vulnerability_alerts: Set to `true` to enable security alerts for vulnerable dependencies. Enabling requires alerts to be enabled on the owner level. (Note for importing: GitHub enables the alerts on public repos but disables them on private repos by default.) See [GitHub Documentation](https://help.github.com/en/github/managing-security-vulnerabilities/about-security-alerts-for-vulnerable-dependencies) for details.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["allow_merge_commit"] = allow_merge_commit
        __props__["allow_rebase_merge"] = allow_rebase_merge
        __props__["allow_squash_merge"] = allow_squash_merge
        __props__["archive_on_destroy"] = archive_on_destroy
        __props__["archived"] = archived
        __props__["auto_init"] = auto_init
        __props__["default_branch"] = default_branch
        __props__["delete_branch_on_merge"] = delete_branch_on_merge
        __props__["description"] = description
        __props__["etag"] = etag
        __props__["full_name"] = full_name
        __props__["git_clone_url"] = git_clone_url
        __props__["gitignore_template"] = gitignore_template
        __props__["has_downloads"] = has_downloads
        __props__["has_issues"] = has_issues
        __props__["has_projects"] = has_projects
        __props__["has_wiki"] = has_wiki
        __props__["homepage_url"] = homepage_url
        __props__["html_url"] = html_url
        __props__["http_clone_url"] = http_clone_url
        __props__["is_template"] = is_template
        __props__["license_template"] = license_template
        __props__["name"] = name
        __props__["node_id"] = node_id
        __props__["pages"] = pages
        __props__["private"] = private
        __props__["repo_id"] = repo_id
        __props__["ssh_clone_url"] = ssh_clone_url
        __props__["svn_url"] = svn_url
        __props__["template"] = template
        __props__["topics"] = topics
        __props__["visibility"] = visibility
        __props__["vulnerability_alerts"] = vulnerability_alerts
        return Repository(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="allowMergeCommit")
    def allow_merge_commit(self) -> pulumi.Output[Optional[bool]]:
        """
        Set to `false` to disable merge commits on the repository.
        """
        return pulumi.get(self, "allow_merge_commit")

    @property
    @pulumi.getter(name="allowRebaseMerge")
    def allow_rebase_merge(self) -> pulumi.Output[Optional[bool]]:
        """
        Set to `false` to disable rebase merges on the repository.
        """
        return pulumi.get(self, "allow_rebase_merge")

    @property
    @pulumi.getter(name="allowSquashMerge")
    def allow_squash_merge(self) -> pulumi.Output[Optional[bool]]:
        """
        Set to `false` to disable squash merges on the repository.
        """
        return pulumi.get(self, "allow_squash_merge")

    @property
    @pulumi.getter(name="archiveOnDestroy")
    def archive_on_destroy(self) -> pulumi.Output[Optional[bool]]:
        """
        Set to `true` to archive the repository instead of deleting on destroy.
        """
        return pulumi.get(self, "archive_on_destroy")

    @property
    @pulumi.getter
    def archived(self) -> pulumi.Output[Optional[bool]]:
        """
        Specifies if the repository should be archived. Defaults to `false`. **NOTE** Currently, the API does not support unarchiving.
        """
        return pulumi.get(self, "archived")

    @property
    @pulumi.getter(name="autoInit")
    def auto_init(self) -> pulumi.Output[Optional[bool]]:
        """
        Set to `true` to produce an initial commit in the repository.
        """
        return pulumi.get(self, "auto_init")

    @property
    @pulumi.getter(name="defaultBranch")
    def default_branch(self) -> pulumi.Output[str]:
        """
        (Deprecated: Use `BranchDefault` resource instead) The name of the default branch of the repository. **NOTE:** This can only be set after a repository has already been created,
        and after a correct reference has been created for the target branch inside the repository. This means a user will have to omit this parameter from the
        initial repository creation and create the target branch inside of the repository prior to setting this attribute.
        """
        return pulumi.get(self, "default_branch")

    @property
    @pulumi.getter(name="deleteBranchOnMerge")
    def delete_branch_on_merge(self) -> pulumi.Output[Optional[bool]]:
        """
        Automatically delete head branch after a pull request is merged. Defaults to `false`.
        """
        return pulumi.get(self, "delete_branch_on_merge")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A description of the repository.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="fullName")
    def full_name(self) -> pulumi.Output[str]:
        """
        A string of the form "orgname/reponame".
        """
        return pulumi.get(self, "full_name")

    @property
    @pulumi.getter(name="gitCloneUrl")
    def git_clone_url(self) -> pulumi.Output[str]:
        """
        URL that can be provided to `git clone` to clone the repository anonymously via the git protocol.
        """
        return pulumi.get(self, "git_clone_url")

    @property
    @pulumi.getter(name="gitignoreTemplate")
    def gitignore_template(self) -> pulumi.Output[Optional[str]]:
        """
        Use the [name of the template](https://github.com/github/gitignore) without the extension. For example, "Haskell".
        """
        return pulumi.get(self, "gitignore_template")

    @property
    @pulumi.getter(name="hasDownloads")
    def has_downloads(self) -> pulumi.Output[Optional[bool]]:
        """
        Set to `true` to enable the (deprecated) downloads features on the repository.
        """
        return pulumi.get(self, "has_downloads")

    @property
    @pulumi.getter(name="hasIssues")
    def has_issues(self) -> pulumi.Output[Optional[bool]]:
        """
        Set to `true` to enable the GitHub Issues features
        on the repository.
        """
        return pulumi.get(self, "has_issues")

    @property
    @pulumi.getter(name="hasProjects")
    def has_projects(self) -> pulumi.Output[Optional[bool]]:
        """
        Set to `true` to enable the GitHub Projects features on the repository. Per the GitHub [documentation](https://developer.github.com/v3/repos/#create) when in an organization that has disabled repository projects it will default to `false` and will otherwise default to `true`. If you specify `true` when it has been disabled it will return an error.
        """
        return pulumi.get(self, "has_projects")

    @property
    @pulumi.getter(name="hasWiki")
    def has_wiki(self) -> pulumi.Output[Optional[bool]]:
        """
        Set to `true` to enable the GitHub Wiki features on
        the repository.
        """
        return pulumi.get(self, "has_wiki")

    @property
    @pulumi.getter(name="homepageUrl")
    def homepage_url(self) -> pulumi.Output[Optional[str]]:
        """
        URL of a page describing the project.
        """
        return pulumi.get(self, "homepage_url")

    @property
    @pulumi.getter(name="htmlUrl")
    def html_url(self) -> pulumi.Output[str]:
        """
        The absolute URL (including scheme) of the rendered Github Pages site e.g. `https://username.github.io`.
        """
        return pulumi.get(self, "html_url")

    @property
    @pulumi.getter(name="httpCloneUrl")
    def http_clone_url(self) -> pulumi.Output[str]:
        """
        URL that can be provided to `git clone` to clone the repository via HTTPS.
        """
        return pulumi.get(self, "http_clone_url")

    @property
    @pulumi.getter(name="isTemplate")
    def is_template(self) -> pulumi.Output[Optional[bool]]:
        """
        Set to `true` to tell GitHub that this is a template repository.
        """
        return pulumi.get(self, "is_template")

    @property
    @pulumi.getter(name="licenseTemplate")
    def license_template(self) -> pulumi.Output[Optional[str]]:
        """
        Use the [name of the template](https://github.com/github/choosealicense.com/tree/gh-pages/_licenses) without the extension. For example, "mit" or "mpl-2.0".
        """
        return pulumi.get(self, "license_template")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the repository.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nodeId")
    def node_id(self) -> pulumi.Output[str]:
        """
        GraphQL global node id for use with v4 API
        """
        return pulumi.get(self, "node_id")

    @property
    @pulumi.getter
    def pages(self) -> pulumi.Output[Optional['outputs.RepositoryPages']]:
        """
        The repository's Github Pages configuration. See Github Pages Configuration below for details.
        """
        return pulumi.get(self, "pages")

    @property
    @pulumi.getter
    def private(self) -> pulumi.Output[bool]:
        """
        Set to `true` to create a private repository.
        Repositories are created as public (e.g. open source) by default.
        """
        return pulumi.get(self, "private")

    @property
    @pulumi.getter(name="repoId")
    def repo_id(self) -> pulumi.Output[int]:
        """
        Github ID for the repository
        """
        return pulumi.get(self, "repo_id")

    @property
    @pulumi.getter(name="sshCloneUrl")
    def ssh_clone_url(self) -> pulumi.Output[str]:
        """
        URL that can be provided to `git clone` to clone the repository via SSH.
        """
        return pulumi.get(self, "ssh_clone_url")

    @property
    @pulumi.getter(name="svnUrl")
    def svn_url(self) -> pulumi.Output[str]:
        """
        URL that can be provided to `svn checkout` to check out the repository via GitHub's Subversion protocol emulation.
        """
        return pulumi.get(self, "svn_url")

    @property
    @pulumi.getter
    def template(self) -> pulumi.Output[Optional['outputs.RepositoryTemplate']]:
        """
        Use a template repository to create this resource. See Template Repositories below for details.
        """
        return pulumi.get(self, "template")

    @property
    @pulumi.getter
    def topics(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The list of topics of the repository.
        """
        return pulumi.get(self, "topics")

    @property
    @pulumi.getter
    def visibility(self) -> pulumi.Output[str]:
        """
        Can be `public` or `private`. If your organization is associated with an enterprise account using GitHub Enterprise Cloud or GitHub Enterprise Server 2.20+, visibility can also be `internal`. The `visibility` parameter overrides the `private` parameter.
        """
        return pulumi.get(self, "visibility")

    @property
    @pulumi.getter(name="vulnerabilityAlerts")
    def vulnerability_alerts(self) -> pulumi.Output[Optional[bool]]:
        """
        Set to `true` to enable security alerts for vulnerable dependencies. Enabling requires alerts to be enabled on the owner level. (Note for importing: GitHub enables the alerts on public repos but disables them on private repos by default.) See [GitHub Documentation](https://help.github.com/en/github/managing-security-vulnerabilities/about-security-alerts-for-vulnerable-dependencies) for details.
        """
        return pulumi.get(self, "vulnerability_alerts")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

