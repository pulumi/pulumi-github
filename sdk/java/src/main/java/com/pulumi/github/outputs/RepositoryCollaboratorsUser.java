// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class RepositoryCollaboratorsUser {
    /**
     * @return The permission of the outside collaborators for the repository.
     * Must be one of `pull`, `push`, `maintain`, `triage` or `admin` or the name of an existing [custom repository role](https://docs.github.com/en/enterprise-cloud@latest/organizations/managing-peoples-access-to-your-organization-with-roles/managing-custom-repository-roles-for-an-organization) within the organization for organization-owned repositories.
     * Must be `push` for personal repositories. Defaults to `push`.
     * 
     */
    private @Nullable String permission;
    /**
     * @return The user to add to the repository as a collaborator.
     * 
     */
    private String username;

    private RepositoryCollaboratorsUser() {}
    /**
     * @return The permission of the outside collaborators for the repository.
     * Must be one of `pull`, `push`, `maintain`, `triage` or `admin` or the name of an existing [custom repository role](https://docs.github.com/en/enterprise-cloud@latest/organizations/managing-peoples-access-to-your-organization-with-roles/managing-custom-repository-roles-for-an-organization) within the organization for organization-owned repositories.
     * Must be `push` for personal repositories. Defaults to `push`.
     * 
     */
    public Optional<String> permission() {
        return Optional.ofNullable(this.permission);
    }
    /**
     * @return The user to add to the repository as a collaborator.
     * 
     */
    public String username() {
        return this.username;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(RepositoryCollaboratorsUser defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String permission;
        private String username;
        public Builder() {}
        public Builder(RepositoryCollaboratorsUser defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.permission = defaults.permission;
    	      this.username = defaults.username;
        }

        @CustomType.Setter
        public Builder permission(@Nullable String permission) {
            this.permission = permission;
            return this;
        }
        @CustomType.Setter
        public Builder username(String username) {
            this.username = Objects.requireNonNull(username);
            return this;
        }
        public RepositoryCollaboratorsUser build() {
            final var _resultValue = new RepositoryCollaboratorsUser();
            _resultValue.permission = permission;
            _resultValue.username = username;
            return _resultValue;
        }
    }
}
