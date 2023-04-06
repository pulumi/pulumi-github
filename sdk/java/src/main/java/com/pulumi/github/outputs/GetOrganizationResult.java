// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;

@CustomType
public final class GetOrganizationResult {
    /**
     * @return The description the organization account
     * 
     */
    private String description;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return The members login
     * 
     */
    private String login;
    /**
     * @return **Deprecated**: use `users` instead by replacing `github_organization.example.members` to `github_organization.example.users[*].login` which will give you the same value, expect this field to be removed in next major version
     * 
     * @deprecated
     * Use `users` instead by replacing `github_organization.example.members` to `github_organization.example.users[*].login`. Expect this field to be removed in next major version.
     * 
     */
    @Deprecated /* Use `users` instead by replacing `github_organization.example.members` to `github_organization.example.users[*].login`. Expect this field to be removed in next major version. */
    private List<String> members;
    /**
     * @return The organization&#39;s public profile name
     * 
     */
    private String name;
    /**
     * @return GraphQL global node id for use with v4 API
     * 
     */
    private String nodeId;
    /**
     * @return The organization&#39;s name as used in URLs and the API
     * 
     */
    private String orgname;
    /**
     * @return The plan name for the organization account
     * 
     */
    private String plan;
    /**
     * @return (`list`) A list with the repositories on the organization
     * 
     */
    private List<String> repositories;
    /**
     * @return (`list`) A list with the members of the organization with following fields:
     * 
     */
    private List<Map<String,String>> users;

    private GetOrganizationResult() {}
    /**
     * @return The description the organization account
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
     * @return The members login
     * 
     */
    public String login() {
        return this.login;
    }
    /**
     * @return **Deprecated**: use `users` instead by replacing `github_organization.example.members` to `github_organization.example.users[*].login` which will give you the same value, expect this field to be removed in next major version
     * 
     * @deprecated
     * Use `users` instead by replacing `github_organization.example.members` to `github_organization.example.users[*].login`. Expect this field to be removed in next major version.
     * 
     */
    @Deprecated /* Use `users` instead by replacing `github_organization.example.members` to `github_organization.example.users[*].login`. Expect this field to be removed in next major version. */
    public List<String> members() {
        return this.members;
    }
    /**
     * @return The organization&#39;s public profile name
     * 
     */
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
     * @return The organization&#39;s name as used in URLs and the API
     * 
     */
    public String orgname() {
        return this.orgname;
    }
    /**
     * @return The plan name for the organization account
     * 
     */
    public String plan() {
        return this.plan;
    }
    /**
     * @return (`list`) A list with the repositories on the organization
     * 
     */
    public List<String> repositories() {
        return this.repositories;
    }
    /**
     * @return (`list`) A list with the members of the organization with following fields:
     * 
     */
    public List<Map<String,String>> users() {
        return this.users;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetOrganizationResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String description;
        private String id;
        private String login;
        private List<String> members;
        private String name;
        private String nodeId;
        private String orgname;
        private String plan;
        private List<String> repositories;
        private List<Map<String,String>> users;
        public Builder() {}
        public Builder(GetOrganizationResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.description = defaults.description;
    	      this.id = defaults.id;
    	      this.login = defaults.login;
    	      this.members = defaults.members;
    	      this.name = defaults.name;
    	      this.nodeId = defaults.nodeId;
    	      this.orgname = defaults.orgname;
    	      this.plan = defaults.plan;
    	      this.repositories = defaults.repositories;
    	      this.users = defaults.users;
        }

        @CustomType.Setter
        public Builder description(String description) {
            this.description = Objects.requireNonNull(description);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder login(String login) {
            this.login = Objects.requireNonNull(login);
            return this;
        }
        @CustomType.Setter
        public Builder members(List<String> members) {
            this.members = Objects.requireNonNull(members);
            return this;
        }
        public Builder members(String... members) {
            return members(List.of(members));
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
        public Builder orgname(String orgname) {
            this.orgname = Objects.requireNonNull(orgname);
            return this;
        }
        @CustomType.Setter
        public Builder plan(String plan) {
            this.plan = Objects.requireNonNull(plan);
            return this;
        }
        @CustomType.Setter
        public Builder repositories(List<String> repositories) {
            this.repositories = Objects.requireNonNull(repositories);
            return this;
        }
        public Builder repositories(String... repositories) {
            return repositories(List.of(repositories));
        }
        @CustomType.Setter
        public Builder users(List<Map<String,String>> users) {
            this.users = Objects.requireNonNull(users);
            return this;
        }
        public GetOrganizationResult build() {
            final var o = new GetOrganizationResult();
            o.description = description;
            o.id = id;
            o.login = login;
            o.members = members;
            o.name = name;
            o.nodeId = nodeId;
            o.orgname = orgname;
            o.plan = plan;
            o.repositories = repositories;
            o.users = users;
            return o;
        }
    }
}
