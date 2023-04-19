// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.github.outputs.GetRepositoryPage;
import com.pulumi.github.outputs.GetRepositoryTemplate;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetRepositoryResult {
    /**
     * @return Whether the repository allows auto-merging pull requests.
     * 
     */
    private Boolean allowAutoMerge;
    /**
     * @return Whether the repository allows merge commits.
     * 
     */
    private Boolean allowMergeCommit;
    /**
     * @return Whether the repository allows rebase merges.
     * 
     */
    private Boolean allowRebaseMerge;
    /**
     * @return Whether the repository allows squash merges.
     * 
     */
    private Boolean allowSquashMerge;
    /**
     * @return Whether the repository is archived.
     * 
     */
    private Boolean archived;
    /**
     * @return The name of the default branch of the repository.
     * 
     */
    private String defaultBranch;
    /**
     * @return A description of the repository.
     * 
     */
    private @Nullable String description;
    /**
     * @return Whether the repository is a fork.
     * 
     */
    private Boolean fork;
    private String fullName;
    /**
     * @return URL that can be provided to `git clone` to clone the repository anonymously via the git protocol.
     * 
     */
    private String gitCloneUrl;
    /**
     * @return Whether the repository has GitHub Discussions enabled.
     * 
     */
    private Boolean hasDiscussions;
    /**
     * @return Whether the repository has Downloads feature enabled.
     * 
     */
    private Boolean hasDownloads;
    /**
     * @return Whether the repository has GitHub Issues enabled.
     * 
     */
    private Boolean hasIssues;
    /**
     * @return Whether the repository has the GitHub Projects enabled.
     * 
     */
    private Boolean hasProjects;
    /**
     * @return Whether the repository has the GitHub Wiki enabled.
     * 
     */
    private Boolean hasWiki;
    /**
     * @return URL of a page describing the project.
     * 
     */
    private @Nullable String homepageUrl;
    /**
     * @return URL to the repository on the web.
     * 
     */
    private String htmlUrl;
    /**
     * @return URL that can be provided to `git clone` to clone the repository via HTTPS.
     * 
     */
    private String httpCloneUrl;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return Whether the repository is a template repository.
     * 
     */
    private Boolean isTemplate;
    /**
     * @return The default value for a merge commit message.
     * 
     */
    private String mergeCommitMessage;
    /**
     * @return The default value for a merge commit title.
     * 
     */
    private String mergeCommitTitle;
    private String name;
    /**
     * @return GraphQL global node id for use with v4 API
     * 
     */
    private String nodeId;
    /**
     * @return The repository&#39;s GitHub Pages configuration.
     * 
     */
    private List<GetRepositoryPage> pages;
    /**
     * @return Whether the repository is private.
     * 
     */
    private Boolean private_;
    /**
     * @return GitHub ID for the repository
     * 
     */
    private Integer repoId;
    /**
     * @return The default value for a squash merge commit message.
     * 
     */
    private String squashMergeCommitMessage;
    /**
     * @return The default value for a squash merge commit title.
     * 
     */
    private String squashMergeCommitTitle;
    /**
     * @return URL that can be provided to `git clone` to clone the repository via SSH.
     * 
     */
    private String sshCloneUrl;
    /**
     * @return URL that can be provided to `svn checkout` to check out the repository via GitHub&#39;s Subversion protocol emulation.
     * 
     */
    private String svnUrl;
    /**
     * @return The repository source template configuration.
     * 
     */
    private GetRepositoryTemplate template;
    /**
     * @return The list of topics of the repository.
     * 
     */
    private List<String> topics;
    /**
     * @return Whether the repository is public, private or internal.
     * 
     */
    private String visibility;

    private GetRepositoryResult() {}
    /**
     * @return Whether the repository allows auto-merging pull requests.
     * 
     */
    public Boolean allowAutoMerge() {
        return this.allowAutoMerge;
    }
    /**
     * @return Whether the repository allows merge commits.
     * 
     */
    public Boolean allowMergeCommit() {
        return this.allowMergeCommit;
    }
    /**
     * @return Whether the repository allows rebase merges.
     * 
     */
    public Boolean allowRebaseMerge() {
        return this.allowRebaseMerge;
    }
    /**
     * @return Whether the repository allows squash merges.
     * 
     */
    public Boolean allowSquashMerge() {
        return this.allowSquashMerge;
    }
    /**
     * @return Whether the repository is archived.
     * 
     */
    public Boolean archived() {
        return this.archived;
    }
    /**
     * @return The name of the default branch of the repository.
     * 
     */
    public String defaultBranch() {
        return this.defaultBranch;
    }
    /**
     * @return A description of the repository.
     * 
     */
    public Optional<String> description() {
        return Optional.ofNullable(this.description);
    }
    /**
     * @return Whether the repository is a fork.
     * 
     */
    public Boolean fork() {
        return this.fork;
    }
    public String fullName() {
        return this.fullName;
    }
    /**
     * @return URL that can be provided to `git clone` to clone the repository anonymously via the git protocol.
     * 
     */
    public String gitCloneUrl() {
        return this.gitCloneUrl;
    }
    /**
     * @return Whether the repository has GitHub Discussions enabled.
     * 
     */
    public Boolean hasDiscussions() {
        return this.hasDiscussions;
    }
    /**
     * @return Whether the repository has Downloads feature enabled.
     * 
     */
    public Boolean hasDownloads() {
        return this.hasDownloads;
    }
    /**
     * @return Whether the repository has GitHub Issues enabled.
     * 
     */
    public Boolean hasIssues() {
        return this.hasIssues;
    }
    /**
     * @return Whether the repository has the GitHub Projects enabled.
     * 
     */
    public Boolean hasProjects() {
        return this.hasProjects;
    }
    /**
     * @return Whether the repository has the GitHub Wiki enabled.
     * 
     */
    public Boolean hasWiki() {
        return this.hasWiki;
    }
    /**
     * @return URL of a page describing the project.
     * 
     */
    public Optional<String> homepageUrl() {
        return Optional.ofNullable(this.homepageUrl);
    }
    /**
     * @return URL to the repository on the web.
     * 
     */
    public String htmlUrl() {
        return this.htmlUrl;
    }
    /**
     * @return URL that can be provided to `git clone` to clone the repository via HTTPS.
     * 
     */
    public String httpCloneUrl() {
        return this.httpCloneUrl;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Whether the repository is a template repository.
     * 
     */
    public Boolean isTemplate() {
        return this.isTemplate;
    }
    /**
     * @return The default value for a merge commit message.
     * 
     */
    public String mergeCommitMessage() {
        return this.mergeCommitMessage;
    }
    /**
     * @return The default value for a merge commit title.
     * 
     */
    public String mergeCommitTitle() {
        return this.mergeCommitTitle;
    }
    public String name() {
        return this.name;
    }
    /**
     * @return GraphQL global node id for use with v4 API
     * 
     */
    public String nodeId() {
        return this.nodeId;
    }
    /**
     * @return The repository&#39;s GitHub Pages configuration.
     * 
     */
    public List<GetRepositoryPage> pages() {
        return this.pages;
    }
    /**
     * @return Whether the repository is private.
     * 
     */
    public Boolean private_() {
        return this.private_;
    }
    /**
     * @return GitHub ID for the repository
     * 
     */
    public Integer repoId() {
        return this.repoId;
    }
    /**
     * @return The default value for a squash merge commit message.
     * 
     */
    public String squashMergeCommitMessage() {
        return this.squashMergeCommitMessage;
    }
    /**
     * @return The default value for a squash merge commit title.
     * 
     */
    public String squashMergeCommitTitle() {
        return this.squashMergeCommitTitle;
    }
    /**
     * @return URL that can be provided to `git clone` to clone the repository via SSH.
     * 
     */
    public String sshCloneUrl() {
        return this.sshCloneUrl;
    }
    /**
     * @return URL that can be provided to `svn checkout` to check out the repository via GitHub&#39;s Subversion protocol emulation.
     * 
     */
    public String svnUrl() {
        return this.svnUrl;
    }
    /**
     * @return The repository source template configuration.
     * 
     */
    public GetRepositoryTemplate template() {
        return this.template;
    }
    /**
     * @return The list of topics of the repository.
     * 
     */
    public List<String> topics() {
        return this.topics;
    }
    /**
     * @return Whether the repository is public, private or internal.
     * 
     */
    public String visibility() {
        return this.visibility;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetRepositoryResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Boolean allowAutoMerge;
        private Boolean allowMergeCommit;
        private Boolean allowRebaseMerge;
        private Boolean allowSquashMerge;
        private Boolean archived;
        private String defaultBranch;
        private @Nullable String description;
        private Boolean fork;
        private String fullName;
        private String gitCloneUrl;
        private Boolean hasDiscussions;
        private Boolean hasDownloads;
        private Boolean hasIssues;
        private Boolean hasProjects;
        private Boolean hasWiki;
        private @Nullable String homepageUrl;
        private String htmlUrl;
        private String httpCloneUrl;
        private String id;
        private Boolean isTemplate;
        private String mergeCommitMessage;
        private String mergeCommitTitle;
        private String name;
        private String nodeId;
        private List<GetRepositoryPage> pages;
        private Boolean private_;
        private Integer repoId;
        private String squashMergeCommitMessage;
        private String squashMergeCommitTitle;
        private String sshCloneUrl;
        private String svnUrl;
        private GetRepositoryTemplate template;
        private List<String> topics;
        private String visibility;
        public Builder() {}
        public Builder(GetRepositoryResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.allowAutoMerge = defaults.allowAutoMerge;
    	      this.allowMergeCommit = defaults.allowMergeCommit;
    	      this.allowRebaseMerge = defaults.allowRebaseMerge;
    	      this.allowSquashMerge = defaults.allowSquashMerge;
    	      this.archived = defaults.archived;
    	      this.defaultBranch = defaults.defaultBranch;
    	      this.description = defaults.description;
    	      this.fork = defaults.fork;
    	      this.fullName = defaults.fullName;
    	      this.gitCloneUrl = defaults.gitCloneUrl;
    	      this.hasDiscussions = defaults.hasDiscussions;
    	      this.hasDownloads = defaults.hasDownloads;
    	      this.hasIssues = defaults.hasIssues;
    	      this.hasProjects = defaults.hasProjects;
    	      this.hasWiki = defaults.hasWiki;
    	      this.homepageUrl = defaults.homepageUrl;
    	      this.htmlUrl = defaults.htmlUrl;
    	      this.httpCloneUrl = defaults.httpCloneUrl;
    	      this.id = defaults.id;
    	      this.isTemplate = defaults.isTemplate;
    	      this.mergeCommitMessage = defaults.mergeCommitMessage;
    	      this.mergeCommitTitle = defaults.mergeCommitTitle;
    	      this.name = defaults.name;
    	      this.nodeId = defaults.nodeId;
    	      this.pages = defaults.pages;
    	      this.private_ = defaults.private_;
    	      this.repoId = defaults.repoId;
    	      this.squashMergeCommitMessage = defaults.squashMergeCommitMessage;
    	      this.squashMergeCommitTitle = defaults.squashMergeCommitTitle;
    	      this.sshCloneUrl = defaults.sshCloneUrl;
    	      this.svnUrl = defaults.svnUrl;
    	      this.template = defaults.template;
    	      this.topics = defaults.topics;
    	      this.visibility = defaults.visibility;
        }

        @CustomType.Setter
        public Builder allowAutoMerge(Boolean allowAutoMerge) {
            this.allowAutoMerge = Objects.requireNonNull(allowAutoMerge);
            return this;
        }
        @CustomType.Setter
        public Builder allowMergeCommit(Boolean allowMergeCommit) {
            this.allowMergeCommit = Objects.requireNonNull(allowMergeCommit);
            return this;
        }
        @CustomType.Setter
        public Builder allowRebaseMerge(Boolean allowRebaseMerge) {
            this.allowRebaseMerge = Objects.requireNonNull(allowRebaseMerge);
            return this;
        }
        @CustomType.Setter
        public Builder allowSquashMerge(Boolean allowSquashMerge) {
            this.allowSquashMerge = Objects.requireNonNull(allowSquashMerge);
            return this;
        }
        @CustomType.Setter
        public Builder archived(Boolean archived) {
            this.archived = Objects.requireNonNull(archived);
            return this;
        }
        @CustomType.Setter
        public Builder defaultBranch(String defaultBranch) {
            this.defaultBranch = Objects.requireNonNull(defaultBranch);
            return this;
        }
        @CustomType.Setter
        public Builder description(@Nullable String description) {
            this.description = description;
            return this;
        }
        @CustomType.Setter
        public Builder fork(Boolean fork) {
            this.fork = Objects.requireNonNull(fork);
            return this;
        }
        @CustomType.Setter
        public Builder fullName(String fullName) {
            this.fullName = Objects.requireNonNull(fullName);
            return this;
        }
        @CustomType.Setter
        public Builder gitCloneUrl(String gitCloneUrl) {
            this.gitCloneUrl = Objects.requireNonNull(gitCloneUrl);
            return this;
        }
        @CustomType.Setter
        public Builder hasDiscussions(Boolean hasDiscussions) {
            this.hasDiscussions = Objects.requireNonNull(hasDiscussions);
            return this;
        }
        @CustomType.Setter
        public Builder hasDownloads(Boolean hasDownloads) {
            this.hasDownloads = Objects.requireNonNull(hasDownloads);
            return this;
        }
        @CustomType.Setter
        public Builder hasIssues(Boolean hasIssues) {
            this.hasIssues = Objects.requireNonNull(hasIssues);
            return this;
        }
        @CustomType.Setter
        public Builder hasProjects(Boolean hasProjects) {
            this.hasProjects = Objects.requireNonNull(hasProjects);
            return this;
        }
        @CustomType.Setter
        public Builder hasWiki(Boolean hasWiki) {
            this.hasWiki = Objects.requireNonNull(hasWiki);
            return this;
        }
        @CustomType.Setter
        public Builder homepageUrl(@Nullable String homepageUrl) {
            this.homepageUrl = homepageUrl;
            return this;
        }
        @CustomType.Setter
        public Builder htmlUrl(String htmlUrl) {
            this.htmlUrl = Objects.requireNonNull(htmlUrl);
            return this;
        }
        @CustomType.Setter
        public Builder httpCloneUrl(String httpCloneUrl) {
            this.httpCloneUrl = Objects.requireNonNull(httpCloneUrl);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder isTemplate(Boolean isTemplate) {
            this.isTemplate = Objects.requireNonNull(isTemplate);
            return this;
        }
        @CustomType.Setter
        public Builder mergeCommitMessage(String mergeCommitMessage) {
            this.mergeCommitMessage = Objects.requireNonNull(mergeCommitMessage);
            return this;
        }
        @CustomType.Setter
        public Builder mergeCommitTitle(String mergeCommitTitle) {
            this.mergeCommitTitle = Objects.requireNonNull(mergeCommitTitle);
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder nodeId(String nodeId) {
            this.nodeId = Objects.requireNonNull(nodeId);
            return this;
        }
        @CustomType.Setter
        public Builder pages(List<GetRepositoryPage> pages) {
            this.pages = Objects.requireNonNull(pages);
            return this;
        }
        public Builder pages(GetRepositoryPage... pages) {
            return pages(List.of(pages));
        }
        @CustomType.Setter("private")
        public Builder private_(Boolean private_) {
            this.private_ = Objects.requireNonNull(private_);
            return this;
        }
        @CustomType.Setter
        public Builder repoId(Integer repoId) {
            this.repoId = Objects.requireNonNull(repoId);
            return this;
        }
        @CustomType.Setter
        public Builder squashMergeCommitMessage(String squashMergeCommitMessage) {
            this.squashMergeCommitMessage = Objects.requireNonNull(squashMergeCommitMessage);
            return this;
        }
        @CustomType.Setter
        public Builder squashMergeCommitTitle(String squashMergeCommitTitle) {
            this.squashMergeCommitTitle = Objects.requireNonNull(squashMergeCommitTitle);
            return this;
        }
        @CustomType.Setter
        public Builder sshCloneUrl(String sshCloneUrl) {
            this.sshCloneUrl = Objects.requireNonNull(sshCloneUrl);
            return this;
        }
        @CustomType.Setter
        public Builder svnUrl(String svnUrl) {
            this.svnUrl = Objects.requireNonNull(svnUrl);
            return this;
        }
        @CustomType.Setter
        public Builder template(GetRepositoryTemplate template) {
            this.template = Objects.requireNonNull(template);
            return this;
        }
        @CustomType.Setter
        public Builder topics(List<String> topics) {
            this.topics = Objects.requireNonNull(topics);
            return this;
        }
        public Builder topics(String... topics) {
            return topics(List.of(topics));
        }
        @CustomType.Setter
        public Builder visibility(String visibility) {
            this.visibility = Objects.requireNonNull(visibility);
            return this;
        }
        public GetRepositoryResult build() {
            final var o = new GetRepositoryResult();
            o.allowAutoMerge = allowAutoMerge;
            o.allowMergeCommit = allowMergeCommit;
            o.allowRebaseMerge = allowRebaseMerge;
            o.allowSquashMerge = allowSquashMerge;
            o.archived = archived;
            o.defaultBranch = defaultBranch;
            o.description = description;
            o.fork = fork;
            o.fullName = fullName;
            o.gitCloneUrl = gitCloneUrl;
            o.hasDiscussions = hasDiscussions;
            o.hasDownloads = hasDownloads;
            o.hasIssues = hasIssues;
            o.hasProjects = hasProjects;
            o.hasWiki = hasWiki;
            o.homepageUrl = homepageUrl;
            o.htmlUrl = htmlUrl;
            o.httpCloneUrl = httpCloneUrl;
            o.id = id;
            o.isTemplate = isTemplate;
            o.mergeCommitMessage = mergeCommitMessage;
            o.mergeCommitTitle = mergeCommitTitle;
            o.name = name;
            o.nodeId = nodeId;
            o.pages = pages;
            o.private_ = private_;
            o.repoId = repoId;
            o.squashMergeCommitMessage = squashMergeCommitMessage;
            o.squashMergeCommitTitle = squashMergeCommitTitle;
            o.sshCloneUrl = sshCloneUrl;
            o.svnUrl = svnUrl;
            o.template = template;
            o.topics = topics;
            o.visibility = visibility;
            return o;
        }
    }
}
