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


public final class RepositoryCollaboratorArgs extends com.pulumi.resources.ResourceArgs {

    public static final RepositoryCollaboratorArgs Empty = new RepositoryCollaboratorArgs();

    /**
     * The permission of the outside collaborator for the repository.
     * Must be one of `pull`, `push`, `maintain`, `triage` or `admin` or the name of an existing [custom repository role](https://docs.github.com/en/enterprise-cloud@latest/organizations/managing-peoples-access-to-your-organization-with-roles/managing-custom-repository-roles-for-an-organization) within the organization for organization-owned repositories.
     * Must be `push` for personal repositories. Defaults to `push`.
     * 
     */
    @Import(name="permission")
    private @Nullable Output<String> permission;

    /**
     * @return The permission of the outside collaborator for the repository.
     * Must be one of `pull`, `push`, `maintain`, `triage` or `admin` or the name of an existing [custom repository role](https://docs.github.com/en/enterprise-cloud@latest/organizations/managing-peoples-access-to-your-organization-with-roles/managing-custom-repository-roles-for-an-organization) within the organization for organization-owned repositories.
     * Must be `push` for personal repositories. Defaults to `push`.
     * 
     */
    public Optional<Output<String>> permission() {
        return Optional.ofNullable(this.permission);
    }

    /**
     * Suppress plan diffs for `triage` and `maintain`.  Defaults to `false`.
     * 
     */
    @Import(name="permissionDiffSuppression")
    private @Nullable Output<Boolean> permissionDiffSuppression;

    /**
     * @return Suppress plan diffs for `triage` and `maintain`.  Defaults to `false`.
     * 
     */
    public Optional<Output<Boolean>> permissionDiffSuppression() {
        return Optional.ofNullable(this.permissionDiffSuppression);
    }

    /**
     * The GitHub repository
     * 
     * &gt; Note: The owner of the repository can be passed as part of the repository name  e.g. `owner-org-name/repo-name`. If owner is not supplied as part of the repository name, it may also be supplied by setting the environment variable `GITHUB_OWNER`.
     * 
     */
    @Import(name="repository", required=true)
    private Output<String> repository;

    /**
     * @return The GitHub repository
     * 
     * &gt; Note: The owner of the repository can be passed as part of the repository name  e.g. `owner-org-name/repo-name`. If owner is not supplied as part of the repository name, it may also be supplied by setting the environment variable `GITHUB_OWNER`.
     * 
     */
    public Output<String> repository() {
        return this.repository;
    }

    /**
     * The user to add to the repository as a collaborator.
     * 
     */
    @Import(name="username", required=true)
    private Output<String> username;

    /**
     * @return The user to add to the repository as a collaborator.
     * 
     */
    public Output<String> username() {
        return this.username;
    }

    private RepositoryCollaboratorArgs() {}

    private RepositoryCollaboratorArgs(RepositoryCollaboratorArgs $) {
        this.permission = $.permission;
        this.permissionDiffSuppression = $.permissionDiffSuppression;
        this.repository = $.repository;
        this.username = $.username;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RepositoryCollaboratorArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RepositoryCollaboratorArgs $;

        public Builder() {
            $ = new RepositoryCollaboratorArgs();
        }

        public Builder(RepositoryCollaboratorArgs defaults) {
            $ = new RepositoryCollaboratorArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param permission The permission of the outside collaborator for the repository.
         * Must be one of `pull`, `push`, `maintain`, `triage` or `admin` or the name of an existing [custom repository role](https://docs.github.com/en/enterprise-cloud@latest/organizations/managing-peoples-access-to-your-organization-with-roles/managing-custom-repository-roles-for-an-organization) within the organization for organization-owned repositories.
         * Must be `push` for personal repositories. Defaults to `push`.
         * 
         * @return builder
         * 
         */
        public Builder permission(@Nullable Output<String> permission) {
            $.permission = permission;
            return this;
        }

        /**
         * @param permission The permission of the outside collaborator for the repository.
         * Must be one of `pull`, `push`, `maintain`, `triage` or `admin` or the name of an existing [custom repository role](https://docs.github.com/en/enterprise-cloud@latest/organizations/managing-peoples-access-to-your-organization-with-roles/managing-custom-repository-roles-for-an-organization) within the organization for organization-owned repositories.
         * Must be `push` for personal repositories. Defaults to `push`.
         * 
         * @return builder
         * 
         */
        public Builder permission(String permission) {
            return permission(Output.of(permission));
        }

        /**
         * @param permissionDiffSuppression Suppress plan diffs for `triage` and `maintain`.  Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder permissionDiffSuppression(@Nullable Output<Boolean> permissionDiffSuppression) {
            $.permissionDiffSuppression = permissionDiffSuppression;
            return this;
        }

        /**
         * @param permissionDiffSuppression Suppress plan diffs for `triage` and `maintain`.  Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder permissionDiffSuppression(Boolean permissionDiffSuppression) {
            return permissionDiffSuppression(Output.of(permissionDiffSuppression));
        }

        /**
         * @param repository The GitHub repository
         * 
         * &gt; Note: The owner of the repository can be passed as part of the repository name  e.g. `owner-org-name/repo-name`. If owner is not supplied as part of the repository name, it may also be supplied by setting the environment variable `GITHUB_OWNER`.
         * 
         * @return builder
         * 
         */
        public Builder repository(Output<String> repository) {
            $.repository = repository;
            return this;
        }

        /**
         * @param repository The GitHub repository
         * 
         * &gt; Note: The owner of the repository can be passed as part of the repository name  e.g. `owner-org-name/repo-name`. If owner is not supplied as part of the repository name, it may also be supplied by setting the environment variable `GITHUB_OWNER`.
         * 
         * @return builder
         * 
         */
        public Builder repository(String repository) {
            return repository(Output.of(repository));
        }

        /**
         * @param username The user to add to the repository as a collaborator.
         * 
         * @return builder
         * 
         */
        public Builder username(Output<String> username) {
            $.username = username;
            return this;
        }

        /**
         * @param username The user to add to the repository as a collaborator.
         * 
         * @return builder
         * 
         */
        public Builder username(String username) {
            return username(Output.of(username));
        }

        public RepositoryCollaboratorArgs build() {
            if ($.repository == null) {
                throw new MissingRequiredPropertyException("RepositoryCollaboratorArgs", "repository");
            }
            if ($.username == null) {
                throw new MissingRequiredPropertyException("RepositoryCollaboratorArgs", "username");
            }
            return $;
        }
    }

}
