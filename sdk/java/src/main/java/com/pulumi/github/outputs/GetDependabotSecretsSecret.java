// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.github.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetDependabotSecretsSecret {
    /**
     * @return Timestamp of the secret creation
     * 
     */
    private String createdAt;
    /**
     * @return The name of the repository.
     * 
     */
    private String name;
    /**
     * @return Timestamp of the secret last update
     * 
     */
    private String updatedAt;

    private GetDependabotSecretsSecret() {}
    /**
     * @return Timestamp of the secret creation
     * 
     */
    public String createdAt() {
        return this.createdAt;
    }
    /**
     * @return The name of the repository.
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

    public static Builder builder(GetDependabotSecretsSecret defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String createdAt;
        private String name;
        private String updatedAt;
        public Builder() {}
        public Builder(GetDependabotSecretsSecret defaults) {
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
        public GetDependabotSecretsSecret build() {
            final var _resultValue = new GetDependabotSecretsSecret();
            _resultValue.createdAt = createdAt;
            _resultValue.name = name;
            _resultValue.updatedAt = updatedAt;
            return _resultValue;
        }
    }
}
