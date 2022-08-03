// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetMembershipResult {
    /**
     * @return An etag representing the membership object.
     * 
     */
    private final String etag;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private final String id;
    private final @Nullable String organization;
    /**
     * @return `admin` or `member` -- the role the user has within the organization.
     * 
     */
    private final String role;
    /**
     * @return The username.
     * 
     */
    private final String username;

    @CustomType.Constructor
    private GetMembershipResult(
        @CustomType.Parameter("etag") String etag,
        @CustomType.Parameter("id") String id,
        @CustomType.Parameter("organization") @Nullable String organization,
        @CustomType.Parameter("role") String role,
        @CustomType.Parameter("username") String username) {
        this.etag = etag;
        this.id = id;
        this.organization = organization;
        this.role = role;
        this.username = username;
    }

    /**
     * @return An etag representing the membership object.
     * 
     */
    public String etag() {
        return this.etag;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public Optional<String> organization() {
        return Optional.ofNullable(this.organization);
    }
    /**
     * @return `admin` or `member` -- the role the user has within the organization.
     * 
     */
    public String role() {
        return this.role;
    }
    /**
     * @return The username.
     * 
     */
    public String username() {
        return this.username;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetMembershipResult defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private String etag;
        private String id;
        private @Nullable String organization;
        private String role;
        private String username;

        public Builder() {
    	      // Empty
        }

        public Builder(GetMembershipResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.etag = defaults.etag;
    	      this.id = defaults.id;
    	      this.organization = defaults.organization;
    	      this.role = defaults.role;
    	      this.username = defaults.username;
        }

        public Builder etag(String etag) {
            this.etag = Objects.requireNonNull(etag);
            return this;
        }
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        public Builder organization(@Nullable String organization) {
            this.organization = organization;
            return this;
        }
        public Builder role(String role) {
            this.role = Objects.requireNonNull(role);
            return this;
        }
        public Builder username(String username) {
            this.username = Objects.requireNonNull(username);
            return this;
        }        public GetMembershipResult build() {
            return new GetMembershipResult(etag, id, organization, role, username);
        }
    }
}
