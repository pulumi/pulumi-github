// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetRepositoryTeamsTeam {
    /**
     * @return The name of the repository.
     * 
     */
    private String name;
    /**
     * @return Team permission
     * 
     */
    private String permission;
    /**
     * @return Team slug
     * 
     */
    private String slug;

    private GetRepositoryTeamsTeam() {}
    /**
     * @return The name of the repository.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return Team permission
     * 
     */
    public String permission() {
        return this.permission;
    }
    /**
     * @return Team slug
     * 
     */
    public String slug() {
        return this.slug;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetRepositoryTeamsTeam defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String name;
        private String permission;
        private String slug;
        public Builder() {}
        public Builder(GetRepositoryTeamsTeam defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.name = defaults.name;
    	      this.permission = defaults.permission;
    	      this.slug = defaults.slug;
        }

        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetRepositoryTeamsTeam", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder permission(String permission) {
            if (permission == null) {
              throw new MissingRequiredPropertyException("GetRepositoryTeamsTeam", "permission");
            }
            this.permission = permission;
            return this;
        }
        @CustomType.Setter
        public Builder slug(String slug) {
            if (slug == null) {
              throw new MissingRequiredPropertyException("GetRepositoryTeamsTeam", "slug");
            }
            this.slug = slug;
            return this;
        }
        public GetRepositoryTeamsTeam build() {
            final var _resultValue = new GetRepositoryTeamsTeam();
            _resultValue.name = name;
            _resultValue.permission = permission;
            _resultValue.slug = slug;
            return _resultValue;
        }
    }
}
