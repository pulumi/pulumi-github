// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.github.outputs.GetTeamRepositoriesDetailed;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetTeamResult {
    /**
     * @return the team&#39;s description.
     * 
     */
    private String description;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return List of team members (list of GitHub usernames). Not returned if `summary_only = true`
     * 
     */
    private List<String> members;
    private @Nullable String membershipType;
    /**
     * @return the team&#39;s full name.
     * 
     */
    private String name;
    /**
     * @return the Node ID of the team.
     * 
     */
    private String nodeId;
    /**
     * @return the team&#39;s permission level.
     * 
     */
    private String permission;
    /**
     * @return the team&#39;s privacy type.
     * 
     */
    private String privacy;
    /**
     * @return List of team repositories (list of repo names). Not returned if `summary_only = true`
     * 
     */
    private List<String> repositories;
    /**
     * @return List of team repositories (list of `repo_id` and `role_name`). Not returned if `summary_only = true`
     * 
     */
    private List<GetTeamRepositoriesDetailed> repositoriesDetaileds;
    private @Nullable Integer resultsPerPage;
    private String slug;
    private @Nullable Boolean summaryOnly;

    private GetTeamResult() {}
    /**
     * @return the team&#39;s description.
     * 
     */
    public String description() {
        return this.description;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return List of team members (list of GitHub usernames). Not returned if `summary_only = true`
     * 
     */
    public List<String> members() {
        return this.members;
    }
    public Optional<String> membershipType() {
        return Optional.ofNullable(this.membershipType);
    }
    /**
     * @return the team&#39;s full name.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return the Node ID of the team.
     * 
     */
    public String nodeId() {
        return this.nodeId;
    }
    /**
     * @return the team&#39;s permission level.
     * 
     */
    public String permission() {
        return this.permission;
    }
    /**
     * @return the team&#39;s privacy type.
     * 
     */
    public String privacy() {
        return this.privacy;
    }
    /**
     * @return List of team repositories (list of repo names). Not returned if `summary_only = true`
     * 
     */
    public List<String> repositories() {
        return this.repositories;
    }
    /**
     * @return List of team repositories (list of `repo_id` and `role_name`). Not returned if `summary_only = true`
     * 
     */
    public List<GetTeamRepositoriesDetailed> repositoriesDetaileds() {
        return this.repositoriesDetaileds;
    }
    public Optional<Integer> resultsPerPage() {
        return Optional.ofNullable(this.resultsPerPage);
    }
    public String slug() {
        return this.slug;
    }
    public Optional<Boolean> summaryOnly() {
        return Optional.ofNullable(this.summaryOnly);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetTeamResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String description;
        private String id;
        private List<String> members;
        private @Nullable String membershipType;
        private String name;
        private String nodeId;
        private String permission;
        private String privacy;
        private List<String> repositories;
        private List<GetTeamRepositoriesDetailed> repositoriesDetaileds;
        private @Nullable Integer resultsPerPage;
        private String slug;
        private @Nullable Boolean summaryOnly;
        public Builder() {}
        public Builder(GetTeamResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.description = defaults.description;
    	      this.id = defaults.id;
    	      this.members = defaults.members;
    	      this.membershipType = defaults.membershipType;
    	      this.name = defaults.name;
    	      this.nodeId = defaults.nodeId;
    	      this.permission = defaults.permission;
    	      this.privacy = defaults.privacy;
    	      this.repositories = defaults.repositories;
    	      this.repositoriesDetaileds = defaults.repositoriesDetaileds;
    	      this.resultsPerPage = defaults.resultsPerPage;
    	      this.slug = defaults.slug;
    	      this.summaryOnly = defaults.summaryOnly;
        }

        @CustomType.Setter
        public Builder description(String description) {
            if (description == null) {
              throw new MissingRequiredPropertyException("GetTeamResult", "description");
            }
            this.description = description;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetTeamResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder members(List<String> members) {
            if (members == null) {
              throw new MissingRequiredPropertyException("GetTeamResult", "members");
            }
            this.members = members;
            return this;
        }
        public Builder members(String... members) {
            return members(List.of(members));
        }
        @CustomType.Setter
        public Builder membershipType(@Nullable String membershipType) {

            this.membershipType = membershipType;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetTeamResult", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder nodeId(String nodeId) {
            if (nodeId == null) {
              throw new MissingRequiredPropertyException("GetTeamResult", "nodeId");
            }
            this.nodeId = nodeId;
            return this;
        }
        @CustomType.Setter
        public Builder permission(String permission) {
            if (permission == null) {
              throw new MissingRequiredPropertyException("GetTeamResult", "permission");
            }
            this.permission = permission;
            return this;
        }
        @CustomType.Setter
        public Builder privacy(String privacy) {
            if (privacy == null) {
              throw new MissingRequiredPropertyException("GetTeamResult", "privacy");
            }
            this.privacy = privacy;
            return this;
        }
        @CustomType.Setter
        public Builder repositories(List<String> repositories) {
            if (repositories == null) {
              throw new MissingRequiredPropertyException("GetTeamResult", "repositories");
            }
            this.repositories = repositories;
            return this;
        }
        public Builder repositories(String... repositories) {
            return repositories(List.of(repositories));
        }
        @CustomType.Setter
        public Builder repositoriesDetaileds(List<GetTeamRepositoriesDetailed> repositoriesDetaileds) {
            if (repositoriesDetaileds == null) {
              throw new MissingRequiredPropertyException("GetTeamResult", "repositoriesDetaileds");
            }
            this.repositoriesDetaileds = repositoriesDetaileds;
            return this;
        }
        public Builder repositoriesDetaileds(GetTeamRepositoriesDetailed... repositoriesDetaileds) {
            return repositoriesDetaileds(List.of(repositoriesDetaileds));
        }
        @CustomType.Setter
        public Builder resultsPerPage(@Nullable Integer resultsPerPage) {

            this.resultsPerPage = resultsPerPage;
            return this;
        }
        @CustomType.Setter
        public Builder slug(String slug) {
            if (slug == null) {
              throw new MissingRequiredPropertyException("GetTeamResult", "slug");
            }
            this.slug = slug;
            return this;
        }
        @CustomType.Setter
        public Builder summaryOnly(@Nullable Boolean summaryOnly) {

            this.summaryOnly = summaryOnly;
            return this;
        }
        public GetTeamResult build() {
            final var _resultValue = new GetTeamResult();
            _resultValue.description = description;
            _resultValue.id = id;
            _resultValue.members = members;
            _resultValue.membershipType = membershipType;
            _resultValue.name = name;
            _resultValue.nodeId = nodeId;
            _resultValue.permission = permission;
            _resultValue.privacy = privacy;
            _resultValue.repositories = repositories;
            _resultValue.repositoriesDetaileds = repositoriesDetaileds;
            _resultValue.resultsPerPage = resultsPerPage;
            _resultValue.slug = slug;
            _resultValue.summaryOnly = summaryOnly;
            return _resultValue;
        }
    }
}
