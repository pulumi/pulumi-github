// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
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
    private String etag;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private @Nullable String organization;
    /**
     * @return `admin` or `member` -- the role the user has within the organization.
     * 
     */
    private String role;
    /**
     * @return `active` or `pending` -- the state of membership within the organization.  `active` if the member has accepted the invite, or `pending` if the invite is still pending.
     * 
     */
    private String state;
    /**
     * @return The username.
     * 
     */
    private String username;

    private GetMembershipResult() {}
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
     * @return `active` or `pending` -- the state of membership within the organization.  `active` if the member has accepted the invite, or `pending` if the invite is still pending.
     * 
     */
    public String state() {
        return this.state;
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
    @CustomType.Builder
    public static final class Builder {
        private String etag;
        private String id;
        private @Nullable String organization;
        private String role;
        private String state;
        private String username;
        public Builder() {}
        public Builder(GetMembershipResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.etag = defaults.etag;
    	      this.id = defaults.id;
    	      this.organization = defaults.organization;
    	      this.role = defaults.role;
    	      this.state = defaults.state;
    	      this.username = defaults.username;
        }

        @CustomType.Setter
        public Builder etag(String etag) {
            if (etag == null) {
              throw new MissingRequiredPropertyException("GetMembershipResult", "etag");
            }
            this.etag = etag;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetMembershipResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder organization(@Nullable String organization) {

            this.organization = organization;
            return this;
        }
        @CustomType.Setter
        public Builder role(String role) {
            if (role == null) {
              throw new MissingRequiredPropertyException("GetMembershipResult", "role");
            }
            this.role = role;
            return this;
        }
        @CustomType.Setter
        public Builder state(String state) {
            if (state == null) {
              throw new MissingRequiredPropertyException("GetMembershipResult", "state");
            }
            this.state = state;
            return this;
        }
        @CustomType.Setter
        public Builder username(String username) {
            if (username == null) {
              throw new MissingRequiredPropertyException("GetMembershipResult", "username");
            }
            this.username = username;
            return this;
        }
        public GetMembershipResult build() {
            final var _resultValue = new GetMembershipResult();
            _resultValue.etag = etag;
            _resultValue.id = id;
            _resultValue.organization = organization;
            _resultValue.role = role;
            _resultValue.state = state;
            _resultValue.username = username;
            return _resultValue;
        }
    }
}
