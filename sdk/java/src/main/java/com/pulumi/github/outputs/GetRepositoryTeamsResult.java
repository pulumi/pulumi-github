// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.github.outputs.GetRepositoryTeamsTeam;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetRepositoryTeamsResult {
    private String fullName;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return Team name
     * 
     */
    private String name;
    /**
     * @return List of teams which have access to the repository
     * 
     */
    private List<GetRepositoryTeamsTeam> teams;

    private GetRepositoryTeamsResult() {}
    public String fullName() {
        return this.fullName;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Team name
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return List of teams which have access to the repository
     * 
     */
    public List<GetRepositoryTeamsTeam> teams() {
        return this.teams;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetRepositoryTeamsResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String fullName;
        private String id;
        private String name;
        private List<GetRepositoryTeamsTeam> teams;
        public Builder() {}
        public Builder(GetRepositoryTeamsResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.fullName = defaults.fullName;
    	      this.id = defaults.id;
    	      this.name = defaults.name;
    	      this.teams = defaults.teams;
        }

        @CustomType.Setter
        public Builder fullName(String fullName) {
            this.fullName = Objects.requireNonNull(fullName);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder teams(List<GetRepositoryTeamsTeam> teams) {
            this.teams = Objects.requireNonNull(teams);
            return this;
        }
        public Builder teams(GetRepositoryTeamsTeam... teams) {
            return teams(List.of(teams));
        }
        public GetRepositoryTeamsResult build() {
            final var o = new GetRepositoryTeamsResult();
            o.fullName = fullName;
            o.id = id;
            o.name = name;
            o.teams = teams;
            return o;
        }
    }
}