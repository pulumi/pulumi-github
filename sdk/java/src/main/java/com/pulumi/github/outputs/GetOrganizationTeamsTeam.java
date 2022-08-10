// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetOrganizationTeamsTeam {
    /**
     * @return the team&#39;s description.
     * 
     */
    private final String description;
    /**
     * @return the ID of the team.
     * 
     */
    private final Integer id;
    /**
     * @return List of team members.
     * 
     */
    private final List<String> members;
    /**
     * @return the team&#39;s full name.
     * 
     */
    private final String name;
    /**
     * @return the Node ID of the team.
     * 
     */
    private final String nodeId;
    /**
     * @return the team&#39;s privacy type.
     * 
     */
    private final String privacy;
    /**
     * @return List of team repositories.
     * 
     */
    private final List<String> repositories;
    /**
     * @return the slug of the team.
     * 
     */
    private final String slug;

    @CustomType.Constructor
    private GetOrganizationTeamsTeam(
        @CustomType.Parameter("description") String description,
        @CustomType.Parameter("id") Integer id,
        @CustomType.Parameter("members") List<String> members,
        @CustomType.Parameter("name") String name,
        @CustomType.Parameter("nodeId") String nodeId,
        @CustomType.Parameter("privacy") String privacy,
        @CustomType.Parameter("repositories") List<String> repositories,
        @CustomType.Parameter("slug") String slug) {
        this.description = description;
        this.id = id;
        this.members = members;
        this.name = name;
        this.nodeId = nodeId;
        this.privacy = privacy;
        this.repositories = repositories;
        this.slug = slug;
    }

    /**
     * @return the team&#39;s description.
     * 
     */
    public String description() {
        return this.description;
    }
    /**
     * @return the ID of the team.
     * 
     */
    public Integer id() {
        return this.id;
    }
    /**
     * @return List of team members.
     * 
     */
    public List<String> members() {
        return this.members;
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
     * @return the team&#39;s privacy type.
     * 
     */
    public String privacy() {
        return this.privacy;
    }
    /**
     * @return List of team repositories.
     * 
     */
    public List<String> repositories() {
        return this.repositories;
    }
    /**
     * @return the slug of the team.
     * 
     */
    public String slug() {
        return this.slug;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetOrganizationTeamsTeam defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private String description;
        private Integer id;
        private List<String> members;
        private String name;
        private String nodeId;
        private String privacy;
        private List<String> repositories;
        private String slug;

        public Builder() {
    	      // Empty
        }

        public Builder(GetOrganizationTeamsTeam defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.description = defaults.description;
    	      this.id = defaults.id;
    	      this.members = defaults.members;
    	      this.name = defaults.name;
    	      this.nodeId = defaults.nodeId;
    	      this.privacy = defaults.privacy;
    	      this.repositories = defaults.repositories;
    	      this.slug = defaults.slug;
        }

        public Builder description(String description) {
            this.description = Objects.requireNonNull(description);
            return this;
        }
        public Builder id(Integer id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        public Builder members(List<String> members) {
            this.members = Objects.requireNonNull(members);
            return this;
        }
        public Builder members(String... members) {
            return members(List.of(members));
        }
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        public Builder nodeId(String nodeId) {
            this.nodeId = Objects.requireNonNull(nodeId);
            return this;
        }
        public Builder privacy(String privacy) {
            this.privacy = Objects.requireNonNull(privacy);
            return this;
        }
        public Builder repositories(List<String> repositories) {
            this.repositories = Objects.requireNonNull(repositories);
            return this;
        }
        public Builder repositories(String... repositories) {
            return repositories(List.of(repositories));
        }
        public Builder slug(String slug) {
            this.slug = Objects.requireNonNull(slug);
            return this;
        }        public GetOrganizationTeamsTeam build() {
            return new GetOrganizationTeamsTeam(description, id, members, name, nodeId, privacy, repositories, slug);
        }
    }
}