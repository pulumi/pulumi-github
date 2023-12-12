// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetActionsEnvironmentSecretsSecret {
    /**
     * @return Timestamp of the secret creation
     * 
     */
    private String createdAt;
    /**
     * @return Name of the secret
     * 
     */
    private String name;
    /**
     * @return Timestamp of the secret last update
     * 
     */
    private String updatedAt;

    private GetActionsEnvironmentSecretsSecret() {}
    /**
     * @return Timestamp of the secret creation
     * 
     */
    public String createdAt() {
        return this.createdAt;
    }
    /**
     * @return Name of the secret
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

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetActionsEnvironmentSecretsSecret defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String createdAt;
        private String name;
        private String updatedAt;
        public Builder() {}
        public Builder(GetActionsEnvironmentSecretsSecret defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.createdAt = defaults.createdAt;
    	      this.name = defaults.name;
    	      this.updatedAt = defaults.updatedAt;
        }

        @CustomType.Setter
        public Builder createdAt(String createdAt) {
            this.createdAt = Objects.requireNonNull(createdAt);
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
        public GetActionsEnvironmentSecretsSecret build() {
            final var _resultValue = new GetActionsEnvironmentSecretsSecret();
            _resultValue.createdAt = createdAt;
            _resultValue.name = name;
            _resultValue.updatedAt = updatedAt;
            return _resultValue;
        }
    }
}
