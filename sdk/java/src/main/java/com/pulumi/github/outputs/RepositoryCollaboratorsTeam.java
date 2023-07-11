// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class RepositoryCollaboratorsTeam {
    /**
     * @return The permission of the outside collaborators for the repository.
     * Must be one of `pull`, `triage`, `push`, `maintain`, `admin` or the name of an existing [custom repository role](https://docs.github.com/en/enterprise-cloud@latest/organizations/managing-peoples-access-to-your-organization-with-roles/managing-custom-repository-roles-for-an-organization) within the organisation. Defaults to `pull`.
     * Must be `push` for personal repositories. Defaults to `push`.
     * 
     */
    private @Nullable String permission;
    /**
     * @return The GitHub team id or the GitHub team slug
     * 
     */
    private String teamId;

    private RepositoryCollaboratorsTeam() {}
    /**
     * @return The permission of the outside collaborators for the repository.
     * Must be one of `pull`, `triage`, `push`, `maintain`, `admin` or the name of an existing [custom repository role](https://docs.github.com/en/enterprise-cloud@latest/organizations/managing-peoples-access-to-your-organization-with-roles/managing-custom-repository-roles-for-an-organization) within the organisation. Defaults to `pull`.
     * Must be `push` for personal repositories. Defaults to `push`.
     * 
     */
    public Optional<String> permission() {
        return Optional.ofNullable(this.permission);
    }
    /**
     * @return The GitHub team id or the GitHub team slug
     * 
     */
    public String teamId() {
        return this.teamId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(RepositoryCollaboratorsTeam defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String permission;
        private String teamId;
        public Builder() {}
        public Builder(RepositoryCollaboratorsTeam defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.permission = defaults.permission;
    	      this.teamId = defaults.teamId;
        }

        @CustomType.Setter
        public Builder permission(@Nullable String permission) {
            this.permission = permission;
            return this;
        }
        @CustomType.Setter
        public Builder teamId(String teamId) {
            this.teamId = Objects.requireNonNull(teamId);
            return this;
        }
        public RepositoryCollaboratorsTeam build() {
            final var o = new RepositoryCollaboratorsTeam();
            o.permission = permission;
            o.teamId = teamId;
            return o;
        }
    }
}
