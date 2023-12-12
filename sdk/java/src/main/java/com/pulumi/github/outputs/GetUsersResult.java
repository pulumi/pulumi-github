// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetUsersResult {
    /**
     * @return list of the user&#39;s publicly visible profile email (will be empty string in case if user decided not to show it).
     * 
     */
    private List<String> emails;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return list of logins of users that could be found.
     * 
     */
    private List<String> logins;
    /**
     * @return list of Node IDs of users that could be found.
     * 
     */
    private List<String> nodeIds;
    /**
     * @return list of logins without matching user.
     * 
     */
    private List<String> unknownLogins;
    private List<String> usernames;

    private GetUsersResult() {}
    /**
     * @return list of the user&#39;s publicly visible profile email (will be empty string in case if user decided not to show it).
     * 
     */
    public List<String> emails() {
        return this.emails;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return list of logins of users that could be found.
     * 
     */
    public List<String> logins() {
        return this.logins;
    }
    /**
     * @return list of Node IDs of users that could be found.
     * 
     */
    public List<String> nodeIds() {
        return this.nodeIds;
    }
    /**
     * @return list of logins without matching user.
     * 
     */
    public List<String> unknownLogins() {
        return this.unknownLogins;
    }
    public List<String> usernames() {
        return this.usernames;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetUsersResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<String> emails;
        private String id;
        private List<String> logins;
        private List<String> nodeIds;
        private List<String> unknownLogins;
        private List<String> usernames;
        public Builder() {}
        public Builder(GetUsersResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.emails = defaults.emails;
    	      this.id = defaults.id;
    	      this.logins = defaults.logins;
    	      this.nodeIds = defaults.nodeIds;
    	      this.unknownLogins = defaults.unknownLogins;
    	      this.usernames = defaults.usernames;
        }

        @CustomType.Setter
        public Builder emails(List<String> emails) {
            this.emails = Objects.requireNonNull(emails);
            return this;
        }
        public Builder emails(String... emails) {
            return emails(List.of(emails));
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder logins(List<String> logins) {
            this.logins = Objects.requireNonNull(logins);
            return this;
        }
        public Builder logins(String... logins) {
            return logins(List.of(logins));
        }
        @CustomType.Setter
        public Builder nodeIds(List<String> nodeIds) {
            this.nodeIds = Objects.requireNonNull(nodeIds);
            return this;
        }
        public Builder nodeIds(String... nodeIds) {
            return nodeIds(List.of(nodeIds));
        }
        @CustomType.Setter
        public Builder unknownLogins(List<String> unknownLogins) {
            this.unknownLogins = Objects.requireNonNull(unknownLogins);
            return this;
        }
        public Builder unknownLogins(String... unknownLogins) {
            return unknownLogins(List.of(unknownLogins));
        }
        @CustomType.Setter
        public Builder usernames(List<String> usernames) {
            this.usernames = Objects.requireNonNull(usernames);
            return this;
        }
        public Builder usernames(String... usernames) {
            return usernames(List.of(usernames));
        }
        public GetUsersResult build() {
            final var _resultValue = new GetUsersResult();
            _resultValue.emails = emails;
            _resultValue.id = id;
            _resultValue.logins = logins;
            _resultValue.nodeIds = nodeIds;
            _resultValue.unknownLogins = unknownLogins;
            _resultValue.usernames = usernames;
            return _resultValue;
        }
    }
}
