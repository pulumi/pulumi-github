// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class TeamRepositoryArgs extends com.pulumi.resources.ResourceArgs {

    public static final TeamRepositoryArgs Empty = new TeamRepositoryArgs();

    /**
     * The permissions of team members regarding the repository.
     * Must be one of `pull`, `triage`, `push`, `maintain`, `admin` or the name of an existing [custom repository role](https://docs.github.com/en/enterprise-cloud{@literal @}latest/organizations/managing-peoples-access-to-your-organization-with-roles/managing-custom-repository-roles-for-an-organization) within the organisation. Defaults to `pull`.
     * 
     */
    @Import(name="permission")
    private @Nullable Output<String> permission;

    /**
     * @return The permissions of team members regarding the repository.
     * Must be one of `pull`, `triage`, `push`, `maintain`, `admin` or the name of an existing [custom repository role](https://docs.github.com/en/enterprise-cloud{@literal @}latest/organizations/managing-peoples-access-to-your-organization-with-roles/managing-custom-repository-roles-for-an-organization) within the organisation. Defaults to `pull`.
     * 
     */
    public Optional<Output<String>> permission() {
        return Optional.ofNullable(this.permission);
    }

    /**
     * The repository to add to the team.
     * 
     */
    @Import(name="repository", required=true)
    private Output<String> repository;

    /**
     * @return The repository to add to the team.
     * 
     */
    public Output<String> repository() {
        return this.repository;
    }

    /**
     * The GitHub team id or the GitHub team slug
     * 
     */
    @Import(name="teamId", required=true)
    private Output<String> teamId;

    /**
     * @return The GitHub team id or the GitHub team slug
     * 
     */
    public Output<String> teamId() {
        return this.teamId;
    }

    private TeamRepositoryArgs() {}

    private TeamRepositoryArgs(TeamRepositoryArgs $) {
        this.permission = $.permission;
        this.repository = $.repository;
        this.teamId = $.teamId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(TeamRepositoryArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private TeamRepositoryArgs $;

        public Builder() {
            $ = new TeamRepositoryArgs();
        }

        public Builder(TeamRepositoryArgs defaults) {
            $ = new TeamRepositoryArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param permission The permissions of team members regarding the repository.
         * Must be one of `pull`, `triage`, `push`, `maintain`, `admin` or the name of an existing [custom repository role](https://docs.github.com/en/enterprise-cloud{@literal @}latest/organizations/managing-peoples-access-to-your-organization-with-roles/managing-custom-repository-roles-for-an-organization) within the organisation. Defaults to `pull`.
         * 
         * @return builder
         * 
         */
        public Builder permission(@Nullable Output<String> permission) {
            $.permission = permission;
            return this;
        }

        /**
         * @param permission The permissions of team members regarding the repository.
         * Must be one of `pull`, `triage`, `push`, `maintain`, `admin` or the name of an existing [custom repository role](https://docs.github.com/en/enterprise-cloud{@literal @}latest/organizations/managing-peoples-access-to-your-organization-with-roles/managing-custom-repository-roles-for-an-organization) within the organisation. Defaults to `pull`.
         * 
         * @return builder
         * 
         */
        public Builder permission(String permission) {
            return permission(Output.of(permission));
        }

        /**
         * @param repository The repository to add to the team.
         * 
         * @return builder
         * 
         */
        public Builder repository(Output<String> repository) {
            $.repository = repository;
            return this;
        }

        /**
         * @param repository The repository to add to the team.
         * 
         * @return builder
         * 
         */
        public Builder repository(String repository) {
            return repository(Output.of(repository));
        }

        /**
         * @param teamId The GitHub team id or the GitHub team slug
         * 
         * @return builder
         * 
         */
        public Builder teamId(Output<String> teamId) {
            $.teamId = teamId;
            return this;
        }

        /**
         * @param teamId The GitHub team id or the GitHub team slug
         * 
         * @return builder
         * 
         */
        public Builder teamId(String teamId) {
            return teamId(Output.of(teamId));
        }

        public TeamRepositoryArgs build() {
            if ($.repository == null) {
                throw new MissingRequiredPropertyException("TeamRepositoryArgs", "repository");
            }
            if ($.teamId == null) {
                throw new MissingRequiredPropertyException("TeamRepositoryArgs", "teamId");
            }
            return $;
        }
    }

}