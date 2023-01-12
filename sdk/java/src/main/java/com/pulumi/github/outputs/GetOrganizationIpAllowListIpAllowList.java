// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetOrganizationIpAllowListIpAllowList {
    /**
     * @return A single IP address or range of IP addresses in CIDR notation.
     * 
     */
    private String allowListValue;
    /**
     * @return Identifies the date and time when the object was created.
     * 
     */
    private String createdAt;
    /**
     * @return The ID of the IP allow list entry.
     * 
     */
    private String id;
    /**
     * @return Whether the entry is currently active.
     * 
     */
    private Boolean isActive;
    /**
     * @return The name of the IP allow list entry.
     * 
     */
    private String name;
    /**
     * @return Identifies the date and time when the object was last updated.
     * 
     */
    private String updatedAt;

    private GetOrganizationIpAllowListIpAllowList() {}
    /**
     * @return A single IP address or range of IP addresses in CIDR notation.
     * 
     */
    public String allowListValue() {
        return this.allowListValue;
    }
    /**
     * @return Identifies the date and time when the object was created.
     * 
     */
    public String createdAt() {
        return this.createdAt;
    }
    /**
     * @return The ID of the IP allow list entry.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Whether the entry is currently active.
     * 
     */
    public Boolean isActive() {
        return this.isActive;
    }
    /**
     * @return The name of the IP allow list entry.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return Identifies the date and time when the object was last updated.
     * 
     */
    public String updatedAt() {
        return this.updatedAt;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetOrganizationIpAllowListIpAllowList defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String allowListValue;
        private String createdAt;
        private String id;
        private Boolean isActive;
        private String name;
        private String updatedAt;
        public Builder() {}
        public Builder(GetOrganizationIpAllowListIpAllowList defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.allowListValue = defaults.allowListValue;
    	      this.createdAt = defaults.createdAt;
    	      this.id = defaults.id;
    	      this.isActive = defaults.isActive;
    	      this.name = defaults.name;
    	      this.updatedAt = defaults.updatedAt;
        }

        @CustomType.Setter
        public Builder allowListValue(String allowListValue) {
            this.allowListValue = Objects.requireNonNull(allowListValue);
            return this;
        }
        @CustomType.Setter
        public Builder createdAt(String createdAt) {
            this.createdAt = Objects.requireNonNull(createdAt);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder isActive(Boolean isActive) {
            this.isActive = Objects.requireNonNull(isActive);
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder updatedAt(String updatedAt) {
            this.updatedAt = Objects.requireNonNull(updatedAt);
            return this;
        }
        public GetOrganizationIpAllowListIpAllowList build() {
            final var o = new GetOrganizationIpAllowListIpAllowList();
            o.allowListValue = allowListValue;
            o.createdAt = createdAt;
            o.id = id;
            o.isActive = isActive;
            o.name = name;
            o.updatedAt = updatedAt;
            return o;
        }
    }
}
