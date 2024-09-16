// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RepositoryFileArgs extends com.pulumi.resources.ResourceArgs {

    public static final RepositoryFileArgs Empty = new RepositoryFileArgs();

    /**
     * Automatically create the branch if it could not be found. Defaults to false. Subsequent reads if the branch is deleted will occur from &#39;autocreate_branch_source_branch&#39;.
     * 
     */
    @Import(name="autocreateBranch")
    private @Nullable Output<Boolean> autocreateBranch;

    /**
     * @return Automatically create the branch if it could not be found. Defaults to false. Subsequent reads if the branch is deleted will occur from &#39;autocreate_branch_source_branch&#39;.
     * 
     */
    public Optional<Output<Boolean>> autocreateBranch() {
        return Optional.ofNullable(this.autocreateBranch);
    }

    /**
     * The branch name to start from, if &#39;autocreate_branch&#39; is set. Defaults to &#39;main&#39;.
     * 
     */
    @Import(name="autocreateBranchSourceBranch")
    private @Nullable Output<String> autocreateBranchSourceBranch;

    /**
     * @return The branch name to start from, if &#39;autocreate_branch&#39; is set. Defaults to &#39;main&#39;.
     * 
     */
    public Optional<Output<String>> autocreateBranchSourceBranch() {
        return Optional.ofNullable(this.autocreateBranchSourceBranch);
    }

    /**
     * The commit hash to start from, if &#39;autocreate_branch&#39; is set. Defaults to the tip of &#39;autocreate_branch_source_branch&#39;. If provided, &#39;autocreate_branch_source_branch&#39; is ignored.
     * 
     */
    @Import(name="autocreateBranchSourceSha")
    private @Nullable Output<String> autocreateBranchSourceSha;

    /**
     * @return The commit hash to start from, if &#39;autocreate_branch&#39; is set. Defaults to the tip of &#39;autocreate_branch_source_branch&#39;. If provided, &#39;autocreate_branch_source_branch&#39; is ignored.
     * 
     */
    public Optional<Output<String>> autocreateBranchSourceSha() {
        return Optional.ofNullable(this.autocreateBranchSourceSha);
    }

    /**
     * Git branch (defaults to the repository&#39;s default branch).
     * The branch must already exist, it will only be created automatically if &#39;autocreate_branch&#39; is set true.
     * 
     */
    @Import(name="branch")
    private @Nullable Output<String> branch;

    /**
     * @return Git branch (defaults to the repository&#39;s default branch).
     * The branch must already exist, it will only be created automatically if &#39;autocreate_branch&#39; is set true.
     * 
     */
    public Optional<Output<String>> branch() {
        return Optional.ofNullable(this.branch);
    }

    /**
     * Committer author name to use. **NOTE:** GitHub app users may omit author and email information so GitHub can verify commits as the GitHub App. This maybe useful when a branch protection rule requires signed commits.
     * 
     */
    @Import(name="commitAuthor")
    private @Nullable Output<String> commitAuthor;

    /**
     * @return Committer author name to use. **NOTE:** GitHub app users may omit author and email information so GitHub can verify commits as the GitHub App. This maybe useful when a branch protection rule requires signed commits.
     * 
     */
    public Optional<Output<String>> commitAuthor() {
        return Optional.ofNullable(this.commitAuthor);
    }

    /**
     * Committer email address to use. **NOTE:** GitHub app users may omit author and email information so GitHub can verify commits as the GitHub App. This may be useful when a branch protection rule requires signed commits.
     * 
     */
    @Import(name="commitEmail")
    private @Nullable Output<String> commitEmail;

    /**
     * @return Committer email address to use. **NOTE:** GitHub app users may omit author and email information so GitHub can verify commits as the GitHub App. This may be useful when a branch protection rule requires signed commits.
     * 
     */
    public Optional<Output<String>> commitEmail() {
        return Optional.ofNullable(this.commitEmail);
    }

    /**
     * The commit message when creating, updating or deleting the managed file.
     * 
     */
    @Import(name="commitMessage")
    private @Nullable Output<String> commitMessage;

    /**
     * @return The commit message when creating, updating or deleting the managed file.
     * 
     */
    public Optional<Output<String>> commitMessage() {
        return Optional.ofNullable(this.commitMessage);
    }

    /**
     * The file content.
     * 
     */
    @Import(name="content", required=true)
    private Output<String> content;

    /**
     * @return The file content.
     * 
     */
    public Output<String> content() {
        return this.content;
    }

    /**
     * The path of the file to manage.
     * 
     */
    @Import(name="file", required=true)
    private Output<String> file;

    /**
     * @return The path of the file to manage.
     * 
     */
    public Output<String> file() {
        return this.file;
    }

    /**
     * Enable overwriting existing files. If set to `true` it will overwrite an existing file with the same name. If set to `false` it will fail if there is an existing file with the same name.
     * 
     */
    @Import(name="overwriteOnCreate")
    private @Nullable Output<Boolean> overwriteOnCreate;

    /**
     * @return Enable overwriting existing files. If set to `true` it will overwrite an existing file with the same name. If set to `false` it will fail if there is an existing file with the same name.
     * 
     */
    public Optional<Output<Boolean>> overwriteOnCreate() {
        return Optional.ofNullable(this.overwriteOnCreate);
    }

    /**
     * The repository to create the file in.
     * 
     */
    @Import(name="repository", required=true)
    private Output<String> repository;

    /**
     * @return The repository to create the file in.
     * 
     */
    public Output<String> repository() {
        return this.repository;
    }

    private RepositoryFileArgs() {}

    private RepositoryFileArgs(RepositoryFileArgs $) {
        this.autocreateBranch = $.autocreateBranch;
        this.autocreateBranchSourceBranch = $.autocreateBranchSourceBranch;
        this.autocreateBranchSourceSha = $.autocreateBranchSourceSha;
        this.branch = $.branch;
        this.commitAuthor = $.commitAuthor;
        this.commitEmail = $.commitEmail;
        this.commitMessage = $.commitMessage;
        this.content = $.content;
        this.file = $.file;
        this.overwriteOnCreate = $.overwriteOnCreate;
        this.repository = $.repository;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RepositoryFileArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RepositoryFileArgs $;

        public Builder() {
            $ = new RepositoryFileArgs();
        }

        public Builder(RepositoryFileArgs defaults) {
            $ = new RepositoryFileArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param autocreateBranch Automatically create the branch if it could not be found. Defaults to false. Subsequent reads if the branch is deleted will occur from &#39;autocreate_branch_source_branch&#39;.
         * 
         * @return builder
         * 
         */
        public Builder autocreateBranch(@Nullable Output<Boolean> autocreateBranch) {
            $.autocreateBranch = autocreateBranch;
            return this;
        }

        /**
         * @param autocreateBranch Automatically create the branch if it could not be found. Defaults to false. Subsequent reads if the branch is deleted will occur from &#39;autocreate_branch_source_branch&#39;.
         * 
         * @return builder
         * 
         */
        public Builder autocreateBranch(Boolean autocreateBranch) {
            return autocreateBranch(Output.of(autocreateBranch));
        }

        /**
         * @param autocreateBranchSourceBranch The branch name to start from, if &#39;autocreate_branch&#39; is set. Defaults to &#39;main&#39;.
         * 
         * @return builder
         * 
         */
        public Builder autocreateBranchSourceBranch(@Nullable Output<String> autocreateBranchSourceBranch) {
            $.autocreateBranchSourceBranch = autocreateBranchSourceBranch;
            return this;
        }

        /**
         * @param autocreateBranchSourceBranch The branch name to start from, if &#39;autocreate_branch&#39; is set. Defaults to &#39;main&#39;.
         * 
         * @return builder
         * 
         */
        public Builder autocreateBranchSourceBranch(String autocreateBranchSourceBranch) {
            return autocreateBranchSourceBranch(Output.of(autocreateBranchSourceBranch));
        }

        /**
         * @param autocreateBranchSourceSha The commit hash to start from, if &#39;autocreate_branch&#39; is set. Defaults to the tip of &#39;autocreate_branch_source_branch&#39;. If provided, &#39;autocreate_branch_source_branch&#39; is ignored.
         * 
         * @return builder
         * 
         */
        public Builder autocreateBranchSourceSha(@Nullable Output<String> autocreateBranchSourceSha) {
            $.autocreateBranchSourceSha = autocreateBranchSourceSha;
            return this;
        }

        /**
         * @param autocreateBranchSourceSha The commit hash to start from, if &#39;autocreate_branch&#39; is set. Defaults to the tip of &#39;autocreate_branch_source_branch&#39;. If provided, &#39;autocreate_branch_source_branch&#39; is ignored.
         * 
         * @return builder
         * 
         */
        public Builder autocreateBranchSourceSha(String autocreateBranchSourceSha) {
            return autocreateBranchSourceSha(Output.of(autocreateBranchSourceSha));
        }

        /**
         * @param branch Git branch (defaults to the repository&#39;s default branch).
         * The branch must already exist, it will only be created automatically if &#39;autocreate_branch&#39; is set true.
         * 
         * @return builder
         * 
         */
        public Builder branch(@Nullable Output<String> branch) {
            $.branch = branch;
            return this;
        }

        /**
         * @param branch Git branch (defaults to the repository&#39;s default branch).
         * The branch must already exist, it will only be created automatically if &#39;autocreate_branch&#39; is set true.
         * 
         * @return builder
         * 
         */
        public Builder branch(String branch) {
            return branch(Output.of(branch));
        }

        /**
         * @param commitAuthor Committer author name to use. **NOTE:** GitHub app users may omit author and email information so GitHub can verify commits as the GitHub App. This maybe useful when a branch protection rule requires signed commits.
         * 
         * @return builder
         * 
         */
        public Builder commitAuthor(@Nullable Output<String> commitAuthor) {
            $.commitAuthor = commitAuthor;
            return this;
        }

        /**
         * @param commitAuthor Committer author name to use. **NOTE:** GitHub app users may omit author and email information so GitHub can verify commits as the GitHub App. This maybe useful when a branch protection rule requires signed commits.
         * 
         * @return builder
         * 
         */
        public Builder commitAuthor(String commitAuthor) {
            return commitAuthor(Output.of(commitAuthor));
        }

        /**
         * @param commitEmail Committer email address to use. **NOTE:** GitHub app users may omit author and email information so GitHub can verify commits as the GitHub App. This may be useful when a branch protection rule requires signed commits.
         * 
         * @return builder
         * 
         */
        public Builder commitEmail(@Nullable Output<String> commitEmail) {
            $.commitEmail = commitEmail;
            return this;
        }

        /**
         * @param commitEmail Committer email address to use. **NOTE:** GitHub app users may omit author and email information so GitHub can verify commits as the GitHub App. This may be useful when a branch protection rule requires signed commits.
         * 
         * @return builder
         * 
         */
        public Builder commitEmail(String commitEmail) {
            return commitEmail(Output.of(commitEmail));
        }

        /**
         * @param commitMessage The commit message when creating, updating or deleting the managed file.
         * 
         * @return builder
         * 
         */
        public Builder commitMessage(@Nullable Output<String> commitMessage) {
            $.commitMessage = commitMessage;
            return this;
        }

        /**
         * @param commitMessage The commit message when creating, updating or deleting the managed file.
         * 
         * @return builder
         * 
         */
        public Builder commitMessage(String commitMessage) {
            return commitMessage(Output.of(commitMessage));
        }

        /**
         * @param content The file content.
         * 
         * @return builder
         * 
         */
        public Builder content(Output<String> content) {
            $.content = content;
            return this;
        }

        /**
         * @param content The file content.
         * 
         * @return builder
         * 
         */
        public Builder content(String content) {
            return content(Output.of(content));
        }

        /**
         * @param file The path of the file to manage.
         * 
         * @return builder
         * 
         */
        public Builder file(Output<String> file) {
            $.file = file;
            return this;
        }

        /**
         * @param file The path of the file to manage.
         * 
         * @return builder
         * 
         */
        public Builder file(String file) {
            return file(Output.of(file));
        }

        /**
         * @param overwriteOnCreate Enable overwriting existing files. If set to `true` it will overwrite an existing file with the same name. If set to `false` it will fail if there is an existing file with the same name.
         * 
         * @return builder
         * 
         */
        public Builder overwriteOnCreate(@Nullable Output<Boolean> overwriteOnCreate) {
            $.overwriteOnCreate = overwriteOnCreate;
            return this;
        }

        /**
         * @param overwriteOnCreate Enable overwriting existing files. If set to `true` it will overwrite an existing file with the same name. If set to `false` it will fail if there is an existing file with the same name.
         * 
         * @return builder
         * 
         */
        public Builder overwriteOnCreate(Boolean overwriteOnCreate) {
            return overwriteOnCreate(Output.of(overwriteOnCreate));
        }

        /**
         * @param repository The repository to create the file in.
         * 
         * @return builder
         * 
         */
        public Builder repository(Output<String> repository) {
            $.repository = repository;
            return this;
        }

        /**
         * @param repository The repository to create the file in.
         * 
         * @return builder
         * 
         */
        public Builder repository(String repository) {
            return repository(Output.of(repository));
        }

        public RepositoryFileArgs build() {
            if ($.content == null) {
                throw new MissingRequiredPropertyException("RepositoryFileArgs", "content");
            }
            if ($.file == null) {
                throw new MissingRequiredPropertyException("RepositoryFileArgs", "file");
            }
            if ($.repository == null) {
                throw new MissingRequiredPropertyException("RepositoryFileArgs", "repository");
            }
            return $;
        }
    }

}
