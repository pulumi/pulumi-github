// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetDependabotOrganizationSecretsSecret {
    /**
     * @return Timestamp of the secret creation
     * 
     */
    private String createdAt;
    /**
     * @return Secret name
     * 
     */
    private String name;
    /**
     * @return Timestamp of the secret last update
     * 
     */
    private String updatedAt;
    /**
     * @return Secret visibility
     * 
     */
    private String visibility;

    private GetDependabotOrganizationSecretsSecret() {}
    /**
     * @return Timestamp of the secret creation
     * 
     */
    public String createdAt() {
        return this.createdAt;
    }
    /**
     * @return Secret name
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return Timestamp of the secret last update
     * 
     */
    public String updatedAt() {
        return this.updatedAt;
    }
    /**
     * @return Secret visibility
     * 
     */
    public String visibility() {
        return this.visibility;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetDependabotOrganizationSecretsSecret defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String createdAt;
        private String name;
        private String updatedAt;
        private String visibility;
        public Builder() {}
        public Builder(GetDependabotOrganizationSecretsSecret defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.createdAt = defaults.createdAt;
    	      this.name = defaults.name;
    	      this.updatedAt = defaults.updatedAt;
    	      this.visibility = defaults.visibility;
        }

        @CustomType.Setter
        public Builder createdAt(String createdAt) {
            if (createdAt == null) {
              throw new MissingRequiredPropertyException("GetDependabotOrganizationSecretsSecret", "createdAt");
            }
            this.createdAt = createdAt;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetDependabotOrganizationSecretsSecret", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder updatedAt(String updatedAt) {
            if (updatedAt == null) {
              throw new MissingRequiredPropertyException("GetDependabotOrganizationSecretsSecret", "updatedAt");
            }
            this.updatedAt = updatedAt;
            return this;
        }
        @CustomType.Setter
        public Builder visibility(String visibility) {
            if (visibility == null) {
              throw new MissingRequiredPropertyException("GetDependabotOrganizationSecretsSecret", "visibility");
            }
            this.visibility = visibility;
            return this;
        }
        public GetDependabotOrganizationSecretsSecret build() {
            final var _resultValue = new GetDependabotOrganizationSecretsSecret();
            _resultValue.createdAt = createdAt;
            _resultValue.name = name;
            _resultValue.updatedAt = updatedAt;
            _resultValue.visibility = visibility;
            return _resultValue;
        }
    }
}