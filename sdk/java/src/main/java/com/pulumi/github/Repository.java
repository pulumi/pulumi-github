// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.github.RepositoryArgs;
import com.pulumi.github.Utilities;
import com.pulumi.github.inputs.RepositoryState;
import com.pulumi.github.outputs.RepositoryPages;
import com.pulumi.github.outputs.RepositorySecurityAndAnalysis;
import com.pulumi.github.outputs.RepositoryTemplate;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * This resource allows you to create and manage repositories within your
 * GitHub organization or personal account.
 * 
 * &gt; Note: When used with GitHub App authentication, even GET requests must have
 * the `contents:write` permission or else the `allow_merge_commit`, `allow_rebase_merge`,
 * and `allow_squash_merge` attributes will be ignored, causing confusing diffs.
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.github.Repository;
 * import com.pulumi.github.RepositoryArgs;
 * import com.pulumi.github.inputs.RepositoryTemplateArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var example = new Repository(&#34;example&#34;, RepositoryArgs.builder()        
 *             .description(&#34;My awesome codebase&#34;)
 *             .template(RepositoryTemplateArgs.builder()
 *                 .includeAllBranches(true)
 *                 .owner(&#34;github&#34;)
 *                 .repository(&#34;terraform-template-module&#34;)
 *                 .build())
 *             .visibility(&#34;public&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ### With GitHub Pages Enabled
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.github.Repository;
 * import com.pulumi.github.RepositoryArgs;
 * import com.pulumi.github.inputs.RepositoryPagesArgs;
 * import com.pulumi.github.inputs.RepositoryPagesSourceArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var example = new Repository(&#34;example&#34;, RepositoryArgs.builder()        
 *             .description(&#34;My awesome web page&#34;)
 *             .pages(RepositoryPagesArgs.builder()
 *                 .source(RepositoryPagesSourceArgs.builder()
 *                     .branch(&#34;master&#34;)
 *                     .path(&#34;/docs&#34;)
 *                     .build())
 *                 .build())
 *             .private_(false)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Repositories can be imported using the `name`, e.g.
 * 
 * ```sh
 * $ pulumi import github:index/repository:Repository terraform terraform
 * ```
 * 
 */
@ResourceType(type="github:index/repository:Repository")
public class Repository extends com.pulumi.resources.CustomResource {
    /**
     * Set to `true` to allow auto-merging pull requests on the repository.
     * 
     */
    @Export(name="allowAutoMerge", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> allowAutoMerge;

    /**
     * @return Set to `true` to allow auto-merging pull requests on the repository.
     * 
     */
    public Output<Optional<Boolean>> allowAutoMerge() {
        return Codegen.optional(this.allowAutoMerge);
    }
    /**
     * Set to `false` to disable merge commits on the repository.
     * 
     */
    @Export(name="allowMergeCommit", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> allowMergeCommit;

    /**
     * @return Set to `false` to disable merge commits on the repository.
     * 
     */
    public Output<Optional<Boolean>> allowMergeCommit() {
        return Codegen.optional(this.allowMergeCommit);
    }
    /**
     * Set to `false` to disable rebase merges on the repository.
     * 
     */
    @Export(name="allowRebaseMerge", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> allowRebaseMerge;

    /**
     * @return Set to `false` to disable rebase merges on the repository.
     * 
     */
    public Output<Optional<Boolean>> allowRebaseMerge() {
        return Codegen.optional(this.allowRebaseMerge);
    }
    /**
     * Set to `false` to disable squash merges on the repository.
     * 
     */
    @Export(name="allowSquashMerge", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> allowSquashMerge;

    /**
     * @return Set to `false` to disable squash merges on the repository.
     * 
     */
    public Output<Optional<Boolean>> allowSquashMerge() {
        return Codegen.optional(this.allowSquashMerge);
    }
    /**
     * Set to `true` to always suggest updating pull request branches.
     * 
     */
    @Export(name="allowUpdateBranch", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> allowUpdateBranch;

    /**
     * @return Set to `true` to always suggest updating pull request branches.
     * 
     */
    public Output<Optional<Boolean>> allowUpdateBranch() {
        return Codegen.optional(this.allowUpdateBranch);
    }
    /**
     * Set to `true` to archive the repository instead of deleting on destroy.
     * 
     */
    @Export(name="archiveOnDestroy", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> archiveOnDestroy;

    /**
     * @return Set to `true` to archive the repository instead of deleting on destroy.
     * 
     */
    public Output<Optional<Boolean>> archiveOnDestroy() {
        return Codegen.optional(this.archiveOnDestroy);
    }
    /**
     * Specifies if the repository should be archived. Defaults to `false`. **NOTE** Currently, the API does not support unarchiving.
     * 
     */
    @Export(name="archived", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> archived;

    /**
     * @return Specifies if the repository should be archived. Defaults to `false`. **NOTE** Currently, the API does not support unarchiving.
     * 
     */
    public Output<Optional<Boolean>> archived() {
        return Codegen.optional(this.archived);
    }
    /**
     * Set to `true` to produce an initial commit in the repository.
     * 
     */
    @Export(name="autoInit", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> autoInit;

    /**
     * @return Set to `true` to produce an initial commit in the repository.
     * 
     */
    public Output<Optional<Boolean>> autoInit() {
        return Codegen.optional(this.autoInit);
    }
    /**
     * (Deprecated: Use `github.BranchDefault` resource instead) The name of the default branch of the repository. **NOTE:** This can only be set after a repository has already been created,
     * and after a correct reference has been created for the target branch inside the repository. This means a user will have to omit this parameter from the
     * initial repository creation and create the target branch inside of the repository prior to setting this attribute.
     * 
     * @deprecated
     * Use the github.BranchDefault resource instead
     * 
     */
    @Deprecated /* Use the github.BranchDefault resource instead */
    @Export(name="defaultBranch", refs={String.class}, tree="[0]")
    private Output<String> defaultBranch;

    /**
     * @return (Deprecated: Use `github.BranchDefault` resource instead) The name of the default branch of the repository. **NOTE:** This can only be set after a repository has already been created,
     * and after a correct reference has been created for the target branch inside the repository. This means a user will have to omit this parameter from the
     * initial repository creation and create the target branch inside of the repository prior to setting this attribute.
     * 
     */
    public Output<String> defaultBranch() {
        return this.defaultBranch;
    }
    /**
     * Automatically delete head branch after a pull request is merged. Defaults to `false`.
     * 
     */
    @Export(name="deleteBranchOnMerge", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> deleteBranchOnMerge;

    /**
     * @return Automatically delete head branch after a pull request is merged. Defaults to `false`.
     * 
     */
    public Output<Optional<Boolean>> deleteBranchOnMerge() {
        return Codegen.optional(this.deleteBranchOnMerge);
    }
    /**
     * A description of the repository.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return A description of the repository.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    @Export(name="etag", refs={String.class}, tree="[0]")
    private Output<String> etag;

    public Output<String> etag() {
        return this.etag;
    }
    /**
     * A string of the form &#34;orgname/reponame&#34;.
     * 
     */
    @Export(name="fullName", refs={String.class}, tree="[0]")
    private Output<String> fullName;

    /**
     * @return A string of the form &#34;orgname/reponame&#34;.
     * 
     */
    public Output<String> fullName() {
        return this.fullName;
    }
    /**
     * URL that can be provided to `git clone` to clone the repository anonymously via the git protocol.
     * 
     */
    @Export(name="gitCloneUrl", refs={String.class}, tree="[0]")
    private Output<String> gitCloneUrl;

    /**
     * @return URL that can be provided to `git clone` to clone the repository anonymously via the git protocol.
     * 
     */
    public Output<String> gitCloneUrl() {
        return this.gitCloneUrl;
    }
    /**
     * Use the [name of the template](https://github.com/github/gitignore) without the extension. For example, &#34;Haskell&#34;.
     * 
     */
    @Export(name="gitignoreTemplate", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> gitignoreTemplate;

    /**
     * @return Use the [name of the template](https://github.com/github/gitignore) without the extension. For example, &#34;Haskell&#34;.
     * 
     */
    public Output<Optional<String>> gitignoreTemplate() {
        return Codegen.optional(this.gitignoreTemplate);
    }
    /**
     * Set to `true` to enable GitHub Discussions on the repository. Defaults to `false`.
     * 
     */
    @Export(name="hasDiscussions", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> hasDiscussions;

    /**
     * @return Set to `true` to enable GitHub Discussions on the repository. Defaults to `false`.
     * 
     */
    public Output<Optional<Boolean>> hasDiscussions() {
        return Codegen.optional(this.hasDiscussions);
    }
    /**
     * Set to `true` to enable the (deprecated) downloads features on the repository.
     * 
     */
    @Export(name="hasDownloads", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> hasDownloads;

    /**
     * @return Set to `true` to enable the (deprecated) downloads features on the repository.
     * 
     */
    public Output<Optional<Boolean>> hasDownloads() {
        return Codegen.optional(this.hasDownloads);
    }
    /**
     * Set to `true` to enable the GitHub Issues features
     * on the repository.
     * 
     */
    @Export(name="hasIssues", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> hasIssues;

    /**
     * @return Set to `true` to enable the GitHub Issues features
     * on the repository.
     * 
     */
    public Output<Optional<Boolean>> hasIssues() {
        return Codegen.optional(this.hasIssues);
    }
    /**
     * Set to `true` to enable the GitHub Projects features on the repository. Per the GitHub [documentation](https://developer.github.com/v3/repos/#create) when in an organization that has disabled repository projects it will default to `false` and will otherwise default to `true`. If you specify `true` when it has been disabled it will return an error.
     * 
     */
    @Export(name="hasProjects", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> hasProjects;

    /**
     * @return Set to `true` to enable the GitHub Projects features on the repository. Per the GitHub [documentation](https://developer.github.com/v3/repos/#create) when in an organization that has disabled repository projects it will default to `false` and will otherwise default to `true`. If you specify `true` when it has been disabled it will return an error.
     * 
     */
    public Output<Optional<Boolean>> hasProjects() {
        return Codegen.optional(this.hasProjects);
    }
    /**
     * Set to `true` to enable the GitHub Wiki features on
     * the repository.
     * 
     */
    @Export(name="hasWiki", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> hasWiki;

    /**
     * @return Set to `true` to enable the GitHub Wiki features on
     * the repository.
     * 
     */
    public Output<Optional<Boolean>> hasWiki() {
        return Codegen.optional(this.hasWiki);
    }
    /**
     * URL of a page describing the project.
     * 
     */
    @Export(name="homepageUrl", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> homepageUrl;

    /**
     * @return URL of a page describing the project.
     * 
     */
    public Output<Optional<String>> homepageUrl() {
        return Codegen.optional(this.homepageUrl);
    }
    /**
     * The absolute URL (including scheme) of the rendered GitHub Pages site e.g. `https://username.github.io`.
     * 
     */
    @Export(name="htmlUrl", refs={String.class}, tree="[0]")
    private Output<String> htmlUrl;

    /**
     * @return The absolute URL (including scheme) of the rendered GitHub Pages site e.g. `https://username.github.io`.
     * 
     */
    public Output<String> htmlUrl() {
        return this.htmlUrl;
    }
    /**
     * URL that can be provided to `git clone` to clone the repository via HTTPS.
     * 
     */
    @Export(name="httpCloneUrl", refs={String.class}, tree="[0]")
    private Output<String> httpCloneUrl;

    /**
     * @return URL that can be provided to `git clone` to clone the repository via HTTPS.
     * 
     */
    public Output<String> httpCloneUrl() {
        return this.httpCloneUrl;
    }
    /**
     * Set to `true` to not call the vulnerability alerts endpoint so the resource can also be used without admin permissions during read.
     * 
     */
    @Export(name="ignoreVulnerabilityAlertsDuringRead", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> ignoreVulnerabilityAlertsDuringRead;

    /**
     * @return Set to `true` to not call the vulnerability alerts endpoint so the resource can also be used without admin permissions during read.
     * 
     */
    public Output<Optional<Boolean>> ignoreVulnerabilityAlertsDuringRead() {
        return Codegen.optional(this.ignoreVulnerabilityAlertsDuringRead);
    }
    /**
     * Set to `true` to tell GitHub that this is a template repository.
     * 
     */
    @Export(name="isTemplate", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> isTemplate;

    /**
     * @return Set to `true` to tell GitHub that this is a template repository.
     * 
     */
    public Output<Optional<Boolean>> isTemplate() {
        return Codegen.optional(this.isTemplate);
    }
    /**
     * Use the [name of the template](https://github.com/github/choosealicense.com/tree/gh-pages/_licenses) without the extension. For example, &#34;mit&#34; or &#34;mpl-2.0&#34;.
     * 
     */
    @Export(name="licenseTemplate", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> licenseTemplate;

    /**
     * @return Use the [name of the template](https://github.com/github/choosealicense.com/tree/gh-pages/_licenses) without the extension. For example, &#34;mit&#34; or &#34;mpl-2.0&#34;.
     * 
     */
    public Output<Optional<String>> licenseTemplate() {
        return Codegen.optional(this.licenseTemplate);
    }
    /**
     * Can be `PR_BODY`, `PR_TITLE`, or `BLANK` for a default merge commit message. Applicable only if `allow_merge_commit` is `true`.
     * 
     */
    @Export(name="mergeCommitMessage", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> mergeCommitMessage;

    /**
     * @return Can be `PR_BODY`, `PR_TITLE`, or `BLANK` for a default merge commit message. Applicable only if `allow_merge_commit` is `true`.
     * 
     */
    public Output<Optional<String>> mergeCommitMessage() {
        return Codegen.optional(this.mergeCommitMessage);
    }
    /**
     * Can be `PR_TITLE` or `MERGE_MESSAGE` for a default merge commit title. Applicable only if `allow_merge_commit` is `true`.
     * 
     */
    @Export(name="mergeCommitTitle", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> mergeCommitTitle;

    /**
     * @return Can be `PR_TITLE` or `MERGE_MESSAGE` for a default merge commit title. Applicable only if `allow_merge_commit` is `true`.
     * 
     */
    public Output<Optional<String>> mergeCommitTitle() {
        return Codegen.optional(this.mergeCommitTitle);
    }
    /**
     * The name of the repository.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the repository.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * GraphQL global node id for use with v4 API
     * 
     */
    @Export(name="nodeId", refs={String.class}, tree="[0]")
    private Output<String> nodeId;

    /**
     * @return GraphQL global node id for use with v4 API
     * 
     */
    public Output<String> nodeId() {
        return this.nodeId;
    }
    /**
     * The repository&#39;s GitHub Pages configuration. See GitHub Pages Configuration below for details.
     * 
     */
    @Export(name="pages", refs={RepositoryPages.class}, tree="[0]")
    private Output</* @Nullable */ RepositoryPages> pages;

    /**
     * @return The repository&#39;s GitHub Pages configuration. See GitHub Pages Configuration below for details.
     * 
     */
    public Output<Optional<RepositoryPages>> pages() {
        return Codegen.optional(this.pages);
    }
    /**
     * The primary language used in the repository.
     * 
     */
    @Export(name="primaryLanguage", refs={String.class}, tree="[0]")
    private Output<String> primaryLanguage;

    /**
     * @return The primary language used in the repository.
     * 
     */
    public Output<String> primaryLanguage() {
        return this.primaryLanguage;
    }
    /**
     * Set to `true` to create a private repository.
     * Repositories are created as public (e.g. open source) by default.
     * 
     * @deprecated
     * use visibility instead
     * 
     */
    @Deprecated /* use visibility instead */
    @Export(name="private", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> private_;

    /**
     * @return Set to `true` to create a private repository.
     * Repositories are created as public (e.g. open source) by default.
     * 
     */
    public Output<Boolean> private_() {
        return this.private_;
    }
    /**
     * GitHub ID for the repository
     * 
     */
    @Export(name="repoId", refs={Integer.class}, tree="[0]")
    private Output<Integer> repoId;

    /**
     * @return GitHub ID for the repository
     * 
     */
    public Output<Integer> repoId() {
        return this.repoId;
    }
    /**
     * The repository&#39;s [security and analysis](https://docs.github.com/en/repositories/managing-your-repositorys-settings-and-features/enabling-features-for-your-repository/managing-security-and-analysis-settings-for-your-repository) configuration. See Security and Analysis Configuration below for details.
     * 
     */
    @Export(name="securityAndAnalysis", refs={RepositorySecurityAndAnalysis.class}, tree="[0]")
    private Output<RepositorySecurityAndAnalysis> securityAndAnalysis;

    /**
     * @return The repository&#39;s [security and analysis](https://docs.github.com/en/repositories/managing-your-repositorys-settings-and-features/enabling-features-for-your-repository/managing-security-and-analysis-settings-for-your-repository) configuration. See Security and Analysis Configuration below for details.
     * 
     */
    public Output<RepositorySecurityAndAnalysis> securityAndAnalysis() {
        return this.securityAndAnalysis;
    }
    /**
     * Can be `PR_BODY`, `COMMIT_MESSAGES`, or `BLANK` for a default squash merge commit message. Applicable only if `allow_squash_merge` is `true`.
     * 
     */
    @Export(name="squashMergeCommitMessage", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> squashMergeCommitMessage;

    /**
     * @return Can be `PR_BODY`, `COMMIT_MESSAGES`, or `BLANK` for a default squash merge commit message. Applicable only if `allow_squash_merge` is `true`.
     * 
     */
    public Output<Optional<String>> squashMergeCommitMessage() {
        return Codegen.optional(this.squashMergeCommitMessage);
    }
    /**
     * Can be `PR_TITLE` or `COMMIT_OR_PR_TITLE` for a default squash merge commit title. Applicable only if `allow_squash_merge` is `true`.
     * 
     */
    @Export(name="squashMergeCommitTitle", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> squashMergeCommitTitle;

    /**
     * @return Can be `PR_TITLE` or `COMMIT_OR_PR_TITLE` for a default squash merge commit title. Applicable only if `allow_squash_merge` is `true`.
     * 
     */
    public Output<Optional<String>> squashMergeCommitTitle() {
        return Codegen.optional(this.squashMergeCommitTitle);
    }
    /**
     * URL that can be provided to `git clone` to clone the repository via SSH.
     * 
     */
    @Export(name="sshCloneUrl", refs={String.class}, tree="[0]")
    private Output<String> sshCloneUrl;

    /**
     * @return URL that can be provided to `git clone` to clone the repository via SSH.
     * 
     */
    public Output<String> sshCloneUrl() {
        return this.sshCloneUrl;
    }
    /**
     * URL that can be provided to `svn checkout` to check out the repository via GitHub&#39;s Subversion protocol emulation.
     * 
     */
    @Export(name="svnUrl", refs={String.class}, tree="[0]")
    private Output<String> svnUrl;

    /**
     * @return URL that can be provided to `svn checkout` to check out the repository via GitHub&#39;s Subversion protocol emulation.
     * 
     */
    public Output<String> svnUrl() {
        return this.svnUrl;
    }
    /**
     * Use a template repository to create this resource. See Template Repositories below for details.
     * 
     */
    @Export(name="template", refs={RepositoryTemplate.class}, tree="[0]")
    private Output</* @Nullable */ RepositoryTemplate> template;

    /**
     * @return Use a template repository to create this resource. See Template Repositories below for details.
     * 
     */
    public Output<Optional<RepositoryTemplate>> template() {
        return Codegen.optional(this.template);
    }
    /**
     * The list of topics of the repository.
     * 
     */
    @Export(name="topics", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> topics;

    /**
     * @return The list of topics of the repository.
     * 
     */
    public Output<List<String>> topics() {
        return this.topics;
    }
    /**
     * Can be `public` or `private`. If your organization is associated with an enterprise account using GitHub Enterprise Cloud or GitHub Enterprise Server 2.20+, visibility can also be `internal`. The `visibility` parameter overrides the `private` parameter.
     * 
     */
    @Export(name="visibility", refs={String.class}, tree="[0]")
    private Output<String> visibility;

    /**
     * @return Can be `public` or `private`. If your organization is associated with an enterprise account using GitHub Enterprise Cloud or GitHub Enterprise Server 2.20+, visibility can also be `internal`. The `visibility` parameter overrides the `private` parameter.
     * 
     */
    public Output<String> visibility() {
        return this.visibility;
    }
    /**
     * Set to `true` to enable security alerts for vulnerable dependencies. Enabling requires alerts to be enabled on the owner level. (Note for importing: GitHub enables the alerts on public repos but disables them on private repos by default.) See [GitHub Documentation](https://help.github.com/en/github/managing-security-vulnerabilities/about-security-alerts-for-vulnerable-dependencies) for details. Note that vulnerability alerts have not been successfully tested on any GitHub Enterprise instance and may be unavailable in those settings.
     * 
     */
    @Export(name="vulnerabilityAlerts", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> vulnerabilityAlerts;

    /**
     * @return Set to `true` to enable security alerts for vulnerable dependencies. Enabling requires alerts to be enabled on the owner level. (Note for importing: GitHub enables the alerts on public repos but disables them on private repos by default.) See [GitHub Documentation](https://help.github.com/en/github/managing-security-vulnerabilities/about-security-alerts-for-vulnerable-dependencies) for details. Note that vulnerability alerts have not been successfully tested on any GitHub Enterprise instance and may be unavailable in those settings.
     * 
     */
    public Output<Optional<Boolean>> vulnerabilityAlerts() {
        return Codegen.optional(this.vulnerabilityAlerts);
    }
    /**
     * Require contributors to sign off on web-based commits. See more [here](https://docs.github.com/en/repositories/managing-your-repositorys-settings-and-features/managing-repository-settings/managing-the-commit-signoff-policy-for-your-repository). Defaults to `false`.
     * 
     */
    @Export(name="webCommitSignoffRequired", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> webCommitSignoffRequired;

    /**
     * @return Require contributors to sign off on web-based commits. See more [here](https://docs.github.com/en/repositories/managing-your-repositorys-settings-and-features/managing-repository-settings/managing-the-commit-signoff-policy-for-your-repository). Defaults to `false`.
     * 
     */
    public Output<Optional<Boolean>> webCommitSignoffRequired() {
        return Codegen.optional(this.webCommitSignoffRequired);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Repository(String name) {
        this(name, RepositoryArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Repository(String name, @Nullable RepositoryArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Repository(String name, @Nullable RepositoryArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/repository:Repository", name, args == null ? RepositoryArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Repository(String name, Output<String> id, @Nullable RepositoryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("github:index/repository:Repository", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static Repository get(String name, Output<String> id, @Nullable RepositoryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Repository(name, id, state, options);
    }
}
